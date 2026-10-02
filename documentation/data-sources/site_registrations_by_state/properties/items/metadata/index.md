---
page_title: "items.metadata"
subcategory: ""
description: "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object."
xcsh_docs: {"aliases": ["items metadata"], "body_bytes": 3096, "body_sha256": "sha256:bd3852d6099e9a4e04eb2f00402df8a1a782bb572dc138d6255f24f96269b579", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata:labels"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "path": "documentation/data-sources/site_registrations_by_state/properties/items/metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0002032302103311-0221203213023231-3122021003312130-1331101011003121-2323222002112200-3002012331110331-1321330031023123-2013130031230103", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "metadata"], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-items--metadata--annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description spec"], "anchor": "schema-items--metadata--description_spec", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable spec"], "anchor": "schema-items--metadata--disable_spec", "description": "Value of true will administratively disable the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "disable_spec"], "syntax": "attribute", "type": "bool"}, {"aliases": ["labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "metadata", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-items--metadata--name", "description": "Name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-items--metadata--namespace", "description": "Defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be ''.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "metadata", "namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/metadata/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.metadata

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
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

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/metadata/labels/): complete subsection reference.

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

- [items.metadata.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/metadata/labels/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
