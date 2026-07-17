# Affinity Go SDK

The official Go SDK for the Affinity API.

> **Status:** This repository is being prepared for its first release. The module is not yet
> published and its public API may change before `v0.1.0`.

The SDK will provide a small, resource-oriented client for Go services connecting healthcare
practices to Affinity's compounder network.

## Planned module

```sh
go get github.com/affinity-health/affinity-go
```

## Intended usage

```go
package main

import (
	"context"
	"os"

	affinity "github.com/affinity-health/affinity-go"
)

func main() {
	client := affinity.NewClient(os.Getenv("AFFINITY_API_KEY"), affinity.ClientOptions{
		APIVersion: "2026-07-09",
	})

	ctx := context.Background()
	catalog, err := client.Catalog.List(ctx, affinity.CatalogListParams{
		Query: "semaglutide",
	})
	if err != nil {
		panic(err)
	}

	_ = catalog
}
```

## Resource model

The initial client surface is planned around these resources:

- `Account` — inspect the authenticated organization and API access
- `Catalog` — search products available through the Affinity network
- `Practices` — create and manage customer practices
- `Orders` — create, submit, inspect, update, and cancel orders
- `Webhooks` — manage endpoints and inspect or replay events

All network operations will accept `context.Context`. Generated transport types will remain
available as an escape hatch, while `NewClient` will be the recommended entry point.

## Safety

Affinity API keys are service-account credentials. Use this SDK only in a trusted backend and load
keys from server-side secret storage. Do not embed a key into distributed client binaries or expose
it through logs.

Requests involving patient, prescription, or fulfillment data may contain protected health
information. Integrators are responsible for their own authorization, logging, retention,
infrastructure, and compliance controls.

## Generation and releases

This SDK will be generated from the versioned contract in
[`affinity-openapi`](https://github.com/affinity-health/affinity-openapi), with a maintained
resource facade layered over the generated transport. Releases will be validated against the same
contract before version tags are published.

## Related projects

- [OpenAPI specification](https://github.com/affinity-health/affinity-openapi)
- [TypeScript SDK](https://github.com/affinity-health/affinity-typescript)

