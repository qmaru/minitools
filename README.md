# minitools

A small, focused collection of Go utility packages (hashing, encoding, crypto, random IDs, file helpers, JSON backends, and time helpers).

Module: `github.com/qmaru/minitools/v2` (Go 1.25)

## Quick start ✅

- Run all unit tests:

```bash
go test ./... -v
```

- Run a single test (useful for reproducing crypto examples):

```bash
go test -run TestSecretAes -v
```

- Run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

- Tidy deps after changes:

```bash
go mod tidy
```

## Highlights & Usage examples 🔧

- Factories: most packages expose a `New()` constructor. Example patterns from `minitools_test.go`:
  - AES GCM/CBC: `gcm.New().Encrypt(plain, key)` and `cbc.New().Encrypt(plain, key, iv)` (keys/IV sizes are important in tests)
  - Dedupe: `dedupe.New().Do("key", func() (any, error) { ... })` (deduplicates concurrent calls)
  - JSON backends: `gojson.New().Json.NewDecoder(...)` and `RawJson2Map` helper
  - Hashing: `sha256.New().Sum256([]byte("...")).ToBase64()`
  - Random: `nanoid.New().New(10)` and `uuid.New().Generate(...)`

Use the tests in `minitools_test.go` as canonical examples — they are intended as documentation.

## Tests & Conventions ⚙️

- APIs are small and Go-idiomatic (short method names, explicit error returns).
- Add a `New()` constructor for new packages and include tests demonstrating expected behavior.
- Formatting and vetting use standard `go fmt` and `go vet` practices.

## Notable files/directories 📁

- `minitools_test.go` — canonical usage examples
- `minitools_bench_test.go` — benchmarks
- `data/json/*` — multiple JSON backends (check compatibility when changing JSON logic)
- `secret/*` — crypto helpers (pay attention to key/nonce sizes in tests)
- `hashx/*`, `encoding/text` — hash & encoding helpers

## Contributing & PRs 📦

- Add/update unit tests that demonstrate typical usage (tests are the preferred documentation source).
- Update `go.mod` and run `go mod tidy` when adding/removing dependencies.
- Ensure tests and benchmarks pass locally before submitting a PR.

## AI agents & local guidance 🤖

See `.github/copilot-instructions.md` for repository-specific guidance for AI coding agents (examples, patterns, and things to watch for).

## License

This project is licensed under the terms in `LICENSE`.
