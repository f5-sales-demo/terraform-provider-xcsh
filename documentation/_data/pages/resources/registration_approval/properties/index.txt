---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration_approval."
xcsh_docs: {"aliases": ["registration approval"], "body_bytes": 3239, "body_sha256": "sha256:733522018a023a208c3635a1e7ef3f061de6d789702c5ef4fb320ddb974b36cd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration_approval:reference", "parent_id": "xcsh-docs:resources:registration_approval:fundamentals", "path": "documentation/resources/registration_approval/properties/index.md", "product": "distributed-cloud", "provider_name": "registration_approval", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0221110020100131-3031023122311011-2222002100222120-2132023011212021-1212033310020221-0302211021123030-3031320113201232-1302122130102300", "registry_path": "docs/guides/resources--registration_approval--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["backup connected region"], "anchor": "schema-backup_connected_region", "description": "backup connected region", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["backup_connected_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["cluster size"], "anchor": "schema-cluster_size", "description": "Number of nodes in the registration's site cluster. Use 1 for a single-node site and the complete node count for an HA site.", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["connected region"], "anchor": "schema-connected_region", "description": "connected region", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connected_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "name", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "namespace", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["preferred active re"], "anchor": "schema-preferred_active_re", "description": "preferred active re", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["preferred_active_re"], "syntax": "attribute", "type": "string"}, {"aliases": ["state"], "anchor": "schema-state", "description": "state", "document_id": "xcsh-docs:resources:registration_approval:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["state"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_registration_approval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_registration_approval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/)
- Property reference

## Direct properties

<a id="schema-backup_connected_region"></a>

### backup_connected_region property

Type: `"string"`. Optional.

<a id="schema-cluster_size"></a>

### cluster_size property

Type: `"number"`. Required.

Number of nodes in the registration's site cluster. Use 1 for a single-node site and the complete
node count for an HA site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.OneOf(1, 3),
}
```

<a id="schema-connected_region"></a>

### connected_region property

Type: `"string"`. Optional.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

<a id="schema-preferred_active_re"></a>

### preferred_active_re property

Type: `"string"`. Optional.

<a id="schema-state"></a>

### state property

Type: `"string"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("APPROVED")
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `backup_connected_region` | [backup_connected_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-backup_connected_region) |
| `cluster_size` | [cluster_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-cluster_size) |
| `connected_region` | [connected_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-connected_region) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-namespace) |
| `preferred_active_re` | [preferred_active_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-preferred_active_re) |
| `state` | [state](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/#schema-state) |
