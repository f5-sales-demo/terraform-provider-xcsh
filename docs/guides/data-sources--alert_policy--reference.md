---
page_title: "Property reference"
subcategory: "Monitoring"
description: "Property reference for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 13965, "body_sha256": "sha256:e687546dc459503cb58ff61afcd1e33f59a3f0cae5f9a92f8a0f361e727ee29b", "canonical_id": "xcsh-docs:data-sources:alert_policy:reference", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "xcsh-docs:data-sources:alert_policy:properties:receivers", "xcsh-docs:data-sources:alert_policy:properties:routes"], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:reference", "parent_id": "xcsh-docs:data-sources:alert_policy:fundamentals", "path": "docs/guides/data-sources--alert_policy--reference.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AlertPolicy.

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

Name of the AlertPolicy.

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

Namespace where the AlertPolicy exists.

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

- [notification_parameters](data-sources--alert_policy--properties--notification_parameters.md): complete subsection reference.

- [receivers](data-sources--alert_policy--properties--receivers.md): complete subsection reference.

- [routes](data-sources--alert_policy--properties--routes.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--alert_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--alert_policy--reference.md#schema-description) |
| `id` | [id](data-sources--alert_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--alert_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--alert_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--alert_policy--reference.md#schema-namespace) |
| `notification_parameters` | [notification_parameters](data-sources--alert_policy--properties--notification_parameters.md#section) |
| `notification_parameters.custom` | [notification_parameters.custom](data-sources--alert_policy--properties--notification_parameters--custom.md#section) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](data-sources--alert_policy--properties--notification_parameters--custom.md#schema-notification_parameters--custom--labels) |
| `notification_parameters.default` | [notification_parameters.default](data-sources--alert_policy--properties--notification_parameters--default.md#section) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](data-sources--alert_policy--properties--notification_parameters.md#schema-notification_parameters--group_interval) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](data-sources--alert_policy--properties--notification_parameters.md#schema-notification_parameters--group_wait) |
| `notification_parameters.individual` | [notification_parameters.individual](data-sources--alert_policy--properties--notification_parameters--individual.md#section) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](data-sources--alert_policy--properties--notification_parameters.md#schema-notification_parameters--repeat_interval) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](data-sources--alert_policy--properties--notification_parameters--ves_io_group.md#section) |
| `receivers` | [receivers](data-sources--alert_policy--properties--receivers.md#section) |
| `receivers.kind` | [receivers.kind](data-sources--alert_policy--properties--receivers.md#schema-receivers--kind) |
| `receivers.name` | [receivers.name](data-sources--alert_policy--properties--receivers.md#schema-receivers--name) |
| `receivers.namespace` | [receivers.namespace](data-sources--alert_policy--properties--receivers.md#schema-receivers--namespace) |
| `receivers.tenant` | [receivers.tenant](data-sources--alert_policy--properties--receivers.md#schema-receivers--tenant) |
| `receivers.uid` | [receivers.uid](data-sources--alert_policy--properties--receivers.md#schema-receivers--uid) |
| `routes` | [routes](data-sources--alert_policy--properties--routes.md#section) |
| `routes.alertname` | [routes.alertname](data-sources--alert_policy--properties--routes.md#schema-routes--alertname) |
| `routes.alertname_regex` | [routes.alertname_regex](data-sources--alert_policy--properties--routes.md#schema-routes--alertname_regex) |
| `routes.any` | [routes.any](data-sources--alert_policy--properties--routes--any.md#section) |
| `routes.custom` | [routes.custom](data-sources--alert_policy--properties--routes--custom.md#section) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](data-sources--alert_policy--properties--routes--custom--alertlabel.md#section) |
| `routes.custom.alertname` | [routes.custom.alertname](data-sources--alert_policy--properties--routes--custom--alertname.md#section) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](data-sources--alert_policy--properties--routes--custom--alertname.md#schema-routes--custom--alertname--exact_match) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](data-sources--alert_policy--properties--routes--custom--alertname.md#schema-routes--custom--alertname--regex_match) |
| `routes.custom.group` | [routes.custom.group](data-sources--alert_policy--properties--routes--custom--group.md#section) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](data-sources--alert_policy--properties--routes--custom--group.md#schema-routes--custom--group--exact_match) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](data-sources--alert_policy--properties--routes--custom--group.md#schema-routes--custom--group--regex_match) |
| `routes.custom.severity` | [routes.custom.severity](data-sources--alert_policy--properties--routes--custom--severity.md#section) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](data-sources--alert_policy--properties--routes--custom--severity.md#schema-routes--custom--severity--exact_match) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](data-sources--alert_policy--properties--routes--custom--severity.md#schema-routes--custom--severity--regex_match) |
| `routes.dont_send` | [routes.dont_send](data-sources--alert_policy--properties--routes--dont_send.md#section) |
| `routes.group` | [routes.group](data-sources--alert_policy--properties--routes--group.md#section) |
| `routes.group.groups` | [routes.group.groups](data-sources--alert_policy--properties--routes--group.md#schema-routes--group--groups) |
| `routes.notification_parameters` | [routes.notification_parameters](data-sources--alert_policy--properties--routes--notification_parameters.md#section) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](data-sources--alert_policy--properties--routes--notification_parameters--custom.md#section) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](data-sources--alert_policy--properties--routes--notification_parameters--custom.md#schema-routes--notification_parameters--custom--labels) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](data-sources--alert_policy--properties--routes--notification_parameters--default.md#section) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](data-sources--alert_policy--properties--routes--notification_parameters.md#schema-routes--notification_parameters--group_interval) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](data-sources--alert_policy--properties--routes--notification_parameters.md#schema-routes--notification_parameters--group_wait) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](data-sources--alert_policy--properties--routes--notification_parameters--individual.md#section) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](data-sources--alert_policy--properties--routes--notification_parameters.md#schema-routes--notification_parameters--repeat_interval) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](data-sources--alert_policy--properties--routes--notification_parameters--ves_io_group.md#section) |
| `routes.send` | [routes.send](data-sources--alert_policy--properties--routes--send.md#section) |
| `routes.severity` | [routes.severity](data-sources--alert_policy--properties--routes--severity.md#section) |
| `routes.severity.severities` | [routes.severity.severities](data-sources--alert_policy--properties--routes--severity.md#schema-routes--severity--severities) |

## Next pages

- [notification_parameters](data-sources--alert_policy--properties--notification_parameters.md)
- [receivers](data-sources--alert_policy--properties--receivers.md)
- [routes](data-sources--alert_policy--properties--routes.md)
- [xcsh_alert_policy](../data-sources/alert_policy.md)
