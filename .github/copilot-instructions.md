# GitHub Copilot / AI Agent Instructions for minitools ✅

Purpose
- Help an AI code agent become productive quickly: understand the repo layout, conventions, test/workflow commands, and integration points.

Big picture (why and structure) 🔧
- Module: `github.com/qmaru/minitools/v2` (Go module, go 1.25).
- This is a small utilities library split into focused packages under directories like `data/`, `encoding/`, `file/`, `hashx/`, `random/`, `secret/`, `time/`, and `utils/convert/`.
- Several packages provide multiple backend implementations (example: `data/json/` contains `common`, `gojson`, `sonic`, `standard`) — expect pluggable implementations and performance experimentation.

Key conventions & patterns (follow these precisely) 🧭
- Factories: most packages expose a `New()` constructor (e.g., `file.New()`, `gcm.New()`, `nanoid.New()`). Use those in examples and tests.
- Methods return `(T, error)` or `error` for operations that can fail; follow existing error handling style from tests.
- Use existing tests as canonical usage examples (tests are the preferred documentation source). See `minitools_test.go` for representative usages (e.g., `gcm.New().Encrypt(plain, key)`, `file.New().RootPath("")`, `dedupe.New().Do(key, fn)`).
- Keep APIs small and Go-idiomatic: short method names (e.g., `Encrypt`, `Decrypt`, `Sum256`, `ToBase64`, `Generate`).

Tests & developer workflow ⚙️
- Run unit tests: `go test ./... -v`
- Run benchmarks: `go test -bench=. -benchmem ./...` (benchmarks exist in `minitools_bench_test.go`).
- Typical debugging: run a single test with `go test -run TestSecretAes -v` to reproduce crypto examples.
- Formatting & vetting: repo uses standard `go fmt`/`go vet` expectations (no custom tooling detected). Use `go mod tidy` after dependency changes.

Important implementation details to keep in mind 🔍
- Crypto key sizes are meaningful in tests: AES examples use a 16-byte key (`"length is 16 bit"`) and an IV the same length; ChaCha20 examples use 32-byte keys.
- Hash packages expose convenient helpers: e.g., `sha256.New().Sum256([]byte("...")).ToBase64()` or Murmur3 returns uint32 via `.ToUint32()`.
- Encoding helpers in `encoding/text` provide encode/decode pairs for base62/base64/hex with decoding helpers that expose `DecodeString()` and `DecodeRaw()`.
- Dedupe pattern: `dedupe.New().Do(key, func() (any, error) { ... })` deduplicates concurrent calls by key (used in `minitools_test.go`).

External dependencies & integration points 🌐
- Notable libs in `go.mod`: `github.com/bytedance/sonic`, `github.com/goccy/go-json`, `lukechampine.com/blake3`, `github.com/matoous/go-nanoid`, `github.com/gofrs/uuid/v5`, `github.com/spaolacci/murmur3`, `github.com/sqids/sqids-go`, `golang.org/x/crypto`.
- When modifying JSON code, check all implementations under `data/json/` for compatibility and tests.

Guidance for changes and PRs 📦
- Add or update tests that show typical usage (unit tests are treated as usage docs).
- If adding new public API, provide a `New()` constructor to match repo style and add tests demonstrating default behavior.
- Update `go.mod` and run `go mod tidy`; ensure unit tests and benchmarks pass locally.
- For breaking changes, follow semantic versioning in the module path (`/v2`).

Files to inspect for examples and patterns 🔎
- `minitools_test.go` — canonical usage examples across packages
- `minitools_bench_test.go` — benchmarks
- `data/json/*` — multiple JSON backends
- `secret/aes`, `secret/chacha20`, `secret/xor` — crypto primitives
- `hashx/*` — hash implementations
- `file/file.go`, `utils/convert/convert.go`

If anything below is unclear or you want more detail (e.g., specific package deep dives, test sample snippets, or missing CI details), please point me to the area to expand. 🙋‍♂️