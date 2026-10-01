---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 20113, "body_sha256": "sha256:3620ba14bb13a5af538d9da9434cdc9234d03947c61056acc2e9673f5091383b", "canonical_id": "xcsh-docs:data-sources:app_setting:reference", "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings"], "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:reference", "parent_id": "xcsh-docs:data-sources:app_setting:fundamentals", "path": "docs/guides/data-sources--app_setting--reference.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md)
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

- [app_type_settings](data-sources--app_setting--properties--app_type_settings.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AppSetting.

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

Name of the AppSetting.

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

Namespace where the AppSetting exists.

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
| `annotations` | [annotations](data-sources--app_setting--reference.md#schema-annotations) |
| `app_type_settings` | [app_type_settings](data-sources--app_setting--properties--app_type_settings.md#section) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#section) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--kind) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--name) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--namespace) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--tenant) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](data-sources--app_setting--properties--app_type_settings--app_type_ref.md#schema-app_type_settings--app_type_ref--uid) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](data-sources--app_setting--properties--app_type_settings--business_logic_markup_setting.md#section) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](data-sources--app_setting--properties--app_type_settings--business_logic_markup_setting--disable_spec.md#section) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](data-sources--app_setting--properties--app_type_settings--business_logic_markup_setting--enable.md#section) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](data-sources--app_setting--properties--app_type_settings--timeseries_analyses_setting.md#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](data-sources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](data-sources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](data-sources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_learning.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--bola_detection_automatic.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--cooling_off_period) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_bola_detection.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_bot_defense_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_failed_login_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_forbidden_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_ip_reputation.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_non_existent_url_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_rate_limit.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_waf_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_bot_defense_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_ip_reputation.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--high.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--low.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_automatic--medium.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom.md#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_rate_limit.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_waf_activity.md#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_learning.md#section) |
| `description` | [description](data-sources--app_setting--reference.md#schema-description) |
| `id` | [id](data-sources--app_setting--reference.md#schema-id) |
| `labels` | [labels](data-sources--app_setting--reference.md#schema-labels) |
| `name` | [name](data-sources--app_setting--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--app_setting--reference.md#schema-namespace) |

## Next pages

- [app_type_settings](data-sources--app_setting--properties--app_type_settings.md)
- [xcsh_app_setting](../data-sources/app_setting.md)
