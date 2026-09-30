---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.path_match.operator"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression.path_match.operator for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 4744, "body_sha256": "sha256:4ef25da1935dd75fb9047b5eb72ecb7d32ec1e7c230262ada3f1ea374ff2fbd6", "canonical_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match:operator", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:path_match", "path": "docs/guides/data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "path_match", "operator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression.path_match.operator for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
- [Property reference](data-sources--cdn_cache_rule--reference.md)
- [cache_rules](data-sources--cdn_cache_rule--properties--cache_rules.md)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md)
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

- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
