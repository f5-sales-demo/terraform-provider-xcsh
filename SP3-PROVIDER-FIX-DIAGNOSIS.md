# SP3 provider codegen fix — diagnosis & design

Branch: `fix/sp3-nested-marker-and-optional-scalar-suppression`
Surfaced by: webapp-api-protection SP3 API-Protection live matrix (issue #41).

## Symptom
Applying `api_protection_rules.api_endpoint_rules[]` (and `validation_custom_list`)
fails with either:
- `Provider produced inconsistent result after apply: ...client_matcher.any_client:
  was absent, but now present` (and `...api_endpoint_method.invert_matcher: was null,
  but now cty.False`), or
- round-trip import-drift on the same fields.

## Root cause (precise)
`renderUnmarshalSingleChild` sets
`preserveWhole = container=="single" && stateBase!="" && !hasObjectReferenceDescendant(attr)`
(render.go:779). `hasObjectReferenceDescendant` returns true if the block contains an
object-ref (a `tenant` child) at ANY depth. `api_protection_rules.api_endpoint_rules[]
.client_matcher` contains object-ref arms in its *deep, rarely-used* branches
(`ip_matcher.prefix_sets`, `asn_matcher.asn_sets` — kind/name/namespace/tenant/uid).
So `hasObjectReferenceDescendant(api_protection_rules)` is true → `preserveWhole=false`
→ the whole block is RECONSTRUCTED from the API response.

The server normalizes `client_matcher` to include `any_client: {}` as a base marker
ALONGSIDE `ip_prefix_list` (verified live: the GET returns both), and returns
`invert_matcher: false`. Reconstruction materializes these server-echoed
defaults, but the plan omitted them → "was absent/null, now present/false".

The same class hits `validation_custom_list` (deep block, server-materialized optionals)
on the import path.

## Why the existing mechanisms don't cover it
- `isImportDefaultSuppressed` guards response-populate with `!isImport` — helps IMPORT
  only, not the apply-time (Create/Update) inconsistency.
- `preserveWhole` is all-or-nothing: any object-ref descendant disables preservation
  for the ENTIRE block, even when the actually-configured arms carry no object-ref.

## Fix design (chosen: preserve + tenant-patch merge)
On the apply path (`!isImport`), preserve the PLANNED block value even when it has
object-ref descendants, then recursively patch ONLY the object-ref Computed leaves
(`tenant`, and `uid`/`kind` where server-derived) from the API response into the
preserved value. This keeps user-set Optional markers/scalars (any_client absence,
invert_matcher) intact while still fixing the #1091 tenant-unknown problem.

Implementation sketch:
- Replace the `preserveWhole` gate so single blocks are preserved whenever configured
  (`stateBase!=""` and planned non-nil), regardless of object-ref descendants.
- Add a generated `patchObjectRefTenants(planned, apiResponse)` walk (per resource,
  over the object-ref descendant paths) invoked in the preserve branch, copying
  server tenant/uid/kind into the preserved value.
- For LIST-nested blocks (api_endpoint_rules), thread the prior-state list positionally
  so element Optional leaves are preserved too (currently stateBase="" for list-child
  reconstruction — render.go:858).
- Regen + `go test ./tools/... ./internal/...` + rebuild + re-run the webapp SP3 matrix
  to full green; release; bump webapp versions.tf.

## Interim module state (webapp)
The SP3 module explicitly sets `invert_matcher=false` and
`request_validation_properties` to satisfy F5 XC; `metadata{name}` added to
api_endpoint_rules + data_guard_rules. Live-clean arms: rate_limit, sensitive_data
(custom), data_guard, api_protection with the default `any_client` matcher. The
`ip_prefix`/`ip_threat` client-matcher on protection rules + `validation_custom_list`
await this provider fix to be round-trip import-clean.
