# Parsing API

## Types

### KeyValue
Represents a simple key-value pair extracted from a string.

```go
type KeyValue struct {
    Key   string // The identifier or key
    Value string // The associated value or text
}
```

## Zero-Alloc Number Parsing from Bytes

For parsing numbers from byte slices without allocations (e.g., JSON parsers):

```go
c := fmt.GetConv()
c.LoadBytes(numBytes)      // load bytes into buffer (0 alloc)
v, err := c.Int64()        // or c.Float64()
c.PutConv()                // return Conv to pool
```

This bypasses `Convert(s ...any)` which boxes the string argument.
The pool ensures Conv reuse after warmup.