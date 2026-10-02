---
page_title: "items"
subcategory: ""
description: "Sessions. List of active sessions."
xcsh_docs: {"aliases": ["items"], "body_bytes": 2456, "body_sha256": "sha256:6751823ee50d1de40d0253d4d811c867db2bf1ebb2a8e0e05f7d05fea0c5a643", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "parent_id": "xcsh-docs:data-sources:access_active_sessions:reference", "path": "documentation/data-sources/access_active_sessions/properties/items/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303232133022210-3100033221123233-0202011333221331-3223033232001012-1213111021112121-1021323200312311-2031213332320103-1011232020311322", "registry_path": "docs/guides/data-sources--access_active_sessions--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items"], "schema_version": 1, "sections": [{"aliases": ["client ip"], "anchor": "schema-items--client_ip", "description": "Client IP of the user that connected via the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "client_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["expiration time"], "anchor": "schema-items--expiration_time", "description": "Expiration time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "expiration_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-items--id", "description": "Session ID. ID of the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["last activity time"], "anchor": "schema-items--last_activity_time", "description": "Last activity time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "last_activity_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy"], "anchor": "schema-items--policy", "description": "Policy. Policy associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["site"], "anchor": "schema-items--site", "description": "Site. Site where the session is created.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "site"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-items--start_time", "description": "Start time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["status"], "anchor": "schema-items--status", "description": "The actual status of an active session Session status is pending Session status is established. Possible values are `PENDING`, `ESTABLISHED`. Defaults to `PENDING`.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "status"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "username"], "anchor": "schema-items--username", "description": "Username used in authentication of this session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server"], "anchor": "schema-items--virtual_server", "description": "Virtual server associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "virtual_server"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/properties/items/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Sessions. List of active sessions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items

Breadcrumbs:

- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Sessions. List of active sessions.

## Direct properties

<a id="schema-items--client_ip"></a>

### client_ip property

Type: `"string"`. Computed.

Client IP of the user that connected via the session.

<a id="schema-items--expiration_time"></a>

### expiration_time property

Type: `"string"`. Computed.

Expiration time of the session in RFC 3339 format.

<a id="schema-items--id"></a>

### id property

Type: `"string"`. Computed.

Session ID. ID of the session.

<a id="schema-items--last_activity_time"></a>

### last_activity_time property

Type: `"string"`. Computed.

Last activity time of the session in RFC 3339 format.

<a id="schema-items--policy"></a>

### policy property

Type: `"string"`. Computed.

Policy. Policy associated with the session.

<a id="schema-items--site"></a>

### site property

Type: `"string"`. Computed.

Site. Site where the session is created.

<a id="schema-items--start_time"></a>

### start_time property

Type: `"string"`. Computed.

Start time of the session in RFC 3339 format.

<a id="schema-items--status"></a>

### status property

Type: `"string"`. Computed.

\[Enum: PENDING|ESTABLISHED\] The actual status of an active session Session status is pending
Session status is established. Possible values are \`PENDING\`, \`ESTABLISHED\`. Defaults to
\`PENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PENDING",
    "ESTABLISHED"),
}
```

<a id="schema-items--username"></a>

### username property

Type: `"string"`. Computed.

Username used in authentication of this session.

<a id="schema-items--virtual_server"></a>

### virtual_server property

Type: `"string"`. Computed.

Virtual server associated with the session.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
