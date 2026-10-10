---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_setting."
xcsh_docs: {"aliases": ["app setting"], "body_bytes": 23074, "body_sha256": "sha256:5bd82ae537a5d96f4587a53cb90d8db4125ed2b1fb503cd334d9dd544a04771d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:reference", "parent_id": "xcsh-docs:data-sources:app_setting:fundamentals", "path": "documentation/data-sources/app_setting/properties/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["app type settings"], "anchor": "section", "description": "List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:app_setting:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_settingCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AppSetting.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Additional upstream details:

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

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-annotations) |
| `app_type_settings` | [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/#section) |
| `app_type_settings.app_type_ref` | [app_type_settings.app_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#section) |
| `app_type_settings.app_type_ref.kind` | [app_type_settings.app_type_ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#schema-app_type_settings--app_type_ref--kind) |
| `app_type_settings.app_type_ref.name` | [app_type_settings.app_type_ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#schema-app_type_settings--app_type_ref--name) |
| `app_type_settings.app_type_ref.namespace` | [app_type_settings.app_type_ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#schema-app_type_settings--app_type_ref--namespace) |
| `app_type_settings.app_type_ref.tenant` | [app_type_settings.app_type_ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#schema-app_type_settings--app_type_ref--tenant) |
| `app_type_settings.app_type_ref.uid` | [app_type_settings.app_type_ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/#schema-app_type_settings--app_type_ref--uid) |
| `app_type_settings.business_logic_markup_setting` | [app_type_settings.business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/#section) |
| `app_type_settings.business_logic_markup_setting.disable_spec` | [app_type_settings.business_logic_markup_setting.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/disable_spec/#section) |
| `app_type_settings.business_logic_markup_setting.enable` | [app_type_settings.business_logic_markup_setting.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/enable/#section) |
| `app_type_settings.timeseries_analyses_setting` | [app_type_settings.timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors` | [app_type_settings.timeseries_analyses_setting.metric_selectors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/#section) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metric` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric) |
| `app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source` | [app_type_settings.timeseries_analyses_setting.metric_selectors.metrics_source](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/#schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source) |
| `app_type_settings.user_behavior_analysis_setting` | [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_detection` | [app_type_settings.user_behavior_analysis_setting.disable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_detection/#section) |
| `app_type_settings.user_behavior_analysis_setting.disable_learning` | [app_type_settings.user_behavior_analysis_setting.disable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_learning/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/bola_detection_automatic/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period` | [app_type_settings.user_behavior_analysis_setting.enable_detection.cooling_off_period](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--cooling_off_period) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bola_detection/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bot_defense_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_failed_login_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_forbidden_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_ip_reputation/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_non_existent_url_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_rate_limit/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_waf_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_bot_defense_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity.login_failures_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity.forbidden_requests_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_forbidden_activity--forbidden_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_ip_reputation/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.high](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/high/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.low](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/low/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic.medium](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/medium/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom.nonexistent_requests_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/#schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_rate_limit/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity` | [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_waf_activity/#section) |
| `app_type_settings.user_behavior_analysis_setting.enable_learning` | [app_type_settings.user_behavior_analysis_setting.enable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_learning/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/#schema-namespace) |
