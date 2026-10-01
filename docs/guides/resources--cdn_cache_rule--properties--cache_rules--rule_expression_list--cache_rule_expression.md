---
page_title: "cache_rules.rule_expression_list.cache_rule_expression"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3379, "body_sha256": "sha256:3579b30accb6960efd006bd5b7fc0002b35f34adec13f4d7d1f70c6c917e64aa", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cookie_matcher", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

The Cache Rule Expression Terms that are ANDed.

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
cache_rule_expression {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_headers](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md): complete subsection reference.

- [cookie_matcher](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md): complete subsection reference.

- [path_match](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md): complete subsection reference.

- [query_parameters](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md): complete subsection reference.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
