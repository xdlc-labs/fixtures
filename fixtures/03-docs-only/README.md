# 03-docs-only

**Expect:** no coding-agent Fix. CI may skip or stay green (`paths-ignore` is not set here, so `go test` still runs on the unchanged `src/`). The agent should treat this as docs-only if your skip rules say so; otherwise it is a no-op green PR.
