---
page_title: "rules"
subcategory: ""
description: "Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 3113, "body_sha256": "sha256:b77e8868ac8bc514a180fedc35ecb58d800fc7eac64dc85dcb63bc85442fd229", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_type:properties:rules:key_pattern", "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "xcsh-docs:resources:data_type:properties:rules:value_pattern"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules", "parent_id": "xcsh-docs:resources:data_type:reference", "path": "documentation/resources/data_type/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_pattern,key_value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_pattern,value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_pattern,key_value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_value_pattern,value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_pattern,value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:key_value_pattern,value_pattern", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules key pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--key_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern:exact_values", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_pattern:exact_values", "type": "conflicts"}], "schema_path": ["rules", "key_pattern"], "syntax": "block", "type": "object"}, {"aliases": ["rules key value pattern"], "anchor": "section", "description": "Search for specific key & value patterns in the specified sections.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern"], "syntax": "block", "type": "object"}, {"aliases": ["rules value pattern"], "anchor": "section", "description": "Test", "document_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "type": "conflicts"}], "schema_path": ["rules", "value_pattern"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure key/value or regex match rules to enable the platform to detect this custom data type in
the API request or response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("key_pattern",
    "key_value_pattern"),
  validators.ConflictingListObjectAttributes("key_pattern",
    "value_pattern"),
  validators.ConflictingListObjectAttributes("key_value_pattern",
    "value_pattern")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/): complete subsection reference.

- [key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/): complete subsection reference.

- [value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/): complete subsection reference.

## Next pages

- [rules.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- [rules.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
