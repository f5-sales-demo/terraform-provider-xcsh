---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 100801, "body_sha256": "sha256:854c457adb914c388664802887d41b3a2e4d7af1dcb922da2ed1db6b75cb263c", "canonical_id": "xcsh-docs:resources:securemesh_site:reference", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:blocked_services", "xcsh-docs:resources:securemesh_site:properties:bond_device_list", "xcsh-docs:resources:securemesh_site:properties:coordinates", "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "xcsh-docs:resources:securemesh_site:properties:default_blocked_services", "xcsh-docs:resources:securemesh_site:properties:default_network_config", "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain", "xcsh-docs:resources:securemesh_site:properties:log_receiver", "xcsh-docs:resources:securemesh_site:properties:logs_streaming_disabled", "xcsh-docs:resources:securemesh_site:properties:master_node_configuration", "xcsh-docs:resources:securemesh_site:properties:no_bond_devices", "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode", "xcsh-docs:resources:securemesh_site:properties:os", "xcsh-docs:resources:securemesh_site:properties:performance_enhancement_mode", "xcsh-docs:resources:securemesh_site:properties:sw", "xcsh-docs:resources:securemesh_site:properties:timeouts", "xcsh-docs:resources:securemesh_site:properties:waf_signatures"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:reference", "parent_id": "xcsh-docs:resources:securemesh_site:fundamentals", "path": "docs/guides/resources--securemesh_site--reference.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

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

- [blocked_services](resources--securemesh_site--properties--blocked_services.md): complete subsection reference.

- [bond_device_list](resources--securemesh_site--properties--bond_device_list.md): complete subsection reference.

- [coordinates](resources--securemesh_site--properties--coordinates.md): complete subsection reference.

- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md): complete subsection reference.

- [default_blocked_services](resources--securemesh_site--properties--default_blocked_services.md): complete subsection reference.

- [default_network_config](resources--securemesh_site--properties--default_network_config.md): complete subsection reference.

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

- [kubernetes_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](resources--securemesh_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](resources--securemesh_site--properties--logs_streaming_disabled.md): complete subsection reference.

- [master_node_configuration](resources--securemesh_site--properties--master_node_configuration.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Securemesh Site. Must be unique within the namespace.

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

Namespace where the Securemesh Site is created.

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

- [no_bond_devices](resources--securemesh_site--properties--no_bond_devices.md): complete subsection reference.

- [offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](resources--securemesh_site--properties--os.md): complete subsection reference.

- [performance_enhancement_mode](resources--securemesh_site--properties--performance_enhancement_mode.md): complete subsection reference.

- [sw](resources--securemesh_site--properties--sw.md): complete subsection reference.

- [timeouts](resources--securemesh_site--properties--timeouts.md): complete subsection reference.

<a id="schema-volterra_certified_hw"></a>

### volterra_certified_hw property

Type: `"string"`. Required.

Name for generic server certified hardware to form this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](resources--securemesh_site--properties--waf_signatures.md): complete subsection reference.

<a id="schema-worker_nodes"></a>

### worker_nodes property

Type: `["list", "string"]`. Optional.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--securemesh_site--reference.md#schema-address) |
| `annotations` | [annotations](resources--securemesh_site--reference.md#schema-annotations) |
| `blocked_services` | [blocked_services](resources--securemesh_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--securemesh_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--securemesh_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--securemesh_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--securemesh_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--securemesh_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `bond_device_list` | [bond_device_list](resources--securemesh_site--properties--bond_device_list.md#section) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](resources--securemesh_site--properties--bond_device_list--bond_devices.md#section) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](resources--securemesh_site--properties--bond_device_list--bond_devices--active_backup.md#section) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](resources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--devices) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](resources--securemesh_site--properties--bond_device_list--bond_devices--lacp.md#section) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](resources--securemesh_site--properties--bond_device_list--bond_devices--lacp.md#schema-bond_device_list--bond_devices--lacp--rate) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](resources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_polling_interval) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](resources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_up_delay) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](resources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--name) |
| `coordinates` | [coordinates](resources--securemesh_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](resources--securemesh_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](resources--securemesh_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_network_config` | [custom_network_config](resources--securemesh_site--properties--custom_network_config.md#section) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies.md#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies.md#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](resources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--name) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](resources--securemesh_site--properties--custom_network_config--active_network_policies.md#section) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](resources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#section) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](resources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--name) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](resources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--namespace) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](resources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--tenant) |
| `custom_network_config.default_config` | [custom_network_config.default_config](resources--securemesh_site--properties--custom_network_config--default_config.md#section) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](resources--securemesh_site--properties--custom_network_config--default_interface_config.md#section) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](resources--securemesh_site--properties--custom_network_config--default_sli_config.md#section) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](resources--securemesh_site--properties--custom_network_config--forward_proxy_allow_all.md#section) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](resources--securemesh_site--properties--custom_network_config--global_network_list.md#section) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](resources--securemesh_site--properties--custom_network_config--interface_list.md#section) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_enabled.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--is_primary.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--monitor.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--monitor_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--node) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--not_primary.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--priority) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--node) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#schema-custom_network_config--interface_list--interfaces--description_spec) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--device) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_client.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_end.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_start.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--first_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--last_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--exclude) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--same_as_dgw.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--host.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_end.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_start.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--is_primary.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--mtu) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--no_ipv6_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--node) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--not_primary.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--priority) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_inside_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--storage_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--untagged.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--vlan_id) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#schema-custom_network_config--interface_list--interfaces--labels) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](resources--securemesh_site--properties--custom_network_config--no_forward_proxy.md#section) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](resources--securemesh_site--properties--custom_network_config--no_global_network.md#section) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](resources--securemesh_site--properties--custom_network_config--no_network_policy.md#section) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](resources--securemesh_site--properties--custom_network_config--sli_config.md#section) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](resources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#section) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](resources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--name) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](resources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--namespace) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](resources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--tenant) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](resources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--labels) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](resources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--nameserver) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](resources--securemesh_site--properties--custom_network_config--sli_config--no_dc_cluster_group.md#section) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--no_static_routes.md#section) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--no_v6_static_routes.md#section) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](resources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--vip) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](resources--securemesh_site--properties--custom_network_config--slo_config.md#section) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](resources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#section) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](resources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--name) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](resources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--namespace) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](resources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--tenant) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](resources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--labels) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](resources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--nameserver) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](resources--securemesh_site--properties--custom_network_config--slo_config--no_dc_cluster_group.md#section) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--no_static_routes.md#section) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--no_v6_static_routes.md#section) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](resources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--vip) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](resources--securemesh_site--properties--custom_network_config--sm_connection_public_ip.md#section) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](resources--securemesh_site--properties--custom_network_config--sm_connection_pvt_ip.md#section) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](resources--securemesh_site--properties--custom_network_config.md#schema-custom_network_config--tunnel_dead_timeout) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](resources--securemesh_site--properties--custom_network_config.md#schema-custom_network_config--vip_vrrp_mode) |
| `default_blocked_services` | [default_blocked_services](resources--securemesh_site--properties--default_blocked_services.md#section) |
| `default_network_config` | [default_network_config](resources--securemesh_site--properties--default_network_config.md#section) |
| `description` | [description](resources--securemesh_site--reference.md#schema-description) |
| `disable` | [disable](resources--securemesh_site--reference.md#schema-disable) |
| `id` | [id](resources--securemesh_site--reference.md#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](resources--securemesh_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](resources--securemesh_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](resources--securemesh_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--securemesh_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--securemesh_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--securemesh_site--properties--logs_streaming_disabled.md#section) |
| `master_node_configuration` | [master_node_configuration](resources--securemesh_site--properties--master_node_configuration.md#section) |
| `master_node_configuration.name` | [master_node_configuration.name](resources--securemesh_site--properties--master_node_configuration.md#schema-master_node_configuration--name) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](resources--securemesh_site--properties--master_node_configuration.md#schema-master_node_configuration--public_ip) |
| `name` | [name](resources--securemesh_site--reference.md#schema-name) |
| `namespace` | [namespace](resources--securemesh_site--reference.md#schema-namespace) |
| `no_bond_devices` | [no_bond_devices](resources--securemesh_site--properties--no_bond_devices.md#section) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](resources--securemesh_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](resources--securemesh_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](resources--securemesh_site--properties--os.md#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--securemesh_site--properties--performance_enhancement_mode.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `sw` | [sw](resources--securemesh_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--securemesh_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--securemesh_site--properties--sw.md#schema-sw--volterra_software_version) |
| `timeouts` | [timeouts](resources--securemesh_site--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--securemesh_site--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--securemesh_site--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--securemesh_site--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--securemesh_site--properties--timeouts.md#schema-timeouts--update) |
| `volterra_certified_hw` | [volterra_certified_hw](resources--securemesh_site--reference.md#schema-volterra_certified_hw) |
| `waf_signatures` | [waf_signatures](resources--securemesh_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--securemesh_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--securemesh_site--properties--waf_signatures--manual.md#section) |
| `worker_nodes` | [worker_nodes](resources--securemesh_site--reference.md#schema-worker_nodes) |

## Next pages

- [blocked_services](resources--securemesh_site--properties--blocked_services.md)
- [bond_device_list](resources--securemesh_site--properties--bond_device_list.md)
- [coordinates](resources--securemesh_site--properties--coordinates.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [default_blocked_services](resources--securemesh_site--properties--default_blocked_services.md)
- [default_network_config](resources--securemesh_site--properties--default_network_config.md)
- [kubernetes_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](resources--securemesh_site--properties--log_receiver.md)
- [logs_streaming_disabled](resources--securemesh_site--properties--logs_streaming_disabled.md)
- [master_node_configuration](resources--securemesh_site--properties--master_node_configuration.md)
- [no_bond_devices](resources--securemesh_site--properties--no_bond_devices.md)
- [offline_survivability_mode](resources--securemesh_site--properties--offline_survivability_mode.md)
- [os](resources--securemesh_site--properties--os.md)
- [performance_enhancement_mode](resources--securemesh_site--properties--performance_enhancement_mode.md)
- [sw](resources--securemesh_site--properties--sw.md)
- [timeouts](resources--securemesh_site--properties--timeouts.md)
- [waf_signatures](resources--securemesh_site--properties--waf_signatures.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
