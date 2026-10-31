#!/usr/bin/env bash
# Generates the Go SDK with Fern from spec/affinity.openapi.json, the committed Affinity contract.
# Fern runs locally in Docker; no Fern account or token is used. Needs Docker, Node (npx) and jq.
#
# Every file at the repository root is generated except the hand-written ones kept below.
set -euo pipefail
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root"

fern_version="$(jq -r .version fern/fern.config.json)"
generator="fernapi/fern-go-sdk"
generator_version="$(awk '/fern-go-sdk/ { getline; print $2 }' fern/generators.yml)"
sdk_version="0.4.0"
output="$root/.fern-output"

# Hand-written files and directories. Everything else at the root is Fern output.
keep=(.git .github .gitignore AGENTS.md CHANGELOG.md CONTRIBUTING.md LICENSE README.md SECURITY.md
  docs fern generation.json scripts spec tests .fern-output)

jq -f scripts/prepare-spec.jq spec/affinity.openapi.json >fern/openapi.json
rm -rf "$output"
env -u FERN_TOKEN npx --yes "fern-api@$fern_version" generate --group go --local --no-prompt \
  --force --version "$sdk_version"
# Fern can exit 0 after its container fails; require the client it writes.
test -f "$output/client/client.go" || { echo "Fern wrote no client/client.go" >&2; exit 1; }

mv "$output/reference.md" docs/reference.md
# Fern's README and CONTRIBUTING are replaced by ours; .fern/metadata.json records the local git
# commit, so it would change on every run.
rm -rf "$output/README.md" "$output/CONTRIBUTING.md" "$output/.fern"

shopt -s dotglob nullglob
for entry in *; do
  skip=false
  for name in "${keep[@]}"; do [[ "$entry" == "$name" ]] && skip=true; done
  $skip || rm -rf "$entry"
done
for entry in "$output"/*; do
  name="$(basename "$entry")"
  for kept in "${keep[@]}"; do
    [[ "$name" == "$kept" ]] && { echo "Fern wrote $name, a hand-written path" >&2; exit 1; }
  done
  cp -R "$entry" "$root/"
done
rm -rf "$output" fern/openapi.json

jq -n \
  --arg contractVersion "$(jq -r .info.version spec/affinity.openapi.json)" \
  --arg contractSha256 "$(sha256sum spec/affinity.openapi.json | cut -d' ' -f1)" \
  --arg fernVersion "$fern_version" \
  --arg generator "$generator" \
  --arg generatorVersion "$generator_version" \
  --argjson operations "$(jq '[.paths | to_entries[] | .key as $path | .value | to_entries[]
      | select(.value | type == "object" and has("operationId"))
      | {key: .value.operationId, value: "\(.key | ascii_upcase) \($path)"}] | from_entries' \
      spec/affinity.openapi.json)" \
  '{contractVersion: $contractVersion, contractSha256: $contractSha256, fernVersion: $fernVersion,
    generator: $generator, generatorVersion: $generatorVersion, operations: $operations}' \
  >generation.json
echo "Generated the Go SDK from API $(jq -r .contractVersion generation.json):" \
  "$(jq '.operations | length' generation.json) operations."
