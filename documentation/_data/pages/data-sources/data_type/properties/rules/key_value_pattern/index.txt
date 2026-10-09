---
page_title: "rules.key_value_pattern"
subcategory: ""
description: "Search for specific key & value patterns in the specified sections."
xcsh_docs: {"aliases": ["rules key value pattern"], "body_bytes": 1168, "body_sha256": "sha256:48e616a4250d4a362225372fd9ffd45afc71df1bc10e376644a0f3a34c97d1b8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules", "path": "documentation/data-sources/data_type/properties/rules/key_value_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern key pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key value pattern value pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Search for specific key & value patterns in the specified sections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["data_typeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/)
- rules.key_value_pattern

<a id="section"></a>

Type: `"single"`. Computed.

Search for specific key &amp; value patterns in the specified sections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/): complete subsection reference.

- [value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/value_pattern/): complete subsection reference.
