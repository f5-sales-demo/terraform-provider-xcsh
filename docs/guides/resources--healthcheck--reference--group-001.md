---
page_title: "xcsh_healthcheck reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck reference."
---

# xcsh_healthcheck reference

<a id="canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dc3fd7b6b78e95ec6979f64b9fc43f97d681da09b605ba44e7657d970018305"></a>

## Property reference — Property reference / d0c93dae80a6 / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- Property reference

<a id="canonical-c8628364ed2f13518367766737bac310825bb725e5d15baec92984ff6c43c9d8"></a>

## Direct properties — Property reference / d0c93dae80a6 / 3

<a id="canonical-a347d6e198673afb13a785a27bfa10b38405eca9dd0a2460f86151778479661b"></a>

<a id="canonical-1683ce7d041bed7ecde015de805e83144104ea5088eb1d385e45409ad9cfb6cc"></a>

## annotations property — Property reference / d0c93dae80a6 / 4

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

- [default_jitter](resources--healthcheck--reference--group-001.md#canonical-0e645d4c54dad36f2289b09099f3bcc3aa811958bf9f3f4dbe5e5f7e13f7afdc): complete subsection reference.

<a id="canonical-c779b2e9dfb1e6d157a1f216f31ae985d3ab63f8054f55102c29850dedd7e87c"></a>

<a id="canonical-7dc8c5872589b4cf9fd77463020ea8fbfdb455e8d9db52075be747e384228ee2"></a>

## description property — Property reference / d0c93dae80a6 / 5

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

<a id="canonical-e70611c13a116e6fd2e543a474d11146d940b23d6aecd13e49428ac09cc98ef3"></a>

<a id="canonical-2112dd825b786cede79da4b1961eb789b1b0cadeb0da92acc1a683fbc735c03f"></a>

## disable property — Property reference / d0c93dae80a6 / 6

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

<a id="canonical-4d1b8da115b125a1d279bd3138549ea4b1155187ae45f1fd10736da8ba59bad7"></a>

<a id="canonical-e5f90fdd356f43de346c3a317b74ca31aef4b5319805213e8de823854230df2d"></a>

## healthy_threshold property — Property reference / d0c93dae80a6 / 7

Type: `"number"`. Required.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy. Recommended: \`3\`.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-264d50a71f517f746d4fd5e2dd5a4ce28acd27c9aa6256a8aace9be4138eea2c): complete subsection reference.

<a id="canonical-aa71ca2198790ae8fd9051069aae9229c598c43556077f22a3997865f147b682"></a>

<a id="canonical-eab1e77d055cf8458f82fd5bd4f5762ffe143eb4e13df314dc0ad4b7ca8b8e8b"></a>

## id property — Property reference / d0c93dae80a6 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-03b2895664b5ee41027dce68a486fe88be3a3dd2646441080a7495478709085c"></a>

<a id="canonical-9ccd2e73676442bc2e7e2f737b7b3013b17e4d2e36a83621a6e2bfa824dc52eb"></a>

## interval property — Property reference / d0c93dae80a6 / 9

Type: `"number"`. Required.

Time interval in seconds between two healthcheck requests. Recommended: \`15\`.

Upstream description:

Time interval in seconds between two healthcheck requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-c006a240a1dbf8650b4d8e5aa1083d386323745addc1062ee8963d65255ae684"></a>

<a id="canonical-215e8019a57ec3390b88b6dc7096e50679b4008b0e35163781efc42455ee0ce4"></a>

## jitter_percent property — Property reference / d0c93dae80a6 / 10

Type: `"number"`. Optional, Computed.

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50. Server applies default when omitted.
Recommended: \`30\`.

Upstream description:

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 50},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Non-contiguous: {0} union [10, 50] — values 1-9 rejected by API",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "multipleOf": 1,
    "ranges": [
      {
        "maximum": 0,
        "minimum": 0
      },
      {
        "maximum": 50,
        "minimum": 10
      }
    ]
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  }
}
```

<a id="canonical-76a91b33c9d947c1d410a9c7fdff416f6ef8b42ea6409ddf1ad7d03335b3730f"></a>

<a id="canonical-2c876cc2bf69b13c6dfbc449ed31ff53e73bd0681bd993866b7fd4109778418c"></a>

## labels property — Property reference / d0c93dae80a6 / 11

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

<a id="canonical-d2a37394439e2fa96d11a641665949cbc1c3263fdfc0188abcad532712939402"></a>

<a id="canonical-1a50f04fc2b0d6e8aa5a25a121debf9dbd1fda563cd3c43943d6c264aa6e1b7f"></a>

## name property — Property reference / d0c93dae80a6 / 12

Type: `"string"`. Required.

Name of the Healthcheck. Must be unique within the namespace.

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

<a id="canonical-e4136ffa2ffc0ea529cc4f65f77baf280b957cd53e346f16b93dcb7ffbed0a5b"></a>

<a id="canonical-7e88eaf1c33c0f6924937c60b16868b9a466e4265f0694296180b9017c6663f8"></a>

## namespace property — Property reference / d0c93dae80a6 / 13

Type: `"string"`. Required.

Namespace where the Healthcheck is created.

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

- [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-d38675190096fcca96cbd7d5419ca6f9c978df26e44617ecbbdfd84a175b6082): complete subsection reference.

<a id="canonical-0e478d7f8af63599882b4d9425203b3db57388bb5357c8955ebb59f1670fe752"></a>

<a id="canonical-020b58007b627ed7d4bf758e9bad03069fcddecfb843a74d7c38e95bbe11d121"></a>

## timeout property — Property reference / d0c93dae80a6 / 14

Type: `"number"`. Required.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure. Recommended: \`3\`.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "API rejects timeout > 600 for healthchecks (global pattern says 3600)",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [timeouts](resources--healthcheck--reference--group-001.md#canonical-474763bc18bcbc2874fce5ed1e0c75fe4e93d1d590d2c75e9c3cb4f644237a1a): complete subsection reference.

- [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-977bbbf706f9971d29c00a24df809f034a7399af65474627fb883cc9320665b3): complete subsection reference.

<a id="canonical-6fefce642e19583c2173f88966499b6d075f48922cb4f36ca566808a9e637f2b"></a>

<a id="canonical-658b125d9aafb005a94a4f66c4677f0d6a613c670e73a59d5e986e79e0f678fb"></a>

## unhealthy_threshold property — Property reference / d0c93dae80a6 / 15

Type: `"number"`. Required.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately. Recommended: \`1\`.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-66eacb98244bb836e5fa863376b1fac626080d367787cfefa4907cb37eed52e4"></a>

## All schema paths — Property reference / d0c93dae80a6 / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--healthcheck--reference--group-001.md#canonical-a347d6e198673afb13a785a27bfa10b38405eca9dd0a2460f86151778479661b) |
| `default_jitter` | [default_jitter](resources--healthcheck--reference--group-001.md#canonical-7df7aefe90dfd288a6588b93bbd7c817e1a77825dad2aac5fc20f27e5a9a03d4) |
| `description` | [description](resources--healthcheck--reference--group-001.md#canonical-c779b2e9dfb1e6d157a1f216f31ae985d3ab63f8054f55102c29850dedd7e87c) |
| `disable` | [disable](resources--healthcheck--reference--group-001.md#canonical-e70611c13a116e6fd2e543a474d11146d940b23d6aecd13e49428ac09cc98ef3) |
| `healthy_threshold` | [healthy_threshold](resources--healthcheck--reference--group-001.md#canonical-4d1b8da115b125a1d279bd3138549ea4b1155187ae45f1fd10736da8ba59bad7) |
| `http_health_check` | [http_health_check](resources--healthcheck--reference--group-001.md#canonical-d769acf4d13d4863271522062894785a4edbd42025acdc1f373f8593a7598e51) |
| `http_health_check.expected_response` | [http_health_check.expected_response](resources--healthcheck--reference--group-001.md#canonical-3fc64778b3950ebf2cca624b08e4816fdddb8328a00603be03619ebe7b5ad571) |
| `http_health_check.expected_status_codes` | [http_health_check.expected_status_codes](resources--healthcheck--reference--group-001.md#canonical-266008cb35c34c86b8b50ae73e442511bfa3ec52b7bcc97dff4191be89f976cb) |
| `http_health_check.headers` | [http_health_check.headers](resources--healthcheck--reference--group-001.md#canonical-a615ee515775ec18732d5e55883f931bde53411ef94a1f0938172cf951ced580) |
| `http_health_check.host_header` | [http_health_check.host_header](resources--healthcheck--reference--group-001.md#canonical-0acf5949fe5233180f1f673bd908d288ac7cd999eb9955cae2622441d16add23) |
| `http_health_check.path` | [http_health_check.path](resources--healthcheck--reference--group-001.md#canonical-7a661de8464e3fe86a736fde3a43db6d1aae0e8e5570c7db37d8c035ecf3260a) |
| `http_health_check.request_headers_to_remove` | [http_health_check.request_headers_to_remove](resources--healthcheck--reference--group-001.md#canonical-9559f64f42de808cd9c8f3d33038b6e55ae74e7cd04df551b95736b7d16ecad6) |
| `http_health_check.use_http2` | [http_health_check.use_http2](resources--healthcheck--reference--group-001.md#canonical-d1ce8a7f2a08d88c9cd428ee03ddc3594c0d6c5aa584f165245221b7ff8d77da) |
| `http_health_check.use_origin_server_name` | [http_health_check.use_origin_server_name](resources--healthcheck--reference--group-001.md#canonical-f3a1e3ddbc83c944046cfba970df6f7b39507fa233b9885a4229846ed831096c) |
| `id` | [id](resources--healthcheck--reference--group-001.md#canonical-aa71ca2198790ae8fd9051069aae9229c598c43556077f22a3997865f147b682) |
| `interval` | [interval](resources--healthcheck--reference--group-001.md#canonical-03b2895664b5ee41027dce68a486fe88be3a3dd2646441080a7495478709085c) |
| `jitter_percent` | [jitter_percent](resources--healthcheck--reference--group-001.md#canonical-c006a240a1dbf8650b4d8e5aa1083d386323745addc1062ee8963d65255ae684) |
| `labels` | [labels](resources--healthcheck--reference--group-001.md#canonical-76a91b33c9d947c1d410a9c7fdff416f6ef8b42ea6409ddf1ad7d03335b3730f) |
| `name` | [name](resources--healthcheck--reference--group-001.md#canonical-d2a37394439e2fa96d11a641665949cbc1c3263fdfc0188abcad532712939402) |
| `namespace` | [namespace](resources--healthcheck--reference--group-001.md#canonical-e4136ffa2ffc0ea529cc4f65f77baf280b957cd53e346f16b93dcb7ffbed0a5b) |
| `tcp_health_check` | [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-52e9c11b3c0c8053e4a0c96c38560852aa4520120b982bf2b236b14bb6ec82e2) |
| `tcp_health_check.expected_response` | [tcp_health_check.expected_response](resources--healthcheck--reference--group-001.md#canonical-550fe923bd7e6f587a8b27364c798b4c56bcb2bac6dfc7d5cc6bd9c45b8af820) |
| `tcp_health_check.send_payload` | [tcp_health_check.send_payload](resources--healthcheck--reference--group-001.md#canonical-c5e7e5c07185a62e7a7ec5af346fb724d050fb5bb5a38a73932878c9f7c3af78) |
| `timeout` | [timeout](resources--healthcheck--reference--group-001.md#canonical-0e478d7f8af63599882b4d9425203b3db57388bb5357c8955ebb59f1670fe752) |
| `timeouts` | [timeouts](resources--healthcheck--reference--group-001.md#canonical-21b7932580686ebb6e1d1879ce4cc8294471d82bb7c40c3c8824f331ec47a2a6) |
| `timeouts.create` | [timeouts.create](resources--healthcheck--reference--group-001.md#canonical-f3331b0224fb10271e26d71705c15382471be6369bbf323bd41f1b6ee66c71b5) |
| `timeouts.delete` | [timeouts.delete](resources--healthcheck--reference--group-001.md#canonical-181f1d0d8c6cda8c568c34905775058f5ce67b9a9c26da804434527179f13036) |
| `timeouts.read` | [timeouts.read](resources--healthcheck--reference--group-001.md#canonical-a3d4298b204428bf49eb4977065b5f4aec39095c5106c2a6ea84b09ea9da679d) |
| `timeouts.update` | [timeouts.update](resources--healthcheck--reference--group-001.md#canonical-f5f63b23e1c4480cbd53ff315067eeff778fa32a556f94138e75d5108363c702) |
| `udp_icmp_health_check` | [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-1c23b167d70b5c3c0b778d8b61f20fc1a03c99a4d519565cf208b859b56f1ff9) |
| `unhealthy_threshold` | [unhealthy_threshold](resources--healthcheck--reference--group-001.md#canonical-6fefce642e19583c2173f88966499b6d075f48922cb4f36ca566808a9e637f2b) |

<a id="canonical-31b3598b27e2a22d8c967c8a730ac5bf994ea97cfad74e04d32bbde3c58e348f"></a>

## Next pages — Property reference / d0c93dae80a6 / 17

- [default_jitter](resources--healthcheck--reference--group-001.md#canonical-0e645d4c54dad36f2289b09099f3bcc3aa811958bf9f3f4dbe5e5f7e13f7afdc)
- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-264d50a71f517f746d4fd5e2dd5a4ce28acd27c9aa6256a8aace9be4138eea2c)
- [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-d38675190096fcca96cbd7d5419ca6f9c978df26e44617ecbbdfd84a175b6082)
- [timeouts](resources--healthcheck--reference--group-001.md#canonical-474763bc18bcbc2874fce5ed1e0c75fe4e93d1d590d2c75e9c3cb4f644237a1a)
- [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-977bbbf706f9971d29c00a24df809f034a7399af65474627fb883cc9320665b3)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-0e645d4c54dad36f2289b09099f3bcc3aa811958bf9f3f4dbe5e5f7e13f7afdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76e463c44a7bd79ff3d4dac0d443d92bbba2ee04e85a4eb265feb5273f2da882"></a>

## default_jitter — default_jitter / fc12da585918 / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- default_jitter

<a id="canonical-7df7aefe90dfd288a6588b93bbd7c817e1a77825dad2aac5fc20f27e5a9a03d4"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

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

- [default_jitter](resources--healthcheck--reference--group-001.md#canonical-7df7aefe90dfd288a6588b93bbd7c817e1a77825dad2aac5fc20f27e5a9a03d4)
- [jitter_percent](resources--healthcheck--reference--group-001.md#canonical-c006a240a1dbf8650b4d8e5aa1083d386323745addc1062ee8963d65255ae684)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_jitter = {}
```

<a id="canonical-e874779aebdebd1d5aa822e7a41e370d28269bdaad49533b65c260ff1336441e"></a>

## Direct properties — default_jitter / fc12da585918 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-646f5ac67e08f432e25eafd92b42f8079b0d433ea5890dea23041ceb2322b785"></a>

## Next pages — default_jitter / fc12da585918 / 4

- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-264d50a71f517f746d4fd5e2dd5a4ce28acd27c9aa6256a8aace9be4138eea2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be5abcb0df1eff8b267e793aa2be961a35fc3530ded7c06a38f88a0816d92657"></a>

## http_health_check — http_health_check / 6a79a1a1812a / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- http_health_check

<a id="canonical-d769acf4d13d4863271522062894785a4edbd42025acdc1f373f8593a7598e51"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, tcp\_health\_check, udp\_icmp\_health\_check\] Healthy if 'GET' method
on URL 'HTTP(s)://&lt;host&gt;/&lt;path&gt;' with optional '&lt;header&gt;' returns success. 'host'
is not used for DNS resolution. It is used as HTTP Header in the request.

Upstream description:

Healthy if "GET" method on URL "HTTP(s)://&lt;host&gt;/&lt;path&gt;" with optional "&lt;header&gt;"
returns success. "host" is not used for DNS resolution. It is used as HTTP Header in the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path"),
  validators.ConflictingObjectAttributes("host_header",
    "use_origin_server_name")}
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
  "x-ves-oneof-field-host_header_choice": "[\"host_header\",\"use_origin_server_name\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-d769acf4d13d4863271522062894785a4edbd42025acdc1f373f8593a7598e51)
- [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-52e9c11b3c0c8053e4a0c96c38560852aa4520120b982bf2b236b14bb6ec82e2)
- [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-1c23b167d70b5c3c0b778d8b61f20fc1a03c99a4d519565cf208b859b56f1ff9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-a00a9629118eb7471e8c141b0aa242d6bb30cbf68297bf1708cdacff5b1b1687"></a>

## Direct properties — http_health_check / 6a79a1a1812a / 3

<a id="canonical-3fc64778b3950ebf2cca624b08e4816fdddb8328a00603be03619ebe7b5ad571"></a>

<a id="canonical-9c11a51ba7b587fac3ec28d3825ef5e5bf2c660baea052d32da8f515de19e518"></a>

## expected_response property — http_health_check / 6a79a1a1812a / 4

Type: `"string"`. Optional, Computed.

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status. Server
applies default when omitted.

Upstream description:

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-266008cb35c34c86b8b50ae73e442511bfa3ec52b7bcc97dff4191be89f976cb"></a>

<a id="canonical-2664ec4d46f6b184d94f4ae8bb469d960e1a17365ab16290457291d1d5391f1c"></a>

## expected_status_codes property — http_health_check / 6a79a1a1812a / 5

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by '-'. Defaults to
\`\[\]\`. Server applies default when omitted.

Upstream description:

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a615ee515775ec18732d5e55883f931bde53411ef94a1f0938172cf951ced580"></a>

<a id="canonical-a916e9f1bc879c71d8e86ee1795ffe1edca7a8fb9343ed9dcc0a39bd93cf2aaf"></a>

## headers property — http_health_check / 6a79a1a1812a / 6

Type: `["map", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs. Defaults to \`map\[\]\`. Server applies default
when omitted.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0acf5949fe5233180f1f673bd908d288ac7cd999eb9955cae2622441d16add23"></a>

<a id="canonical-62b247ca34a5beba882fbec634149feb49e6c4b4560a1791ab8f8841c207bb46"></a>

## host_header property — http_health_check / 6a79a1a1812a / 7

Type: `"string"`. Optional.

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Upstream description:

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-7a661de8464e3fe86a736fde3a43db6d1aae0e8e5570c7db37d8c035ecf3260a"></a>

<a id="canonical-0ad0ab2d5c81b4228c8e48f7d45f174341601f6e45b7fed85d1b3af2834138ed"></a>

## path property — http_health_check / 6a79a1a1812a / 8

Type: `"string"`. Optional.

Specifies the HTTP path that will be requested during health checking. Recommended: \`/\`.

Upstream description:

Specifies the HTTP path that will be requested during health checking.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-9559f64f42de808cd9c8f3d33038b6e55ae74e7cd04df551b95736b7d16ecad6"></a>

<a id="canonical-a219b47289c6d4acafbd323ec97fbdd2f0bbb2dcfce6b4cad538da94e9f11c8f"></a>

## request_headers_to_remove property — http_health_check / 6a79a1a1812a / 9

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-d1ce8a7f2a08d88c9cd428ee03ddc3594c0d6c5aa584f165245221b7ff8d77da"></a>

<a id="canonical-415ab239d139392e86b7cb1be0523e05a4efda1c30ca693434ea308a349d018f"></a>

## use_http2 property — http_health_check / 6a79a1a1812a / 10

Type: `"bool"`. Optional, Computed.

If set, health checks will be made using HTTP/2. Defaults to \`false\`. Server applies default when
omitted. Recommended: \`false\`.

Upstream description:

If set, health checks will be made using HTTP/2.

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

- [use_origin_server_name](resources--healthcheck--reference--group-001.md#canonical-9e7eb8b2e65ee6a13a811b5ccb21a5e96759fd97b743881a254d6d37306d2dcc): complete subsection reference.

<a id="canonical-5c6e2214b7ff2212f22731275c6c96b94cc3983342878bfc0239c6784465934b"></a>

## Next pages — http_health_check / 6a79a1a1812a / 11

- [http_health_check.use_origin_server_name](resources--healthcheck--reference--group-001.md#canonical-9e7eb8b2e65ee6a13a811b5ccb21a5e96759fd97b743881a254d6d37306d2dcc)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-9e7eb8b2e65ee6a13a811b5ccb21a5e96759fd97b743881a254d6d37306d2dcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11be601b4852fccd7ef9e24d76a2ca69d2a42f70c0b1e5fb92e7b316a2e69e35"></a>

## http_health_check.use_origin_server_name — http_health_check.use_origin_server_name / a1c219c1b3ef / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-264d50a71f517f746d4fd5e2dd5a4ce28acd27c9aa6256a8aace9be4138eea2c)
- http_health_check.use_origin_server_name

<a id="canonical-f3a1e3ddbc83c944046cfba970df6f7b39507fa233b9885a4229846ed831096c"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_origin_server_name = {}
```

<a id="canonical-0030d23f0e7995eb6dba6f5413ac5163d15cc52b494b55e475ba954b71af6ca3"></a>

## Direct properties — http_health_check.use_origin_server_name / a1c219c1b3ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb11b2f2cbe3e52f20dc3bc0c47bd94fe6581f85d3a632445d5226850177e68b"></a>

## Next pages — http_health_check.use_origin_server_name / a1c219c1b3ef / 4

- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-264d50a71f517f746d4fd5e2dd5a4ce28acd27c9aa6256a8aace9be4138eea2c)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-d38675190096fcca96cbd7d5419ca6f9c978df26e44617ecbbdfd84a175b6082"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ece9783d1c1ac31fe8e26886cc53e351a3a203cfb82e6a834eb1be67d54315d"></a>

## tcp_health_check — tcp_health_check / 7824d39099c6 / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- tcp_health_check

<a id="canonical-52e9c11b3c0c8053e4a0c96c38560852aa4520120b982bf2b236b14bb6ec82e2"></a>

Type: `"object"`. single nested block, Optional.

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;.

Upstream description:

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-a1db0295c2b459a471b3cfa9c95877fb4b49c175283a47e85385d4e2f8eda4e7"></a>

## Direct properties — tcp_health_check / 7824d39099c6 / 3

<a id="canonical-550fe923bd7e6f587a8b27364c798b4c56bcb2bac6dfc7d5cc6bd9c45b8af820"></a>

<a id="canonical-c246eb9f724377b34ef480d3287060fc8da788aa2fc541f00558d3efd5e28d8c"></a>

## expected_response property — tcp_health_check / 7824d39099c6 / 4

Type: `"string"`. Optional.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Upstream description:

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-c5e7e5c07185a62e7a7ec5af346fb724d050fb5bb5a38a73932878c9f7c3af78"></a>

<a id="canonical-d14951445c2ada5b85487f90981c0ec12cbc1a12d2cf9d69e189c3b3ff21de4f"></a>

## send_payload property — tcp_health_check / 7824d39099c6 / 5

Type: `"string"`. Optional.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Upstream description:

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-229110a9cc2feea3c8d821f15fc0f32106bb98eb21635332925b84ef64803e09"></a>

## Next pages — tcp_health_check / 7824d39099c6 / 6

- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-474763bc18bcbc2874fce5ed1e0c75fe4e93d1d590d2c75e9c3cb4f644237a1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f29e6928b5c68e4951e970ac17a29a3e9f3016c87db5b3ba985e3fff3d6ba70"></a>

## timeouts — timeouts / e0d1982e2596 / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- timeouts

<a id="canonical-21b7932580686ebb6e1d1879ce4cc8294471d82bb7c40c3c8824f331ec47a2a6"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2636524531073083924337947a338b66cfaa7d4beff0a3cd8793154ad1bcb1db"></a>

## Direct properties — timeouts / e0d1982e2596 / 3

<a id="canonical-f3331b0224fb10271e26d71705c15382471be6369bbf323bd41f1b6ee66c71b5"></a>

<a id="canonical-e3a557d3252818952314f2c1253f6bc8d03e0092413e1e8e63c0c791c9dafc03"></a>

## create property — timeouts / e0d1982e2596 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-181f1d0d8c6cda8c568c34905775058f5ce67b9a9c26da804434527179f13036"></a>

<a id="canonical-35b0db1b6e28054f46c18d74f2a9c40c0bc731c403fbc7264cd7be0123b13a19"></a>

## delete property — timeouts / e0d1982e2596 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-a3d4298b204428bf49eb4977065b5f4aec39095c5106c2a6ea84b09ea9da679d"></a>

<a id="canonical-a57b86a9efe5ee55c7ae1038eeaad88b774f89ea302f2714e7ec24e7fda0ca37"></a>

## read property — timeouts / e0d1982e2596 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-f5f63b23e1c4480cbd53ff315067eeff778fa32a556f94138e75d5108363c702"></a>

<a id="canonical-442e204da09738f324d5120d6aea37d3c67723a3a012daee2273f934d1fddb3f"></a>

## update property — timeouts / e0d1982e2596 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-98a62855c98b80f63ecc08a0889578cae0a0af5e7a5a1cc7f9d4b117835ac152"></a>

## Next pages — timeouts / e0d1982e2596 / 8

- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)

<a id="canonical-977bbbf706f9971d29c00a24df809f034a7399af65474627fb883cc9320665b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-814408ffe1fb68dda26a4b9ed016ea900aeb0f3f82121483bb949a70e57ff287"></a>

## udp_icmp_health_check — udp_icmp_health_check / 824f8c765a4a / 2

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- udp_icmp_health_check

<a id="canonical-1c23b167d70b5c3c0b778d8b61f20fc1a03c99a4d519565cf208b859b56f1ff9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for udp icmp health check.

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
udp_icmp_health_check = {}
```

<a id="canonical-bc9c6290231a12e30eb58b0c5dddebbf945f3c93740ab69a1641df49e52bf85f"></a>

## Direct properties — udp_icmp_health_check / 824f8c765a4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17d602dd74a4d5cb9b2f6f51d9bd01ea9df93c40cc8e131d52c24d53546a63fe"></a>

## Next pages — udp_icmp_health_check / 824f8c765a4a / 4

- [Property reference](resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [xcsh_healthcheck](../resources/healthcheck.md#canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610)
