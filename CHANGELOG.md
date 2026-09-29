# Changelog

## Unreleased

### Changed

- Use direct API-key constructors, short nested resource methods, and separate request options across all 73 API operations.
- Resolve practice keys automatically and support immutable scoped clients or per-request practice context.
- Generate keys for routine writes, including patient deletion; require persisted keys for consequential order actions.
- Add lazy pagination iterators and consistent API error metadata.
- Accept patient status `archived` as an alias for the API's `inactive` status.
- Default to API `2026-09-28`, a 60-second timeout, and no automatic retries.
- Verify the language guide examples against the implemented client.

## 0.1.0 - 2026-09-28

- Initial Go SDK preview for Affinity API version `2026-09-28`.
- Generated resource clients, request and response types, and transport verification.
