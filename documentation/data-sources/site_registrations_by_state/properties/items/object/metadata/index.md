---
page_title: "items.object.metadata"
subcategory: ""
description: "ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["items object metadata"], "body_bytes": 3519, "body_sha256": "sha256:46e156fcaef053e8e2a27ca8fb0c95e51afa997e1db0eac34136e5372a0cecd7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3311000213200301-0000010010112023-3110111210100303-0221203320133311-1001000010211122-0232212110321002-1010201033112110-0200221123110123", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "metadata"], "schema_version": 1, "sections": [{"aliases": ["items object metadata annotations"], "anchor": "schema-items--object--metadata--annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["items object metadata description spec"], "anchor": "schema-items--object--metadata--description_spec", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": true, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object metadata disable spec"], "anchor": "schema-items--object--metadata--disable_spec", "description": "Value of true will administratively disable the object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "disable_spec"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items object metadata labels"], "anchor": "schema-items--object--metadata--labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["items object metadata name"], "anchor": "schema-items--object--metadata--name", "description": "Name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object metadata namespace"], "anchor": "schema-items--object--metadata--namespace", "description": "Defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be ''.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object metadata uid", "succeeded", "success", "successful"], "anchor": "schema-items--object--metadata--uid", "description": "Uid is the unique in time and space value for this object. Object create will fail if provided by the client and the value exists in the system. Typically generated by the server on successful creation of an object and is not allowed to change once populated.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "metadata", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/metadata/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.metadata

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- items.object.metadata

<a id="section"></a>

Type: `"single"`. Computed.

ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have.
The information in this type can be specified by user during create and replace APIs.

## Direct properties

<a id="schema-items--object--metadata--annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

<a id="schema-items--object--metadata--description_spec"></a>

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

<a id="schema-items--object--metadata--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Value of true will administratively disable the object.

<a id="schema-items--object--metadata--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

<a id="schema-items--object--metadata--name"></a>

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

<a id="schema-items--object--metadata--namespace"></a>

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

<a id="schema-items--object--metadata--uid"></a>

### uid property

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. Object create will fail if provided by
the client and the value exists in the system. Typically generated by the server on successful
creation of an object and is not allowed to change once populated.
