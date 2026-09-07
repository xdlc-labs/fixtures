# fixtures

Scratch repo for [xdlc-agent](https://github.com/xdlc-labs/xdlc-agent) CI Fix tests.

[![CI](https://github.com/xdlc-labs/fixtures/actions/workflows/ci.yml/badge.svg)](https://github.com/xdlc-labs/fixtures/actions/workflows/ci.yml)

The app under review is `src/`. Each fixture overlays files onto that tree and opens a real GitHub PR. Expected outcomes live **only** in the fixture `README.md`, which is never shipped in the PR.

The public battleground for GitOps / `/healthz` is [example-service](https://github.com/xdlc-labs/example-service). This repo is the graded scratch pad so `example-service` `main` stays clean.

## Bait sterility

Everything the agent can see must read like a real engineer's PR. If the diff, title, or body says "fixture", "E2E", or "planted bug", the model excuses the defect and the fixture tests nothing.

- Overlay comments are production comments. Never annotate the defect.
- `meta.env` `TITLE` / `BODY` are the PR title/body — natural change descriptions.
- Grading stays in `fixtures/<id>/README.md`.

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

## Fixture index

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

Docs: [Getting started](https://xdlc-labs.github.io/documentation/xdlc-agent/getting-started/).

## License

[MIT](LICENSE)
