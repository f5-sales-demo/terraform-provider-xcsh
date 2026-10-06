---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": ["bot suggest values"], "body_bytes": 4271, "body_sha256": "sha256:88d075925cb40848dc1b3d8c4558abae602baab2730ea688e4e4b0fb43be7770", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:properties:items", "xcsh-docs:data-sources:bot_suggest_values:properties:request_body"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:reference", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:fundamentals", "path": "documentation/data-sources/bot_suggest_values/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0133311012330202-3003220301112011-3121011301000332-0123101221021012-1323031211013031-1311312133320211-3120303313001021-1311000323111302", "registry_path": "docs/guides/data-sources--bot_suggest_values--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["field path"], "anchor": "schema-field_path", "description": "JSON path of the field for which the suggested values are being requested.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["field_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["items"], "anchor": "section", "description": "Suggested Items. List of suggested items.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items"], "syntax": "attribute", "type": "object"}, {"aliases": ["match value"], "anchor": "schema-match_value", "description": "Substring that must be present in either the value or description of each SuggestedItem in the response.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["match_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace in which the suggestions are scoped.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["request body"], "anchor": "section", "description": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["request_body"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_bot_suggest_values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
- Property reference

## Direct properties

<a id="schema-field_path"></a>

### field_path property

Type: `"string"`. Optional.

JSON path of the field for which the suggested values are being requested.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/): complete subsection reference.

<a id="schema-match_value"></a>

### match_value property

Type: `"string"`. Optional.

Substring that must be present in either the value or description of each SuggestedItem in the
response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace in which the suggestions are scoped.

- [request_body](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/request_body/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `field_path` | [field_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/#schema-field_path) |
| `items` | [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/#section) |
| `items.description_spec` | [items.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/#schema-items--description_spec) |
| `items.ref_value` | [items.ref_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/#section) |
| `items.ref_value.name` | [items.ref_value.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/#schema-items--ref_value--name) |
| `items.ref_value.namespace` | [items.ref_value.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/#schema-items--ref_value--namespace) |
| `items.ref_value.tenant` | [items.ref_value.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/#schema-items--ref_value--tenant) |
| `items.str_value` | [items.str_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/#schema-items--str_value) |
| `items.title` | [items.title](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/#schema-items--title) |
| `items.value` | [items.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/#schema-items--value) |
| `match_value` | [match_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/#schema-match_value) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/#schema-namespace) |
| `request_body` | [request_body](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/request_body/#section) |
| `request_body.type_url` | [request_body.type_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/request_body/#schema-request_body--type_url) |
| `request_body.value` | [request_body.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/request_body/#schema-request_body--value) |
