---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 14089, "body_sha256": "sha256:72d8de4a567255972741fa06c2e22b828b88289cad32756a1a7f06ef6f3d508c", "canonical_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:timeouts"], "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "path": "docs/guides/resources--bot_defense_app_infrastructure--reference.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
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

- [cloud_hosted](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md): complete subsection reference.

- [data_center_hosted](resources--bot_defense_app_infrastructure--properties--data_center_hosted.md): complete subsection reference.

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

<a id="schema-environment_type"></a>

### environment_type property

Type: `"string"`. Optional, Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PRODUCTION",
    "TESTING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PRODUCTION",
  "enum": [
    "PRODUCTION",
    "TESTING"
  ],
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

Name of the Bot Defense App Infrastructure. Must be unique within the namespace.

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

Namespace where the Bot Defense App Infrastructure is created.

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

- [timeouts](resources--bot_defense_app_infrastructure--properties--timeouts.md): complete subsection reference.

<a id="schema-traffic_type"></a>

### traffic_type property

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("WEB",
    "MOBILE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
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
| `annotations` | [annotations](resources--bot_defense_app_infrastructure--reference.md#schema-annotations) |
| `cloud_hosted` | [cloud_hosted](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md#section) |
| `cloud_hosted.egress` | [cloud_hosted.egress](resources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md#section) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md#schema-cloud_hosted--egress--ip_address) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](resources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md#schema-cloud_hosted--egress--location) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md#schema-cloud_hosted--infra_host_name) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](resources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md#section) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md#schema-cloud_hosted--ingress--host_name) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md#schema-cloud_hosted--ingress--ip_address) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](resources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md#schema-cloud_hosted--ingress--location) |
| `cloud_hosted.region` | [cloud_hosted.region](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md#schema-cloud_hosted--region) |
| `data_center_hosted` | [data_center_hosted](resources--bot_defense_app_infrastructure--properties--data_center_hosted.md#section) |
| `data_center_hosted.egress` | [data_center_hosted.egress](resources--bot_defense_app_infrastructure--properties--data_center_hosted--egress.md#section) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](resources--bot_defense_app_infrastructure--properties--data_center_hosted--egress.md#schema-data_center_hosted--egress--ip_address) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](resources--bot_defense_app_infrastructure--properties--data_center_hosted--egress.md#schema-data_center_hosted--egress--location) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](resources--bot_defense_app_infrastructure--properties--data_center_hosted.md#schema-data_center_hosted--infra_host_name) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](resources--bot_defense_app_infrastructure--properties--data_center_hosted--ingress.md#section) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](resources--bot_defense_app_infrastructure--properties--data_center_hosted--ingress.md#schema-data_center_hosted--ingress--host_name) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](resources--bot_defense_app_infrastructure--properties--data_center_hosted--ingress.md#schema-data_center_hosted--ingress--ip_address) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](resources--bot_defense_app_infrastructure--properties--data_center_hosted--ingress.md#schema-data_center_hosted--ingress--location) |
| `data_center_hosted.region` | [data_center_hosted.region](resources--bot_defense_app_infrastructure--properties--data_center_hosted.md#schema-data_center_hosted--region) |
| `description` | [description](resources--bot_defense_app_infrastructure--reference.md#schema-description) |
| `disable` | [disable](resources--bot_defense_app_infrastructure--reference.md#schema-disable) |
| `environment_type` | [environment_type](resources--bot_defense_app_infrastructure--reference.md#schema-environment_type) |
| `id` | [id](resources--bot_defense_app_infrastructure--reference.md#schema-id) |
| `labels` | [labels](resources--bot_defense_app_infrastructure--reference.md#schema-labels) |
| `name` | [name](resources--bot_defense_app_infrastructure--reference.md#schema-name) |
| `namespace` | [namespace](resources--bot_defense_app_infrastructure--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--bot_defense_app_infrastructure--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--bot_defense_app_infrastructure--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--bot_defense_app_infrastructure--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--bot_defense_app_infrastructure--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--bot_defense_app_infrastructure--properties--timeouts.md#schema-timeouts--update) |
| `traffic_type` | [traffic_type](resources--bot_defense_app_infrastructure--reference.md#schema-traffic_type) |

## Next pages

- [cloud_hosted](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md)
- [data_center_hosted](resources--bot_defense_app_infrastructure--properties--data_center_hosted.md)
- [timeouts](resources--bot_defense_app_infrastructure--properties--timeouts.md)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
