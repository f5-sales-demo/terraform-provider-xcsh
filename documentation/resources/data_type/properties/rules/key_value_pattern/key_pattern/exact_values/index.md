---
page_title: "rules.key_value_pattern.key_pattern.exact_values"
subcategory: ""
description: "List of exact values to match."
xcsh_docs: {"aliases": ["rules key value pattern key pattern exact values"], "body_bytes": 2834, "body_sha256": "sha256:40dc76f929c0ff07b7d9d35f179a200503bcc07b22ee8c80d3eaba1c50a5b384", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "parent_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "path": "documentation/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0302333232221013-0011212100001221-0331122032202233-3103232223100202-0033213022221110-1010002333011320-2110330202033321-1022320220100233", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "schema-rules--key_value_pattern--key_pattern--exact_values--exact_values", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern.exact_values:RequiredObjectAttributes:exact_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "key_pattern", "exact_values"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern key pattern exact values exact values"], "anchor": "schema-rules--key_value_pattern--key_pattern--exact_values--exact_values", "description": "List of exact values to match.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "exact_values", "exact_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of exact values to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.key_pattern.exact_values

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- [rules.key_value_pattern.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/)
- rules.key_value_pattern.key_pattern.exact_values

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
```

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

Terraform syntax:

```terraform
exact_values {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--key_value_pattern--key_pattern--exact_values--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rules.key_value_pattern.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
