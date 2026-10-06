---
page_title: "items.object.system_metadata.initializers.pending"
subcategory: ""
description: "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients."
xcsh_docs: {"aliases": ["items object system metadata initializers pending"], "body_bytes": 1945, "body_sha256": "sha256:6c41100af030f3fdc2f82cc0820dba399ffff3df628865e3d5b6638616c611d6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:pending", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/pending/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3301123013322303-3333000210020013-1203220200121020-1110012131110302-3200230130333001-1203202321003023-1002302230113003-3023311232320231", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers", "pending"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata initializers pending name"], "anchor": "schema-items--object--system_metadata--initializers--pending--name", "description": "Name of the service that is responsible for initializing this object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:pending", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "pending", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/pending/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.initializers.pending

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/)
- items.object.system_metadata.initializers.pending

<a id="section"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

## Direct properties

<a id="schema-items--object--system_metadata--initializers--pending--name"></a>

### name property

Type: `"string"`. Computed.

Name of the service that is responsible for initializing this object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```
