---
page_title: "Property reference"
subcategory: "Monitoring"
description: "Property reference for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 18876, "body_sha256": "sha256:37df70b2276d7da5317260d24ff619b8173ed483aa0582c57a4fa2fe26c180b2", "child_ids": ["xcsh-docs:resources:alert_policy:properties:notification_parameters", "xcsh-docs:resources:alert_policy:properties:receivers", "xcsh-docs:resources:alert_policy:properties:routes", "xcsh-docs:resources:alert_policy:properties:timeouts"], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:reference", "parent_id": "xcsh-docs:resources:alert_policy:fundamentals", "path": "documentation/resources/alert_policy/properties/index.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
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

Name of the Alert Policy. Must be unique within the namespace.

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

Namespace where the Alert Policy is created.

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

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/): complete subsection reference.

- [receivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/): complete subsection reference.

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/#schema-namespace) |
| `notification_parameters` | [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/#section) |
| `notification_parameters.custom` | [notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/custom/#section) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/custom/#schema-notification_parameters--custom--labels) |
| `notification_parameters.default` | [notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/default/#section) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/#schema-notification_parameters--group_interval) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/#schema-notification_parameters--group_wait) |
| `notification_parameters.individual` | [notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/individual/#section) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/#schema-notification_parameters--repeat_interval) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/ves_io_group/#section) |
| `receivers` | [receivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#section) |
| `receivers.kind` | [receivers.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#schema-receivers--kind) |
| `receivers.name` | [receivers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#schema-receivers--name) |
| `receivers.namespace` | [receivers.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#schema-receivers--namespace) |
| `receivers.tenant` | [receivers.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#schema-receivers--tenant) |
| `receivers.uid` | [receivers.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/#schema-receivers--uid) |
| `routes` | [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/#section) |
| `routes.alertname` | [routes.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/#schema-routes--alertname) |
| `routes.alertname_regex` | [routes.alertname_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/#schema-routes--alertname_regex) |
| `routes.any` | [routes.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/any/#section) |
| `routes.custom` | [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/#section) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertlabel/#section) |
| `routes.custom.alertname` | [routes.custom.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertname/#section) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertname/#schema-routes--custom--alertname--exact_match) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/alertname/#schema-routes--custom--alertname--regex_match) |
| `routes.custom.group` | [routes.custom.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/group/#section) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/group/#schema-routes--custom--group--exact_match) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/group/#schema-routes--custom--group--regex_match) |
| `routes.custom.severity` | [routes.custom.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/severity/#section) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/severity/#schema-routes--custom--severity--exact_match) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/severity/#schema-routes--custom--severity--regex_match) |
| `routes.dont_send` | [routes.dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/dont_send/#section) |
| `routes.group` | [routes.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/group/#section) |
| `routes.group.groups` | [routes.group.groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/group/#schema-routes--group--groups) |
| `routes.notification_parameters` | [routes.notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/#section) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/custom/#section) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/custom/#schema-routes--notification_parameters--custom--labels) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/default/#section) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--group_interval) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--group_wait) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/individual/#section) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--repeat_interval) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/notification_parameters/ves_io_group/#section) |
| `routes.send` | [routes.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/send/#section) |
| `routes.severity` | [routes.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/severity/#section) |
| `routes.severity.severities` | [routes.severity.severities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/severity/#schema-routes--severity--severities) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/)
- [receivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/receivers/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/timeouts/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
