---
page_title: "xcsh_usb_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy reference."
---

# xcsh_usb_policy reference

<a id="canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71bb8ccfd7ea38b3938e42f9b51c58e469ea920891174a8fa75f375ce9140113"></a>

## Property reference — Property reference / 08da30d5c1f1 / 2

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
- Property reference

<a id="canonical-2c844109c3c2f8b3883922700c6d3ea2dd9f85c899e6f6a63206109acabfb566"></a>

## Direct properties — Property reference / 08da30d5c1f1 / 3

- [allowed_devices](resources--usb_policy--reference--group-001.md#canonical-f8b7c37e954d3ea566ecf987149dcc17091b60b06a9a9e970b7fcaadba73db28): complete subsection reference.

<a id="canonical-5a48b4764a0f162b4fcb7ba1616cd4b59cf39290ecddde1c3841a84135c6cf68"></a>

<a id="canonical-e0567e33eb8ea3e79bc537261b632a0a9c816ced28ccd95db48a87e0bb6555a5"></a>

## annotations property — Property reference / 08da30d5c1f1 / 4

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

<a id="canonical-237c5ec96da00d3b15a41d3de7d93c61369030d2b029235e4119d1509481bfbd"></a>

<a id="canonical-ebe30825348404a9924059cfbea1bca4105a7654cd63595fbe27053e8d775e4c"></a>

## description property — Property reference / 08da30d5c1f1 / 5

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

<a id="canonical-fb10a154493ad9feaa5398e404791a97f509543385d5a015932268f2bf591b97"></a>

<a id="canonical-d3c6dd86ed02883f53a3e3250d78276e0bee92ac70e976aa13576e43189ac65b"></a>

## disable property — Property reference / 08da30d5c1f1 / 6

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

<a id="canonical-c4563d4b4e2b55abcf9cb5e0c9c9c0317aba1114ac3c977eee4ded8a4e5f1074"></a>

<a id="canonical-ad5e9e96e4e448cb874c6dd0ff368f8c99207279eee13790176964866df48f81"></a>

## id property — Property reference / 08da30d5c1f1 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6db5ca4c03d6d3877da576b19a08238a7cf5bbc5bfdbb4f4b5a9067442041790"></a>

<a id="canonical-941c1030c9cac18ae4833fdc7ee47bd51392aa7930190efdff0373c1382c3460"></a>

## labels property — Property reference / 08da30d5c1f1 / 8

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

<a id="canonical-9ef732a3ec5611b68a0ec6ffdae69526e91f5d45240595bcf8c63e8b154a0e91"></a>

<a id="canonical-94e2778a69249c82472a191fb5531d5d3f862567bb3d081f8d23b1a07306772a"></a>

## name property — Property reference / 08da30d5c1f1 / 9

Type: `"string"`. Required.

Name of the Usb Policy. Must be unique within the namespace.

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

<a id="canonical-a193856d2c5bce5504f269ecb1f7e7686ebebb8b01d15addcca874c56435ff71"></a>

<a id="canonical-18d7477807d5aca98bf7dc41c49897af12b6adb00819c61310e20733849a47b2"></a>

## namespace property — Property reference / 08da30d5c1f1 / 10

Type: `"string"`. Required.

Namespace where the Usb Policy is created.

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

- [timeouts](resources--usb_policy--reference--group-001.md#canonical-f075229b428c3260403de90e8cebf597503f4637a1d8124c4342a95e03114316): complete subsection reference.

<a id="canonical-daf9c07facf1507bab8d70a3675f7758bc70d0ea2964de8ad046b0db5acfc2c3"></a>

## All schema paths — Property reference / 08da30d5c1f1 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_devices` | [allowed_devices](resources--usb_policy--reference--group-001.md#canonical-e442a1a93c748b0c041f8e3f96ccd6470d7020396ea248ff012ab52a50de31bb) |
| `allowed_devices.b_device_class` | [allowed_devices.b_device_class](resources--usb_policy--reference--group-001.md#canonical-757e900e8fb2afca5dea8556d09d021ff3af5d47ceb52c8f1cd1fb2bf75b9ffc) |
| `allowed_devices.b_device_protocol` | [allowed_devices.b_device_protocol](resources--usb_policy--reference--group-001.md#canonical-880a9813ed96012670a89a52bfd3fcb46220135b5913ec204da8db2de658774a) |
| `allowed_devices.b_device_sub_class` | [allowed_devices.b_device_sub_class](resources--usb_policy--reference--group-001.md#canonical-0c538404d2c60f73a68ae9f2450439c1a79dd3f4280a368e7d94cad21ed33877) |
| `allowed_devices.i_serial` | [allowed_devices.i_serial](resources--usb_policy--reference--group-001.md#canonical-73e90cbd4a507ae2ab15fd8883b7fa34c0bdc967af9b41b8e75e6c4c4f046ec7) |
| `allowed_devices.id_product` | [allowed_devices.id_product](resources--usb_policy--reference--group-001.md#canonical-c14b44f2c7e22c4f83cf09bc23a37c166402890f2d2862d3a3faef628a7e40be) |
| `allowed_devices.id_vendor` | [allowed_devices.id_vendor](resources--usb_policy--reference--group-001.md#canonical-3770d4d230ad01d7881ab2637922cb99e26f6d015f4d831cdbd7c37b73c23859) |
| `annotations` | [annotations](resources--usb_policy--reference--group-001.md#canonical-5a48b4764a0f162b4fcb7ba1616cd4b59cf39290ecddde1c3841a84135c6cf68) |
| `description` | [description](resources--usb_policy--reference--group-001.md#canonical-237c5ec96da00d3b15a41d3de7d93c61369030d2b029235e4119d1509481bfbd) |
| `disable` | [disable](resources--usb_policy--reference--group-001.md#canonical-fb10a154493ad9feaa5398e404791a97f509543385d5a015932268f2bf591b97) |
| `id` | [id](resources--usb_policy--reference--group-001.md#canonical-c4563d4b4e2b55abcf9cb5e0c9c9c0317aba1114ac3c977eee4ded8a4e5f1074) |
| `labels` | [labels](resources--usb_policy--reference--group-001.md#canonical-6db5ca4c03d6d3877da576b19a08238a7cf5bbc5bfdbb4f4b5a9067442041790) |
| `name` | [name](resources--usb_policy--reference--group-001.md#canonical-9ef732a3ec5611b68a0ec6ffdae69526e91f5d45240595bcf8c63e8b154a0e91) |
| `namespace` | [namespace](resources--usb_policy--reference--group-001.md#canonical-a193856d2c5bce5504f269ecb1f7e7686ebebb8b01d15addcca874c56435ff71) |
| `timeouts` | [timeouts](resources--usb_policy--reference--group-001.md#canonical-44565af17189128f8978d6efee251a444a225eb3dffdeeb969c8cd32f59cf75e) |
| `timeouts.create` | [timeouts.create](resources--usb_policy--reference--group-001.md#canonical-ca60a7ab9c2327e70847c4c7d7ba303bf207fa07d7d89da8453bb6a93bd106e6) |
| `timeouts.delete` | [timeouts.delete](resources--usb_policy--reference--group-001.md#canonical-264e3600116da2e526835f665f2c1efc6a19adcaf52ea0c28725f0f24ffef5f1) |
| `timeouts.read` | [timeouts.read](resources--usb_policy--reference--group-001.md#canonical-cd5b272b85e34b52d74450a1bf1041088649f67556404d1abe9b4c679aad92c8) |
| `timeouts.update` | [timeouts.update](resources--usb_policy--reference--group-001.md#canonical-d1dad712f1480e967ffc778b77068a1da4d607fddb805f15649eafb72208d861) |

<a id="canonical-bd255597356d8980f512ca35436143574da4365b9e7edaaa2ad875f75a054e0f"></a>

## Next pages — Property reference / 08da30d5c1f1 / 12

- [allowed_devices](resources--usb_policy--reference--group-001.md#canonical-f8b7c37e954d3ea566ecf987149dcc17091b60b06a9a9e970b7fcaadba73db28)
- [timeouts](resources--usb_policy--reference--group-001.md#canonical-f075229b428c3260403de90e8cebf597503f4637a1d8124c4342a95e03114316)
- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)

<a id="canonical-f8b7c37e954d3ea566ecf987149dcc17091b60b06a9a9e970b7fcaadba73db28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f299c4b2f6d8575ad037fa27a16928a68c172f07e090f344e0e76f8f11a7f1"></a>

## allowed_devices — allowed_devices / 4a0113e5808f / 2

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
- [Property reference](resources--usb_policy--reference--group-001.md#canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d)
- allowed_devices

<a id="canonical-e442a1a93c748b0c041f8e3f96ccd6470d7020396ea248ff012ab52a50de31bb"></a>

Type: `"object"`. list nested block, Optional.

Allowed USB devices. List of allowed USB devices.

Upstream description:

List of allowed USB devices.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
allowed_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-245b878f3ab58683117d62103d13ac5b0c72e4fad504519c1b7a64c263d5a2f2"></a>

## Direct properties — allowed_devices / 4a0113e5808f / 3

<a id="canonical-757e900e8fb2afca5dea8556d09d021ff3af5d47ceb52c8f1cd1fb2bf75b9ffc"></a>

<a id="canonical-dfd7474a7357b2056220fc35c6ef7cf80e9477f3addff239b60efa620dffc7dc"></a>

## b_device_class property — allowed_devices / 4a0113e5808f / 4

Type: `"string"`. Optional.

Class. The class of this device.

Upstream description:

The class of this device.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-880a9813ed96012670a89a52bfd3fcb46220135b5913ec204da8db2de658774a"></a>

<a id="canonical-dadb6b164e2b32b7dbd3707ac541ed8d7a5ccc591612c2469d894282c9c9c4e4"></a>

## b_device_protocol property — allowed_devices / 4a0113e5808f / 5

Type: `"string"`. Optional.

The protocol (within the sub-class) of this device.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0c538404d2c60f73a68ae9f2450439c1a79dd3f4280a368e7d94cad21ed33877"></a>

<a id="canonical-447bd1000b871755fa2682d3f6bb2c0ebf9d27a382e73e525b6a2d26fb41ffe6"></a>

## b_device_sub_class property — allowed_devices / 4a0113e5808f / 6

Type: `"string"`. Optional.

The sub-class (within the class) of this device.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-73e90cbd4a507ae2ab15fd8883b7fa34c0bdc967af9b41b8e75e6c4c4f046ec7"></a>

<a id="canonical-45335458d0e910ee75d137dc34efb870db57b5206a5c966d96c963c46be833f6"></a>

## i_serial property — allowed_devices / 4a0113e5808f / 7

Type: `"string"`. Optional.

Index of Serial Number String Descriptor.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c14b44f2c7e22c4f83cf09bc23a37c166402890f2d2862d3a3faef628a7e40be"></a>

<a id="canonical-5e4589ee807453e50e223fa3487c7802f28c01b3a17c44419b1d8c695f086267"></a>

## id_product property — allowed_devices / 4a0113e5808f / 8

Type: `"string"`. Optional.

Product ID (Assigned by Manufacturer) in hex.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3770d4d230ad01d7881ab2637922cb99e26f6d015f4d831cdbd7c37b73c23859"></a>

<a id="canonical-b322538b8246d606751438afa0f9edb94b6bb3f6fad73d266faafc63b22c2b87"></a>

## id_vendor property — allowed_devices / 4a0113e5808f / 9

Type: `"string"`. Optional.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

Upstream description:

Vendor ID (Assigned by USB Org) in hex.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-37bad065470bee8f8f88a5f063fd9c71a015f33a96cbfa7b9c9d4ebdb1262dfd"></a>

## Next pages — allowed_devices / 4a0113e5808f / 10

- [Property reference](resources--usb_policy--reference--group-001.md#canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d)
- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)

<a id="canonical-f075229b428c3260403de90e8cebf597503f4637a1d8124c4342a95e03114316"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b2c17a2fc5713cb0261d4ce4fb899e9e773b3b634ff6506d8fc570e8889284e"></a>

## timeouts — timeouts / 432091436e3a / 2

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
- [Property reference](resources--usb_policy--reference--group-001.md#canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d)
- timeouts

<a id="canonical-44565af17189128f8978d6efee251a444a225eb3dffdeeb969c8cd32f59cf75e"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f92a3e6975f708817329954014aa47e702bb139a135df973ac939d817c3060a3"></a>

## Direct properties — timeouts / 432091436e3a / 3

<a id="canonical-ca60a7ab9c2327e70847c4c7d7ba303bf207fa07d7d89da8453bb6a93bd106e6"></a>

<a id="canonical-a9db94dbd997574fd025b4588d0d3a5ab64c54265bdbb7135fc7eac69e951f6a"></a>

## create property — timeouts / 432091436e3a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-264e3600116da2e526835f665f2c1efc6a19adcaf52ea0c28725f0f24ffef5f1"></a>

<a id="canonical-79fb4c2dfd39825e43cf8a9f0ce3348813b7a12c227438266f1f69a3958ba5cc"></a>

## delete property — timeouts / 432091436e3a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-cd5b272b85e34b52d74450a1bf1041088649f67556404d1abe9b4c679aad92c8"></a>

<a id="canonical-e76ed73b39cd41de776e6ae4e8aed8b6deb1b398107d3570d9a055dfb6183197"></a>

## read property — timeouts / 432091436e3a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d1dad712f1480e967ffc778b77068a1da4d607fddb805f15649eafb72208d861"></a>

<a id="canonical-38534b5e9979693dd4ebc0d6b5707503f4c2c3fcc0b98967715b9cfe9b240bc4"></a>

## update property — timeouts / 432091436e3a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-be6cb90472a52e2e9af6d3fea7fa5161bf8b59a79177474109b114aa76e9a45b"></a>

## Next pages — timeouts / 432091436e3a / 8

- [Property reference](resources--usb_policy--reference--group-001.md#canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d)
- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
