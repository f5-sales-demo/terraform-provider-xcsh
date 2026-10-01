---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-8a03a6d147f6473c9c3a5b6320c797555c75951249f34fbde8a7dfa4634a5af3"></a>

## url property — response_headers_to_add.secret_value.clear_secret_info / 221066718a6a / 5

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

<a id="canonical-556308ed6695b74378edc12e7db5d8097f9bb0f6a205e83c943dd37fda4d31a7"></a>

## Next pages — response_headers_to_add.secret_value.clear_secret_info / 221066718a6a / 6

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d8937146e8606cccb8a5d0c883571ec51d383d5cf570f87b7edaa3bf109ee3e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53638250dee9dc78f309122ba3dd8c5d9fb8e12a9445a7161ac4a53a6492cffd"></a>

## retry_policy — retry_policy / 35fff4f35738 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- retry_policy

<a id="canonical-8c83455fcc3830d8c2b0807cf1af96b9766bfde824f7032efb14f33c75704eea"></a>

Type: `"object"`. single nested block, Optional.

Retry policy configuration for route destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("retry_condition")}
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
retry_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-a49b0b6b227204868a04e808b237cdad119c45a804b3bf634660a354375d3699"></a>

## Direct properties — retry_policy / 35fff4f35738 / 3

- [back_off](resources--virtual_host--reference--group-003.md#canonical-ee8c5e5e60a2adf02add7015a4b195054fb75a81e721c93bfc0ee1eeb1dce418): complete subsection reference.

<a id="canonical-f946855524cefd20fa278e7bfcad851b463b36d3aaf3420c79375fe88d7269b4"></a>

<a id="canonical-ed4c055b58618ca0cd5bd904b5d7ea0b304f18e93d98a391025675ecbfdf1d22"></a>

## num_retries property — retry_policy / 35fff4f35738 / 4

Type: `"number"`. Optional.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-3c61f242a030156d4413297c547d3f86a6843a6518f4ed844960977ec75ad4c0"></a>

<a id="canonical-c23da74d774a25d4a34ab5153907b036236c31273109b13ecc1b55059fe8a447"></a>

## per_try_timeout property — retry_policy / 35fff4f35738 / 5

Type: `"number"`. Optional.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-93bce8999d10370e971e402a55534617241cdd40fbf791614cf983a0c38a1797"></a>

<a id="canonical-8e33cfe0f735377cabe153c6a5b90822b8711ec503b5e062d610d581e11657a0"></a>

## retriable_status_codes property — retry_policy / 35fff4f35738 / 6

Type: `["list", "number"]`. Optional.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b7d27dc2efa392cc1f44c31134c2b1d05b9083741f7771affc83ad7c853544a6"></a>

<a id="canonical-334397e9ebf580f18c3bd8f848b8fc304559f06c77af90143052332ab7890877"></a>

## retry_condition property — retry_policy / 35fff4f35738 / 7

Type: `["list", "string"]`. Optional.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Upstream description:

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7470fac8995c18b26184024cda2da915bc0590cc4d993a9e592473fe1f6d25cd"></a>

## Next pages — retry_policy / 35fff4f35738 / 8

- [retry_policy.back_off](resources--virtual_host--reference--group-003.md#canonical-ee8c5e5e60a2adf02add7015a4b195054fb75a81e721c93bfc0ee1eeb1dce418)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ee8c5e5e60a2adf02add7015a4b195054fb75a81e721c93bfc0ee1eeb1dce418"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c0cab940c79aa01b2734f851b71d61f6f5a78fe0c3678ce55700eb2b1480d92"></a>

## retry_policy.back_off — retry_policy.back_off / 60faf79e8220 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-d8937146e8606cccb8a5d0c883571ec51d383d5cf570f87b7edaa3bf109ee3e0)
- retry_policy.back_off

<a id="canonical-8a89e21a37a2e132fba037408d32ffecdda2b3e105a86f9616e2c1528b743cec"></a>

Type: `"object"`. single nested block, Optional.

Specifies parameters that control retry back off.

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
back_off {
  # Configure direct properties listed below.
}
```

<a id="canonical-5efb640315382af148be46f0d7cac6cc4a649a07032ae6c4c7c2db18fb84175c"></a>

## Direct properties — retry_policy.back_off / 60faf79e8220 / 3

<a id="canonical-00b2fde48849acecf981836a73288d1447dfbb757f0e7a48c1d4c7b1c5a8d93e"></a>

<a id="canonical-e72ea288f6bd14a3fddd841e880f06b410ae25cf5be8f6b0fe73e3e3c1d57385"></a>

## base_interval property — retry_policy.back_off / 60faf79e8220 / 4

Type: `"number"`. Optional.

Specifies the base interval between retries in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-ed756f6c8a09ba24d7163c62b2f34ca17ff962fcb4ef01957e2e50d2ce9587fd"></a>

<a id="canonical-3cd9dfd45a7289ce6c55ac1e88f63bb4f58f33c79314745f9a7abc1687a42b9d"></a>

## max_interval property — retry_policy.back_off / 60faf79e8220 / 5

Type: `"number"`. Optional.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

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

<a id="canonical-950931cac77c2c8657fbe24846cfc13be336fd254eadaf7b67145e3a304c6bd3"></a>

## Next pages — retry_policy.back_off / 60faf79e8220 / 6

- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-d8937146e8606cccb8a5d0c883571ec51d383d5cf570f87b7edaa3bf109ee3e0)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-6c636d254e1181e17ba1647911e7ca1e748a8724cdbd5ecd9bfbcfb08ca38ceb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63f85893f65ac5fccab05784dd42a5699036a08919a41aff8b9dd784244d5027"></a>

## routes — routes / cb7295a010dd / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- routes

<a id="canonical-41b5509dc4a9a734c09b499ebed9ff21f44dfeeddab239772d62ba9847107338"></a>

Type: `"object"`. list nested block, Optional.

HTTP routing rules that match incoming requests based on path, headers, or query parameters and
forward them to appropriate backend origin pools.

Upstream description:

The list of routes that will be matched, in order, for incoming requests. The first route that
matches will be used. Currently route object is redundant in case of TCP proxy but required. For
TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts, the route object only specifies the
cluster/weighted-cluster as route destination without any match condition. In other words, match
condition in route object is ignored for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts.
Routes used for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts cannot have DirectResponse
or Redirect as actions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-39c521b61cdad85565041e74056da1b9c751f43073721a3915a04651c7362320"></a>

## Direct properties — routes / cb7295a010dd / 3

<a id="canonical-7eca30a37b7561b2e6739bddaa22db86168a2624034019faabeb5dc286803c42"></a>

<a id="canonical-1ff83757693bff0055cf1d6e50eadf8c6c7d01b6ef1c5b9b526cb76ffb98e44d"></a>

## kind property — routes / cb7295a010dd / 4

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

<a id="canonical-ef811af33dfcd0a5eaac3e197ab79dc61932b9eb3a7a8eeac982375f58be3360"></a>

<a id="canonical-7b3d6cfcd9597b353cdbc6afc1f068931e6ff16a0a73f916dc543ea07e850939"></a>

## name property — routes / cb7295a010dd / 5

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

<a id="canonical-bfbd04cbf024f588cd60e5c81a52b6c70279b64bcf1ecc77e586b1dd287de7ea"></a>

<a id="canonical-21cf82ddae11193194cc66bc0370ec19bfde2ada6785d41924cc5b657ed05267"></a>

## namespace property — routes / cb7295a010dd / 6

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

<a id="canonical-13722e8427c7026bd607f6ef1396cf15794f9883312c4bb9733878818e324452"></a>

<a id="canonical-232eeaa088d4f05669d09503d329a59083545f3db03317447241a3d618394432"></a>

## tenant property — routes / cb7295a010dd / 7

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

<a id="canonical-92fa30813beaf4607d0fedd4eb1b07ee67690e31889f49b3a0071d9e7547f274"></a>

<a id="canonical-dfae4d0efaa1538b88ce8c3605adb23daf1822d7f8b2dee18da364ffb340e504"></a>

## uid property — routes / cb7295a010dd / 8

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

<a id="canonical-0a453a6783f3019adbca16e1eb33de58e10b893ed10ff7f197b360beb50996ad"></a>

## Next pages — routes / cb7295a010dd / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-578cb808eb1ecb7ba626831d3422f85119ed4ee3daff27fd256fecfc5c29a224"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31577b38c6862235acc386033d5dcf6fc1fb320f0bf3be4e887061d24216db6b"></a>

## sensitive_data_policy — sensitive_data_policy / 0839f496b230 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- sensitive_data_policy

<a id="canonical-2a6d7a8d87b4164e518fce36ae94769fe2d960040666c4879153b1fec7a6026d"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

References to sensitive\_data\_policy objects.

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
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7daaa6f242c1b459ab986fbd582fb0e08fbf7b5db7a0c77e125dc2aa8e6610b"></a>

## Direct properties — sensitive_data_policy / 0839f496b230 / 3

<a id="canonical-1d019c85a59158e2e2591f0af7f1919a89369d04d5571d10141f4c2ceafe2e82"></a>

<a id="canonical-ccf86ed9418f8f31becbe22a8d4a08e0740330930cbc6c49ae0818af2436d880"></a>

## kind property — sensitive_data_policy / 0839f496b230 / 4

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

<a id="canonical-efeaeb8b2391cf3d338119409e587bdbf8b6019ba4a377d2ae37ee99d0b5eeee"></a>

<a id="canonical-79723f4223aaf7208d61ec5bbbf0d96e57f156ec5674b43300ac0343d086abd2"></a>

## name property — sensitive_data_policy / 0839f496b230 / 5

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

<a id="canonical-d41abdb37770c5883d42e63bf9063c4320bea817bb7c49707805065bafdf3450"></a>

<a id="canonical-dcf6d59e5c7346247bb2bec32e4bdb76553a9a9363af039630c35f88a2843130"></a>

## namespace property — sensitive_data_policy / 0839f496b230 / 6

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

<a id="canonical-178693dfdcfa963f4e6fc8da9a44620ed388abc9641b220e22798d9fada2218d"></a>

<a id="canonical-c9daf9e798d7b2c29b77b017437e400cbf2c72a9deabdbdf6c0390e978079518"></a>

## tenant property — sensitive_data_policy / 0839f496b230 / 7

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

<a id="canonical-94f79d1068cb6968b79e07fe73a531df61ee81e962248d9ea06d5c1c133146fc"></a>

<a id="canonical-461b768628baec807b528b40bb7781508fb1165b468c5f5949393e0f7cf57207"></a>

## uid property — sensitive_data_policy / 0839f496b230 / 8

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

<a id="canonical-1c5073cb550636a8903a0d31f570f4ad5847bcb1db8b0238029e390cd316ab0f"></a>

## Next pages — sensitive_data_policy / 0839f496b230 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-bfd8c69950f38f4e58ba7969dd63cf2da90dd42a6c9450294fe6276246fe52f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b8018b2130def6f741ee89b5977639a853432bf354a0c113768bc9cdd3bcf6c"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 7e22ab15a0ea / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- slow_ddos_mitigation

<a id="canonical-27587dc92427d66a0d0fef3287eb4a83b206273d0bb9fa4abc44a8b59e2a1236"></a>

Type: `"object"`. single nested block, Optional.

'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-788c6e3f50e1565a3b93e093c4ae50c4fcd179f81cff16a784a740e53e58a029"></a>

## Direct properties — slow_ddos_mitigation / 7e22ab15a0ea / 3

- [disable_request_timeout](resources--virtual_host--reference--group-003.md#canonical-32582d1189edf854b2656c2a8a8329b94889e77849097221b8ba9b53e6d88d72): complete subsection reference.

<a id="canonical-34663318a254c795ec1eb2052e85fb96e4eb454248878667a57b26c3030d8196"></a>

<a id="canonical-12a5a7e94a22513566d0e0dbf2768dd798207a9a7d6277b62c76f40d2db4bc19"></a>

## request_headers_timeout property — slow_ddos_mitigation / 7e22ab15a0ea / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-44f6d2bd5ea4f845ec7d1aa987679b03cd0e325b74d02eee86c55b3ec52ed9d6"></a>

<a id="canonical-5fd3163872f0902bcbde13542254002cafb1e7e7171209c081a6405a906aae19"></a>

## request_timeout property — slow_ddos_mitigation / 7e22ab15a0ea / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-710700971493e9e6137a064353d76fae8258bc02fb51a0d1e4972c6a2838f979"></a>

## Next pages — slow_ddos_mitigation / 7e22ab15a0ea / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--virtual_host--reference--group-003.md#canonical-32582d1189edf854b2656c2a8a8329b94889e77849097221b8ba9b53e6d88d72)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-32582d1189edf854b2656c2a8a8329b94889e77849097221b8ba9b53e6d88d72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b82e36b58af675d1755d5aa01045dcd721d8bb92161e3b183f81d28634e0c7eb"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / fb6372508672 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-bfd8c69950f38f4e58ba7969dd63cf2da90dd42a6c9450294fe6276246fe52f1)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-0e63684fe038d15a057d6cd14ff450f2eee6bfbecabaed086ebf41f9219697ff"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

<a id="canonical-e3a29e3d07d80098d4cc783410b0c5001b4d2289643fe43fd62c99d2bb498cc1"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / fb6372508672 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cd30f04d4817b92e63d885df85fd3e3a97672381ead1839f867aaef954473d9"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / fb6372508672 / 4

- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-bfd8c69950f38f4e58ba7969dd63cf2da90dd42a6c9450294fe6276246fe52f1)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9fd42ea62fb46c99d5c861f931e67aeb0b60a2b0c3ecb3f21d9fea3e4fdd59b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9893e3772e9e5b30aae0a3d497fb3e46978b4c91404382cc6cdbeff6e95744af"></a>

## timeouts — timeouts / af164c66f8b1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- timeouts

<a id="canonical-6bcc0454c557758d4246132541c3e88f3d4b56202f2fcea983fdc6a5f4469b24"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f2d1df3a3d1f5c0a425baba850dfaf8118f69e49438546404908c6fe820c06d"></a>

## Direct properties — timeouts / af164c66f8b1 / 3

<a id="canonical-57f890124a18538cb7a04f07ffd3600b5a5dcde84756919b0e0bc47cc672649e"></a>

<a id="canonical-c82b69c4f2c2c3861b903ab70f5ed0c9284b3628e4dd245febf70d5b0acaefe2"></a>

## create property — timeouts / af164c66f8b1 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ccbef58964b5aaa10a93faf707bfbeda05e51e865892f785f30182c8b74efa67"></a>

<a id="canonical-658ef84ff1fd63fff74014ad6490979802686933f7d3b87d22ad5f4126d18a91"></a>

## delete property — timeouts / af164c66f8b1 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-724705d1ae1d8a4e9c0556c3452e7232356159bf6b9407ff240d4a2cec7a54cc"></a>

<a id="canonical-12a8c58554d0ea142c79e38343f662d6624a74f6898e8ae6324a551024fa46dc"></a>

## read property — timeouts / af164c66f8b1 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-744cc76a7ea111414623d33a7b51d9636d0362127dece7a75f19590ac12e8b0d"></a>

<a id="canonical-f467af0297348a81a07719d7c634f14e4cbff66b3e9e6d39ed39f31fa8280e23"></a>

## update property — timeouts / af164c66f8b1 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6cbad6ede58be05412c59e3b60da03ded9ef3183873ad3f2b290c5b146d4771b"></a>

## Next pages — timeouts / af164c66f8b1 / 8

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bde4d03dce1fa14080b05187941383288efe6ed12f0ac062293e8f68a92a9b9"></a>

## tls_cert_params — tls_cert_params / d3fbe7892956 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- tls_cert_params

<a id="canonical-9bf14422de2f877266665152d496cd59590f4c859819e3fb92fcda32939edafd"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: tls\_cert\_params, tls\_parameters\] Certificate Parameters for authentication, TLS
ciphers, and trust store.

Upstream description:

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
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

OneOf alternatives in this subsection:

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-9bf14422de2f877266665152d496cd59590f4c859819e3fb92fcda32939edafd)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-2910e5ae02d5215342ecba466be0588450d1c2ee31bd899f0945f871a1b7bdff)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-72f6ad2f0239fe23895e237dbc9664a9b101472093f71feb79ce1211f5147c7e"></a>

## Direct properties — tls_cert_params / d3fbe7892956 / 3

- [certificates](resources--virtual_host--reference--group-003.md#canonical-7f7ac098a85e634d50dd047896ed48989381aab2915d221bc8227b3f739ce9aa): complete subsection reference.

<a id="canonical-feba6908239bd7267067749969afa176ff290e459c0dc488fe23728b39cba97d"></a>

<a id="canonical-8b2d5827744b18e39aaece5193b08fdc5521908499130cde14e0a929f238ebfd"></a>

## cipher_suites property — tls_cert_params / d3fbe7892956 / 4

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

- [client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-a8d6bd0075cffa5d1b9945ead71e4196fde82e7d1b9af66f83f775513111fb1c): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-432e17818dac61c20f375ef7401b6c14e841cfcf84d606052b209edc26be76ad): complete subsection reference.

<a id="canonical-e293788e9390ea27bfdf414bf800a45d0edcebab112d08dcb8fe25b103ab2ef0"></a>

<a id="canonical-d280e111c4873dfae5f255f0bfa7352a20e08354c299cc70f9a81b9d3aeebc38"></a>

## maximum_protocol_version property — tls_cert_params / d3fbe7892956 / 5

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

<a id="canonical-d73b9f5f7b646366e7237047f250e0b1a7df7e992c2a12d1beecb92a876b0d44"></a>

<a id="canonical-3ade402435e1d970e19c3cd79287890ee22c815d9b37988b1cd57071d7969e62"></a>

## minimum_protocol_version property — tls_cert_params / d3fbe7892956 / 6

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

- [no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-9179f4574a62b76bf459504e7442cd1cc80637f925880730c68313c2fbaeddb9): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-003.md#canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096): complete subsection reference.

<a id="canonical-e0bf5b230e74820603e05c3c9093d5188b2fe5e78762b4583777a518ca312e75"></a>

<a id="canonical-5dc18ee5fc4fa3208f524b223cf3bf55d5c889b8184c230847c67d4b1fa5cbe7"></a>

## xfcc_header_elements property — tls_cert_params / d3fbe7892956 / 7

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

<a id="canonical-e3a3d1e71f9e6953565081b59d3428b17abc2e2f0b679503f1a29e8c93457b8b"></a>

## Next pages — tls_cert_params / d3fbe7892956 / 8

- [tls_cert_params.certificates](resources--virtual_host--reference--group-003.md#canonical-7f7ac098a85e634d50dd047896ed48989381aab2915d221bc8227b3f739ce9aa)
- [tls_cert_params.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-a8d6bd0075cffa5d1b9945ead71e4196fde82e7d1b9af66f83f775513111fb1c)
- [tls_cert_params.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-432e17818dac61c20f375ef7401b6c14e841cfcf84d606052b209edc26be76ad)
- [tls_cert_params.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-9179f4574a62b76bf459504e7442cd1cc80637f925880730c68313c2fbaeddb9)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-7f7ac098a85e634d50dd047896ed48989381aab2915d221bc8227b3f739ce9aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce652c33dd682fa13f02d8ac2c1d34bd68e613094af281010483d6a178c5e42f"></a>

## tls_cert_params.certificates — tls_cert_params.certificates / 9dc4e860e3b8 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- tls_cert_params.certificates

<a id="canonical-42f37b46de16d2b3bce96e08ffac4b28ac23b329f145127601a596e7ea8bc40a"></a>

Type: `"object"`. list nested block, Optional.

Certificates. Set of certificates.

Upstream description:

Set of certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-54d7423d53e46464fced00879cdd6094d99d1f2f576cf2a83657d0848008301e"></a>

## Direct properties — tls_cert_params.certificates / 9dc4e860e3b8 / 3

<a id="canonical-8c5bd782d7d1e0cebdb90b0020b05ee70a4dcd0f53725a7783c6f8de30c6af1f"></a>

<a id="canonical-bc8face96d19f867589cdaf29c1bc51f4e60cbd7fe96359cce21b553ceb55f3f"></a>

## kind property — tls_cert_params.certificates / 9dc4e860e3b8 / 4

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

<a id="canonical-804375d171de4405a0daf69aa9e33efa123b13350cc9ef6cbe88ab6b162efff9"></a>

<a id="canonical-352e3fab0cb14e65fb64080b0b4da31b6d2b63c496f92e3c92befc00990009a6"></a>

## name property — tls_cert_params.certificates / 9dc4e860e3b8 / 5

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

<a id="canonical-b88a2c93e341db6da23f430bfb748feea29e2ca6e4abcc40b8c04fcc4765a715"></a>

<a id="canonical-b9dedf037ac561f7d7dea05b40a7fa372310cf321a1596a86bb1027fe3f91761"></a>

## namespace property — tls_cert_params.certificates / 9dc4e860e3b8 / 6

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

<a id="canonical-244f9036d1288cc22a39d0a631f33f381486b630bbbf3bc3694f6098605a846f"></a>

<a id="canonical-790288aaf7ab74629496b3fad5d15fd1e7782d3d706f1e428dcca46a93656837"></a>

## tenant property — tls_cert_params.certificates / 9dc4e860e3b8 / 7

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

<a id="canonical-0f4ff91add3d278b93002e35a1ae6132830a3decb5fc91b4c421533ee6608c53"></a>

<a id="canonical-24ac51ff36f7e725fa173c639b740ee43de217007223fadc155f84df146d4adf"></a>

## uid property — tls_cert_params.certificates / 9dc4e860e3b8 / 8

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

<a id="canonical-be82b8c1caa0d10d06980f9592ec75730ab60069012ae5a44f66906cae3496d6"></a>

## Next pages — tls_cert_params.certificates / 9dc4e860e3b8 / 9

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-a8d6bd0075cffa5d1b9945ead71e4196fde82e7d1b9af66f83f775513111fb1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7a6ff0d80d0b1e555d9e525430bd8d3607d7f56859dba53e6edf499fd200ea2"></a>

## tls_cert_params.client_certificate_optional — tls_cert_params.client_certificate_optional / 652166104b05 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- tls_cert_params.client_certificate_optional

<a id="canonical-6bcaf662a7d595683aa98e82bff36fc931472e007d22f2a131ba28cb7a3cb653"></a>

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

<a id="canonical-bc5a22694f03509be4326c8a6dad0fbd379d16e621d638657c375d2a8aef42a6"></a>

## Direct properties — tls_cert_params.client_certificate_optional / 652166104b05 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be00f9bc06c24e0dad2f460c3abc1ff13b4cabd2b400b7b20d47106fcbd1664c"></a>

## Next pages — tls_cert_params.client_certificate_optional / 652166104b05 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-432e17818dac61c20f375ef7401b6c14e841cfcf84d606052b209edc26be76ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3258bf266cae3a91a3a2e5e323b659ba539a16eda4bbea0aba39cc9d971d547c"></a>

## tls_cert_params.client_certificate_required — tls_cert_params.client_certificate_required / 41b68fbce5a4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- tls_cert_params.client_certificate_required

<a id="canonical-92106834e14776aacb3887b41b88aefe28938c122cac350cd0e42da030028a7d"></a>

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

<a id="canonical-4efbf2e648b475b8b0abcf00bc675922f96cb565012fada41880cf1b89fd2dbc"></a>

## Direct properties — tls_cert_params.client_certificate_required / 41b68fbce5a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a935d8ef39b0761a7b17f656452446ea735ccfe7e9be7366fb482d9f9b3b36d"></a>

## Next pages — tls_cert_params.client_certificate_required / 41b68fbce5a4 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9179f4574a62b76bf459504e7442cd1cc80637f925880730c68313c2fbaeddb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baad51ca025b36cd27a144725c8445e3e3938a7e9b507a8c480d1b81b4ad0761"></a>

## tls_cert_params.no_client_certificate — tls_cert_params.no_client_certificate / 5bab619a88e1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- tls_cert_params.no_client_certificate

<a id="canonical-f49c06c4de642692d3b1e616f142089df494175cbe753559041b4f3b5eee2593"></a>

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

<a id="canonical-e07b1efdbc0f634e0398ca276facfba41ff91816f82c51c78d483a7494dfb16a"></a>

## Direct properties — tls_cert_params.no_client_certificate / 5bab619a88e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-530a6a853a255ec0eb90c86d5ebb591f91d87c2722d0345bc369d176e6072098"></a>

## Next pages — tls_cert_params.no_client_certificate / 5bab619a88e1 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e1c259562483bd76a3dcd8534cd4dd0afc7f9d8c4b9780f98dc6b0d144d04a9"></a>

## tls_cert_params.validation_params — tls_cert_params.validation_params / 2d72eea342a3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- tls_cert_params.validation_params

<a id="canonical-4e8b2c25c1ffc84f962c974999294c0daf6cfa1dce7d1a80ca65e2aebd9554bd"></a>

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

<a id="canonical-21c778938007b38e80ca6c8aa78a129e65dcf4715c56e1ebb061b856cd4b392a"></a>

## Direct properties — tls_cert_params.validation_params / 2d72eea342a3 / 3

<a id="canonical-14f9e5639f989f3ce643fe230a9d98c82f003e9ce551064e3662c72e7ba9642f"></a>

<a id="canonical-9252f171b37628209bfbf93bac6a1bedda3c809c663cb0cd61bd4145628034b6"></a>

## skip_hostname_verification property — tls_cert_params.validation_params / 2d72eea342a3 / 4

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

- [trusted_ca](resources--virtual_host--reference--group-003.md#canonical-d2ea348e440472e4869ceb8ac6a305a1663e69ccc90719057c82a83a008c4b5d): complete subsection reference.

<a id="canonical-12b8f30b17f86840764bbfc98d7e534df7237e3f70bbd332f21eeefb45624ed3"></a>

<a id="canonical-fb693b2cc3ab926c47b699b08a9e6241d9400c473705145ae0cb93160606527e"></a>

## trusted_ca_url property — tls_cert_params.validation_params / 2d72eea342a3 / 5

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

<a id="canonical-7cf787ddec29db407850fd42070a43a8e18246c5e447945dc8c1a21cc2c71421"></a>

<a id="canonical-2f34b7a45e72273a2fa146f177f37d1c4e310314732090369d7038ffa3f5fead"></a>

## verify_subject_alt_names property — tls_cert_params.validation_params / 2d72eea342a3 / 6

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

<a id="canonical-4d05cb53da493cfcb72c6635df1f2533acbc613fb3ac795a72dfc9ec2873dd16"></a>

## Next pages — tls_cert_params.validation_params / 2d72eea342a3 / 7

- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-d2ea348e440472e4869ceb8ac6a305a1663e69ccc90719057c82a83a008c4b5d)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d2ea348e440472e4869ceb8ac6a305a1663e69ccc90719057c82a83a008c4b5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db3f249068eb64df6792cc299882a0de7f794b0627398a645f79ae9d64108d61"></a>

## tls_cert_params.validation_params.trusted_ca — tls_cert_params.validation_params.trusted_ca / 9189283de3f9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-423dd4682577d152ea99df8ccc5a04cc99f052117f8da5183aec2ddd5d222e52"></a>

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

<a id="canonical-5c4303dbfe8233891600201d2d0f0c076107512b8cb8fcba44e37053fc5da59f"></a>

## Direct properties — tls_cert_params.validation_params.trusted_ca / 9189283de3f9 / 3

- [trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-2dbbed030800be108261da399c67126c21f873f05c3972e456fcc6b1e1dd790c): complete subsection reference.

<a id="canonical-d9ccbcb2af1cda5eacfaac8ceb0271c9c74bfafd3ac013451fac606dcdc848ef"></a>

## Next pages — tls_cert_params.validation_params.trusted_ca / 9189283de3f9 / 4

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-2dbbed030800be108261da399c67126c21f873f05c3972e456fcc6b1e1dd790c)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-2dbbed030800be108261da399c67126c21f873f05c3972e456fcc6b1e1dd790c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4986fd027d5cf0845ee091c673a8e895ced10305fa4058b62c5c911a4e34905a"></a>

## tls_cert_params.validation_params.trusted_ca.trusted_ca_list — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-959b2628118c17345fb71090d9460a0e399f7426a02735c90b09618ecaf5a096)
- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-d2ea348e440472e4869ceb8ac6a305a1663e69ccc90719057c82a83a008c4b5d)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-554594813952b6b9ac77a4683386b61fdbb2c23cee86ec0a010b9cbf429ffd83"></a>

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

<a id="canonical-847532cd9aad600f5d8d61d60242c1ba83432b12b49e331d06f332b5257c587c"></a>

## Direct properties — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 3

<a id="canonical-7f58bea9f7411fdf42e3929328e8496ee0ced1fce1ff419a3d07d6d02e2573a2"></a>

<a id="canonical-75a4d66f3dc24016106adf86ed8fe4a315721b77c7c4e8ace28304be0017a772"></a>

## kind property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 4

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

<a id="canonical-fb74d5ccd7912de8d3b8a89cf3302b29fb7e51c9dcac0307c5ae4f251b96d43b"></a>

<a id="canonical-6a35ebf2b61298864984bf043b5a543cdd2e3bc58b340a517a286168aaeab4f4"></a>

## name property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 5

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

<a id="canonical-d931c1bca4f2537f5ed06b27535c6912a7382232246d9a4c94d7d03f42f7e045"></a>

<a id="canonical-f99250d44198a9ce7f0d6247a762fd3632ee7a0044ec77cc521029fc18511818"></a>

## namespace property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 6

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

<a id="canonical-50846b2f27c6f631cfed959f4ba905ea0285370950dd7a1766868568633d20b5"></a>

<a id="canonical-d3d249f2c79f3c556590a2900c9ef6712dbe1f7f8e807ad1336a638dbe3c8bc3"></a>

## tenant property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 7

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

<a id="canonical-2f7cb6edc37b5ee8946112cffbcf471f583266756e81ba43597238e20490583a"></a>

<a id="canonical-76b7c96148e02ae820031b3743ca89eb2680f6e78b6dced2529e4549e87128bf"></a>

## uid property — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 8

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

<a id="canonical-05223ff39afafacb69f956dafb931d7932017818797a02daeed2b14f0828a177"></a>

## Next pages — tls_cert_params.validation_params.trusted_ca.trusted_ca_list / 34743bc40a80 / 9

- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-d2ea348e440472e4869ceb8ac6a305a1663e69ccc90719057c82a83a008c4b5d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62c31c0921b31484df9cd389745a6f8b8c23de944127876baf4ac84a4e707dec"></a>

## tls_parameters — tls_parameters / dab8f0ae53e1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- tls_parameters

<a id="canonical-2910e5ae02d5215342ecba466be0588450d1c2ee31bd899f0945f871a1b7bdff"></a>

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

<a id="canonical-b0b5c90e222bb817cb9c588c52980b7792202cfeaab7d9ec660d81d31bacc74c"></a>

## Direct properties — tls_parameters / dab8f0ae53e1 / 3

- [client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-ba7f7eeade78e775a0f87507b5ef52589cc9c12c1a4872afc7d8c26a99091565): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-e2702437f169b1453ce988b33f16ce626b913d3bb3394873c600caddedd5f45a): complete subsection reference.

- [common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9): complete subsection reference.

- [no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-a79a15f02ba7959db01ddd30bc7373c1c306bceec0d2da8091624c9ab0b4804d): complete subsection reference.

<a id="canonical-2783469c11c3d084d2beafc39a7bc24ce1506901dac126020d1ad1a0b042ca40"></a>

<a id="canonical-de2a72b30ebf5d2d2401cb250af9a20d771b72143c82d4f995f8de32dad1e1e6"></a>

## xfcc_header_elements property — tls_parameters / dab8f0ae53e1 / 4

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

<a id="canonical-54ce724d2878a495a3ccfd7386f4742ba49c6126b40aa6d0517dffe2d65782b3"></a>

## Next pages — tls_parameters / dab8f0ae53e1 / 5

- [tls_parameters.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-ba7f7eeade78e775a0f87507b5ef52589cc9c12c1a4872afc7d8c26a99091565)
- [tls_parameters.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-e2702437f169b1453ce988b33f16ce626b913d3bb3394873c600caddedd5f45a)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-a79a15f02ba7959db01ddd30bc7373c1c306bceec0d2da8091624c9ab0b4804d)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ba7f7eeade78e775a0f87507b5ef52589cc9c12c1a4872afc7d8c26a99091565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ca187601b9fa256b595531a2689e0cda1b97c9a5e5cb40c52564158d3772b6f"></a>

## tls_parameters.client_certificate_optional — tls_parameters.client_certificate_optional / a96d7987dc9c / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- tls_parameters.client_certificate_optional

<a id="canonical-3f655a1f3b234364ff15174de1671876e5948d7f1b4b8b10cec8bfc2c4029462"></a>

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

<a id="canonical-b92eaeeb2a0b701e7f47a45524df5ac498304d66c3651cc5048d4d712c803d70"></a>

## Direct properties — tls_parameters.client_certificate_optional / a96d7987dc9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74238a69e4d5b98c56549b88ae822cf7152166708d0b591e60ade437c21e4d64"></a>

## Next pages — tls_parameters.client_certificate_optional / a96d7987dc9c / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-e2702437f169b1453ce988b33f16ce626b913d3bb3394873c600caddedd5f45a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94a4c4a5e279a8d19fdec0453e61f279e8b80c75316e23c7c65b4a5ecc26c1e5"></a>

## tls_parameters.client_certificate_required — tls_parameters.client_certificate_required / 74334dccbc40 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- tls_parameters.client_certificate_required

<a id="canonical-dcb8ece32e17faf90c5ddeb537118789410100f199939d9d87d6041d77cba500"></a>

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

<a id="canonical-77f37b8c6f4b458c2058257c1a7e0ebfbb4b7ab91fe577f9f388425b2b57d859"></a>

## Direct properties — tls_parameters.client_certificate_required / 74334dccbc40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce9e652d678c29d09bda602018e06bf7e059992e5a2c1af27ebc106921e55d11"></a>

## Next pages — tls_parameters.client_certificate_required / 74334dccbc40 / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d4fc5251192d8c9ab4ef68cda843f7db76e5ac4ef64b00b3227b66035408ee1"></a>

## tls_parameters.common_params — tls_parameters.common_params / cba7c7e471ee / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- tls_parameters.common_params

<a id="canonical-3361f7ef1ea137fa4acadf36723f081e3135c28bc782367ccd89835c042d7719"></a>

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

<a id="canonical-71c110d1fa78f5a3ba9f93710b4e87d7cb03f742840a29190d1eb7b64e2550d4"></a>

## Direct properties — tls_parameters.common_params / cba7c7e471ee / 3

<a id="canonical-b4e707c32c0133936d0d88a7d702791a5d5fd054ea6a8ef566a878bf27d9e421"></a>

<a id="canonical-2a8c009b8d2a8690e0bebe17d1ee7d1f2a5a2a477bd50ab536a62f55e6d9ddb0"></a>

## cipher_suites property — tls_parameters.common_params / cba7c7e471ee / 4

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

<a id="canonical-dffbc60bba666e379dd58dcbaaa41b08e95ba833eda099c00a41a0debfffeed4"></a>

<a id="canonical-01c2ad9933506f92685ffbfb2864dd27d6ce10bf532ad7f182eeb08b4215042a"></a>

## maximum_protocol_version property — tls_parameters.common_params / cba7c7e471ee / 5

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

<a id="canonical-278f4578e3c4a057236a88c0e0c9b52aede56db873ff36bbeed6c3b63bd47379"></a>

<a id="canonical-f198b3afa66af9df41464da7a3fb798c2848a31bbf7521e807870d4270673f4a"></a>

## minimum_protocol_version property — tls_parameters.common_params / cba7c7e471ee / 6

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

- [tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-003.md#canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b): complete subsection reference.

<a id="canonical-b0625db1d7209481a891420b8e6861dd0136fa662f4e5e6e53218c2d54eb969e"></a>

## Next pages — tls_parameters.common_params / cba7c7e471ee / 7

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ad8f8c6b4984a21b5364e4eec5e0fb1f15b6ed02bc6c0ab8fd7550a6af70c7b"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / e1b0403caeb6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- tls_parameters.common_params.tls_certificates

<a id="canonical-45eeded18167500449953058a3c3596e2db4f63c12f3bf8a41df81306ad9d260"></a>

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

<a id="canonical-05f00c032e09ea363b36d919519fd7e6d7130b7f8a3ef65ce43de268aeba43f4"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / e1b0403caeb6 / 3

<a id="canonical-f295dbfd63d722f6326764ae8f868e830fc85b54a556f036155a505a8691bbc2"></a>

<a id="canonical-270f3c7e0b44907cde251add5eecf92d280e1d763be5501d7e2c18c26098f891"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / e1b0403caeb6 / 4

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

- [custom_hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-4015f10ad55da8e415af70b01f2898d43c5512fa6c2d42446b8a63007825ece6): complete subsection reference.

<a id="canonical-4513e9bb77db3b600f35a7bccff34bc9cc08406f9beefeca98a71f31baa26ae0"></a>

<a id="canonical-29cae590cf7a9928251de29afe018edbc9a10a244fbb6b0a6c4b43e1f3f002bd"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / e1b0403caeb6 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--virtual_host--reference--group-003.md#canonical-eccfe038103362be66fcd1dd735652d3a1a9c99c2c339f6d03db0e94033a9fbf): complete subsection reference.

- [private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7): complete subsection reference.

- [use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-79e751023f71af6a2f1244ea48c0f7d7283c8a928d43cc95eb9cb8ec488dbc63): complete subsection reference.

<a id="canonical-8a1597fd2ee6e0c724de2f8849429475fe5229299db383da4b137138b4b24efc"></a>

## Next pages — tls_parameters.common_params.tls_certificates / e1b0403caeb6 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-4015f10ad55da8e415af70b01f2898d43c5512fa6c2d42446b8a63007825ece6)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--virtual_host--reference--group-003.md#canonical-eccfe038103362be66fcd1dd735652d3a1a9c99c2c339f6d03db0e94033a9fbf)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-79e751023f71af6a2f1244ea48c0f7d7283c8a928d43cc95eb9cb8ec488dbc63)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-4015f10ad55da8e415af70b01f2898d43c5512fa6c2d42446b8a63007825ece6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b52ebbcf14bb6e90690cbcd36a21fe84cf6281477a4a04a57ff1515c27c6d24"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 0a8c7ec972f3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-71f409a81bd3389c626aa47709ab2477fc7ef40ac74d3f59824397324de2119f"></a>

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

<a id="canonical-8820950936405ac1518d81a65034ba52aef27f7ccd19328d0eac63c4da70ac97"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 0a8c7ec972f3 / 3

<a id="canonical-91918049f6c58ecb50442bd8faa73b7cf803578d7cadc4c1f2107960fd342ff5"></a>

<a id="canonical-dfb1f21d23561eac7e79686382bb48926225059404f03cfc999e13d7745861ad"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 0a8c7ec972f3 / 4

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

<a id="canonical-596ce34506d385960565a4d3d933a2d6bb46883d26bb97bd3329c3fba03fcda5"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 0a8c7ec972f3 / 5

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-eccfe038103362be66fcd1dd735652d3a1a9c99c2c339f6d03db0e94033a9fbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ae5df77a971df6962174fdb49bf6614107f8d6da5b27800e9c5aa6f1a08c5dc"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 7eff09bf235b / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-11c8400dcb15d587af4cb7859697a86a3f30795256c569b17b7c32b66af9f0cf"></a>

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

<a id="canonical-f8bad72b03b7459e6daa123c3805d33feeebf00c1b4572174bd0fc7e23817a51"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 7eff09bf235b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6715a6d9bf03a7198e9184024bba245535313238a983ec8868eb58b4e6a38ee9"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 7eff09bf235b / 4

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e108517799b7afdc099a9dc0019288ecc9552f38c9e0d9bb34850744ee066854"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / eeb0d4952424 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-b298967600db69cf12af6b236de25221726b249bdc3358b76461f3d2c653d667"></a>

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

<a id="canonical-5f59497aea5564237b2f1d2fd9201d47af99bc213d344cc29f5270625769e52f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / eeb0d4952424 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-b3ae6eda95608794b3d5eaa5da2186e1718281cc19c76928e23e3789845fd2ff): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-b91ab66f38a8e4acbcaa81ed505b6c53b664276bf21adf109cc31165d81b67e6): complete subsection reference.

<a id="canonical-421fad92bc40ea16417e25e4cfe14db54d373b7726d3e188a96ce289e5602ed4"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / eeb0d4952424 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-b3ae6eda95608794b3d5eaa5da2186e1718281cc19c76928e23e3789845fd2ff)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-b91ab66f38a8e4acbcaa81ed505b6c53b664276bf21adf109cc31165d81b67e6)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-b3ae6eda95608794b3d5eaa5da2186e1718281cc19c76928e23e3789845fd2ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c204403b2927c157c1c3955848688844b0b8cae3d331c7cc5b08e177a8db8a28"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-ed45b6b86142e48b849add86415f897ee702baaa873520f19ddd2c2bcbe549cf"></a>

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

<a id="canonical-230ec81b0cd90e1493fc4435915ab4a39812172e26270ade0a8bb55c3f84e07f"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 3

<a id="canonical-ccb2ebb6f35915b8e6304d0e441d026ccd69a7cfa8fd5b4883691849279fed68"></a>

<a id="canonical-afc02ce74fdfdf61aff6ff672ac1e51708496253713e75a89bd95f1d052f398b"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 4

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

<a id="canonical-d0f566b5e8baf55b39c7a7118207d5fea2c83837fa382134bace13cf9730f18b"></a>

<a id="canonical-f6962c297378091bc6c42abc0b97dd7e5f9d11e88a30fcb7e46a0dfdc89e5e10"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 5

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

<a id="canonical-33144ceee90fbd34d45599d41bf2be5fbbde5d5a6ff7c6700efeb28c2e313d63"></a>

<a id="canonical-5ec7f36e97dadcd3b258eb9189b2bbb2cfcc1d194f8e163290f6e9f14d446231"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 6

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

<a id="canonical-51c8fd2fbb6193fb9776136e3197369d14c6d2494b904f495e3d47eff7ca1d70"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 6d11b8d37bd9 / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-b91ab66f38a8e4acbcaa81ed505b6c53b664276bf21adf109cc31165d81b67e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e535aa74dd3fb2066f300ec6a2c07616a6f30a347acabaf8901b48d66631070c"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / e19b000f1846 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3122101c7343ae3dab6220750a904a61d8253ea3c262d72c827be90f0ca76685"></a>

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

<a id="canonical-87b183b500d0b93680609b75781b285c50fc5f34b18493d6951e6ea924514a4a"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / e19b000f1846 / 3

<a id="canonical-f713c46bc014cc1077e4b6356ca34a448555f5ee7cc90b901d94ffda1203973e"></a>

<a id="canonical-c5b401297bebbb26a13507f5b740f41c767348ee60dfac648d5049bf6e6d9270"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / e19b000f1846 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4a87f94fd2eb5900fb4fb36267c88465bc4f34c122015777eba5efae4ded24a6"></a>

<a id="canonical-1c126cc5b085b3d05b47278ab26ae3cfd834e1557f272bb72410ba25ee3d5887"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / e19b000f1846 / 5

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

<a id="canonical-e376b2e823556aa61e6d150a4976b1cae7e19f361253c51cb4bae611ff53697b"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / e19b000f1846 / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-10441f456e2bbeadabb804b62750ff4e39327285f088caf9555084b86cb15bd7)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-79e751023f71af6a2f1244ea48c0f7d7283c8a928d43cc95eb9cb8ec488dbc63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0fa9161c021f4e96b786cb5b7ae0089cd148d602c3e347406fab5fc1d39952d"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / dee5ad235df2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-efcda848f285218a82b1e8894adf645bb85e26a26bcd1d5fbb921d261e7973a3"></a>

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

<a id="canonical-80ee140680bc71f938655715e419a013281e7f6ebde7b28e270b2575d70d5b82"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / dee5ad235df2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9db793bd0b016a7cd2df66ff80ca9c73059b99aac09a75d03a6ad19d3f92dde4"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / dee5ad235df2 / 4

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-200edaca79ee56ef1be52c137193f5e6553b0067f7eb07e7f5f8d8bf9489ebe3)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1f3184384ffda2b5af607723c3ec66942b78cf52e465f50a6e0a067781f8193"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / 70fff96101b6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- tls_parameters.common_params.validation_params

<a id="canonical-8d4ee58b7177a6f9fa13c59abf709eaf5ceedec1ff0c993fd3777804b4e516b4"></a>

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

<a id="canonical-66842a908ae71edd9bef9c676f5556ec0a436fa5fc5a86b910e517bcdde065d8"></a>

## Direct properties — tls_parameters.common_params.validation_params / 70fff96101b6 / 3

<a id="canonical-b0252ef02943c090ff9e0e658d6de362d7c445832b4bf845f3d52d1ca7f8a981"></a>

<a id="canonical-563fb8e76cc8af2b1431e50b5574c2a2d2cbeef9be0f8233759808ecf820b730"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / 70fff96101b6 / 4

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

- [trusted_ca](resources--virtual_host--reference--group-003.md#canonical-9986323fb0b420fbdc7cdb70ae0c454fef9f1cc94c65178990322e8988ae1d75): complete subsection reference.

<a id="canonical-a1f010dc5317da2f2a78e2b42271461bf71f176e90185a34b573c3dc34c053c7"></a>

<a id="canonical-fc24204a7b128b9e40bf42dfdad7db447914cc7d71bd84ebbba6a6ba9abc2ff2"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / 70fff96101b6 / 5

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

<a id="canonical-ba32d9898f625ee930f3fde08cdff8b0114df8b49223b685d89b280a27657507"></a>

<a id="canonical-53f3cee2b913c4ca1e0bef3157294ab76b3040d7bf85a81ef9e1a795f9a1fb10"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / 70fff96101b6 / 6

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

<a id="canonical-5fe6881fd0da0518692c40e121a8406eddd04e1cf3d9deb04591e70794148653"></a>

## Next pages — tls_parameters.common_params.validation_params / 70fff96101b6 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-9986323fb0b420fbdc7cdb70ae0c454fef9f1cc94c65178990322e8988ae1d75)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9986323fb0b420fbdc7cdb70ae0c454fef9f1cc94c65178990322e8988ae1d75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b61c4cf1b57ca41c77ed8b2749c8a858029c3e825a2a3e6c36efac41e6b7c99f"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / 952556498ca9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-9209c3fbbc62abaf32cab4acb9680b3e82a68a142a21a92a349d703069d27363"></a>

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

<a id="canonical-b6cfa4c1315ba55861b2db1008a11560608c4db01112c2ea617c0b86502ca3ad"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / 952556498ca9 / 3

- [trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-03fc4e99dadabc1e29b36179cd6058e288819c5cccb05834d7fcd742c4863535): complete subsection reference.

<a id="canonical-fd62d188aec4229efa190101612453019c4b809cadb32b8f6f3c73ec6daf9362"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / 952556498ca9 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-03fc4e99dadabc1e29b36179cd6058e288819c5cccb05834d7fcd742c4863535)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-03fc4e99dadabc1e29b36179cd6058e288819c5cccb05834d7fcd742c4863535"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46d7a260b34099379387b6411e40a049f82ceb120c2b25f9403bf15da1f183d4"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3f9735b8aedbd23ce0065f977e0cae6f1bd4b3e3414e12f760bfe9b9dc1158e9)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-507b3dafb2f6b8bbaeb62595af6aa852f866fec30afed3a26414719e4a2c199b)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-9986323fb0b420fbdc7cdb70ae0c454fef9f1cc94c65178990322e8988ae1d75)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-c6f256f24a3eb76ed761f1c000a638dbe489cfbce41c3046a2b8777a6ea991da"></a>

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

<a id="canonical-3961074f36fe906c42b9c0dee798928183db10f96f609359b88db4ef7738a576"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 3

<a id="canonical-e60aea31f7102649835da76b95c34bc80aa33385ac295251180d77b19cac6a6f"></a>

<a id="canonical-e16910986b47b6e8a62dd8f1a77f144f7b7b59360badb50077118c007c27de7d"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 4

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

<a id="canonical-b75214e91d0164a77bb472704a987ce62106d5246d0fecb15d22011c86f79f10"></a>

<a id="canonical-76f62c2a8021cd305e6cc0ea950ce7e2799e90528fab66e401f8fa82adf0f96e"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 5

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

<a id="canonical-4b2ef68cfcbe3b7c1e413d57cd1170fd366108ca6135363c15a59c12b89c00af"></a>

<a id="canonical-76fc4a1de6e1e1f1c7a3ce5d1035b10ad0f71ffb70f2c0b71990306bd84a5297"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 6

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

<a id="canonical-cd9036b164eb1ca226dc36a897d9ef00886c0ab35917552493b8cdc3c16932fa"></a>

<a id="canonical-a1498a6f3b53b8df93d764f0cbde9ba17b24d945aef9af602c34d0ac28cbed03"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 7

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

<a id="canonical-fb36d170a925e7190f03f3c5d0f125947e64c1f4d7b9bbc28bfb94d958e79660"></a>

<a id="canonical-a460872ed099ed4f82babecbd2ced15b496d265e5bac60b315df85cf778fc89e"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 8

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

<a id="canonical-9bfd05d608bbd2743d89fcfbdc1f582adacd26d4b59818fd1e3e54326e698616"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 3ab24de90050 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-9986323fb0b420fbdc7cdb70ae0c454fef9f1cc94c65178990322e8988ae1d75)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-a79a15f02ba7959db01ddd30bc7373c1c306bceec0d2da8091624c9ab0b4804d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cd0490b226f6df6480aaed70de8919c9c660504daa9f366127c2fc077f67608"></a>

## tls_parameters.no_client_certificate — tls_parameters.no_client_certificate / fd1b61c27a3e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- tls_parameters.no_client_certificate

<a id="canonical-bb5b2d7c60df78ea3b36937fd521914f2fb7279df3ff501900130a7fa722a10d"></a>

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

<a id="canonical-ba7b507a3dea6b8575425375184a21266dc9aa44986558f9f7597575ac721e86"></a>

## Direct properties — tls_parameters.no_client_certificate / fd1b61c27a3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62cbabe3aa8c3a4bcb9f1a50841cfc46a0810b76bf724c611fdca7063a074577"></a>

## Next pages — tls_parameters.no_client_certificate / fd1b61c27a3e / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-46844342ae47b9c04a530b6e2c0292516326b045b1dd72188227e28335b02bb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e977a8720bb8bd75dba49e900e1879d7891c37ed8061f2f3ea2480112cce8ed"></a>

## user_identification — user_identification / 1c69a39365e5 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- user_identification

<a id="canonical-dbb523abe25819e8bb81bae3d99b15d764b7f9255b96454aeb275d56294fe312"></a>

Type: `"object"`. list nested block, Optional.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited.

Upstream description:

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

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
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2507a7082b5aaed63a81b8fa7ca889ae83911053e51306e2e2d19ea8e541dff1"></a>

## Direct properties — user_identification / 1c69a39365e5 / 3

<a id="canonical-18a95d2a38fcc67452c08ded4c7313fbe9455318669ff9f0636c40c825343c0f"></a>

<a id="canonical-59b125d110114b20110fcd4b1b9c3c80d0d307413ed41121837f9e444d7fbace"></a>

## kind property — user_identification / 1c69a39365e5 / 4

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

<a id="canonical-9869dac62bc714e6eb7487342ebb4c0ea909a3bbee9511562dad9985e5427753"></a>

<a id="canonical-8e6f5245a4aafad48cf18352fb490f202ec8e2d7b827d09b017f2de51b4c0fc4"></a>

## name property — user_identification / 1c69a39365e5 / 5

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

<a id="canonical-23a40709550f16e153f7555295f94632d25ebf1d676ec2b376ffe0a8e3e0a7d2"></a>

<a id="canonical-aebdbfc97d3726517661b772cb021123e160f2f60d93599c12bcebec1eae83be"></a>

## namespace property — user_identification / 1c69a39365e5 / 6

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

<a id="canonical-51a9e06f95220dc2419041eb04d9c9611773537b3534dbaf9f1621efa4c0bd28"></a>

<a id="canonical-de8f6f29444e81e4c25c064b8ee8e71db777902dcbddf452c0bf64fb83ce7e1a"></a>

## tenant property — user_identification / 1c69a39365e5 / 7

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

<a id="canonical-3a64ef21e12d15b69149f7b6fe9ae130a19a08af42aa42ce89419bd33e72e0d2"></a>

<a id="canonical-be7b513737f0a34a4e1b894d7e309fc88d3ed728e345dfa152d4119005165d66"></a>

## uid property — user_identification / 1c69a39365e5 / 8

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

<a id="canonical-710285129ed1975dfe4f6824cdb28010624ed30f95406c377ba746e185167392"></a>

## Next pages — user_identification / 1c69a39365e5 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-949480136811c183090cfa67010a87b841ce6c38a1cd3a8c7e4bf9448423c63a"></a>

## waf_type — waf_type / fe1fa8e9662e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- waf_type

<a id="canonical-ec1c32c04f956ddacb473484f769acc06292a25d4df21e8433fb6559f2e68938"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
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
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-b2ed04b637cd3efe60d4c1c917bdd51063dcd040c84ec42ac99a86af115903ef"></a>

## Direct properties — waf_type / fe1fa8e9662e / 3

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-c766f377dd26e9aaeab67bb17f27ea07424c43e514f4ac327ce4593d9d5bd14d): complete subsection reference.

- [disable_waf](resources--virtual_host--reference--group-003.md#canonical-201ec42670840ef9105005a05225be6643369dbde7e84c66da6be3acad0c6a53): complete subsection reference.

- [inherit_waf](resources--virtual_host--reference--group-003.md#canonical-74e227e6a93f32320580c00c8459eea86e2dfae357c661b05e8017eb7327fae0): complete subsection reference.

<a id="canonical-1a479f96ca246b7b818ecc1358dd3bef1226eb9c09f359665128713a7d988d05"></a>

## Next pages — waf_type / fe1fa8e9662e / 4

- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-c766f377dd26e9aaeab67bb17f27ea07424c43e514f4ac327ce4593d9d5bd14d)
- [waf_type.disable_waf](resources--virtual_host--reference--group-003.md#canonical-201ec42670840ef9105005a05225be6643369dbde7e84c66da6be3acad0c6a53)
- [waf_type.inherit_waf](resources--virtual_host--reference--group-003.md#canonical-74e227e6a93f32320580c00c8459eea86e2dfae357c661b05e8017eb7327fae0)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c766f377dd26e9aaeab67bb17f27ea07424c43e514f4ac327ce4593d9d5bd14d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b623ba3a4aea271462f35a524e39c4d3929cea73f84b402053fc2158b94297e"></a>

## waf_type.app_firewall — waf_type.app_firewall / 0688061a57cf / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- waf_type.app_firewall

<a id="canonical-e728666a02f1d08cf43ef8ff4ff2553c9814007ca124886f3e6c0e38a623cc7f"></a>

Type: `"object"`. single nested block, Optional.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-a9ab3997d91d71dfef600217c589533cb4cd9e7e2adb925cb56156379fda47d3"></a>

## Direct properties — waf_type.app_firewall / 0688061a57cf / 3

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-c274eee577ec3ac14cabc601e844468d275ea973f6da0ed15e652360f6b20763): complete subsection reference.

<a id="canonical-5949fb6d771d85c34b9bfef9bcd90e838e9842ca3673302af0e7d5a43a1e332a"></a>

## Next pages — waf_type.app_firewall / 0688061a57cf / 4

- [waf_type.app_firewall.app_firewall](resources--virtual_host--reference--group-003.md#canonical-c274eee577ec3ac14cabc601e844468d275ea973f6da0ed15e652360f6b20763)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c274eee577ec3ac14cabc601e844468d275ea973f6da0ed15e652360f6b20763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-257ed5b90bb8b1e703c0cd619fc743e7305576cbeb98ce3aaf940a0120c2baaa"></a>

## waf_type.app_firewall.app_firewall — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-c766f377dd26e9aaeab67bb17f27ea07424c43e514f4ac327ce4593d9d5bd14d)
- waf_type.app_firewall.app_firewall

<a id="canonical-9baf36fd3935fcfd8aa4c52631f0f0f546fe3fdbec53120c6915260124879e7f"></a>

Type: `"object"`. list nested block, Optional.

References to an Application Firewall configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2f09fb089f2915d7618f041775caf5d978eb0c5e9385a6d6d76c2a714009396"></a>

## Direct properties — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 3

<a id="canonical-043861fd334adfe6e17dd10bdf049520a281f0381cb12d5d7baeb70091eb1ea0"></a>

<a id="canonical-7d3fad70122dc64e2b96d3c569819d1bc4c28eecf5dc7b9668e27ef3aa5a952e"></a>

## kind property — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 4

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

<a id="canonical-6037f3460bca75d2599b4da3f4e3fb4a39517a848cc8863e73bacfc144ebb6d0"></a>

<a id="canonical-1bbaa009519d1925c8977b50627518e955acd835d315daeab8f43f4f5712f2e0"></a>

## name property — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 5

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

<a id="canonical-8d6d5e7dfee3b786eef87f6566f553e8039b8e8159d98f6d6a9aea39b7e6328e"></a>

<a id="canonical-7a30d6dad1c6d4abdb5c5603d43131dfeec68084433f1235c4b3fabef6e8f6bd"></a>

## namespace property — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 6

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

<a id="canonical-d1eaa2db01df304d0f8ff02803bc1008a6883d1076e222f361d968e21da83f99"></a>

<a id="canonical-506178c960d04149519c58cafa36a0d41bb36fabd5b0a49a9afc2a857857b346"></a>

## tenant property — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 7

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

<a id="canonical-5bce65512d6c3c884de73014da477c3a0c57e72a696fd1fd86628187fe750a42"></a>

<a id="canonical-d37d4e0b2474a7bdaa08db4713bc1cb0f9cbba3674b5f14919bad0d4e227d82e"></a>

## uid property — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 8

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

<a id="canonical-e14be552604d59289798f19673325e3f34b5eda5e2bf0a6cfcbced16b341323a"></a>

## Next pages — waf_type.app_firewall.app_firewall / 22ca3b36ceb2 / 9

- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-c766f377dd26e9aaeab67bb17f27ea07424c43e514f4ac327ce4593d9d5bd14d)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-201ec42670840ef9105005a05225be6643369dbde7e84c66da6be3acad0c6a53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec3e4d84130d3c1348d1e7e84806910347075e29bd4ec23da657eb9631b9e661"></a>

## waf_type.disable_waf — waf_type.disable_waf / 5a8e9b4e06a7 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- waf_type.disable_waf

<a id="canonical-32d2630735fdc0531a934161612f54962c3d67d8c19905addb2c9a7828926d24"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

<a id="canonical-33857bd3a2ef9b8e9937037b1d5cd6ea41648c20d828dd7986b2ca0a0817bf17"></a>

## Direct properties — waf_type.disable_waf / 5a8e9b4e06a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd2ca71a2952ebd58c0df0ed995927590bde624e8027613a0bcf070680abfb99"></a>

## Next pages — waf_type.disable_waf / 5a8e9b4e06a7 / 4

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-74e227e6a93f32320580c00c8459eea86e2dfae357c661b05e8017eb7327fae0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70c43dd8fe9e62963b15fc2d8e92209c79d342673edee52adcadb2922f3ac5cc"></a>

## waf_type.inherit_waf — waf_type.inherit_waf / 7530a2d9da0f / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- waf_type.inherit_waf

<a id="canonical-05758b14baf08eb645c717336bea474329c00789729f0ea82de3b12db9a8873d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

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
inherit_waf = {}
```

<a id="canonical-e20f858fb640d48f112047086fa478434579308cfa144657a76e85029197aef1"></a>

## Direct properties — waf_type.inherit_waf / 7530a2d9da0f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd16e033c80d54732be4c2278edaf1d8c57f4a3f9006ab12e5cfaf9255ff25c9"></a>

## Next pages — waf_type.inherit_waf / 7530a2d9da0f / 4

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
