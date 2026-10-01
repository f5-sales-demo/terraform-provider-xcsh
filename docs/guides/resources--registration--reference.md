---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 28461, "body_sha256": "sha256:f481c3911dbcf78be13ac05076fced3ffa9566e4dd533f2639fc9600d94cd37a", "canonical_id": "xcsh-docs:resources:registration:reference", "child_ids": ["xcsh-docs:resources:registration:properties:infra", "xcsh-docs:resources:registration:properties:passport", "xcsh-docs:resources:registration:properties:timeouts"], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:reference", "parent_id": "xcsh-docs:resources:registration:fundamentals", "path": "docs/guides/resources--registration--reference.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_registration](../resources/registration.md)
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

- [infra](resources--registration--properties--infra.md): complete subsection reference.

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

Name of the Registration. Must be unique within the namespace.

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

Namespace where the Registration is created.

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

- [passport](resources--registration--properties--passport.md): complete subsection reference.

- [timeouts](resources--registration--properties--timeouts.md): complete subsection reference.

<a id="schema-token"></a>

### token property

Type: `"string"`. Required.

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
| `annotations` | [annotations](resources--registration--reference.md#schema-annotations) |
| `description` | [description](resources--registration--reference.md#schema-description) |
| `disable` | [disable](resources--registration--reference.md#schema-disable) |
| `id` | [id](resources--registration--reference.md#schema-id) |
| `infra` | [infra](resources--registration--properties--infra.md#section) |
| `infra.availability_zone` | [infra.availability_zone](resources--registration--properties--infra.md#schema-infra--availability_zone) |
| `infra.bond_config` | [infra.bond_config](resources--registration--properties--infra--bond_config.md#section) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](resources--registration--properties--infra--bond_config.md#schema-infra--bond_config--interfaces) |
| `infra.bond_config.mode` | [infra.bond_config.mode](resources--registration--properties--infra--bond_config.md#schema-infra--bond_config--mode) |
| `infra.bond_config.name` | [infra.bond_config.name](resources--registration--properties--infra--bond_config.md#schema-infra--bond_config--name) |
| `infra.certified_hw` | [infra.certified_hw](resources--registration--properties--infra.md#schema-infra--certified_hw) |
| `infra.domain` | [infra.domain](resources--registration--properties--infra.md#schema-infra--domain) |
| `infra.hostname` | [infra.hostname](resources--registration--properties--infra.md#schema-infra--hostname) |
| `infra.hugepages` | [infra.hugepages](resources--registration--properties--infra--hugepages.md#section) |
| `infra.hugepages.free` | [infra.hugepages.free](resources--registration--properties--infra--hugepages.md#schema-infra--hugepages--free) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](resources--registration--properties--infra--hugepages.md#schema-infra--hugepages--page_size) |
| `infra.hugepages.total` | [infra.hugepages.total](resources--registration--properties--infra--hugepages.md#schema-infra--hugepages--total) |
| `infra.hw_info` | [infra.hw_info](resources--registration--properties--infra--hw_info.md#section) |
| `infra.hw_info.bios` | [infra.hw_info.bios](resources--registration--properties--infra--hw_info--bios.md#section) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](resources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--date) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](resources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--vendor) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](resources--registration--properties--infra--hw_info--bios.md#schema-infra--hw_info--bios--version) |
| `infra.hw_info.board` | [infra.hw_info.board](resources--registration--properties--infra--hw_info--board.md#section) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](resources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--asset_tag) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](resources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--name) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](resources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--serial) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](resources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--vendor) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](resources--registration--properties--infra--hw_info--board.md#schema-infra--hw_info--board--version) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](resources--registration--properties--infra--hw_info--chassis.md#section) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](resources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--asset_tag) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](resources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--serial) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](resources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--type) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](resources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--vendor) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](resources--registration--properties--infra--hw_info--chassis.md#schema-infra--hw_info--chassis--version) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](resources--registration--properties--infra--hw_info--cpu.md#section) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cache) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cores) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--cpus) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--model) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--speed) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--threads) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](resources--registration--properties--infra--hw_info--cpu.md#schema-infra--hw_info--cpu--vendor) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](resources--registration--properties--infra--hw_info--gpu.md#section) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](resources--registration--properties--infra--hw_info--gpu.md#schema-infra--hw_info--gpu--cuda_version) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](resources--registration--properties--infra--hw_info--gpu.md#schema-infra--hw_info--gpu--driver_version) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](resources--registration--properties--infra--hw_info--gpu--gpu_device.md#section) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](resources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--id) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](resources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--processes) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](resources--registration--properties--infra--hw_info--gpu--gpu_device.md#schema-infra--hw_info--gpu--gpu_device--product_name) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](resources--registration--properties--infra--hw_info--kernel.md#section) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](resources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--architecture) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](resources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--release) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](resources--registration--properties--infra--hw_info--kernel.md#schema-infra--hw_info--kernel--version) |
| `infra.hw_info.memory` | [infra.hw_info.memory](resources--registration--properties--infra--hw_info--memory.md#section) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](resources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--size_mb) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](resources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--speed) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](resources--registration--properties--infra--hw_info--memory.md#schema-infra--hw_info--memory--type) |
| `infra.hw_info.network` | [infra.hw_info.network](resources--registration--properties--infra--hw_info--network.md#section) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--driver) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--ip_address) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--link_quality) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--link_type) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--mac_address) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--name) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--port) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](resources--registration--properties--infra--hw_info--network.md#schema-infra--hw_info--network--speed) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](resources--registration--properties--infra--hw_info.md#schema-infra--hw_info--numa_nodes) |
| `infra.hw_info.os` | [infra.hw_info.os](resources--registration--properties--infra--hw_info--os.md#section) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](resources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--architecture) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](resources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--name) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](resources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--release) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](resources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--vendor) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](resources--registration--properties--infra--hw_info--os.md#schema-infra--hw_info--os--version) |
| `infra.hw_info.product` | [infra.hw_info.product](resources--registration--properties--infra--hw_info--product.md#section) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](resources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--name) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](resources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--serial) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](resources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--vendor) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](resources--registration--properties--infra--hw_info--product.md#schema-infra--hw_info--product--version) |
| `infra.hw_info.storage` | [infra.hw_info.storage](resources--registration--properties--infra--hw_info--storage.md#section) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--driver) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--model) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--name) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--serial) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--size_gb) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](resources--registration--properties--infra--hw_info--storage.md#schema-infra--hw_info--storage--vendor) |
| `infra.hw_info.usb` | [infra.hw_info.usb](resources--registration--properties--infra--hw_info--usb.md#section) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--address) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_class) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_protocol) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_device_sub_class) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--b_max_packet_size) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bcd_device) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bcd_usb) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--bus) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--description_spec) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_manufacturer) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_product) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--i_serial) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--id_product) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--id_vendor) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--port) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--product_name) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--speed) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--usb_type) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](resources--registration--properties--infra--hw_info--usb.md#schema-infra--hw_info--usb--vendor_name) |
| `infra.instance_id` | [infra.instance_id](resources--registration--properties--infra.md#schema-infra--instance_id) |
| `infra.interfaces` | [infra.interfaces](resources--registration--properties--infra--interfaces.md#section) |
| `infra.internet_proxy` | [infra.internet_proxy](resources--registration--properties--infra--internet_proxy.md#section) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](resources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--http_proxy) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](resources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--https_proxy) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](resources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--no_proxy) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](resources--registration--properties--infra--internet_proxy.md#schema-infra--internet_proxy--proxy_cacert_url) |
| `infra.is_slo_static` | [infra.is_slo_static](resources--registration--properties--infra.md#schema-infra--is_slo_static) |
| `infra.machine_id` | [infra.machine_id](resources--registration--properties--infra.md#schema-infra--machine_id) |
| `infra.provider_ref` | [infra.provider_ref](resources--registration--properties--infra.md#schema-infra--provider_ref) |
| `infra.sw_info` | [infra.sw_info](resources--registration--properties--infra--sw_info.md#section) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](resources--registration--properties--infra--sw_info.md#schema-infra--sw_info--sw_version) |
| `infra.timestamp` | [infra.timestamp](resources--registration--properties--infra.md#schema-infra--timestamp) |
| `infra.zone` | [infra.zone](resources--registration--properties--infra.md#schema-infra--zone) |
| `labels` | [labels](resources--registration--reference.md#schema-labels) |
| `name` | [name](resources--registration--reference.md#schema-name) |
| `namespace` | [namespace](resources--registration--reference.md#schema-namespace) |
| `passport` | [passport](resources--registration--properties--passport.md#section) |
| `passport.cluster_name` | [passport.cluster_name](resources--registration--properties--passport.md#schema-passport--cluster_name) |
| `passport.cluster_size` | [passport.cluster_size](resources--registration--properties--passport.md#schema-passport--cluster_size) |
| `passport.cluster_type` | [passport.cluster_type](resources--registration--properties--passport.md#schema-passport--cluster_type) |
| `passport.default_os_version` | [passport.default_os_version](resources--registration--properties--passport--default_os_version.md#section) |
| `passport.default_sw_version` | [passport.default_sw_version](resources--registration--properties--passport--default_sw_version.md#section) |
| `passport.latitude` | [passport.latitude](resources--registration--properties--passport.md#schema-passport--latitude) |
| `passport.longitude` | [passport.longitude](resources--registration--properties--passport.md#schema-passport--longitude) |
| `passport.operating_system_version` | [passport.operating_system_version](resources--registration--properties--passport.md#schema-passport--operating_system_version) |
| `passport.private_network_name` | [passport.private_network_name](resources--registration--properties--passport.md#schema-passport--private_network_name) |
| `passport.volterra_software_version` | [passport.volterra_software_version](resources--registration--properties--passport.md#schema-passport--volterra_software_version) |
| `timeouts` | [timeouts](resources--registration--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--registration--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--registration--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--registration--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--registration--properties--timeouts.md#schema-timeouts--update) |
| `token` | [token](resources--registration--reference.md#schema-token) |

## Next pages

- [infra](resources--registration--properties--infra.md)
- [passport](resources--registration--properties--passport.md)
- [timeouts](resources--registration--properties--timeouts.md)
- [xcsh_registration](../resources/registration.md)
