# webtyp/fmt
<img src="docs/img/badges.svg">

fmt is a lightweight Go library that provides comprehensive string manipulation, type conversion, formatting, and multilingual error handling with a fluid API, specifically designed for small devices and web applications using TinyGo as the target compiler.


## Key Features

- 🚀 **Fluid and chainable API** - Easy to use and readable operations
- 📝 **Complete string toolkit** - Transformations, conversions, formatting, and error handling
- 🌍 **Multilingual error messages** - Opt-in: importing `webtyp.com/lang` installs the translator
- 🧵 **Concurrency safe** - Thread-safe operations for concurrent environments
- 📦 **Zero dependencies** - No `fmt`, `strings`, `strconv`, or `errors` imports
- 🎯 **TinyGo optimized** - Manual implementations for minimal binary size
- 🌐 **WebAssembly-first** - Designed for modern web deployment
- 🔄 **Universal type support** - Works with strings, numbers, booleans, and slices
- ⚡ **Performance focused** - Predictable allocations and custom optimizations

## [Why fmt?](docs/WHY.md)

## Installation

```bash
go get webtyp.com/fmt
```

## Usage

```go
import . "webtyp.com/fmt"

// Conversion and transformation
text := Convert("Hola Mundo").ToLower().String() // out: "hola mundo"
numText := Convert(42).String()                  // out: "42"
boolText := Convert(true).String()               // out: "true"

// Memory-efficient approach using string pointers
original := "el murcielago rapido"
Convert(&original).CamelUp().Apply()
// original is now: "ElMurcielagoRapido"

// Efficient builder and chaining
items := []string{"  APPLE  ", "  banana  ", "  piñata  "}
builder := Convert() // without params reused buffer = optimal performance
for i, item := range items {
    processed := Convert(item).TrimSpace().ToLower().Capitalize().String()
    builder.Write(processed)
    if i < len(items)-1 {
        builder.Write(" - ")
    }
}
out := builder.String() // out: "Apple - Banana - Piñata"

// Formatting, quoting and errors
msg := Sprintf("%s has %d items", "cart", 3) // out: "cart has 3 items"
q := Sprintf("%q", "say \"hi\"")             // out: "\"say \\\"hi\\\"\""
err := Err("invalid", "email")                // translated when webtyp.com/lang is imported

// Multi-term searching (AND) with internal normalization
found := Matches("Hello World", "hello", "world") // out: true
```

## Moved out of fmt

`fmt` only replaces `fmt`, `strings`, `strconv` and `errors`, plus its own core. Other concerns
live in their own repos (fmt v1.1.0):

| Was in fmt | Now |
|---|---|
| `PathJoin/PathBase/PathExt/PathShort/PathRelativeTo` | `webtyp.com/filepath`: `Join/Base/Ext/Short/RelativeTo` (+ `Tilde`) |
| `SetPathBase` | no successor: tests use `t.Chdir` |
| `webtyp.com/fmt/lang` | `webtyp.com/lang` |
| `MessageType`, `Msg`, `StringType` | `webtyp.com/msgtype`: `Type`, constants, `Detect` |
| `EscapeHTML`, `EscapeAttr` | `escape.HTML` (`webtyp.com/escape`) |
| `JSONEscape` | `escape.JSON` |
| `Html` | `Sprintf` |
| `Convert(x).Quote()` | `Sprintf("%q", x)` |

Deleted (no users): `TagPairs`, `TagValue`, `ExtractValue`, `Thousands`, `HasUpperPrefix`,
`CamelLow`, `SnakeUp`, `ReplaceN`, `MatchesAny`, `IsZero`, `TruncateName`, `Tilde`, `GetKind`,
`IDorPrimaryKey`. Now private: `Quote`, `StringErr`, `GetStringZeroCopy`.

## Documentation

### String Manipulation & Conversion

- [Fmt Package Equivalents](docs/API_FMT.md) - Replace fmt package functions
- [Strings Package Equivalents](docs/API_STRINGS.md) - Replace strings package functions
- [Strconv Package Equivalents](docs/API_STRCONV.md) - Replace strconv package functions
- [Errors Package Equivalents](docs/API_ERRORS.md) - Replace errors package functions

### Utilities & Helpers

- [Key-Value type and zero-alloc parsing](docs/API_PARSING.md)
- [Smart Truncation](docs/TRUNCATION.md) - Text truncation utilities


## Testing

Tests live in `tests/` (package `fmt_test`, public API only). A test stays at the root only
when it needs an unexported identifier, and starts with a `// Root-level test (justified):` line.
Run `gotest` (vet, race, WASM).

## Benchmarking
- [Benchmarking](benchmark/README.md)

## Examples

- [WebAssembly Html Code](example/web/client.go)

---
## [Contributing](https://github.com/webtyp/cdvelop/blob/main/CONTRIBUTING.md)
---
## [License](LICENSE)