# capstack render

Generate visualizations from a capability stack.

## Synopsis

```bash
capstack render <file> [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--format`, `-f` | Output format: `svg`, `d2`, `html`, `lit`, or `json` | `d2` |
| `--output`, `-o` | Output file path | stdout |
| `--layout` | SVG layout: `stack` or `hub` (svg only) | `stack` |
| `--substrate` | Substrate text for SVG (svg only) | (empty) |
| `--standalone` | Include full HTML document (html/lit) | `true` |
| `--dark` | Use dark theme (HTML/lit) | `false` |
| `--legend` | Show status legend/filters | `true` |
| `--style` | D2 style: `default` or `grid` (d2 only) | `default` |
| `--view` | View mode: `by-layer` or `by-category` (lit/json) | `by-layer` |
| `--component-path` | Path to prism-ui.js component (lit) | `prism-ui.js` |

## Examples

### SVG (Native)

```bash
# Generate layered stack diagram (default)
capstack render my-stack.json --format svg -o stack.svg

# Generate hexagonal hub-and-spoke (for smaller stacks)
capstack render my-stack.json --format svg --layout hub -o hub.svg

# Customize substrate text
capstack render my-stack.json --format svg --substrate "Go · Kubernetes · 15+ providers" -o stack.svg

# Hub with custom substrate
capstack render my-stack.json --format svg --layout hub --substrate "Built on cloud-native infrastructure" -o platform.svg
```

### D2

```bash
# D2 diagram
capstack render my-stack.json --format d2 -o stack.d2

# D2 grid style (executive view)
capstack render my-stack.json --format d2 --style grid -o stack-grid.d2

# Convert D2 to SVG (requires d2 installed)
d2 stack.d2 stack.svg
```

### HTML

```bash
# HTML output to file
capstack render my-stack.json --format html -o stack.html

# HTML fragment (embeddable)
capstack render my-stack.json --format html --standalone=false

# Dark theme
capstack render my-stack.json --format html --dark -o stack-dark.html
```

### Lit Web Component

```bash
# Lit web component (interactive)
capstack render my-stack.json --format lit -o stack.html

# Lit with category view and dark theme
capstack render my-stack.json --format lit --view by-category --dark -o stack.html
```

### JSON

```bash
# JSON data for custom integration
capstack render my-stack.json --format json -o stack-data.json
```

## SVG Output

The native SVG renderer generates portable, themeable diagrams in two layouts:

### Stack Layout (default)

Layered marketecture diagram best for comprehensive architecture documentation:

- Horizontal bands per layer with layer names and accent stripes
- Capability cells with smart text wrapping
- "Applications" framing band at top
- "Substrate" band at bottom for infrastructure/providers
- Per-layer accent color cycling
- No external dependencies (pure Go + stdlib)

Use for README files, wikis, documentation sites, and presentations.

### Hub Layout

Hexagonal hub-and-spoke diagram best for smaller stacks (4-8 layers):

- Central hexagon representing the core system or platform
- One node per layer arranged on a ring around the hub
- Spokes connecting layers to the center
- Substrate band at bottom
- Tooling names extracted from layer capabilities (up to 3 per node)
- Fallback to capability count if tooling is unavailable

Use for executive summaries, hero diagrams, and slide decks.

### Features

- **Themeable**: Customizable colors, accents, typography via Go API
- **Responsive**: Viewbox-based scaling, readable at any size
- **Secure**: Pure vector SVG with no scripts or external resources
- **No CLI dependency**: Unlike D2, SVG generation requires no external tools

## HTML Output

The HTML renderer generates:

- Layered capability visualization
- Status color coding
- Interactive filters (status, maturity)
- Tooltips with capability details
- Optional legend/filter controls

## D2 Output

The D2 renderer generates diagrams showing:

- Capabilities organized by layer
- Category color coding
- Dependency relationships (default style)
- Grid layout (grid style)
- Optional badge overlays

## Lit Output

The Lit renderer generates interactive HTML with:

- Modern Lit web component (`<maturity-grid>`)
- View toggle between layer and category grouping
- Interactive filtering by status
- Sorting options
- Dark/light theme support
- Tooltips with capability details

## JSON Output

The JSON renderer outputs structured data for custom Lit component integration:

- `layers`: Layer definitions with IDs and names
- `categories`: Category definitions
- `capabilities`: Capability data with layer/category associations
- `maturity`: Optional maturity level data

Use JSON output when integrating with custom web applications or when you need programmatic access to the rendered data structure.
