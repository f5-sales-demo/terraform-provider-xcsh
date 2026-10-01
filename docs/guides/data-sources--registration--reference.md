---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 27276, "body_sha256": "sha256:66a1906420abc5a04395b7c45dcb42dd79de36a5cdc4751d44693c42dd5a16e2", "canonical_id": "xcsh-docs:data-sources:registration:reference", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra", "xcsh-docs:data-sources:registration:properties:passport"], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:reference", "parent_id": "xcsh-docs:data-sources:registration:fundamentals", "path": "docs/guides/data-sources--registration--reference.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
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

Description of the Registration.

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

- [infra](data-sources--registration--properties--infra.md): complete subsection reference.

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

Name of the Registration.

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

Namespace where the Registration exists.

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

- [passport](data-sources--registration--properties--passport.md): complete subsection reference.

<a id="schema-token"></a>

### token property

Type: `"string"`. Computed.

Token is used for machine and tenant identification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "security",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20
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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--registration--reference.md#schema-annotations) |
| `description` | [description](data-sources--registration--reference.md#schema-description) |
| `id` | [id](data-sources--registration--reference.md#schema-id) |
| `infra` | [infra](data-sources--registration--properties--infra.md#section) |
| `infra.availability_zone` | [infra.availability_zone](data-sources--registration--properties--infra.md#schema-infra--availability_zone) |
| `infra.bond_config` | [infra.bond_config](data-sources--registration--properties--infra--bond_config.md#section) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](data-sources--registration--properties--infra--bond_config.md#schema-infra--bond_config--interfaces) |
| `infra.bond_config.mode` | [infra.bond_config.mode](data-sources--registration--properties--infra--bond_config.md#schema-infra--bond_config--mode) |
| `infra.bond_config.name` | [infra.bond_config.name](data-sources--registration--properties--infra--bond_config.md#schema-infra--bond_config--name) |
| `infra.certified_hw` | [infra.certified_hw](data-sources--registration--properties--infra.md#schema-infra--certified_hw) |
| `infra.domain` | [infra.domain](data-sources--registration--properties--infra.md#schema-infra--domain) |
| `infra.hostname` | [infra.hostname](data-sources--registration--properties--infra.md#schema-infra--hostname) |
| `infra.hugepages` | [infra.hugepages](data-sources--registration--properties--infra--hugepages.md#section) |
| `infra.hugepages.free` | [infra.hugepages.free](data-sources--registration--properties--infra--hugepages.md#schema-infra--hugepages--free) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](data-sources--registration--properties--infra--hugepages.md#schema-infra--hugepages--page_size) |
| `infra.hugepages.total` | [infra.hugepages.total](data-sources--registration--properties--infra--hugepages.md#schema-infra--hugepages--total) |
| `infra.hw_info` | [infra.hw_info](data-sources--registration--properties--infra--hw_info.md#section) |
| `infra.hw_info.bios` | [infra.hw_info.bios](data-sources--registration--properties--infra--hw_info--bios.md#section) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](data-sources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--date) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](data-sources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--vendor) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](data-sources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--version) |
| `infra.hw_info.board` | [infra.hw_info.board](data-sources--registration--properties--infra--hw_info--board.md#section) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](data-sources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--asset_tag) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](data-sources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--name) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](data-sources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--serial) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](data-sources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--vendor) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](data-sources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--version) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](data-sources--registration--properties--infra--hw_info--chassis.md#section) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](data-sources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--asset_tag) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](data-sources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--serial) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](data-sources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--type) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](data-sources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--vendor) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](data-sources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--version) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](data-sources--registration--properties--infra--hw_info--cpu.md#section) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cache) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cores) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cpus) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--model) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--speed) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--threads) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](data-sources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--vendor) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](data-sources--registration--properties--infra--hw_info--gpu.md#section) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](data-sources--registration--properties--infra--hw_info--gpu.md#schema-infra--hw_info--gpu--cuda_version) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](data-sources--registration--properties--infra--hw_info--gpu.md#schema-infra--hw_info--gpu--driver_version) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md#section) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--id) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--processes) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](data-sources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--product_name) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](data-sources--registration--properties--infra--hw_info--kernel.md#section) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](data-sources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--architecture) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](data-sources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--release) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](data-sources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--version) |
| `infra.hw_info.memory` | [infra.hw_info.memory](data-sources--registration--properties--infra--hw_info--memory.md#section) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](data-sources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--size_mb) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](data-sources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--speed) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](data-sources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--type) |
| `infra.hw_info.network` | [infra.hw_info.network](data-sources--registration--properties--infra--hw_info--network.md#section) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--driver) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--ip_address) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--link_quality) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--link_type) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--mac_address) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--name) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--port) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](data-sources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--speed) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](data-sources--registration--properties--infra--hw_info.md#schema-infra--hw_info--numa_nodes) |
| `infra.hw_info.os` | [infra.hw_info.os](data-sources--registration--properties--infra--hw_info--os.md#section) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](data-sources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--architecture) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](data-sources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--name) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](data-sources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--release) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](data-sources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--vendor) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](data-sources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--version) |
| `infra.hw_info.product` | [infra.hw_info.product](data-sources--registration--properties--infra--hw_info--product.md#section) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](data-sources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--name) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](data-sources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--serial) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](data-sources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--vendor) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](data-sources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--version) |
| `infra.hw_info.storage` | [infra.hw_info.storage](data-sources--registration--properties--infra--hw_info--storage.md#section) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--driver) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--model) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--name) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--serial) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--size_gb) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](data-sources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--vendor) |
| `infra.hw_info.usb` | [infra.hw_info.usb](data-sources--registration--properties--infra--hw_info--usb.md#section) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--address) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_class) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_protocol) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_sub_class) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_max_packet_size) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bcd_device) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bcd_usb) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bus) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--description_spec) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_manufacturer) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_product) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_serial) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--id_product) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--id_vendor) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--port) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--product_name) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--speed) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--usb_type) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](data-sources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--vendor_name) |
| `infra.instance_id` | [infra.instance_id](data-sources--registration--properties--infra.md#schema-infra--instance_id) |
| `infra.interfaces` | [infra.interfaces](data-sources--registration--properties--infra--interfaces.md#section) |
| `infra.internet_proxy` | [infra.internet_proxy](data-sources--registration--properties--infra--internet_proxy.md#section) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](data-sources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--http_proxy) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](data-sources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--https_proxy) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](data-sources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--no_proxy) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](data-sources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--proxy_cacert_url) |
| `infra.is_slo_static` | [infra.is_slo_static](data-sources--registration--properties--infra.md#schema-infra--is_slo_static) |
| `infra.machine_id` | [infra.machine_id](data-sources--registration--properties--infra.md#schema-infra--machine_id) |
| `infra.provider_ref` | [infra.provider_ref](data-sources--registration--properties--infra.md#schema-infra--provider_ref) |
| `infra.sw_info` | [infra.sw_info](data-sources--registration--properties--infra--sw_info.md#section) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](data-sources--registration--properties--infra--sw_info.md#schema-infra--sw_info--sw_version) |
| `infra.timestamp` | [infra.timestamp](data-sources--registration--properties--infra.md#schema-infra--timestamp) |
| `infra.zone` | [infra.zone](data-sources--registration--properties--infra.md#schema-infra--zone) |
| `labels` | [labels](data-sources--registration--reference.md#schema-labels) |
| `name` | [name](data-sources--registration--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--registration--reference.md#schema-namespace) |
| `passport` | [passport](data-sources--registration--properties--passport.md#section) |
| `passport.cluster_name` | [passport.cluster_name](data-sources--registration--properties--passport.md#schema-passport--cluster_name) |
| `passport.cluster_size` | [passport.cluster_size](data-sources--registration--properties--passport.md#schema-passport--cluster_size) |
| `passport.cluster_type` | [passport.cluster_type](data-sources--registration--properties--passport.md#schema-passport--cluster_type) |
| `passport.default_os_version` | [passport.default_os_version](data-sources--registration--properties--passport--default_os_version.md#section) |
| `passport.default_sw_version` | [passport.default_sw_version](data-sources--registration--properties--passport--default_sw_version.md#section) |
| `passport.latitude` | [passport.latitude](data-sources--registration--properties--passport.md#schema-passport--latitude) |
| `passport.longitude` | [passport.longitude](data-sources--registration--properties--passport.md#schema-passport--longitude) |
| `passport.operating_system_version` | [passport.operating_system_version](data-sources--registration--properties--passport.md#schema-passport--operating_system_version) |
| `passport.private_network_name` | [passport.private_network_name](data-sources--registration--properties--passport.md#schema-passport--private_network_name) |
| `passport.volterra_software_version` | [passport.volterra_software_version](data-sources--registration--properties--passport.md#schema-passport--volterra_software_version) |
| `token` | [token](data-sources--registration--reference.md#schema-token) |

## Next pages

- [infra](data-sources--registration--properties--infra.md)
- [passport](data-sources--registration--properties--passport.md)
- [xcsh_registration](../data-sources/registration.md)
