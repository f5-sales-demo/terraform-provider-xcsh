---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd74c84818e0416cb90ffef902de0b22a962f3dcf09d2ba2135ec1560728e892"></a>

## Property reference — Property reference / a5cba4f9b81e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- Property reference

<a id="canonical-721b71495caed0dbd7740055e0b5fe82af6ed1a8a1bb5101813c967f3339e83a"></a>

## Direct properties — Property reference / a5cba4f9b81e / 3

<a id="canonical-eaa7006da3fc7b072a1d82dce39b7e70eecff7592bc016bf85806e2bbf984f07"></a>

<a id="canonical-5b4740430c094df23f171380661a31e3c79f2f1f615fa528a596779431fa761e"></a>

## address property — Property reference / a5cba4f9b81e / 4

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

- [admin_password](resources--azure_vnet_site--reference--group-003.md#canonical-a4c5e0742f0e1519317ff9370a7c911f0407b57cdd804bb4d857c8681333a73b): complete subsection reference.

<a id="canonical-ac02f1dcf923e92c1e7c8cdcce802c7fe460e3c5a42d38e616d44c76870b1805"></a>

<a id="canonical-417a1f22bdd768a0b8a8fcf46440a4a4582b02d1d471eef63578429715a98fa2"></a>

## alternate_region property — Property reference / a5cba4f9b81e / 5

Type: `"string"`. Optional, Computed.

\[OneOf: alternate\_region, azure\_region\] Exclusive with \[azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[azure\_region\] Name of the Azure region which does not support availability zones.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](resources--azure_vnet_site--reference--group-001.md#canonical-ac02f1dcf923e92c1e7c8cdcce802c7fe460e3c5a42d38e616d44c76870b1805)
- [azure_region](resources--azure_vnet_site--reference--group-001.md#canonical-060a5f63c8d746697c36600dd4fb60f3ac65f552feabbe66a9ca76e7378c2b83)

Select alternatives according to the provider validators above.

<a id="canonical-38cea25ffe8e84b24be42c560738d5530a4766d5d0b672bc475395cc0be8b382"></a>

<a id="canonical-e537ee46f66cb577614fc0e937e606316e5dd24bf6d3e0ab9aa7a5630fda1f67"></a>

## annotations property — Property reference / a5cba4f9b81e / 6

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

- [azure_cred](resources--azure_vnet_site--reference--group-003.md#canonical-ec1403c2fbb4d6e4a2a1dfc59ffa231f1257f4e63712477f80e802b634b52705): complete subsection reference.

<a id="canonical-060a5f63c8d746697c36600dd4fb60f3ac65f552feabbe66a9ca76e7378c2b83"></a>

<a id="canonical-40b16c89d6c22300babf4e91c10e6b23a8d7a4ca7b87818db0f24d66ed449b06"></a>

## azure_region property — Property reference / a5cba4f9b81e / 7

Type: `"string"`. Optional, Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](resources--azure_vnet_site--reference--group-003.md#canonical-12c9c041e0e21d24be09d73e98ec59771a2bd10683b852adf888552d77975d82): complete subsection reference.

- [blocked_services](resources--azure_vnet_site--reference--group-003.md#canonical-dc456181bdc0c6e7e58a998f9312360811b3dce00018b881d9e7d984e5bcc916): complete subsection reference.

- [coordinates](resources--azure_vnet_site--reference--group-003.md#canonical-73a266db759dc5a559e48f8d1b3ffd9a7eae765928bba683b8e1bf7d9ca44dc4): complete subsection reference.

- [custom_dns](resources--azure_vnet_site--reference--group-003.md#canonical-e2da174b1ab504b3d50fc290e07b8d862735bd75525efb5441ae9ddcb1f7cb26): complete subsection reference.

- [default_blocked_services](resources--azure_vnet_site--reference--group-003.md#canonical-43717785b38c0704b1f14f27398a55a7f3967e52d00d4f0e2c27100a4020122a): complete subsection reference.

<a id="canonical-1f1a1d855aee038d81bb8182d8ad967395f97121a847c22922b79d64150d41bb"></a>

<a id="canonical-a144a599ed9f6d93aa748ee690d371911760851c9f9d2ad60f8805de88b4133f"></a>

## description property — Property reference / a5cba4f9b81e / 8

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

<a id="canonical-38eca63315fdcf6ae9ac44af278e1598ecd1ccf1a9a4fd9610fc866a385784a8"></a>

<a id="canonical-8e85a2d8f7fb2d230c58ee292b146e0124f69b0f50a6e04f0b3b21f4adf96611"></a>

## disable property — Property reference / a5cba4f9b81e / 9

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

- [disable_encryption](resources--azure_vnet_site--reference--group-003.md#canonical-e4d301250d4d9b29ec227ac0573f23fe067ba1ad11d8b21ffc94183a96ac6275): complete subsection reference.

<a id="canonical-e01c91cdcdd4dd9b535cfbe047d371ae96168b04ca1bba05498bd0b6185943fe"></a>

<a id="canonical-5fb30291d44a99a4a7925170e6c3b83f341c25667d77c4398e9c50f89d23d095"></a>

## disk_size property — Property reference / a5cba4f9b81e / 10

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](resources--azure_vnet_site--reference--group-003.md#canonical-40d3a078b52d7085d32f3127c6c25340698a2f734377f03a67cbf910b683d270): complete subsection reference.

<a id="canonical-974aa6ac420ade05573fba37758dd704c84b627041a237678cc30a8285bda7fb"></a>

<a id="canonical-c9e43cc46914bd996790bd2dea01dd8bbcfaaa23b47dab9f35ad6affa960549f"></a>

## id property — Property reference / a5cba4f9b81e / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218): complete subsection reference.

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5): complete subsection reference.

- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228): complete subsection reference.

- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4): complete subsection reference.

- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-b5c1af0f284a4e0425831c6c3358f56df9913afd1d87a0cde4ebdc2d992213f0): complete subsection reference.

<a id="canonical-c7bff45ebffb2de10a1c2c56ffa5afe99975c6a5ccefb7f0bac7fccbd976054d"></a>

<a id="canonical-2a9f50ac5635dd5b06e8b8e30b69675509d2ed8fa7a5a00f6a00f91b202a81a7"></a>

## labels property — Property reference / a5cba4f9b81e / 12

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

- [log_receiver](resources--azure_vnet_site--reference--group-008.md#canonical-56780c660093e40a35bccfb572d45434d392039269f7dfe6edd188bdde47f2aa): complete subsection reference.

- [logs_streaming_disabled](resources--azure_vnet_site--reference--group-008.md#canonical-25258cc231a9d4026b2cde1f1ec58b3aee1f285ea58854a696755c09feb1df65): complete subsection reference.

<a id="canonical-3c2e06aee94ed2d8fe7d7637b0315f751d4dd98ae229d4b9a0d07323d8f9bca0"></a>

<a id="canonical-bc21193eda5eb72f941903df3686dd723ecfbe011ac3ccd7c8d7f228dca3b196"></a>

## machine_type property — Property reference / a5cba4f9b81e / 13

Type: `"string"`. Required.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-a4a646831b24b1b67a89af89e1d936cc4388dfe028eda37e6e79b655ce7e432b"></a>

<a id="canonical-58b28241bc8466491bde75617308d8e9c83e7f62ff51634d46c3e2bbe65d601b"></a>

## name property — Property reference / a5cba4f9b81e / 14

Type: `"string"`. Required.

Name of the Azure VNET Site. Must be unique within the namespace.

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

<a id="canonical-1d278c70f741d050597ecf1e8dd14ac8d4b441d1b09223f172ac9c0d68ffe5f4"></a>

<a id="canonical-5d5152abe269487b6320007af13e1eaa1aa54a26a8dc35932e67d6e656422af7"></a>

## namespace property — Property reference / a5cba4f9b81e / 15

Type: `"string"`. Optional, Computed.

Namespace for the Azure VNET Site. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [no_worker_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-12091dabc447e067174432d59d031f10b3730bf83bc590517c664f1ebc67620e): complete subsection reference.

<a id="canonical-f7414ad3e0bf098c5c7e53b85428673559f45d751bb2fc3f22541b66596c634a"></a>

<a id="canonical-34be88f5570ff3ee159630e2e1215e90ec8095e819aed9c26213743bc561821e"></a>

## nodes_per_az property — Property reference / a5cba4f9b81e / 16

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-6723ac4ac38dc4b1691aa0cf9dabdebad2791a8eb4fac00fded397135cac5838): complete subsection reference.

- [os](resources--azure_vnet_site--reference--group-008.md#canonical-70dd020b208d2a6edaf4a3aae30f4ff67ffceee27ffc85af22c92dc13e8b70af): complete subsection reference.

<a id="canonical-1c87220f8a95c4e5930c0c9c6cde90ccd1a42a55c7bf0090afba117e12455cb8"></a>

<a id="canonical-820465d96184e191ed4475da1684f9efe5c00632d2286417cb0791f9db0c7ac2"></a>

## resource_group property — Property reference / a5cba4f9b81e / 17

Type: `"string"`. Required.

Azure resource group for resources that will be created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d510fb394e074f880f3004741764dd7ecbc6ef12295445434a2444257290854f"></a>

<a id="canonical-d83b9edbf7a29dd5150102b079fe1f5bda30e6135296979d5422173ec65f2e91"></a>

## ssh_key property — Property reference / a5cba4f9b81e / 18

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--azure_vnet_site--reference--group-008.md#canonical-44c595767a0eea838d43115b7d32a7c06c846d485a06c5ffc924da3ceec773d7): complete subsection reference.

<a id="canonical-094657e3a259e2bc8db26ac12d69264a794275ff5a995072277f1eba94448acc"></a>

<a id="canonical-1a3754cff8b3102b910ed50b90ce29210d30aa4f000cab7c9bb2d6e6c6f38a19"></a>

## tags property — Property reference / a5cba4f9b81e / 19

Type: `["map", "string"]`. Optional, Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [timeouts](resources--azure_vnet_site--reference--group-008.md#canonical-7aaf0ee72da98f9252c44a72ca2e351fbad6c9e97e7f0920851c965169a6da21): complete subsection reference.

<a id="canonical-2ee62b44b3416fa22f983c0aa54c0693978eade311abfd14a43cbecfaff46541"></a>

<a id="canonical-903fc4374232517bdfb4b7ad450c589feeaf774321a4331d45d945a6e816dc0c"></a>

## total_nodes property — Property reference / a5cba4f9b81e / 20

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-8c0e3483e56222f4feb0fd6a8c47c326590ec63eab2337454d8ded6295ce7d93): complete subsection reference.

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314): complete subsection reference.

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f): complete subsection reference.

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb): complete subsection reference.
