---
page_title: "secret-policy"
subcategory: ""
description: "Maintained secret-policy guide."
xcsh_docs: {"aliases": ["adopt existing object", "import existing resource", "secret-policy"], "body_bytes": 2467, "body_sha256": "sha256:27771cb2b95c085c08347ffffff6ff28e10ad4d8b2d4e72a448c11fc40aee341", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved"}, "collection_id": "xcsh-docs:guides:secret-policy:collection", "completeness": "complete", "id": "xcsh-docs:guides:secret-policy:overview", "parent_id": "xcsh-docs:provider:xcsh:navigation", "path": "documentation/guides/secret-policy/index.md", "product": "distributed-cloud", "provider_name": "secret-policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "guides", "registry_path": "docs/guides/secret-policy.md", "relationships": [], "retrieval_version": 1, "role": "overview", "schema_path": [], "schema_version": 1, "sections": [], "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Maintained secret-policy guide.", "tasks": ["configuration"]}
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
