---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 24043, "body_sha256": "sha256:c93f46b60c3f1b1b1af21cfba1103b6a25478a4c7d95a0fa1d25dfa524bdf3ec", "canonical_id": "xcsh-docs:data-sources:app_firewall:reference", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:allow_all_response_codes", "xcsh-docs:data-sources:app_firewall:properties:allowed_response_codes", "xcsh-docs:data-sources:app_firewall:properties:blocking", "xcsh-docs:data-sources:app_firewall:properties:blocking_page", "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization", "xcsh-docs:data-sources:app_firewall:properties:default_anonymization", "xcsh-docs:data-sources:app_firewall:properties:default_bot_setting", "xcsh-docs:data-sources:app_firewall:properties:default_detection_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "xcsh-docs:data-sources:app_firewall:properties:disable_ai_enhancements", "xcsh-docs:data-sources:app_firewall:properties:disable_anonymization", "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "xcsh-docs:data-sources:app_firewall:properties:monitoring", "xcsh-docs:data-sources:app_firewall:properties:use_default_blocking_page"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:reference", "parent_id": "xcsh-docs:data-sources:app_firewall:fundamentals", "path": "docs/guides/data-sources--app_firewall--reference.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- Property reference

## Direct properties

- [allow_all_response_codes](data-sources--app_firewall--properties--allow_all_response_codes.md): complete subsection reference.

- [allowed_response_codes](data-sources--app_firewall--properties--allowed_response_codes.md): complete subsection reference.

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

- [blocking](data-sources--app_firewall--properties--blocking.md): complete subsection reference.

- [blocking_page](data-sources--app_firewall--properties--blocking_page.md): complete subsection reference.

- [bot_protection_setting](data-sources--app_firewall--properties--bot_protection_setting.md): complete subsection reference.

- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md): complete subsection reference.

- [default_anonymization](data-sources--app_firewall--properties--default_anonymization.md): complete subsection reference.

- [default_bot_setting](data-sources--app_firewall--properties--default_bot_setting.md): complete subsection reference.

- [default_detection_settings](data-sources--app_firewall--properties--default_detection_settings.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AppFirewall.

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

- [detection_settings](data-sources--app_firewall--properties--detection_settings.md): complete subsection reference.

- [disable_ai_enhancements](data-sources--app_firewall--properties--disable_ai_enhancements.md): complete subsection reference.

- [disable_anonymization](data-sources--app_firewall--properties--disable_anonymization.md): complete subsection reference.

- [enable_ai_enhancements](data-sources--app_firewall--properties--enable_ai_enhancements.md): complete subsection reference.

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

- [monitoring](data-sources--app_firewall--properties--monitoring.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AppFirewall.

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

Namespace where the AppFirewall exists.

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

- [use_default_blocking_page](data-sources--app_firewall--properties--use_default_blocking_page.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](data-sources--app_firewall--properties--allow_all_response_codes.md#section) |
| `allowed_response_codes` | [allowed_response_codes](data-sources--app_firewall--properties--allowed_response_codes.md#section) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](data-sources--app_firewall--properties--allowed_response_codes.md#schema-allowed_response_codes--response_code) |
| `annotations` | [annotations](data-sources--app_firewall--reference.md#schema-annotations) |
| `blocking` | [blocking](data-sources--app_firewall--properties--blocking.md#section) |
| `blocking_page` | [blocking_page](data-sources--app_firewall--properties--blocking_page.md#section) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](data-sources--app_firewall--properties--blocking_page.md#schema-blocking_page--blocking_page) |
| `blocking_page.response_code` | [blocking_page.response_code](data-sources--app_firewall--properties--blocking_page.md#schema-blocking_page--response_code) |
| `bot_protection_setting` | [bot_protection_setting](data-sources--app_firewall--properties--bot_protection_setting.md#section) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](data-sources--app_firewall--properties--bot_protection_setting.md#schema-bot_protection_setting--good_bot_action) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](data-sources--app_firewall--properties--bot_protection_setting.md#schema-bot_protection_setting--malicious_bot_action) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--properties--bot_protection_setting.md#schema-bot_protection_setting--suspicious_bot_action) |
| `custom_anonymization` | [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md#section) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](data-sources--app_firewall--properties--custom_anonymization--anonymization_config.md#section) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--cookie.md#section) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--cookie.md#schema-custom_anonymization--anonymization_config--cookie--cookie_name) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--http_header.md#section) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--http_header.md#schema-custom_anonymization--anonymization_config--http_header--header_name) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--query_parameter.md#section) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](data-sources--app_firewall--properties--custom_anonymization--anonymization_config--query_parameter.md#schema-custom_anonymization--anonymization_config--query_parameter--query_param_name) |
| `default_anonymization` | [default_anonymization](data-sources--app_firewall--properties--default_anonymization.md#section) |
| `default_bot_setting` | [default_bot_setting](data-sources--app_firewall--properties--default_bot_setting.md#section) |
| `default_detection_settings` | [default_detection_settings](data-sources--app_firewall--properties--default_detection_settings.md#section) |
| `description` | [description](data-sources--app_firewall--reference.md#schema-description) |
| `detection_settings` | [detection_settings](data-sources--app_firewall--properties--detection_settings.md#section) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](data-sources--app_firewall--properties--detection_settings--bot_protection_setting.md#section) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](data-sources--app_firewall--properties--detection_settings--bot_protection_setting.md#schema-detection_settings--bot_protection_setting--good_bot_action) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](data-sources--app_firewall--properties--detection_settings--bot_protection_setting.md#schema-detection_settings--bot_protection_setting--malicious_bot_action) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--properties--detection_settings--bot_protection_setting.md#schema-detection_settings--bot_protection_setting--suspicious_bot_action) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](data-sources--app_firewall--properties--detection_settings--default_bot_setting.md#section) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](data-sources--app_firewall--properties--detection_settings--default_violation_settings.md#section) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](data-sources--app_firewall--properties--detection_settings--disable_staging.md#section) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](data-sources--app_firewall--properties--detection_settings--disable_suppression.md#section) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](data-sources--app_firewall--properties--detection_settings--disable_threat_campaigns.md#section) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](data-sources--app_firewall--properties--detection_settings--enable_suppression.md#section) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](data-sources--app_firewall--properties--detection_settings--enable_threat_campaigns.md#section) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](data-sources--app_firewall--properties--detection_settings--signature_selection_setting.md#section) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md#section) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md#schema-detection_settings--signature_selection_setting--attack_type_settings--disabled_attack_types) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md#section) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_signature_setting.md#section) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_accuracy_signatures.md#section) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_low_accuracy_signatures.md#section) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--only_high_accuracy_signatures.md#section) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md#section) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--low_accuracy_action) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md#schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--medium_accuracy_action) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](data-sources--app_firewall--properties--detection_settings--stage_new_and_updated_signatures.md#section) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](data-sources--app_firewall--properties--detection_settings--stage_new_and_updated_signatures.md#schema-detection_settings--stage_new_and_updated_signatures--staging_period) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](data-sources--app_firewall--properties--detection_settings--stage_new_signatures.md#section) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](data-sources--app_firewall--properties--detection_settings--stage_new_signatures.md#schema-detection_settings--stage_new_signatures--staging_period) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](data-sources--app_firewall--properties--detection_settings--violation_settings.md#section) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](data-sources--app_firewall--properties--detection_settings--violation_settings.md#schema-detection_settings--violation_settings--disabled_violation_types) |
| `detection_settings.violations_view` | [detection_settings.violations_view](data-sources--app_firewall--properties--detection_settings--violations_view.md#section) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](data-sources--app_firewall--properties--detection_settings--violations_view.md#schema-detection_settings--violations_view--description_spec) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](data-sources--app_firewall--properties--detection_settings--violations_view.md#schema-detection_settings--violations_view--enabled) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](data-sources--app_firewall--properties--detection_settings--violations_view.md#schema-detection_settings--violations_view--enabled_by_default) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](data-sources--app_firewall--properties--detection_settings--violations_view.md#schema-detection_settings--violations_view--name) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](data-sources--app_firewall--properties--detection_settings--violations_view.md#schema-detection_settings--violations_view--title) |
| `disable_ai_enhancements` | [disable_ai_enhancements](data-sources--app_firewall--properties--disable_ai_enhancements.md#section) |
| `disable_anonymization` | [disable_anonymization](data-sources--app_firewall--properties--disable_anonymization.md#section) |
| `enable_ai_enhancements` | [enable_ai_enhancements](data-sources--app_firewall--properties--enable_ai_enhancements.md#section) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_medium_risk_action.md#section) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](data-sources--app_firewall--properties--enable_ai_enhancements--mitigate_high_risk_action.md#section) |
| `id` | [id](data-sources--app_firewall--reference.md#schema-id) |
| `labels` | [labels](data-sources--app_firewall--reference.md#schema-labels) |
| `monitoring` | [monitoring](data-sources--app_firewall--properties--monitoring.md#section) |
| `name` | [name](data-sources--app_firewall--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--app_firewall--reference.md#schema-namespace) |
| `use_default_blocking_page` | [use_default_blocking_page](data-sources--app_firewall--properties--use_default_blocking_page.md#section) |

## Next pages

- [allow_all_response_codes](data-sources--app_firewall--properties--allow_all_response_codes.md)
- [allowed_response_codes](data-sources--app_firewall--properties--allowed_response_codes.md)
- [blocking](data-sources--app_firewall--properties--blocking.md)
- [blocking_page](data-sources--app_firewall--properties--blocking_page.md)
- [bot_protection_setting](data-sources--app_firewall--properties--bot_protection_setting.md)
- [custom_anonymization](data-sources--app_firewall--properties--custom_anonymization.md)
- [default_anonymization](data-sources--app_firewall--properties--default_anonymization.md)
- [default_bot_setting](data-sources--app_firewall--properties--default_bot_setting.md)
- [default_detection_settings](data-sources--app_firewall--properties--default_detection_settings.md)
- [detection_settings](data-sources--app_firewall--properties--detection_settings.md)
- [disable_ai_enhancements](data-sources--app_firewall--properties--disable_ai_enhancements.md)
- [disable_anonymization](data-sources--app_firewall--properties--disable_anonymization.md)
- [enable_ai_enhancements](data-sources--app_firewall--properties--enable_ai_enhancements.md)
- [monitoring](data-sources--app_firewall--properties--monitoring.md)
- [use_default_blocking_page](data-sources--app_firewall--properties--use_default_blocking_page.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
