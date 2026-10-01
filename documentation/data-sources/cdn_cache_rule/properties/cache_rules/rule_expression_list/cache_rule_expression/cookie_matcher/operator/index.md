---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 5216, "body_sha256": "sha256:5d702c24d721108142fa0bd35589bec2629576e02132dc5b91913129b7667f70", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cookie_matcher:operator", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cookie_matcher", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/index.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "cookie_matcher", "operator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

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

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--contains"></a>

### contains property

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The cookie value must include the specified value as a substring.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_contain"></a>

### does_not_contain property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not include the specified value as a substring.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_end_with"></a>

### does_not_end_with property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not end with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_equal"></a>

### does_not_equal property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not match the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_start_with"></a>

### does_not_start_with property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The cookie value must not begin with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--endswith"></a>

### endswith property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The cookie value must end with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--equals"></a>

### equals property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The cookie value must exactly match the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--match_regex"></a>

### match_regex property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The cookie value must match the specified regular expression pattern in PCRE
format.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--startswith"></a>

### startswith property

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The cookie value must begin with the specified value.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
