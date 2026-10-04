---
page_title: "items.metadata"
subcategory: ""
description: "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object."
xcsh_docs: {"aliases": ["items metadata"], "body_bytes": 3087, "body_sha256": "sha256:949f260aba7c7f07894b901af714fc8778bcd95053889a834702b18268574fdd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata:labels"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "path": "documentation/data-sources/site_registrations_by_site/properties/items/metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3010130210021221-2130212223231102-2111103331300023-1321312302002131-2011120230333201-1100132311312032-3110103122201302-2132013132320130", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "metadata"], "schema_version": 1, "sections": [{"aliases": ["items metadata annotations"], "anchor": "schema-items--metadata--annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["items metadata description spec"], "anchor": "schema-items--metadata--description_spec", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items metadata disable spec"], "anchor": "schema-items--metadata--disable_spec", "description": "Value of true will administratively disable the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "disable_spec"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items metadata labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "metadata", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items metadata name"], "anchor": "schema-items--metadata--name", "description": "Name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items metadata namespace"], "anchor": "schema-items--metadata--namespace", "description": "Defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be ''.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/metadata/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.metadata

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
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
Validators: []validator.String{
  stringvalidator.LengthAtMost(1200),
}
```

<a id="schema-items--metadata--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Value of true will administratively disable the object.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/metadata/labels/): complete subsection reference.

<a id="schema-items--metadata--name"></a>

### name property

Type: `"string"`. Computed.

Name of configuration object. It has to be unique within the namespace. It can only be specified
during create API and cannot be changed during replace API.

Provider validators and defaults (from schema source):

```go
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
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

## Next pages

- [items.metadata.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/metadata/labels/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
