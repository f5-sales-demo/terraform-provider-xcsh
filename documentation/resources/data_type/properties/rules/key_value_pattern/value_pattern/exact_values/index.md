---
page_title: "rules.key_value_pattern.value_pattern.exact_values"
subcategory: ""
description: "List of exact values to match."
xcsh_docs: {"aliases": ["rules key value pattern value pattern exact values"], "body_bytes": 2848, "body_sha256": "sha256:54f2021dc7bc17f60161fd48e51bd8b84acce9d25847513b7bde4ca86670187d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "parent_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "path": "documentation/resources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2102300102311332-0011000011222313-2313223022000202-3003210011003022-1020323323302331-2101230130212220-1133231231131021-1301002320310200", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "schema-rules--key_value_pattern--value_pattern--exact_values--exact_values", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern.exact_values:RequiredObjectAttributes:exact_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "value_pattern", "exact_values"], "schema_version": 1, "sections": [{"aliases": ["exact values"], "anchor": "schema-rules--key_value_pattern--value_pattern--exact_values--exact_values", "description": "List of exact values to match.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern", "exact_values", "exact_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of exact values to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.value_pattern.exact_values

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- [rules.key_value_pattern.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/)
- rules.key_value_pattern.value_pattern.exact_values

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

<a id="schema-rules--key_value_pattern--value_pattern--exact_values--exact_values"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [rules.key_value_pattern.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
