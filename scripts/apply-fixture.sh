#!/usr/bin/env bash
# Overlay a fixture onto src/, push a branch, open a PR.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
ID="${1:-}"
[[ -n "$ID" ]] || { echo "usage: $0 <fixture-id>" >&2; exit 1; }
DIR="$ROOT/fixtures/$ID"
[[ -d "$DIR/overlay" ]] || { echo "missing $DIR/overlay" >&2; exit 1; }
# TITLE/BODY often contain colons and spaces (`fix: include high byte`).
# `source` treats `TITLE=fix: include …` as TITLE=fix: plus a command.
while IFS= read -r line || [[ -n "$line" ]]; do
  [[ -z "$line" || "$line" == \#* ]] && continue
  key="${line%%=*}"
  val="${line#*=}"
  if [[ ${#val} -ge 2 && "$val" == \'*\' ]]; then
    val="${val:1:${#val}-2}"
  elif [[ ${#val} -ge 2 && "$val" == \"*\" ]]; then
    val="${val:1:${#val}-2}"
  fi
  printf -v "$key" '%s' "$val"
done < "$DIR/meta.env"

cd "$ROOT"
git fetch origin main
git checkout main
git reset --hard origin/main
git checkout -B "$BRANCH"
cp -R "$DIR/overlay/." "$ROOT/src/"
git add -A src
git status --short
git commit -m "$TITLE"
git push -u origin "$BRANCH" --force
gh pr create --title "$TITLE" --body "$BODY" --base main --head "$BRANCH"
echo "opened PR for $ID — grade against $DIR/README.md (not the PR body)"
