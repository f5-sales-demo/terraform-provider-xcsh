---
page_title: "xcsh_endpoint reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint reference."
---

# xcsh_endpoint reference

<a id="canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0e65ab7d0481b9e30c2e107b20b3bd8b4dd1ae168ab326fa58b03aa1185a68a"></a>

## Property reference — Property reference / 8c453f7252bd / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- Property reference

<a id="canonical-366811980dd6bc2627584c3fa2bc2ce84ad0480e59e9dbee7abaa6c242744d82"></a>

## Direct properties — Property reference / 8c453f7252bd / 3

<a id="canonical-ca9d39a5e54245e7417b748ae708a7238daca1f6d05021830be7abe392f42c1b"></a>

<a id="canonical-2427b47b9748528ed0895d58bd8750f8f97499e4975b571bdcf6075b0b72816a"></a>

## annotations property — Property reference / 8c453f7252bd / 4

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

<a id="canonical-0852c210e6218383495c91b48d954ca421cfd007ab47828c4a354a067aac5f8e"></a>

<a id="canonical-1ece0c6138532cedbca00ae72f0c5a5f3cf03fa397fe28d187f5b0d9f18d78bc"></a>

## description property — Property reference / 8c453f7252bd / 5

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

<a id="canonical-b17ccd1212fedd917e7f09772be2e2fdba3afd62ff698db4547841d9e6474e21"></a>

<a id="canonical-c1bc38a8fd6e495fd8c79d5225accada3f5b3840eec537f52736f88628fcc9a6"></a>

## disable property — Property reference / 8c453f7252bd / 6

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

<a id="canonical-f405aa3e2203d5212ee63b5718c0cbf61738c92cb61bc0a7cdf675b5bfdd10c1"></a>

<a id="canonical-ede88a22ca13f371f34780368e0bfe8543ffef45659314ca4f1108e5bf26595c"></a>

## dns_name property — Property reference / 8c453f7252bd / 7

Type: `"string"`. Optional, Computed.

\[OneOf: dns\_name, dns\_name\_advanced, ip, service\_info\] Exclusive with \[dns\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Upstream description:

Exclusive with \[dns\_name\_advanced IP service\_info\] Endpoint's IP address is discovered using
DNS name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [dns_name](resources--endpoint--reference--group-001.md#canonical-f405aa3e2203d5212ee63b5718c0cbf61738c92cb61bc0a7cdf675b5bfdd10c1)
- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-6afe56c59488739a6bbf8b387e2d5a0635e0692edc0c5faa55f2abc62b4104ef)
- [ip](resources--endpoint--reference--group-001.md#canonical-6a10702f689c849631eef95a9786c2348328f2df35530823eeca86fe2861673f)
- [service_info](resources--endpoint--reference--group-001.md#canonical-48260fbfdf06ddd1b083110173b85b10cadeae09fc1b97fe74e263a4b1b877a9)

Select alternatives according to the provider validators above.

- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-cd11f7bdea51dbff2578a79d9945c86607d7fcc356f8e30f5cecab3b8fdcfb59): complete subsection reference.

<a id="canonical-d28d126a1b6c3cd1348123b97b0b72acca32625d411ca8ed1fab7df7b744d87a"></a>

<a id="canonical-977120ef7daa6bd6b42716eed0b8fa42c5b1e896bc985007a646a96488e270a9"></a>

## health_check_port property — Property reference / 8c453f7252bd / 8

Type: `"number"`. Optional, Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Upstream description:

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(65535),
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
    }
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

<a id="canonical-b3468be7b5f1c5bbb0c0cbffb7e89a0f8302ecb6819323addf76fccb2897c070"></a>

<a id="canonical-270946131ef27a45cda5878701088aafe540dd20b024ca5f32b31776cec6481a"></a>

## id property — Property reference / 8c453f7252bd / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6a10702f689c849631eef95a9786c2348328f2df35530823eeca86fe2861673f"></a>

<a id="canonical-bb6b675d3f568be69b148b44969eced152c8dbe00e0f7656583a0ea6867c4601"></a>

## ip property — Property reference / 8c453f7252bd / 10

Type: `"string"`. Optional, Computed.

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Upstream description:

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

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

<a id="canonical-209959b20d5db8e10bc787813815d0097f02a16b7e97f560b098a76148602a99"></a>

<a id="canonical-f19cb7d8f6f231a99f9283a74bc8b856dcd40889562e244535647be3f25d31c5"></a>

## labels property — Property reference / 8c453f7252bd / 11

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

<a id="canonical-e9f063a59f7f8f18f4d20d2a4a9afa9cfd9469506dfbc18369acda5712ff52e6"></a>

<a id="canonical-aadf0331a06bdce12234b94a81a5c45957d7fce912bd3406ab52920db492bbed"></a>

## name property — Property reference / 8c453f7252bd / 12

Type: `"string"`. Required.

Name of the Endpoint. Must be unique within the namespace.

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

<a id="canonical-85e6d18827a7c56c706a3cbe1ca46a645683439720424d1dd7d86d325410c941"></a>

<a id="canonical-8e1779b468c3d637d2e93e6afcc78378af8c2d49049ea2dbef2341efc4a51b51"></a>

## namespace property — Property reference / 8c453f7252bd / 13

Type: `"string"`. Required.

Namespace where the Endpoint is created.

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

<a id="canonical-cde649e581eed2284e147bc6c429adcc81458924da908a54072d50247b7aa7e3"></a>

<a id="canonical-64ffc10efcade31d288ec422750fe77624be86a973d78b9bd4b51e16ef48d25d"></a>

## port property — Property reference / 8c453f7252bd / 14

Type: `"number"`. Optional, Computed.

Endpoint service is available on this port.

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

<a id="canonical-6253a9e23707bd09afa66a76ea6ed5d0525d5a2945bc0664f5962de66760c9ab"></a>

<a id="canonical-c000a70231c7c4bb25d0458656dcf5cb46badf548274acf539141081d88489cc"></a>

## protocol property — Property reference / 8c453f7252bd / 15

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.

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

- [service_info](resources--endpoint--reference--group-001.md#canonical-cbfadd1795e5aa517039634b1f11091982818ee1faea3b56f1f02b62ca4cded4): complete subsection reference.

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638): complete subsection reference.

- [timeouts](resources--endpoint--reference--group-001.md#canonical-88e7657f6415c5899fa119005b29f2ebd5e02465c5d71ce19baa7148a452b9ac): complete subsection reference.

- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06): complete subsection reference.

<a id="canonical-af726a309e79bdda17399b345281f9061b4081aee3d165d90598976515bc2cae"></a>

## All schema paths — Property reference / 8c453f7252bd / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--endpoint--reference--group-001.md#canonical-ca9d39a5e54245e7417b748ae708a7238daca1f6d05021830be7abe392f42c1b) |
| `description` | [description](resources--endpoint--reference--group-001.md#canonical-0852c210e6218383495c91b48d954ca421cfd007ab47828c4a354a067aac5f8e) |
| `disable` | [disable](resources--endpoint--reference--group-001.md#canonical-b17ccd1212fedd917e7f09772be2e2fdba3afd62ff698db4547841d9e6474e21) |
| `dns_name` | [dns_name](resources--endpoint--reference--group-001.md#canonical-f405aa3e2203d5212ee63b5718c0cbf61738c92cb61bc0a7cdf675b5bfdd10c1) |
| `dns_name_advanced` | [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-6afe56c59488739a6bbf8b387e2d5a0635e0692edc0c5faa55f2abc62b4104ef) |
| `dns_name_advanced.name` | [dns_name_advanced.name](resources--endpoint--reference--group-001.md#canonical-604596753872063d83dfd4129615aeec186b66410ae80fc6bf62c14c0f837ff2) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](resources--endpoint--reference--group-001.md#canonical-912902d3fd85272ae6715ff00f735f07e74ef4012f69f81a6931e2b9ba66ad98) |
| `health_check_port` | [health_check_port](resources--endpoint--reference--group-001.md#canonical-d28d126a1b6c3cd1348123b97b0b72acca32625d411ca8ed1fab7df7b744d87a) |
| `id` | [id](resources--endpoint--reference--group-001.md#canonical-b3468be7b5f1c5bbb0c0cbffb7e89a0f8302ecb6819323addf76fccb2897c070) |
| `ip` | [ip](resources--endpoint--reference--group-001.md#canonical-6a10702f689c849631eef95a9786c2348328f2df35530823eeca86fe2861673f) |
| `labels` | [labels](resources--endpoint--reference--group-001.md#canonical-209959b20d5db8e10bc787813815d0097f02a16b7e97f560b098a76148602a99) |
| `name` | [name](resources--endpoint--reference--group-001.md#canonical-e9f063a59f7f8f18f4d20d2a4a9afa9cfd9469506dfbc18369acda5712ff52e6) |
| `namespace` | [namespace](resources--endpoint--reference--group-001.md#canonical-85e6d18827a7c56c706a3cbe1ca46a645683439720424d1dd7d86d325410c941) |
| `port` | [port](resources--endpoint--reference--group-001.md#canonical-cde649e581eed2284e147bc6c429adcc81458924da908a54072d50247b7aa7e3) |
| `protocol` | [protocol](resources--endpoint--reference--group-001.md#canonical-6253a9e23707bd09afa66a76ea6ed5d0525d5a2945bc0664f5962de66760c9ab) |
| `service_info` | [service_info](resources--endpoint--reference--group-001.md#canonical-48260fbfdf06ddd1b083110173b85b10cadeae09fc1b97fe74e263a4b1b877a9) |
| `service_info.discovery_type` | [service_info.discovery_type](resources--endpoint--reference--group-001.md#canonical-070393b826470fefd9a4659a749077c680821da000f7dcd4b6aa4a11bdfc91b7) |
| `service_info.service_name` | [service_info.service_name](resources--endpoint--reference--group-001.md#canonical-ed6ca6089532539b0b17f5ec156b7a0a8d55fbe677c053f04c52e9a68b12137d) |
| `service_info.service_selector` | [service_info.service_selector](resources--endpoint--reference--group-001.md#canonical-3e00acd0a3529c336c0012342514ff5227045fbb687b13ae0881f3f2be604100) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](resources--endpoint--reference--group-001.md#canonical-14d39e6e3746a09b85c950f60d0018a135e83b2a78241df7d3e76f91b64ccc1d) |
| `snat_pool` | [snat_pool](resources--endpoint--reference--group-001.md#canonical-b53a13f7c2b30d25b62867f8618f315aff1f75367499c6e674d0672123f1bd16) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](resources--endpoint--reference--group-001.md#canonical-305ca4b6cc6fe742f341b0c2991fce2d8cb645f18a21980abdd3b5402631c901) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](resources--endpoint--reference--group-001.md#canonical-016d5fbafe714d8bbe0c373edc54c56fb80f5b0dedf8a17c7e4db370743a0428) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](resources--endpoint--reference--group-001.md#canonical-2c76bd44ee99abb2f6d51ba0dbfa93c5444d105630f51e6b9f5e305c91a718f7) |
| `timeouts` | [timeouts](resources--endpoint--reference--group-001.md#canonical-3ab31fa5332709a80f926082c2286fdb91ce11922420057478b36731103fcfa6) |
| `timeouts.create` | [timeouts.create](resources--endpoint--reference--group-001.md#canonical-03e4cca9616ece776acc66a9fdabed32c746a0fa3f9077ebdf77bcc96b6315e9) |
| `timeouts.delete` | [timeouts.delete](resources--endpoint--reference--group-001.md#canonical-c26c57a4fd9b3e8f81feb5c0303cceb8365cb83b3b5ab9946442723b4a0654ba) |
| `timeouts.read` | [timeouts.read](resources--endpoint--reference--group-001.md#canonical-f3ce9547c42ef79992b80648959911c62d0c72973ccf921323bfd50bfc0e37dd) |
| `timeouts.update` | [timeouts.update](resources--endpoint--reference--group-001.md#canonical-e6afdc58461b6afbdaca4c8387bcfd9e43608912b7a1972feaef4e824cd8f11c) |
| `where` | [where](resources--endpoint--reference--group-001.md#canonical-1a821da9039aadd0a8013f9dd8cf885d8e17a820026486ef5ea820c0e7555ed0) |
| `where.site` | [where.site](resources--endpoint--reference--group-001.md#canonical-f24f3e056090b1cdb4cf996fc27315ea3c4ee619440d7e020f2cd663a56377cf) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-afcf8034577dc4c2f44ae9d11af613b5c9f37a9a5d632ddfdd9ff7fb3ec34726) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-09321b578b12bd96148e19a6fb8bd08ab2738ccbf82726c34971ffb2f30be389) |
| `where.site.network_type` | [where.site.network_type](resources--endpoint--reference--group-001.md#canonical-d3170dbae14dfb761baf62f3619e9636dc56210d91bf97080b67419b48028b20) |
| `where.site.ref` | [where.site.ref](resources--endpoint--reference--group-001.md#canonical-0720c567ff440093162cff366f699ea7309106c4e2c996dea1d2cb6354c2866e) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--endpoint--reference--group-001.md#canonical-6558d29c5cea5fef31218ab361840ab2a093e2cdd765e168daf72755249814b1) |
| `where.site.ref.name` | [where.site.ref.name](resources--endpoint--reference--group-001.md#canonical-6ad7023f3b585ae7b81e6bc82192354d69d507aa3681d3d1430cdf72db9a8e22) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--endpoint--reference--group-001.md#canonical-426fef7c1743c5554c9f1c8a272f155f844704615e622fa4209b3b3a2a3cc11e) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--endpoint--reference--group-001.md#canonical-21371410601bfff4a33782588091f45180abcd822f85c3b61d433b61b8f19287) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--endpoint--reference--group-001.md#canonical-457acb45bca2f3a7b72a6818a7f05b5a11a03820bd6fbfe15536d2ded306a8f0) |
| `where.virtual_network` | [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-9189105f43056be39da52d34d21836a50e462a94476b3d382bfe34539fc1b228) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--endpoint--reference--group-001.md#canonical-6cfff5375e9b43e73f1e8b7d5c3b10c599ae87f644219f3fc5b19d69c95e5886) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--endpoint--reference--group-001.md#canonical-cb01b91f1e461388001fd14d854513e8cc388bc1849d43684be64672c5fd2640) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--endpoint--reference--group-001.md#canonical-1bb2493634c4a1e4575c3aec56e82162972dbc82e1506da170bb3deef0f5aa97) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--endpoint--reference--group-001.md#canonical-5c6e0d48aaeef69b09bda5c8c9f784c795d1e8c229fce4e7b975e1f126314424) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--endpoint--reference--group-001.md#canonical-b6793e4ba2f1f28ec58f57793016f3b75d4559d997bc5306f96d57879793a8d0) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--endpoint--reference--group-001.md#canonical-f94ea40fb03536ae65882cf1060da587db93a4fb2415241efcc90e3e0af6ca26) |
| `where.virtual_site` | [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-da638557ee38c7b41d6dd48b4218af6e035a85c904819fc009a8f1268f37a215) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-aee8903abb80fcf5c77bd70697a4324c42a62f30a90a565278d7016b7ced7010) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-fb911b5fd2a170562c31fc94b3df459f6d6c462b4830a6448706c05f43ce8aa0) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--endpoint--reference--group-001.md#canonical-f9259148ea66302cf054a575de6455c1dadad574fb12ea2c964fa82b5f166e77) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--endpoint--reference--group-001.md#canonical-c4b95fc9f2f42e773f7ac708e0c8e249845691bafd0c947e18a2ca17703a6f7a) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--endpoint--reference--group-001.md#canonical-47a2753dde2960304e0acbd0e392f667f7cc64a4bf5e6bf4935b34a5b2acdc86) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--endpoint--reference--group-001.md#canonical-ab78562de57644d646412b45f44fa36ef2e57993629eeb93006f8c137ec0e75e) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--endpoint--reference--group-001.md#canonical-8c3e968ce86b5feb32beadba068655198532cb5166fbea86f05dd69bc1416bfb) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--endpoint--reference--group-001.md#canonical-1328cadce1a63330dbd72769a24061496fd490617992096f00adbfd843bb331e) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--endpoint--reference--group-001.md#canonical-c14556a2c4633841d50e422997dd057b699ee77671f0a7f99e0f50a9700086a9) |

<a id="canonical-62a9a3b2f74402e966c3b753eaad697df4d6f5edabaa08330f6d9e886498ccc1"></a>

## Next pages — Property reference / 8c453f7252bd / 17

- [dns_name_advanced](resources--endpoint--reference--group-001.md#canonical-cd11f7bdea51dbff2578a79d9945c86607d7fcc356f8e30f5cecab3b8fdcfb59)
- [service_info](resources--endpoint--reference--group-001.md#canonical-cbfadd1795e5aa517039634b1f11091982818ee1faea3b56f1f02b62ca4cded4)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638)
- [timeouts](resources--endpoint--reference--group-001.md#canonical-88e7657f6415c5899fa119005b29f2ebd5e02465c5d71ce19baa7148a452b9ac)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-cd11f7bdea51dbff2578a79d9945c86607d7fcc356f8e30f5cecab3b8fdcfb59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c052ba0538efe239bcdd933ce79148f6225cf47883fa7874a012306db0c76c29"></a>

## dns_name_advanced — dns_name_advanced / 8b019946f87b / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- dns_name_advanced

<a id="canonical-6afe56c59488739a6bbf8b387e2d5a0635e0692edc0c5faa55f2abc62b4104ef"></a>

Type: `"object"`. single nested block, Optional.

Specifies name and TTL used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ttl_choice": "[\"refresh_interval\"]"
}
```

Terraform syntax:

```terraform
dns_name_advanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1a93f73782f19da58a2ece5f5bdf1192cbeba307a556cd52d1685dbb22807046"></a>

## Direct properties — dns_name_advanced / 8b019946f87b / 3

<a id="canonical-604596753872063d83dfd4129615aeec186b66410ae80fc6bf62c14c0f837ff2"></a>

<a id="canonical-ea64b6aed0338f2d3afb0b2ec3d9b4192c4dff2c8ab430acc765e286e400a7a2"></a>

## name property — dns_name_advanced / 8b019946f87b / 4

Type: `"string"`. Optional.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-912902d3fd85272ae6715ff00f735f07e74ef4012f69f81a6931e2b9ba66ad98"></a>

<a id="canonical-b23db616c37378e8074560db18a69b0e292d7b46495954c58692e39a8a2a2182"></a>

## refresh_interval property — dns_name_advanced / 8b019946f87b / 5

Type: `"number"`. Optional.

Exclusive with \[\] Interval for DNS refresh in seconds.

Upstream description:

Exclusive with \[\] Interval for DNS refresh in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 604800),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  }
}
```

<a id="canonical-1a22a4c78c7914cd0b8b05b31703f435bb7c8d4836131242561a8979f14bc2b3"></a>

## Next pages — dns_name_advanced / 8b019946f87b / 6

- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-cbfadd1795e5aa517039634b1f11091982818ee1faea3b56f1f02b62ca4cded4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccb2c1d9238c424bfca84791f3d4a3da0508e34fd3e21ffe9e9d4aa1b6dd2415"></a>

## service_info — service_info / dcb64d60f490 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- service_info

<a id="canonical-48260fbfdf06ddd1b083110173b85b10cadeae09fc1b97fe74e263a4b1b877a9"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether endpoint service is discovered by name or labels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("service_name",
    "service_selector")}
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
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

Terraform syntax:

```terraform
service_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6d75d708e29d4dedfd0b07b655a5c41d8db38bc150f210b3c7ba403fd752cd8"></a>

## Direct properties — service_info / dcb64d60f490 / 3

<a id="canonical-070393b826470fefd9a4659a749077c680821da000f7dcd4b6aa4a11bdfc91b7"></a>

<a id="canonical-5d11fc2672b1b033e21041f7de2e76c51977371c29840d624ab775475b87e74e"></a>

## discovery_type property — service_info / dcb64d60f490 / 4

Type: `"string"`. Optional.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Upstream description:

Specifies the type of discovery

Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover
from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID_DISCOVERY",
  "enum": [
    "INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ed6ca6089532539b0b17f5ec156b7a0a8d55fbe677c053f04c52e9a68b12137d"></a>

<a id="canonical-d8d9dfe468aa115dcb09623fccaca2c403ac6612b681f9d555c66f0390def7ec"></a>

## service_name property — service_info / dcb64d60f490 / 5

Type: `"string"`. Optional.

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the..

Upstream description:

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the cluster identifier matches. If cluster identifier is not specified then discovery will be
done in all discovery objects of the site.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [service_selector](resources--endpoint--reference--group-001.md#canonical-891b5c00f845feac509ba93ef71532170c2012ad0b80091f28269d846c043d2c): complete subsection reference.

<a id="canonical-c8b9da3c082a0fc740a161145de64152ccf9e9320e04e2f7352860771104834b"></a>

## Next pages — service_info / dcb64d60f490 / 6

- [service_info.service_selector](resources--endpoint--reference--group-001.md#canonical-891b5c00f845feac509ba93ef71532170c2012ad0b80091f28269d846c043d2c)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-891b5c00f845feac509ba93ef71532170c2012ad0b80091f28269d846c043d2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d3afa85f17f264180e51633c4d7a261e7d1ac455007b6155532b945aa0951bd"></a>

## service_info.service_selector — service_info.service_selector / 598d855a67fe / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [service_info](resources--endpoint--reference--group-001.md#canonical-cbfadd1795e5aa517039634b1f11091982818ee1faea3b56f1f02b62ca4cded4)
- service_info.service_selector

<a id="canonical-3e00acd0a3529c336c0012342514ff5227045fbb687b13ae0881f3f2be604100"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
service_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-461705b2b191349ccbfbad9d5156a35ecf443b35c5d03286ad7f6aa2d2f904fe"></a>

## Direct properties — service_info.service_selector / 598d855a67fe / 3

<a id="canonical-14d39e6e3746a09b85c950f60d0018a135e83b2a78241df7d3e76f91b64ccc1d"></a>

<a id="canonical-eca0a30ae7d54dbfd02a9ba06c22c779309ed29fe06a737aa5df081d6606406c"></a>

## expressions property — service_info.service_selector / 598d855a67fe / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-e6f150e3348103215aeb7ca8f288b6264cff34c90f53a9442655e5e07aed7b67"></a>

## Next pages — service_info.service_selector / 598d855a67fe / 5

- [service_info](resources--endpoint--reference--group-001.md#canonical-cbfadd1795e5aa517039634b1f11091982818ee1faea3b56f1f02b62ca4cded4)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-365f290ca0e404698cf18c21cbd5ec7ce6391f4c47841aa61a4e606d86bc4471"></a>

## snat_pool — snat_pool / 6c20579a90be / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- snat_pool

<a id="canonical-b53a13f7c2b30d25b62867f8618f315aff1f75367499c6e674d0672123f1bd16"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-996bdfad3a9bc0a391b122496207f2ecbfb2ef60fa94ff471d61f40f85a46c28"></a>

## Direct properties — snat_pool / 6c20579a90be / 3

- [no_snat_pool](resources--endpoint--reference--group-001.md#canonical-a8a69c7004a3ad8b92c12440d126984f582445f251da96d98a4597f12e4c1f10): complete subsection reference.

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-d59e140225c7de5a08892142f7641b3064a6c90f0208aadf860042dbaa37d441): complete subsection reference.

<a id="canonical-6c1a93fb42f314abcfe733ab4aaaa44da884b236acceff2121e3425143762b45"></a>

## Next pages — snat_pool / 6c20579a90be / 4

- [snat_pool.no_snat_pool](resources--endpoint--reference--group-001.md#canonical-a8a69c7004a3ad8b92c12440d126984f582445f251da96d98a4597f12e4c1f10)
- [snat_pool.snat_pool](resources--endpoint--reference--group-001.md#canonical-d59e140225c7de5a08892142f7641b3064a6c90f0208aadf860042dbaa37d441)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-a8a69c7004a3ad8b92c12440d126984f582445f251da96d98a4597f12e4c1f10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-759038841f5293ea8beff1d74c5a0f2deb62235bb3cdd2122333f2734c545b1b"></a>

## snat_pool.no_snat_pool — snat_pool.no_snat_pool / be1ed4bb0b09 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638)
- snat_pool.no_snat_pool

<a id="canonical-305ca4b6cc6fe742f341b0c2991fce2d8cb645f18a21980abdd3b5402631c901"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-bc639bc15b9c3f293936b893377b969e4e747828f88a5d07301f6b36c774c805"></a>

## Direct properties — snat_pool.no_snat_pool / be1ed4bb0b09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31e404487801cab3da3aa992b6e52a8d6c002a7f28d5911ce619f6b5851f638b"></a>

## Next pages — snat_pool.no_snat_pool / be1ed4bb0b09 / 4

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-d59e140225c7de5a08892142f7641b3064a6c90f0208aadf860042dbaa37d441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-951d44d3adf770bb9f0db9d60ddad9bb0034e818baa16bf7c9fb1ef6fa01bba7"></a>

## snat_pool.snat_pool — snat_pool.snat_pool / 40c2e61b9aa7 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638)
- snat_pool.snat_pool

<a id="canonical-016d5fbafe714d8bbe0c373edc54c56fb80f5b0dedf8a17c7e4db370743a0428"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-274430fa15ff9b66ba7a97ff676b7fca6e767ba208596c48014ffc3add2bfdc0"></a>

## Direct properties — snat_pool.snat_pool / 40c2e61b9aa7 / 3

<a id="canonical-2c76bd44ee99abb2f6d51ba0dbfa93c5444d105630f51e6b9f5e305c91a718f7"></a>

<a id="canonical-42267e7ba7639d68703335c62fd8dba075abffdc6b5aa1b3d88f2fb27049fd01"></a>

## prefixes property — snat_pool.snat_pool / 40c2e61b9aa7 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-725c52e49a002fcf54964d561f1ac4b8bfabc301fe1cf3968210777890f254e4"></a>

## Next pages — snat_pool.snat_pool / 40c2e61b9aa7 / 5

- [snat_pool](resources--endpoint--reference--group-001.md#canonical-b88256a4e5ca0df3c8cc9e384d03b7bad7ce3205cc1582057c96495f74f15638)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-88e7657f6415c5899fa119005b29f2ebd5e02465c5d71ce19baa7148a452b9ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12a42f49b8393eeb8108ce74d11fe1f9636c73f70037fd77a57b99942d49ba6d"></a>

## timeouts — timeouts / a9c4fd58abb9 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- timeouts

<a id="canonical-3ab31fa5332709a80f926082c2286fdb91ce11922420057478b36731103fcfa6"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e0ea3711fd0d211d88a58b02ef04fe57abe84ed7effb32e0addad1cfe0fae0ef"></a>

## Direct properties — timeouts / a9c4fd58abb9 / 3

<a id="canonical-03e4cca9616ece776acc66a9fdabed32c746a0fa3f9077ebdf77bcc96b6315e9"></a>

<a id="canonical-80d016f98eb4a8ad97977548762d015f5994cfe552cefa9bfae2b096d2f76eea"></a>

## create property — timeouts / a9c4fd58abb9 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c26c57a4fd9b3e8f81feb5c0303cceb8365cb83b3b5ab9946442723b4a0654ba"></a>

<a id="canonical-60a1232b66743c3f4494cf76efc1a2d0b87038f75f78b5579c4da0f959335048"></a>

## delete property — timeouts / a9c4fd58abb9 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-f3ce9547c42ef79992b80648959911c62d0c72973ccf921323bfd50bfc0e37dd"></a>

<a id="canonical-9b9941d3d4449085ff9823e22020cf376cfac5ce8ec7d6ce27e22b8ea8aafc78"></a>

## read property — timeouts / a9c4fd58abb9 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-e6afdc58461b6afbdaca4c8387bcfd9e43608912b7a1972feaef4e824cd8f11c"></a>

<a id="canonical-103adb0b366557c88cc317b630eb4c51f540c5372838b12aecf8be93da0aadb2"></a>

## update property — timeouts / a9c4fd58abb9 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9c03deb2cdff555360e96ff1fcbd1d0296332cb9e6ecd82749c9ac61ca2248fc"></a>

## Next pages — timeouts / a9c4fd58abb9 / 8

- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e8f5cab293a152980be0b7f6959933cf1a57c344a08381f7eca2b68029e5db4"></a>

## where — where / 0cc0ca7ed4e8 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- where

<a id="canonical-1a821da9039aadd0a8013f9dd8cf885d8e17a820026486ef5ea820c0e7555ed0"></a>

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

<a id="canonical-eed95538c480e3e48db59801a705d504bbe47a1bf4b952c734194ca3d1bbe6ed"></a>

## Direct properties — where / 0cc0ca7ed4e8 / 3

- [site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475): complete subsection reference.

- [virtual_network](resources--endpoint--reference--group-001.md#canonical-f58ee68b17aa327522249f33e7ab65a11cc5bca7476867681251d83064e5b991): complete subsection reference.

- [virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a): complete subsection reference.

<a id="canonical-a2f324ed2ff75bb3b9b3f1aeeae5f8f611a0eed8b6320bafdd9881960be0569b"></a>

## Next pages — where / 0cc0ca7ed4e8 / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-f58ee68b17aa327522249f33e7ab65a11cc5bca7476867681251d83064e5b991)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71d96d1142bbd0090d175da2cfaa1337d1ca3aa8fbbb09eb708e56661ebf39fe"></a>

## where.site — where.site / 5e2b52ced9d5 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- where.site

<a id="canonical-f24f3e056090b1cdb4cf996fc27315ea3c4ee619440d7e020f2cd663a56377cf"></a>

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

<a id="canonical-c7383119f3cba5fb685529174790b17bf1a78978f19b3e71939ac8069a16b298"></a>

## Direct properties — where.site / 5e2b52ced9d5 / 3

- [disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-66ed8a8da025dc632d551f5e661d7235904e16c86caf253edc43fb84966294c4): complete subsection reference.

- [enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-c685b102a6f3cff0bf0eb29cd503c4dc71f8b02ff3a329287eae47466896dd1e): complete subsection reference.

<a id="canonical-d3170dbae14dfb761baf62f3619e9636dc56210d91bf97080b67419b48028b20"></a>

<a id="canonical-2f03f47f978081e0b4920e9fca0be19117f2c6c417f966f3cdd212c0765278ef"></a>

## network_type property — where.site / 5e2b52ced9d5 / 4

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

- [ref](resources--endpoint--reference--group-001.md#canonical-5468f388308d67d0447e5f6f3aed7127be6d53b5c1812c847d8323b8937ce0e5): complete subsection reference.

<a id="canonical-032088673f8bcdffe1dac551c91f4f639fecde777ffca177ceb02377924ba7d6"></a>

## Next pages — where.site / 5e2b52ced9d5 / 5

- [where.site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-66ed8a8da025dc632d551f5e661d7235904e16c86caf253edc43fb84966294c4)
- [where.site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-c685b102a6f3cff0bf0eb29cd503c4dc71f8b02ff3a329287eae47466896dd1e)
- [where.site.ref](resources--endpoint--reference--group-001.md#canonical-5468f388308d67d0447e5f6f3aed7127be6d53b5c1812c847d8323b8937ce0e5)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-66ed8a8da025dc632d551f5e661d7235904e16c86caf253edc43fb84966294c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60347501a087c060dac87b6e2923f96e71b5d8dc816c0531ea8fb3a7833f1c53"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 92173cd7002b / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- where.site.disable_internet_vip

<a id="canonical-afcf8034577dc4c2f44ae9d11af613b5c9f37a9a5d632ddfdd9ff7fb3ec34726"></a>

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

<a id="canonical-fb271cd7486040c53c1b3cac915d4915dd33a388c689eb2bf2966f4a28ff7c27"></a>

## Direct properties — where.site.disable_internet_vip / 92173cd7002b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91206c1c5bccd28c6b4fafbff775fa30cdb85c9741dab906ada31e5d9fab974b"></a>

## Next pages — where.site.disable_internet_vip / 92173cd7002b / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-c685b102a6f3cff0bf0eb29cd503c4dc71f8b02ff3a329287eae47466896dd1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d084edaf461a798ebbb099ee21384411d6051ec2c0a3b26b97c3fec8268278f"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 5e844502b650 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- where.site.enable_internet_vip

<a id="canonical-09321b578b12bd96148e19a6fb8bd08ab2738ccbf82726c34971ffb2f30be389"></a>

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

<a id="canonical-73044213ece4e85a9c86d780d5eae0373848062b58e815ca12f73a5f017a0c5b"></a>

## Direct properties — where.site.enable_internet_vip / 5e844502b650 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89d008f5206ce2c50ebdb9c77320a864824e8b786058b5f789f15146db3664b4"></a>

## Next pages — where.site.enable_internet_vip / 5e844502b650 / 4

- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-5468f388308d67d0447e5f6f3aed7127be6d53b5c1812c847d8323b8937ce0e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5ab44c9780f1048b98aee35afac8150a64ceb955f0242509547fd605453abe8"></a>

## where.site.ref — where.site.ref / d84f6984de7e / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- where.site.ref

<a id="canonical-0720c567ff440093162cff366f699ea7309106c4e2c996dea1d2cb6354c2866e"></a>

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

<a id="canonical-4d599bcc503f47fc21ecd69f087e1300d1d6338fcc6dc1262f94ffad99d84091"></a>

## Direct properties — where.site.ref / d84f6984de7e / 3

<a id="canonical-6558d29c5cea5fef31218ab361840ab2a093e2cdd765e168daf72755249814b1"></a>

<a id="canonical-e15a5fe9491cb9f2e8a680ddca93f9129d3ce605343c792e8b83c792db4e7ee1"></a>

## kind property — where.site.ref / d84f6984de7e / 4

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

<a id="canonical-6ad7023f3b585ae7b81e6bc82192354d69d507aa3681d3d1430cdf72db9a8e22"></a>

<a id="canonical-3cb4c7818ff664671244e1d3d38d0c760223680df5e60832fb18229d06c256c2"></a>

## name property — where.site.ref / d84f6984de7e / 5

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

<a id="canonical-426fef7c1743c5554c9f1c8a272f155f844704615e622fa4209b3b3a2a3cc11e"></a>

<a id="canonical-84b379339ca189ee15881f5d041dcc89bef4c5d159b82cf606a97ff74eb31618"></a>

## namespace property — where.site.ref / d84f6984de7e / 6

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

<a id="canonical-21371410601bfff4a33782588091f45180abcd822f85c3b61d433b61b8f19287"></a>

<a id="canonical-ecac903379cf3702590613df83038bef9f90b9b95dad13e0b4175216d5a7012e"></a>

## tenant property — where.site.ref / d84f6984de7e / 7

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

<a id="canonical-457acb45bca2f3a7b72a6818a7f05b5a11a03820bd6fbfe15536d2ded306a8f0"></a>

<a id="canonical-c0035ea2e5d44d0f0f6ff74b6dd0091667876fe8d8e5d819b23f921563e4ec33"></a>

## uid property — where.site.ref / d84f6984de7e / 8

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

<a id="canonical-ff3608add39e59c617f04ec2c2f5eb332434fee781eefe8413e6edc92e856ce9"></a>

## Next pages — where.site.ref / d84f6984de7e / 9

- [where.site](resources--endpoint--reference--group-001.md#canonical-5e9352393dbe8b23fa6b54e8fa95c0a8cebdbff47470e011e3265215128c6475)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-f58ee68b17aa327522249f33e7ab65a11cc5bca7476867681251d83064e5b991"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fae003a4921dd8a3bf136cbc9b08a38872e515417b4dfd98ace6ebac9986d7b"></a>

## where.virtual_network — where.virtual_network / 84f0d1230447 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- where.virtual_network

<a id="canonical-9189105f43056be39da52d34d21836a50e462a94476b3d382bfe34539fc1b228"></a>

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

<a id="canonical-ec304511a3b9975b027d7af94317c4ee9c083ca916c753d306b70aebf5a5b28e"></a>

## Direct properties — where.virtual_network / 84f0d1230447 / 3

- [ref](resources--endpoint--reference--group-001.md#canonical-9c24f4b84259bb0b959c2164a222d6d21b3fb16e6ebadc9891d66b4b413b8c31): complete subsection reference.

<a id="canonical-0bc2c7bd486232d5e4731ccad2aa2a74b7ce51cb51ffb95a2405ddce411e45ae"></a>

## Next pages — where.virtual_network / 84f0d1230447 / 4

- [where.virtual_network.ref](resources--endpoint--reference--group-001.md#canonical-9c24f4b84259bb0b959c2164a222d6d21b3fb16e6ebadc9891d66b4b413b8c31)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-9c24f4b84259bb0b959c2164a222d6d21b3fb16e6ebadc9891d66b4b413b8c31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-111a18c16c678e6f5a98bff98186683640a5ccd06a6872a5395db66ef545f2a4"></a>

## where.virtual_network.ref — where.virtual_network.ref / 3ad3418560d1 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-f58ee68b17aa327522249f33e7ab65a11cc5bca7476867681251d83064e5b991)
- where.virtual_network.ref

<a id="canonical-6cfff5375e9b43e73f1e8b7d5c3b10c599ae87f644219f3fc5b19d69c95e5886"></a>

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

<a id="canonical-a9d44c93c63be03744b61d4f2651493fe8bc101ab847c4df8cdede23b520c11a"></a>

## Direct properties — where.virtual_network.ref / 3ad3418560d1 / 3

<a id="canonical-cb01b91f1e461388001fd14d854513e8cc388bc1849d43684be64672c5fd2640"></a>

<a id="canonical-0ccc57be85fd58d4a1751a73a94d025b2df2ed5f9864aea260c778d8fcb5fb8b"></a>

## kind property — where.virtual_network.ref / 3ad3418560d1 / 4

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

<a id="canonical-1bb2493634c4a1e4575c3aec56e82162972dbc82e1506da170bb3deef0f5aa97"></a>

<a id="canonical-16621b4838a01927f91f06d74fdc9c92f387aa353e34c5199acbaf07abfecf88"></a>

## name property — where.virtual_network.ref / 3ad3418560d1 / 5

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

<a id="canonical-5c6e0d48aaeef69b09bda5c8c9f784c795d1e8c229fce4e7b975e1f126314424"></a>

<a id="canonical-7cd0293e145ce88d5f6edbcf04d65cf8cd00d2b4bc42932cc368d501dc9ee7c8"></a>

## namespace property — where.virtual_network.ref / 3ad3418560d1 / 6

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

<a id="canonical-b6793e4ba2f1f28ec58f57793016f3b75d4559d997bc5306f96d57879793a8d0"></a>

<a id="canonical-adf881aef868a8e1181e5efe2db6d109070ee59fe99460b9b0d8543bcbb478ef"></a>

## tenant property — where.virtual_network.ref / 3ad3418560d1 / 7

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

<a id="canonical-f94ea40fb03536ae65882cf1060da587db93a4fb2415241efcc90e3e0af6ca26"></a>

<a id="canonical-61462aa1d94c51f3551352c2778ca30563ce5fa661672869c26477214a47fc10"></a>

## uid property — where.virtual_network.ref / 3ad3418560d1 / 8

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

<a id="canonical-20368f09855af7f329441f0e70a841e4d2136206ea060b53917f3851e9bcf190"></a>

## Next pages — where.virtual_network.ref / 3ad3418560d1 / 9

- [where.virtual_network](resources--endpoint--reference--group-001.md#canonical-f58ee68b17aa327522249f33e7ab65a11cc5bca7476867681251d83064e5b991)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a777c9220ea93f006217cbca651a06eafc8747bce79b84eb4483c7f52fe8261"></a>

## where.virtual_site — where.virtual_site / dd46f128d129 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- where.virtual_site

<a id="canonical-da638557ee38c7b41d6dd48b4218af6e035a85c904819fc009a8f1268f37a215"></a>

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

<a id="canonical-9e4796388c36b00b743ed95300751d5b94f187f1c241ec902eb22ab0303d892a"></a>

## Direct properties — where.virtual_site / dd46f128d129 / 3

- [disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-39dc8265ae0e353ba07dbe10d3b61077494e24be2baea592e9fc8cc39597b281): complete subsection reference.

- [enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-cb575f918b8e87d01093ef15bfb4d141a27b25da9d8c08e4129852e0812e2a7e): complete subsection reference.

<a id="canonical-f9259148ea66302cf054a575de6455c1dadad574fb12ea2c964fa82b5f166e77"></a>

<a id="canonical-92e7353e2b647597473dac0022e78558c619170b89f5673f047928981a783152"></a>

## network_type property — where.virtual_site / dd46f128d129 / 4

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

- [ref](resources--endpoint--reference--group-001.md#canonical-3ebb3f5ece20d3d533266ddde5912674c84b6d5b116703200a46a1be584a9f66): complete subsection reference.

<a id="canonical-213b0859973a92e4e85a963c2f0e9661555fd4f95e1d2b395c84280cc58e04a5"></a>

## Next pages — where.virtual_site / dd46f128d129 / 5

- [where.virtual_site.disable_internet_vip](resources--endpoint--reference--group-001.md#canonical-39dc8265ae0e353ba07dbe10d3b61077494e24be2baea592e9fc8cc39597b281)
- [where.virtual_site.enable_internet_vip](resources--endpoint--reference--group-001.md#canonical-cb575f918b8e87d01093ef15bfb4d141a27b25da9d8c08e4129852e0812e2a7e)
- [where.virtual_site.ref](resources--endpoint--reference--group-001.md#canonical-3ebb3f5ece20d3d533266ddde5912674c84b6d5b116703200a46a1be584a9f66)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-39dc8265ae0e353ba07dbe10d3b61077494e24be2baea592e9fc8cc39597b281"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6abe8a9ac95fff09c7e8ce94d95e17bda86d44fd0657e39fc5d0f6f0cabb42c6"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 95edce8a6f43 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- where.virtual_site.disable_internet_vip

<a id="canonical-aee8903abb80fcf5c77bd70697a4324c42a62f30a90a565278d7016b7ced7010"></a>

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

<a id="canonical-2cac08cc60b288c3c6db82e70a589194b6909676ae4397d7241aa197b3007b94"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 95edce8a6f43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f4031477ff96818aacbc487211d149058fb878c016814520a911748ccbf8892"></a>

## Next pages — where.virtual_site.disable_internet_vip / 95edce8a6f43 / 4

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-cb575f918b8e87d01093ef15bfb4d141a27b25da9d8c08e4129852e0812e2a7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26f540f69c8579f04356f13aa5334c59e6f554807c2078417eb318db58cfe718"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / f13eb00b247f / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- where.virtual_site.enable_internet_vip

<a id="canonical-fb911b5fd2a170562c31fc94b3df459f6d6c462b4830a6448706c05f43ce8aa0"></a>

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

<a id="canonical-4c17b609ea76e5d7fa4f73a38aafdfee1960a39b6a633b7a16caa0ab08ab4854"></a>

## Direct properties — where.virtual_site.enable_internet_vip / f13eb00b247f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf355ea0a524399da4a1cefe6733d4249e70166c6f6090d770058b8c410af512"></a>

## Next pages — where.virtual_site.enable_internet_vip / f13eb00b247f / 4

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-3ebb3f5ece20d3d533266ddde5912674c84b6d5b116703200a46a1be584a9f66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caef756ac1cf6c13db2f015ea418647133b78fbe13818099b2f91dc603f68a56"></a>

## where.virtual_site.ref — where.virtual_site.ref / a721a4547548 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Property reference](resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [where](resources--endpoint--reference--group-001.md#canonical-28c620071468ed4ac7a4682f756f4d149971bf0d8bd4d5d2e6b649913a764b06)
- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- where.virtual_site.ref

<a id="canonical-c4b95fc9f2f42e773f7ac708e0c8e249845691bafd0c947e18a2ca17703a6f7a"></a>

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

<a id="canonical-dcbb88e4544336fac12d86983c676f47206adf2fa3bdf43f426141a9427f187f"></a>

## Direct properties — where.virtual_site.ref / a721a4547548 / 3

<a id="canonical-47a2753dde2960304e0acbd0e392f667f7cc64a4bf5e6bf4935b34a5b2acdc86"></a>

<a id="canonical-d26e57dd92f946f6edbe66b34805c8fdd88828180962c4aee7b113eb7f3e4eca"></a>

## kind property — where.virtual_site.ref / a721a4547548 / 4

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

<a id="canonical-ab78562de57644d646412b45f44fa36ef2e57993629eeb93006f8c137ec0e75e"></a>

<a id="canonical-569293c087e55b81c0b6c892ce77820dbc242a1159e50d27b4d26b2daba58a55"></a>

## name property — where.virtual_site.ref / a721a4547548 / 5

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

<a id="canonical-8c3e968ce86b5feb32beadba068655198532cb5166fbea86f05dd69bc1416bfb"></a>

<a id="canonical-811ea8dd1016e44c370ca5a0cdd7ee56cdfe9b04b25df7f5f0b267a2053b6dea"></a>

## namespace property — where.virtual_site.ref / a721a4547548 / 6

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

<a id="canonical-1328cadce1a63330dbd72769a24061496fd490617992096f00adbfd843bb331e"></a>

<a id="canonical-448ccc6d184a80ac99a1d158d39a12e6532df53da1b09e07f0d583a620dd4025"></a>

## tenant property — where.virtual_site.ref / a721a4547548 / 7

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

<a id="canonical-c14556a2c4633841d50e422997dd057b699ee77671f0a7f99e0f50a9700086a9"></a>

<a id="canonical-fa4dad78bd193b8d532b6265ef06fad384cd58704d658ce6f3edf069186dbd70"></a>

## uid property — where.virtual_site.ref / a721a4547548 / 8

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

<a id="canonical-17a51e91bf0cd11be347602bbc2abfea426c106d894174713d1453be2320d28d"></a>

## Next pages — where.virtual_site.ref / a721a4547548 / 9

- [where.virtual_site](resources--endpoint--reference--group-001.md#canonical-2f39bd8208eb559b77aedf1334e72b7b6856ea7481b5cde4d46d10af9501088a)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
