---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.query_parameters"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression.query_parameters for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3432, "body_sha256": "sha256:3ee00f9617ccaed91add6878f6a6969fca68719d90910011bb3f1a826e0bc982", "canonical_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters:operator"], "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "path": "docs/guides/data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "query_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression.query_parameters for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.query_parameters

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
- [Property reference](data-sources--cdn_cache_rule--reference.md)
- [cache_rules](data-sources--cdn_cache_rule--properties--cache_rules.md)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="section"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--key"></a>

### key property

Type: `"string"`. Computed.

The name of the query parameter to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md): complete subsection reference.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
