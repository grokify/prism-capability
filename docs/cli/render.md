# capstack render

Generate visualizations from a capability stack.

## Synopsis

```bash
capstack render <file> [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--format`, `-f` | Output format: `d2`, `html`, `lit`, or `json` | `d2` |
| `--output`, `-o` | Output file path | stdout |
| `--standalone` | Include full HTML document | `true` |
| `--dark` | Use dark theme (HTML/lit) | `false` |
| `--legend` | Show status legend/filters | `true` |
| `--style` | D2 style: `default` or `grid` | `default` |
| `--view` | View mode: `by-layer` or `by-category` (lit/json) | `by-layer` |
| `--component-path` | Path to prism-ui.js component (lit) | `prism-ui.js` |

## Examples

```bash
# HTML output to file
capstack render my-stack.json --format html -o stack.html

# HTML fragment (embeddable)
capstack render my-stack.json --format html --standalone=false

# Dark theme
capstack render my-stack.json --format html --dark -o stack-dark.html

# D2 diagram
capstack render my-stack.json --format d2 -o stack.d2

# D2 grid style (executive view)
capstack render my-stack.json --format d2 --style grid -o stack-grid.d2

# Convert D2 to SVG (requires d2 installed)
d2 stack.d2 stack.svg

# Lit web component (interactive)
capstack render my-stack.json --format lit -o stack.html

# Lit with category view and dark theme
capstack render my-stack.json --format lit --view by-category --dark -o stack.html

# JSON data for custom integration
capstack render my-stack.json --format json -o stack-data.json
```

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
