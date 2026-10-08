---
page_title: "items"
subcategory: ""
description: "Items represents the collection in response."
xcsh_docs: {"aliases": ["items"], "body_bytes": 3066, "body_sha256": "sha256:adbc9e2352a6f5bc33c58075701d853f7076b50a819e0568f86180c1d256f93f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:annotations", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:labels", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:owner_view", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:reference", "path": "documentation/data-sources/site_registrations_by_state/properties/items/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items"], "schema_version": 1, "sections": [{"aliases": ["items annotations"], "anchor": "section", "description": "The set of annotations present on this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:annotations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "annotations"], "syntax": "attribute", "type": "object"}, {"aliases": ["items description spec"], "anchor": "schema-items--description_spec", "description": "The description set for this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items disabled"], "anchor": "schema-items--disabled", "description": "Value of true indicates registration is administratively disabled.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "disabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items get spec"], "anchor": "section", "description": "GET Registration. GET registration specification.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["items labels"], "anchor": "section", "description": "The set of labels present on this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items metadata"], "anchor": "section", "description": "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items name"], "anchor": "schema-items--name", "description": "Name. The name of this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items namespace"], "anchor": "schema-items--namespace", "description": "Namespace. The namespace this item belongs to.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object"], "anchor": "section", "description": "Registration object stores node registration and information regarding the node.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object"], "syntax": "attribute", "type": "object"}, {"aliases": ["items owner view"], "anchor": "section", "description": "ViewRefType represents a reference to a view.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "owner_view"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata"], "anchor": "section", "description": "SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items tenant"], "anchor": "schema-items--tenant", "description": "Tenant. The tenant this item belongs to.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items uid"], "anchor": "schema-items--uid", "description": "UID. The unique uid of this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Items represents the collection in response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Items represents the collection in response.

## Direct properties

- [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/annotations/): complete subsection reference.

<a id="schema-items--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

The description set for this registration.

<a id="schema-items--disabled"></a>

### disabled property

Type: `"bool"`. Computed.

Value of true indicates registration is administratively disabled.

- [get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/labels/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/metadata/): complete subsection reference.

<a id="schema-items--name"></a>

### name property

Type: `"string"`. Computed.

Name. The name of this registration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--namespace"></a>

### namespace property

Type: `"string"`. Computed.

Namespace. The namespace this item belongs to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

- [object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/): complete subsection reference.

- [owner_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/owner_view/): complete subsection reference.

- [system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/): complete subsection reference.

<a id="schema-items--tenant"></a>

### tenant property

Type: `"string"`. Computed.

Tenant. The tenant this item belongs to.

<a id="schema-items--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. The unique uid of this registration.
