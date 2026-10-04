---
page_title: "rules.key_value_pattern"
subcategory: ""
description: "Search for specific key & value patterns in the specified sections."
xcsh_docs: {"aliases": ["rules key value pattern"], "body_bytes": 1732, "body_sha256": "sha256:567f9d93835ced35bdf48324c4b7033cf5fd4332b636e5963efb8961e94db13d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules", "path": "documentation/data-sources/data_type/properties/rules/key_value_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern key pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key value pattern value pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Search for specific key & value patterns in the specified sections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

## Next pages

- [rules.key_value_pattern.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/)
- [rules.key_value_pattern.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/value_pattern/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
