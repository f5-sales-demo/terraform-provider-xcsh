---
page_title: "cache_rules.rule_expression_list"
subcategory: ""
description: "cache_rules.rule_expression_list for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3527, "body_sha256": "sha256:a2bf0f90d20b7de9b0f980f223e34b670d4a2d52523792b4e6aedfd30c5e8d0e", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- cache_rules.rule_expression_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Upstream description:

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs..

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cache_rule_expression",
    "expression_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rule_expression_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_rule_expression](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md): complete subsection reference.

<a id="schema-cache_rules--rule_expression_list--expression_name"></a>

### expression_name property

Type: `"string"`. Optional.

Name of the Expressions items that are ANDed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
