---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa281a70a238fde8b973d2a04c05244789ac77b9db7260ebad3a287dd5203556"></a>

## Property reference — Property reference / b53bca66dce6 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- Property reference

<a id="canonical-59a895005896b5a5dccc504fe1eee048295c26bc156e5265e4cfcf9444f5abfc"></a>

## Direct properties — Property reference / b53bca66dce6 / 3

<a id="canonical-8fc1fcdea495da5e9e480b53d44276f9b3438095a20d87c50e956eb02b33a2cb"></a>

<a id="canonical-968018456e342511fe5bf3e82e194b9aba8923719cd40de04fe7801cbcf63fce"></a>

## annotations property — Property reference / b53bca66dce6 / 4

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

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-bab48a0ae1745a91d573591e50c3736ddce7a2865341efff0ab28a096a3d6c2c): complete subsection reference.

- [circuit_breaker](resources--cluster--reference--group-001.md#canonical-a922e60dc11e555e59e92a6a368f4c84058d1a3e7b1ac34bacf1c3bdef49e8b7): complete subsection reference.

<a id="canonical-caff0f8947455f51240cca4d0ca306f5af7627ffe5276ca7870242e7d97e0829"></a>

<a id="canonical-96978a23fee267da221be239dbe54cad11be6ec90002da37be360f9df5770a76"></a>

## connection_timeout property — Property reference / b53bca66dce6 / 5

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_subset](resources--cluster--reference--group-001.md#canonical-a37165ac8f46cd776e132350689d8ee540f26d075bd365ec9b1894595191044c): complete subsection reference.

<a id="canonical-ad569a86286331f9d267457bc475b33a1f9ac8b3a871b5bbfe8ccd78b5af8913"></a>

<a id="canonical-8c96440aa445a2cf8a6167beeb243d9da8f6831dabf01877adbfb762198e9ac2"></a>

## description property — Property reference / b53bca66dce6 / 6

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

<a id="canonical-78427797c74f3a93dc8a0acb0eb648841ff7cbbe48a1766eced4cff4eae571fa"></a>

<a id="canonical-19fff202bf1a8675a709dc38917d741d37a86b823355a498978340549fb08114"></a>

## disable property — Property reference / b53bca66dce6 / 7

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

- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-7eb599f330f69eb467d8de9a13f3426f2cbd0fc9d106c199360d68182d532705): complete subsection reference.

<a id="canonical-1edbef0ff9565f846a2dd0b354b39a2a66ef700871593b2b208007af87b22d41"></a>

<a id="canonical-703b43a8faa994164b21b3109b26e3dcde4aef44b62bc24126c2d7f1f45a9a5e"></a>

## endpoint_selection property — Property reference / b53bca66dce6 / 8

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-1c0aa2012e7f4d0cf35d1bac7244000b92a45d48cd22a94ea808079bc16856f3): complete subsection reference.

- [endpoints](resources--cluster--reference--group-001.md#canonical-5a3e03f0a6bd6fb117b6d668c5434b31a5b7be74ba20f1e881e6f1e8450e282b): complete subsection reference.

<a id="canonical-2fb6a670847a603ab65e2ced05868cf045d432a3e84d34286c044261448004ee"></a>

<a id="canonical-6918191cb76f6e4516c8633bb8adcf16100e327a721065de211c89e5bab5b02d"></a>

## fallback_policy property — Property reference / b53bca66dce6 / 9

Type: `"string"`. Optional, Computed.

\[Enum: NO\_FALLBACK|ANY\_ENDPOINT|DEFAULT\_SUBSET\] Enumeration for SubsetFallbackPolicy if subset
match is not found. The request fails as if the cluster had no endpoint matching the subset policy
Any cluster endpoint may be selected if the cluster had no endpoint matching the subset policy Load
balancing is done over endpoints matching.. Possible values are \`NO\_FALLBACK\`, \`ANY\_ENDPOINT\`,
\`DEFAULT\_SUBSET\`. Defaults to \`NO\_FALLBACK\`.

Upstream description:

Enumeration for SubsetFallbackPolicy if subset match is not found.

The request fails as if the cluster had no endpoint matching the subset policy Any cluster endpoint
may be selected if the cluster had no endpoint matching the subset policy Load balancing is done
over endpoints matching default\_subset if the cluster had no endpoint matching the subset policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NO_FALLBACK",
  "enum": [
    "NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [health_checks](resources--cluster--reference--group-001.md#canonical-ff58a11a5ba61e5cfea0689ee64449fce911c8d1afffbddbe44ed167426fb810): complete subsection reference.

- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df): complete subsection reference.

- [http2_options](resources--cluster--reference--group-001.md#canonical-5a4340c3d7f582f647657ffb7085cd625450fe296019861bdd5b6d1c377e0777): complete subsection reference.

<a id="canonical-d4eadb95d3c21b52df9729306ee2feede6a04348e55cd83c5b0af0d898786b99"></a>

<a id="canonical-7cdd05455a6378f75a8e7f650287d1f318773bc1d8be4584b32a411b263246d9"></a>

## http_idle_timeout property — Property reference / b53bca66dce6 / 10

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

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

<a id="canonical-67549938df20d23c8f4ed205d1cb01b7e77ef6bc66be39c327785bfecabedc49"></a>

<a id="canonical-45b247f17bef9170166493e03dfec3297a8e450784d892722ccf3386c79ff1d2"></a>

## id property — Property reference / b53bca66dce6 / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4a8e7e474761b7c38036e06d91257ae8559e17847159159c58be0ab82d5068b2"></a>

<a id="canonical-995b7a4ebff631f456f1ae86cdc82408b400b2912386e43fa3fbb3f68f6008bc"></a>

## labels property — Property reference / b53bca66dce6 / 12

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

<a id="canonical-bce28e47a877f7b25985fade01808367c833b652a52288841a146f808a244fbb"></a>

<a id="canonical-362b6be2c5ef06c6462cd5070c18da0778cae7fde3cb3d26d42f6cdb183ee775"></a>

## loadbalancer_algorithm property — Property reference / b53bca66dce6 / 13

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e1e8a4480c8ffed6c0da9f9eaa33a4d087a240f0c3056230b4c7aa671222078f"></a>

<a id="canonical-36283bb52f75833d452ab61a3a2c67a423f35a8e0f88b38a1e7a68c7e4b95ece"></a>

## max_requests_per_connection property — Property reference / b53bca66dce6 / 14

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\]

Sets the maximum number of requests allowed per connection to the origin server. Enter a value
&gt;=1 to define the request limit per connection.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

OneOf alternatives in this subsection:

- [max_requests_per_connection](resources--cluster--reference--group-001.md#canonical-e1e8a4480c8ffed6c0da9f9eaa33a4d087a240f0c3056230b4c7aa671222078f)
- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-b782d87536617bf6cf2eb48c4214a20810fd35a28dcab745323a8bfb706c5518)

Select alternatives according to the provider validators above.

<a id="canonical-273a0bb0fd52c4ae3200d94d2f48b6acc437465350950aafc268eed5e5c5b533"></a>

<a id="canonical-a5f3c89fc58f050d4ef8a281136f4295db359da9093b86b5d9e55043e93f8967"></a>

## name property — Property reference / b53bca66dce6 / 15

Type: `"string"`. Required.

Name of the Cluster. Must be unique within the namespace.

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

<a id="canonical-13fe36c67e2b6cd56dc33acb57cf4675977cf9278c82dd89248fb832f240c1eb"></a>

<a id="canonical-88dae2006d1b38196d9d7106f9c8e6082c18762def2a7218b34f634bdd7a76bd"></a>

## namespace property — Property reference / b53bca66dce6 / 16

Type: `"string"`. Required.

Namespace where the Cluster is created.

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

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-c3a3695a1c7b081f038e7cd5ed3d1c217bb9cd9239af60a986eb5c4c98feb63a): complete subsection reference.

- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-6a3ec1c603dcd6c9a1200adf9cdb795b80a906d4a7fc21c796f92507d72a5cda): complete subsection reference.

- [outlier_detection](resources--cluster--reference--group-001.md#canonical-7e2b475404eb58b71efe89676a31ca555f4a8f25da60029de259509bf7fbd314): complete subsection reference.

<a id="canonical-6ae632f1a1b7c68a21072b5ff1753fc41f50111305f7e77b342f1326ac9cd4a0"></a>

<a id="canonical-c9ac51b7ff6a2aabb832b144dd600f2b1fa2a648551110017f257216c559768a"></a>

## panic_threshold property — Property reference / b53bca66dce6 / 17

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for loadbalancing ignoring its health status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0faea09edc5f6f80d46593aa359b0a3d9532955253883a1fec198705ea86c6a0): complete subsection reference.

- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-4cb0ace12438d5ea44c00d424f1cf3656df4abb50a9478f84670c01f1f5a5c8a): complete subsection reference.

- [timeouts](resources--cluster--reference--group-001.md#canonical-bc30be729566d8a5ea5e028084a80da67eedaf1c5d039840a9ab5d9c30f60e84): complete subsection reference.

- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522): complete subsection reference.

<a id="canonical-f7262f68bc9ec4adae6aa9003b1329dcfc2693275218bea4ddd6a843976c2cb8"></a>

## All schema paths — Property reference / b53bca66dce6 / 18

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cluster--reference--group-001.md#canonical-8fc1fcdea495da5e9e480b53d44276f9b3438095a20d87c50e956eb02b33a2cb) |
| `auto_http_config` | [auto_http_config](resources--cluster--reference--group-001.md#canonical-fb59821ecefeba08b71f19cb636e8745f1786751a56a8187ad9d71336580e418) |
| `circuit_breaker` | [circuit_breaker](resources--cluster--reference--group-001.md#canonical-31e45ab347b0edddeb48bab54509d12a9e26655550bef6f24c295d02b07a81d2) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](resources--cluster--reference--group-001.md#canonical-cb9cf7622c482b3b29129e3045bd347bb5b16aedbd932ad3eca56f617ae7f0f6) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](resources--cluster--reference--group-001.md#canonical-162dd1e4f5ded0d35d1040e9045d5c7917e96ce3a595544a12f2fc480ef11849) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](resources--cluster--reference--group-001.md#canonical-68a90acb3fffa1a5538af4311b3cef52db981b6b321482058ba8c490720023e3) |
| `circuit_breaker.priority` | [circuit_breaker.priority](resources--cluster--reference--group-001.md#canonical-e9d0a188619fbbd4a0b463472aed96cc9e6e11735699dcb63c2b913ab48b504f) |
| `circuit_breaker.retries` | [circuit_breaker.retries](resources--cluster--reference--group-001.md#canonical-f0a5a0505ccd00ae3e678d8fedf2ed9b32c2b4306701d32b34987472a6def801) |
| `connection_timeout` | [connection_timeout](resources--cluster--reference--group-001.md#canonical-caff0f8947455f51240cca4d0ca306f5af7627ffe5276ca7870242e7d97e0829) |
| `default_subset` | [default_subset](resources--cluster--reference--group-001.md#canonical-fc94230e0b77bd94d6298105fc701b97e5907042a847f36ae65f320aa64171c3) |
| `description` | [description](resources--cluster--reference--group-001.md#canonical-ad569a86286331f9d267457bc475b33a1f9ac8b3a871b5bbfe8ccd78b5af8913) |
| `disable` | [disable](resources--cluster--reference--group-001.md#canonical-78427797c74f3a93dc8a0acb0eb648841ff7cbbe48a1766eced4cff4eae571fa) |
| `disable_proxy_protocol` | [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-96686c9aa56a3e0cfb61d93226724d10e317488e446ba083978b693b93d7f0bb) |
| `endpoint_selection` | [endpoint_selection](resources--cluster--reference--group-001.md#canonical-1edbef0ff9565f846a2dd0b354b39a2a66ef700871593b2b208007af87b22d41) |
| `endpoint_subsets` | [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-2a4ec87ac8fd21dd0dd0905113eca80aa1b3466c8f576992e7932eba2ee2da2c) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](resources--cluster--reference--group-001.md#canonical-2e55675e3e743d67fdb1389b067688a9fbca64163b2f30a2530f9771e834d82f) |
| `endpoints` | [endpoints](resources--cluster--reference--group-001.md#canonical-b7c767fa5117f128ee884ded7613a1aa0c98f2648865e087b6d48ceecb39ec43) |
| `endpoints.kind` | [endpoints.kind](resources--cluster--reference--group-001.md#canonical-6b071c75dd03e16a3091188bbe2a1b6cdb1efe05382e2da39fdef907672f66b7) |
| `endpoints.name` | [endpoints.name](resources--cluster--reference--group-001.md#canonical-f17e004f477fa6535f6407576f712b9ee9b81da593501e4c6560d408c20724e5) |
| `endpoints.namespace` | [endpoints.namespace](resources--cluster--reference--group-001.md#canonical-d2e93f3698a53d35c09d0846179d5e1e3fd76c7c64e789e2bbf3751e63abd188) |
| `endpoints.tenant` | [endpoints.tenant](resources--cluster--reference--group-001.md#canonical-6d1703e21c80b82a42ea124a1686b855d93e6805b4421b19459a0813e37d5a96) |
| `endpoints.uid` | [endpoints.uid](resources--cluster--reference--group-001.md#canonical-ca0431d20d21570b8ac03804343d3e36b48d210c7a3f0fc8b1d80ea2f1f05ce8) |
| `fallback_policy` | [fallback_policy](resources--cluster--reference--group-001.md#canonical-2fb6a670847a603ab65e2ced05868cf045d432a3e84d34286c044261448004ee) |
| `health_checks` | [health_checks](resources--cluster--reference--group-001.md#canonical-d5eb8413a271297ff1c1dc01af0acaae90f350c8f2fc1d3e68a272f2af058fe2) |
| `health_checks.kind` | [health_checks.kind](resources--cluster--reference--group-001.md#canonical-d0d8a5c5871b983eac300a44822e8c4a06ad128165844e798f9fdaf2def8d403) |
| `health_checks.name` | [health_checks.name](resources--cluster--reference--group-001.md#canonical-b531ae2f5332f58b7607e492b550bc4246d6df5daeb42b4d197463bfa23a5c21) |
| `health_checks.namespace` | [health_checks.namespace](resources--cluster--reference--group-001.md#canonical-16c1f7377f3b73532ce9c5bac999dc8d34eda5fad57bba16439c7d43979e0fbd) |
| `health_checks.tenant` | [health_checks.tenant](resources--cluster--reference--group-001.md#canonical-97a4ddc50ec1a36f83b9e7fb4d66991f8a555c4daff94eec639df1cfe18c9cc8) |
| `health_checks.uid` | [health_checks.uid](resources--cluster--reference--group-001.md#canonical-be5746370065e17a6cc89b76cebd3c8747d276f2b0c3c38498cde1dbaa59f7be) |
| `http1_config` | [http1_config](resources--cluster--reference--group-001.md#canonical-b68310e03edecb418fe71b6fd903274d09acb3fd0b03eb85e4378141f90cd3e0) |
| `http1_config.header_transformation` | [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-056167a11d5b4b7207318c7e3296ed8a0c0fd65102e18449d43acdfb8c185624) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](resources--cluster--reference--group-001.md#canonical-fd5705c372bd8edcbec65e294ae4ac094f7def3fd27fc05c348ee4a43f5736cf) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-0ea1363f6b1dee97dfdd1558afb95cdd38d84b1fa26512c1da515d0b5966d8dd) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-7c5428cf322cfe37b945aa1fc88be9efca4608723400793d09aee6d5969b01ed) |
| `http2_options` | [http2_options](resources--cluster--reference--group-001.md#canonical-91abadd82ba6567818ad5263633ae205e99430f96f8588f4af74c0e2015defb4) |
| `http2_options.enabled` | [http2_options.enabled](resources--cluster--reference--group-001.md#canonical-3c2e14c6515b2318c9418a7da04cbc34f71d50db079ca3a648efc7c23145d8b7) |
| `http_idle_timeout` | [http_idle_timeout](resources--cluster--reference--group-001.md#canonical-d4eadb95d3c21b52df9729306ee2feede6a04348e55cd83c5b0af0d898786b99) |
| `id` | [id](resources--cluster--reference--group-001.md#canonical-67549938df20d23c8f4ed205d1cb01b7e77ef6bc66be39c327785bfecabedc49) |
| `labels` | [labels](resources--cluster--reference--group-001.md#canonical-4a8e7e474761b7c38036e06d91257ae8559e17847159159c58be0ab82d5068b2) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](resources--cluster--reference--group-001.md#canonical-bce28e47a877f7b25985fade01808367c833b652a52288841a146f808a244fbb) |
| `max_requests_per_connection` | [max_requests_per_connection](resources--cluster--reference--group-001.md#canonical-e1e8a4480c8ffed6c0da9f9eaa33a4d087a240f0c3056230b4c7aa671222078f) |
| `name` | [name](resources--cluster--reference--group-001.md#canonical-273a0bb0fd52c4ae3200d94d2f48b6acc437465350950aafc268eed5e5c5b533) |
| `namespace` | [namespace](resources--cluster--reference--group-001.md#canonical-13fe36c67e2b6cd56dc33acb57cf4675977cf9278c82dd89248fb832f240c1eb) |
| `no_panic_threshold` | [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-0c5ae0df7005c08e4764e302ed0d8ec65685adb59eab3170438ed49002d774e9) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-b782d87536617bf6cf2eb48c4214a20810fd35a28dcab745323a8bfb706c5518) |
| `outlier_detection` | [outlier_detection](resources--cluster--reference--group-001.md#canonical-80c771b6adb3d962e905e91ac1ff6291de88b10e015b37de851a3ef372d45d33) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](resources--cluster--reference--group-001.md#canonical-607b84026e14ce33ebfda8f2dfc6e331a7e19ff07584d430f7848c74b97b3782) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](resources--cluster--reference--group-001.md#canonical-c6c64f871cb9f46bb6e2baaa429e5aaf4b4c53bc286ef7eda44514311312667b) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](resources--cluster--reference--group-001.md#canonical-073eb12937573aa8f5d6ec0770f62d14c104bd0d146e7bd6841d17644d0751f2) |
| `outlier_detection.interval` | [outlier_detection.interval](resources--cluster--reference--group-001.md#canonical-c9d3004d7d0fb237cb239e82d3cf637e970dca80b0d353421188d19e68c2d531) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](resources--cluster--reference--group-001.md#canonical-7cd129b1dc4618dc46a0807f663ad584374972dacf9ce000e0af478cfc1d72f4) |
| `panic_threshold` | [panic_threshold](resources--cluster--reference--group-001.md#canonical-6ae632f1a1b7c68a21072b5ff1753fc41f50111305f7e77b342f1326ac9cd4a0) |
| `proxy_protocol_v1` | [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-1b2bd886a3e494ea5ed6bfb99fef6bd969ce630197665b939d5d03accdc9c161) |
| `proxy_protocol_v2` | [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-d236d9064833a6e6d9b63c6b247c73347808661797bf8fa18261a01e9f9568ff) |
| `timeouts` | [timeouts](resources--cluster--reference--group-001.md#canonical-056060cdd509edf92db81cbb6c10d9fd0d52fa118432c1eb62cd5e21a7167f58) |
| `timeouts.create` | [timeouts.create](resources--cluster--reference--group-001.md#canonical-5ced1feb443a7e7f5fc2e5e3dacff54054253e899955ae95da9f7d684e143cee) |
| `timeouts.delete` | [timeouts.delete](resources--cluster--reference--group-001.md#canonical-c6591149383f3dbd5f34297409873d2a78f93f060d2fca9912d4a77ec6ce9472) |
| `timeouts.read` | [timeouts.read](resources--cluster--reference--group-001.md#canonical-9daa3efe5d66a34af899a7950b09ce296205dfe776d6c455fa2a0b420ff2c3c0) |
| `timeouts.update` | [timeouts.update](resources--cluster--reference--group-001.md#canonical-e03f31c6597718922718999d58a3215e7ff3102d96551ab4c60e9cce644fabfc) |
| `tls_parameters` | [tls_parameters](resources--cluster--reference--group-001.md#canonical-dd8a87d2474107719e6379d039e437638cf0d967563ce1f71351d27a58a3c998) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-b9998fdc13a0a5f04aa41ef2f7bae90630f498a9d562833dd2a0b0035068e04d) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](resources--cluster--reference--group-001.md#canonical-addfdb5e04a31cb8510b45dbfb6991e8912bd3c6fdf08f2fa2ae266fec96a6f8) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](resources--cluster--reference--group-001.md#canonical-fe3ed84b554f1463f9fa01ac3323fbc59694824d57f11422bed145b446a585c6) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](resources--cluster--reference--group-001.md#canonical-7f157af08efecd392add8982040ebfc33ed57dcc9fc457c2dcaf92653146c57a) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](resources--cluster--reference--group-001.md#canonical-e70485500f55d0440fdc71fb625c88bab065201b73d701bbcd7ad7d55ec564d3) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](resources--cluster--reference--group-001.md#canonical-a0bd99152dc53c229da2ff7fd0c10c861f0189c574308d74eb60b673373bf007) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](resources--cluster--reference--group-001.md#canonical-e9526204eaf32121a6ba8a7857f08be21dc19c6121eb84c0a437071c21c2e4d2) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](resources--cluster--reference--group-001.md#canonical-7a608f048ad7ab2b8b246a27b581a648d8f3746dd623e583383f75b7b3220689) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](resources--cluster--reference--group-001.md#canonical-4507b54c6416365dd3f9c583578c242a54fb6c3b4cef87dd2e2abedd0550e45f) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](resources--cluster--reference--group-001.md#canonical-af9ae12f906508dd7e7391441a1d95e42c028c5b00ea7915d228b8499500e35d) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](resources--cluster--reference--group-001.md#canonical-b4e18753ce6a35fe86e27f747b5153df3793919d726a5c92fab3e4778abc1ccb) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-55a6a6ca75db984cfcfa83156007bbccec48479557a7888cde4e1274608ebdea) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](resources--cluster--reference--group-001.md#canonical-19f6cc099d61755427cb138854779f6f41ff6c7776ead57f0a733e87c94b512b) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-99cbd2465519e92554e8b0173bca484558991933c34bedfedb82067aebd5c8e9) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-001.md#canonical-e038d8bdf3d308533a63be452f07751999fbd9578cb92694fbdd302852360126) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](resources--cluster--reference--group-001.md#canonical-b28bc2125656e380d023cf07f0b4e98e9ab29c22a96129d56b152c5de06a1477) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](resources--cluster--reference--group-001.md#canonical-d658b4abb22243d4cdb6cf5ec48145002ae88d7496dfd8811ef98209ddb3ef4c) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](resources--cluster--reference--group-001.md#canonical-83aaadc188014eb638c199fe9d9d56de15962f432c533e24e04ec3fa6b89ad3a) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](resources--cluster--reference--group-001.md#canonical-66fefdbed0c9d84ad65f69414cbb581262a98892ba892df1621ca82e050a707f) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](resources--cluster--reference--group-001.md#canonical-526fa0ba78e9014c04e1ebf0ed02ec6e973913c48a18a5c9ac73818b208b155c) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](resources--cluster--reference--group-001.md#canonical-036e3bdb93960f7dea43753a63cde4b65e00e4052efe62bec44334603f2b3979) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](resources--cluster--reference--group-001.md#canonical-c98d20c59e86848bd3918cc22c9966ea0646279a20708f29a90834bb975247e9) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-37ae27aa12d3c9435aab0eae126cf6010fdf68e43dbdf6070f5581ff809e19ae) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-1f277e092d4675646bf7274ea5256c891590f9a536ed3baa1c42c2cf524cd24d) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--cluster--reference--group-001.md#canonical-5f95547a893f9d83522b5fb05beaf2d03b884c0c9397ee42d15ef6f15dde05aa) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--cluster--reference--group-001.md#canonical-38b6cab40d0433ec6128b2119fd8195c6d3c5bdeca040c0dee25529f233b8267) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--cluster--reference--group-001.md#canonical-b4d1f704fb77a685cb192e40f3a244e798d9d70b2038a2252a60a9478b29b75a) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-10435dbffdb135aeef9f3ad66f251def795099d76ea3b9259e708b9f0b5f4c9e) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--cluster--reference--group-001.md#canonical-d1d448e38bcd862fa131035b28fc7baa0cd407aba326df88a12fe695c2ce3486) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-d137a891e686c4cfa26f512991ae1e88f10aeae5db0efd0453017ff58d5a779b) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--cluster--reference--group-001.md#canonical-6cd45ba48734d0a007b4b4b3dcbe3bedf108cb71e06560f569d6718df4524f07) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--cluster--reference--group-001.md#canonical-4a2c005adc93db0f5a82ee57ae9828a40da60ca2605fcdb15243aeabc9c4ee0d) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-42c450dcb16715afee2f4ccaa2e4a84b78a417def91b50f65e4f86d21d897db6) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-a0eb9ffb350218a4ce7eb459feb35176c778aa8f7eaa75a10ce9ad820a4c74c2) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-f5922691f2a61281fdee3596800c37b1f47c700926abd368033e9f9cd6185b45) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--cluster--reference--group-001.md#canonical-5fec8c3b96656f1cc2044ed285cd8837df3d27caf852085548f388695251392c) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--cluster--reference--group-001.md#canonical-dc957c9510dc06a5bd1357d077eebcdbd65669e9f8aebb0b5d9e1b14303fadf6) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--cluster--reference--group-001.md#canonical-1d0ba3ae987e39cddf0f2eeb9f2f37a024a521327eeff7b5a190f7c12f290a77) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--reference--group-001.md#canonical-3bb3c7d0bb9cdf37c9616b285ea1a8010ce3830faaa5560f7df7b31fe9cff1eb) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--cluster--reference--group-001.md#canonical-f393ffb4e24c0c8456621e693e683dec58a9baba84b3201fdef5a69b102a65f3) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--cluster--reference--group-001.md#canonical-4bc7f207f1b3472a4d49aa6fa10d0784a8d4502dcc047feba4375c29098e5eb1) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--cluster--reference--group-001.md#canonical-180f1d45f5c145e8f1b7af1abd357916daacc391c4d3b36dfe9cd1cb536b17b8) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--cluster--reference--group-001.md#canonical-1f7d2d917efbfa082ebb6ba809a3ba490978c528eaa80b39e160363679fffa50) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--cluster--reference--group-001.md#canonical-a50f6ae38cc2aba5ead3c14d78f8bcf2127708cf2ed03e6c3594fc03cc7ffc14) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-170719cf37f958c5e9afa9de0a731719b8f06e83d203b46282fa1023b230bbd1) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-002.md#canonical-776a519ebb4c104526135872694f216de76fef866b5faef2133c4fecb549e86f) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--cluster--reference--group-002.md#canonical-c64f97c45f7f33804dd32009753dea14692f433582bb4601f82fbddc91e8b2ae) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--cluster--reference--group-002.md#canonical-91fdb9da0b3d68ab472bc939127f2db621af3d9bc63a8e682e1049f4b28e4c04) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--cluster--reference--group-002.md#canonical-404eeb2124262c10cc48f2073d86ab7ab56b6cfde2297c7c0561d8748e2d1974) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--cluster--reference--group-002.md#canonical-1a4beef95ca51a95e6d8fdaa82cb9f18d1843c97cfb82d9478eb83361a936527) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--cluster--reference--group-002.md#canonical-f37dca176e42eb6bcbb30d7ada0b4d8c74bb67e30d088771ea8bbbf43fd9419a) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--cluster--reference--group-001.md#canonical-5fe454fe1cd63ad24fcd52bbc90441efe88964153b70dc9d6f96af3ae6645ea6) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--cluster--reference--group-001.md#canonical-94d256d3924032c843d8cfc827d45f82179ea7c1b0b316fdbe08cc3b52cdceb0) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](resources--cluster--reference--group-002.md#canonical-332a750670cdb85df34e3b3b0a6f7e1d2490570aebb0a1e175e7bf2f449b66f2) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-b3c4b2e7f5fd767142ae13637e73d0cf3a5ddc7fc745da2eef570b922f6b6245) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](resources--cluster--reference--group-002.md#canonical-c5351cb89ee4e58ee13243ba03bd70661fb273440ddc816f0103f43283b09e27) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](resources--cluster--reference--group-001.md#canonical-624ba90a9f547b00a7825255798b78b590da24a62db543fb160ea5edf5fdd1fd) |
| `tls_parameters.sni` | [tls_parameters.sni](resources--cluster--reference--group-001.md#canonical-4ebecbeb32da21925b86b69e08390ba8a2ef960458baba4c926f6012385b7ca9) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-869a973133fb15ecd6f9aa17a836a35914a65f15d528f0ae86823e6f84734bfb) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9bfa6c40a0109eebbc9d72592925401c0144c4f71f455e533ac7d3d7190009a8) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-37eef0fa3226f945d5b34dacd3c189172682ff9112adbe283ddc80b8d7bf353c) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-99e4afe2af583441efae158d61e4138c88183f8ff71309dc9b02fff1370fc0df) |

<a id="canonical-04f60f3d16de927a4311c9c40702d023a77f4b2f9430696b714824ac07e80ec9"></a>

## Next pages — Property reference / b53bca66dce6 / 19

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-bab48a0ae1745a91d573591e50c3736ddce7a2865341efff0ab28a096a3d6c2c)
- [circuit_breaker](resources--cluster--reference--group-001.md#canonical-a922e60dc11e555e59e92a6a368f4c84058d1a3e7b1ac34bacf1c3bdef49e8b7)
- [default_subset](resources--cluster--reference--group-001.md#canonical-a37165ac8f46cd776e132350689d8ee540f26d075bd365ec9b1894595191044c)
- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-7eb599f330f69eb467d8de9a13f3426f2cbd0fc9d106c199360d68182d532705)
- [endpoint_subsets](resources--cluster--reference--group-001.md#canonical-1c0aa2012e7f4d0cf35d1bac7244000b92a45d48cd22a94ea808079bc16856f3)
- [endpoints](resources--cluster--reference--group-001.md#canonical-5a3e03f0a6bd6fb117b6d668c5434b31a5b7be74ba20f1e881e6f1e8450e282b)
- [health_checks](resources--cluster--reference--group-001.md#canonical-ff58a11a5ba61e5cfea0689ee64449fce911c8d1afffbddbe44ed167426fb810)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- [http2_options](resources--cluster--reference--group-001.md#canonical-5a4340c3d7f582f647657ffb7085cd625450fe296019861bdd5b6d1c377e0777)
- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-c3a3695a1c7b081f038e7cd5ed3d1c217bb9cd9239af60a986eb5c4c98feb63a)
- [no_request_limit_per_connection](resources--cluster--reference--group-001.md#canonical-6a3ec1c603dcd6c9a1200adf9cdb795b80a906d4a7fc21c796f92507d72a5cda)
- [outlier_detection](resources--cluster--reference--group-001.md#canonical-7e2b475404eb58b71efe89676a31ca555f4a8f25da60029de259509bf7fbd314)
- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-0faea09edc5f6f80d46593aa359b0a3d9532955253883a1fec198705ea86c6a0)
- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-4cb0ace12438d5ea44c00d424f1cf3656df4abb50a9478f84670c01f1f5a5c8a)
- [timeouts](resources--cluster--reference--group-001.md#canonical-bc30be729566d8a5ea5e028084a80da67eedaf1c5d039840a9ab5d9c30f60e84)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-9670852cd96a827e1246fe6fc573c2dedb956c77b7b1717fa85237418a3d9522)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-bab48a0ae1745a91d573591e50c3736ddce7a2865341efff0ab28a096a3d6c2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b4379a852d4e20b283ef2dd604c152541e027d73c90e2ccb40b9f4ba4ba6926"></a>

## auto_http_config — auto_http_config / 9219c75fdd03 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- auto_http_config

<a id="canonical-fb59821ecefeba08b71f19cb636e8745f1786751a56a8187ad9d71336580e418"></a>

Type: `["object", {}]`. Optional.

\[OneOf: auto\_http\_config, http1\_config, http2\_options\] Enable this option

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

- [auto_http_config](resources--cluster--reference--group-001.md#canonical-fb59821ecefeba08b71f19cb636e8745f1786751a56a8187ad9d71336580e418)
- [http1_config](resources--cluster--reference--group-001.md#canonical-b68310e03edecb418fe71b6fd903274d09acb3fd0b03eb85e4378141f90cd3e0)
- [http2_options](resources--cluster--reference--group-001.md#canonical-91abadd82ba6567818ad5263633ae205e99430f96f8588f4af74c0e2015defb4)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
auto_http_config = {}
```

<a id="canonical-1bca992ac90797bd730de1d8527b0981d78a61f8e9c576d7a3e997502a4479da"></a>

## Direct properties — auto_http_config / 9219c75fdd03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37874f203eb91c93b31d8013c8befacf23876d312a92a0723d01fa9a6751b74b"></a>

## Next pages — auto_http_config / 9219c75fdd03 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-a922e60dc11e555e59e92a6a368f4c84058d1a3e7b1ac34bacf1c3bdef49e8b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d65299f6334fb9455e07db7ecb7a8785116e168e6c8f1a38774ec28f66408dc1"></a>

## circuit_breaker — circuit_breaker / 604693530529 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- circuit_breaker

<a id="canonical-31e45ab347b0edddeb48bab54509d12a9e26655550bef6f24c295d02b07a81d2"></a>

Type: `"object"`. single nested block, Optional.

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Upstream description:

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

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
circuit_breaker {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa1dce9d569c654441ae7e193795ada02a78854975f14031ab313de3d18a0284"></a>

## Direct properties — circuit_breaker / 604693530529 / 3

<a id="canonical-cb9cf7622c482b3b29129e3045bd347bb5b16aedbd932ad3eca56f617ae7f0f6"></a>

<a id="canonical-712161cbcf6fcb0383cf3bf1204e722efc4061f3d33551c3192d7886cc67f7f2"></a>

## connection_limit property — circuit_breaker / 604693530529 / 4

Type: `"number"`. Optional.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-162dd1e4f5ded0d35d1040e9045d5c7917e96ce3a595544a12f2fc480ef11849"></a>

<a id="canonical-1e5a59c2667b5927683c1a73174a1644c0ed03b03073ea2a0c57ac4df40a841d"></a>

## max_requests property — circuit_breaker / 604693530529 / 5

Type: `"number"`. Optional.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-68a90acb3fffa1a5538af4311b3cef52db981b6b321482058ba8c490720023e3"></a>

<a id="canonical-93774665bac1ef45960da3725dd456216203c5ccb348c76bbd62ce772f9aaea7"></a>

## pending_requests property — circuit_breaker / 604693530529 / 6

Type: `"number"`. Optional.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately..

Upstream description:

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-e9d0a188619fbbd4a0b463472aed96cc9e6e11735699dcb63c2b913ab48b504f"></a>

<a id="canonical-142b20373cd456c9fca4edb75b358f0b8f438712d21bfcf517e0d35828549877"></a>

## priority property — circuit_breaker / 604693530529 / 7

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f0a5a0505ccd00ae3e678d8fedf2ed9b32c2b4306701d32b34987472a6def801"></a>

<a id="canonical-2ecf3b15948ba8c78ccc2b7b443124e24c827f63f73cb24ea3210053ea84de3f"></a>

## retries property — circuit_breaker / 604693530529 / 8

Type: `"number"`. Optional.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0e44b41927fdfc7cea187301f35d7234fcc433a24f4c92751720136860595fc8"></a>

## Next pages — circuit_breaker / 604693530529 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-a37165ac8f46cd776e132350689d8ee540f26d075bd365ec9b1894595191044c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd7d7219d1f131c98e7e8d96339395b745fe292df2ae0cde7080295132176a07"></a>

## default_subset — default_subset / 885fa7a3919b / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- default_subset

<a id="canonical-fc94230e0b77bd94d6298105fc701b97e5907042a847f36ae65f320aa64171c3"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

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
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

Terraform syntax:

```terraform
default_subset {}
```

<a id="canonical-e8c44bcbf4389f122935bc4ce940418d2a495b8a0f16caa0d410fa77980eef23"></a>

## Direct properties — default_subset / 885fa7a3919b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9ecf41e9c81803af7ac354d5fbb632ce26f3dec732ef2c5d49d1833256f50d1"></a>

## Next pages — default_subset / 885fa7a3919b / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-7eb599f330f69eb467d8de9a13f3426f2cbd0fc9d106c199360d68182d532705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ef188920bec8ca60b1319bbde2fbb5aecfa14b47efba2150309aa695d06c781"></a>

## disable_proxy_protocol — disable_proxy_protocol / a4190303e17d / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- disable_proxy_protocol

<a id="canonical-96686c9aa56a3e0cfb61d93226724d10e317488e446ba083978b693b93d7f0bb"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

- [disable_proxy_protocol](resources--cluster--reference--group-001.md#canonical-96686c9aa56a3e0cfb61d93226724d10e317488e446ba083978b693b93d7f0bb)
- [proxy_protocol_v1](resources--cluster--reference--group-001.md#canonical-1b2bd886a3e494ea5ed6bfb99fef6bd969ce630197665b939d5d03accdc9c161)
- [proxy_protocol_v2](resources--cluster--reference--group-001.md#canonical-d236d9064833a6e6d9b63c6b247c73347808661797bf8fa18261a01e9f9568ff)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_proxy_protocol = {}
```

<a id="canonical-8d2341ab8bb714cf1e1c10b1c2355bde8bb8f375155cacd9c785d8adc40ecd67"></a>

## Direct properties — disable_proxy_protocol / a4190303e17d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-101f4675714aa4eabfe349e8ee270150be5c1022511f59cb65df92d8523b52ca"></a>

## Next pages — disable_proxy_protocol / a4190303e17d / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-1c0aa2012e7f4d0cf35d1bac7244000b92a45d48cd22a94ea808079bc16856f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f61b22684b93cb8a81875329a83b9284dd24d0cb4f92e7645d13acdef0462dc"></a>

## endpoint_subsets — endpoint_subsets / 05754767d99f / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- endpoint_subsets

<a id="canonical-2a4ec87ac8fd21dd0dd0905113eca80aa1b3466c8f576992e7932eba2ee2da2c"></a>

Type: `"object"`. list nested block, Optional.

Configure endpoint groups based on metadata labels for traffic routing. Supports weighted
distribution and session affinity across labeled endpoints.

Upstream description:

Cluster may be configured to divide its endpoints into subsets based on metadata attached to the
endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected
by the load balancer.

Endpoint\_subsets is list of subsets for this cluster. Each entry in this list has definition for a
subset (which is collection of keys)

During routing, the route’s metadata match configuration is used to find a specific subset. If there
is a subset with the exact keys and values specified by the route, the subset is used for load
balancing. Otherwise, the fallback policy is used. The cluster’s subset configuration must,
therefore, contain a definition that has the same keys as a given route in order for subset load
balancing to occur.

Example:

RouteConfig

routes: &#8203;- match: &#8203;- headers: \[\] path: path: /1.log query\_params: \[\]
routeDestination: destinations: &#8203;- cluster: &#8203;- kind: cluster.object uid:
00000000-0000-4000-8000-0b50b89d07a2 endpointSubsets: site: india

EndpointConfig

metadata: labels: deployment: debug site: india name: end-1 uid: end-1

ClusterConfig

gcSpec: defaultSubset: stage: production fallbackPolicy: DEFAULT\_SUBSET endpointSubsets: &#8203;-
keys: &#8203;- site &#8203;- keys: &#8203;- stage &#8203;- app

Assume the below endpoints are defined and associated with the cluster.

Endpoint Labels -------- ------

ep1 stage: production, site: india ep2 stage: deployment, site: us ep3 stage: production, app: hr
ep4 site: india

The following table describes some routes and the result of their application to the cluster. The
subset definition for cluster is assumed to be same as given above in the ClusterConfig section

RouteMatch Criteria Subset Reason ------------------- ------ ------

site: india ep1, ep4 Subset of endpoints selected site: us ep2 Subset of endpoints selected app: hr
ep1, ep3 Fallback: No subset selector for "app" alone stage: production, app: hr ep3 Subset of
endpoints selected other: x ep1, ep3 Fallback: No subset selector for “other” (none) ep1, ep3
Fallback: No subset requested.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("keys")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-decf46596638721eda8a0e575414fa69d76d23ea26a80fcba2a07bc164d7fdc6"></a>

## Direct properties — endpoint_subsets / 05754767d99f / 3

<a id="canonical-2e55675e3e743d67fdb1389b067688a9fbca64163b2f30a2530f9771e834d82f"></a>

<a id="canonical-496fac3b74400c3de6aff44c057ef5dfb8ef2dd5480f8dc1a24e534cf89b3913"></a>

## keys property — endpoint_subsets / 05754767d99f / 4

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3c7cf5586c5ce3eb63a1ec9889b9386ef5335b61efd8b4953e19e44970d28355"></a>

## Next pages — endpoint_subsets / 05754767d99f / 5

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-5a3e03f0a6bd6fb117b6d668c5434b31a5b7be74ba20f1e881e6f1e8450e282b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13cd8e298f7b6ecff6489fe8362a7bb1537bbfd466c81578414e018d734344c1"></a>

## endpoints — endpoints / 21b51ba14c38 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- endpoints

<a id="canonical-b7c767fa5117f128ee884ded7613a1aa0c98f2648865e087b6d48ceecb39ec43"></a>

Type: `"object"`. list nested block, Optional.

List of endpoints for this cluster.

Upstream description:

List of references to all endpoint objects that belong to this cluster.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8418b927689e036c738d03c7eb0171827d8564635ecf12f7f8ce437acafafbd"></a>

## Direct properties — endpoints / 21b51ba14c38 / 3

<a id="canonical-6b071c75dd03e16a3091188bbe2a1b6cdb1efe05382e2da39fdef907672f66b7"></a>

<a id="canonical-3fe99366ada164e8ff0dfd778d8c6a39ef67c8978e73a54f3807dad66908f464"></a>

## kind property — endpoints / 21b51ba14c38 / 4

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

<a id="canonical-f17e004f477fa6535f6407576f712b9ee9b81da593501e4c6560d408c20724e5"></a>

<a id="canonical-6c7c2f9bfad59da86a0730df7274f09526f1d04a32690fd6bddca1fd9ccd1506"></a>

## name property — endpoints / 21b51ba14c38 / 5

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

<a id="canonical-d2e93f3698a53d35c09d0846179d5e1e3fd76c7c64e789e2bbf3751e63abd188"></a>

<a id="canonical-87fbe77a10233dc39b87981b085eb6708ceaebf08ca18668ee8e255c51b5d9c8"></a>

## namespace property — endpoints / 21b51ba14c38 / 6

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

<a id="canonical-6d1703e21c80b82a42ea124a1686b855d93e6805b4421b19459a0813e37d5a96"></a>

<a id="canonical-e87df3836f9d60dbe3ad6b72ac30af480453d5459c9638897813ab996157d4fd"></a>

## tenant property — endpoints / 21b51ba14c38 / 7

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

<a id="canonical-ca0431d20d21570b8ac03804343d3e36b48d210c7a3f0fc8b1d80ea2f1f05ce8"></a>

<a id="canonical-53169cb76bab73402ee6443e518c950def40e0766b1925255d6cbc4a6e269cc5"></a>

## uid property — endpoints / 21b51ba14c38 / 8

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

<a id="canonical-4581cb9e0e966f6890a0337753c0e6faa29c6699d96535dca68afaf369d23402"></a>

## Next pages — endpoints / 21b51ba14c38 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-ff58a11a5ba61e5cfea0689ee64449fce911c8d1afffbddbe44ed167426fb810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dae9e56fcd8255ce1e09bcf035701700f21e4e64de11db57d606cb137f3b1407"></a>

## health_checks — health_checks / e1a924fa941a / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- health_checks

<a id="canonical-d5eb8413a271297ff1c1dc01af0acaae90f350c8f2fc1d3e68a272f2af058fe2"></a>

Type: `"object"`. list nested block, Optional.

Health check configuration for backend monitoring.

Upstream description:

List of references to healthcheck object for this cluster.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-a472639d3a5598c5bd9f229930189559bbc362885fe2f4de11908c4006b322c3"></a>

## Direct properties — health_checks / e1a924fa941a / 3

<a id="canonical-d0d8a5c5871b983eac300a44822e8c4a06ad128165844e798f9fdaf2def8d403"></a>

<a id="canonical-e75f2ab81249d18cb0036edfe357ee6944e81b913ae5fdf9c4fbdc1882248325"></a>

## kind property — health_checks / e1a924fa941a / 4

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

<a id="canonical-b531ae2f5332f58b7607e492b550bc4246d6df5daeb42b4d197463bfa23a5c21"></a>

<a id="canonical-97f545e6a7c06dad5ea23a9d83501c54d6344659ee79384c272e9fba60343fec"></a>

## name property — health_checks / e1a924fa941a / 5

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

<a id="canonical-16c1f7377f3b73532ce9c5bac999dc8d34eda5fad57bba16439c7d43979e0fbd"></a>

<a id="canonical-6ce7183c3f8bec010d24f4024aac02fc79b76cf4cd7f03d33b00b6848f25ff23"></a>

## namespace property — health_checks / e1a924fa941a / 6

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

<a id="canonical-97a4ddc50ec1a36f83b9e7fb4d66991f8a555c4daff94eec639df1cfe18c9cc8"></a>

<a id="canonical-584608dcc92379eef1834f84f0e57ce129b0ae0a98d8d59a428ba525117075ac"></a>

## tenant property — health_checks / e1a924fa941a / 7

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

<a id="canonical-be5746370065e17a6cc89b76cebd3c8747d276f2b0c3c38498cde1dbaa59f7be"></a>

<a id="canonical-e44ad6bec451271e41bc7720847bd225b25b6502055380d90b08bf16661d6f4b"></a>

## uid property — health_checks / e1a924fa941a / 8

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

<a id="canonical-c9ae71c474df41e23af1b5ede40d9904edb93563cb6ec8840f41e88b73a6f663"></a>

## Next pages — health_checks / e1a924fa941a / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42d5ac09d706fbeeb2bc28986849ead9edf310ca77e82d4a860112a54ccfe320"></a>

## http1_config — http1_config / f25e9c495402 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- http1_config

<a id="canonical-b68310e03edecb418fe71b6fd903274d09acb3fd0b03eb85e4378141f90cd3e0"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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
http1_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3324020c473ab409f937cba2884ca531e4531d7d8a84bacc6442575c550edd1e"></a>

## Direct properties — http1_config / f25e9c495402 / 3

- [header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285): complete subsection reference.

<a id="canonical-e82a996ece1049a732458bc3f046d41f41e5675a972b0511580054184f13affe"></a>

## Next pages — http1_config / f25e9c495402 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34cffd914a153f53b1ee26fb0671119d48b2d258a3837146cdab60ac11ba2836"></a>

## http1_config.header_transformation — http1_config.header_transformation / 3fddf4f4d565 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- http1_config.header_transformation

<a id="canonical-056167a11d5b4b7207318c7e3296ed8a0c0fd65102e18449d43acdfb8c185624"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fbc9e6870181c085573593393ad59a0db22013e825acd80823b547dfbc49c95"></a>

## Direct properties — http1_config.header_transformation / 3fddf4f4d565 / 3

- [default_header_transformation](resources--cluster--reference--group-001.md#canonical-e456117ff3bd11c168cb1deb82fafaca3de6673fb5d4bfd3f027b27a3beb5261): complete subsection reference.

- [preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-af30163a20dd7ea0773998fcbae339ea6768cec4473eb5283d0192f562f3139d): complete subsection reference.

- [proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-ad439f2364891cb0730ea39936dc5fc9e520916a9a94ff608247eb9a622094d8): complete subsection reference.

<a id="canonical-0105c95f666ae18cbf157ecbb11da2fe67a00c19a22d730ce9b24108f68c95ca"></a>

## Next pages — http1_config.header_transformation / 3fddf4f4d565 / 4

- [http1_config.header_transformation.default_header_transformation](resources--cluster--reference--group-001.md#canonical-e456117ff3bd11c168cb1deb82fafaca3de6673fb5d4bfd3f027b27a3beb5261)
- [http1_config.header_transformation.preserve_case_header_transformation](resources--cluster--reference--group-001.md#canonical-af30163a20dd7ea0773998fcbae339ea6768cec4473eb5283d0192f562f3139d)
- [http1_config.header_transformation.proper_case_header_transformation](resources--cluster--reference--group-001.md#canonical-ad439f2364891cb0730ea39936dc5fc9e520916a9a94ff608247eb9a622094d8)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-e456117ff3bd11c168cb1deb82fafaca3de6673fb5d4bfd3f027b27a3beb5261"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cdd3435705740d0e8b215a24b934910e445cd7f281d72a88c9de61f795e938c"></a>

## http1_config.header_transformation.default_header_transformation — http1_config.header_transformation.default_header_transformation / b807ee27f2b3 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- http1_config.header_transformation.default_header_transformation

<a id="canonical-fd5705c372bd8edcbec65e294ae4ac094f7def3fd27fc05c348ee4a43f5736cf"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

<a id="canonical-85c1299c4c0059ad6c8cf7ad0a591ad9553fc837479bf7c395b0471e85983bc7"></a>

## Direct properties — http1_config.header_transformation.default_header_transformation / b807ee27f2b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5cc464bd0621f903dbc92e26a48105614a51c414f0be46ef16a346cf1a56bea4"></a>

## Next pages — http1_config.header_transformation.default_header_transformation / b807ee27f2b3 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-af30163a20dd7ea0773998fcbae339ea6768cec4473eb5283d0192f562f3139d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fca1d84ddc1bd9170c3ef718029eca613af90ce0003093d31a4d1bf82b5df323"></a>

## http1_config.header_transformation.preserve_case_header_transformation — http1_config.header_transformation.preserve_case_header_transformation / 70285e545bf0 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-0ea1363f6b1dee97dfdd1558afb95cdd38d84b1fa26512c1da515d0b5966d8dd"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

<a id="canonical-02f6bc119ae3e4a95beea75841898b37e6cc24f392fb599c497a99ea588fcf71"></a>

## Direct properties — http1_config.header_transformation.preserve_case_header_transformation / 70285e545bf0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f323dfdf42c06a5d35fbe5ee9a77076b91800f5bb9baf8bcce0beb43edaf2bdf"></a>

## Next pages — http1_config.header_transformation.preserve_case_header_transformation / 70285e545bf0 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-ad439f2364891cb0730ea39936dc5fc9e520916a9a94ff608247eb9a622094d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8669704bc00dfdfc1a76deb03569bbb6559f6d244fa514a4fc66a38fc8d54e27"></a>

## http1_config.header_transformation.proper_case_header_transformation — http1_config.header_transformation.proper_case_header_transformation / 7c919cac2bf7 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [http1_config](resources--cluster--reference--group-001.md#canonical-26fb8e8cdad4b45dde1fb3ac15ac3aeb9ccb2d672bd7e3ec641f110aa683e0df)
- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-7c5428cf322cfe37b945aa1fc88be9efca4608723400793d09aee6d5969b01ed"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

<a id="canonical-b255c3f85fa8936bd94d61fae28ce999ecfc5a3da9935b44df5cc7cb65d13cb2"></a>

## Direct properties — http1_config.header_transformation.proper_case_header_transformation / 7c919cac2bf7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66a27be2440264518060701332b243c4f89159c3e5269dd113e60e7e04371a58"></a>

## Next pages — http1_config.header_transformation.proper_case_header_transformation / 7c919cac2bf7 / 4

- [http1_config.header_transformation](resources--cluster--reference--group-001.md#canonical-e3a496fbee840d3868cedf28a9cc98b57a61d2f7eb23e01be432c775b3dc9285)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-5a4340c3d7f582f647657ffb7085cd625450fe296019861bdd5b6d1c377e0777"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d34d86ba660afd1c6613e7384ad05c2cd5fc05c7a4d445ae6cacba4a4929754e"></a>

## http2_options — http2_options / 35e760243d05 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- http2_options

<a id="canonical-91abadd82ba6567818ad5263633ae205e99430f96f8588f4af74c0e2015defb4"></a>

Type: `"object"`. single nested block, Optional.

Http2 Protocol OPTIONS for upstream connections.

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
http2_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-30668eff9db4b99506dc35ba96fa1831958f6295bb23fb840599cf7a0ce7aef6"></a>

## Direct properties — http2_options / 35e760243d05 / 3

<a id="canonical-3c2e14c6515b2318c9418a7da04cbc34f71d50db079ca3a648efc7c23145d8b7"></a>

<a id="canonical-1ca2f204843ff7afbf8360b62840670c2f242fa0499e92a4d36203f34c339734"></a>

## enabled property — http2_options / 35e760243d05 / 4

Type: `"bool"`. Optional.

Enable/disable HTTP2 Protocol for upstream connections.

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

<a id="canonical-6bae1fd17bd0b74bb15bb001ac7bc7b8d73fbec04bde8ab9a6821e7b2fa920fc"></a>

## Next pages — http2_options / 35e760243d05 / 5

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-c3a3695a1c7b081f038e7cd5ed3d1c217bb9cd9239af60a986eb5c4c98feb63a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae9a270d9f770467eb3f4866078beb9419f971ced972cf999f301264737eeb2e"></a>

## no_panic_threshold — no_panic_threshold / 95124c157fcd / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- no_panic_threshold

<a id="canonical-0c5ae0df7005c08e4764e302ed0d8ec65685adb59eab3170438ed49002d774e9"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_panic\_threshold, panic\_threshold; Default: no\_panic\_threshold\] Configuration
parameter for no panic threshold.

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

- [no_panic_threshold](resources--cluster--reference--group-001.md#canonical-0c5ae0df7005c08e4764e302ed0d8ec65685adb59eab3170438ed49002d774e9)
- [panic_threshold](resources--cluster--reference--group-001.md#canonical-6ae632f1a1b7c68a21072b5ff1753fc41f50111305f7e77b342f1326ac9cd4a0)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_panic_threshold = {}
```

<a id="canonical-8c945030c98351d32bd0dc47e5a4ab7d2cddefedf24303782d1b2195677dc567"></a>

## Direct properties — no_panic_threshold / 95124c157fcd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9bc869a815ec3e738cf4268e2cafcd24f28b4f560d0beb00e0d8760b8d6288f"></a>

## Next pages — no_panic_threshold / 95124c157fcd / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-6a3ec1c603dcd6c9a1200adf9cdb795b80a906d4a7fc21c796f92507d72a5cda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-132c297b299a75a965b044f44023cf0de2b36019e01567353c4c466b65734ff3"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / a40ba4dec755 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- no_request_limit_per_connection

<a id="canonical-b782d87536617bf6cf2eb48c4214a20810fd35a28dcab745323a8bfb706c5518"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-795de37546ebab528debca45557655ea9f0285484370bc8f90e2dccd4a608e1f"></a>

## Direct properties — no_request_limit_per_connection / a40ba4dec755 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d32af973e30a5be83481ce4a8ba22edb92e7d8387838afbd1c0ddc0664aa8ca3"></a>

## Next pages — no_request_limit_per_connection / a40ba4dec755 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-7e2b475404eb58b71efe89676a31ca555f4a8f25da60029de259509bf7fbd314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa450d6c2789a3fa33bd7fef0c77b79825545ce39aa997b847d6e28afdb70508"></a>

## outlier_detection — outlier_detection / 1faa068c1bc1 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- outlier_detection

<a id="canonical-80c771b6adb3d962e905e91ac1ff6291de88b10e015b37de851a3ef372d45d33"></a>

Type: `"object"`. single nested block, Optional.

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking. Algorithm 1.

Upstream description:

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking.

Algorithm

&#8203;1. A endpoint is determined to be an outlier (based on configured number of consecutive\_5xx
or consecutive\_gateway\_failures) . &#8203;2. If no endpoints have been ejected, loadbalancer will
eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below
the allowed threshold (specified via max\_ejection\_percent setting). If the number of ejected hosts
is above the threshold, the host is not ejected. &#8203;3. The endpoint is ejected for some number
of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during
load balancing. The number of milliseconds is equal to the base\_ejection\_time value multiplied by
the number of times the host has been ejected. &#8203;4. An ejected endpoint will automatically be
brought back into service after the ejection time has been satisfied.

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
outlier_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-a912613c3c0dc5179ab6439a65348c4506d48a9ce2154497735fb1826ab46488"></a>

## Direct properties — outlier_detection / 1faa068c1bc1 / 3

<a id="canonical-607b84026e14ce33ebfda8f2dfc6e331a7e19ff07584d430f7848c74b97b3782"></a>

<a id="canonical-0add3cc090fe4cf10e4d5a8ee61ff33e13f1fc98e174335cd2bb2cf59d0f5b7a"></a>

## base_ejection_time property — outlier_detection / 1faa068c1bc1 / 4

Type: `"number"`. Optional.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail.

Upstream description:

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-c6c64f871cb9f46bb6e2baaa429e5aaf4b4c53bc286ef7eda44514311312667b"></a>

<a id="canonical-1c1af3a99c245e41394717ef6f57358b1ddb98960d15ba0a92562d98acee2502"></a>

## consecutive_5xx property — outlier_detection / 1faa068c1bc1 / 5

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates
the..

Upstream description:

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates the
number of consecutive 5xx responses required before a consecutive 5xx ejection occurs. Defaults to
5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-073eb12937573aa8f5d6ec0770f62d14c104bd0d146e7bd6841d17644d0751f2"></a>

<a id="canonical-a9b9297887e3e37c65360731ba2caaf3867cab7073891f16ce069638359a2c99"></a>

## consecutive_gateway_failure property — outlier_detection / 1faa068c1bc1 / 6

Type: `"number"`. Optional.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.)..

Upstream description:

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-c9d3004d7d0fb237cb239e82d3cf637e970dca80b0d353421188d19e68c2d531"></a>

<a id="canonical-dd97a8dad321b4fa9d90d6aa16ac3b1a2bf8fea0c65f456fe40574e9faa1bacf"></a>

## interval property — outlier_detection / 1faa068c1bc1 / 7

Type: `"number"`. Optional.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Upstream description:

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to 10000ms or 10s. Specified in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600000),
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-7cd129b1dc4618dc46a0807f663ad584374972dacf9ce000e0af478cfc1d72f4"></a>

<a id="canonical-f89817971100b12e362ff01fb49b3e32922d9ec18a87a49bb028b7d37ba9fc9a"></a>

## max_ejection_percent property — outlier_detection / 1faa068c1bc1 / 8

Type: `"number"`. Optional.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Upstream description:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-308032d1ebd5fd7605525351aee1b06092326eca0d531544d560db9e39cdb6cb"></a>

## Next pages — outlier_detection / 1faa068c1bc1 / 9

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-0faea09edc5f6f80d46593aa359b0a3d9532955253883a1fec198705ea86c6a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-694c2b82c753e0cdd1f90e4b035d1f756a39c2b2911185b3cbc3e82dd2b696aa"></a>

## proxy_protocol_v1 — proxy_protocol_v1 / 8f930388b1f4 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- proxy_protocol_v1

<a id="canonical-1b2bd886a3e494ea5ed6bfb99fef6bd969ce630197665b939d5d03accdc9c161"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v1.

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
proxy_protocol_v1 = {}
```

<a id="canonical-55cb19333a07a47055959b94f8741c78bfa8c18aacd4b1f5c749b7aa585dad4a"></a>

## Direct properties — proxy_protocol_v1 / 8f930388b1f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06d0e8f0aa2a71018440316ff193f1f97e148950f3fbec559a1e253cdc00d5a3"></a>

## Next pages — proxy_protocol_v1 / 8f930388b1f4 / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-4cb0ace12438d5ea44c00d424f1cf3656df4abb50a9478f84670c01f1f5a5c8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a60cb169f3b9327f24fab15427826e17941c6bd41eb95a7cece0e4ec7ab0dd5"></a>

## proxy_protocol_v2 — proxy_protocol_v2 / 9c52ea23cbdd / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- proxy_protocol_v2

<a id="canonical-d236d9064833a6e6d9b63c6b247c73347808661797bf8fa18261a01e9f9568ff"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for proxy protocol v2.

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
proxy_protocol_v2 = {}
```

<a id="canonical-9355769857bd492bd12a49944f8a4d57f7198a8ace5436d926def1a40b94d6ff"></a>

## Direct properties — proxy_protocol_v2 / 9c52ea23cbdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78cf0dd6b5ce367aff02756e385442d22f00783f5a29152f244e497f80737c35"></a>

## Next pages — proxy_protocol_v2 / 9c52ea23cbdd / 4

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-bc30be729566d8a5ea5e028084a80da67eedaf1c5d039840a9ab5d9c30f60e84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92112f2bd2c1946f588969f0d739b6c3356d18ece93495cd95059461408c3e97"></a>

## timeouts — timeouts / c784307ad63b / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- timeouts

<a id="canonical-056060cdd509edf92db81cbb6c10d9fd0d52fa118432c1eb62cd5e21a7167f58"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-d31b7822a343cabe1922cf349605d4fdbb141b359661f5968613a5c88013de40"></a>

## Direct properties — timeouts / c784307ad63b / 3

<a id="canonical-5ced1feb443a7e7f5fc2e5e3dacff54054253e899955ae95da9f7d684e143cee"></a>

<a id="canonical-e6d9783f87dbaf8977c3473a6acb01767f3d1be5faf40518e16776b87b2ff536"></a>

## create property — timeouts / c784307ad63b / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c6591149383f3dbd5f34297409873d2a78f93f060d2fca9912d4a77ec6ce9472"></a>

<a id="canonical-a88584b4e12b3f8d175cea7dc4e313a72567be9c63fa2deb439f2b11696e49ef"></a>

## delete property — timeouts / c784307ad63b / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-9daa3efe5d66a34af899a7950b09ce296205dfe776d6c455fa2a0b420ff2c3c0"></a>

<a id="canonical-12de5e45394a39c492947533f617bc3f2f45e8fb8b26733ebdcde8e6ca14b7d0"></a>

## read property — timeouts / c784307ad63b / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-e03f31c6597718922718999d58a3215e7ff3102d96551ab4c60e9cce644fabfc"></a>

<a id="canonical-4145a0cb969d2678b32f221aeb638402f6b360fb7e4ab2a35807c8a34908386b"></a>

## update property — timeouts / c784307ad63b / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e55160077310c16051513006e381f7bc6e3081ecca2c755e2f64a4441b5e5849"></a>

## Next pages — timeouts / c784307ad63b / 8

- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5371691f0375620b2fe36a93acb363f492bd8039d20c88298a99a20d6bc12984"></a>

## tls_parameters — tls_parameters / e5ae3270508e / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- tls_parameters

<a id="canonical-dd8a87d2474107719e6379d039e437638cf0d967563ce1f71351d27a58a3c998"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-00a7a934936a4aaa12b8d27f5da2a279996aee709271a97c49e1f078e0f6a7d0"></a>

## Direct properties — tls_parameters / e5ae3270508e / 3

- [cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a): complete subsection reference.

- [common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd): complete subsection reference.

- [default_session_key_caching](resources--cluster--reference--group-002.md#canonical-d5938793cbc4f992c49a04eb7cd82d52a334f576743046a32b760b559c56a81f): complete subsection reference.

- [disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-92ec7b9b9a70b93cab15f9d7c462b9b4cfa2d511834f0888f5363db22418f5f4): complete subsection reference.

- [disable_sni](resources--cluster--reference--group-002.md#canonical-6d84cd0eae7dcfacd28f41de2d59fbe2960aa31c3666b0f7722e3f2d3416ee9e): complete subsection reference.

<a id="canonical-624ba90a9f547b00a7825255798b78b590da24a62db543fb160ea5edf5fdd1fd"></a>

<a id="canonical-2afa030bea66d47661d6fa00dd8ed277fdffe07c9dbabe9317f3df5b331e6365"></a>

## max_session_keys property — tls_parameters / e5ae3270508e / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="canonical-4ebecbeb32da21925b86b69e08390ba8a2ef960458baba4c926f6012385b7ca9"></a>

<a id="canonical-1d5a14c374c69ef5af41ef03aadae88a5f12b492a852684fb26ce4958737868b"></a>

## sni property — tls_parameters / e5ae3270508e / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-f3317e3237accb9a6c0b6e3b1684f389cecf00eacc6e283dce291bbf3cf7e30c): complete subsection reference.

<a id="canonical-ec0920afdabd7764fb9a725f60260cdfd55fdd5c99deb795631d440132cc326e"></a>

## Next pages — tls_parameters / e5ae3270508e / 6

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.default_session_key_caching](resources--cluster--reference--group-002.md#canonical-d5938793cbc4f992c49a04eb7cd82d52a334f576743046a32b760b559c56a81f)
- [tls_parameters.disable_session_key_caching](resources--cluster--reference--group-002.md#canonical-92ec7b9b9a70b93cab15f9d7c462b9b4cfa2d511834f0888f5363db22418f5f4)
- [tls_parameters.disable_sni](resources--cluster--reference--group-002.md#canonical-6d84cd0eae7dcfacd28f41de2d59fbe2960aa31c3666b0f7722e3f2d3416ee9e)
- [tls_parameters.use_host_header_as_sni](resources--cluster--reference--group-002.md#canonical-f3317e3237accb9a6c0b6e3b1684f389cecf00eacc6e283dce291bbf3cf7e30c)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af0b18a8697b79e2fd1ceab3f59f8be1229c9ea69c0c3a2a9c2d040b13b790ed"></a>

## tls_parameters.cert_params — tls_parameters.cert_params / ea95b1bcf255 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.cert_params

<a id="canonical-b9998fdc13a0a5f04aa41ef2f7bae90630f498a9d562833dd2a0b0035068e04d"></a>

Type: `"object"`. single nested block, Optional.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "tls_validation_params"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("tls_validation_params",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2f114a4b6d94e0ebd72b6f77b92baba1639b4a90db31739841c33792dca932c9"></a>

## Direct properties — tls_parameters.cert_params / ea95b1bcf255 / 3

- [certificates](resources--cluster--reference--group-001.md#canonical-eda56529aaaec6a91b326f54d27fc7a97b829a27bd663379ba1de0676d3c53f3): complete subsection reference.

<a id="canonical-7a608f048ad7ab2b8b246a27b581a648d8f3746dd623e583383f75b7b3220689"></a>

<a id="canonical-1e35141409cc60ee2b5aecfed5369e702b6e81497fc56b6830d46f2f7b8573e0"></a>

## cipher_suites property — tls_parameters.cert_params / ea95b1bcf255 / 4

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

<a id="canonical-4507b54c6416365dd3f9c583578c242a54fb6c3b4cef87dd2e2abedd0550e45f"></a>

<a id="canonical-4782643b71f2a6a8b2e5083ce22ac348c3deffb2cc96ebf9759903bc8ce85808"></a>

## maximum_protocol_version property — tls_parameters.cert_params / ea95b1bcf255 / 5

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

<a id="canonical-af9ae12f906508dd7e7391441a1d95e42c028c5b00ea7915d228b8499500e35d"></a>

<a id="canonical-7129ab86d28cf22c7963b80ff518322313c2eb42f595078a1fda2b2b4113716d"></a>

## minimum_protocol_version property — tls_parameters.cert_params / ea95b1bcf255 / 6

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

- [skip_server_verification](resources--cluster--reference--group-001.md#canonical-ec5e6a1475d964f25adf151ab044b4b0b5c2bf61f2e3d2cb503085fa8889d335): complete subsection reference.

- [tls_validation_params](resources--cluster--reference--group-001.md#canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5): complete subsection reference.

- [volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-2e051f67c57df12737f0cd3b7509fc1670165fd8fb0ad5dc3043ec4f35434658): complete subsection reference.

<a id="canonical-6112d49389d45e2f02de661699c3114cfef42cf5c3fde2b90dbaf9c75379f803"></a>

## Next pages — tls_parameters.cert_params / ea95b1bcf255 / 7

- [tls_parameters.cert_params.certificates](resources--cluster--reference--group-001.md#canonical-eda56529aaaec6a91b326f54d27fc7a97b829a27bd663379ba1de0676d3c53f3)
- [tls_parameters.cert_params.skip_server_verification](resources--cluster--reference--group-001.md#canonical-ec5e6a1475d964f25adf151ab044b4b0b5c2bf61f2e3d2cb503085fa8889d335)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5)
- [tls_parameters.cert_params.volterra_trusted_ca](resources--cluster--reference--group-001.md#canonical-2e051f67c57df12737f0cd3b7509fc1670165fd8fb0ad5dc3043ec4f35434658)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-eda56529aaaec6a91b326f54d27fc7a97b829a27bd663379ba1de0676d3c53f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-736f86e527cbf2aa96e7c8ad5ecd391272ffa6897624b57cdfdbd6ca07fc900b"></a>

## tls_parameters.cert_params.certificates — tls_parameters.cert_params.certificates / 9edfa37df951 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- tls_parameters.cert_params.certificates

<a id="canonical-addfdb5e04a31cb8510b45dbfb6991e8912bd3c6fdf08f2fa2ae266fec96a6f8"></a>

Type: `"object"`. list nested block, Optional.

Client TLS Certificate required for mTLS authentication.

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
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
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

<a id="canonical-c3929b7b916f45b01293160cdadb62f48d47ccf22e00aae4e21d039667aa533a"></a>

## Direct properties — tls_parameters.cert_params.certificates / 9edfa37df951 / 3

<a id="canonical-fe3ed84b554f1463f9fa01ac3323fbc59694824d57f11422bed145b446a585c6"></a>

<a id="canonical-ae08298f7405a54c0606de9ee071b0247e0f23d4bdc1f9b7b49a3be3ff1ad8bf"></a>

## kind property — tls_parameters.cert_params.certificates / 9edfa37df951 / 4

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

<a id="canonical-7f157af08efecd392add8982040ebfc33ed57dcc9fc457c2dcaf92653146c57a"></a>

<a id="canonical-46137875ed59d6bc7e19d0986c0ad62489a559967694dd0aab6d79267193cb5d"></a>

## name property — tls_parameters.cert_params.certificates / 9edfa37df951 / 5

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

<a id="canonical-e70485500f55d0440fdc71fb625c88bab065201b73d701bbcd7ad7d55ec564d3"></a>

<a id="canonical-350be09d3621528e8234fab52602585f2efa693f22633f23aab137f73351deda"></a>

## namespace property — tls_parameters.cert_params.certificates / 9edfa37df951 / 6

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

<a id="canonical-a0bd99152dc53c229da2ff7fd0c10c861f0189c574308d74eb60b673373bf007"></a>

<a id="canonical-9f6f0a6837e60e69cc06379e40703eec3a9288aacc75e8d7b4b7ddb4d3fa49b2"></a>

## tenant property — tls_parameters.cert_params.certificates / 9edfa37df951 / 7

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

<a id="canonical-e9526204eaf32121a6ba8a7857f08be21dc19c6121eb84c0a437071c21c2e4d2"></a>

<a id="canonical-0459e737b7decd121c9304265a3ed82f94450b284953eb88ef51f23ab1d7ef49"></a>

## uid property — tls_parameters.cert_params.certificates / 9edfa37df951 / 8

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

<a id="canonical-e3de5aae42b4249a538350de707a170471889a4133e17d5862cc94f36ab7f7e3"></a>

## Next pages — tls_parameters.cert_params.certificates / 9edfa37df951 / 9

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-ec5e6a1475d964f25adf151ab044b4b0b5c2bf61f2e3d2cb503085fa8889d335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e612e29980d477007ac0e1f88544ec35f4ba5eda183c174542e4f6071170230"></a>

## tls_parameters.cert_params.skip_server_verification — tls_parameters.cert_params.skip_server_verification / aae1fde1c191 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- tls_parameters.cert_params.skip_server_verification

<a id="canonical-b4e18753ce6a35fe86e27f747b5153df3793919d726a5c92fab3e4778abc1ccb"></a>

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
skip_server_verification = {}
```

<a id="canonical-b3e56a81dadd348aa127013abfff258b7b30d129f057ddb8ff9459571f70a7d3"></a>

## Direct properties — tls_parameters.cert_params.skip_server_verification / aae1fde1c191 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2369f7afec1f27dd51fd3aa387e041a6e29bee914ea72b3cf5177c1753b1c50a"></a>

## Next pages — tls_parameters.cert_params.skip_server_verification / aae1fde1c191 / 4

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9eec6296760abab1a72088fc0917d7de787418a3f67da94af83fb3dc5557ef05"></a>

## tls_parameters.cert_params.tls_validation_params — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- tls_parameters.cert_params.tls_validation_params

<a id="canonical-55a6a6ca75db984cfcfa83156007bbccec48479557a7888cde4e1274608ebdea"></a>

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
tls_validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a461e3080b93ae245d11c490977aa916845f4c0107ee5ebca500bf7f5e7f218"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 3

<a id="canonical-19f6cc099d61755427cb138854779f6f41ff6c7776ead57f0a733e87c94b512b"></a>

<a id="canonical-2f98e7502e1972b7306de69cbf1236a996fc848c385deab299248f64d26636aa"></a>

## skip_hostname_verification property — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 4

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

- [trusted_ca](resources--cluster--reference--group-001.md#canonical-8932490d6d2c71d21ae11386b2fbf5573f92cf27f240175fdfcf727cd37e877a): complete subsection reference.

<a id="canonical-036e3bdb93960f7dea43753a63cde4b65e00e4052efe62bec44334603f2b3979"></a>

<a id="canonical-2cc3b86a70ce47636da39d37a4640d9deae320bf7adc31aadc6516b440d94e38"></a>

## trusted_ca_url property — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 5

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

<a id="canonical-c98d20c59e86848bd3918cc22c9966ea0646279a20708f29a90834bb975247e9"></a>

<a id="canonical-833f3f3e104890a7664a3215590849089f120e7463f8c565023432c9ad970195"></a>

## verify_subject_alt_names property — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 6

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

<a id="canonical-c23cfda234861eb85ded9578a1645659849b410a0a0c793e52029e5f82c8d8cb"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params / 47b1d23d3cd7 / 7

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-8932490d6d2c71d21ae11386b2fbf5573f92cf27f240175fdfcf727cd37e877a)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-8932490d6d2c71d21ae11386b2fbf5573f92cf27f240175fdfcf727cd37e877a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18866b855a7dbf63a672b1c8fc77a1f85be22a42c1063541f9563dafa6aff2ce"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca — tls_parameters.cert_params.tls_validation_params.trusted_ca / 09bf22aa89fa / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="canonical-99cbd2465519e92554e8b0173bca484558991933c34bedfedb82067aebd5c8e9"></a>

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

<a id="canonical-8e005d66e9875e8817d1cd52d3acac4d1a8be8f2a339bcce86fcfaf4a9e1936d"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params.trusted_ca / 09bf22aa89fa / 3

- [trusted_ca_list](resources--cluster--reference--group-001.md#canonical-e4886beaaa546d5032824de254b599619fd34e365b6caa15751510eab231ade5): complete subsection reference.

<a id="canonical-6c22e6eb9f9e1cb48cb70b744d6cc78785c780579cdfb521a23587a955216a04"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params.trusted_ca / 09bf22aa89fa / 4

- [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-001.md#canonical-e4886beaaa546d5032824de254b599619fd34e365b6caa15751510eab231ade5)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-e4886beaaa546d5032824de254b599619fd34e365b6caa15751510eab231ade5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f51494f0ddf595c61b1e95fbe4d8dde040e1830bf1bcf5b364fa2e0f7dc1eed"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-001.md#canonical-91141632f5b12bed9b78a57e7cfb9b5fb301382d5c4f07fad2a3ebda20bd62d5)
- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-8932490d6d2c71d21ae11386b2fbf5573f92cf27f240175fdfcf727cd37e877a)
- tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-e038d8bdf3d308533a63be452f07751999fbd9578cb92694fbdd302852360126"></a>

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

<a id="canonical-1d32be51ac33ef7f87b1369680372c07cb10333b2b9dc1343e8e7c2dc0d95f96"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 3

<a id="canonical-b28bc2125656e380d023cf07f0b4e98e9ab29c22a96129d56b152c5de06a1477"></a>

<a id="canonical-e411ba9252d0da21838abbdacc96fa4783150c99ff789a0312d253b24226acfb"></a>

## kind property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 4

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

<a id="canonical-d658b4abb22243d4cdb6cf5ec48145002ae88d7496dfd8811ef98209ddb3ef4c"></a>

<a id="canonical-6bf3609497a1038d92449de7b6da8dd9033c8210c0bbcc39b32b3d30417eb2b8"></a>

## name property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 5

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

<a id="canonical-83aaadc188014eb638c199fe9d9d56de15962f432c533e24e04ec3fa6b89ad3a"></a>

<a id="canonical-4a8a6ad38f73ca1176e0864ec388bf8d9d1edd2236f7519643d3470ca0bc77d4"></a>

## namespace property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 6

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

<a id="canonical-66fefdbed0c9d84ad65f69414cbb581262a98892ba892df1621ca82e050a707f"></a>

<a id="canonical-9b8251bdfac324dfe29091190eb0e3845abf8de2f2964ca0cef48d4c064140ee"></a>

## tenant property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 7

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

<a id="canonical-526fa0ba78e9014c04e1ebf0ed02ec6e973913c48a18a5c9ac73818b208b155c"></a>

<a id="canonical-16999f5e62b26c4208cdc1316b832211045028d7777f74f49ad18d30bd4d530c"></a>

## uid property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 8

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

<a id="canonical-36118941a8129c22bfac4f7a8f38b4e46215b8e4706793c01ccd4c47f878ff24"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / 97ae74ea6901 / 9

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-8932490d6d2c71d21ae11386b2fbf5573f92cf27f240175fdfcf727cd37e877a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-2e051f67c57df12737f0cd3b7509fc1670165fd8fb0ad5dc3043ec4f35434658"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5347568f7db4278314279857e560f581e9eac680c0e52266896d31be64df1f30"></a>

## tls_parameters.cert_params.volterra_trusted_ca — tls_parameters.cert_params.volterra_trusted_ca / e40e63cbf5c9 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="canonical-37ae27aa12d3c9435aab0eae126cf6010fdf68e43dbdf6070f5581ff809e19ae"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-b601166f75111a60b898017fbf7a61d9a6bcfa8f8b58a2ca8e54b6f0cff66936"></a>

## Direct properties — tls_parameters.cert_params.volterra_trusted_ca / e40e63cbf5c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9aa510c27e1f52802d7994714212ca2101cd910420716e8f2ce28f7c9283063d"></a>

## Next pages — tls_parameters.cert_params.volterra_trusted_ca / e40e63cbf5c9 / 4

- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-508c271a0acc2a1fd357ba80bcf6b50d9cde9f0acd190935ba2c52826182110a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0aebb72edbfc6f0284b10adc5e477c376879f7479f3ce97d6f2c79702df3008"></a>

## tls_parameters.common_params — tls_parameters.common_params / 42c5aa2f498c / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- tls_parameters.common_params

<a id="canonical-1f277e092d4675646bf7274ea5256c891590f9a536ed3baa1c42c2cf524cd24d"></a>

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

<a id="canonical-2897d644f16cd6f5b800a828360dd4f3d14c990d422085082c0e1b26791fc3c4"></a>

## Direct properties — tls_parameters.common_params / 42c5aa2f498c / 3

<a id="canonical-5f95547a893f9d83522b5fb05beaf2d03b884c0c9397ee42d15ef6f15dde05aa"></a>

<a id="canonical-210c1c05a867ace3322c09959b710d795081553fe3e01c7abd53dce81c4261bb"></a>

## cipher_suites property — tls_parameters.common_params / 42c5aa2f498c / 4

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

<a id="canonical-38b6cab40d0433ec6128b2119fd8195c6d3c5bdeca040c0dee25529f233b8267"></a>

<a id="canonical-5ebda1753231245a010a634f372c51bff5afcd26a146464215962c75d9deb59e"></a>

## maximum_protocol_version property — tls_parameters.common_params / 42c5aa2f498c / 5

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

<a id="canonical-b4d1f704fb77a685cb192e40f3a244e798d9d70b2038a2252a60a9478b29b75a"></a>

<a id="canonical-67b5606e2d0a3582e63499126b0d38d589d488af038705f9f5779c77ba8c6326"></a>

## minimum_protocol_version property — tls_parameters.common_params / 42c5aa2f498c / 6

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

- [tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42): complete subsection reference.

- [validation_params](resources--cluster--reference--group-001.md#canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93): complete subsection reference.

<a id="canonical-9d6c5dd56ac297e2e668705fab07fa0f3c29074e0c0b43bb6cb53178a8a0ab38"></a>

## Next pages — tls_parameters.common_params / 42c5aa2f498c / 7

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-001.md#canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dd07b742ae4b29f2ac06760c2ad0f072007dbcd859845f9def1b62acfe6c2d4"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / 1d724ff35698 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- tls_parameters.common_params.tls_certificates

<a id="canonical-10435dbffdb135aeef9f3ad66f251def795099d76ea3b9259e708b9f0b5f4c9e"></a>

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

<a id="canonical-5c37799656c1c7b9edcce002086afa5f45d211db3dc91edcca27ed016dcdb727"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / 1d724ff35698 / 3

<a id="canonical-d1d448e38bcd862fa131035b28fc7baa0cd407aba326df88a12fe695c2ce3486"></a>

<a id="canonical-e779854825d2b81a3d21108d921a03b7fa131f538baa52e00998719f5fb79ce2"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / 1d724ff35698 / 4

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

- [custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-d032288bfe49b906e5a4558f4487ad8be1c13649ce3e4ddb44fad383a8fca9b6): complete subsection reference.

<a id="canonical-4a2c005adc93db0f5a82ee57ae9828a40da60ca2605fcdb15243aeabc9c4ee0d"></a>

<a id="canonical-2dc846d6e95ec9d0a626e98f6f11e453be93b2e4afea1b6ac59c4ce07073b78d"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / 1d724ff35698 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-5719deb83275bbd8d1a127e3977d1703ae736498076f2005e08796d8cee18045): complete subsection reference.

- [private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd): complete subsection reference.

- [use_system_defaults](resources--cluster--reference--group-001.md#canonical-f66df9c28f53a0356f6325d0feb7cf42abdc8feffa2e4bf041741e75e4e25092): complete subsection reference.

<a id="canonical-22755ecbe4cd732bf61eb98e0718cdfce67c9b3f931e2b296839b888f68022cd"></a>

## Next pages — tls_parameters.common_params.tls_certificates / 1d724ff35698 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--cluster--reference--group-001.md#canonical-d032288bfe49b906e5a4558f4487ad8be1c13649ce3e4ddb44fad383a8fca9b6)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--cluster--reference--group-001.md#canonical-5719deb83275bbd8d1a127e3977d1703ae736498076f2005e08796d8cee18045)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--cluster--reference--group-001.md#canonical-f66df9c28f53a0356f6325d0feb7cf42abdc8feffa2e4bf041741e75e4e25092)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-d032288bfe49b906e5a4558f4487ad8be1c13649ce3e4ddb44fad383a8fca9b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9663c93e074a8091d5e2e5cd4efe787b2abaf5314330de7c4de6819a6ceef9d4"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 41b19be84e62 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-d137a891e686c4cfa26f512991ae1e88f10aeae5db0efd0453017ff58d5a779b"></a>

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

<a id="canonical-bbde8a6b265f0b31848e9b925fca667e89731a4113ebb97c6fe109cd478fc039"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 41b19be84e62 / 3

<a id="canonical-6cd45ba48734d0a007b4b4b3dcbe3bedf108cb71e06560f569d6718df4524f07"></a>

<a id="canonical-604b4098a214ddfe50e5822d7fc9a6f570a5730bff3b9234a04aa3742bec62d5"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 41b19be84e62 / 4

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

<a id="canonical-a3e4ededd3c422c5526704b5f924296aeb529447eb209f21ead48096bd865994"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 41b19be84e62 / 5

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-5719deb83275bbd8d1a127e3977d1703ae736498076f2005e08796d8cee18045"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07fab5287ea5cf5b319cda1f6301963b1a8001b13b0ecd8edfcb987f86a4f0ee"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 27abed215f7b / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-42c450dcb16715afee2f4ccaa2e4a84b78a417def91b50f65e4f86d21d897db6"></a>

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

<a id="canonical-94e24790ee6dee7cebdcba77ee7f9ed9fc8e002275d1344791e2f416d4548fd0"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 27abed215f7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac21685066a69b2a20753ce98f78276f0370718f47842352a5614b1405bb789a"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 27abed215f7b / 4

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20cba0c3e031d48cc2c35cedaabc6d2f5c13efa92b53d2eec469090c6d416791"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / 17e421ae2ef1 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-a0eb9ffb350218a4ce7eb459feb35176c778aa8f7eaa75a10ce9ad820a4c74c2"></a>

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

<a id="canonical-ddcf78212f1bd569287cda54befa73c6014643d641bb6915ce05069e6ea80e31"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / 17e421ae2ef1 / 3

- [blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-55841998903387d2d7a27f15b6e9df0a50c45da124ec19bffa89e1461ef0a9f0): complete subsection reference.

- [clear_secret_info](resources--cluster--reference--group-001.md#canonical-433184695e3d4bc7bf9f651921be04b6ce5830979ef77d429138b3b5d9813c84): complete subsection reference.

<a id="canonical-bf8ce193bcf4b0cec237cbcecac51c87a55dbade3d0b81aa44108bb609b40c58"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / 17e421ae2ef1 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--reference--group-001.md#canonical-55841998903387d2d7a27f15b6e9df0a50c45da124ec19bffa89e1461ef0a9f0)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--reference--group-001.md#canonical-433184695e3d4bc7bf9f651921be04b6ce5830979ef77d429138b3b5d9813c84)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-55841998903387d2d7a27f15b6e9df0a50c45da124ec19bffa89e1461ef0a9f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89abe475b1bb923d09927f4a6ace84b7b4def0b08418974bd8c7cab591d82ee7"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-f5922691f2a61281fdee3596800c37b1f47c700926abd368033e9f9cd6185b45"></a>

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

<a id="canonical-b0dc75e6939cda79a664d469abe52d4a1967fcc293a20768189c87247f8bdae7"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 3

<a id="canonical-5fec8c3b96656f1cc2044ed285cd8837df3d27caf852085548f388695251392c"></a>

<a id="canonical-9be2fd175eff48f7198b0f0f43942ae8f0021076b6b85c4ba6bbf90c585c9e0c"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 4

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

<a id="canonical-dc957c9510dc06a5bd1357d077eebcdbd65669e9f8aebb0b5d9e1b14303fadf6"></a>

<a id="canonical-0af8037931cf5d298828f04a18b898a127e2d2c4d81f6bd5c91d1126b3748abe"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 5

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

<a id="canonical-1d0ba3ae987e39cddf0f2eeb9f2f37a024a521327eeff7b5a190f7c12f290a77"></a>

<a id="canonical-b1330883a506d69af44e7880aaf6a141a53ffdd9eebab151f138c969ac375157"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 6

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

<a id="canonical-54050970d5f3d6dc867317a83c64693cefdf001f29397c16184ca4207116c331"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 3fc4a7ae477d / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-433184695e3d4bc7bf9f651921be04b6ce5830979ef77d429138b3b5d9813c84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c849988b2d69137abdc426d531120f2251af9fa84132924acc173ed76b295b4b"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 51e8f31c9e94 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3bb3c7d0bb9cdf37c9616b285ea1a8010ce3830faaa5560f7df7b31fe9cff1eb"></a>

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

<a id="canonical-66cbfbbd420b556be8e6ae9f9d1aa70cb684f5c7b020e928c2a43b727bb13964"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 51e8f31c9e94 / 3

<a id="canonical-f393ffb4e24c0c8456621e693e683dec58a9baba84b3201fdef5a69b102a65f3"></a>

<a id="canonical-b6d1701f9f1afc99409e899794eeb062dd81a7f1620783024a50ac16a3e1b3e6"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 51e8f31c9e94 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4bc7f207f1b3472a4d49aa6fa10d0784a8d4502dcc047feba4375c29098e5eb1"></a>

<a id="canonical-6422275ba66101222e5ac16d3bc2875113421def7397c5fe039e289c5d6a0185"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 51e8f31c9e94 / 5

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

<a id="canonical-a47462d9ebe9bab9ba647bf750c7e7069a878bf9139b39b93aa2c99ec99c1c4d"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 51e8f31c9e94 / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-001.md#canonical-16e99117ae37b1517d523f8e53248139730fd40d50642b20160b155ea8a39ffd)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-f66df9c28f53a0356f6325d0feb7cf42abdc8feffa2e4bf041741e75e4e25092"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cecac50839766ee9149e95bb8e5e58d5d70f772fe3325d75c7c967de1f1d68c0"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / f83416ee7d66 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-180f1d45f5c145e8f1b7af1abd357916daacc391c4d3b36dfe9cd1cb536b17b8"></a>

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

<a id="canonical-a2ba2937b4178547c45f1de635510a07b6daa6240f5e2ffe97be4e7133cbc689"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / f83416ee7d66 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12fa0b7aa088ecb5be08fe85fa58970c317424f9efb841726a1923fa686a98eb"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / f83416ee7d66 / 4

- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-001.md#canonical-1d1135aeed9bf24c83adace9347998ef5f78061c085336e95a72912fd9bd0a42)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d74e7d64f2d28fb066754f478af4408f6544b202e408191252a2f879fb9d0562"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / d7062adf944b / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- tls_parameters.common_params.validation_params

<a id="canonical-1f7d2d917efbfa082ebb6ba809a3ba490978c528eaa80b39e160363679fffa50"></a>

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

<a id="canonical-ae41046c1e380736f85ec48782f76a2abbd83e086fc731f5066016c6e467bfdf"></a>

## Direct properties — tls_parameters.common_params.validation_params / d7062adf944b / 3

<a id="canonical-a50f6ae38cc2aba5ead3c14d78f8bcf2127708cf2ed03e6c3594fc03cc7ffc14"></a>

<a id="canonical-316804382fb5961414ddb8c49851822bd400a6700090f5ebb3dfbf257ba3cc7c"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / d7062adf944b / 4

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

- [trusted_ca](resources--cluster--reference--group-001.md#canonical-9b219cc99447098cc906f098d54d9b188b41109365a80a24aed55db0b7aa516c): complete subsection reference.

<a id="canonical-5fe454fe1cd63ad24fcd52bbc90441efe88964153b70dc9d6f96af3ae6645ea6"></a>

<a id="canonical-31183ecfa4f608e137efcaa615d6e72766c8e7be66cd39fff2ef2c2e2edac1e6"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / d7062adf944b / 5

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

<a id="canonical-94d256d3924032c843d8cfc827d45f82179ea7c1b0b316fdbe08cc3b52cdceb0"></a>

<a id="canonical-f338a8a51cf7c14ce20f5fd70847854a311799e1bb60f0fe85e782ba66b98612"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / d7062adf944b / 6

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

<a id="canonical-516b4d1f9342af35bbb934bd578529db1ad6d8cb23b91b0db32d400a87ea004c"></a>

## Next pages — tls_parameters.common_params.validation_params / d7062adf944b / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-001.md#canonical-9b219cc99447098cc906f098d54d9b188b41109365a80a24aed55db0b7aa516c)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-9b219cc99447098cc906f098d54d9b188b41109365a80a24aed55db0b7aa516c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95302f16283a29eca9b3e522991525566634529864f1374820f6a127c387bcc0"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / 3c594bcd01da / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Property reference](resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-e891a37ad95c6abd2732ec43824318ac4c40db0eaa681d108a9e72134d43fd0d)
- [tls_parameters.common_params](resources--cluster--reference--group-001.md#canonical-c0f92b4dae58a6cd257ee8960e68808fa66b0542b22cc0c62afe84cb381f66cd)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-001.md#canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-170719cf37f958c5e9afa9de0a731719b8f06e83d203b46282fa1023b230bbd1"></a>

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

<a id="canonical-ffcf44f4568b8b557bb8b47defa2e8191712750fc5bba58afb6577ca87b08d86"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / 3c594bcd01da / 3

- [trusted_ca_list](resources--cluster--reference--group-001.md#canonical-8df8cf6e5365f2d9457f7c85a4702e28d9ce0da3af890484469ac5623f000640): complete subsection reference.

<a id="canonical-8ef91d837d1a3a21f97c637e34188ed28a099d1158e1c3dae302b43963702b05"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / 3c594bcd01da / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--cluster--reference--group-001.md#canonical-8df8cf6e5365f2d9457f7c85a4702e28d9ce0da3af890484469ac5623f000640)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-001.md#canonical-5df2cdb334c28e2a539dd2549ae97ddc8ccd52d91c6283b4ff928859fd0ccc93)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-8df8cf6e5365f2d9457f7c85a4702e28d9ce0da3af890484469ac5623f000640"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
