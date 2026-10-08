# AGENTS.md — webtyp/fmt

Working notes for AI agents operating in this library. For end-user docs see [README.md](README.md).
Plans for code changes live in [docs/PLAN.md](docs/PLAN.md) and link back here for the standing
rules below (do not duplicate them in plans).

## Mission

`webtyp/fmt` is the **foundation** of the ecosystem: string manipulation, type conversion,
formatting, error handling, schema contracts, and the typed serialization codec — all
**reflection-free** and **TinyGo/WASM-optimized**. It is the stdlib replacement: nothing in the
ecosystem imports `fmt`, `strings`, `strconv`, or `errors` — they use this package.

## Ecosystem restrictions (do NOT violate)

- **`fmt` IS the stdlib replacement.** Inside this package, do not import `strings`, `strconv`,
  `errors`, or the stdlib `fmt`. Use the internal `Conv` buffer/primitives.
- **No `reflect` in WASM builds.** Reflection (~tens of KB of type tables) belongs only in
  `//go:build !wasm` files (e.g. the `*.back.go` reflection fallbacks). The agnostic/`wasm` side
  must be reflection-free.
- **No `map` on hot paths / WASM.** TinyGo pulls in the hashmap runtime; it inflates the binary.
- **Build-tag split convention:** `*.back.go` (`//go:build !wasm`) for backend-only code
  (`sync`, `reflect`); `*.front.go` (`//go:build wasm`) for the WASM counterpart. Keep both in
  sync behind a shared unexported API.
- **0-allocation is a first-class goal.** Reuse `Conv` buffers; never return `[]any`/`[]Field`/
  `map` on a hot path. Note: `GetConv()` **allocates in WASM** (no pool) — reuse a single `Conv`
  across a serialization pass rather than calling `GetConv()` per field.

## Key contracts owned by this package

- **`Conv` / `Builder`** — the chainable buffer and its pool (`GetConv`/`PutConv`), `KeyValue`,
  `Kind`.
- **The translation hook** (`SetTranslator`) — i18n is opt-in: `webtyp.com/lang` installs it.
  `fmt` never imports `lang`; `Err`, `Println` and `Sprintf("%L")` translate only when `lang` is
  imported.
- Exported for the split-out repos, never remove: `SetTranslator`, `GetConv`, `Conv`, `BuffOut`,
  `BuffDest`, `(*Conv).WrString`, `Sprintf`, `Split`, `IsWordSeparator` (lang); `ToLower`,
  `Contains` (msgtype); `Builder` (escape); `Index`, `HasPrefix`, `HasSuffix` (filepath).

## Rules

- **fmt only grows by stdlib equivalents** (`fmt`, `strings`, `strconv`, `errors`). Anything
  else is a new repo (paths → `filepath`, i18n → `lang`, message types → `msgtype`, escaping →
  `escape`).
- **Tests live in `tests/`** (package `fmt_test`, public API only). A root-level test only when it
  needs an unexported identifier, with a first line `// Root-level test (justified): …`. **Never
  export a symbol so a test can reach it.**

## Testing

```bash
go install webtyp.com/devflow/cmd/gotest@latest   # once
gotest            # vet + race + cover + wasm + badges (NOT `go test`)
gotest -run TestX
```

Binary-size / memory benchmarks (standard library vs webtyp/fmt) live in `benchmark/`
(`build-and-measure.sh`, `benchmark_results.md`). Update them when the hot path changes.

## Publishing

```bash
gopush 'message'   # tests + tag + push + dependency bumps (NOT git commit/push directly)
```

## Related

- [`webtyp/orm`](https://github.com/webtyp/orm) — `ormc` generates `Schema()`/`Pointers()`
  and the codec's `EncodeFields`/`DecodeFields`.
- [`webtyp/json`](https://github.com/webtyp/json), [`webtyp/jsvalue`](https://github.com/webtyp/jsvalue)
  — concrete codec encoders (JSON, JS).
