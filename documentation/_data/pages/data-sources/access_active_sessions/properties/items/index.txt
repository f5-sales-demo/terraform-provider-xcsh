---
page_title: "items"
subcategory: ""
description: "Sessions. List of active sessions."
xcsh_docs: {"aliases": ["items"], "body_bytes": 2446, "body_sha256": "sha256:754e9ad1c757009e87ae76cb213c2a2baf2604d1f6d6f0f417735938c984975f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "parent_id": "xcsh-docs:data-sources:access_active_sessions:reference", "path": "documentation/data-sources/access_active_sessions/properties/items/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1303232133022210-3100033221123233-0202011333221331-3223033232001012-1213111021112121-1021323200312311-2031213332320103-1011232020311322", "registry_path": "docs/guides/data-sources--access_active_sessions--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items"], "schema_version": 1, "sections": [{"aliases": ["items client ip"], "anchor": "schema-items--client_ip", "description": "Client IP of the user that connected via the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "client_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["items expiration time"], "anchor": "schema-items--expiration_time", "description": "Expiration time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "expiration_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["items id"], "anchor": "schema-items--id", "description": "Session ID. ID of the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items last activity time"], "anchor": "schema-items--last_activity_time", "description": "Last activity time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "last_activity_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["items policy"], "anchor": "schema-items--policy", "description": "Policy. Policy associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items site"], "anchor": "schema-items--site", "description": "Site. Site where the session is created.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "site"], "syntax": "attribute", "type": "string"}, {"aliases": ["items start time"], "anchor": "schema-items--start_time", "description": "Start time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["items status"], "anchor": "schema-items--status", "description": "The actual status of an active session Session status is pending Session status is established. Possible values are `PENDING`, `ESTABLISHED`. Defaults to `PENDING`.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ESTABLISHED", "PENDING"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "status"], "syntax": "attribute", "type": "string"}, {"aliases": ["items username"], "anchor": "schema-items--username", "description": "Username used in authentication of this session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["items virtual server"], "anchor": "schema-items--virtual_server", "description": "Virtual server associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "virtual_server"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/properties/items/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Sessions. List of active sessions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ESTABLISHED","PENDING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
