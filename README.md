# fixtures

Sample pull requests that fail CI, so you can watch [xdlc-agent](https://github.com/xdlc-labs/xdlc-agent) open a **Fix**.

[![CI](https://github.com/xdlc-labs/fixtures/actions/workflows/ci.yml/badge.svg)](https://github.com/xdlc-labs/fixtures/actions/workflows/ci.yml)

Point the daemon at this repo, apply a fixture, wait for CI, then see a Fix (or no Fix). The app under review is `src/`. Each fixture copies files onto that tree and opens a real GitHub PR.

Want a running HTTP demo (`/healthz`, `/metrics`) instead? Use [example-service](https://github.com/xdlc-labs/example-service). Leave that repo's `main` alone. Try failing PRs here.

## Keep the PR looking real

The coding agent sees the PR. If the title, body, or diff says "fixture", "E2E", or "planted bug", the model excuses the defect and you learn nothing.

Expected outcomes live **only** in `fixtures/<id>/README.md`, which is never shipped in the PR.

- Overlay comments are production comments. Never annotate the defect.
- `meta.env` `TITLE` / `BODY` are the PR title and body. Write them like a normal change.

## Layout

```
src/                         baseline Go service
fixtures/<NN-name>/
  overlay/                   files copied onto src/
  meta.env                   BRANCH / TITLE / BODY
  README.md                  expected outcomes (not in the PR)
scripts/apply-fixture.sh
```

## Workflow

```bash
./scripts/apply-fixture.sh 01-ci-fail
# wait for CI + xdlc Fix, then:
gh pr close <N> --delete-branch
```

`apply-fixture.sh` always resets to `main` before overlaying.

## What each fixture does

| ID | Name | Behavior | Expect |
|----|------|----------|--------|
| 01 | `01-ci-fail` | Checksum helper off-by-one → red CI | Fix |
| 02 | `02-clean` | Comment-only change → green CI | no Fix |
| 03 | `03-docs-only` | Markdown only | skip / no agent |
| 04 | `04-unfixable` | Test needs an upstream that is not there | `BACKLOG.md`, no silent green |

## Point xdlc at this repo

```yaml
repos:
  - name: fixtures
    github: xdlc-labs/fixtures
    gates: [ci]
```

Then [install xdlc](https://xdlc.dev/agent/docs) and run `xdlc daemon`. Walkthrough: [Getting started](https://xdlc.dev/agent/docs/getting-started).

## License

[MIT](LICENSE)
