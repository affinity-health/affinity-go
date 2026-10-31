# Changelog

All notable changes to `github.com/affinity-health/affinity-go` are documented here. This project
uses [Semantic Versioning](https://semver.org/spec/v2.0.0.html); before 1.0, a minor version can
break the interface.

## Unreleased

## 0.4.1 - 2026-10-31

### Added

- `client.New(key, options...)`: a client with a 60-second timeout per attempt
  (`client.DefaultTimeout`) and 2 retries (`client.DefaultMaxAttempts` = 3) for 408, 429 and 5xx.
  Retries send the same request, with the same idempotency key. Options after the key replace
  the defaults. `client.NewClient` is unchanged: no timeout and 1 retry.

## 0.4.0 - 2026-10-31

This version is regenerated with Fern from API version `2026-10-28`. It replaces the whole
interface of 0.3.0.

### Changed

- Create the client with `client.NewClient(option.WithToken(key))` from
  `github.com/affinity-health/affinity-go/client`. `affinity.NewClient(key)` is removed.
- Resources are fields of the client: `Account`, `Feedback`, `Medications`, `Orders`,
  `Practices` and `Webhooks`. Each method takes a `context.Context` and one request struct, for
  example `c.Orders.Get(ctx, &affinity.GetOrdersRequest{OrderID: id})`.
- Every write requires `IdempotencyKey` in its request struct. The SDK no longer generates one.
  Save the key with the reviewed request and reuse it to retry.
- List methods return a `*core.Page` with `Iterator()`, which reads every page.
- Errors are typed by status, for example `*affinity.UnprocessableEntityError`, whose
  `GetBody().GetIssues()` names each refused field.
- Requests send `Affinity-Version: 2026-10-28`.
- Retries: at most 2 attempts in total for 408, 429 and 5xx, honoring `Retry-After`. 0.3.0 did
  not retry. The same request, with the same idempotency key, is sent again.
- Timeout: the client has no timeout of its own. 0.3.0 used 60 seconds. Set a context deadline
  or pass `option.WithHTTPClient`.

### Added

- Medications (`List`, `Get` with `Expand: variants`), feedback, practice create, get, update and
  list, webhook endpoints (`List`, `Create`, `Update`, `Remove`, `ResendFailed`), and order
  create, list, get and cancel.
- `docs/reference.md`, the reference for every method.
- A webhook signature verification example in the README.
- `tests/smoke`, which calls every operation against a real API in Test mode.

### Removed

- Patients, catalog items, pharmacies, team, locations, API keys, signing and submission
  routes, the practice-scoped client and the hand-written facade layer. API version
  `2026-10-28` does not have them.

## 0.3.0 - 2026-10-07

- Regenerate from the deployed API contract used by TypeScript SDK 1.16.0 (API `2026-09-28`).

## 0.1.0 - 2026-09-28

- Initial Go SDK preview for Affinity API version `2026-09-28`.
