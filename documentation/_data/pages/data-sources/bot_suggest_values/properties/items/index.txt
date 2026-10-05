---
page_title: "items"
subcategory: ""
description: "Suggested Items. List of suggested items."
xcsh_docs: {"aliases": ["items"], "body_bytes": 1757, "body_sha256": "sha256:a0ccb6849152f99888690f8e709a29ceb866e99ba499bf41fe569d09ccb9363f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:properties:items:ref_value"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "path": "documentation/data-sources/bot_suggest_values/properties/items/index.md", "product": "distributed-cloud", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1111021110211130-3122023222133231-2011033032111212-3233113022331010-3001220323221102-2221321210220331-1220320022123212-1333000002221221", "registry_path": "docs/guides/data-sources--bot_suggest_values--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items"], "schema_version": 1, "sections": [{"aliases": ["items description spec"], "anchor": "schema-items--description_spec", "description": "Optional description for the suggested value.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["items ref value"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items:ref_value", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "ref_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["items str value"], "anchor": "schema-items--str_value", "description": "String. Exclusive with", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "str_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["items title"], "anchor": "schema-items--title", "description": "Optional title to be displayed instead of the actual value. Used when pure value doesn't contain all the needed information to be displayed, or when display titles should be customized.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "title"], "syntax": "attribute", "type": "string"}, {"aliases": ["items value"], "anchor": "schema-items--value", "description": "Suggested value for the field. Should use value_choice.str_value instead.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:items", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/items/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Suggested Items. List of suggested items.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
