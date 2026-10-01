---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-8c6703beecfa29477132ea0f0dfd71dad9b395fd27ce05d3ff0630490a1fed09"></a>

## default_pool.advanced_options.circuit_breaker — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.circuit_breaker

<a id="canonical-19ad4b26d98c7fce7e3c70b5bd5a057a65fa7b76bedb979cb01052faaf2f80f4"></a>

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

<a id="canonical-b4b7fe4a8bda6d2c5f82999960190df9d82dc9b0d479ac63ba5e0bca99f4fab2"></a>

## Direct properties — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 3

<a id="canonical-3bf09ecd9a5a31658661826430148698d4a7a644bfa9fa73eec5289724663629"></a>

<a id="canonical-f239d7f7e5d5096fb06487f874e1eeb00307475ee4343c4c3ca2180f860370d6"></a>

## connection_limit property — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 4

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

<a id="canonical-758eba5d678c9e6f2415fadb9209bd789e9d4d43b39663affc7ec5d56f69dcd4"></a>

<a id="canonical-5ec709b3a828d0290cc1b31ba26dab95e24004d047deb460bd8ff66f5b63777e"></a>

## max_requests property — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 5

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

<a id="canonical-433051a4b806a06267667488a00aa1253c926cc63bd219ccf8e448920d5c0439"></a>

<a id="canonical-44fce78447ef24ce0f7f7c900e20223b5092644bcc49a4f0bdf45911b1cd8c9c"></a>

## pending_requests property — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 6

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

<a id="canonical-24e7ab40f27952a6c740f73c3f4abf164d022dfbb0d6545eb9597cb746d46b67"></a>

<a id="canonical-5ab5123c106a93132f0a7dabdc50645626ac11cd2f2a5db63337c23a5678f64f"></a>

## priority property — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 7

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

<a id="canonical-09f15729c5b4b2de694d4438c3155465024f7fb51edcc3631e5a682cec535031"></a>

<a id="canonical-de83553b20790c447e0562a7482231afe88d16a68922c1877f1617417d28a2c0"></a>

## retries property — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 8

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

<a id="canonical-0da605236e0b009587b8efce2a3fad5b5ab2e08916370bef68e868c4b19d4d3a"></a>

## Next pages — default_pool.advanced_options.circuit_breaker / 67f5174b783c / 9

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5e8dcb1009b2ac8b463ae599d6c4ad6f44cfcf5319123f179023e6e372d6bc83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aa5cfad803063d709e2454c24a9b0202ba9e0dd6ca93c24b70eba378839e2e2"></a>

## default_pool.advanced_options.default_circuit_breaker — default_pool.advanced_options.default_circuit_breaker / 1822eed3c062 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.default_circuit_breaker

<a id="canonical-fbfe72a1ef5dea6aa6ca1d9301eb9f845c952947c9d5953085e46d638a23ee02"></a>

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

<a id="canonical-e6d97141c02b067a67299b27cada3c3b281f266efdc8af12d20e2b78f664f95f"></a>

## Direct properties — default_pool.advanced_options.default_circuit_breaker / 1822eed3c062 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13fa707b4147cdf1eaaaa87257bdc5e4896e62b05f3aa8920d59107dc39c5d64"></a>

## Next pages — default_pool.advanced_options.default_circuit_breaker / 1822eed3c062 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5369d22b2f383f736398706d1dab0723c9ce1bc365c39df59263882d2109c493"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e83cbbaa21aeac450e4509aa4dbd8d3dbf2b7ab7458c7ab2029b6918830e434"></a>

## default_pool.advanced_options.disable_circuit_breaker — default_pool.advanced_options.disable_circuit_breaker / 1c1cff2fb1ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.disable_circuit_breaker

<a id="canonical-1fd5068094eedbf1f4b9b53c979e1419caf3a2434fbb54e4f25c0fc838412602"></a>

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

<a id="canonical-42eaabdfb6fa636e0403603c93a751a9b73d70e9f978f4eb82ce7644631b3532"></a>

## Direct properties — default_pool.advanced_options.disable_circuit_breaker / 1c1cff2fb1ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3d0f3cb6a2930da8baad04024b94aa681c78730be4aa3f43b12d002e55f2eeb"></a>

## Next pages — default_pool.advanced_options.disable_circuit_breaker / 1c1cff2fb1ce / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c8311f1076a3c891150774c73d1cc155d6b02aa5569075dbe87888088d9aa6db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-078228628a586bf234c98d1a1ec49b42dabe0dcc3d4c548aedb1a13d40a04981"></a>

## default_pool.advanced_options.disable_lb_source_ip_persistence — default_pool.advanced_options.disable_lb_source_ip_persistence / 2505ffa9dc14 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.disable_lb_source_ip_persistence

<a id="canonical-556f40d7ec0d001dcbfd5d646a87298e15e2a999d546bca98b2cf27def84844a"></a>

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

<a id="canonical-7b95009709d83ba00ab97c3e0c0fb01b41c61d28876f5264557d2d9e2d03a7d6"></a>

## Direct properties — default_pool.advanced_options.disable_lb_source_ip_persistence / 2505ffa9dc14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e4e16e4d8512b7705a768a06dc92b25dc6808f91c3293f0383b4eabe46acba5"></a>

## Next pages — default_pool.advanced_options.disable_lb_source_ip_persistence / 2505ffa9dc14 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7e1b53d3792b5137119e9bae1e0d70a7c949e29228523daff1dd68e445a13d1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-decdf516f9ebd997556156d3ce07a3f1088141e279ebc6b371a1e04a38b65373"></a>

## default_pool.advanced_options.disable_outlier_detection — default_pool.advanced_options.disable_outlier_detection / 9c3c5eb31ea8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.disable_outlier_detection

<a id="canonical-15ef3ae531f653463d41e992b9febd2d4c78b84eff41bf79db23f53a6db3b5aa"></a>

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

<a id="canonical-2c35e2d01f648ffb89a141e3772dd962411d376a6e6ecb0b97f76b6e76001b07"></a>

## Direct properties — default_pool.advanced_options.disable_outlier_detection / 9c3c5eb31ea8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fefbe1e67da2e1ba2fa4f01ddd1dd842a84c91449372909bd1b6e56f3c37ffb"></a>

## Next pages — default_pool.advanced_options.disable_outlier_detection / 9c3c5eb31ea8 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4cdfee2073154ae2f880711b670a1a723ae3292bc5effa46459f1e554c76c26c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea9264e1810b4597af5dae8f63a0c4445db9f116cd47795ce88f82481c00fad2"></a>

## default_pool.advanced_options.disable_proxy_protocol — default_pool.advanced_options.disable_proxy_protocol / 928cacafa79c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.disable_proxy_protocol

<a id="canonical-96e94ad76db3c43e2a107af435cb5c4769b794eed7ce9173a3c7ae0ff1647872"></a>

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

<a id="canonical-608f05b6dfae495d523bb07f265edad057d08df20dd986d243c77e04c0d83f48"></a>

## Direct properties — default_pool.advanced_options.disable_proxy_protocol / 928cacafa79c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f01d13ca62c9582ccf13245fcc60e66285fcc47909fc00e9d5e7554faf38c61a"></a>

## Next pages — default_pool.advanced_options.disable_proxy_protocol / 928cacafa79c / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-39fb095beca65af1fd5b74a0538b5d96d45255975d62c2b096a5c0d8f482526a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4af24f54a00e938717b32288ac96d0fa0dddecf31384c7a2a555c8a94b0a466f"></a>

## default_pool.advanced_options.disable_subsets — default_pool.advanced_options.disable_subsets / da5fe12780e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.disable_subsets

<a id="canonical-aafdd5171c7fc142bdf132ad45ea0400a47d7c0acd30ce8bd49a91f448869b93"></a>

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

<a id="canonical-de2842dba7eb8dda0f6426fa034969850f92817400e7caf22dc970007218f45a"></a>

## Direct properties — default_pool.advanced_options.disable_subsets / da5fe12780e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9ec44ca4e77e5d27b55aae64af7f7e05194466f5b19eedbb7cd49609b4b3ed5"></a>

## Next pages — default_pool.advanced_options.disable_subsets / da5fe12780e6 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c89ce3e50e5d3bca2b0c257fc95ee303b081911127a64045ce78329a94954aae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7863b35d0fdd39c642a7243c7eb7d7fc8a15e074a1f23752766c52a9474fe30"></a>

## default_pool.advanced_options.enable_lb_source_ip_persistence — default_pool.advanced_options.enable_lb_source_ip_persistence / da281fb9ae08 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.enable_lb_source_ip_persistence

<a id="canonical-94f0bdb09772b9290ba285f3894f27e6c4e91869d0c7f4f95020a8f4a1113c83"></a>

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

<a id="canonical-39a76e978aa527c56efb185459d9853c15a9fe5a3ecfce1b690bc67e28059a04"></a>

## Direct properties — default_pool.advanced_options.enable_lb_source_ip_persistence / da281fb9ae08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-311703d6ae00b130bd11af0e823f51c0b98a811bd4d635d82726d1ec026d1d00"></a>

## Next pages — default_pool.advanced_options.enable_lb_source_ip_persistence / da281fb9ae08 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e3ed5435311f51ca09e5b31beb39fee1870ba1759778be48a506708784c3d81"></a>

## default_pool.advanced_options.enable_subsets — default_pool.advanced_options.enable_subsets / c2b44786ff83 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.enable_subsets

<a id="canonical-0ae2a8e3a22444bbd7c51c83aa73af9e2dd7d88648a0a12382636f18d4ad0ae1"></a>

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

<a id="canonical-6eb7499776d306354d6cec552c2f6ec0c36ade07d486ec0bbd3e83511639f258"></a>

## Direct properties — default_pool.advanced_options.enable_subsets / c2b44786ff83 / 3

- [any_endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-0f0249274f8342233bf507cc80a26eb9d3a54a0c18edef9272d7c83983e1ab16): complete subsection reference.

- [default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-e93223b5cba3da610417fa83dc4c20fa58d7b7299630b2d63dce2955229fa166): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-99e8a9667fc907cc3186cfd7ce1201ad593411282f9002a560ec2d9a9e09ffab): complete subsection reference.

- [fail_request](data-sources--http_loadbalancer--reference--group-015.md#canonical-494c033a1db79cd4df63fcec47e37247784c33241ced4f970a061754c3eb87a7): complete subsection reference.

<a id="canonical-2c97f2881c10834925a7f42fe9593287eca041e4f6b69be227a11be3f8626f4f"></a>

## Next pages — default_pool.advanced_options.enable_subsets / c2b44786ff83 / 4

- [default_pool.advanced_options.enable_subsets.any_endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-0f0249274f8342233bf507cc80a26eb9d3a54a0c18edef9272d7c83983e1ab16)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-e93223b5cba3da610417fa83dc4c20fa58d7b7299630b2d63dce2955229fa166)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-99e8a9667fc907cc3186cfd7ce1201ad593411282f9002a560ec2d9a9e09ffab)
- [default_pool.advanced_options.enable_subsets.fail_request](data-sources--http_loadbalancer--reference--group-015.md#canonical-494c033a1db79cd4df63fcec47e37247784c33241ced4f970a061754c3eb87a7)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0f0249274f8342233bf507cc80a26eb9d3a54a0c18edef9272d7c83983e1ab16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70c7acb81a51682d5e5e174e9d78dab836bb439fe6cb1850e29817689b717c2c"></a>

## default_pool.advanced_options.enable_subsets.any_endpoint — default_pool.advanced_options.enable_subsets.any_endpoint / b337aea7c0a4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- default_pool.advanced_options.enable_subsets.any_endpoint

<a id="canonical-35c27ad5a18990a91a007b4033a4371f9c6ed6a9e885189a1a6569d3b9c390ca"></a>

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

<a id="canonical-bc6ade6463d035d53087ffef1fc39151112c2330d24ea4a8576f4e2e149fa23a"></a>

## Direct properties — default_pool.advanced_options.enable_subsets.any_endpoint / b337aea7c0a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa87ac9f1619b1b03b4f6a5e1683ebb339028b15826914ee2b120890373a9912"></a>

## Next pages — default_pool.advanced_options.enable_subsets.any_endpoint / b337aea7c0a4 / 4

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e93223b5cba3da610417fa83dc4c20fa58d7b7299630b2d63dce2955229fa166"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-406197d2f152665363cba6424307b9c7b8cce2801afd8794aa1599eef6af44b4"></a>

## default_pool.advanced_options.enable_subsets.default_subset — default_pool.advanced_options.enable_subsets.default_subset / 381cbc28308a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="canonical-a89554669f6b128cf8a5f45fb6826f25dc6ca8f956c68323aced94d0e50b3aa3"></a>

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

<a id="canonical-18d4d5cadd474583eea72aae09d8cfb10b981cd6aa6fc46dac1943a9894e66ab"></a>

## Direct properties — default_pool.advanced_options.enable_subsets.default_subset / 381cbc28308a / 3

- [default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-18938b90a66e2aa9e785fc2459dec965fdb0f983b924990ecb314dc758d0e998): complete subsection reference.

<a id="canonical-e808cca0e808ead84dd32de14bca1db1affba8ac4f5e583a419304a19dc1aeca"></a>

## Next pages — default_pool.advanced_options.enable_subsets.default_subset / 381cbc28308a / 4

- [default_pool.advanced_options.enable_subsets.default_subset.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-18938b90a66e2aa9e785fc2459dec965fdb0f983b924990ecb314dc758d0e998)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-18938b90a66e2aa9e785fc2459dec965fdb0f983b924990ecb314dc758d0e998"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a63a545e62f42e9e9f0c45e5835b7dd0ea4bf9be9308df97fbc97b23fec475ca"></a>

## default_pool.advanced_options.enable_subsets.default_subset.default_subset — default_pool.advanced_options.enable_subsets.default_subset.default_subset / 59e2cae86a82 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-e93223b5cba3da610417fa83dc4c20fa58d7b7299630b2d63dce2955229fa166)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-7160bbdb1427e9aa7a667750b9eb06af4611444986b21231f20bc2ade956b3bf"></a>

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

<a id="canonical-a295d6304cd454082902b18ba1e69188d46947427f690d6b645fc63e780b18be"></a>

## Direct properties — default_pool.advanced_options.enable_subsets.default_subset.default_subset / 59e2cae86a82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03206a8276b4b45e540018f2343ee1d2a6d26db9b7176b5419d8ae2cda611d78"></a>

## Next pages — default_pool.advanced_options.enable_subsets.default_subset.default_subset / 59e2cae86a82 / 4

- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-e93223b5cba3da610417fa83dc4c20fa58d7b7299630b2d63dce2955229fa166)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-99e8a9667fc907cc3186cfd7ce1201ad593411282f9002a560ec2d9a9e09ffab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dc116173ab9ac8435cb1136ba448a8603f9de7e1eb79130ae6466f81d5cfb4e"></a>

## default_pool.advanced_options.enable_subsets.endpoint_subsets — default_pool.advanced_options.enable_subsets.endpoint_subsets / f7e924a72326 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-e13e73c63ee5f8f3ec073b77fd1e5f87c49ae64d9a346f895f7a82f41a80b901"></a>

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

<a id="canonical-7ad79bddf2bd7bc7c22fc1243b5fe3542321d91ab91ae76d44638c34b54e86a1"></a>

## Direct properties — default_pool.advanced_options.enable_subsets.endpoint_subsets / f7e924a72326 / 3

<a id="canonical-5accc3a98af3e33452d13f73071f88f4f7a4cd0d92644996ed52009c941bb915"></a>

<a id="canonical-ee00947ad342d82df608ceacbd36e338e6baa3e03001c7827a0076c14fc4e8de"></a>

## keys property — default_pool.advanced_options.enable_subsets.endpoint_subsets / f7e924a72326 / 4

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

<a id="canonical-73411b9b916ee17f37626b9868b3d27beefb0185f75d06858fadeb22ab137095"></a>

## Next pages — default_pool.advanced_options.enable_subsets.endpoint_subsets / f7e924a72326 / 5

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-494c033a1db79cd4df63fcec47e37247784c33241ced4f970a061754c3eb87a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a85da6070c797c8f2dddddc2aa2d888d55c59c11e308a2434b220ceac1c4b3e"></a>

## default_pool.advanced_options.enable_subsets.fail_request — default_pool.advanced_options.enable_subsets.fail_request / 284b4ddda8fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- default_pool.advanced_options.enable_subsets.fail_request

<a id="canonical-b7cc3e56009343b2a579fa8f5e410c88dd1914b765dead1ba479bab01a6a2910"></a>

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

<a id="canonical-c3eea998eed2cc4f4d23f854a37a6184000fa3adbb001534a242116f0eac0f34"></a>

## Direct properties — default_pool.advanced_options.enable_subsets.fail_request / 284b4ddda8fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aec09656c1d0839457d0e22da1bc28ce9234d35890e220cfb4144fe8209636ec"></a>

## Next pages — default_pool.advanced_options.enable_subsets.fail_request / 284b4ddda8fb / 4

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-a42b09a9e56fe1e56d2e00a366cd6ed7cffd5b50c918287e6a810b2d86253c83)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69744071c74e1db21b974ef30396544350d59733d242da5c83cddf63ff31b6fc"></a>

## default_pool.advanced_options.http1_config — default_pool.advanced_options.http1_config / f2840d1c6886 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.http1_config

<a id="canonical-d99ece91f8cdba69c3b61e1b97f6d68fcf8c1ec12518306430047f3332a5041a"></a>

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

<a id="canonical-15dea861dad571ff8cf733b997f329e578b78576cb94a1fbcc6ef985c3aad04f"></a>

## Direct properties — default_pool.advanced_options.http1_config / f2840d1c6886 / 3

- [header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968): complete subsection reference.

<a id="canonical-093538f9f7bc12a356b6c7af3b4ce233f29d4ec7b04f2f6aa41ce059596264a9"></a>

## Next pages — default_pool.advanced_options.http1_config / f2840d1c6886 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9286957e73f26073e48f865f6cc90fc21a56e629b835cc0d4a8ccd73dee44bf"></a>

## default_pool.advanced_options.http1_config.header_transformation — default_pool.advanced_options.http1_config.header_transformation / 815a99d66d17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- default_pool.advanced_options.http1_config.header_transformation

<a id="canonical-e6f52549527909cdfdda587355872fe4f257dcd3af3d75cc5de9ee0a8251ddd0"></a>

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

<a id="canonical-969dd3df6d023b2a1325cd9f16c2d98eb9ae9bd102e5e88c279b75f69c47b05f"></a>

## Direct properties — default_pool.advanced_options.http1_config.header_transformation / 815a99d66d17 / 3

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-4dc5ff67b0195ac09c82aa1423d6cab29e77b1154e81d69ffe964b447fee8ec0): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-ba50bc0a56322fc56837b0580b4145b47d504b75add2fb28f479bbe4067b3d61): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-7b3ef7d6a997fe07caf7b9ea02e5ec88dd28f51c2ed1c9c83d114a16693dab35): complete subsection reference.

<a id="canonical-64549a0fc1895f56489d505b24aa49c8f97da1aef7c5e75c502a8a26c451088d"></a>

## Next pages — default_pool.advanced_options.http1_config.header_transformation / 815a99d66d17 / 4

- [default_pool.advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-4dc5ff67b0195ac09c82aa1423d6cab29e77b1154e81d69ffe964b447fee8ec0)
- [default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-ba50bc0a56322fc56837b0580b4145b47d504b75add2fb28f479bbe4067b3d61)
- [default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-7b3ef7d6a997fe07caf7b9ea02e5ec88dd28f51c2ed1c9c83d114a16693dab35)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4dc5ff67b0195ac09c82aa1423d6cab29e77b1154e81d69ffe964b447fee8ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bb44f83b0edac602eeed68c7d53f011c4251a1ff78dbb9e41e87cf36a085687"></a>

## default_pool.advanced_options.http1_config.header_transformation.default_header_transformation — default_pool.advanced_options.http1_config.header_transformation.default_header_ / 89fe638755d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- default_pool.advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-6c4eaf650fac8ed0271bc82b626abd4ae47da9d65d1c7c78ad2fd13f5531394c"></a>

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

<a id="canonical-5c24bfa7b088ef6414821f5c55c7a5e5d774492b23a4616a31d63434d3b3ee99"></a>

## Direct properties — default_pool.advanced_options.http1_config.header_transformation.default_header_ / 89fe638755d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1bde482993d83cb3fabbb0acaca963feebfbbb6e4b837ccec74c3ec2d72e4e1"></a>

## Next pages — default_pool.advanced_options.http1_config.header_transformation.default_header_ / 89fe638755d3 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ba50bc0a56322fc56837b0580b4145b47d504b75add2fb28f479bbe4067b3d61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abdce83f0db19f5bf0ec0b20f80cf75593cd10cb73c145b59bd64f239f558ee5"></a>

## default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation — default_pool.advanced_options.http1_config.header_transformation.preserve_case_h / 5e6af920b1bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-b76bb2c8a988600410b6863436b5a9022d1cff8bade645c463884ad1c96b26d4"></a>

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

<a id="canonical-ef824f8acb530778120614d75ef8c6bab5a876e06cfde7c0a699d017e2b2e17f"></a>

## Direct properties — default_pool.advanced_options.http1_config.header_transformation.preserve_case_h / 5e6af920b1bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24afc0717306024d541252dfd4c0ed7fc7ddabac2ee74a443889d5337f364b7b"></a>

## Next pages — default_pool.advanced_options.http1_config.header_transformation.preserve_case_h / 5e6af920b1bf / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7b3ef7d6a997fe07caf7b9ea02e5ec88dd28f51c2ed1c9c83d114a16693dab35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1767ea0d82c504db51c91dd1ffec5076bd93e675cae3a92a18424ba8c34b6f85"></a>

## default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation — default_pool.advanced_options.http1_config.header_transformation.proper_case_hea / df468f1cf0c7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-868058d67c9ebb526444dd67146132159600476ec72b691f7a06ed687c4b3c4b)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-31128df10082014b0ede5b413aa68bb2ec47d8778f0d3d02d26dcea441b237b7"></a>

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

<a id="canonical-affb7c1c849b09dcfb96dbc19491e66253fdb3c1334890aa06447a6c94bd80a7"></a>

## Direct properties — default_pool.advanced_options.http1_config.header_transformation.proper_case_hea / df468f1cf0c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a89c095f9c0ccd7bec73041e742d7e67d82e77e41c95654a0c7a5d10b1cc9a15"></a>

## Next pages — default_pool.advanced_options.http1_config.header_transformation.proper_case_hea / df468f1cf0c7 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-a4ba7a4143185b0d2bfbe254606772e5fd7d70942463691e6fbbeec7b6e4c968)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-324e48dbb083a33bf32131888d62b13bb782e82de3fe5e1f8f8e84454ba903c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4221bee81fdbb2ec7b0ebfb363791656a23bbb499f32ce265e3a6518fc0531d"></a>

## default_pool.advanced_options.http2_options — default_pool.advanced_options.http2_options / d457ca92a761 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.http2_options

<a id="canonical-4dd2e8f15396e1274c66135a427d35e41e57b0f8a7fb26b59bfa1d84fcefcb03"></a>

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

<a id="canonical-fdf32234914f347aee85e6a6911d8adf30877f335588900407a5f74e9ff057a3"></a>

## Direct properties — default_pool.advanced_options.http2_options / d457ca92a761 / 3

<a id="canonical-cf7866f1b137c04dcfb0e8ec98d10aae474fcc06c055b529293d5addab2b921b"></a>

<a id="canonical-92865a6c447635cfe72f605cc7aef87447c68360972082045253f5964e778f98"></a>

## enabled property — default_pool.advanced_options.http2_options / d457ca92a761 / 4

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

<a id="canonical-4dcba6db92fa2b25ce727166f0e3a7c6c1ee9ebcd7b146503b705239d34b05a7"></a>

## Next pages — default_pool.advanced_options.http2_options / d457ca92a761 / 5

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-261f8654cf935740e1982b5a54cbe968169ccc9bb826f97314724c097b01556b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca70ccb3c940dc55c1df5fb8e3443de28e7c11b8194c02b44214d3c1dfa0f8e3"></a>

## default_pool.advanced_options.no_panic_threshold — default_pool.advanced_options.no_panic_threshold / b09760f1e0e7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.no_panic_threshold

<a id="canonical-9205ee789df81c38a92cc870f1f3fce4f6a02c529ad0935ee85f167073d988c1"></a>

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

<a id="canonical-2206986618fec06eb70cd6b3d6855d390c519f997f0c8e8bd5606f9a24a830a6"></a>

## Direct properties — default_pool.advanced_options.no_panic_threshold / b09760f1e0e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45476faff5c93a2fc5184754b7128a0ddee9296ebb715bbb54c35c4ca8e1f183"></a>

## Next pages — default_pool.advanced_options.no_panic_threshold / b09760f1e0e7 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-791916a592038a377dd3ab4e636f8dd575739e3348f6dd937ccd095d0e248230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b1f779d95e69c1be66ba748cba77d55ac86b0f63164931425b70188bb2538cf"></a>

## default_pool.advanced_options.no_request_limit_per_connection — default_pool.advanced_options.no_request_limit_per_connection / 1d4b68b9fe02 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.no_request_limit_per_connection

<a id="canonical-e7dc74ff4490864dff3a37260624c9a77fbcdad0a1166016e6f9283de992eca6"></a>

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

<a id="canonical-5564587898365c91cee0b2a2bb28f01cedd32b4d050fafec6a57a5be474c7a0d"></a>

## Direct properties — default_pool.advanced_options.no_request_limit_per_connection / 1d4b68b9fe02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8163d7079cf4f7ff04d70839c3aef3b287b3711ac4d45a0726f29021b515caff"></a>

## Next pages — default_pool.advanced_options.no_request_limit_per_connection / 1d4b68b9fe02 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2f40095c75e6ce28684437106211d9abf1384088ed8542adc2abebb8a7933358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d963f4abccd07397d801484a5b443a2c594b7065d15f4cc152a076c3bfb07276"></a>

## default_pool.advanced_options.outlier_detection — default_pool.advanced_options.outlier_detection / ea11d393e368 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.outlier_detection

<a id="canonical-8967c45d0ba52e6152562438a792bb67f384553489be5c13719fd0d63a71bee2"></a>

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

<a id="canonical-d5e5a7924fac02948c179c185db5f1becdf616eba1313a5a5fd65aced017cb8c"></a>

## Direct properties — default_pool.advanced_options.outlier_detection / ea11d393e368 / 3

<a id="canonical-396e93b959996fa053d977e6a58e823da650b6211bdf248529712d472c7bf7f4"></a>

<a id="canonical-75bd4a94d4b18c87c533300ba26446c88ca4cc11a3c484350e9558bdeb61c37f"></a>

## base_ejection_time property — default_pool.advanced_options.outlier_detection / ea11d393e368 / 4

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

<a id="canonical-92bcc7757523881548dc3cae64c781d61a68960f1a1b8cbe599e96f30d6db467"></a>

<a id="canonical-5984b55032c0b4d48d10aff6b8d1c78722e6222798756c83f8e4b461a4d9dbe7"></a>

## consecutive_5xx property — default_pool.advanced_options.outlier_detection / ea11d393e368 / 5

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

<a id="canonical-ba92fe3eec9e8538583d5793556f60bc018cf6ed30b368268a2f44b5a4178362"></a>

<a id="canonical-09e335f4bd066928ad461c99681a7f82ec79ecc29822d5f2d15eab09a430c2fb"></a>

## consecutive_gateway_failure property — default_pool.advanced_options.outlier_detection / ea11d393e368 / 6

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

<a id="canonical-0a82cdb2fb83853f3c3574d9a6b4957b7f404e227dc3e1887ab867b57446fa12"></a>

<a id="canonical-7d32f46720e2c5c5258f3d6af7c02bcf3f34924c8d60b457fc955c5e53f6ab39"></a>

## interval property — default_pool.advanced_options.outlier_detection / ea11d393e368 / 7

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

<a id="canonical-951dfc484d400622a4306b9a422eabd53e1c977c76d1fd67b0e5791debff9218"></a>

<a id="canonical-15be600bad1a094cdc489f1417d83c4bef4fcd9704014fafc4ece6d38995a8f3"></a>

## max_ejection_percent property — default_pool.advanced_options.outlier_detection / ea11d393e368 / 8

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

<a id="canonical-c58f58b86547d980424e34ad9da52e472b2fd307ea83f9358cce6d3782523bb9"></a>

## Next pages — default_pool.advanced_options.outlier_detection / ea11d393e368 / 9

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-78d8a16b1644f794bace42e17864703c7de6fd05e8552ea0c9d86d3d31bce0c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99f3ad64bd8437afec48b79055abe1984648fadb8986a9be8a508fbb5595587a"></a>

## default_pool.advanced_options.proxy_protocol_v1 — default_pool.advanced_options.proxy_protocol_v1 / 21b67b93c8a2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.proxy_protocol_v1

<a id="canonical-4ac34c5447aac07379ad968df292386ee2149800c231c44cb1fe3488294d4e08"></a>

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

<a id="canonical-3384f6982b1c5aaedb0e97f3c84b3a10f426b23d801a6b4fe39831e788347c6e"></a>

## Direct properties — default_pool.advanced_options.proxy_protocol_v1 / 21b67b93c8a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65a41fa941a8a4dde38a8397bc7c0023f88ac5afdd39d6ed7ec5f6452160074a"></a>

## Next pages — default_pool.advanced_options.proxy_protocol_v1 / 21b67b93c8a2 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-254f7e4e21689359b5440db0ac52ce34e7a08e0b840191efb5144931fb63bcfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90a793ed17340e3688adf0546a146c45973e121604f8bb9eb9b845cd6cb97742"></a>

## default_pool.advanced_options.proxy_protocol_v2 — default_pool.advanced_options.proxy_protocol_v2 / 14565a51fd3e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- default_pool.advanced_options.proxy_protocol_v2

<a id="canonical-a9be18a1f9d3b320bd32905a452592589bdd1529572336fe4e890eb6fa3b2551"></a>

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

<a id="canonical-33984dd119193dc5c6c4a6b444ee4386d09965601ddd250f49b392d18498bb7c"></a>

## Direct properties — default_pool.advanced_options.proxy_protocol_v2 / 14565a51fd3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2bdfcbcbae3c5c63941010a25718593ed48ab9bc6d9b7c6c5d28923c1092798"></a>

## Next pages — default_pool.advanced_options.proxy_protocol_v2 / 14565a51fd3e / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-014.md#canonical-8e76b7d2133ce1c25e3d12a1d8444e5f80bbb7d6b1dc8b95b2127575eb520e47)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ff857da8254cb22d6d28b25edcf09e8945fc7f1069c626765bdb72560ecc69c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14308be65906c0c2975863d6589f11772ed854006b19a97a7b41eb9de84195c4"></a>

## default_pool.automatic_port — default_pool.automatic_port / 7585d58dd141 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.automatic_port

<a id="canonical-808ac2eba971570e3c5bf5e6862b1aeacbf773ce7963fbe44ff3cc3bc153cd72"></a>

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

<a id="canonical-9910f37c94287f71973cc3dd87ff0a1cb231495376ee69941886c1f58e3873fe"></a>

## Direct properties — default_pool.automatic_port / 7585d58dd141 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84cfd0cafc128ee4617c2af08d22e4cef688d9b243ed10e998b10458cf3635a4"></a>

## Next pages — default_pool.automatic_port / 7585d58dd141 / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-62f7672bfb3c3340703a1c669f486ca846258513ad74e17527fa62228935afc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f22303cb0fd16d959df032ffdaae372ea3ddd4e7d3dd3a4cf66552a03bb811d"></a>

## default_pool.healthcheck — default_pool.healthcheck / c9146e6f9e7b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.healthcheck

<a id="canonical-2af77fd550c028b0748db8c5e51e401831e10d29e05831ac1ec599985a0f04dd"></a>

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

<a id="canonical-fb7b3edc0701531856c1245de7c29ae9b0c3bd8ce005821953e5cdbc90ff6380"></a>

## Direct properties — default_pool.healthcheck / c9146e6f9e7b / 3

<a id="canonical-9f1a8231ff5e6b48ad4b089858ecbf5aa0b9b45f693c9e55bd3966399c6d7583"></a>

<a id="canonical-3b916e2122064f8e516160e56f6591765e7f2c78cc910ae479274e5317782e34"></a>

## name property — default_pool.healthcheck / c9146e6f9e7b / 4

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

<a id="canonical-354461dcf3e6c89f97bf6b34d80bb6a7bca14e6cea8444eca98547621d57cc41"></a>

<a id="canonical-b42e6fa034628e266ce446e6ca55b9e448eebb489bc957b2276a3be9c96a98b1"></a>

## namespace property — default_pool.healthcheck / c9146e6f9e7b / 5

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

<a id="canonical-d45ca34e39dc3c37bf4e7f63c982ac8c68601ead9c7763ac1f9e9a80655bdb34"></a>

<a id="canonical-1f6e6989ec5b6cdd7dbc4b8fb38635b2496af6f3973154babffd7b1b603105ad"></a>

## tenant property — default_pool.healthcheck / c9146e6f9e7b / 6

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

<a id="canonical-de1b6f1ac9cb4618e291a2534be57a5b6fb11296fd1cfcbad23ca55aa363b5c4"></a>

## Next pages — default_pool.healthcheck / c9146e6f9e7b / 7

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9ad8ab5b9aa41cd455cbf362882fa6bd081274d406011ea56066430878b87477"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caa3bfca3de356000afa8af1a7fb72a1c64f43e04b6bde046172ebccdab6e569"></a>

## default_pool.lb_port — default_pool.lb_port / b6528246247d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.lb_port

<a id="canonical-7b3aeb4f584a60f8eafbf2bfa31425ce2d3e8373ba9a27ff69dc13725c91bde1"></a>

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

<a id="canonical-1cd3130b65f39de20a55967f1ba44ddeb62c493d10e52408b408c14cb9229a08"></a>

## Direct properties — default_pool.lb_port / b6528246247d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-398211f4fe2f64b9d7c985d7aae2bbdc78c8e0d9c374271a3965385d088ac90a"></a>

## Next pages — default_pool.lb_port / b6528246247d / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-529a47404046c3510b6d8e6b1b3c029d2ae5619501b1270b94491c2b636a29de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b4853e32192d55e3e47879d15e5cf77346b69574086f3ba0d065b5c4e0f6ca0"></a>

## default_pool.no_tls — default_pool.no_tls / db53bfaa0a92 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.no_tls

<a id="canonical-841398cd5e0dcbb3f56011facedad69ee7e4dfa59bc190b3567976d0e977dc33"></a>

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

<a id="canonical-dfd6964bca74a3c44fb6f627f646eb6edbdc528c672b2abdad85e80262efef86"></a>

## Direct properties — default_pool.no_tls / db53bfaa0a92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2ad3f770bf79ece794d2bc24e4f34224173cd17e85fdc99e2368423f2980b9e"></a>

## Next pages — default_pool.no_tls / db53bfaa0a92 / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eb3da8aed12f8cd486ce89d46d694d1e5088b632390cb3812885de73bd20ac1"></a>

## default_pool.origin_servers — default_pool.origin_servers / f3b1d31a03a8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.origin_servers

<a id="canonical-e0a64a8645f26ebc45cc92355cb4bda4ce77b1d9482b72718040ce7e6352c10a"></a>

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

<a id="canonical-c822f2922c2f434921b93b15b345d12a32aba79e26c072d90b49bff93b3f3f5d"></a>

## Direct properties — default_pool.origin_servers / f3b1d31a03a8 / 3

- [cbip_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-5052793e3d8883a156f15b4de20f4e811b5cc9c36fe7357226dcd23dd696f186): complete subsection reference.

- [consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046): complete subsection reference.

- [custom_endpoint_object](data-sources--http_loadbalancer--reference--group-015.md#canonical-055a91496ac2639599fbdb0334c75702aca7edbbbe0ec456aaac14812156fa93): complete subsection reference.

- [k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9): complete subsection reference.

<a id="canonical-e81a6b921a159202a96191abb1dd9478e96742ef5dadb7aeaed6633080e878f0"></a>

<a id="canonical-b3b675e0c4e32387a1c9e36f35fb1829170cc28468c9d4d6e75042456775bfac"></a>

## labels property — default_pool.origin_servers / f3b1d31a03a8 / 4

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

- [private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f): complete subsection reference.

- [private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa): complete subsection reference.

- [public_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-ecb9b44d81af36c8674aac3d22ff09119e51f5fc66f0d2b9426df39ca56d60b3): complete subsection reference.

- [public_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-e7dee21ac44d2e1f9a83eb7e70c863336725b13d8bb07f428db0db52c8b898ae): complete subsection reference.

- [vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-25ede9f0af75f70a6978df18cd382af126dd908afc4ffbf058f5808d071e3b56): complete subsection reference.

- [vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-04ae16beb927798a2b09499d45d1b859c7e1e1e6ee5ba2be614139e7bd8049f6): complete subsection reference.

<a id="canonical-86e7fecb1f62cf45464d5ece0b55e729bc4c09615bedadc2797398fbd78bc3e5"></a>

## Next pages — default_pool.origin_servers / f3b1d31a03a8 / 5

- [default_pool.origin_servers.cbip_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-5052793e3d8883a156f15b4de20f4e811b5cc9c36fe7357226dcd23dd696f186)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [default_pool.origin_servers.custom_endpoint_object](data-sources--http_loadbalancer--reference--group-015.md#canonical-055a91496ac2639599fbdb0334c75702aca7edbbbe0ec456aaac14812156fa93)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [default_pool.origin_servers.public_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-ecb9b44d81af36c8674aac3d22ff09119e51f5fc66f0d2b9426df39ca56d60b3)
- [default_pool.origin_servers.public_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-e7dee21ac44d2e1f9a83eb7e70c863336725b13d8bb07f428db0db52c8b898ae)
- [default_pool.origin_servers.vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-25ede9f0af75f70a6978df18cd382af126dd908afc4ffbf058f5808d071e3b56)
- [default_pool.origin_servers.vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-04ae16beb927798a2b09499d45d1b859c7e1e1e6ee5ba2be614139e7bd8049f6)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5052793e3d8883a156f15b4de20f4e811b5cc9c36fe7357226dcd23dd696f186"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecf86adbbaaa575255f88bea960f1e8e5dcd9507767da13d3ea3afab9d52d765"></a>

## default_pool.origin_servers.cbip_service — default_pool.origin_servers.cbip_service / c138e9a2ecb4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.cbip_service

<a id="canonical-4b39b1866bc4ed71d3171808278040e1e18dbf53b1d5968204ba1a0ae95290ce"></a>

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

<a id="canonical-542ae0ea598f016bb8750e53a65170c2704446798bd4bc8981bdae8af1cb706b"></a>

## Direct properties — default_pool.origin_servers.cbip_service / c138e9a2ecb4 / 3

<a id="canonical-b13b2ecc4b55029c7d0ff35d721910f92d2ffd5c9d3b11e8b03dc02a50ca1cd1"></a>

<a id="canonical-f7ef53a62566c4f4c3598b9e6d19d4c9356d836d441f30b0122ca123a89ccfe2"></a>

## service_name property — default_pool.origin_servers.cbip_service / c138e9a2ecb4 / 4

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

<a id="canonical-f8693001071206753264ecdc7219bf729742cfc45d855cf15219acd443f02d96"></a>

## Next pages — default_pool.origin_servers.cbip_service / c138e9a2ecb4 / 5

- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b352d8e19c7350fe9c3dea35f5e552f27987b7751eea96198427abba346183d8"></a>

## default_pool.origin_servers.consul_service — default_pool.origin_servers.consul_service / a0e762663e4a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.consul_service

<a id="canonical-454b310de00de2018e13ed524129193b9c8286a1517344d35afaec1f77d525b9"></a>

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

<a id="canonical-00d7962301ef279c8bf9811d1170a053ab204491ec09ccb0bc3b1771360def90"></a>

## Direct properties — default_pool.origin_servers.consul_service / a0e762663e4a / 3

- [inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-b31404d962e5d35c7638c6c110e841765308310ba3bd5ea67e1ae8d85b23d6d3): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-9b0dd09991402f462dcba8351d8adefce415d003283645e3437dc4d8bdb63c34): complete subsection reference.

<a id="canonical-6d721f08f887d80523ea6b10e22e18b4a8430500f086720bbbd48eb303ceaf03"></a>

<a id="canonical-2bbf2b06b3739128d610d8985050b6b35cecf4b76cef1248eec719a33b6eb288"></a>

## service_name property — default_pool.origin_servers.consul_service / a0e762663e4a / 4

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

- [site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd): complete subsection reference.

<a id="canonical-12527910fc4e60b3968f34588bca62b91fc10d063fef5a9e7f9c83734f2f14dc"></a>

## Next pages — default_pool.origin_servers.consul_service / a0e762663e4a / 5

- [default_pool.origin_servers.consul_service.inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-b31404d962e5d35c7638c6c110e841765308310ba3bd5ea67e1ae8d85b23d6d3)
- [default_pool.origin_servers.consul_service.outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-9b0dd09991402f462dcba8351d8adefce415d003283645e3437dc4d8bdb63c34)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b31404d962e5d35c7638c6c110e841765308310ba3bd5ea67e1ae8d85b23d6d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd765c97fa379d4732974fdee03aeb75bfcf2bbd86717894d559e50115f74984"></a>

## default_pool.origin_servers.consul_service.inside_network — default_pool.origin_servers.consul_service.inside_network / adf32b99a52b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- default_pool.origin_servers.consul_service.inside_network

<a id="canonical-b39f69d8622e0ec8a75bb8539e00ddba50021bac874f7156012393c249f3b221"></a>

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

<a id="canonical-de631b7793fb40d04794b6cf08643ac9f50fb28897782daeb8ebb750d2d40a12"></a>

## Direct properties — default_pool.origin_servers.consul_service.inside_network / adf32b99a52b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-850c4968856a8af94dbfc9ba9f19bc710701ed2bf584d8706ea17c4a712667a2"></a>

## Next pages — default_pool.origin_servers.consul_service.inside_network / adf32b99a52b / 4

- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9b0dd09991402f462dcba8351d8adefce415d003283645e3437dc4d8bdb63c34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c05ab30af1e1a165c0d1bb81802174ef23c85aaec8a18d4453e5c58d56c37d25"></a>

## default_pool.origin_servers.consul_service.outside_network — default_pool.origin_servers.consul_service.outside_network / c9f155a0429a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- default_pool.origin_servers.consul_service.outside_network

<a id="canonical-6dd0f3c029a2ce6c375429d8a5544534bf697dc1ebd3c3ac60be8e28ffdf6360"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-6b0b108ad2247054b4ef369ee43a3f6c552c82055efeda2cc82a33628dc95d22"></a>

## Direct properties — default_pool.origin_servers.consul_service.outside_network / c9f155a0429a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c71a324c57da0be9f80154fd72736e58e7fb83cb635f3fba6565259f1544c8d"></a>

## Next pages — default_pool.origin_servers.consul_service.outside_network / c9f155a0429a / 4

- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae87daa56b1f167cbe3032f0700ce57629b43d90ecb06b620b7112ad79363654"></a>

## default_pool.origin_servers.consul_service.site_locator — default_pool.origin_servers.consul_service.site_locator / cf5d6b311512 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- default_pool.origin_servers.consul_service.site_locator

<a id="canonical-f50948f3e9d5d1a380323be292e13443c9edd1922f3bd5fb9e02c6d6f82edbc1"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-340fac6a087b5e06fccc24f7c015fe11eaa88b9f583b1764b1d38494218ad1dd"></a>

## Direct properties — default_pool.origin_servers.consul_service.site_locator / cf5d6b311512 / 3

- [site](data-sources--http_loadbalancer--reference--group-015.md#canonical-177e710d2999c9d651192c04c84b4c5652264c30a8daac925bb782ada4dc31ae): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-2597e173ef6e1b75f2633e97cd17f8056f9cd570c8dd6dbd8b00ffa61e55dde5): complete subsection reference.

<a id="canonical-513b33d472c64e644f381fbc6e8819ac974e2ebd19567c15e9a3081089ccc235"></a>

## Next pages — default_pool.origin_servers.consul_service.site_locator / cf5d6b311512 / 4

- [default_pool.origin_servers.consul_service.site_locator.site](data-sources--http_loadbalancer--reference--group-015.md#canonical-177e710d2999c9d651192c04c84b4c5652264c30a8daac925bb782ada4dc31ae)
- [default_pool.origin_servers.consul_service.site_locator.virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-2597e173ef6e1b75f2633e97cd17f8056f9cd570c8dd6dbd8b00ffa61e55dde5)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-177e710d2999c9d651192c04c84b4c5652264c30a8daac925bb782ada4dc31ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd7fc9644fcc553a4b0fa79551efb4e0ae57d89571be5935f58c2ea30e055738"></a>

## default_pool.origin_servers.consul_service.site_locator.site — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1)
- default_pool.origin_servers.consul_service.site_locator.site

<a id="canonical-cc0446d990ed13359967a7c99afbc4c1b8eef6c67016ce84e0a8d81183651eb1"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-ef3710ea4b3fa500503085e46fedbe648b21adc84e62a0644413b01ded19b484"></a>

## Direct properties — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 3

<a id="canonical-3337278a9a81a8c40d4969b313eb93d20f9a2a49c6aa2566d6982ce83c644eb0"></a>

<a id="canonical-fe75bce4f56b92cd230d8368ff728e9d4b72383fb4851aeffc9db1b5e48543ec"></a>

## name property — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 4

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

<a id="canonical-eb5f8979d2dd6690c2a767ce66bf5c631e85fcdc62478dcfbff46d45a0ab1bba"></a>

<a id="canonical-8f5c2af6a63ab32b7c21dd83f5b02e1245bfae584055060dcbeaa7db12f8bef1"></a>

## namespace property — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 5

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

<a id="canonical-34ceb706062b41977699d72531a9f3c6a320049de42c6f8544e70bfe44bd329c"></a>

<a id="canonical-a41335c394521a6578ca96b7002cfd6cf99cc5bcc8c80711579f448c5b3b54a4"></a>

## tenant property — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 6

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

<a id="canonical-fc79a18a2c2227f1bb4854e6c3adf11bd7932c0ca6315ed72e75d931331b0577"></a>

## Next pages — default_pool.origin_servers.consul_service.site_locator.site / f12f50c0e4b2 / 7

- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2597e173ef6e1b75f2633e97cd17f8056f9cd570c8dd6dbd8b00ffa61e55dde5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-816b417cd75092a8eed659fab08e13b738a42b3ef74cbe4865d4850c900a0a2e"></a>

## default_pool.origin_servers.consul_service.site_locator.virtual_site — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1)
- default_pool.origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-4a0b352021183cb060f6e6f90fdda4b046d25ca31db8c57824e2919b8dc45f0f"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-5cd624d933867d4f8b916c8affd4fed8055cf5e19759a62a2dfc469316592e76"></a>

## Direct properties — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 3

<a id="canonical-87f79094f8bca2029b5692e7ff35e1a72b625c49ce49e1e6a3b64330de56e681"></a>

<a id="canonical-3015cdab2eb35c85326fe592cc8895e8b0cb8f8c7702478d7deeccc305ffd98b"></a>

## name property — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 4

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

<a id="canonical-5de9dfb32c390e17a1d32e04c1960d607687bfa07d3a05e107af97ee3919f87d"></a>

<a id="canonical-2fc52c54762e1f377603aa0246fc37e42721ed042a149a88ad52a0879a2359ed"></a>

## namespace property — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 5

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

<a id="canonical-cd58da90d9cc69b6628536ba014e225dbf31a41bb2e70221a6d9ad00e1c89740"></a>

<a id="canonical-24f95b3d4505fd3bff1883afc4de42cea9566da6a81e2f5a6508ef129e9bbfc8"></a>

## tenant property — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 6

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

<a id="canonical-5e10d34c61a443cede43a120d7c7f272ef24a8257de2bd2e88a095bbc4c55726"></a>

## Next pages — default_pool.origin_servers.consul_service.site_locator.virtual_site / a389b09bfd9a / 7

- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-a5c83aa39e58f5335c1dfcbca66689ce8b8582bdb0e4af55e36a9ffb094315b1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0a0beda8ca5ded9e4d3760a7ec55a2bd48171ee8a5ef943d97196dea9617852"></a>

## default_pool.origin_servers.consul_service.snat_pool — default_pool.origin_servers.consul_service.snat_pool / 58ae13905892 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- default_pool.origin_servers.consul_service.snat_pool

<a id="canonical-0b2580fde6c4435faad418038a8d5b0b6833a99fe27a7c9592208f87f17c76bc"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-178375847c1b0bfb568239143903056bc56514b26ccaedb007a10fdcb7e3bbcb"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool / 58ae13905892 / 3

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-3b951036bb89253e4ed7dd9c188d9bfa1cbe7d21e3a39947bec72e47fe03fdb0): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-49fae10a6ddfe4c11e515e4de0047496980576379499687787bb18ffc45522b7): complete subsection reference.

<a id="canonical-bc23f4b125952cfc4994ad06e73a7c2ed8ee1973677578c29e7bce1ae239691b"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool / 58ae13905892 / 4

- [default_pool.origin_servers.consul_service.snat_pool.no_snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-3b951036bb89253e4ed7dd9c188d9bfa1cbe7d21e3a39947bec72e47fe03fdb0)
- [default_pool.origin_servers.consul_service.snat_pool.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-49fae10a6ddfe4c11e515e4de0047496980576379499687787bb18ffc45522b7)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3b951036bb89253e4ed7dd9c188d9bfa1cbe7d21e3a39947bec72e47fe03fdb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12a5a9258493a28dbfe20599bbbb3e43962f4c9f8a3428e2859cb71365cbc270"></a>

## default_pool.origin_servers.consul_service.snat_pool.no_snat_pool — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / bdcac1a0195c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd)
- default_pool.origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-0ba1bcbf7c82b8f85076702fc093f49538c1cafa9633c3481dfb3fa5857d9509"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-6e713a8ef08745afb7e507ba99de2d0fe395a6020799608465e3a09809b9c6f0"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / bdcac1a0195c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d27b1758cff9a9bf597ea669bc16d0a38cdca0ce53f0e86d785e7fc8b4a1471c"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / bdcac1a0195c / 4

- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-49fae10a6ddfe4c11e515e4de0047496980576379499687787bb18ffc45522b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8772e0ac405646406bc71daf5bed3b81fac6c9c6e468dc38c5dc5774e4eb42d"></a>

## default_pool.origin_servers.consul_service.snat_pool.snat_pool — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 0cd2b9ae25ea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-564ec7426092bbe0204794866b25e52fd5419266a87702ce680895eb2bcbf046)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd)
- default_pool.origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-15953e17939869c7bea36389791a25f5028af934f63a232a0b1ea0b39b13ccdd"></a>

Type: `"single"`. Computed.

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

<a id="canonical-d7c60f07388d00625e1823b45e65d5b6c1fcb891c47b54bb2f068efd66fe4f36"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 0cd2b9ae25ea / 3

<a id="canonical-b9b9a4cc816c55866166532f36051fc8665947994a43be4716b9376087ce380a"></a>

<a id="canonical-94e748d7286272cd9526e3352e1dc0db1997b675709cf9cd8b64d34ec56e8234"></a>

## prefixes property — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 0cd2b9ae25ea / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-b880ddffa4594142e39eea91b5fee05bff84e08d666baf71acffe7339a7833ce"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 0cd2b9ae25ea / 5

- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-485224535332e81d4e3bbf835d9b415a98c093c2d3aa8494a5b647a7bd5a92dd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-055a91496ac2639599fbdb0334c75702aca7edbbbe0ec456aaac14812156fa93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19bbb90383fedbca9db40520727e7fa0571f342b6d51553a437f8668bdb1b500"></a>

## default_pool.origin_servers.custom_endpoint_object — default_pool.origin_servers.custom_endpoint_object / 5dd91160e962 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.custom_endpoint_object

<a id="canonical-d3bfa81ddabc6d7527f88a4d112d1806c91f5cd267b71000bf0da72227e2d9d0"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

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

<a id="canonical-342eb46de3e458db54aa837a7a40105af1928df7998a5ffce14af4f87c55e6c3"></a>

## Direct properties — default_pool.origin_servers.custom_endpoint_object / 5dd91160e962 / 3

- [endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-122cdebc7bc7d04a87ef0b3cb2dc5b0d3d66e250bba25d724ceadc41e53e0cd1): complete subsection reference.

<a id="canonical-0cceb0e86737d00440af74b525005dcec2725a6d257797adeb745bcc3f66b81f"></a>

## Next pages — default_pool.origin_servers.custom_endpoint_object / 5dd91160e962 / 4

- [default_pool.origin_servers.custom_endpoint_object.endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-122cdebc7bc7d04a87ef0b3cb2dc5b0d3d66e250bba25d724ceadc41e53e0cd1)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-122cdebc7bc7d04a87ef0b3cb2dc5b0d3d66e250bba25d724ceadc41e53e0cd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42ddffa2c56f54a94346c75a215da897d245f580ba6e469a56a39c878fb424f9"></a>

## default_pool.origin_servers.custom_endpoint_object.endpoint — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.custom_endpoint_object](data-sources--http_loadbalancer--reference--group-015.md#canonical-055a91496ac2639599fbdb0334c75702aca7edbbbe0ec456aaac14812156fa93)
- default_pool.origin_servers.custom_endpoint_object.endpoint

<a id="canonical-3d45834c34a12a3f90aa0a57ed6c065d3cdcc080d6ece1995067b13f5d0215ce"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-317c6ac98c51061d90ff7e8f3e072b1bdd562affa01d0f18e1faa74cb8edb582"></a>

## Direct properties — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 3

<a id="canonical-25f3d9fbe6e6f82392470e747ce587ce193f8c4bcf965ee1b33a81642875e4eb"></a>

<a id="canonical-4ac170173897efb057ce7f3288f13b268193f36cf3d4497bf7efb09c5dcb1796"></a>

## name property — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 4

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

<a id="canonical-50ce2e46555dc12935bce60d4175ee187ba93ee4b8d8fd862ede3bcff059e688"></a>

<a id="canonical-8331e4539c1c1ca18506bc6b9acdb6130fc4b89cf27a79432348bba3dde71604"></a>

## namespace property — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 5

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

<a id="canonical-c83b9b77ebfb961824721d75f07f27812b420936dfcdbed8c006f55d79ced9df"></a>

<a id="canonical-8bf3321bb571bb672e6e8c10e619bc98621b96cbbb4549ae8de99b1d64f2cd0c"></a>

## tenant property — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 6

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

<a id="canonical-5760e5167822c99223ee7f342c39de0b4ca8147d7d80d5bc429c15414676b573"></a>

## Next pages — default_pool.origin_servers.custom_endpoint_object.endpoint / 240c9e009bf1 / 7

- [default_pool.origin_servers.custom_endpoint_object](data-sources--http_loadbalancer--reference--group-015.md#canonical-055a91496ac2639599fbdb0334c75702aca7edbbbe0ec456aaac14812156fa93)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b14f4dcfd2c11fa0f542fd95ccc1387c06892e73d4aa450bbfca1ed1d5bbf3be"></a>

## default_pool.origin_servers.k8s_service — default_pool.origin_servers.k8s_service / 9a0f2278db1e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.k8s_service

<a id="canonical-4c0e31661bd6c5fc08dc349a728bc882bfd2dd5fcecccb35736f67e1c1df5e08"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-61982c08125c7355ac1886c991409652a41a4070bd323a195adf176035c25932"></a>

## Direct properties — default_pool.origin_servers.k8s_service / 9a0f2278db1e / 3

- [inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-02daef22e98bc43b42008b3dc8c64a78f9644ad503b139761c12ad3119a02be3): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-7976d6d3c2aaa9d358f04a2acf4696fa1d911a4ffa9b6554eb1ed82c14164491): complete subsection reference.

<a id="canonical-34f658100844ccb2bcea31f6c9e809128795dac3d0365731e410722d9fe0b533"></a>

<a id="canonical-dbe69d3250cbf0119ce0b9c5a474fa0402dd4499f4471a9f8f024c0fbd2c27ab"></a>

## protocol property — default_pool.origin_servers.k8s_service / 9a0f2278db1e / 4

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-63881b9d54a70436e0f3f7eee6285a3a1b5e0e66ebb6c1d0855faccc12c3be9d"></a>

<a id="canonical-d7e6e869bc73c790131b1ab80d83f2699eb9fc610dd627eda95dbd78972a50d3"></a>

## service_name property — default_pool.origin_servers.k8s_service / 9a0f2278db1e / 5

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f): complete subsection reference.

- [vk8s_networks](data-sources--http_loadbalancer--reference--group-015.md#canonical-48bda19fd0f98bccf1d4a4d169a1d17ca9062bb1eda73bef4265cb632daa8329): complete subsection reference.

<a id="canonical-d24aa703b5fb1fe0a155b7a9fa8ce4ef7271f98f976c730ca3052697cb7351ff"></a>

## Next pages — default_pool.origin_servers.k8s_service / 9a0f2278db1e / 6

- [default_pool.origin_servers.k8s_service.inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-02daef22e98bc43b42008b3dc8c64a78f9644ad503b139761c12ad3119a02be3)
- [default_pool.origin_servers.k8s_service.outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-7976d6d3c2aaa9d358f04a2acf4696fa1d911a4ffa9b6554eb1ed82c14164491)
- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3)
- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f)
- [default_pool.origin_servers.k8s_service.vk8s_networks](data-sources--http_loadbalancer--reference--group-015.md#canonical-48bda19fd0f98bccf1d4a4d169a1d17ca9062bb1eda73bef4265cb632daa8329)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-02daef22e98bc43b42008b3dc8c64a78f9644ad503b139761c12ad3119a02be3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26e29942ae1305a3bea402e4bebc1a430bd0ba39cde7a5c10b999437ced912b3"></a>

## default_pool.origin_servers.k8s_service.inside_network — default_pool.origin_servers.k8s_service.inside_network / 0f99cd336548 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- default_pool.origin_servers.k8s_service.inside_network

<a id="canonical-b76acdd50d1f9324c24f9880828f593d15583d38c398326103899633fbf57b38"></a>

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

<a id="canonical-2938e9d962f30648103f03a82fc8a26b0e7bbedb5b55913e59d3e6e07fbbba94"></a>

## Direct properties — default_pool.origin_servers.k8s_service.inside_network / 0f99cd336548 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74686bd7e3ea2a931b2e4b028bdce2c8193a59f2e9d94694139fa5ae9fb8653d"></a>

## Next pages — default_pool.origin_servers.k8s_service.inside_network / 0f99cd336548 / 4

- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7976d6d3c2aaa9d358f04a2acf4696fa1d911a4ffa9b6554eb1ed82c14164491"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b514f17adcf8786b3304f30e9d6b3a525b02700dee8a68c2819bbcf768b4755"></a>

## default_pool.origin_servers.k8s_service.outside_network — default_pool.origin_servers.k8s_service.outside_network / 7daf37faee87 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- default_pool.origin_servers.k8s_service.outside_network

<a id="canonical-8a1f156e6fdd70ca9da269d5e685ed3f777be74e2181d0fd048ee281524c309c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-eb34279f30a2ecaaa9e81f6bbb8b7f00f0952027e13252f2bf5ff2aa3bb50fc0"></a>

## Direct properties — default_pool.origin_servers.k8s_service.outside_network / 7daf37faee87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-751d6867f62f34b1109e3988e8f0cea9927e8b1a966762a0934c6c787715ac11"></a>

## Next pages — default_pool.origin_servers.k8s_service.outside_network / 7daf37faee87 / 4

- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f82b26508d1d45038c45e239ee44b039f27ea3d1144316a9cad765f7ac1ab3ec"></a>

## default_pool.origin_servers.k8s_service.site_locator — default_pool.origin_servers.k8s_service.site_locator / 9f81cb0e3efd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- default_pool.origin_servers.k8s_service.site_locator

<a id="canonical-f37719448ac660a52672e46889a1d1d89c23be054e1ad1b560f625fa246aa92d"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-d73b98993cd7533940191c10d77ba64e1cb8d6e3ed769a82932f7e15e7644daf"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator / 9f81cb0e3efd / 3

- [site](data-sources--http_loadbalancer--reference--group-015.md#canonical-f10ff0426f645e6f6a38f313ec44a42b04e4198e306290f1a146cf7eee8ac3b3): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-1bca05dd31f21b3a871297192e87c6c2d389b634b4456be1e4a249b50d94943e): complete subsection reference.

<a id="canonical-40c66b359211cdf0aa9173cfa9c9fee2cd100eca26ce1af21afa9ce1a3062ead"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator / 9f81cb0e3efd / 4

- [default_pool.origin_servers.k8s_service.site_locator.site](data-sources--http_loadbalancer--reference--group-015.md#canonical-f10ff0426f645e6f6a38f313ec44a42b04e4198e306290f1a146cf7eee8ac3b3)
- [default_pool.origin_servers.k8s_service.site_locator.virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-1bca05dd31f21b3a871297192e87c6c2d389b634b4456be1e4a249b50d94943e)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f10ff0426f645e6f6a38f313ec44a42b04e4198e306290f1a146cf7eee8ac3b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1e4a04d5b9f3fccce2def31628840ddd30cb9734fbb0f50230cab910a14cc12"></a>

## default_pool.origin_servers.k8s_service.site_locator.site — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3)
- default_pool.origin_servers.k8s_service.site_locator.site

<a id="canonical-2607b9a691c760fd34c4a2c2590e4c107999fd9e731573c35fe0642fe6632a7b"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-ac4cfd778c503d80ebfb697cf77e9f24646111fc21f6d535a9d2f133c5a47dd2"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 3

<a id="canonical-e83f099733b52bc40475163c0c2dbdb433f59d8e80920d55da2be19b4c95dc04"></a>

<a id="canonical-0a7a513f6ad71b775eab5b05c847467113f9a7842acc4d6cf2b081636ec44281"></a>

## name property — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 4

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

<a id="canonical-fa091abc4e4f46725d1f015ad7c0465296be27bac68f6dbe91ffcc62ef84580d"></a>

<a id="canonical-4dc361a69b705d899f486c9be3ed440ce01af815e74984dd88be06312ae1dceb"></a>

## namespace property — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 5

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

<a id="canonical-99ffc8bf765e32e3c99150db98dc5f6edf5ebad436a050f06feebae5c16695fc"></a>

<a id="canonical-dd66d5c59997135cffe20ab1faee676a26413af44f054bdee611008386775460"></a>

## tenant property — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 6

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

<a id="canonical-7b3a45b3f138ac018823a6cba51898f62e781f3b8cdcce149b970907083a31d2"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator.site / 180c29f4d95b / 7

- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1bca05dd31f21b3a871297192e87c6c2d389b634b4456be1e4a249b50d94943e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd19872b9577395651e172daaf93b9ee3852cbd0e6da7433a7657a3b1d87a843"></a>

## default_pool.origin_servers.k8s_service.site_locator.virtual_site — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3)
- default_pool.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-d606a5fe07bb398d472e25fba5cf575994cf5015c47e767b92ef130dbb28c6ec"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-9917df349cecbd95890aaffc4a1001dfe628901c1f7708828067624555b75f2e"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 3

<a id="canonical-b092792a93ef1fc5234e79cb6d2cb9644de3bf87be54fe078e867d65cf96dd69"></a>

<a id="canonical-9fcbb7f681966c5bc84d4d145b6ae9c448863deb3c545d8bfa49d9f5428ac759"></a>

## name property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 4

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

<a id="canonical-d53ac457012b061cfe18a5bee2eaef3986f3d35d3c4de489a925088030416260"></a>

<a id="canonical-6ee86cee629b537df1fc98bea7e2dfead617f7de014295b0dfbab7b68e602ea0"></a>

## namespace property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 5

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

<a id="canonical-fe4591a24cd0390b8937765de7899684248a6a15c3e95df20b9d0c2538da22b9"></a>

<a id="canonical-17494e4b6fa97838d694af8a6c609c1ebbb8001f125399b0fe75e42ef1d0f77d"></a>

## tenant property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 6

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

<a id="canonical-cabbfde0527e4b86427e12d4e8459e32f7c94731faad78a259b4fe77aa545a61"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator.virtual_site / 88b50ff12145 / 7

- [default_pool.origin_servers.k8s_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-37cacd9e719bf4e0fe9b1d14ee8d5bfcdfdd53f1966c90f95f1a1efbb6fcfbb3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4d2fc7f5438b627f65c145f69578a317eebbd42af5bfbfb85a968be47301a27"></a>

## default_pool.origin_servers.k8s_service.snat_pool — default_pool.origin_servers.k8s_service.snat_pool / 9ec6f3d025bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="canonical-7bb878518f906d8dcb9145863265fa8b449bdff1b245d1d662e0ae9a333e78ee"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-1578bc8ab11d358ec6811ddb3b7b4dea8db3c8a0308008f96f60d02eed6e4a16"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool / 9ec6f3d025bf / 3

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-4cf19005e21b0fe1d026f02be6f993055cd0e183065e235d12eb7cb13a3fad1a): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-80e5666f940417c0366bb7e2639641e50c237998599bec94645e8e1a0034d92b): complete subsection reference.

<a id="canonical-5b445ca4e3a8ce5a97d060175a460d2a9c9c24409ddc370758cf3a21ffa41586"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool / 9ec6f3d025bf / 4

- [default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-4cf19005e21b0fe1d026f02be6f993055cd0e183065e235d12eb7cb13a3fad1a)
- [default_pool.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-80e5666f940417c0366bb7e2639641e50c237998599bec94645e8e1a0034d92b)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4cf19005e21b0fe1d026f02be6f993055cd0e183065e235d12eb7cb13a3fad1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ece37bf1071cc14aca47ef28cecf744e6b23fd1b592a424720a295e1a5e5d88"></a>

## default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / ab295f442d56 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f)
- default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-cddd6c77e0254a89a2b4dc1c703e3908fa22ab6ceb870411f653e79d9f0b57a6"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7081462aa4af466d7780592b3c1002be76c4230d68a74d234bfd2c8282ae9f27"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / ab295f442d56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c476e46565a8a6214ea85068b963d9d1692c0fc364ef53a058d3d8b41e3aa176"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / ab295f442d56 / 4

- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-80e5666f940417c0366bb7e2639641e50c237998599bec94645e8e1a0034d92b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc6ede23e497b56a8d7b4944d7e73062139ce9406d2f53c31b6e705cadf50026"></a>

## default_pool.origin_servers.k8s_service.snat_pool.snat_pool — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / 717313503913 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f)
- default_pool.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-168a44741397a13e27dda25e44607fdfeb4b947844889741f0ebd382c6dcf900"></a>

Type: `"single"`. Computed.

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

<a id="canonical-e2d510206dddfb718b7dd6867ee5b5b0ed1727838c1defb5e0dc1c302f9549a8"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / 717313503913 / 3

<a id="canonical-48c934b099353a5e096072082107fca983644e842b0561a090f3d480e279fae8"></a>

<a id="canonical-711907f041d33d79a7a49442de029d3d619ef136a27a8ae94453b97a5378f5b8"></a>

## prefixes property — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / 717313503913 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-7c2a0e3b03aa67194b09064af14bf3c40dc441cf4dde586922b7323429f39117"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / 717313503913 / 5

- [default_pool.origin_servers.k8s_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-67d7328901f595078668abf886f3b519af6d62446ab4f46b7ee9dbc3014a199f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-48bda19fd0f98bccf1d4a4d169a1d17ca9062bb1eda73bef4265cb632daa8329"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22cb653cd7aa84c77bfc56db7315274f97a04675ccfa4130c0ee549ace2a646d"></a>

## default_pool.origin_servers.k8s_service.vk8s_networks — default_pool.origin_servers.k8s_service.vk8s_networks / 27a5e95051da / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- default_pool.origin_servers.k8s_service.vk8s_networks

<a id="canonical-ccf4a7aaeff29c6bc90552e4a1026e162487f6dc62131dfe9c2cbe7267ca40fa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

<a id="canonical-5a472ac3bb16b6496e35d7deeccd1f9d02c39af043b3cb4b30d5f1562e2a5e78"></a>

## Direct properties — default_pool.origin_servers.k8s_service.vk8s_networks / 27a5e95051da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ab1ec88127348087deb84599b409cce60e24b62c4d25d956dab9bec0450f3bd"></a>

## Next pages — default_pool.origin_servers.k8s_service.vk8s_networks / 27a5e95051da / 4

- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-a06412987620ac998825c37ef9203341370119c97736966f239fd0e0ce7e24b9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e97c3ab7c776757599322df2ae725b871987845059d9328d46e42c26bd69a625"></a>

## default_pool.origin_servers.private_ip — default_pool.origin_servers.private_ip / a66db9752f2f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.private_ip

<a id="canonical-c3e857480236298c126e9b2cfffbefc90a67daebd1a4999fd349e8659c6c2af4"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-f1099353a6d5adc3db1f1b230c03a2cc50475c2aed8198558dc2955d81f18635"></a>

## Direct properties — default_pool.origin_servers.private_ip / a66db9752f2f / 3

- [inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-5f87b53b5070e49e2544cda8f6777c208d2831179946e2bc957ab87c20ca49f8): complete subsection reference.

<a id="canonical-0a2ef0d85c1d7a4d4faaeabf7e4a116a6b5b0a780843fc983a82e7a8a098d040"></a>

<a id="canonical-8065c8a1a8c06d46d48f1cc2c756f58c814e789c7c163af612b2e3c96468ee33"></a>

## ip property — default_pool.origin_servers.private_ip / a66db9752f2f / 4

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-fd3c1ced9d30ae2d85b7edd2499db475084fc14b8d7e25780788a7de467ff675): complete subsection reference.

- [segment](data-sources--http_loadbalancer--reference--group-015.md#canonical-f2a883ae7a7f127de721c067c2f2f8b4883e2a4ead49c1332dc563f79ac5d459): complete subsection reference.

- [site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42): complete subsection reference.

<a id="canonical-e7d752a8f9648f5f0f21b7f28473bbc3597f2e01a44882d92e3a4f5a42a02ff5"></a>

## Next pages — default_pool.origin_servers.private_ip / a66db9752f2f / 5

- [default_pool.origin_servers.private_ip.inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-5f87b53b5070e49e2544cda8f6777c208d2831179946e2bc957ab87c20ca49f8)
- [default_pool.origin_servers.private_ip.outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-fd3c1ced9d30ae2d85b7edd2499db475084fc14b8d7e25780788a7de467ff675)
- [default_pool.origin_servers.private_ip.segment](data-sources--http_loadbalancer--reference--group-015.md#canonical-f2a883ae7a7f127de721c067c2f2f8b4883e2a4ead49c1332dc563f79ac5d459)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5f87b53b5070e49e2544cda8f6777c208d2831179946e2bc957ab87c20ca49f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8422f1dc1e58a744bf6a6fff2fa1a85154f76b5d712a19dc1a71329c259299a2"></a>

## default_pool.origin_servers.private_ip.inside_network — default_pool.origin_servers.private_ip.inside_network / f89f1fb5647e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- default_pool.origin_servers.private_ip.inside_network

<a id="canonical-e312b201ec98408cdb2465806760a16d0234bd0ac2c33fdb8452720ca9ae2681"></a>

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

<a id="canonical-e22043d163a6f0ed1d2a247aaea353d90d0ea13a13ecbd93cd8d6b87e27b3887"></a>

## Direct properties — default_pool.origin_servers.private_ip.inside_network / f89f1fb5647e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7ec3984d6a9805a5e8ae8bccafc6e5c029a3e452b17ba4b81c3b909f8b2b5a7"></a>

## Next pages — default_pool.origin_servers.private_ip.inside_network / f89f1fb5647e / 4

- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fd3c1ced9d30ae2d85b7edd2499db475084fc14b8d7e25780788a7de467ff675"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-334ab07e4bf72a4a5c9c6a4164113ceb5d6ff34c1baca39f7e58f6eb3ab52f8f"></a>

## default_pool.origin_servers.private_ip.outside_network — default_pool.origin_servers.private_ip.outside_network / d101530208fe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- default_pool.origin_servers.private_ip.outside_network

<a id="canonical-ef3ad2e8ead6ae0bd69815fe43aece13826367e8c576aa1a17ac86daddcdde61"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-15e919edb561d1acb48bdc54b168830d7dd9af4d490590d4b5a8895f4fc09dbc"></a>

## Direct properties — default_pool.origin_servers.private_ip.outside_network / d101530208fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f191da9866a7685ccef0e59b33154e9cdfb66d5fc45064a1c9ba79d31c910d5"></a>

## Next pages — default_pool.origin_servers.private_ip.outside_network / d101530208fe / 4

- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f2a883ae7a7f127de721c067c2f2f8b4883e2a4ead49c1332dc563f79ac5d459"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d897ba199422621a80b22b94859f63643a6f2869012987a80b0c505208687200"></a>

## default_pool.origin_servers.private_ip.segment — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- default_pool.origin_servers.private_ip.segment

<a id="canonical-b06f387015a9d2c9e978f489d8c1cbbb139dd12a8460b01d9441a2975c95cee8"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3be9724df0a771bdb3e876aef580dcf13496cecff6e515858587f2b2058ba24a"></a>

## Direct properties — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 3

<a id="canonical-2aeca992bc708290859af7e5bdb6c11c719d57a0f83eab3f3d5da35ba9b498a7"></a>

<a id="canonical-32a1436d466f0741b9a5a20a87cb2cb4b9b7c2f5f13933cc90bcd855e0b7f01f"></a>

## name property — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 4

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

<a id="canonical-20e7c948e947aad3df9b2460b053dbe526d113d65f09ae2c63af0ad8913c6e09"></a>

<a id="canonical-07ace7cb614a753fa03f5310411b94fe2a5762d31719a0200926857c6ea35cf6"></a>

## namespace property — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 5

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

<a id="canonical-f9d550ad9be34ac999e4f9532ca5994034c5e6cc24d7d67760bafd8da7b61268"></a>

<a id="canonical-01c777a14474cecdac19c17fe47ccdbd4ade80c238c22c407c306919c146996e"></a>

## tenant property — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 6

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

<a id="canonical-36178f25cf9291457c959ebedca383b50d7770bb62b8c8060489573df8edf4a5"></a>

## Next pages — default_pool.origin_servers.private_ip.segment / 14518ea0efde / 7

- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10f808013a840dfd82b5e971ded3d90d3a368da03c8786ecb70c2fc0e739d31e"></a>

## default_pool.origin_servers.private_ip.site_locator — default_pool.origin_servers.private_ip.site_locator / 02a7ed928798 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- default_pool.origin_servers.private_ip.site_locator

<a id="canonical-717eacb7874d67558d0f84745cc38dd149f41a4f07b900248cc73ead967d8008"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```
