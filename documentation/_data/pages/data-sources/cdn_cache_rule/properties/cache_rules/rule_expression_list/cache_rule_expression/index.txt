---
page_title: "cache_rules.rule_expression_list.cache_rule_expression"
subcategory: ""
description: "The Cache Rule Expression Terms that are ANDed."
xcsh_docs: {"aliases": ["cache rules rule expression list cache rule expression"], "body_bytes": 3981, "body_sha256": "sha256:5ed0df265ea53c8630de22f820171a64aa72d941c1d6197925b4212d29b7b32d", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cookie_matcher", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression"], "schema_version": 1, "sections": [{"aliases": ["cache rules rule expression list cache rule expression cache headers"], "anchor": "section", "description": "Configure cache rule headers to match the criteria.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "cache_headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules rule expression list cache rule expression cookie matcher"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cookie_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "cookie_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules rule expression list cache rule expression path match"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules rule expression list cache rule expression query parameters"], "anchor": "section", "description": "List of (key, value) query parameters.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:query_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "query_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "The Cache Rule Expression Terms that are ANDed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

- [cache_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/): complete subsection reference.

- [cookie_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/): complete subsection reference.

- [path_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/): complete subsection reference.

- [query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/): complete subsection reference.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
