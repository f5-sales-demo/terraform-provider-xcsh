---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": ["bot suggest values"], "body_bytes": 4595, "body_sha256": "sha256:96b40f2d531f19788f761a6c18858de90861f7f48b5e7f16533c20a5072fdd19", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:properties:items", "xcsh-docs:data-sources:bot_suggest_values:properties:request_body"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:reference", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:fundamentals", "path": "documentation/data-sources/bot_suggest_values/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0133311012330202-3003220301112011-3121011301000332-0123101221021012-1323031211013031-1311312133320211-3120303313001021-1311000323111302", "registry_path": "docs/guides/data-sources--bot_suggest_values--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["field path"], "anchor": "schema-field_path", "description": "JSON path of the field for which the suggested values are being requested.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["field_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["items"], "anchor": "section", "description": "Suggested Items. List of suggested items.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items"], "syntax": "attribute", "type": "object"}, {"aliases": ["match value"], "anchor": "schema-match_value", "description": "Substring that must be present in either the value or description of each SuggestedItem in the response.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["match_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace in which the suggestions are scoped.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["request body"], "anchor": "section", "description": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["request_body"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Property reference for xcsh_bot_suggest_values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": [], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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

## Next pages

- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/)
- [request_body](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/request_body/)
- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
