---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-b44a5f0e013d4eb6d4920f0f584203f6ac0f4c183e3290e0cb0290edd0ea436b"></a>

## Next pages — captcha_challenge / 4bbc8cbfcf5d / 6

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7271caa7f1b7bb7c445b90437bded4399c011685b1462a0bad37a5770d5ab076"></a>

## coalescing_options — coalescing_options / f9ab8a71023b / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- coalescing_options

<a id="canonical-c32e9c2c91a02111b18820cf076c03b93d03a47dbb3f60320151bbfd9fa07fe1"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-6ff7280f4d67a69ed5160c120b15f247c5890bfaee74ebc573d36561a1106519"></a>

## Direct properties — coalescing_options / f9ab8a71023b / 3

- [default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-048fc51421dbbf93ae1fea03e94c403679fa7f70d4eeb3dc728211718f6b8388): complete subsection reference.

- [strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-4fac4a25c0252217c336ab079e0700b60c76bf41255efd64454963a23bc5a17b): complete subsection reference.

<a id="canonical-bc82de1a04963dd60cdd46674f9221bcf5d00bf6d1e9cd6d73eb9343cb3e8339"></a>

## Next pages — coalescing_options / f9ab8a71023b / 4

- [coalescing_options.default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-048fc51421dbbf93ae1fea03e94c403679fa7f70d4eeb3dc728211718f6b8388)
- [coalescing_options.strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-4fac4a25c0252217c336ab079e0700b60c76bf41255efd64454963a23bc5a17b)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-048fc51421dbbf93ae1fea03e94c403679fa7f70d4eeb3dc728211718f6b8388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8f89a8c54b8bf483566f56d2219681243e746b7b5620a3fb7ebcdd382267737"></a>

## coalescing_options.default_coalescing — coalescing_options.default_coalescing / e631ed4bd9dc / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705)
- coalescing_options.default_coalescing

<a id="canonical-3e141c92a8b3542f5a934497060e5c156eaade44f25b9aa40b5917412391d941"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-a70204448a9322abb210f1a8934f2d98269150e3b8a8a5f5134d319cd6e2df88"></a>

## Direct properties — coalescing_options.default_coalescing / e631ed4bd9dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d97046407a2921d9a06af7c658d826094c70a5d05232a9594de9d98b956ce53"></a>

## Next pages — coalescing_options.default_coalescing / e631ed4bd9dc / 4

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4fac4a25c0252217c336ab079e0700b60c76bf41255efd64454963a23bc5a17b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b8c99654e505addd010fd311644c5e08cf46f2a8688eccabc5243ee4c4833e5"></a>

## coalescing_options.strict_coalescing — coalescing_options.strict_coalescing / e56d9669c508 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705)
- coalescing_options.strict_coalescing

<a id="canonical-f93a8970153609fab7ae74d14e444c62c2a215e08b02f13b3247f25e29f30bee"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-f1dd42c1c4515aaf2e1893054d41146abac9fea55d1ac45a84cac51c63e6d49d"></a>

## Direct properties — coalescing_options.strict_coalescing / e56d9669c508 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6f87c1a92fd25569d7264993b95eb51bdaff47f638c81d85010e3754296e76d"></a>

## Next pages — coalescing_options.strict_coalescing / e56d9669c508 / 4

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-175e663c8201caf980e21bdc50984a06a8b563d7d07c88dfceb91ced493567cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac5dc2024732ca4eec08c850e6147b99d26c9592552365ee3e15770af2de30a1"></a>

## compression_params — compression_params / 75c87527cf84 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- compression_params

<a id="canonical-371420dea8f387a112f8045b268676bd8537814a14e5c594eab27903140dce19"></a>

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

<a id="canonical-6219ca084694f36c08e6727e1475ce76c4615e663be9c879faa965bd44de20c7"></a>

## Direct properties — compression_params / 75c87527cf84 / 3

<a id="canonical-6290ae6c6709f11f19438b38fd66d40436b1fb2937cbf04c3dfbf2a9ea373223"></a>

<a id="canonical-42bf82c080828128833eff45fac857a64d46728941a62e01133a6e1d2efd20ba"></a>

## content_length property — compression_params / 75c87527cf84 / 4

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

<a id="canonical-d977c3567a3a070c269757e513b5ef7c90a7d747129ab0b47bb62f1b208a3e11"></a>

<a id="canonical-1eaa17eb0844c23f7afb816f8b2d1461050a83119c3d930b9f3bb77f57a46584"></a>

## content_type property — compression_params / 75c87527cf84 / 5

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

<a id="canonical-8d823396df1bc90ad06f0272c574aae4e320929beba249e756b5e16967f61f94"></a>

<a id="canonical-ca7f0541871ef34f12448b57ee314911c607e05f4f0ee887babe9ce68572d2cb"></a>

## disable_on_etag_header property — compression_params / 75c87527cf84 / 6

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

<a id="canonical-7842ee8f6db8a9c5d3337853b35502a1bf9fa5a401bebfe132db88bad7801677"></a>

<a id="canonical-de58ece1491c0e0e7d795a4d91386556eee3e2574d1048dbaeee75a4368fbfa6"></a>

## remove_accept_encoding_header property — compression_params / 75c87527cf84 / 7

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

<a id="canonical-4f8e26652b8eb68fd2c8ac3e0083d1f232f2a9c2221ce58d58d58b525aaf6184"></a>

## Next pages — compression_params / 75c87527cf84 / 8

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-a4ecd779378b174be7fa310c13ecda6cf2402208e2a96236ae26a14b61393c5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45fc1da0b414b8eaa674f4697da0d87b0168a04b168c725beb6bdaf598fc1970"></a>

## cors_policy — cors_policy / 8ed924f6fcd0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- cors_policy

<a id="canonical-8029f76942051a56649d4af7cfb097baa2d2b95163103c177bf9ac7bdcabc600"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8ed8207c04197b0a33b26ad27d2da230d8ee30c3f61518ca213c46fca4dad25e"></a>

## Direct properties — cors_policy / 8ed924f6fcd0 / 3

<a id="canonical-84829179233191f2cb737ea0aeb5f0c87ec1be9a5f6f913b93011f1f735f7e8e"></a>

<a id="canonical-3c64955738884d1861532ca2e623b0fed0aab733550b64b11276db5f2ca50afe"></a>

## allow_credentials property — cors_policy / 8ed924f6fcd0 / 4

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ff53a860e055ee96273bae6bb4c34cc5cb8d29f141e34129ca0aea3ecca00dd4"></a>

<a id="canonical-a12a8ddf04999a42b564d0e226029d0b6d6201213ac2e5afa92ac4c77a0215b0"></a>

## allow_headers property — cors_policy / 8ed924f6fcd0 / 5

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-c50f9fbf03cc3fe20f50d762d93d29fd19a737f48cb5f7b28ead3ecc4552c43b"></a>

<a id="canonical-2720a837ef19338260be44ab7a3c109a3707d21e2a9de51631bdce95a2ff9629"></a>

## allow_methods property — cors_policy / 8ed924f6fcd0 / 6

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-4bb5aea485eb7effdff0031033a673dbc0262b899a4e4e49b8011b3a05111834"></a>

<a id="canonical-41690f612be5cb41d8e06d58a96e42fb200b201073ffd1f37da33c00ee0b69fd"></a>

## allow_origin property — cors_policy / 8ed924f6fcd0 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4cb72b296e1985049897043ce56a3d24cabb822e6739537cd7c2c2aa05b0654e"></a>

<a id="canonical-ead107381a5886d608bb32cf70844ac7e9b5c9c154b7730a095309b1ed76477c"></a>

## allow_origin_regex property — cors_policy / 8ed924f6fcd0 / 8

Type: `["list", "string"]`. Computed.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e4797bd6b552dfa313b4d5eb82d6f3614f420cfa148a0bab291de8026a8a9415"></a>

<a id="canonical-5d0c9611605e3d9d7374633b005000180f97efa305c4c51d8cb46de7fa09f15b"></a>

## disabled property — cors_policy / 8ed924f6fcd0 / 9

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bc1a8a5df6d1f0a3ec1f313d07c52ed36e25f2754d563fa1e412df3a52975e54"></a>

<a id="canonical-ba1f8cddbbd43e2a5a678c9f936355fe039d3c4bb401bf13ef37a5a3c4dd6e91"></a>

## expose_headers property — cors_policy / 8ed924f6fcd0 / 10

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-fa2a338e7cd18a2e1dfc2979f53a7626d2ef38b97a5e57f48485b1e2f78c9998"></a>

<a id="canonical-09151968e19039d0aa4aa33aaa2d3575cfcc923fa706de9c3aab2a1c90d1e344"></a>

## maximum_age property — cors_policy / 8ed924f6fcd0 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-c9f3f6c32f900e46176b7642f15b1ee099912e027d8efb400394034c72040f5e"></a>

## Next pages — cors_policy / 8ed924f6fcd0 / 12

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20c5d706cc295cdcc1af44a9f8a5debeb9fd870fc21bee64350172788dfd0b77"></a>

## csrf_policy — csrf_policy / 895ceb6f9198 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- csrf_policy

<a id="canonical-cdadf81518d323a24c95b4a540fcc368734f947f5e5277f695854487b8b7b400"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-e56498341ab4674d9de68334e8309a3e7017f798cfd8511a6a095a8880449fbc"></a>

## Direct properties — csrf_policy / 895ceb6f9198 / 3

- [all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-eb74a18bb4fe665817f1bec0ae41acb9a2d0940ae9b8fdffd978d306d6a2f184): complete subsection reference.

- [custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-cc6c25ab1e1e8494959d6aba2c79ff0e618d7a7c15dfd5b25954f45826b99ab3): complete subsection reference.

- [disabled](data-sources--virtual_host--reference--group-002.md#canonical-808e20cabaae3c21d26c3fba2a36e4d48ff4f3ac0b4f916ec63e3a44c6eb0c58): complete subsection reference.

<a id="canonical-01943169fae44f6eaaafe26ea08f005b0bd7fab0d29953b5ebdc84282cdc7027"></a>

## Next pages — csrf_policy / 895ceb6f9198 / 4

- [csrf_policy.all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-eb74a18bb4fe665817f1bec0ae41acb9a2d0940ae9b8fdffd978d306d6a2f184)
- [csrf_policy.custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-cc6c25ab1e1e8494959d6aba2c79ff0e618d7a7c15dfd5b25954f45826b99ab3)
- [csrf_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-808e20cabaae3c21d26c3fba2a36e4d48ff4f3ac0b4f916ec63e3a44c6eb0c58)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-eb74a18bb4fe665817f1bec0ae41acb9a2d0940ae9b8fdffd978d306d6a2f184"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-190a8417d085cd5190fb90fd8f085f1c93fcecc691aedd5e3a8fa48580ffc1b5"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / ecd21618bbe0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- csrf_policy.all_load_balancer_domains

<a id="canonical-5ac40d12ac1fe90bcc33d902c65d860f2fc9abffd31636611b3fa5bb80f57249"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

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

<a id="canonical-d2618da9fff76ee9372b0463b5bdc64c377918b809679175583609dbe53e9c16"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / ecd21618bbe0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47f50f0ad712d01d6cfc8c395ef220b7cdd38be4afb32770b14795557b47694e"></a>

## Next pages — csrf_policy.all_load_balancer_domains / ecd21618bbe0 / 4

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-cc6c25ab1e1e8494959d6aba2c79ff0e618d7a7c15dfd5b25954f45826b99ab3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad49b8ae92863e361a08070294fd486c588e7d3d6a7c7c2dcbd67e39a630aec3"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / 97b0457a15f0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- csrf_policy.custom_domain_list

<a id="canonical-778a7081ef1667154c4681cb407fd72b0276d3ec3293af36f7baca1494ec9b19"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d7eddc4c6cae9e4adf92ddf070901385857cd679cb6224d403cd1fb89157bc7a"></a>

## Direct properties — csrf_policy.custom_domain_list / 97b0457a15f0 / 3

<a id="canonical-17c09f70e87d350f220aef66cca398ca17a6a17d797ec069d5560c8c36c58a9e"></a>

<a id="canonical-ec2e41aabb82e0a6273f1bf4af5822895babb68f2b7c692decab8207d10a2ff0"></a>

## domains property — csrf_policy.custom_domain_list / 97b0457a15f0 / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-79e1f1dde973eb31112541a44b0b8c38423117ccd57172d97ee773227f8147c8"></a>

## Next pages — csrf_policy.custom_domain_list / 97b0457a15f0 / 5

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-808e20cabaae3c21d26c3fba2a36e4d48ff4f3ac0b4f916ec63e3a44c6eb0c58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9696cbd87b9259dba07910b4f40c3773d68c31682a298da963f20cfdf6684e43"></a>

## csrf_policy.disabled — csrf_policy.disabled / 04b308d172a3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- csrf_policy.disabled

<a id="canonical-3b3eb6a71fc1dbaba2504303bef46edf3f7dbb756b68df9ecac48b74387c31fe"></a>

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

<a id="canonical-648b8a70d8233c101eb66ba60e114310dee5220ed28dbf3bdc18c94e56a4634f"></a>

## Direct properties — csrf_policy.disabled / 04b308d172a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ada641010ab8b7a3f24e440ce95ba565e7edc6013cfa4f371dd1c7f57ebac33c"></a>

## Next pages — csrf_policy.disabled / 04b308d172a3 / 4

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-2ba5db0f8ddd27b58c29095c0986c994d4d2c03100a06100429e44c738a20fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eae2fedfe9f5574236fcda53b5d99d2563d487674a738e72f20536b8f21a7209"></a>

## default_header — default_header / c3bb3353a51f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- default_header

<a id="canonical-1db43c5339da33f1c2a8f4536a86d2286e9309f41d9b07058244c35b01325b12"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-b89620b6ef92598f989a4c9b788c6609728f292db31164d836915bd6bf63242a"></a>

## Direct properties — default_header / c3bb3353a51f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6ee5194ef0f0056e0cfa98049eaaa8a9647698f8e6eaad3c01016bb6f849034"></a>

## Next pages — default_header / c3bb3353a51f / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-be86a1bd888037fdc3242c0d59516ab12221888e99a192b774b252f4409bcb2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcaea92fe90aade79f7b4524924a7d00e941623c153e16738da222cffb78be09"></a>

## default_loadbalancer — default_loadbalancer / 30209d056ac9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- default_loadbalancer

<a id="canonical-2393bca9603456e5118af941fbb8a44800e59cf3bfa0ba99f0e7fed04e27f438"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_loadbalancer, non\_default\_loadbalancer; Default: default\_loadbalancer\]
Configuration parameter for default loadbalancer.

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

- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2393bca9603456e5118af941fbb8a44800e59cf3bfa0ba99f0e7fed04e27f438)
- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-4c2f3a3c6f3ec034edc006c8cbff7413f0ac22898e8d8ea3779a2d86f956d9c3)

Select alternatives according to the provider validators above.

<a id="canonical-011f56aa3df83a362a7b9763c471cc4c4d156b7de0fc220c2d48e22a5df704cf"></a>

## Direct properties — default_loadbalancer / 30209d056ac9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df1eb554a9c6bcd9ac8a40714f67e366cf2b5cdd2fe77f59508aa80a6e345d5e"></a>

## Next pages — default_loadbalancer / 30209d056ac9 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-fbdf7fd14ab032fa59e6f0edf908260deddc86d23ea5ac6cb183d78a9e97ca96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffd894f0c96f4d10359d0b7f0c2749af212839650e04ac465ae630559d1b00e2"></a>

## disable_path_normalize — disable_path_normalize / 5cd819958434 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- disable_path_normalize

<a id="canonical-f1495d990daaf02cc2afa0a7d57e14e4e49d847696ade3fe5b735d2a24114f2c"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_path\_normalize, enable\_path\_normalize; Default: disable\_path\_normalize\]
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

OneOf alternatives in this subsection:

- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-f1495d990daaf02cc2afa0a7d57e14e4e49d847696ade3fe5b735d2a24114f2c)
- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-212b027873cc3a831d39cc1391c20838522a550bc2066de82e60509f37e8dd36)

Select alternatives according to the provider validators above.

<a id="canonical-6abee2fff0f9111831d0901bb64790ac3005c3f3fecb08acd5a8f4bc2c797548"></a>

## Direct properties — disable_path_normalize / 5cd819958434 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2326cd8b0827438e4c6a41507bf8b8e297b129129306fd1db611c4c8c9cd1f18"></a>

## Next pages — disable_path_normalize / 5cd819958434 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-b8e8c7fcb7321f9bdf837de330392a7f4162a7567cf7f68cb61c12c9040bbf3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79da8405b4dbf01b8c64346a339c83eac5af4d6a4c33153c8e450bd74c25a660"></a>

## dynamic_reverse_proxy — dynamic_reverse_proxy / 60a946224b54 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- dynamic_reverse_proxy

<a id="canonical-f6043a214957e33bb842cdc42d48bdeced74065fac3db0a95d4b8b7525ac0d2e"></a>

Type: `"single"`. Computed.

In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic
resolution is done using a predefined field in the request. This predefined field depends on the
ProxyType configured on the Virtual Host.

Upstream description:

In this mode of proxy, virtual host will resolve the destination endpoint dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cb368df40ebaa89720c6dfce5cf14a3d9b74628610f60ca90a56141f479ff891"></a>

## Direct properties — dynamic_reverse_proxy / 60a946224b54 / 3

<a id="canonical-e5a1e7b3d67d861fe2d459e042014203b8b0925e74099b4b5e7858883db3e19c"></a>

<a id="canonical-6c20df35134d46de7d7765f662ff62fae2a37716ef9d58ead200b68b7cb03e90"></a>

## connection_timeout property — dynamic_reverse_proxy / 60a946224b54 / 4

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

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

- [resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-f0c2c0c57fd92f1734eb1583703b944b947b46b14f013f468ffdba7facd4d5db): complete subsection reference.

<a id="canonical-c47035729c92719f9f6b7f688971663d09b481dbd9155844a3157c983fe93717"></a>

<a id="canonical-d9917964190c43c08a72c74c83c9c6ec5ca1d77a6f31c3eb966ebe54cdd9372e"></a>

## resolution_network_type property — dynamic_reverse_proxy / 60a946224b54 / 5

Type: `"string"`. Computed.

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

<a id="canonical-950fdfc61208072977d6af5de3fdbaf966ce149ba87804eadc5f3a4e60882029"></a>

<a id="canonical-bc47563b09b2c57b10adce875ce5c55d8067e6357a3d407ce93d85078214d748"></a>

## resolve_endpoint_dynamically property — dynamic_reverse_proxy / 60a946224b54 / 6

Type: `"bool"`. Computed.

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically. The dynamic resolution is done using a predefined field in the request. This predefined
field depends on the ProxyType configured on the Virtual Host.

Upstream description:

X-example : true In this mode of proxy, virtual host will resolve the destination endpoint
dynamically.

The dynamic resolution is done using a predefined field in the request. This predefined field
depends on the ProxyType configured on the Virtual Host.

For HTTP traffic, i.e. With ProxyType as HTTP\_PROXY or HTTPS\_PROXY, virtual host will use the
"HOST" HTTP header from the request and perform DNS resolution to select destination endpoint.

For TCP traffic with SNI, (If the ProxyType is TCP\_PROXY\_WITH\_SNI), virtual host will perform DNS
resolution using the SNI.

The DNS resolution is performed in the virtual network specified in outside\_network\_type or
outside\_network

In both modes of operation(either using Host header or SNI), the DNS resolution could return
multiple addresses. First IPv4 address from such returned list is used as endpoint for the request.
The DNS response is cached for 60s by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0efca7ffcf92344132099d2683d6dd29c0a95ceaec4ad5810699cca0867b28c9"></a>

## Next pages — dynamic_reverse_proxy / 60a946224b54 / 7

- [dynamic_reverse_proxy.resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-f0c2c0c57fd92f1734eb1583703b944b947b46b14f013f468ffdba7facd4d5db)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f0c2c0c57fd92f1734eb1583703b944b947b46b14f013f468ffdba7facd4d5db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa46644ff8952ee23551b6369922e28d7743b262f669e13dd388147ce9cdcbb3"></a>

## dynamic_reverse_proxy.resolution_network — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-b8e8c7fcb7321f9bdf837de330392a7f4162a7567cf7f68cb61c12c9040bbf3e)
- dynamic_reverse_proxy.resolution_network

<a id="canonical-cf1d00a7a61beb8a113ba4d1e32f7bb283d45cd78545778b536153cd25931693"></a>

Type: `"list"`. Computed.

Reference to virtual network where the endpoint is resolved. Reference is valid only when the
network type is VIRTUAL\_NETWORK\_PER\_SITE or VIRTUAL\_NETWORK\_GLOBAL. It is ignored for all other
network types.

Upstream description:

Reference to virtual network where the endpoint is resolved. Reference is valid only when the
network type is VIRTUAL\_NETWORK\_PER\_SITE or VIRTUAL\_NETWORK\_GLOBAL. It is ignored for all other
network types.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f08d3ec7099f2e40940f7b1ede7d5f6ffdaa9c282c76cf6a4067cd0d153c7459"></a>

## Direct properties — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 3

<a id="canonical-f07149fa8f2d9cdd3a298e448d433bdc0f751f56d0f725911b39005d53a4f784"></a>

<a id="canonical-17c964cc5f1aeb924403ee9a612cdec9fc024c26a0495446f952e61c15fbff9e"></a>

## kind property — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 4

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

<a id="canonical-3ec1ac8ea14ff42ee1e21e69cf21b5b0129c5e2188f5cedcbd040fbef5b2f3a9"></a>

<a id="canonical-61ae629d771b778ccfe81cff6081ecd43112ed47c988692112f53f5e54c0145d"></a>

## name property — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 5

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

<a id="canonical-cc347ae2de0c9ba2ce2089216708fa54aa46aa705a0aab508c052a82d25e998c"></a>

<a id="canonical-3ccbdead157fe51fc6dcdab7754171883854cb134b8a06b8b6e2b3a66a3dabca"></a>

## namespace property — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 6

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

<a id="canonical-a967b1bf8cd5e74019bab82995c55b7862ee2e0dc55b165079c7d4a0b3df56f3"></a>

<a id="canonical-cb89fd70dc339c451c5bdc12695a7a5023a741263893ff2a3d417b7cb66ff880"></a>

## tenant property — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 7

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

<a id="canonical-5e2dd412116bff9bb84497ecac6c8abc4aa9ddfec2054a6ad5f1af68da403e58"></a>

<a id="canonical-6898dc874cb12d7fad468adcc302633392456f32bd5f4c507fabfba0835fb37e"></a>

## uid property — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 8

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

<a id="canonical-4f3ff0b19cf6cdd711fcb99833286b964dbc2b29b41bcddf37afee0d293bfa65"></a>

## Next pages — dynamic_reverse_proxy.resolution_network / c28266e029f8 / 9

- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-b8e8c7fcb7321f9bdf837de330392a7f4162a7567cf7f68cb61c12c9040bbf3e)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-1171379e2864a9edbff9631c19170289cd393aea92e23cbd8fbf34b9e78ce7e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c22c293a5c0523d21ae7443ff42422b247e75798d136104175f2d512826caeed"></a>

## enable_path_normalize — enable_path_normalize / f1f9181f1220 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- enable_path_normalize

<a id="canonical-212b027873cc3a831d39cc1391c20838522a550bc2066de82e60509f37e8dd36"></a>

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

<a id="canonical-5bc15fd706689125a3a64b708d15c965392ed052af342362be332e165a2cf24b"></a>

## Direct properties — enable_path_normalize / f1f9181f1220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa3f0338a1769a2b3c42727ab8a57a07c05be9444fac19fc982015d3588f96c7"></a>

## Next pages — enable_path_normalize / f1f9181f1220 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa081e08ee24e661a37cf513733a49a0c70411c459b9620256bc2786d13133b6"></a>

## http_protocol_options — http_protocol_options / 07af33c02e60 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- http_protocol_options

<a id="canonical-b45f4704e2098f4f35efec5d0287a0adbcf1b6ca5844e04cb7e65a575ff1f297"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-88b0860debd39ac248fa227427a71e449f60889beea9a80c2c80bb1fdf26485f"></a>

## Direct properties — http_protocol_options / 07af33c02e60 / 3

- [http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-e4615db49948d7f7dc2b1e8986626f9b1717d65ef69de767979a2ed9c5f55d7f): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-4ed09f400e59cb797d729ddf580fbf8d0e82aa1027b500586c5ba49a39784795): complete subsection reference.

<a id="canonical-8b0d00cb0ffd4f1da7c0227a6e243b0f00703c54c2090135abbb2474ebf7424a"></a>

## Next pages — http_protocol_options / 07af33c02e60 / 4

- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- [http_protocol_options.http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-e4615db49948d7f7dc2b1e8986626f9b1717d65ef69de767979a2ed9c5f55d7f)
- [http_protocol_options.http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-4ed09f400e59cb797d729ddf580fbf8d0e82aa1027b500586c5ba49a39784795)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe8c0a8f15e8292f00476aa72731f78fe5d2298b9676f1ed6ced4a7bf012b43f"></a>

## http_protocol_options.http_protocol_enable_v1_only — http_protocol_options.http_protocol_enable_v1_only / bb3128490d40 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-9833a9ba874394cd9e5a7523b486c7b43695475ffa606ec8c3e547d734d06aa1"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-60ebd700745393fa500486b6a8e7bde7f2431fa51e508f878ffeb30375055847"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only / bb3128490d40 / 3

- [header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0): complete subsection reference.

<a id="canonical-94cdc8edf7ba459d75bc5fffc586db628834fec641708cda66bca5a4409051ef"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only / bb3128490d40 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9435fce6b98d6052ba9b077834392a211ad246ea36781ff95c0bf1d5ddf8ef27"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 991fc9a00707 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-18fadd1917758db465a449070d23ff247f1c36cb5795e9aa020f738bb4d29fed"></a>

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

<a id="canonical-ced0aa96ba44bf02357d1cc81824d328c53b5c97d6f39547bd619e7da1053495"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 991fc9a00707 / 3

- [default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-41cb8bc5c2a61ce033cff92e51a9533a0e88987bed6978e114868474c55d8e70): complete subsection reference.

- [preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1c0d5cb862a0d57860cae9fe1ef095f640e996c12e909be00935f87aef1069c8): complete subsection reference.

- [proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-607e7d05879303692c6bfa4eab7c422d7b9a747ac47ce76bb585243887d0be06): complete subsection reference.

<a id="canonical-8f0dff0984b95bc9b12d638a5deb5e232c0745d2de97e50b90f6e1cfb7e5d4c0"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 991fc9a00707 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-41cb8bc5c2a61ce033cff92e51a9533a0e88987bed6978e114868474c55d8e70)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1c0d5cb862a0d57860cae9fe1ef095f640e996c12e909be00935f87aef1069c8)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-607e7d05879303692c6bfa4eab7c422d7b9a747ac47ce76bb585243887d0be06)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-41cb8bc5c2a61ce033cff92e51a9533a0e88987bed6978e114868474c55d8e70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae2bc2084b9835823c1988e3b806699850e510a7f0ef962e99c210f30e077a8f"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / b243001d4db4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0c6687e127d595e890513e3869fd455127341febe2cfffb2c49d783d250f5773"></a>

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

<a id="canonical-eefa9338ebf5bba00a70eeb82fc887bd9e6f974ab53a820f8a06edee8f39cc8f"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / b243001d4db4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c6d6ec712aefb9086d3b6f31616da4e0df00d0474715ba6ae708d80b27fec9e"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / b243001d4db4 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-1c0d5cb862a0d57860cae9fe1ef095f640e996c12e909be00935f87aef1069c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af6f3416044a40b9cca45c49bb147410a7ae39c4f164b7515b9844b10bbc2c4d"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / b0dd290c94cc / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-aff6b9d1cb32a877251a31259db00b624d56832f65fadfd9446c1d766ad9190b"></a>

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

<a id="canonical-e94a2c8439f191f43fb2d2045478f3c10d41354f6dad2a15e82456fbd1709c9f"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / b0dd290c94cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c361a82cab1a681cc505dfb390449e9dc0a5565d81ccc35001d75ff06a8007b0"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / b0dd290c94cc / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-607e7d05879303692c6bfa4eab7c422d7b9a747ac47ce76bb585243887d0be06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bec2af4cb47f1621c8927e226ebe437d144989a6611b158761be0ead8a50b29c"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / e34b10636027 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9c9834a7469389ea3bddb844767983bb98ff7f6d59075a6b38d03830ca24b663)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-6c82bf72820a713ac74edbccd8cea0198b52713a73de406a11fb5dddf2c6061b"></a>

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

<a id="canonical-47f55800df92148f1486d7046aa842bbdd660d39a9854d8ecda241ac249dec1a"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / e34b10636027 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d77fd610927000ea8cbf9e83779732805190ec9611e804833fe1eb46f75ff0b"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / e34b10636027 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-1ced60db1e4c104ff5541537f5222408725a15d7403e49e41c562c4e695171e0)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-e4615db49948d7f7dc2b1e8986626f9b1717d65ef69de767979a2ed9c5f55d7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-684e90b4538155070fa6038e70a1d1f1deee176cdcfde707a3ed7b7abe246db8"></a>

## http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_options.http_protocol_enable_v1_v2 / f4abe590d9dd / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-f3864f1653d1cbfa6f9f59ddef8d0f91c7b05dc135119d8066c8cee45587be64"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-2d4cf47b4dc7449a5a2476cb9aaef1d891c28918e2423fc0b0cad3d68b6b0255"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_v2 / f4abe590d9dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df1a6ec42886e84460a1225a2bbc9b32d2d0e89fb33181105971bbc6a5af6c18"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_v2 / f4abe590d9dd / 4

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4ed09f400e59cb797d729ddf580fbf8d0e82aa1027b500586c5ba49a39784795"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88465396276a1c6e8cf3dc8e1c3d346145f6ba6931c572dd6c72252d05b98bff"></a>

## http_protocol_options.http_protocol_enable_v2_only — http_protocol_options.http_protocol_enable_v2_only / db29c0425afd / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-a540c67759efc43fe558b6900c86cf50d3322032559d16284126f2258ae6a92b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-43c016290ef34b93e795933146a2539ea1d33f7762f27b5c4b1de901ed35fd47"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v2_only / db29c0425afd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64d416a3c3157071e0f3bb5fffad7a580261ae268a2e3b53937d381443287cb5"></a>

## Next pages — http_protocol_options.http_protocol_enable_v2_only / db29c0425afd / 4

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-35d4e92967faff2a93e6f02288638f65dd0c817f65630593c95b78f48dc76894"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8200c9ac67d1f96169542f49efc382740827b2e6b75970bc18d1243905c9e23d"></a>

## js_challenge — js_challenge / 48555fa5a117 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- js_challenge

<a id="canonical-6fead0d412383114980e7657466e05e6825f8ea21f77f0cfff01ac7f11f7cf40"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c61bb215311ad6d3f1fdeb36421509444ba6c554b7acb0beae5610c8013a34ad"></a>

## Direct properties — js_challenge / 48555fa5a117 / 3

<a id="canonical-b75ff2bc483ebbc39c86f7c34209b74f2722de6a8f099e345f6e5b3b45c4b421"></a>

<a id="canonical-33f08854886d78b36ab77e47e6a43d085dae38f229bf67c7efa1c89e037a327e"></a>

## cookie_expiry property — js_challenge / 48555fa5a117 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-72b10fbea2de566816df88f765778ed2f47b9a6374dc48679741c29304d5eee5"></a>

<a id="canonical-b81b97bae8ec501fa79dd63d9e94ee36dbe4f19ef6d47b6127a21cc186ded4dd"></a>

## custom_page property — js_challenge / 48555fa5a117 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-13628b4d5e0b945a16120939f3d9f9e772006c9b0439ff155a3b3cb356efabac"></a>

<a id="canonical-50fac4a9dcd94423d014a35e8bcd6e771389845c8ae5418e533e23d9a926000f"></a>

## js_script_delay property — js_challenge / 48555fa5a117 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-ae02e8a4838fa27df06a49c777bf283eb42fb3959376ffabc408611ff02f489d"></a>

## Next pages — js_challenge / 48555fa5a117 / 7

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-6fd9adf9858ac044d27f69605c0da0afab5ae1521dd18568d820e2c4370a249b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e99993fec0e8625d21ba4cc029732feccb077490bcdd570638e74ef6263715f"></a>

## no_authentication — no_authentication / 76886c10f6c1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- no_authentication

<a id="canonical-524320b0777e1f50d78a30949d30b94597acbf8a2389b0c8271854f1eef75fd4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-4f92c940c414a3360a317d9c469e216923b4ef6a8f95b5c8764cae08e25b9aa4"></a>

## Direct properties — no_authentication / 76886c10f6c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c0244050864385ff33ae889a6157b11074463317b7cf833b84473f21dd260fc"></a>

## Next pages — no_authentication / 76886c10f6c1 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-d44525cc14415a80f6e329d02112bdf067ba82809fab4500bda9866ea672ad6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c5c6412254800e0ff70328f26ea982022432a026975c05ab67ff3a325ffc61f"></a>

## no_challenge — no_challenge / 1f65a31c164f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- no_challenge

<a id="canonical-3f0bdb5749658baa498a33d910c44024358aa16e77e1a62fe2a2b49bfe09053c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-67019a4d94f7c54fa680d95271e7b820ddde84af0d5d4cc51b0a75ac57c8385e"></a>

## Direct properties — no_challenge / 1f65a31c164f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b51072b75cf38e4f3e0554e363f613d1c60fec873f653f04f2b791c64e76c882"></a>

## Next pages — no_challenge / 1f65a31c164f / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-545aca29b311fea151c0f133136e8abac4d81e741b6298394ac83aad08e117f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d0a768744b3b5860c41ab364cd54d7894d8067b75faa9378cddb55d19868e77"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / f2a5ff002326 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- no_request_limit_per_connection

<a id="canonical-aab4430cdc26268dd04c870fd9a8040e5206acac957535669f5b3b0ec91639e3"></a>

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

<a id="canonical-777bbfad370a53517bf0a3465b6cccf239031818d3c19433c6081365bce9bea7"></a>

## Direct properties — no_request_limit_per_connection / f2a5ff002326 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-146f2dbdc2b53e7576bf2d34169e6d03015c671120f9e497efa518a6c90dee1d"></a>

## Next pages — no_request_limit_per_connection / f2a5ff002326 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-8b5b859bf066db8c50680f61439eb6236af261b0f24ab581d2fa842959be914d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63515cb86ad271ebfe84a5f01ffe1e2166292468d0da4c7b9ad9a5525f3c2afd"></a>

## non_default_loadbalancer — non_default_loadbalancer / fb41eae8af3c / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- non_default_loadbalancer

<a id="canonical-4c2f3a3c6f3ec034edc006c8cbff7413f0ac22898e8d8ea3779a2d86f956d9c3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-c9de42337ac185cc9f486fd9bee4a1ce0ccc0ec6b2aba4f7086f789013cafbcf"></a>

## Direct properties — non_default_loadbalancer / fb41eae8af3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d742bd5d5a31cbe76ccbc492e54dfd2c588991363b5244f8ca42f3b183e90bb"></a>

## Next pages — non_default_loadbalancer / fb41eae8af3c / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-62dbd01f9aad94f490c77cc7b4d61c54a4fe212c25d4796a116058d8f3256475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f80ed0538157b5f6ec26b9cae9b9a289148d77cd22917fe6b67573ee0b940702"></a>

## pass_through — pass_through / 8e10fd46fb17 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- pass_through

<a id="canonical-85b1046ead71ab45112df5f0a65c1bb1a1473365aa0626737f0354f3b3cc8339"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-9e4b1ac6dfd464c0fcce92e250ad3cc83c986e7277f24e5c0a17f1a81b8796cb"></a>

## Direct properties — pass_through / 8e10fd46fb17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6905d2324ab530e2d2a1d6a821a7bbb00ce476dbde754d0caa7624918d25e9ba"></a>

## Next pages — pass_through / 8e10fd46fb17 / 4

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5fd02a3e53a4014bceaf5d9ab95dd66a94fd084b4552ea6670553171736272ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc1f6d892626e3e6a5da8bcab894dfe6976304fdbaef086e4c31a2db67d39d9c"></a>

## rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- rate_limiter_allowed_prefixes

<a id="canonical-3df10ea042418c3f98638797c2a3b88c40e51e055f8603cc7a55497bc592e0d4"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

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

<a id="canonical-2b37ae45c1b206f320603cb57d858d6312ac12104d35caf597033ec7403d3e22"></a>

## Direct properties — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 3

<a id="canonical-c32a11b6486b38446933a6e0cee529eb85881650a40dbb74c17caccbccbd1bcd"></a>

<a id="canonical-c2769792d21c1ab3c87d5a9896903dd2dd57b41c41b2281cc9dafae7935b0b04"></a>

## kind property — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 4

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

<a id="canonical-b031c297824dadaf3e516dd6e2ddbe026cd7a1c24de5502b90a06fe0c32f16f3"></a>

<a id="canonical-4cc2354ab26643d22530f2e3bb2b0bfb1dcaff74157ab8c6dc87aac91dc1cf31"></a>

## name property — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 5

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

<a id="canonical-1bfb76bfafe95aa8ae00f1e222a4b40c38f1b93219ec9ce365095142c488984b"></a>

<a id="canonical-0be2a00310db4a432924d5d8b0854683d195cdff8093e9b977147beb84ee61e9"></a>

## namespace property — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 6

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

<a id="canonical-3e7b9f690462aeb52e5e7ace32ee8a75e3249673b13336a421cbde62b57cf228"></a>

<a id="canonical-bd3a07b31e24a4919a81042d06fe93e65397f5873422a4f9e96a5709ff4c40b9"></a>

## tenant property — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 7

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

<a id="canonical-83723f6796b6275b05d97800db3fe6b8204312861395c8453e4b6c113a37263b"></a>

<a id="canonical-20f6497d7c5cd7e8f3aa2df309fcea8808939bbb14d932932f7992962e8b9705"></a>

## uid property — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 8

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

<a id="canonical-931da60bd9fc9712dfa1dd58b77910a0111e712c9f440841ab034fa73e484ca2"></a>

## Next pages — rate_limiter_allowed_prefixes / 36ad8e5dc081 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae8f6e65011f685c24bca3709661c52cf21abfde273f1570a231198e0c86b5c6"></a>

## request_cookies_to_add — request_cookies_to_add / 4d56eccd5396 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- request_cookies_to_add

<a id="canonical-776363ebebcb6241a454d9d20e337a98552f267a3928361042306ab6e36a8d62"></a>

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

<a id="canonical-e147941278baf34a7904bfa0a87c03c1215b6957ee76d61d01c9f73a84e43e10"></a>

## Direct properties — request_cookies_to_add / 4d56eccd5396 / 3

<a id="canonical-e45782640e78e648f2df05a60e62701df398125c24406ef4a1833b98fc219d60"></a>

<a id="canonical-86657e55759d8ecc78dada7b7242024bd95a3ed2d7116411be2fcf237f7936ce"></a>

## name property — request_cookies_to_add / 4d56eccd5396 / 4

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

<a id="canonical-db16ce1f8cb2e7429ec56120c0a6bb4fb75baa8a4892dac0b05d43f9e7445a25"></a>

<a id="canonical-b993630a0e3ba295fef2a0d3bb2256f85cb3e858cc8742a731e9b755a96c268c"></a>

## overwrite property — request_cookies_to_add / 4d56eccd5396 / 5

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827): complete subsection reference.

<a id="canonical-ca57d9fe157c450b5949a03bf5077280f9394dc61b739388b856ecf32d670b43"></a>

<a id="canonical-f3ad310ef050a2b80da00f1a271481f56e6762e8d3acb05aa5349a8fc8c289e4"></a>

## value property — request_cookies_to_add / 4d56eccd5396 / 6

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

<a id="canonical-05ec6c8772a4874c76e6ede7e918a615f2e7bc7491edaea87a543d1e8e13d6f6"></a>

## Next pages — request_cookies_to_add / 4d56eccd5396 / 7

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90fb73c509d6ef423f9a349cad8de46e265810f0e16a58643b1486625aef34a5"></a>

## request_cookies_to_add.secret_value — request_cookies_to_add.secret_value / 0d22eec30166 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a)
- request_cookies_to_add.secret_value

<a id="canonical-3ace17affd7cf4b8d293f5a6990cd1108ac8924d2361f5293ad4dbc71bbec36a"></a>

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

<a id="canonical-f9fe942cc964f26766af069636d3bb3dcb0f6590be7fafa0395a951b1b237a8f"></a>

## Direct properties — request_cookies_to_add.secret_value / 0d22eec30166 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-fa0e803c81eb65e2ca639498d65a8e95f3f8139aad2f1982988b54ce11ce7d0c): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-4901b76acfce9c6f9c8c13299efb9d0447c1af0834abbc094683d1316016cb5e): complete subsection reference.

<a id="canonical-e1243e097531c02fe3db119792c5de444b9366f530f47133ccf840753aa80e2d"></a>

## Next pages — request_cookies_to_add.secret_value / 0d22eec30166 / 4

- [request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-fa0e803c81eb65e2ca639498d65a8e95f3f8139aad2f1982988b54ce11ce7d0c)
- [request_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-4901b76acfce9c6f9c8c13299efb9d0447c1af0834abbc094683d1316016cb5e)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-fa0e803c81eb65e2ca639498d65a8e95f3f8139aad2f1982988b54ce11ce7d0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc4c2e4d4eec2d3d5b5db49fbffdc0cfb244c05f3042393a796933f0072804c8"></a>

## request_cookies_to_add.secret_value.blindfold_secret_info — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827)
- request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-52f95755b6912b1e92c4e0d76abeb5a1133956a8c87efcda1b23cbbf5a73a8c1"></a>

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

<a id="canonical-a09fb7672cf285d89eb2e2c46a2fe37c472e5d9e4d57cde4e940538e6818707f"></a>

## Direct properties — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 3

<a id="canonical-05581b4822299bb277086ad9fe01b41d47c077324caf2195d0f31508f9392c63"></a>

<a id="canonical-5acedfea4209fa4c2fa8b71c096788d8827b69305a70fab31dc2f877c706612a"></a>

## decryption_provider property — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 4

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

<a id="canonical-e381fb78cd600e45a9715ff4208ec325ad9208dc008782e0860c23a0668a09ef"></a>

<a id="canonical-cbb76e15e04ef4409fa0348dd82a7a176dcc49fba800408873ba851729a8c4bf"></a>

## location property — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 5

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

<a id="canonical-3df4cef01ebaa6c38fc3863a4ee397bd40adc35a695c4a8fdf0a67de20b31920"></a>

<a id="canonical-df698715c92ab062a6b0d825709b1dab601f392cd92dae5e8f82d135a5bd057c"></a>

## store_provider property — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 6

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

<a id="canonical-257f5eeccc0e900d80d72259159d7a6f8161ec67ea2f0302f240228df755d81a"></a>

## Next pages — request_cookies_to_add.secret_value.blindfold_secret_info / 4b232dc5479a / 7

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4901b76acfce9c6f9c8c13299efb9d0447c1af0834abbc094683d1316016cb5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8041535229f2e08c3ad15781c2be0080eed72d4486e40700fbfcb59aa7baec39"></a>

## request_cookies_to_add.secret_value.clear_secret_info — request_cookies_to_add.secret_value.clear_secret_info / 4a2f4f9caf5e / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a)
- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827)
- request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-c8da37600829c08f39ee7d0ab75668cd8c969595274f315c7b0875bef889f4a8"></a>

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

<a id="canonical-1ca51b02413e8de8692203a4a23ef9837c1eac651106327427be94ecf9995030"></a>

## Direct properties — request_cookies_to_add.secret_value.clear_secret_info / 4a2f4f9caf5e / 3

<a id="canonical-ee9c03b5f191210deab87d5a2f5b939c98eb689e3c03bb63a2b674684349d07c"></a>

<a id="canonical-bd79e4c250e0ac5376d6e919aadc0d75de60bf0f8f39495a5088fb7bed5aace3"></a>

## provider_ref property — request_cookies_to_add.secret_value.clear_secret_info / 4a2f4f9caf5e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-30e6d50dbd1aeb401bf4a9ceccca5e3f17b0c7166168d3ce25ee05acef773de5"></a>

<a id="canonical-50083344fe9b152c32bb7d8aaa71bf29252da78d413cf4267c6d0532cc2c4228"></a>

## url property — request_cookies_to_add.secret_value.clear_secret_info / 4a2f4f9caf5e / 5

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

<a id="canonical-27f743ad87a52ae7744fe45e1f10eb5cfbd4f7a235ca5ecf847904be00fac7e1"></a>

## Next pages — request_cookies_to_add.secret_value.clear_secret_info / 4a2f4f9caf5e / 6

- [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-a9186495b6512e9113ebccb27ba1a6f2ce7e2932d981d4347a9113229be71827)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efe13d2b93f834bf24dc53ce376b3b02e5fa2e9fe95212a458deb3f730e1e8c4"></a>

## request_headers_to_add — request_headers_to_add / 600960a4776e / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- request_headers_to_add

<a id="canonical-26f55a0b540c1767ac9b07c40e8f4d69de61898fa7f681101abb116b5799e1be"></a>

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

<a id="canonical-aa5b3011644ba1619d94fcda565e2b38714fcc9a88550bed63bf0e1c5ca8f61f"></a>

## Direct properties — request_headers_to_add / 600960a4776e / 3

<a id="canonical-150fa57524d7a8ed0fb6638718868e11fac55a371cdad49f444470143b1ec33b"></a>

<a id="canonical-1cdffaafe5fd04721b988f6a81787dd5df19d7f0193a592001132f3461bdcfae"></a>

## append property — request_headers_to_add / 600960a4776e / 4

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

<a id="canonical-8713a7dbf6b637110187b3da378c7c2a820ad1793813663875a19162c9f91a20"></a>

<a id="canonical-309acc6d1e54f89bb16cb3204ff2b8a5538903c9b1a59fbe42e44324a1c2484d"></a>

## name property — request_headers_to_add / 600960a4776e / 5

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237): complete subsection reference.

<a id="canonical-02a786784f681246896724d233dbb7d4ad6a8af1d80faca325b6b737a3161557"></a>

<a id="canonical-cc0c003ddecb7210c5343c1ffc489a6a74698a1bd045f24636a734e126762be0"></a>

## value property — request_headers_to_add / 600960a4776e / 6

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

<a id="canonical-f01ce5fafe7a12d881a4640bbbe19f4403e56972ecfa4e3b715a5375365c56da"></a>

## Next pages — request_headers_to_add / 600960a4776e / 7

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b89d82b9f9512fac9aeebf99ab537aca2dd1368c4d980d554a8a951ca7fa6038"></a>

## request_headers_to_add.secret_value — request_headers_to_add.secret_value / 11eed3915629 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17)
- request_headers_to_add.secret_value

<a id="canonical-5740343d46afc111b586fdf4a27bb2d742613296ba2a753a429a2cb20a2d184e"></a>

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

<a id="canonical-b8226dd355bd5ba47f7a5141bdd18cb836fb2a200df22bd2ffd38b94572adaee"></a>

## Direct properties — request_headers_to_add.secret_value / 11eed3915629 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-8f743425336a9e443fd9567906359c267f14a164e0a2e03dc8771bfcabce0f0a): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-c5de48ce6aab4993d57f779a6a5a3bd93a1870a69155fba411348ceb0b4ed92b): complete subsection reference.

<a id="canonical-a4c58b752670f375dd1a1460c1a38caf72222e7d5adce8f5eec32b425488bd86"></a>

## Next pages — request_headers_to_add.secret_value / 11eed3915629 / 4

- [request_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-8f743425336a9e443fd9567906359c267f14a164e0a2e03dc8771bfcabce0f0a)
- [request_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-c5de48ce6aab4993d57f779a6a5a3bd93a1870a69155fba411348ceb0b4ed92b)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-8f743425336a9e443fd9567906359c267f14a164e0a2e03dc8771bfcabce0f0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b33f553dd1090bc8bfb8c50586c0bd7105fadcfb5fb839e4623a7e36e6d3f0e5"></a>

## request_headers_to_add.secret_value.blindfold_secret_info — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237)
- request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-54225ff1210bf59e5bee19125a3975bc0f933b09f5994560ba87649a13132fa4"></a>

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

<a id="canonical-11ca88c73172733b0a900730b44222daed753e0ab20a69ef6440ca524a07def3"></a>

## Direct properties — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 3

<a id="canonical-e9c01de1a21d6586e4061e0525a4a9164a9ae4d77405ae6fe4b1ba58fbdb3114"></a>

<a id="canonical-a1ea8bb8e9326fd5ffa59603ec79f73df730dc8dfce3bf4efde7b0f3bf17d6e4"></a>

## decryption_provider property — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 4

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

<a id="canonical-5c114e6cd26190fa44a798355e58f43301c62dc7dc9ac76974b0fa34e0aa7276"></a>

<a id="canonical-6eb92ad6b2d62844dcf06fa36f4428f8662e4acb5a83435fb970e571fdc91205"></a>

## location property — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 5

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

<a id="canonical-f349c5437f0c0da90b078d429fe8e4bc6cb2a0c910fa462ab930f9e91c6af3c9"></a>

<a id="canonical-3526b50a9a568335617273a2907cb5a3120a049f36ca1d90878b5a0ab97b2b51"></a>

## store_provider property — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 6

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

<a id="canonical-3c75daafbb9ccdc52aecb50fac0ee26c0b3759a45e32fd006d4996c654e6e62c"></a>

## Next pages — request_headers_to_add.secret_value.blindfold_secret_info / 2607d58bad3b / 7

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-c5de48ce6aab4993d57f779a6a5a3bd93a1870a69155fba411348ceb0b4ed92b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd2def864f59d26777ffa5eecedae750025e928676a9b3dc67a5b1c8c99fd652"></a>

## request_headers_to_add.secret_value.clear_secret_info — request_headers_to_add.secret_value.clear_secret_info / 4e8dbf68ae47 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17)
- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237)
- request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-d95670bb6d4565add1fce8be201c50beab20bb7f5f3e2e7b8d0cb20a534595f1"></a>

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

<a id="canonical-6fafca8ab81dfa5e03468b7d9833888165255e7d93c9dd5cfc79294d891a7ce5"></a>

## Direct properties — request_headers_to_add.secret_value.clear_secret_info / 4e8dbf68ae47 / 3

<a id="canonical-60e847c190b9013c5fcb1e18e177d86929a2a95f18e11cfce57f110bacf6e198"></a>

<a id="canonical-3ed8676fc993f84a7273f87eabedbcd8104c5eeaa35b4ca04350eb3acb16387e"></a>

## provider_ref property — request_headers_to_add.secret_value.clear_secret_info / 4e8dbf68ae47 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bc2757555e16a04d04a9c067567a0c9e3a6fc6a6685fc5a364d1c0a0a5961a21"></a>

<a id="canonical-f0f068ee066607a5fdc33298ed70bc19419c50cb60c517c4cab4668470e7c1c1"></a>

## url property — request_headers_to_add.secret_value.clear_secret_info / 4e8dbf68ae47 / 5

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

<a id="canonical-9c5c05bcff604ed4145eed9a24b2a5a1e193c6f5b84bb3a711d5be82cdcb1040"></a>

## Next pages — request_headers_to_add.secret_value.clear_secret_info / 4e8dbf68ae47 / 6

- [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-7c39b8ee87886e7c0ff1dc26677823961aae1f765b7bb4749e0973f57d17f237)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c32dbd54789be4f166612ece66efbeb8854581c7aa0ccc7dae53d0ef225cb045"></a>

## response_cookies_to_add — response_cookies_to_add / 14623d8ff481 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- response_cookies_to_add

<a id="canonical-17981b85dfe84272d2b1d9bc88eacf351c9df395da7ca5a8e0ba26aef6c236ec"></a>

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

<a id="canonical-c5abb8d364481ff5de21a628222d058f51f6565abc415397c8484d91abb00cc4"></a>

## Direct properties — response_cookies_to_add / 14623d8ff481 / 3

<a id="canonical-011487866e5765de3cc2bc6d450b7b0deea30f43b0534d45d91c55e00726a158"></a>

<a id="canonical-135f1de5a0280cbbab5e1cfc41f2c186de92cf02645b23fb705a990e5d0d8544"></a>

## add_domain property — response_cookies_to_add / 14623d8ff481 / 4

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

<a id="canonical-3bf67dabc90db0a5629e82dec41e68054af9f24bfbc7b6a73318e698a19b4a42"></a>

<a id="canonical-0c1b3ffa1cbf7e98feaacd79b48d82fd1d35b3db4ea1e686e12f1a9b14e2e3cf"></a>

## add_expiry property — response_cookies_to_add / 14623d8ff481 / 5

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

- [add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-df6445affadfd15e9efd02775342309f3ad27296dcb98c885cc4f4d19d317352): complete subsection reference.

- [add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-70180bf2b7f2ec3bbc8644918e761495463c65413347faa4d3d940aad4325735): complete subsection reference.

<a id="canonical-2536a2a964bafba71b5445f16485b16af98f557e08d7576da5bee42e0a52138c"></a>

<a id="canonical-78b9e43fb2c5c8091e4a32cbc1e82e1ccd6aadd7dc0249290a5334d448a610df"></a>

## add_path property — response_cookies_to_add / 14623d8ff481 / 6

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

- [add_secure](data-sources--virtual_host--reference--group-002.md#canonical-5d56dadfb40e1a504d8ef98256ab959b7a1f10afa38907fa075bb2c3a50fd515): complete subsection reference.

- [ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-dcd1fb425d0375f68c2a082f85c82cc7f3161668e28dfcf2858dd747be9684f3): complete subsection reference.

- [ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-8c84c2a48ee05e55166489cb3f8056f6fff46f654b5435afc4450f7025ead3a0): complete subsection reference.

- [ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-97d76ba907d1d7a6c478e5ce616f82dd1ac3143b48bf622bad3c78518e293100): complete subsection reference.

- [ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-677437dc5e40c011762254304b49f58ffbf91a902d636e887a0325cbcbad3247): complete subsection reference.

- [ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-c713dd8098a7a2a206de31fa4bb1aa3045c2852e849d40ded49049859fed3884): complete subsection reference.

- [ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-52976862d593914fe989839d746b46627e4cd481a48fae17d4b27b2ea6684fe8): complete subsection reference.

- [ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-3cba5f5b1260bfb1ed7a0ea1f6a54485741b16bb405ea1e9fc29a31f36b33ab3): complete subsection reference.

- [ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-63fc5e5dc86dd4c8193879c07b95eb9ce203acf07c49f8456627338b106ab052): complete subsection reference.

- [ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-1373b8ded53f8b9262b40dee124c7fc2bec4a0c88881f4a2d33cc70cfd062542): complete subsection reference.

<a id="canonical-6bfac342c9a702da62e016157179940809832e204d8d3b1b6641f2024b7c4a0c"></a>

<a id="canonical-b10d45df2b890a50cd0f0b9e13fd9cfb968f28005c9741088240228087dfd737"></a>

## max_age_value property — response_cookies_to_add / 14623d8ff481 / 7

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

<a id="canonical-0f93d81a288807991fbff83e22dfafaadecc8e404c08f1c090aac28b45789ec7"></a>

<a id="canonical-e23a3c7dbb53ceddb9f66f95411b36c6a068b1b92bba5b6e3e7b4981a6e87051"></a>

## name property — response_cookies_to_add / 14623d8ff481 / 8

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

<a id="canonical-1d53b0e3f0d5cdba0424ecd2cf062970df72e7a2e374fefbb49aee2e3a41ee14"></a>

<a id="canonical-e1d2813171e49d33a592a8285bbf84a7bfe21401b5b581f8d45f07d701c0d725"></a>

## overwrite property — response_cookies_to_add / 14623d8ff481 / 9

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

- [samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-4ed6efcb1b920e92bea7379735d5bdb71c5c3adf6a09a60171bae81d97e4e3d9): complete subsection reference.

- [samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-af582aaeab33f9bfe179152325f15f5629b2a2a30467f3e41ab826b674bfcceb): complete subsection reference.

- [samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-7a8677ea4bef03cd2d3328ef1f10a0ddf2a11738e483554dc10c4d7b5e115480): complete subsection reference.

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e): complete subsection reference.

<a id="canonical-1e9a9fa0163187ad338fe527297e6786d6aabe1cf5159f7ab0cf1f392e18e517"></a>

<a id="canonical-e17d4a81ee36f98b94b60880400a337afca609f5992ada9bb3b5b79542b653eb"></a>

## value property — response_cookies_to_add / 14623d8ff481 / 10

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

<a id="canonical-f068938ca654c53f398faebfee564e9ac4493c9e06bc702a8ec4252a12cb6897"></a>

## Next pages — response_cookies_to_add / 14623d8ff481 / 11

- [response_cookies_to_add.add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-df6445affadfd15e9efd02775342309f3ad27296dcb98c885cc4f4d19d317352)
- [response_cookies_to_add.add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-70180bf2b7f2ec3bbc8644918e761495463c65413347faa4d3d940aad4325735)
- [response_cookies_to_add.add_secure](data-sources--virtual_host--reference--group-002.md#canonical-5d56dadfb40e1a504d8ef98256ab959b7a1f10afa38907fa075bb2c3a50fd515)
- [response_cookies_to_add.ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-dcd1fb425d0375f68c2a082f85c82cc7f3161668e28dfcf2858dd747be9684f3)
- [response_cookies_to_add.ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-8c84c2a48ee05e55166489cb3f8056f6fff46f654b5435afc4450f7025ead3a0)
- [response_cookies_to_add.ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-97d76ba907d1d7a6c478e5ce616f82dd1ac3143b48bf622bad3c78518e293100)
- [response_cookies_to_add.ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-677437dc5e40c011762254304b49f58ffbf91a902d636e887a0325cbcbad3247)
- [response_cookies_to_add.ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-c713dd8098a7a2a206de31fa4bb1aa3045c2852e849d40ded49049859fed3884)
- [response_cookies_to_add.ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-52976862d593914fe989839d746b46627e4cd481a48fae17d4b27b2ea6684fe8)
- [response_cookies_to_add.ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-3cba5f5b1260bfb1ed7a0ea1f6a54485741b16bb405ea1e9fc29a31f36b33ab3)
- [response_cookies_to_add.ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-63fc5e5dc86dd4c8193879c07b95eb9ce203acf07c49f8456627338b106ab052)
- [response_cookies_to_add.ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-1373b8ded53f8b9262b40dee124c7fc2bec4a0c88881f4a2d33cc70cfd062542)
- [response_cookies_to_add.samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-4ed6efcb1b920e92bea7379735d5bdb71c5c3adf6a09a60171bae81d97e4e3d9)
- [response_cookies_to_add.samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-af582aaeab33f9bfe179152325f15f5629b2a2a30467f3e41ab826b674bfcceb)
- [response_cookies_to_add.samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-7a8677ea4bef03cd2d3328ef1f10a0ddf2a11738e483554dc10c4d7b5e115480)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-df6445affadfd15e9efd02775342309f3ad27296dcb98c885cc4f4d19d317352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf6b21181f4c1de1c3ec23e9f7a24eaa0a0c6ad066b1c3852d00786b5710bd90"></a>

## response_cookies_to_add.add_httponly — response_cookies_to_add.add_httponly / 564775f1dca3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.add_httponly

<a id="canonical-fd83576e95d9adfcf52962b924aae8efba61eaa93cc123ef02a56646ad0eae1b"></a>

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

<a id="canonical-89a43937dc4f43e78c590711f9c3cca5b2d41285d3d382f2a59c7cd90f9fb545"></a>

## Direct properties — response_cookies_to_add.add_httponly / 564775f1dca3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28f188b0c768cabcb6c3cd79d1e8d35aa4b10ebb5b8b1a384ce2f29a866ddfe7"></a>

## Next pages — response_cookies_to_add.add_httponly / 564775f1dca3 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-70180bf2b7f2ec3bbc8644918e761495463c65413347faa4d3d940aad4325735"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90e5cbc8425d037d06b3ce3f5881040deb8d82d753ca807028c9a19117cc7771"></a>

## response_cookies_to_add.add_partitioned — response_cookies_to_add.add_partitioned / 4481e5c61be8 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.add_partitioned

<a id="canonical-8cd7289149232ac0bc0c748d4dbce4eee4c86b46e6720c2540aaff0e22ef8360"></a>

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

<a id="canonical-c0bce3130e21e1070c4206575edcfd4dbd1cfd88bb530ab4f934f72f77d66db5"></a>

## Direct properties — response_cookies_to_add.add_partitioned / 4481e5c61be8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38f46b74ae162002964113f55d7f154d9c2a3e419bb552c458a5af87e3793965"></a>

## Next pages — response_cookies_to_add.add_partitioned / 4481e5c61be8 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5d56dadfb40e1a504d8ef98256ab959b7a1f10afa38907fa075bb2c3a50fd515"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cd2de17d0a90cb19dfd60835ac1e7755c1ccf93f928a26a7de2da76ade40698"></a>

## response_cookies_to_add.add_secure — response_cookies_to_add.add_secure / ae826160a760 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.add_secure

<a id="canonical-c3930e9d64b1409e15f88f20a6d5cbf23c146efa5e95e788e8477b564b719404"></a>

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

<a id="canonical-0c18d4d8bb3a27dfe58338b947aec6914cb4ec847af75060f2f4c3eaa46552cb"></a>

## Direct properties — response_cookies_to_add.add_secure / ae826160a760 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a89202e95a1615dffd6656deede174a3380bccdf31a118eb603ab699393bc023"></a>

## Next pages — response_cookies_to_add.add_secure / ae826160a760 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-dcd1fb425d0375f68c2a082f85c82cc7f3161668e28dfcf2858dd747be9684f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15735b853061f28d64811017892e2ec8f54ffaa2a162467bbfa1042b86f463ce"></a>

## response_cookies_to_add.ignore_domain — response_cookies_to_add.ignore_domain / abf9c3199f14 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_domain

<a id="canonical-b066e999199324b697316b24e7884e28c05f2cab883479d3c00dd2cef2677dd6"></a>

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

<a id="canonical-c74cfe2b180b420e27c5eaff35fbc09bcccdf422ebad89a37cd5ec24a60e2275"></a>

## Direct properties — response_cookies_to_add.ignore_domain / abf9c3199f14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c47e556afa5ba9dee9430f5dbdb3bcd41e29f1819bd38657e5d2747eec4e6b7c"></a>

## Next pages — response_cookies_to_add.ignore_domain / abf9c3199f14 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-8c84c2a48ee05e55166489cb3f8056f6fff46f654b5435afc4450f7025ead3a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e1584d6cd209eb5241be8390a6077fdad7cf500c17c34dd857e57e16e460db6"></a>

## response_cookies_to_add.ignore_expiry — response_cookies_to_add.ignore_expiry / 44c05e3d26b3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_expiry

<a id="canonical-1dc8ff42b29f50cf3871db46b643f8920b41ba8d67b1ae77d92e59b9d1bc7b91"></a>

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

<a id="canonical-cf8e9562a4293ad0ec7d794f1f1fa3303ffab600e571ee973ddef7bf3378a557"></a>

## Direct properties — response_cookies_to_add.ignore_expiry / 44c05e3d26b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42f9d989238cc76523f2cc4c33b2de0328cafe3f60e3d92939d13fe80fb31f00"></a>

## Next pages — response_cookies_to_add.ignore_expiry / 44c05e3d26b3 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-97d76ba907d1d7a6c478e5ce616f82dd1ac3143b48bf622bad3c78518e293100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa0727d0c7012e8d8c01c7854a6fbb4c53f997538faaa68019e776077974a0be"></a>

## response_cookies_to_add.ignore_httponly — response_cookies_to_add.ignore_httponly / cf99c9340b63 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_httponly

<a id="canonical-3132e64a1698f2d98eac60ebac4bd742fc48de7affaa7d8ed8d33d963c57e7b6"></a>

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

<a id="canonical-e2e7a36d6226f30c3f8eb8e24670d6a4ae3e1d76f0db44fb28e9fc256ff132a6"></a>

## Direct properties — response_cookies_to_add.ignore_httponly / cf99c9340b63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9811bb9894336d19f50132473fe43c9970782c81ed0ca794731f212a36fb28c3"></a>

## Next pages — response_cookies_to_add.ignore_httponly / cf99c9340b63 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-677437dc5e40c011762254304b49f58ffbf91a902d636e887a0325cbcbad3247"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-543b5d0c88a85d5deaf796f668eedaefd3e0d0b24c328ef6f8d75261b905c710"></a>

## response_cookies_to_add.ignore_max_age — response_cookies_to_add.ignore_max_age / ea7c909bf8ab / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_max_age

<a id="canonical-76ffce866c875853bf94bf4fa1a3c66a79f771b474cb2f443f6450745934362c"></a>

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

<a id="canonical-8d15505750332e4415219fa940706b57dc1d32c9e8c0bbe603dd8c190a64e3ac"></a>

## Direct properties — response_cookies_to_add.ignore_max_age / ea7c909bf8ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c1821938e028a4e2cb6da9b5b1b8a9707321819f9a2c33b715e4f8059bd837be"></a>

## Next pages — response_cookies_to_add.ignore_max_age / ea7c909bf8ab / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-c713dd8098a7a2a206de31fa4bb1aa3045c2852e849d40ded49049859fed3884"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c106c831d68e397c9dbac324fd91476f99a4b82bfc70aaf9232c84a4efcfbc45"></a>

## response_cookies_to_add.ignore_partitioned — response_cookies_to_add.ignore_partitioned / 993709084abb / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_partitioned

<a id="canonical-d48dd46198ec4b23225c1b7cc7c746d6827e9ca9113da41d7d706e9fb1ab6eae"></a>

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

<a id="canonical-78d316e00cae7f1dc4dc89ddbbf5704bfb8e88df6042bbc938b5c0205c98835b"></a>

## Direct properties — response_cookies_to_add.ignore_partitioned / 993709084abb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49182643848528c4ef501aa166214319a49524f8e1326e827c77e8fd74ce4c9c"></a>

## Next pages — response_cookies_to_add.ignore_partitioned / 993709084abb / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-52976862d593914fe989839d746b46627e4cd481a48fae17d4b27b2ea6684fe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ad4f3d8e41e99b223a8964c827afe14bdd1972207da704f0f5dc29720156c24"></a>

## response_cookies_to_add.ignore_path — response_cookies_to_add.ignore_path / 2e85b05ad86f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_path

<a id="canonical-9ce19d337e3b49d0aa5d71e7b32ea5e176d2c3601a2cf7dd6cbfeb63ec7462fe"></a>

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

<a id="canonical-354db91cae737eccb493b71f8db886dafab87d01cfc7403262f4f3096fb96368"></a>

## Direct properties — response_cookies_to_add.ignore_path / 2e85b05ad86f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a3a5e4296d15d4cd7b2dfe9f82fb3ec03e6bf9d7b6af934a040d9c14cfbabf7"></a>

## Next pages — response_cookies_to_add.ignore_path / 2e85b05ad86f / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-3cba5f5b1260bfb1ed7a0ea1f6a54485741b16bb405ea1e9fc29a31f36b33ab3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e366ab329fa17ca443c1d8546aea9fcda68523fa6f8a330d421527f6adeb4d1"></a>

## response_cookies_to_add.ignore_samesite — response_cookies_to_add.ignore_samesite / f5d9989f9e66 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_samesite

<a id="canonical-c9e199a0e79fdfc4d4d76c5e2da178ea9d4909db78064b7fdd838aeb56eb9741"></a>

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

<a id="canonical-bc2958142a0acff4a388e44f3464e69522a2e278d7eee6959b6d6b320ebfd578"></a>

## Direct properties — response_cookies_to_add.ignore_samesite / f5d9989f9e66 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-486f20542269931c5db1dba771cd3bd5a9a753e681ef5cef8a307e91bc0b0f64"></a>

## Next pages — response_cookies_to_add.ignore_samesite / f5d9989f9e66 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-63fc5e5dc86dd4c8193879c07b95eb9ce203acf07c49f8456627338b106ab052"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17414098e55a531550a15db91e3b0b733f87def05db004b80eed3fc990c8abf3"></a>

## response_cookies_to_add.ignore_secure — response_cookies_to_add.ignore_secure / 9430547cf7b3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_secure

<a id="canonical-b20668cc5c116e9a8f819f6ee02c8f3f52a5556e29778696858fa93a4c8a994e"></a>

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

<a id="canonical-12265b8c8f44ae274f58ec785ce429037ec0a225188dcff0c66b2fb5e037736d"></a>

## Direct properties — response_cookies_to_add.ignore_secure / 9430547cf7b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe39897ffb4ff73ea2bd945b6922c11acaec5d83e05605214dc2136397d951ec"></a>

## Next pages — response_cookies_to_add.ignore_secure / 9430547cf7b3 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-1373b8ded53f8b9262b40dee124c7fc2bec4a0c88881f4a2d33cc70cfd062542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78c6eda6c82e2f62ea0140e3c7272742a6d272647501e0b4afb4dc7cba2e439c"></a>

## response_cookies_to_add.ignore_value — response_cookies_to_add.ignore_value / dbf9fda2a755 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.ignore_value

<a id="canonical-eadc1c8a157361ab2d134135565ab4f9a1cf303efaa0fa4f711d092c089d2c3d"></a>

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

<a id="canonical-cb316c4f4a8705b6aabe4bd7d60a71b434782606a497390a2055aae52b9c385e"></a>

## Direct properties — response_cookies_to_add.ignore_value / dbf9fda2a755 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b07e255f8603640d050fc8e610b7dcb2d3aab25618cf00bd5d88d15ff9120977"></a>

## Next pages — response_cookies_to_add.ignore_value / dbf9fda2a755 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4ed6efcb1b920e92bea7379735d5bdb71c5c3adf6a09a60171bae81d97e4e3d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6750d69cbcfdf09ce00f6daf1fa668639857e5a47b888aac72ca5c7df4113f32"></a>

## response_cookies_to_add.samesite_lax — response_cookies_to_add.samesite_lax / 57ac74b59647 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.samesite_lax

<a id="canonical-8f6665d32428f36ebf9519e3b83ff2250dc648a553546931270f3d4539fc8072"></a>

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

<a id="canonical-a194ae5bc52464d6454f47fea92981a2b5e3099a6c8884031944f2c51c6579a3"></a>

## Direct properties — response_cookies_to_add.samesite_lax / 57ac74b59647 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-995ac87f0519f39afc3066f82f64a1ee15da24d5596e3531dc6c72fd46f684ef"></a>

## Next pages — response_cookies_to_add.samesite_lax / 57ac74b59647 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-af582aaeab33f9bfe179152325f15f5629b2a2a30467f3e41ab826b674bfcceb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65238420e1c960f97d304d32b1ec60b7e8f08daaccfac4b89a0c315794dc5535"></a>

## response_cookies_to_add.samesite_none — response_cookies_to_add.samesite_none / 46774e2a0795 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.samesite_none

<a id="canonical-e8e123bb1258d25d91711149df2d072d31fa7cfd2e20b0ca2098dc20e25442c0"></a>

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

<a id="canonical-547da51c9d54f71e00b120bf64417c921a6ba71514dbf9d1e344139a20bd8053"></a>

## Direct properties — response_cookies_to_add.samesite_none / 46774e2a0795 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-560aa1d367a826374441c317381a7a4e1024b9a97696ee6eb23b07ce0a3134f4"></a>

## Next pages — response_cookies_to_add.samesite_none / 46774e2a0795 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-7a8677ea4bef03cd2d3328ef1f10a0ddf2a11738e483554dc10c4d7b5e115480"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab9e192ea981b381d015967432c581f37eb3e10132bc719e901a75b3f81cd8ba"></a>

## response_cookies_to_add.samesite_strict — response_cookies_to_add.samesite_strict / 5319744a1d23 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.samesite_strict

<a id="canonical-087a07a0b62cc9931167e5e530af13abddd1b61e4c8a1975cfafc4f04112d7ed"></a>

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

<a id="canonical-94d5f3985c53c8e6718ca2f0d1b84e37b56386adf4950e9f2133a6aa92042cfa"></a>

## Direct properties — response_cookies_to_add.samesite_strict / 5319744a1d23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a7340b919403dfdbbf4570df249997d87d24041f18a88cc0ad9594bdb2c8ebc"></a>

## Next pages — response_cookies_to_add.samesite_strict / 5319744a1d23 / 4

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-273bcbfa29571b0dd800d9d0680bce42ae6416122cbc707008b82a1ba2e5039e"></a>

## response_cookies_to_add.secret_value — response_cookies_to_add.secret_value / 054fb72850de / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- response_cookies_to_add.secret_value

<a id="canonical-2ca057e45985a0698098d6f385c06b7a2aad949748ba83aea794394def5630a6"></a>

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

<a id="canonical-238345d78191548d7bde1daa7f056af6c32e874960afda3db31ef4aeb3c68ec9"></a>

## Direct properties — response_cookies_to_add.secret_value / 054fb72850de / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-f0795365701d228ffe8381608648a45a8165b86878fb2794f9d46d637c28cc53): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-129efa49d8426d313689bd71fffd4fddcc835b5c5c4acc26dfacd1b596e14a89): complete subsection reference.

<a id="canonical-6416fa4a29510d628170f66ad744963a4555a700c1ff5eb5729224252c4b0a90"></a>

## Next pages — response_cookies_to_add.secret_value / 054fb72850de / 4

- [response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-f0795365701d228ffe8381608648a45a8165b86878fb2794f9d46d637c28cc53)
- [response_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-129efa49d8426d313689bd71fffd4fddcc835b5c5c4acc26dfacd1b596e14a89)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f0795365701d228ffe8381608648a45a8165b86878fb2794f9d46d637c28cc53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-158cff970d759bdae48fb1b8d020a6f064fa08861cdf822e1a1bd5640060fb89"></a>

## response_cookies_to_add.secret_value.blindfold_secret_info — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e)
- response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-19093bc3844ce31e8a3ab853fa57c213aa26ed286bbf25692ca2885a4a0bd762"></a>

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

<a id="canonical-69f268f718d994f34883c56725aa470719965286ecee631d0b9afc3134c90343"></a>

## Direct properties — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 3

<a id="canonical-065db219601c23acfcea759748e9ec121416bcfe971d30d6b0dbd67cefadf4c2"></a>

<a id="canonical-bad90930dae7513966f7854361ca47511d3d1e9f28564dd64ce68d3edb94b405"></a>

## decryption_provider property — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 4

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

<a id="canonical-326a745d61583c35108a2badf56536c6bbb59656902ec89c3f359b407eebc247"></a>

<a id="canonical-5f0c55a7fa4c3314e853f3463fccbebd26e5f76ac4c663fab0d76194b5fd7d8b"></a>

## location property — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 5

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

<a id="canonical-48da01f788c114232ee5bf78d95c6d206dd0234b3e633443165636811b0bb6a1"></a>

<a id="canonical-41fc1826d9335ad0e5e9c27b87ebdc858b9be0c701d3b0fb06bb8e9c61dd80ce"></a>

## store_provider property — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 6

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

<a id="canonical-ddad7a81d54a4c690d4c8aad36d67b7134aee8f51f5e01d01731a9769cbe5335"></a>

## Next pages — response_cookies_to_add.secret_value.blindfold_secret_info / 0e3eedf2b552 / 7

- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-129efa49d8426d313689bd71fffd4fddcc835b5c5c4acc26dfacd1b596e14a89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da677e83db00ea9bb7f26a77b5977c9f4510afe2350e07e6611636a0f6a68bab"></a>

## response_cookies_to_add.secret_value.clear_secret_info — response_cookies_to_add.secret_value.clear_secret_info / da147a47442a / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-f7a61bad0e5adc48004bd76fac09c27970f693dabd18642affd7d98567a441f2"></a>

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

<a id="canonical-9b3d16e7cbdf1b580b29c420f390a764e4b2c84ff8d2add29d66e80653fbc4da"></a>

## Direct properties — response_cookies_to_add.secret_value.clear_secret_info / da147a47442a / 3

<a id="canonical-9c410b6c15fe7c5aab1e431b24ad655e7e1ff17a3cac42a311e1eb0acf60bbbd"></a>

<a id="canonical-00eea4bdedc6af42d4472cb428cf51cca597495adb83b09403cc0fd924734bf8"></a>

## provider_ref property — response_cookies_to_add.secret_value.clear_secret_info / da147a47442a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bac287308f8c015a9581a10cbd104827ab61ebfad306700be94aabbf26b644a2"></a>

<a id="canonical-f7c175cb7d13abc33166c4a4d0279fdf0197288b49670a2c825cca9e6daf2210"></a>

## url property — response_cookies_to_add.secret_value.clear_secret_info / da147a47442a / 5

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

<a id="canonical-fc8d62439a504de585c7930ceecd130c258bb34d5f774213184298c08e4c53ea"></a>

## Next pages — response_cookies_to_add.secret_value.clear_secret_info / da147a47442a / 6

- [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-f7a6dc08794b99c07a7bbed7702c5ae0b03f9721464a0d3dd16a16e37c86bf5e)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed03db55074128718f30451ffd43eed42ebd5311ae9a2ef80b302f505692f3d2"></a>

## response_headers_to_add — response_headers_to_add / 227dcdc24582 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- response_headers_to_add

<a id="canonical-11881322be9bcd571d044f37f73bc11d8882cbc7f5d8624d981b5ed46b1ca462"></a>

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

<a id="canonical-d6e86b73ce466c3136a3e05660434c558d0cad358db4b6e5f22513f3054c3d14"></a>

## Direct properties — response_headers_to_add / 227dcdc24582 / 3

<a id="canonical-f30f676fcc8389dbd06d4858dde7008bafd2e05784b5969af8f0bce5cc37a7a3"></a>

<a id="canonical-c0a80d46a91bed3948967f589aa85fb7babdbe318bda3c554b94a4475983f944"></a>

## append property — response_headers_to_add / 227dcdc24582 / 4

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

<a id="canonical-835c0dc5e4dfd8f6a6ef3faedd57e29f6d0988f388033bc8a44c1facfed9cbf4"></a>

<a id="canonical-0f7cd6ee0c29316db2b91c52dbbffb56c5bb12818cb7e6731e4b610dc06bf2f6"></a>

## name property — response_headers_to_add / 227dcdc24582 / 5

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

- [secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98): complete subsection reference.

<a id="canonical-fadbfd5f0d1ea02eaf762a6750864c3df116f170c5744c17604f3edc245c39f3"></a>

<a id="canonical-eb1d32b00391f35a28cb6f93597fff964fc2e848ece58acfe6b3a0c1441d4590"></a>

## value property — response_headers_to_add / 227dcdc24582 / 6

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

<a id="canonical-231dfe75b4c7c17eabfa8ae2dc31836881dfc12d93f14446456568771862e3a5"></a>

## Next pages — response_headers_to_add / 227dcdc24582 / 7

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da12cac4758f6d3acc7160ab08b735fd07487ef08b3f59412c17dbf6b1c342da"></a>

## response_headers_to_add.secret_value — response_headers_to_add.secret_value / 04fa90fc5ee8 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597)
- response_headers_to_add.secret_value

<a id="canonical-d03a128181625d7bb378cd4b6bfda94f95fde6b90e53990e73477f7128ec1552"></a>

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

<a id="canonical-d78a0015d853a9b3eec362bb9357af6f86e641f0f04f8853d6b3068927886f49"></a>

## Direct properties — response_headers_to_add.secret_value / 04fa90fc5ee8 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-866b56ef7b9da9b671d6cf4905518ad7b37093c458e1a36098f3e0abd394c5ca): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-346f350d338ff03bab753fbd2037658858ada529b6b57def04242e62109efcf3): complete subsection reference.

<a id="canonical-870afd00020ac443f443fef0babedbfb650c867a24db5b3e093135a31a7c287e"></a>

## Next pages — response_headers_to_add.secret_value / 04fa90fc5ee8 / 4

- [response_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-866b56ef7b9da9b671d6cf4905518ad7b37093c458e1a36098f3e0abd394c5ca)
- [response_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-346f350d338ff03bab753fbd2037658858ada529b6b57def04242e62109efcf3)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-866b56ef7b9da9b671d6cf4905518ad7b37093c458e1a36098f3e0abd394c5ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df5aa5f2864f19afeaa6b9e132d0ea0391d8d91b658cfcd18117cf242370aaf7"></a>

## response_headers_to_add.secret_value.blindfold_secret_info — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-386b75b7fd8f7280081c41be1fcdb60631afb5c37b8a2d5713592c88d5da4d6f"></a>

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

<a id="canonical-4341e3c4e0e9002d29f8cf25753520dcef02a3f3d7310f48c2bd11e388b18c56"></a>

## Direct properties — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 3

<a id="canonical-b9a23ffd647f88ae03e760dcc8452bd7bc8ac42d931e6b8fa803233c6f16b0fd"></a>

<a id="canonical-323b0ccd6e8a2e3b8becb1a78435766340247cee33a21f006bc83b6981a69cc0"></a>

## decryption_provider property — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 4

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

<a id="canonical-c0fa0232605aeb5695b7eefd27ad0834230706328e89c335614947964e6cbeb4"></a>

<a id="canonical-261c5904016cf53de02b245f8281f072e513fe74195419eba227c28f6f955946"></a>

## location property — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 5

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

<a id="canonical-52d08406d242b8f5d41b8c841b40a8a60f3957da13de6c32e24386f391f502d0"></a>

<a id="canonical-1963a8b370299b84db0b24502f6199adeed3db286fedecacfb806b684df1b596"></a>

## store_provider property — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 6

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

<a id="canonical-ce7a789f671b83070eddba2a5691d2dda42119a5d62303b53ab126d1741e19f8"></a>

## Next pages — response_headers_to_add.secret_value.blindfold_secret_info / 7b8c784118f9 / 7

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-346f350d338ff03bab753fbd2037658858ada529b6b57def04242e62109efcf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cbb89dacb5740d25d8a8a03a3c78376a9d277942e3332dfd1742b2d089d57cc"></a>

## response_headers_to_add.secret_value.clear_secret_info — response_headers_to_add.secret_value.clear_secret_info / cab4e7ad7e9a / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-dc289188951f310f7fe23d2757c0cc29c4c73dce1d00bd497516d40a0fa1cf6f"></a>

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

<a id="canonical-89309b39dc1cdc2b16261a7d28169380ce63a853816961a64e31e491cc2795b2"></a>

## Direct properties — response_headers_to_add.secret_value.clear_secret_info / cab4e7ad7e9a / 3

<a id="canonical-dcd6c6b972da1cf0a1ffb013a98e77b069e2f82e223d3d215c3ee6e548da1287"></a>

<a id="canonical-f4ea8a7a76eebfcd56debbb815c79bc5b85e40e1547bfc27f42ee5e8eacd83fc"></a>

## provider_ref property — response_headers_to_add.secret_value.clear_secret_info / cab4e7ad7e9a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c94c007a7e846c42bbfaf244eb21ea7af9fc29b60259362c0269258c805939fa"></a>

<a id="canonical-24f735f111bfb350a62afa61cf63a2bcc63e58d720053198b972e7c2a26d8f1a"></a>

## url property — response_headers_to_add.secret_value.clear_secret_info / cab4e7ad7e9a / 5

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

<a id="canonical-9ad745c141bd2c4eb85c29c3e53b5a4f3966a983aeac734835f3e42d5be2eb7e"></a>

## Next pages — response_headers_to_add.secret_value.clear_secret_info / cab4e7ad7e9a / 6

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d1b39a8efb9a6396009107b0a08b25f7e23ebe21d7a87b0a58d364f0fc0b8c98)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-4934cf1290ecb6b5c3257959b6affcf382a8a9611d4070478d2b7a96a328bd49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3497aa6befd53bec1e8e040ce46595ea3697b8596534c8d04f07d05bd7cf214b"></a>

## retry_policy — retry_policy / 293ed874fffe / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- retry_policy

<a id="canonical-d9c0aee6ce40c260e86d51530ab8d2bba6f46075b3bf88617d4fa05b846d9051"></a>

Type: `"single"`. Computed.

Retry policy configuration for route destination.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cea6ee7a5a067af9356e1e970c583c3f5b6808f684b1983b03d4ceea76268b31"></a>

## Direct properties — retry_policy / 293ed874fffe / 3

- [back_off](data-sources--virtual_host--reference--group-002.md#canonical-e706c1ce08a49ebf7ef3150208cc45a5f95ecffae3fe6e7e23fc93a4d670fbb9): complete subsection reference.

<a id="canonical-68daadb14e4d16961d60a4e2f7e551d42a5f58561c959fe0549a1f0d524f13de"></a>

<a id="canonical-88ce950a569c7f2920f1c592e567744dc722278e1030be48035b293019275267"></a>

## num_retries property — retry_policy / 293ed874fffe / 4

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

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

<a id="canonical-3894889d54a7b1b23f4013773a2f2f1821abd0328c35a5017c0df1a2d9900cbb"></a>

<a id="canonical-e9b72df5787422a2299af23a4bb5377f8baeb0ce0536d31ef5298571e09ecabb"></a>

## per_try_timeout property — retry_policy / 293ed874fffe / 5

Type: `"number"`. Computed.

Specifies a non-zero timeout per retry attempt. In milliseconds.

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

<a id="canonical-9928b281c0798e6c8b32d624cb409569496fb93c8293b1f1d4c47acb1f2b6d50"></a>

<a id="canonical-2be9ca06a2d3bf19e3cb807215808d12ed5e3fe787f4d353d15433bb2e524828"></a>

## retriable_status_codes property — retry_policy / 293ed874fffe / 6

Type: `["list", "number"]`. Computed.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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

<a id="canonical-7c75f6412705bde23f13c3b0a02dd161cebf3e523ad6ab8fc58a90bb55abd7b6"></a>

<a id="canonical-c88f9fb82ab2248bb74d6a1007d24756c1fad4011e47376af91fb0ec7e79d07e"></a>

## retry_condition property — retry_policy / 293ed874fffe / 7

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3f30443f8994c76ffbf1c6e73def16de0a6a6176672f57e1ac8ffb4e4796f6da"></a>

## Next pages — retry_policy / 293ed874fffe / 8

- [retry_policy.back_off](data-sources--virtual_host--reference--group-002.md#canonical-e706c1ce08a49ebf7ef3150208cc45a5f95ecffae3fe6e7e23fc93a4d670fbb9)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-e706c1ce08a49ebf7ef3150208cc45a5f95ecffae3fe6e7e23fc93a4d670fbb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8052c6fc7cff4d2ff61e1eaf8232de255bdb4c034be3bad2333ba880ccd4fae1"></a>

## retry_policy.back_off — retry_policy.back_off / 2ae1840cf403 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-4934cf1290ecb6b5c3257959b6affcf382a8a9611d4070478d2b7a96a328bd49)
- retry_policy.back_off

<a id="canonical-90137b1f1dd9c158c985781b156e4912b77324a7f9b5cd9c9c92e0651e9f5b6f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-71e000aa6e870548fa5c431af8d94f1d61156bfac2d839fab6c9198fb5110658"></a>

## Direct properties — retry_policy.back_off / 2ae1840cf403 / 3

<a id="canonical-a378be87197252a3eecd1438eec0aac7d9f5ba3889c3e863e1088f675c0cbd61"></a>

<a id="canonical-f553d0225fb6a90219e31e63e8fb46813444b1c7bc8787c79d25e429c0fd07a2"></a>

## base_interval property — retry_policy.back_off / 2ae1840cf403 / 4

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

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

<a id="canonical-ac68f6a4f19a9e0e9129e1b937d73e308c1f52ad45dfc34b1ad23da891c32052"></a>

<a id="canonical-8084123bfe9deff584c7f876401e9d63df82089445b28d5db0cabbc889a87435"></a>

## max_interval property — retry_policy.back_off / 2ae1840cf403 / 5

Type: `"number"`. Computed.

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

<a id="canonical-ba4e25c006d648858a2cad3627ad8fcd0324ddddad2ba66e1f6556379c55b934"></a>

## Next pages — retry_policy.back_off / 2ae1840cf403 / 6

- [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-4934cf1290ecb6b5c3257959b6affcf382a8a9611d4070478d2b7a96a328bd49)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-e51606ec3c960457562b22f8d862e9f0eb68742c034a1b4373e9a7b926a9d694"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec4912e6caff6a510cc338b58852da66d7df36694f5b9f1b1f0a46618c76df0b"></a>

## routes — routes / 41209979738f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- routes

<a id="canonical-825bb61962da8698fa4c81ae326a69be2af94e3f6712315b000abdfd2fd05dbe"></a>

Type: `"list"`. Computed.

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

<a id="canonical-e3e927a755eb9d3150d629a042f412fed0a97f066ba53c5a4a74aac2aa72c8ae"></a>

## Direct properties — routes / 41209979738f / 3

<a id="canonical-6b232fa48bd96afb2945ddf3cb824d058d3c9ba6462543821c3a6b08eaa2c083"></a>

<a id="canonical-fa879cce75762d15faeea9cf35938ce3656fe7634222de0fcddeeafd8383474b"></a>

## kind property — routes / 41209979738f / 4

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

<a id="canonical-a2c384541714a0ae946ef06a5d333077eec8a1c5ebbd45e48625f62f552ce378"></a>

<a id="canonical-73dfa1b947fa86653dd28affbf8f36c086360a80596076cc4347cd9fd9907530"></a>

## name property — routes / 41209979738f / 5

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

<a id="canonical-245a46803a727013bfb5d6cc8cd5d667cdab7d6fdbea2ab4a55c6b945b327cda"></a>

<a id="canonical-cce19764f6ffb816f9e81be4a3cbb05bab14d1ca879743cfba6c155acc4f61a7"></a>

## namespace property — routes / 41209979738f / 6

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

<a id="canonical-f8ce44d497bbb69aacec032256df3fe040ad350767792a4d9f4bc73a4d58f3e6"></a>

<a id="canonical-e0a3ab2e61b49e578f3b165623c9f3dab60417d967baa31475bb2d26987a31f0"></a>

## tenant property — routes / 41209979738f / 7

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

<a id="canonical-dbf4e4162c3876182e186c9ad869252efecf79988a20498f7aeb3b029b3ba2e8"></a>
