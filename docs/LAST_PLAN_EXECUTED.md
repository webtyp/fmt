---
PLAN: "refactor!: fmt reduced to its minimal layer; tests moved to tests/"
TAG: v1.1.0
EXECUTOR: claude (local; Jules session 4340317710125413898 vanished, 404)
REVIEWER: none
STATUS: done — tag v1.1.0 pending until mjosefa-cms drops webtyp.com/fmt/lang (roadmap H1)
---

> This plan is dispatched via the CodeJob workflow. See skill: agents-workflow.

# Plan — fmt: the minimal layer

Phase **F (last)** of the master plan `SOURCE_SELECTION_MASTER_PLAN.md` (orchestration only —
everything this plan needs is inline). **BREAKING CHANGE, no backward compatibility, no
deprecation shims.** Tag `v1.1.0` is a deliberate choice of the maintainer: every consumer
migrates before this tag exists.

Standing rules: [AGENTS.md](../AGENTS.md). In short: this package compiles to TinyGo/WASM; no
`strings`/`strconv`/`errors`/stdlib `fmt` inside it; no `map`; `os`/`sync`/`reflect` only in
`*.back.go` / `*.stlib.go` (`//go:build !wasm`) files.

## Why

`fmt` is in every binary of the ecosystem. Every change to it forces an update of the whole
ecosystem. So it keeps only what replaces the standard library's `fmt`, `strings`, `strconv` and
`errors`, plus its own core (`Conv`/`Builder`, `KeyValue`, `Kind`, the translation hook). Anything
with a different responsibility, or that keeps growing, lives in its own repo. Anything nobody uses is
deleted.

TinyGo already drops unused functions from the binary. The gain here is not bytes: it is fewer
reasons for `fmt` to change, and a smaller surface to learn.

## State when this plan starts (already done by the maintainer, 2026-10-06)

The code was physically moved out. It no longer exists in this repo:

| Moved | Now in |
|---|---|
| `filepath.go`, `filepath.stlib.go`, `filepath.wasm.go`, their tests, `docs/API_FILEPATH.md` | `https://github.com/webtyp/filepath` |
| `lang/` (whole package), `docs/TRANSLATE.md` | `https://github.com/webtyp/lang` |
| `messagetype.go`, its test, `docs/MESSAGE_TYPES.md`, two `docs/issues/*MESSAGE*` files | `https://github.com/webtyp/msgtype` |
| `html.go` (`EscapeHTML`, `EscapeAttr`, `Html`), `JSONEscape` (cut from `quote.go`), their tests, `docs/API_HTML.md`, `docs/API_JSON_ESCAPE.md` | `https://github.com/webtyp/escape` |
| `IDorPrimaryKey.go`, its test, `docs/ID_PRIMARY_KEY.md` | deleted (zero users) |
| every `*_test.go` of the root | `tests/` (moved with `git mv`, contents unchanged) |

The library builds (backend and WASM). The tests do not: 32 files in `tests/` still say
`package fmt`.

## Design gate

1. **Prior art.** Go itself splits `fmt`, `strings`, `strconv`, `errors`, `path/filepath`, `html`,
   `encoding/json` and `log` into separate packages. Rust's `std::fmt` only formats. Node splits
   `util.format`, `path` and `querystring`. Every one of them keeps formatting separate from paths,
   escaping, i18n and log classification.
2. **Novice-name test.** No new names in `fmt`. The new homes are `filepath`, `lang`, `msgtype` and
   `escape`, each named after its concern.
3. **Complexity ledger.**
   ```
   Concepts the developer must learn   −20 exported symbols in fmt (listed in Stage 2)
   Files they must touch to do X       +0 / −0
   Lines at the call site              ±0 (consumers change an import, plus a name where the API moved)
   Ways to do the same thing           −3  (Html ≡ Sprintf; Escape* duplicated; lang.Println ≡ fmt.Println)
   ```
4. **Where it belongs.** Each moved piece is in the repo named after its concern (table above).
5. **What it deletes.** Stage 2.

## Stage 1 — tests in `tests/`

RULE for this repo from now on: tests live in `tests/`. A test stays at the root only when it needs
an unexported identifier, with a written justification. **Never export a symbol so a test can
reach it.**

1. Every file in `tests/` becomes `package fmt_test` and adds `import . "webtyp.com/fmt"`. The
   dot import keeps existing call sites (`Convert(...)`, `Sprintf(...)`) unchanged. The 5 files that
   already say `package fmt_test` keep their imports.
2. A test that uses an unexported identifier (`getConv`, `putConv`, `wrString`, `kindNames`,
   `isWasm`, buffer internals, …): first try to assert the same behaviour through the exported API.
   If that is not possible, move that test (only the cases that need internals, split into their own
   file if needed) back to the root as `package fmt`. Its first lines are
   `// Root-level test (justified): exercises <unexported identifiers> — <why the behaviour is not
   observable through the exported API>.` Never export a symbol just for a test. List every
   root-level file and its justification in the PR description.
3. Add `tests/` to whatever the README says about running tests.
4. `benchmark/bench-memory-alloc/*/main_test.go` are separate benchmark programs, not package
   tests: leave them where they are.

## Stage 2 — delete the unused API (zero users in the workspace, verified 2026-10-06)

Delete each symbol, its now-dead private helpers, its test cases and its doc mentions:

- `parse.go`: `TagPairs`, `TagValue`, `ExtractValue` (keep `KeyValue`; it is used everywhere).
  If `docs/STRUCT_TAGS.md` only documents these, delete it.
- `fmt_number.go`: `Thousands`.
- `capitalize.go`: `HasUpperPrefix`, `CamelLow`, `SnakeUp` (keep `Capitalize`, `CamelUp`,
  `SnakeLow`).
- `strings.go`: `ReplaceN`.
- `search.go`: `MatchesAny`.
- `convert.go`: `IsZero`.
- `truncate.go`: `TruncateName` (keep `Truncate`).
- `mapping.go`: `Tilde`, `isASCIIOnlyOut`, `tildeUnicodeOptimized` and any table only they use.
  Keep the tables that `toUpperRune`/`toLowerRune` use.
- `GetKind` (find it with `grep -n "GetKind" *.go`).

Make these private (used only inside this package): `Quote` → `quote`, `StringErr` → `stringErr`,
`GetStringZeroCopy` → `getStringZeroCopy`.

`Quote` stays **in fmt**: it is the counterpart of the stdlib's `strconv.Quote` (it builds a Go
string literal), which is this package's charter. It is the engine of `Sprintf("%q", …)`, and
that is the one public way to quote. Two paths (`Convert(s).Quote()` and `%q`) for one intent is
one too many, and nothing outside fmt calls `Quote()`. It does NOT belong in `webtyp.com/escape`:
escape makes text safe for an output context (HTML, JSON), while `%q` produces Go syntax. The test
cases in `tests/quote_test.go` and `tests/concurrency_test.go` that call `.Quote()` are rewritten
as `Sprintf("%q", x)` with the same expectations, not deleted.

### Stage 2b — bug: `%q` emits raw control bytes (red test first)

Verified 2026-10-06: `Sprintf("%q", "a\x01b")` returns `"a` + byte 0x01 + `b"`, while Go returns
`"a\x01b"`. The same happens with `\a` and `\x7f`. The result is not a valid Go literal, and an
invisible byte reaches logs.
1. Add to `tests/fmt_q_escape_test.go` a table that compares against the expected literal text
   (hard-code the expected strings; the test must not import the stdlib `fmt`):
   `"a\x01b"` → `"a\x01b"`, `"bell\a"` → `"bell\a"`, `"\b\f\v"` → `"\b\f\v"`,
   `"x\x7f"` → `"x\x7f"`, `"tab\there"` → `"tab\there"`, `"ñandú"` → `"ñandú"` (printable
   UTF-8 stays as is, as in Go). Run it; it fails.
2. Fix in `quote.go`: besides `\" \\ \n \r \t`, write `\a \b \f \v` for those bytes, and
   `\xHH` (lowercase hex, two digits) for every other byte `< 0x20` and for `0x7f`. Bytes `>= 0x80`
   pass through unchanged. Do NOT add Unicode printability tables: they would cost binary size.
   Write in the function comment that non-printable non-ASCII runes are not escaped, unlike Go's `%q`.

The maintainer verified zero users outside this repo for every symbol above. If a test in `tests/`
is the only user, delete those test cases together with the symbol.

## Stage 3 — the translation hook stays

`translate_hook.go` (`SetTranslator`, `tr`) STAYS. `webtyp.com/lang` installs the translator from
outside. `fmt` never imports `lang`, and `Err`, `Println` and `Sprintf("%L")` keep translating when
`lang` is imported. Update the comment in `translate_hook.go`: "It is typically called by the
webtyp.com/lang package during its initialization."

The new repos use these exported symbols. They MUST stay exported, and Stage 2 must not touch them:
- `webtyp.com/lang`: `SetTranslator`, `GetConv`, `Conv`, `BuffOut`, `BuffDest`,
  `(*Conv).WrString`, `Sprintf`, `Split`, `IsWordSeparator`.
- `webtyp.com/msgtype`: `ToLower`, `Contains`.
- `webtyp.com/escape`: `Builder` (`WriteString`, `WriteByte`).
- `webtyp.com/filepath`: `Index`, `HasPrefix`, `HasSuffix`.

## Stage 4 — docs

- `README.md`: remove links to the moved docs. Add a section "Moved out of fmt" with this mapping:
  `PathJoin/PathBase/PathExt/PathShort/PathRelativeTo/SetPathBase` → `webtyp.com/filepath`
  (`Join/Base/Ext/Short/RelativeTo`; `SetPathBase` has no successor: tests use `t.Chdir`); `webtyp.com/fmt/lang` → `webtyp.com/lang`;
  `MessageType/Msg/StringType` → `webtyp.com/msgtype` (`Type`, constants, `Detect`);
  `EscapeHTML/EscapeAttr` → `escape.HTML`; `JSONEscape` → `escape.JSON`; `Html` → `Sprintf`.
- `AGENTS.md`: under "Key contracts", drop the `Field`/`Fielder`/codec bullets if those files no
  longer exist in this repo (`grep -ln "type Fielder\|type FieldWriter" *.go`). Add the rule
  "tests live in `tests/`; a root-level test only with a written justification; never export a symbol for a test" and the rule "fmt only grows by stdlib equivalents; anything else is
  a new repo".
- `docs/issues/` and `docs/betterment/`: delete files that only discuss moved or deleted features.
- `benchmark/benchmark_results.md`: re-run `benchmark/build-and-measure.sh` and update the numbers
  (after Stage 5).

## Stage 5 — the old name `tinystring` becomes `webtyp`, and the size programs use fmt only

`fmt` was called `tinystring` before the rebrand. The name survives in code comments, benchmark
directories, binaries, scripts and docs (`grep -rli tinystring . | grep -v .git/` lists them).
1. Rename directories with `git mv`: `benchmark/bench-binary-size/tinystring-lib` →
   `benchmark/bench-binary-size/webtyp-lib`, and `benchmark/bench-memory-alloc/tinystring` →
   `benchmark/bench-memory-alloc/webtyp`. Their `go.mod` module names `tinystring-example`/… become
   `webtyp-example`/….
2. In every script (`build-and-measure.sh`, `clean-all.sh`, `update-readme.sh`,
   `memory-benchmark.sh`, `run-all-benchmarks.sh`) and Go file (`analyzer.go`, `reporter.go`,
   `common.go`): the output binary names `tinystring*` → `webtyp*`, variables such as
   `TINYSTRING_RESULTS` → `WEBTYP_RESULTS`, and labels in generated reports → `webtyp`. Update
   `.gitignore` to the new paths.
3. `benchmark/phase13-analysis.sh` and `benchmark/phase13-analysis/` hard-code a Windows path
   (`/c/Users/Cesar/Packages/Internal/tinystring`) and belong to a finished optimization phase:
   delete them.
4. Comments in `convert.go` and `kind.go`, and `AGENTS.md`: `tinystring` → `webtyp/fmt`.
5. `benchmark/bench-memory-alloc/webtyp/main.go` calls `Tilde()`, which Stage 2 deletes: remove that
   call. The size and memory programs must use **only** `webtyp.com/fmt`, never `lang` or another
   webtyp package, because they exist to measure fmt alone. The same holds for
   `example/web/client.go`, which the maintainer already rewrote on 2026-10-06: minimal, fmt only,
   75 KB with TinyGo `-opt=z -no-debug`.
6. Docs in `docs/betterment/` and `docs/issues/` whose subject is gone are deleted (Stage 4). In
   the ones kept, replace `tinystring` with `webtyp/fmt`.

Acceptance for this stage: `grep -rni "tinystring" . --exclude-dir=.git` → empty.

## Acceptance

- `gotest` passes (vet, race, WASM).
- Every `*_test.go` left at the root starts with a `// Root-level test (justified):` comment:
  `for f in *_test.go; do head -1 "$f" | grep -q "Root-level test (justified)" || echo "$f"; done` → empty.
- `grep -rnE "PathJoin|PathShort|StringType|MessageType|EscapeHTML|EscapeAttr|JSONEscape|func Html|TagPairs|Thousands|IDorPrimaryKey|TruncateName|MatchesAny|ReplaceN|HasUpperPrefix|CamelLow|SnakeUp|\) Tilde\(" --include='*.go' .` → empty.
- The PR description lists every removed and every unexported symbol under "BREAKING CHANGE".

## Stages

| # | Stage | Files |
|---|---|---|
| 1 | Tests in `tests/` | `tests/*.go` |
| 2 | Delete / unexport | `parse.go`, `fmt_number.go`, `capitalize.go`, `strings.go`, `search.go`, `convert.go`, `truncate.go`, `mapping.go`, `quote.go`, `error.go`, `memory.go` (+ callers) |
| 2b | `%q` control bytes | `quote.go`, `tests/fmt_q_escape_test.go` |
| 3 | Hook comment | `translate_hook.go` |
| 4 | Docs | `README.md`, `AGENTS.md`, `docs/**`, `benchmark/benchmark_results.md` |
| 5 | Old name + size programs | `benchmark/**`, `convert.go`, `kind.go`, `.gitignore`, `AGENTS.md` |

## Executor notes (2026-10-07, local run)

- Root-level test: `memory_zero_copy_test.go` (justified: pointer identity with the internal
  buffers). Everything else is in `tests/` as `fmt_test`; internals replaced by `Conv.Error()`,
  `runtime.GOARCH`, and a hard-coded kind-name table.
- `Sprintf("%q", x)` for a non-string quotes its string form (`"123"`), documented as differing
  from Go; `%q` now escapes `\a \b \f \v` and `\xHH` for other control bytes and 0x7f.
- Also deleted dead private helpers found by `staticcheck -checks U1000` (both targets):
  `hasUpperPrefix`, `isWasm`, `addThousandSeparatorsCustom`, `removeTrailingZeros`, `spaces`,
  `bL`/`bU`, `isASCIIOnlyOut`, `tildeUnicodeOptimized`, `addRuneToWork`,
  `bufferContainsPattern`, `bytesContain`.
- Memory benchmarks: both programs dropped the accent folding and thousands separators so they
  stay equivalent. Size (ultra WASM): webtyp 26.1 KB vs standard 156.2 KB.
