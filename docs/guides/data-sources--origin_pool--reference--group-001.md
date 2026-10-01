---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34266efa4a4ec1596d2dc5f1722ddf5c5e4b69c94b03f2a322d2d21d310fcf87"></a>

## Property reference — Property reference / 3ff6f7b85ba2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- Property reference

<a id="canonical-9f3e15a0f0e8a742685736f744974db6ea21b666a84403d0768df66e28592b40"></a>

## Direct properties — Property reference / 3ff6f7b85ba2 / 3

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c): complete subsection reference.

<a id="canonical-84adc83a65eb7275552209ece12e0749bacfc302a2c679c82c22e90bee575ed7"></a>

<a id="canonical-ec381fbe5e4f0f30b9c89ef50b023e92958efe60c3e130dd03a62477fbaae997"></a>

## annotations property — Property reference / 3ff6f7b85ba2 / 4

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

- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-ee83b1c5c66570d25195e6adf3441a9ddee69a5606e8092b84e05fbfa6b99ddd): complete subsection reference.

<a id="canonical-09d00da307bccc6c0bd0a4ec58b974a98c8721f957cdfc5078f152358473819d"></a>

<a id="canonical-8f67d0a9ad8b0cc55302e6a9549e276163fa8f55824c8a63910d219201168e18"></a>

## description property — Property reference / 3ff6f7b85ba2 / 5

Type: `"string"`. Computed.

Description of the OriginPool.

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

<a id="canonical-964e26863d906d0eb568f47994e62aba4df50e922bb86338363504f70621481d"></a>

<a id="canonical-b06c6e2f384c1a479aad59583292e452a3a4ed5d289bfba846953ce8cfd51575"></a>

## endpoint_selection property — Property reference / 3ff6f7b85ba2 / 6

Type: `"string"`. Computed.

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

<a id="canonical-75a26e08d5790f58031011d4836ae53e946c3b87ea1a0e56d6bee8ffb002e842"></a>

<a id="canonical-cced8ab36aa088fc9f30eb218ba8e486837c46c5c9a2895eb15cafb690f17624"></a>

## health_check_port property — Property reference / 3ff6f7b85ba2 / 7

Type: `"number"`. Computed.

\[OneOf: health\_check\_port, same\_as\_endpoint\_port\] Exclusive with \[same\_as\_endpoint\_port\]
Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

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

- [health_check_port](data-sources--origin_pool--reference--group-001.md#canonical-75a26e08d5790f58031011d4836ae53e946c3b87ea1a0e56d6bee8ffb002e842)
- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-ec881b452d0dd978f0b3bf07f4f243573b579dc9584c3238af9df0f90788304c)

Select alternatives according to the provider validators above.

- [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-f9c5329002cb51dedaf6893faf1e6fdf7bf89fef384fc6c4c762f123a10d47ca): complete subsection reference.

<a id="canonical-fb45314d5ee83c7ff4d5531846bdd2c0fc3df9858088dfcbe7bd67002822231f"></a>

<a id="canonical-79299d2f6b9779e1b849fdd598e3f60ba917ef1b31ab271c5e99e069ee0d5de4"></a>

## id property — Property reference / 3ff6f7b85ba2 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-806bce32df9c6015f5ca9250b1060ebf6ff0a05de817438a0df9364fd513cc93"></a>

<a id="canonical-3204ade4a5154e4ef1dc879276f93d16b8b4a714b68b940baead559e028a88ee"></a>

## labels property — Property reference / 3ff6f7b85ba2 / 9

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

- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-b9f7675bc66fd490a145bca6473577ae0bfa83dfe9c0163a5d8280a3990293f4): complete subsection reference.

<a id="canonical-cfac1e1e004e450d3f16a6de7b87098957133792384e96823ce286215ac44cf1"></a>

<a id="canonical-527249398b4ec1c2e96e22604f78dd78d7181bc916b772add75874432dc01016"></a>

## loadbalancer_algorithm property — Property reference / 3ff6f7b85ba2 / 10

Type: `"string"`. Computed.

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

<a id="canonical-0267b5e2ae1d80a9c01e2d0f74b0e3cd7f0149b72293f774305cbacc51c565c2"></a>

<a id="canonical-f2a20b9324e734d1afabf41af4805adaa4b43985c54859837d54d082b168eb24"></a>

## name property — Property reference / 3ff6f7b85ba2 / 11

Type: `"string"`. Required.

Name of the OriginPool.

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

<a id="canonical-c629f794aff74c7432cf07e8e0586b0454dfdb8334695877cb96b655f1560694"></a>

<a id="canonical-ef9548e2962c7fa34afb8a2fd64f24d3c2010f1b8ef944ac6c17756bef09400f"></a>

## namespace property — Property reference / 3ff6f7b85ba2 / 12

Type: `"string"`. Required.

Namespace where the OriginPool exists.

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

- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-c30a60b3744a38686e8b5447b85e7d12379a5a560a12b3df726b08a42f2337ad): complete subsection reference.

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5): complete subsection reference.

<a id="canonical-7dbca99b2e206e11371db9038b1de32467fbe5fc28a8ef6a5f255c4e857a58ef"></a>

<a id="canonical-7e84970d2a6a9308bd88fcefd31be41109b5348b058fc73b3a8f21b7aeddb55b"></a>

## port property — Property reference / 3ff6f7b85ba2 / 13

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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

- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-dc185cb410008c05bf9d16e9bbc2d76dec9df69be9a6270d1af105dc52e365cf): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d): complete subsection reference.

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1): complete subsection reference.

<a id="canonical-d6640d2881434e74f8e03ad9876d8fa03cc0d9da5466bb155994217703c0d1fa"></a>

## All schema paths — Property reference / 3ff6f7b85ba2 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-8fb7305b2a6e2af4f974021dd9ab10c88e4dfbdca38fb097b3031d93a0c5905b) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-0bcfa6bc148453eed2b487d7dac507e02d633059d58f3f6f219ad6e9a047dbdc) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-21d129a0e1218ab13bc5bba3d3f977f7cb51bda7d1419b11f50157c61a3c77fd) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](data-sources--origin_pool--reference--group-001.md#canonical-8f1eef20cf3810746218cb1a11ffc6525b899205dcea24448153b2c150cc0c10) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](data-sources--origin_pool--reference--group-001.md#canonical-5aeb5f1b65860cce94e5d4f4e32f63d90bb391879234705df672f9a0b17c7d2f) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](data-sources--origin_pool--reference--group-001.md#canonical-32d9651ad5c7806a6315f1a606c78cfd9fe2a8e55d924c05cb5b9ca198ab5bb8) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](data-sources--origin_pool--reference--group-001.md#canonical-fe6116f956e2cd338bcd33feae12289d00121d2380584537f5bc09a04ac8a954) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](data-sources--origin_pool--reference--group-001.md#canonical-fa7580fdf035fceeafea77a2508e86bb848a3031fe2d3b78f664277a18674a36) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](data-sources--origin_pool--reference--group-001.md#canonical-77c6697663a6df35c99dfe7bf03305e534e8966e08c3a2a617a37a7cb7b3623c) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-c3910697548b92080875ad784d45ae7afacd56c0c692d7a0cfe92be3fe1819d3) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-93f5318a19f6fe10488f85aee4afa5576735150a111be8ca67ff56e91a93e6d2) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-2b46888ae150275e1a1285adf1631a41265c2a30cfd86e04dba0d9868a0b986a) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-1e29d688841fd33bb4c39d09cd38a7cfbd2a33008ea4362270a66c649cdda8dd) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-35575d1e85a664a2ea3db58132409e7d524a497ef50a638e144fef27dc4d93bf) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-578b4c93c72262ab48338f7bc2e9cb22cb21d2b7573f3fafdfb1f13c59724c60) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-47b2898b4f7db1a1d9c79fe5e0bb53da0a6e3f94462a2aeccfcbdc6e86c37582) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-f062a7977ff29c521df4ff8a29a59b91ec6ff9a78c44cc9e321e968f50ca7c00) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-d1ab76efbce8a83ed914182fd949b1ecaad4f24c3c17acd75d21f71265c64473) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-7c9eac209c1a1c5978b2e2b0de593ab08e6472486b240ebb43c142eba7627aa8) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-34bf523ebd2b41a4d7494408bd6b479a2ac2a86ab6f54b97c692f56cf47a3852) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-c16e77f4811e2f57f61341504d4ead775d0b9bdbc559695c32da5887c1f3933d) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](data-sources--origin_pool--reference--group-001.md#canonical-ba6d8cb0250c38edb63e91c18aa359ea448bad032c742166802f320f9a6277d5) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](data-sources--origin_pool--reference--group-001.md#canonical-6ad2ec89c86cdbffce3a1f5b9cec0a8e157a3797e098f3291198317dbcd52d0b) |
| `advanced_options.http1_config` | [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-c5d52e9c7bb6b93c1a8818e18371259a26f87752628c0a68cba88851a0ff392f) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-2b56e5a9b027819ebb6b1c9b9d46fd8472d910be7e3712c2d71fc7796afc04c7) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-379fea9e430966e7634063171037501bdf9612d6767fac39b554c40b924d0882) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-a844de717418d170c67a391d0500213a322e7ded3b417fb670d847f8c4050edf) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-c0b8075badb0fd1a070575bd25cc732cdfb8c5585ad5fab64a9eebb9c3d50bf7) |
| `advanced_options.http2_options` | [advanced_options.http2_options](data-sources--origin_pool--reference--group-001.md#canonical-0c6ba7f9710f04c3f1769372c4da32a230b0abc6dd3f752096fd3a0c8d2f864a) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](data-sources--origin_pool--reference--group-001.md#canonical-f24c85debcca24f05c2e5fa9777d45c69c332090d6b1ce811af74dbb377e6798) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](data-sources--origin_pool--reference--group-001.md#canonical-01e5247d9c36c70f3c837574133c46ae5c46c6adef76dd9c182731a9d0828623) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-cd36be8d654dd7594b87a429befef57aedee39ca1db16f30a77563d45ca3d8e7) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-777ad7871f97f3e18f0701f62b913df640d033d07d9b18be62a574add13718b8) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-7cf591cdfa324e39d375cf7fe7982b6a6e62737fd0dd6e1050c06996d2fce92b) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-18b0101fe3d2b956de7c504582b90667a20f09d64b1af25432b45f55674f4cad) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](data-sources--origin_pool--reference--group-001.md#canonical-4eb2a3672b4a1d7fa6de2c9af803f8199b16a7dbf64d7bb6262e1fd9ab293799) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](data-sources--origin_pool--reference--group-001.md#canonical-f450a244a53fb6a16e25a0fed044a62a52f126c3ac39a0257a81d9077c0501b1) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](data-sources--origin_pool--reference--group-001.md#canonical-89a600e09b35e868d2a52c357e50cf1ae808abd6b5cb6a4f6ab58a88171b8d9f) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](data-sources--origin_pool--reference--group-001.md#canonical-17e02e39c8e9fbd1a185cb55ac3ee7168269b8668086543a2be324b523da1fd0) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](data-sources--origin_pool--reference--group-001.md#canonical-e5c283bea864abccb2f8d1e7ba975d25a4955dbc8eb0e6f15287b4db70a28fa1) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-4c425fd00fee84b3ab00998291e1573ad50c9130756f3a895e5b60138e4d9e2f) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-dd24f25afec3de6ca7d1ceed3e20ceb5f45b96b1cbbac8e6e4bc37a3959e88b9) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-f3380002659b371edf17b16c64cb4b8ea348cebb25a05ab27a97b4de84e92917) |
| `annotations` | [annotations](data-sources--origin_pool--reference--group-001.md#canonical-84adc83a65eb7275552209ece12e0749bacfc302a2c679c82c22e90bee575ed7) |
| `automatic_port` | [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-fc3119de4b057057bb0c3e9ac36bd7545e10cade5df8ab15ded77f579cde2036) |
| `description` | [description](data-sources--origin_pool--reference--group-001.md#canonical-09d00da307bccc6c0bd0a4ec58b974a98c8721f957cdfc5078f152358473819d) |
| `endpoint_selection` | [endpoint_selection](data-sources--origin_pool--reference--group-001.md#canonical-964e26863d906d0eb568f47994e62aba4df50e922bb86338363504f70621481d) |
| `health_check_port` | [health_check_port](data-sources--origin_pool--reference--group-001.md#canonical-75a26e08d5790f58031011d4836ae53e946c3b87ea1a0e56d6bee8ffb002e842) |
| `healthcheck` | [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-4775a28fd85744a41aaedcd0428c2df0d6da8f161ba7268f950e2782a0a2ec41) |
| `healthcheck.name` | [healthcheck.name](data-sources--origin_pool--reference--group-001.md#canonical-4888fa14ccc4d9af7d972b650a186dd97703689ae7eeb6ac07a05dd93819c005) |
| `healthcheck.namespace` | [healthcheck.namespace](data-sources--origin_pool--reference--group-001.md#canonical-9e70c1098ac96b5d6f83756de984505528bd0a1b8ffb1f20b7a1e5811a68f6b3) |
| `healthcheck.tenant` | [healthcheck.tenant](data-sources--origin_pool--reference--group-001.md#canonical-ff6ae18922ad492377cbe9e05a957f6ce918783c122f0084d51356bb4f1bdd37) |
| `id` | [id](data-sources--origin_pool--reference--group-001.md#canonical-fb45314d5ee83c7ff4d5531846bdd2c0fc3df9858088dfcbe7bd67002822231f) |
| `labels` | [labels](data-sources--origin_pool--reference--group-001.md#canonical-806bce32df9c6015f5ca9250b1060ebf6ff0a05de817438a0df9364fd513cc93) |
| `lb_port` | [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-fee3cd22c1f53ee4104f51d80243405380ff4038c42a4d9c573be480a4b1c80a) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](data-sources--origin_pool--reference--group-001.md#canonical-cfac1e1e004e450d3f16a6de7b87098957133792384e96823ce286215ac44cf1) |
| `name` | [name](data-sources--origin_pool--reference--group-001.md#canonical-0267b5e2ae1d80a9c01e2d0f74b0e3cd7f0149b72293f774305cbacc51c565c2) |
| `namespace` | [namespace](data-sources--origin_pool--reference--group-001.md#canonical-c629f794aff74c7432cf07e8e0586b0454dfdb8334695877cb96b655f1560694) |
| `no_tls` | [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-c31d7aa8a56d1aa9f182420c47cb0998925f802ffbb1db8676c2d1136f44eaaa) |
| `origin_servers` | [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6202f7dfc72ee06927396d7ea6f61bca050257a7ce50c5ac98d6ca4c8159907d) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-2f3bdc3bea748f677670a901a6f2e6f24f63d137a12048f50150080a91065a66) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](data-sources--origin_pool--reference--group-001.md#canonical-ff1a0428e5163c350b7c56c2c0406ce65a520014bc8d623489e009076e56bd12) |
| `origin_servers.consul_service` | [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-24d102c4c185943baf5675605435c6d4478cc07836f4082b8b652a711ac9830b) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](data-sources--origin_pool--reference--group-001.md#canonical-78b6575736f10dd8007ef49b2eda47706267b8d2fe9b6bdd3eafe2f917adadd5) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-81ea009a6dc5bfba306456579588e2349076791a29016f5bf23436335a4a4a97) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](data-sources--origin_pool--reference--group-001.md#canonical-7065af5a5b0467e48088de787b049905650745d76c244023b6dfae255695aead) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-a10a3225aa809327b3dbf2f473b0d5b29becd296c8380b6f08ecc37ba57f5571) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-770fbe12b39f6cacf251e79d542924ca63d04ceaecb7ca3efc8c2ae22e175049) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-c06af92907fb6d738166468b6cf7c8f815ec6a8ed41c0f9813a751738e698e8c) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-031295adf16a424db5ff2e932b3fc109c9a548a051caa82d0f838fdd85aecb4c) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-e59a2319e0d769a2ee4927ccc39e1d76b1e565d0940b619783560bf6070a61e6) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-08cbe9c225ba413e24b0be53c6a48fefc9e49dfbbd7466d015c99b01ba662ccb) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-f5b5c6745711fc4dd42fc696058518c1a667091c4a3be6836b1c12733dbd117c) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3bd4c7f24836d8b7b07ffe63c647e1951105326af2d25bab328147f2c6258cc4) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-8d142dbe166a658889dc26686d5ffa32c24274a78b7dc2725774be568d0375d3) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-82d90332c22aed68f17bd76604d3856d579c476ccb2e472aeabf12dee32a52a2) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-e284b4b38610bf18f682b4b0f93c66f640e1ee35a32be96eecddc615cd96e794) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-caa372535c0c9df868fc35fc3e5bbc07792b4148201a2f76838be632e3d22cfe) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-247bb92d9c8b81bc539b01e573827c038572521197379c17977d5b93d786722a) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-ff272def9ed381d6177b5f7d9b90bc088cc5c5f3f480f000ecce5dd8e8d0dd23) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](data-sources--origin_pool--reference--group-002.md#canonical-6753b96d3fede410153cbb36986f18bf4e520de52b22f425ba3c88f9ff4257e1) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](data-sources--origin_pool--reference--group-002.md#canonical-ddebee53ef7d2931ef60589f026d50e8bab2edbd618712151cb89f87ee3405bb) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](data-sources--origin_pool--reference--group-002.md#canonical-38680cd0a56b05d06d016b835056fb84303ffbbfd8f2cddbab426fb5f8a6e89a) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](data-sources--origin_pool--reference--group-002.md#canonical-df322039f1103e3dd4e8adca81b781d69e088ea2b120ad35c38612513521ae1a) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-6bbc8291e89644837af72bf4e4db993e8ade6f2bce94ccea4b613fe057f9e1f6) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-a3d7d40aeab39cec9818d99d6036c8b3313a717b2f0e843f1545adbd70aacc9e) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-c626bf8dbd344e7a792c172a5472dd6a4cca2e4804a3abe420080d770c6427ad) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](data-sources--origin_pool--reference--group-002.md#canonical-8f512e6d85f32ea26f508afca96a30933008f04df84b2f2f3e091bc9c3ece285) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](data-sources--origin_pool--reference--group-002.md#canonical-a33f8a81932d56ea9bbe2fca72c4ccd70ed1535fed8ef592b75bc4c504b57987) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0a5b4f0a144525708ab23af749c2d49078da616928d6619809e499af5a2dcb61) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-5eb033314d7e24a1d5016bdc9504fe0f3b76572fadd0a58328dabfba237502b5) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-88866a55c5f3712389880e69537e79185b272f1e0c69fc04565dc3ed1e9ad058) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-22be4be707b1debb1916b3aac6629032180a22891c0755671b6c86fa600cba39) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-aa1b7d240a38740465b3c62155cb1eb1d76a1f69c12cb0e44d6602035dafffe5) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-343baa270be35fdaf017735efe0ff309e28fccd478ba56ef97b9ff9ab259e4e0) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-d964e1b4fba535b3cf49891676171f534530c157f790bc01f564d247ff36a01d) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-d001a426353b052d92d11ff0d187b72dbc66b55388c578cd5a36d28867285cee) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-4c76abe91de8873a4725e2159f8bb07a51a6e65f899c998d8f71cb432ecf6c22) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b079615f41a92b55493c597c479e2fdc4df23a9a4da404542e6d929695c25382) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-967f2dda11227aba91a2d6b47ced624d8c83d3c84d1f740187b9826b38988dfa) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-31ecf11385680a6d231dbec295a7e18cbb1ff6fc33c6d2e2baeac84c5a680f74) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-60c6f910f75a32f6904d9bb5a2106a51bce0d304cda26872dea717824c281909) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-59e29f02dc7ded37db556319a022b3f6f28a0bf16342300a97fef5760af24395) |
| `origin_servers.labels` | [origin_servers.labels](data-sources--origin_pool--reference--group-001.md#canonical-525320e2af59490223d3f3006ecc645ae0fafea96079e1a703539e97d7cd1f91) |
| `origin_servers.private_ip` | [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-7b78e4b127ca4a304866f4cb48d6a55649dd1d975060949e6faf28ad1e6f74f9) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-7efded0ac379797b3a61bb817de8824bfa40d0c8c7fe589acc5ef510790342cc) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-38655241a32d9f14ca608624a90ec0b6d704a1e6a7699c307a1d6521fea7d0fa) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-cd3d4f5ee3e20d12fcff65f7235488d512f998f70115727ede576e2cff4813ec) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](data-sources--origin_pool--reference--group-002.md#canonical-e154ac5ae6b19c866aeb2fc4b1b542c0b48ec028e136abd954537f989bf65150) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](data-sources--origin_pool--reference--group-002.md#canonical-59f84e843d0671e19dd02970aae2ad8492fef7dd5cd2e8b70a1c567934a0b6f7) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](data-sources--origin_pool--reference--group-002.md#canonical-f2d908fafd11e8ed432737cb085ca223c5a07fe0af414f817c884ccc22679f85) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](data-sources--origin_pool--reference--group-002.md#canonical-eff8404654d58067d8a85acb3faae66500a0f418298b896dbf0c62345b2af5e6) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-5a6dd14d378cb947f1b9d4aa0e7defc1e79ad4250e833007a9a780131c3935f3) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-245205a82bc0d92de0161f7c07a6edfffd724d1ff8a762fd31a42428566fa030) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-85fdfd0fd5587e05f583b01f04ddc8bc8658d3c054576e0a6aa4f29b562cb6d2) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-6a0b9421518717fd4eedfa962795fdaf0a0e6eeaf63f0ca9c0929eb432e32aa0) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-066e5585ec4e5d77e75f699c4a8090d2789e82103ef6b8f1a968a4bca326af5c) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-baaf62dce8415ebabda8f37b88487db0493f5acf7d4b64df19655eb935c37852) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-9aa1aecfc5dee5d476b533a9c9534a212328eca6374451a4f5cff55695d4d0fc) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-c4f2d0fc7e52e2b65eb4b574071c3e4651704dbfe9065b881dfa99f470a1f960) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-1c020a8223b05ac8133256fbdac84cbfe03822058e6e2e7875e2c2cb584d2486) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-6c277abbd82cd7606f5f1221aea8416b0e276fa5d07b6c19a2f3ac777394d73a) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-03cb97e9dcba6f935a9010bd047b3d434831b28606f0299eae5dd8f38935391a) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-aeffe7333c62117e0a957a43be80b3041178db66de2179e59433bbcd1f3131ec) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-ab71ad81e970f7eec2dfff64a2122b6730f280d8b47f828162b1b6bab13b2887) |
| `origin_servers.private_name` | [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-667792fdfba67427d732773d4a07976b305fdba8299bb49a1eb73821116abd00) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-0f529d7a2e0a68cc2a322a192f4f56a0b9214c9d189b9ef5c3dc795954baeee7) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-f61fc861ded55b3aaeaca9bf478fd806cfc1b8ca0038f20f8cff1b617cda6bf0) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-17d12717d08771d18c828494df50dd981a88224e4f07d613c6f9f280b9d7eedb) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](data-sources--origin_pool--reference--group-002.md#canonical-04c72a2b1a5aa9f776b5733a3f786acbc5ab531b218e45f54a8ad4cc928d1a39) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](data-sources--origin_pool--reference--group-002.md#canonical-6904524a0f14570289a82147d89bf02b434923439c9f94eab0c059d42a4d62ba) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](data-sources--origin_pool--reference--group-002.md#canonical-4f89f4147a17b07cd992f4dbc0aeac320a99ee65f10c23f33d71da6090dd7d84) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](data-sources--origin_pool--reference--group-002.md#canonical-b96fa63736b2d770e31bbff123dd65be7b5f843d8a113721a45fb7cf06a7d8c4) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](data-sources--origin_pool--reference--group-002.md#canonical-648bda1664e6fcc2c9debf64b530eb850e135fc20479cddf9acf764767b1e46b) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-7b3ccef1409ee4a76a5e010daf8828653625a1a8c5ef9e936abc9607e6f7621b) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-9a4db3929dff81fa270c7131f4866042e502e3175158e4953ded11a8c3044661) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](data-sources--origin_pool--reference--group-002.md#canonical-65bce5e841889dee3589cff5b5c32d52add7b53801918c3247446512dc530204) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-5839b765023d267a3dc1b58077bd510fd57971ecd204668d58b5df83d4db9b58) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-6f8fc465f4a258f47d7039fb8427c3acc6b4b84715207bbbedb7850adde7454d) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-4c79d144666db1e529185b266b67efba4eba2299c5f749d2133f917b06c47b1f) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](data-sources--origin_pool--reference--group-002.md#canonical-d43cc47befc9042c20c508601dbae18eac429611daef50a92ac133ba2a405cda) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](data-sources--origin_pool--reference--group-002.md#canonical-3fa422982493f467854901c9eca3b2b3559ad7c92e4c9e232ddeda183fbd9b30) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](data-sources--origin_pool--reference--group-002.md#canonical-6068d2cb7dd17488712770a55c9c07be7a657b14b849ebe8c7ba805f2eddd00b) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-7dbe6978bc6f62e71f4e56661845e2910cfbf09873e5066cda24b7ce71d99c27) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-6d87f71e424905304d5cce2e2003cbce06a4d6edd847bde0da709e7ee3117f6b) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-8060bd568736191dc8507aba98f9bab9ae42285bf7c1d750613a493346491fb0) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](data-sources--origin_pool--reference--group-002.md#canonical-2f4e5b5f94a83b91387f1802cc1a79784f1423ddf054beec65d0f9b6d8c32c76) |
| `origin_servers.public_ip` | [origin_servers.public_ip](data-sources--origin_pool--reference--group-002.md#canonical-5b50a97014f839c7e6097061273397bc043c16d29c282c9c4e363e0c398d2e8f) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-96df61bbb42bcbf4d6d570d273b1b1027d155ea13f888458e14b768c105a6b9a) |
| `origin_servers.public_name` | [origin_servers.public_name](data-sources--origin_pool--reference--group-002.md#canonical-6fb411cd8555e196f633d30e145238f4efd924bbdc6c77917cbb92e6ac41c14b) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-60990dda2ca9a53fce1e5e9fab6c1f1553d9cced4c81931f31eeb8faf25cbf3e) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](data-sources--origin_pool--reference--group-002.md#canonical-eaa734c3a1f00cd790dc29553de4f0e89cd444c1e063f41cfb643b0f0bbdccc9) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-615972742a9e0fc94f9b8822ece675ed1f58f6ec4a0c83b89a96ad8650f3d09b) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](data-sources--origin_pool--reference--group-002.md#canonical-83afeac1eed5b3b4760440b08c09eac022157e45e8c376eeb6629c4789738a8c) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-0225eaf983978cc253679db5ab9305bd30e56815dfcda7bb3bf124c9cc946cfe) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](data-sources--origin_pool--reference--group-002.md#canonical-461fcca72d51e6fc9f9dd2f18bf68ec3eefbc90d9c29e2bab0d481c0e34ce521) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](data-sources--origin_pool--reference--group-002.md#canonical-ed4a74dbc0196ab42390a63dba4a755320d697b5f8106d4c5990833788c424db) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](data-sources--origin_pool--reference--group-002.md#canonical-f6abe4c886be357bccb89294b6f194ad77afe6036749473e4a8655a3eb852cf0) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-f6c0b230a9fe56ffbcfe1e0af871b97796562a72ab579c4bd86ec2ca0d1c55ed) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](data-sources--origin_pool--reference--group-002.md#canonical-51d7dccc0e1b4ae14c5c1250eb838b76b0a6de54e530a0863bfe8c3058d929f7) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](data-sources--origin_pool--reference--group-002.md#canonical-8a04b86b67d07902bfdaf1485f1285772e6062aeba087bcbf9bcbabecc47f304) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](data-sources--origin_pool--reference--group-002.md#canonical-dd31a30eab24cd74c8153d2c1011ebf0fe38acfa37a38368f548e5c8bfb9d805) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](data-sources--origin_pool--reference--group-002.md#canonical-50698aa361d95dd078fa56fc4bb4953fba7cba0c88883d1a7794e6681f23c094) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](data-sources--origin_pool--reference--group-002.md#canonical-9a7dfabf7da9cb20c871ef085905da302cf463f5fd7ff34646cae2dd996a8862) |
| `port` | [port](data-sources--origin_pool--reference--group-001.md#canonical-7dbca99b2e206e11371db9038b1de32467fbe5fc28a8ef6a5f255c4e857a58ef) |
| `same_as_endpoint_port` | [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-ec881b452d0dd978f0b3bf07f4f243573b579dc9584c3238af9df0f90788304c) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-6bd239adb6868a830e65cede41e9b2d8cc481cdf1d8d2d182f468c43e4a30c7e) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-695f2671b7219f78298ff6fb9b82bc027c354e9ff781f4367457ae429ba7e127) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-73d8e778899a16458998e6722757d7982183e42821eabcd79c62776ff361305e) |
| `use_tls` | [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-02a9345e3df45300299a1b27753191de2b20bdfeb558d4a0de38bf539a781e09) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-ba888a9c6e50be9fff457057e4ba0b7d6d61bfb8c4c4f9b6d31a5d2f359e437b) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-831352753e23156bdbb91496d35c11133f82f0f292149cf556d8aeb77215cb30) |
| `use_tls.disable_sni` | [use_tls.disable_sni](data-sources--origin_pool--reference--group-002.md#canonical-666a0984c79e72a9c9d9e7ed49dc7115cdb6a54f25c3a980bf9ea6cfa8183720) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](data-sources--origin_pool--reference--group-002.md#canonical-2f4d1a3848d7351750d559796df92c62b839caae0f5575a890646f791c5ac7c4) |
| `use_tls.no_mtls` | [use_tls.no_mtls](data-sources--origin_pool--reference--group-002.md#canonical-99cbd2d8d8b63b3855a361be4c6a5827bc479ddd4b76aa5230de7d4d154d83db) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](data-sources--origin_pool--reference--group-002.md#canonical-fd23590b59a2ce32728a54efb4eb626edea2dfb7f9f5bd53e0cf6d9781716379) |
| `use_tls.sni` | [use_tls.sni](data-sources--origin_pool--reference--group-002.md#canonical-f718cb2b088a501a8843aae5b022fceee2aac73a6f9d81ef0218a964397a929b) |
| `use_tls.tls_config` | [use_tls.tls_config](data-sources--origin_pool--reference--group-002.md#canonical-a9649983f595f34dfab45b4b894164595c3bbdf94ade57504f20db1abee483e3) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](data-sources--origin_pool--reference--group-002.md#canonical-d39a13a0de6ff191f3cc7c0c97a4c2b6d556d3555c6174fc91555c97c3e97753) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](data-sources--origin_pool--reference--group-002.md#canonical-5fd087290530d3962cc09485ff02b6bf92943da2cf7ad2a9c74bc0b0d72c0aec) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](data-sources--origin_pool--reference--group-002.md#canonical-fbdd660f964ef55eded06a691021838f69b5e0c8c63a4f0c0b610634eba165e0) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](data-sources--origin_pool--reference--group-003.md#canonical-6aea64e3f4141fb6afed6bca2e593f3ae94692b01eb6937ecef1e09d060ad803) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](data-sources--origin_pool--reference--group-003.md#canonical-c5750f2ebdac13019022f0410f5a4cbddbd24e5dcba7f0ecd2b14cb5fee40e51) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](data-sources--origin_pool--reference--group-003.md#canonical-c13090001ba1f5b8c2bbf8fb98758498ecd3dba66575c48ff25fcdfca2d3634d) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](data-sources--origin_pool--reference--group-003.md#canonical-6d9c5f7bcc6d4877d32a0f35f62570e59b3cef5fbb17b90053d3b8b95e181cee) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-c20b08bc3ea0a185ca2886e94f8bfbd6e4be658be476dd5466f33bc62b712007) |
| `use_tls.use_mtls` | [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-0f346614982ab588cb8d9246cbc97e266e36e0c8366073f7430f747b56cc15fe) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-a29af6c5b34ccbcdc00e3100b01209edc7c556e83645109a6c7b5d1b798fe4be) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](data-sources--origin_pool--reference--group-003.md#canonical-6bf4457816c58543f658047eef705cf665bc5b979c4289e9977ecc39b440c046) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](data-sources--origin_pool--reference--group-003.md#canonical-684b88b355351a4a22cce93e1d8c5b03175c68bd2b20c1d3db8ebc83b8031bce) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--origin_pool--reference--group-003.md#canonical-681a9b2747b2968c16530872de047c272e7db9e36c2855533b6133bccf449080) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](data-sources--origin_pool--reference--group-003.md#canonical-4d9fece768216abeb0c8758bc61e52e44bf71ce8ee0876f5d9e52e06fa6ab939) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](data-sources--origin_pool--reference--group-003.md#canonical-149fbfb6cd2c91dc9a3b8cd1ce7ec2d5eb578cb77e2a6df6ec3cdd2b706d11a5) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](data-sources--origin_pool--reference--group-003.md#canonical-b1732eb9dd6a32b5d31b815d62921ae1a28a36980c48d58e3e0c03b1864df707) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-0bc2c65a471923680f4f12d039bdefa3feb01c50911905d43f5c7a712aaec55c) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--origin_pool--reference--group-003.md#canonical-31cdd92d4a02b81e01641424eafe7a427cf3bc093dd434a6b9ac3a2ef9a6b19c) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](data-sources--origin_pool--reference--group-003.md#canonical-40058bad8044c39aa39ea7306cbaa56a291adbe4d7cb11f4b5ca7ad88b4a827f) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--origin_pool--reference--group-003.md#canonical-b02fc2c42495e8d2c7e43f89bb8439a05c839c926130dde9965304f8228db9b1) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-b35db8718c720029c6d5d29bbc4545d28b6a4366d002602072bba872dd92595c) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--origin_pool--reference--group-003.md#canonical-1463db5e88b3a0b63020cce221f33ba3efa7e7d2ae9b48156ca66f7a9b3b935a) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](data-sources--origin_pool--reference--group-003.md#canonical-2e8d8abd9bf87deb4d4a2c069114da28278d4ae7cfea35521a4fe58e7d86fcbb) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](data-sources--origin_pool--reference--group-003.md#canonical-0dd264dc74dd02d6c6772ad179c25d09b07186f944ac68805712e989ed35438a) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-e9b82aaf4d40d19d0dd6129e5295cb2354db76e8a3fdf1a2d8995ca824db1b02) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](data-sources--origin_pool--reference--group-003.md#canonical-1d3173eb5b6d404e7fcff9829de97af321daafbd2a05078a13e047418a2434c4) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](data-sources--origin_pool--reference--group-003.md#canonical-50a019137ac6b00d0cc7f139008452ef9a762b58846017f1eebc0a281a9b4deb) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](data-sources--origin_pool--reference--group-003.md#canonical-30d23ba80f9859e8cbe84697ed88d8ceb8f2c4a44d9db7dd8e402aa1c76f0be7) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-bebb4a53c8164ad7c4c02190a8dcd00c3c1ace20f0fd356de1e486c6d19ab187) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-bff194d9658aaf88a849ab751fb12c652207312511980351bc97903175f16314) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](data-sources--origin_pool--reference--group-003.md#canonical-d113973a5129a55f5add759edd047decb0b30a371b8cde00f579818ad8337201) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](data-sources--origin_pool--reference--group-003.md#canonical-cb97b259f27fab6a07b05e2ebf66d5783bd889bdf5cb099fd90ca44cea41d3be) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](data-sources--origin_pool--reference--group-003.md#canonical-1d487250e05a8dc4961fc0761087200cbea168e251f49c9dbf02a720389f579d) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](data-sources--origin_pool--reference--group-003.md#canonical-96c01ceb7f4ebbd48a8596cbdd103e13566d9d7836b5f24cc413dd27006847d6) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-b69cb6f54b4f0012f752ba3b6616c5dde146fbb3d468b2764403169e85c0329b) |

<a id="canonical-36f9a5a1f2d0752b833c6103d6e6cf16f233991971f79c84d1029675cc24879d"></a>

## Next pages — Property reference / 3ff6f7b85ba2 / 15

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-ee83b1c5c66570d25195e6adf3441a9ddee69a5606e8092b84e05fbfa6b99ddd)
- [healthcheck](data-sources--origin_pool--reference--group-001.md#canonical-f9c5329002cb51dedaf6893faf1e6fdf7bf89fef384fc6c4c762f123a10d47ca)
- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-b9f7675bc66fd490a145bca6473577ae0bfa83dfe9c0163a5d8280a3990293f4)
- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-c30a60b3744a38686e8b5447b85e7d12379a5a560a12b3df726b08a42f2337ad)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [same_as_endpoint_port](data-sources--origin_pool--reference--group-002.md#canonical-dc185cb410008c05bf9d16e9bbc2d76dec9df69be9a6270d1af105dc52e365cf)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-feca7059ddb95c1c0f9ac26182a6d9ce1d3170127d7aaa5ebd34b58fe31124bb"></a>

## advanced_options — advanced_options / 8ba982632054 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- advanced_options

<a id="canonical-8fb7305b2a6e2af4f974021dd9ab10c88e4dfbdca38fb097b3031d93a0c5905b"></a>

Type: `"single"`. Computed.

Configure Advanced OPTIONS for origin pool.

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

<a id="canonical-e0f75309c96d2036afb3972653721f0fe848227d6e474e7a4bd4051224980d63"></a>

## Direct properties — advanced_options / 8ba982632054 / 3

- [auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-a034a08c70c96b68070bddd4163de04949808c75443f42f986fe2a3feb176214): complete subsection reference.

- [circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-f394ef52c077ecae95ff54fd0af060877bc0f42b35ed4c524f45dc990c32e245): complete subsection reference.

<a id="canonical-77c6697663a6df35c99dfe7bf03305e534e8966e08c3a2a617a37a7cb7b3623c"></a>

<a id="canonical-d985098245794b0b5ba495252aaa93bba58a3ae3995ac68c59db09a091cc1767"></a>

## connection_timeout property — advanced_options / 8ba982632054 / 4

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-aad9b1db2839f37be9ed97be187fb380d6a6412ffd2fe53d332f2bfa8295a5d6): complete subsection reference.

- [disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-b3f69eee9d3c5e0876de3b2dce47a1049e0a1d1e9b7d7124918ed11b92f8c4e7): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-673bdb812b8bdcb41f25ffa60b8629586d57b64d7030da2dfe82d759337e37a3): complete subsection reference.

- [disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-3dfb4e0eeda2e0c022061b3c7746d17e45fec1b9eca80527b60275b06623670b): complete subsection reference.

- [disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-6532745d42c120687b41b30233da6efdcb320b1d5fb3eb0f189cd46ec5d5b137): complete subsection reference.

- [disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-e103fbb4533e36ac5cd01e2f5ddddfd2e2a4058c5768c0fc3c91ab79589cf6e6): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-8df8bf15037f7f68b0996e582e7a063d4a25b00c22117284ff847f2f0c832d89): complete subsection reference.

- [enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9): complete subsection reference.

- [http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e): complete subsection reference.

- [http2_options](data-sources--origin_pool--reference--group-001.md#canonical-35c66efb1a9f6d3b66df53c268e20fab473e17ddcaec020b65fc929c17e17594): complete subsection reference.

<a id="canonical-01e5247d9c36c70f3c837574133c46ae5c46c6adef76dd9c182731a9d0828623"></a>

<a id="canonical-1f93508b6a0e6f872bbf26cbbd5c207bb2557266e193e5d25850dd1c97dfc8e8"></a>

## http_idle_timeout property — advanced_options / 8ba982632054 / 5

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

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

<a id="canonical-cd36be8d654dd7594b87a429befef57aedee39ca1db16f30a77563d45ca3d8e7"></a>

<a id="canonical-88492d64d9d52e7c939ef32d545f75ed8c6bce717af95123c122a3369de5b70a"></a>

## max_requests_per_connection property — advanced_options / 8ba982632054 / 6

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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

- [no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-3e6ceb308815c54afa432281eaafd413c9bf6b3fef9f3e0f133f00344b0a247f): complete subsection reference.

- [no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-2cb4c492bd00d43b6ef85b09b4744dfae5135302840b7d60550fbea368334b52): complete subsection reference.

- [outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-486e541a55ab36e83bf842dbc7f6e9ca96b493e234eee61fa51bc4e4cb296a1a): complete subsection reference.

<a id="canonical-4c425fd00fee84b3ab00998291e1573ad50c9130756f3a895e5b60138e4d9e2f"></a>

<a id="canonical-14cd22105c365490097f32ec6fe8ea3c9cb5f0b6edadce8ddbc83656db172696"></a>

## panic_threshold property — advanced_options / 8ba982632054 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-96caf8eeb67295ec326a12d6d3081e58ddd958aa850de88bc6c390ac7708630d): complete subsection reference.

- [proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-7444cdbf043a326c1ebaf2de6cd1bb94268f6768d978b68c13410ddb317156f8): complete subsection reference.

<a id="canonical-e29f25a161fd80caef19593e2277e2082696b96c1b3d8d1aeb539709062d2b2b"></a>

## Next pages — advanced_options / 8ba982632054 / 8

- [advanced_options.auto_http_config](data-sources--origin_pool--reference--group-001.md#canonical-a034a08c70c96b68070bddd4163de04949808c75443f42f986fe2a3feb176214)
- [advanced_options.circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-f394ef52c077ecae95ff54fd0af060877bc0f42b35ed4c524f45dc990c32e245)
- [advanced_options.default_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-aad9b1db2839f37be9ed97be187fb380d6a6412ffd2fe53d332f2bfa8295a5d6)
- [advanced_options.disable_circuit_breaker](data-sources--origin_pool--reference--group-001.md#canonical-b3f69eee9d3c5e0876de3b2dce47a1049e0a1d1e9b7d7124918ed11b92f8c4e7)
- [advanced_options.disable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-673bdb812b8bdcb41f25ffa60b8629586d57b64d7030da2dfe82d759337e37a3)
- [advanced_options.disable_outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-3dfb4e0eeda2e0c022061b3c7746d17e45fec1b9eca80527b60275b06623670b)
- [advanced_options.disable_proxy_protocol](data-sources--origin_pool--reference--group-001.md#canonical-6532745d42c120687b41b30233da6efdcb320b1d5fb3eb0f189cd46ec5d5b137)
- [advanced_options.disable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-e103fbb4533e36ac5cd01e2f5ddddfd2e2a4058c5768c0fc3c91ab79589cf6e6)
- [advanced_options.enable_lb_source_ip_persistence](data-sources--origin_pool--reference--group-001.md#canonical-8df8bf15037f7f68b0996e582e7a063d4a25b00c22117284ff847f2f0c832d89)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- [advanced_options.http2_options](data-sources--origin_pool--reference--group-001.md#canonical-35c66efb1a9f6d3b66df53c268e20fab473e17ddcaec020b65fc929c17e17594)
- [advanced_options.no_panic_threshold](data-sources--origin_pool--reference--group-001.md#canonical-3e6ceb308815c54afa432281eaafd413c9bf6b3fef9f3e0f133f00344b0a247f)
- [advanced_options.no_request_limit_per_connection](data-sources--origin_pool--reference--group-001.md#canonical-2cb4c492bd00d43b6ef85b09b4744dfae5135302840b7d60550fbea368334b52)
- [advanced_options.outlier_detection](data-sources--origin_pool--reference--group-001.md#canonical-486e541a55ab36e83bf842dbc7f6e9ca96b493e234eee61fa51bc4e4cb296a1a)
- [advanced_options.proxy_protocol_v1](data-sources--origin_pool--reference--group-001.md#canonical-96caf8eeb67295ec326a12d6d3081e58ddd958aa850de88bc6c390ac7708630d)
- [advanced_options.proxy_protocol_v2](data-sources--origin_pool--reference--group-001.md#canonical-7444cdbf043a326c1ebaf2de6cd1bb94268f6768d978b68c13410ddb317156f8)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-a034a08c70c96b68070bddd4163de04949808c75443f42f986fe2a3feb176214"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd2fc095e67edb872146f383c01c92981460b2ace4c3cf40f5a0002eabbe960c"></a>

## advanced_options.auto_http_config — advanced_options.auto_http_config / e9bf4bf946d3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.auto_http_config

<a id="canonical-0bcfa6bc148453eed2b487d7dac507e02d633059d58f3f6f219ad6e9a047dbdc"></a>

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

<a id="canonical-feb390f73d83d5feeef1bc7cba23a7e3c4a8f99e156c9d3f56787a7434018661"></a>

## Direct properties — advanced_options.auto_http_config / e9bf4bf946d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec50abfb11baa5057b92f8842737c0008cf5cb9ec13ed0b16bddf03ce0a60918"></a>

## Next pages — advanced_options.auto_http_config / e9bf4bf946d3 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-f394ef52c077ecae95ff54fd0af060877bc0f42b35ed4c524f45dc990c32e245"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c78a35809fce21d100ac68b66f3a9214bcca58d4235de5e7ea96a48afbe54938"></a>

## advanced_options.circuit_breaker — advanced_options.circuit_breaker / 511bbef395a3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.circuit_breaker

<a id="canonical-21d129a0e1218ab13bc5bba3d3f977f7cb51bda7d1419b11f50157c61a3c77fd"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0d452ea50f4389f61e6fad40f05b87ccdeebf6c1dc65200a22cfba24208188b0"></a>

## Direct properties — advanced_options.circuit_breaker / 511bbef395a3 / 3

<a id="canonical-8f1eef20cf3810746218cb1a11ffc6525b899205dcea24448153b2c150cc0c10"></a>

<a id="canonical-0ff38b644cdc14f7a1cd051dfad50d03cb5b6820f6d265bb5efd3f1391cf0a6c"></a>

## connection_limit property — advanced_options.circuit_breaker / 511bbef395a3 / 4

Type: `"number"`. Computed.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

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

<a id="canonical-5aeb5f1b65860cce94e5d4f4e32f63d90bb391879234705df672f9a0b17c7d2f"></a>

<a id="canonical-7fcc57f04c7e813cb0539de4482deb6b22642a4291750a5ba3b0028101b0de09"></a>

## max_requests property — advanced_options.circuit_breaker / 511bbef395a3 / 5

Type: `"number"`. Computed.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

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

<a id="canonical-32d9651ad5c7806a6315f1a606c78cfd9fe2a8e55d924c05cb5b9ca198ab5bb8"></a>

<a id="canonical-b6dd2f754addd8a0a7b04474737bd505be2a49e657807ac643d2a22d3a1c54d7"></a>

## pending_requests property — advanced_options.circuit_breaker / 511bbef395a3 / 6

Type: `"number"`. Computed.

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

<a id="canonical-fe6116f956e2cd338bcd33feae12289d00121d2380584537f5bc09a04ac8a954"></a>

<a id="canonical-342269472e37d8abbe9064ea7dc5d9732e74816707374538c56718fc69a1eb31"></a>

## priority property — advanced_options.circuit_breaker / 511bbef395a3 / 7

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

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

<a id="canonical-fa7580fdf035fceeafea77a2508e86bb848a3031fe2d3b78f664277a18674a36"></a>

<a id="canonical-f23ae698a5db345b00affc43adfc57f2cfa971d958c8e5254b22392ea5ec35ed"></a>

## retries property — advanced_options.circuit_breaker / 511bbef395a3 / 8

Type: `"number"`. Computed.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

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

<a id="canonical-a5535ee8525b8a80a256486c23c47fe1d7f655a1b1c71274bc0caec64c1ecb57"></a>

## Next pages — advanced_options.circuit_breaker / 511bbef395a3 / 9

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-aad9b1db2839f37be9ed97be187fb380d6a6412ffd2fe53d332f2bfa8295a5d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db8479bb4c988d0fc95affe21f9c967fd4b6fb47f04da236dff061454c03cc1b"></a>

## advanced_options.default_circuit_breaker — advanced_options.default_circuit_breaker / c3d4974c6c19 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.default_circuit_breaker

<a id="canonical-c3910697548b92080875ad784d45ae7afacd56c0c692d7a0cfe92be3fe1819d3"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ae61204f7302b610f421e3bc24fe20fd1418c0abd4424bfd6f570b4e3476fd6b"></a>

## Direct properties — advanced_options.default_circuit_breaker / c3d4974c6c19 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c42c7a2bf38f4f7c08ed62609759baf17a5b71c84d9250d0df76a1ab2711a122"></a>

## Next pages — advanced_options.default_circuit_breaker / c3d4974c6c19 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-b3f69eee9d3c5e0876de3b2dce47a1049e0a1d1e9b7d7124918ed11b92f8c4e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5652062c89325c9703886c60d0f8204fbbc30e2606485d042f66f4b7e115370a"></a>

## advanced_options.disable_circuit_breaker — advanced_options.disable_circuit_breaker / 27dc2681d56d / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.disable_circuit_breaker

<a id="canonical-93f5318a19f6fe10488f85aee4afa5576735150a111be8ca67ff56e91a93e6d2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1100b393abf0652925974b0aa4152c4d96aabb50dab5f5260a0130e14974ec4f"></a>

## Direct properties — advanced_options.disable_circuit_breaker / 27dc2681d56d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a016ca764a30ec485be598937497d83efbcee31dd717909bd31ec162454b81fb"></a>

## Next pages — advanced_options.disable_circuit_breaker / 27dc2681d56d / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-673bdb812b8bdcb41f25ffa60b8629586d57b64d7030da2dfe82d759337e37a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af2947cb813a1b3c092c2ea69af59513166013d93e0198aef86582b33536ed34"></a>

## advanced_options.disable_lb_source_ip_persistence — advanced_options.disable_lb_source_ip_persistence / f98f1160c6c8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.disable_lb_source_ip_persistence

<a id="canonical-2b46888ae150275e1a1285adf1631a41265c2a30cfd86e04dba0d9868a0b986a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e01526014e0f183fa92d80426cfe6831a0654689f5cef97f9e1b7aa8116485c6"></a>

## Direct properties — advanced_options.disable_lb_source_ip_persistence / f98f1160c6c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10f732f6395d3eb8ce9808a8eeb9a6f91b2d09050a86b1aba066481224f10087"></a>

## Next pages — advanced_options.disable_lb_source_ip_persistence / f98f1160c6c8 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-3dfb4e0eeda2e0c022061b3c7746d17e45fec1b9eca80527b60275b06623670b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c69a1c3999f8081dd22d3b898e0425d71cc085d0ff2189a59af8e2b63045ffb"></a>

## advanced_options.disable_outlier_detection — advanced_options.disable_outlier_detection / 11548cefc5b0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.disable_outlier_detection

<a id="canonical-1e29d688841fd33bb4c39d09cd38a7cfbd2a33008ea4362270a66c649cdda8dd"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-36f050d59714d83ecb18a679aeb894e5276299fdf956e72bd0701d2b67ddd600"></a>

## Direct properties — advanced_options.disable_outlier_detection / 11548cefc5b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c85618c4fa8d11bb95ccf02c4e5e3c58413338564560cde0419f9163395091cd"></a>

## Next pages — advanced_options.disable_outlier_detection / 11548cefc5b0 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-6532745d42c120687b41b30233da6efdcb320b1d5fb3eb0f189cd46ec5d5b137"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa65cdf028cf36d46733642fc30f7dc9ff929a112551495723c59caf4592b805"></a>

## advanced_options.disable_proxy_protocol — advanced_options.disable_proxy_protocol / 1e8c976e8e3a / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.disable_proxy_protocol

<a id="canonical-35575d1e85a664a2ea3db58132409e7d524a497ef50a638e144fef27dc4d93bf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f95787a8865ccd8f80b08a6986b5cc4cd3b3677e905e28f89fa3b3aecbc24938"></a>

## Direct properties — advanced_options.disable_proxy_protocol / 1e8c976e8e3a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2412c85249afc2b7dc5440f6c82f6770600f565e69ee7cf77b1876f3d92f19b8"></a>

## Next pages — advanced_options.disable_proxy_protocol / 1e8c976e8e3a / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-e103fbb4533e36ac5cd01e2f5ddddfd2e2a4058c5768c0fc3c91ab79589cf6e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ffe99e0fc5f394b7f4ed5ee3dd96e092b2e22532dc6d6dd2b2e0c31f4293205"></a>

## advanced_options.disable_subsets — advanced_options.disable_subsets / 400119212c8f / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.disable_subsets

<a id="canonical-578b4c93c72262ab48338f7bc2e9cb22cb21d2b7573f3fafdfb1f13c59724c60"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-38d83c89e7c6d926c50c52cfe5eaa18eb2bf3c8ab29f9bdbfe20b993485233be"></a>

## Direct properties — advanced_options.disable_subsets / 400119212c8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4fed7b9adb77aa9faa5080b62d7af3b2c26c587aabb1599648504495b5ca8383"></a>

## Next pages — advanced_options.disable_subsets / 400119212c8f / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-8df8bf15037f7f68b0996e582e7a063d4a25b00c22117284ff847f2f0c832d89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42a5b6da0606a5ad52d8669f45b6dabe5c1368fe94b3a26851064ca773000345"></a>

## advanced_options.enable_lb_source_ip_persistence — advanced_options.enable_lb_source_ip_persistence / d687e30e1dc8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.enable_lb_source_ip_persistence

<a id="canonical-47b2898b4f7db1a1d9c79fe5e0bb53da0a6e3f94462a2aeccfcbdc6e86c37582"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-312d75848c3b5b11f7ed2a398e23fc9eabb72dd58ea825ee6e635a5d036f2595"></a>

## Direct properties — advanced_options.enable_lb_source_ip_persistence / d687e30e1dc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28ca6b13c8d637f1a36acbd22d2f478a8580a78240966603d688f5ec324da253"></a>

## Next pages — advanced_options.enable_lb_source_ip_persistence / d687e30e1dc8 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-568e08d8fc2f8cedc9574753309c6353e2ad2a2ea608433c75d4225bc9232d8a"></a>

## advanced_options.enable_subsets — advanced_options.enable_subsets / 70a9be275874 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.enable_subsets

<a id="canonical-f062a7977ff29c521df4ff8a29a59b91ec6ff9a78c44cc9e321e968f50ca7c00"></a>

Type: `"single"`. Computed.

Configure subset OPTIONS for origin pool.

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

<a id="canonical-c24d5f68e10b88a3eab8d0b64a9860777d602cead541b3c79d3e154965a0d7bc"></a>

## Direct properties — advanced_options.enable_subsets / 70a9be275874 / 3

- [any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-82ea0192cf0a6fafaaa046638d435fdcc539d08c6125b0d3828ff5281d569675): complete subsection reference.

- [default_subset](data-sources--origin_pool--reference--group-001.md#canonical-25a51496aeddcdd25c69fab2f86e28400a789736aff1cd061617c9e2b795241e): complete subsection reference.

- [endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-67b8ab9d961fbf2f759fc61ef319420ed6635d401401fe3fa56ddb837f4ea9a4): complete subsection reference.

- [fail_request](data-sources--origin_pool--reference--group-001.md#canonical-be91bb597de8967360d040fd768850bcf267bc9bed37b13ecece9ea5de2fe144): complete subsection reference.

<a id="canonical-406387888f607662d4cbc7fe75cb9b521f7b8fd6dc974880a60cbf93d0e939ce"></a>

## Next pages — advanced_options.enable_subsets / 70a9be275874 / 4

- [advanced_options.enable_subsets.any_endpoint](data-sources--origin_pool--reference--group-001.md#canonical-82ea0192cf0a6fafaaa046638d435fdcc539d08c6125b0d3828ff5281d569675)
- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-25a51496aeddcdd25c69fab2f86e28400a789736aff1cd061617c9e2b795241e)
- [advanced_options.enable_subsets.endpoint_subsets](data-sources--origin_pool--reference--group-001.md#canonical-67b8ab9d961fbf2f759fc61ef319420ed6635d401401fe3fa56ddb837f4ea9a4)
- [advanced_options.enable_subsets.fail_request](data-sources--origin_pool--reference--group-001.md#canonical-be91bb597de8967360d040fd768850bcf267bc9bed37b13ecece9ea5de2fe144)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-82ea0192cf0a6fafaaa046638d435fdcc539d08c6125b0d3828ff5281d569675"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32377596d37a1f8e865ddd1de9cf2340f6bdeb38e76575662a1f067c4c9e76f2"></a>

## advanced_options.enable_subsets.any_endpoint — advanced_options.enable_subsets.any_endpoint / d9f524df7fa3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- advanced_options.enable_subsets.any_endpoint

<a id="canonical-d1ab76efbce8a83ed914182fd949b1ecaad4f24c3c17acd75d21f71265c64473"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7e0e376a2d10ac90085aebc61a280489e87b8986f09c78dd2a2ed957f65324f5"></a>

## Direct properties — advanced_options.enable_subsets.any_endpoint / d9f524df7fa3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aee40dc575af8cb379205519e2e9df6eb2efa5d97eb99e23a0acd7b978075932"></a>

## Next pages — advanced_options.enable_subsets.any_endpoint / d9f524df7fa3 / 4

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-25a51496aeddcdd25c69fab2f86e28400a789736aff1cd061617c9e2b795241e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce8427aca8f77b9fb7a1dbb1cd5bdbc50bf624a017d0f07043354e5816f77009"></a>

## advanced_options.enable_subsets.default_subset — advanced_options.enable_subsets.default_subset / 625c85939297 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- advanced_options.enable_subsets.default_subset

<a id="canonical-7c9eac209c1a1c5978b2e2b0de593ab08e6472486b240ebb43c142eba7627aa8"></a>

Type: `"single"`. Computed.

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

<a id="canonical-141dc828eabdaadefa6ea04596db4f2d275a1f0a02433d65b8026e6174d3cadd"></a>

## Direct properties — advanced_options.enable_subsets.default_subset / 625c85939297 / 3

- [default_subset](data-sources--origin_pool--reference--group-001.md#canonical-4197fc3186a8a828c6cc5d9bb0b4f478bd88adaa8ef2f1648fd6aa280e671972): complete subsection reference.

<a id="canonical-46d904ef6e05a5d9c15b665f78c24b0362b130397302e2bcc1102fb692a848d3"></a>

## Next pages — advanced_options.enable_subsets.default_subset / 625c85939297 / 4

- [advanced_options.enable_subsets.default_subset.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-4197fc3186a8a828c6cc5d9bb0b4f478bd88adaa8ef2f1648fd6aa280e671972)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-4197fc3186a8a828c6cc5d9bb0b4f478bd88adaa8ef2f1648fd6aa280e671972"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4216cce6f40617e66bf6cb6828130a5de3e873e86bc488d875273067921ac449"></a>

## advanced_options.enable_subsets.default_subset.default_subset — advanced_options.enable_subsets.default_subset.default_subset / e8d4d14e8986 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-25a51496aeddcdd25c69fab2f86e28400a789736aff1cd061617c9e2b795241e)
- advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-34bf523ebd2b41a4d7494408bd6b479a2ac2a86ab6f54b97c692f56cf47a3852"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c6635a5733182c14d122b34cd1b37485e54aaa5d094d8544b348c83ef251547f"></a>

## Direct properties — advanced_options.enable_subsets.default_subset.default_subset / e8d4d14e8986 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2eb629593407756959d76f26991d262d67a22ac05952bcac6a5a1031a293a0d"></a>

## Next pages — advanced_options.enable_subsets.default_subset.default_subset / e8d4d14e8986 / 4

- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--reference--group-001.md#canonical-25a51496aeddcdd25c69fab2f86e28400a789736aff1cd061617c9e2b795241e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-67b8ab9d961fbf2f759fc61ef319420ed6635d401401fe3fa56ddb837f4ea9a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8dcf68107df0f6387b04169fac07a8c4556009af3751b7281d50597c1818032"></a>

## advanced_options.enable_subsets.endpoint_subsets — advanced_options.enable_subsets.endpoint_subsets / f76ce7df9f2c / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-c16e77f4811e2f57f61341504d4ead775d0b9bdbc559695c32da5887c1f3933d"></a>

Type: `"list"`. Computed.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

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

<a id="canonical-005aa5bb0e2a9eb3a3546eaae94852a27f3fdc86c414c0cb89a2ac737621c14c"></a>

## Direct properties — advanced_options.enable_subsets.endpoint_subsets / f76ce7df9f2c / 3

<a id="canonical-ba6d8cb0250c38edb63e91c18aa359ea448bad032c742166802f320f9a6277d5"></a>

<a id="canonical-703c79c1c94f05b570c40dfe315b46b150ee91cc09be1669d77b86729366cc74"></a>

## keys property — advanced_options.enable_subsets.endpoint_subsets / f76ce7df9f2c / 4

Type: `["list", "string"]`. Computed.

List of keys that define a cluster subset class.

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

<a id="canonical-d07a0118d9439c977d0008592a912433d3d238624f547edf8ce04382b49277b1"></a>

## Next pages — advanced_options.enable_subsets.endpoint_subsets / f76ce7df9f2c / 5

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-be91bb597de8967360d040fd768850bcf267bc9bed37b13ecece9ea5de2fe144"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3c9e834670d6076c7f1f41d8d6a6dfb3be95b8243080dc2e9dc9499a640e51a"></a>

## advanced_options.enable_subsets.fail_request — advanced_options.enable_subsets.fail_request / 622899a3ed83 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- advanced_options.enable_subsets.fail_request

<a id="canonical-6ad2ec89c86cdbffce3a1f5b9cec0a8e157a3797e098f3291198317dbcd52d0b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ff1bba8d1de187549cb992804e0557faaf428d12bfb8f8d316cb93df5c46aeb5"></a>

## Direct properties — advanced_options.enable_subsets.fail_request / 622899a3ed83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c88fa82d6b0303b0d0cd54fe482711a66eeb47f4dab56431db39ca53d44d80e9"></a>

## Next pages — advanced_options.enable_subsets.fail_request / 622899a3ed83 / 4

- [advanced_options.enable_subsets](data-sources--origin_pool--reference--group-001.md#canonical-1438b00c9ffd2fa8576a742d76bf0a61155b1f66f7df58010c88c8a683fda5b9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1d450147bcc5687a1d6a6244f67799cd2357f0c82839e81468104fbbc4721ff"></a>

## advanced_options.http1_config — advanced_options.http1_config / 6d880122243d / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.http1_config

<a id="canonical-c5d52e9c7bb6b93c1a8818e18371259a26f87752628c0a68cba88851a0ff392f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9e92e9e787860dcfb13e7225fd7333aa05c1d44b9fd58faef853d27c7aeb74eb"></a>

## Direct properties — advanced_options.http1_config / 6d880122243d / 3

- [header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e): complete subsection reference.

<a id="canonical-a5ccc4e4305b15b9944506f6540e426e0940dba40da1e5a1884181330a5a7899"></a>

## Next pages — advanced_options.http1_config / 6d880122243d / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f962fa75522fb34b77aefa9676beeaacad2e09af0c6c44303299cbe03e4c1af7"></a>

## advanced_options.http1_config.header_transformation — advanced_options.http1_config.header_transformation / b8b0599cfac1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- advanced_options.http1_config.header_transformation

<a id="canonical-2b56e5a9b027819ebb6b1c9b9d46fd8472d910be7e3712c2d71fc7796afc04c7"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

<a id="canonical-f48c541359eefb47b4dfce98f6a0c5c24ba726aa4ec63a107cb7c1381aa0c557"></a>

## Direct properties — advanced_options.http1_config.header_transformation / b8b0599cfac1 / 3

- [default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-acca87e7695da89d92156f16c7d1e993429aaa024f15fd25b084bee61a2d525f): complete subsection reference.

- [preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-233aaa8a1fc85fae944a18c334d1f934b2854d2687eefd428f8be63d64ec6652): complete subsection reference.

- [proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-3f7dcd356c0b560878953152e03b90faf42b0de4cf6bf24e50240e976eca36b0): complete subsection reference.

<a id="canonical-07fde395956bc92b156cb584cbce4424933e2d0401144aab2e310da12490490d"></a>

## Next pages — advanced_options.http1_config.header_transformation / b8b0599cfac1 / 4

- [advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-acca87e7695da89d92156f16c7d1e993429aaa024f15fd25b084bee61a2d525f)
- [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-233aaa8a1fc85fae944a18c334d1f934b2854d2687eefd428f8be63d64ec6652)
- [advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-3f7dcd356c0b560878953152e03b90faf42b0de4cf6bf24e50240e976eca36b0)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-acca87e7695da89d92156f16c7d1e993429aaa024f15fd25b084bee61a2d525f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78b3e9bfee4018afec32480812c8dd1ea2f17f9f24b19513924b8622fcdfaaad"></a>

## advanced_options.http1_config.header_transformation.default_header_transformation — advanced_options.http1_config.header_transformation.default_header_transformatio / 102ae02c062a / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-379fea9e430966e7634063171037501bdf9612d6767fac39b554c40b924d0882"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-c30f24e7f0152c3477e195244e9c2a0252844f5356dc898e103bb71c32749e29"></a>

## Direct properties — advanced_options.http1_config.header_transformation.default_header_transformatio / 102ae02c062a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06df6343de1ebf815ae8bd8d654043486aebd055e317de0968ee1cb98acd8aa4"></a>

## Next pages — advanced_options.http1_config.header_transformation.default_header_transformatio / 102ae02c062a / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-233aaa8a1fc85fae944a18c334d1f934b2854d2687eefd428f8be63d64ec6652"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-012006a08a4c2638e6ac9143a5071a9147dac0eef75d8e3bcfd32035b3d0120b"></a>

## advanced_options.http1_config.header_transformation.preserve_case_header_transformation — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / 02fbb13f4fe3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-a844de717418d170c67a391d0500213a322e7ded3b417fb670d847f8c4050edf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b599c689737a41aad88452f68900e05580f6ecfb6b3481b0a830de093fc8d17b"></a>

## Direct properties — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / 02fbb13f4fe3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c1a543d43b1c8e33a03dc6d0b2b1ce9f8db2ef09f09a9230f77959e93b057308"></a>

## Next pages — advanced_options.http1_config.header_transformation.preserve_case_header_transfo / 02fbb13f4fe3 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-3f7dcd356c0b560878953152e03b90faf42b0de4cf6bf24e50240e976eca36b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-522f3d227d39d278b29c3a6c4dafeaee0f52c66355db061a267b17983d6c59fd"></a>

## advanced_options.http1_config.header_transformation.proper_case_header_transformation — advanced_options.http1_config.header_transformation.proper_case_header_transform / ffa62cd51d30 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [advanced_options.http1_config](data-sources--origin_pool--reference--group-001.md#canonical-bf880910ab6e6d70195bee23f7476a8cd8f76be6e7c497789244b58661509f8e)
- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-c0b8075badb0fd1a070575bd25cc732cdfb8c5585ad5fab64a9eebb9c3d50bf7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-9d06e3a73eabb30fa345cd1d1298d07229de15cebd1fef8eaa45327c71a4330a"></a>

## Direct properties — advanced_options.http1_config.header_transformation.proper_case_header_transform / ffa62cd51d30 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afa2582a7ca0d3e1940bf978cb90cfe745f4cf44ec377ffa30a56a707ee82e44"></a>

## Next pages — advanced_options.http1_config.header_transformation.proper_case_header_transform / ffa62cd51d30 / 4

- [advanced_options.http1_config.header_transformation](data-sources--origin_pool--reference--group-001.md#canonical-7831eb0ea82d8610b3408a11e301e446a31bd5f1cba2275cb257541cf464041e)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-35c66efb1a9f6d3b66df53c268e20fab473e17ddcaec020b65fc929c17e17594"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08fc4e21ac6752472b102a67824f7d59d2f3e1d5544ceb85608c7edf558f7765"></a>

## advanced_options.http2_options — advanced_options.http2_options / 678b33c44c4d / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.http2_options

<a id="canonical-0c6ba7f9710f04c3f1769372c4da32a230b0abc6dd3f752096fd3a0c8d2f864a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-e76355a5a2104dea73fe5a3232ac35dbd29b7876cf6e11eced42e9a082d9a9ab"></a>

## Direct properties — advanced_options.http2_options / 678b33c44c4d / 3

<a id="canonical-f24c85debcca24f05c2e5fa9777d45c69c332090d6b1ce811af74dbb377e6798"></a>

<a id="canonical-7964e146f79b359af71d973e5df0b61dff9d8858d0cb80113b2ac6fe757f969c"></a>

## enabled property — advanced_options.http2_options / 678b33c44c4d / 4

Type: `"bool"`. Computed.

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

<a id="canonical-65bb44494629a10fdd6104018f46d2a503703cc1f9aa869254c352afdd15fae5"></a>

## Next pages — advanced_options.http2_options / 678b33c44c4d / 5

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-3e6ceb308815c54afa432281eaafd413c9bf6b3fef9f3e0f133f00344b0a247f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baf8f81d19429262fff1ddeef947a615f0a415e800b115ec024adc2e6702a217"></a>

## advanced_options.no_panic_threshold — advanced_options.no_panic_threshold / b93655d1cb61 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.no_panic_threshold

<a id="canonical-777ad7871f97f3e18f0701f62b913df640d033d07d9b18be62a574add13718b8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-43c91118d02c05cb41787caba586da12946a1908e1becb8f44245ecc2f5ebfa0"></a>

## Direct properties — advanced_options.no_panic_threshold / b93655d1cb61 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60c3fb0d2a1d4c695f79e15c0a28ca4991ef4be900ec655bd775752f31babc05"></a>

## Next pages — advanced_options.no_panic_threshold / b93655d1cb61 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-2cb4c492bd00d43b6ef85b09b4744dfae5135302840b7d60550fbea368334b52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b9eb28eb17b2d96afb9cfb176343d5a3b20d585794bfeaf7ead9a1546720b32"></a>

## advanced_options.no_request_limit_per_connection — advanced_options.no_request_limit_per_connection / 24fb2944b831 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.no_request_limit_per_connection

<a id="canonical-7cf591cdfa324e39d375cf7fe7982b6a6e62737fd0dd6e1050c06996d2fce92b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-37b8d17e19a1427c1505bd0b12d02f41034aa6466f81408770cfc90ac994d643"></a>

## Direct properties — advanced_options.no_request_limit_per_connection / 24fb2944b831 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-edb51b9a99bc1376b21f1e0d45f791919e1cb8a96ed899ba72c1a37d1e80bb1b"></a>

## Next pages — advanced_options.no_request_limit_per_connection / 24fb2944b831 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-486e541a55ab36e83bf842dbc7f6e9ca96b493e234eee61fa51bc4e4cb296a1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fcb07b818182d6ac8e48272ce1798e97028a2d16d68b40a9fc149c1de064c2e"></a>

## advanced_options.outlier_detection — advanced_options.outlier_detection / b5b3ec503d4d / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.outlier_detection

<a id="canonical-18b0101fe3d2b956de7c504582b90667a20f09d64b1af25432b45f55674f4cad"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0d545199545ef5a6372409721878891168e91d4b64567eb8be61c46398fdc056"></a>

## Direct properties — advanced_options.outlier_detection / b5b3ec503d4d / 3

<a id="canonical-4eb2a3672b4a1d7fa6de2c9af803f8199b16a7dbf64d7bb6262e1fd9ab293799"></a>

<a id="canonical-0c1b55ed9ef4e420278b1611ba977db85883c5e51e3a617689972c85f5cf371d"></a>

## base_ejection_time property — advanced_options.outlier_detection / b5b3ec503d4d / 4

Type: `"number"`. Computed.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail.

Upstream description:

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

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

<a id="canonical-f450a244a53fb6a16e25a0fed044a62a52f126c3ac39a0257a81d9077c0501b1"></a>

<a id="canonical-43fe60b0cc005dc37553b09d9c022ce7fc3b2ee3f426d432f8901aa0bb6a274e"></a>

## consecutive_5xx property — advanced_options.outlier_detection / b5b3ec503d4d / 5

Type: `"number"`. Computed.

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

<a id="canonical-89a600e09b35e868d2a52c357e50cf1ae808abd6b5cb6a4f6ab58a88171b8d9f"></a>

<a id="canonical-db1e9e020442d8cde60bb6433346cf478eb5677537c04b748546d1dcde656338"></a>

## consecutive_gateway_failure property — advanced_options.outlier_detection / b5b3ec503d4d / 6

Type: `"number"`. Computed.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.)..

Upstream description:

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

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

<a id="canonical-17e02e39c8e9fbd1a185cb55ac3ee7168269b8668086543a2be324b523da1fd0"></a>

<a id="canonical-32035ef739b2cf8b98e846c381cc614a1d001515e7af8a2174f3174ee7582655"></a>

## interval property — advanced_options.outlier_detection / b5b3ec503d4d / 7

Type: `"number"`. Computed.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Upstream description:

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to 10000ms or 10s. Specified in milliseconds.

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

<a id="canonical-e5c283bea864abccb2f8d1e7ba975d25a4955dbc8eb0e6f15287b4db70a28fa1"></a>

<a id="canonical-19a4d6aad6e3bcd570a47f61d3ad3af85dd22d1efc4c10e1faf8fd261bd7a6d7"></a>

## max_ejection_percent property — advanced_options.outlier_detection / b5b3ec503d4d / 8

Type: `"number"`. Computed.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Upstream description:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

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

<a id="canonical-c7bfa30e5b6f7a7d524acdebe3607926454679db88ae638d8b02550ed9f04f1c"></a>

## Next pages — advanced_options.outlier_detection / b5b3ec503d4d / 9

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-96caf8eeb67295ec326a12d6d3081e58ddd958aa850de88bc6c390ac7708630d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0fe136b54b7095a9e416c0f1f3de8a92a94e54c1f8f0beeb21286735aa39300"></a>

## advanced_options.proxy_protocol_v1 — advanced_options.proxy_protocol_v1 / 18fece0f0ab4 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.proxy_protocol_v1

<a id="canonical-dd24f25afec3de6ca7d1ceed3e20ceb5f45b96b1cbbac8e6e4bc37a3959e88b9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b3b0e98d2fbf244155e64c4c0c74702e740cbb8bd1766f286612fb5cf7f854c3"></a>

## Direct properties — advanced_options.proxy_protocol_v1 / 18fece0f0ab4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9019763d5a9dae9933171340ff1af10eb0c6be1958a18ebfe7b1b96bc723f218"></a>

## Next pages — advanced_options.proxy_protocol_v1 / 18fece0f0ab4 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7444cdbf043a326c1ebaf2de6cd1bb94268f6768d978b68c13410ddb317156f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56d7f8421b7380de554cd625d1d312833d5a6df8f6bb4aed693225b217dfa844"></a>

## advanced_options.proxy_protocol_v2 — advanced_options.proxy_protocol_v2 / f97e7d246db5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- advanced_options.proxy_protocol_v2

<a id="canonical-f3380002659b371edf17b16c64cb4b8ea348cebb25a05ab27a97b4de84e92917"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1b9d1c70499a3a2aa45b7af5597beb1240d9fde6062ce31bf5a2cf5da76ea8b4"></a>

## Direct properties — advanced_options.proxy_protocol_v2 / f97e7d246db5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a791541b2475aba83bec0b9d4b24b7430efb4611413126b6026fb545448bdfc"></a>

## Next pages — advanced_options.proxy_protocol_v2 / f97e7d246db5 / 4

- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-274bce8e43774d8ed0546235f5deb00ecc778fa3915c9922bf31bb9be3c23f5c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-ee83b1c5c66570d25195e6adf3441a9ddee69a5606e8092b84e05fbfa6b99ddd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a9babf4c2a523d062ab0774b20cab1e1f967e2a629ff3ed02c33416b26355ea"></a>

## automatic_port — automatic_port / 4abf44fea776 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- automatic_port

<a id="canonical-fc3119de4b057057bb0c3e9ac36bd7545e10cade5df8ab15ded77f579cde2036"></a>

Type: `["object", {}]`. Computed.

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

- [automatic_port](data-sources--origin_pool--reference--group-001.md#canonical-fc3119de4b057057bb0c3e9ac36bd7545e10cade5df8ab15ded77f579cde2036)
- [lb_port](data-sources--origin_pool--reference--group-001.md#canonical-fee3cd22c1f53ee4104f51d80243405380ff4038c42a4d9c573be480a4b1c80a)
- [port](data-sources--origin_pool--reference--group-001.md#canonical-7dbca99b2e206e11371db9038b1de32467fbe5fc28a8ef6a5f255c4e857a58ef)

Select alternatives according to the provider validators above.

<a id="canonical-b9d8f3ce3daa2b2e3e162f33dc302db2cd7b28fab3c829b70d0c4db8dba0cf17"></a>

## Direct properties — automatic_port / 4abf44fea776 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93997b81ddbc879cfedb3fee25e7a771004bec5ef1cec97b700f042d45024cdf"></a>

## Next pages — automatic_port / 4abf44fea776 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-f9c5329002cb51dedaf6893faf1e6fdf7bf89fef384fc6c4c762f123a10d47ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55bab4f123ea8db3037b10df7cc648cac40b5cf5228b41511aeaad63ad74461e"></a>

## healthcheck — healthcheck / c913880dcad5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- healthcheck

<a id="canonical-4775a28fd85744a41aaedcd0428c2df0d6da8f161ba7268f950e2782a0a2ec41"></a>

Type: `"list"`. Computed.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Reference to healthcheck configuration objects.

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

<a id="canonical-f50fa7b07950ee34aca3bae4d1a514d93b402dc24beed892c789511c3ca55843"></a>

## Direct properties — healthcheck / c913880dcad5 / 3

<a id="canonical-4888fa14ccc4d9af7d972b650a186dd97703689ae7eeb6ac07a05dd93819c005"></a>

<a id="canonical-0dc60b2f754be1e915bac3d56418c6043551a96366e27688750c9909aa294d67"></a>

## name property — healthcheck / c913880dcad5 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-9e70c1098ac96b5d6f83756de984505528bd0a1b8ffb1f20b7a1e5811a68f6b3"></a>

<a id="canonical-298d8839ef9dcc93b56b5395518ddab61d22cdcadf8cd2a74455dd69f6951e9e"></a>

## namespace property — healthcheck / c913880dcad5 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-ff6ae18922ad492377cbe9e05a957f6ce918783c122f0084d51356bb4f1bdd37"></a>

<a id="canonical-2fdc01f1a0be985b85ad95c917d887f086821de9449400b5d17a165915fb46a3"></a>

## tenant property — healthcheck / c913880dcad5 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2ee5f8b6fcd9a4fe212ddd0c59e78ceb8c326d3e9826d8d261f360f143b93f35"></a>

## Next pages — healthcheck / c913880dcad5 / 7

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-b9f7675bc66fd490a145bca6473577ae0bfa83dfe9c0163a5d8280a3990293f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-666b0f7d8f1335502655d1a02fdfacb5e55ac5437a2451bea4acbf9c5d23f98b"></a>

## lb_port — lb_port / 3995ce564848 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- lb_port

<a id="canonical-fee3cd22c1f53ee4104f51d80243405380ff4038c42a4d9c573be480a4b1c80a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-23ddcd5bf612ffaf4edc604025c8533a48e6bc11f17ccf68344afd0788221a9f"></a>

## Direct properties — lb_port / 3995ce564848 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57e275e941e5eede7d241c5abb1fd226c69c2b215d6d8bbd313d8bc5e6644f86"></a>

## Next pages — lb_port / 3995ce564848 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-c30a60b3744a38686e8b5447b85e7d12379a5a560a12b3df726b08a42f2337ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-756fbb420845596ce0473407227e072da67cd71b1a995cdcb788512fb0d9fb61"></a>

## no_tls — no_tls / 1763565b61db / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- no_tls

<a id="canonical-c31d7aa8a56d1aa9f182420c47cb0998925f802ffbb1db8676c2d1136f44eaaa"></a>

Type: `["object", {}]`. Computed.

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

- [no_tls](data-sources--origin_pool--reference--group-001.md#canonical-c31d7aa8a56d1aa9f182420c47cb0998925f802ffbb1db8676c2d1136f44eaaa)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-02a9345e3df45300299a1b27753191de2b20bdfeb558d4a0de38bf539a781e09)

Select alternatives according to the provider validators above.

<a id="canonical-4c9ce6728f66524afc0c0c17ae408ca3af5fa9c012cc94ba171393afdfd60d44"></a>

## Direct properties — no_tls / 1763565b61db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73e5f3ff57665a10dc98134379dad96c22b05d3651e470c9b31049ad72da3646"></a>

## Next pages — no_tls / 1763565b61db / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0358e2dfb8cbb3319cc1bb6b85e92f244d31824ca2b15b73e314e16a49bc32ae"></a>

## origin_servers — origin_servers / 8a158590f277 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- origin_servers

<a id="canonical-6202f7dfc72ee06927396d7ea6f61bca050257a7ce50c5ac98d6ca4c8159907d"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

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

<a id="canonical-42e3798c2937c85b0f3a0824bf033b1488336ccbd35be88085c896f962618f32"></a>

## Direct properties — origin_servers / 8a158590f277 / 3

- [cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-5d2abc591750e1b6d4e5c5d475aff7740c27241f6bcae6ca79d2fa2dd3b4627d): complete subsection reference.

- [consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811): complete subsection reference.

- [custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-db084b1865c022fee81a4c1d788a9ac1e8c81ee38a27f7ce6dae2ea733e679ad): complete subsection reference.

- [k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1): complete subsection reference.

<a id="canonical-525320e2af59490223d3f3006ecc645ae0fafea96079e1a703539e97d7cd1f91"></a>

<a id="canonical-2d567ad90600acb6e40d92c7e4f1f36ff49926026bfb219104fcb445ec7b1d43"></a>

## labels property — origin_servers / 8a158590f277 / 4

Type: `["map", "string"]`. Computed.

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

- [private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52): complete subsection reference.

- [private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4): complete subsection reference.

- [public_ip](data-sources--origin_pool--reference--group-002.md#canonical-efafe219d02588b88db03465cb41b4d228fdc93ba156e95a4f9774d7c13e45f9): complete subsection reference.

- [public_name](data-sources--origin_pool--reference--group-002.md#canonical-cb646c3857fcf36e43739596634a67807483b7ec8de38ca2d9dbe38dbb706755): complete subsection reference.

- [vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-5edbfdaafe6b2fa0b48a8773b6ef085a8ce8bee684f9d6c33db4a37c37bd7262): complete subsection reference.

- [vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-fa5838e570f4c510e55c3914c3f94319a049dfc3fb3ca9a70f4f2a3a098cae3c): complete subsection reference.

<a id="canonical-304cd6644f066d0008512d377c43d51c2bc297c97cf1b34056063ec01a082852"></a>

## Next pages — origin_servers / 8a158590f277 / 5

- [origin_servers.cbip_service](data-sources--origin_pool--reference--group-001.md#canonical-5d2abc591750e1b6d4e5c5d475aff7740c27241f6bcae6ca79d2fa2dd3b4627d)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-db084b1865c022fee81a4c1d788a9ac1e8c81ee38a27f7ce6dae2ea733e679ad)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [origin_servers.public_ip](data-sources--origin_pool--reference--group-002.md#canonical-efafe219d02588b88db03465cb41b4d228fdc93ba156e95a4f9774d7c13e45f9)
- [origin_servers.public_name](data-sources--origin_pool--reference--group-002.md#canonical-cb646c3857fcf36e43739596634a67807483b7ec8de38ca2d9dbe38dbb706755)
- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-5edbfdaafe6b2fa0b48a8773b6ef085a8ce8bee684f9d6c33db4a37c37bd7262)
- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-fa5838e570f4c510e55c3914c3f94319a049dfc3fb3ca9a70f4f2a3a098cae3c)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-5d2abc591750e1b6d4e5c5d475aff7740c27241f6bcae6ca79d2fa2dd3b4627d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49dff350c2da4e1d807420b57f371f8d27e30fba7f28a6f34f7380ca05a8b422"></a>

## origin_servers.cbip_service — origin_servers.cbip_service / 0b141ad77c89 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.cbip_service

<a id="canonical-2f3bdc3bea748f677670a901a6f2e6f24f63d137a12048f50150080a91065a66"></a>

Type: `"single"`. Computed.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

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

<a id="canonical-30967f6a046be2c1ffbac17930dae27563a8f1c763143132a1b4d0b639fcbc3c"></a>

## Direct properties — origin_servers.cbip_service / 0b141ad77c89 / 3

<a id="canonical-ff1a0428e5163c350b7c56c2c0406ce65a520014bc8d623489e009076e56bd12"></a>

<a id="canonical-30cdc4e5feeeb862068aa4073a9c8523338d53a36df00ac844b30ea9067a26aa"></a>

## service_name property — origin_servers.cbip_service / 0b141ad77c89 / 4

Type: `"string"`. Computed.

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

<a id="canonical-0c0049f8545f7a7ba5dbfef2d53b62a90f7806fd82e90e386aa9d0ae7bfef2fe"></a>

## Next pages — origin_servers.cbip_service / 0b141ad77c89 / 5

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06cc5ca4d94b235fbb734361020683bd7a3fdba30cf2617f495dac717170f84e"></a>

## origin_servers.consul_service — origin_servers.consul_service / 5dc0e7ab30db / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.consul_service

<a id="canonical-24d102c4c185943baf5675605435c6d4478cc07836f4082b8b652a711ac9830b"></a>

Type: `"single"`. Computed.

Specify origin server with HashiCorp Consul service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

<a id="canonical-5225b768f3029e852f90832dc70e904a431c7d58818362a254bfc87323775774"></a>

## Direct properties — origin_servers.consul_service / 5dc0e7ab30db / 3

- [inside_network](data-sources--origin_pool--reference--group-001.md#canonical-75a6f45909da2037ed57e48809ba9447d1ad84d3dc5527bcdebacd67145d765a): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-001.md#canonical-5c28b654e798ad91ffa4c35b7928ec292d4d0b9d2099490c62d58c076ed98cd7): complete subsection reference.

<a id="canonical-7065af5a5b0467e48088de787b049905650745d76c244023b6dfae255695aead"></a>

<a id="canonical-4258932457b7779097bb53d2e7eaa8e614d6e4053441964c770dc29b49e04cb3"></a>

## service_name property — origin_servers.consul_service / 5dc0e7ab30db / 4

Type: `"string"`. Computed.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9): complete subsection reference.

<a id="canonical-eb6194fb60435363a44fa585b1996e028fa63adaf6f4516c01dd6aad4541b5d1"></a>

## Next pages — origin_servers.consul_service / 5dc0e7ab30db / 5

- [origin_servers.consul_service.inside_network](data-sources--origin_pool--reference--group-001.md#canonical-75a6f45909da2037ed57e48809ba9447d1ad84d3dc5527bcdebacd67145d765a)
- [origin_servers.consul_service.outside_network](data-sources--origin_pool--reference--group-001.md#canonical-5c28b654e798ad91ffa4c35b7928ec292d4d0b9d2099490c62d58c076ed98cd7)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-75a6f45909da2037ed57e48809ba9447d1ad84d3dc5527bcdebacd67145d765a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d37b8464f1e3af3044cdc7b17df0d9c11f7a1a5b6dd727b47ed2799b7d2d609"></a>

## origin_servers.consul_service.inside_network — origin_servers.consul_service.inside_network / 56dd50f658cb / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- origin_servers.consul_service.inside_network

<a id="canonical-78b6575736f10dd8007ef49b2eda47706267b8d2fe9b6bdd3eafe2f917adadd5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-b1c552bb2f2d0ea8ed254434a1a7f65b02097f8be3eb34d6d6848594eb7eb6cd"></a>

## Direct properties — origin_servers.consul_service.inside_network / 56dd50f658cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09fac9a04915871b398c3ab2717cb34c132b7df2282d241fe67f7033245b72bc"></a>

## Next pages — origin_servers.consul_service.inside_network / 56dd50f658cb / 4

- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-5c28b654e798ad91ffa4c35b7928ec292d4d0b9d2099490c62d58c076ed98cd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
