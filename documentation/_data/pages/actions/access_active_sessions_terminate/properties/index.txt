---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": ["access active sessions terminate"], "body_bytes": 1305, "body_sha256": "sha256:0fb30f264f4a4dc6b3df07fb5c6eeb2fd8908217f88e98af2761d5ebbc42a058", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:fundamentals", "path": "documentation/actions/access_active_sessions_terminate/properties/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "actions", "registry_anchor": "canonical-1011310031122232-1101123230020033-3121112303102021-3303203200231201-2030111000030313-0022320120213221-2333100310201032-0210322102310330", "registry_path": "docs/guides/actions--access_active_sessions_terminate--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["ids"], "anchor": "schema-ids", "description": "List of session IDs to terminate (maximum 100 per request).", "document_id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ids"], "syntax": "attribute", "type": "list"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace of the App type for the current request (path parameter)", "document_id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_access_active_sessions_terminate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
- Property reference

## Direct properties

<a id="schema-ids"></a>

### ids property

Type: `["list", "string"]`. Optional.

List of session IDs to terminate (maximum 100 per request).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request (path parameter)

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ids` | [ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/#schema-ids) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/#schema-namespace) |
