---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 30354, "body_sha256": "sha256:dea9c2cb8773db901fa09700dc4b105a4665fae53fe650a8b9dbed96aab28399", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:reference", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "xcsh-docs:resources:cdn_cache_rule:properties:timeouts"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:reference", "parent_id": "xcsh-docs:resources:cdn_cache_rule:fundamentals", "path": "docs/guides/resources--cdn_cache_rule--reference.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the CDN Cache Rule. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace where the CDN Cache Rule is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--cdn_cache_rule--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_cache_rule--reference.md#schema-annotations) |
| `cache_rules` | [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md#section) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](resources--cdn_cache_rule--properties--cache_rules--cache_bypass.md#section) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_request_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--ignore_response_cookie) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#section) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md#schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--name) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--cookie_matcher--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--path_match--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--path_match--operator--startswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--key) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#section) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--contains) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_contain) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_end_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_equal) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--does_not_start_with) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--endswith) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--equals) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--match_regex) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator.md#schema-cache_rules--rule_expression_list--cache_rule_expression--query_parameters--operator--startswith) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md#schema-cache_rules--rule_expression_list--expression_name) |
| `cache_rules.rule_name` | [cache_rules.rule_name](resources--cdn_cache_rule--properties--cache_rules.md#schema-cache_rules--rule_name) |
| `description` | [description](resources--cdn_cache_rule--reference.md#schema-description) |
| `disable` | [disable](resources--cdn_cache_rule--reference.md#schema-disable) |
| `id` | [id](resources--cdn_cache_rule--reference.md#schema-id) |
| `labels` | [labels](resources--cdn_cache_rule--reference.md#schema-labels) |
| `name` | [name](resources--cdn_cache_rule--reference.md#schema-name) |
| `namespace` | [namespace](resources--cdn_cache_rule--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--cdn_cache_rule--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--cdn_cache_rule--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_cache_rule--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--cdn_cache_rule--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--cdn_cache_rule--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [timeouts](resources--cdn_cache_rule--properties--timeouts.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
