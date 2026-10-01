---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-313f20672fa41d7f6e7d10e63161622d52ebd6c7a718c4d91b047c3f0ee98609"></a>

## custom_errors property — http_proxy.more_option / ee2c20307560 / 4

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-3af4f345df0380e1dad3b9e42c99bd2f4af2421a5e839046b1a40a6c2ab831d3"></a>

<a id="canonical-50a2f148784eccd838fc676e1dc7631d11de42520ce37469ce5e2438969220b7"></a>

## disable_default_error_pages property — http_proxy.more_option / ee2c20307560 / 5

Type: `"bool"`. Computed.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-43905b3dbb9fca069ecd975381d003d9b864e3b221434dcd22cf9cb4d41d8b58): complete subsection reference.

- [enable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-6a7d154cadcacc3b4253efe066f1a7fb90aaf0db457c1107d94fde113dbf84ac): complete subsection reference.

<a id="canonical-ec43193bbd4ee38eb6271912b3b7a00cc7b571183f62322f3409b5899e24cf03"></a>

<a id="canonical-a3b24b3ca62fa1185cfad99c317d17395daae09f3c97827e26bfed9f431bbcf7"></a>

## idle_timeout property — http_proxy.more_option / ee2c20307560 / 6

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-66d4b7c92424dd6df0d03fb2a42cebc43f5da119ebacac7dd7af3fef9a763e6f"></a>

<a id="canonical-c3c4d1885e67eb06c885e5aba61c3c18fc749ce4a0349cab92e826bb07c3e6f7"></a>

## max_request_header_size property — http_proxy.more_option / ee2c20307560 / 7

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-0dbd9fe9f8de2a836e706451a48427f57311234ee1be16c5c40f240e6d312f6d"></a>

<a id="canonical-a0b1747e35cd944cedf74d36d89ba23630595f82fd211fecdf35b9585389b75b"></a>

## max_requests_per_connection property — http_proxy.more_option / ee2c20307560 / 8

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](data-sources--proxy--reference--group-004.md#canonical-62ed4b67570fb363bbe41fccd4c7d5bb10a7d6d475b70977302597c39a6e0680): complete subsection reference.

- [request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7): complete subsection reference.

<a id="canonical-7f5a175007c7c16464d2c5fe5e02e1beb84d22ca9f85089e479e3c577183a7b8"></a>

<a id="canonical-ec6356b527edea8646f8a45dacc1a55f50482a14f08ee5047ecc66e32030d4cc"></a>

## request_cookies_to_remove property — http_proxy.more_option / ee2c20307560 / 9

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a): complete subsection reference.

<a id="canonical-c2f1f2d369e4613e929115e24307a991330dfc8367c7510aa734e136df461220"></a>

<a id="canonical-9900ab69f363607011d5553d9cb96adac6f0168208110044a4cd22cd7160d415"></a>

## request_headers_to_remove property — http_proxy.more_option / ee2c20307560 / 10

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba): complete subsection reference.

<a id="canonical-97687c194baf78dbd355870bf736b97c9fac3001fdc402037ee63342944c7b1e"></a>

<a id="canonical-8f61980b2ebf184393355ba2345bbbbe7477bcdd4e12410162084e6db622e62e"></a>

## response_cookies_to_remove property — http_proxy.more_option / ee2c20307560 / 11

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397): complete subsection reference.

<a id="canonical-226e79e4b1ff6553c7343ed9f1f691234a61105003385001794abac631b4c97a"></a>

<a id="canonical-01c79313a8a90efc0c3a30d10645fe2e009392c9a4ac09b3185bf7c4039e5d2d"></a>

## response_headers_to_remove property — http_proxy.more_option / ee2c20307560 / 12

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6135ed3bddde094354843e5a7ede58c80f57c0a194ad494c3751efc62ba9e685"></a>

## Next pages — http_proxy.more_option / ee2c20307560 / 13

- [http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-004.md#canonical-01ab23d3ce197bde44da917f78effc43ea5044104ff2f20f10815e3199dc5daa)
- [http_proxy.more_option.compression_params](data-sources--proxy--reference--group-004.md#canonical-a6178e8b5129d7011cc321ea3cba253ea073e572a2a7b19ea4c669678829c51c)
- [http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-43905b3dbb9fca069ecd975381d003d9b864e3b221434dcd22cf9cb4d41d8b58)
- [http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-6a7d154cadcacc3b4253efe066f1a7fb90aaf0db457c1107d94fde113dbf84ac)
- [http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-004.md#canonical-62ed4b67570fb363bbe41fccd4c7d5bb10a7d6d475b70977302597c39a6e0680)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-01ab23d3ce197bde44da917f78effc43ea5044104ff2f20f10815e3199dc5daa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ea93db49cfd4011a9da51c6f4aacc34b10d99f0b9d0fb003a2e93502a33e8ba"></a>

## http_proxy.more_option.buffer_policy — http_proxy.more_option.buffer_policy / ab5756d9e884 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.buffer_policy

<a id="canonical-797be26098d1bd1213e80419c0319702abe14660d6740170b2a985d317acbe07"></a>

Type: `"single"`. Computed.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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

<a id="canonical-33674eb5559a6504f6d2068a85a7951530f49f17f02db7e7e07bf617886a46e4"></a>

## Direct properties — http_proxy.more_option.buffer_policy / ab5756d9e884 / 3

<a id="canonical-37a644aae40d4b47f0f74387df28cb17dcced8a32c3ad7d5c29801db267d7c68"></a>

<a id="canonical-7ea4cdb37356724a3e88537b3af0e4f1844e6bef43dee2658d8efce56b6f1e0e"></a>

## disabled property — http_proxy.more_option.buffer_policy / ab5756d9e884 / 4

Type: `"bool"`. Computed.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-4a7be4dc905ed483af2e6b607c87989eb1c9076511d88b286c1b443b77000dda"></a>

<a id="canonical-afde8c2667894723b27e694f7f689c58679e353ac04d1a7894b469840e818978"></a>

## max_request_bytes property — http_proxy.more_option.buffer_policy / ab5756d9e884 / 5

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-2967147ca521a875450e5c5a98d9843bc0c9eb5b5de93de7264549bce9ea2641"></a>

## Next pages — http_proxy.more_option.buffer_policy / ab5756d9e884 / 6

- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-a6178e8b5129d7011cc321ea3cba253ea073e572a2a7b19ea4c669678829c51c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6794bcf276ff52354e2d3a49681aad1d471dddd39148ad1016f71882b83ac07"></a>

## http_proxy.more_option.compression_params — http_proxy.more_option.compression_params / d7acf98d707d / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.compression_params

<a id="canonical-373d60d59d7778b43434e690b39787abfd4ee67b38f6a9908b41a0a51dc7f673"></a>

Type: `"single"`. Computed.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/javascript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

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

<a id="canonical-8f0b649bb0ccea02cafecedb55c1b6e57b01984bd40975e194101e273b6ee32e"></a>

## Direct properties — http_proxy.more_option.compression_params / d7acf98d707d / 3

<a id="canonical-7a702f962cb9b53b2e19df7f105db02feba197cbe9f80cfc2a62e8280494c062"></a>

<a id="canonical-506441a6247bb51ea6e97cea712515400f39e8455c8a7136dde27813ad5efee2"></a>

## content_length property — http_proxy.more_option.compression_params / d7acf98d707d / 4

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

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
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-4d234b4bba106a581acf470a4c13dc585f3b44404bd6266015524b5a44718535"></a>

<a id="canonical-0748de9842677f39af4d9ee18d08256c5d5cff16bf8a464f73a15d5986403c45"></a>

## content_type property — http_proxy.more_option.compression_params / d7acf98d707d / 5

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/javascript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/javascript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ce0ce0f749dcf02fca28d7595c6611049806462053fa5841da955f38c912abdd"></a>

<a id="canonical-3c8bfe2f0ecc4112d14b9b7e18debdd09a6beb5e34ab512ecba2185095ef16d5"></a>

## disable_on_etag_header property — http_proxy.more_option.compression_params / d7acf98d707d / 6

Type: `"bool"`. Computed.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-ae89c1a8ffb7bf4189a7745afbf7654933347ff445e021975a8b1fe1a25f7e67"></a>

<a id="canonical-683d3050795712fa63ba9209f3ef69a133548a98ac35d0ca584ef84d39566459"></a>

## remove_accept_encoding_header property — http_proxy.more_option.compression_params / d7acf98d707d / 7

Type: `"bool"`. Computed.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-d601c5c3a87b9cfef27fc84a668bf234f9e809b0d05b4a23713857f806884210"></a>

## Next pages — http_proxy.more_option.compression_params / d7acf98d707d / 8

- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-43905b3dbb9fca069ecd975381d003d9b864e3b221434dcd22cf9cb4d41d8b58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82a088c3684a0997a5be8083fdb293934bff9a726c7a6d4d91ba4836e2da0c5d"></a>

## http_proxy.more_option.disable_path_normalize — http_proxy.more_option.disable_path_normalize / fe1f827c4b35 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.disable_path_normalize

<a id="canonical-662018c04261d03c5c0a2bf828570ad1c8aae2ea18d569f4be5dadebd64ae6d7"></a>

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

<a id="canonical-5f5e98f3965c97a1d3d3c77c2dc6b374d430eaa30ea834330ec2f618ab5d5700"></a>

## Direct properties — http_proxy.more_option.disable_path_normalize / fe1f827c4b35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2ab004e1ac2d599caf3c1c6b80513bc5075dacc4b8a27f3da6281028ea02e09"></a>

## Next pages — http_proxy.more_option.disable_path_normalize / fe1f827c4b35 / 4

- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-6a7d154cadcacc3b4253efe066f1a7fb90aaf0db457c1107d94fde113dbf84ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-462c9fba9c6eb4b33db8a3bdf2792e0319aec297c793861b2dfc3a71a04c196d"></a>

## http_proxy.more_option.enable_path_normalize — http_proxy.more_option.enable_path_normalize / 560a71f0e7b8 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.enable_path_normalize

<a id="canonical-c874d73b43fff1c70dd958ec34f9705fa22d506f993cb3d4f0d79e1856be04ae"></a>

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

<a id="canonical-aa7d9f402bdd6c5facf6543c548d8659c25b2f698684f1568dc6e6a355b292fb"></a>

## Direct properties — http_proxy.more_option.enable_path_normalize / 560a71f0e7b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ea3122d1155a4fd79a67ad533619ce29532b5a556b614bb15f9c386868dad77"></a>

## Next pages — http_proxy.more_option.enable_path_normalize / 560a71f0e7b8 / 4

- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-62ed4b67570fb363bbe41fccd4c7d5bb10a7d6d475b70977302597c39a6e0680"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-853e84abadb1e05b0e9722675096ff3236c90be60f674f72a0abdef8ee0acbf0"></a>

## http_proxy.more_option.no_request_limit_per_connection — http_proxy.more_option.no_request_limit_per_connection / 50e7ac94b786 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-d4fa92c3041899578bc60733ea6251893d9961f7d7baf110a4b2bb5c24f6a79f"></a>

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

<a id="canonical-f293b434ffcd61dc6c27408c7a984ca09f1a9b7d576d40236f362b1b1f883895"></a>

## Direct properties — http_proxy.more_option.no_request_limit_per_connection / 50e7ac94b786 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0988d69bb26f03e47729312c5285478520e255fe2d4622a4758821418bf3eeb9"></a>

## Next pages — http_proxy.more_option.no_request_limit_per_connection / 50e7ac94b786 / 4

- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80c21f263f1da85d8bea814377c73626a651db79f28d4299f64e06a1683f45c0"></a>

## http_proxy.more_option.request_cookies_to_add — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.request_cookies_to_add

<a id="canonical-5714fe10d064e6d9ef3591d24d71eba7e26ffc2ab1871716bb9b63e3efa2d851"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7d7ea6e7becf3c51efcca4492fe91dc933cada4376b1a389a8585c5344dbf954"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 3

<a id="canonical-cf03af0067ba07f8656b10c3865f0b560497181c88d3e52f00a8987d494aa0a5"></a>

<a id="canonical-ddb0f0343dbbbfeb377a2aef3c145dc828998d7b224b50884c4804076f0edded"></a>

## name property — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "format": "dns-label",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-44ba8c83e88a037bbf0a76c4554d0ed20b94e2cfcec087b5793ad33b354d59dc"></a>

<a id="canonical-4b228446207a1b4656154ba15a8e2a8b3bb84f10705b54f3c995882a803d2263"></a>

## overwrite property — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 5

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b): complete subsection reference.

<a id="canonical-fc1ef5716282e7ffe0ffaaed931b64870f2bc83773a772f675f2f560caaa9110"></a>

<a id="canonical-14567d850c1070d9bb069eb7a699ab2c7c780d6da5c117f91c6a5e47b5adc541"></a>

## value property — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-99418e58367717740cd5ef9e34f2807101c8114463786adf8739692d3902c3de"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add / d4635d4036fa / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cb1bf37cf52e7e1e851e51c0d0ba9427b661c3adba7431144eb3c873c83b5be"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value — http_proxy.more_option.request_cookies_to_add.secret_value / 5193f4655054 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7)
- http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-233cecf31b6368dc87b82d23ea2c093b5d2b7211198027e02d70f78a3c497981"></a>

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

<a id="canonical-63d77e1d4c1e44c9b68109614a69cf244e51788f08fb23e3101e6d2ddc7e9c20"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value / 5193f4655054 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-eaf161e26145bd0d0d182424d7ecc8d0fd467a6dc4b6c74383dd0754c084331f): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-d660aef40c846beba3ac0978883d8c2b31b43f6756e82f8cd8a1ada99dcc3526): complete subsection reference.

<a id="canonical-088d408299d81c41e1ad41ccafd479a4548cf139087137e27601975a82c6e24b"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value / 5193f4655054 / 4

- [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-eaf161e26145bd0d0d182424d7ecc8d0fd467a6dc4b6c74383dd0754c084331f)
- [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-d660aef40c846beba3ac0978883d8c2b31b43f6756e82f8cd8a1ada99dcc3526)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-eaf161e26145bd0d0d182424d7ecc8d0fd467a6dc4b6c74383dd0754c084331f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68595746ac4a5a6bdd360b15f79a051fd972f66411c999febbc6e701638e792e"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7)
- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b)
- http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-203c60fb07e9cc10851474628cab27577ab690ffbd06cc2130c4f7c324396257"></a>

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

<a id="canonical-4016271f3acc42e44182420272623088094871dbf132e8ebcc01e0d73dc187ee"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 3

<a id="canonical-1c07dcddffd3f8d87a98a9dc33481ff14df02968150ec78a31c550378a4a1c02"></a>

<a id="canonical-da012d864c45e3f059b785382e2a7e30d22860c624498a54b06e156f40b7486a"></a>

## decryption_provider property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 4

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

<a id="canonical-45f84cc367d99979a78ba877736fc243e60f80bbbbb96a7c909cf9a383b7f05b"></a>

<a id="canonical-a327f5fa1a4501863dd9fafbcef53702927fcfca0dbb9dde5a71f464b9018bdd"></a>

## location property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 5

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

<a id="canonical-4ad3c6f8d321bdfc0a526101c729bfa062b0c4b0605b7057f329d6bf2f0e8031"></a>

<a id="canonical-d93cb9a78a869646169a96ee28004c873a9190995b6dff05f6df2e8b14d84a95"></a>

## store_provider property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 6

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

<a id="canonical-91346e6be365ddaa21c9e74c28346feffb3fa27bd79487a1f883e43b1be44149"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / b99b209b80f3 / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-d660aef40c846beba3ac0978883d8c2b31b43f6756e82f8cd8a1ada99dcc3526"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f92b32da292c65f4eaa48143cf8949a0de6c25c57fcb240c20e61135151c9ee"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / bcbfdf74f6a9 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-bf6024c009bd46b528356eb81b0ef15301b75ad52e068c2867e411aa83f28fa7)
- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b)
- http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0092b1585cbe918a136e069b189f33b14bac367056ff71efc9ea4396b1b4ac9f"></a>

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

<a id="canonical-6ebd7da95245253d176a9dfa0a3f39edb230879c1ce04f2793e514767821d3cd"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / bcbfdf74f6a9 / 3

<a id="canonical-01330ea03709f123ad56298dee474aa9bfa154a9f4d1b75c1f1a17e141a3e750"></a>

<a id="canonical-28b0eeb190b18ed43106d44d911aedf35757424dc23cfe2c3fd778047baf2b75"></a>

## provider_ref property — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / bcbfdf74f6a9 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ccce5965f4466b2f7bf20869142ce84294cdcc4cd9c37109728c2bc26d01f2e4"></a>

<a id="canonical-2e229fc14f855fe5fd3ea50a4da1a8538590fc0cb64b191ab201e5ad8397fab8"></a>

## url property — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / bcbfdf74f6a9 / 5

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

<a id="canonical-af541150099ac97e6ef722346b4f059ba3c40686a52bf64e1e67aca5b3c318a9"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / bcbfdf74f6a9 / 6

- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-e3fb8218df35edaa725a71ec5055bf7078d245064a60ecf4bc0ffba3a5e3d17b)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a83e0758d807e7d344c71a51566137e0c54f8caf33fabd0a527fa113eb998b6"></a>

## http_proxy.more_option.request_headers_to_add — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.request_headers_to_add

<a id="canonical-064d597b3f5a05dd4e0b1f5c8b99db866ec1d6f375f3293a39f3b72ddd75d14e"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3e6acf6c18aec184db3937c96e01be7f999f11b6ca2dc046e653b3669e52c0a8"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 3

<a id="canonical-7b0c6d4b245e59af2a8e293f1668164290010e6e725b15a0baedee76e6de2f80"></a>

<a id="canonical-b302da7e0e35d01d32d606667fb44d89beb4ea44446dec0207457116051a44c7"></a>

## append property — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-5de0022c006fa4073802c032ac7e6bcc7ac46568cf7b8ba11e51027f2b48f82d"></a>

<a id="canonical-863a1c81f79f4fb6cc9b696e00bd222e2c322f8be8accb6e8a169ae0987a96b7"></a>

## name property — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "format": "dns-label",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116): complete subsection reference.

<a id="canonical-e7da28cfc6a0b4e46d69f6a8842d3402f6cc2008be2c488c277a72a5ca95be66"></a>

<a id="canonical-459aa9be0dfaa9bdce41c79eecdec5e5fed63868336b2d00d443e78d4c6383aa"></a>

## value property — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-a0423e38994a1d83451438fb1c29721d8a0d2892006b5518abc44679e357425e"></a>

## Next pages — http_proxy.more_option.request_headers_to_add / 18521f8c2192 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ce78c9a08d00170cc24a9b4182199de64082404bbd9ee84ea1b296b538d7831"></a>

## http_proxy.more_option.request_headers_to_add.secret_value — http_proxy.more_option.request_headers_to_add.secret_value / 0483c5d62c98 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a)
- http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-067fd92d4f617afb48c258987785c620a695029061af534250407947ce61480c"></a>

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

<a id="canonical-0b62667da89220781f89e4d528659212ee7bd91e8b839359e99b50a45095bf67"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value / 0483c5d62c98 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-ba4cb26cbf262c19107976dbf6f0afa7b9c2bacdbe78872a93383d26e01717a7): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-e811420a1b970d6fc862aebc0e9250146c3a2f0e9ce096f8e306d73e104171fb): complete subsection reference.

<a id="canonical-ca324c6543928ead4f645d5952a0cb48286e93af943cafdc5d022ea85adc3b20"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value / 0483c5d62c98 / 4

- [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-ba4cb26cbf262c19107976dbf6f0afa7b9c2bacdbe78872a93383d26e01717a7)
- [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-e811420a1b970d6fc862aebc0e9250146c3a2f0e9ce096f8e306d73e104171fb)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-ba4cb26cbf262c19107976dbf6f0afa7b9c2bacdbe78872a93383d26e01717a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fe7e10c2d3d361e5ae99bdc6b077ef914841d226b21854697d8adf25dd1f75f"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a)
- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116)
- http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-100c6d6207a86d34c896fbc10f521d8ad81151ce7969178f96e4c1672a1f5763"></a>

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

<a id="canonical-1667d04d76b69e80c79848ff70f3ccce302917598689a302c48389811c9de4a7"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 3

<a id="canonical-05d4afcc62eda94638a388400566fbcda75a9d8d26d5287ff8b074b13ba48f04"></a>

<a id="canonical-a9fc529f5879c5554093352fda5a01562748852731c7d3c8f0ffcbbdb637e4be"></a>

## decryption_provider property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 4

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

<a id="canonical-3017d7c4e22f40fb3aa439113764ce488fffeac36c13ec626de23641a8379a51"></a>

<a id="canonical-7de0fbb41f95a3bc88ffd87db987c3578d76ef6ebbb09710b553a925e6c1c088"></a>

## location property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 5

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

<a id="canonical-f82fa1be55e65a8b038bb2efa5b3f4085524f2ea82f4eca694f42dc04ff2b76b"></a>

<a id="canonical-2f3a8443b261c406b55eaf5f9124241ad92ed406fde416887bcbe65beb408e00"></a>

## store_provider property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 6

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

<a id="canonical-7080507d3b9a37805a2f5c8813f12d69a84e100b8e5e02d0f8f3127fed85f9f5"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 911534656831 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-e811420a1b970d6fc862aebc0e9250146c3a2f0e9ce096f8e306d73e104171fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b0e2e4da36574c2e027fe2b8676506373fc65fde0d4fd1d4429d73ccdddcba1"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / a8e892bb16ff / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-a71a1b8f39f222f9bc37fcf997912b63ee364082a14623fb0d2b457bfecb728a)
- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116)
- http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2194f960e0d3a166c57ecf2dacec7788b3a5696a48cb266d434cdc4bbadcecb2"></a>

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

<a id="canonical-102a68e00a037898cb3b23db643b08e31f7874f04c9c0c917c3bb9746704adc5"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / a8e892bb16ff / 3

<a id="canonical-0499a2b89e7723a38768c9e7cb7afb55d694a60ca5bcd69ad69b0422c1d1f552"></a>

<a id="canonical-432489e04f6b41ff7c49f79345c7a33dc714b627db92085a2476a6b5a3575b2d"></a>

## provider_ref property — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / a8e892bb16ff / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6fe0a207e2bfdb7731df37eb1ce350fef6c8bab6ae2c521c3517833f1976b468"></a>

<a id="canonical-2e3809f34637dc1b6c91023cca7cdb22b588360c552f4068b2e8673af5884a79"></a>

## url property — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / a8e892bb16ff / 5

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

<a id="canonical-74b17193f81bf843982a237da5f1c5ea1ad0ea5fa4db72ec1513dbc5335ca52b"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / a8e892bb16ff / 6

- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-af2d821bb84f9f72282121be8574b5c9cac7cfb916a3b105b667fafd82c09116)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8fb4c9102f33579e177306a9fc12fe8c3cda57a3b4fd4647c0a678ca33bcf98"></a>

## http_proxy.more_option.response_cookies_to_add — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.response_cookies_to_add

<a id="canonical-606ce31f751ada6c13e19bd0f50712c892514c6d47d97683e3ef0dd646ec4b77"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-93f64515ee250911a447ec0819c5e56507ccd140107047309b59c2d2e2839141"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 3

<a id="canonical-97b4604d5cc755ee1a53d9f8ee9388a0abe4a7bf90dcc2e3ffb48085ecfb8fe3"></a>

<a id="canonical-77c5c36d8e87156093f46a40a97be24a8f802389212f78a74fa33d2892671c9c"></a>

## add_domain property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2431521ae261600884774a7152b68e447e4999d2246d4945d555a065993ac3fa"></a>

<a id="canonical-d597f30bd92760f1f5d3840d3d0f6786249597af6ddf921d8269bfc41bd7e141"></a>

## add_expiry property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](data-sources--proxy--reference--group-004.md#canonical-0cb83a30e714b451592077998963584cac36140d50a940e253a1a2c82091e3ca): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-004.md#canonical-b59da45c5dfefcdb407f8667fab9963d5f5e2479003dc3023df49e96deda9e02): complete subsection reference.

<a id="canonical-373f164ec3eb4c79ab5acdb75e8d93562fc867b3ecbf0d051571f0ff38f25c7b"></a>

<a id="canonical-6906e33afb2fa110ec17869535fa3d27f809b76ffdfca73f11dd5776fa0be982"></a>

## add_path property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--proxy--reference--group-004.md#canonical-53322e28e730c265e6efabd4c9a5e6838208ace254f76fec63c885ab8bcf906e): complete subsection reference.

- [ignore_domain](data-sources--proxy--reference--group-004.md#canonical-87f336a5f6cf2dec1dacdd2cd50c95aa660c5160aedd085a29f940d6ead99eed): complete subsection reference.

- [ignore_expiry](data-sources--proxy--reference--group-004.md#canonical-f3de4646a7d636f483536a21d0c51a5298a4c6efe2851054e1b6961aa2b70314): complete subsection reference.

- [ignore_httponly](data-sources--proxy--reference--group-004.md#canonical-9329c3911dfade3f1b3ec260815c62c60aad258ec273a319cfed7dfbd3e656c4): complete subsection reference.

- [ignore_max_age](data-sources--proxy--reference--group-004.md#canonical-092341c2987e11079961f5bd5da7f6ded84e2f4ec1c6edc5db668659ea90e062): complete subsection reference.

- [ignore_partitioned](data-sources--proxy--reference--group-004.md#canonical-cf701a38a9782593f33928d5d9ec04ce2a07b49fee5c88324bf0c15eb53c8c0b): complete subsection reference.

- [ignore_path](data-sources--proxy--reference--group-004.md#canonical-799729b5d91436cd632d31444b3049f03798123592f450b9beeb11f6af2d3b3c): complete subsection reference.

- [ignore_samesite](data-sources--proxy--reference--group-004.md#canonical-03de797d1d84c38de00c72fc3ca8ebb7e5695266c8c8f76a7fbe0d067ac4d32e): complete subsection reference.

- [ignore_secure](data-sources--proxy--reference--group-004.md#canonical-64df96a440c14d90bad6b824f20c4068a9a7035b42699ce41bf8430f06f859f2): complete subsection reference.

- [ignore_value](data-sources--proxy--reference--group-004.md#canonical-c8f3cf3c9b83bea4fb7b67ee8ccc789277ee9a458747bbcdfcee7b1215ea4765): complete subsection reference.

<a id="canonical-ec43c45e4d102d71d3c1bd219bdcd45dec0bd3912b0dc7af7901647f3c0c4a0e"></a>

<a id="canonical-af25775451e419ac5f6a8b3e32953b4e0f7c491b0a87dc3dc533bdea1f932e27"></a>

## max_age_value property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-a3ad9b64684aed33e7cb929b0f7bc7ba88b247bed4c8e459052a1f6b81bd36d1"></a>

<a id="canonical-57bdc4c06d9f6872bab9392a91a2c1bcdb99ba93c85b48be4f64e8a0003dc4c6"></a>

## name property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "format": "dns-label",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-cefe23f09ddfaf006c3008e79d1f3b6b2b7040be733bc585c2771638ae96d893"></a>

<a id="canonical-b3768cfab53c142cc54c8123e1e351c3ca710173724f7e34f23aedf3ad2c1527"></a>

## overwrite property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 9

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [samesite_lax](data-sources--proxy--reference--group-004.md#canonical-f0b98c00557cd6eb544fc5a8ce07b567c2dfc94bc9d22b0efd26b9b26901685f): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-004.md#canonical-588ce994ece5f4acb99ef536f737991a1225da0d3265d24d41ecc1d672338fd0): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-004.md#canonical-cc6162865fe695da0f2fe256c4f3ba8523bdcdd0222430b434f0a99609f1a9dd): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd): complete subsection reference.

<a id="canonical-2a31a9c9b667f41f8c1d9e9d3b62547e1ad1b5334035c60d6c96f0a476592325"></a>

<a id="canonical-b25ffb5715b64830dcbc780304975248fcaa5f14b26fc87b3c5d6c076a707d93"></a>

## value property — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-44cb5684dd46e67a4b732c74dfffdafba84cfa08b851056ae24a0db656cb07f7"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add / 38bcd7fbc788 / 11

- [http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-004.md#canonical-0cb83a30e714b451592077998963584cac36140d50a940e253a1a2c82091e3ca)
- [http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-004.md#canonical-b59da45c5dfefcdb407f8667fab9963d5f5e2479003dc3023df49e96deda9e02)
- [http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-004.md#canonical-53322e28e730c265e6efabd4c9a5e6838208ace254f76fec63c885ab8bcf906e)
- [http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-004.md#canonical-87f336a5f6cf2dec1dacdd2cd50c95aa660c5160aedd085a29f940d6ead99eed)
- [http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-004.md#canonical-f3de4646a7d636f483536a21d0c51a5298a4c6efe2851054e1b6961aa2b70314)
- [http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-004.md#canonical-9329c3911dfade3f1b3ec260815c62c60aad258ec273a319cfed7dfbd3e656c4)
- [http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-004.md#canonical-092341c2987e11079961f5bd5da7f6ded84e2f4ec1c6edc5db668659ea90e062)
- [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-004.md#canonical-cf701a38a9782593f33928d5d9ec04ce2a07b49fee5c88324bf0c15eb53c8c0b)
- [http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-004.md#canonical-799729b5d91436cd632d31444b3049f03798123592f450b9beeb11f6af2d3b3c)
- [http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-004.md#canonical-03de797d1d84c38de00c72fc3ca8ebb7e5695266c8c8f76a7fbe0d067ac4d32e)
- [http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-004.md#canonical-64df96a440c14d90bad6b824f20c4068a9a7035b42699ce41bf8430f06f859f2)
- [http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-004.md#canonical-c8f3cf3c9b83bea4fb7b67ee8ccc789277ee9a458747bbcdfcee7b1215ea4765)
- [http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-004.md#canonical-f0b98c00557cd6eb544fc5a8ce07b567c2dfc94bc9d22b0efd26b9b26901685f)
- [http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-004.md#canonical-588ce994ece5f4acb99ef536f737991a1225da0d3265d24d41ecc1d672338fd0)
- [http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-004.md#canonical-cc6162865fe695da0f2fe256c4f3ba8523bdcdd0222430b434f0a99609f1a9dd)
- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-0cb83a30e714b451592077998963584cac36140d50a940e253a1a2c82091e3ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0072ebf7534634345a7e93d910e8a136815c0a96120e5a430e0eb90cb7a0a1e0"></a>

## http_proxy.more_option.response_cookies_to_add.add_httponly — http_proxy.more_option.response_cookies_to_add.add_httponly / 95299ed81112 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-7d08a8fd4fff4e0f81875904393935656fa3556b704ab6d5483101426125b0fe"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-d29c3b973f2aaa3cb053f707a97972ab1ad8264a491c6adcc8732fbad096126c"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_httponly / 95299ed81112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69cd812205b493b240f8cc0e42d0c389743e17751f6bcdd02b37a740b840a28a"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_httponly / 95299ed81112 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-b59da45c5dfefcdb407f8667fab9963d5f5e2479003dc3023df49e96deda9e02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2c5467e25229c69b2ded9eb5290335aa2dba9cd38e2182fa52d1ae09f1e90cd"></a>

## http_proxy.more_option.response_cookies_to_add.add_partitioned — http_proxy.more_option.response_cookies_to_add.add_partitioned / feb8a7c9107a / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-08209eb1bb5bc53fde378771cc5a3a120d8442c3d4bff80a29f3bd48982e7d4e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-94016722ba39ed1ca3d1303b0e30636e540fc9d944058216c7d9586a3fe404bb"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_partitioned / feb8a7c9107a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-024e58166c51f527b6427523c6f239c517470cc5e4763cb41867819c0e68425c"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_partitioned / feb8a7c9107a / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-53322e28e730c265e6efabd4c9a5e6838208ace254f76fec63c885ab8bcf906e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b40b9117bfa8a33e94500ecd0f0287e22c7175cc1e52ef41c6bb433a39a95ed4"></a>

## http_proxy.more_option.response_cookies_to_add.add_secure — http_proxy.more_option.response_cookies_to_add.add_secure / dcac024ee2aa / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-c5a604f16a90387d354be6183e10d67ebc34c0ba92caf7435d454269033a532a"></a>

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

<a id="canonical-a6b9547870498b89603f3b79b18504b824182764fb0c39dcecaa55d8f49a62ec"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_secure / dcac024ee2aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c36e7aa44e9312dd8592225d53670b592440db83e874b385c75b80e4e76f52bd"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_secure / dcac024ee2aa / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-87f336a5f6cf2dec1dacdd2cd50c95aa660c5160aedd085a29f940d6ead99eed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd4409212d3ec624e9685083376b91958f7ccfa74905504bbf844e1f417d8028"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_domain — http_proxy.more_option.response_cookies_to_add.ignore_domain / f5d3da1821fa / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-5574b1e3163d905c719a6a66724eb1afa8fb921c6f8cd73ca1d0a47951e9a63e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-45a89a7620b22136c46189bede953bb926d9cdb8b7118fd8fbb6802ff2cc2965"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_domain / f5d3da1821fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07584eaac43c7200e1027e63819f184f0a33a6ee2bb76c9117fc08394a3c9ac1"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_domain / f5d3da1821fa / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-f3de4646a7d636f483536a21d0c51a5298a4c6efe2851054e1b6961aa2b70314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3c85f3fb8ac7719beafb96fa2b893267cecb2edbc4420b7cc4c0194127e4317"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_expiry — http_proxy.more_option.response_cookies_to_add.ignore_expiry / c3cfe38d1a6a / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-968217cd0f38d5f0175478c695cb6a68bc48c3096bf118b83deaa8278517859e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-9659bf997a2896a826850986730d3b80c29a7c2a1ae66f51dba073779aad5513"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_expiry / c3cfe38d1a6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3242cdd38ffa9ddcb473b402b2638c2dece6e2b50997e5f07b13e9669aa4c748"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_expiry / c3cfe38d1a6a / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-9329c3911dfade3f1b3ec260815c62c60aad258ec273a319cfed7dfbd3e656c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd72f246e58662bea2b0a3f26e521bf402463e79ad0c580b16c22d66f798c12a"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_httponly — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 9102fb8afb43 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-b87a9f030516e794cdc783b7ffd408f200e1585544c48ead55bbeb66b239f609"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-17d5fde5fa40ac04c05a6b48370cfc1080e6790ed1095c2a20e92ae9c31bd5bc"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 9102fb8afb43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f5e69774754151981f7bf9e77d468d0da545d23f2691f31232a3fb78053e6dde"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 9102fb8afb43 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-092341c2987e11079961f5bd5da7f6ded84e2f4ec1c6edc5db668659ea90e062"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-107982ba829cf6c49ea937356123bf7e8d6db371ec794d5a925111b29a771acc"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_max_age — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 5ef4bb00d917 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-65e48b18ba1c33277ab226bfa78e4a167cfc954ba239c9dfc24c00a769cb28d3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-c5aab8e67957299e36f9868b2b17236756fe85b0af0d47d18a8f75d9bfba1c99"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 5ef4bb00d917 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a661096f874c14bfd4d21f29c218b88b90b134c29bc73af62cba5c006c6696f"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 5ef4bb00d917 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-cf701a38a9782593f33928d5d9ec04ce2a07b49fee5c88324bf0c15eb53c8c0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f65d5d1461becd1a3d514e66aa2b11a31f1c742c622eb7c4407c3fd65b95a49c"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_partitioned — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / eb87c679e93d / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-1311df8cd4adf5e3c517a62d056b9ffab03dea838cd48ec285bf9ceea3e3c65b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-35071ce4b748e9b77363683c4bd6b6b91a1d05279325c428b003ba75c45aa8b5"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / eb87c679e93d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-804236a895e7e98fe3912e6fe1110aa9e8399703ba50b967d34577e980e9ebbc"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / eb87c679e93d / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-799729b5d91436cd632d31444b3049f03798123592f450b9beeb11f6af2d3b3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d28b2506e59c2091eace07fc52ed87349f2e6435aa57d2b1e06d792c008cc900"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_path — http_proxy.more_option.response_cookies_to_add.ignore_path / fca28fc199fb / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-7806ae27347760bbaa8fe8b6070d4158f96b94451041dccbbdafb485857ae3ba"></a>

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

<a id="canonical-9415bbf197e2c7b433b5294b9a12e4b3745366ae8a3d2e3a69eb84b7c9e8c240"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_path / fca28fc199fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5b6e228286082fecc4c61b3ddca4592f2eed40a6bb12b0ff0810414e7094712"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_path / fca28fc199fb / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-03de797d1d84c38de00c72fc3ca8ebb7e5695266c8c8f76a7fbe0d067ac4d32e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-095c12d4e984485b6e7ff6c4a0ff984e431816480c82ed882807521831e3d6c6"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_samesite — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4c037b763d8e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-fc446e341ea494d1d3da7d78ea75d8eb523e50e84f693c2cc4117f1d0c7e1456"></a>

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

<a id="canonical-76e53a0ad046d25d5b615c50da1577cfe670672de328832b0c8b19e58c1acccf"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4c037b763d8e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64e1824e4d09824f265aa2e16b237941e65c47d46c4c67b673dc051bbeae5210"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4c037b763d8e / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-64df96a440c14d90bad6b824f20c4068a9a7035b42699ce41bf8430f06f859f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac66ebc29ed66b58ff8f67b826e9396c24599b9b587d4c185ec0fdf60aa57d8b"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_secure — http_proxy.more_option.response_cookies_to_add.ignore_secure / 25345cb9b1b0 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-3006de3e9e1fec4acfbc33100aa82f4386e8895425fd7c5abd1254d06417eaa2"></a>

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

<a id="canonical-e4974338669e1b6d0ae1b12cfd2ececc7697593f29dbf539db9907fb181dfb97"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_secure / 25345cb9b1b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae2306d82d552628ee15c027ce9a71224bfe05e3abf6a5d0ac48ccd9dddd3733"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_secure / 25345cb9b1b0 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-c8f3cf3c9b83bea4fb7b67ee8ccc789277ee9a458747bbcdfcee7b1215ea4765"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a04b1d23764908552962c289dee7696e6d6410bc18accf98ba1dc927e142596"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_value — http_proxy.more_option.response_cookies_to_add.ignore_value / cf83edc1db03 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-c9399f9b4bbea3db90e48be592b52e73762861f138362c0978f4cadb3e17e1d6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore value.

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

<a id="canonical-567338c3901e932c77fd75f3c3ad6cebba9fbf0c74d1889ce38e68ffc3e4cb77"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_value / cf83edc1db03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9be8552ad8a721cb5d80f086d91dd20f11683a9b9a98132d72bfafc31c5d2069"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_value / cf83edc1db03 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-f0b98c00557cd6eb544fc5a8ce07b567c2dfc94bc9d22b0efd26b9b26901685f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e19420b471489d8ba038343d3ca8ed8a81109da3c5ec3866a85698b5f8ce257f"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_lax — http_proxy.more_option.response_cookies_to_add.samesite_lax / fd5006c2c53e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-10f7676365a1ed7f92205b8952c8c861c68c9b31dfe326ee259eafab95d3ee22"></a>

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

<a id="canonical-0f34564dce0ee55da8ee2eb65e093695fc9ac464e54a396d121a58b3b15fa479"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_lax / fd5006c2c53e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5b8a725c4abf44c69bcf0eed8937b870d999579d9a093d01a4051145e59e184"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_lax / fd5006c2c53e / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-588ce994ece5f4acb99ef536f737991a1225da0d3265d24d41ecc1d672338fd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9489a08fdfe49f7c9fc16b7fadf327b1b2e3f61078e4661413f8f80d017f38ac"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_none — http_proxy.more_option.response_cookies_to_add.samesite_none / fc6b528e38fd / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-e55b8b25e285e522c9ac718464b87a0c3d5d822f6f81d9379e64f6ac0c118cf1"></a>

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

<a id="canonical-b59c87e3fc12c28507669f0284eeee78fea248e761947086756fdfffc15b2c3e"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_none / fc6b528e38fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-987c32bed3e49b85ccf4e485385cd3f5f24a7c32eb13fb95413c79582fb25f39"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_none / fc6b528e38fd / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-cc6162865fe695da0f2fe256c4f3ba8523bdcdd0222430b434f0a99609f1a9dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b03575779e10447a787d49d554c9d6299b2d859eb8a6b74d317ae1dbea83775"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_strict — http_proxy.more_option.response_cookies_to_add.samesite_strict / 39edcd07f471 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-f4c7a0f2084d4f30028be9ae31192bc20d89cdcc5eef83cbe12250d10888fba4"></a>

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

<a id="canonical-13676f87336f2948c125d9b3a31e0617f7e73c15081a6edff0f037295a0a1859"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_strict / 39edcd07f471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98292a4f65190e67da8514ac28d6e0c2c1ffa5229718b029a1a98cad117e1091"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_strict / 39edcd07f471 / 4

- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ca5009063e2a236f74f03040fe1c477a80f2c254d43e27ce61876d3ea342680"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value — http_proxy.more_option.response_cookies_to_add.secret_value / 41c83771012b / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-3c6bd7fe5e5aabd1986d3620a8886e3d70c48affd1e2089314bbb022239a182c"></a>

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

<a id="canonical-b30a1d1df4df380f093d85d317a6ab56ce6c9c7eafd24c9aa6f028e8527701aa"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value / 41c83771012b / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-b4b185bf7638ba9af3412df2c980fb3e7987c1531de23eba254364d921d0d042): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-70de69e5732269fe93dc010a290ed93b380a5000c1b001872b94fed97c283b26): complete subsection reference.

<a id="canonical-b26e6c49855572b882564fb3f5d13df24baab32fb7ce378ef35cb5c710a236a2"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value / 41c83771012b / 4

- [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-b4b185bf7638ba9af3412df2c980fb3e7987c1531de23eba254364d921d0d042)
- [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-70de69e5732269fe93dc010a290ed93b380a5000c1b001872b94fed97c283b26)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-b4b185bf7638ba9af3412df2c980fb3e7987c1531de23eba254364d921d0d042"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccf736ae3e493827e674dbb2a485359c5488137bb83dad604a82364a94c25729"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd)
- http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-56ca5041376353dc6292d3fb0eb2d13cf8e4dd29af5ed22213a9b060b3899b88"></a>

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

<a id="canonical-03549360f196e221c3cf27825c7b2ee0cb2eab2671c8bb7a08e22619ba83afc9"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 3

<a id="canonical-409cc1d0c056e031b8e662e0cbda2b8c5e24662351c456e5f3f7d86d407e4b94"></a>

<a id="canonical-474d05cf01c4ae130243a26b3850eb2455d41b698349f8e401a128bdcd1ae24a"></a>

## decryption_provider property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 4

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

<a id="canonical-0449e63f6dd52d5537800600931d541338a8a44ec46165c2f4d851b0befe470e"></a>

<a id="canonical-e9f70b9077704051bde9ed07b626b45736aaa905b733c36964b178db9c196dcb"></a>

## location property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 5

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

<a id="canonical-e6b18a14d4a6d573262ed7d32d10b9f1a6a05297925162e3b232178f663d1874"></a>

<a id="canonical-5a644855dec84cad507f4d8503d54bc76eb07bff59d3e1c3c012dc6696b30e4d"></a>

## store_provider property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 6

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

<a id="canonical-ebebef9c807fdf565d197109b68158a2081ca018254caefd355dad7d61207035"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 8f73d60de4e0 / 7

- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-70de69e5732269fe93dc010a290ed93b380a5000c1b001872b94fed97c283b26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d12df6580fb078bae7311f0fb7322367a21ab72020dded67713a51204447a28"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / f4324a282d97 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-06e3e6c13e4445eec1255bec07a39482c027b650c137c676387d59b0d0a60aba)
- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd)
- http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-bce1d7c3a37015c5fdc5163c6a7f1d66985309cebf5c3190d44664e8d26fa3dc"></a>

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

<a id="canonical-6003ed7e745db73847adbdd0d8600e4d00786a445633e241323218340c5532f0"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / f4324a282d97 / 3

<a id="canonical-4e04bfcad35014799d972d6a7f565366b619a99a10132a6830c1a026d885cdc6"></a>

<a id="canonical-85398f37572b2c6272c6b13e1dbcc87f65f6f24ebecfe73768ce16fd87e29e33"></a>

## provider_ref property — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / f4324a282d97 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8f762dc3e39543ac7fa7efe104fbb4da91195c75a0c1641f29ac8d4e8b7482d0"></a>

<a id="canonical-7892bc996f1c7b25410374cac15a7936b69352e6cedbb31d9e7ffa5c6ac2fc35"></a>

## url property — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / f4324a282d97 / 5

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

<a id="canonical-175d764c221519777f40a85ac046a0149be8b02144211c2cd3e1bcc62e8a6d05"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / f4324a282d97 / 6

- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0c4f89734e048a9e98a7064cd6411de43b0ffb84ba8adce2f1c87ea5384c66fd)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97380b21531cf1ae9a0ea3db6d1d15726424f29c3fe2460c28c37e7fbd4ac49b"></a>

## http_proxy.more_option.response_headers_to_add — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- http_proxy.more_option.response_headers_to_add

<a id="canonical-b04d9684c1c4f12d1f9ffb66d166cc5a7c6ec270742a6d588d7735649e528dc1"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-852b49c68318609f7cb9de98664ee6e229248af933e84282a9d2ee8dbb67fe23"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 3

<a id="canonical-f7197d71e7516cacc49090c1a933cc4e92acdd904746a295b904dad5dff1c064"></a>

<a id="canonical-b4dbd17853392812062cc5ce8b9d231e490533a309aa45311699e9865767ade5"></a>

## append property — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-1beb3234f82b5a81312aa53bccaf850930e7bced68f50106c75cb364ae8a0121"></a>

<a id="canonical-554625dbeb0326897efca75b42568debd0da3a6133cfa0f13b1c04d6a181db98"></a>

## name property — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "format": "dns-label",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305): complete subsection reference.

<a id="canonical-33ba97422707d69d098c749818e1b4920354f787482d1ef729830dfc7469845c"></a>

<a id="canonical-daf6f011e95ea133621770e9c57b2dc73a96c93a761a2048b3c68ebb966c81c5"></a>

## value property — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-c4118e40a40289e3996fce42153849194ef9b987f8c1d485c5d2b5e14ce4a707"></a>

## Next pages — http_proxy.more_option.response_headers_to_add / 91d9bcf3b17f / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4e347ccff3a370aea8e6bfc786575acab6d75208bdd35800a0f8e8a9733af80"></a>

## http_proxy.more_option.response_headers_to_add.secret_value — http_proxy.more_option.response_headers_to_add.secret_value / 42adbeecda0d / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-c129bb9958164828e280461cbb6f48105958856a5735157b575c323f31ef83f7"></a>

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

<a id="canonical-c066cad9a339fb9cfb6f9407cbf096c85eaf1da059156255abd1466c9189852e"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value / 42adbeecda0d / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-f2a5648db3253f89fe68f13646a2f0a0c39d85183b2fdfee22e0f3057cc4b316): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-4835f6a68de6bc848d2a95613999085931440c71d75a6ab62fd548295026038f): complete subsection reference.

<a id="canonical-1ec6e08b83ac1d91f861f21b28930e0e123679a65c0de50c9d544743d7cbf9e6"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value / 42adbeecda0d / 4

- [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-f2a5648db3253f89fe68f13646a2f0a0c39d85183b2fdfee22e0f3057cc4b316)
- [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-4835f6a68de6bc848d2a95613999085931440c71d75a6ab62fd548295026038f)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-f2a5648db3253f89fe68f13646a2f0a0c39d85183b2fdfee22e0f3057cc4b316"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c779dc4cf6bda0d2e822b0d640bfecb1a5f65fceaa6ec66f87ad4c5c0b7641f"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397)
- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305)
- http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-d096c1d687241792d52f6e08f1b5c5943911cd26e7a91a3cc983e4e5f12c52e5"></a>

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

<a id="canonical-f70f10bf8294ada4bbd962f541ae62fee471bee4be9cfe3fd25488f8f08524dc"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 3

<a id="canonical-b6bdb3fbd4d7254a16444072e589bd595953e5452e21615db94bd36453defc92"></a>

<a id="canonical-6871005ffa491e83099883b92f0a66705348e1fad948ccbda264c47242007a0a"></a>

## decryption_provider property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 4

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

<a id="canonical-46db97ad85f2dcfe32d4b8b2834d1774256c3ad7cd04d48fff4496861eecbaec"></a>

<a id="canonical-ac678dc0b35904eed43418ee75614b6ded6e32435b8d63b9c3fab26750c04852"></a>

## location property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 5

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

<a id="canonical-5a35879f8e4f9c3600281a792c9bed88c9e84146b5732e6dfc5ef15f831e975b"></a>

<a id="canonical-1bb05ac3eda310db259d2110a2f2900b2fd4bd4375a2349f97454cea2a11d92e"></a>

## store_provider property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 6

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

<a id="canonical-1263b0b9a53a85c5ea4a7c7683e5552e80784ecf203bacafedebd6984af1a0dc"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d817f5a49120 / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4835f6a68de6bc848d2a95613999085931440c71d75a6ab62fd548295026038f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-791fc390337d1c12d93c5f89853fa87a639c8194502ae2a49b29f9f008ea8bce"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 875faea73c3c / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-094f977f1a2f69ae26c51c813a5009ff019d85cd7eda77abf9334acef961a2ef)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-e87c8f0a6ecd2174573dd507d0a20366c98f58779a27953814b99c8eb57ba397)
- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305)
- http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1309170c61368ac41bd54fe8e8f452f19e6f8d3346abaad62c9244961b664425"></a>

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

<a id="canonical-1a621b1c9e64feb70368e4e965d409415a6175b3fb920e72b1b2e8d4f5796395"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 875faea73c3c / 3

<a id="canonical-0056204a972b55057afc6b0b90b45f2dae15d41697f22ba6d58f264c47962e0c"></a>

<a id="canonical-51255edfaebe9554549f70fa5d36a2b39af0516d0cba043d191cd099382eaa42"></a>

## provider_ref property — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 875faea73c3c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2881c4bdbe7c8d28912788467589edc1e151c3a96de5e72dcf42f7432ff5cbad"></a>

<a id="canonical-7eb8f2a3488880097312a79d4bd398dc6feed47776c5b45705f4b58ff6b69b17"></a>

## url property — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 875faea73c3c / 5

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

<a id="canonical-d7be0192d63d183bfa88b9cbcb9a8db0f85fbcc3babb3662dbaf56ea161a0128"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 875faea73c3c / 6

- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0bdfc8e448c801a870998587f4e4bf58183830b4afbd93ae5b25fb9be4cc3305)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-354eed661a275799f673f7dd3b219dd0d9c148389186b9bd151c99738179e6c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e06a6e479aded0e5fc1a5fa7759948c4cbd804b0611331bcfcf6c37e8c7b8682"></a>

## no_forward_proxy_policy — no_forward_proxy_policy / dc5d7069a897 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- no_forward_proxy_policy

<a id="canonical-b4151087faedbfb9198d1954ec55c0ac5c869e873ff01d3dc8c4c316eac84f12"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-06ddb5bc6befa52e8074aebd4d99d92034e3c9b057117cdb4035ec217023a3ec"></a>

## Direct properties — no_forward_proxy_policy / dc5d7069a897 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-022e1433e1829af7165e1a368252bcc331069edc47d3136752eb4a3ecc70d3d5"></a>

## Next pages — no_forward_proxy_policy / dc5d7069a897 / 4

- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4cb287101d7cff732a91ce780d465332c4fc488b430e1c588c72d8d91fff4f3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bad12b260194eda2b7d9e42248c877bf8975d9fec6012eac9edd7660e06acb8"></a>

## no_interception — no_interception / 29d304063bc8 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- no_interception

<a id="canonical-472d438ff2425ac3470397b8ccc7afe368e24274f0df9164bff253696d9d3e23"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_interception, tls\_intercept; Default: no\_interception\] Configuration parameter for
no interception.

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

- [no_interception](data-sources--proxy--reference--group-004.md#canonical-472d438ff2425ac3470397b8ccc7afe368e24274f0df9164bff253696d9d3e23)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-99bb678c4909c834195d7cefc40f5c76cdae7ac5abe2820b0fc8a25451848886)

Select alternatives according to the provider validators above.

<a id="canonical-8c148af2b4aab9877c0804a86a3d9911b6a599088b375541c8264508c4b1157d"></a>

## Direct properties — no_interception / 29d304063bc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4001017120f14af4bd50d6f0318ff543e09ea907d54f8d683555b6c8c18a9a8a"></a>

## Next pages — no_interception / 29d304063bc8 / 4

- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-511ab4dca177ca5d7d292f2efc016dc5ab5e4f01c63811a03340a2195219d67a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5608ea4a4ef2719c1b1b00abe82fc020e11bcce28d1f1666247b18ff38ee7eb7"></a>

## site_local_inside_network — site_local_inside_network / 25f857dd9d8c / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- site_local_inside_network

<a id="canonical-c5fc7b71f5016f4b6f581708b16103c722f0351bd6c961311178ce76691801e2"></a>

Type: `["object", {}]`. Computed.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

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

- [site_local_inside_network](data-sources--proxy--reference--group-004.md#canonical-c5fc7b71f5016f4b6f581708b16103c722f0351bd6c961311178ce76691801e2)
- [site_local_network](data-sources--proxy--reference--group-004.md#canonical-98d9bb15e9aa2e43fb2f0517b15cb21a33e860eb2bd9b73f722d31e9a62f5399)

Select alternatives according to the provider validators above.

<a id="canonical-c732faa2f5d7495d015940d0c573c5294005a6d6bd4ff44a2cb6935fd7e01836"></a>

## Direct properties — site_local_inside_network / 25f857dd9d8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25d74c96ee6932ecd88e4e39172a35e75c5066d22b9061ba81a8b85985c8b22b"></a>

## Next pages — site_local_inside_network / 25f857dd9d8c / 4

- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-7772efa695c7a5d2abe78ba96890f5435cd65223acd8167714b620ad5f4bda05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebdbf13ca7cf290d3205cc19efef58789f6cc6dc9bf5693d77f73d3cf550321b"></a>

## site_local_network — site_local_network / 621e20587a6e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- site_local_network

<a id="canonical-98d9bb15e9aa2e43fb2f0517b15cb21a33e860eb2bd9b73f722d31e9a62f5399"></a>

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

<a id="canonical-07335a00491bf8389c7955c6ab79a3c1b31dc68869a63dc6efe6e4ce69cad486"></a>

## Direct properties — site_local_network / 621e20587a6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a3bde36799a0a0c725199ac399236cd43749bd1daddddd8ec19c15f78ad31200"></a>

## Next pages — site_local_network / 621e20587a6e / 4

- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-218153d3cf1552ea2210032431b21458d8cdf24b3895304dd628940c2dbce6c1"></a>

## site_virtual_sites — site_virtual_sites / f93577427be2 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- site_virtual_sites

<a id="canonical-16f7c6653a49b2c94edcdda8f045fd6277695fc329d537369a57aa90a3e7193b"></a>

Type: `"single"`. Computed.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

<a id="canonical-4a412a8cd75b7d6d3e1fcfdcce88f3be9ba60c9e9e4bb63e0d7992b9afddb61c"></a>

## Direct properties — site_virtual_sites / f93577427be2 / 3

- [advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175): complete subsection reference.

<a id="canonical-5a4748c9f2ae1c619564f2704fdaf1927a9968f938523479c0b79e04bb3bcbd4"></a>

## Next pages — site_virtual_sites / f93577427be2 / 4

- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c2b4fe11adb41f69c44447a93fb074f84bfc9603944fc1e05ad6c37fe7732aa"></a>

## site_virtual_sites.advertise_where — site_virtual_sites.advertise_where / b5d194c96e27 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- site_virtual_sites.advertise_where

<a id="canonical-60ab02db3202b6f59ba79bfa44fef9a6b98b2c3a86e664bdf34d4c10f3e18e16"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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

<a id="canonical-59a0e61bcf931b4072bb146a5453c58bbfc5a71193f7d0d390072399657e2934"></a>

## Direct properties — site_virtual_sites.advertise_where / b5d194c96e27 / 3

<a id="canonical-43c271745d0ad1f68128d1f6199146d06460b1b097c25228700d1c8a07a7a43d"></a>

<a id="canonical-e470902937a88f857682c74c5818f897dc05a0ffd723ae29f1b18584c3757f26"></a>

## port property — site_virtual_sites.advertise_where / b5d194c96e27 / 4

Type: `"number"`. Computed.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Upstream description:

Exclusive with \[use\_default\_port\] TCP port to Listen.

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

- [site](data-sources--proxy--reference--group-004.md#canonical-7f89a60b89cdf29fb420f748aecb6779d5995cfb0aa96c0a6441bec12232429b): complete subsection reference.

- [use_default_port](data-sources--proxy--reference--group-004.md#canonical-fb12d69618786a97fe90759af145754e82ee41c579c8dc15d20706e2955063f6): complete subsection reference.

- [virtual_site](data-sources--proxy--reference--group-004.md#canonical-0503ef9cdc1ad966e25baa97d5b1e92b90def581d6994baa8d24180df4d369c1): complete subsection reference.

<a id="canonical-54daae34b556a6d22038b0a82b04ad9e9960ba6ea5ecdf481174d8dedd74c2a5"></a>

## Next pages — site_virtual_sites.advertise_where / b5d194c96e27 / 5

- [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-004.md#canonical-7f89a60b89cdf29fb420f748aecb6779d5995cfb0aa96c0a6441bec12232429b)
- [site_virtual_sites.advertise_where.use_default_port](data-sources--proxy--reference--group-004.md#canonical-fb12d69618786a97fe90759af145754e82ee41c579c8dc15d20706e2955063f6)
- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-0503ef9cdc1ad966e25baa97d5b1e92b90def581d6994baa8d24180df4d369c1)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-7f89a60b89cdf29fb420f748aecb6779d5995cfb0aa96c0a6441bec12232429b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65a2c95d55d91617a11ddbb4ccd0ecf5ed04448fff1dd87c56e1f4d61ec8bb48"></a>

## site_virtual_sites.advertise_where.site — site_virtual_sites.advertise_where.site / 8f4e9984023a / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- site_virtual_sites.advertise_where.site

<a id="canonical-86b9b90e9c0653df335ad2284dcd2ee45a51f6baa871260a86abc18f85c3a78c"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-8f43511566e5086b9024ad372b362d10cf983e7023a7ef885ace7a69cc7ff0ab"></a>

## Direct properties — site_virtual_sites.advertise_where.site / 8f4e9984023a / 3

<a id="canonical-7277efd852e6f788b2dab7eb2c8565ff8b29102bd8849643985887d7ce2ab29f"></a>

<a id="canonical-221718578ea5dc7dc8363fdf7a3df283639558d4aeb38ec33c70230af02c67d3"></a>

## ip property — site_virtual_sites.advertise_where.site / 8f4e9984023a / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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

<a id="canonical-5afbbcf4f6bcc13fa9401d812a80e1b977b8080e60ffd270427db09109ea17ff"></a>

<a id="canonical-2112fa7789a09b04f8892549ab6271d8c83cdb7125df9d0fa4ee4cb55021211c"></a>

## network property — site_virtual_sites.advertise_where.site / 8f4e9984023a / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--proxy--reference--group-004.md#canonical-3f2efe133f4814c53cb85efd300d3dd948df970c18504c27c1d01a8a4a04a1e4): complete subsection reference.

<a id="canonical-d7620d20dfbdc73c78a7be6f554f2240a646d6428f2fba470ec339f84722e9a5"></a>

## Next pages — site_virtual_sites.advertise_where.site / 8f4e9984023a / 6

- [site_virtual_sites.advertise_where.site.site](data-sources--proxy--reference--group-004.md#canonical-3f2efe133f4814c53cb85efd300d3dd948df970c18504c27c1d01a8a4a04a1e4)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-3f2efe133f4814c53cb85efd300d3dd948df970c18504c27c1d01a8a4a04a1e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-358b74b2c5d909dead4e11c42745be88620043e675b719a9a75b2988e88b6fd8"></a>

## site_virtual_sites.advertise_where.site.site — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-004.md#canonical-7f89a60b89cdf29fb420f748aecb6779d5995cfb0aa96c0a6441bec12232429b)
- site_virtual_sites.advertise_where.site.site

<a id="canonical-865e01e4b5ab47086ab0d6c006b7412356adfc7c7c2b80335d9cdc8d7059cb2b"></a>

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

<a id="canonical-3a59d037ea2dfb073e059574ddf6ae373fb306ae9eb07105a9b77e19b2a6e4ca"></a>

## Direct properties — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 3

<a id="canonical-06ad849fe759bf38a7fb5b83bd5cd20be0d38e2b57a70f34e94f54e0bf6468fb"></a>

<a id="canonical-42de419a0a713d7f4b1472c116b2d6ff12c228630e74d0e5043f3d6880e02258"></a>

## name property — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 4

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

<a id="canonical-8f6a9c2d16fae0e6a960088009cfe057916bbeb5a7d3bdf94bd81a709fe37132"></a>

<a id="canonical-9c618e281c495e5b63b0371d941f4248d2f904d93ff50115c793e630cfe2892c"></a>

## namespace property — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 5

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

<a id="canonical-7d783636b4a51589d0fd7fd3f783c0fe9fe9bfc076d99f4d9daba22641cdde4c"></a>

<a id="canonical-e6c5fa682a0d87e2b0bd3ace3c23dbd4320fd5025866409f72137c5736d33d21"></a>

## tenant property — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 6

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

<a id="canonical-66c30698a861db8d4e8820e60cb857d9fb7044172d08fbcf6875840cd6f9b2e0"></a>

## Next pages — site_virtual_sites.advertise_where.site.site / 2f1fb26f4586 / 7

- [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-004.md#canonical-7f89a60b89cdf29fb420f748aecb6779d5995cfb0aa96c0a6441bec12232429b)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-fb12d69618786a97fe90759af145754e82ee41c579c8dc15d20706e2955063f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad5be8a845e932613dd0ade31114b49836489134976079cab28f13122f4bd101"></a>

## site_virtual_sites.advertise_where.use_default_port — site_virtual_sites.advertise_where.use_default_port / c997752c489b / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- site_virtual_sites.advertise_where.use_default_port

<a id="canonical-d7bfda3ac2b63892d619223bc212a3574847654bc8eb079baebc35ed5f94335b"></a>

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

<a id="canonical-1850db679d49c214ed0a9a8a2e5ed46e5d7024eecad0bbbbbf577b546b0989b7"></a>

## Direct properties — site_virtual_sites.advertise_where.use_default_port / c997752c489b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4dbdd9113a0960d409871d6d48d71c00320706b9ff2f73530ecfcac40af5b9de"></a>

## Next pages — site_virtual_sites.advertise_where.use_default_port / c997752c489b / 4

- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-0503ef9cdc1ad966e25baa97d5b1e92b90def581d6994baa8d24180df4d369c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6db350f9d1ae6b51cac33c390605032383b762f0059ddb27f3ef16b06cbc3b94"></a>

## site_virtual_sites.advertise_where.virtual_site — site_virtual_sites.advertise_where.virtual_site / 6eb5ce37d172 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- site_virtual_sites.advertise_where.virtual_site

<a id="canonical-1ac96376af346de2aaffa5294e4160946b7624b783ebe5b2858ef6701d115231"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-120d6a4752edcdcdd943164c89f4b8c94c0a97cdc289787265c02447d957b706"></a>

## Direct properties — site_virtual_sites.advertise_where.virtual_site / 6eb5ce37d172 / 3

<a id="canonical-2a60f5302a604100d643dd0a1b6cae6d4a603f278113a9f4a6a13aecf2f2daaf"></a>

<a id="canonical-33cb2c122f8ec030f86b3114eac5bc78c19021dd6292f2e2db34fe8f38663874"></a>

## network property — site_virtual_sites.advertise_where.virtual_site / 6eb5ce37d172 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--proxy--reference--group-004.md#canonical-723fd08e6fbb3eea8f0ed638ae841ec73923247e29fe06d15b8497e34aca1a11): complete subsection reference.

<a id="canonical-75d927d0e4af061dd8fb295798fde532c8b86656566212c5c49877644c040644"></a>

## Next pages — site_virtual_sites.advertise_where.virtual_site / 6eb5ce37d172 / 5

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--reference--group-004.md#canonical-723fd08e6fbb3eea8f0ed638ae841ec73923247e29fe06d15b8497e34aca1a11)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-723fd08e6fbb3eea8f0ed638ae841ec73923247e29fe06d15b8497e34aca1a11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b861ceb7f722ba82c37bc155e4e2aa2cd2af0442d0831e5b9058fd3b552efdc8"></a>

## site_virtual_sites.advertise_where.virtual_site.virtual_site — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-51d59d36201d2a9d332561cb208e30ed69b3506d2c82f66273c9ba601588b175)
- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-0503ef9cdc1ad966e25baa97d5b1e92b90def581d6994baa8d24180df4d369c1)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-e17c031195aa8d766b737be803a5d552c1fdfbfe55dfaeac8ea88f22e2c7863d"></a>

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

<a id="canonical-b7f29bd8c3d3294de74f058f37c5884dec366095cf27a69ad59c7454ba1fe27b"></a>

## Direct properties — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 3

<a id="canonical-61ccd381416a92fdf80db8d75e635aa63aff6593106ccc0970a9fc16ccb3ef4c"></a>

<a id="canonical-07993f7c8092a4ab2d294ef0a7ef6412a63f71b593cc46b227c76dc225e37e23"></a>

## name property — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 4

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

<a id="canonical-a63e4366f453758acb00762d9f0c4e7a356907f9f9e8d36885f09a1005d2b8da"></a>

<a id="canonical-4d07bd1c8fdeb7d27e182fad743be8785e6e346e1ca335463ade6f498e512bc4"></a>

## namespace property — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 5

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

<a id="canonical-3e269c202ab856d32feed39266d976c96f39aebc3e0e483a7f1ba0aace24c741"></a>

<a id="canonical-4c069f56f7e3b8cdfebab1fb66c51ed4215c58a905faaf39c768759dffcb94e4"></a>

## tenant property — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 6

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

<a id="canonical-d834f0970b6568de6d008032b918bfba3b059cba167eeac6393df3001d4c1d56"></a>

## Next pages — site_virtual_sites.advertise_where.virtual_site.virtual_site / dcd653f476de / 7

- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-0503ef9cdc1ad966e25baa97d5b1e92b90def581d6994baa8d24180df4d369c1)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cde6f15c7a02f8f5520e62cad159e948eaab5394e00fea3098af83b43dab5fb5"></a>

## tls_intercept — tls_intercept / 1c964e9e4f5c / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- tls_intercept

<a id="canonical-99bb678c4909c834195d7cefc40f5c76cdae7ac5abe2820b0fc8a25451848886"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-f63abc940cedf65e84c9129f19b68e6f9f7a38e626cdc0249bc2501197cc7d4c"></a>

## Direct properties — tls_intercept / 1c964e9e4f5c / 3

- [custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28): complete subsection reference.

- [enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-ecb2a177ffba54be67bf4286a0deaed73e261813227e5337de84698479edc76b): complete subsection reference.

- [policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110): complete subsection reference.

<a id="canonical-127c8112a718a6ebf430734723e5cc76a264429ac483ab7669aa621e8f60a855"></a>

<a id="canonical-289c2d02667e12f93482a94348f308b7daa8fb2753d0192f7f06b1a6ba497644"></a>

## trusted_ca_url property — tls_intercept / 1c964e9e4f5c / 4

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-c81186c0316cc9113970ea7b9b30ccea8316c7a3acae7a83d09ea2fc188a7dd9): complete subsection reference.

- [volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-20a058a6b9bcb6d8622f44922a7417163865c4159331dfa258fabb72315c782d): complete subsection reference.

<a id="canonical-050f5f8006be7aa8aac2450117a6241f473494851557c58d91ed1a17ccc9e9fc"></a>

## Next pages — tls_intercept / 1c964e9e4f5c / 5

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [tls_intercept.enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-ecb2a177ffba54be67bf4286a0deaed73e261813227e5337de84698479edc76b)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- [tls_intercept.volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-c81186c0316cc9113970ea7b9b30ccea8316c7a3acae7a83d09ea2fc188a7dd9)
- [tls_intercept.volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-20a058a6b9bcb6d8622f44922a7417163865c4159331dfa258fabb72315c782d)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5ea75ff409c7fbef61296d60dd7c29b6b2539ed477ae0e8316e2069e8bc55d5"></a>

## tls_intercept.custom_certificate — tls_intercept.custom_certificate / a1f7cb5f635c / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- tls_intercept.custom_certificate

<a id="canonical-ee865bcdf218dd5bc12859c7cc11ed1a461708d498845c81b6fc2a6e87e79fb2"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

<a id="canonical-7d8bdddcc0b1bd9225ee53ebcd1143a5ebba8d5bc364449d0aa56fff1be1c7cf"></a>

## Direct properties — tls_intercept.custom_certificate / a1f7cb5f635c / 3

<a id="canonical-ce5c1057088b0ffa89916b33330d2e03323f0266384f98702efc6eecf4f083a4"></a>

<a id="canonical-8d5e8b8a67a1223aeeeeb96b72b24782ddc9fb15b2831a3e5eeb913f6f3b130a"></a>

## certificate_url property — tls_intercept.custom_certificate / a1f7cb5f635c / 4

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

- [custom_hash_algorithms](data-sources--proxy--reference--group-004.md#canonical-4b9ccfa1b640f113448d3e5b150533101efd432cf7816a1762c830ef49d39bb4): complete subsection reference.

<a id="canonical-88a01b3474ff2c4e16e277479dd5d0636f32168e2936ed51bdfbe1098a3d3f51"></a>

<a id="canonical-f42c65cfc2642d6d076d5a4c35953c67b6def2ffad645aca1668085cb7fe6644"></a>

## description_spec property — tls_intercept.custom_certificate / a1f7cb5f635c / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--reference--group-004.md#canonical-4cbededae3ed8863b04a9ab4d04a8b014f693530148277b8df84ff9f4baa4cdc): complete subsection reference.

- [private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7): complete subsection reference.

- [use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-72ba0999c99d0ccb921817e6c9637334daa45195b2895807ee7fcd77e78d66ae): complete subsection reference.

<a id="canonical-7ea31e666c095ce1e0649130f1243e280d4870151014576e5de3834ccf2329d5"></a>

## Next pages — tls_intercept.custom_certificate / a1f7cb5f635c / 6

- [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--reference--group-004.md#canonical-4b9ccfa1b640f113448d3e5b150533101efd432cf7816a1762c830ef49d39bb4)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--reference--group-004.md#canonical-4cbededae3ed8863b04a9ab4d04a8b014f693530148277b8df84ff9f4baa4cdc)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7)
- [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-72ba0999c99d0ccb921817e6c9637334daa45195b2895807ee7fcd77e78d66ae)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4b9ccfa1b640f113448d3e5b150533101efd432cf7816a1762c830ef49d39bb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72302527f36bc08069e6eb36d5baed77d0e9ed00bf778d67c301a291eb33d955"></a>

## tls_intercept.custom_certificate.custom_hash_algorithms — tls_intercept.custom_certificate.custom_hash_algorithms / 14e1ef38b464 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-1cc1628f0cd934728afb3199ab654f3c46a3beade4baa8572bc873dc1ace6fc9"></a>

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

<a id="canonical-c6592fbc7127752d7ba46ac12fe7f8126d72dc9cae69816c8fec2e0616e88d9b"></a>

## Direct properties — tls_intercept.custom_certificate.custom_hash_algorithms / 14e1ef38b464 / 3

<a id="canonical-e0ef1598d1cd8415899dea4be66554ef4a1205fb9c87c4f5a212f8468f535c95"></a>

<a id="canonical-eb998842e26c4e48f0fa501dd71f7e1fb64a0f677047e261f39f7b88cd82979c"></a>

## hash_algorithms property — tls_intercept.custom_certificate.custom_hash_algorithms / 14e1ef38b464 / 4

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

<a id="canonical-b90a2005a6abda6b2e5fe64290c5bd9fc7be0a4ac7bd8f4bbd8779b6955468b1"></a>

## Next pages — tls_intercept.custom_certificate.custom_hash_algorithms / 14e1ef38b464 / 5

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4cbededae3ed8863b04a9ab4d04a8b014f693530148277b8df84ff9f4baa4cdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ead29c685251b7382beb6161dc4c1395f1cfe5b9c1692779a9abeab22a9ca627"></a>

## tls_intercept.custom_certificate.disable_ocsp_stapling — tls_intercept.custom_certificate.disable_ocsp_stapling / 113b3e8bb2c2 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-e1ebbdb2d1d9dda88a86fd700968497de1b3fe1369c448835014e64ce977e661"></a>

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

<a id="canonical-79df916d64226159f841c3d3c120d00cf59445a9e5c95a030f32394883365d4e"></a>

## Direct properties — tls_intercept.custom_certificate.disable_ocsp_stapling / 113b3e8bb2c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d0183ceb21cb5d94de60a5a85ed027e71962a80abe61474e30ab8f94eb4c2ca"></a>

## Next pages — tls_intercept.custom_certificate.disable_ocsp_stapling / 113b3e8bb2c2 / 4

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abafbb704a9cb805f51d7c75e057649fbd565e2d42de13b23206c956eabba290"></a>

## tls_intercept.custom_certificate.private_key — tls_intercept.custom_certificate.private_key / 1a3e99cf0cba / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- tls_intercept.custom_certificate.private_key

<a id="canonical-7951ef69fb554e3ce16cab4e48e30838205179d360a4659305092e9bb5d7615e"></a>

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
