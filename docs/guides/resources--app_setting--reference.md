---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 21552, "body_sha256": "sha256:677e5afe2b48e9966c22f70ea1ee31690688d75cd5e51a1af5112f4d640bef1d", "canonical_id": "xcsh-docs:resources:app_setting:reference", "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings", "xcsh-docs:resources:app_setting:properties:timeouts"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:reference", "parent_id": "xcsh-docs:resources:app_setting:fundamentals", "path": "docs/guides/resources--app_setting--reference.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
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

- [app_type_settings](resources--app_setting--properties--app_type_settings.md): complete subsection reference.

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

Name of the App Setting. Must be unique within the namespace.

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

Namespace where the App Setting is created.

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

- [timeouts](resources--app_setting--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_setting--reference.md#schema-annotations) |
| `app_type_settings` | [app_type_settings](resources--app_setting--properties--app_type_settings.md#section) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](resources--app_setting--properties--app_type_settings--app_type_ref.md#section) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](resources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--kind) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](resources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--name) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](resources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--namespace) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](resources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--tenant) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](resources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--uid) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](resources--app_setting--properties--app_type_settings--business_logic_markup_setting.md#section) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](resources--app_setting--properties--app_type_settings--business_logic_markup_setting--disable_spec.md#section) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](resources--app_setting--properties--app_type_settings--business_logic_markup_setting--enable.md#section) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting.md#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_learning.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--bola_detection_automatic.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--cooling_off_period) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_bola_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_bot_defense_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_failed_login_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_forbidden_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_ip_reputation.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_non_existent_url_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_rate_limit.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_waf_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_bot_defense_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_ip_reputation.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--high.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--low.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--medium.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_rate_limit.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_waf_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_learning.md#section) |
| `description` | [description](resources--app_setting--reference.md#schema-description) |
| `disable` | [disable](resources--app_setting--reference.md#schema-disable) |
| `id` | [id](resources--app_setting--reference.md#schema-id) |
| `labels` | [labels](resources--app_setting--reference.md#schema-labels) |
| `name` | [name](resources--app_setting--reference.md#schema-name) |
| `namespace` | [namespace](resources--app_setting--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--app_setting--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--app_setting--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--app_setting--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--app_setting--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--app_setting--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- [timeouts](resources--app_setting--properties--timeouts.md)
- [xcsh_app_setting](../resources/app_setting.md)
