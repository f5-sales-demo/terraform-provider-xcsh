---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 4595, "body_sha256": "sha256:96b40f2d531f19788f761a6c18858de90861f7f48b5e7f16533c20a5072fdd19", "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:properties:items", "xcsh-docs:data-sources:bot_suggest_values:properties:request_body"], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:reference", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:fundamentals", "path": "documentation/data-sources/bot_suggest_values/properties/index.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
