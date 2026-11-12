---
page_title: "rules"
subcategory: ""
description: "Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 2036, "body_sha256": "sha256:0476ba1b1258b36e82bbef1fcaea27d7935f11dd66f966b969c324e769b9f3f7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_pattern", "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "xcsh-docs:data-sources:data_type:properties:rules:value_pattern"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules", "parent_id": "xcsh-docs:data-sources:data_type:reference", "path": "documentation/data-sources/data_type/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules key pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_pattern"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key value pattern"], "anchor": "section", "description": "Search for specific key & value patterns in the specified sections.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules value pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:value_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "value_pattern"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["data_typeCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- rules

<a id="section"></a>

Type: `"list"`. Computed.

Configure key/value or regex match rules to enable the platform to detect this custom data type in
the API request or response.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_pattern/): complete subsection reference.

- [key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/): complete subsection reference.

- [value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/value_pattern/): complete subsection reference.
