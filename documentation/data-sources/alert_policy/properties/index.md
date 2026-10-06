---
page_title: "Property reference"
subcategory: "Monitoring"
description: "Property reference for xcsh_alert_policy."
xcsh_docs: {"aliases": ["alert policy"], "body_bytes": 17138, "body_sha256": "sha256:78c73961a520a88f9b16687b35dfb25097897b7cf34858d3af49037a7804c9e8", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "xcsh-docs:data-sources:alert_policy:properties:receivers", "xcsh-docs:data-sources:alert_policy:properties:routes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:reference", "parent_id": "xcsh-docs:data-sources:alert_policy:fundamentals", "path": "documentation/data-sources/alert_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:alert_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["notification parameters"], "anchor": "section", "description": "Set of notification parameters to decide how and when the alert notifications should be sent to the receivers.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["notification_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["receivers"], "anchor": "section", "description": "List of Alert Receivers where the alerts will be sent.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:receivers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["receivers"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes"], "anchor": "section", "description": "Set of routes to match the incoming alert. The routes are evaluated in the specified order and terminates on the first match.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_alert_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AlertPolicy.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Name of the AlertPolicy.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/): complete subsection reference.

- [receivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/): complete subsection reference.

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/#schema-namespace) |
| `notification_parameters` | [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/#section) |
| `notification_parameters.custom` | [notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/custom/#section) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/custom/#schema-notification_parameters--custom--labels) |
| `notification_parameters.default` | [notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/default/#section) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/#schema-notification_parameters--group_interval) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/#schema-notification_parameters--group_wait) |
| `notification_parameters.individual` | [notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/individual/#section) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/#schema-notification_parameters--repeat_interval) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/notification_parameters/ves_io_group/#section) |
| `receivers` | [receivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#section) |
| `receivers.kind` | [receivers.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#schema-receivers--kind) |
| `receivers.name` | [receivers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#schema-receivers--name) |
| `receivers.namespace` | [receivers.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#schema-receivers--namespace) |
| `receivers.tenant` | [receivers.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#schema-receivers--tenant) |
| `receivers.uid` | [receivers.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/receivers/#schema-receivers--uid) |
| `routes` | [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/#section) |
| `routes.alertname` | [routes.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/#schema-routes--alertname) |
| `routes.alertname_regex` | [routes.alertname_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/#schema-routes--alertname_regex) |
| `routes.any` | [routes.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/any/#section) |
| `routes.custom` | [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/#section) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertlabel/#section) |
| `routes.custom.alertname` | [routes.custom.alertname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertname/#section) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertname/#schema-routes--custom--alertname--exact_match) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/alertname/#schema-routes--custom--alertname--regex_match) |
| `routes.custom.group` | [routes.custom.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/group/#section) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/group/#schema-routes--custom--group--exact_match) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/group/#schema-routes--custom--group--regex_match) |
| `routes.custom.severity` | [routes.custom.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/severity/#section) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/severity/#schema-routes--custom--severity--exact_match) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/severity/#schema-routes--custom--severity--regex_match) |
| `routes.dont_send` | [routes.dont_send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/dont_send/#section) |
| `routes.group` | [routes.group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/group/#section) |
| `routes.group.groups` | [routes.group.groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/group/#schema-routes--group--groups) |
| `routes.notification_parameters` | [routes.notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/#section) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/custom/#section) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/custom/#schema-routes--notification_parameters--custom--labels) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/default/#section) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--group_interval) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--group_wait) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/individual/#section) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/#schema-routes--notification_parameters--repeat_interval) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/ves_io_group/#section) |
| `routes.send` | [routes.send](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/send/#section) |
| `routes.severity` | [routes.severity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/severity/#section) |
| `routes.severity.severities` | [routes.severity.severities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/severity/#schema-routes--severity--severities) |
