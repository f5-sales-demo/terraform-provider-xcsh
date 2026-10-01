---
page_title: "cache_rules"
subcategory: ""
description: "cache_rules for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3469, "body_sha256": "sha256:367305d6abe9fd2612b4373cf2888201f2a5a8004c260900182729098353e97e", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "parent_id": "xcsh-docs:resources:cdn_cache_rule:reference", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/index.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cache_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- cache_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Cache Rule. This defines a CDN Cache Rule.

Upstream description:

This defines a CDN Cache Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule_expression_list",
    "rule_name"),
  validators.ConflictingObjectAttributes("cache_bypass",
    "eligible_for_cache")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_bypass\",\"eligible_for_cache\"]"
}
```

Terraform syntax:

```terraform
cache_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/cache_bypass/): complete subsection reference.

- [eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/): complete subsection reference.

- [rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/): complete subsection reference.

<a id="schema-cache_rules--rule_name"></a>

### rule_name property

Type: `"string"`. Optional.

Rule Name. Name of the Cache Rule.

Upstream description:

Name of the Cache Rule.

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

- [cache_rules.cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/cache_bypass/)
- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
