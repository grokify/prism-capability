package render

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	capstack "github.com/grokify/prism-capability"
)

// SVGLayout selects a marketecture diagram layout for native SVG rendering.
type SVGLayout string

const (
	// SVGLayoutStack renders layers as horizontal bands with capability cells,
	// framed by an applications band on top and a substrate band on the bottom.
	SVGLayoutStack SVGLayout = "stack"

	// SVGLayoutHub renders a hub-and-spoke diagram: a central node surrounded by
	// one node per layer arranged on a ring, over a substrate band. Best for a
	// small number of layers (roughly 4-8).
	SVGLayoutHub SVGLayout = "hub"
)

// SVGTheme holds the colors used when rendering a capability stack to SVG.
// Accents are cycled per layer; HubGradient colors the central hub.
type SVGTheme struct {
	Background  string   // page background (ignored when SVGOptions.Transparent)
	Surface     string   // band and node fill
	Text        string   // primary text
	TextMuted   string   // secondary/substrate text
	CellText    string   // capability/tool text inside nodes
	Accents     []string // per-layer accent colors, cycled
	HubGradient []string // gradient stops for the hub center and app band
}

// DefaultSVGTheme returns a dark theme with a cyan/purple/pink accent palette.
func DefaultSVGTheme() SVGTheme {
	return SVGTheme{
		Background:  "#0a0e1a",
		Surface:     "#1e293b",
		Text:        "#f1f5f9",
		TextMuted:   "#94a3b8",
		CellText:    "#cbd5e1",
		Accents:     []string{"#06b6d4", "#8b5cf6", "#ec4899", "#a855f7", "#0891b2", "#7c3aed", "#db2777"},
		HubGradient: []string{"#06b6d4", "#8b5cf6", "#ec4899"},
	}
}

// SVGOptions configures native SVG rendering of a capability stack.
type SVGOptions struct {
	// Layout selects the diagram layout. Defaults to SVGLayoutStack.
	Layout SVGLayout

	// Title is the diagram title (top-left). Defaults to the document title or name.
	Title string

	// Subtitle is rendered under the title.
	Subtitle string

	// TopBandLabel is the label of the top framing band in the stack layout.
	// Defaults to "Applications".
	TopBandLabel string

	// CenterLabel is the primary label of the hub center. Defaults to the document name.
	CenterLabel string

	// CenterSubLabel is the secondary label of the hub center.
	CenterSubLabel string

	// Substrate is the text of the bottom band (shared by both layouts).
	Substrate string

	// MaxToolsPerNode caps how many tool names appear in each hub layer node.
	// Defaults to 3. Ignored by the stack layout.
	MaxToolsPerNode int

	// Transparent omits the background rectangle so the SVG can be embedded
	// inline over an arbitrary page background.
	Transparent bool

	// Theme holds the color palette. A zero Theme falls back to DefaultSVGTheme.
	Theme SVGTheme
}

// DefaultSVGOptions returns options for a stack-layout diagram with the default theme.
func DefaultSVGOptions() SVGOptions {
	return SVGOptions{
		Layout:          SVGLayoutStack,
		MaxToolsPerNode: 3,
		Theme:           DefaultSVGTheme(),
	}
}

// RenderSVG renders a capability stack to native SVG.
func RenderSVG(w io.Writer, doc *capstack.CapabilityStack, opts SVGOptions) error {
	if doc == nil {
		return fmt.Errorf("render: nil capability stack")
	}
	opts = normalizeSVGOptions(doc, opts)

	var svg string
	switch opts.Layout {
	case SVGLayoutHub:
		svg = renderSVGHub(doc, opts)
	case SVGLayoutStack, "":
		svg = renderSVGStack(doc, opts)
	default:
		return fmt.Errorf("render: unknown SVG layout %q", opts.Layout)
	}

	if _, err := io.WriteString(w, svg); err != nil {
		return fmt.Errorf("render: write SVG: %w", err)
	}
	return nil
}

// RenderSVGString renders a capability stack to an SVG string.
func RenderSVGString(doc *capstack.CapabilityStack, opts SVGOptions) (string, error) {
	var b strings.Builder
	if err := RenderSVG(&b, doc, opts); err != nil {
		return "", err
	}
	return b.String(), nil
}

func normalizeSVGOptions(doc *capstack.CapabilityStack, opts SVGOptions) SVGOptions {
	if len(opts.Theme.Accents) == 0 {
		opts.Theme = DefaultSVGTheme()
	}
	if opts.MaxToolsPerNode <= 0 {
		opts.MaxToolsPerNode = 3
	}
	if opts.Title == "" {
		if doc.Metadata.Title != "" {
			opts.Title = doc.Metadata.Title
		} else {
			opts.Title = doc.Metadata.Name
		}
	}
	if opts.TopBandLabel == "" {
		opts.TopBandLabel = "Applications"
	}
	if opts.CenterLabel == "" {
		opts.CenterLabel = doc.Metadata.Name
	}
	return opts
}

// sortedLayers returns the document's layers sorted by Order ascending.
func sortedLayers(doc *capstack.CapabilityStack) []capstack.Layer {
	layers := make([]capstack.Layer, len(doc.Layers))
	copy(layers, doc.Layers)
	sort.SliceStable(layers, func(i, j int) bool { return layers[i].Order < layers[j].Order })
	return layers
}

// layerToolNames collects up to max distinct tool names across a layer's capabilities.
func layerToolNames(doc *capstack.CapabilityStack, layerID string, max int) []string {
	var names []string
	seen := map[string]bool{}
	for _, c := range doc.CapabilitiesByLayer(layerID) {
		for _, t := range c.Tooling {
			if t.Name == "" || seen[t.Name] {
				continue
			}
			seen[t.Name] = true
			names = append(names, t.Name)
			if len(names) >= max {
				return names
			}
		}
	}
	return names
}

// --- shared SVG helpers ---

func esc(s string) string {
	var b strings.Builder
	if err := xml.EscapeText(&b, []byte(s)); err != nil {
		panic(err) // EscapeText into a strings.Builder cannot fail
	}
	return b.String()
}

// textWidth estimates rendered width of a string at a font size (sans-serif).
func textWidth(s string, fontSize float64) float64 {
	return float64(len([]rune(s))) * fontSize * 0.55
}

func (t SVGTheme) accent(i int) string {
	return t.Accents[i%len(t.Accents)]
}

func gradientDef(id string, stops []string, x2, y2 string) string {
	var b strings.Builder
	fmt.Fprintf(&b, `<linearGradient id=%q x1="0" y1="0" x2=%q y2=%q>`, id, x2, y2)
	n := len(stops)
	for i, c := range stops {
		off := 0.0
		if n > 1 {
			off = float64(i) / float64(n-1)
		}
		fmt.Fprintf(&b, `<stop offset="%.2f" stop-color=%q/>`, off, c)
	}
	b.WriteString(`</linearGradient>`)
	return b.String()
}

// --- stack layout ---

const (
	stackCanvasW = 1200.0
	stackMargin  = 28.0
	stackLabelW  = 210.0
	stackCellH   = 28.0
	stackCellGap = 8.0
	stackCellPad = 11.0
	stackBandPad = 12.0
	stackBandGap = 6.0
	stackFrameH  = 44.0
	stackTitleH  = 92.0
)

func wrapCells(names []string, avail, fontSize float64) [][]string {
	var rows [][]string
	var row []string
	x := 0.0
	for _, n := range names {
		w := textWidth(n, fontSize) + 2*stackCellPad
		if len(row) > 0 && x+w > avail {
			rows = append(rows, row)
			row, x = nil, 0
		}
		row = append(row, n)
		x += w + stackCellGap
	}
	if len(row) > 0 {
		rows = append(rows, row)
	}
	return rows
}

func renderSVGStack(doc *capstack.CapabilityStack, opts SVGOptions) string {
	th := opts.Theme
	layers := sortedLayers(doc)
	avail := stackCanvasW - 2*stackMargin - stackLabelW - 2*stackBandPad

	type laid struct {
		layer capstack.Layer
		rows  [][]string
		h     float64
	}
	var rows []laid
	total := stackTitleH + stackFrameH + stackBandGap
	for _, l := range layers {
		var names []string
		for _, c := range doc.CapabilitiesByLayer(l.ID) {
			names = append(names, c.Name)
		}
		wrapped := wrapCells(names, avail, 12.5)
		if len(wrapped) == 0 {
			wrapped = [][]string{{}}
		}
		h := float64(len(wrapped))*stackCellH + float64(len(wrapped)-1)*stackCellGap + 2*stackBandPad
		rows = append(rows, laid{l, wrapped, h})
		total += h + stackBandGap
	}
	total += stackFrameH + stackMargin

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %.0f %.0f" font-family="Inter, system-ui, -apple-system, sans-serif">`+"\n", stackCanvasW, total)
	b.WriteString(`<defs>` + gradientDef("brand", th.HubGradient, "1", "0") + `</defs>` + "\n")
	if !opts.Transparent {
		fmt.Fprintf(&b, `<rect width="%.0f" height="%.0f" fill=%q/>`+"\n", stackCanvasW, total, th.Background)
	}

	// Title
	fmt.Fprintf(&b, `<text x="%.0f" y="46" font-size="28" font-weight="700" fill=%q>%s</text>`+"\n", stackMargin, th.Text, esc(opts.Title))
	if opts.Subtitle != "" {
		fmt.Fprintf(&b, `<text x="%.0f" y="72" font-size="14" fill=%q>%s</text>`+"\n", stackMargin, th.TextMuted, esc(opts.Subtitle))
	}

	bandW := stackCanvasW - 2*stackMargin
	y := stackTitleH

	// Applications frame band
	fmt.Fprintf(&b, `<rect x="%.0f" y="%.0f" width="%.0f" height="%.0f" rx="8" fill="url(#brand)" opacity="0.18"/>`+"\n", stackMargin, y, bandW, stackFrameH)
	fmt.Fprintf(&b, `<text x="%.0f" y="%.0f" font-size="15" font-weight="600" fill=%q text-anchor="middle">%s</text>`+"\n", stackCanvasW/2, y+stackFrameH/2+5, th.Text, esc(opts.TopBandLabel))
	y += stackFrameH + stackBandGap

	// Layer bands
	for i, r := range rows {
		accent := th.accent(i)
		fmt.Fprintf(&b, `<rect x="%.0f" y="%.1f" width="%.0f" height="%.1f" rx="8" fill=%q/>`+"\n", stackMargin, y, bandW, r.h, th.Surface)
		fmt.Fprintf(&b, `<rect x="%.0f" y="%.1f" width="4" height="%.1f" rx="2" fill=%q/>`+"\n", stackMargin, y, r.h, accent)
		fmt.Fprintf(&b, `<text x="%.0f" y="%.1f" font-size="16" font-weight="700" fill=%q>%s</text>`+"\n", stackMargin+18, y+r.h/2+5.5, th.Text, esc(r.layer.Name))

		cy := y + stackBandPad
		for _, cells := range r.rows {
			cx := stackMargin + stackLabelW
			for _, name := range cells {
				w := textWidth(name, 12.5) + 2*stackCellPad
				fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="%.1f" height="%.0f" rx="6" fill=%q stroke=%q stroke-opacity="0.45"/>`+"\n", cx, cy, w, stackCellH, th.Background, accent)
				fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" font-size="12.5" fill=%q text-anchor="middle">%s</text>`+"\n", cx+w/2, cy+stackCellH/2+4.5, accent, esc(name))
				cx += w + stackCellGap
			}
			cy += stackCellH + stackCellGap
		}
		y += r.h + stackBandGap
	}

	// Substrate band
	fmt.Fprintf(&b, `<rect x="%.0f" y="%.1f" width="%.0f" height="%.0f" rx="8" fill=%q/>`+"\n", stackMargin, y, bandW, stackFrameH, th.Surface)
	if opts.Substrate != "" {
		fmt.Fprintf(&b, `<text x="%.0f" y="%.1f" font-size="13" fill=%q text-anchor="middle">%s</text>`+"\n", stackCanvasW/2, y+stackFrameH/2+4.5, th.TextMuted, esc(opts.Substrate))
	}

	b.WriteString(`</svg>` + "\n")
	return b.String()
}

// --- hub layout ---

const (
	hubCanvasW = 1000.0
	hubRingR   = 248.0
	hubNodeW   = 228.0
	hubNodeH   = 78.0
	hubHexR    = 96.0
	hubTitleH  = 92.0
	hubFrameH  = 48.0
)

func renderSVGHub(doc *capstack.CapabilityStack, opts SVGOptions) string {
	th := opts.Theme
	layers := sortedLayers(doc)
	n := len(layers)
	if n == 0 {
		return renderSVGStack(doc, opts) // nothing to arrange
	}

	cx := hubCanvasW / 2
	cy := hubTitleH + 8 + hubRingR + hubNodeH/2
	total := cy + hubRingR + hubNodeH/2 + 18 + hubFrameH + 20

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %.0f %.0f" font-family="Inter, system-ui, -apple-system, sans-serif">`+"\n", hubCanvasW, total)
	b.WriteString(`<defs>` + gradientDef("hub", th.HubGradient, "1", "1") + `</defs>` + "\n")
	if !opts.Transparent {
		fmt.Fprintf(&b, `<rect width="%.0f" height="%.0f" fill=%q/>`+"\n", hubCanvasW, total, th.Background)
	}

	// Title
	fmt.Fprintf(&b, `<text x="%.0f" y="46" font-size="28" font-weight="700" fill=%q>%s</text>`+"\n", 40.0, th.Text, esc(opts.Title))
	if opts.Subtitle != "" {
		fmt.Fprintf(&b, `<text x="%.0f" y="72" font-size="14" fill=%q>%s</text>`+"\n", 40.0, th.TextMuted, esc(opts.Subtitle))
	}

	// Node positions on the ring, starting at top (90 deg) going clockwise.
	type pos struct{ x, y float64 }
	nodes := make([]pos, n)
	for i := range layers {
		angle := (90.0 - float64(i)*360.0/float64(n)) * math.Pi / 180.0
		nodes[i] = pos{cx + hubRingR*math.Cos(angle), cy - hubRingR*math.Sin(angle)}
	}

	// Spokes (drawn first, covered by hub + nodes)
	for i := range layers {
		fmt.Fprintf(&b, `<line x1="%.0f" y1="%.0f" x2="%.1f" y2="%.1f" stroke=%q stroke-opacity="0.35" stroke-width="1.5"/>`+"\n",
			cx, cy, nodes[i].x, nodes[i].y, th.accent(i))
	}

	// Center hub (pointy-top hexagon)
	var pts []string
	for _, deg := range []float64{90, 150, 210, 270, 330, 30} {
		r := deg * math.Pi / 180.0
		pts = append(pts, fmt.Sprintf("%.1f,%.1f", cx+hubHexR*math.Cos(r), cy-hubHexR*math.Sin(r)))
	}
	fmt.Fprintf(&b, `<polygon points=%q fill="url(#hub)" fill-opacity="0.18" stroke="url(#hub)" stroke-width="2"/>`+"\n", strings.Join(pts, " "))
	fmt.Fprintf(&b, `<text x="%.0f" y="%.0f" text-anchor="middle" font-size="21" font-weight="700" fill=%q>%s</text>`+"\n", cx, cy-4, th.Text, esc(opts.CenterLabel))
	if opts.CenterSubLabel != "" {
		fmt.Fprintf(&b, `<text x="%.0f" y="%.0f" text-anchor="middle" font-size="12" fill=%q>%s</text>`+"\n", cx, cy+20, th.TextMuted, esc(opts.CenterSubLabel))
	}

	// Layer nodes
	for i, l := range layers {
		accent := th.accent(i)
		nx, ny := nodes[i].x, nodes[i].y
		fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="%.0f" height="%.0f" rx="12" fill=%q stroke=%q stroke-width="1.5"/>`+"\n",
			nx-hubNodeW/2, ny-hubNodeH/2, hubNodeW, hubNodeH, th.Surface, accent)
		fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" text-anchor="middle" font-size="15" font-weight="700" fill=%q>%s</text>`+"\n",
			nx, ny-8, accent, esc(l.Name))

		sub := strings.Join(layerToolNames(doc, l.ID, opts.MaxToolsPerNode), " · ")
		if sub == "" {
			if c := len(doc.CapabilitiesByLayer(l.ID)); c > 0 {
				sub = fmt.Sprintf("%d capabilities", c)
			}
		}
		if sub != "" {
			fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" text-anchor="middle" font-size="10.5" fill=%q>%s</text>`+"\n",
				nx, ny+15, th.CellText, esc(sub))
		}
	}

	// Substrate band
	subW := hubRingR*2 + hubNodeW - 20
	subX := cx - subW/2
	subY := cy + hubRingR + hubNodeH/2 + 18
	fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="%.1f" height="%.0f" rx="10" fill=%q stroke="#334155" stroke-width="1"/>`+"\n", subX, subY, subW, hubFrameH, th.Surface)
	if opts.Substrate != "" {
		fmt.Fprintf(&b, `<text x="%.0f" y="%.1f" text-anchor="middle" font-size="12.5" fill=%q>%s</text>`+"\n", cx, subY+hubFrameH/2+4.5, th.TextMuted, esc(opts.Substrate))
	}

	b.WriteString(`</svg>` + "\n")
	return b.String()
}
