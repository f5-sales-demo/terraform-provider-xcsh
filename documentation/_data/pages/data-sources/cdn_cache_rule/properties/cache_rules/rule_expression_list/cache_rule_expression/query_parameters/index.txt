---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.query_parameters"
subcategory: ""
description: "List of (key, value) query parameters."
xcsh_docs: {"aliases": ["cache rules rule expression list cache rule expression query parameters"], "body_bytes": 3217, "body_sha256": "sha256:8def6e087747a18790274b5afc688f0da0921f484258fb7bf068eff0f6f9c1ff", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters:operator"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "query_parameters"], "schema_version": 1, "sections": [{"aliases": ["cache rules rule expression list cache rule expression query parameters key"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--key", "description": "The name of the query parameter to match.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "query_parameters", "key"], "syntax": "attribute", "type": "string"}, {"aliases": ["cache rules rule expression list cache rule expression query parameters operator"], "anchor": "section", "description": "Operator", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters:operator", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "query_parameters", "operator"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of (key, value) query parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.query_parameters

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="section"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [operator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/): complete subsection reference.
