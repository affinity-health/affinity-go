# Affinity Go SDK agent guide

This repository is the official Go SDK for the Affinity public API. Module
`github.com/affinity-health/affinity-go`, package `affinity`.

## Source of truth

- `spec/affinity.openapi.json` is the API contract, copied unchanged from the TypeScript SDK
  (`affinity-health/affinity-typescript`, `spec/affinity.openapi.json`). Do not edit it here.
- Fern generates every Go file outside `tests/`, plus `go.mod`, `go.sum`, `wiremock/` and
  `docs/reference.md`. Never hand-edit them. Change the contract upstream, copy the new snapshot,
  and run `scripts/generate.sh`. Fix generation problems in `fern/generators.yml`,
  `scripts/prepare-spec.jq` or `scripts/generate.sh`.
- `fern/generators.yml` pins the Fern CLI (via `fern/fern.config.json`) and the
  `fernapi/fern-go-sdk` version. `generation.json` records the inputs of the last generation.

## Hand-written files

`README.md`, `CHANGELOG.md`, `AGENTS.md`, `CONTRIBUTING.md`, `SECURITY.md`, `LICENSE`,
`.gitignore`, `.github/`, `fern/`, `scripts/`, `spec/` and `tests/` (`tests/smoke` and
`tests/webhooksig`). `scripts/generate.sh` keeps exactly these and replaces everything else at
the root. The webhook verifier in `tests/webhooksig/verify.go` must stay identical to the README
example.

## Commands

```sh
scripts/generate.sh         # regenerate (Docker, Node, jq)
scripts/generate-check.sh   # regenerate and fail on any diff
docker run --rm -v "$PWD":/src -w /src golang:1.23 \
  sh -c 'go build ./... && go vet ./... && go test $(go list ./... | grep -v "_test$")'
docker run --rm --env-file smoke.env -v "$PWD":/src -w /src golang:1.23 go run ./tests/smoke
```

## Rules

- Do not publish or tag a release without explicit authorization. The module is installed from
  GitHub tags only.
- Run the smoke test with Test keys (`sk_test_…`) only. Never call production with it. Never
  print, log or commit API keys.
- Keep examples synthetic and Test mode only. The SDK is for trusted server-side code; never
  encourage API keys in browser or mobile code.
- A provider is a clinician or prescriber; a practice is the customer organization. Do not add a
  `clinic` resource.
