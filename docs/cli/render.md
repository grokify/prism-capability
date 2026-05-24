# capstack render

Generate visualizations from a capability stack.

## Synopsis

```bash
capstack render <file> [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--format`, `-f` | Output format: `html` or `d2` | `html` |
| `--output`, `-o` | Output file path | stdout |
| `--standalone` | Include full HTML document | `true` |
| `--dark` | Use dark theme | `false` |
| `--legend` | Show status legend/filters | `true` |
| `--style` | D2 style: `default` or `grid` | `default` |

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
