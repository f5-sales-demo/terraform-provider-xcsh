---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 32429, "body_sha256": "sha256:9238b169876682a20e890b1a2def5d6db03a3099dbd234eb1831b24895d9ae28", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules"], "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:reference", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:fundamentals", "path": "documentation/data-sources/cdn_cache_rule/properties/index.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the CDNCacheRule.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the CDNCacheRule.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the CDNCacheRule exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-annotations) |
| `cache_rules` | [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/#section) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/cache_bypass/#section) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--ignore_response_cookie) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cookie_matcher/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/path_match/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--key) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/query_parameters/operator/#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--startswith) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/#schema-cache_rules--rule_expression_list--expression_name) |
| `cache_rules.rule_name` | [cache_rules.rule_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/#schema-cache_rules--rule_name) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/#schema-namespace) |

## Next pages

- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
