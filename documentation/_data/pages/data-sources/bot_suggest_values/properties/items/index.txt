---
page_title: "items"
subcategory: ""
description: "items for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 1757, "body_sha256": "sha256:a0ccb6849152f99888690f8e709a29ceb866e99ba499bf41fe569d09ccb9363f", "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:properties:items:ref_value"], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "path": "documentation/data-sources/bot_suggest_values/properties/items/index.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/items/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items

Breadcrumbs:

- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Suggested Items. List of suggested items.

## Direct properties

<a id="schema-items--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Optional description for the suggested value.

- [ref_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/): complete subsection reference.

<a id="schema-items--str_value"></a>

### str_value property

Type: `"string"`. Computed.

String. Exclusive with \[ref\_value\]

<a id="schema-items--title"></a>

### title property

Type: `"string"`. Computed.

Optional title to be displayed instead of the actual value. Used when pure value doesn't contain all
the needed information to be displayed, or when display titles should be customized.

<a id="schema-items--value"></a>

### value property

Type: `"string"`. Computed.

Suggested value for the field. Should use value\_choice.str\_value instead.

## Next pages

- [items.ref_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/items/ref_value/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/)
- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
