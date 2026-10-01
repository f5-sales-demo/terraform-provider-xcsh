---
page_title: "xcsh_advertise_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy reference."
---

# xcsh_advertise_policy reference

<a id="canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c9e7a7df7683a2e7a45950b96c00c0c2fc691e740f685523ad8aabf3e6fc8e1"></a>

## Property reference — Property reference / d121f1992063 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- Property reference

<a id="canonical-3ad0ff25038a698a64f86d333205fac9f9761ffbe00fbd53409ec947cfeb1131"></a>

## Direct properties — Property reference / d121f1992063 / 3

<a id="canonical-f0a1afb87c03578ee99d3dea806efa1eeefe027a32506c8eacbf6594bd30199c"></a>

<a id="canonical-7e4bf33bc9a139548b6f083c013e136a76c96a4002c6c8f9e574b8364ef44486"></a>

## address property — Property reference / d121f1992063 / 4

Type: `"string"`. Optional, Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Upstream description:

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
"where" is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site object will be used based on the network type. If inside\_vip/outside\_vip is not configured in
the site object, system use interface IP in the respected networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-b8fcefe4276cabfb0b0f2bdbf66868f069275019f5008e005caba43233a3be9d"></a>

<a id="canonical-3ae149157f89a79a13b15253a5de3fbbda544d6b85b19c2a18aadc2ab0d6a344"></a>

## annotations property — Property reference / d121f1992063 / 5

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

<a id="canonical-2f76e5e857bd8bdfcadf440b4addb26f2039fd2ea412e70fe8244c3dcfd7a90b"></a>

<a id="canonical-5aec13659fe70459bb0258c09b62476f5d320fb9806022b77c3cb04336a6bae6"></a>

## description property — Property reference / d121f1992063 / 6

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

<a id="canonical-eafbd6602dff47fd0b6c4e849b0a3422818676d061f5929c5c573b6c922aac67"></a>

<a id="canonical-e791124d953ac134cf68a0281705f1a9278577d2d8c44680fa9c46d0a90427aa"></a>

## disable property — Property reference / d121f1992063 / 7

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

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-dc6d846a688e96e876ec4bb21eb0adf7eb7d22f5ad4753352f7a7003e78b5079): complete subsection reference.

<a id="canonical-cf15574019fc7570b7a344441c62007c092c44e7766a6d9e24626614103ca72e"></a>

<a id="canonical-4d50291c24e3d542c6c5c19af89f74d79112269552d6709304e4d9a68d094680"></a>

## id property — Property reference / d121f1992063 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-322e87a9f7c05a42845c949dccc855f050551a4586253edc350fc2a65685348c): complete subsection reference.

- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-988e666ede6f194642cb5fdb4ae9db08ae4e80f89ea4d03c2a1d1debea52f26a): complete subsection reference.

<a id="canonical-3288f687384a0a33895b8b42e85d638f858bf78fc2e46538966f0ed0913fc03f"></a>

<a id="canonical-5e589f4162668cb251fc5eb999d24d03a6bdd6618a356b4c862decde2a38a026"></a>

## labels property — Property reference / d121f1992063 / 9

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

<a id="canonical-18c720ec380224934dce3973b210a9ed1c4adc09dd198af006ed402405e9677d"></a>

<a id="canonical-f99774b62e787b5522f98b7f372f833e8d28add666380f14479303e482a6375a"></a>

## name property — Property reference / d121f1992063 / 10

Type: `"string"`. Required.

Name of the Advertise Policy. Must be unique within the namespace.

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

<a id="canonical-1ea871d578d7f4507d8b2f6bfd83385be2dd5249b6b2b7d82b6286ee60de6465"></a>

<a id="canonical-db67d2a43831fec0483ac54ba78df66146274318a86b336d122b32c118d90d7e"></a>

## namespace property — Property reference / d121f1992063 / 11

Type: `"string"`. Required.

Namespace where the Advertise Policy is created.

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

<a id="canonical-4da024f8f7c1319c8c33c08c2fb71f619d83581f5e88ecd757848946373ac393"></a>

<a id="canonical-00b9f4fcdf1099c5cc2d42088542d3a2da576d0d7bc92dd9a5c43eeefd77fb1f"></a>

## port property — Property reference / d121f1992063 / 12

Type: `"number"`. Optional, Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Upstream description:

Exclusive with \[port\_ranges\] Port to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [port](resources--advertise_policy--reference--group-001.md#canonical-4da024f8f7c1319c8c33c08c2fb71f619d83581f5e88ecd757848946373ac393)
- [port_ranges](resources--advertise_policy--reference--group-001.md#canonical-d44ac9cd93bf6772f3efbb00250e18b910b2a2359ab923e7dad34fef112f96b4)

Select alternatives according to the provider validators above.

<a id="canonical-d44ac9cd93bf6772f3efbb00250e18b910b2a2359ab923e7dad34fef112f96b4"></a>

<a id="canonical-870ff1b6a1ecef1232f0d0af6a927c8ee0e55d6b6877b5268a32b6d60bed3c69"></a>

## port_ranges property — Property reference / d121f1992063 / 13

Type: `"string"`. Optional, Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-7830dce4866edd4ab5a709451acc8c7457081c5f419ada010af2b58ba8956b4b"></a>

<a id="canonical-0e5912a2078f7d2431291b57e7775b665a110c960e22429a8b5e8274d6bd4186"></a>

## protocol property — Property reference / d121f1992063 / 14

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Protocol to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TCP",
    "UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](resources--advertise_policy--reference--group-001.md#canonical-b3c548b5aa3a8348d9cff3799377ea9fd555cd8938e9759c3199ab15ab8efe86): complete subsection reference.

<a id="canonical-d801c7e0b5534005108714ddcada0a37371c001689bd482680357120bb7bfa97"></a>

<a id="canonical-3be12c0708ff4a76520d50ce79ed21f45ad771ff59c16cc59ea4f336d3d1e72e"></a>

## skip_xff_append property — Property reference / d121f1992063 / 15

Type: `"bool"`. Optional, Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

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

- [timeouts](resources--advertise_policy--reference--group-001.md#canonical-43f161effa9d2d50609bf9656eb56f7c91b255f77e6da0c988f8f6aeb3ff54af): complete subsection reference.

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc): complete subsection reference.

- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411): complete subsection reference.

<a id="canonical-9dadccbba64a9b5c68fa611d294243ebb6bbf45373f9596b20a010e5a4716589"></a>

## All schema paths — Property reference / d121f1992063 / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--advertise_policy--reference--group-001.md#canonical-f0a1afb87c03578ee99d3dea806efa1eeefe027a32506c8eacbf6594bd30199c) |
| `annotations` | [annotations](resources--advertise_policy--reference--group-001.md#canonical-b8fcefe4276cabfb0b0f2bdbf66868f069275019f5008e005caba43233a3be9d) |
| `description` | [description](resources--advertise_policy--reference--group-001.md#canonical-2f76e5e857bd8bdfcadf440b4addb26f2039fd2ea412e70fe8244c3dcfd7a90b) |
| `disable` | [disable](resources--advertise_policy--reference--group-001.md#canonical-eafbd6602dff47fd0b6c4e849b0a3422818676d061f5929c5c573b6c922aac67) |
| `dualstack` | [dualstack](resources--advertise_policy--reference--group-001.md#canonical-abca3e3112c23f36c03125bf3813d6b797ef2a2fedce8bc1acac85454aa02682) |
| `id` | [id](resources--advertise_policy--reference--group-001.md#canonical-cf15574019fc7570b7a344441c62007c092c44e7766a6d9e24626614103ca72e) |
| `ipv4` | [ipv4](resources--advertise_policy--reference--group-001.md#canonical-ff3c64084adc23b304c255e203a0a1d1f8f7c686147af8597dcea4b5a8c02590) |
| `ipv6` | [ipv6](resources--advertise_policy--reference--group-001.md#canonical-8063764a258576f3b1041c7967150279d747250bf67d8ae7ffc6e2431269245f) |
| `labels` | [labels](resources--advertise_policy--reference--group-001.md#canonical-3288f687384a0a33895b8b42e85d638f858bf78fc2e46538966f0ed0913fc03f) |
| `name` | [name](resources--advertise_policy--reference--group-001.md#canonical-18c720ec380224934dce3973b210a9ed1c4adc09dd198af006ed402405e9677d) |
| `namespace` | [namespace](resources--advertise_policy--reference--group-001.md#canonical-1ea871d578d7f4507d8b2f6bfd83385be2dd5249b6b2b7d82b6286ee60de6465) |
| `port` | [port](resources--advertise_policy--reference--group-001.md#canonical-4da024f8f7c1319c8c33c08c2fb71f619d83581f5e88ecd757848946373ac393) |
| `port_ranges` | [port_ranges](resources--advertise_policy--reference--group-001.md#canonical-d44ac9cd93bf6772f3efbb00250e18b910b2a2359ab923e7dad34fef112f96b4) |
| `protocol` | [protocol](resources--advertise_policy--reference--group-001.md#canonical-7830dce4866edd4ab5a709451acc8c7457081c5f419ada010af2b58ba8956b4b) |
| `public_ip` | [public_ip](resources--advertise_policy--reference--group-001.md#canonical-3ab6f41bb4ff56cf6c012c2e7c3e7f2a7fb8376c026c373dfa026962dbb13813) |
| `public_ip.kind` | [public_ip.kind](resources--advertise_policy--reference--group-001.md#canonical-8036e915b82d1ee4e2f9f564aa1961d5d784f22570505fcfafda98f5ff4f9822) |
| `public_ip.name` | [public_ip.name](resources--advertise_policy--reference--group-001.md#canonical-94b1db7c3c88d5c6d08234b962a8ea701f4c03b580503a346abafd7589d59610) |
| `public_ip.namespace` | [public_ip.namespace](resources--advertise_policy--reference--group-001.md#canonical-bf413663821d9d27e34778e6b66e5e791d1d1eaf4fd6cd2d113dbb40a735a40c) |
| `public_ip.tenant` | [public_ip.tenant](resources--advertise_policy--reference--group-001.md#canonical-5bd6e6253f2d4b6696d209e9923547fb87cb3fe661f52c7dd79809052b3a4c98) |
| `public_ip.uid` | [public_ip.uid](resources--advertise_policy--reference--group-001.md#canonical-36373ad6e4c95d30b4eb7331e8764c30dc3385ed3955ea3dfbda0a50725766ca) |
| `skip_xff_append` | [skip_xff_append](resources--advertise_policy--reference--group-001.md#canonical-d801c7e0b5534005108714ddcada0a37371c001689bd482680357120bb7bfa97) |
| `timeouts` | [timeouts](resources--advertise_policy--reference--group-001.md#canonical-82dcd1a25cfe842fc2922d447732b7c81e2def103faced0af84d99ceb9a7adee) |
| `timeouts.create` | [timeouts.create](resources--advertise_policy--reference--group-001.md#canonical-2dd9df4c8edfbe33bf842efd77aa54198d1320639d9f2443b3989ecf9714d487) |
| `timeouts.delete` | [timeouts.delete](resources--advertise_policy--reference--group-001.md#canonical-54c5490b99db02b78c6933a241d47d4e1980dfbdbad7223a0b85e9173d4dfd32) |
| `timeouts.read` | [timeouts.read](resources--advertise_policy--reference--group-001.md#canonical-1c21c07ecaa378ca64ee45f8d178b2f0adff2c18baace0e8741c1913b20814f4) |
| `timeouts.update` | [timeouts.update](resources--advertise_policy--reference--group-001.md#canonical-09a1d58ed82884688812c42cce69b0c22248d6fc39b9bc563c46413db42dea90) |
| `tls_parameters` | [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-ba3f99861b8f62bbb57f7b244f2500c6bbdbb205dd30bc3200977bb887051ffc) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-ac5aed29bee713f23e41a9e43822ba72b04ad5211009fc055155e33aefd1982c) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-f17ca6a37f12d48f846627c7229e2c4684d6e11160da0dadfd02dacb57b0a7cf) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-78553486a97769af0d0cf6292c63a02841f78d3b92fd08a3bd0275229e69b1f1) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--advertise_policy--reference--group-001.md#canonical-445d1a45d9a92eda5415062a350e6c4b3f071d25e98b06b91e62cd7fbde23ab9) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--advertise_policy--reference--group-001.md#canonical-5d02c076de7935405fec394711d6d721a253fd96eee8915a962057fa58067145) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--advertise_policy--reference--group-001.md#canonical-bd24bf70683e71f1ff9763d2f1a63686516ab9a2045cace9a74bd4a59cc77003) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-036a51231b3596e1a2ad06ebcb22428fd7690987af13bc2862f88f84df3d14ac) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--advertise_policy--reference--group-001.md#canonical-e72c0c0e743443fe319ea2dbafec7cefe8fa106184b2efde3d14f08cca077d79) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-bc8ba1de8503b8e4eb768edc4ab37ab75e472c9b38e6a52c8bfd6d286d262634) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-c4fc211e86779274c0af80af7a4cf7403ddb9eff0059f0d888408b314954d1cc) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--advertise_policy--reference--group-001.md#canonical-467bb0992b6ac6300a92018e05b09bb11fa4ca56d39dfee55d07d71d0ad296c4) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-d1e3230eb1c51ae3ad771209b9066b3ea5ff2aaeb82e3097ba29634cbf640b0c) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-65f6d614197e97d71f5dc85eff9fc20a47b63a464f15cdf096bb6b8952026781) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-088e8fb7b82928465a58fc7b359cd68767794213741586cc5ae24e93a2310745) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--advertise_policy--reference--group-001.md#canonical-0344e8541f3aa3cb3cccf8d8e2c36c5c9e6ba223eddecd59c66e0f80f92d1de0) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--advertise_policy--reference--group-001.md#canonical-e30125bd2b140d778fadbff1462d7fa0808ba55bc74e41302c951f258b1e3569) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--advertise_policy--reference--group-001.md#canonical-7bb1801cebbbab6c44bb4bd8eebb016a61056e165ff103356e495d12427b805b) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-0471ea66d19eeb0eb26bfdda83644dd344ea7aee65d395b55b739d5fc258ded8) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--advertise_policy--reference--group-001.md#canonical-2da7f872eaa1bd328273021579e9a22e7eca208f08e259101bf798ecbd11f4ed) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--advertise_policy--reference--group-001.md#canonical-08ef70ccf34736f3336cff20c48ff58cdb1c93afcbf09ce5128a8259335d2733) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-9a1660f6c4ad29e7cf5ec131451d458c598ddb9730db1655697b71b2bfb475d5) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-0bd5d3978b4b708c7627cd4653b26be1fc6ee05df922a795162bae4f1074aa86) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--advertise_policy--reference--group-001.md#canonical-91cf0f4a2629e57f039c9f15e98190b09f553c5f820de2a55bc052a637aedb2a) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-2cd58d2f3d9c71c5eb1117c9a8bee021eaf3c5ceef5500d80643e02a329f0f13) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-df1fdb9857d77330136c66bd191b824762f1606f5cc16d0f15680aaba53d8bf1) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--advertise_policy--reference--group-001.md#canonical-6b586509dbc5b24f253f73adea82a48341f935879d6ee633b95bd648520d1291) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--advertise_policy--reference--group-001.md#canonical-df9fbc2ae1a684c6b7f01f2d516d94da35914f8658e9898b66b1c3adc87b223c) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--advertise_policy--reference--group-001.md#canonical-98c3fd8e86986a761f97366542217f29514ec1b955553563ce2395ddad09d904) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--advertise_policy--reference--group-001.md#canonical-64ca1633cb781045f98e1be1380773aa3eb4998d5abdd95b8b24f3372f72afee) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--advertise_policy--reference--group-001.md#canonical-ab86cf97ff9a782ec218ba8d8a9c6083d58abae044b7f5738d8642afeb3415c6) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--advertise_policy--reference--group-001.md#canonical-d01f727a3d9013e04ca315bfe4cf10e501a35d0215c8922e2b3f395e705ee40a) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--advertise_policy--reference--group-001.md#canonical-48013132d0a028faedc7c5fc5cfaeef5d11afa6797cb9c0ad0f4b7286328fabe) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-e126f3a35d654726dca8999e770ce24bf031a75bd37e88be4ec6fdf65f841dc1) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](resources--advertise_policy--reference--group-001.md#canonical-6fe0d7df1fd169183dbfcada617cc775348a243312ae443a234042053a4a2920) |
| `where` | [where](resources--advertise_policy--reference--group-001.md#canonical-b67e098c0f131e81a4df6aa2165a1ea4e1810e0c6cfd2fe2c2694d9777fba9e3) |
| `where.site` | [where.site](resources--advertise_policy--reference--group-001.md#canonical-7405e90b6286dc2dd5c5fbf44633668431c12701992b7de2e2917a003bfcc34f) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-a89ecb867bb76dc0aa5c370f0ccf3c55e772a399e732ffb07705f7f93f52ce72) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-8e109be107ee0c88a350c359e650352b1fad75b93b6a89f7b2341c4971793d93) |
| `where.site.network_type` | [where.site.network_type](resources--advertise_policy--reference--group-001.md#canonical-68aaabd962f82d6a22eb36c633f78dfad854631e747a95ee1faa936dd112a4fd) |
| `where.site.ref` | [where.site.ref](resources--advertise_policy--reference--group-001.md#canonical-cc4f4d9ee002aae545297623609971f086fe762bb7f79033dfba38b372ec912b) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-6f0dda2b10cc2211e940756aa11e0f976ae83f959e87882e5f8467653ebe2c0a) |
| `where.site.ref.name` | [where.site.ref.name](resources--advertise_policy--reference--group-001.md#canonical-756f7fe9952e7252becc78c4a96f7c7b110e18c19a05a9eb85889f038af968d1) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-f7aecbb5d5554316c652781dcc8bb6886c3151cfd37ae96ef3890298688c9a46) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-06a125b38aa3c13bd790518ceb2e5d536bae7650bf58cc347ad5c72555525691) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-d689678160fb0a8222d53689761575ccdda3c2546c6a4d51bb23b8cb9d8085dd) |
| `where.virtual_network` | [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-89eff51ac09c925639bb5c0caabe32eb488e5ba7915f8fed461d5665c1b9ce51) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--advertise_policy--reference--group-001.md#canonical-0b7fab83ab388fc65b0253ffccd1d3ffeaa247fdcd558880d8c393c9bd0f709d) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-960382c63ed975c6cd7355705328e8f94c46ff8aef2505ceb4762b2d0aa86f9f) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--advertise_policy--reference--group-001.md#canonical-85aff7941c6bc1a9bc0469a44c3c7e645579fe909b1168482b63a0edd71ae7a6) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-972a7f0a88e41ff6a18f9e1f3465f8596d163bb66ae18f28ea127160369df5af) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-d60afe71144940ba683a280254c4e3faed8e2a4725b03ea050db4b72abc7aa38) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-7685d926af2f37c7d80c31b30f9eb0fe031758a66b8971757c6c0464ec0d37da) |
| `where.virtual_site` | [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-d56fe5403c1ea3067df5e372451c9b22c32c09d3cd6f914f64caed5abadf4c3b) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-0da4170db352dc1a1a4030ba52b1e230a3ab4545e530c659f177b1a0310b7d5c) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-32b84dc60c99a5fb9579f61f6de44031f9c7affb8f42ec28e0f9837f633ec936) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--advertise_policy--reference--group-001.md#canonical-66ac43ecaf3263fe9ea6dec651d83f3f7f44358f78ff224366884c585cdfc6eb) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--advertise_policy--reference--group-001.md#canonical-ac0819755bae21b8dc46e194767ebbc439ff179ece5e34a10ce87bb804d81908) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--advertise_policy--reference--group-001.md#canonical-e574a2c04fa9174a801f6bc0ec48678d14e4d041a326a8d1497d3a4c445f0107) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--advertise_policy--reference--group-001.md#canonical-737297559d7ca54fea2f517e0f2e6581cde745c9dcaf94e15a6bb0a5678b3ee8) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--advertise_policy--reference--group-001.md#canonical-9fa8f5cf1f4fcdfee2884c66d2906c41e0f87660104ced2ab1553788790350f5) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--advertise_policy--reference--group-001.md#canonical-da28c1b814541d5abdcbf2f2fe63e4cb356d5d22def2b9f16468a2bfe1631b37) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--advertise_policy--reference--group-001.md#canonical-771f9dc7ca509e37c1ad6fc35ef0724a4aaec7c25c562bb67093c862b1427607) |

<a id="canonical-02ef879aed44496934558359d4e61d1cd31fa3fd996b5a51d6cd5fa827cc437b"></a>

## Next pages — Property reference / d121f1992063 / 17

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-dc6d846a688e96e876ec4bb21eb0adf7eb7d22f5ad4753352f7a7003e78b5079)
- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-322e87a9f7c05a42845c949dccc855f050551a4586253edc350fc2a65685348c)
- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-988e666ede6f194642cb5fdb4ae9db08ae4e80f89ea4d03c2a1d1debea52f26a)
- [public_ip](resources--advertise_policy--reference--group-001.md#canonical-b3c548b5aa3a8348d9cff3799377ea9fd555cd8938e9759c3199ab15ab8efe86)
- [timeouts](resources--advertise_policy--reference--group-001.md#canonical-43f161effa9d2d50609bf9656eb56f7c91b255f77e6da0c988f8f6aeb3ff54af)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-dc6d846a688e96e876ec4bb21eb0adf7eb7d22f5ad4753352f7a7003e78b5079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e945a9bae836cd5414cb5fd5496956f1a8593c0e9a7a407b67c396b51a99dc2d"></a>

## dualstack — dualstack / e02e9e773871 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- dualstack

<a id="canonical-abca3e3112c23f36c03125bf3813d6b797ef2a2fedce8bc1acac85454aa02682"></a>

Type: `["object", {}]`. Optional.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

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

- [dualstack](resources--advertise_policy--reference--group-001.md#canonical-abca3e3112c23f36c03125bf3813d6b797ef2a2fedce8bc1acac85454aa02682)
- [ipv4](resources--advertise_policy--reference--group-001.md#canonical-ff3c64084adc23b304c255e203a0a1d1f8f7c686147af8597dcea4b5a8c02590)
- [ipv6](resources--advertise_policy--reference--group-001.md#canonical-8063764a258576f3b1041c7967150279d747250bf67d8ae7ffc6e2431269245f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dualstack = {}
```

<a id="canonical-f4a8a8c212b5bb9b27da607df4b898fa2a7865e1fb35c6163ff1335dc94fd480"></a>

## Direct properties — dualstack / e02e9e773871 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-264732976d45dadf1c51c4196bf9c55893dc1e153b0978adffd08dce97404882"></a>

## Next pages — dualstack / e02e9e773871 / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-322e87a9f7c05a42845c949dccc855f050551a4586253edc350fc2a65685348c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ab9973afc5359140741fc68ced47cf605e3750fdfe25a0a09829bb7b715d876"></a>

## ipv4 — ipv4 / a27982f3a7a7 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- ipv4

<a id="canonical-ff3c64084adc23b304c255e203a0a1d1f8f7c686147af8597dcea4b5a8c02590"></a>

Type: `["object", {}]`. Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

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
ipv4 = {}
```

<a id="canonical-ea9125934d29d939226f2fa3424cfc34ddb1d1a027c5b228eacae95a9289bf89"></a>

## Direct properties — ipv4 / a27982f3a7a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa38d24206d697fee3e0acd5a2d78e9baefa60f9dbf71c1bb874ec4bed47483a"></a>

## Next pages — ipv4 / a27982f3a7a7 / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-988e666ede6f194642cb5fdb4ae9db08ae4e80f89ea4d03c2a1d1debea52f26a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64e2bf8068080dea6f6d5a61099947336a97c89aceedfaab937860e9938ffab9"></a>

## ipv6 — ipv6 / 5d1f18f50bcc / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- ipv6

<a id="canonical-8063764a258576f3b1041c7967150279d747250bf67d8ae7ffc6e2431269245f"></a>

Type: `["object", {}]`. Optional.

IPv6 address in colon-separated hexadecimal format.

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
ipv6 = {}
```

<a id="canonical-0c6bdb32b3145e259422ffd3f44ec3fd02bc25457a71f87bbf8e25bf15c35385"></a>

## Direct properties — ipv6 / 5d1f18f50bcc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2dac032b0ecbcc9b1dc651dc3f82fbdd5d49b8282112946af347d3d99a5b5b25"></a>

## Next pages — ipv6 / 5d1f18f50bcc / 4

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-b3c548b5aa3a8348d9cff3799377ea9fd555cd8938e9759c3199ab15ab8efe86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-942274330a08c233e1254d2b8fa39fed576dc92499d226f2080993e9ff4a7941"></a>

## public_ip — public_ip / aca5f4c846f9 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- public_ip

<a id="canonical-3ab6f41bb4ff56cf6c012c2e7c3e7f2a7fb8376c026c373dfa026962dbb13813"></a>

Type: `"object"`. list nested block, Optional.

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

Upstream description:

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-fd92ff7f2c583210d7321d47fc9d5f2d1e6060c60eefb91115b3ba741cda44dc"></a>

## Direct properties — public_ip / aca5f4c846f9 / 3

<a id="canonical-8036e915b82d1ee4e2f9f564aa1961d5d784f22570505fcfafda98f5ff4f9822"></a>

<a id="canonical-147a2c971bd426173642bedf40b11e4007ed728dc94fc88f23da72315b936e8f"></a>

## kind property — public_ip / aca5f4c846f9 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-94b1db7c3c88d5c6d08234b962a8ea701f4c03b580503a346abafd7589d59610"></a>

<a id="canonical-a838d906edda79bc5b9d88378f1b7d529e5772373716c763e3de6e1d3b8078f1"></a>

## name property — public_ip / aca5f4c846f9 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-bf413663821d9d27e34778e6b66e5e791d1d1eaf4fd6cd2d113dbb40a735a40c"></a>

<a id="canonical-bcf23c95ee3d2f631d5da4c073f33bc03e7eb9de005fb0b5b781ed7e211cf33d"></a>

## namespace property — public_ip / aca5f4c846f9 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-5bd6e6253f2d4b6696d209e9923547fb87cb3fe661f52c7dd79809052b3a4c98"></a>

<a id="canonical-87ce8db4fd030b6b82678acf89dc36bf67b1c9d43c40486b4fc00a3879583456"></a>

## tenant property — public_ip / aca5f4c846f9 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-36373ad6e4c95d30b4eb7331e8764c30dc3385ed3955ea3dfbda0a50725766ca"></a>

<a id="canonical-a082a073771cee4b4a0be0c31037eaa148b6c4cc8c6a7523d8ac06eb22d7017f"></a>

## uid property — public_ip / aca5f4c846f9 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2cbb21081c3cd48c1266d2d12e877870edf31360787d70f6a223cada99182907"></a>

## Next pages — public_ip / aca5f4c846f9 / 9

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-43f161effa9d2d50609bf9656eb56f7c91b255f77e6da0c988f8f6aeb3ff54af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b1e96d7c561fd389eef0bcf55d53b3e9757ac47fc34eeab0647b762add42eb4"></a>

## timeouts — timeouts / 9f5714a8f6f7 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- timeouts

<a id="canonical-82dcd1a25cfe842fc2922d447732b7c81e2def103faced0af84d99ceb9a7adee"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b46d04bf71817e52cb0af31a61aadba9bd679efdd2852cbf5dc3a61eb9ee24aa"></a>

## Direct properties — timeouts / 9f5714a8f6f7 / 3

<a id="canonical-2dd9df4c8edfbe33bf842efd77aa54198d1320639d9f2443b3989ecf9714d487"></a>

<a id="canonical-4165d68f49b8c1e863a01bf1fdc2938d0def47a3efb9206fa9f2fdbf1f452995"></a>

## create property — timeouts / 9f5714a8f6f7 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-54c5490b99db02b78c6933a241d47d4e1980dfbdbad7223a0b85e9173d4dfd32"></a>

<a id="canonical-4ca48a1af5835b457c9a6384e456ca663832990ae60b77dfeed5be7f0c1685aa"></a>

## delete property — timeouts / 9f5714a8f6f7 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1c21c07ecaa378ca64ee45f8d178b2f0adff2c18baace0e8741c1913b20814f4"></a>

<a id="canonical-16cdca2dd3a92b8743cd8aa41eba800d4bcb2d4c55f5de3ca2bdc8e3be3dda30"></a>

## read property — timeouts / 9f5714a8f6f7 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-09a1d58ed82884688812c42cce69b0c22248d6fc39b9bc563c46413db42dea90"></a>

<a id="canonical-21be6cda1c6fbc8fb31692f78ac13c883e88c1c10637d18234d0176d6a79a3e3"></a>

## update property — timeouts / 9f5714a8f6f7 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d0db47d97f13029bbf85a58fc59fabb64d2c0a03075520b58fc72ab2f95d5bf4"></a>

## Next pages — timeouts / 9f5714a8f6f7 / 8

- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80325551191acc970a409e058feeb1cab99499ff0404d7ea936ac0ceaafbbccc"></a>

## tls_parameters — tls_parameters / 4948a1a4b1a7 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- tls_parameters

<a id="canonical-ba3f99861b8f62bbb57f7b244f2500c6bbdbb205dd30bc3200977bb887051ffc"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("client_certificate_optional",
    "client_certificate_required"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "no_client_certificate"),
  validators.ConflictingObjectAttributes("client_certificate_required",
    "no_client_certificate")}
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
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-58e7b42e0fd09bc80b095d4a8179f3085e142d8f1feb23edcf7b219f129ea3fd"></a>

## Direct properties — tls_parameters / 4948a1a4b1a7 / 3

- [client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-c5940e5963942d0a719b7657785b4182850d7369d2cd76ec53ceeaf59eca6193): complete subsection reference.

- [client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-9b8342284dd36cefe291177f9b6d48e6c8d95a58f3176194ed7b2a2d8e772082): complete subsection reference.

- [common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c): complete subsection reference.

- [no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-5b5c89d91df5888df611b5527ef5deb8dfdfaa14d36041f4a1f8b6d48059a7c9): complete subsection reference.

<a id="canonical-6fe0d7df1fd169183dbfcada617cc775348a243312ae443a234042053a4a2920"></a>

<a id="canonical-eb6d21a0487925ffc11785a0f594e8a85f0ae28e3d5d12d441481f80eef62bad"></a>

## xfcc_header_elements property — tls_parameters / 4948a1a4b1a7 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b01ed089d419ea9ebeba10c3f6ddddba68b42a657f5d43c7a635d0eea186324c"></a>

## Next pages — tls_parameters / 4948a1a4b1a7 / 5

- [tls_parameters.client_certificate_optional](resources--advertise_policy--reference--group-001.md#canonical-c5940e5963942d0a719b7657785b4182850d7369d2cd76ec53ceeaf59eca6193)
- [tls_parameters.client_certificate_required](resources--advertise_policy--reference--group-001.md#canonical-9b8342284dd36cefe291177f9b6d48e6c8d95a58f3176194ed7b2a2d8e772082)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.no_client_certificate](resources--advertise_policy--reference--group-001.md#canonical-5b5c89d91df5888df611b5527ef5deb8dfdfaa14d36041f4a1f8b6d48059a7c9)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-c5940e5963942d0a719b7657785b4182850d7369d2cd76ec53ceeaf59eca6193"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2185a9604c2d6846c2067e210019a223a337ddefcf57ef5d4bd6a504b0a5ae45"></a>

## tls_parameters.client_certificate_optional — tls_parameters.client_certificate_optional / 895cc1a4a06c / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- tls_parameters.client_certificate_optional

<a id="canonical-ac5aed29bee713f23e41a9e43822ba72b04ad5211009fc055155e33aefd1982c"></a>

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
client_certificate_optional = {}
```

<a id="canonical-b08b3747ced84e7db427f926e1c76c1e9849ee401f8c783d371ef94fc7a923ff"></a>

## Direct properties — tls_parameters.client_certificate_optional / 895cc1a4a06c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc5bebc82be161ab1795ead637862edc7cb516bb74aa16ff4fe309e9c48d77cf"></a>

## Next pages — tls_parameters.client_certificate_optional / 895cc1a4a06c / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-9b8342284dd36cefe291177f9b6d48e6c8d95a58f3176194ed7b2a2d8e772082"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0735f99dcdc63da7c6536e89fc1dae77ab7d96f03e9d8921ef6036d926b9e8fa"></a>

## tls_parameters.client_certificate_required — tls_parameters.client_certificate_required / 5dc2f97bcfe0 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- tls_parameters.client_certificate_required

<a id="canonical-f17ca6a37f12d48f846627c7229e2c4684d6e11160da0dadfd02dacb57b0a7cf"></a>

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
client_certificate_required = {}
```

<a id="canonical-347683f59e66968b55d1e63542557d96dff38fc33e638c53d179a784edd41d2a"></a>

## Direct properties — tls_parameters.client_certificate_required / 5dc2f97bcfe0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6782daf3c28495140d67f29d0c0b93bf7eb3056ae8f897e15e879c544242a4ad"></a>

## Next pages — tls_parameters.client_certificate_required / 5dc2f97bcfe0 / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c9914e90553ebffda78f6db8619fb3d1a89d415bec070bcfe99d9f93407bec"></a>

## tls_parameters.common_params — tls_parameters.common_params / 7eefa3a78673 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- tls_parameters.common_params

<a id="canonical-78553486a97769af0d0cf6292c63a02841f78d3b92fd08a3bd0275229e69b1f1"></a>

Type: `"object"`. single nested block, Optional.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef2c8b4b4565aeb2f3c25e7bd58be445aa76f1fd04f268e0a35a490e76a23baf"></a>

## Direct properties — tls_parameters.common_params / 7eefa3a78673 / 3

<a id="canonical-445d1a45d9a92eda5415062a350e6c4b3f071d25e98b06b91e62cd7fbde23ab9"></a>

<a id="canonical-9c7bf0a18b6941199f866cedf68a2912753f1e1bd12cae83390b39f1641c1ed1"></a>

## cipher_suites property — tls_parameters.common_params / 7eefa3a78673 / 4

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5d02c076de7935405fec394711d6d721a253fd96eee8915a962057fa58067145"></a>

<a id="canonical-565a10fdd522091f82e041f763cc1196c1c6b89a49a638e1363e85b6873f1854"></a>

## maximum_protocol_version property — tls_parameters.common_params / 7eefa3a78673 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bd24bf70683e71f1ff9763d2f1a63686516ab9a2045cace9a74bd4a59cc77003"></a>

<a id="canonical-c061b3c710471a6188f1a8f816b3d6b1a54c07c971f749bb55b079bab81e2f76"></a>

## minimum_protocol_version property — tls_parameters.common_params / 7eefa3a78673 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f): complete subsection reference.

- [validation_params](resources--advertise_policy--reference--group-001.md#canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2): complete subsection reference.

<a id="canonical-ab58f4cb559afa52f34610b0ce39783df1ba8c4b1e3d3efc7da6a41c9cff50c6"></a>

## Next pages — tls_parameters.common_params / 7eefa3a78673 / 7

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-248ca42a9d55705122e29fc421f78da7b83f819c47fa57286d9927552843d935"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / d0ce32233ec8 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- tls_parameters.common_params.tls_certificates

<a id="canonical-036a51231b3596e1a2ad06ebcb22428fd7690987af13bc2862f88f84df3d14ac"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3320b277b618b0651c1a6c20564f508089bedf18ecc50d1e5dfe42b4957cd66"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / d0ce32233ec8 / 3

<a id="canonical-e72c0c0e743443fe319ea2dbafec7cefe8fa106184b2efde3d14f08cca077d79"></a>

<a id="canonical-71ba0eba99ab59712552fe6ed8dc6e55f90e1f37eed69c08743dac1e6ad4ce67"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / d0ce32233ec8 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-feeb3e6cdd7f7bb776c511cd48580e1c76a5455ad93fff544af7d3d6b8dfd9df): complete subsection reference.

<a id="canonical-467bb0992b6ac6300a92018e05b09bb11fa4ca56d39dfee55d07d71d0ad296c4"></a>

<a id="canonical-e820883276ea8c125302591a98a835473f97c75d41399e15d40832ccec33116e"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / d0ce32233ec8 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-62f4040f4865f654b2fafac0592bd05e269b88d7f1eaa1764db57a24ec1b8e34): complete subsection reference.

- [private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2): complete subsection reference.

- [use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-505c421018457efa7ae21289fdf86e9a8226ac6d430121ebbb919f092b247c16): complete subsection reference.

<a id="canonical-7d8bd447fa9c3f10ae86d76abcca4259aa2feb2ffeb756dda8f4071f70a42a60"></a>

## Next pages — tls_parameters.common_params.tls_certificates / d0ce32233ec8 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--advertise_policy--reference--group-001.md#canonical-feeb3e6cdd7f7bb776c511cd48580e1c76a5455ad93fff544af7d3d6b8dfd9df)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--advertise_policy--reference--group-001.md#canonical-62f4040f4865f654b2fafac0592bd05e269b88d7f1eaa1764db57a24ec1b8e34)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--advertise_policy--reference--group-001.md#canonical-505c421018457efa7ae21289fdf86e9a8226ac6d430121ebbb919f092b247c16)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-feeb3e6cdd7f7bb776c511cd48580e1c76a5455ad93fff544af7d3d6b8dfd9df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3138c049d49b39303cd9cef0675cec4cd4760db467263357fac4b3abece9d2d7"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / e987d992a260 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-bc8ba1de8503b8e4eb768edc4ab37ab75e472c9b38e6a52c8bfd6d286d262634"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2c968dfd39a05f407e28591a97b53fb84d6d96a3c77ff864aa3d42a07b21b02b"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / e987d992a260 / 3

<a id="canonical-c4fc211e86779274c0af80af7a4cf7403ddb9eff0059f0d888408b314954d1cc"></a>

<a id="canonical-0ef7eb8f05ec6249f476c34ab77770d806bdae2094e57334a3d0f9d3099c08fb"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / e987d992a260 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cb4d3026b312c0bae9a405b0c91c916056b969260b38b360c9c14fb622e81ab6"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / e987d992a260 / 5

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-62f4040f4865f654b2fafac0592bd05e269b88d7f1eaa1764db57a24ec1b8e34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ede1d3c63d405da37f7a9d5798392f6f4aae3eb4dcae44c0471b1108af30cb82"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 05b0874f70ae / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-d1e3230eb1c51ae3ad771209b9066b3ea5ff2aaeb82e3097ba29634cbf640b0c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-4d570a0f02312902eb226708bf3e2eba7da98774e97bdb4a1ced9c226dc02731"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 05b0874f70ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa1dd52dba93798e84ddcb07b53981964663aa01ed79deeeed841cf5bc28e0e3"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 05b0874f70ae / 4

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93d5c3de0cf96a839bdd00aa6e5c6209bed08e6a2de256d6a4dfcbed33280c54"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / 04976f4aa150 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-65f6d614197e97d71f5dc85eff9fc20a47b63a464f15cdf096bb6b8952026781"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c1b482e3fc9e4875a4f982751792d969101638a3803ad87e05c5b2ed620c409"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / 04976f4aa150 / 3

- [blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-9b7563668a33d339fcdcc1dbf6ffea949ef78e3d3fd0bb14872846737128aed4): complete subsection reference.

- [clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-35bcc458f119a46541944d0fe2a913785c145ef2bf77a3c0eaf198b08b23a941): complete subsection reference.

<a id="canonical-fecc9f9d677d417ad0af30f317c1072f98efbf4c5484deccd2d01935e01b1123"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / 04976f4aa150 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--reference--group-001.md#canonical-9b7563668a33d339fcdcc1dbf6ffea949ef78e3d3fd0bb14872846737128aed4)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--reference--group-001.md#canonical-35bcc458f119a46541944d0fe2a913785c145ef2bf77a3c0eaf198b08b23a941)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-9b7563668a33d339fcdcc1dbf6ffea949ef78e3d3fd0bb14872846737128aed4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32e3521c5ede9d395783ea0f0c59d13621345f8530fe69a1b551cf54fb0a199c"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-088e8fb7b82928465a58fc7b359cd68767794213741586cc5ae24e93a2310745"></a>

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

<a id="canonical-b22fcd75207a72eed8e79ce37ce2d51ae80501d8f72a8a9df88df582c038195f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 3

<a id="canonical-0344e8541f3aa3cb3cccf8d8e2c36c5c9e6ba223eddecd59c66e0f80f92d1de0"></a>

<a id="canonical-18f38352e58bee46e57e81a38f2938190f6e9d3e0d795d560dce38172277614e"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 4

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

<a id="canonical-e30125bd2b140d778fadbff1462d7fa0808ba55bc74e41302c951f258b1e3569"></a>

<a id="canonical-1c89ae001bd74856fdae8901d87e02de57b7b04267395c741675d05abd0f537c"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 5

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

<a id="canonical-7bb1801cebbbab6c44bb4bd8eebb016a61056e165ff103356e495d12427b805b"></a>

<a id="canonical-36f850b776302c185babe350c0a75fae4010b117dbf0ca2c63747638054521dc"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 6

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

<a id="canonical-9ebf552b0a11c836616358bbac779d64ac3bea05c98ec8dc415e1d691d9fd2fb"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / a1829f12b27f / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-35bcc458f119a46541944d0fe2a913785c145ef2bf77a3c0eaf198b08b23a941"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb31ed28f6f2fd8aacf3282d9d511b1fac7ac40aac47fd6df30c0a72431bb8e1"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / f9e581ec30da / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0471ea66d19eeb0eb26bfdda83644dd344ea7aee65d395b55b739d5fc258ded8"></a>

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

<a id="canonical-13e63a530346cf83bba0be245011871ade77d21ee869c09f1946edfe17b58693"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / f9e581ec30da / 3

<a id="canonical-2da7f872eaa1bd328273021579e9a22e7eca208f08e259101bf798ecbd11f4ed"></a>

<a id="canonical-f6f3583cbf726bed3b60d9ea117b17d7207e757772b0ea2a1ffded595862286e"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / f9e581ec30da / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-08ef70ccf34736f3336cff20c48ff58cdb1c93afcbf09ce5128a8259335d2733"></a>

<a id="canonical-21fea3606f019bbc57652a3ae9a689075622cc69f208b14a25405634323c6b71"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / f9e581ec30da / 5

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

<a id="canonical-c554f645416c4ba6fb8fb4dd50fd277e86fd302ad70bcec4e0c007f0da0c408d"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / f9e581ec30da / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--reference--group-001.md#canonical-116bd677452b961865c66c5d9f73215f6d03f7e2be07927c40d1c32c595a97a2)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-505c421018457efa7ae21289fdf86e9a8226ac6d430121ebbb919f092b247c16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4233a82f4140b8580cd9c2902af4e1b8023f65fa05633f710139c0ecbc36fcdc"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / bae46810224e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-9a1660f6c4ad29e7cf5ec131451d458c598ddb9730db1655697b71b2bfb475d5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-321842c9d7a418eb12f3c4ff53367d70c2de2ba9e81cd1d24bb4328a0719b678"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / bae46810224e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f011c1cdd6ae1890cdb4f971180d0b201cfa80fc8cacda674d825bded8723456"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / bae46810224e / 4

- [tls_parameters.common_params.tls_certificates](resources--advertise_policy--reference--group-001.md#canonical-6946e0e095c6ad4ff9abeb7abd08aeec203353132a3a61a89f6d8f838026745f)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8fddbd124a1f3a2b4b5ad58ad198cd3a61a9186ee456e87b73f0e02775c6b06"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / 0a63dd2d64da / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- tls_parameters.common_params.validation_params

<a id="canonical-0bd5d3978b4b708c7627cd4653b26be1fc6ee05df922a795162bae4f1074aa86"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-c4a840f4056c90917905bd867a073181c416fc73daeb1d9088c6ab3aabcfbe02"></a>

## Direct properties — tls_parameters.common_params.validation_params / 0a63dd2d64da / 3

<a id="canonical-91cf0f4a2629e57f039c9f15e98190b09f553c5f820de2a55bc052a637aedb2a"></a>

<a id="canonical-9b7f51854e6ca9ebe545a58f3fe8ac7cd7651262e89d40b5100c868f0fad50ca"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / 0a63dd2d64da / 4

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-2ff11f015ed86e4253c3a84da55969a5d83b49dbb396c67c6879f0de4136d462): complete subsection reference.

<a id="canonical-d01f727a3d9013e04ca315bfe4cf10e501a35d0215c8922e2b3f395e705ee40a"></a>

<a id="canonical-06f5cec9a9aff0e67daf6b258bcd74ad6c7535fb735730c17106eed3fb960735"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / 0a63dd2d64da / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-48013132d0a028faedc7c5fc5cfaeef5d11afa6797cb9c0ad0f4b7286328fabe"></a>

<a id="canonical-e1060747761787bceec69e911dbb07f0633c3fa8e82f664f9041c99126b3cb60"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / 0a63dd2d64da / 6

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-26b0ab82a9ab3393fbe09eaa68d234a9f12c02b816b42bd80e8c545628cb4f9a"></a>

## Next pages — tls_parameters.common_params.validation_params / 0a63dd2d64da / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-2ff11f015ed86e4253c3a84da55969a5d83b49dbb396c67c6879f0de4136d462)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-2ff11f015ed86e4253c3a84da55969a5d83b49dbb396c67c6879f0de4136d462"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1202b42e832fa57a53ffd6e5efa7fd5bd78f9eca6a5a254ccf7261a0266bd7b"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / 771068be338c / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-2cd58d2f3d9c71c5eb1117c9a8bee021eaf3c5ceef5500d80643e02a329f0f13"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-af7ad60a54b712767c7a602105f56e15251ae3181f990786520eb83a5e7f1510"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / 771068be338c / 3

- [trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-0ef664ec4da3f6241f93d3d168c01862d2840ad1b7174f02261767ab2215d5b2): complete subsection reference.

<a id="canonical-20fdbdfc10abb9e92eeaf39624e9ee74711b4e423aef18757d135cdd20ebb2fd"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / 771068be338c / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--advertise_policy--reference--group-001.md#canonical-0ef664ec4da3f6241f93d3d168c01862d2840ad1b7174f02261767ab2215d5b2)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-0ef664ec4da3f6241f93d3d168c01862d2840ad1b7174f02261767ab2215d5b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b17b0e71b9fcc2edda2889cf876a1ec96fecddf97372af32f5fdd3426d5f142"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [tls_parameters.common_params](resources--advertise_policy--reference--group-001.md#canonical-3724c3bb1464307485d45d50419ba17695acf70ca30f49c80ebe9a9c4f883d8c)
- [tls_parameters.common_params.validation_params](resources--advertise_policy--reference--group-001.md#canonical-58c99e4504ec180896ccda0cac7683a73c152ff6dee551c3eb53518a27de4dc2)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-2ff11f015ed86e4253c3a84da55969a5d83b49dbb396c67c6879f0de4136d462)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-df1fdb9857d77330136c66bd191b824762f1606f5cc16d0f15680aaba53d8bf1"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-d02762082e7a21e688357fd1136f0f7929f72f72165d939aa9a4dce65a050d61"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 3

<a id="canonical-6b586509dbc5b24f253f73adea82a48341f935879d6ee633b95bd648520d1291"></a>

<a id="canonical-6b0449a22b16501767eaff52bdbab701503543ab6cb1970e358e5e55d3483b37"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-df9fbc2ae1a684c6b7f01f2d516d94da35914f8658e9898b66b1c3adc87b223c"></a>

<a id="canonical-6d30e24860c03e052991170627446de211c2ee9ae77dad1b582c6d7b8e0b2a33"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-98c3fd8e86986a761f97366542217f29514ec1b955553563ce2395ddad09d904"></a>

<a id="canonical-ef098cd0528306a8b2c30a8ac126510c890c306192af65fe8d6159dc2e421d66"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-64ca1633cb781045f98e1be1380773aa3eb4998d5abdd95b8b24f3372f72afee"></a>

<a id="canonical-85b8b9a84c892e95cbd3c4e87fc3e46cb23b13f12d3eadf257b8d3fc7c850b3b"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-ab86cf97ff9a782ec218ba8d8a9c6083d58abae044b7f5738d8642afeb3415c6"></a>

<a id="canonical-6d71c860012148e040ccb84e7ff511de9ac8de1cdbb0a23e86d25a5f07aa276c"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-b9dba8b5392bfe03deebf2f7ee425bceee90352bdab8b9e49c120f743d739d75"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 901ccbf60b2e / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--reference--group-001.md#canonical-2ff11f015ed86e4253c3a84da55969a5d83b49dbb396c67c6879f0de4136d462)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-5b5c89d91df5888df611b5527ef5deb8dfdfaa14d36041f4a1f8b6d48059a7c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58dc6c45d97db76af7631b4452c49ae12ac633469834ef019e85562bf124e79c"></a>

## tls_parameters.no_client_certificate — tls_parameters.no_client_certificate / eead9716c42e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- tls_parameters.no_client_certificate

<a id="canonical-e126f3a35d654726dca8999e770ce24bf031a75bd37e88be4ec6fdf65f841dc1"></a>

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
no_client_certificate = {}
```

<a id="canonical-731d1860e6af5cd3e63a8495c39117a1168736d9974671263bae33ec60fc2327"></a>

## Direct properties — tls_parameters.no_client_certificate / eead9716c42e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10bcfb514f65fa348feec2c4512e0e1304e8665e27e28bd7ae821afa33820d35"></a>

## Next pages — tls_parameters.no_client_certificate / eead9716c42e / 4

- [tls_parameters](resources--advertise_policy--reference--group-001.md#canonical-3efd73f08ba44265d8edb3f4ae33ad99514642cda37532307c7b4b69882cc3fc)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23678c84109532664cb2c050fe081d06577107ceb446458e87f621ff374eb179"></a>

## where — where / 96d95b4149b3 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- where

<a id="canonical-b67e098c0f131e81a4df6aa2165a1ea4e1810e0c6cfd2fe2c2694d9777fba9e3"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc0a99fe4abb8b6ebb51646f3d93083af55b97b02c5f39e50a71a0054af58e40"></a>

## Direct properties — where / 96d95b4149b3 / 3

- [site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8): complete subsection reference.

- [virtual_network](resources--advertise_policy--reference--group-001.md#canonical-da205a514b2b5bf451bb2a95130cecea3bbe1b7bdf7bda9a31df3811592e9e7d): complete subsection reference.

- [virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02): complete subsection reference.

<a id="canonical-5e4753d3bdb59c683973825de3d98264746742e0407399d3624c659aea9e2c2a"></a>

## Next pages — where / 96d95b4149b3 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-da205a514b2b5bf451bb2a95130cecea3bbe1b7bdf7bda9a31df3811592e9e7d)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efe6cf5da3e9327d4e7b4514033f65e35c4c0fbf9fbe050373878fa08ffcce69"></a>

## where.site — where.site / b6609afa39c4 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- where.site

<a id="canonical-7405e90b6286dc2dd5c5fbf44633668431c12701992b7de2e2917a003bfcc34f"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-e187a12947b4f2b32d86a025c112d538e6b0ed56740d1a498ce44b0c231d79e0"></a>

## Direct properties — where.site / b6609afa39c4 / 3

- [disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-7809a09efe6aa675fdd1a5fe14c5a02a80e328187ccc18448ee613ddb748401a): complete subsection reference.

- [enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-bc9019adc559154e4e87307127a0fcd94251bf7c83727b46fe4ea7c1c6d881be): complete subsection reference.

<a id="canonical-68aaabd962f82d6a22eb36c633f78dfad854631e747a95ee1faa936dd112a4fd"></a>

<a id="canonical-503c1ac51ccd10ba3802ef2fdb9dabdba810bc77570374c3ae4b6cfdb8654c61"></a>

## network_type property — where.site / b6609afa39c4 / 4

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

- [ref](resources--advertise_policy--reference--group-001.md#canonical-ed6dd73d9a80e55a78770291e4fe3ad62ec867ce8114f50ccc45b7c29c73d379): complete subsection reference.

<a id="canonical-bfcf91bc6714824b5259054f5e0e8df578d5ba39a4b71b592fc2aff092a19bbf"></a>

## Next pages — where.site / b6609afa39c4 / 5

- [where.site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-7809a09efe6aa675fdd1a5fe14c5a02a80e328187ccc18448ee613ddb748401a)
- [where.site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-bc9019adc559154e4e87307127a0fcd94251bf7c83727b46fe4ea7c1c6d881be)
- [where.site.ref](resources--advertise_policy--reference--group-001.md#canonical-ed6dd73d9a80e55a78770291e4fe3ad62ec867ce8114f50ccc45b7c29c73d379)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-7809a09efe6aa675fdd1a5fe14c5a02a80e328187ccc18448ee613ddb748401a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbb6ccbf6d7fbc6f4b301aaa4445fa886074ea475c9b1d8d1a32bcb7abcfe496"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 2935f16b6c51 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- where.site.disable_internet_vip

<a id="canonical-a89ecb867bb76dc0aa5c370f0ccf3c55e772a399e732ffb07705f7f93f52ce72"></a>

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
disable_internet_vip = {}
```

<a id="canonical-3355daa1a1c71d603b8c8e6d43229ff8c675f86a04429fa287ef7bd64386e440"></a>

## Direct properties — where.site.disable_internet_vip / 2935f16b6c51 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-290a1058a7d1210bcfc352a342083b454d16ef242e250d4c8ac8764d83a9a262"></a>

## Next pages — where.site.disable_internet_vip / 2935f16b6c51 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-bc9019adc559154e4e87307127a0fcd94251bf7c83727b46fe4ea7c1c6d881be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5e4e5ac4606d0ed186c92f587a6b88450939df8a28252d55259d93457c096fd"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / aca7ba0bc965 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- where.site.enable_internet_vip

<a id="canonical-8e109be107ee0c88a350c359e650352b1fad75b93b6a89f7b2341c4971793d93"></a>

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
enable_internet_vip = {}
```

<a id="canonical-8c9a8c6e1a5a394ec315dddbfbe65955674a4b21e246a0c536b3b355305ea1cd"></a>

## Direct properties — where.site.enable_internet_vip / aca7ba0bc965 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df8e377db057d2d0c1994d540949c47b5ffe3656e99e068c06a51283ae816673"></a>

## Next pages — where.site.enable_internet_vip / aca7ba0bc965 / 4

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-ed6dd73d9a80e55a78770291e4fe3ad62ec867ce8114f50ccc45b7c29c73d379"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8eab2dc1bf24d801e37853d01ca0cd74027b352c1305b437a17d424a69d77888"></a>

## where.site.ref — where.site.ref / b91d5c5c6c68 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- where.site.ref

<a id="canonical-cc4f4d9ee002aae545297623609971f086fe762bb7f79033dfba38b372ec912b"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4301ed4063d97be0c0c8002c9de8b3cb9afa1ed4506ce3927030c0fc836f563"></a>

## Direct properties — where.site.ref / b91d5c5c6c68 / 3

<a id="canonical-6f0dda2b10cc2211e940756aa11e0f976ae83f959e87882e5f8467653ebe2c0a"></a>

<a id="canonical-50821776b6dc484e4ab9526feb71400c7d6b81b4592dcdf0cca9e93f4330d9c8"></a>

## kind property — where.site.ref / b91d5c5c6c68 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-756f7fe9952e7252becc78c4a96f7c7b110e18c19a05a9eb85889f038af968d1"></a>

<a id="canonical-355bff49d8a8792ba82b3dcb3acd5719e5293e2177b4abeeb262d8ae18b00b74"></a>

## name property — where.site.ref / b91d5c5c6c68 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-f7aecbb5d5554316c652781dcc8bb6886c3151cfd37ae96ef3890298688c9a46"></a>

<a id="canonical-f91d1442393dc0908c6e8f51d1587457e7bdafe174f4fa97b198bf9af3631237"></a>

## namespace property — where.site.ref / b91d5c5c6c68 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-06a125b38aa3c13bd790518ceb2e5d536bae7650bf58cc347ad5c72555525691"></a>

<a id="canonical-7a2d0c13ca3992f68eb6b04a0df451890ce027165cbc98df0fe1834605dc95ae"></a>

## tenant property — where.site.ref / b91d5c5c6c68 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-d689678160fb0a8222d53689761575ccdda3c2546c6a4d51bb23b8cb9d8085dd"></a>

<a id="canonical-a5bd71119dcfd1bf032ef018346877701751c6adcd5df7b16a32078b41c4291e"></a>

## uid property — where.site.ref / b91d5c5c6c68 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-b06b80c69de83a7b8d47373443ed8fa0934776693aa31f5426616aea26376f62"></a>

## Next pages — where.site.ref / b91d5c5c6c68 / 9

- [where.site](resources--advertise_policy--reference--group-001.md#canonical-2a464c203b0ba62b750757e87eb744988bbbb8b36dd429b2a284e90bfdb2b6c8)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-da205a514b2b5bf451bb2a95130cecea3bbe1b7bdf7bda9a31df3811592e9e7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5234da9319dd1f78c1784bb53a115b4f0345938131b3b8bae13b6dcbc0ec0d9c"></a>

## where.virtual_network — where.virtual_network / 5aa099df69c0 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- where.virtual_network

<a id="canonical-89eff51ac09c925639bb5c0caabe32eb488e5ba7915f8fed461d5665c1b9ce51"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-42f5f7cf8ebb1a15624551178bcdce1f94f073d3d4c432ac8b1005f1ef61a36b"></a>

## Direct properties — where.virtual_network / 5aa099df69c0 / 3

- [ref](resources--advertise_policy--reference--group-001.md#canonical-78c1c6fc6968c2d2e4d9e9aa72299a7bab471d46ca44397b51c7cf6bad769022): complete subsection reference.

<a id="canonical-99e935ac0b1e8086f64bf7a12fdfe6991ac2078883fda0a13e3aac2f24f69679"></a>

## Next pages — where.virtual_network / 5aa099df69c0 / 4

- [where.virtual_network.ref](resources--advertise_policy--reference--group-001.md#canonical-78c1c6fc6968c2d2e4d9e9aa72299a7bab471d46ca44397b51c7cf6bad769022)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-78c1c6fc6968c2d2e4d9e9aa72299a7bab471d46ca44397b51c7cf6bad769022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c1df61d69e84ac3d0a3c17eb0b2c556096e7d93941a47b526e26fc4f29c8174"></a>

## where.virtual_network.ref — where.virtual_network.ref / 11344b0fa9c2 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-da205a514b2b5bf451bb2a95130cecea3bbe1b7bdf7bda9a31df3811592e9e7d)
- where.virtual_network.ref

<a id="canonical-0b7fab83ab388fc65b0253ffccd1d3ffeaa247fdcd558880d8c393c9bd0f709d"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc85612d5c022d00e5d9e74101d0dff72e79f2cb9c2e2fe05c9dbe8d404793a1"></a>

## Direct properties — where.virtual_network.ref / 11344b0fa9c2 / 3

<a id="canonical-960382c63ed975c6cd7355705328e8f94c46ff8aef2505ceb4762b2d0aa86f9f"></a>

<a id="canonical-ece90572a5563f2039833adba4f17ab70b08b47c6a0514089fc09c9e1b88598c"></a>

## kind property — where.virtual_network.ref / 11344b0fa9c2 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-85aff7941c6bc1a9bc0469a44c3c7e645579fe909b1168482b63a0edd71ae7a6"></a>

<a id="canonical-b12f6ff718de1bb91f57218b3fdd6488f882549859be7269a4e5c62b9143215a"></a>

## name property — where.virtual_network.ref / 11344b0fa9c2 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-972a7f0a88e41ff6a18f9e1f3465f8596d163bb66ae18f28ea127160369df5af"></a>

<a id="canonical-44e3f6a64ab8177383e8d3f37ee01dea04c53a047437ad1298e8f8a4aa5a749f"></a>

## namespace property — where.virtual_network.ref / 11344b0fa9c2 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-d60afe71144940ba683a280254c4e3faed8e2a4725b03ea050db4b72abc7aa38"></a>

<a id="canonical-3329565c6f7826107b82867198f5563c1ee8c5bdb5d58977e8920d992132bf7a"></a>

## tenant property — where.virtual_network.ref / 11344b0fa9c2 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-7685d926af2f37c7d80c31b30f9eb0fe031758a66b8971757c6c0464ec0d37da"></a>

<a id="canonical-52fc56626e55fc635c079b8d624364d42433f1b64eb8e806157d618abfb1dace"></a>

## uid property — where.virtual_network.ref / 11344b0fa9c2 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2b7fc638b143c39fee92b92cae848e9b5ec511ae33e342c51a91382cd2e0e36b"></a>

## Next pages — where.virtual_network.ref / 11344b0fa9c2 / 9

- [where.virtual_network](resources--advertise_policy--reference--group-001.md#canonical-da205a514b2b5bf451bb2a95130cecea3bbe1b7bdf7bda9a31df3811592e9e7d)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f792e4697d5650ff13ce5bc97fccb9896af0ac4cdd24dafe70068205d3a97f3"></a>

## where.virtual_site — where.virtual_site / 72e07f5d446e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- where.virtual_site

<a id="canonical-d56fe5403c1ea3067df5e372451c9b22c32c09d3cd6f914f64caed5abadf4c3b"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-17299bcec5d674bd2ca59980e131755db45c72b87b2483a06be92399abdddada"></a>

## Direct properties — where.virtual_site / 72e07f5d446e / 3

- [disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-7fc9294b2829e14da5c8af740d608bfda0208d9a017879580041848fa3b679a7): complete subsection reference.

- [enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-fd921648d1f92f62ac496fe6928080f7324a560f8fe46c54e97ce35b7286ac6c): complete subsection reference.

<a id="canonical-66ac43ecaf3263fe9ea6dec651d83f3f7f44358f78ff224366884c585cdfc6eb"></a>

<a id="canonical-9230644beeb5662b7a9db7b38235740347a12222915f0ab69184b478d6bde535"></a>

## network_type property — where.virtual_site / 72e07f5d446e / 4

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

- [ref](resources--advertise_policy--reference--group-001.md#canonical-3c7a687790aad91ac48b2ee26e9e8b74eb286d4192151c1bbd52d697d9ebc8b7): complete subsection reference.

<a id="canonical-62b97563ba0a22e8f3ea2fcd07696c85b9d3cb5ebfb61c1258c5daccfd61cfb0"></a>

## Next pages — where.virtual_site / 72e07f5d446e / 5

- [where.virtual_site.disable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-7fc9294b2829e14da5c8af740d608bfda0208d9a017879580041848fa3b679a7)
- [where.virtual_site.enable_internet_vip](resources--advertise_policy--reference--group-001.md#canonical-fd921648d1f92f62ac496fe6928080f7324a560f8fe46c54e97ce35b7286ac6c)
- [where.virtual_site.ref](resources--advertise_policy--reference--group-001.md#canonical-3c7a687790aad91ac48b2ee26e9e8b74eb286d4192151c1bbd52d697d9ebc8b7)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-7fc9294b2829e14da5c8af740d608bfda0208d9a017879580041848fa3b679a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-875dfaaea34f551144ff04bc5020d1cc5127d9efe8c45d112fa5b39c434e7846"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 27ab6ece7950 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- where.virtual_site.disable_internet_vip

<a id="canonical-0da4170db352dc1a1a4030ba52b1e230a3ab4545e530c659f177b1a0310b7d5c"></a>

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
disable_internet_vip = {}
```

<a id="canonical-dc6956c42659d0e7dcb69d6ba7de2eca1c68904dcfb23a6d53f57d05cb46bcd7"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 27ab6ece7950 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5700797855ffffc82133f5e2351e184150fc9bdffffb89cf063ce8b55d2d6415"></a>

## Next pages — where.virtual_site.disable_internet_vip / 27ab6ece7950 / 4

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-fd921648d1f92f62ac496fe6928080f7324a560f8fe46c54e97ce35b7286ac6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f16adbeeff466abe561b202c14ae0d6810092442f1ae10746b815fa9f529de69"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / 58492694aa7f / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- where.virtual_site.enable_internet_vip

<a id="canonical-32b84dc60c99a5fb9579f61f6de44031f9c7affb8f42ec28e0f9837f633ec936"></a>

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
enable_internet_vip = {}
```

<a id="canonical-81ec9e32e16c31332b2dd2e06eabc98b1c2bf51844ad64c44fdbdde5de12a83e"></a>

## Direct properties — where.virtual_site.enable_internet_vip / 58492694aa7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4cb0f5defafd0a39d7ae0a4fb73c0ba5ee11826aebc128723f202e447214c092"></a>

## Next pages — where.virtual_site.enable_internet_vip / 58492694aa7f / 4

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-3c7a687790aad91ac48b2ee26e9e8b74eb286d4192151c1bbd52d697d9ebc8b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-337796e549105f2d7160718691771d74956a1ed6867e47f87e51f780312a9321"></a>

## where.virtual_site.ref — where.virtual_site.ref / 0f9566791e8b / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Property reference](resources--advertise_policy--reference--group-001.md#canonical-b5cc12d2efbedb6a5c13e201fb876d994e92fd89722dbac50b1058709e65057d)
- [where](resources--advertise_policy--reference--group-001.md#canonical-5e2451c9413da4b5b509da871159db12fc431df4bdddd9487bd3528b469f7411)
- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- where.virtual_site.ref

<a id="canonical-ac0819755bae21b8dc46e194767ebbc439ff179ece5e34a10ce87bb804d81908"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c2bebbe0aaa7a03026a29bc34bd56584e101d882feabd1af395da73dc7cc402"></a>

## Direct properties — where.virtual_site.ref / 0f9566791e8b / 3

<a id="canonical-e574a2c04fa9174a801f6bc0ec48678d14e4d041a326a8d1497d3a4c445f0107"></a>

<a id="canonical-7f17a29649f21bc817513566769c040bf5b002259f9668b6698f852e5153de3f"></a>

## kind property — where.virtual_site.ref / 0f9566791e8b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-737297559d7ca54fea2f517e0f2e6581cde745c9dcaf94e15a6bb0a5678b3ee8"></a>

<a id="canonical-8718e311a6959adbef522214249d7d17e1c14de4705d2401623a2091fab8cf8b"></a>

## name property — where.virtual_site.ref / 0f9566791e8b / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-9fa8f5cf1f4fcdfee2884c66d2906c41e0f87660104ced2ab1553788790350f5"></a>

<a id="canonical-682bbc35742bec8e80bd8d77f41d8e614f693ec50c2ec2e78840c53411cd980f"></a>

## namespace property — where.virtual_site.ref / 0f9566791e8b / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-da28c1b814541d5abdcbf2f2fe63e4cb356d5d22def2b9f16468a2bfe1631b37"></a>

<a id="canonical-954f672a5a0e3ae3a2758e4389eee21087badaa5f216ed80b7022ffacbec02ea"></a>

## tenant property — where.virtual_site.ref / 0f9566791e8b / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-771f9dc7ca509e37c1ad6fc35ef0724a4aaec7c25c562bb67093c862b1427607"></a>

<a id="canonical-12c5dc38c288706736c7ecf5392c7a2ca9d2cae3e2a3b3c913939ad223074039"></a>

## uid property — where.virtual_site.ref / 0f9566791e8b / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-ebb3f11c626e54fec96505e8903b88ac36fe9c8f2890e2269fdaf9f31bf17f29"></a>

## Next pages — where.virtual_site.ref / 0f9566791e8b / 9

- [where.virtual_site](resources--advertise_policy--reference--group-001.md#canonical-c12885c07b6929391b5523639a194f4e6caad532bae143958f3a238c005e1a02)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
