---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_registration."
xcsh_docs: {"aliases": ["registration"], "body_bytes": 34183, "body_sha256": "sha256:4301402c031aa3f24d1a47c74a777d6956240a73e9583ebc04881ce2a3593be2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra", "xcsh-docs:data-sources:registration:properties:passport"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:reference", "parent_id": "xcsh-docs:data-sources:registration:fundamentals", "path": "documentation/data-sources/registration/properties/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra"], "anchor": "section", "description": "InfraMetadata stores information about instance infrastructure.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["passport"], "anchor": "section", "description": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "document_id": "xcsh-docs:data-sources:registration:properties:passport", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["passport"], "syntax": "attribute", "type": "object"}, {"aliases": ["token"], "anchor": "schema-token", "description": "Token is used for machine and tenant identification.", "document_id": "xcsh-docs:data-sources:registration:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["token"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
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

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/): complete subsection reference.

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

- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/): complete subsection reference.

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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-id) |
| `infra` | [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#section) |
| `infra.availability_zone` | [infra.availability_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--availability_zone) |
| `infra.bond_config` | [infra.bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/#section) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/#schema-infra--bond_config--interfaces) |
| `infra.bond_config.mode` | [infra.bond_config.mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/#schema-infra--bond_config--mode) |
| `infra.bond_config.name` | [infra.bond_config.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/#schema-infra--bond_config--name) |
| `infra.certified_hw` | [infra.certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--certified_hw) |
| `infra.domain` | [infra.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--domain) |
| `infra.hostname` | [infra.hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--hostname) |
| `infra.hugepages` | [infra.hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/#section) |
| `infra.hugepages.free` | [infra.hugepages.free](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/#schema-infra--hugepages--free) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/#schema-infra--hugepages--page_size) |
| `infra.hugepages.total` | [infra.hugepages.total](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/#schema-infra--hugepages--total) |
| `infra.hw_info` | [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/#section) |
| `infra.hw_info.bios` | [infra.hw_info.bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/bios/#section) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/bios/#schema-infra--hw_info--bios--date) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/bios/#schema-infra--hw_info--bios--vendor) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/bios/#schema-infra--hw_info--bios--version) |
| `infra.hw_info.board` | [infra.hw_info.board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#section) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#schema-infra--hw_info--board--asset_tag) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#schema-infra--hw_info--board--name) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#schema-infra--hw_info--board--serial) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#schema-infra--hw_info--board--vendor) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/#schema-infra--hw_info--board--version) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#section) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#schema-infra--hw_info--chassis--asset_tag) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#schema-infra--hw_info--chassis--serial) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#schema-infra--hw_info--chassis--type) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#schema-infra--hw_info--chassis--vendor) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/#schema-infra--hw_info--chassis--version) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#section) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--cache) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--cores) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--cpus) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--model) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--speed) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--threads) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/#schema-infra--hw_info--cpu--vendor) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/#section) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/#schema-infra--hw_info--gpu--cuda_version) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/#schema-infra--hw_info--gpu--driver_version) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/#section) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/#schema-infra--hw_info--gpu--gpu_device--id) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/#schema-infra--hw_info--gpu--gpu_device--processes) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/gpu_device/#schema-infra--hw_info--gpu--gpu_device--product_name) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/kernel/#section) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/kernel/#schema-infra--hw_info--kernel--architecture) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/kernel/#schema-infra--hw_info--kernel--release) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/kernel/#schema-infra--hw_info--kernel--version) |
| `infra.hw_info.memory` | [infra.hw_info.memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/memory/#section) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/memory/#schema-infra--hw_info--memory--size_mb) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/memory/#schema-infra--hw_info--memory--speed) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/memory/#schema-infra--hw_info--memory--type) |
| `infra.hw_info.network` | [infra.hw_info.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#section) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--driver) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--ip_address) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--link_quality) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--link_type) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--mac_address) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--name) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--port) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/#schema-infra--hw_info--network--speed) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/#schema-infra--hw_info--numa_nodes) |
| `infra.hw_info.os` | [infra.hw_info.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#section) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#schema-infra--hw_info--os--architecture) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#schema-infra--hw_info--os--name) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#schema-infra--hw_info--os--release) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#schema-infra--hw_info--os--vendor) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/#schema-infra--hw_info--os--version) |
| `infra.hw_info.product` | [infra.hw_info.product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/#section) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/#schema-infra--hw_info--product--name) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/#schema-infra--hw_info--product--serial) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/#schema-infra--hw_info--product--vendor) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/#schema-infra--hw_info--product--version) |
| `infra.hw_info.storage` | [infra.hw_info.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#section) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--driver) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--model) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--name) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--serial) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--size_gb) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/#schema-infra--hw_info--storage--vendor) |
| `infra.hw_info.usb` | [infra.hw_info.usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#section) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--address) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--b_device_class) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--b_device_protocol) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--b_device_sub_class) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--b_max_packet_size) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--bcd_device) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--bcd_usb) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--bus) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--description_spec) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--i_manufacturer) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--i_product) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--i_serial) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--id_product) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--id_vendor) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--port) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--product_name) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--speed) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--usb_type) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/#schema-infra--hw_info--usb--vendor_name) |
| `infra.instance_id` | [infra.instance_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--instance_id) |
| `infra.interfaces` | [infra.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/interfaces/#section) |
| `infra.internet_proxy` | [infra.internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/#section) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/#schema-infra--internet_proxy--http_proxy) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/#schema-infra--internet_proxy--https_proxy) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/#schema-infra--internet_proxy--no_proxy) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/#schema-infra--internet_proxy--proxy_cacert_url) |
| `infra.is_slo_static` | [infra.is_slo_static](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--is_slo_static) |
| `infra.machine_id` | [infra.machine_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--machine_id) |
| `infra.provider_ref` | [infra.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--provider_ref) |
| `infra.sw_info` | [infra.sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/sw_info/#section) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/sw_info/#schema-infra--sw_info--sw_version) |
| `infra.timestamp` | [infra.timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--timestamp) |
| `infra.zone` | [infra.zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/#schema-infra--zone) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-namespace) |
| `passport` | [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#section) |
| `passport.cluster_name` | [passport.cluster_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--cluster_name) |
| `passport.cluster_size` | [passport.cluster_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--cluster_size) |
| `passport.cluster_type` | [passport.cluster_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--cluster_type) |
| `passport.default_os_version` | [passport.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/default_os_version/#section) |
| `passport.default_sw_version` | [passport.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/default_sw_version/#section) |
| `passport.latitude` | [passport.latitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--latitude) |
| `passport.longitude` | [passport.longitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--longitude) |
| `passport.operating_system_version` | [passport.operating_system_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--operating_system_version) |
| `passport.private_network_name` | [passport.private_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--private_network_name) |
| `passport.volterra_software_version` | [passport.volterra_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/#schema-passport--volterra_software_version) |
| `token` | [token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/#schema-token) |

## Next pages

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/passport/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
