---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.path_match.operator"
subcategory: ""
description: "Operator"
xcsh_docs: {"aliases": ["cache rules rule expression list cache rule expression path match operator"], "body_bytes": 5244, "body_sha256": "sha256:05aaaee7c51c45a0e16c1c60f5730b149e6ff81159fa1b1ab42b6b87cf8e3113", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1010230310032130-2010303023131230-0130113213231301-0231333131032033-2230300113001323-0123300113030322-3123010001230231-3332320100131020", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator"], "schema_version": 1, "sections": [{"aliases": ["contains"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--contains", "description": "Exclusive with The path must include the specified value as a substring, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "contains"], "syntax": "attribute", "type": "string"}, {"aliases": ["does not contain"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_contain", "description": "Exclusive with The path must not include the specified value as a substring, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "does_not_contain"], "syntax": "attribute", "type": "string"}, {"aliases": ["does not end with"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_end_with", "description": "Exclusive with The path must not end with the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "does_not_end_with"], "syntax": "attribute", "type": "string"}, {"aliases": ["does not equal"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_equal", "description": "Exclusive with The path must not match the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "does_not_equal"], "syntax": "attribute", "type": "string"}, {"aliases": ["does not start with"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_start_with", "description": "Exclusive with The path must not begin with the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "does_not_start_with"], "syntax": "attribute", "type": "string"}, {"aliases": ["endswith"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--endswith", "description": "Exclusive with The path must end with the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "endswith"], "syntax": "attribute", "type": "string"}, {"aliases": ["equals"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--equals", "description": "Exclusive with The path must exactly match the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "equals"], "syntax": "attribute", "type": "string"}, {"aliases": ["match regex"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--match_regex", "description": "Exclusive with The path must match the specified regular expression pattern in PCRE format.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "match_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["startswith"], "anchor": "schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--startswith", "description": "Exclusive with The path must begin with the specified value, up to the filename.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator", "startswith"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Operator", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/)
- cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

<a id="section"></a>

Type: `"single"`. Computed.

Operator

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

## Direct properties

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--contains"></a>

### contains property

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The path must include the specified value as a substring, up to the
filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_contain"></a>

### does_not_contain property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not include the specified value as a substring, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_end_with"></a>

### does_not_end_with property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not end with the specified value, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_equal"></a>

### does_not_equal property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not match the specified value, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_start_with"></a>

### does_not_start_with property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The path must not begin with the specified value, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--endswith"></a>

### endswith property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The path must end with the specified value, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--equals"></a>

### equals property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The path must exactly match the specified value, up to the filename.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--match_regex"></a>

### match_regex property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The path must match the specified regular expression pattern in PCRE format.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--startswith"></a>

### startswith property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The path must begin with the specified value, up to the filename.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.path_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
