---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2da9d72f8a9f57a167349d6c303f7806b1409d781fc4f4256b29be5f45ce8de3"></a>

## Property reference — Property reference / 410000157222 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- Property reference

<a id="canonical-72d7e2c3b84eee809c620d58d630a0baa110d1121553cb632b22fab02ae86ca0"></a>

## Direct properties — Property reference / 410000157222 / 3

<a id="canonical-51e33a3a8bd8dbdae0edc7eab347ad0dca1cf67e33539a08a39fef0b39667727"></a>

<a id="canonical-1c7de4c3ea555d456466852f05402b45e1e085d88fef1e8f1220897198224243"></a>

## annotations property — Property reference / 410000157222 / 4

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

- [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-19a0f045eccbf77d58a1cc50bd892ef1d3ba500b6191887622c5dafa890ea5bd): complete subsection reference.

- [circuit_breaker](data-sources--cluster--reference--group-001.md#canonical-4cd7a43609dc5f0405313ba87f5e77c02b259b77a73a09576d491775ff087126): complete subsection reference.

<a id="canonical-80e4ac82210f6a25e53a6342c6475b67a19cf035f0ebc65682af4ba69b85984a"></a>

<a id="canonical-77a9c6f957801d1f6cda3b60838ba1c7f8a70a35425773bad6b9f7fac4543d87"></a>

## connection_timeout property — Property reference / 410000157222 / 5

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

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

- [default_subset](data-sources--cluster--reference--group-001.md#canonical-e540f982fbd34173ed201ef4dfe41db5cd6d2dcbf918641eed0b683d735e7d9d): complete subsection reference.

<a id="canonical-28f8ff3633709e422fc2d0bb121bdb8e060fab202550b9d170b8987de0e775bf"></a>

<a id="canonical-1cfcdacd27201b1b5ca3124d8aef73188fd4f777271f47f4002c31ef09c5ec81"></a>

## description property — Property reference / 410000157222 / 6

Type: `"string"`. Computed.

Description of the Cluster.

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

- [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-ce27ba8b97f4471f68fe562176a2b75d7a3f32f9abc1c43241a36abe475b6f9d): complete subsection reference.

<a id="canonical-0dc39279ccaf0e769559806eaa82c252593806ed1e79e5df090e76a3f74fbe5b"></a>

<a id="canonical-522eee54f7430cd316667e8c4c9792efa17fd1ed235aa81d6769cfc73dc4a34c"></a>

## endpoint_selection property — Property reference / 410000157222 / 7

Type: `"string"`. Computed.

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

- [endpoint_subsets](data-sources--cluster--reference--group-001.md#canonical-087125b8ea746e0097e7c3c1c62aa7046c5b2b2fb3383aae7b04dbd18304075c): complete subsection reference.

- [endpoints](data-sources--cluster--reference--group-001.md#canonical-16174666140d7d2abb9ea2cf39956ceca78806b911a2c604e94b31bcb80bca35): complete subsection reference.

<a id="canonical-223280ef5d433bd255b84c2c1f747a8a27ebc915aa8c26ff8a5df56634f89bba"></a>

<a id="canonical-f34845fe844e56fa73de599d145fb7f1797198e9d7c0f330a8994f6f5a4949f5"></a>

## fallback_policy property — Property reference / 410000157222 / 8

Type: `"string"`. Computed.

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

- [health_checks](data-sources--cluster--reference--group-001.md#canonical-e219769449ba4dbfd720971cb93fe586f9e87dd8306f8ff4b7f8aeda6a4c6cc3): complete subsection reference.

- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819): complete subsection reference.

- [http2_options](data-sources--cluster--reference--group-001.md#canonical-4c12d80afa74f6a46c3c8d8315de2962fac86fac65a7b37e5bb1b14aee8f9393): complete subsection reference.

<a id="canonical-7e64b8ffdd4065b905d48e43a41bdcf014c009923040aa74dea924160e9a0449"></a>

<a id="canonical-22b7679d3fa8878314dee07c781b3ae39c2acee455d51d8411889f9e7ba4785e"></a>

## http_idle_timeout property — Property reference / 410000157222 / 9

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed.

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

<a id="canonical-3ea65953b19b1bce1c5d76b4babf60b1cf2f5d36875cb6173bdfc912bda14d60"></a>

<a id="canonical-d56207150750a04f734cead72edf29c8bc3f958d03b3b7d40b0010c4635bd2b3"></a>

## id property — Property reference / 410000157222 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0c8dfadc9e4cc3a30dbce664f8d782e5f50c5517478b86f99b7ba14ab552e90d"></a>

<a id="canonical-bdac3cb9a9c338dd8fec5274bff0bfae06fd2c5fc4597d6bb213e2a68de7b8de"></a>

## labels property — Property reference / 410000157222 / 11

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

<a id="canonical-d694b659cf42c0aebcb3e8874491bc4127c1a2190f8c453656d981845151dd48"></a>

<a id="canonical-6286e819ef9c1ca0521077e5a4e435994884888450f4c269ef7c94d1e6f56414"></a>

## loadbalancer_algorithm property — Property reference / 410000157222 / 12

Type: `"string"`. Computed.

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

<a id="canonical-94529a582a791aab6fd139233539b80403445be1dd95ae8b7660e2e9df1ca040"></a>

<a id="canonical-088d61841d7efad8be0e25b670c87243b61933d42b1d91af06f005b5d86acd54"></a>

## max_requests_per_connection property — Property reference / 410000157222 / 13

Type: `"number"`. Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\]

Sets the maximum number of requests allowed per connection to the origin server. Enter a value
&gt;=1 to define the request limit per connection.

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

- [max_requests_per_connection](data-sources--cluster--reference--group-001.md#canonical-94529a582a791aab6fd139233539b80403445be1dd95ae8b7660e2e9df1ca040)
- [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-c881f977264bc3f5c8916cc056038d53400508147bce98c3b7896994b3b8d82b)

Select alternatives according to the provider validators above.

<a id="canonical-bd7fad4f94774a774b11eebd5a846cf11522b83c2203ab1af0c039c134a96008"></a>

<a id="canonical-6ecac97145498f655687d598a4bd924b5fd03226a85de09bea51bea63aae8c05"></a>

## name property — Property reference / 410000157222 / 14

Type: `"string"`. Required.

Name of the Cluster.

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

<a id="canonical-00503c7def92cd96c6534d81f85ce8e9e74fd06b5237e89329f357e74b943b5c"></a>

<a id="canonical-a91d058f008bb2f7a6165ec768921f3119b1cce7d3196eb65fc554de87b3c81d"></a>

## namespace property — Property reference / 410000157222 / 15

Type: `"string"`. Required.

Namespace where the Cluster exists.

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

- [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-a5c6c47882277a20f493b391674b78e550ac8038309d9aac1fc22383ee9b372b): complete subsection reference.

- [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-a3e2c033a9e939aba655d167ae9a63caecc22a780fabb5f6832778a96b1d7085): complete subsection reference.

- [outlier_detection](data-sources--cluster--reference--group-001.md#canonical-691efd6b990d2daf89ff433afca8d8621938dda6c0a41534b50b5c58ee27985a): complete subsection reference.

<a id="canonical-7dc9c57c5bb749ce3a077daa7c03365564f59523ff9f6ab1e3bd766b3dadb6e5"></a>

<a id="canonical-ad5832f880837117c27d9298b83754cc115f96632c817090eb39f521eca9c6c9"></a>

## panic_threshold property — Property reference / 410000157222 / 16

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for loadbalancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-50f7ac5e3835512d3e8263a0b38860e0b7d4afc37242d802ebb4c1289e2b2f0a): complete subsection reference.

- [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-4b1a740d2ea2ad06df11a445b663bb3634c120248925bca17952965c68455255): complete subsection reference.

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506): complete subsection reference.

<a id="canonical-89bfc3880b4181e45028612de4d910ca6fc952c71d47f185805093ff4da385d8"></a>

## All schema paths — Property reference / 410000157222 / 17

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cluster--reference--group-001.md#canonical-51e33a3a8bd8dbdae0edc7eab347ad0dca1cf67e33539a08a39fef0b39667727) |
| `auto_http_config` | [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-05c0f72ce523444f6387e232ea93df7ec59fa5799da6af29ce39b2c4056014ea) |
| `circuit_breaker` | [circuit_breaker](data-sources--cluster--reference--group-001.md#canonical-c002904960e87415bf20beaa1c0c96466d77c627ae47530f46c5c32ef095e6a3) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](data-sources--cluster--reference--group-001.md#canonical-250cd53ac439d9d26d1598cc17924e6867770c78cfc08d956de147c4426a0c25) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](data-sources--cluster--reference--group-001.md#canonical-98f4ec39405570fdec01ec4115f34b76d0fdb6cb0091a521d679a3ab8cc1303e) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](data-sources--cluster--reference--group-001.md#canonical-846aaf19ba0d3690beda07341e59802f49d83995c5047c0a97546502cb5ca523) |
| `circuit_breaker.priority` | [circuit_breaker.priority](data-sources--cluster--reference--group-001.md#canonical-d7a5ba78b10f89be7b19c00b24732fa06cea61f3f0a16268886a8cb54d8a096a) |
| `circuit_breaker.retries` | [circuit_breaker.retries](data-sources--cluster--reference--group-001.md#canonical-ea0acc91a15bbe2cc599c77ee6bf89820e1f1a28bf7b5fe4654352b976380651) |
| `connection_timeout` | [connection_timeout](data-sources--cluster--reference--group-001.md#canonical-80e4ac82210f6a25e53a6342c6475b67a19cf035f0ebc65682af4ba69b85984a) |
| `default_subset` | [default_subset](data-sources--cluster--reference--group-001.md#canonical-5b23da34d01e491fba9ea8c203ef3097b68fa29628d3314cf20cf1e36b877f90) |
| `description` | [description](data-sources--cluster--reference--group-001.md#canonical-28f8ff3633709e422fc2d0bb121bdb8e060fab202550b9d170b8987de0e775bf) |
| `disable_proxy_protocol` | [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-f5764ad6af0d173469d6de38e1a92b1241ca5e080e94740a0c0cf5c9406d8346) |
| `endpoint_selection` | [endpoint_selection](data-sources--cluster--reference--group-001.md#canonical-0dc39279ccaf0e769559806eaa82c252593806ed1e79e5df090e76a3f74fbe5b) |
| `endpoint_subsets` | [endpoint_subsets](data-sources--cluster--reference--group-001.md#canonical-8b80d59e8770ab9ca00b0085176242f2a0ede81e228ca735499b406193bffa5e) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](data-sources--cluster--reference--group-001.md#canonical-ba597c16d35fced31c925c0365ad50b0fd35062ee666b766ae4a3e6881102330) |
| `endpoints` | [endpoints](data-sources--cluster--reference--group-001.md#canonical-c2671d11b0acc21a641a905fc06c4f4bb7843ed57179850aa2be8dc8d6a22195) |
| `endpoints.kind` | [endpoints.kind](data-sources--cluster--reference--group-001.md#canonical-6699fc9e71399f4fcb48b1aca6d6e71b8785a014b325084271674c307413942c) |
| `endpoints.name` | [endpoints.name](data-sources--cluster--reference--group-001.md#canonical-a261b16c3b624d07e558bdee25cc10d2eecca3ab8eff6621c135be71ad3889cd) |
| `endpoints.namespace` | [endpoints.namespace](data-sources--cluster--reference--group-001.md#canonical-65b50f0e1b865473aced0b8e4cd6d96fde8177fde27f4f03b682bc170ba6b264) |
| `endpoints.tenant` | [endpoints.tenant](data-sources--cluster--reference--group-001.md#canonical-3504aaf396a462d37fbe0696219b69ad4043520ab64d99759a815e36471cad26) |
| `endpoints.uid` | [endpoints.uid](data-sources--cluster--reference--group-001.md#canonical-485b9cf7d48bd4c6168ee15637a4b5c236a6a93bfdbd1027d845dda68b2e2526) |
| `fallback_policy` | [fallback_policy](data-sources--cluster--reference--group-001.md#canonical-223280ef5d433bd255b84c2c1f747a8a27ebc915aa8c26ff8a5df56634f89bba) |
| `health_checks` | [health_checks](data-sources--cluster--reference--group-001.md#canonical-bcb012d1ce101cdbbbab71e39461414d5e6d8b1585b9acbde5674aec27241a79) |
| `health_checks.kind` | [health_checks.kind](data-sources--cluster--reference--group-001.md#canonical-8505d9cd6532f919c0d5c7795183cbeb122021b08635a6a5620d252453191c91) |
| `health_checks.name` | [health_checks.name](data-sources--cluster--reference--group-001.md#canonical-55364b519e43fd1755b0b56992d1e92ea3b14a05b6a6d71326b7ffda802457b3) |
| `health_checks.namespace` | [health_checks.namespace](data-sources--cluster--reference--group-001.md#canonical-995645541f074edc3d86ff9ec41194a1a4e368cf587efe8d68ba5a1e5564b408) |
| `health_checks.tenant` | [health_checks.tenant](data-sources--cluster--reference--group-001.md#canonical-428aabfd4f1590107964fc151730a4da08ba49f188034ee7e4e99a08871b5b99) |
| `health_checks.uid` | [health_checks.uid](data-sources--cluster--reference--group-001.md#canonical-329ec17b44cce77510a7cc3fca7fe03718468c83172a259a3ecc6cec1b21acfb) |
| `http1_config` | [http1_config](data-sources--cluster--reference--group-001.md#canonical-b10b49903e0995eb4b34a646a943f7d6d4d313131dcfe5e5c697f357be68f88a) |
| `http1_config.header_transformation` | [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-98ae929fb3137cb4ad6d90ac43239873ec9adb35c03c27a35d32463ec09b4573) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](data-sources--cluster--reference--group-001.md#canonical-8dd997164d95d69700d57856ad2ada4bc0b04e99296fea6d4616b9024d38049a) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-dfef2fd9755cb8fe6b9ec1458b3d28512d47a14acdb24e312ebae0a81e2a8b0e) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-ebf5b40aaeb1c2006ae308359822ae8fbc21a7bbc412140c6fcbf3af19dacdcb) |
| `http2_options` | [http2_options](data-sources--cluster--reference--group-001.md#canonical-105814196f78099b8f1ff2d96d06d39aec0a92d9bb06d8f28cdb6496dbb31cd9) |
| `http2_options.enabled` | [http2_options.enabled](data-sources--cluster--reference--group-001.md#canonical-131328843d2d642a67a7ca9a784fc78fa120d2c5c6f7831164b6813e3b8f6cf6) |
| `http_idle_timeout` | [http_idle_timeout](data-sources--cluster--reference--group-001.md#canonical-7e64b8ffdd4065b905d48e43a41bdcf014c009923040aa74dea924160e9a0449) |
| `id` | [id](data-sources--cluster--reference--group-001.md#canonical-3ea65953b19b1bce1c5d76b4babf60b1cf2f5d36875cb6173bdfc912bda14d60) |
| `labels` | [labels](data-sources--cluster--reference--group-001.md#canonical-0c8dfadc9e4cc3a30dbce664f8d782e5f50c5517478b86f99b7ba14ab552e90d) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](data-sources--cluster--reference--group-001.md#canonical-d694b659cf42c0aebcb3e8874491bc4127c1a2190f8c453656d981845151dd48) |
| `max_requests_per_connection` | [max_requests_per_connection](data-sources--cluster--reference--group-001.md#canonical-94529a582a791aab6fd139233539b80403445be1dd95ae8b7660e2e9df1ca040) |
| `name` | [name](data-sources--cluster--reference--group-001.md#canonical-bd7fad4f94774a774b11eebd5a846cf11522b83c2203ab1af0c039c134a96008) |
| `namespace` | [namespace](data-sources--cluster--reference--group-001.md#canonical-00503c7def92cd96c6534d81f85ce8e9e74fd06b5237e89329f357e74b943b5c) |
| `no_panic_threshold` | [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-5f50c57db8ff143926dc81d1100e5167e21f070d6f2fe135f7b254c2e0f9fade) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-c881f977264bc3f5c8916cc056038d53400508147bce98c3b7896994b3b8d82b) |
| `outlier_detection` | [outlier_detection](data-sources--cluster--reference--group-001.md#canonical-ceb085b69b61e782c6ccb0d238a1220bcf2fa1e04c8b9c952bc50d77d7852c95) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](data-sources--cluster--reference--group-001.md#canonical-9df355eb42c74e60d1080b0a33a09735e738b70f3df82391cab0a62d6723c2de) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](data-sources--cluster--reference--group-001.md#canonical-f13aeeade8e34843527a9a2c95d11478f376024d9a6e7e255f923cee24f5abdd) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](data-sources--cluster--reference--group-001.md#canonical-397172d9b01d0f7ce418f1edc85a1085c5b618526e3515cdb25b706774329f10) |
| `outlier_detection.interval` | [outlier_detection.interval](data-sources--cluster--reference--group-001.md#canonical-91495d21758336140a3431f4adffa32bcd550ccfec2cdcb728f623ae76d8e8e7) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](data-sources--cluster--reference--group-001.md#canonical-87a9e0684ac32fdb2692d94c38ea559cf032585ab34e951328db158536a3bb89) |
| `panic_threshold` | [panic_threshold](data-sources--cluster--reference--group-001.md#canonical-7dc9c57c5bb749ce3a077daa7c03365564f59523ff9f6ab1e3bd766b3dadb6e5) |
| `proxy_protocol_v1` | [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-5d055633d9beb9112d9d729823689fc31938ed8f2175e25c30bc5158c3c71510) |
| `proxy_protocol_v2` | [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-4bd28d46bc08da311db2092a491c3b34320bed8ff5de463e8b4eba047fa2f8e5) |
| `tls_parameters` | [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-cc11fbd86e34a75f02bff9d6a9166018ec6f3a84d459ecb1c9bdfbe20c14ba38) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-c8be4d6f67b89339fd9ff4e95bb2ea09b55d4bf6551cf5e4e989f8126850863b) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](data-sources--cluster--reference--group-001.md#canonical-caf128a4dbf28d97367ae09c55b7b2ce6e158e01e3a24bde7562e002bae810a4) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](data-sources--cluster--reference--group-001.md#canonical-6af87fd9879f0ee1e48a2b970807abdfdad898a939f4619980823f56433078dc) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](data-sources--cluster--reference--group-001.md#canonical-997a67d5f0a6741ff7e2dc0581306a55ebff56f9020511bf57c79402113e83d9) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](data-sources--cluster--reference--group-001.md#canonical-abdd7a069d6d048a1de4b73f884a51f22402f851c89f2926150fcd9c433f6efe) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](data-sources--cluster--reference--group-001.md#canonical-12f7b61ec159e5fe65b06176d8bd95e98d21ea0b9bfa92e8b4a57c6358f216f4) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](data-sources--cluster--reference--group-001.md#canonical-5356bdd988ee2fcf9a183d7d2f727508c0149f92cd1af4d3b656aceb9f02abe0) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](data-sources--cluster--reference--group-001.md#canonical-3b73c46e2cc45d841f0db4226850b46ac0d8e0b28c5cb96616043b3db3c087f0) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](data-sources--cluster--reference--group-001.md#canonical-847c51c5b90bfd398a5f35e2f0e2302f0ab05f51404613611c89ba7963c4647d) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](data-sources--cluster--reference--group-001.md#canonical-62a7638249601deb2eb7794fea29e37401a6da511474cd066dbc72ced03984c0) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](data-sources--cluster--reference--group-001.md#canonical-b1cf5558ced0916f832630c1d5a3fed3fd48432473f46160782dd1be9aa03e72) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-45445938906ef13399bdb089516d3dc65199399dedd524446fc47d99721f85d0) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](data-sources--cluster--reference--group-001.md#canonical-77daa7d910bca75ed79dc552119881752d05885c5329eb458906537ec726014a) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-262c9dc2e6384045e55baa195eca61683d2b71b9f85f96627fe19163c86adc1a) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-2bece8a7b7122d192dd801bc01a4a6d3ef4c2fb4ced99057aed80fe64fbaf12e) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--reference--group-001.md#canonical-0b9bcf9880c3fbf9a85df4efec0d45129d988945db1dd4eb5353d89a1d2f8b9b) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--reference--group-001.md#canonical-f840c7aca4eb807679c7f4fb3bfcb389786f21f7d40625e6e2d70469a4d578bc) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--reference--group-001.md#canonical-d9919f337ef2a26ee3d38b8fee468f7202e271273acc55925ab9d6244c10f1ce) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--reference--group-001.md#canonical-45bec31760a3546b0e6aa7ba0ecf9e95376e426fa63b698b2100124eceeff2ce) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--reference--group-001.md#canonical-ec09b4e495caf65e9e00c4c587ae6098e584b43964fd431202d13509a796e4c0) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](data-sources--cluster--reference--group-001.md#canonical-5342de805dcfe0cbb68b04bfffabcd4add5d07367a20f7e6bad740c8b59c69be) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--cluster--reference--group-001.md#canonical-bb5f2c99e183c4f64379dcfe29929dd4a30b0ac6170c28fca1159f07bc738214) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](data-sources--cluster--reference--group-001.md#canonical-df6e2c87bd7d88a4ec18b2d01d0e9c2052d0fae95354991872fdbf59fce178bc) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-dbe7f1daee70112cdec190e5f57f5170cc64d9419bf256641494938a31c2f2f2) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--cluster--reference--group-001.md#canonical-377f4ba52b57a8fa78a2c1aa7792429cd2d7be789b30fa68222fecbb722fdd49) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--cluster--reference--group-001.md#canonical-2576235b4a618ceb3207e4e116f0232f9a47264cdf173ab99e072bb41127c35f) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--cluster--reference--group-001.md#canonical-8bcc3374debd1b423cf8b846c1d50f691deb39bf320ce36c96c35dcc08f4b968) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-74d96801150f653469a0e02bbf6998a8638eb18581e2d370ad1a0f7db726af24) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--cluster--reference--group-001.md#canonical-9934ea420113e2e439265bcec487c482da3f505d38b763158a1ffaee0cfbda98) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--cluster--reference--group-001.md#canonical-4a24126452f401606025014552eca75642c695875fc75131029be87c3ecef969) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--cluster--reference--group-001.md#canonical-10560f3db676ddd374ca3e6a83931cbbddeead3a4747269f40d92f2c80c4d339) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--cluster--reference--group-001.md#canonical-8a833864af8fddca4b7e35fe3e6f7c643fa3791004f26065e2f68fb7f27f16c7) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--cluster--reference--group-001.md#canonical-2c77c47d67addfd6df181df84791442ecab710caf62e2209cb9c1a9b2ba44553) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-68c7e1c5c88e04a5f65913ced3316cef3545992c656d2109b33f81bf5210fb70) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cluster--reference--group-001.md#canonical-330d2f8dd2dbfea27a66a3cbac6b48c2a44d1c4e21fb7ad9f130e401458b5f80) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--cluster--reference--group-001.md#canonical-30e6000739722640cc4cd01b000cd77293797b3c281b0657708bb5a015715d50) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--cluster--reference--group-001.md#canonical-545696c16763d703722fc6c1eec947acaecf80e144909d46cfa1a3afe656d21c) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--cluster--reference--group-001.md#canonical-ea4d5b54817f18c97828a54ed9b32c9ce4b975a0ac8805c78285da48d6b46c61) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--cluster--reference--group-001.md#canonical-ec7c2ae8fa9037eb41b411a7eebd89c742294653f2ac1fd3ebfe76d3ad259892) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--cluster--reference--group-001.md#canonical-b9bfbe571fa09839096dfccf5bb28d1cae2fb9b9cd2c1796f3a63ca3d5391b67) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--cluster--reference--group-001.md#canonical-b7b47fac51f89f393c9f735eeccb22ada36a944aa3bcda8bd98664b5e7d1de18) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--cluster--reference--group-001.md#canonical-2aa2e7476eea83b2ed9dc85bd2a1286bdd2b27f35576827c65139969dccdeae8) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-001.md#canonical-3a5178596b076c9ce9a08ff592658d7f61db09ae4126e2cea686b18522dfe42c) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--cluster--reference--group-001.md#canonical-87f02294e9bf854a9780eac99d43fd068a473f87966683b04545f33cc66960e3) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-4128640e942a3603c1af012764f1a723b5adb69b4610873bff03df811d185451) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-d69cf5b4f900daf02b05f2be6f6b2f7432455a484d00aa0e800a8fd7d9eb548d) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--reference--group-001.md#canonical-62e3c9f24b24545e4d85dba3209a22a9b340589d7719379b37fd09fc2f6d8502) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--reference--group-001.md#canonical-cc957b0bedbee33704a102dd604b01946a296d8278efb449a5966fa6c3e0ef3e) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--reference--group-001.md#canonical-e77510910d9dd22322aa4b090bad04f38cf6e2a2530a3618b90059dc3df07783) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--reference--group-001.md#canonical-3175c21952e72a40a76a74101434ff4f566f6237faa06c922d7b1cd90547575b) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--reference--group-001.md#canonical-bbe591597b4b319db05f854b2a8f8a056e6269fb47e05f0731346393fbe8e964) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--cluster--reference--group-001.md#canonical-a348d04a6ef48deea124eb907b96adcfd45bf96a1e437f84077ffef72b7edfb5) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--cluster--reference--group-001.md#canonical-17eda3d4cfd86c78fa92e9753262c6c690150163e187208fed30947ed6a0a03c) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-087ba200a8871664dc5de5699bd62a0244de8ad228d9ed1d7987683ddc511daf) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-0749d29a3af4dff45fc7a83bdc0161db779b88104b1a6fb05098917b04b29a19) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](data-sources--cluster--reference--group-001.md#canonical-ab0be30e4681ec95e20e3fc32c31e32429fdb6c2b7d9cbc159d830a29e5d67b8) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](data-sources--cluster--reference--group-001.md#canonical-c0e14d795f2ef821abd61ec11f690bbe390d352e02a834c1e533108116b3631a) |
| `tls_parameters.sni` | [tls_parameters.sni](data-sources--cluster--reference--group-001.md#canonical-5308596ac2825b5da64b4dff372cdc957db210b1797a142c8ceb524fac3c8423) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](data-sources--cluster--reference--group-002.md#canonical-8753f5686a4b82b8e9d0088e463c4bfc5f50f7d4daeb556935641dbc909f8315) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-0aa2f695104716438b0339fd7260249546d56960117d588d5fa935ff6994a028) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-25bc61fac320825cf64fc9d6ef46709c9d0e0486e5f462d396c1f552258cee24) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-557ed43439c68e3f9658e13df87c0faac86a857df1eab5065ca634aeaba0e771) |

<a id="canonical-a43ed2342fd7618124f51454376ae1fbe1ac4e16343a42a406578636fdc8de41"></a>

## Next pages — Property reference / 410000157222 / 18

- [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-19a0f045eccbf77d58a1cc50bd892ef1d3ba500b6191887622c5dafa890ea5bd)
- [circuit_breaker](data-sources--cluster--reference--group-001.md#canonical-4cd7a43609dc5f0405313ba87f5e77c02b259b77a73a09576d491775ff087126)
- [default_subset](data-sources--cluster--reference--group-001.md#canonical-e540f982fbd34173ed201ef4dfe41db5cd6d2dcbf918641eed0b683d735e7d9d)
- [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-ce27ba8b97f4471f68fe562176a2b75d7a3f32f9abc1c43241a36abe475b6f9d)
- [endpoint_subsets](data-sources--cluster--reference--group-001.md#canonical-087125b8ea746e0097e7c3c1c62aa7046c5b2b2fb3383aae7b04dbd18304075c)
- [endpoints](data-sources--cluster--reference--group-001.md#canonical-16174666140d7d2abb9ea2cf39956ceca78806b911a2c604e94b31bcb80bca35)
- [health_checks](data-sources--cluster--reference--group-001.md#canonical-e219769449ba4dbfd720971cb93fe586f9e87dd8306f8ff4b7f8aeda6a4c6cc3)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- [http2_options](data-sources--cluster--reference--group-001.md#canonical-4c12d80afa74f6a46c3c8d8315de2962fac86fac65a7b37e5bb1b14aee8f9393)
- [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-a5c6c47882277a20f493b391674b78e550ac8038309d9aac1fc22383ee9b372b)
- [no_request_limit_per_connection](data-sources--cluster--reference--group-001.md#canonical-a3e2c033a9e939aba655d167ae9a63caecc22a780fabb5f6832778a96b1d7085)
- [outlier_detection](data-sources--cluster--reference--group-001.md#canonical-691efd6b990d2daf89ff433afca8d8621938dda6c0a41534b50b5c58ee27985a)
- [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-50f7ac5e3835512d3e8263a0b38860e0b7d4afc37242d802ebb4c1289e2b2f0a)
- [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-4b1a740d2ea2ad06df11a445b663bb3634c120248925bca17952965c68455255)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-19a0f045eccbf77d58a1cc50bd892ef1d3ba500b6191887622c5dafa890ea5bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abc06c2fb7f861ea8c028b945a80bffd6480db291ce34d497f0fd9dc59854aa7"></a>

## auto_http_config — auto_http_config / 2fd1db7713d2 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- auto_http_config

<a id="canonical-05c0f72ce523444f6387e232ea93df7ec59fa5799da6af29ce39b2c4056014ea"></a>

Type: `["object", {}]`. Computed.

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

- [auto_http_config](data-sources--cluster--reference--group-001.md#canonical-05c0f72ce523444f6387e232ea93df7ec59fa5799da6af29ce39b2c4056014ea)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-b10b49903e0995eb4b34a646a943f7d6d4d313131dcfe5e5c697f357be68f88a)
- [http2_options](data-sources--cluster--reference--group-001.md#canonical-105814196f78099b8f1ff2d96d06d39aec0a92d9bb06d8f28cdb6496dbb31cd9)

Select alternatives according to the provider validators above.

<a id="canonical-9cd4e97931b1fa6636150f779af79f5891d6a7f3c943a77d0a8d31bdc2a76830"></a>

## Direct properties — auto_http_config / 2fd1db7713d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37fa6f42068b3bc37ede1798675fc5e4243e2f601a48e09bb3f46b4d33ef8ea5"></a>

## Next pages — auto_http_config / 2fd1db7713d2 / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-4cd7a43609dc5f0405313ba87f5e77c02b259b77a73a09576d491775ff087126"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61060a3bdc14112495148f7a0f18a2b1ee6185763718d6a1410a1f7ab2d08158"></a>

## circuit_breaker — circuit_breaker / 9b331e80422d / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- circuit_breaker

<a id="canonical-c002904960e87415bf20beaa1c0c96466d77c627ae47530f46c5c32ef095e6a3"></a>

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

<a id="canonical-6b3c6268a9a2bed0789b234098d966853a32813ecea26211001d36f8e781de7e"></a>

## Direct properties — circuit_breaker / 9b331e80422d / 3

<a id="canonical-250cd53ac439d9d26d1598cc17924e6867770c78cfc08d956de147c4426a0c25"></a>

<a id="canonical-90a94e91178f941084553a6af1653525fd5011056ca081d40980a3eb8137b2b5"></a>

## connection_limit property — circuit_breaker / 9b331e80422d / 4

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

<a id="canonical-98f4ec39405570fdec01ec4115f34b76d0fdb6cb0091a521d679a3ab8cc1303e"></a>

<a id="canonical-50767992c8579026c6afc76ce62f6b3d03c134db181e26d5e7455bb0729b17b5"></a>

## max_requests property — circuit_breaker / 9b331e80422d / 5

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

<a id="canonical-846aaf19ba0d3690beda07341e59802f49d83995c5047c0a97546502cb5ca523"></a>

<a id="canonical-8b1ecb0a2fabeac287b2c6526f492dee58070d78568ea88a8f373c257b824ac2"></a>

## pending_requests property — circuit_breaker / 9b331e80422d / 6

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

<a id="canonical-d7a5ba78b10f89be7b19c00b24732fa06cea61f3f0a16268886a8cb54d8a096a"></a>

<a id="canonical-c38169cd549ca84e9ee6caf4453f742f1a2d22fbfbd5ef771584ec220164b7e7"></a>

## priority property — circuit_breaker / 9b331e80422d / 7

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

<a id="canonical-ea0acc91a15bbe2cc599c77ee6bf89820e1f1a28bf7b5fe4654352b976380651"></a>

<a id="canonical-59917914a89ee1fb40056dae6ecb225011b7eef9e51465e7bb7e670321d03b04"></a>

## retries property — circuit_breaker / 9b331e80422d / 8

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

<a id="canonical-d8c5fc9034ca348aecafb2b57d83a53d18a680ddf45925ee2e205454e1a7eeed"></a>

## Next pages — circuit_breaker / 9b331e80422d / 9

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-e540f982fbd34173ed201ef4dfe41db5cd6d2dcbf918641eed0b683d735e7d9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f109fe33fbf0658e0b5723cd220a81cbae68bbc09204acf8f08f4f85b5b6d338"></a>

## default_subset — default_subset / 3f3f1f36af54 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- default_subset

<a id="canonical-5b23da34d01e491fba9ea8c203ef3097b68fa29628d3314cf20cf1e36b877f90"></a>

Type: `"single"`. Computed.

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

<a id="canonical-fa757fc38af67ce5b7393db50dde7b46bed88d5f43b40d580591d50135dc305f"></a>

## Direct properties — default_subset / 3f3f1f36af54 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38c09d5c89b61b483611715f1fce904a16f659111e1cc9896900111e77409884"></a>

## Next pages — default_subset / 3f3f1f36af54 / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-ce27ba8b97f4471f68fe562176a2b75d7a3f32f9abc1c43241a36abe475b6f9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a82f832b1bcc0ede441cd0be009e3d1dcc6aae7e8de5ff86891f4927640d3433"></a>

## disable_proxy_protocol — disable_proxy_protocol / dddc6598e3ef / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- disable_proxy_protocol

<a id="canonical-f5764ad6af0d173469d6de38e1a92b1241ca5e080e94740a0c0cf5c9406d8346"></a>

Type: `["object", {}]`. Computed.

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

- [disable_proxy_protocol](data-sources--cluster--reference--group-001.md#canonical-f5764ad6af0d173469d6de38e1a92b1241ca5e080e94740a0c0cf5c9406d8346)
- [proxy_protocol_v1](data-sources--cluster--reference--group-001.md#canonical-5d055633d9beb9112d9d729823689fc31938ed8f2175e25c30bc5158c3c71510)
- [proxy_protocol_v2](data-sources--cluster--reference--group-001.md#canonical-4bd28d46bc08da311db2092a491c3b34320bed8ff5de463e8b4eba047fa2f8e5)

Select alternatives according to the provider validators above.

<a id="canonical-28e08fd736d87e7adc76b5f25ced15cb72456b6faffb42c6ae7e082d40d9cf20"></a>

## Direct properties — disable_proxy_protocol / dddc6598e3ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0142c5a3bf10541bb2e80c8e64b143835ef421c071659958ca5cf7ebe675facc"></a>

## Next pages — disable_proxy_protocol / dddc6598e3ef / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-087125b8ea746e0097e7c3c1c62aa7046c5b2b2fb3383aae7b04dbd18304075c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2696a8c900c98cdab851c9438adca88ba844be5f8c2f004b5c69ff7286ed8d56"></a>

## endpoint_subsets — endpoint_subsets / c1394ba0334d / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- endpoint_subsets

<a id="canonical-8b80d59e8770ab9ca00b0085176242f2a0ede81e228ca735499b406193bffa5e"></a>

Type: `"list"`. Computed.

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

<a id="canonical-80c38407b0488942cf953edaec4d4a39e508efc7adcc9e0f56f53cdeaeb164f8"></a>

## Direct properties — endpoint_subsets / c1394ba0334d / 3

<a id="canonical-ba597c16d35fced31c925c0365ad50b0fd35062ee666b766ae4a3e6881102330"></a>

<a id="canonical-9904762ce9ff356032b1a63a1817c90c1d06e58216ee62db7ff98fb75421f9ca"></a>

## keys property — endpoint_subsets / c1394ba0334d / 4

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

<a id="canonical-75aaf8913f06eba292a5f379d1dc3039099fa7c09e86df3c411d67bcd00bf12e"></a>

## Next pages — endpoint_subsets / c1394ba0334d / 5

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-16174666140d7d2abb9ea2cf39956ceca78806b911a2c604e94b31bcb80bca35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cac40668a173181eec09e418d1bc454647073e3f9f9bf18ffcadeb44720933ed"></a>

## endpoints — endpoints / f29ef0a27768 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- endpoints

<a id="canonical-c2671d11b0acc21a641a905fc06c4f4bb7843ed57179850aa2be8dc8d6a22195"></a>

Type: `"list"`. Computed.

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

<a id="canonical-e464fce91d4575c1f39938428bb38042a2581391fbdd3160895c98075039eca6"></a>

## Direct properties — endpoints / f29ef0a27768 / 3

<a id="canonical-6699fc9e71399f4fcb48b1aca6d6e71b8785a014b325084271674c307413942c"></a>

<a id="canonical-b6f192d33c69b99d76a2f2d237a9bc123edac6228c5806334f5732a98e503d38"></a>

## kind property — endpoints / f29ef0a27768 / 4

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

<a id="canonical-a261b16c3b624d07e558bdee25cc10d2eecca3ab8eff6621c135be71ad3889cd"></a>

<a id="canonical-a15e3ab2e98723aa4f044dc8cc3521013454406b599c2a13fc9eb005652186f9"></a>

## name property — endpoints / f29ef0a27768 / 5

Type: `"string"`. Computed.

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

<a id="canonical-65b50f0e1b865473aced0b8e4cd6d96fde8177fde27f4f03b682bc170ba6b264"></a>

<a id="canonical-e1bedf570d74cb8bc8ad2b1a987cdb161a7d75d6defeeb2d877c18c2ddc0cb02"></a>

## namespace property — endpoints / f29ef0a27768 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3504aaf396a462d37fbe0696219b69ad4043520ab64d99759a815e36471cad26"></a>

<a id="canonical-6aa25506e17e0f4d7a53a33aae8a3c3f68ab81c1a8f2197e193903800816e6a5"></a>

## tenant property — endpoints / f29ef0a27768 / 7

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

<a id="canonical-485b9cf7d48bd4c6168ee15637a4b5c236a6a93bfdbd1027d845dda68b2e2526"></a>

<a id="canonical-0b8791e1597d2502dcecdc99f1560be24880572c5d47dead997cba35b9d73f8c"></a>

## uid property — endpoints / f29ef0a27768 / 8

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

<a id="canonical-ec3d24a8f4f9e656c76d907c502a29267600d66be30972defc338abce2dbb3ab"></a>

## Next pages — endpoints / f29ef0a27768 / 9

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-e219769449ba4dbfd720971cb93fe586f9e87dd8306f8ff4b7f8aeda6a4c6cc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d23160a320361b9ed050f8e1980a6706586cd06e9a0c23139e394b87275d0b0"></a>

## health_checks — health_checks / b3b2fad2199e / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- health_checks

<a id="canonical-bcb012d1ce101cdbbbab71e39461414d5e6d8b1585b9acbde5674aec27241a79"></a>

Type: `"list"`. Computed.

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

<a id="canonical-8b5482596fb9c5a204ada2fe825572d5e3685a9fe645818005219a2b4c9576d3"></a>

## Direct properties — health_checks / b3b2fad2199e / 3

<a id="canonical-8505d9cd6532f919c0d5c7795183cbeb122021b08635a6a5620d252453191c91"></a>

<a id="canonical-94abc50184ef36738c994cd890714378380a96902b8ef4bdd02c95eba849c37b"></a>

## kind property — health_checks / b3b2fad2199e / 4

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

<a id="canonical-55364b519e43fd1755b0b56992d1e92ea3b14a05b6a6d71326b7ffda802457b3"></a>

<a id="canonical-1e5ea10398e0cb776edf1575738ff36b4a886d0137894a430106ca26d5d07e87"></a>

## name property — health_checks / b3b2fad2199e / 5

Type: `"string"`. Computed.

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

<a id="canonical-995645541f074edc3d86ff9ec41194a1a4e368cf587efe8d68ba5a1e5564b408"></a>

<a id="canonical-d14a51d2c1881bd0b8eac5d282d53ede35b45d87960000673dff1b74804e222e"></a>

## namespace property — health_checks / b3b2fad2199e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-428aabfd4f1590107964fc151730a4da08ba49f188034ee7e4e99a08871b5b99"></a>

<a id="canonical-6a994cf5a3e4eeb3b60361a1c6715d3f48fa9dd32cd3162361035a1ce2243bed"></a>

## tenant property — health_checks / b3b2fad2199e / 7

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

<a id="canonical-329ec17b44cce77510a7cc3fca7fe03718468c83172a259a3ecc6cec1b21acfb"></a>

<a id="canonical-4e0cf30c24c004231f5c367c0f177db052b2f07de17e6920f69ba64947cb4a40"></a>

## uid property — health_checks / b3b2fad2199e / 8

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

<a id="canonical-70b4543c7273c3a45219b5ef40eb60325a902d38e1b37c86f2bc65da19dfa352"></a>

## Next pages — health_checks / b3b2fad2199e / 9

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1014a6c520be4e0f34c0087c3364543cffea158360a705a8b1412cf363e1286b"></a>

## http1_config — http1_config / 05ba10a5c9a6 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- http1_config

<a id="canonical-b10b49903e0995eb4b34a646a943f7d6d4d313131dcfe5e5c697f357be68f88a"></a>

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

<a id="canonical-b289536cbb2cd13769f97ee1419e14b5a1affc1c9765dd99c23f9952fe7ab73b"></a>

## Direct properties — http1_config / 05ba10a5c9a6 / 3

- [header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555): complete subsection reference.

<a id="canonical-6dd1b4ff4a1756f0b2111baf8855c52c74fe152f0216d53e8d5544c6096da438"></a>

## Next pages — http1_config / 05ba10a5c9a6 / 4

- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2b7075d264adf1615f18664f9b8781cfce673ec3cbf84dedc3cedbcbee6f5a1"></a>

## http1_config.header_transformation — http1_config.header_transformation / 2ff62705ac25 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- http1_config.header_transformation

<a id="canonical-98ae929fb3137cb4ad6d90ac43239873ec9adb35c03c27a35d32463ec09b4573"></a>

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

<a id="canonical-8d8b67e5220498554da5472ba65f7eb50ff109564693b52bbb3adf64f52d9f07"></a>

## Direct properties — http1_config.header_transformation / 2ff62705ac25 / 3

- [default_header_transformation](data-sources--cluster--reference--group-001.md#canonical-77134cb3acbe4cf5c546ebb28bb5d358356bb66ad0e985c0bd7b27c2c27506d3): complete subsection reference.

- [preserve_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-a8fb3dcdb094d96610a97bc4d32cf07d432c4759a5cb7489a3e5a7af10457b87): complete subsection reference.

- [proper_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-7983147fb3cf52d60ee8ae63e32cd41dfb5f94c87b3cb66837baba9f0862ffa2): complete subsection reference.

<a id="canonical-7fe6cc651224b4c76063e483bda61532e2b5603c626fbff1708fa9b81c6c19c6"></a>

## Next pages — http1_config.header_transformation / 2ff62705ac25 / 4

- [http1_config.header_transformation.default_header_transformation](data-sources--cluster--reference--group-001.md#canonical-77134cb3acbe4cf5c546ebb28bb5d358356bb66ad0e985c0bd7b27c2c27506d3)
- [http1_config.header_transformation.preserve_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-a8fb3dcdb094d96610a97bc4d32cf07d432c4759a5cb7489a3e5a7af10457b87)
- [http1_config.header_transformation.proper_case_header_transformation](data-sources--cluster--reference--group-001.md#canonical-7983147fb3cf52d60ee8ae63e32cd41dfb5f94c87b3cb66837baba9f0862ffa2)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-77134cb3acbe4cf5c546ebb28bb5d358356bb66ad0e985c0bd7b27c2c27506d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-559072945a3d93f3c1aeeee805171d4c859f57eb505c8d55b22f2fddcebfbb54"></a>

## http1_config.header_transformation.default_header_transformation — http1_config.header_transformation.default_header_transformation / 8a3b7f50f6aa / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- http1_config.header_transformation.default_header_transformation

<a id="canonical-8dd997164d95d69700d57856ad2ada4bc0b04e99296fea6d4616b9024d38049a"></a>

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

<a id="canonical-a45f8f06b80c68cbc8f2d3ddc9cee983e2ba058673ffccca2eb70877e2f73716"></a>

## Direct properties — http1_config.header_transformation.default_header_transformation / 8a3b7f50f6aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aabc896d643336bb8fbfe27c62a14e93acccc93ac95ee871fb8f60e6df945606"></a>

## Next pages — http1_config.header_transformation.default_header_transformation / 8a3b7f50f6aa / 4

- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-a8fb3dcdb094d96610a97bc4d32cf07d432c4759a5cb7489a3e5a7af10457b87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e988c7c62d70bcb17f1c19c4d08a6fec4e4a094dce67b7c8d3c0a95dc0eda9cc"></a>

## http1_config.header_transformation.preserve_case_header_transformation — http1_config.header_transformation.preserve_case_header_transformation / 4e19ff618084 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-dfef2fd9755cb8fe6b9ec1458b3d28512d47a14acdb24e312ebae0a81e2a8b0e"></a>

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

<a id="canonical-0c9192c6a156006d268b823b17fb63f7b10f17ba23894a54e098f3bf3415d910"></a>

## Direct properties — http1_config.header_transformation.preserve_case_header_transformation / 4e19ff618084 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c837d247f0fc07c7f5a8c2ab9fc39d9c59e6d6b7cfeaddb4a6e563940dc0efa"></a>

## Next pages — http1_config.header_transformation.preserve_case_header_transformation / 4e19ff618084 / 4

- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-7983147fb3cf52d60ee8ae63e32cd41dfb5f94c87b3cb66837baba9f0862ffa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96d865dc2d8b22b2bac3891a7e8c31ce8bf2243588cdc86239c2fcfd95e90576"></a>

## http1_config.header_transformation.proper_case_header_transformation — http1_config.header_transformation.proper_case_header_transformation / b0ed61927184 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [http1_config](data-sources--cluster--reference--group-001.md#canonical-2da4f60ef4b96e94a7d48a97edf2dbda3034af03a9335cd35a0bccfdb1d76819)
- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-ebf5b40aaeb1c2006ae308359822ae8fbc21a7bbc412140c6fcbf3af19dacdcb"></a>

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

<a id="canonical-1d41f07d4e591b6daaa5807e30c547e6b45981ed6c4197ae5037f5ef13f3da93"></a>

## Direct properties — http1_config.header_transformation.proper_case_header_transformation / b0ed61927184 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-202c498f2a13df0403f5963f058f0647035d6dee1c903c17efbcd42acbc0c3f0"></a>

## Next pages — http1_config.header_transformation.proper_case_header_transformation / b0ed61927184 / 4

- [http1_config.header_transformation](data-sources--cluster--reference--group-001.md#canonical-d3a029a88104827f92b744dfd7f7556caff4ccfbe8b9d3ad108bce0f7bab7555)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-4c12d80afa74f6a46c3c8d8315de2962fac86fac65a7b37e5bb1b14aee8f9393"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02c78b685a42b4fa6e0418e796f5f3f2327748825aeb37d00be237f7282a7bb4"></a>

## http2_options — http2_options / 60b6b2259b38 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- http2_options

<a id="canonical-105814196f78099b8f1ff2d96d06d39aec0a92d9bb06d8f28cdb6496dbb31cd9"></a>

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

<a id="canonical-7c1324d56c7372a60874ce5775a7ba04297870bbd9723628a21729188abfc9eb"></a>

## Direct properties — http2_options / 60b6b2259b38 / 3

<a id="canonical-131328843d2d642a67a7ca9a784fc78fa120d2c5c6f7831164b6813e3b8f6cf6"></a>

<a id="canonical-58f675e01e3b83847561cb92025d3cdb04a01b4215d5626985ad744139a23285"></a>

## enabled property — http2_options / 60b6b2259b38 / 4

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

<a id="canonical-675899c61484412273bf4299fed43fae733bcd1e63104dc9251d0ada13fc34ff"></a>

## Next pages — http2_options / 60b6b2259b38 / 5

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-a5c6c47882277a20f493b391674b78e550ac8038309d9aac1fc22383ee9b372b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-576fd7f3aa8123637bec8a380bf87c1be1e5c09963cdf439e82fcdca94a5bfdf"></a>

## no_panic_threshold — no_panic_threshold / 28f201dc0006 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- no_panic_threshold

<a id="canonical-5f50c57db8ff143926dc81d1100e5167e21f070d6f2fe135f7b254c2e0f9fade"></a>

Type: `["object", {}]`. Computed.

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

- [no_panic_threshold](data-sources--cluster--reference--group-001.md#canonical-5f50c57db8ff143926dc81d1100e5167e21f070d6f2fe135f7b254c2e0f9fade)
- [panic_threshold](data-sources--cluster--reference--group-001.md#canonical-7dc9c57c5bb749ce3a077daa7c03365564f59523ff9f6ab1e3bd766b3dadb6e5)

Select alternatives according to the provider validators above.

<a id="canonical-c3b271560221ab06ebab41064b63c10aa3331c8b63e13e89d302cea2eeb84e56"></a>

## Direct properties — no_panic_threshold / 28f201dc0006 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-211b2d02db99414a874901f2444333fb1418f1dbca7c43778b582cce4c918a44"></a>

## Next pages — no_panic_threshold / 28f201dc0006 / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-a3e2c033a9e939aba655d167ae9a63caecc22a780fabb5f6832778a96b1d7085"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e034e53af4e19e3ba918f5907a4af91afd0bf362142399df860a361287fcc909"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / 75db9208e60b / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- no_request_limit_per_connection

<a id="canonical-c881f977264bc3f5c8916cc056038d53400508147bce98c3b7896994b3b8d82b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-396c809b80960fe3cc9349dc5feb1a08130b014ee77142398397b9b840f44ffd"></a>

## Direct properties — no_request_limit_per_connection / 75db9208e60b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cf17e9b339495c714806c8ef8c797017c939b15aec2e045e96f42d74c1d91b2"></a>

## Next pages — no_request_limit_per_connection / 75db9208e60b / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-691efd6b990d2daf89ff433afca8d8621938dda6c0a41534b50b5c58ee27985a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20f03a91e219c9d9992a48ce10c717c624e1d6a1477a468ff03e18d06c151870"></a>

## outlier_detection — outlier_detection / 3de726c9d508 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- outlier_detection

<a id="canonical-ceb085b69b61e782c6ccb0d238a1220bcf2fa1e04c8b9c952bc50d77d7852c95"></a>

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

<a id="canonical-29ae4130552b839c8b92a22030a797cff79fd02bbc629bfc32407246a71339b5"></a>

## Direct properties — outlier_detection / 3de726c9d508 / 3

<a id="canonical-9df355eb42c74e60d1080b0a33a09735e738b70f3df82391cab0a62d6723c2de"></a>

<a id="canonical-c1f02ef15f3f7727ec31137851e5faf0b269e8a59be09b9b07e9013eaa0672c9"></a>

## base_ejection_time property — outlier_detection / 3de726c9d508 / 4

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

<a id="canonical-f13aeeade8e34843527a9a2c95d11478f376024d9a6e7e255f923cee24f5abdd"></a>

<a id="canonical-40681e3cc93c5810c4666e9fab1ca042e613efe3e4f0c76a9208dd732ef6f4f7"></a>

## consecutive_5xx property — outlier_detection / 3de726c9d508 / 5

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

<a id="canonical-397172d9b01d0f7ce418f1edc85a1085c5b618526e3515cdb25b706774329f10"></a>

<a id="canonical-b45d4f07ca8f6a4f44809d1a7a59afbf2802de85c1bb33e69dfee5c5d03b75db"></a>

## consecutive_gateway_failure property — outlier_detection / 3de726c9d508 / 6

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

<a id="canonical-91495d21758336140a3431f4adffa32bcd550ccfec2cdcb728f623ae76d8e8e7"></a>

<a id="canonical-261337218b1f6dc698f0cc242dd8356d0e501bf817a6979644cfacaacc803691"></a>

## interval property — outlier_detection / 3de726c9d508 / 7

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

<a id="canonical-87a9e0684ac32fdb2692d94c38ea559cf032585ab34e951328db158536a3bb89"></a>

<a id="canonical-9ef18821604fff3fa739028add94099924322f3af7d573339ecf719701112938"></a>

## max_ejection_percent property — outlier_detection / 3de726c9d508 / 8

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

<a id="canonical-56eeb0eac13311a3c6a59180ccd7d61e6518e67ebc36f6dc70a44176624f7f1b"></a>

## Next pages — outlier_detection / 3de726c9d508 / 9

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-50f7ac5e3835512d3e8263a0b38860e0b7d4afc37242d802ebb4c1289e2b2f0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7c1e508744d3566a266642b23c82b650995f8177f5f7a1f40a77279ea22247f"></a>

## proxy_protocol_v1 — proxy_protocol_v1 / 8f7f9e0fd4f2 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- proxy_protocol_v1

<a id="canonical-5d055633d9beb9112d9d729823689fc31938ed8f2175e25c30bc5158c3c71510"></a>

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

<a id="canonical-8c8eefb86cd7b6f72e5e6d4c252e3fed37c828b5d3c558d11e89bb7c7a8f9a8d"></a>

## Direct properties — proxy_protocol_v1 / 8f7f9e0fd4f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed6ddf8d9bdb517ee31d42e337b5b57073c9c825ab180176228118c1183cab6b"></a>

## Next pages — proxy_protocol_v1 / 8f7f9e0fd4f2 / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-4b1a740d2ea2ad06df11a445b663bb3634c120248925bca17952965c68455255"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bd81796497272ad1971fe079549de94285a632c3e30b038b8d927aaf6e91281"></a>

## proxy_protocol_v2 — proxy_protocol_v2 / 27ba5431bbad / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- proxy_protocol_v2

<a id="canonical-4bd28d46bc08da311db2092a491c3b34320bed8ff5de463e8b4eba047fa2f8e5"></a>

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

<a id="canonical-4e41538c3861e2dd340ceb8b739b2976d39cf076ac15012d6c41ae19ae930fec"></a>

## Direct properties — proxy_protocol_v2 / 27ba5431bbad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4d917f7590c4689f63a75674a288d2a16e8241c9238f31a6d0790adf48b3989"></a>

## Next pages — proxy_protocol_v2 / 27ba5431bbad / 4

- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bf084e087c1f817702b438299339369903cc143772a2dcca2d68e1f969ec7d4"></a>

## tls_parameters — tls_parameters / 0aa2779bd65b / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- tls_parameters

<a id="canonical-cc11fbd86e34a75f02bff9d6a9166018ec6f3a84d459ecb1c9bdfbe20c14ba38"></a>

Type: `"single"`. Computed.

TLS configuration for upstream connections.

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

<a id="canonical-e2a41d6981a05743a3ad88cb22a51e0abc72f03bdf0950272ca3447a584e9595"></a>

## Direct properties — tls_parameters / 0aa2779bd65b / 3

- [cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202): complete subsection reference.

- [common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577): complete subsection reference.

- [default_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-0cce19f682bc8502139c38b72d2d5a96277693617252a9b9e0434c5aeefba0fc): complete subsection reference.

- [disable_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-dd0f740b8566bfe025fb9497f27593b6715710fcc8de31046777a5646e4f5848): complete subsection reference.

- [disable_sni](data-sources--cluster--reference--group-001.md#canonical-5c71f82444989d7f956b83513fbf71093b9fc7069a9481be26d63bd4376f5931): complete subsection reference.

<a id="canonical-c0e14d795f2ef821abd61ec11f690bbe390d352e02a834c1e533108116b3631a"></a>

<a id="canonical-37fed7998bc2aabc0d1febe81f9ae894c5a7607dac079705fe988bb1419e4082"></a>

## max_session_keys property — tls_parameters / 0aa2779bd65b / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

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

<a id="canonical-5308596ac2825b5da64b4dff372cdc957db210b1797a142c8ceb524fac3c8423"></a>

<a id="canonical-6cc01f2fb1624c64e6ba728d016dbd840893f9d7f20e3e544579d8b75dd2d664"></a>

## sni property — tls_parameters / 0aa2779bd65b / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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

- [use_host_header_as_sni](data-sources--cluster--reference--group-001.md#canonical-ab8b5e8d1ced3c8ce6cf61795c1e0ba598e0153324ae6ab8d8ee46fbed25b8dc): complete subsection reference.

<a id="canonical-eef546190401d17af78a89aab45bfb6897fdb20329e8576a7b837f16d1bf5a3d"></a>

## Next pages — tls_parameters / 0aa2779bd65b / 6

- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.default_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-0cce19f682bc8502139c38b72d2d5a96277693617252a9b9e0434c5aeefba0fc)
- [tls_parameters.disable_session_key_caching](data-sources--cluster--reference--group-001.md#canonical-dd0f740b8566bfe025fb9497f27593b6715710fcc8de31046777a5646e4f5848)
- [tls_parameters.disable_sni](data-sources--cluster--reference--group-001.md#canonical-5c71f82444989d7f956b83513fbf71093b9fc7069a9481be26d63bd4376f5931)
- [tls_parameters.use_host_header_as_sni](data-sources--cluster--reference--group-001.md#canonical-ab8b5e8d1ced3c8ce6cf61795c1e0ba598e0153324ae6ab8d8ee46fbed25b8dc)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c5f36401c665d378259d71ac30b5e65038ac05a0036da8309af7a5a0f607d6a"></a>

## tls_parameters.cert_params — tls_parameters.cert_params / 7be2f85e4e34 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.cert_params

<a id="canonical-c8be4d6f67b89339fd9ff4e95bb2ea09b55d4bf6551cf5e4e989f8126850863b"></a>

Type: `"single"`. Computed.

Certificate Parameters for authentication, TLS ciphers, and trust store.

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

<a id="canonical-cc56e34765008c88a5ff90bcec2c6c6c41fa78a20bf7bbd7a5866ed24885b177"></a>

## Direct properties — tls_parameters.cert_params / 7be2f85e4e34 / 3

- [certificates](data-sources--cluster--reference--group-001.md#canonical-e0bfceb02ad76b97e7e0331b7a1e93a6bac79aef2adc86adc3b0e47bca7256fe): complete subsection reference.

<a id="canonical-3b73c46e2cc45d841f0db4226850b46ac0d8e0b28c5cb96616043b3db3c087f0"></a>

<a id="canonical-1b67b97a5fb4b59f50e47749bef4ad0375356a46ca06dd54062871fe3745dabe"></a>

## cipher_suites property — tls_parameters.cert_params / 7be2f85e4e34 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-847c51c5b90bfd398a5f35e2f0e2302f0ab05f51404613611c89ba7963c4647d"></a>

<a id="canonical-0f025df538a158f97592ea5b8c607441b087df93a30ca7abbcd172636f8302d3"></a>

## maximum_protocol_version property — tls_parameters.cert_params / 7be2f85e4e34 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-62a7638249601deb2eb7794fea29e37401a6da511474cd066dbc72ced03984c0"></a>

<a id="canonical-03ef971a261f8f8beeae2f7ca27eadd842eca5d4db541e3e6d72ba61c238ad7e"></a>

## minimum_protocol_version property — tls_parameters.cert_params / 7be2f85e4e34 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

- [skip_server_verification](data-sources--cluster--reference--group-001.md#canonical-a3dfe6976b1a44221dd92ebec7b82dfbe51598797eb5e8abb66054163569abe4): complete subsection reference.

- [tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc): complete subsection reference.

- [volterra_trusted_ca](data-sources--cluster--reference--group-001.md#canonical-c9f0f8287cd1abdbed4abe4ccbabac5a5ae4d263b9aeed915513550fcd166fa6): complete subsection reference.

<a id="canonical-e06d44528f0092dd5bbbcb4491a685726b70c58bad6a09f76d6b7812b235905d"></a>

## Next pages — tls_parameters.cert_params / 7be2f85e4e34 / 7

- [tls_parameters.cert_params.certificates](data-sources--cluster--reference--group-001.md#canonical-e0bfceb02ad76b97e7e0331b7a1e93a6bac79aef2adc86adc3b0e47bca7256fe)
- [tls_parameters.cert_params.skip_server_verification](data-sources--cluster--reference--group-001.md#canonical-a3dfe6976b1a44221dd92ebec7b82dfbe51598797eb5e8abb66054163569abe4)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc)
- [tls_parameters.cert_params.volterra_trusted_ca](data-sources--cluster--reference--group-001.md#canonical-c9f0f8287cd1abdbed4abe4ccbabac5a5ae4d263b9aeed915513550fcd166fa6)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-e0bfceb02ad76b97e7e0331b7a1e93a6bac79aef2adc86adc3b0e47bca7256fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e98496cc6b94bb08576faa50b6534a540fb15ca28432a77e76dcc032bb07c97a"></a>

## tls_parameters.cert_params.certificates — tls_parameters.cert_params.certificates / 914d5b8ff04c / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- tls_parameters.cert_params.certificates

<a id="canonical-caf128a4dbf28d97367ae09c55b7b2ce6e158e01e3a24bde7562e002bae810a4"></a>

Type: `"list"`. Computed.

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

<a id="canonical-0ab693ee80a77824b8a1996aa501b4c5ad06bfc96ae70d7dcad791e707ebacbe"></a>

## Direct properties — tls_parameters.cert_params.certificates / 914d5b8ff04c / 3

<a id="canonical-6af87fd9879f0ee1e48a2b970807abdfdad898a939f4619980823f56433078dc"></a>

<a id="canonical-3b82077d9323562525d1d1d57232b00c7ac4086b6018c9328f405ecfd1f842cb"></a>

## kind property — tls_parameters.cert_params.certificates / 914d5b8ff04c / 4

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

<a id="canonical-997a67d5f0a6741ff7e2dc0581306a55ebff56f9020511bf57c79402113e83d9"></a>

<a id="canonical-7ac966bff33f88e77fe59fb1fc75fa4df796561083d1a9c447381b5eca4c7580"></a>

## name property — tls_parameters.cert_params.certificates / 914d5b8ff04c / 5

Type: `"string"`. Computed.

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

<a id="canonical-abdd7a069d6d048a1de4b73f884a51f22402f851c89f2926150fcd9c433f6efe"></a>

<a id="canonical-e6bef0d26d34f15afb347605fedb8910d076c7a18d568983657972a2bb0bf756"></a>

## namespace property — tls_parameters.cert_params.certificates / 914d5b8ff04c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-12f7b61ec159e5fe65b06176d8bd95e98d21ea0b9bfa92e8b4a57c6358f216f4"></a>

<a id="canonical-902badd55bce989011169bc0ce5406fb880f386a81341e648c0361e213914d83"></a>

## tenant property — tls_parameters.cert_params.certificates / 914d5b8ff04c / 7

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

<a id="canonical-5356bdd988ee2fcf9a183d7d2f727508c0149f92cd1af4d3b656aceb9f02abe0"></a>

<a id="canonical-675ffc471b797a3d5cbec739eca56e994556d21593f02f546acb71f48569dacb"></a>

## uid property — tls_parameters.cert_params.certificates / 914d5b8ff04c / 8

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

<a id="canonical-522d23852e5d6f0c351ed80fd160c9a53f45067321862c18c2e842c422fe6b2c"></a>

## Next pages — tls_parameters.cert_params.certificates / 914d5b8ff04c / 9

- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-a3dfe6976b1a44221dd92ebec7b82dfbe51598797eb5e8abb66054163569abe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68e6a550e006602cf3e0eaa596dc8b9576c4c1e77aed5ad9e55bcc0a95e0f9cd"></a>

## tls_parameters.cert_params.skip_server_verification — tls_parameters.cert_params.skip_server_verification / da6cbc511a95 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- tls_parameters.cert_params.skip_server_verification

<a id="canonical-b1cf5558ced0916f832630c1d5a3fed3fd48432473f46160782dd1be9aa03e72"></a>

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

<a id="canonical-154e46629ab16ee3365a1aa43d6c37a16bc1840aabe47f173935157cc3dbfb49"></a>

## Direct properties — tls_parameters.cert_params.skip_server_verification / da6cbc511a95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03a628ac35a7773acf7bd00ed899bc4744d78ac8fa5b62e9d7c5cc687a2504da"></a>

## Next pages — tls_parameters.cert_params.skip_server_verification / da6cbc511a95 / 4

- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2480165c8d719a85ad80b370e5ab529ce148f81dc4e47d7c2468aaa89f9e2b9"></a>

## tls_parameters.cert_params.tls_validation_params — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- tls_parameters.cert_params.tls_validation_params

<a id="canonical-45445938906ef13399bdb089516d3dc65199399dedd524446fc47d99721f85d0"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

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

<a id="canonical-4764491ff591e6a11eea64f5a57f0e05eae23e1a5549d23b71ea495ef819670e"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 3

<a id="canonical-77daa7d910bca75ed79dc552119881752d05885c5329eb458906537ec726014a"></a>

<a id="canonical-e4c64facd70112e1e35afd73ac2babb932592c7a48dccc88084a070969d1ad4b"></a>

## skip_hostname_verification property — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 4

Type: `"bool"`. Computed.

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

- [trusted_ca](data-sources--cluster--reference--group-001.md#canonical-8eb6f6fe3ba760ddf8c4704a39b38d66a7535211e1d8d5a34fdcf57cb677760f): complete subsection reference.

<a id="canonical-5342de805dcfe0cbb68b04bfffabcd4add5d07367a20f7e6bad740c8b59c69be"></a>

<a id="canonical-6941a2de2cf8a61d3a4bf35a18219f9230a19ddaa4f040c00a54007373efcc8c"></a>

## trusted_ca_url property — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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

<a id="canonical-bb5f2c99e183c4f64379dcfe29929dd4a30b0ac6170c28fca1159f07bc738214"></a>

<a id="canonical-4e720ef4eec49682fb62ffbade3f404ddbdafddf81975e055935a8f9d8acf442"></a>

## verify_subject_alt_names property — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 6

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3ac0921ef85bca9d50bd4df0f17551493fb2fab006fc64a264894ef7ba3aab65"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params / 8adc5fae7c1e / 7

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-8eb6f6fe3ba760ddf8c4704a39b38d66a7535211e1d8d5a34fdcf57cb677760f)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-8eb6f6fe3ba760ddf8c4704a39b38d66a7535211e1d8d5a34fdcf57cb677760f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-feae1ba09cc7de074b4c6e08a7539ae8c73f48549af20cac29906528e96b4922"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca — tls_parameters.cert_params.tls_validation_params.trusted_ca / 3e6a490baf8e / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="canonical-262c9dc2e6384045e55baa195eca61683d2b71b9f85f96627fe19163c86adc1a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0540a11e10ca69cf96a836b25581408a5d202e665d5fadc47620cf3ce6641dcd"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params.trusted_ca / 3e6a490baf8e / 3

- [trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-cd7f2ac2ee8fe9fa66abc9b23d52f0a1b7fbddbede996bacc46955ab29602bc7): complete subsection reference.

<a id="canonical-60fbc652788983924d2a28208d0119f7be995f6fb0404707482d5074b372fd3b"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params.trusted_ca / 3e6a490baf8e / 4

- [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-cd7f2ac2ee8fe9fa66abc9b23d52f0a1b7fbddbede996bacc46955ab29602bc7)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-cd7f2ac2ee8fe9fa66abc9b23d52f0a1b7fbddbede996bacc46955ab29602bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67ee3491967080c8d3f86f990cc5885bca51fce9099d553297688cf398ff33e7"></a>

## tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-001.md#canonical-35f32700535949ff851ba6b5af2ed2ff2b67b77b3945408a8b964bb0205a29cc)
- [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-8eb6f6fe3ba760ddf8c4704a39b38d66a7535211e1d8d5a34fdcf57cb677760f)
- tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-2bece8a7b7122d192dd801bc01a4a6d3ef4c2fb4ced99057aed80fe64fbaf12e"></a>

Type: `"list"`. Computed.

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

<a id="canonical-6cdfc6c1eebc7521681615a5944c764366c5f75d5e3e66fd41941482b794468c"></a>

## Direct properties — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 3

<a id="canonical-0b9bcf9880c3fbf9a85df4efec0d45129d988945db1dd4eb5353d89a1d2f8b9b"></a>

<a id="canonical-208c81bc680ef8ca4d8149edaf9e4799d2c8b3820853394577f26797719fbfa5"></a>

## kind property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 4

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

<a id="canonical-f840c7aca4eb807679c7f4fb3bfcb389786f21f7d40625e6e2d70469a4d578bc"></a>

<a id="canonical-6105d5d0a9afdbce5c473fac71bc571839fe79a524e9b9b78a7165ab56a50fa3"></a>

## name property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 5

Type: `"string"`. Computed.

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

<a id="canonical-d9919f337ef2a26ee3d38b8fee468f7202e271273acc55925ab9d6244c10f1ce"></a>

<a id="canonical-5e0ac9e2379952a09b41b3375c3f410651c73171b36285445e82abd6620be2c1"></a>

## namespace property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-45bec31760a3546b0e6aa7ba0ecf9e95376e426fa63b698b2100124eceeff2ce"></a>

<a id="canonical-7c9b3f7ab1db2e1c86e69ddc568d4641ee623eb0a88e22926ac83652004c5a57"></a>

## tenant property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 7

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

<a id="canonical-ec09b4e495caf65e9e00c4c587ae6098e584b43964fd431202d13509a796e4c0"></a>

<a id="canonical-69645cceb3b875ff390bb0cae3fc84f581762016acf3f182fcc49e2bedbcf2b5"></a>

## uid property — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 8

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

<a id="canonical-729833e43fcd9c9ceda191938de9f46934986f34a48cc318943e3ac568b98f77"></a>

## Next pages — tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list / b79af6de50d5 / 9

- [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-8eb6f6fe3ba760ddf8c4704a39b38d66a7535211e1d8d5a34fdcf57cb677760f)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-c9f0f8287cd1abdbed4abe4ccbabac5a5ae4d263b9aeed915513550fcd166fa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e7dbb171457365d0c13a676646905ccd40161fdb198d3f7516176fe60d95086"></a>

## tls_parameters.cert_params.volterra_trusted_ca — tls_parameters.cert_params.volterra_trusted_ca / 97a903dd2d0c / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="canonical-df6e2c87bd7d88a4ec18b2d01d0e9c2052d0fae95354991872fdbf59fce178bc"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-bf4e9196178cb383cb33bb5f9d0c2b2dec994bcb2d8623822cfd9b0eb1221235"></a>

## Direct properties — tls_parameters.cert_params.volterra_trusted_ca / 97a903dd2d0c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffb26c207d65bf399fec1b3fd771c2d0ad45e3e7a806beb8f20efa9e2cbbda5c"></a>

## Next pages — tls_parameters.cert_params.volterra_trusted_ca / 97a903dd2d0c / 4

- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-72e6ee6ce0624bbbf3232830011cc24dfde690daa1f6e7aadf7fa8260bf71202)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac0f9aec3730713d2af926f9639c111c60e456c78c44746f79de063275928560"></a>

## tls_parameters.common_params — tls_parameters.common_params / 6ed7893ab361 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.common_params

<a id="canonical-dbe7f1daee70112cdec190e5f57f5170cc64d9419bf256641494938a31c2f2f2"></a>

Type: `"single"`. Computed.

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

<a id="canonical-7f7c61d66c37bcebacb4e1f51b5aeeff0e72ee92ed7a7fe28878afe78b9f6df2"></a>

## Direct properties — tls_parameters.common_params / 6ed7893ab361 / 3

<a id="canonical-377f4ba52b57a8fa78a2c1aa7792429cd2d7be789b30fa68222fecbb722fdd49"></a>

<a id="canonical-dd89b3f932df343dfd35aa414fbe8e5a2d5d6cf27a09fff2b55a00a6cc414f8e"></a>

## cipher_suites property — tls_parameters.common_params / 6ed7893ab361 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2576235b4a618ceb3207e4e116f0232f9a47264cdf173ab99e072bb41127c35f"></a>

<a id="canonical-b0a9dc070059b564f46d15e46f41c94827fe28dfc5ed99d16796a4a9d926d75d"></a>

## maximum_protocol_version property — tls_parameters.common_params / 6ed7893ab361 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-8bcc3374debd1b423cf8b846c1d50f691deb39bf320ce36c96c35dcc08f4b968"></a>

<a id="canonical-752fed2af9cb5ecc63a84a94153df2e2296d96bb3333ea7f520db12a0b10bec9"></a>

## minimum_protocol_version property — tls_parameters.common_params / 6ed7893ab361 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

- [tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3): complete subsection reference.

- [validation_params](data-sources--cluster--reference--group-001.md#canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35): complete subsection reference.

<a id="canonical-23c1812d15d3219a3219101c92c6635df675f3da65ea31190503fa50ae3b9e70"></a>

## Next pages — tls_parameters.common_params / 6ed7893ab361 / 7

- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-001.md#canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fa53add1bfd14bfa35b4773b2f82ae91dee5efdbe91e69fbcb2b69801c3fde2"></a>

## tls_parameters.common_params.tls_certificates — tls_parameters.common_params.tls_certificates / 5fde1ed67e47 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- tls_parameters.common_params.tls_certificates

<a id="canonical-74d96801150f653469a0e02bbf6998a8638eb18581e2d370ad1a0f7db726af24"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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

<a id="canonical-6f9678a448a3a205b58beafff4c134aac92cf7b3f65ec74e660e5da80d6baa5b"></a>

## Direct properties — tls_parameters.common_params.tls_certificates / 5fde1ed67e47 / 3

<a id="canonical-9934ea420113e2e439265bcec487c482da3f505d38b763158a1ffaee0cfbda98"></a>

<a id="canonical-f9c275366dcf1797811e8c10286336e00a6de7190e7369e368f43bc4eedcec73"></a>

## certificate_url property — tls_parameters.common_params.tls_certificates / 5fde1ed67e47 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--cluster--reference--group-001.md#canonical-e5a2d2cc259440ca8ea5032350b08a027199e1e634b9032eda063f3a8fefe8e8): complete subsection reference.

<a id="canonical-8a833864af8fddca4b7e35fe3e6f7c643fa3791004f26065e2f68fb7f27f16c7"></a>

<a id="canonical-3a5c99d9c7ce098395c3b3ce23a359af10604f2ce03d17445c7356e2adc232cd"></a>

## description_spec property — tls_parameters.common_params.tls_certificates / 5fde1ed67e47 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cluster--reference--group-001.md#canonical-3f2817d5471f5da409352fc802cdaa50a9bc91506798a63d11b1ef769b599a2b): complete subsection reference.

- [private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee): complete subsection reference.

- [use_system_defaults](data-sources--cluster--reference--group-001.md#canonical-63d011c9edc38aea2581b5716d884f616ebc3c2e6f8f0b4cf49027efbdc69dfb): complete subsection reference.

<a id="canonical-24dc1dffcb75ac7b14fcac104cadcc59cbaed7e4911657af2721a896d5362f0f"></a>

## Next pages — tls_parameters.common_params.tls_certificates / 5fde1ed67e47 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--cluster--reference--group-001.md#canonical-e5a2d2cc259440ca8ea5032350b08a027199e1e634b9032eda063f3a8fefe8e8)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--cluster--reference--group-001.md#canonical-3f2817d5471f5da409352fc802cdaa50a9bc91506798a63d11b1ef769b599a2b)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--cluster--reference--group-001.md#canonical-63d011c9edc38aea2581b5716d884f616ebc3c2e6f8f0b4cf49027efbdc69dfb)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-e5a2d2cc259440ca8ea5032350b08a027199e1e634b9032eda063f3a8fefe8e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-635e4b5d2e7ec76eff0a2f1b9a11e2351bb159942e044447962298e917573e6f"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 67a515845161 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-4a24126452f401606025014552eca75642c695875fc75131029be87c3ecef969"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-3d74ac75ec4043f6bcda63a1421a4794623240934b6b6877a69dc5576da9505e"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 67a515845161 / 3

<a id="canonical-10560f3db676ddd374ca3e6a83931cbbddeead3a4747269f40d92f2c80c4d339"></a>

<a id="canonical-90aaf4f2c7e371100d9d2ab059bfb8ad1599b49b162b979a40753f1552392658"></a>

## hash_algorithms property — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 67a515845161 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-4747ea086e9867d4008491fd3aa2362803df959b946bc2ba09bec3d7f9c4126c"></a>

## Next pages — tls_parameters.common_params.tls_certificates.custom_hash_algorithms / 67a515845161 / 5

- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-3f2817d5471f5da409352fc802cdaa50a9bc91506798a63d11b1ef769b599a2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e249e389e46ddf71523f18c337bebcae3e34f578bf53eecc27a17de0d3a73448"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 126ebfec97e4 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-2c77c47d67addfd6df181df84791442ecab710caf62e2209cb9c1a9b2ba44553"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-750d371b8256aef6f8e74fffecac5504e6a1ed49451b86e128e9fe3ff32c00d3"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 126ebfec97e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6efac0d2bad93bc3e047324eb952a0b27d284f526f36092db5f3db4e6c9e2c8"></a>

## Next pages — tls_parameters.common_params.tls_certificates.disable_ocsp_stapling / 126ebfec97e4 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aac24d483f06c94a3d45609d3f28f0f25657e6aafd7c4de61d92bd0e33f4a968"></a>

## tls_parameters.common_params.tls_certificates.private_key — tls_parameters.common_params.tls_certificates.private_key / c37cedc6905d / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-68c7e1c5c88e04a5f65913ced3316cef3545992c656d2109b33f81bf5210fb70"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-fca418fc8a2c405866b128026a9e8adfd0c47ee5c59343ffca7cffec30fa3c2a"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key / c37cedc6905d / 3

- [blindfold_secret_info](data-sources--cluster--reference--group-001.md#canonical-49ebed77e8670cfad788952b94392509cebe45273171f14ed97d594f3f28a711): complete subsection reference.

- [clear_secret_info](data-sources--cluster--reference--group-001.md#canonical-ef1de58e126ff53c1d4646ddca3273c60e01cea47a001d7452b7132332720b92): complete subsection reference.

<a id="canonical-c30f81c90b52c88ef950c8ebadede3f341986eea1673f4a0a8ae8cf69f978260"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key / c37cedc6905d / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cluster--reference--group-001.md#canonical-49ebed77e8670cfad788952b94392509cebe45273171f14ed97d594f3f28a711)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--cluster--reference--group-001.md#canonical-ef1de58e126ff53c1d4646ddca3273c60e01cea47a001d7452b7132332720b92)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-49ebed77e8670cfad788952b94392509cebe45273171f14ed97d594f3f28a711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8e7e99bd624950e656a31b55804d4d42bbe1359a44006014baf9a42874996e5"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-330d2f8dd2dbfea27a66a3cbac6b48c2a44d1c4e21fb7ad9f130e401458b5f80"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-52230be616460ce217b0bf79bcedb8887e4965b589c6a9a16bd54f668dca4fcd"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 3

<a id="canonical-30e6000739722640cc4cd01b000cd77293797b3c281b0657708bb5a015715d50"></a>

<a id="canonical-d9c752a8e5db59c83efc5750ef66e784450af72cc10a62e6232f2b7a4ee6fd5b"></a>

## decryption_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 4

Type: `"string"`. Computed.

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

<a id="canonical-545696c16763d703722fc6c1eec947acaecf80e144909d46cfa1a3afe656d21c"></a>

<a id="canonical-08ed72828b2ab502cdb10926f2479f299145e2f961d96482f49c9c6f3aa46f8b"></a>

## location property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-ea4d5b54817f18c97828a54ed9b32c9ce4b975a0ac8805c78285da48d6b46c61"></a>

<a id="canonical-4bfde9a008dae7d241f97168a1328e6421219607e49bea6035e5152cdaea1dd7"></a>

## store_provider property — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 6

Type: `"string"`. Computed.

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

<a id="canonical-9040ba0b9465b7686783895778abb5e8b2f92b7b18ce1752d32acf819d3534d1"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info / 1b5128ce9681 / 7

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-ef1de58e126ff53c1d4646ddca3273c60e01cea47a001d7452b7132332720b92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80929e0cfe724338cc41f36630b0848e54dd015985b24ec82ce539250b7c5fbe"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 6a0801c497fc / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-ec7c2ae8fa9037eb41b411a7eebd89c742294653f2ac1fd3ebfe76d3ad259892"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-8c7ef57e152c632b674fbcd09b79810b75a33fb507f7109859f1cb82038047a9"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 6a0801c497fc / 3

<a id="canonical-b9bfbe571fa09839096dfccf5bb28d1cae2fb9b9cd2c1796f3a63ca3d5391b67"></a>

<a id="canonical-82f04a98d17ad80417349d61d6f77e98b6992d69833a3e4254480fa3f92598d3"></a>

## provider_ref property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 6a0801c497fc / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b7b47fac51f89f393c9f735eeccb22ada36a944aa3bcda8bd98664b5e7d1de18"></a>

<a id="canonical-6ef3f432d5a55b1e69ce70bf197ac090dde0781bfdea421d44e514bee55bb411"></a>

## url property — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 6a0801c497fc / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-db74230848e9c5ae1e615f7a0c8228597c702f183890ce88ca1ea307ae4e9216"></a>

## Next pages — tls_parameters.common_params.tls_certificates.private_key.clear_secret_info / 6a0801c497fc / 6

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-001.md#canonical-fe24897edb18b7469155b34b4d896ed9acc08b5027c5d7d49698ba15c13898ee)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-63d011c9edc38aea2581b5716d884f616ebc3c2e6f8f0b4cf49027efbdc69dfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87f898cadcf9dd7e93adf0a972f39449d744b56df2befa0fe4253f0a732d656d"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — tls_parameters.common_params.tls_certificates.use_system_defaults / 9c4c2c0a1355 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-2aa2e7476eea83b2ed9dc85bd2a1286bdd2b27f35576827c65139969dccdeae8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b8bb9cc989f6c1d8e1d48f9e7698411524dc21edea055416b1a9985bda6f5ec6"></a>

## Direct properties — tls_parameters.common_params.tls_certificates.use_system_defaults / 9c4c2c0a1355 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-818ccd5593d2a9926022cb196a5dd8723593afc4ebf6a8b4cef7861580e8af1a"></a>

## Next pages — tls_parameters.common_params.tls_certificates.use_system_defaults / 9c4c2c0a1355 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-001.md#canonical-bae65215bb61cd2f1a358469c018bfd2bc3bc13de0a80801c80e1b4741f06af3)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77976049d3e69623fa6b1e7da39241d55b128bf5a9c80a49579b44c35854cf6f"></a>

## tls_parameters.common_params.validation_params — tls_parameters.common_params.validation_params / 8221708336fc / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- tls_parameters.common_params.validation_params

<a id="canonical-3a5178596b076c9ce9a08ff592658d7f61db09ae4126e2cea686b18522dfe42c"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

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

<a id="canonical-81d33d15361f429e650ad8bd42e5195c7777bf7eabc695bf559e16406be92d9d"></a>

## Direct properties — tls_parameters.common_params.validation_params / 8221708336fc / 3

<a id="canonical-87f02294e9bf854a9780eac99d43fd068a473f87966683b04545f33cc66960e3"></a>

<a id="canonical-92d2f94730aabede26fc06aeaf1029a3233ee332250d5673f427475d96bd9254"></a>

## skip_hostname_verification property — tls_parameters.common_params.validation_params / 8221708336fc / 4

Type: `"bool"`. Computed.

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

- [trusted_ca](data-sources--cluster--reference--group-001.md#canonical-9d30e499ae71979faabf84a0c8e6efd1f8f422b52c3346935a3973086132ccad): complete subsection reference.

<a id="canonical-a348d04a6ef48deea124eb907b96adcfd45bf96a1e437f84077ffef72b7edfb5"></a>

<a id="canonical-13d304c8e76f927db3be9a938d4038a354ff4d5a599346b8dd421f1221bbdc3a"></a>

## trusted_ca_url property — tls_parameters.common_params.validation_params / 8221708336fc / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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

<a id="canonical-17eda3d4cfd86c78fa92e9753262c6c690150163e187208fed30947ed6a0a03c"></a>

<a id="canonical-3a5bc2b7077e30cb24d70ab5123a2433ab6d0d55c80f341928fe95707728f65f"></a>

## verify_subject_alt_names property — tls_parameters.common_params.validation_params / 8221708336fc / 6

Type: `["list", "string"]`. Computed.

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

<a id="canonical-885179910a70382d5d61589ca53f28935ea5f2716193948b8b42da66beaef0dc"></a>

## Next pages — tls_parameters.common_params.validation_params / 8221708336fc / 7

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-9d30e499ae71979faabf84a0c8e6efd1f8f422b52c3346935a3973086132ccad)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-9d30e499ae71979faabf84a0c8e6efd1f8f422b52c3346935a3973086132ccad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-427ca0750c3c6231927f7191ca3c138249aeaeec0bf507db62a5864d8fd94cfc"></a>

## tls_parameters.common_params.validation_params.trusted_ca — tls_parameters.common_params.validation_params.trusted_ca / ded26ac7bd38 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-001.md#canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-4128640e942a3603c1af012764f1a723b5adb69b4610873bff03df811d185451"></a>

Type: `"single"`. Computed.

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

<a id="canonical-8a3b3bd0a73409cd6ec8c96c66524385ab7c372f483d79b43715a15ccd925e8d"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca / ded26ac7bd38 / 3

- [trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-722ed113657a15fc980c0a7652b9067a7eea0905360848670ab1825fd3599b3c): complete subsection reference.

<a id="canonical-e706052e1ad977a5ecd066103488eb1d8134508823c3b2ca3f0468d2765949a8"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca / ded26ac7bd38 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--reference--group-001.md#canonical-722ed113657a15fc980c0a7652b9067a7eea0905360848670ab1825fd3599b3c)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-001.md#canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-722ed113657a15fc980c0a7652b9067a7eea0905360848670ab1825fd3599b3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f164931f544f3927f33e6e1cc527ff6fa9be65276388fad5607fbe2559fa7f2d"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [tls_parameters.common_params](data-sources--cluster--reference--group-001.md#canonical-88c4343548b8783b787e6ce1e1bd0da7a5eee148264aa732ede58ef667917577)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-001.md#canonical-20292347aa5ad3b5587cb42dedd7bb8f6463ed9ee1cc0c02a400c2b8fbba2f35)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-9d30e499ae71979faabf84a0c8e6efd1f8f422b52c3346935a3973086132ccad)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-d69cf5b4f900daf02b05f2be6f6b2f7432455a484d00aa0e800a8fd7d9eb548d"></a>

Type: `"list"`. Computed.

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

<a id="canonical-7407e2394eaf440d815d627046b59dba44e30abc5d78533fb7fc3063b4a2ffb5"></a>

## Direct properties — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 3

<a id="canonical-62e3c9f24b24545e4d85dba3209a22a9b340589d7719379b37fd09fc2f6d8502"></a>

<a id="canonical-c60c1012dab5f26a92f339f1737213dd2e78fde2b87c8c221db2c8bd0b037cf5"></a>

## kind property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 4

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

<a id="canonical-cc957b0bedbee33704a102dd604b01946a296d8278efb449a5966fa6c3e0ef3e"></a>

<a id="canonical-399ab06f29563d374abd6e3451f59f794e0f414bff649a0c1420ebd5d8fb3b79"></a>

## name property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 5

Type: `"string"`. Computed.

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

<a id="canonical-e77510910d9dd22322aa4b090bad04f38cf6e2a2530a3618b90059dc3df07783"></a>

<a id="canonical-5c98197f0f8fe766c017b8e0e3164de90a107a8e1f13996a1a8ddf8807ef788d"></a>

## namespace property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3175c21952e72a40a76a74101434ff4f566f6237faa06c922d7b1cd90547575b"></a>

<a id="canonical-5808db445e35d317a1412f0dbf9de6be4475db785e2ce2dc82eae6e64d3e6a66"></a>

## tenant property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 7

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

<a id="canonical-bbe591597b4b319db05f854b2a8f8a056e6269fb47e05f0731346393fbe8e964"></a>

<a id="canonical-9cc4dfb0fc3d294699d490f4972060827a79893c3c119d134536a96fa081508d"></a>

## uid property — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 8

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

<a id="canonical-422949c565e3e60e47693bbff5edba39804fc6fe77539e416347933da8e3ba28"></a>

## Next pages — tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list / 6b30364c80a7 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-001.md#canonical-9d30e499ae71979faabf84a0c8e6efd1f8f422b52c3346935a3973086132ccad)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-0cce19f682bc8502139c38b72d2d5a96277693617252a9b9e0434c5aeefba0fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c460972d3955a8d386e6b6dc79032386a3dc7917b484d1adc7056997f94ffa2d"></a>

## tls_parameters.default_session_key_caching — tls_parameters.default_session_key_caching / 64bb782fc1f1 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.default_session_key_caching

<a id="canonical-087ba200a8871664dc5de5699bd62a0244de8ad228d9ed1d7987683ddc511daf"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching.

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

<a id="canonical-598eff04ffbf960a87315327a0e4b9a8070d00458c9406043e3f670b9539a169"></a>

## Direct properties — tls_parameters.default_session_key_caching / 64bb782fc1f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7ec4cbe3058d9834ed5f0c5e6d90b56ef2b27f4ef5458415052a3ec72159abd"></a>

## Next pages — tls_parameters.default_session_key_caching / 64bb782fc1f1 / 4

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-dd0f740b8566bfe025fb9497f27593b6715710fcc8de31046777a5646e4f5848"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e05bc530bf857f5752a1f6d580dbbf77f9028f045963269ce8a8fd91c09e9cd8"></a>

## tls_parameters.disable_session_key_caching — tls_parameters.disable_session_key_caching / 1ef540dea5ed / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.disable_session_key_caching

<a id="canonical-0749d29a3af4dff45fc7a83bdc0161db779b88104b1a6fb05098917b04b29a19"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-f506366239e811b6d56300cd09344489f0d50a2a8cbc844e2f8a460fe99fa2fe"></a>

## Direct properties — tls_parameters.disable_session_key_caching / 1ef540dea5ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5acb24c6d4cc72c3dac8c41542ec59bd2d6a425c8bca1a8cec0d014d95de28a0"></a>

## Next pages — tls_parameters.disable_session_key_caching / 1ef540dea5ed / 4

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-5c71f82444989d7f956b83513fbf71093b9fc7069a9481be26d63bd4376f5931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7f6de15fd43f2ac1ebea9f510a99e8776a2e4788b16f1e08ae870ed65c4e575"></a>

## tls_parameters.disable_sni — tls_parameters.disable_sni / ba1e989f6e20 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.disable_sni

<a id="canonical-ab0be30e4681ec95e20e3fc32c31e32429fdb6c2b7d9cbc159d830a29e5d67b8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-ec73d5054d30e458a8a51031fe48121d9d59ee0648ce5b8e4d097378bd850070"></a>

## Direct properties — tls_parameters.disable_sni / ba1e989f6e20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6adb9c767ab31ccad08309ee2dd8be4e8f46614dc64f75a6dbcbdee94c218033"></a>

## Next pages — tls_parameters.disable_sni / ba1e989f6e20 / 4

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-ab8b5e8d1ced3c8ce6cf61795c1e0ba598e0153324ae6ab8d8ee46fbed25b8dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
