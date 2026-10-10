---
page_title: "items.metadata"
subcategory: ""
description: "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object."
xcsh_docs: {"aliases": ["items metadata"], "body_bytes": 2704, "body_sha256": "sha256:a5afe21121a17b96736dae09c04a2b33512bd4a1f5bca6418b51d5c5d1458e26", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:metadata:labels"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "documentation/data-sources/site_registrations/properties/items/metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0102013231012023-1332123221311130-3331232330103022-2131320202330100-1102112023302003-2130211120333213-3012012312011111-3320012332022122", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "metadata"], "schema_version": 1, "sections": [{"aliases": ["items metadata annotations"], "anchor": "schema-items--metadata--annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["items metadata description spec"], "anchor": "schema-items--metadata--description_spec", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items metadata disable spec"], "anchor": "schema-items--metadata--disable_spec", "description": "Value of true will administratively disable the object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "disable_spec"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items metadata labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "metadata", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items metadata name"], "anchor": "schema-items--metadata--name", "description": "Name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items metadata namespace"], "anchor": "schema-items--metadata--namespace", "description": "Defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be ''.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/metadata/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.metadata

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- items.metadata

<a id="section"></a>

Type: `"single"`. Computed.

ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.

## Direct properties

<a id="schema-items--metadata--annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

<a id="schema-items--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Human readable description for the object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1200),
}
```

<a id="schema-items--metadata--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Value of true will administratively disable the object.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/metadata/labels/): complete subsection reference.

<a id="schema-items--metadata--name"></a>

### name property

Type: `"string"`. Computed.

Name of configuration object. It has to be unique within the namespace. It can only be specified
during create API and cannot be changed during replace API.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--metadata--namespace"></a>

### namespace property

Type: `"string"`. Computed.

Defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ''.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```
