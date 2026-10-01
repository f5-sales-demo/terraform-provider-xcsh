---
page_title: "xcsh_usb_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy reference."
---

# xcsh_usb_policy reference

<a id="canonical-407b1cbd64e749afbe04a859df95b1539bfd8fdf4088bc42bfa8bc99410033c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-773ed41a3d884f25a9377a6736ca191c92463a6e15faa385cd7fe3b7909589ce"></a>

## Property reference — Property reference / 40393a4a96d5 / 2

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
- Property reference

<a id="canonical-3d69ad0213864dcd307ad48d969957edf4941130b1516a6085af21802507f1af"></a>

## Direct properties — Property reference / 40393a4a96d5 / 3

- [allowed_devices](data-sources--usb_policy--reference--group-001.md#canonical-6add43490fd10761c3595c2046e33d64dab304bfbc4ef1c53556fae32b902160): complete subsection reference.

<a id="canonical-669fe75c42c6935c23719b45d15aca0ae98c9329cdd2efeac774461f99640c4e"></a>

<a id="canonical-89a4ebe93bc9a9bd306b03f3e18cbef97d9cefc85346d8e4d785c049b97a6f11"></a>

## annotations property — Property reference / 40393a4a96d5 / 4

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

<a id="canonical-9cda623166595520b80ce1ba2306244011a07eabbf213da3896f553b6cb1f853"></a>

<a id="canonical-987c3ba7e025fc6bb4fd615e5f978b90e18ecf0cf8e4d68d5915fb539e91eac1"></a>

## description property — Property reference / 40393a4a96d5 / 5

Type: `"string"`. Computed.

Description of the UsbPolicy.

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

<a id="canonical-e1e53e57c2221e062fae700052d505ad1ee04ad3bec8f9bcde553c34b168503b"></a>

<a id="canonical-6da0f4e083bacd460ce4c05a601805e5022a716e3d696f47672dc37b590e0f0e"></a>

## id property — Property reference / 40393a4a96d5 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d3925cd050fa8a7eb3ccdd943036dc69dd74d96ae5a19fe1d97815f9fcad7eb6"></a>

<a id="canonical-db8976e29ce0a5ff2b89a2262fa44d984f181d03090af16dab2504d1b52f6e3a"></a>

## labels property — Property reference / 40393a4a96d5 / 7

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

<a id="canonical-68e0a06435be0764bf7f4cb4144de7ed01726100eaa5278c25bf81babd837781"></a>

<a id="canonical-4f3f271767dab76784636bbee403105e75ed6c2fbaef53005a731b7e2a2b817d"></a>

## name property — Property reference / 40393a4a96d5 / 8

Type: `"string"`. Required.

Name of the UsbPolicy.

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

<a id="canonical-a55d974bc49bcbacaed64a3647ee19bfeb758295d1ae7ea2ac00e4fb785e67ad"></a>

<a id="canonical-ea69b5e5c08b09500cf6b7315e0b337f720e43bbcdd73f4dee46e89ff5c3fb7f"></a>

## namespace property — Property reference / 40393a4a96d5 / 9

Type: `"string"`. Required.

Namespace where the UsbPolicy exists.

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

<a id="canonical-6e70c963d193c0114998231b4011c1d12ea708c5a94d3d717e128595df7afe2a"></a>

## All schema paths — Property reference / 40393a4a96d5 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_devices` | [allowed_devices](data-sources--usb_policy--reference--group-001.md#canonical-da1fbe01bf54bfdc3e2153822b1da4339b2c3f94363bdc260027278b43679f82) |
| `allowed_devices.b_device_class` | [allowed_devices.b_device_class](data-sources--usb_policy--reference--group-001.md#canonical-8af35a71a8b4e81e6c30c84313242fa8228faf4e9603130e6e441a2d776fd169) |
| `allowed_devices.b_device_protocol` | [allowed_devices.b_device_protocol](data-sources--usb_policy--reference--group-001.md#canonical-5f20885a44fece104f03168c1d97a5f53b65ebb5571aca9296da7e1403581f24) |
| `allowed_devices.b_device_sub_class` | [allowed_devices.b_device_sub_class](data-sources--usb_policy--reference--group-001.md#canonical-bfaa165cadd252ef60e891a3e0ec5c327506f2fdcf2381ae64fbad228b64cfad) |
| `allowed_devices.i_serial` | [allowed_devices.i_serial](data-sources--usb_policy--reference--group-001.md#canonical-1f1f6b0067d092397f54d86937bec3afcdada8b062fb31f7dc49260c15f0a7dd) |
| `allowed_devices.id_product` | [allowed_devices.id_product](data-sources--usb_policy--reference--group-001.md#canonical-15edf64c75b83e4263bfb696b87251b2077ff85281965684e8a727e8e908a2cf) |
| `allowed_devices.id_vendor` | [allowed_devices.id_vendor](data-sources--usb_policy--reference--group-001.md#canonical-71d7151c9c743f53b9995dc9ebd02a43cc9679abca835e35e0081356e5e9b18f) |
| `annotations` | [annotations](data-sources--usb_policy--reference--group-001.md#canonical-669fe75c42c6935c23719b45d15aca0ae98c9329cdd2efeac774461f99640c4e) |
| `description` | [description](data-sources--usb_policy--reference--group-001.md#canonical-9cda623166595520b80ce1ba2306244011a07eabbf213da3896f553b6cb1f853) |
| `id` | [id](data-sources--usb_policy--reference--group-001.md#canonical-e1e53e57c2221e062fae700052d505ad1ee04ad3bec8f9bcde553c34b168503b) |
| `labels` | [labels](data-sources--usb_policy--reference--group-001.md#canonical-d3925cd050fa8a7eb3ccdd943036dc69dd74d96ae5a19fe1d97815f9fcad7eb6) |
| `name` | [name](data-sources--usb_policy--reference--group-001.md#canonical-68e0a06435be0764bf7f4cb4144de7ed01726100eaa5278c25bf81babd837781) |
| `namespace` | [namespace](data-sources--usb_policy--reference--group-001.md#canonical-a55d974bc49bcbacaed64a3647ee19bfeb758295d1ae7ea2ac00e4fb785e67ad) |

<a id="canonical-b6234964d4a8255f7479c5fb7f76be3345a95a754b4d4ce222bffb9a66546d67"></a>

## Next pages — Property reference / 40393a4a96d5 / 11

- [allowed_devices](data-sources--usb_policy--reference--group-001.md#canonical-6add43490fd10761c3595c2046e33d64dab304bfbc4ef1c53556fae32b902160)
- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)

<a id="canonical-6add43490fd10761c3595c2046e33d64dab304bfbc4ef1c53556fae32b902160"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2173a93c01ac0ddedb34f04e7aa312c62cdd05961b644d090b68488ed7088cc"></a>

## allowed_devices — allowed_devices / 410b14067896 / 2

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
- [Property reference](data-sources--usb_policy--reference--group-001.md#canonical-407b1cbd64e749afbe04a859df95b1539bfd8fdf4088bc42bfa8bc99410033c0)
- allowed_devices

<a id="canonical-da1fbe01bf54bfdc3e2153822b1da4339b2c3f94363bdc260027278b43679f82"></a>

Type: `"list"`. Computed.

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

<a id="canonical-dda8cfdbfdacd07fc97f55f7fc93ad1ead8075ecba5d4632172d83b420c9b5f8"></a>

## Direct properties — allowed_devices / 410b14067896 / 3

<a id="canonical-8af35a71a8b4e81e6c30c84313242fa8228faf4e9603130e6e441a2d776fd169"></a>

<a id="canonical-04a600c4f4d08c08ee886191bfb2f7c1fb9cd5ac2c05a3a5f4223f7f4680d999"></a>

## b_device_class property — allowed_devices / 410b14067896 / 4

Type: `"string"`. Computed.

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

<a id="canonical-5f20885a44fece104f03168c1d97a5f53b65ebb5571aca9296da7e1403581f24"></a>

<a id="canonical-e1ec8137b1b010665ae9f4fa77818f499c4c8e1e47ffafdd86ee6497295122c4"></a>

## b_device_protocol property — allowed_devices / 410b14067896 / 5

Type: `"string"`. Computed.

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

<a id="canonical-bfaa165cadd252ef60e891a3e0ec5c327506f2fdcf2381ae64fbad228b64cfad"></a>

<a id="canonical-63bb7d546c8366d013291e7afdb24f471a22661478aa62177abfcf65b3ebc4bc"></a>

## b_device_sub_class property — allowed_devices / 410b14067896 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1f1f6b0067d092397f54d86937bec3afcdada8b062fb31f7dc49260c15f0a7dd"></a>

<a id="canonical-a2c5e6f4bc94504c09a68660f8e3cb07434a8473789c8343643e8eabee815027"></a>

## i_serial property — allowed_devices / 410b14067896 / 7

Type: `"string"`. Computed.

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

<a id="canonical-15edf64c75b83e4263bfb696b87251b2077ff85281965684e8a727e8e908a2cf"></a>

<a id="canonical-3c31f04b00dce21ee83604ab571eeb87b527b391b7cc43107b9ff2b234c737bd"></a>

## id_product property — allowed_devices / 410b14067896 / 8

Type: `"string"`. Computed.

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

<a id="canonical-71d7151c9c743f53b9995dc9ebd02a43cc9679abca835e35e0081356e5e9b18f"></a>

<a id="canonical-1eb8dbb7198b529fd585123b6f00c955932c661b2be3a400952635d27b6772e3"></a>

## id_vendor property — allowed_devices / 410b14067896 / 9

Type: `"string"`. Computed.

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

<a id="canonical-b6ab01e5e0dccc0a8b17d99bf306df787b3988611c61c8a309fd223637b2a987"></a>

## Next pages — allowed_devices / 410b14067896 / 10

- [Property reference](data-sources--usb_policy--reference--group-001.md#canonical-407b1cbd64e749afbe04a859df95b1539bfd8fdf4088bc42bfa8bc99410033c0)
- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
