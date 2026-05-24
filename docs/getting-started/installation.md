# Installation

## CLI

Install the `capstack` CLI tool:

```bash
go install github.com/grokify/prism-capability/cmd/capstack@latest
```

Verify installation:

```bash
capstack --help
```

## Library

Add prism-capability as a dependency:

```bash
go get github.com/grokify/prism-capability
```

Import in your code:

```go
import (
    "github.com/grokify/prism-capability"
    "github.com/grokify/prism-capability/render"
)
```

## Requirements

- Go 1.24 or later
- For D2 diagram rendering: [D2](https://d2lang.com/) installed (optional)
