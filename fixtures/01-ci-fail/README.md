# 01-ci-fail

**Expect:** GitHub Actions red. xdlc-agent opens a **Fix** (or a Fix PR, depending on `agent.fix.mode`).

Overlay makes `Fold` return `uint16(n & 0xff)` and replaces `TestFold` with a sum that exceeds 255 (`[]byte{200, 200}` wants 400). The PR must not mention fixtures.

After a good Fix: `go test ./...` in `src/` is green again.
