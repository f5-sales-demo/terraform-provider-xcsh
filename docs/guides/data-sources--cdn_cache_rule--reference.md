---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 29047, "body_sha256": "sha256:ddd2b3860927c9e360262e86c58c461183c33dbafe0fed20bb80c8c451e89170", "canonical_id": "xcsh-docs:data-sources:cdn_cache_rule:reference", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules"], "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:reference", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:fundamentals", "path": "docs/guides/data-sources--cdn_cache_rule--reference.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
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

- [cache_rules](data-sources--cdn_cache_rule--properties--cache_rules.md): complete subsection reference.

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
| `annotations` | [annotations](data-sources--cdn_cache_rule--reference.md#schema-annotations) |
| `cache_rules` | [cache_rules](data-sources--cdn_cache_rule--properties--cache_rules.md#section) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](data-sources--cdn_cache_rule--properties--cache_rules--cache_bypass.md#section) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--ignore_response_cookie) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](data-sources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--key) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--startswith) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](data-sources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md#schema-cache_rules--rule_expression_list--expression_name) |
| `cache_rules.rule_name` | [cache_rules.rule_name](data-sources--cdn_cache_rule--properties--cache_rules.md#schema-cache_rules--rule_name) |
| `description` | [description](data-sources--cdn_cache_rule--reference.md#schema-description) |
| `id` | [id](data-sources--cdn_cache_rule--reference.md#schema-id) |
| `labels` | [labels](data-sources--cdn_cache_rule--reference.md#schema-labels) |
| `name` | [name](data-sources--cdn_cache_rule--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cdn_cache_rule--reference.md#schema-namespace) |

## Next pages

- [cache_rules](data-sources--cdn_cache_rule--properties--cache_rules.md)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md)
