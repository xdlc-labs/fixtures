# 01-ci-fail

**Expect:** GitHub Actions red. xdlc-agent opens a **Fix** (or a Fix PR, depending on `agent.fix.mode`).

Overlay makes `Fold` return `uint16(n & 0xff)` so `TestFold` fails. The PR must not mention fixtures.

After a good Fix: `go test ./...` in `src/` is green again.
