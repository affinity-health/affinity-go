# Affinity Go SDK

Server-side client for the Affinity API. Requires Go 1.21+.

The new interface is implemented in this source update and has not been published to a registry yet.

## Install from source

```sh
go get github.com/affinity-health/affinity-go@main
```

## Use

Set `AFFINITY_API_KEY` to a Test practice key on your server. Keep API keys out of browser and mobile code.

```go
api := affinity.NewClient(os.Getenv("AFFINITY_API_KEY"))
patients, err := api.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
```

Import the module as `affinity "github.com/affinity-health/affinity-go"`. Use `context.Context` for requests and handle `err` before reading the result.

Practice keys identify their practice automatically. Platform keys pass a practice ID in request options or use a scoped client.

See the [SDK guide](docs/guide.md) for platform requests, patient updates, signing, submission, pagination, and errors.
Routine patient writes generate an idempotency key. Persist your own keys for order creation, signing, and submission.

Defaults: API `2026-09-28`, a 60-second timeout, and no automatic retries.

## Verify

```sh
go build ./...
./tests/with-fixtures.sh go test ./smoke
```

The tests use synthetic fixtures on loopback. The fixture runner requires Python 3; Docker runs the language toolchain for the `scripts/check.sh` commands.

Generated with Cloudflare Forge, Fern, and Affinity's facade generator. [generation.json](generation.json) records the pinned inputs. Fix the generator in the Affinity monorepo before regenerating client code.
