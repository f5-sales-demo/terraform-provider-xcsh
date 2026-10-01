---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9a06194fb50b5bbfc673a872188d0cf5f82ad9e77fa9a415edf7b4978ea5dd9"></a>

## Property reference — Property reference / 5edb7d60145d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- Property reference

<a id="canonical-6cb8093e95c472c5fc6468b320da14cfe4afb2abbe633537b0ad21830148e14c"></a>

## Direct properties — Property reference / 5edb7d60145d / 3

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad): complete subsection reference.

<a id="canonical-90eb5a6258b4cae70898213eac45102e7b19580561e8c8c9b42431d74f6852ac"></a>

<a id="canonical-5bffd32fccebbc423388328c6028859406847ee8694438fa5db8111d10abc177"></a>

## annotations property — Property reference / 5edb7d60145d / 4

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

- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-562a3acefba6f1de598f335e49c1ccd5b6fc8bf55ed03dcb7c1b8592e7bb66c5): complete subsection reference.

<a id="canonical-0080089b97f11467549d7760832a8b32c8cd46cf36a6ab314da72a240d87dabe"></a>

<a id="canonical-5092f9291db98e6da95ca8f359eddb810670bde14a079f96df1f1e8de38d3c1a"></a>

## description property — Property reference / 5edb7d60145d / 5

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

<a id="canonical-5149533cf2075800b406322c288c63fa1d5ab62d177a16cb9ae636230c48a20b"></a>

<a id="canonical-8299ca6d64edb83a8e54186df08825c50f32eb6a247a9122a0593c5b7062f36c"></a>

## disable property — Property reference / 5edb7d60145d / 6

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

<a id="canonical-4f80e13c5d133ca735a8bf01d8b9d1faa1e9c74cfc96a1c498b8553436f92d4e"></a>

<a id="canonical-72b2f31308aef5448ff4e3a7838dbdd121867a5f1c8bdf97ec4ac166c573cbb1"></a>

## endpoint_selection property — Property reference / 5edb7d60145d / 7

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

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

<a id="canonical-e713bcf99ffbf391f5165247bdc69f43cda28865a1c2a365bd5b6877e97b400a"></a>

<a id="canonical-da15f5edcd64da8b4cdf0bbe9779738a6320d3f418aeab31cce6f1f254471749"></a>

## health_check_port property — Property reference / 5edb7d60145d / 8

Type: `"number"`. Optional, Computed.

\[OneOf: health\_check\_port, same\_as\_endpoint\_port\] Exclusive with \[same\_as\_endpoint\_port\]
Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [health_check_port](resources--origin_pool--reference--group-001.md#canonical-e713bcf99ffbf391f5165247bdc69f43cda28865a1c2a365bd5b6877e97b400a)
- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-83612ed238eaabdb14d2ae66732f0b85ad5ce12a3af85da13e3e876582fa0658)

Select alternatives according to the provider validators above.

- [healthcheck](resources--origin_pool--reference--group-001.md#canonical-28cdf1ceb2efb9bb358019dbc9717f35a6f260c3edd34c16aa1a76490ce84607): complete subsection reference.

<a id="canonical-5352575b4180862b87741ba2e8107decdecfa4471af6cccc57ffd526ad084dd7"></a>

<a id="canonical-2f5709fb86bf389a4ec3d99abc053dd23fd79387a87e1619fdcde849775dd96f"></a>

## id property — Property reference / 5edb7d60145d / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0c777a749b05344e2548204a3b887a8963fb9f186806c7bc5633693d73e6e9d8"></a>

<a id="canonical-b190e8ef4af80304372e6c14be12fcd94010ccf194bf00fde893d19e1eeaa8ba"></a>

## labels property — Property reference / 5edb7d60145d / 10

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

- [lb_port](resources--origin_pool--reference--group-001.md#canonical-08c2f998734b62d52645d8fd722945970f71788b1cc66b844a568ecb08b11050): complete subsection reference.

<a id="canonical-44483ec971598bcc6f4a7ad69b21b813b38e1af92e09f5b5b527867e1da7f2d1"></a>

<a id="canonical-a689a3094c2220aec4114dff917266c6dc05f3c359d919668f274e6aef451dda"></a>

## loadbalancer_algorithm property — Property reference / 5edb7d60145d / 11

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

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

<a id="canonical-20dc117acb41c2ff688980fe213916884a6d2cc41a4187898a5174d9efb5242e"></a>

<a id="canonical-baec473e12a28e4790cd7bda390ebf96753380a6c7a9316f057b13d18cf4b558"></a>

## name property — Property reference / 5edb7d60145d / 12

Type: `"string"`. Required.

Name of the Origin Pool. Must be unique within the namespace.

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

<a id="canonical-41e3730e4294a20248dbc059ff4fe5257d01ff852ae43c30bb4fc13354512a40"></a>

<a id="canonical-cdca4275daead2e0d7cc415e2b9a33164e45032c9ad2eb5e506d57af46ba01da"></a>

## namespace property — Property reference / 5edb7d60145d / 13

Type: `"string"`. Required.

Namespace where the Origin Pool is created.

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

- [no_tls](resources--origin_pool--reference--group-001.md#canonical-ee8653d9b863ea4392eec42b70dbfc2bfef2c0f4270f8332c7945741c9026b4d): complete subsection reference.

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918): complete subsection reference.

<a id="canonical-be501128f5e3d84e69b9ae443dc6dd371a24c754218e96c6650995b7ad5e7793"></a>

<a id="canonical-87f83d36afe4622539210269b9e2bc07dc92b2e5b89d922c49fa5f7328a274b3"></a>

## port property — Property reference / 5edb7d60145d / 14

Type: `"number"`. Optional, Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-d4a1eba2ebd5b3aefffcc5c49f834145c4fab929d7024ff96905398f2406158f): complete subsection reference.

- [timeouts](resources--origin_pool--reference--group-002.md#canonical-7ae7ce0a89d4f66924dedb8dfdf24bdabfc1c4ceecaf7e928bf3c2815469bcce): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38): complete subsection reference.

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf): complete subsection reference.

<a id="canonical-94679bb717b2f8e6464a89f33efe9c48bf72b7800ba5150030713534699d097f"></a>

## All schema paths — Property reference / 5edb7d60145d / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](resources--origin_pool--reference--group-001.md#canonical-069d803663c2dff549e696fe896c1e215bfb42702b749e2e0163568030200cea) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](resources--origin_pool--reference--group-001.md#canonical-dff138ba4b4a18b9f751174eacd46415025bec6a1349938d2b7e6ecb15959f4d) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-d96543fb7ca049b91c32ada1c3bddad6508e24c7c473b1fcc5f480f6d829bbaf) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](resources--origin_pool--reference--group-001.md#canonical-92dc47ff68a57a711dd8c00052fe46d28ae3d8ae05a79f17083a3956287d6f97) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](resources--origin_pool--reference--group-001.md#canonical-cdc4e7ec612d321352a5850ee0ac1cf00aa8e20bd89cafb92c2400154a40a319) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](resources--origin_pool--reference--group-001.md#canonical-778da4b2f71ab6c5859f29a665aaaffa0598713dc6f7aac9075b7ec35d2e320e) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](resources--origin_pool--reference--group-001.md#canonical-8aa7150541ba59c8bcbbec14de4fd30296a4767dda6b72d3cebfffd7a9ee0258) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](resources--origin_pool--reference--group-001.md#canonical-b11f1cce184dc3b78cab91b475eb0c028a20331a47fdb95c7ca2628c82b4921a) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](resources--origin_pool--reference--group-001.md#canonical-7ec75d3722e6a32356d556c4bba2276b3bfe62c023e4d939c43b468ee6d965d2) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-b5bf4681807d97cd1ccc3bf302884f3118c343d2c661de52bd1e76f27fe1bcd9) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-61ae33c0268245c29b18a4e27bd4c394c0fd514a4b871ec7527225e338d994c3) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-f69735a59a089253569aecbd9e57d28762147af5f2026c261e5d1f0419f59bb6) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-a687e4d6b20484c75edf80ef309f69b61f8a95069ce1ab8736887ee7c2eebbd1) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-6b9e1532e32de91ef9099c0dda8d7c878c59a89de20d3f053cd5f39f44f6933f) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](resources--origin_pool--reference--group-001.md#canonical-a90990e214edae3b36360e951f89226f6102f0d3ee54f01da37bedba49dae827) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-f05c13a51338eccfb51f68ffc0e9d68c8b71cce299b62737da16d9a9eafbc8c7) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1c9d9c9fb2d77bc2df160529b6359f2d89bcd1153c486d7b98c906800bab4b9a) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--reference--group-001.md#canonical-164d18bbc9dedbe728f4ddf8124bc4c20bdcbee7c83573b7883e6579e46f6861) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-d94349b6f66afcb94183c30aeef2bf502e90648822f75ed7595deb69e2130d4b) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--reference--group-001.md#canonical-109cfc5f2f1b9fe16b37024a667d88b32acb0b4e1428e9a271e0167688a4588b) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-c34dc32e2c00d1a77135106fa451810babf4c625596dfc1b3ba6c14b991808b6) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](resources--origin_pool--reference--group-001.md#canonical-61c56b0675d48a41be3d1ced33484227d68f443e7342e6efaedb29effacfcb87) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](resources--origin_pool--reference--group-001.md#canonical-d9aaf050cd7efb283bcf612f0475e378cfa8cd1c9b49d9aec809f859a353caa5) |
| `advanced_options.http1_config` | [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-9e2071cfbb46fd77e5cf7698d2c8a43f0a88a5e286385e6db2f136b20a39d6ea) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-d93e7ffed66ef55089937ffea6c3f8d0c816fc35011e5cdbd701eb9b052f9cd1) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-4cdf1882f711fca1a6a63c42b653d11194f484a88f479713a1303fc29ebc5f82) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-9726d3fc6ab1cfd7eef5af0335424b456a420548efc971dbe062e834c0b5dc60) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-233f160e72d7d8191fde0b491ddbebb0d64a50383411e0f638e5e216db39cb67) |
| `advanced_options.http2_options` | [advanced_options.http2_options](resources--origin_pool--reference--group-001.md#canonical-abd94c4e89f169f6cec459cb7436b52e898f357c06f95b513269f553820802b1) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](resources--origin_pool--reference--group-001.md#canonical-013beaac6e4990222b79607f317dd79bf237d85766c5401b10d336b9282a2d26) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](resources--origin_pool--reference--group-001.md#canonical-79121eaa5c7c1865d913abc4fce7abb9df84433dcfa5fd5af686d092a7f1dbf6) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](resources--origin_pool--reference--group-001.md#canonical-1cb726dd77c9c0b258d6dc205ad6893eab0929a01dcf7625f53d1e4091154465) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-ca955333f3916b4ec9eaaf6716904060f1b22bac7fa59355f420305be93e8289) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-3fe443dcf8df2c689a15077fe76fea3dae05616858f0f6fe32bb318d6f17c946) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](resources--origin_pool--reference--group-001.md#canonical-f3022a07cdb8faf479ec3e61440229376295b1ca336dac85f009d12b509a5668) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](resources--origin_pool--reference--group-001.md#canonical-ea1eb8a6c9f435765a22222f6bd25277ff7be74a5530cf5ab7821f89f9eee1fc) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](resources--origin_pool--reference--group-001.md#canonical-a3e66ecb4fabe054a6d1c5f8c8ec08caf4ecf1cd327ab533b7e75dfb13e9f957) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](resources--origin_pool--reference--group-001.md#canonical-8e91cacb924765ddce8466549fc27dd6cb0698a7c7c02d0d1fed0055fbdd54e6) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](resources--origin_pool--reference--group-001.md#canonical-9c95adf3d0db128c97bdf6d3b9e72ad9fc6586898dab5bc927bf85ba4906e46b) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](resources--origin_pool--reference--group-001.md#canonical-9a67ebcd546ee902c655da1d05023f47b2fdaaa2d6ebddc9b4a40f7064c7bcd6) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](resources--origin_pool--reference--group-001.md#canonical-16f4cb0ee11ca3be584c34687ba431396c58c823d9fe4c5d88a2e8e1194cd7d8) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-5610b07c5be24333551d364fbf5e0a1eb9e533a72cb042e66ab2bbaf7231888f) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-5cf925e609f8ebd8a04a23310c1ea4b9dbc5787e2ea85cf29042ad7143a6ac5d) |
| `annotations` | [annotations](resources--origin_pool--reference--group-001.md#canonical-90eb5a6258b4cae70898213eac45102e7b19580561e8c8c9b42431d74f6852ac) |
| `automatic_port` | [automatic_port](resources--origin_pool--reference--group-001.md#canonical-7328a554a0db646d4173b6b522cf6c6b3633699941810d8295576afdfff78961) |
| `description` | [description](resources--origin_pool--reference--group-001.md#canonical-0080089b97f11467549d7760832a8b32c8cd46cf36a6ab314da72a240d87dabe) |
| `disable` | [disable](resources--origin_pool--reference--group-001.md#canonical-5149533cf2075800b406322c288c63fa1d5ab62d177a16cb9ae636230c48a20b) |
| `endpoint_selection` | [endpoint_selection](resources--origin_pool--reference--group-001.md#canonical-4f80e13c5d133ca735a8bf01d8b9d1faa1e9c74cfc96a1c498b8553436f92d4e) |
| `health_check_port` | [health_check_port](resources--origin_pool--reference--group-001.md#canonical-e713bcf99ffbf391f5165247bdc69f43cda28865a1c2a365bd5b6877e97b400a) |
| `healthcheck` | [healthcheck](resources--origin_pool--reference--group-001.md#canonical-f5007421860a7fab444859ff097418c99a182bdea5969f38c3c31d717f658a3b) |
| `healthcheck.name` | [healthcheck.name](resources--origin_pool--reference--group-001.md#canonical-d2705f7d646a3d90a03f2bac485fa62a1dc7d0f6e57bc3542f4fcf6775887181) |
| `healthcheck.namespace` | [healthcheck.namespace](resources--origin_pool--reference--group-001.md#canonical-ee7245ab3d3850f1df1bc5d5e4187fc9d842be089d9cf27bb80f51ed65163a46) |
| `healthcheck.tenant` | [healthcheck.tenant](resources--origin_pool--reference--group-001.md#canonical-53fcfea637bb1320f4494f5438a0f9cb753d4ac7a413f5671d7aef53f140e32a) |
| `id` | [id](resources--origin_pool--reference--group-001.md#canonical-5352575b4180862b87741ba2e8107decdecfa4471af6cccc57ffd526ad084dd7) |
| `labels` | [labels](resources--origin_pool--reference--group-001.md#canonical-0c777a749b05344e2548204a3b887a8963fb9f186806c7bc5633693d73e6e9d8) |
| `lb_port` | [lb_port](resources--origin_pool--reference--group-001.md#canonical-63fd7950004f798a70fb424fc7458fb474d3b1282aac03cc1481bc5625e83548) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](resources--origin_pool--reference--group-001.md#canonical-44483ec971598bcc6f4a7ad69b21b813b38e1af92e09f5b5b527867e1da7f2d1) |
| `name` | [name](resources--origin_pool--reference--group-001.md#canonical-20dc117acb41c2ff688980fe213916884a6d2cc41a4187898a5174d9efb5242e) |
| `namespace` | [namespace](resources--origin_pool--reference--group-001.md#canonical-41e3730e4294a20248dbc059ff4fe5257d01ff852ae43c30bb4fc13354512a40) |
| `no_tls` | [no_tls](resources--origin_pool--reference--group-001.md#canonical-0d9835930e24bec33e65513b60c6666097a8ff7e52b1d06c06fe840193d9ee98) |
| `origin_servers` | [origin_servers](resources--origin_pool--reference--group-001.md#canonical-81ea379dca927997f61834000c27f38a27f279dff31737cda01b2f67600d8705) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](resources--origin_pool--reference--group-001.md#canonical-ae17b2cd9b8b80dac1ad748930b26182b22dd9b415b22d0a3c85ddabf0ba0395) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](resources--origin_pool--reference--group-001.md#canonical-9ae2544e893e7bfc0e051a53dbc69fbcd5daa06c9f63c8026ddde25f1ff536d9) |
| `origin_servers.consul_service` | [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-bf5d8bd66042946317ef306d10cb327b81bc31ac2cdf4fd0b5c8a9edeb14cbcd) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-540eceffddbc3ea2cc8076910438777def6b70f45d683e0808f5f6eff8fd507a) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-ae04953eb23f15b858ed3eec07a9f6f2af80d44dc20f00dcff0a47fb3df7d1a8) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](resources--origin_pool--reference--group-002.md#canonical-fd653f0d5a6d35504ac6864c8fc0e556003b620622f1fc3e614d8c3cba12d203) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-b4ade0bd48844219b6cfc73d1378472a73d4c422da4f3295adc499bd27d361c4) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-1283d35fa0c67f378ab125a6e0b7dc239a89ffc8890e1056e525de055fb5030c) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-2c7dd3770855afa108bc3cf50f1dae7f952aa15bc8b8a75be3e2834d13e39446) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-78e24096f0cce7a0dcdac8dba9ad9d7a9998111209730a1aab0646e97ecc7daa) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-6ab849b8a33c78fc34cba9bf9224accfce32c9d6b6e8b2edb3119b0d858f1fc2) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-49bc0070edca322e346f277f37e3cd5a85c017c289ab58ac0d02fccfa3582856) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-134f49b25a1bba16a7762ea30eb4ed38d114344d88420fb468fe01612901b230) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-e11c62cc0590dab52b5ce8e1c6c7a238ca57f55e5d35eaec97e0e6f54680f25f) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-3c358daa92a4638e7e857dd0511715b7240b1d724545823465df21f568da1c5e) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-f45b62c68cfeb0b8eb377807de1cd77e058c708438cf963b8bcc1ad4184396be) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-65146b3867f400aafe696cfe3850fc25bdb046ac6a0d018df51aa51d5770a369) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-25866dce6896ba9da923277acfc55e21dad742ff6276874c393bc29708b4bf40) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-88611a1e5f9ed9876c7072c00ea1dbc1a2e8531d31d3559e86e0939ed0fb2ec8) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-722a6f33a3ea2749dafa392f2bc314fc4f553f9e07d654e5d971d65338b5e171) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--reference--group-002.md#canonical-a81bd5af47328b0ba2600822e2823fac1084cd286c7ddbb3cadb49fc6b713d59) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](resources--origin_pool--reference--group-002.md#canonical-03e165f02fd524e60e9393437558b9b089b3ab08b5f0e6012d92fec8191ae8d9) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](resources--origin_pool--reference--group-002.md#canonical-674455876f4f17b283c3099a5fa280a57f5175a96fa6e03730575c2e360d8496) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](resources--origin_pool--reference--group-002.md#canonical-fcb66f47914386d9b69e938f50be5d9538f02a8de039c44bdcc45686a4461f29) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-4cc36d072e3b836f3963a4006394f79d05773937f6f93137562620853f55262f) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-4bc32f8a429bf8d320f06fe432f84355781c6d641025e83f4ac211efb7bcba3d) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-209c35de8baadf74ac5fbc62a5a8d7850ffa035c8763beb97c48ebeed19fd304) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](resources--origin_pool--reference--group-002.md#canonical-356706e2a9910b0a6fc13a0b8729a852116f168e187474761622561f6ce8b1b3) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](resources--origin_pool--reference--group-002.md#canonical-d90bd5cdec234deaa49aaadee54f0268019d2f893292dd563355ac649c77f322) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-24511af17e95e4c39a461030ee8a80cc3b7a0e894d32cc50c7cdef56710ad4be) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-de98207845ff9afe786dba67d5eeb362c67d98d522cc23dcddfb0f41bfebc6ca) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-6dcb4caa518c1bae5ccb8e4f00e1c4ec010ad1acf5719ccdce53bf6b571c6bc6) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-801f1e4681b1603c215908c3579fc78e706417121ac01d7c02fe6f9c05431a68) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-a4101f3b356aedb695c6bfa7a04807f00ee83ebde582664fbf0b4bd3b88b5e26) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-a269e26df9220ee07488cdf47cb035e08b612288d1ec4e39387ca101f678bf1d) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-55fd434ee75ed6b9deb1e21b5bcbfbf8a15d738ef11f761525749453c5d6ec2e) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-36315ceb239d1acc4110673cc4a483090dc86d5bb1f7925662886c6a154c9d06) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-97e6fd5e5d826832b836f76255a771a41d51bb79e1ccc3a82298237d63a8b930) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-77fb38331d336a35ec65521c63972927f194ef30063f660cf1dbaad94edaa74f) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-e4133e83c926000ffb9985337e03d2e65fbd226c7a9adfc6a8773734d8885c63) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-b10aa12219775810e3f7abbbf722b8e4919e11a797318f420756d478e2d8108d) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-2e8de694041bb17a4fc323a0aa6311f4744a4aaa0fca7feb0ba7e9ea5b99ce1b) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-eecf2bd28d2c5c29d426b5660a620aa0359c03e8a185038510282c61a9e23688) |
| `origin_servers.labels` | [origin_servers.labels](resources--origin_pool--reference--group-001.md#canonical-2aeb5c46329d7c19737e8417b48452075f2b1f2b65033b76ae3aacf0add861fb) |
| `origin_servers.private_ip` | [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-1e683ee35788385ef9e0a366ace6cd6aa447b671f350432ef3503bc033caa643) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](resources--origin_pool--reference--group-002.md#canonical-44d9e9329af27e38c264066bb9569ff8efa1260bd5a083d90a203a218395954b) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](resources--origin_pool--reference--group-002.md#canonical-dde746a7871eabf587417e36556e9e9ac730f68e8e7e0d25ab79840fb174c7b9) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](resources--origin_pool--reference--group-002.md#canonical-c86d92076fac7c55c1f9d1dfb32578c1d8e6bc4d9644516a291f6b30cc1e74b9) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](resources--origin_pool--reference--group-002.md#canonical-1a87e0c6dfebeb8a6494e3e957494237de3635f050c384b854586ed3907667f4) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](resources--origin_pool--reference--group-002.md#canonical-6b09762dada6a9bbe039141b10c40ff174c5a1ee1c866b0559ff36be65f06f7e) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](resources--origin_pool--reference--group-002.md#canonical-022cb5b3c00a2c719a585a24b0935c5674cdaff1cbd16abcd622e7d0c7b70af1) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](resources--origin_pool--reference--group-002.md#canonical-a5b7619cfc4b755e049c2fb5ffea1623e4d63e6ba266b9e024da3f9ada3657b4) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-b0b7442b5f15d7ee4a872faf21b86f61b43065a1765293bb469c9fe85b474373) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-bd22e9cbcd50d5d0e415c28d7470de00712657f0385803ce87e156905f2cda37) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-27a120f7b92510cf16ec0a95f6985c757c0ee44368d97ae57d665ef7d9187c76) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-34fb467206054e3e9ff9fd01f2f4da665a2e3344d91a8c8805142181fda62422) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-ff706414df112f3eac19c45c8e54f5d89192cec2d948e50c0da3c1824ac8b295) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-aab2d25f74faa8dea78cb9e823f2f2dd1a44d1b123ccee54ac26fa3c9925ba6a) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-47908698d7d0f70ba3bc47706e61752c46638713a403904bae4f52cc7de4a185) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-a0a086e3cf9f4d8217a683f0532d48e1ce032f4a8174b8ef209f4d401f443299) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-3b200b4ab9c3d42b6974ab6dfa7d940be2e0baac144ad61622d1a6e528d0b6e1) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-6c29470f1ac3ba3bafcf81989b04dec33d03a1cecd98f3a5240e3837d4007baa) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-6700f16253d3439c437e941e88b05d685356df542755de6cbc112ad25e7aa9f2) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0f1080232b6c72ab201fbb30f4412277d08a1979c06b4abfeff77d8615294f23) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-b4f99cdbffb1fb4d25630fd762eeea320e9c5cd4775e1a6cb736a41e9344f783) |
| `origin_servers.private_name` | [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-cac0931ea8a860ed4ecd5bada698762abb726b9716e0be57b7b242e1c7d368c6) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-0ce08da9969c53ad44ad88ce0ce37314bc8cf80480ff7d84407f0c7e31bc7bab) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](resources--origin_pool--reference--group-002.md#canonical-bad88409213fccb90ae3551ca3c95d1c9656693e0c182e2088a3fb4a102099dd) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](resources--origin_pool--reference--group-002.md#canonical-07b4373dd86cb4c9b876536febf5a449936d669c88bc1bc007edc3f285c5046d) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](resources--origin_pool--reference--group-002.md#canonical-41e47e31ed71fce7467d096222d035ba381f0ccba41d2022f08b81bd07b481ec) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](resources--origin_pool--reference--group-002.md#canonical-5f44b1ed95084cbb9e57465f510de16e746e12e6ff21e85a8007f48320e11787) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](resources--origin_pool--reference--group-002.md#canonical-dcc30c1a89676e67a30c75fc76e9b9769ef317bf03a1d30e44d4b85e335c1278) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](resources--origin_pool--reference--group-002.md#canonical-a78ba70e724860151c1573d43e0dec12c7e152391b887893a012aa8277cde738) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](resources--origin_pool--reference--group-002.md#canonical-087b31e7008bad871cad30e1b82a54de451014d5e49f76893b03c5be2cd2529a) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-48068257e132a5c4ef6f21d48f7063e93c689f5057cb6b94b637ac9bd78d3332) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-b453632d3b5064536ddde51bfa0d20dd78a02c0824e1a949ef893b7036680baf) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](resources--origin_pool--reference--group-002.md#canonical-6ffbc1239345816cf22b5942dd2796114c1543cde154db1bfa4e80920b4d7a3b) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](resources--origin_pool--reference--group-002.md#canonical-8c52c68519dadf8713ce2aff2258ef3315c306420827072cf276718772a756b6) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](resources--origin_pool--reference--group-002.md#canonical-dad70ef825aa2825d2f02e32ffd33fa766897faaeaf6ec2faeec0eeb4fe474cb) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-348a7140d5a4fc4900d6c0c1df538ffebece08085f08fad45754eff833ae9d9a) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](resources--origin_pool--reference--group-002.md#canonical-ec235a5ff88ebe29c3978ffd105211e3077825e74f0a81200ee73e1cc38809a2) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](resources--origin_pool--reference--group-002.md#canonical-65897b90931e6191eb37f2ac88a9e6bc47fe44f69f59ab8a2d7298f4713a3e2a) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](resources--origin_pool--reference--group-002.md#canonical-ff63d2601a69a5a94fbb57358104d6467b459829cdaefe7a48ba29ff0e375ac7) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-46ea5444232de51039fd93f759d4f94c067b19c65f703b180872d1fbd761511f) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-178f6d27948ecb0ba6979e87924b25adc765d40398cc373b2ff3b3ba2b7507d9) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-4629448c3ad00f4559b859ca2b593ef693a00eead7de02d6d8075b570d997f01) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](resources--origin_pool--reference--group-002.md#canonical-0a943bd15bda44430af3858613ecfa8c6ab8d76a48a0d95e0ff0e468be8020e8) |
| `origin_servers.public_ip` | [origin_servers.public_ip](resources--origin_pool--reference--group-002.md#canonical-bc493ef6b56befe695eb3e4962b5b785ace5850f7b092275fa7a27eddc3c7590) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](resources--origin_pool--reference--group-002.md#canonical-315ece931a77f943a0a7e78202ee0e183a2d95fe8cf067cf1d4bf14c69c91b9e) |
| `origin_servers.public_name` | [origin_servers.public_name](resources--origin_pool--reference--group-002.md#canonical-2ac9f93881a630c4cac32514c0fd0313b7c46513a69187474672a6b7e2d8b5bc) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-956b384aaf91f37038cd01af534efd898eb6cd3958c1e351823f9b966fd167ee) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](resources--origin_pool--reference--group-002.md#canonical-c94bbdb691fcbe68524a58d9dfc318fbe7d53071999ca46b4b571e644510a22a) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-66c53a132a7f5352b43080f4c539c11c3f5d98ca5983171459fdfb679acfa1aa) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](resources--origin_pool--reference--group-002.md#canonical-33347e5af6084f36f080e2d54a9859d7ae8b669e1dcb1a01d18856e036c472cb) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](resources--origin_pool--reference--group-002.md#canonical-039b0ea2ea020688e310705175ee00e092efd4df2cdf28f0f2970db3e2fb5c5c) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](resources--origin_pool--reference--group-002.md#canonical-008b25925619102ff21f64000ba4a844215fc82183853ea47c318b118290eebd) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](resources--origin_pool--reference--group-002.md#canonical-2bad1689a77e515554cf53e6a29b48466e6b93b50a81f0b86357c43afe36fe18) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](resources--origin_pool--reference--group-002.md#canonical-15eb1e303a855ee36633106025d09ac6db8593ba8c13ececadef427d18edfe9f) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-5ff6ebc37855c1b55bab1719486a4cb87a505a36d4d079305c8b0d237e5df01d) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](resources--origin_pool--reference--group-002.md#canonical-c230a39ef5c4565ea779a73a220a341021d716124800b64488748640b2a60f80) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](resources--origin_pool--reference--group-002.md#canonical-5fcf92e681c0d64a45c65b72012593c9bbf458930aa5d2426b9406e602a684ec) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](resources--origin_pool--reference--group-002.md#canonical-251847ef6c4773eda3f378812c9c185b0babfed849c69511dc476fc40aee398c) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](resources--origin_pool--reference--group-002.md#canonical-c13596a18905004c6b7398499384af1f0d322c5283f8cd4c47c4236c3f6f6e30) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](resources--origin_pool--reference--group-002.md#canonical-e38c952d151a6c7f2d8b92cd202f98de9aee065c62ba4ba696ed7ac2eab26bb8) |
| `port` | [port](resources--origin_pool--reference--group-001.md#canonical-be501128f5e3d84e69b9ae443dc6dd371a24c754218e96c6650995b7ad5e7793) |
| `same_as_endpoint_port` | [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-83612ed238eaabdb14d2ae66732f0b85ad5ce12a3af85da13e3e876582fa0658) |
| `timeouts` | [timeouts](resources--origin_pool--reference--group-002.md#canonical-00b59cfc8389916908c08b416bc90e9669ff9c56fd9a0f8395f82fc5e547c57d) |
| `timeouts.create` | [timeouts.create](resources--origin_pool--reference--group-002.md#canonical-3a25e579d6cc337c1c535600bb5a5dd98645d0edb4f17d3438cb8485a978edde) |
| `timeouts.delete` | [timeouts.delete](resources--origin_pool--reference--group-002.md#canonical-5b7a596926eb1935c1123ba766458e40aad2d71e6d2e73030d7ceba3e628c57b) |
| `timeouts.read` | [timeouts.read](resources--origin_pool--reference--group-002.md#canonical-164ad8c4429ef13ac7e81f756085680fe07432b196580ebaae7d1d932fd9621b) |
| `timeouts.update` | [timeouts.update](resources--origin_pool--reference--group-002.md#canonical-14dd3f317d46bac30ac3cc2d3b8f5150fa07e802f41a3467a15bce1d35fccbd8) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-885fb983433829a1ec57ecde5e41de6772c9483c958d2f1dcf40b58091de1622) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-72fac9903e12f1d1f04b145b02857d415c3de6822142310971f0fcc9c0a62d4c) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--reference--group-002.md#canonical-914460952669ffb6a52ec16dbbdd4fe3d5431886f236e055408d05b20f63e8d1) |
| `use_tls` | [use_tls](resources--origin_pool--reference--group-002.md#canonical-da6a539edede34014702388d9e0abd5ba500bd242fc3a4a40f407fdf4ae46046) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-ccba32fca58c04f30c6ea094640d7e46ebccb976307af9bf364cf292100b52bb) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-1ff14ff691bc8669668470d527c065bcfe4584069a545f66929c609ae63a2723) |
| `use_tls.disable_sni` | [use_tls.disable_sni](resources--origin_pool--reference--group-003.md#canonical-d9a6a90c052f6d4027b4763ff9f1cba15a55f24c3aa09bcf7118abce3d56e6a1) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](resources--origin_pool--reference--group-002.md#canonical-44fb1c637240b30606d9859d2b11546d31bd40ce512bc43cbd045707a0e71d28) |
| `use_tls.no_mtls` | [use_tls.no_mtls](resources--origin_pool--reference--group-003.md#canonical-c58eebd52426320bce8e844d54523f7bdfb4078e111256756b2327a147bb666a) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-aabc3415320c442785d10159c2d66dbeb5df5ddca5b58f9f6bde8ba917846db1) |
| `use_tls.sni` | [use_tls.sni](resources--origin_pool--reference--group-002.md#canonical-68805d3a775008d51de2db1476c038b589d9fc6063322a25af75808f74a65fa0) |
| `use_tls.tls_config` | [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-48d80a573e5abff2da2251bc8116d41690f34e990bd6b2b0627dbbb7a2a1b7f0) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](resources--origin_pool--reference--group-003.md#canonical-23cc1d74ba0a67be4461bb6e5702355aa1ba5d0a4e6128af3e2f16127bcb7fde) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](resources--origin_pool--reference--group-003.md#canonical-e31e33f893cb760d94c69a584cf4704d395b8d26484c3ef982390073364e673c) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](resources--origin_pool--reference--group-003.md#canonical-1a9d483511b01fb81c22169ed7e12503ebd39d9eee7e4fd7bbe1d56cd108e786) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](resources--origin_pool--reference--group-003.md#canonical-6fd72f8fd305bd79f7a14c37c6daf4fb65dab7221b073426afe7118bff8def62) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](resources--origin_pool--reference--group-003.md#canonical-a0b2316e46d2d46972a8367a009cfcf47943e976e28bece8d1ef78c17dc5543e) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](resources--origin_pool--reference--group-003.md#canonical-b094541db67703b9df23ab03ccee1430080c2a084d9929f4a23026c70eeb56a7) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](resources--origin_pool--reference--group-003.md#canonical-e88bbf350348e3ba2b3d193987f8c144a726cefa327ca141bc74a4b348cf7254) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-c4bd4b873837edc5c2fe5b13eee5d85a6cd0d5231fa78694dd638250b3b95474) |
| `use_tls.use_mtls` | [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-45af1ade87992d9a2998f6d5fd1be5d866864b3c6370f93333249e5b3f293625) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-fab8c0218b7f95836c557b571310e6300576d4ef8640357c499f1a1398118ceb) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](resources--origin_pool--reference--group-003.md#canonical-48fff4f0e30652c53d4363ce7dee2613433b3defbafcbd7403893a56511ba1c4) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-a7ae552bc9c62cf9c0475f71145d74000d32c0e8921159d83e83c046b4175cd6) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-3951c323dbb70c7a10f00141e35bd0d2c9f717fa30f6a31bad718c18d9a7cfcf) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](resources--origin_pool--reference--group-003.md#canonical-b712e06a7926c4ea400834c0d529ccbba58f523d7888479a4708fe0edad0dff5) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-ffc2196f4ec59c808aafe69aab57d4a37b4fe4c8f2b0c6ae66718e49558f0f3e) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-b66b0564b88effdb1412803f9fa66e98fe6f99ef96c340068eb48d8e01e67294) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-793231fc84bd04d21af688065f28b33c17fd0b6bdd4917cfb3568e07c732f6e4) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--origin_pool--reference--group-003.md#canonical-87f71613a458f22636023170d80dd4e7e31ed324fc36283cef0a61c931ad7681) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](resources--origin_pool--reference--group-003.md#canonical-d2a40c00442c74541e0116d85d5ec572a7cbf3ee8c85a909a7bca28dcc06604c) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--origin_pool--reference--group-003.md#canonical-0fd46ad65c6978409bdd43e53cacbaffa63fd305e3eb6ee7ae3c3859c7c9dd22) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-957b7d05b684c6015b27ea4e91117987e62b80b3b7f710fcbd91fde0cb771d8e) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](resources--origin_pool--reference--group-003.md#canonical-0f3a90a79aec881baee92fee3b87e838da274eee2c1ad313de93926586837d0f) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](resources--origin_pool--reference--group-003.md#canonical-65f96d206e089ac3d81604af7e54ea32ba8a49c8208a5e0d84a04ca329c0d94c) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-31d81d41c7a01d5944274ca77fc0db24145fb1c7e18607e9feb33a2af328836f) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-3520e5e6634d0ff0d817cf349f31d5f4607e7cc113f149b83596a1bb40c78808) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](resources--origin_pool--reference--group-003.md#canonical-2d95cc8ea2403adddf24a1184da8ed67f4545f96d3a481e195360cbc3b337fc1) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](resources--origin_pool--reference--group-003.md#canonical-e482b2717b9434b2eceefc9fe9480552eb61cea581b8b99a2fd961746fb02e58) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](resources--origin_pool--reference--group-003.md#canonical-e0528a56de892151cbc81f6a383a37469e4c6c3711eb71466f16d0157f1723d2) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-6229777aebfda37586b438d4a060f4f7336d718149760aa4f2a091b2992fdacd) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](resources--origin_pool--reference--group-003.md#canonical-e5614161190c60fde0652313d0982a4dd055fc8f2696bf9d032e8eeb33f18bb2) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](resources--origin_pool--reference--group-003.md#canonical-b001e6cf93141e04038668b59d889c1910c7e35cd03630d5e70397dde80553b9) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](resources--origin_pool--reference--group-003.md#canonical-19956175ca098843607e4098d3d16c397dfbc1e97a52f759a9ed779e3c4e7c3d) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](resources--origin_pool--reference--group-003.md#canonical-fb45b5ebd6c0dde13bbe7f2a5a18934167d6dd64791bb758685f1df1ae6b4f84) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](resources--origin_pool--reference--group-003.md#canonical-687cb5851f4fbe47acb950c0753c5f5b82facb5e24dff2b65fb896aa35fc7d78) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-42cb4c427226d97d0b39813647ae5c7b7b6687048b45cad7cba573a986cd3b8f) |

<a id="canonical-f34b77d29203680572b34bed6f7dd57d7c534e9442bdffb2bb01607a48891a10"></a>

## Next pages — Property reference / 5edb7d60145d / 16

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-562a3acefba6f1de598f335e49c1ccd5b6fc8bf55ed03dcb7c1b8592e7bb66c5)
- [healthcheck](resources--origin_pool--reference--group-001.md#canonical-28cdf1ceb2efb9bb358019dbc9717f35a6f260c3edd34c16aa1a76490ce84607)
- [lb_port](resources--origin_pool--reference--group-001.md#canonical-08c2f998734b62d52645d8fd722945970f71788b1cc66b844a568ecb08b11050)
- [no_tls](resources--origin_pool--reference--group-001.md#canonical-ee8653d9b863ea4392eec42b70dbfc2bfef2c0f4270f8332c7945741c9026b4d)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [same_as_endpoint_port](resources--origin_pool--reference--group-002.md#canonical-d4a1eba2ebd5b3aefffcc5c49f834145c4fab929d7024ff96905398f2406158f)
- [timeouts](resources--origin_pool--reference--group-002.md#canonical-7ae7ce0a89d4f66924dedb8dfdf24bdabfc1c4ceecaf7e928bf3c2815469bcce)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-002.md#canonical-94f24e1eba33680896ec00dce0eb25ca3c9a2d809ad8bfb56c911d0bc8905f38)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74e9aba7d6bd76f83a13e6303ac9f593e335e2832d08fa6e6ab3edf80e21651b"></a>

## advanced_options — advanced_options / d95794cc6fae / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- advanced_options

<a id="canonical-069d803663c2dff549e696fe896c1e215bfb42702b749e2e0163568030200cea"></a>

Type: `"object"`. single nested block, Optional.

Configure Advanced OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_http_config",
    "http1_config"),
  validators.ConflictingObjectAttributes("auto_http_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "default_circuit_breaker"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("default_circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("disable_lb_source_ip_persistence",
    "enable_lb_source_ip_persistence"),
  validators.ConflictingObjectAttributes("disable_outlier_detection",
    "outlier_detection"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v1"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v2"),
  validators.ConflictingObjectAttributes("disable_subsets",
    "enable_subsets"),
  validators.ConflictingObjectAttributes("http1_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection"),
  validators.ConflictingObjectAttributes("no_panic_threshold",
    "panic_threshold"),
  validators.ConflictingObjectAttributes("proxy_protocol_v1",
    "proxy_protocol_v2")}
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
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-847e9b342c1ff65733070b5ffeb963e97ce35439fc04547eff6938d24f0ecea3"></a>

## Direct properties — advanced_options / d95794cc6fae / 3

- [auto_http_config](resources--origin_pool--reference--group-001.md#canonical-fd506cb3bc845b0bdb4d18cec39628640685662a661c4ad90488ac1ae4c5a532): complete subsection reference.

- [circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-0f5c40b26ffd883486afb92af7cad06cf156e49af54e0e3f3c9cb533178e6208): complete subsection reference.

<a id="canonical-7ec75d3722e6a32356d556c4bba2276b3bfe62c023e4d939c43b468ee6d965d2"></a>

<a id="canonical-55aae58058699ec17137a0616600e2394d5fa0657778b9c668c1f08d4d47fa79"></a>

## connection_timeout property — advanced_options / d95794cc6fae / 4

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-d1604207fc5e9081a8810b785f1eb9df0a682dbdf3b4c91e0779e800269f7ff5): complete subsection reference.

- [disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-408d237246c3693378d50046c3a2f3c129bae4506c95bf90f05cf849da72daa5): complete subsection reference.

- [disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-543f4569354cccebe5ce4a1bcf1587d08d4b7874b2a7c1e6a5efdc97bc2c5d6c): complete subsection reference.

- [disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-3354d67cccb8501cc0362e64f78351b82fa5bc9fe0e427faeaba705af18858dd): complete subsection reference.

- [disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-a7c0fbfde308686b48c0e75ad48693586f3f547e40ddd2e99da0201f80c6be9d): complete subsection reference.

- [disable_subsets](resources--origin_pool--reference--group-001.md#canonical-1e4a46b5f72d6455877040730db9353cc4852d228e3f8924335e7724f1be696c): complete subsection reference.

- [enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-030e9bc674ac9cd8a51df892c1eae1981bc6dcc48aad7aa1f58401542c2b92ec): complete subsection reference.

- [enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953): complete subsection reference.

- [http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140): complete subsection reference.

- [http2_options](resources--origin_pool--reference--group-001.md#canonical-d4a7cbb40f338c150efd9c1aa32a259d82650fcadc56fad64c1aee3b1a158e18): complete subsection reference.

<a id="canonical-79121eaa5c7c1865d913abc4fce7abb9df84433dcfa5fd5af686d092a7f1dbf6"></a>

<a id="canonical-0dc77defa4700ca0caf6fb0766deedf47a3e6eab3130e583399cdeece41e4a68"></a>

## http_idle_timeout property — advanced_options / d95794cc6fae / 5

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

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

<a id="canonical-1cb726dd77c9c0b258d6dc205ad6893eab0929a01dcf7625f53d1e4091154465"></a>

<a id="canonical-9720c22847e08b4b35466049e7237e47a61d6f983801f3e473b774f4c8c16c15"></a>

## max_requests_per_connection property — advanced_options / d95794cc6fae / 6

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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

- [no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-225988483612c01d5033a2edee9b550e6893aa93bea70072d43b2421241fca7c): complete subsection reference.

- [no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-7c9ffdd4d03712899b720dfee11ea21dab3199cca862377d69440dc7564c3017): complete subsection reference.

- [outlier_detection](resources--origin_pool--reference--group-001.md#canonical-c0ba35094db276855bfcba569678fa692f571ea65a98c4dc0174455fdfa1b4da): complete subsection reference.

<a id="canonical-16f4cb0ee11ca3be584c34687ba431396c58c823d9fe4c5d88a2e8e1194cd7d8"></a>

<a id="canonical-28ba8fd2bb3dc6ab714f5a711313a3274d9e2a01a2871477b2700367a3a75429"></a>

## panic_threshold property — advanced_options / d95794cc6fae / 7

Type: `"number"`. Optional.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-adc87d7eb2230911c724e6307a23b54149da57db0192b566a4a73f1b9ecbc346): complete subsection reference.

- [proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-38347ca8fe427a759b8b14be843c18a94ef87875e55a3d45c19d7d67b87f2505): complete subsection reference.

<a id="canonical-0fbe8d5527e90bcf2fcdf4ef448dbd524ec9b0e5bc909b762bbf0562c161c9bb"></a>

## Next pages — advanced_options / d95794cc6fae / 8

- [advanced_options.auto_http_config](resources--origin_pool--reference--group-001.md#canonical-fd506cb3bc845b0bdb4d18cec39628640685662a661c4ad90488ac1ae4c5a532)
- [advanced_options.circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-0f5c40b26ffd883486afb92af7cad06cf156e49af54e0e3f3c9cb533178e6208)
- [advanced_options.default_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-d1604207fc5e9081a8810b785f1eb9df0a682dbdf3b4c91e0779e800269f7ff5)
- [advanced_options.disable_circuit_breaker](resources--origin_pool--reference--group-001.md#canonical-408d237246c3693378d50046c3a2f3c129bae4506c95bf90f05cf849da72daa5)
- [advanced_options.disable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-543f4569354cccebe5ce4a1bcf1587d08d4b7874b2a7c1e6a5efdc97bc2c5d6c)
- [advanced_options.disable_outlier_detection](resources--origin_pool--reference--group-001.md#canonical-3354d67cccb8501cc0362e64f78351b82fa5bc9fe0e427faeaba705af18858dd)
- [advanced_options.disable_proxy_protocol](resources--origin_pool--reference--group-001.md#canonical-a7c0fbfde308686b48c0e75ad48693586f3f547e40ddd2e99da0201f80c6be9d)
- [advanced_options.disable_subsets](resources--origin_pool--reference--group-001.md#canonical-1e4a46b5f72d6455877040730db9353cc4852d228e3f8924335e7724f1be696c)
- [advanced_options.enable_lb_source_ip_persistence](resources--origin_pool--reference--group-001.md#canonical-030e9bc674ac9cd8a51df892c1eae1981bc6dcc48aad7aa1f58401542c2b92ec)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- [advanced_options.http2_options](resources--origin_pool--reference--group-001.md#canonical-d4a7cbb40f338c150efd9c1aa32a259d82650fcadc56fad64c1aee3b1a158e18)
- [advanced_options.no_panic_threshold](resources--origin_pool--reference--group-001.md#canonical-225988483612c01d5033a2edee9b550e6893aa93bea70072d43b2421241fca7c)
- [advanced_options.no_request_limit_per_connection](resources--origin_pool--reference--group-001.md#canonical-7c9ffdd4d03712899b720dfee11ea21dab3199cca862377d69440dc7564c3017)
- [advanced_options.outlier_detection](resources--origin_pool--reference--group-001.md#canonical-c0ba35094db276855bfcba569678fa692f571ea65a98c4dc0174455fdfa1b4da)
- [advanced_options.proxy_protocol_v1](resources--origin_pool--reference--group-001.md#canonical-adc87d7eb2230911c724e6307a23b54149da57db0192b566a4a73f1b9ecbc346)
- [advanced_options.proxy_protocol_v2](resources--origin_pool--reference--group-001.md#canonical-38347ca8fe427a759b8b14be843c18a94ef87875e55a3d45c19d7d67b87f2505)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-fd506cb3bc845b0bdb4d18cec39628640685662a661c4ad90488ac1ae4c5a532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f9e05f4261b698fe334ad6cb4f6fa71672f690397157997ba30ca29702d23af"></a>

## advanced_options.auto_http_config — advanced_options.auto_http_config / b5e6eb6cd716 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.auto_http_config

<a id="canonical-dff138ba4b4a18b9f751174eacd46415025bec6a1349938d2b7e6ecb15959f4d"></a>

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
auto_http_config = {}
```

<a id="canonical-c9bcb1fa48f3059a26e1675be330e861bd8745dfde6f71efcda1eecb90a56e24"></a>

## Direct properties — advanced_options.auto_http_config / b5e6eb6cd716 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba9c4ca0cbe89f5b154dce31d742a8277d26bac44ed2621ac58d675742fc8b49"></a>

## Next pages — advanced_options.auto_http_config / b5e6eb6cd716 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-0f5c40b26ffd883486afb92af7cad06cf156e49af54e0e3f3c9cb533178e6208"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e10b41cd6cc86452340efa426e3f18c5532c5da963beb1f6369e35fc39be512"></a>

## advanced_options.circuit_breaker — advanced_options.circuit_breaker / 5ce3d1fa9166 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.circuit_breaker

<a id="canonical-d96543fb7ca049b91c32ada1c3bddad6508e24c7c473b1fcc5f480f6d829bbaf"></a>

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

<a id="canonical-1af4d76d5d409ad9a493400ff37089a15c9a904afaf53988db298834446bf38d"></a>

## Direct properties — advanced_options.circuit_breaker / 5ce3d1fa9166 / 3

<a id="canonical-92dc47ff68a57a711dd8c00052fe46d28ae3d8ae05a79f17083a3956287d6f97"></a>

<a id="canonical-3f9b87e32a69c48c22a2ebb08a3a0b07ddfcb5d70e3217f869cffd2448f1f37d"></a>

## connection_limit property — advanced_options.circuit_breaker / 5ce3d1fa9166 / 4

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

<a id="canonical-cdc4e7ec612d321352a5850ee0ac1cf00aa8e20bd89cafb92c2400154a40a319"></a>

<a id="canonical-b81bac081213368a2e76af5e03c0e4b2fea806c4d1238a46d66d4d0c151febeb"></a>

## max_requests property — advanced_options.circuit_breaker / 5ce3d1fa9166 / 5

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

<a id="canonical-778da4b2f71ab6c5859f29a665aaaffa0598713dc6f7aac9075b7ec35d2e320e"></a>

<a id="canonical-b3ecb942fb50368115ed1980a0985e6645a293e4856495832bf6d1ff7d9f0fd3"></a>

## pending_requests property — advanced_options.circuit_breaker / 5ce3d1fa9166 / 6

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

<a id="canonical-8aa7150541ba59c8bcbbec14de4fd30296a4767dda6b72d3cebfffd7a9ee0258"></a>

<a id="canonical-e31299f04d7a9d180d310bfd485853a4033e07f02db796c9da0e6a5066febdff"></a>

## priority property — advanced_options.circuit_breaker / 5ce3d1fa9166 / 7

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

<a id="canonical-b11f1cce184dc3b78cab91b475eb0c028a20331a47fdb95c7ca2628c82b4921a"></a>

<a id="canonical-af3c764d6023c6a3f7741a827e7f7f384738b47c859375f9115e6542190452e0"></a>

## retries property — advanced_options.circuit_breaker / 5ce3d1fa9166 / 8

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

<a id="canonical-f6ed3315fa70e2203bc0b3df149cc57b51877bf2b6ca78663ac625376e180276"></a>

## Next pages — advanced_options.circuit_breaker / 5ce3d1fa9166 / 9

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d1604207fc5e9081a8810b785f1eb9df0a682dbdf3b4c91e0779e800269f7ff5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51f41bfaede1734a5ccb7d3ea0ff67a8d23175d60346836c5c74d0d69c5c48df"></a>

## advanced_options.default_circuit_breaker — advanced_options.default_circuit_breaker / 3bf3fd897fd0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.default_circuit_breaker

<a id="canonical-b5bf4681807d97cd1ccc3bf302884f3118c343d2c661de52bd1e76f27fe1bcd9"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
default_circuit_breaker = {}
```

<a id="canonical-b710de92bafab791a46117e59e5dd6c91fcf367e82b9d4538e37c349242c9f4b"></a>

## Direct properties — advanced_options.default_circuit_breaker / 3bf3fd897fd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5517c0a7be9265d58ceef7157778704fef6427667d546cee5ee5d4e7ca2f8a20"></a>

## Next pages — advanced_options.default_circuit_breaker / 3bf3fd897fd0 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-408d237246c3693378d50046c3a2f3c129bae4506c95bf90f05cf849da72daa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f425641c2c70331ed2a9ecbe5d05323e49fdd1da81cd336da69eeb2bae00875"></a>

## advanced_options.disable_circuit_breaker — advanced_options.disable_circuit_breaker / 217803333c7b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.disable_circuit_breaker

<a id="canonical-61ae33c0268245c29b18a4e27bd4c394c0fd514a4b871ec7527225e338d994c3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable circuit breaker.

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
disable_circuit_breaker = {}
```

<a id="canonical-6b1f99f08c21679ef8ddece796f3e06cbe0973f3970b414fba86c3b95a88bd32"></a>

## Direct properties — advanced_options.disable_circuit_breaker / 217803333c7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6cec7051915bf2f1df31561c5bfc936d9c4042828f3461b51e555895a05be680"></a>

## Next pages — advanced_options.disable_circuit_breaker / 217803333c7b / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-543f4569354cccebe5ce4a1bcf1587d08d4b7874b2a7c1e6a5efdc97bc2c5d6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fcaf2d20d4d4ca6e706f9a2c2a3743cdd39985256793f30d9ba419072f6d33b"></a>

## advanced_options.disable_lb_source_ip_persistence — advanced_options.disable_lb_source_ip_persistence / 1c36f470033c / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.disable_lb_source_ip_persistence

<a id="canonical-f69735a59a089253569aecbd9e57d28762147af5f2026c261e5d1f0419f59bb6"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
disable_lb_source_ip_persistence = {}
```

<a id="canonical-e1276bc53b61490e9bbf53ecd7c93b26fe7cc58c73c97af6dfd6df85be81817e"></a>

## Direct properties — advanced_options.disable_lb_source_ip_persistence / 1c36f470033c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8aa07c39af3c54a359d6fb9ecbb17ba91fe4638f17bcf3c0e024f414e97a8077"></a>

## Next pages — advanced_options.disable_lb_source_ip_persistence / 1c36f470033c / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-3354d67cccb8501cc0362e64f78351b82fa5bc9fe0e427faeaba705af18858dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f4b22f062e70cb631d8b50936e0f869968709f30e6b39895c1ccaad13d9284e"></a>

## advanced_options.disable_outlier_detection — advanced_options.disable_outlier_detection / 8d9c2ec62786 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.disable_outlier_detection

<a id="canonical-a687e4d6b20484c75edf80ef309f69b61f8a95069ce1ab8736887ee7c2eebbd1"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
disable_outlier_detection = {}
```

<a id="canonical-c2b083a7d11d9041141350ac9d93a45238207d41d60aa23f2652b0bdf6fa6bfc"></a>

## Direct properties — advanced_options.disable_outlier_detection / 8d9c2ec62786 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e0146db3a29891328a3792dd68e2137271a7cf67a400f6cb8fb29da55423448"></a>

## Next pages — advanced_options.disable_outlier_detection / 8d9c2ec62786 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a7c0fbfde308686b48c0e75ad48693586f3f547e40ddd2e99da0201f80c6be9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22e3efafc63bd006f42d130c84de19eb24063633215916cc72524780ed7202bf"></a>

## advanced_options.disable_proxy_protocol — advanced_options.disable_proxy_protocol / 9f368077ffde / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.disable_proxy_protocol

<a id="canonical-6b9e1532e32de91ef9099c0dda8d7c878c59a89de20d3f053cd5f39f44f6933f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable proxy protocol.

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
disable_proxy_protocol = {}
```

<a id="canonical-977cc3b269dc26de419bbdec4659940e8ddb85d6e1a77cd96aa3142bb458ebdc"></a>

## Direct properties — advanced_options.disable_proxy_protocol / 9f368077ffde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3777b24b978fa537871b7855c9ee11c54d06f5ce3facb0e258ad9fe1885c26c3"></a>

## Next pages — advanced_options.disable_proxy_protocol / 9f368077ffde / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1e4a46b5f72d6455877040730db9353cc4852d228e3f8924335e7724f1be696c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cae537d91788816a8151e664cd98dd9835ef89afdbd681cb7096dc0dd933e504"></a>

## advanced_options.disable_subsets — advanced_options.disable_subsets / 6449414ef460 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.disable_subsets

<a id="canonical-a90990e214edae3b36360e951f89226f6102f0d3ee54f01da37bedba49dae827"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_subsets = {}
```

<a id="canonical-c2ee09b05513153453a5d2d1ea31d85d73c00d3022bf0f44ae0714d1ca28e500"></a>

## Direct properties — advanced_options.disable_subsets / 6449414ef460 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0302603cc987568ae6a3f91e74051818939d8e9672a0d2a1b877a2c5f554d83"></a>

## Next pages — advanced_options.disable_subsets / 6449414ef460 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-030e9bc674ac9cd8a51df892c1eae1981bc6dcc48aad7aa1f58401542c2b92ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7b56b861dd8bbf4543c50334917de1a047b8ef91734f6c3521e315f5e7006a3"></a>

## advanced_options.enable_lb_source_ip_persistence — advanced_options.enable_lb_source_ip_persistence / 8db1ff00eb17 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.enable_lb_source_ip_persistence

<a id="canonical-f05c13a51338eccfb51f68ffc0e9d68c8b71cce299b62737da16d9a9eafbc8c7"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
enable_lb_source_ip_persistence = {}
```

<a id="canonical-185764317d167ce39180124943fb6bcc9e7b2c938fe471762183ddf8b9089bdb"></a>

## Direct properties — advanced_options.enable_lb_source_ip_persistence / 8db1ff00eb17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76964852347dde8f87e91af5f200e80f81cd04d167d3546ab60bbb87670c07cb"></a>

## Next pages — advanced_options.enable_lb_source_ip_persistence / 8db1ff00eb17 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d2f1d42deba3c27059d28ac405fbd19503cda6c0dc000e82d42d88d03322243"></a>

## advanced_options.enable_subsets — advanced_options.enable_subsets / 39da7819fb42 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.enable_subsets

<a id="canonical-1c9d9c9fb2d77bc2df160529b6359f2d89bcd1153c486d7b98c906800bab4b9a"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint_subsets"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "default_subset"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "fail_request"),
  validators.ConflictingObjectAttributes("default_subset",
    "fail_request")}
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
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

Terraform syntax:

```terraform
enable_subsets {
  # Configure direct properties listed below.
}
```

<a id="canonical-b2dbf738c00d21e6cc3b27d4fd19fc52b0124a84b9b4641cf1aa8af0549e1dad"></a>

## Direct properties — advanced_options.enable_subsets / 39da7819fb42 / 3

- [any_endpoint](resources--origin_pool--reference--group-001.md#canonical-59ce3c6f3c74a779bd9dcc46d1fcc308f19b45c8cb9cf60fc56364d82a609968): complete subsection reference.

- [default_subset](resources--origin_pool--reference--group-001.md#canonical-794cd9c735ac9b0fe665c62a02994ad9d4676381ddcab7593e0398b0c810c1ff): complete subsection reference.

- [endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-ff96b79f5eb8465dfba82667cfb3263aad7b727fdcfed41cc20e83f36bc0e499): complete subsection reference.

- [fail_request](resources--origin_pool--reference--group-001.md#canonical-a811096b80d019cf13d705bcb5c63ed28ec4c429986cd8ba624aea44dbe9b57c): complete subsection reference.

<a id="canonical-97b17b6772b00a0e5849033a3c9308f4f72bce80c30bcccf855b04bfbdac22ec"></a>

## Next pages — advanced_options.enable_subsets / 39da7819fb42 / 4

- [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--reference--group-001.md#canonical-59ce3c6f3c74a779bd9dcc46d1fcc308f19b45c8cb9cf60fc56364d82a609968)
- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-794cd9c735ac9b0fe665c62a02994ad9d4676381ddcab7593e0398b0c810c1ff)
- [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--reference--group-001.md#canonical-ff96b79f5eb8465dfba82667cfb3263aad7b727fdcfed41cc20e83f36bc0e499)
- [advanced_options.enable_subsets.fail_request](resources--origin_pool--reference--group-001.md#canonical-a811096b80d019cf13d705bcb5c63ed28ec4c429986cd8ba624aea44dbe9b57c)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-59ce3c6f3c74a779bd9dcc46d1fcc308f19b45c8cb9cf60fc56364d82a609968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a7ebfbfc1364f5533e8f8850deb0e710fba9fd80fc74fb7af9e06e147fb0cce"></a>

## advanced_options.enable_subsets.any_endpoint — advanced_options.enable_subsets.any_endpoint / 96ab144e462b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- advanced_options.enable_subsets.any_endpoint

<a id="canonical-164d18bbc9dedbe728f4ddf8124bc4c20bdcbee7c83573b7883e6579e46f6861"></a>

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
any_endpoint = {}
```

<a id="canonical-d5233aeda409c0f6768fbfb533ee3a330ba0bbbac5c1748ad3bf858c0b9bfaf1"></a>

## Direct properties — advanced_options.enable_subsets.any_endpoint / 96ab144e462b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-634e6606a1574a61421fd9dd3a4f1a796f99942b12996dcab8fc36d0b0340444"></a>

## Next pages — advanced_options.enable_subsets.any_endpoint / 96ab144e462b / 4

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-794cd9c735ac9b0fe665c62a02994ad9d4676381ddcab7593e0398b0c810c1ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1058f52d2e9a4ee700a08baa0c13ce8fc5c405d1338051d1485b9d50ac6a6c1"></a>

## advanced_options.enable_subsets.default_subset — advanced_options.enable_subsets.default_subset / cd0f65402835 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- advanced_options.enable_subsets.default_subset

<a id="canonical-d94349b6f66afcb94183c30aeef2bf502e90648822f75ed7595deb69e2130d4b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

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
default_subset {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5b58f63326ef27090db07013cbc12fe1351285d2ee62c3d292314bb2e0529cb"></a>

## Direct properties — advanced_options.enable_subsets.default_subset / cd0f65402835 / 3

- [default_subset](resources--origin_pool--reference--group-001.md#canonical-e18670b1e3d39655159698ae36e53ddc6a7adba0a2097cb68f12a3446e179f82): complete subsection reference.

<a id="canonical-e0ca6c2ca7ad02d1970b9b7af9e7b9d250e0d1a0656d1ef77113c5833c2fc712"></a>

## Next pages — advanced_options.enable_subsets.default_subset / cd0f65402835 / 4

- [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--reference--group-001.md#canonical-e18670b1e3d39655159698ae36e53ddc6a7adba0a2097cb68f12a3446e179f82)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-e18670b1e3d39655159698ae36e53ddc6a7adba0a2097cb68f12a3446e179f82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3a1062a3ecbfa8bf6d3daca8d74a6815b1801bea6e3be126243665663982249"></a>

## advanced_options.enable_subsets.default_subset.default_subset — advanced_options.enable_subsets.default_subset.default_subset / e7840648ff54 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-794cd9c735ac9b0fe665c62a02994ad9d4676381ddcab7593e0398b0c810c1ff)
- advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-109cfc5f2f1b9fe16b37024a667d88b32acb0b4e1428e9a271e0167688a4588b"></a>

Type: `"object"`. single nested block, Optional.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

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

<a id="canonical-c19b3cf2970bd68b1db3c98751cf49ad5151eb829c4ca2e596e8877201e6f775"></a>

## Direct properties — advanced_options.enable_subsets.default_subset.default_subset / e7840648ff54 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f45404d37d12e9ab7eb5d2be40dfb0b228c5982856b30b640b2a049e7259ce9c"></a>

## Next pages — advanced_options.enable_subsets.default_subset.default_subset / e7840648ff54 / 4

- [advanced_options.enable_subsets.default_subset](resources--origin_pool--reference--group-001.md#canonical-794cd9c735ac9b0fe665c62a02994ad9d4676381ddcab7593e0398b0c810c1ff)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-ff96b79f5eb8465dfba82667cfb3263aad7b727fdcfed41cc20e83f36bc0e499"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87bdec1dc2fffa0ecef17c2e42f9ff2c5e8030e9b370151f77a1789ed65657d9"></a>

## advanced_options.enable_subsets.endpoint_subsets — advanced_options.enable_subsets.endpoint_subsets / 8d9758c97d81 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-c34dc32e2c00d1a77135106fa451810babf4c625596dfc1b3ba6c14b991808b6"></a>

Type: `"object"`. list nested block, Optional.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-26481e25ef51cf97c8a1c458717d5f634ab8cfc035a43ff8edc055545aa001c9"></a>

## Direct properties — advanced_options.enable_subsets.endpoint_subsets / 8d9758c97d81 / 3

<a id="canonical-61c56b0675d48a41be3d1ced33484227d68f443e7342e6efaedb29effacfcb87"></a>

<a id="canonical-db4a80204d26e83c481b370cc07ebd926ba638c780df178ab5006f38f03db26d"></a>

## keys property — advanced_options.enable_subsets.endpoint_subsets / 8d9758c97d81 / 4

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

<a id="canonical-487b4d4440fe28bfac11f90646cffcfb3cf8255291074778f1bf4e2a9c775fc4"></a>

## Next pages — advanced_options.enable_subsets.endpoint_subsets / 8d9758c97d81 / 5

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a811096b80d019cf13d705bcb5c63ed28ec4c429986cd8ba624aea44dbe9b57c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01e0735c985729783ed6fccb8ab9e11767411fae1edff0284167e3c8a39fb0fb"></a>

## advanced_options.enable_subsets.fail_request — advanced_options.enable_subsets.fail_request / dcfc098e3c81 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- advanced_options.enable_subsets.fail_request

<a id="canonical-d9aaf050cd7efb283bcf612f0475e378cfa8cd1c9b49d9aec809f859a353caa5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fail request.

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
fail_request = {}
```

<a id="canonical-6eca75068b828a35315582e8b90666b3fc2979085d6d274cb739641a61cddf03"></a>

## Direct properties — advanced_options.enable_subsets.fail_request / dcfc098e3c81 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8317cc25ac37fdcdf569ab584382966ca3710bb75f01eb6ab7be765bd8838772"></a>

## Next pages — advanced_options.enable_subsets.fail_request / dcfc098e3c81 / 4

- [advanced_options.enable_subsets](resources--origin_pool--reference--group-001.md#canonical-1a4021ed818e9aedd4eb928379933fd2363464d2e756568dc1fe4458d3fa7953)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0276788982e1e899256e97565aa466ea0c7e017f17aab6f2628c0afde577537"></a>

## advanced_options.http1_config — advanced_options.http1_config / 60ea83f45734 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.http1_config

<a id="canonical-9e2071cfbb46fd77e5cf7698d2c8a43f0a88a5e286385e6db2f136b20a39d6ea"></a>

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

<a id="canonical-4d01c4c56d5a047a9ebfb37c124b65fba93ca277c7f7b7543c60190ca1b06d03"></a>

## Direct properties — advanced_options.http1_config / 60ea83f45734 / 3

- [header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7): complete subsection reference.

<a id="canonical-8c5a3e2db4a8e13fb303a4b87b91714b699c2919ec43bc57fd4b4e6df51f49c2"></a>

## Next pages — advanced_options.http1_config / 60ea83f45734 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bc9c5012bbf7811b4554f470b08c2b04ecaa4572a4c0d302c6993f6037e95c2"></a>

## advanced_options.http1_config.header_transformation — advanced_options.http1_config.header_transformation / 715b0a3e13bb / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- advanced_options.http1_config.header_transformation

<a id="canonical-d93e7ffed66ef55089937ffea6c3f8d0c816fc35011e5cdbd701eb9b052f9cd1"></a>

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

<a id="canonical-abc3ef08b709d8edd7a1ef7d889081de4a011e368e1769af03ed05e58d2ff98a"></a>

## Direct properties — advanced_options.http1_config.header_transformation / 715b0a3e13bb / 3

- [default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0046c714880df0ccf689f0ed0d6b96e1441524dcd386fad04b43572b030189fa): complete subsection reference.

- [preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-3c2435697e6b08f41b0060db2b6646aced9518eeb6530f169b058c4bcc4b06e8): complete subsection reference.

- [proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-b95715a1c150591a35df9f8da4d1da5cdbef72956d1a69a0badf2d92ad157015): complete subsection reference.

<a id="canonical-84a992b15aeff59e27d02005a0f8d5b08d89870cc7f24e2bf87eb17a1980e338"></a>

## Next pages — advanced_options.http1_config.header_transformation / 715b0a3e13bb / 4

- [advanced_options.http1_config.header_transformation.default_header_transformation](resources--origin_pool--reference--group-001.md#canonical-0046c714880df0ccf689f0ed0d6b96e1441524dcd386fad04b43572b030189fa)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-3c2435697e6b08f41b0060db2b6646aced9518eeb6530f169b058c4bcc4b06e8)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](resources--origin_pool--reference--group-001.md#canonical-b95715a1c150591a35df9f8da4d1da5cdbef72956d1a69a0badf2d92ad157015)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-0046c714880df0ccf689f0ed0d6b96e1441524dcd386fad04b43572b030189fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80ecdba632c0ed0050310b2ac976cde94d855e0ee037c3a78ee57e8616a09e72"></a>

## advanced_options.http1_config.header_transformation.default_header_transformation — advanced_options.http1_config.header_transformation.default_header_transformatio / 1c6f248a52bb / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-4cdf1882f711fca1a6a63c42b653d11194f484a88f479713a1303fc29ebc5f82"></a>

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

<a id="canonical-8dc800fe21c54b79cbfa1029e11378961dbab8a18c22f1983c5fd81e091b945b"></a>

## Direct properties — advanced_options.http1_config.header_transformation.default_header_transformatio / 1c6f248a52bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8eb7946ad11357a288e0194bd090ad5fc2417192a3efd9c39087bb6b21c1fdca"></a>

## Next pages — advanced_options.http1_config.header_transformation.default_header_transformatio / 1c6f248a52bb / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-3c2435697e6b08f41b0060db2b6646aced9518eeb6530f169b058c4bcc4b06e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0329c2013c0d45a72c644e930e88cd92a544f77a7800914c7d826e1bd2d1db"></a>

## advanced_options.http1_config.header_transformation.preserve_case_header_transformation — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / db47854eb8b2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-9726d3fc6ab1cfd7eef5af0335424b456a420548efc971dbe062e834c0b5dc60"></a>

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

<a id="canonical-1fe31d64fcabe3a5d14ab64e4d221657a2824808e15cabde18ba3bbd1091d30b"></a>

## Direct properties — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / db47854eb8b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04187ace911754f7fb38e065f49316ffde2e2c91ed59763a1cca8ffbf0e893e1"></a>

## Next pages — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / db47854eb8b2 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-b95715a1c150591a35df9f8da4d1da5cdbef72956d1a69a0badf2d92ad157015"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a03f15bd3af7b7f5a851a3db1ede77090f3be9269474224b4469bbd94d89b39"></a>

## advanced_options.http1_config.header_transformation.proper_case_header_transformation — advanced_options.http1_config.header_transformation.proper_case_header_transform / ec42f564ecc5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [advanced_options.http1_config](resources--origin_pool--reference--group-001.md#canonical-587bc94f06e54d0405256c02cc02618abfc41bfe8ab08f6a175a53347afa1140)
- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-233f160e72d7d8191fde0b491ddbebb0d64a50383411e0f638e5e216db39cb67"></a>

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

<a id="canonical-5cc32d39251ee3ad7dac0283094ab4d72e3bb100c045160297bd90299f95bb16"></a>

## Direct properties — advanced_options.http1_config.header_transformation.proper_case_header_transform / ec42f564ecc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7716830691f71df31ba5f2d2b6e032c20b42c529d626a2414bd221649c0bdf3"></a>

## Next pages — advanced_options.http1_config.header_transformation.proper_case_header_transform / ec42f564ecc5 / 4

- [advanced_options.http1_config.header_transformation](resources--origin_pool--reference--group-001.md#canonical-b336bff42e77de7aebb3fbd6dd9b3ec6d7bab1f944db9e6f0edd97554f63cdd7)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d4a7cbb40f338c150efd9c1aa32a259d82650fcadc56fad64c1aee3b1a158e18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-746e290c9b89bf1bb81e5b50b9873dd2b46965e6e03b864246ebad2d71ccad48"></a>

## advanced_options.http2_options — advanced_options.http2_options / 271ba61f357a / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.http2_options

<a id="canonical-abd94c4e89f169f6cec459cb7436b52e898f357c06f95b513269f553820802b1"></a>

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

<a id="canonical-1a7371d51c7412450928a74905b914a4e00f8daf5e44d2c6a4f63324086ca56f"></a>

## Direct properties — advanced_options.http2_options / 271ba61f357a / 3

<a id="canonical-013beaac6e4990222b79607f317dd79bf237d85766c5401b10d336b9282a2d26"></a>

<a id="canonical-cd4fc68b39df63eb95e2a4bad9f4935e9faa446ef16286e91dbc2c6a28c95fad"></a>

## enabled property — advanced_options.http2_options / 271ba61f357a / 4

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

<a id="canonical-1fd8aede731b5d08f24a1a921952a35143758c7a91b84ab0958ef4b086e5e69e"></a>

## Next pages — advanced_options.http2_options / 271ba61f357a / 5

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-225988483612c01d5033a2edee9b550e6893aa93bea70072d43b2421241fca7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-911d922b8be90bd528bb2b5b26f2f669ee9478214ab34f07a3418180d7d66788"></a>

## advanced_options.no_panic_threshold — advanced_options.no_panic_threshold / dfe779a5a9ef / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.no_panic_threshold

<a id="canonical-ca955333f3916b4ec9eaaf6716904060f1b22bac7fa59355f420305be93e8289"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no panic threshold. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_panic_threshold = {}
```

<a id="canonical-9486d78eb93e9519d86627ca7ab8dd1195091af7853fa764c0bb1f88ffa23d33"></a>

## Direct properties — advanced_options.no_panic_threshold / dfe779a5a9ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efc71949f1bea9bdbd372b2894ac5db88992fd7fdc25c85020f3493a8daa64fd"></a>

## Next pages — advanced_options.no_panic_threshold / dfe779a5a9ef / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-7c9ffdd4d03712899b720dfee11ea21dab3199cca862377d69440dc7564c3017"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06ed7a15456f398d977424af87b1fe3520945fbf209778f15b22904917393334"></a>

## advanced_options.no_request_limit_per_connection — advanced_options.no_request_limit_per_connection / e5b7d5bc9c29 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.no_request_limit_per_connection

<a id="canonical-3fe443dcf8df2c689a15077fe76fea3dae05616858f0f6fe32bb318d6f17c946"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-d792b9f4e7c669d6afa188d29f84a325beac229ceaa06de6e1f81618901ab646"></a>

## Direct properties — advanced_options.no_request_limit_per_connection / e5b7d5bc9c29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c36da9964dbcbded90c01bd29360b8262c1c235e67e4e308fc1eaa53a5468221"></a>

## Next pages — advanced_options.no_request_limit_per_connection / e5b7d5bc9c29 / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-c0ba35094db276855bfcba569678fa692f571ea65a98c4dc0174455fdfa1b4da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7ea026e22bf62cb25f2e4a4ad466ca11c8c897757c91e0b9b0b315537ff81d0"></a>

## advanced_options.outlier_detection — advanced_options.outlier_detection / 0324f1ed68aa / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.outlier_detection

<a id="canonical-f3022a07cdb8faf479ec3e61440229376295b1ca336dac85f009d12b509a5668"></a>

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

<a id="canonical-496c1919e0b8157ecf6dbb8649aa807f8048fa9434ee6403fe0a2947ed9e6efd"></a>

## Direct properties — advanced_options.outlier_detection / 0324f1ed68aa / 3

<a id="canonical-ea1eb8a6c9f435765a22222f6bd25277ff7be74a5530cf5ab7821f89f9eee1fc"></a>

<a id="canonical-0bc432f30b70ced117e7203394727a4e6231f99f3b9b5692ea7b7958b46a917f"></a>

## base_ejection_time property — advanced_options.outlier_detection / 0324f1ed68aa / 4

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

<a id="canonical-a3e66ecb4fabe054a6d1c5f8c8ec08caf4ecf1cd327ab533b7e75dfb13e9f957"></a>

<a id="canonical-dd61d405b9d297e1405849617271ee8db523170069987377bbc64034f21643be"></a>

## consecutive_5xx property — advanced_options.outlier_detection / 0324f1ed68aa / 5

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

<a id="canonical-8e91cacb924765ddce8466549fc27dd6cb0698a7c7c02d0d1fed0055fbdd54e6"></a>

<a id="canonical-0835c1b9e5148c10c4dc9ca5267978d7f50d7247e65bf3a76c15c1671ff4b72b"></a>

## consecutive_gateway_failure property — advanced_options.outlier_detection / 0324f1ed68aa / 6

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

<a id="canonical-9c95adf3d0db128c97bdf6d3b9e72ad9fc6586898dab5bc927bf85ba4906e46b"></a>

<a id="canonical-cfc9c0892f7bbfcf20d3f934f179674fe35d0fe39e3da419c9f20e1f696b6d47"></a>

## interval property — advanced_options.outlier_detection / 0324f1ed68aa / 7

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

<a id="canonical-9a67ebcd546ee902c655da1d05023f47b2fdaaa2d6ebddc9b4a40f7064c7bcd6"></a>

<a id="canonical-936693e5438f27f639699b6d0c49ff074ec71b28fdaed25ea59236ba012bc59c"></a>

## max_ejection_percent property — advanced_options.outlier_detection / 0324f1ed68aa / 8

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

<a id="canonical-3c3d706a7ccd4184a96dbb02b522f6528b096f112da754d0a8ea7f56d24c7167"></a>

## Next pages — advanced_options.outlier_detection / 0324f1ed68aa / 9

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-adc87d7eb2230911c724e6307a23b54149da57db0192b566a4a73f1b9ecbc346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3c8ad00e4f66f74fde0a313907a0e6058c9c8a383857b3a9fe3142e8ceb06f6"></a>

## advanced_options.proxy_protocol_v1 — advanced_options.proxy_protocol_v1 / f0b16fb9bf6d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.proxy_protocol_v1

<a id="canonical-5610b07c5be24333551d364fbf5e0a1eb9e533a72cb042e66ab2bbaf7231888f"></a>

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

<a id="canonical-0c134022fa25b9ea7774c9a6b1ffa3419e6b24b9b078c78c8a83bc86ba24853f"></a>

## Direct properties — advanced_options.proxy_protocol_v1 / f0b16fb9bf6d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78ec5af45d3a96d22f07d4f8bbcb3fc026993e7938d1292b442035f5d1fe11be"></a>

## Next pages — advanced_options.proxy_protocol_v1 / f0b16fb9bf6d / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-38347ca8fe427a759b8b14be843c18a94ef87875e55a3d45c19d7d67b87f2505"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edf8b064636af37f511a36893eb126620735277a619cbd2a3a06f4dd61e9c8d7"></a>

## advanced_options.proxy_protocol_v2 — advanced_options.proxy_protocol_v2 / 9bdebcea2cbc / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- advanced_options.proxy_protocol_v2

<a id="canonical-5cf925e609f8ebd8a04a23310c1ea4b9dbc5787e2ea85cf29042ad7143a6ac5d"></a>

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

<a id="canonical-0a50b88e0b4529f7fe1c4baf76fecbca88a686cac3fa94409eaafb71de6c45fe"></a>

## Direct properties — advanced_options.proxy_protocol_v2 / 9bdebcea2cbc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e8a93a478f3cc1f785cfbb929ab0aec01b0b7b1a6519c3a133678e1c405c681"></a>

## Next pages — advanced_options.proxy_protocol_v2 / 9bdebcea2cbc / 4

- [advanced_options](resources--origin_pool--reference--group-001.md#canonical-4bd6850d2de0a7490f991cc4f76833162eb76bfb1646864d1894d395e0d859ad)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-562a3acefba6f1de598f335e49c1ccd5b6fc8bf55ed03dcb7c1b8592e7bb66c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-559b894c1d2401188a3607bbeccbea5edd6b656d46bf3fab5ecf428414b4215c"></a>

## automatic_port — automatic_port / 8ec0e739b58e / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- automatic_port

<a id="canonical-7328a554a0db646d4173b6b522cf6c6b3633699941810d8295576afdfff78961"></a>

Type: `["object", {}]`. Optional.

\[OneOf: automatic\_port, lb\_port, port\] Enable this option

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

- [automatic_port](resources--origin_pool--reference--group-001.md#canonical-7328a554a0db646d4173b6b522cf6c6b3633699941810d8295576afdfff78961)
- [lb_port](resources--origin_pool--reference--group-001.md#canonical-63fd7950004f798a70fb424fc7458fb474d3b1282aac03cc1481bc5625e83548)
- [port](resources--origin_pool--reference--group-001.md#canonical-be501128f5e3d84e69b9ae443dc6dd371a24c754218e96c6650995b7ad5e7793)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
automatic_port = {}
```

<a id="canonical-cef532b52bbbf62d316099ca3a68c21de4115f0b10fd7ab737bce258deffdcfb"></a>

## Direct properties — automatic_port / 8ec0e739b58e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9de29d798fb785af696aeacded9b22e0980e5dd0d50ba7872020432d60e87154"></a>

## Next pages — automatic_port / 8ec0e739b58e / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-28cdf1ceb2efb9bb358019dbc9717f35a6f260c3edd34c16aa1a76490ce84607"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b344d261d154c96c29dd0cd3db6778e4195c0e7f963bbbc95c948838240e28dd"></a>

## healthcheck — healthcheck / 05b49fdcb79d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- healthcheck

<a id="canonical-f5007421860a7fab444859ff097418c99a182bdea5969f38c3c31d717f658a3b"></a>

Type: `"object"`. list nested block, Optional.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Reference to healthcheck configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
healthcheck {
  # Configure direct properties listed below.
}
```

<a id="canonical-00ebb399a2b2a334b91d0f648718532c7da0f848a16e1ebfc949547008bb7f1f"></a>

## Direct properties — healthcheck / 05b49fdcb79d / 3

<a id="canonical-d2705f7d646a3d90a03f2bac485fa62a1dc7d0f6e57bc3542f4fcf6775887181"></a>

<a id="canonical-cc9b0ec308b09470b335047285e7180ca76fbbf6e1e582361aabf57bf768fe90"></a>

## name property — healthcheck / 05b49fdcb79d / 4

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

<a id="canonical-ee7245ab3d3850f1df1bc5d5e4187fc9d842be089d9cf27bb80f51ed65163a46"></a>

<a id="canonical-a2d8b52ba535e8c48a0d5876b61c97bbc82e720937afe08782d8807fe0f2a6a5"></a>

## namespace property — healthcheck / 05b49fdcb79d / 5

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

<a id="canonical-53fcfea637bb1320f4494f5438a0f9cb753d4ac7a413f5671d7aef53f140e32a"></a>

<a id="canonical-81bad19c9e68026c30891eaf7f4185613619ef757cac89433ad1409c76135ff3"></a>

## tenant property — healthcheck / 05b49fdcb79d / 6

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

<a id="canonical-14d0456fda01d2e6f2cc94b7535f5ab78c1a752714f8a945fdd163060d2598cb"></a>

## Next pages — healthcheck / 05b49fdcb79d / 7

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-08c2f998734b62d52645d8fd722945970f71788b1cc66b844a568ecb08b11050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8ffd74bc5e8efc99f97ec240e5d8cab034aa2ae4b9a3c949f58ba74f2e974a7"></a>

## lb_port — lb_port / 5c706863cde7 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- lb_port

<a id="canonical-63fd7950004f798a70fb424fc7458fb474d3b1282aac03cc1481bc5625e83548"></a>

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
lb_port = {}
```

<a id="canonical-bedd4983210cc57b011b43902fe4af324edab0783f14985aba309e74b856b9c1"></a>

## Direct properties — lb_port / 5c706863cde7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84af5e37d80fd9fc20e07fc711acf0009afcf6bdad716e1dd48424a021bca960"></a>

## Next pages — lb_port / 5c706863cde7 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-ee8653d9b863ea4392eec42b70dbfc2bfef2c0f4270f8332c7945741c9026b4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a69dc4514e54351eb5737be005e6459b642ad5c8aa4d3152983a8fe4a284f90"></a>

## no_tls — no_tls / 73e7d934dc12 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- no_tls

<a id="canonical-0d9835930e24bec33e65513b60c6666097a8ff7e52b1d06c06fe840193d9ee98"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: no\_tls, use\_tls; Default: no\_tls\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [no_tls](resources--origin_pool--reference--group-001.md#canonical-0d9835930e24bec33e65513b60c6666097a8ff7e52b1d06c06fe840193d9ee98)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-da6a539edede34014702388d9e0abd5ba500bd242fc3a4a40f407fdf4ae46046)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_tls = {}
```

<a id="canonical-7b1248e0af5b350a5f0863e9b5790e4ec575ecd2c2d91154082c35833a021bcf"></a>

## Direct properties — no_tls / 73e7d934dc12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62fe5873e26921b3cc47674443e6c7146f876a18f68736f3ac874fae108bd49b"></a>

## Next pages — no_tls / 73e7d934dc12 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffc9567c30f52032af29432eb92f5bf02c055512deb2785b26c67531ff07a602"></a>

## origin_servers — origin_servers / fbff96f07045 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- origin_servers

<a id="canonical-81ea379dca927997f61834000c27f38a27f279dff31737cda01b2f67600d8705"></a>

Type: `"object"`. list nested block, Optional.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cbip_service",
    "consul_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("consul_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("vn_private_ip",
    "vn_private_name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-86852e0bb6dfcf7bc815ccfd887a81aaa55c6b0f162a3492d66bceb3004bcc1c"></a>

## Direct properties — origin_servers / fbff96f07045 / 3

- [cbip_service](resources--origin_pool--reference--group-001.md#canonical-44f52dbbaddf46d774b4043dd44361b2fd543d949218ca3368e03764cec38501): complete subsection reference.

- [consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d): complete subsection reference.

- [custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-4707dd4589e4425912ca819bdba0f57eaa059d66a65782ad2e52450b8c12e6f9): complete subsection reference.

- [k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0): complete subsection reference.

<a id="canonical-2aeb5c46329d7c19737e8417b48452075f2b1f2b65033b76ae3aacf0add861fb"></a>

<a id="canonical-33e3ab89ecf49a6cefc196e33b77a93f659e01cad9fe6e155d9aec1f9e6eccc8"></a>

## labels property — origin_servers / fbff96f07045 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this origin server, these labels can be used to form subset.

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

- [private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43): complete subsection reference.

- [private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b): complete subsection reference.

- [public_ip](resources--origin_pool--reference--group-002.md#canonical-86892a0b8abf53409d722e2d600702358eeb6ed794eb16cede78c0fd203677c6): complete subsection reference.

- [public_name](resources--origin_pool--reference--group-002.md#canonical-116ce7a9a539106140d695639f56717a3f212c77d02f1eac26ab9efaa3734e7b): complete subsection reference.

- [vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-9c6aea85b0e4c5054d34baa041360d879809b176747d1ec0c14e7d1b8beb9df7): complete subsection reference.

- [vn_private_name](resources--origin_pool--reference--group-002.md#canonical-4790c620301dfed425ab1a0395b5baee9c0e9940049d0b38fa5a5ed992615e56): complete subsection reference.

<a id="canonical-b31d45248b7cb2542c01e7c036edb6f68e12d0f9474a70d8a9b4919edeeb6106"></a>

## Next pages — origin_servers / fbff96f07045 / 5

- [origin_servers.cbip_service](resources--origin_pool--reference--group-001.md#canonical-44f52dbbaddf46d774b4043dd44361b2fd543d949218ca3368e03764cec38501)
- [origin_servers.consul_service](resources--origin_pool--reference--group-001.md#canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d)
- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-4707dd4589e4425912ca819bdba0f57eaa059d66a65782ad2e52450b8c12e6f9)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-807eedd9067c3254887005dfeaa77d3a564cb3cacb261d63c3ed9cb49a836cf0)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-2c6da3442d0b6ec67494cb089009cadef7abf00026bdc05ae26a0c0b45758b43)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-5f29de2edbca1bf266d4175c8611bfdf1235ff2e40f228e3a30fc1e25889003b)
- [origin_servers.public_ip](resources--origin_pool--reference--group-002.md#canonical-86892a0b8abf53409d722e2d600702358eeb6ed794eb16cede78c0fd203677c6)
- [origin_servers.public_name](resources--origin_pool--reference--group-002.md#canonical-116ce7a9a539106140d695639f56717a3f212c77d02f1eac26ab9efaa3734e7b)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-9c6aea85b0e4c5054d34baa041360d879809b176747d1ec0c14e7d1b8beb9df7)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-4790c620301dfed425ab1a0395b5baee9c0e9940049d0b38fa5a5ed992615e56)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-44f52dbbaddf46d774b4043dd44361b2fd543d949218ca3368e03764cec38501"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78bfdada82932d586a26dd8a9764d2700bcf16b5ff40382917b5009188394be1"></a>

## origin_servers.cbip_service — origin_servers.cbip_service / 8e8916e5a3aa / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- origin_servers.cbip_service

<a id="canonical-ae17b2cd9b8b80dac1ad748930b26182b22dd9b415b22d0a3c85ddabf0ba0395"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name")}
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
cbip_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-443063a8179db024b6c01e0c352ebaac07fe295735b801a2eacd669b40dc16a0"></a>

## Direct properties — origin_servers.cbip_service / 8e8916e5a3aa / 3

<a id="canonical-9ae2544e893e7bfc0e051a53dbc69fbcd5daa06c9f63c8026ddde25f1ff536d9"></a>

<a id="canonical-baae0bb75a3b15f8fb67d39e8ea10555c9d99a529e5865bf57e3b99529fcb2fe"></a>

## service_name property — origin_servers.cbip_service / 8e8916e5a3aa / 4

Type: `"string"`. Optional.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

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

<a id="canonical-076d55f010775f72bf8926d552019ceb1f9b773f41fa7635f63541b4b09201d8"></a>

## Next pages — origin_servers.cbip_service / 8e8916e5a3aa / 5

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-e79254f217492d7ef12376b7b0d4967cd4cfdec4c2c8b3b6fdcfa04c9f004918)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-e7953cdd2d65579ba3e25ce6ffa198ed7447a59c86e7fccde05f11a418f1b53d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
