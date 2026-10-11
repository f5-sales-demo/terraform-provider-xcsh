---
page_title: "items.system_metadata.owner_view"
subcategory: ""
description: "ViewRefType represents a reference to a view."
xcsh_docs: {"aliases": ["items system metadata owner view"], "body_bytes": 2138, "body_sha256": "sha256:5d268d51c670d5d53ac9197367ca97bf0bc5033ec339493f204b1717175ef65a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "path": "documentation/data-sources/site_registrations/properties/items/system_metadata/owner_view/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3111020003212130-3110210330100332-0230123233010131-2202330330132003-1330000001131331-3303211232103033-0103311301132200-2313033033323231", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "system_metadata", "owner_view"], "schema_version": 1, "sections": [{"aliases": ["items system metadata owner view kind"], "anchor": "schema-items--system_metadata--owner_view--kind", "description": "Kind. Kind of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata owner view name"], "anchor": "schema-items--system_metadata--owner_view--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata owner view namespace"], "anchor": "schema-items--system_metadata--owner_view--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata owner view uid"], "anchor": "schema-items--system_metadata--owner_view--uid", "description": "UID. UID of the view object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "owner_view", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/system_metadata/owner_view/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ViewRefType represents a reference to a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata.owner_view

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/)
- items.system_metadata.owner_view

<a id="section"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

## Direct properties

<a id="schema-items--system_metadata--owner_view--kind"></a>

### kind property

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="schema-items--system_metadata--owner_view--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--system_metadata--owner_view--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--system_metadata--owner_view--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. UID of the view object.
