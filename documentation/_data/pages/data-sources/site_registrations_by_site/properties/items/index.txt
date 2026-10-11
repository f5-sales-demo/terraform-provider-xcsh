---
page_title: "items"
subcategory: ""
description: "Items represents the collection in response."
xcsh_docs: {"aliases": ["items"], "body_bytes": 3056, "body_sha256": "sha256:b9f391d85be30709bb831e1806e59d1b2bb22bf2065c2f8244303610ac4b4076", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:annotations", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:labels", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:owner_view", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:reference", "path": "documentation/data-sources/site_registrations_by_site/properties/items/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0203000012213013-2023113310133033-2023302133132132-2122132031003002-3223313321231231-1110001122110330-0313203032121021-0313221120120330", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items"], "schema_version": 1, "sections": [{"aliases": ["items annotations"], "anchor": "section", "description": "The set of annotations present on this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:annotations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "annotations"], "syntax": "attribute", "type": "object"}, {"aliases": ["items description spec"], "anchor": "schema-items--description_spec", "description": "The description set for this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items disabled"], "anchor": "schema-items--disabled", "description": "Value of true indicates registration is administratively disabled.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "disabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items get spec"], "anchor": "section", "description": "GET Registration. GET registration specification.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["items labels"], "anchor": "section", "description": "The set of labels present on this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items metadata"], "anchor": "section", "description": "ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items name"], "anchor": "schema-items--name", "description": "Name. The name of this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items namespace"], "anchor": "schema-items--namespace", "description": "Namespace. The namespace this item belongs to.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object"], "anchor": "section", "description": "Registration object stores node registration and information regarding the node.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object"], "syntax": "attribute", "type": "object"}, {"aliases": ["items owner view"], "anchor": "section", "description": "ViewRefType represents a reference to a view.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "owner_view"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata"], "anchor": "section", "description": "SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items tenant"], "anchor": "schema-items--tenant", "description": "Tenant. The tenant this item belongs to.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items uid"], "anchor": "schema-items--uid", "description": "UID. The unique uid of this registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Items represents the collection in response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Items represents the collection in response.

## Direct properties

- [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/annotations/): complete subsection reference.

<a id="schema-items--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

The description set for this registration.

<a id="schema-items--disabled"></a>

### disabled property

Type: `"bool"`. Computed.

Value of true indicates registration is administratively disabled.

- [get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/labels/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/metadata/): complete subsection reference.

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

- [object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/): complete subsection reference.

- [owner_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/owner_view/): complete subsection reference.

- [system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/system_metadata/): complete subsection reference.

<a id="schema-items--tenant"></a>

### tenant property

Type: `"string"`. Computed.

Tenant. The tenant this item belongs to.

<a id="schema-items--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. The unique uid of this registration.
