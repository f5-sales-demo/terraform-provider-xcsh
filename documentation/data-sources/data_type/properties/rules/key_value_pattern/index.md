---
page_title: "rules.key_value_pattern"
subcategory: ""
description: "Search for specific key & value patterns in the specified sections."
xcsh_docs: {"aliases": ["rules key value pattern"], "body_bytes": 1168, "body_sha256": "sha256:48e616a4250d4a362225372fd9ffd45afc71df1bc10e376644a0f3a34c97d1b8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules", "path": "documentation/data-sources/data_type/properties/rules/key_value_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern key pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key value pattern value pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Search for specific key & value patterns in the specified sections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["data_typeCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
