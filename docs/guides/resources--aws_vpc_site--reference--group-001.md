---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e4aeeaf034cd6bcd068068f7a6696441b8ba9a38ebd62b44a01def508a8f2fd"></a>

## Property reference — Property reference / 4a6ce98b27f7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- Property reference

<a id="canonical-46694ab0bf625c00d96ace2c501789abf16679a123e7ebfe79c90c4ae4ea7539"></a>

## Direct properties — Property reference / 4a6ce98b27f7 / 3

<a id="canonical-4832ec3e92528e8f71d65e59ab4f4cb9941519a1417453c0f759385936bc53b3"></a>

<a id="canonical-3b06e34f6f344457070e453ece8667e6d06392ecbb4680d82d018a6e7252d318"></a>

## address property — Property reference / 4a6ce98b27f7 / 4

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

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324): complete subsection reference.

<a id="canonical-f4880a47443049fa387fb2e0b5678bffb9a8c9b2378f4e28a9d4319852044572"></a>

<a id="canonical-1c6f2c78a776e3d13e98c3506a5cb98294aed4e02635574066a2c01f2521b270"></a>

## annotations property — Property reference / 4a6ce98b27f7 / 5

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

- [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-f22b110b7edcf29bcdeff7d93aaf213d150a00a137b70fcb1683452fcbb381a9): complete subsection reference.

<a id="canonical-e69461432e90bd5f8cec5e179a38e375e97814bb217e929cf23533b38bd0dab8"></a>

<a id="canonical-92fa6c190b2d58a8c6c944bd2bee25282d36837e596f2586dda1f80c51596cea"></a>

## aws_region property — Property reference / 4a6ce98b27f7 / 6

Type: `"string"`. Required.

AWS Region. Name for AWS Region.

Upstream description:

Name for AWS Region.

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

- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-40d9a23f5ec89a1ea8665a4ad0cd2fd10880b880cd230a59a4e82048dff0ce29): complete subsection reference.

- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc): complete subsection reference.

- [coordinates](resources--aws_vpc_site--reference--group-001.md#canonical-d3c6e286629185efed02ddf04e4cc3b7e45b916196fbf8c57c7e92aae9adfd97): complete subsection reference.

- [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-09c7d204935c84c5375a63cbc14e06bf5157500ed84092fa1f923683ce3bca7f): complete subsection reference.

- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-5b7a3935b1838b49ee2d084b292388d079eab2dc745cc84a61f53f9b2aadde8e): complete subsection reference.

- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-4b76809e41769ed01b26b6c290952e2476c7e3775f91c338121b0d7d0ad3d43d): complete subsection reference.

<a id="canonical-7576b42b26f6a9fbdf021708bcb1af442a3f2131386a0cc19d3177670f245021"></a>

<a id="canonical-00b7adf392de906204be65371e5e164b0e569e69987b71af5b9b682c1a149a63"></a>

## description property — Property reference / 4a6ce98b27f7 / 7

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

- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-63bd3819b7e4dec701590266727d665f4d46dae0ae26408d7540c82ce07a8838): complete subsection reference.

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256): complete subsection reference.

<a id="canonical-59577a2cd939bd0f8c41dd5a6020d10847f85426bb604bf5f40398467a9945f0"></a>

<a id="canonical-13df67f90d8e09aaf7ee2abe68f118ba656382545b714ce625ab139479456043"></a>

## disable property — Property reference / 4a6ce98b27f7 / 8

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

- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-6223edc061be8cc81e48316841a7cf498750d8a0c00acc2bf767f5d931febc4d): complete subsection reference.

- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-4068adc6014448bc07d7ab63746cf6ce80a3fc063b63c7e93e64491beecbb31c): complete subsection reference.

<a id="canonical-2aab600cec4a6ca30bfd6a1e5f8849b859b4179f4c3e91ec6ad09eb9fb76d31e"></a>

<a id="canonical-bfb346c18b9c66919835ceba824cb55ad90fa93022ddbdad86c36a2f506a06ad"></a>

## disk_size property — Property reference / 4a6ce98b27f7 / 9

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2048,
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
    "ves.io.schema.rules.uint32.lte": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  }
}
```

- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-662cb1c766ab5942d9d12ddd323a9ab6cc30eb0ff961e565556e21ceab133049): complete subsection reference.

- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-9efdb256997148a2eef145984b59ba9a87faa0ba2632b4025d7d8dc6d5d78c74): complete subsection reference.

- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-175073f7a4cdbf9b1f82a592583d94135e809165c4307b252469a6fb34e853f3): complete subsection reference.

- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-da48ecd52f76e8a8414bfa8792082b40c74f6ff78395c63e44830dbee422729a): complete subsection reference.

- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-f26a15bffe5278b54999ccbe58e244df51238820933712ed60f8437927c3253a): complete subsection reference.

- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-2f43fbe97c023e023323ddbe139b16cf5e26356d9691ae04b777fca69e39128c): complete subsection reference.

- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-1a305484da30c6bdf1b10ab03e7ffaa9215fd2dd126b99629e0535a72ce56821): complete subsection reference.

<a id="canonical-eb84054199d95de0771a2d019c5e63f7173d22940b5a76a456ecb4ab2476d984"></a>

<a id="canonical-266cfb2f379bbfdbeee1c82cbd5d8f6af3eab4d6a851bc771272f6511732606f"></a>

## id property — Property reference / 4a6ce98b27f7 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a): complete subsection reference.

- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb): complete subsection reference.

<a id="canonical-27e906288106eccdecae6b0ca149ff006888cc7aa691aff70846f1303a575e64"></a>

<a id="canonical-d43172dc6ebd9d034a8e3fed216381ff82439bdba75b4bf889a0f35cf0589c08"></a>

## instance_type property — Property reference / 4a6ce98b27f7 / 11

Type: `"string"`. Required.

Select Instance size based on performance needed.

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

- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a): complete subsection reference.

<a id="canonical-b3800af37de22ba00acba49ba739e163704fd42d42c12e67f4b22d06936df5f2"></a>

<a id="canonical-37c4c76309fe6ffb201ef4225424a3973da032fa193e3254fb9ba2bd48b5183b"></a>

## labels property — Property reference / 4a6ce98b27f7 / 12

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

- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-e8bd1fa12942769f9086e9fc72f0d6cf4d396489fde894196a56996dfd0a2061): complete subsection reference.

- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-8e2259a9b2d220f4369c5133a7a4d2120a94e6dd0544e37a326c6fd85e416ca1): complete subsection reference.

- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-befcbfaedf144bd307a72ed3a7799f5b3d479e264e8d1c71aec48e5aa0e3add9): complete subsection reference.

<a id="canonical-3ac6bdb90c460a6f5aa9fc862fe646af8077a388061ba5a813ca58aab374b4ef"></a>

<a id="canonical-3a59ed56132bb2116cf133d79eb9403fae9a85ad14e6383ce1c0b7c7a265ab8d"></a>

## name property — Property reference / 4a6ce98b27f7 / 13

Type: `"string"`. Required.

Name of the AWS VPC Site. Must be unique within the namespace.

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

<a id="canonical-477467148a265d51f00d2f31d91f2d0bd1da0736f4886e5cac4b1e259ff1be5f"></a>

<a id="canonical-644c92673feb09d816efca0d599502424b1ca1669a85e55682024cb295b1cc7e"></a>

## namespace property — Property reference / 4a6ce98b27f7 / 14

Type: `"string"`. Required.

Namespace where the AWS VPC Site is created.

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

- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-3b567624fc53d7fd5be1fbd116ca3550fbd6c572ba3c41a433a5459854ca051d): complete subsection reference.

<a id="canonical-ee913788c85c8a86da740e494d8f1f0f8873ffbea734dc351e781ade07360004"></a>

<a id="canonical-d7f08277a41a3f96f7dd9c4296d1c4d80f72fafd026cc172c31bc53f88c11780"></a>

## nodes_per_az property — Property reference / 4a6ce98b27f7 / 15

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

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de): complete subsection reference.

- [os](resources--aws_vpc_site--reference--group-004.md#canonical-d9fbbd4f678898f88afff40d99873b7424d6d2c4373892f5fd2ada2490f1581f): complete subsection reference.

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1): complete subsection reference.

<a id="canonical-5a65c726751bb3f831acddfa1565ddfe69c70b7d0181073c86586cf728c30343"></a>

<a id="canonical-d6c6350140df9ec0c30aa10637af48ecfe01821d64e8bdea33241d731013c0a7"></a>

## ssh_key property — Property reference / 4a6ce98b27f7 / 16

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

- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-789f7d6e3d560f13cd6566c9347e630b0ad42ebf4f749e81c7ac27bccff98aa0): complete subsection reference.

<a id="canonical-454d05bd7d7b8d9b0f368721d565f8900969abea31347afa73c9e35716af84eb"></a>

<a id="canonical-6d126c6eef68c820b0c240c0b27eb3c35100fb7fb91dd41f9a81100b3587aa98"></a>

## tags property — Property reference / 4a6ce98b27f7 / 17

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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

- [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-e8515c5d0f67a8340f079b06dd5c6221e19fdcba0f0ec762690045b2b82602de): complete subsection reference.

<a id="canonical-7e1d45130ca6ad34af48bed23dd9606cab15fee04d7f9fefacf4afe53898b470"></a>

<a id="canonical-cccd95a4ba652baa4afd40506b91019ddf708dacf3ebb67512de9e0a8e125087"></a>

## total_nodes property — Property reference / 4a6ce98b27f7 / 18

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

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843): complete subsection reference.

- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-83c61a0cb454c5394bfd86d4c8768ac4ec0c77cf4ba07d1bfe82cfc8a808ae04): complete subsection reference.

- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-942adc18dc91b7727d9a120ebfcc7a63039428d6b5dbbe857c031f627628d598): complete subsection reference.

<a id="canonical-117b8a778f477d9c857adbc54f6cf023eb768f3a0f439008a01bf7a424340ce3"></a>

## All schema paths — Property reference / 4a6ce98b27f7 / 19

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--aws_vpc_site--reference--group-001.md#canonical-4832ec3e92528e8f71d65e59ab4f4cb9941519a1417453c0f759385936bc53b3) |
| `admin_password` | [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-594265fb9a55cbf05717f0968ff0ba6b113a4f6f887b95819d7498840c926568) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-c1625a496fc9c0d25b855245101741a1a359e1e3639d55326c95231af65728d6) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](resources--aws_vpc_site--reference--group-001.md#canonical-a985a5a601c89a1730fe53bc4bd40d078be3ce4aecb0ffc71f843e98b86251ed) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](resources--aws_vpc_site--reference--group-001.md#canonical-63635d5529ad064725223847dfedcc919e5fafb9f26bb41fc29422a0b1db6ff4) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](resources--aws_vpc_site--reference--group-001.md#canonical-a85828a996b2636be6d88c23c5c437566b6acd99b6aa83ee523706af23ce2b27) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-e92feb01480f0ec8e29d6495a2393f0f561f8c11d95eb9c0b9703425d32cba12) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](resources--aws_vpc_site--reference--group-001.md#canonical-580255315b9aaaf83d115cd6b67686bc9862d9dbffdf5f864ca6f6f94ca66af8) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](resources--aws_vpc_site--reference--group-001.md#canonical-bd77b5ef90f6ec1336b3276060d3808354c2be657377b92fbadab5631644f23a) |
| `annotations` | [annotations](resources--aws_vpc_site--reference--group-001.md#canonical-f4880a47443049fa387fb2e0b5678bffb9a8c9b2378f4e28a9d4319852044572) |
| `aws_cred` | [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-5e1ad08068f195e8b3011f54fb6e714dab99daf99f331c6fdf2804e27ef2355b) |
| `aws_cred.name` | [aws_cred.name](resources--aws_vpc_site--reference--group-001.md#canonical-ef48fb855aed633addfd47fc6c3e6eae7c92deb7470a8b127fee98eeab8d7d38) |
| `aws_cred.namespace` | [aws_cred.namespace](resources--aws_vpc_site--reference--group-001.md#canonical-d4ce77ddb40e0a2f73e70c63125111ee17c21190ad138d5330d371993fdce435) |
| `aws_cred.tenant` | [aws_cred.tenant](resources--aws_vpc_site--reference--group-001.md#canonical-e6c508e7f0d5ea6f537ca17309ce19a0e2c796fce4c6fe94708c18a74d4dee13) |
| `aws_region` | [aws_region](resources--aws_vpc_site--reference--group-001.md#canonical-e69461432e90bd5f8cec5e179a38e375e97814bb217e929cf23533b38bd0dab8) |
| `block_all_services` | [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-546cd3845738edd8580214f550b9bd2504a129df848460cf5a110e1fc330a7f3) |
| `blocked_services` | [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-5171d89de983f1ac85acd83d59c560774f1b25972698034a0984af11ff3f47a2) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0b6f233c3ad6987f004e2d39a2b4fa8c8c05df7079269a45fcc80100ce9f8387) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_vpc_site--reference--group-001.md#canonical-da2ac4282bd3721d46e4f86031bdef495decb60f6bedeab0fedae98de0705a62) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_vpc_site--reference--group-001.md#canonical-7b9647712101bee70270aaf328a1dc35b705c5a03b3ad1d237207b804784e2ab) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_vpc_site--reference--group-001.md#canonical-bff4a9cbcb5413e5def266630d906611213a892a7f94929a14e932ccbb895e36) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_vpc_site--reference--group-001.md#canonical-8265bca38ad4d0dcbf6ea4bee631e6a97c6a85b8d6ffe739e6c1b68da9e1b565) |
| `coordinates` | [coordinates](resources--aws_vpc_site--reference--group-001.md#canonical-49edcdb6337b6c2a618e2f7da69600ac0e53b8b9eb6c97984bf80fe8267bf6c4) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_vpc_site--reference--group-001.md#canonical-a1ff2ba11c8a4ffd731a4bb515bfc5c9ba0c7838be6f2175473d4abde785774e) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_vpc_site--reference--group-002.md#canonical-56f3f87e4e5bfef252c15a29db34f605f7512f797dcf8eb527b861b8b87a6545) |
| `custom_dns` | [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-c79fec7f021e02efd47b16fb9f88c76f9c5d56bc3e59006c71a1acb879fdcdeb) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_vpc_site--reference--group-002.md#canonical-7a845f568fa98625bd6a456a50fa611073f8774467481718c33370da5c016fef) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_vpc_site--reference--group-002.md#canonical-7e034b024b412634c1d2bdd397d48792a1030416dfae8a2e39adc18f382d2021) |
| `custom_security_group` | [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-9274af8148c1ca3ab69710d7190b671ae70a85d8f46291700143952cfbf04e45) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](resources--aws_vpc_site--reference--group-002.md#canonical-c24437fc82e6c4efbcdeb62d86076bc82c88e0d79eed39b1599f04d3f439d077) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](resources--aws_vpc_site--reference--group-002.md#canonical-650ca3159c41242c8a6c1f9540c04597083d762f7aa0c5d6492569185230177b) |
| `default_blocked_services` | [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-107d8daa2a0b4ba72fa0da3c56e81d32d6c424b7b554a889c0ef35890a717bb2) |
| `description` | [description](resources--aws_vpc_site--reference--group-001.md#canonical-7576b42b26f6a9fbdf021708bcb1af442a3f2131386a0cc19d3177670f245021) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-f847caf11fdd93d20920f7737442df073adfab1e70aa7fe26bc02898d8970fc6) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-e3f04534e6c45f7507bd9c708a1e38b4496d0bc44167320cb22145f214117f81) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-fd773522f917ce202475ab2522843d013941f9d0ae37272456339fda758ca735) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_vpc_site--reference--group-002.md#canonical-04dd771165c335d4e8b254472bb55facaca39b8e6c1f850efeef1ca64f655658) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-5136bb3655ebd77b20c235dc359d9f098aa9d6a17ab77c19f0b8ccc3830fcfcc) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-2f5e7703ea0be6222221e1b3960c1b6956521c7ccc0bf342a943a1ccdd77a652) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_vpc_site--reference--group-002.md#canonical-1669a14dfb29df6733f772876be4e065fab5ed4d6df9fd40403d8f60e4004a13) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-99ad18e46cc0b1697dc0fbb2ee4cb36ed3c89a9c3bbdb3895d8118a930be2e5c) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-bf8529e422458d385edb52fd5338b95eb687124a211f74141e4e8c565fefd9a8) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_vpc_site--reference--group-002.md#canonical-5c0854eaff3d40a8586fe569264c1baf2bb4487a5f136beafb5a8196ccdb1f8f) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-6e94970acc1c3eb03e5603645cc9a2a72f7a3d76add735040d556d469d34cd7d) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_vpc_site--reference--group-002.md#canonical-071c6604bf5a556c5455833357811d71373b575678d68534081624aaa2360121) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-df7e3a71460cdacbbb9ead82212834343803f9b79c808011b8c50f20d6d8488a) |
| `disable` | [disable](resources--aws_vpc_site--reference--group-001.md#canonical-59577a2cd939bd0f8c41dd5a6020d10847f85426bb604bf5f40398467a9945f0) |
| `disable_encryption` | [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-3bb62a0311a7a7d3e6fefedd77c0b189c55466526ef345b933577675a7f16b9e) |
| `disable_internet_vip` | [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-74bf589be459697a74a7431450d870406155041507e6aa4c4444f0b53716d00b) |
| `disk_size` | [disk_size](resources--aws_vpc_site--reference--group-001.md#canonical-2aab600cec4a6ca30bfd6a1e5f8849b859b4179f4c3e91ec6ad09eb9fb76d31e) |
| `egress_gateway_default` | [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-2b78c4d1a5f62696fbb388133b7f1e001e499e40c66ef19deb9a329d736e90dc) |
| `egress_nat_gw` | [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-b44c33f0c52cfc87b4bf392a8aa7159ccd82fcabe7f1c773379d84edfe584f08) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](resources--aws_vpc_site--reference--group-002.md#canonical-2bd12ca9b7eac7bddd41ec8f458300144f610775715470b9dc7d5e1c9773baa6) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-9a3715518fb6c71afdfc766cf2841c92a8348809446875906f74e91fe63f629e) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](resources--aws_vpc_site--reference--group-002.md#canonical-7b3a5500401f0a42abd87af588b7750e9de12ca834aa25b5863a1dc17201b226) |
| `enable_encryption` | [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-8cd5cbc5f3879c41d2173dfceebcde1ec00a37777e55755ee4a6c8a0b473d868) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](resources--aws_vpc_site--reference--group-002.md#canonical-eb724f49d998087076908938e93cd781f280488f23ae27dfaf15f2101780154d) |
| `enable_internet_vip` | [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-f9e38a102bd0e36baa826ae5153a21b93d95c091608aa478d8b8a0ecff2049e3) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-6012e1886fc593a833a34f70ca376543a6289b992b21d7587c409f3707b4954b) |
| `f5xc_security_group` | [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-d5baaebb13185ddb9f08cf03b5a2186aeaf97eabc9c0cfc2c226fe71a81f2fef) |
| `id` | [id](resources--aws_vpc_site--reference--group-001.md#canonical-eb84054199d95de0771a2d019c5e63f7173d22940b5a76a456ecb4ab2476d984) |
| `ingress_egress_gw` | [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-d713a098ee4dcbcdee803063f8fb7e3d8ef97c713bc4cac861138c989fc8595a) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-3239a16f6add5e06e358345b4282167501cf54c45a9525099d11af18a0b279ca) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-60e1a76442302d65cbf7fb30a8ce21ecaf6084bd5e44e0b5c529861e25f227ce) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-8d63eb0b9afd40248dc62bb716283c78b98f1bdd3928ff6f62f307c7f55ee985) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-1b8810b53cfec9a5a1c1d0693ca4df0bfbd1ea8d1297037dc9bff336f9487a5b) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-31e88374ff5953c806f3d41baa81f40f13a497c9d3ca1b3c3e7e2afe7867c8c6) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1f26b244ff524d37ca8adb44433a02ba1a461a7765bf812ebdf4f4bb72dbc5f1) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-dfeccd1538d946594d84c36aeecb62a0cb8c8efbfcdf2e5871c8072edbaaff1d) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-f4063e4aa0a9070117e6ec193c9cb8927c5f221311b191acc9d68ac348eb7f4f) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-e335b990e410395aa0c8cf610061607ffe147d5293b86e8ffce2d997823159bf) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-fe15a7fb67884c940a499a007ebe5b775af25c25f856f61e9ea69feda3ec9abc) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-bbc2cc1cd2580a81127e40bb960c68f1884ac396b2c7ad20ec7a94d18457f05b) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-cc2258bd1d73f083126572fa371d6d999fe8eee88738b30f8fbccdb955fb7f1f) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-c2256aca2f4af87695789fe2c8e6bb246c728b2daed0b4e2cf1e6a4aade49229) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-df4612a6cc220269e7889153de554e4dcdee19c3e4c5fb870fa925d78754f6fa) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-9e77eae0ed4fba5eb45c8dae8555bb33c21fc7703ad673c122a34d3bc115f4eb) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-8dccbf736c9289beda749760cb60532d5edf37fc96240dc3b0ddbe8764615682) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-f63f9832c0fe7c1c7e610830d444a06c58f03bcc914f0d7afd58219b74383b7b) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-002.md#canonical-0bd6657cb2877c62a67bce2326a30cc00ce555f8c0e323ecdd6b9ba5f77977d6) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-b00985accbc0eb02b082516319fd46904b30687dbab9be0159797a9e49d2658b) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-a30e7606f4316465c822a1bd2d13edd80963c6b9ded136f398916ce52ae14c81) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-4134ae38c2c18ae124832722eba6c8b7f006d3b68843ea020133a0b0c49c5297) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-d182c7daa3febbb0473b55d8e80d8abdf8de167ad20207133c5563f8322a2e7b) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-194c8b543f177b8df15d9bbb4d35808592b800c1e97a0b865ace4e4194082853) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-c7480beda7f70c03cd5a7c0ef927f4ca81d01d295a7a97e7ae1e1b98104fd807) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-002.md#canonical-6995e62b6656d27675fe68990e83f2396eead2b5c36bb7d226434b87bc01a9c5) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-edda45df2e5c28d7858d94765b215cae18a53f86454ba80a1b614d53e54d73a0) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-0a086d169fd327ac7971585af0cfc65a86b59fd093d84f07c697f7364c9f09c4) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-7b18c927a105c086cb6ac2421d68b2148ae3c9794b216ccd7f315f11e86fd958) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-b47b948d225b54fa4228f0dc0a85e45191b0ad4f144bb0c052a8703bd5f44f1d) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](resources--aws_vpc_site--reference--group-002.md#canonical-b6fec99e2cbb49efd440d62abdeaed8d079e2b08590524b6e54ee49b48dce9a4) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-c37f11b7ce3f1c15608232489609493a40d07c22758b99363b6ecf253ad9b273) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-002.md#canonical-79d88448ed88ef60eb1ad367805eda3925e0874b2224e74ff01c289158869654) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2f77f6ea17af2aee161cd5ffb55219d59a0483d0514a86ad4cafea2d672463e3) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-0590169f208510cb5f339968988bf9391c12c2e9b9a1287973ee82d9d39adeeb) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-d02b9a3cdff42365ad8160161c64c782b26884aae9ea05138227022654fa15bf) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-55af322486f01b721cc2645d01935d33d40d0f6e149740d56952e01d58f9e9f9) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-7ad5b08697ac2716d64e6c9514f9910e84cb0788fd2cf86a0a0dec7f1895c8bd) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-14636ffe9d0465f7109ab1ca431e84c4b357f2588db70b306f114b945adc41ea) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-9cd7e2c159349ef3761983231b06d77c33306b40882f63a7535b149bcf6258eb) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-da07c752c5c6c13aeb56ebd85c76898a3e8f49265c1c576545b464eaefd934fa) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-56b5bdbf6d9d0355fadbb3d89827393b7f1cb26b6fc1809d4202dd5b9593556f) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-cf055d6e18a05490bd7e2ad6ec3c61a8cc47116d16e1bd1f9c2410a02a99a470) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-da1ac5ef0cd6941d9fd8750797aafdb4bb0bbe382659a602b2b4c98c8ed4cf50) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-c52d753f12d51d6d63da352338e4c9b1f1f89365d7da5185e375133490c2f42e) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-bfa9913d02b0f54eb57fb808dbf75379818efe1f4b82d7b8e54551294579e64d) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-5b308535412116d3dd70851d34bde75d1dfe333ec37cc6f9f361375333dd6895) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-669526408deb9a178563868cd459d9d73f5851c1fc41d22dfcc7a1fb3fed55db) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-866fc46bc686e0a377996f8972b34a13aa177e33ff0226404d86bacfed21e48e) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-2add9ca3c5407426812e25c930fc65b313fa04f14155f20e749ba16c4d0f6739) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-d11e6e57e9d4e99d82c96dddc56641d7e99d4540edea60c212bd1890a6550aee) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-2e1c9e22c4f18f4ecc3f0b40aec88e73c2cac79aa39538d83820e1f8d504210f) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-51696de63176d949276085e49a6e185bb6b69adf5d77d04cbf4fbc2bb6cc35de) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-ea1f0309d0e5a2a655a5c7b804c11d02ef020ed5372082bcd69d023217f24b83) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-c0345d2378a430d8870a5e4f94ec1650cfeed217a1583a5b4789e768608b7f94) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-b5c0e95591f50bbff5f178b79a718d9240e896ac1522bd406082b1b02dcc7887) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-a605a318fdaa12e16e233a002ffc79d2931865d82f3c690412d696759737439c) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-51fcafcc196e64c296afcb97b06b11b94c58a61f06e87a6e28286ffd6cd901ab) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-2887107343287c2746dc652fadf606931aa71f0838e68e086593b92fb2c00fbc) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-1738ca21fc112f6710fe0b854b4386250fc094b3d94dd9d3234d57a484be0f81) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-964c21925d45389bd14dd491e7eeb739b740ed8c8c7975bda0404914e28ec372) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-57a89557d2cd276f20c3d60c57339fcc7335172218c49e887de08e2c4c43c975) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-61c965dcf99fe256cd4f12c04be05d49acf6c8d8e5e9c7819d4760b0d704f519) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-87e531582e6dea504dcde9d921a6ccbc68184327b7a919ea0a117513a3c64684) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-24623e2c08e70f2c5cc3d3d00195d871606ade0677c8d48dc08fd91ec946c4c5) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-ea898a2769d92abbc086aed10aa402c1e4c020ef888085adf6dea81fc4db9041) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-4550d01e61fe43de53fafc4e08b8fbb729ee6f5ed27fbb0e93d3b8ef13d12886) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-a5606b1c89d944d06ecaabc8f41092a81119f5574f205b165cd6b225c2e23f5f) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c50d100519e47b3228832b654ea4551c1587ac15cb708723e4c481a9a114a90e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-6cc64cebefac3bf9ac7eaeff2ae84b123b49472482c4827c4aa3b84e0f97cbed) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-003.md#canonical-71e910f387a21fa1ce9a2c1bdcef21f6c772f644ffd1938c7b6acb819f46067e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-99ad057a07128bf1a92c30779bf02d1ee2676e7ff763569612e330d40ef440fb) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-182bd0bee04b6a9e71c7b729458084e59a6192bf956686d79576fd4cbe151961) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-ee4805351a57f430acd51e1d18019899709fcd74843ba8446ddc9d260ded627f) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-003.md#canonical-67f8cd4c959a9f8da0883bd05d6509d8101185f64a79d9a91c2e8f20857ef910) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-003.md#canonical-bd079428cf3085f25be045f04ef3761b51220c969e0d97d785374525a8f75e07) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-003.md#canonical-bcdf9c7bcbdf4ac7e1ac9c3c1093ca743b31fc9ee554864186909c0dee29a897) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-162708a0a727c1351244d41d16ffa804979c98eea50894236870bcaf395291d5) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-003.md#canonical-313a4ee9deb415f24a011b4d83a318f46302b8e021a644b8ad5473e16f342b80) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-bec0fecdbafac0a8825f3f1b6913196a7459f909e8652f460175362fd05ab060) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-35d9444279af1d592ce8b96610d2fc64490d56a1b29e1cc2bf6e6057e2d100f7) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-a92468aceaa83423ed319a5e0680122244e8333aeea5d64dc384eb4f96b3c3da) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-a1bc9647e2eb715cdb46cc3d01d562afd93b2906521b733c3d1c3c807b8aaeb6) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-381edbdcecfe3151472d8526cdf6a4fe720937176990736dca336a2b285fd9c3) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-a1ae9e5e9129f2ddb43b674a538308440e441c4909f5b36653aa81a4640bf648) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-31b1809dc53fd0bffe5be7691b8b410fe84993fbad7ad3ba965bb8c813bd609d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-767b8430aeb5df5879074ea43f73cc6f3517fdb92e80638c048b18ee3ad742bf) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-7e45b8b14b7faaa45a1ffbc7b9fe8ad158b8b257c6461bc34815c56896f79714) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-3d4fb5b8d2b870ed33efbd3ac235371023b07357b224b41cd81d2b24f4471f06) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-003.md#canonical-7b56a56c928bb2775179dbb8b1aca9eb1a1b731e9252feb3075645fe349b301d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-943228c37b9e8d5b74cd3a9ece22f578ada06421ee425c92942209ca029405db) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-8d5223e129a1aec3639c924d68aa2a7f74545113f2f6fccd4b977d1aff3119ad) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-003.md#canonical-0141159f9caa23578ce3f4f65d98e3d7193327d671567cc415d05fee10e13bd7) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-140ff040033ffa09be72f8b57b30b528ce9fef250fb0c10c6267739f395c6f4a) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-8a84201813ee223da47e2dced75059518443e6a95e5e3b6490a59f662025dc75) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-003.md#canonical-86762f21cc00edf46fd6f4a3480782549106a4c1484446addcf336755983d969) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-596951496729ecf35c4293b67f66291f0b8daed0884d91aa5a46c96d4a9e85ea) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-bb87f072a634453a021eac7bb787dc033adbc3a0847ce10d031f96cc5e65469f) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-d6773f99d42ee2de2b1e848da04bfbdcc52131ab61cb4fcb6069b38b847299cb) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-26cf3c496a59097682f36cce9bba072eb0216cda05d2945b5031c2e8d114a5d9) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-7880605a489b3619ce26517dd1c5092ffa1c25b32a4e04e83b02b56658bd8f1f) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-190dab5db21f8b7a151c256c051a3d766ef6353fb68e52ad93a31d79c4fcc6f3) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-84b626f40401db77af6111d03bbdf845c3e9f0e5e3ad34f34252686db60f5032) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-a2103abfb7a59716f79ff41db6489204399823c85742c4412864bd9e9c8a457b) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-ffd7063c32f5664763e7d19290177c80948eda25b25c2caf7b29f5fd553a1ad8) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3df3a03596be2ce68ba780d4bdee89340cfc298ca293e24f00ad74732ec79b12) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-d244c24557b2ee4400d26cf0bcf20962677945b17d676f77d51d0bc232e68a0e) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-003.md#canonical-65259d049942a3dc6a9b772b1afb811404b6dbce21e021dfc80d56995afec8f8) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-dea97b8d643120497762ccd4cc7fb6afbe329833eb2968cb8e383cacdcd88033) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-c06dd6dec234901f45d0ec9cf91961922280a5aa07e7504e652490ccd476c896) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-45fb85ba639b337583e617ab21d6b6571b6ad7f915da24e3c25f3783b9a907b4) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-003.md#canonical-965dc1f7372427da2ecba4e3af9944da279332e00c11f4248fa39e230ef82d46) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-003.md#canonical-22e00cbe093a9d255a3f0f3f6d212f63f96da3c255a5ca5244e2067768150925) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-003.md#canonical-991be531ebdca6fb47711465e70e986dfeb28b8245d11182e20cc89e37c8f271) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-4d098eeddf8077b640703b0974197a61dd30fc286860ea16122711aceb0f1191) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-003.md#canonical-4b808fd6dcf10c03da556d92eabc460679104ccfcc17d1058c531c20aa2fb957) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-215f73204f16bac87d1b9b5abceba8068a1c42137e34cd4592959d6217f74e11) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-a2f5b97f3e7bc049d9a8c28458802614e465696da2d9ad7f8bd30ac409ce4be8) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-3bd5d9fe4ca016cf65f1afb1cca07203e031d6b8d9df4bd8f0a20fac5f227b0c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-f1ce4c631268d57a97e2cfdc675a9e7982e377cd9ca2e34f21b465179497e420) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-597d8929af9fe9ec5df440071a1f0e6b279c618df863f0bf2f40d71d5f03cfaf) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-7c337cec75ea276d0862e71ba582270a0bd79b71f99549e248535c38ae32492f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-1d8a1a2671855cdabe3325ef93a8519aaf2a76017c82e3f7c29dccdfe0cf1a5c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-f6267bedbec12a85c3cb23cb8cf438811a50cf809ba6367c163635439082e730) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-8c4def8d382958f7ca1496bccca59dbac1f4af6b136f1f151697cd57705f4c2b) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-66e85147055518049801c20bb408ee177af87528a7af8e16f14c69a5eecdb722) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-003.md#canonical-8c6cd3f779d74dbc55c5b8b273d876e2b6c11f0190d08bb27f90e908a242471e) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-3e6899f1d0812d39deb9b8b952c82fa774e26b1475a9a3536270293dca13fb5f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-728cde474a74a3ca766bfb85c37780dc5b3006316a024db6f96727cc6a3c7c76) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-003.md#canonical-93e4a3c0b93465869c57c442210eca3dca88f1b1555a8e1ad86bdff508a7801b) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-f65720438783c81ae21e41f61dafd5e457abf5ddc2c9e86ec6311c8d20f4bad7) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-e0eda02837310212f8b8752aa89c87a9faa2b168f99ec58abbbb6d7977c8ba93) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-003.md#canonical-305815cb695398bc8e03096c55fbb7b288c355fece775f9f8cf25c10a1ef0869) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-589adb9e3b6a96dea4231b9a76361866d893f874beea3e76f0ee5ba25e875bc5) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-f0894809dd82ab5c0360d4db6f012b77c551951d3d2fb8f4b20378d73500944e) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-b06dd12e5b5cd5e86f9564c8f50a91a9212db50adb646f28a80aaa0fca75147e) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d488a275cfd0d912bda8c1083137ea5d5a60ee8a3e679ba46a68d2e930ee77fb) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-668e586a47df6cf8006bcbf478fd3849c4c5e76b47211b1affc1098570c6c198) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-cbd193e1a954636156b52e0a1495b39d2d5bef036228778b324a99a85c43b62a) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-91199c3fd0ce7467d7b9c2db23e862a5727a55a178826d824626537b50c8cdf9) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-867572b25cc62986bba5858f3dbb56fe635993e64face12bd65a437252b63483) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-33bbc6f224d7aa8abdaecdb13103f06c0168e42696baf2e5316e193a94eff624) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-4f419771b4a9131069cad1b1b5864a8be99ec0dd2aac52e3687e1800ea84de03) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-f8f8c373817ef37f923920090b76b35ac8b8c72c9e990f2ebc8213062ab160d4) |
| `ingress_gw` | [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-673064a09e6a2668846219b579341e6df28ae8d7a84d2c33fa7479cf571b751e) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-c90fa5a4224d92f808c329b82eec9b69a8d6b6017ecf70129f7b1152412e8583) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-003.md#canonical-2da8adf8c1e911e99f25f17dbdbf761c1f6b6a346830baf523b33e696db8f664) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-003.md#canonical-0ca6a1abbbe8fadda423c937c017f88a463527493fe0fde4c80b31d78105d182) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-0fa474b2bb8bfce99b97a7de49e681b4f6942efae1fa670447ff7428102be086) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-003.md#canonical-f77a87fbb6d87fe365edf63021f384dc34a2818ea25ba67518bff2df4950d3af) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-003.md#canonical-8bf2432399d25b296b89ea498eb8e67b31c4a291770e51f4be6cb28f8fd884e0) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-b4117099a0196cd06ebacd71fc3714a03147a56c60ba28ceea70d32227bdcb1c) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](resources--aws_vpc_site--reference--group-003.md#canonical-19b85bdb753522a6d4c17aa65abde9e10d52ddc913ad2c3605c16d389c05dd5b) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-88ccc2221b792b3a820c2621b124bfafe8ebfc64901be55fbf60aab179116472) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-004.md#canonical-f68a641310fbdb4878575aa0028558df4b65b8357db8e12c32eddd883f3ee032) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-b3e4287d40043082c5d3667dc2d00321d1e39a0b65baacb554b52fbb06902c20) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-004.md#canonical-cd930ea66a8ec459283f9f841150f297f6b164527b4866211ea6ebb23311e6f5) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-9fb242d2561af114e71edac0e0d3f6a77dbc1b484f2c9ae204b8056740a36da6) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-004.md#canonical-9a60cb3417f1667ee1cd12aefe1936e0895dcd1d14fa2d9e37622bdbea701338) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-026eec207699d0b590341fb8e14158e1fe24dfd10a147b649280cc7a9d3eeea0) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-682ce2c6cc9a7df315850e09dba04d4755fb297dde63965f22229d1de0477084) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-34bfa530c2440489a217aeb7f35a1775bc6de23fe5d444831babb18038670686) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-bd80bfbe2550c6a5123d269a560e376a6c615e412e6579d28ff89a4dba3783a3) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9e347e2c67f3aef6c447f630769b61f4fa023686cb6fd9278ac955e8bf9805c5) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-dd2f6f1550004868db1a32696811ab5cd663787de4f33a11811c9b3229c21a73) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-e270bfa7e387fbdaf0c1023af21b36ac928b601abab6589b5cf920ddae850970) |
| `instance_type` | [instance_type](resources--aws_vpc_site--reference--group-001.md#canonical-27e906288106eccdecae6b0ca149ff006888cc7aa691aff70846f1303a575e64) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-bbbe103d57e01fe88d1deb31dead94903cbc4c2f3f577734e220421bf26f1199) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-bb699eedb294da3fa3647539c2af9dd838e27f59ea9e3a4c12beccfdc3440776) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-ff1fe051e8206594051ad1484106a0f9ea8d436e2287d594630896752a0d198c) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0e7343ad4575d1041ab6103ad5fe8bfe1d56c6933027c632f1e650036e4f7fa0) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_vpc_site--reference--group-004.md#canonical-3366a04ccf400626abee7830a93ac63e27487bf583dd1ba6833fd34b32958df2) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_vpc_site--reference--group-004.md#canonical-f85058ece4cd9bd38675a20ca0113096d1f9447e7a7a1a026803b9d7727e7508) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_vpc_site--reference--group-004.md#canonical-c861593d3eda4ee3499c5436d02bc775fd107523e0adfc61cd9494d210af083c) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d8533ca08156bf9655f300bde1cb5297be84bc5838ba1dfe8734b4cf3679299f) |
| `labels` | [labels](resources--aws_vpc_site--reference--group-001.md#canonical-b3800af37de22ba00acba49ba739e163704fd42d42c12e67f4b22d06936df5f2) |
| `log_receiver` | [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-839cc3a1006b8375f716fe9c10910a865c449ba4b01872bdfee2b13f4df2b2ac) |
| `log_receiver.name` | [log_receiver.name](resources--aws_vpc_site--reference--group-004.md#canonical-3fc544cc704271d19d066e7e223cb63d703e82af5548762e980f5798bb5b2361) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-4bb5f6fda9aa625c155b83064f9bd120a8ffd1f0f57f2c5cc8285b76ebe184d9) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-fec36b6f52a30bedf0276045cb785300bbf21c60229518bd688286d4dea113d5) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-82a912f582da1406a56e5a0614190eaf3ab538b3ac531c200be7b2efa7d9d193) |
| `manual_routing` | [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-c864d3ad531ccb13cd7fc450b170b6ea6f1c9458d3bbd77e0a81ece6574449a6) |
| `name` | [name](resources--aws_vpc_site--reference--group-001.md#canonical-3ac6bdb90c460a6f5aa9fc862fe646af8077a388061ba5a813ca58aab374b4ef) |
| `namespace` | [namespace](resources--aws_vpc_site--reference--group-001.md#canonical-477467148a265d51f00d2f31d91f2d0bd1da0736f4886e5cac4b1e259ff1be5f) |
| `no_worker_nodes` | [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-10f20b1b2571ba4f64da21213f026d249076dbe797ad9a63849f43bcad496801) |
| `nodes_per_az` | [nodes_per_az](resources--aws_vpc_site--reference--group-001.md#canonical-ee913788c85c8a86da740e494d8f1f0f8873ffbea734dc351e781ade07360004) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-7384ccf528ad5878419cf695fde4acb9020a041f00ad21fb507fa7355013db00) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0bcacd0223f947a47a34190e9b4b5993111160abbf60280a2363d6132573ade2) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0701ab95e79b5872db237799df2b2e4c35fdc34c2d3ccc465af61bc2acf934ad) |
| `os` | [os](resources--aws_vpc_site--reference--group-004.md#canonical-655ec6c3a5c03ca120af2c88da6d11ae517be8eb1e9c3342be4d6473300afa58) |
| `os.default_os_version` | [os.default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-002a82370040155c6062f63027dbefda2bf7554afbe43c59348ef8fb00ff6fcc) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_vpc_site--reference--group-004.md#canonical-6f49e0d9f0d0817455827ae9b44d8edfd6958cc80e811be5ebcd1f4fa6269c9b) |
| `private_connectivity` | [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-f209eb63cc8a411874e064df721409acd5caf5601b01bf4bc66a44a3062759ea) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-b663721f1cd9316618e077d17c442bb0a35a40404a907b1db466557691cace73) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_vpc_site--reference--group-004.md#canonical-d7326ee3f16add41b9adaf007a6bb4e24913f678837934c416d58b7c6d0f1146) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-15d477802a1c22a0f5c9041fece65489c53dd3a98f6ef3545e10d67c63a02607) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-691d236ee7dc35894c759c71987790792a85124eb4bf73cdb1797e8c7359288c) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_vpc_site--reference--group-004.md#canonical-9bfdc6ca54536065b47ec6de8108a6396cb9c894d72b341e53c649dc64a31d07) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_vpc_site--reference--group-004.md#canonical-46495d930a8858ae3c2515c4952cddce574383b1d622203f6f1845ca377a25d0) |
| `ssh_key` | [ssh_key](resources--aws_vpc_site--reference--group-001.md#canonical-5a65c726751bb3f831acddfa1565ddfe69c70b7d0181073c86586cf728c30343) |
| `sw` | [sw](resources--aws_vpc_site--reference--group-004.md#canonical-f265dad63b8d2fdae9a9fbba4db18760f919f55c4952590194693c507341afa0) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-a96a7b1832b2e919ca11e006657c2afbd4a6f04963cb81c935cb4b9bfb52beb7) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_vpc_site--reference--group-004.md#canonical-3cf5fdd8fb094823637a106183cc53ab4ca19f12abcc8fb421fac72f4a45d063) |
| `tags` | [tags](resources--aws_vpc_site--reference--group-001.md#canonical-454d05bd7d7b8d9b0f368721d565f8900969abea31347afa73c9e35716af84eb) |
| `timeouts` | [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-ea4b02153fb0ce5ab671520380a1c243fad8654a7b55f64910d86b20c68cc803) |
| `timeouts.create` | [timeouts.create](resources--aws_vpc_site--reference--group-004.md#canonical-543e4915d9f1736a25fba5aa471847f04ab0a74057aa5ffc9ce35258ca2a8ce3) |
| `timeouts.delete` | [timeouts.delete](resources--aws_vpc_site--reference--group-004.md#canonical-092043c5570a6c8990495f1053f3095a4fd602dee01b365fd1f7faf220964635) |
| `timeouts.read` | [timeouts.read](resources--aws_vpc_site--reference--group-004.md#canonical-8d4de19d384ec2b914cfd2e47f7f690846fef51bb4e623ca4bb8c978631e11be) |
| `timeouts.update` | [timeouts.update](resources--aws_vpc_site--reference--group-004.md#canonical-1ef485fb63bd324511a40fb639980e7959378c3d35d111f9e21aad99eeba9a97) |
| `total_nodes` | [total_nodes](resources--aws_vpc_site--reference--group-001.md#canonical-7e1d45130ca6ad34af48bed23dd9606cab15fee04d7f9fefacf4afe53898b470) |
| `voltstack_cluster` | [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-4141c9dfecf713cd3dfd532024e078d75bda39e24b2abdbeace7b643dc73f1d6) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0611cf8d680bce3c22018f73ac6b5ee536e36f2669f47b464822c679a441c8a2) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-8e202d3883782fb0a9ab00a823287ad400b094cbc5cc67bb126a7259d4287d51) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-fdfbcdce328259ff3f1edc81761e743715ae020b4b5298153da67ec1314ecfb8) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-e34e4c468bfbd5ee913339286fa38b19ac469e46e72a2ce788f479548cf48c2b) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-e99589e99d0ae84ddc4e792d14412db0a792fadcab278c30a0c569d9ce33acbc) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-92db8098b32adafeadca027df3310047bdec2750decb80ae5c7495f635be3984) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-d7944c56cd6179fae3aab4a04b1ef5549f6f689147d24685aeb138e04226fb38) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-a823a85fb83502e1e21cc1ca95528fadf7eea46250ec4f86ec9a2fc665346311) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-7180b76d2f722f6b6b33ee631e84637aa05fc326dc1ada6abf79a8ca9e441427) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-8fc797ba07b560ac4d1e188a748bcadbee6562ebd4bf0e99f83ca64d5b38df98) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-59e346f6bb4ab3e23ce74cef0c0212489b55e2605896323b109764b980219085) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-29a3dabd22a119f5c1f8080de6582a07d7a09c883fac7aeea4b95138e0d63ae1) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-c0d0dc70dfa599b771250759848a5342be7bc8202596b43cd05298d1c6ed3701) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-1e854af2e2d07dcd695682cb15505423c4b40aa1eae7b4fe9ae44db31421243a) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-efbccace5df3811cd703efee2ecc344e7aaa6669fe90a11bf66c598a76118b01) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-8baaa6e298805ac2e9913aa390ca9a66d1a5e05a505a498662e93d09ab3f9a25) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-80b19c12a55f074a30654feac97bb202a0d484efb8b59cb05540b18d3080cb06) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-004.md#canonical-c867a76d6a1b711150f2d34184a7ec4f0e2dd869a718fcc99847b42dbbbd08fa) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-b9831c905909a3e67fac2d00ffdf089325f90b8a63e97497621b48a77bf4e8e4) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3129a6cb9a91a78a6e4668715904eed63aef1bb6092f0c578e549a44fc30e4bd) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-bca41cf6e80923903a0771b4cd34fbc4dbc8ebc0aecbb53b9007c40ff1c3dfde) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-bfab9db0e30066753d3ac21482202fb92e91317196e162d4aa2112a2b58e823c) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](resources--aws_vpc_site--reference--group-004.md#canonical-0699156be774ac8da6e4ceba8b1ea3ffcc65836ec752ca4f5a4b3d59feb2588a) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-98f5e24803464bb6b06272b41e718daf0f212192beaafc31f6ae684fa3f831a6) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-004.md#canonical-3f3c202c6f858536cac3b0188e06f0d0ac7c748b0f036a160c14d1cdc89317c5) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-4ce9f218fae841feb3ef8a430a4cdf64f0162bab1f4e90734fe353e336513a0e) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-004.md#canonical-7a2692f991efbd0ae1f84d73c7ecfcc9c7d53ca768c74797db018acf8b4f18cc) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-d5311464d37cb019cb48b18e090590765bfe6e41b604f226b0f8740ba0a393bc) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-004.md#canonical-e61dab5b02b7d137254322b8ea823b63befc9337090a87e9ce42a553d58fe550) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-4993b4df2ccabd0a9775cd25d8c8cc0ae8538a6556798638b0430a0fd3813bc1) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](resources--aws_vpc_site--reference--group-004.md#canonical-70c5b447464b1d83da51df9912921d7edf048a80634f0dccea2f0045482cde42) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-37f05279794cc6281003303ceed13fc6b50911e9855deb686306cd30a8f47828) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-c1bd4fd16a3cab0fce104a16eb6b78a11435a04142c8c8db776fb72945f2ccf2) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-8eba3fe8e1f69e8b585e01ddf342c87e4f48e6bf0ba6f16e5caea36232a7ea54) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-95f860311be77962e580d0fe0c55610dcc97d701306fe3561392cb398d17d6ac) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-e1d4954dac5d28ec4cec557541679b4e545f4c2838c6d99511ef3d65e1fa0114) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-765cccf3b3e176562e20b796cb20d633b16dc0527fd6da32cdbcb2e4c01a8d4e) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-d3bc23e6e873f20ab21d23a9dfe2264d1399c9e599f6e3b6b8a9b6d79c878022) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-f44b85c33882d89c530d15f1cb07e6d343948dd113d69cf8f02ce640311e4a3d) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-004.md#canonical-a1e0dfd13bc4015a55507630b29898248784fb14786c5af37d35f9990a8149f7) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-c442d3e50f58ae2eda211d50746acdbaf3c3fe5a56db520bc9d803f9f2997276) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-70ee482d803c4e8f23b28e0c49dfcb0bb1681ae9b9c0206aa9674bd8b5cabded) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-49901eec3ff4f801d14a0985b83784970df1fe075eb078d388220387fc6a62d5) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-10607dbdf3c7916e6e34ecec605cfca0b6bd37da6add237b73c9d2b5aaf4e0ea) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-004.md#canonical-c12fca647de0420e92e7e3f64f55c987a35aff638cda4ed93f57bd89d440a0c6) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-0447d4c1e12fc0f6106bf860682b657a14bd1767a97d2fa7c80e97228828e404) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-26c4db3943e9ba1ff0ccc2f29b0184c08b412c69c21e7c110afaef8518a4f1d8) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-9b8ac3341b1a0ff84de6cfe1c550cdc3725dbca9f07b34bfc73db737381c28db) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](resources--aws_vpc_site--reference--group-004.md#canonical-2108e0b5ad37ae8e44fa15cd7f623f4022b8940531df7759e45ea117186d42e0) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-04c21825666451762d37f51929b9f63b39d6857d94c0ac2af05870c8c1d3aa2c) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-d5072b7fe7def1509a0e13db6095613c4fd29c92bb1ad86c280256dc7457bae7) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-06013a72e64ecd35bca2a9c69923d8aec11e636fd527fce9795f9792aaf7290f) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-277a72114f34188f6ce05902512935c70149aa6172e490e3d6feb7e092a2f191) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-632bb0df527b39c34e2b9c428c2e75ce7ee06f03bcff16def5699244560dc7fe) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-ec5ae296ba1a9a4b10424d02eabf241ff9b9525340cb23a8ade65b4efed96ec2) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-61f6ff7447dffcb5a7f4002c6eec32a2a03adea9853b937461ce22ed485721e2) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-f1a102293403e5eb4097dd6b6f12ce603412bb1a74fd1642051e4a3bcd2fb449) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-e5e77e034f183b655b0c4174f01998e9ec7b8a0474c362bd9bd6866f16604e22) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-a2d58706208cddb5fa5ee562b67fbf6e5a2fcbef7f297414e99f3d1fcd09c75d) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-a27547d3f73e07f0a160abb78c8804d991519356957ee0f637ad27e11528377d) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-005.md#canonical-530c0939c4bf8edff101705adb856b3c4640268b0237e2a004199098db6b7ad6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-005.md#canonical-ca8a67f9022b0895d642c5b00b8fcef3668b9b868e36cb5846839537226748b2) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-0421da6e28d2a6865b1996ca12450ca09e5c34f6a9a1834b47701569ab0fc443) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-005.md#canonical-d038b1cfa4112b01dee6936a45be0b80f5c91d3c3c42b435a13a6fed14b2c224) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-005.md#canonical-cef58680158e7a3d36f7f4e5df633a96de008247fbceb39833f0fb593c7f1093) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-005.md#canonical-a2320ed813e609f8ac3ef696169287acab6e6bdf0fa554c986fa23a3ad048bba) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-005.md#canonical-1dd1ebd4d70b8a3d896455722096421201cca87aacdfb034e82957408419ede1) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-005.md#canonical-458d900ab0c0ae1acdd6e40bf8a2047435bdb4d59c4811ba7285d3bb6f5debd9) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-005.md#canonical-3d2c5c87bb62634016ae46ac8e07eda309bf077f6c489248115d2f0b12997222) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-7db8812f40531cf8b8ac23f03f5edf53b0d423470fc62a791749c353e4081c5f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-eca2ac31f112eef36e941d1bbff4f066e53426e5e8df115f625805095d2a8466) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-76402b1dfee922988f7083ec7e1455dd95840d3c1ae8748ba46200adc2910d15) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-005.md#canonical-2104b409b0c2f6a5f254893fa06333c5d0cfb8f1196d72e8d4cdaac340efe618) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-ff3c734c67aa2e79c713b3b3343f21e2b5cc6d64034922972a9c675224fac348) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-005.md#canonical-a1f38709101dbfc3f74923a702f5fc183fd2a41f1a7b2d7625910dac15d2fd8f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-26fe804478ef500a988f3590943d2e7827797d5a47908a2567593c1a040b0816) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-005.md#canonical-7212b388caf383575f689c0829973f73e8398a825df6de44b72278c78c39c26d) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-982e38751f59e370894c3845478b90651a027f38791729989428c8da66832e74) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-005.md#canonical-1476d194cd11ae7f9348d4cbcb3d9721a47091ed141227fbd703d4942766a6b0) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-005.md#canonical-024e9ecd8d6467a01e964f0c77ae7778ae159105f07ad29c26630971330e480e) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-582c23a0c7989e59ca8af915f46a7205e6161ca7f7828c39f2edfdd28a7fff59) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-e756886328e10ff6985d26326343cae2a4e507a23da6012c57917e9b2b3807ef) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-005.md#canonical-3c9a61dc5d3d67235066c9759b1302f466e75cfce8c410ba474abc4f513e1111) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-005.md#canonical-f2f00821f7ae1a98f4aaa9ace0ed05c929ffc94e32281605515e5e00a7222e35) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-e0c9e60bd73d18c54ee30e6bfbfe224a281916c2b082785aff03e70683805e4f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-005.md#canonical-245afd38c4cdb7e956e43f4dbb8210cede8ec2dc6c853788fb8896fbbb92ee45) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-005.md#canonical-307935869997527235d0bee0627d4e4db7a19d96190098bdc8b853292a059394) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-2729a81e8b78e714bedd02d2d3f25b9acab734bb131e7cba970bef58aad57d9e) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-afe3801b8aa9ad38714753f0ac53618c8e183551abdd708e2257699e3b55cb67) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-84e5069709cc4e62b27e01b63583c0c838796fdc051df248f426cec6a25eee77) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-62832277c1913f0175a962333be6f1b0be35dd555d9b090293ad84d4532c9a47) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](resources--aws_vpc_site--reference--group-005.md#canonical-5706fc00f772ae8f12abfac27dfb4b40144509b3610f140cfa737b1204d44971) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](resources--aws_vpc_site--reference--group-005.md#canonical-7ee5861b2ef7a1999256839fd301b96a348c5483ac16b49e3306eed00dd5fa8c) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](resources--aws_vpc_site--reference--group-005.md#canonical-6d11719d985a91a63a280d7d01bdffff702867b99a770010770fd6832ce161c3) |
| `vpc` | [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-c17bad2667895dcdc52b14168a4aa14f4c0f5b93508e966139f2b4e54ad96650) |
| `vpc.new_vpc` | [vpc.new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-40470bd5ebb42f185b4be4a803d64805870ac86c9120b16e5ecac1a59d68e356) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](resources--aws_vpc_site--reference--group-005.md#canonical-ea2b28ae9a9115343a22ac49551e6a3a08d876a8741e9e540a7756809c25d741) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](resources--aws_vpc_site--reference--group-005.md#canonical-fbc255be280c013d033b5b83216b7e6aa71a7191ce233606fcc22456be0f570e) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-6d8239eee0f1557a409f6c2826811a834ab6bea624b014106115c42201c0901b) |
| `vpc.vpc_id` | [vpc.vpc_id](resources--aws_vpc_site--reference--group-005.md#canonical-2ffdde9e27233f2d9af940f50c96fab305fba442097d695e33339ad170420848) |
| `waf_signatures` | [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-855a734d7f8fdb3a3363a90c0a8160c047b912896ae2769f333d2ff841c61418) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_vpc_site--reference--group-005.md#canonical-09c1abaeeee35f2089f23028316110add73b1a327107190306fe8510652a3361) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_vpc_site--reference--group-005.md#canonical-ca366d94372de6cad2826a5dd24a1187b3be609c5329be5f12a88022e8fdb39b) |

<a id="canonical-6c91eb9aa708cde8bb9b4cfd0698a6c5a60ce8a33d96058b1dd9ab5a6316676f"></a>

## Next pages — Property reference / 4a6ce98b27f7 / 20

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324)
- [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-f22b110b7edcf29bcdeff7d93aaf213d150a00a137b70fcb1683452fcbb381a9)
- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-40d9a23f5ec89a1ea8665a4ad0cd2fd10880b880cd230a59a4e82048dff0ce29)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- [coordinates](resources--aws_vpc_site--reference--group-001.md#canonical-d3c6e286629185efed02ddf04e4cc3b7e45b916196fbf8c57c7e92aae9adfd97)
- [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-09c7d204935c84c5375a63cbc14e06bf5157500ed84092fa1f923683ce3bca7f)
- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-5b7a3935b1838b49ee2d084b292388d079eab2dc745cc84a61f53f9b2aadde8e)
- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-4b76809e41769ed01b26b6c290952e2476c7e3775f91c338121b0d7d0ad3d43d)
- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-63bd3819b7e4dec701590266727d665f4d46dae0ae26408d7540c82ce07a8838)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-1c1a902e75fcacf108197c273a15463837fcd165745e7ccc21c340cdc5f41256)
- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-6223edc061be8cc81e48316841a7cf498750d8a0c00acc2bf767f5d931febc4d)
- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-4068adc6014448bc07d7ab63746cf6ce80a3fc063b63c7e93e64491beecbb31c)
- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-662cb1c766ab5942d9d12ddd323a9ab6cc30eb0ff961e565556e21ceab133049)
- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-9efdb256997148a2eef145984b59ba9a87faa0ba2632b4025d7d8dc6d5d78c74)
- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-175073f7a4cdbf9b1f82a592583d94135e809165c4307b252469a6fb34e853f3)
- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-da48ecd52f76e8a8414bfa8792082b40c74f6ff78395c63e44830dbee422729a)
- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-f26a15bffe5278b54999ccbe58e244df51238820933712ed60f8437927c3253a)
- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-2f43fbe97c023e023323ddbe139b16cf5e26356d9691ae04b777fca69e39128c)
- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-1a305484da30c6bdf1b10ab03e7ffaa9215fd2dd126b99629e0535a72ce56821)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-e8bd1fa12942769f9086e9fc72f0d6cf4d396489fde894196a56996dfd0a2061)
- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-8e2259a9b2d220f4369c5133a7a4d2120a94e6dd0544e37a326c6fd85e416ca1)
- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-befcbfaedf144bd307a72ed3a7799f5b3d479e264e8d1c71aec48e5aa0e3add9)
- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-3b567624fc53d7fd5be1fbd116ca3550fbd6c572ba3c41a433a5459854ca051d)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de)
- [os](resources--aws_vpc_site--reference--group-004.md#canonical-d9fbbd4f678898f88afff40d99873b7424d6d2c4373892f5fd2ada2490f1581f)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-789f7d6e3d560f13cd6566c9347e630b0ad42ebf4f749e81c7ac27bccff98aa0)
- [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-e8515c5d0f67a8340f079b06dd5c6221e19fdcba0f0ec762690045b2b82602de)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-83c61a0cb454c5394bfd86d4c8768ac4ec0c77cf4ba07d1bfe82cfc8a808ae04)
- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-942adc18dc91b7727d9a120ebfcc7a63039428d6b5dbbe857c031f627628d598)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6fe7427614d50457d28d04832e7d95f9eda874ef2871276171dbee69fdde915"></a>

## admin_password — admin_password / d79e93128ffc / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- admin_password

<a id="canonical-594265fb9a55cbf05717f0968ff0ba6b113a4f6f887b95819d7498840c926568"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-c3bfad19ba48688cc0c8685b5b2a25f708c6f69f4e84bf760141eb8a712278d7"></a>

## Direct properties — admin_password / d79e93128ffc / 3

- [blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-80458833ad46d37ccea5c044a3598297aa50e6239eda339003cb3d8d61355c79): complete subsection reference.

- [clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-7ce0145bd3c6bc49f01f00dfc24582ec9c92913bc02b614be0bda1897955a49b): complete subsection reference.

<a id="canonical-1ec4ddfd2e327cfc90231600b571614dcb969092518da124383d1ad9da50501f"></a>

## Next pages — admin_password / d79e93128ffc / 4

- [admin_password.blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-80458833ad46d37ccea5c044a3598297aa50e6239eda339003cb3d8d61355c79)
- [admin_password.clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-7ce0145bd3c6bc49f01f00dfc24582ec9c92913bc02b614be0bda1897955a49b)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-80458833ad46d37ccea5c044a3598297aa50e6239eda339003cb3d8d61355c79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bf8c9d1010ed2243a3e27ece00a18c82a072ead31bbd473f435e72f5ac9fc7e"></a>

## admin_password.blindfold_secret_info — admin_password.blindfold_secret_info / 264b3b687018 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324)
- admin_password.blindfold_secret_info

<a id="canonical-c1625a496fc9c0d25b855245101741a1a359e1e3639d55326c95231af65728d6"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-65a001871721f70059ebdc101642ab0b5504d0b0074dd9d1f836833b0f81fe2f"></a>

## Direct properties — admin_password.blindfold_secret_info / 264b3b687018 / 3

<a id="canonical-a985a5a601c89a1730fe53bc4bd40d078be3ce4aecb0ffc71f843e98b86251ed"></a>

<a id="canonical-7f23745e9e3edcbe2a546e81b5fd0e85c41ac3281e67ad712be06764f810289e"></a>

## decryption_provider property — admin_password.blindfold_secret_info / 264b3b687018 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-63635d5529ad064725223847dfedcc919e5fafb9f26bb41fc29422a0b1db6ff4"></a>

<a id="canonical-803fe66c23f39b14716c2fcb6580774e87ff3d38c795f1fe373180ff3741049d"></a>

## location property — admin_password.blindfold_secret_info / 264b3b687018 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-a85828a996b2636be6d88c23c5c437566b6acd99b6aa83ee523706af23ce2b27"></a>

<a id="canonical-842b4a41191b5ccd0e520a1165cdf248677f8b6dbaa61fe3fabea692753b5f01"></a>

## store_provider property — admin_password.blindfold_secret_info / 264b3b687018 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1a31d84029ae24c0d43f9dedc9a1b51d51013bddafd62ef317635efb3c81230a"></a>

## Next pages — admin_password.blindfold_secret_info / 264b3b687018 / 7

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-7ce0145bd3c6bc49f01f00dfc24582ec9c92913bc02b614be0bda1897955a49b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ba5a4990f1ba5cad96e4016bfb4b9f1661c32f1971a75d98c1f81a2efcfa5df"></a>

## admin_password.clear_secret_info — admin_password.clear_secret_info / b964d81d0766 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324)
- admin_password.clear_secret_info

<a id="canonical-e92feb01480f0ec8e29d6495a2393f0f561f8c11d95eb9c0b9703425d32cba12"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-5956c8674bd43c6bdf4882bc54d6bf9b04bc173750e93cb14c2c00ddbdd116ae"></a>

## Direct properties — admin_password.clear_secret_info / b964d81d0766 / 3

<a id="canonical-580255315b9aaaf83d115cd6b67686bc9862d9dbffdf5f864ca6f6f94ca66af8"></a>

<a id="canonical-1061fe7ebf9dd741985786e183c9ebcfe1f244977b61fecf97c8887940a55e14"></a>

## provider_ref property — admin_password.clear_secret_info / b964d81d0766 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bd77b5ef90f6ec1336b3276060d3808354c2be657377b92fbadab5631644f23a"></a>

<a id="canonical-a797ae7cfed3b64a1de2cb60864531ccf1d33b758f28a4e1201be96c6d6c11c5"></a>

## url property — admin_password.clear_secret_info / b964d81d0766 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-27e1360128b4f9933415c26febe0d0502eca00ea36aa62211087a92cc5e47171"></a>

## Next pages — admin_password.clear_secret_info / b964d81d0766 / 6

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-eb4786af49ca9fabddb9ac6bf8cb892af88162032877fda62d93489b61f52324)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f22b110b7edcf29bcdeff7d93aaf213d150a00a137b70fcb1683452fcbb381a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f425ade1b81d3fe37259df2b1878d727e4610c22fd4e6817e7ec16adf14001d"></a>

## aws_cred — aws_cred / 8739ae259822 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- aws_cred

<a id="canonical-5e1ad08068f195e8b3011f54fb6e714dab99daf99f331c6fdf2804e27ef2355b"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-c364a7d87f8ac8849fbe1b490a96f8de6db26dab9bf9144f74de7fd4de38468a"></a>

## Direct properties — aws_cred / 8739ae259822 / 3

<a id="canonical-ef48fb855aed633addfd47fc6c3e6eae7c92deb7470a8b127fee98eeab8d7d38"></a>

<a id="canonical-56680d3b7176379062680d3e5928a2d57a1c26276873c6fb0c4f8d92139669ba"></a>

## name property — aws_cred / 8739ae259822 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-d4ce77ddb40e0a2f73e70c63125111ee17c21190ad138d5330d371993fdce435"></a>

<a id="canonical-161dfb279ebd63bf692ba848cb5937373ed03013ab12953dae5a30a9150c015f"></a>

## namespace property — aws_cred / 8739ae259822 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-e6c508e7f0d5ea6f537ca17309ce19a0e2c796fce4c6fe94708c18a74d4dee13"></a>

<a id="canonical-304c6367b10e8ac3f481a4b766b14c8ba8258a135aeeb7729aa875ac66ab1a4c"></a>

## tenant property — aws_cred / 8739ae259822 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "byteLength": {
      "max": 64
    },
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-b97f0d9c5e31878ecc2a02e937d298d7d5c0d6f64ddabeee1069425d76a1e13a"></a>

## Next pages — aws_cred / 8739ae259822 / 7

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-40d9a23f5ec89a1ea8665a4ad0cd2fd10880b880cd230a59a4e82048dff0ce29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8479b9a9908866397cfa6850c1a156102157b0ebdb96c8b950f500d301b8da50"></a>

## block_all_services — block_all_services / e74f3548ef1d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- block_all_services

<a id="canonical-546cd3845738edd8580214f550b9bd2504a129df848460cf5a110e1fc330a7f3"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-546cd3845738edd8580214f550b9bd2504a129df848460cf5a110e1fc330a7f3)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-5171d89de983f1ac85acd83d59c560774f1b25972698034a0984af11ff3f47a2)
- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-107d8daa2a0b4ba72fa0da3c56e81d32d6c424b7b554a889c0ef35890a717bb2)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-32eeb6617967157a3d344b6c9fb6b1209f20054153acae9742d08dc5a876dcde"></a>

## Direct properties — block_all_services / e74f3548ef1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d7ef2dd0dfd1ddf158f8a8113159847b505487b65d62be6f34972b55285aa69"></a>

## Next pages — block_all_services / e74f3548ef1d / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae620c7a094473f5d058d66a0c1bb6251cee0b6dbe6cc2c3444ed7b409e37163"></a>

## blocked_services — blocked_services / 06267d92fe6b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- blocked_services

<a id="canonical-5171d89de983f1ac85acd83d59c560774f1b25972698034a0984af11ff3f47a2"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef191d535625e5cbec8c604961578d4ddcfd55b2c05b0ca0ba37c53f2ae57d5c"></a>

## Direct properties — blocked_services / 06267d92fe6b / 3

- [blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e): complete subsection reference.

<a id="canonical-44fcfda0bef66ed83fce5727dd4f037b90e5791052683c0a2a4b01aac9bc19ca"></a>

## Next pages — blocked_services / 06267d92fe6b / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fded389410330d4914909a8be0d343839a62cdfe06f37e7ac94a32ad775e035"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / 22de820cb270 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- blocked_services.blocked_service

<a id="canonical-0b6f233c3ad6987f004e2d39a2b4fa8c8c05df7079269a45fcc80100ce9f8387"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-fce21a52f9cf71bef815a4d83b6e97bf6521cc43b23591f95edeea17ab114c8d"></a>

## Direct properties — blocked_services.blocked_service / 22de820cb270 / 3

- [dns](resources--aws_vpc_site--reference--group-001.md#canonical-f294a633f422487174d561169ecff9813c57bd9d0263f104dcf6d53e0e207b1d): complete subsection reference.

<a id="canonical-7b9647712101bee70270aaf328a1dc35b705c5a03b3ad1d237207b804784e2ab"></a>

<a id="canonical-29b403a44ee5ba61cb21e6446624d5c2d38ffdd1d4bfd0f31e4c219fb91a7d7b"></a>

## network_type property — blocked_services.blocked_service / 22de820cb270 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--aws_vpc_site--reference--group-001.md#canonical-86ada5d0c4e01b375da57ee15f6b00bab529b7d474d37ac0dbaaa0eb29768c0b): complete subsection reference.

- [web_user_interface](resources--aws_vpc_site--reference--group-001.md#canonical-bcfacf90dc5fe58be6c58b74ee6973a6d6e2748d49fe0c26d387cf3d2a94c2eb): complete subsection reference.

<a id="canonical-25a06b7a3105edcf2390c13772834e5675968d78d311c83d0560d2af967e0ab6"></a>

## Next pages — blocked_services.blocked_service / 22de820cb270 / 5

- [blocked_services.blocked_service.dns](resources--aws_vpc_site--reference--group-001.md#canonical-f294a633f422487174d561169ecff9813c57bd9d0263f104dcf6d53e0e207b1d)
- [blocked_services.blocked_service.ssh](resources--aws_vpc_site--reference--group-001.md#canonical-86ada5d0c4e01b375da57ee15f6b00bab529b7d474d37ac0dbaaa0eb29768c0b)
- [blocked_services.blocked_service.web_user_interface](resources--aws_vpc_site--reference--group-001.md#canonical-bcfacf90dc5fe58be6c58b74ee6973a6d6e2748d49fe0c26d387cf3d2a94c2eb)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f294a633f422487174d561169ecff9813c57bd9d0263f104dcf6d53e0e207b1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a23fec6718f8d615de71e824df295e131ade2669aa9e4091c94f5cb714e1902"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / 64035e4c09da / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- blocked_services.blocked_service.dns

<a id="canonical-da2ac4282bd3721d46e4f86031bdef495decb60f6bedeab0fedae98de0705a62"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
dns = {}
```

<a id="canonical-64a793cb70af7901cc15b65a6b25c54a1700095db2398a7e6f410a354d47763a"></a>

## Direct properties — blocked_services.blocked_service.dns / 64035e4c09da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bbbee0a52987a48246fb9556a59dcd05b6b1b69d68fe8358bbd6b618921e788"></a>

## Next pages — blocked_services.blocked_service.dns / 64035e4c09da / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-86ada5d0c4e01b375da57ee15f6b00bab529b7d474d37ac0dbaaa0eb29768c0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-042bd28256d66dea0325a473c79fb4015f41ada775072575831db2a6bc578880"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / fc1d82a82c6d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- blocked_services.blocked_service.ssh

<a id="canonical-bff4a9cbcb5413e5def266630d906611213a892a7f94929a14e932ccbb895e36"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
ssh = {}
```

<a id="canonical-358a50171add658049dd7cfffda33f0ec8997b148eba111850ad1a84179d3f2c"></a>

## Direct properties — blocked_services.blocked_service.ssh / fc1d82a82c6d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd4062c5a5c70cec23b70a842a932c3b7df95a35310462b5491b6623e0b01c35"></a>

## Next pages — blocked_services.blocked_service.ssh / fc1d82a82c6d / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bcfacf90dc5fe58be6c58b74ee6973a6d6e2748d49fe0c26d387cf3d2a94c2eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9dd5b20b83d4e0f92fab6903cb758f905d6ef6d41cd9c57612d0ae11b676986"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / af7df1786919 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-fb3aa901f26128d2e4251b73cbf87f56e03b6edcebc52523d084e597746c36bc)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-8265bca38ad4d0dcbf6ea4bee631e6a97c6a85b8d6ffe739e6c1b68da9e1b565"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
web_user_interface = {}
```

<a id="canonical-f9a6b4d5a1ff35c6f2119c3450a3091550faaf7ec10515961c30cc2977ab16f7"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / af7df1786919 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-926f674b022669d85d76a9726830f897142ffd805d8dab4d3e3f776a9814cb0c"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / af7df1786919 / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-12e8c3a9762144fa36db8f76bb8eadafdcf3aa28fae29a84354ad8dea73cb18e)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d3c6e286629185efed02ddf04e4cc3b7e45b916196fbf8c57c7e92aae9adfd97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f4311496be5cfc5a02c2abcb768ca67e07587bb04c2b965ada12dec8d590cfc"></a>

## coordinates — coordinates / 4f412f35ce63 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- coordinates

<a id="canonical-49edcdb6337b6c2a618e2f7da69600ac0e53b8b9eb6c97984bf80fe8267bf6c4"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1199fffb39cd801ba11d3d32fd4bb0cf6ecdf82fa91f508643acb7d2851e8618"></a>

## Direct properties — coordinates / 4f412f35ce63 / 3

<a id="canonical-a1ff2ba11c8a4ffd731a4bb515bfc5c9ba0c7838be6f2175473d4abde785774e"></a>
