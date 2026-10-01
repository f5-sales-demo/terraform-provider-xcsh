---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22a1f3b7c8a5601cd71d3f3e1dd1bf106e5e989dc492b9d9b9a8c6e85e37d587"></a>

## Property reference — Property reference / 8d6f482b2745 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- Property reference

<a id="canonical-a2590396bfd23b97d899d5eb3a7de0f8b7a922104713b7995bd035b608212abc"></a>

## Direct properties — Property reference / 8d6f482b2745 / 3

<a id="canonical-83f1ad24129440abbd5f746fdf73a6754635b70d18f8ea18884fa7647892cbdc"></a>

<a id="canonical-04a064904b6f4750503581f2de88c9cc7c9ec4591c597213e3c0615062018c4f"></a>

## address property — Property reference / 8d6f482b2745 / 4

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

- [allow_all_usb](resources--voltstack_site--reference--group-003.md#canonical-0c609c579b7b8f47584a039bea90606a48bec15f7b34acca35be08e606d1e931): complete subsection reference.

<a id="canonical-cfad95aa40139be4a2c82ba87605181555d1b404a94cc8087517f1a46e978d78"></a>

<a id="canonical-17901e856dbf370dff78bd4e0c068af79038c3d51b2beffb0f24ea0a5579de67"></a>

## annotations property — Property reference / 8d6f482b2745 / 5

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

- [blocked_services](resources--voltstack_site--reference--group-003.md#canonical-88791dccd7a1bd662a7a8b811db816d13ddb3de65298be7195d033be16bdac73): complete subsection reference.

- [bond_device_list](resources--voltstack_site--reference--group-003.md#canonical-cc4c0b47524d5ee03d8caee7e924c84497cf5cd8c3ce2e9f9076ffe78ee2721c): complete subsection reference.

- [coordinates](resources--voltstack_site--reference--group-003.md#canonical-8cff2240da4a0e65474e320c84f9877bfe5ec8be2661086834ed97869bb9028b): complete subsection reference.

- [custom_dns](resources--voltstack_site--reference--group-003.md#canonical-c46192b03e24734904ebeb4a094c941577239e3018481579880885c3805232f2): complete subsection reference.

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-875fc1888b0d1e9356602920438ee6e95e0b97b44f0f5a6e97f623a33321d8c1): complete subsection reference.

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3): complete subsection reference.

- [default_blocked_services](resources--voltstack_site--reference--group-008.md#canonical-9c7aa305f2b7aefe8a0fa849f813608d94870678f3ce3c6e4bbd6f1d8c322462): complete subsection reference.

- [default_network_config](resources--voltstack_site--reference--group-008.md#canonical-16760fb2699e8de90ea1777349cbc988e39c3cd99c0776eb893327576b74a787): complete subsection reference.

- [default_sriov_interface](resources--voltstack_site--reference--group-008.md#canonical-3f9a876eae72912459859917f4f83155dc5f3166eedd2432f05b48ad0350cff5): complete subsection reference.

- [default_storage_config](resources--voltstack_site--reference--group-008.md#canonical-2dc1fe90a2f31c9e060c766a43fd8e42dbff1fb1558962379ea1f37c7eefffd2): complete subsection reference.

- [deny_all_usb](resources--voltstack_site--reference--group-008.md#canonical-05ff4a257b87219168061775c8e9079bf4b1eec1ee3d319d1568fd80421eddd8): complete subsection reference.

<a id="canonical-a47a611416300c1e935864f20431120189ede8d0c69c208668241cabdc544d6c"></a>

<a id="canonical-9f300bb110efc75780dd64d85bf1a5ff466ad27c3d63e305dda481030a205aad"></a>

## description property — Property reference / 8d6f482b2745 / 6

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

<a id="canonical-38c1113e93a3aca8267d6ddc10c882d1bc7a3cb29ff2a71743455dbbe237cb1a"></a>

<a id="canonical-2e4bc281601d130b71fa8a8309e72782b63c29817dcf069944b4608c4168d911"></a>

## disable property — Property reference / 8d6f482b2745 / 7

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

- [disable_gpu](resources--voltstack_site--reference--group-008.md#canonical-8149b6f14e2befec8f8c1fd2b131fe6a2dcdb1b9b7dc0f595905cd1ba5e9a418): complete subsection reference.

- [disable_vm](resources--voltstack_site--reference--group-008.md#canonical-b383f3c053581f48bc6990012676b4f4780feb328d7908bd262404edcf5dba3d): complete subsection reference.

- [enable_gpu](resources--voltstack_site--reference--group-008.md#canonical-d43202a2253010a9d1d7bfe20f370d399532c015052461bf8ae6d4eb218601d7): complete subsection reference.

- [enable_vgpu](resources--voltstack_site--reference--group-008.md#canonical-29ba2a83c6e25fe2dfe15a0afa21faf5f1753feda6c22fb30112e405e0b74924): complete subsection reference.

- [enable_vm](resources--voltstack_site--reference--group-008.md#canonical-a34ba3ffb822aafbc24178b2e50a1661ea64f10546fb642900eaf65d1b75d838): complete subsection reference.

<a id="canonical-dddb7cb7df40fddb68fb5ecf64f97fc991946eb5f492134306e15d2ba2fdb98d"></a>

<a id="canonical-7f033459282d36724a6ea815d3d35a99a5a1d1414f78b43ccddc33fe1694381d"></a>

## id property — Property reference / 8d6f482b2745 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster](resources--voltstack_site--reference--group-009.md#canonical-0fb68b89f5f7a81145ed1211b5db6d166121e37f3e7e50e5ac8e2e05951d60ac): complete subsection reference.

- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-39b4acd2c70fbaf081efb6ea65069207caa0eb11c0a9136651ee1681cde07815): complete subsection reference.

<a id="canonical-3ae126972c9cbeedbd4445382bb98d17a06f824cab1c29bc658b6070374dc1bc"></a>

<a id="canonical-4683b0ed2ef8e419f1cac0da505222656676dab3cb8e29ecc9349c307bf1670b"></a>

## labels property — Property reference / 8d6f482b2745 / 9

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

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-5ee3efeca95fecf8bce448308b6b448b4a801fc5f36a5a715c7240f45f94c151): complete subsection reference.

- [log_receiver](resources--voltstack_site--reference--group-009.md#canonical-f427be209dfb9b1b092fec98b6fc07c003c09f3ccd6188933e00188986bba470): complete subsection reference.

- [logs_streaming_disabled](resources--voltstack_site--reference--group-009.md#canonical-870e433337c76cd152ff0fdeb29b4d50767116fb7e6727361a766d0987cd9cee): complete subsection reference.

- [master_node_configuration](resources--voltstack_site--reference--group-009.md#canonical-ad1c28b9743413718a3b177c526c28dd11e9734d12ec0a1899837c724ececc4f): complete subsection reference.

<a id="canonical-04d379d8920219306e410b5f85efa8f30714aad38f65d7733a72352c78322b04"></a>

<a id="canonical-f9372dfb20655a41ec7e6d24a05c14cf1b5646a2cf084e18fb189f6c6edb91c5"></a>

## name property — Property reference / 8d6f482b2745 / 10

Type: `"string"`. Required.

Name of the Voltstack Site. Must be unique within the namespace.

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

<a id="canonical-b67fdfff57979f12617bbd160cb7a53d9065b4b44ca8863588929b3fd94d3486"></a>

<a id="canonical-9a719c48e19c2147c83f5d49309036100d21908e2cd69941203686d2d6d1da78"></a>

## namespace property — Property reference / 8d6f482b2745 / 11

Type: `"string"`. Required.

Namespace where the Voltstack Site is created.

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

- [no_bond_devices](resources--voltstack_site--reference--group-009.md#canonical-15682f23e72c547d6211938bb42701baedc08ac036e147a1efa0950d69f09bbd): complete subsection reference.

- [no_k8s_cluster](resources--voltstack_site--reference--group-009.md#canonical-54d6d51504ff10e4b354b972db67901b7e303b239c786fda5988906e6b72c9eb): complete subsection reference.

- [no_local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-6671aef9498eb983a7034a3bc3ed9724f6fe735664ca5250753bef797a70f967): complete subsection reference.

- [offline_survivability_mode](resources--voltstack_site--reference--group-009.md#canonical-186e14d0d96fbfdebbd08044ee03b5772cde1a87e273a1f99fcc0fe6163d5171): complete subsection reference.

- [os](resources--voltstack_site--reference--group-010.md#canonical-acafd9e483552322116dcf73798aa1ebffd5871a6d89ddb562916c57e59f1683): complete subsection reference.

- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-e288685eed32403a1b38a053cd958daafe67a3c468f5a4c75bbb371860409f6f): complete subsection reference.

- [sw](resources--voltstack_site--reference--group-010.md#canonical-745052ba3f940f79c0a1e2dd6e9d5edafe93562c98b4f624e7b3234156449238): complete subsection reference.

- [timeouts](resources--voltstack_site--reference--group-010.md#canonical-5c6a3b09311bf2ef7cafbadfd87aa53d9413906a85e90eac3b032f350aa4daaf): complete subsection reference.

- [usb_policy](resources--voltstack_site--reference--group-010.md#canonical-fed8e4ab7457c4c6764b430be7605274aed28a3dfd7652d7cb4f1ce547934f33): complete subsection reference.

<a id="canonical-3fa79bc72100168fd972ac3ac5bf3b2b29bdb7b1f1dd61ce925d99c908151399"></a>

<a id="canonical-0acb047502b5c9222508a2064d49e0b2cb20f0f471576ecadddc8fb47cef37fe"></a>

## volterra_certified_hw property — Property reference / 8d6f482b2745 / 12

Type: `"string"`. Required.

Name for generic server certified hardware to form this App Stack site.

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

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-c19033b27a37f0dcddc37a2716822743c64449717b58b7d4d99a91468d93821e): complete subsection reference.

<a id="canonical-faa671f7102a61f38dca98e0cb34ab0a85b64e74b08119fedd5740e812d6a2fc"></a>

<a id="canonical-05231fa0b7992c5a074665f2fefb6669179fce9ac4589804c25e9b01f08b4a68"></a>

## worker_nodes property — Property reference / 8d6f482b2745 / 13

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
