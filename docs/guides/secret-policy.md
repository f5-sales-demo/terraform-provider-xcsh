---
page_title: "Native Secret Policy lifecycle"
---

# Native Secret Policy lifecycle

Use `xcsh_secret_policy` to own ordered inline ALLOW/DENY rules through the published
`/api/secret_management/namespaces/{namespace}/secret_policys` API. A selector-bound ALLOW
followed by an unconditional DENY scopes Wingman access to the application. Each selector
string combines its comma-separated predicates with AND. Rules evaluate top to bottom;
The provider verifies `algo: FIRST_MATCH` using the full policy response; it does not send
an extra algorithm field in the published CRUD request. A separate replace-form read supplies
normalized inline rules without losing full-response deletion status.

The resource deliberately supports inline selector rules rather than policy chaining or name
matchers. Unsupported stored predicates must be reviewed before adoption; do not replace them
with an empty selector. F5's [how-to](https://docs.cloud.f5.com/docs-v2/multi-cloud-network-connect/how-tos/secret-mgmt/pp-secrets-using-wingman)
and [Wingman API reference](https://docs.cloud.f5.com/docs-v2/platform/reference/wingman-api-reference)
remain authoritative for the service and runtime protocol.

Import a reviewed owned policy using `terraform import xcsh_secret_policy.application demo-app/demo-wingman-policy`.
The provider reads the replace form and resolves legacy child references to inline rules. It sends
one create or replace attempt and verifies named readback. Metadata updates preserve unrelated
annotations and labels. Keep credentials in the provider environment and private TLS inputs
outside this resource.

Deletion is soft. `marked_for_delete` and `deletion_time` are observed API values. Destroy
retains state and reports a pending tombstone until an exact GET returns 404. Repeating destroy
for a marked policy checks absence without sending DELETE again. A 403 is a failure, never absence.
A changed UID stops mutation and deletion so a replacement owner's policy is preserved.

Before deleting a policy, remove its application routes and consumers. Preserve ciphertext and
provenance during recovery; a newly created policy of the same name need not have the same policy
ID and cannot be assumed to decrypt old material. Wait for the recorded tombstone expiry and
verify absence. Do not rotate retained certificate material or edit the policy's read-only deletion
fields to clear a failed destroy. Record actual deletion, credential revocation, and cleanup results
separately from local schema tests.
