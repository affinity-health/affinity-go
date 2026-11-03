# Changelog

All notable changes to `github.com/affinity-health/affinity-go` are documented here. This project
uses [Semantic Versioning](https://semver.org/spec/v2.0.0.html); before 1.0, a minor version can
break the interface.

## Unreleased

## 0.7.0 - 2026-11-03

Generated from the corrected API `2026-11-01` document. Requests and responses on the wire are
unchanged.

### Removed

- `affinity.MethodNotAllowedError`. The API does not answer `405` to an SDK call. A `405` from a proxy raises a `*core.APIError`.
- The `ContentTooLargeErrorBody`, `ServiceUnavailableErrorBody` and their `…Code` and `…CodeCode` types and constants. `ContentTooLargeError.Body` and `ServiceUnavailableError.Body` are now `*Problem`, like the other errors (before, `ContentTooLargeError.Body` was `any`).

### Changed

- Clearer descriptions in the types and reference. Every enum member keeps its name.

## 0.6.0 - 2026-11-02

This release is generated from the API `2026-11-01` contract that production serves. That contract
marks each response object and response value list as open: a later API version can add a field or
a value.

### Changed

- Response values keep their known-value types, such as `OrderStatus` and `AllergyCategory`. A response with an
  unknown value or an added field, at the root or nested, decodes without an error: the unknown value is kept as a string.
- Every response type has a public `ExtraProperties` map with the fields this SDK does not know
  (before, only the webhook types had it). Marshaling a response type writes them back.

## 0.5.0 - 2026-11-01

This release targets API `2026-11-01`, which breaks `2026-10-28`.

### Changed

- Every request sends `Affinity-Version: 2026-11-01`.
- A descending sort is `<field>_desc`, such as `affinity.ListPracticesRequestSortCreatedDesc`: `created_desc` and `patient_desc`
  for orders, and `name_desc`, `location_desc`, `mode_desc` and `created_desc` for practices. The
  `-created` values are gone; the API refuses them.

### Added

- The `awaiting_approval` order status, in orders, webhook orders and the `orders.list` status
  filter.

## 0.4.2 - 2026-11-01

### Fixed

- Descending sorts are in the sort types: `-created` and `-patient` for orders, and `-name`,
  `-location`, `-mode` and `-created` for practices. Before, the generator dropped each `-` value,
  so a typed call could not reverse a sort.

### Removed

- The `EndingBefore` list parameter. Lists page forward only. After a page fetched with
  `EndingBefore`, the next page also sent the forward cursor, and the API refused it with a 400.

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
