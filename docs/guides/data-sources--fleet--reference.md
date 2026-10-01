---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 150993, "body_sha256": "sha256:cfe86d20b5e8967b451429a41d6f47b00178b5e069b5e6c2e697703c998f7548", "canonical_id": "xcsh-docs:data-sources:fleet:reference", "child_ids": ["xcsh-docs:data-sources:fleet:properties:allow_all_usb", "xcsh-docs:data-sources:fleet:properties:blocked_services", "xcsh-docs:data-sources:fleet:properties:bond_device_list", "xcsh-docs:data-sources:fleet:properties:dc_cluster_group", "xcsh-docs:data-sources:fleet:properties:dc_cluster_group_inside", "xcsh-docs:data-sources:fleet:properties:default_config", "xcsh-docs:data-sources:fleet:properties:default_sriov_interface", "xcsh-docs:data-sources:fleet:properties:default_storage_class", "xcsh-docs:data-sources:fleet:properties:deny_all_usb", "xcsh-docs:data-sources:fleet:properties:device_list", "xcsh-docs:data-sources:fleet:properties:disable_gpu", "xcsh-docs:data-sources:fleet:properties:disable_log_anonymization", "xcsh-docs:data-sources:fleet:properties:disable_vm", "xcsh-docs:data-sources:fleet:properties:enable_gpu", "xcsh-docs:data-sources:fleet:properties:enable_log_anonymization", "xcsh-docs:data-sources:fleet:properties:enable_vgpu", "xcsh-docs:data-sources:fleet:properties:enable_vm", "xcsh-docs:data-sources:fleet:properties:inside_virtual_network", "xcsh-docs:data-sources:fleet:properties:interface_list", "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:fleet:properties:log_receiver", "xcsh-docs:data-sources:fleet:properties:logs_streaming_disabled", "xcsh-docs:data-sources:fleet:properties:network_connectors", "xcsh-docs:data-sources:fleet:properties:network_firewall", "xcsh-docs:data-sources:fleet:properties:no_bond_devices", "xcsh-docs:data-sources:fleet:properties:no_dc_cluster_group", "xcsh-docs:data-sources:fleet:properties:no_storage_device", "xcsh-docs:data-sources:fleet:properties:no_storage_interfaces", "xcsh-docs:data-sources:fleet:properties:no_storage_static_routes", "xcsh-docs:data-sources:fleet:properties:outside_virtual_network", "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode", "xcsh-docs:data-sources:fleet:properties:sriov_interfaces", "xcsh-docs:data-sources:fleet:properties:storage_class_list", "xcsh-docs:data-sources:fleet:properties:storage_device_list", "xcsh-docs:data-sources:fleet:properties:storage_interface_list", "xcsh-docs:data-sources:fleet:properties:storage_static_routes", "xcsh-docs:data-sources:fleet:properties:usb_policy"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:reference", "parent_id": "xcsh-docs:data-sources:fleet:fundamentals", "path": "docs/guides/data-sources--fleet--reference.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- Property reference

## Direct properties

- [allow_all_usb](data-sources--fleet--properties--allow_all_usb.md): complete subsection reference.

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

- [blocked_services](data-sources--fleet--properties--blocked_services.md): complete subsection reference.

- [bond_device_list](data-sources--fleet--properties--bond_device_list.md): complete subsection reference.

- [dc_cluster_group](data-sources--fleet--properties--dc_cluster_group.md): complete subsection reference.

- [dc_cluster_group_inside](data-sources--fleet--properties--dc_cluster_group_inside.md): complete subsection reference.

- [default_config](data-sources--fleet--properties--default_config.md): complete subsection reference.

- [default_sriov_interface](data-sources--fleet--properties--default_sriov_interface.md): complete subsection reference.

- [default_storage_class](data-sources--fleet--properties--default_storage_class.md): complete subsection reference.

- [deny_all_usb](data-sources--fleet--properties--deny_all_usb.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Fleet.

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

- [device_list](data-sources--fleet--properties--device_list.md): complete subsection reference.

- [disable_gpu](data-sources--fleet--properties--disable_gpu.md): complete subsection reference.

- [disable_log_anonymization](data-sources--fleet--properties--disable_log_anonymization.md): complete subsection reference.

- [disable_vm](data-sources--fleet--properties--disable_vm.md): complete subsection reference.

<a id="schema-enable_default_fleet_config_download"></a>

### enable_default_fleet_config_download property

Type: `"bool"`. Computed.

Enable default fleet config, It must be set for storage config and GPU config.

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

- [enable_gpu](data-sources--fleet--properties--enable_gpu.md): complete subsection reference.

- [enable_log_anonymization](data-sources--fleet--properties--enable_log_anonymization.md): complete subsection reference.

- [enable_vgpu](data-sources--fleet--properties--enable_vgpu.md): complete subsection reference.

- [enable_vm](data-sources--fleet--properties--enable_vm.md): complete subsection reference.

<a id="schema-fleet_label"></a>

### fleet_label property

Type: `"string"`. Computed.

Fleet\_label value is used to create known\_label 'F5 XC/fleet=&lt;fleet\_label&gt;' The
known\_label is created in the 'shared' namespace for the tenant. A virtual\_site object with name
&lt;fleet\_label&gt; is also created in 'shared' namespace for tenant. The virtual\_site object will
select all sites..

Upstream description:

Fleet\_label value is used to create known\_label "F5 XC/fleet=&lt;fleet\_label&gt;" The
known\_label is created in the "shared" namespace for the tenant.

A virtual\_site object with name &lt;fleet\_label&gt; is also created in "shared" namespace for
tenant. The virtual\_site object will select all sites configured with the known\_label above
fleet\_label with "sfo" will create a known\_label "F5 XC/fleet=sfo" in tenant for the fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.k8s_label_value": "true"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [inside_virtual_network](data-sources--fleet--properties--inside_virtual_network.md): complete subsection reference.

- [interface_list](data-sources--fleet--properties--interface_list.md): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain.md): complete subsection reference.

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

- [log_receiver](data-sources--fleet--properties--log_receiver.md): complete subsection reference.

- [logs_streaming_disabled](data-sources--fleet--properties--logs_streaming_disabled.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Fleet.

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

Namespace where the Fleet exists.

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

- [network_connectors](data-sources--fleet--properties--network_connectors.md): complete subsection reference.

- [network_firewall](data-sources--fleet--properties--network_firewall.md): complete subsection reference.

- [no_bond_devices](data-sources--fleet--properties--no_bond_devices.md): complete subsection reference.

- [no_dc_cluster_group](data-sources--fleet--properties--no_dc_cluster_group.md): complete subsection reference.

- [no_storage_device](data-sources--fleet--properties--no_storage_device.md): complete subsection reference.

- [no_storage_interfaces](data-sources--fleet--properties--no_storage_interfaces.md): complete subsection reference.

- [no_storage_static_routes](data-sources--fleet--properties--no_storage_static_routes.md): complete subsection reference.

<a id="schema-operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Computed.

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

Upstream description:

Desired Operating System version that is applied to all sites that are member of the fleet. Current
Operating System version can be overridden via site config.

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

- [outside_virtual_network](data-sources--fleet--properties--outside_virtual_network.md): complete subsection reference.

- [performance_enhancement_mode](data-sources--fleet--properties--performance_enhancement_mode.md): complete subsection reference.

- [sriov_interfaces](data-sources--fleet--properties--sriov_interfaces.md): complete subsection reference.

- [storage_class_list](data-sources--fleet--properties--storage_class_list.md): complete subsection reference.

- [storage_device_list](data-sources--fleet--properties--storage_device_list.md): complete subsection reference.

- [storage_interface_list](data-sources--fleet--properties--storage_interface_list.md): complete subsection reference.

- [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md): complete subsection reference.

- [usb_policy](data-sources--fleet--properties--usb_policy.md): complete subsection reference.

<a id="schema-volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

Upstream description:

F5XC software version is human readable string matching released set of version components. The
given software version is applied to all sites that are member of the fleet. Current software
installed can be overridden via site config.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_usb` | [allow_all_usb](data-sources--fleet--properties--allow_all_usb.md#section) |
| `annotations` | [annotations](data-sources--fleet--reference.md#schema-annotations) |
| `blocked_services` | [blocked_services](data-sources--fleet--properties--blocked_services.md#section) |
| `blocked_services.dns` | [blocked_services.dns](data-sources--fleet--properties--blocked_services--dns.md#section) |
| `blocked_services.network_type` | [blocked_services.network_type](data-sources--fleet--properties--blocked_services.md#schema-blocked_services--network_type) |
| `blocked_services.ssh` | [blocked_services.ssh](data-sources--fleet--properties--blocked_services--ssh.md#section) |
| `blocked_services.web_user_interface` | [blocked_services.web_user_interface](data-sources--fleet--properties--blocked_services--web_user_interface.md#section) |
| `bond_device_list` | [bond_device_list](data-sources--fleet--properties--bond_device_list.md#section) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--fleet--properties--bond_device_list--bond_devices.md#section) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--fleet--properties--bond_device_list--bond_devices--active_backup.md#section) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--fleet--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--devices) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--fleet--properties--bond_device_list--bond_devices--lacp.md#section) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--fleet--properties--bond_device_list--bond_devices--lacp.md#schema-bond_device_list--bond_devices--lacp--rate) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--fleet--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_polling_interval) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--fleet--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--link_up_delay) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--fleet--properties--bond_device_list--bond_devices.md#schema-bond_device_list--bond_devices--name) |
| `dc_cluster_group` | [dc_cluster_group](data-sources--fleet--properties--dc_cluster_group.md#section) |
| `dc_cluster_group.name` | [dc_cluster_group.name](data-sources--fleet--properties--dc_cluster_group.md#schema-dc_cluster_group--name) |
| `dc_cluster_group.namespace` | [dc_cluster_group.namespace](data-sources--fleet--properties--dc_cluster_group.md#schema-dc_cluster_group--namespace) |
| `dc_cluster_group.tenant` | [dc_cluster_group.tenant](data-sources--fleet--properties--dc_cluster_group.md#schema-dc_cluster_group--tenant) |
| `dc_cluster_group_inside` | [dc_cluster_group_inside](data-sources--fleet--properties--dc_cluster_group_inside.md#section) |
| `dc_cluster_group_inside.name` | [dc_cluster_group_inside.name](data-sources--fleet--properties--dc_cluster_group_inside.md#schema-dc_cluster_group_inside--name) |
| `dc_cluster_group_inside.namespace` | [dc_cluster_group_inside.namespace](data-sources--fleet--properties--dc_cluster_group_inside.md#schema-dc_cluster_group_inside--namespace) |
| `dc_cluster_group_inside.tenant` | [dc_cluster_group_inside.tenant](data-sources--fleet--properties--dc_cluster_group_inside.md#schema-dc_cluster_group_inside--tenant) |
| `default_config` | [default_config](data-sources--fleet--properties--default_config.md#section) |
| `default_sriov_interface` | [default_sriov_interface](data-sources--fleet--properties--default_sriov_interface.md#section) |
| `default_storage_class` | [default_storage_class](data-sources--fleet--properties--default_storage_class.md#section) |
| `deny_all_usb` | [deny_all_usb](data-sources--fleet--properties--deny_all_usb.md#section) |
| `description` | [description](data-sources--fleet--reference.md#schema-description) |
| `device_list` | [device_list](data-sources--fleet--properties--device_list.md#section) |
| `device_list.devices` | [device_list.devices](data-sources--fleet--properties--device_list--devices.md#section) |
| `device_list.devices.name` | [device_list.devices.name](data-sources--fleet--properties--device_list--devices.md#schema-device_list--devices--name) |
| `device_list.devices.network_device` | [device_list.devices.network_device](data-sources--fleet--properties--device_list--devices--network_device.md#section) |
| `device_list.devices.network_device.interface` | [device_list.devices.network_device.interface](data-sources--fleet--properties--device_list--devices--network_device--interface.md#section) |
| `device_list.devices.network_device.interface.kind` | [device_list.devices.network_device.interface.kind](data-sources--fleet--properties--device_list--devices--network_device--interface.md#schema-device_list--devices--network_device--interface--kind) |
| `device_list.devices.network_device.interface.name` | [device_list.devices.network_device.interface.name](data-sources--fleet--properties--device_list--devices--network_device--interface.md#schema-device_list--devices--network_device--interface--name) |
| `device_list.devices.network_device.interface.namespace` | [device_list.devices.network_device.interface.namespace](data-sources--fleet--properties--device_list--devices--network_device--interface.md#schema-device_list--devices--network_device--interface--namespace) |
| `device_list.devices.network_device.interface.tenant` | [device_list.devices.network_device.interface.tenant](data-sources--fleet--properties--device_list--devices--network_device--interface.md#schema-device_list--devices--network_device--interface--tenant) |
| `device_list.devices.network_device.interface.uid` | [device_list.devices.network_device.interface.uid](data-sources--fleet--properties--device_list--devices--network_device--interface.md#schema-device_list--devices--network_device--interface--uid) |
| `device_list.devices.network_device.use` | [device_list.devices.network_device.use](data-sources--fleet--properties--device_list--devices--network_device.md#schema-device_list--devices--network_device--use) |
| `device_list.devices.owner` | [device_list.devices.owner](data-sources--fleet--properties--device_list--devices.md#schema-device_list--devices--owner) |
| `disable_gpu` | [disable_gpu](data-sources--fleet--properties--disable_gpu.md#section) |
| `disable_log_anonymization` | [disable_log_anonymization](data-sources--fleet--properties--disable_log_anonymization.md#section) |
| `disable_vm` | [disable_vm](data-sources--fleet--properties--disable_vm.md#section) |
| `enable_default_fleet_config_download` | [enable_default_fleet_config_download](data-sources--fleet--reference.md#schema-enable_default_fleet_config_download) |
| `enable_gpu` | [enable_gpu](data-sources--fleet--properties--enable_gpu.md#section) |
| `enable_log_anonymization` | [enable_log_anonymization](data-sources--fleet--properties--enable_log_anonymization.md#section) |
| `enable_vgpu` | [enable_vgpu](data-sources--fleet--properties--enable_vgpu.md#section) |
| `enable_vgpu.feature_type` | [enable_vgpu.feature_type](data-sources--fleet--properties--enable_vgpu.md#schema-enable_vgpu--feature_type) |
| `enable_vgpu.server_address` | [enable_vgpu.server_address](data-sources--fleet--properties--enable_vgpu.md#schema-enable_vgpu--server_address) |
| `enable_vgpu.server_port` | [enable_vgpu.server_port](data-sources--fleet--properties--enable_vgpu.md#schema-enable_vgpu--server_port) |
| `enable_vm` | [enable_vm](data-sources--fleet--properties--enable_vm.md#section) |
| `fleet_label` | [fleet_label](data-sources--fleet--reference.md#schema-fleet_label) |
| `id` | [id](data-sources--fleet--reference.md#schema-id) |
| `inside_virtual_network` | [inside_virtual_network](data-sources--fleet--properties--inside_virtual_network.md#section) |
| `inside_virtual_network.kind` | [inside_virtual_network.kind](data-sources--fleet--properties--inside_virtual_network.md#schema-inside_virtual_network--kind) |
| `inside_virtual_network.name` | [inside_virtual_network.name](data-sources--fleet--properties--inside_virtual_network.md#schema-inside_virtual_network--name) |
| `inside_virtual_network.namespace` | [inside_virtual_network.namespace](data-sources--fleet--properties--inside_virtual_network.md#schema-inside_virtual_network--namespace) |
| `inside_virtual_network.tenant` | [inside_virtual_network.tenant](data-sources--fleet--properties--inside_virtual_network.md#schema-inside_virtual_network--tenant) |
| `inside_virtual_network.uid` | [inside_virtual_network.uid](data-sources--fleet--properties--inside_virtual_network.md#schema-inside_virtual_network--uid) |
| `interface_list` | [interface_list](data-sources--fleet--properties--interface_list.md#section) |
| `interface_list.interfaces` | [interface_list.interfaces](data-sources--fleet--properties--interface_list--interfaces.md#section) |
| `interface_list.interfaces.name` | [interface_list.interfaces.name](data-sources--fleet--properties--interface_list--interfaces.md#schema-interface_list--interfaces--name) |
| `interface_list.interfaces.namespace` | [interface_list.interfaces.namespace](data-sources--fleet--properties--interface_list--interfaces.md#schema-interface_list--interfaces--namespace) |
| `interface_list.interfaces.tenant` | [interface_list.interfaces.tenant](data-sources--fleet--properties--interface_list--interfaces.md#schema-interface_list--interfaces--tenant) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--fleet--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--fleet--reference.md#schema-labels) |
| `log_receiver` | [log_receiver](data-sources--fleet--properties--log_receiver.md#section) |
| `log_receiver.name` | [log_receiver.name](data-sources--fleet--properties--log_receiver.md#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--fleet--properties--log_receiver.md#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--fleet--properties--log_receiver.md#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--fleet--properties--logs_streaming_disabled.md#section) |
| `name` | [name](data-sources--fleet--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--fleet--reference.md#schema-namespace) |
| `network_connectors` | [network_connectors](data-sources--fleet--properties--network_connectors.md#section) |
| `network_connectors.kind` | [network_connectors.kind](data-sources--fleet--properties--network_connectors.md#schema-network_connectors--kind) |
| `network_connectors.name` | [network_connectors.name](data-sources--fleet--properties--network_connectors.md#schema-network_connectors--name) |
| `network_connectors.namespace` | [network_connectors.namespace](data-sources--fleet--properties--network_connectors.md#schema-network_connectors--namespace) |
| `network_connectors.tenant` | [network_connectors.tenant](data-sources--fleet--properties--network_connectors.md#schema-network_connectors--tenant) |
| `network_connectors.uid` | [network_connectors.uid](data-sources--fleet--properties--network_connectors.md#schema-network_connectors--uid) |
| `network_firewall` | [network_firewall](data-sources--fleet--properties--network_firewall.md#section) |
| `network_firewall.kind` | [network_firewall.kind](data-sources--fleet--properties--network_firewall.md#schema-network_firewall--kind) |
| `network_firewall.name` | [network_firewall.name](data-sources--fleet--properties--network_firewall.md#schema-network_firewall--name) |
| `network_firewall.namespace` | [network_firewall.namespace](data-sources--fleet--properties--network_firewall.md#schema-network_firewall--namespace) |
| `network_firewall.tenant` | [network_firewall.tenant](data-sources--fleet--properties--network_firewall.md#schema-network_firewall--tenant) |
| `network_firewall.uid` | [network_firewall.uid](data-sources--fleet--properties--network_firewall.md#schema-network_firewall--uid) |
| `no_bond_devices` | [no_bond_devices](data-sources--fleet--properties--no_bond_devices.md#section) |
| `no_dc_cluster_group` | [no_dc_cluster_group](data-sources--fleet--properties--no_dc_cluster_group.md#section) |
| `no_storage_device` | [no_storage_device](data-sources--fleet--properties--no_storage_device.md#section) |
| `no_storage_interfaces` | [no_storage_interfaces](data-sources--fleet--properties--no_storage_interfaces.md#section) |
| `no_storage_static_routes` | [no_storage_static_routes](data-sources--fleet--properties--no_storage_static_routes.md#section) |
| `operating_system_version` | [operating_system_version](data-sources--fleet--reference.md#schema-operating_system_version) |
| `outside_virtual_network` | [outside_virtual_network](data-sources--fleet--properties--outside_virtual_network.md#section) |
| `outside_virtual_network.kind` | [outside_virtual_network.kind](data-sources--fleet--properties--outside_virtual_network.md#schema-outside_virtual_network--kind) |
| `outside_virtual_network.name` | [outside_virtual_network.name](data-sources--fleet--properties--outside_virtual_network.md#schema-outside_virtual_network--name) |
| `outside_virtual_network.namespace` | [outside_virtual_network.namespace](data-sources--fleet--properties--outside_virtual_network.md#schema-outside_virtual_network--namespace) |
| `outside_virtual_network.tenant` | [outside_virtual_network.tenant](data-sources--fleet--properties--outside_virtual_network.md#schema-outside_virtual_network--tenant) |
| `outside_virtual_network.uid` | [outside_virtual_network.uid](data-sources--fleet--properties--outside_virtual_network.md#schema-outside_virtual_network--uid) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--fleet--properties--performance_enhancement_mode.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--fleet--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md#section) |
| `sriov_interfaces` | [sriov_interfaces](data-sources--fleet--properties--sriov_interfaces.md#section) |
| `sriov_interfaces.sriov_interface` | [sriov_interfaces.sriov_interface](data-sources--fleet--properties--sriov_interfaces--sriov_interface.md#section) |
| `sriov_interfaces.sriov_interface.interface_name` | [sriov_interfaces.sriov_interface.interface_name](data-sources--fleet--properties--sriov_interfaces--sriov_interface.md#schema-sriov_interfaces--sriov_interface--interface_name) |
| `sriov_interfaces.sriov_interface.number_of_vfio_vfs` | [sriov_interfaces.sriov_interface.number_of_vfio_vfs](data-sources--fleet--properties--sriov_interfaces--sriov_interface.md#schema-sriov_interfaces--sriov_interface--number_of_vfio_vfs) |
| `sriov_interfaces.sriov_interface.number_of_vfs` | [sriov_interfaces.sriov_interface.number_of_vfs](data-sources--fleet--properties--sriov_interfaces--sriov_interface.md#schema-sriov_interfaces--sriov_interface--number_of_vfs) |
| `storage_class_list` | [storage_class_list](data-sources--fleet--properties--storage_class_list.md#section) |
| `storage_class_list.storage_classes` | [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md#section) |
| `storage_class_list.storage_classes.advanced_storage_parameters` | [storage_class_list.storage_classes.advanced_storage_parameters](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--advanced_storage_parameters) |
| `storage_class_list.storage_classes.allow_volume_expansion` | [storage_class_list.storage_classes.allow_volume_expansion](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--allow_volume_expansion) |
| `storage_class_list.storage_classes.custom_storage` | [storage_class_list.storage_classes.custom_storage](data-sources--fleet--properties--storage_class_list--storage_classes--custom_storage.md#section) |
| `storage_class_list.storage_classes.custom_storage.yaml` | [storage_class_list.storage_classes.custom_storage.yaml](data-sources--fleet--properties--storage_class_list--storage_classes--custom_storage.md#schema-storage_class_list--storage_classes--custom_storage--yaml) |
| `storage_class_list.storage_classes.default_storage_class` | [storage_class_list.storage_classes.default_storage_class](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--default_storage_class) |
| `storage_class_list.storage_classes.description_spec` | [storage_class_list.storage_classes.description_spec](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--description_spec) |
| `storage_class_list.storage_classes.hpe_storage` | [storage_class_list.storage_classes.hpe_storage](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#section) |
| `storage_class_list.storage_classes.hpe_storage.allow_mutations` | [storage_class_list.storage_classes.hpe_storage.allow_mutations](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--allow_mutations) |
| `storage_class_list.storage_classes.hpe_storage.allow_overrides` | [storage_class_list.storage_classes.hpe_storage.allow_overrides](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--allow_overrides) |
| `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` | [storage_class_list.storage_classes.hpe_storage.dedupe_enabled](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--dedupe_enabled) |
| `storage_class_list.storage_classes.hpe_storage.description_spec` | [storage_class_list.storage_classes.hpe_storage.description_spec](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--description_spec) |
| `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` | [storage_class_list.storage_classes.hpe_storage.destroy_on_delete](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--destroy_on_delete) |
| `storage_class_list.storage_classes.hpe_storage.encrypted` | [storage_class_list.storage_classes.hpe_storage.encrypted](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--encrypted) |
| `storage_class_list.storage_classes.hpe_storage.folder` | [storage_class_list.storage_classes.hpe_storage.folder](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--folder) |
| `storage_class_list.storage_classes.hpe_storage.limit_iops` | [storage_class_list.storage_classes.hpe_storage.limit_iops](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--limit_iops) |
| `storage_class_list.storage_classes.hpe_storage.limit_mbps` | [storage_class_list.storage_classes.hpe_storage.limit_mbps](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--limit_mbps) |
| `storage_class_list.storage_classes.hpe_storage.performance_policy` | [storage_class_list.storage_classes.hpe_storage.performance_policy](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--performance_policy) |
| `storage_class_list.storage_classes.hpe_storage.pool` | [storage_class_list.storage_classes.hpe_storage.pool](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--pool) |
| `storage_class_list.storage_classes.hpe_storage.protection_template` | [storage_class_list.storage_classes.hpe_storage.protection_template](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--protection_template) |
| `storage_class_list.storage_classes.hpe_storage.secret_name` | [storage_class_list.storage_classes.hpe_storage.secret_name](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--secret_name) |
| `storage_class_list.storage_classes.hpe_storage.secret_namespace` | [storage_class_list.storage_classes.hpe_storage.secret_namespace](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--secret_namespace) |
| `storage_class_list.storage_classes.hpe_storage.sync_on_detach` | [storage_class_list.storage_classes.hpe_storage.sync_on_detach](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--sync_on_detach) |
| `storage_class_list.storage_classes.hpe_storage.thick` | [storage_class_list.storage_classes.hpe_storage.thick](data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md#schema-storage_class_list--storage_classes--hpe_storage--thick) |
| `storage_class_list.storage_classes.netapp_trident` | [storage_class_list.storage_classes.netapp_trident](data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident.md#section) |
| `storage_class_list.storage_classes.netapp_trident.selector` | [storage_class_list.storage_classes.netapp_trident.selector](data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident--selector.md#section) |
| `storage_class_list.storage_classes.netapp_trident.storage_pools` | [storage_class_list.storage_classes.netapp_trident.storage_pools](data-sources--fleet--properties--storage_class_list--storage_classes--netapp_trident.md#schema-storage_class_list--storage_classes--netapp_trident--storage_pools) |
| `storage_class_list.storage_classes.pure_service_orchestrator` | [storage_class_list.storage_classes.pure_service_orchestrator](data-sources--fleet--properties--storage_class_list--storage_classes--pure_service_orchestrator.md#section) |
| `storage_class_list.storage_classes.pure_service_orchestrator.backend` | [storage_class_list.storage_classes.pure_service_orchestrator.backend](data-sources--fleet--properties--storage_class_list--storage_classes--pure_service_orchestrator.md#schema-storage_class_list--storage_classes--pure_service_orchestrator--backend) |
| `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit](data-sources--fleet--properties--storage_class_list--storage_classes--pure_service_orchestrator.md#schema-storage_class_list--storage_classes--pure_service_orchestrator--bandwidth_limit) |
| `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` | [storage_class_list.storage_classes.pure_service_orchestrator.iops_limit](data-sources--fleet--properties--storage_class_list--storage_classes--pure_service_orchestrator.md#schema-storage_class_list--storage_classes--pure_service_orchestrator--iops_limit) |
| `storage_class_list.storage_classes.reclaim_policy` | [storage_class_list.storage_classes.reclaim_policy](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--reclaim_policy) |
| `storage_class_list.storage_classes.storage_class_name` | [storage_class_list.storage_classes.storage_class_name](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--storage_class_name) |
| `storage_class_list.storage_classes.storage_device` | [storage_class_list.storage_classes.storage_device](data-sources--fleet--properties--storage_class_list--storage_classes.md#schema-storage_class_list--storage_classes--storage_device) |
| `storage_device_list` | [storage_device_list](data-sources--fleet--properties--storage_device_list.md#section) |
| `storage_device_list.storage_devices` | [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md#section) |
| `storage_device_list.storage_devices.advanced_advanced_parameters` | [storage_device_list.storage_devices.advanced_advanced_parameters](data-sources--fleet--properties--storage_device_list--storage_devices.md#schema-storage_device_list--storage_devices--advanced_advanced_parameters) |
| `storage_device_list.storage_devices.custom_storage` | [storage_device_list.storage_devices.custom_storage](data-sources--fleet--properties--storage_device_list--storage_devices--custom_storage.md#section) |
| `storage_device_list.storage_devices.hpe_storage` | [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#section) |
| `storage_device_list.storage_devices.hpe_storage.api_server_port` | [storage_device_list.storage_devices.hpe_storage.api_server_port](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#schema-storage_device_list--storage_devices--hpe_storage--api_server_port) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password.md#section) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info--url) |
| `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` | [storage_device_list.storage_devices.hpe_storage.iscsi_chap_user](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_user) |
| `storage_device_list.storage_devices.hpe_storage.password` | [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password.md#section) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--password--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--clear_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--password--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` | [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage--password--clear_secret_info.md#schema-storage_device_list--storage_devices--hpe_storage--password--clear_secret_info--url) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` | [storage_device_list.storage_devices.hpe_storage.storage_server_ip_address](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#schema-storage_device_list--storage_devices--hpe_storage--storage_server_ip_address) |
| `storage_device_list.storage_devices.hpe_storage.storage_server_name` | [storage_device_list.storage_devices.hpe_storage.storage_server_name](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#schema-storage_device_list--storage_devices--hpe_storage--storage_server_name) |
| `storage_device_list.storage_devices.hpe_storage.username` | [storage_device_list.storage_devices.hpe_storage.username](data-sources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md#schema-storage_device_list--storage_devices--hpe_storage--username) |
| `storage_device_list.storage_devices.netapp_trident` | [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_cidrs.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_cidrs.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_cidrs--prefixes) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--backend_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_certificate) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_ip) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--labels) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_aggregate_usage) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_volume_size) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_dns_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_ip) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--nfs_mount_options) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--password--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--region) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--labels) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--adaptive_qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--encryption) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--export_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--no_qos.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--security_style) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--snapshot_dir) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--snapshot_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--snapshot_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--space_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--split_on_clone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--tiering_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--unix_permissions) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--zone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_prefix) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--svm) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--trusted_ca_certificate) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--username) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--adaptive_qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--encryption) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--export_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--no_qos.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--security_style) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--snapshot_dir) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--snapshot_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--snapshot_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--space_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--split_on_clone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--tiering_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--unix_permissions) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_certificate) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_dns_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_ip) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--igroup_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--labels) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_aggregate_usage) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_volume_size) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_dns_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_ip) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--no_chap.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--region) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--labels) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--adaptive_qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--encryption) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--export_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--no_qos.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--security_style) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_dir) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--space_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--split_on_clone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--tiering_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--unix_permissions) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--zone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_driver_name) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_prefix) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--svm) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--trusted_ca_certificate) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info--url) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_username) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_username) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--username) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--adaptive_qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--encryption) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--export_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--no_qos.md#section) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--qos_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--security_style) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_dir) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--space_reserve) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--split_on_clone) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--tiering_policy) |
| `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` | [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults.md#schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--unix_permissions) |
| `storage_device_list.storage_devices.pure_service_orchestrator` | [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_opt) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_type) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_mount_opts) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--disable_preempt_attachments) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--clear_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--clear_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--clear_secret_info--url) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--labels) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--mgmt_dns_name) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--mgmt_ip) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--iscsi_login_timeout) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--san_type) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--export_rules) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info--decryption_provider) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info--location) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info--store_provider) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md#section) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info--provider_ref) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info--url) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--labels) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_dns_name) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_ip) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_dns_name) |
| `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` | [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_ip) |
| `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` | [storage_device_list.storage_devices.pure_service_orchestrator.cluster_id](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_storage_topology) |
| `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` | [storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md#schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_strict_topology) |
| `storage_device_list.storage_devices.storage_device` | [storage_device_list.storage_devices.storage_device](data-sources--fleet--properties--storage_device_list--storage_devices.md#schema-storage_device_list--storage_devices--storage_device) |
| `storage_interface_list` | [storage_interface_list](data-sources--fleet--properties--storage_interface_list.md#section) |
| `storage_interface_list.interfaces` | [storage_interface_list.interfaces](data-sources--fleet--properties--storage_interface_list--interfaces.md#section) |
| `storage_interface_list.interfaces.name` | [storage_interface_list.interfaces.name](data-sources--fleet--properties--storage_interface_list--interfaces.md#schema-storage_interface_list--interfaces--name) |
| `storage_interface_list.interfaces.namespace` | [storage_interface_list.interfaces.namespace](data-sources--fleet--properties--storage_interface_list--interfaces.md#schema-storage_interface_list--interfaces--namespace) |
| `storage_interface_list.interfaces.tenant` | [storage_interface_list.interfaces.tenant](data-sources--fleet--properties--storage_interface_list--interfaces.md#schema-storage_interface_list--interfaces--tenant) |
| `storage_static_routes` | [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md#section) |
| `storage_static_routes.storage_routes` | [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md#section) |
| `storage_static_routes.storage_routes.attrs` | [storage_static_routes.storage_routes.attrs](data-sources--fleet--properties--storage_static_routes--storage_routes.md#schema-storage_static_routes--storage_routes--attrs) |
| `storage_static_routes.storage_routes.labels` | [storage_static_routes.storage_routes.labels](data-sources--fleet--properties--storage_static_routes--storage_routes--labels.md#section) |
| `storage_static_routes.storage_routes.nexthop` | [storage_static_routes.storage_routes.nexthop](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop.md#section) |
| `storage_static_routes.storage_routes.nexthop.interface` | [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#section) |
| `storage_static_routes.storage_routes.nexthop.interface.kind` | [storage_static_routes.storage_routes.nexthop.interface.kind](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#schema-storage_static_routes--storage_routes--nexthop--interface--kind) |
| `storage_static_routes.storage_routes.nexthop.interface.name` | [storage_static_routes.storage_routes.nexthop.interface.name](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#schema-storage_static_routes--storage_routes--nexthop--interface--name) |
| `storage_static_routes.storage_routes.nexthop.interface.namespace` | [storage_static_routes.storage_routes.nexthop.interface.namespace](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#schema-storage_static_routes--storage_routes--nexthop--interface--namespace) |
| `storage_static_routes.storage_routes.nexthop.interface.tenant` | [storage_static_routes.storage_routes.nexthop.interface.tenant](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#schema-storage_static_routes--storage_routes--nexthop--interface--tenant) |
| `storage_static_routes.storage_routes.nexthop.interface.uid` | [storage_static_routes.storage_routes.nexthop.interface.uid](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--interface.md#schema-storage_static_routes--storage_routes--nexthop--interface--uid) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address` | [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv4.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv4.md#schema-storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv6.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv6.md#schema-storage_static_routes--storage_routes--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4.addr](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4.md#schema-storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4--addr) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv6.md#section) |
| `storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr` | [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6.addr](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv6.md#schema-storage_static_routes--storage_routes--nexthop--nexthop_address--ipv6--addr) |
| `storage_static_routes.storage_routes.nexthop.type` | [storage_static_routes.storage_routes.nexthop.type](data-sources--fleet--properties--storage_static_routes--storage_routes--nexthop.md#schema-storage_static_routes--storage_routes--nexthop--type) |
| `storage_static_routes.storage_routes.subnets` | [storage_static_routes.storage_routes.subnets](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets.md#section) |
| `storage_static_routes.storage_routes.subnets.ipv4` | [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md#section) |
| `storage_static_routes.storage_routes.subnets.ipv4.plen` | [storage_static_routes.storage_routes.subnets.ipv4.plen](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md#schema-storage_static_routes--storage_routes--subnets--ipv4--plen) |
| `storage_static_routes.storage_routes.subnets.ipv4.prefix` | [storage_static_routes.storage_routes.subnets.ipv4.prefix](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md#schema-storage_static_routes--storage_routes--subnets--ipv4--prefix) |
| `storage_static_routes.storage_routes.subnets.ipv6` | [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md#section) |
| `storage_static_routes.storage_routes.subnets.ipv6.plen` | [storage_static_routes.storage_routes.subnets.ipv6.plen](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md#schema-storage_static_routes--storage_routes--subnets--ipv6--plen) |
| `storage_static_routes.storage_routes.subnets.ipv6.prefix` | [storage_static_routes.storage_routes.subnets.ipv6.prefix](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md#schema-storage_static_routes--storage_routes--subnets--ipv6--prefix) |
| `usb_policy` | [usb_policy](data-sources--fleet--properties--usb_policy.md#section) |
| `usb_policy.name` | [usb_policy.name](data-sources--fleet--properties--usb_policy.md#schema-usb_policy--name) |
| `usb_policy.namespace` | [usb_policy.namespace](data-sources--fleet--properties--usb_policy.md#schema-usb_policy--namespace) |
| `usb_policy.tenant` | [usb_policy.tenant](data-sources--fleet--properties--usb_policy.md#schema-usb_policy--tenant) |
| `volterra_software_version` | [volterra_software_version](data-sources--fleet--reference.md#schema-volterra_software_version) |

## Next pages

- [allow_all_usb](data-sources--fleet--properties--allow_all_usb.md)
- [blocked_services](data-sources--fleet--properties--blocked_services.md)
- [bond_device_list](data-sources--fleet--properties--bond_device_list.md)
- [dc_cluster_group](data-sources--fleet--properties--dc_cluster_group.md)
- [dc_cluster_group_inside](data-sources--fleet--properties--dc_cluster_group_inside.md)
- [default_config](data-sources--fleet--properties--default_config.md)
- [default_sriov_interface](data-sources--fleet--properties--default_sriov_interface.md)
- [default_storage_class](data-sources--fleet--properties--default_storage_class.md)
- [deny_all_usb](data-sources--fleet--properties--deny_all_usb.md)
- [device_list](data-sources--fleet--properties--device_list.md)
- [disable_gpu](data-sources--fleet--properties--disable_gpu.md)
- [disable_log_anonymization](data-sources--fleet--properties--disable_log_anonymization.md)
- [disable_vm](data-sources--fleet--properties--disable_vm.md)
- [enable_gpu](data-sources--fleet--properties--enable_gpu.md)
- [enable_log_anonymization](data-sources--fleet--properties--enable_log_anonymization.md)
- [enable_vgpu](data-sources--fleet--properties--enable_vgpu.md)
- [enable_vm](data-sources--fleet--properties--enable_vm.md)
- [inside_virtual_network](data-sources--fleet--properties--inside_virtual_network.md)
- [interface_list](data-sources--fleet--properties--interface_list.md)
- [kubernetes_upgrade_drain](data-sources--fleet--properties--kubernetes_upgrade_drain.md)
- [log_receiver](data-sources--fleet--properties--log_receiver.md)
- [logs_streaming_disabled](data-sources--fleet--properties--logs_streaming_disabled.md)
- [network_connectors](data-sources--fleet--properties--network_connectors.md)
- [network_firewall](data-sources--fleet--properties--network_firewall.md)
- [no_bond_devices](data-sources--fleet--properties--no_bond_devices.md)
- [no_dc_cluster_group](data-sources--fleet--properties--no_dc_cluster_group.md)
- [no_storage_device](data-sources--fleet--properties--no_storage_device.md)
- [no_storage_interfaces](data-sources--fleet--properties--no_storage_interfaces.md)
- [no_storage_static_routes](data-sources--fleet--properties--no_storage_static_routes.md)
- [outside_virtual_network](data-sources--fleet--properties--outside_virtual_network.md)
- [performance_enhancement_mode](data-sources--fleet--properties--performance_enhancement_mode.md)
- [sriov_interfaces](data-sources--fleet--properties--sriov_interfaces.md)
- [storage_class_list](data-sources--fleet--properties--storage_class_list.md)
- [storage_device_list](data-sources--fleet--properties--storage_device_list.md)
- [storage_interface_list](data-sources--fleet--properties--storage_interface_list.md)
- [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md)
- [usb_policy](data-sources--fleet--properties--usb_policy.md)
- [xcsh_fleet](../data-sources/fleet.md)
