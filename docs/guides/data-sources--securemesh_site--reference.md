---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 99793, "body_sha256": "sha256:8239148fccd16a6bc3b678db43c75c074787b857c32b31563f1acda7e9490dcb", "canonical_id": "xcsh-docs:data-sources:securemesh_site:reference", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:blocked_services", "xcsh-docs:data-sources:securemesh_site:properties:bond_device_list", "xcsh-docs:data-sources:securemesh_site:properties:coordinates", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "xcsh-docs:data-sources:securemesh_site:properties:default_blocked_services", "xcsh-docs:data-sources:securemesh_site:properties:default_network_config", "xcsh-docs:data-sources:securemesh_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:securemesh_site:properties:log_receiver", "xcsh-docs:data-sources:securemesh_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:securemesh_site:properties:master_node_configuration", "xcsh-docs:data-sources:securemesh_site:properties:no_bond_devices", "xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:securemesh_site:properties:os", "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode", "xcsh-docs:data-sources:securemesh_site:properties:sw", "xcsh-docs:data-sources:securemesh_site:properties:waf_signatures"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:reference", "parent_id": "xcsh-docs:data-sources:securemesh_site:fundamentals", "path": "docs/guides/data-sources--securemesh_site--reference.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

- [blocked_services](data-sources--securemesh_site--properties--blocked_services.md): complete subsection reference.

- [bond_device_list](data-sources--securemesh_site--properties--bond_device_list.md): complete subsection reference.

- [coordinates](data-sources--securemesh_site--properties--coordinates.md): complete subsection reference.

- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md): complete subsection reference.

- [default_blocked_services](data-sources--securemesh_site--properties--default_blocked_services.md): complete subsection reference.

- [default_network_config](data-sources--securemesh_site--properties--default_network_config.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the SecuremeshSite.

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

- [kubernetes_upgrade_drain](data-sources--securemesh_site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](data-sources--securemesh_site--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](data-sources--securemesh_site--properties--logs_streaming_disabled.md): complete subsection reference.

- [master_node_configuration](data-sources--securemesh_site--properties--master_node_configuration.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the SecuremeshSite.

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

Namespace where the SecuremeshSite exists.

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

- [no_bond_devices](data-sources--securemesh_site--properties--no_bond_devices.md): complete subsection reference.

- [offline_survivability_mode](data-sources--securemesh_site--properties--offline_survivability_mode.md): complete subsection reference.

- [os](data-sources--securemesh_site--properties--os.md): complete subsection reference.

- [performance_enhancement_mode](data-sources--securemesh_site--properties--performance_enhancement_mode.md): complete subsection reference.

- [sw](data-sources--securemesh_site--properties--sw.md): complete subsection reference.

<a id="schema-volterra_certified_hw"></a>

### volterra_certified_hw property

Type: `"string"`. Computed.

Name for generic server certified hardware to form this Secure Mesh site.

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

- [waf_signatures](data-sources--securemesh_site--properties--waf_signatures.md): complete subsection reference.

<a id="schema-worker_nodes"></a>

### worker_nodes property

Type: `["list", "string"]`. Computed.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

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
| `address` | [address](data-sources--securemesh_site--reference.md#schema-address) |
| `annotations` | [annotations](data-sources--securemesh_site--reference.md#schema-annotations) |
| `blocked_services` | [blocked_services](data-sources--securemesh_site--properties--blocked_services.md#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--securemesh_site--properties--blocked_services--blocked_service.md#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--securemesh_site--properties--blocked_services--blocked_service--dns.md#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--securemesh_site--properties--blocked_services--blocked_service.md#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--securemesh_site--properties--blocked_services--blocked_service--ssh.md#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site--properties--blocked_services--blocked_service--web_user_interface.md#section) |
| `bond_device_list` | [bond_device_list](data-sources--securemesh_site--properties--bond_device_list.md#section) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--securemesh_site--properties--bond_device_list--bond_devices.md#section) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--securemesh_site--properties--bond_device_list--bond_devices--active_backup.md#section) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--devices) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--securemesh_site--properties--bond_device_list--bond_devices--lacp.md#section) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--securemesh_site--properties--bond_device_list--bond_devices--lacp.md#schema-bond_device_list--bond_devices--lacp--rate) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_polling_interval) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_up_delay) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--securemesh_site--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--name) |
| `coordinates` | [coordinates](data-sources--securemesh_site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--securemesh_site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--securemesh_site--properties--coordinates.md#schema-coordinates--longitude) |
| `custom_network_config` | [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md#section) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies.md#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies.md#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](data-sources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--name) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--securemesh_site--properties--custom_network_config--active_forward_proxy_policies--forward_proxy_policies.md#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](data-sources--securemesh_site--properties--custom_network_config--active_network_policies.md#section) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#section) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--name) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--namespace) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](data-sources--securemesh_site--properties--custom_network_config--active_network_policies--network_policies.md#schema-custom_network_config--active_network_policies--network_policies--tenant) |
| `custom_network_config.default_config` | [custom_network_config.default_config](data-sources--securemesh_site--properties--custom_network_config--default_config.md#section) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](data-sources--securemesh_site--properties--custom_network_config--default_interface_config.md#section) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](data-sources--securemesh_site--properties--custom_network_config--default_sli_config.md#section) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](data-sources--securemesh_site--properties--custom_network_config--forward_proxy_allow_all.md#section) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](data-sources--securemesh_site--properties--custom_network_config--global_network_list.md#section) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr.md#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](data-sources--securemesh_site--properties--custom_network_config--interface_list.md#section) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dc_cluster_group_connectivity_interface_enabled.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--is_primary.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--monitor.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--monitor_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--node) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface--not_primary.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_interface--priority) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--dedicated_management_interface.md#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--node) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#schema-custom_network_config--interface_list--interfaces--description_spec) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--cluster.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--device) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_client.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_end.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_start.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--first_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--last_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--exclude) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--same_as_dgw.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--host.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_end.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_start.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--is_primary.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor_disabled.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--mtu) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--no_ipv6_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--node) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--not_primary.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--priority) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_inside_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--storage_network.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--untagged.md#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md#schema-custom_network_config--interface_list--interfaces--ethernet_interface--vlan_id) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md#schema-custom_network_config--interface_list--interfaces--labels) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](data-sources--securemesh_site--properties--custom_network_config--no_forward_proxy.md#section) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](data-sources--securemesh_site--properties--custom_network_config--no_global_network.md#section) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](data-sources--securemesh_site--properties--custom_network_config--no_network_policy.md#section) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](data-sources--securemesh_site--properties--custom_network_config--sli_config.md#section) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](data-sources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#section) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](data-sources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--name) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](data-sources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--namespace) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](data-sources--securemesh_site--properties--custom_network_config--sli_config--dc_cluster_group.md#schema-custom_network_config--sli_config--dc_cluster_group--tenant) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](data-sources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--labels) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](data-sources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--nameserver) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](data-sources--securemesh_site--properties--custom_network_config--sli_config--no_dc_cluster_group.md#section) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--no_static_routes.md#section) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--no_v6_static_routes.md#section) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes.md#schema-custom_network_config--sli_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_routes--static_routes--node_interface--list.md#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--properties--custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list.md#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](data-sources--securemesh_site--properties--custom_network_config--sli_config.md#schema-custom_network_config--sli_config--vip) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](data-sources--securemesh_site--properties--custom_network_config--slo_config.md#section) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](data-sources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#section) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](data-sources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--name) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](data-sources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--namespace) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](data-sources--securemesh_site--properties--custom_network_config--slo_config--dc_cluster_group.md#schema-custom_network_config--slo_config--dc_cluster_group--tenant) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](data-sources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--labels) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](data-sources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--nameserver) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](data-sources--securemesh_site--properties--custom_network_config--slo_config--no_dc_cluster_group.md#section) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--no_static_routes.md#section) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--no_v6_static_routes.md#section) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md#schema-custom_network_config--slo_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--default_gateway.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list.md#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](data-sources--securemesh_site--properties--custom_network_config--slo_config.md#schema-custom_network_config--slo_config--vip) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](data-sources--securemesh_site--properties--custom_network_config--sm_connection_public_ip.md#section) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](data-sources--securemesh_site--properties--custom_network_config--sm_connection_pvt_ip.md#section) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](data-sources--securemesh_site--properties--custom_network_config.md#schema-custom_network_config--tunnel_dead_timeout) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](data-sources--securemesh_site--properties--custom_network_config.md#schema-custom_network_config--vip_vrrp_mode) |
| `default_blocked_services` | [default_blocked_services](data-sources--securemesh_site--properties--default_blocked_services.md#section) |
| `default_network_config` | [default_network_config](data-sources--securemesh_site--properties--default_network_config.md#section) |
| `description` | [description](data-sources--securemesh_site--reference.md#schema-description) |
| `id` | [id](data-sources--securemesh_site--reference.md#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--securemesh_site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--securemesh_site--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](data-sources--securemesh_site--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](data-sources--securemesh_site--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--securemesh_site--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--securemesh_site--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--securemesh_site--properties--logs_streaming_disabled.md#section) |
| `master_node_configuration` | [master_node_configuration](data-sources--securemesh_site--properties--master_node_configuration.md#section) |
| `master_node_configuration.name` | [master_node_configuration.name](data-sources--securemesh_site--properties--master_node_configuration.md#schema-master_node_configuration--name) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](data-sources--securemesh_site--properties--master_node_configuration.md#schema-master_node_configuration--public_ip) |
| `name` | [name](data-sources--securemesh_site--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--securemesh_site--reference.md#schema-namespace) |
| `no_bond_devices` | [no_bond_devices](data-sources--securemesh_site--properties--no_bond_devices.md#section) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--securemesh_site--properties--offline_survivability_mode.md#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--securemesh_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--securemesh_site--properties--offline_survivability_mode--no_offline_survivability_mode.md#section) |
| `os` | [os](data-sources--securemesh_site--properties--os.md#section) |
| `os.default_os_version` | [os.default_os_version](data-sources--securemesh_site--properties--os--default_os_version.md#section) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--securemesh_site--properties--os.md#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--securemesh_site--properties--performance_enhancement_mode.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `sw` | [sw](data-sources--securemesh_site--properties--sw.md#section) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--securemesh_site--properties--sw--default_sw_version.md#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--securemesh_site--properties--sw.md#schema-sw--volterra_software_version) |
| `volterra_certified_hw` | [volterra_certified_hw](data-sources--securemesh_site--reference.md#schema-volterra_certified_hw) |
| `waf_signatures` | [waf_signatures](data-sources--securemesh_site--properties--waf_signatures.md#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--securemesh_site--properties--waf_signatures--automatic.md#section) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--securemesh_site--properties--waf_signatures--manual.md#section) |
| `worker_nodes` | [worker_nodes](data-sources--securemesh_site--reference.md#schema-worker_nodes) |

## Next pages

- [blocked_services](data-sources--securemesh_site--properties--blocked_services.md)
- [bond_device_list](data-sources--securemesh_site--properties--bond_device_list.md)
- [coordinates](data-sources--securemesh_site--properties--coordinates.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
- [default_blocked_services](data-sources--securemesh_site--properties--default_blocked_services.md)
- [default_network_config](data-sources--securemesh_site--properties--default_network_config.md)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--properties--kubernetes_upgrade_drain.md)
- [log_receiver](data-sources--securemesh_site--properties--log_receiver.md)
- [logs_streaming_disabled](data-sources--securemesh_site--properties--logs_streaming_disabled.md)
- [master_node_configuration](data-sources--securemesh_site--properties--master_node_configuration.md)
- [no_bond_devices](data-sources--securemesh_site--properties--no_bond_devices.md)
- [offline_survivability_mode](data-sources--securemesh_site--properties--offline_survivability_mode.md)
- [os](data-sources--securemesh_site--properties--os.md)
- [performance_enhancement_mode](data-sources--securemesh_site--properties--performance_enhancement_mode.md)
- [sw](data-sources--securemesh_site--properties--sw.md)
- [waf_signatures](data-sources--securemesh_site--properties--waf_signatures.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
