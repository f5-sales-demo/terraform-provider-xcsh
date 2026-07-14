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

## Fix design (implemented: state-threaded reconstruction)
The originally-sketched "preserve whole + separate patchObjectRefTenants walk" was
replaced by a cleaner, lower-risk equivalent that reuses the provider's existing,
already-tested per-leaf preserve logic (scalar preserve at render.go:612, empty-marker
preserve at render.go:740) instead of generating a new recursive patch walker per
resource. Same semantics — off-spine Optional markers/scalars reflect the PLAN, object
references reconstruct their Computed `tenant` from the API — but achieved by threading
the prior-state accessor down into the reconstruction rather than post-patching a
preserved copy.

Two changes in `tools/pkg/codegen/render.go`:

1. **`renderUnmarshalSingleChild` — thread state into a "spine" block's children.**
   `preserveWhole` is unchanged (still only for blocks with NO object-ref descendant, so
   direct/nested refs still reconstruct — #1079/#1091 preserved). When a block merely
   CONTAINS a reference on one arm and so cannot be preserved whole, its children are now
   reconstructed with the prior-state accessor threaded in (`childStateBase =
   stateBase.<Field>`), UNLESS the block *is* itself an object reference (all its leaves
   are server-derived). Off-spine children then hit their existing preserve paths (return
   the planned marker/scalar); the reference arm, when recursion reaches it, reads its
   Computed tenant from the API.

2. **`renderUnmarshalListChild` — positional element-state threading.** A list configured
   inside a single block now loads its prior-state elements (`<Field>.ElementsAs` — nested
   lists are always `types.List`) and threads `existing[idx]` + `len(existing) > idx` into
   each element's children, mirroring `renderUnmarshalTopLevelList`. This preserves element
   Optional leaves (api_endpoint_method.invert_matcher, client_matcher.any_client) that
   were previously reconstructed from the API (stateBase="" at the old render.go:857).

Why this is correct and low-risk: it makes NESTED reconstruction consistent with the
TOP-LEVEL list renderer, which already threads positional state. The three existing
object-ref tests (`ObjectRefReadsFromAPI`, `NonRefPreserves`, `NestedObjectRefReconstructs`)
still pass unchanged — they exercise ref leaves only, whose behavior is untouched. New
TDD tests: `TestRenderUnmarshalSingleChild_SpinePreservesOffSpineLeaves`,
`TestRenderUnmarshalListChild_PreservesElementStatePositionally`.

Verification: `go test ./tools/... ./internal/...` green; all 128 resources regenerate
and compile; webapp defaults plan 0-change; SP3 live matrix to full green; then release +
bump webapp versions.tf.

## Interim module state (webapp)
The SP3 module explicitly sets `invert_matcher=false` and
`request_validation_properties` to satisfy F5 XC; `metadata{name}` added to
api_endpoint_rules + data_guard_rules. Live-clean arms: rate_limit, sensitive_data
(custom), data_guard, api_protection with the default `any_client` matcher. The
`ip_prefix`/`ip_threat` client-matcher on protection rules + `validation_custom_list`
await this provider fix to be round-trip import-clean.
