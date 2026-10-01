---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 8333, "body_sha256": "sha256:8a65d3e7cfd9c0c88c2499cbc3387dcb3886e1a15a7cb7c16be1f8de26c4f47d", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers:operator", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/index.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "cache_headers", "operator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/)
- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--contains"></a>

### contains property

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The header value must include the specified value as a substring.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_contain"></a>

### does_not_contain property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not include the specified value as a substring.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_end_with"></a>

### does_not_end_with property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not end with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_equal"></a>

### does_not_equal property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not match the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_start_with"></a>

### does_not_start_with property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The header value must not begin with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--endswith"></a>

### endswith property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The header value must end with the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--equals"></a>

### equals property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The header value must exactly match the specified value.

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--match_regex"></a>

### match_regex property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The header value must match the specified regular expression pattern.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--startswith"></a>

### startswith property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The header value must begin with the specified value.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
