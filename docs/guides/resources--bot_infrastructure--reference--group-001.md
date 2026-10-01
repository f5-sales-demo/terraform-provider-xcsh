---
page_title: "xcsh_bot_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure reference."
---

# xcsh_bot_infrastructure reference

<a id="canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efaf0954bd696d5c1252ce04e40d0bceadfd36ed510c7793f32e33c5a8bf3b63"></a>

## Property reference — Property reference / 2c78a1cd1d9f / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- Property reference

<a id="canonical-ef919b63e84051cc254fa4b0d198584eda04ed6d222f835883ae3435d6e55c01"></a>

## Direct properties — Property reference / 2c78a1cd1d9f / 3

<a id="canonical-1f115ae62e731b73e8af7f855a8657f8eb7a8ce3fc377cf5525709cb5ac1fb36"></a>

<a id="canonical-c7398634be04a9c956f3cb8af08f85e87eb29b028523b265ed65962eb20a5113"></a>

## annotations property — Property reference / 2c78a1cd1d9f / 4

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

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a): complete subsection reference.

<a id="canonical-80e9fefefbe7373f75217de421b034b6fb9a992a7da013e126ecb42b0db9c2ee"></a>

<a id="canonical-bc00a3bd2dccdb2ff8e714f41e54edb31e645046078c23c174cb0b005cb14cbd"></a>

## description property — Property reference / 2c78a1cd1d9f / 5

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

<a id="canonical-38ba4b369890a8d98a5233643f5666b7cf3f06cfb38dabca64fb799ba8a8e120"></a>

<a id="canonical-61d9f71183587c4597da999ba20c0cc02689e62ac4f38ccc0128f35449a7ff09"></a>

## disable property — Property reference / 2c78a1cd1d9f / 6

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

<a id="canonical-86d1a91894871259d9d3efa1e49fcc28655173cc4a7b7077ec228a8d09d84192"></a>

<a id="canonical-3166f530175718306d36dbab657a1fafd7031cdb19d5c3020a56a6487b30df26"></a>

## id property — Property reference / 2c78a1cd1d9f / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d749594fd0d7721d12b2febe2538de64f4ef9a2112b2bcd4a8478998a8a2bdec"></a>

<a id="canonical-6874f340c97b39abd1de37b5b5f003df8d938bc74f23743996b50a2f6a70274e"></a>

## labels property — Property reference / 2c78a1cd1d9f / 8

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

<a id="canonical-6de6c0256d6d0162311a9ecac42527631caa9c2635abfd266323e2ba5af5a1d5"></a>

<a id="canonical-75889f1d3a5721db80c062882970a89d0dfdc7da8b8409e9e1a01a90cc8f1645"></a>

## name property — Property reference / 2c78a1cd1d9f / 9

Type: `"string"`. Required.

Name of the Bot Infrastructure. Must be unique within the namespace.

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

<a id="canonical-7176871b0c16da33b683c92e06312c5a327c43e19a3aed9467f848154fc8e41b"></a>

<a id="canonical-e3bec4b055cee69673909feca3fd9fe756a3ee4cacf8e046025ec1b85864d8de"></a>

## namespace property — Property reference / 2c78a1cd1d9f / 10

Type: `"string"`. Required.

Namespace where the Bot Infrastructure is created.

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

- [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-2aa4a2d11badc359557c777ae2901e45795ad604e20bf5da708ea72663b8cba6): complete subsection reference.

<a id="canonical-87fcca5b80a0ffabd9727141b03476bf122067d2479702e2040f53784f76c782"></a>

<a id="canonical-b54b480870df43ef688f3cfa1f2627d5c2dc02750bdea111a1962424f0a8afa3"></a>

## traffic_type property — Property reference / 2c78a1cd1d9f / 11

Type: `"string"`. Optional, Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

The type of traffic that is routed to and processed by this infrastructure (Web or Mobile).

Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot
Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are
routed through this Bot Defense infrastructure.

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

<a id="canonical-11231c4f0a5d07556f6a07e286cb03ab1f41019f33bc093dc2456239768a4201"></a>

## All schema paths — Property reference / 2c78a1cd1d9f / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bot_infrastructure--reference--group-001.md#canonical-1f115ae62e731b73e8af7f855a8657f8eb7a8ce3fc377cf5525709cb5ac1fb36) |
| `create_cloud_hosted` | [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-8611913bd27c43135ec99ab5637f81a4538d11fba6f4652957713ba2032e217e) |
| `create_cloud_hosted.ip_addresses` | [create_cloud_hosted.ip_addresses](resources--bot_infrastructure--reference--group-001.md#canonical-33ddb7f4484abfb5f1552a71a31d2cd81e64a66d03a74f8e91d4b1716d60366c) |
| `create_cloud_hosted.production` | [create_cloud_hosted.production](resources--bot_infrastructure--reference--group-001.md#canonical-5d005a7462a38ec38f5aea9a99967cf9713c07631ce77becfbb2ff40cf9217a3) |
| `create_cloud_hosted.production.region_1` | [create_cloud_hosted.production.region_1](resources--bot_infrastructure--reference--group-001.md#canonical-31d42932693de71556b269e7a89265e9a447e6a49d7d26b32aadc52e14dc49ee) |
| `create_cloud_hosted.production.region_2` | [create_cloud_hosted.production.region_2](resources--bot_infrastructure--reference--group-001.md#canonical-c3dcde37ea95551e995a1b848bb349b8aae7769010c259dd8eaa504e51e55c92) |
| `create_cloud_hosted.testing` | [create_cloud_hosted.testing](resources--bot_infrastructure--reference--group-001.md#canonical-66250bab00925bd502ab8e8de07333eb6f75c3d70721a0f06b7d67708d45bef8) |
| `create_cloud_hosted.testing.region_1` | [create_cloud_hosted.testing.region_1](resources--bot_infrastructure--reference--group-001.md#canonical-62d03ade132deec67d304bf1bbccfff93a57b8ccb3d6dc77cbf2157f37169e34) |
| `description` | [description](resources--bot_infrastructure--reference--group-001.md#canonical-80e9fefefbe7373f75217de421b034b6fb9a992a7da013e126ecb42b0db9c2ee) |
| `disable` | [disable](resources--bot_infrastructure--reference--group-001.md#canonical-38ba4b369890a8d98a5233643f5666b7cf3f06cfb38dabca64fb799ba8a8e120) |
| `id` | [id](resources--bot_infrastructure--reference--group-001.md#canonical-86d1a91894871259d9d3efa1e49fcc28655173cc4a7b7077ec228a8d09d84192) |
| `labels` | [labels](resources--bot_infrastructure--reference--group-001.md#canonical-d749594fd0d7721d12b2febe2538de64f4ef9a2112b2bcd4a8478998a8a2bdec) |
| `name` | [name](resources--bot_infrastructure--reference--group-001.md#canonical-6de6c0256d6d0162311a9ecac42527631caa9c2635abfd266323e2ba5af5a1d5) |
| `namespace` | [namespace](resources--bot_infrastructure--reference--group-001.md#canonical-7176871b0c16da33b683c92e06312c5a327c43e19a3aed9467f848154fc8e41b) |
| `timeouts` | [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-0874c27d958927b875b010c23f5bdc0299b5942987620bee60958a8d43222250) |
| `timeouts.create` | [timeouts.create](resources--bot_infrastructure--reference--group-001.md#canonical-f720ed9d982fdbcefaf83b65084bc04b37e296674147f65dcafa64b4cd97f78a) |
| `timeouts.delete` | [timeouts.delete](resources--bot_infrastructure--reference--group-001.md#canonical-41188793224fbc80c888b8ba6fa9d47b2bd7d5f93e116bc73a3c75634445c9c2) |
| `timeouts.read` | [timeouts.read](resources--bot_infrastructure--reference--group-001.md#canonical-c1294db1f808a70d215864d60a37d6cc388021a167c964e82c03f2b791ad1f37) |
| `timeouts.update` | [timeouts.update](resources--bot_infrastructure--reference--group-001.md#canonical-d0752ec949a1b69d1320d6e8c276be288dac41170760e3a80121c07d6c3a1f25) |
| `traffic_type` | [traffic_type](resources--bot_infrastructure--reference--group-001.md#canonical-87fcca5b80a0ffabd9727141b03476bf122067d2479702e2040f53784f76c782) |

<a id="canonical-6c425e70cbfc4d4f726814bf9f2804b2194d061558587b60b2320309d113dae7"></a>

## Next pages — Property reference / 2c78a1cd1d9f / 13

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a)
- [timeouts](resources--bot_infrastructure--reference--group-001.md#canonical-2aa4a2d11badc359557c777ae2901e45795ad604e20bf5da708ea72663b8cba6)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)

<a id="canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-600ee18589b82942693778dab7c38e33851d79307668cff6f885008a35dda367"></a>

## create_cloud_hosted — create_cloud_hosted / 818640c2c9f2 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- create_cloud_hosted

<a id="canonical-8611913bd27c43135ec99ab5637f81a4538d11fba6f4652957713ba2032e217e"></a>

Type: `"object"`. single nested block, Optional.

F5 Cloud Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("production",
    "testing")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

Terraform syntax:

```terraform
create_cloud_hosted {
  # Configure direct properties listed below.
}
```

<a id="canonical-1407898db20fe91c2d2860921ce51acc6cb0076c128c80b5600b361c9e5f392f"></a>

## Direct properties — create_cloud_hosted / 818640c2c9f2 / 3

<a id="canonical-33ddb7f4484abfb5f1552a71a31d2cd81e64a66d03a74f8e91d4b1716d60366c"></a>

<a id="canonical-831c79dff51b618e57140e5a423bbc4b8f5c8f3a1287c8906f5f7edd62ac829d"></a>

## ip_addresses property — create_cloud_hosted / 818640c2c9f2 / 4

Type: `["list", "string"]`. Optional.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](resources--bot_infrastructure--reference--group-001.md#canonical-7db05f9764030ee961d28d0d2bc27e1cdf0cb6626bac4632b207f3054b8c9dc3): complete subsection reference.

- [testing](resources--bot_infrastructure--reference--group-001.md#canonical-cfe5fcda3d623ce1824102a5e568dc75f9ffdff324942fc6e90c9d99682a8b67): complete subsection reference.

<a id="canonical-49faf6e13b502d31b4b57115ed62efce2e63bac72597ed8e4727accfcf817e46"></a>

## Next pages — create_cloud_hosted / 818640c2c9f2 / 5

- [create_cloud_hosted.production](resources--bot_infrastructure--reference--group-001.md#canonical-7db05f9764030ee961d28d0d2bc27e1cdf0cb6626bac4632b207f3054b8c9dc3)
- [create_cloud_hosted.testing](resources--bot_infrastructure--reference--group-001.md#canonical-cfe5fcda3d623ce1824102a5e568dc75f9ffdff324942fc6e90c9d99682a8b67)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)

<a id="canonical-7db05f9764030ee961d28d0d2bc27e1cdf0cb6626bac4632b207f3054b8c9dc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6560f9e7ce92a8fb97e56d0d1d370bc5207f465f46394a9e1799a4e391fd0ee"></a>

## create_cloud_hosted.production — create_cloud_hosted.production / 683618ca88b7 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a)
- create_cloud_hosted.production

<a id="canonical-5d005a7462a38ec38f5aea9a99967cf9713c07631ce77becfbb2ff40cf9217a3"></a>

Type: `"object"`. single nested block, Optional.

Production.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1",
    "region_2")}
```

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

Terraform syntax:

```terraform
production {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7264317c99c4357f914f823565fe9abc3cdbb281fb9506ebac80011e94e13e8"></a>

## Direct properties — create_cloud_hosted.production / 683618ca88b7 / 3

<a id="canonical-31d42932693de71556b269e7a89265e9a447e6a49d7d26b32aadc52e14dc49ee"></a>

<a id="canonical-802bdf34ef30c83f971958ea7b394be2dc0af299369cf1ed8af2258119d58bc9"></a>

## region_1 property — create_cloud_hosted.production / 683618ca88b7 / 4

Type: `"string"`. Optional.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-c3dcde37ea95551e995a1b848bb349b8aae7769010c259dd8eaa504e51e55c92"></a>

<a id="canonical-94549b666ecb4aefc8027b305b36ffb9b25b554e9c7141d6d89fa4194784982a"></a>

## region_2 property — create_cloud_hosted.production / 683618ca88b7 / 5

Type: `"string"`. Optional.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-b87dc7a26c19a7da25cf0972a31e33f615466d16267712417f6922558f578392"></a>

## Next pages — create_cloud_hosted.production / 683618ca88b7 / 6

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)

<a id="canonical-cfe5fcda3d623ce1824102a5e568dc75f9ffdff324942fc6e90c9d99682a8b67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9972b8ac9856c4567811370da359db15164b4982ccff05fa0b5245171fc893a2"></a>

## create_cloud_hosted.testing — create_cloud_hosted.testing / de226a23dad0 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a)
- create_cloud_hosted.testing

<a id="canonical-66250bab00925bd502ab8e8de07333eb6f75c3d70721a0f06b7d67708d45bef8"></a>

Type: `"object"`. single nested block, Optional.

Testing

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1")}
```

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

Terraform syntax:

```terraform
testing {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b7654a710678499be6028f121b9391ce43cdcf14fe4322b2652216ade1aa2e3"></a>

## Direct properties — create_cloud_hosted.testing / de226a23dad0 / 3

<a id="canonical-62d03ade132deec67d304bf1bbccfff93a57b8ccb3d6dc77cbf2157f37169e34"></a>

<a id="canonical-0bb774174ad6b2751c096c0bfbf5d756e46b70d13f84535c76f4f068b82a8441"></a>

## region_1 property — create_cloud_hosted.testing / de226a23dad0 / 4

Type: `"string"`. Optional.

Active-Passive Infrastructure configuration where traffic is routed to a single region.

Upstream description:

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-6952e848988d483d28a9128a76c6add70dfe6e0d57d3b5282ad9fcda81bd4f5f"></a>

## Next pages — create_cloud_hosted.testing / de226a23dad0 / 5

- [create_cloud_hosted](resources--bot_infrastructure--reference--group-001.md#canonical-bcd0449083a90079dba285cce8f529732e021082b9be2a5ed131f42f4395a60a)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)

<a id="canonical-2aa4a2d11badc359557c777ae2901e45795ad604e20bf5da708ea72663b8cba6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88b944e0b006e02caacd1fd7e0b545780931e69f23c2d28ee8655909b457a308"></a>

## timeouts — timeouts / 7ad2cdc21a8c / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- timeouts

<a id="canonical-0874c27d958927b875b010c23f5bdc0299b5942987620bee60958a8d43222250"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-853d890c2aebe36d0e312050195af4bfc5a6ae8e03b59b2db1f56ffb5deffb7b"></a>

## Direct properties — timeouts / 7ad2cdc21a8c / 3

<a id="canonical-f720ed9d982fdbcefaf83b65084bc04b37e296674147f65dcafa64b4cd97f78a"></a>

<a id="canonical-36d28c6f4353d1887aff5568cb8a991a799598ac2330e6541327a63f0cf87aef"></a>

## create property — timeouts / 7ad2cdc21a8c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-41188793224fbc80c888b8ba6fa9d47b2bd7d5f93e116bc73a3c75634445c9c2"></a>

<a id="canonical-10875b0685d193f7c212d9b8e2c34eea3a0c99576108b5e354a91b19dfa9b833"></a>

## delete property — timeouts / 7ad2cdc21a8c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-c1294db1f808a70d215864d60a37d6cc388021a167c964e82c03f2b791ad1f37"></a>

<a id="canonical-b14e394d4037e77552ce22a1eb30ee257c516e0107983c0f89af6e24a330703b"></a>

## read property — timeouts / 7ad2cdc21a8c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d0752ec949a1b69d1320d6e8c276be288dac41170760e3a80121c07d6c3a1f25"></a>

<a id="canonical-20302b3da648aa887f1a7bafabf2de8f8400e185df4913eb20e8a4b6d3912739"></a>

## update property — timeouts / 7ad2cdc21a8c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d6f3f8f7f8330363691fa9e7a0b9661d33a255a5e3660fe5fba0baa773152b23"></a>

## Next pages — timeouts / 7ad2cdc21a8c / 8

- [Property reference](resources--bot_infrastructure--reference--group-001.md#canonical-e561a7e41a50a82e651979a59c0ecb0c612a900a06a77b3600482ca3210b9251)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-aae65caac0c4aab59621523325e7f8203c0205db6a352330cd7f03cb7cbed28f)
