#!/usr/bin/env bash
# Regenerates the SDK from spec/affinity.openapi.json and fails if anything differs from the
# committed files: generated code is never edited by hand. Needs Docker, Node and jq, and a
# repository without uncommitted changes to tracked files.
set -euo pipefail
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root"
bash scripts/generate.sh >/dev/null
git add -N .
git diff --exit-code --stat
