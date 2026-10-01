---
page_title: "xcsh_healthcheck reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck reference."
---

# xcsh_healthcheck reference

<a id="canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a37da72910a9ee6f522ab33c8ed56522d269b4ceb8a039769f1f7d9bdbd249c2"></a>

## Property reference — Property reference / 414d0a47ba6f / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- Property reference

<a id="canonical-fb467c10a6fbe8ad36e5f95d4e6dcc913af0ad5de97304b2d8b3136b3a087fee"></a>

## Direct properties — Property reference / 414d0a47ba6f / 3

<a id="canonical-6e95afda7b4c1c52ba653aae5dc265d773bf071e1aff2a60d703bbbfe9ac8708"></a>

<a id="canonical-1e09f0c2a2d5cd053b359f925a0a0ba8799d0457073267d82d67bf52c5162bec"></a>

## annotations property — Property reference / 414d0a47ba6f / 4

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

- [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-3a384028ff1e718966c67bfab252894a89df3f3512e390640c4d3477f8556bfb): complete subsection reference.

<a id="canonical-39af24329ff1328bd138ac37d3cda43969ea5f6f66bc1e2fccf7a8735f31c873"></a>

<a id="canonical-bd62acb94c2be4b572867ad759990d47fcb7e9334ed028abcf53b976f93f332f"></a>

## description property — Property reference / 414d0a47ba6f / 5

Type: `"string"`. Computed.

Description of the Healthcheck.

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

<a id="canonical-16e2a6d7a80b794a25fa181513f587a9da58dac71c83083d26a3da4330885e9e"></a>

<a id="canonical-939d6d98c175ebecf65b81f7fa33cbf6520f2cae2b320373b744bd7b38fff3d5"></a>

## healthy_threshold property — Property reference / 414d0a47ba6f / 6

Type: `"number"`. Computed.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy. Recommended: \`3\`.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

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

- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-83cc441036af77ca9aca9b2f4a34d969890f60c41814d6bd2f0731cd93b4f8d9): complete subsection reference.

<a id="canonical-510f3d44274bcb301b8f97493ebbe89317f234709c119d55aaaad4b318ba7e6b"></a>

<a id="canonical-5639b6f2a5fddd2056ac1b56e26bed2331f866dcb3b2fee8f3190b6507a34a3d"></a>

## id property — Property reference / 414d0a47ba6f / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a9e2f9992933d9a5dcd2ac6800bad2c404fc39179beefdfcce02d2edfa543bcb"></a>

<a id="canonical-8a619d3290f1fdbe72b16db7f5b7de5ea5ed155b731499de74527561448fb24b"></a>

## interval property — Property reference / 414d0a47ba6f / 8

Type: `"number"`. Computed.

Time interval in seconds between two healthcheck requests. Recommended: \`15\`.

Upstream description:

Time interval in seconds between two healthcheck requests.

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

<a id="canonical-c3049e9b4a473e6255e69333d71265aa35cdd129b22df81469365bfeadb9c1f1"></a>

<a id="canonical-227dd78e477e454cbb6f8d78f960271a9dc3d59812344ff1676ef35046f94bc6"></a>

## jitter_percent property — Property reference / 414d0a47ba6f / 9

Type: `"number"`. Computed.

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50. Server applies default when omitted.
Recommended: \`30\`.

Upstream description:

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50.

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

<a id="canonical-2e15106d7b7ab778e2ce8322e8eab1402a454bff2a71f152f864949307cbbc2a"></a>

<a id="canonical-7ccc62cb61c862d8fc762f11150f7b1282da95f29e7c0775d764ed171f22e4a5"></a>

## labels property — Property reference / 414d0a47ba6f / 10

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

<a id="canonical-c2a9355f880048283fccbda77d3208f0581272878518fb4b4b4d7ac391957736"></a>

<a id="canonical-eaef6cf4f378f94ecff6f63cc2fc00c1da7333173a3626c1c7e9b970b0b454fe"></a>

## name property — Property reference / 414d0a47ba6f / 11

Type: `"string"`. Required.

Name of the Healthcheck.

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

<a id="canonical-3dca69fdd9fda665934f0e54f01d327d7b3b56dd6e50e4d6cc0000e7bb0ac739"></a>

<a id="canonical-070fcdc9eca507503a4bb5fd7e5091eb2e7562cfe3232145d4d23c3198b96f21"></a>

## namespace property — Property reference / 414d0a47ba6f / 12

Type: `"string"`. Required.

Namespace where the Healthcheck exists.

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

- [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-7cc63fd84f59db2fbdb2ab13e4eb35522dff93d87fb58c595e3f6109eee12b94): complete subsection reference.

<a id="canonical-46ed8f500f25c586357b184b777efc030e403e49c3576f843b7bce878e71132a"></a>

<a id="canonical-afff727c78b98ae0784c01b57271e8ddd7b2c58f238e5b1855a10d7ea4776a51"></a>

## timeout property — Property reference / 414d0a47ba6f / 13

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure. Recommended: \`3\`.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

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

- [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-8012b193094025ee22fe96956b0ce2a2cc4f203185a13906083b01136028ac2e): complete subsection reference.

<a id="canonical-b487a0d45a154da42cc7b6b7b9ca8c70ac64b0050324c50a716fd0a9a68ceda9"></a>

<a id="canonical-ebefbf3938760c22dff317f01ba7162b47c84d96b76651d09b2d8abb12142bc8"></a>

## unhealthy_threshold property — Property reference / 414d0a47ba6f / 14

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately. Recommended: \`1\`.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

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

<a id="canonical-3456f872706320b025ab18157ae3bc2187224ac56ab0e3cf76fa86317739076d"></a>

## All schema paths — Property reference / 414d0a47ba6f / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--healthcheck--reference--group-001.md#canonical-6e95afda7b4c1c52ba653aae5dc265d773bf071e1aff2a60d703bbbfe9ac8708) |
| `default_jitter` | [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-5a61da86009a27fb6e06dbaf0c2f9546ab38d28c9e3c049af06913fb3ed7c3b1) |
| `description` | [description](data-sources--healthcheck--reference--group-001.md#canonical-39af24329ff1328bd138ac37d3cda43969ea5f6f66bc1e2fccf7a8735f31c873) |
| `healthy_threshold` | [healthy_threshold](data-sources--healthcheck--reference--group-001.md#canonical-16e2a6d7a80b794a25fa181513f587a9da58dac71c83083d26a3da4330885e9e) |
| `http_health_check` | [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-bdf272eb8533589ca83ebbecb387b22e19a3346bece1dafa9e880ab1e57966d0) |
| `http_health_check.expected_response` | [http_health_check.expected_response](data-sources--healthcheck--reference--group-001.md#canonical-a71e1ffcfd387372f22eb308edbf582c95d33651cf3cf09f53d3825f25b5a15a) |
| `http_health_check.expected_status_codes` | [http_health_check.expected_status_codes](data-sources--healthcheck--reference--group-001.md#canonical-35c00265419003c9017a12f5d803dcffaa4ed95fd4b4db3c60ef15e241982ecd) |
| `http_health_check.headers` | [http_health_check.headers](data-sources--healthcheck--reference--group-001.md#canonical-de8a9ea40c20379a21494db37183ac95f648ae1d7980dd422756ed6ce7165ed5) |
| `http_health_check.host_header` | [http_health_check.host_header](data-sources--healthcheck--reference--group-001.md#canonical-a4930e080530e863cda95d5669854e7c7edc3dea95bd879b3a22ce59fddf9e66) |
| `http_health_check.path` | [http_health_check.path](data-sources--healthcheck--reference--group-001.md#canonical-e4fdf5dcd8432a8dc992ceb55a5fa58cf6a9f7232d7ae07db5e4fc0070166735) |
| `http_health_check.request_headers_to_remove` | [http_health_check.request_headers_to_remove](data-sources--healthcheck--reference--group-001.md#canonical-3bdb190f6f3000a2e051aa0fc009d5384a3db50fc0220182ae9971f0d12fdacd) |
| `http_health_check.use_http2` | [http_health_check.use_http2](data-sources--healthcheck--reference--group-001.md#canonical-659b1a97aedf7fc1e12343e64041323ae5df01babd0b57a4004db796078567a5) |
| `http_health_check.use_origin_server_name` | [http_health_check.use_origin_server_name](data-sources--healthcheck--reference--group-001.md#canonical-f9c80d1486e8b54bca425fea7882f1c81a81bb33e74be1b83a42d066b21cac40) |
| `id` | [id](data-sources--healthcheck--reference--group-001.md#canonical-510f3d44274bcb301b8f97493ebbe89317f234709c119d55aaaad4b318ba7e6b) |
| `interval` | [interval](data-sources--healthcheck--reference--group-001.md#canonical-a9e2f9992933d9a5dcd2ac6800bad2c404fc39179beefdfcce02d2edfa543bcb) |
| `jitter_percent` | [jitter_percent](data-sources--healthcheck--reference--group-001.md#canonical-c3049e9b4a473e6255e69333d71265aa35cdd129b22df81469365bfeadb9c1f1) |
| `labels` | [labels](data-sources--healthcheck--reference--group-001.md#canonical-2e15106d7b7ab778e2ce8322e8eab1402a454bff2a71f152f864949307cbbc2a) |
| `name` | [name](data-sources--healthcheck--reference--group-001.md#canonical-c2a9355f880048283fccbda77d3208f0581272878518fb4b4b4d7ac391957736) |
| `namespace` | [namespace](data-sources--healthcheck--reference--group-001.md#canonical-3dca69fdd9fda665934f0e54f01d327d7b3b56dd6e50e4d6cc0000e7bb0ac739) |
| `tcp_health_check` | [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-87dd96d8a965d018044db580cbee4ba5057a0092d2e7bd1e5373905d28f667f5) |
| `tcp_health_check.expected_response` | [tcp_health_check.expected_response](data-sources--healthcheck--reference--group-001.md#canonical-81fc06c9d84d4a0447f77d7906a9cdcbdc7a7c79d745d2c28f8fd0e3fece472c) |
| `tcp_health_check.send_payload` | [tcp_health_check.send_payload](data-sources--healthcheck--reference--group-001.md#canonical-26df65095bb6c8decd39f05984cf8e9ee509fcc6e514a29ebad63b1b8652c4c6) |
| `timeout` | [timeout](data-sources--healthcheck--reference--group-001.md#canonical-46ed8f500f25c586357b184b777efc030e403e49c3576f843b7bce878e71132a) |
| `udp_icmp_health_check` | [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-4fcd6931eef4e9bcf5ecf0c83233fce4483edec091c14727e70a992892916163) |
| `unhealthy_threshold` | [unhealthy_threshold](data-sources--healthcheck--reference--group-001.md#canonical-b487a0d45a154da42cc7b6b7b9ca8c70ac64b0050324c50a716fd0a9a68ceda9) |

<a id="canonical-cbac6906a66f0d8e95436b9ebe1b73f9691379beba94f684ab4c9b7d1a4e31a3"></a>

## Next pages — Property reference / 414d0a47ba6f / 16

- [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-3a384028ff1e718966c67bfab252894a89df3f3512e390640c4d3477f8556bfb)
- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-83cc441036af77ca9aca9b2f4a34d969890f60c41814d6bd2f0731cd93b4f8d9)
- [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-7cc63fd84f59db2fbdb2ab13e4eb35522dff93d87fb58c595e3f6109eee12b94)
- [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-8012b193094025ee22fe96956b0ce2a2cc4f203185a13906083b01136028ac2e)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-3a384028ff1e718966c67bfab252894a89df3f3512e390640c4d3477f8556bfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d47762554c4674d071256530cb9a8d06fd0dbb368677983ec706b77ca233b859"></a>

## default_jitter — default_jitter / ca1577444835 / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- default_jitter

<a id="canonical-5a61da86009a27fb6e06dbaf0c2f9546ab38d28c9e3c049af06913fb3ed7c3b1"></a>

Type: `["object", {}]`. Computed.

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

- [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-5a61da86009a27fb6e06dbaf0c2f9546ab38d28c9e3c049af06913fb3ed7c3b1)
- [jitter_percent](data-sources--healthcheck--reference--group-001.md#canonical-c3049e9b4a473e6255e69333d71265aa35cdd129b22df81469365bfeadb9c1f1)

Select alternatives according to the provider validators above.

<a id="canonical-6c5c109c9b9eee8a86570fc0cc7ab217a2a7a996a30062e89be2d62c3177a4ff"></a>

## Direct properties — default_jitter / ca1577444835 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4ad0f13a87e51a5ec81c44e92ef818c9c16bcc129f4a5a111a9319bbcba0b38"></a>

## Next pages — default_jitter / ca1577444835 / 4

- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-83cc441036af77ca9aca9b2f4a34d969890f60c41814d6bd2f0731cd93b4f8d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-091fb8a90c856eaaa3d0600809381b79614f7ecc5185ddddd77d428c07663f12"></a>

## http_health_check — http_health_check / f8856a5ec0d2 / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- http_health_check

<a id="canonical-bdf272eb8533589ca83ebbecb387b22e19a3346bece1dafa9e880ab1e57966d0"></a>

Type: `"single"`. Computed.

\[OneOf: http\_health\_check, tcp\_health\_check, udp\_icmp\_health\_check\] Healthy if 'GET' method
on URL 'HTTP(s)://&lt;host&gt;/&lt;path&gt;' with optional '&lt;header&gt;' returns success. 'host'
is not used for DNS resolution. It is used as HTTP Header in the request.

Upstream description:

Healthy if "GET" method on URL "HTTP(s)://&lt;host&gt;/&lt;path&gt;" with optional "&lt;header&gt;"
returns success. "host" is not used for DNS resolution. It is used as HTTP Header in the request.

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

- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-bdf272eb8533589ca83ebbecb387b22e19a3346bece1dafa9e880ab1e57966d0)
- [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-87dd96d8a965d018044db580cbee4ba5057a0092d2e7bd1e5373905d28f667f5)
- [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-4fcd6931eef4e9bcf5ecf0c83233fce4483edec091c14727e70a992892916163)

Select alternatives according to the provider validators above.

<a id="canonical-63c5f4352e11da054aae5be07924d0796fb094eb8b33d5e1250cc73b3d9c3b8b"></a>

## Direct properties — http_health_check / f8856a5ec0d2 / 3

<a id="canonical-a71e1ffcfd387372f22eb308edbf582c95d33651cf3cf09f53d3825f25b5a15a"></a>

<a id="canonical-6101e62bfe16f761fe413dd7de62efdc3cb8c64ba0217571be94504cbd7f5aa1"></a>

## expected_response property — http_health_check / f8856a5ec0d2 / 4

Type: `"string"`. Computed.

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status. Server
applies default when omitted.

Upstream description:

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status.

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

<a id="canonical-35c00265419003c9017a12f5d803dcffaa4ed95fd4b4db3c60ef15e241982ecd"></a>

<a id="canonical-940cd89191885a9c7ca1aa3fba135f4705ab71823261d74c227fcf132141c4b9"></a>

## expected_status_codes property — http_health_check / f8856a5ec0d2 / 5

Type: `["list", "string"]`. Computed.

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by '-'. Defaults to
\`\[\]\`. Server applies default when omitted.

Upstream description:

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by "-".

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

<a id="canonical-de8a9ea40c20379a21494db37183ac95f648ae1d7980dd422756ed6ce7165ed5"></a>

<a id="canonical-6d89f5b9c46263d2fe2581925d432e5e86fbec5754bbf7dcb5dfbe02379364b5"></a>

## headers property — http_health_check / f8856a5ec0d2 / 6

Type: `["map", "string"]`. Computed.

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

<a id="canonical-a4930e080530e863cda95d5669854e7c7edc3dea95bd879b3a22ce59fddf9e66"></a>

<a id="canonical-b332e1adbfcd854e6b7fb4463918d87ac4329d4cdb309c2fa4ca81708443807d"></a>

## host_header property — http_health_check / f8856a5ec0d2 / 7

Type: `"string"`. Computed.

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Upstream description:

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

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

<a id="canonical-e4fdf5dcd8432a8dc992ceb55a5fa58cf6a9f7232d7ae07db5e4fc0070166735"></a>

<a id="canonical-05beb2081da5ba5a5d81138e68f974e3ab9031dea61f1ba091dd683551d2f7d1"></a>

## path property — http_health_check / f8856a5ec0d2 / 8

Type: `"string"`. Computed.

Specifies the HTTP path that will be requested during health checking. Recommended: \`/\`.

Upstream description:

Specifies the HTTP path that will be requested during health checking.

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

<a id="canonical-3bdb190f6f3000a2e051aa0fc009d5384a3db50fc0220182ae9971f0d12fdacd"></a>

<a id="canonical-e0661b1870700507147f3992e72df1f13347441ee97af600739f2bab53843666"></a>

## request_headers_to_remove property — http_health_check / f8856a5ec0d2 / 9

Type: `["list", "string"]`. Computed.

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers.

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

<a id="canonical-659b1a97aedf7fc1e12343e64041323ae5df01babd0b57a4004db796078567a5"></a>

<a id="canonical-363579545ea196f5537f5b8b75632bd5ba58eb48c7bcb74803d429ac721b5f5c"></a>

## use_http2 property — http_health_check / f8856a5ec0d2 / 10

Type: `"bool"`. Computed.

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

- [use_origin_server_name](data-sources--healthcheck--reference--group-001.md#canonical-b60aeee1cd650225ddbda61f3801e1a4da5818e925f40e339285f58b30fe7728): complete subsection reference.

<a id="canonical-66628a46933cd7f84ccb372706d8f9be23b2374263cd39c99b33b1568a489d90"></a>

## Next pages — http_health_check / f8856a5ec0d2 / 11

- [http_health_check.use_origin_server_name](data-sources--healthcheck--reference--group-001.md#canonical-b60aeee1cd650225ddbda61f3801e1a4da5818e925f40e339285f58b30fe7728)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-b60aeee1cd650225ddbda61f3801e1a4da5818e925f40e339285f58b30fe7728"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcf6880b784e5f24862787644c1e525e323c3eb088ecadecb286d27104967467"></a>

## http_health_check.use_origin_server_name — http_health_check.use_origin_server_name / 1c05c6a8d4b6 / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-83cc441036af77ca9aca9b2f4a34d969890f60c41814d6bd2f0731cd93b4f8d9)
- http_health_check.use_origin_server_name

<a id="canonical-f9c80d1486e8b54bca425fea7882f1c81a81bb33e74be1b83a42d066b21cac40"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2e584f6b8c2f33177f0479275af6bdb07c125ebdec599f40815580d1027bf10e"></a>

## Direct properties — http_health_check.use_origin_server_name / 1c05c6a8d4b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dff8ed19ae26114abc18c511f849f30d3401f1207cda9beabab154ada5059c29"></a>

## Next pages — http_health_check.use_origin_server_name / 1c05c6a8d4b6 / 4

- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-83cc441036af77ca9aca9b2f4a34d969890f60c41814d6bd2f0731cd93b4f8d9)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-7cc63fd84f59db2fbdb2ab13e4eb35522dff93d87fb58c595e3f6109eee12b94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d845f44af37ab031419c37ee75d61b9707f1d3aeaa0f9525dafe5d2dc4e37abc"></a>

## tcp_health_check — tcp_health_check / 0ec2255be689 / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- tcp_health_check

<a id="canonical-87dd96d8a965d018044db580cbee4ba5057a0092d2e7bd1e5373905d28f667f5"></a>

Type: `"single"`. Computed.

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

<a id="canonical-e5dbbcc42927a4bbc5052039a45ccb9d771620774a542787fb749788fc323442"></a>

## Direct properties — tcp_health_check / 0ec2255be689 / 3

<a id="canonical-81fc06c9d84d4a0447f77d7906a9cdcbdc7a7c79d745d2c28f8fd0e3fece472c"></a>

<a id="canonical-3df6fd91ffcacd95d402fd95bf9c017d08d4067b07049b28c84787f0c653a55f"></a>

## expected_response property — tcp_health_check / 0ec2255be689 / 4

Type: `"string"`. Computed.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Upstream description:

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

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

<a id="canonical-26df65095bb6c8decd39f05984cf8e9ee509fcc6e514a29ebad63b1b8652c4c6"></a>

<a id="canonical-446c76b1218bfa572a4b8b77c56f7a729104e289601fa2ec5babc5c358dbf90d"></a>

## send_payload property — tcp_health_check / 0ec2255be689 / 5

Type: `"string"`. Computed.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Upstream description:

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

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

<a id="canonical-a1fb2642e106a862962e332cb6b95ec774cd03fe878dfb2a4a1d0d2b7e7ae84c"></a>

## Next pages — tcp_health_check / 0ec2255be689 / 6

- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-8012b193094025ee22fe96956b0ce2a2cc4f203185a13906083b01136028ac2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cc11f57cce99c7d3fd881e6a3539b89fbe0af9850e4d9d27b790d7b3a1d66b6"></a>

## udp_icmp_health_check — udp_icmp_health_check / e14142d5d62f / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- udp_icmp_health_check

<a id="canonical-4fcd6931eef4e9bcf5ecf0c83233fce4483edec091c14727e70a992892916163"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-918a8359f67854e2a8ca9ee157603d14abdd8e049480d8f6be1e10401415ee4e"></a>

## Direct properties — udp_icmp_health_check / e14142d5d62f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-229cf60d743e27c5620ef16491832d58e5521652ff6303724f166867fb5d3065"></a>

## Next pages — udp_icmp_health_check / e14142d5d62f / 4

- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
