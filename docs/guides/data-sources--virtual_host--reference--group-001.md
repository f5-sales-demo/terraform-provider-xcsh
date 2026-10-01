---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4a11f5d7e44f89caa74a09328bc67e10a1193ca5348f707f412a2615f8a3ea2"></a>

## Property reference — Property reference / 4ebaf40b9c3f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- Property reference

<a id="canonical-9da0e4e0bfb7dcf5075575bb4aea5866b0f4d3e9fa56d6aa56a8e7ab3da8145e"></a>

## Direct properties — Property reference / 4ebaf40b9c3f / 3

<a id="canonical-4c76d0e23eac1ec0f39f17987a8c2ce6ca00f963a1f58d4726b7e56003a45e56"></a>

<a id="canonical-91dca04137d2f9b417e6cfd22065041426ae4129e99e7c19500c106a974ef799"></a>

## add_location property — Property reference / 4ebaf40b9c3f / 4

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-1021cfe9bb5566ddd87b255a7eccf4493464e67a5a1f176e3667d41f4f67d516): complete subsection reference.

<a id="canonical-a4e2ec726d8c0568aaa80e5e912e98f1d2914fd2aec364168e7f101dc9d8b82f"></a>

<a id="canonical-285d93fa1693206db0c098c2a79ebb18c12128a6e78be1d9dff0af08ea8f911b"></a>

## annotations property — Property reference / 4ebaf40b9c3f / 5

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

<a id="canonical-3b5e1c2baafdbb27b339acc6521212d06ff95a39c4cd403097fcc2395aaa73e3"></a>

<a id="canonical-103207491a0ac639d6fbd2e0c1cf39ec1b070a5bdbdf4e7f073f0fba1ff00adf"></a>

## append_server_name property — Property reference / 4ebaf40b9c3f / 6

Type: `"string"`. Computed.

\[OneOf: append\_server\_name, default\_header, pass\_through, server\_name; Default:
default\_header\] Exclusive with \[default\_header pass\_through server\_name\] Specifies the value
to be used for Server header if it is not already present. If Server Header is already present it is
not overwritten. It is just passed.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Specifies the value to be used for
Server header if it is not already present. If Server Header is already present it is not
overwritten. It is just passed.

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

OneOf alternatives in this subsection:

- [append_server_name](data-sources--virtual_host--reference--group-001.md#canonical-3b5e1c2baafdbb27b339acc6521212d06ff95a39c4cd403097fcc2395aaa73e3)
- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-1db43c5339da33f1c2a8f4536a86d2286e9309f41d9b07058244c35b01325b12)
- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-85b1046ead71ab45112df5f0a65c1bb1a1473365aa0626737f0354f3b3cc8339)
- [server_name](data-sources--virtual_host--reference--group-001.md#canonical-807e60475c19e2d90c79afea33ef50293340f687899143ddbdfe578428d0af2b)

Select alternatives according to the provider validators above.

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660): complete subsection reference.

- [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-67c6bb4057c626103a8169b0ef713a2bdb3c67563949f25de13828f8d848654f): complete subsection reference.

- [captcha_challenge](data-sources--virtual_host--reference--group-001.md#canonical-cce1978f17446fe7bf3d9c0dd2a62ce4ea2a4f0e13f8fd7f2ffff4c33710c445): complete subsection reference.

- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705): complete subsection reference.

- [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-175e663c8201caf980e21bdc50984a06a8b563d7d07c88dfceb91ced493567cd): complete subsection reference.

<a id="canonical-5d032c217cda296ab82cdb0ed7f1575bcf1d816f114761007ec8b424ee561cc6"></a>

<a id="canonical-bef653d47fca40a8211b6e9f099927a563470cb6ec413c7fa999e4ba2449c35e"></a>

## connection_idle_timeout property — Property reference / 4ebaf40b9c3f / 7

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-a4ecd779378b174be7fa310c13ecda6cf2402208e2a96236ae26a14b61393c5b): complete subsection reference.

- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6): complete subsection reference.

<a id="canonical-35fd48a5e8c713e17130bbfddbf55b4d16fb23b4406be1c67fe7c6728afa2069"></a>

<a id="canonical-785624200382a742fd1c8acafcf5412fb0da29fb1259a00d311598d4d022ea20"></a>

## custom_errors property — Property reference / 4ebaf40b9c3f / 8

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value is the uri\_ref. Currently supported URL schemes
is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify
this message as base64 encoded plain text message e.g. "Access Denied" or it can be HTML paragraph
or a body string encoded as base64 string E.g. "&lt;p&gt; Access Denied &lt;/p&gt;". Base64 encoded
string for this HTML is "PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==" Specific response code takes preference
when both response code and response code class matches for a request.

The configured custom errors are only applicable for loadbalancer generated errors. Errors returned
from upstream server is propagated as is.

F5XC provides default error pages for the errors generated by the loadbalancer. Content of these
pages are not editable. User has an option to disable the use of default F5XC error pages.

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
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-2ba5db0f8ddd27b58c29095c0986c994d4d2c03100a06100429e44c738a20fe4): complete subsection reference.

- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-be86a1bd888037fdc3242c0d59516ab12221888e99a192b774b252f4409bcb2d): complete subsection reference.

<a id="canonical-46948f8280ae95e32d371a1172e2c9d9acd18c6e01002a4cbb3efb752518b30b"></a>

<a id="canonical-443e64d2ff2b226a1bd36de03ac54d06c504674afede9598d701c5886ca41b84"></a>

## description property — Property reference / 4ebaf40b9c3f / 9

Type: `"string"`. Computed.

Description of the VirtualHost.

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

<a id="canonical-f16c72ae3bcc3d8cac5f7843160cdc623973296cb1b09a90a138d19ec34d17f2"></a>

<a id="canonical-3326c12877f1aaf7e2cd135457bafbe99049b7fc38aead0e51f9fbd9763f2ae1"></a>

## disable_default_error_pages property — Property reference / 4ebaf40b9c3f / 10

Type: `"bool"`. Computed.

Option to specify whether to disable using default F5XC error pages.

Upstream description:

An option to specify whether to disable using default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-99a1e2187f0839f820e3f970b5b1f890ac850359518999444e89fca7d1478ae7"></a>

<a id="canonical-0efc1fa1cf2a17cae0bf236f093eacd379834e0296f39eb83c7927bd72ebb3af"></a>

## disable_dns_resolve property — Property reference / 4ebaf40b9c3f / 11

Type: `"bool"`. Computed.

Disable DNS resolution for domains specified in the virtual host When the virtual host is configured
as Dynamive Resolve Proxy (DRP), disable DNS resolution for domains configured. This configuration
is suitable for HTTP CONNECT proxy.

Upstream description:

Disable DNS resolution for domains specified in the virtual host

When the virtual host is configured as Dynamive Resolve Proxy (DRP), disable DNS resolution for
domains configured. This configuration is suitable for HTTP CONNECT proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-fbdf7fd14ab032fa59e6f0edf908260deddc86d23ea5ac6cb183d78a9e97ca96): complete subsection reference.

<a id="canonical-814abfd8d9cf74c76f884a731cfe874889873929cba0b6326d61bd8ee6593631"></a>

<a id="canonical-9c59d87b3b672c46a402c86c4b5580292c852d162d2345eba16ca75bcc3295ce"></a>

## domains property — Property reference / 4ebaf40b9c3f / 12

Type: `["list", "string"]`. Computed.

List of domain names matched to this virtual host for routing incoming requests. Supports wildcard
patterns like \*.example.com for subdomain matching.

Upstream description:

A list of Domains (host/authority header) that will be matched to this Virtual Host. Wildcard hosts
are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the virtual host proxy type is
TCP\_PROXY\_WITH\_SNI/HTTPS\_PROXY Domains also indicate the list of names for which DNS resolution
will be automatically resolved to IP addresses by the system.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 33,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 33,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-b8e8c7fcb7321f9bdf837de330392a7f4162a7567cf7f68cb61c12c9040bbf3e): complete subsection reference.

- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-1171379e2864a9edbff9631c19170289cd393aea92e23cbd8fbf34b9e78ce7e2): complete subsection reference.

- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a): complete subsection reference.

<a id="canonical-1b808917fd4454059d2eda60727a70b2399a47a295ba874467131f5bf7a0ab73"></a>

<a id="canonical-081ff71b08c61fefb1562756a643692e81a279fcb9fcf7f381f8234264f4874c"></a>

## id property — Property reference / 4ebaf40b9c3f / 13

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-8e2b523893827a3093e3ebadae8845b18aaaba93d9c3f8dd8f0ee8da00e7ceeb"></a>

<a id="canonical-13d9b2ff82d7dd5d2bb4fd705be0dd08411ca480147475b0215c82c648cdcf25"></a>

## idle_timeout property — Property reference / 4ebaf40b9c3f / 14

Type: `"number"`. Computed.

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity. Idle timeout and Proxy Type: HTTP\_PROXY, HTTPS\_PROXY: Idle timer
is started when the first byte is received on the connection. Each time an encode/decode event for..

Upstream description:

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity.

Idle timeout and Proxy Type:

HTTP\_PROXY, HTTPS\_PROXY: Idle timer is started when the first byte is received on the connection.
Each time an encode/decode event for headers or data is processed for the stream, the timer will be
reset. If the timeout fires, the stream is terminated with a 504 (Gateway Timeout) error code if no
upstream response header has been received, otherwise a stream reset occurs. The default idle
timeout is 30 seconds

TCP PROXY, TCP\_PROXY\_WITH\_SNI, SMA\_PROXY: The idle timeout is defined as the period in which
there are no bytes sent or received on either the upstream or downstream connection. The default
idle timeout is 1 hour.

UDP PROXY: The idle timeout for sessions. Idle timeout is defined as the period in which there are
no datagrams sent or received on the session. The default if not specified is 1 minute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-35d4e92967faff2a93e6f02288638f65dd0c817f65630593c95b78f48dc76894): complete subsection reference.

<a id="canonical-b943740f4c151f3ba4125e1dd5c831b0abad9648e22dd96639ab1184d184afc5"></a>

<a id="canonical-4bb9cb1c266782f739a11c74d73c158d03a2d730e74d5acdfe1713c8c83c6411"></a>

## labels property — Property reference / 4ebaf40b9c3f / 15

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

<a id="canonical-56c468ea0fe91490925c5123a2e75de8db686219727436d2e739f60f938bc28f"></a>

<a id="canonical-04055a1aaa1fed27be907fc6be11831e030a72b06f1f00b5ad2704414eeac340"></a>

## max_request_header_size property — Property reference / 4ebaf40b9c3f / 16

Type: `"number"`. Computed.

The maximum request header size in KiB for incoming connections. If un-configured, the default max
request headers allowed is 60 KiB. Requests that exceed this limit will receive a 431 response.

Upstream description:

The maximum request header size in KiB for incoming connections.

If un-configured, the default max request headers allowed is 60 KiB.

Requests that exceed this limit will receive a 431 response.

The max configurable limit is 96 KiB, based on current implementation constraints.

Note: a. This configuration parameter is applicable only for HTTP\_PROXY and HTTPS\_PROXY b. When
multiple HTTP\_PROXY virtual hosts share the same advertise policy, the effective "maximum request
header size" for such virtual hosts is the highest value configured on any of the virtual hosts.

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

<a id="canonical-9836c09b914be170027f862463b99a79363a3720c3f2586e1632509bf7d3658f"></a>

<a id="canonical-375be9842cf7562dd1f5fa8d14d0eb5694d009615293b74bfd794217ed6e77d3"></a>

## max_requests_per_connection property — Property reference / 4ebaf40b9c3f / 17

Type: `"number"`. Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests a downstream client can send over a single connection to Envoy. Enter
a value &gt;=1 to define the request limit per connection.

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

OneOf alternatives in this subsection:

- [max_requests_per_connection](data-sources--virtual_host--reference--group-001.md#canonical-9836c09b914be170027f862463b99a79363a3720c3f2586e1632509bf7d3658f)
- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-aab4430cdc26268dd04c870fd9a8040e5206acac957535669f5b3b0ec91639e3)

Select alternatives according to the provider validators above.

<a id="canonical-1f10a3621eb1c513986c298ddb6f85bff8d360d5130151826fe97cd8fee8ce11"></a>

<a id="canonical-d95d1462494eb38ae9698ebee05450f12e443eda500bf27f514daeee580b3882"></a>

## name property — Property reference / 4ebaf40b9c3f / 18

Type: `"string"`. Required.

Name of the VirtualHost.

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

<a id="canonical-53a039a7b5c06f6febd76ca4ccc5eb4ad29f8d07e4645ce10e2c4df03f967767"></a>

<a id="canonical-0ff8af1fdd3edead943c3ded79624b85abc227415b582bb5312bc4c945dfe92b"></a>

## namespace property — Property reference / 4ebaf40b9c3f / 19

Type: `"string"`. Required.

Namespace where the VirtualHost exists.

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

- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-6fd9adf9858ac044d27f69605c0da0afab5ae1521dd18568d820e2c4370a249b): complete subsection reference.

- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-d44525cc14415a80f6e329d02112bdf067ba82809fab4500bda9866ea672ad6f): complete subsection reference.

- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-545aca29b311fea151c0f133136e8abac4d81e741b6298394ac83aad08e117f3): complete subsection reference.

- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-8b5b859bf066db8c50680f61439eb6236af261b0f24ab581d2fa842959be914d): complete subsection reference.

- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-62dbd01f9aad94f490c77cc7b4d61c54a4fe212c25d4796a116058d8f3256475): complete subsection reference.

<a id="canonical-8f57e684607e734a5ccff4418e117cdd235cf71faacf03b03f00cf4b88c86e7d"></a>

<a id="canonical-e1ff0d2f56ef04c6b28168ba3035b88fdb4fe837cb52c0dd2288023944eea218"></a>

## proxy property — Property reference / 4ebaf40b9c3f / 20

Type: `"string"`. Computed.

\[Enum:
UDP\_PROXY|SMA\_PROXY|DNS\_PROXY|ZTNA\_PROXY|UZTNA\_PROXY|TMM\_HTTP\_PROXY|TMM\_HTTPS\_PROXY|TMM\_TCP\_PROXY|TMM\_UDP\_PROXY|TMM\_QUIC\_PROXY\]
ProxyType tells the type of proxy to install for the virtual host. Only the following combination of
VirtualHosts within same AdvertisePolicy is permitted (None of them should have '\*' in domains when
used with other VirtualHosts in same AdvertisePolicy) 1. Multiple TCP\_PROXY\_WITH\_SNI and..
Possible values are \`UDP\_PROXY\`, \`SMA\_PROXY\`, \`DNS\_PROXY\`, \`ZTNA\_PROXY\`,
\`UZTNA\_PROXY\`, \`TMM\_HTTP\_PROXY\`, \`TMM\_HTTPS\_PROXY\`, \`TMM\_TCP\_PROXY\`,
\`TMM\_UDP\_PROXY\`, \`TMM\_QUIC\_PROXY\`.

Upstream description:

ProxyType tells the type of proxy to install for the virtual host.

Only the following combination of VirtualHosts within same AdvertisePolicy is permitted (None of
them should have "\*" in domains when used with other VirtualHosts in same AdvertisePolicy)
&#8203;1. Multiple TCP\_PROXY\_WITH\_SNI and multiple HTTPS\_PROXY &#8203;2. Multiple HTTP\_PROXY
&#8203;3. Multiple HTTPS\_PROXY &#8203;4. Multiple TCP\_PROXY\_WITH\_SNI

HTTPS\_PROXY without TLS parameters is not permitted
HTTP\_PROXY/HTTPS\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY with empty domains is not permitted
TCP\_PROXY\_WITH\_SNI/SMA\_PROXY should not have "\*" in domains

&#8203;- HTTP\_PROXY: HTTP\_PROXY

Install HTTP proxy. HTTP Proxy is the default proxy installed. &#8203;- TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- TLS\_TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TLS\_TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- HTTPS\_PROXY: HTTPS\_PROXY

Install HTTPS proxy &#8203;- UDP\_PROXY: UDP\_PROXY

Install UDP proxy &#8203;- SMA\_PROXY: SMA\_PROXY

Install Secret Management Access proxy &#8203;- DNS\_PROXY: DNS\_PROXY

Install DNS proxy &#8203;- ZTNA\_PROXY: ZTNA\_PROXY

Install ZTNA proxy.this is going to be deprecated with UZTNA\_PROXY. &#8203;- UZTNA\_PROXY:
UZTNA\_PROXY

Install UZTNA proxy &#8203;- TMM\_HTTP\_PROXY: TMM\_HTTP\_PROXY

Install TMM HTTP proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_HTTPS\_PROXY: TMM\_HTTPS\_PROXY

Install TMM HTTPS proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_TCP\_PROXY: TMM\_TCP\_PROXY

Install TMM TCP proxy for TCP traffic. Used by TMM proxy type. &#8203;- TMM\_UDP\_PROXY:
TMM\_UDP\_PROXY

Install TMM UDP proxy for UDP traffic. Used by TMM proxy type. &#8203;- TMM\_QUIC\_PROXY:
TMM\_QUIC\_PROXY

Install TMM QUIC proxy for HTTP/3 traffic. Used by TMM proxy type.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-5fd02a3e53a4014bceaf5d9ab95dd66a94fd084b4552ea6670553171736272ca): complete subsection reference.

- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a): complete subsection reference.

<a id="canonical-d9f8c7954f38a11d2d16afcf05ffcc400572468f198e1e960443db5efd45fe32"></a>

<a id="canonical-d05e1e1cd438c609b3430c94fe2b46479226f2fd8782c123e868fefa3c6700a8"></a>

## request_cookies_to_remove property — Property reference / 4ebaf40b9c3f / 21

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

- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17): complete subsection reference.

<a id="canonical-1eb0b6363b51fa35a0417472f20cc4ad60afc1c6c6e3d301423ce48a993582d4"></a>

<a id="canonical-38018ad4f3b117374162a68a57af126200b15594218e58c066e7953f6e3d9d05"></a>

## request_headers_to_remove property — Property reference / 4ebaf40b9c3f / 22

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca): complete subsection reference.

<a id="canonical-8617633d0f2d1ae4771d64a479397b4a8363a20e7ed095581c299971fae384c0"></a>

<a id="canonical-1b3e80c73348da6af4440db196f38c608df5ea5ec1cfc17fc6e9327122d9d395"></a>

## response_cookies_to_remove property — Property reference / 4ebaf40b9c3f / 23

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

- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597): complete subsection reference.

<a id="canonical-0ffb469e9294b334fdde156a38b7328d2ea14f3faf5649fe75b0e54f4cd541e0"></a>

<a id="canonical-d9e1031baf28b8898fac729f491fe7605487c2854c433b72d5f134cbccaa1289"></a>

## response_headers_to_remove property — Property reference / 4ebaf40b9c3f / 24

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-4934cf1290ecb6b5c3257959b6affcf382a8a9611d4070478d2b7a96a328bd49): complete subsection reference.

- [routes](data-sources--virtual_host--reference--group-002.md#canonical-e51606ec3c960457562b22f8d862e9f0eb68742c034a1b4373e9a7b926a9d694): complete subsection reference.

- [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-5570059452e7a42407d2b110e7733a91dd1e1471402407dd0b9bd2b2184c3e65): complete subsection reference.

<a id="canonical-807e60475c19e2d90c79afea33ef50293340f687899143ddbdfe578428d0af2b"></a>

<a id="canonical-b6c1ef74f77740c34b6f2195e9a066e4e112bbf1637511ed760790a4eaad5d1e"></a>

## server_name property — Property reference / 4ebaf40b9c3f / 25

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

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

- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-88cb41571444c3a26fbfd8e039fa894468f209f6f9fb8b8397263ebd3c25f6e5): complete subsection reference.

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575): complete subsection reference.

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741): complete subsection reference.

- [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-42caf6807b285559de11020a7412783936bce26b48dc1c3b083cd332528d2cb1): complete subsection reference.

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1): complete subsection reference.

<a id="canonical-4fba7ae8f957ce10eb12c44c40c0dfbad122b193552681519eebb439aa51d039"></a>

## All schema paths — Property reference / 4ebaf40b9c3f / 26

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `add_location` | [add_location](data-sources--virtual_host--reference--group-001.md#canonical-4c76d0e23eac1ec0f39f17987a8c2ce6ca00f963a1f58d4726b7e56003a45e56) |
| `advertise_policies` | [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-05494ad126ed1d3d8ba1432041ec4189d78e08a936c599cded20b69ba316aa25) |
| `advertise_policies.kind` | [advertise_policies.kind](data-sources--virtual_host--reference--group-001.md#canonical-c7b24cfc89036f33952d2fd6110ca9692c01fba3b742a6a66bad6b7fe9759928) |
| `advertise_policies.name` | [advertise_policies.name](data-sources--virtual_host--reference--group-001.md#canonical-c2ccbc9a8a9e9ba0c51d2d9c344fd40c818dba19fbedc83d5c7e5dda5402e09e) |
| `advertise_policies.namespace` | [advertise_policies.namespace](data-sources--virtual_host--reference--group-001.md#canonical-d9eb86e615062680ce1019603c57547768149d62cb983cdfdc0cb423040af98f) |
| `advertise_policies.tenant` | [advertise_policies.tenant](data-sources--virtual_host--reference--group-001.md#canonical-f7223ce543c2df2c1540a8e008b1e6acde62ab3d3c49c63effcbe42fffbe8005) |
| `advertise_policies.uid` | [advertise_policies.uid](data-sources--virtual_host--reference--group-001.md#canonical-c96b2044c69182ac5db579740ac35f3ed0dc7b7a06d6f9d6c68aaf4c58fafa2b) |
| `annotations` | [annotations](data-sources--virtual_host--reference--group-001.md#canonical-a4e2ec726d8c0568aaa80e5e912e98f1d2914fd2aec364168e7f101dc9d8b82f) |
| `append_server_name` | [append_server_name](data-sources--virtual_host--reference--group-001.md#canonical-3b5e1c2baafdbb27b339acc6521212d06ff95a39c4cd403097fcc2395aaa73e3) |
| `authentication` | [authentication](data-sources--virtual_host--reference--group-001.md#canonical-545d089bac57d022dd93f2f3a9229a961d1ed6d51b8979e737d8f6507e2aa40e) |
| `authentication.auth_config` | [authentication.auth_config](data-sources--virtual_host--reference--group-001.md#canonical-432e29e671ba2a926c08990f17d6508c1241450e6e6b14a0682008e22d75d3c9) |
| `authentication.auth_config.kind` | [authentication.auth_config.kind](data-sources--virtual_host--reference--group-001.md#canonical-2b283c1a1d34b26c560827e967be9b9d7427d7b6cbb1a2548e9fd5daf1e87e78) |
| `authentication.auth_config.name` | [authentication.auth_config.name](data-sources--virtual_host--reference--group-001.md#canonical-93216c981497701474c8bb072c2e7d2d840b948c81d14b9502535ced6ce18b1d) |
| `authentication.auth_config.namespace` | [authentication.auth_config.namespace](data-sources--virtual_host--reference--group-001.md#canonical-ebf7f8d4f9606c5a964019f9b31958b4b94af96ae491d855775592378db883ed) |
| `authentication.auth_config.tenant` | [authentication.auth_config.tenant](data-sources--virtual_host--reference--group-001.md#canonical-f7831c240b154a2a979c7e4b426e09e32b9beb61d27086d011855baad5d048df) |
| `authentication.auth_config.uid` | [authentication.auth_config.uid](data-sources--virtual_host--reference--group-001.md#canonical-7340a4ba0928ae551f1808689d2893076110559fbf13b564a5afdf087b1412d7) |
| `authentication.cookie_params` | [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-3f8c10ebccfe81c63b92370a682e1acc15a154076d82be89c809465b66f4fb6f) |
| `authentication.cookie_params.auth_hmac` | [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-aef7595d637d8c179f4df76f8bdea5911b7189356015cd5feef2a633b30ec7ef) |
| `authentication.cookie_params.auth_hmac.prim_key` | [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-7bddd40bc4d2064bef3369cd7299ca145681cf66e50eea985f2870fb0661f4e4) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-c0341d65c47e460f596a8091270d8400c47ae79bf344553fcf337852cfc5007c) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-001.md#canonical-4c5f77d1a5dbd3b7c97d8ffa77201c71efdc5e953b5f03f78fffea3065503bd0) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-001.md#canonical-0188008038a59052d496e926585d0c6ce6461346ececa785c4643d4e74cb1431) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-001.md#canonical-ca9343ebb91dbb37a78509a4b79b83270cea1648be48656d5de5b16f6c1c5ad8) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-09355611c22c3b0f9ae0e9772b3994622e8db963d66e7a9b530ba5823f3497e5) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-001.md#canonical-3544944703b5e2476d775b2431e4c242c0f892bf95dd62cb1df8fc5eebae309b) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url](data-sources--virtual_host--reference--group-001.md#canonical-c4236e2869ef013afb8a35f5ef58a0d1b371aae075ed16a82bf2a56643725288) |
| `authentication.cookie_params.auth_hmac.prim_key_expiry` | [authentication.cookie_params.auth_hmac.prim_key_expiry](data-sources--virtual_host--reference--group-001.md#canonical-f11adc1269cc8a3196f37f64f73aad6c6d3e58b7e903187ca59dee2ef73a2c54) |
| `authentication.cookie_params.auth_hmac.sec_key` | [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a992fac48d63bcf50835428af8cc39e519ab2ea96e2a6e10c5fed98d015092f7) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-0ca8b20ed853e4aaa551092427fd6dd3f3aaceaa5201dcff7be414c15bd785d3) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-001.md#canonical-21ad37a05736aeb27698a9e5937fccf037e0e9beddbadfecda8b5cce2e5557a7) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-001.md#canonical-c791032b62709107e5e5022a153d5e282e44f6d0b23e41e945bc803d58c3dab1) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-001.md#canonical-b2b7869e53882c67bf684e140ba6a2e4c3fa2764cc7349f3b9a95b6c54d70a9a) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-7d85ecda68dbb92d34f1a32fc48486b062ed2410061dd28f04a57bdbbb4b2742) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-001.md#canonical-f4c899bdcd373efcc298575bb1247eb7606b139157da8a4a321178d74c2cc779) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url](data-sources--virtual_host--reference--group-001.md#canonical-38bfcbee7fc3b6a2e43a666f44f4d250aec9343b5c40b2e9ada4137bf50515f1) |
| `authentication.cookie_params.auth_hmac.sec_key_expiry` | [authentication.cookie_params.auth_hmac.sec_key_expiry](data-sources--virtual_host--reference--group-001.md#canonical-d573333500f65ae28a46a28eb4613068fc9474d56542f5087e48db6c71c67d6b) |
| `authentication.cookie_params.cookie_expiry` | [authentication.cookie_params.cookie_expiry](data-sources--virtual_host--reference--group-001.md#canonical-034b5f1e3779a60c7bc855c5470479941faf8df4dca3e3c48ccbc43283878f7b) |
| `authentication.cookie_params.cookie_refresh_interval` | [authentication.cookie_params.cookie_refresh_interval](data-sources--virtual_host--reference--group-001.md#canonical-0660651eedcf79a126ac2ab14aa6c68c010766b222fa49ee182c580ef4b1127b) |
| `authentication.cookie_params.kms_key_hmac` | [authentication.cookie_params.kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-0bf2c24828e319c09f6253cc7cc80a51008086f3085993f4365caed4edf773e1) |
| `authentication.cookie_params.session_expiry` | [authentication.cookie_params.session_expiry](data-sources--virtual_host--reference--group-001.md#canonical-fb85439fef7c781a9158124c8819940106ffb7351b41919ce09f759b3664577b) |
| `authentication.redirect_dynamic` | [authentication.redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-2355951834bb37b5c69f3115561b18f77c431cb74e70a148410180684b3358e1) |
| `authentication.redirect_url` | [authentication.redirect_url](data-sources--virtual_host--reference--group-001.md#canonical-09f89f9369443451dc474e3f055fa399b3dbf6e24b754f3aeefb127b5d3cc6b7) |
| `authentication.use_auth_object_config` | [authentication.use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-e435948b4c0aa0bf6868c064699abbad12773800cf5acccab7a1a5403bc7b4e5) |
| `buffer_policy` | [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-571f6f41d81c6e65e1d841cceb9d5ecdc4ec0e5603c345cbadcdf3ca663f1e90) |
| `buffer_policy.disabled` | [buffer_policy.disabled](data-sources--virtual_host--reference--group-001.md#canonical-accdec73b60de3448f2e425634ac3f50af8bbd97b3c198235a0f274d7f81a096) |
| `buffer_policy.max_request_bytes` | [buffer_policy.max_request_bytes](data-sources--virtual_host--reference--group-001.md#canonical-ae0dc59b0331ea3e2799c396aa7f986ca394b3e4272cfc0e8ec7f26134e4e437) |
| `captcha_challenge` | [captcha_challenge](data-sources--virtual_host--reference--group-001.md#canonical-3bff95472b997489389cf2f7c2c5183a9a42f9ac1cdcfee69125a1f00eae7087) |
| `captcha_challenge.cookie_expiry` | [captcha_challenge.cookie_expiry](data-sources--virtual_host--reference--group-001.md#canonical-b738d882bb776eb6864073de272b09b745e5bfc608c0c9dbd1d771508d5f6d19) |
| `captcha_challenge.custom_page` | [captcha_challenge.custom_page](data-sources--virtual_host--reference--group-001.md#canonical-4bcc08a2b4476c9745ad1f75c41a0270c7b6dbb917dd7b7706f7c943a3762939) |
| `coalescing_options` | [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-c32e9c2c91a02111b18820cf076c03b93d03a47dbb3f60320151bbfd9fa07fe1) |
| `coalescing_options.default_coalescing` | [coalescing_options.default_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-3e141c92a8b3542f5a934497060e5c156eaade44f25b9aa40b5917412391d941) |
| `coalescing_options.strict_coalescing` | [coalescing_options.strict_coalescing](data-sources--virtual_host--reference--group-002.md#canonical-f93a8970153609fab7ae74d14e444c62c2a215e08b02f13b3247f25e29f30bee) |
| `compression_params` | [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-371420dea8f387a112f8045b268676bd8537814a14e5c594eab27903140dce19) |
| `compression_params.content_length` | [compression_params.content_length](data-sources--virtual_host--reference--group-002.md#canonical-6290ae6c6709f11f19438b38fd66d40436b1fb2937cbf04c3dfbf2a9ea373223) |
| `compression_params.content_type` | [compression_params.content_type](data-sources--virtual_host--reference--group-002.md#canonical-d977c3567a3a070c269757e513b5ef7c90a7d747129ab0b47bb62f1b208a3e11) |
| `compression_params.disable_on_etag_header` | [compression_params.disable_on_etag_header](data-sources--virtual_host--reference--group-002.md#canonical-8d823396df1bc90ad06f0272c574aae4e320929beba249e756b5e16967f61f94) |
| `compression_params.remove_accept_encoding_header` | [compression_params.remove_accept_encoding_header](data-sources--virtual_host--reference--group-002.md#canonical-7842ee8f6db8a9c5d3337853b35502a1bf9fa5a401bebfe132db88bad7801677) |
| `connection_idle_timeout` | [connection_idle_timeout](data-sources--virtual_host--reference--group-001.md#canonical-5d032c217cda296ab82cdb0ed7f1575bcf1d816f114761007ec8b424ee561cc6) |
| `cors_policy` | [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-8029f76942051a56649d4af7cfb097baa2d2b95163103c177bf9ac7bdcabc600) |
| `cors_policy.allow_credentials` | [cors_policy.allow_credentials](data-sources--virtual_host--reference--group-002.md#canonical-84829179233191f2cb737ea0aeb5f0c87ec1be9a5f6f913b93011f1f735f7e8e) |
| `cors_policy.allow_headers` | [cors_policy.allow_headers](data-sources--virtual_host--reference--group-002.md#canonical-ff53a860e055ee96273bae6bb4c34cc5cb8d29f141e34129ca0aea3ecca00dd4) |
| `cors_policy.allow_methods` | [cors_policy.allow_methods](data-sources--virtual_host--reference--group-002.md#canonical-c50f9fbf03cc3fe20f50d762d93d29fd19a737f48cb5f7b28ead3ecc4552c43b) |
| `cors_policy.allow_origin` | [cors_policy.allow_origin](data-sources--virtual_host--reference--group-002.md#canonical-4bb5aea485eb7effdff0031033a673dbc0262b899a4e4e49b8011b3a05111834) |
| `cors_policy.allow_origin_regex` | [cors_policy.allow_origin_regex](data-sources--virtual_host--reference--group-002.md#canonical-4cb72b296e1985049897043ce56a3d24cabb822e6739537cd7c2c2aa05b0654e) |
| `cors_policy.disabled` | [cors_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-e4797bd6b552dfa313b4d5eb82d6f3614f420cfa148a0bab291de8026a8a9415) |
| `cors_policy.expose_headers` | [cors_policy.expose_headers](data-sources--virtual_host--reference--group-002.md#canonical-bc1a8a5df6d1f0a3ec1f313d07c52ed36e25f2754d563fa1e412df3a52975e54) |
| `cors_policy.maximum_age` | [cors_policy.maximum_age](data-sources--virtual_host--reference--group-002.md#canonical-fa2a338e7cd18a2e1dfc2979f53a7626d2ef38b97a5e57f48485b1e2f78c9998) |
| `csrf_policy` | [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-cdadf81518d323a24c95b4a540fcc368734f947f5e5277f695854487b8b7b400) |
| `csrf_policy.all_load_balancer_domains` | [csrf_policy.all_load_balancer_domains](data-sources--virtual_host--reference--group-002.md#canonical-5ac40d12ac1fe90bcc33d902c65d860f2fc9abffd31636611b3fa5bb80f57249) |
| `csrf_policy.custom_domain_list` | [csrf_policy.custom_domain_list](data-sources--virtual_host--reference--group-002.md#canonical-778a7081ef1667154c4681cb407fd72b0276d3ec3293af36f7baca1494ec9b19) |
| `csrf_policy.custom_domain_list.domains` | [csrf_policy.custom_domain_list.domains](data-sources--virtual_host--reference--group-002.md#canonical-17c09f70e87d350f220aef66cca398ca17a6a17d797ec069d5560c8c36c58a9e) |
| `csrf_policy.disabled` | [csrf_policy.disabled](data-sources--virtual_host--reference--group-002.md#canonical-3b3eb6a71fc1dbaba2504303bef46edf3f7dbb756b68df9ecac48b74387c31fe) |
| `custom_errors` | [custom_errors](data-sources--virtual_host--reference--group-001.md#canonical-35fd48a5e8c713e17130bbfddbf55b4d16fb23b4406be1c67fe7c6728afa2069) |
| `default_header` | [default_header](data-sources--virtual_host--reference--group-002.md#canonical-1db43c5339da33f1c2a8f4536a86d2286e9309f41d9b07058244c35b01325b12) |
| `default_loadbalancer` | [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-2393bca9603456e5118af941fbb8a44800e59cf3bfa0ba99f0e7fed04e27f438) |
| `description` | [description](data-sources--virtual_host--reference--group-001.md#canonical-46948f8280ae95e32d371a1172e2c9d9acd18c6e01002a4cbb3efb752518b30b) |
| `disable_default_error_pages` | [disable_default_error_pages](data-sources--virtual_host--reference--group-001.md#canonical-f16c72ae3bcc3d8cac5f7843160cdc623973296cb1b09a90a138d19ec34d17f2) |
| `disable_dns_resolve` | [disable_dns_resolve](data-sources--virtual_host--reference--group-001.md#canonical-99a1e2187f0839f820e3f970b5b1f890ac850359518999444e89fca7d1478ae7) |
| `disable_path_normalize` | [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-f1495d990daaf02cc2afa0a7d57e14e4e49d847696ade3fe5b735d2a24114f2c) |
| `domains` | [domains](data-sources--virtual_host--reference--group-001.md#canonical-814abfd8d9cf74c76f884a731cfe874889873929cba0b6326d61bd8ee6593631) |
| `dynamic_reverse_proxy` | [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-f6043a214957e33bb842cdc42d48bdeced74065fac3db0a95d4b8b7525ac0d2e) |
| `dynamic_reverse_proxy.connection_timeout` | [dynamic_reverse_proxy.connection_timeout](data-sources--virtual_host--reference--group-002.md#canonical-e5a1e7b3d67d861fe2d459e042014203b8b0925e74099b4b5e7858883db3e19c) |
| `dynamic_reverse_proxy.resolution_network` | [dynamic_reverse_proxy.resolution_network](data-sources--virtual_host--reference--group-002.md#canonical-cf1d00a7a61beb8a113ba4d1e32f7bb283d45cd78545778b536153cd25931693) |
| `dynamic_reverse_proxy.resolution_network.kind` | [dynamic_reverse_proxy.resolution_network.kind](data-sources--virtual_host--reference--group-002.md#canonical-f07149fa8f2d9cdd3a298e448d433bdc0f751f56d0f725911b39005d53a4f784) |
| `dynamic_reverse_proxy.resolution_network.name` | [dynamic_reverse_proxy.resolution_network.name](data-sources--virtual_host--reference--group-002.md#canonical-3ec1ac8ea14ff42ee1e21e69cf21b5b0129c5e2188f5cedcbd040fbef5b2f3a9) |
| `dynamic_reverse_proxy.resolution_network.namespace` | [dynamic_reverse_proxy.resolution_network.namespace](data-sources--virtual_host--reference--group-002.md#canonical-cc347ae2de0c9ba2ce2089216708fa54aa46aa705a0aab508c052a82d25e998c) |
| `dynamic_reverse_proxy.resolution_network.tenant` | [dynamic_reverse_proxy.resolution_network.tenant](data-sources--virtual_host--reference--group-002.md#canonical-a967b1bf8cd5e74019bab82995c55b7862ee2e0dc55b165079c7d4a0b3df56f3) |
| `dynamic_reverse_proxy.resolution_network.uid` | [dynamic_reverse_proxy.resolution_network.uid](data-sources--virtual_host--reference--group-002.md#canonical-5e2dd412116bff9bb84497ecac6c8abc4aa9ddfec2054a6ad5f1af68da403e58) |
| `dynamic_reverse_proxy.resolution_network_type` | [dynamic_reverse_proxy.resolution_network_type](data-sources--virtual_host--reference--group-002.md#canonical-c47035729c92719f9f6b7f688971663d09b481dbd9155844a3157c983fe93717) |
| `dynamic_reverse_proxy.resolve_endpoint_dynamically` | [dynamic_reverse_proxy.resolve_endpoint_dynamically](data-sources--virtual_host--reference--group-002.md#canonical-950fdfc61208072977d6af5de3fdbaf966ce149ba87804eadc5f3a4e60882029) |
| `enable_path_normalize` | [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-212b027873cc3a831d39cc1391c20838522a550bc2066de82e60509f37e8dd36) |
| `http_protocol_options` | [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-b45f4704e2098f4f35efec5d0287a0adbcf1b6ca5844e04cb7e65a575ff1f297) |
| `http_protocol_options.http_protocol_enable_v1_only` | [http_protocol_options.http_protocol_enable_v1_only](data-sources--virtual_host--reference--group-002.md#canonical-9833a9ba874394cd9e5a7523b486c7b43695475ffa606ec8c3e547d734d06aa1) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-18fadd1917758db465a449070d23ff247f1c36cb5795e9aa020f738bb4d29fed) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-0c6687e127d595e890513e3869fd455127341febe2cfffb2c49d783d250f5773) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-aff6b9d1cb32a877251a31259db00b624d56832f65fadfd9446c1d766ad9190b) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--virtual_host--reference--group-002.md#canonical-6c82bf72820a713ac74edbccd8cea0198b52713a73de406a11fb5dddf2c6061b) |
| `http_protocol_options.http_protocol_enable_v1_v2` | [http_protocol_options.http_protocol_enable_v1_v2](data-sources--virtual_host--reference--group-002.md#canonical-f3864f1653d1cbfa6f9f59ddef8d0f91c7b05dc135119d8066c8cee45587be64) |
| `http_protocol_options.http_protocol_enable_v2_only` | [http_protocol_options.http_protocol_enable_v2_only](data-sources--virtual_host--reference--group-002.md#canonical-a540c67759efc43fe558b6900c86cf50d3322032559d16284126f2258ae6a92b) |
| `id` | [id](data-sources--virtual_host--reference--group-001.md#canonical-1b808917fd4454059d2eda60727a70b2399a47a295ba874467131f5bf7a0ab73) |
| `idle_timeout` | [idle_timeout](data-sources--virtual_host--reference--group-001.md#canonical-8e2b523893827a3093e3ebadae8845b18aaaba93d9c3f8dd8f0ee8da00e7ceeb) |
| `js_challenge` | [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-6fead0d412383114980e7657466e05e6825f8ea21f77f0cfff01ac7f11f7cf40) |
| `js_challenge.cookie_expiry` | [js_challenge.cookie_expiry](data-sources--virtual_host--reference--group-002.md#canonical-b75ff2bc483ebbc39c86f7c34209b74f2722de6a8f099e345f6e5b3b45c4b421) |
| `js_challenge.custom_page` | [js_challenge.custom_page](data-sources--virtual_host--reference--group-002.md#canonical-72b10fbea2de566816df88f765778ed2f47b9a6374dc48679741c29304d5eee5) |
| `js_challenge.js_script_delay` | [js_challenge.js_script_delay](data-sources--virtual_host--reference--group-002.md#canonical-13628b4d5e0b945a16120939f3d9f9e772006c9b0439ff155a3b3cb356efabac) |
| `labels` | [labels](data-sources--virtual_host--reference--group-001.md#canonical-b943740f4c151f3ba4125e1dd5c831b0abad9648e22dd96639ab1184d184afc5) |
| `max_request_header_size` | [max_request_header_size](data-sources--virtual_host--reference--group-001.md#canonical-56c468ea0fe91490925c5123a2e75de8db686219727436d2e739f60f938bc28f) |
| `max_requests_per_connection` | [max_requests_per_connection](data-sources--virtual_host--reference--group-001.md#canonical-9836c09b914be170027f862463b99a79363a3720c3f2586e1632509bf7d3658f) |
| `name` | [name](data-sources--virtual_host--reference--group-001.md#canonical-1f10a3621eb1c513986c298ddb6f85bff8d360d5130151826fe97cd8fee8ce11) |
| `namespace` | [namespace](data-sources--virtual_host--reference--group-001.md#canonical-53a039a7b5c06f6febd76ca4ccc5eb4ad29f8d07e4645ce10e2c4df03f967767) |
| `no_authentication` | [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-524320b0777e1f50d78a30949d30b94597acbf8a2389b0c8271854f1eef75fd4) |
| `no_challenge` | [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3f0bdb5749658baa498a33d910c44024358aa16e77e1a62fe2a2b49bfe09053c) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-aab4430cdc26268dd04c870fd9a8040e5206acac957535669f5b3b0ec91639e3) |
| `non_default_loadbalancer` | [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-4c2f3a3c6f3ec034edc006c8cbff7413f0ac22898e8d8ea3779a2d86f956d9c3) |
| `pass_through` | [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-85b1046ead71ab45112df5f0a65c1bb1a1473365aa0626737f0354f3b3cc8339) |
| `proxy` | [proxy](data-sources--virtual_host--reference--group-001.md#canonical-8f57e684607e734a5ccff4418e117cdd235cf71faacf03b03f00cf4b88c86e7d) |
| `rate_limiter_allowed_prefixes` | [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-3df10ea042418c3f98638797c2a3b88c40e51e055f8603cc7a55497bc592e0d4) |
| `rate_limiter_allowed_prefixes.kind` | [rate_limiter_allowed_prefixes.kind](data-sources--virtual_host--reference--group-002.md#canonical-c32a11b6486b38446933a6e0cee529eb85881650a40dbb74c17caccbccbd1bcd) |
| `rate_limiter_allowed_prefixes.name` | [rate_limiter_allowed_prefixes.name](data-sources--virtual_host--reference--group-002.md#canonical-b031c297824dadaf3e516dd6e2ddbe026cd7a1c24de5502b90a06fe0c32f16f3) |
| `rate_limiter_allowed_prefixes.namespace` | [rate_limiter_allowed_prefixes.namespace](data-sources--virtual_host--reference--group-002.md#canonical-1bfb76bfafe95aa8ae00f1e222a4b40c38f1b93219ec9ce365095142c488984b) |
| `rate_limiter_allowed_prefixes.tenant` | [rate_limiter_allowed_prefixes.tenant](data-sources--virtual_host--reference--group-002.md#canonical-3e7b9f690462aeb52e5e7ace32ee8a75e3249673b13336a421cbde62b57cf228) |
| `rate_limiter_allowed_prefixes.uid` | [rate_limiter_allowed_prefixes.uid](data-sources--virtual_host--reference--group-002.md#canonical-83723f6796b6275b05d97800db3fe6b8204312861395c8453e4b6c113a37263b) |
| `request_cookies_to_add` | [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-776363ebebcb6241a454d9d20e337a98552f267a3928361042306ab6e36a8d62) |
| `request_cookies_to_add.name` | [request_cookies_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-e45782640e78e648f2df05a60e62701df398125c24406ef4a1833b98fc219d60) |
| `request_cookies_to_add.overwrite` | [request_cookies_to_add.overwrite](data-sources--virtual_host--reference--group-002.md#canonical-db16ce1f8cb2e7429ec56120c0a6bb4fb75baa8a4892dac0b05d43f9e7445a25) |
| `request_cookies_to_add.secret_value` | [request_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3ace17affd7cf4b8d293f5a6990cd1108ac8924d2361f5293ad4dbc71bbec36a) |
| `request_cookies_to_add.secret_value.blindfold_secret_info` | [request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-52f95755b6912b1e92c4e0d76abeb5a1133956a8c87efcda1b23cbbf5a73a8c1) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-05581b4822299bb277086ad9fe01b41d47c077324caf2195d0f31508f9392c63) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.location` | [request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-e381fb78cd600e45a9715ff4208ec325ad9208dc008782e0860c23a0668a09ef) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-3df4cef01ebaa6c38fc3863a4ee397bd40adc35a695c4a8fdf0a67de20b31920) |
| `request_cookies_to_add.secret_value.clear_secret_info` | [request_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-c8da37600829c08f39ee7d0ab75668cd8c969595274f315c7b0875bef889f4a8) |
| `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-ee9c03b5f191210deab87d5a2f5b939c98eb689e3c03bb63a2b674684349d07c) |
| `request_cookies_to_add.secret_value.clear_secret_info.url` | [request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-30e6d50dbd1aeb401bf4a9ceccca5e3f17b0c7166168d3ce25ee05acef773de5) |
| `request_cookies_to_add.value` | [request_cookies_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-ca57d9fe157c450b5949a03bf5077280f9394dc61b739388b856ecf32d670b43) |
| `request_cookies_to_remove` | [request_cookies_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-d9f8c7954f38a11d2d16afcf05ffcc400572468f198e1e960443db5efd45fe32) |
| `request_headers_to_add` | [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-26f55a0b540c1767ac9b07c40e8f4d69de61898fa7f681101abb116b5799e1be) |
| `request_headers_to_add.append` | [request_headers_to_add.append](data-sources--virtual_host--reference--group-002.md#canonical-150fa57524d7a8ed0fb6638718868e11fac55a371cdad49f444470143b1ec33b) |
| `request_headers_to_add.name` | [request_headers_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-8713a7dbf6b637110187b3da378c7c2a820ad1793813663875a19162c9f91a20) |
| `request_headers_to_add.secret_value` | [request_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-5740343d46afc111b586fdf4a27bb2d742613296ba2a753a429a2cb20a2d184e) |
| `request_headers_to_add.secret_value.blindfold_secret_info` | [request_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-54225ff1210bf59e5bee19125a3975bc0f933b09f5994560ba87649a13132fa4) |
| `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-e9c01de1a21d6586e4061e0525a4a9164a9ae4d77405ae6fe4b1ba58fbdb3114) |
| `request_headers_to_add.secret_value.blindfold_secret_info.location` | [request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-5c114e6cd26190fa44a798355e58f43301c62dc7dc9ac76974b0fa34e0aa7276) |
| `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-f349c5437f0c0da90b078d429fe8e4bc6cb2a0c910fa462ab930f9e91c6af3c9) |
| `request_headers_to_add.secret_value.clear_secret_info` | [request_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-d95670bb6d4565add1fce8be201c50beab20bb7f5f3e2e7b8d0cb20a534595f1) |
| `request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-60e847c190b9013c5fcb1e18e177d86929a2a95f18e11cfce57f110bacf6e198) |
| `request_headers_to_add.secret_value.clear_secret_info.url` | [request_headers_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-bc2757555e16a04d04a9c067567a0c9e3a6fc6a6685fc5a364d1c0a0a5961a21) |
| `request_headers_to_add.value` | [request_headers_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-02a786784f681246896724d233dbb7d4ad6a8af1d80faca325b6b737a3161557) |
| `request_headers_to_remove` | [request_headers_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-1eb0b6363b51fa35a0417472f20cc4ad60afc1c6c6e3d301423ce48a993582d4) |
| `response_cookies_to_add` | [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-17981b85dfe84272d2b1d9bc88eacf351c9df395da7ca5a8e0ba26aef6c236ec) |
| `response_cookies_to_add.add_domain` | [response_cookies_to_add.add_domain](data-sources--virtual_host--reference--group-002.md#canonical-011487866e5765de3cc2bc6d450b7b0deea30f43b0534d45d91c55e00726a158) |
| `response_cookies_to_add.add_expiry` | [response_cookies_to_add.add_expiry](data-sources--virtual_host--reference--group-002.md#canonical-3bf67dabc90db0a5629e82dec41e68054af9f24bfbc7b6a73318e698a19b4a42) |
| `response_cookies_to_add.add_httponly` | [response_cookies_to_add.add_httponly](data-sources--virtual_host--reference--group-002.md#canonical-fd83576e95d9adfcf52962b924aae8efba61eaa93cc123ef02a56646ad0eae1b) |
| `response_cookies_to_add.add_partitioned` | [response_cookies_to_add.add_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-8cd7289149232ac0bc0c748d4dbce4eee4c86b46e6720c2540aaff0e22ef8360) |
| `response_cookies_to_add.add_path` | [response_cookies_to_add.add_path](data-sources--virtual_host--reference--group-002.md#canonical-2536a2a964bafba71b5445f16485b16af98f557e08d7576da5bee42e0a52138c) |
| `response_cookies_to_add.add_secure` | [response_cookies_to_add.add_secure](data-sources--virtual_host--reference--group-002.md#canonical-c3930e9d64b1409e15f88f20a6d5cbf23c146efa5e95e788e8477b564b719404) |
| `response_cookies_to_add.ignore_domain` | [response_cookies_to_add.ignore_domain](data-sources--virtual_host--reference--group-002.md#canonical-b066e999199324b697316b24e7884e28c05f2cab883479d3c00dd2cef2677dd6) |
| `response_cookies_to_add.ignore_expiry` | [response_cookies_to_add.ignore_expiry](data-sources--virtual_host--reference--group-002.md#canonical-1dc8ff42b29f50cf3871db46b643f8920b41ba8d67b1ae77d92e59b9d1bc7b91) |
| `response_cookies_to_add.ignore_httponly` | [response_cookies_to_add.ignore_httponly](data-sources--virtual_host--reference--group-002.md#canonical-3132e64a1698f2d98eac60ebac4bd742fc48de7affaa7d8ed8d33d963c57e7b6) |
| `response_cookies_to_add.ignore_max_age` | [response_cookies_to_add.ignore_max_age](data-sources--virtual_host--reference--group-002.md#canonical-76ffce866c875853bf94bf4fa1a3c66a79f771b474cb2f443f6450745934362c) |
| `response_cookies_to_add.ignore_partitioned` | [response_cookies_to_add.ignore_partitioned](data-sources--virtual_host--reference--group-002.md#canonical-d48dd46198ec4b23225c1b7cc7c746d6827e9ca9113da41d7d706e9fb1ab6eae) |
| `response_cookies_to_add.ignore_path` | [response_cookies_to_add.ignore_path](data-sources--virtual_host--reference--group-002.md#canonical-9ce19d337e3b49d0aa5d71e7b32ea5e176d2c3601a2cf7dd6cbfeb63ec7462fe) |
| `response_cookies_to_add.ignore_samesite` | [response_cookies_to_add.ignore_samesite](data-sources--virtual_host--reference--group-002.md#canonical-c9e199a0e79fdfc4d4d76c5e2da178ea9d4909db78064b7fdd838aeb56eb9741) |
| `response_cookies_to_add.ignore_secure` | [response_cookies_to_add.ignore_secure](data-sources--virtual_host--reference--group-002.md#canonical-b20668cc5c116e9a8f819f6ee02c8f3f52a5556e29778696858fa93a4c8a994e) |
| `response_cookies_to_add.ignore_value` | [response_cookies_to_add.ignore_value](data-sources--virtual_host--reference--group-002.md#canonical-eadc1c8a157361ab2d134135565ab4f9a1cf303efaa0fa4f711d092c089d2c3d) |
| `response_cookies_to_add.max_age_value` | [response_cookies_to_add.max_age_value](data-sources--virtual_host--reference--group-002.md#canonical-6bfac342c9a702da62e016157179940809832e204d8d3b1b6641f2024b7c4a0c) |
| `response_cookies_to_add.name` | [response_cookies_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-0f93d81a288807991fbff83e22dfafaadecc8e404c08f1c090aac28b45789ec7) |
| `response_cookies_to_add.overwrite` | [response_cookies_to_add.overwrite](data-sources--virtual_host--reference--group-002.md#canonical-1d53b0e3f0d5cdba0424ecd2cf062970df72e7a2e374fefbb49aee2e3a41ee14) |
| `response_cookies_to_add.samesite_lax` | [response_cookies_to_add.samesite_lax](data-sources--virtual_host--reference--group-002.md#canonical-8f6665d32428f36ebf9519e3b83ff2250dc648a553546931270f3d4539fc8072) |
| `response_cookies_to_add.samesite_none` | [response_cookies_to_add.samesite_none](data-sources--virtual_host--reference--group-002.md#canonical-e8e123bb1258d25d91711149df2d072d31fa7cfd2e20b0ca2098dc20e25442c0) |
| `response_cookies_to_add.samesite_strict` | [response_cookies_to_add.samesite_strict](data-sources--virtual_host--reference--group-002.md#canonical-087a07a0b62cc9931167e5e530af13abddd1b61e4c8a1975cfafc4f04112d7ed) |
| `response_cookies_to_add.secret_value` | [response_cookies_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-2ca057e45985a0698098d6f385c06b7a2aad949748ba83aea794394def5630a6) |
| `response_cookies_to_add.secret_value.blindfold_secret_info` | [response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-19093bc3844ce31e8a3ab853fa57c213aa26ed286bbf25692ca2885a4a0bd762) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-065db219601c23acfcea759748e9ec121416bcfe971d30d6b0dbd67cefadf4c2) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.location` | [response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-326a745d61583c35108a2badf56536c6bbb59656902ec89c3f359b407eebc247) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-48da01f788c114232ee5bf78d95c6d206dd0234b3e633443165636811b0bb6a1) |
| `response_cookies_to_add.secret_value.clear_secret_info` | [response_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-f7a61bad0e5adc48004bd76fac09c27970f693dabd18642affd7d98567a441f2) |
| `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-9c410b6c15fe7c5aab1e431b24ad655e7e1ff17a3cac42a311e1eb0acf60bbbd) |
| `response_cookies_to_add.secret_value.clear_secret_info.url` | [response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-bac287308f8c015a9581a10cbd104827ab61ebfad306700be94aabbf26b644a2) |
| `response_cookies_to_add.value` | [response_cookies_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-1e9a9fa0163187ad338fe527297e6786d6aabe1cf5159f7ab0cf1f392e18e517) |
| `response_cookies_to_remove` | [response_cookies_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-8617633d0f2d1ae4771d64a479397b4a8363a20e7ed095581c299971fae384c0) |
| `response_headers_to_add` | [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-11881322be9bcd571d044f37f73bc11d8882cbc7f5d8624d981b5ed46b1ca462) |
| `response_headers_to_add.append` | [response_headers_to_add.append](data-sources--virtual_host--reference--group-002.md#canonical-f30f676fcc8389dbd06d4858dde7008bafd2e05784b5969af8f0bce5cc37a7a3) |
| `response_headers_to_add.name` | [response_headers_to_add.name](data-sources--virtual_host--reference--group-002.md#canonical-835c0dc5e4dfd8f6a6ef3faedd57e29f6d0988f388033bc8a44c1facfed9cbf4) |
| `response_headers_to_add.secret_value` | [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-d03a128181625d7bb378cd4b6bfda94f95fde6b90e53990e73477f7128ec1552) |
| `response_headers_to_add.secret_value.blindfold_secret_info` | [response_headers_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-386b75b7fd8f7280081c41be1fcdb60631afb5c37b8a2d5713592c88d5da4d6f) |
| `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-002.md#canonical-b9a23ffd647f88ae03e760dcc8452bd7bc8ac42d931e6b8fa803233c6f16b0fd) |
| `response_headers_to_add.secret_value.blindfold_secret_info.location` | [response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--virtual_host--reference--group-002.md#canonical-c0fa0232605aeb5695b7eefd27ad0834230706328e89c335614947964e6cbeb4) |
| `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-002.md#canonical-52d08406d242b8f5d41b8c841b40a8a60f3957da13de6c32e24386f391f502d0) |
| `response_headers_to_add.secret_value.clear_secret_info` | [response_headers_to_add.secret_value.clear_secret_info](data-sources--virtual_host--reference--group-002.md#canonical-dc289188951f310f7fe23d2757c0cc29c4c73dce1d00bd497516d40a0fa1cf6f) |
| `response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-002.md#canonical-dcd6c6b972da1cf0a1ffb013a98e77b069e2f82e223d3d215c3ee6e548da1287) |
| `response_headers_to_add.secret_value.clear_secret_info.url` | [response_headers_to_add.secret_value.clear_secret_info.url](data-sources--virtual_host--reference--group-002.md#canonical-c94c007a7e846c42bbfaf244eb21ea7af9fc29b60259362c0269258c805939fa) |
| `response_headers_to_add.value` | [response_headers_to_add.value](data-sources--virtual_host--reference--group-002.md#canonical-fadbfd5f0d1ea02eaf762a6750864c3df116f170c5744c17604f3edc245c39f3) |
| `response_headers_to_remove` | [response_headers_to_remove](data-sources--virtual_host--reference--group-001.md#canonical-0ffb469e9294b334fdde156a38b7328d2ea14f3faf5649fe75b0e54f4cd541e0) |
| `retry_policy` | [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-d9c0aee6ce40c260e86d51530ab8d2bba6f46075b3bf88617d4fa05b846d9051) |
| `retry_policy.back_off` | [retry_policy.back_off](data-sources--virtual_host--reference--group-002.md#canonical-90137b1f1dd9c158c985781b156e4912b77324a7f9b5cd9c9c92e0651e9f5b6f) |
| `retry_policy.back_off.base_interval` | [retry_policy.back_off.base_interval](data-sources--virtual_host--reference--group-002.md#canonical-a378be87197252a3eecd1438eec0aac7d9f5ba3889c3e863e1088f675c0cbd61) |
| `retry_policy.back_off.max_interval` | [retry_policy.back_off.max_interval](data-sources--virtual_host--reference--group-002.md#canonical-ac68f6a4f19a9e0e9129e1b937d73e308c1f52ad45dfc34b1ad23da891c32052) |
| `retry_policy.num_retries` | [retry_policy.num_retries](data-sources--virtual_host--reference--group-002.md#canonical-68daadb14e4d16961d60a4e2f7e551d42a5f58561c959fe0549a1f0d524f13de) |
| `retry_policy.per_try_timeout` | [retry_policy.per_try_timeout](data-sources--virtual_host--reference--group-002.md#canonical-3894889d54a7b1b23f4013773a2f2f1821abd0328c35a5017c0df1a2d9900cbb) |
| `retry_policy.retriable_status_codes` | [retry_policy.retriable_status_codes](data-sources--virtual_host--reference--group-002.md#canonical-9928b281c0798e6c8b32d624cb409569496fb93c8293b1f1d4c47acb1f2b6d50) |
| `retry_policy.retry_condition` | [retry_policy.retry_condition](data-sources--virtual_host--reference--group-002.md#canonical-7c75f6412705bde23f13c3b0a02dd161cebf3e523ad6ab8fc58a90bb55abd7b6) |
| `routes` | [routes](data-sources--virtual_host--reference--group-002.md#canonical-825bb61962da8698fa4c81ae326a69be2af94e3f6712315b000abdfd2fd05dbe) |
| `routes.kind` | [routes.kind](data-sources--virtual_host--reference--group-002.md#canonical-6b232fa48bd96afb2945ddf3cb824d058d3c9ba6462543821c3a6b08eaa2c083) |
| `routes.name` | [routes.name](data-sources--virtual_host--reference--group-002.md#canonical-a2c384541714a0ae946ef06a5d333077eec8a1c5ebbd45e48625f62f552ce378) |
| `routes.namespace` | [routes.namespace](data-sources--virtual_host--reference--group-002.md#canonical-245a46803a727013bfb5d6cc8cd5d667cdab7d6fdbea2ab4a55c6b945b327cda) |
| `routes.tenant` | [routes.tenant](data-sources--virtual_host--reference--group-002.md#canonical-f8ce44d497bbb69aacec032256df3fe040ad350767792a4d9f4bc73a4d58f3e6) |
| `routes.uid` | [routes.uid](data-sources--virtual_host--reference--group-002.md#canonical-dbf4e4162c3876182e186c9ad869252efecf79988a20498f7aeb3b029b3ba2e8) |
| `sensitive_data_policy` | [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-971a796701c3056e00203ca23b35d9ea84378d278b5024c02f0bf4a704cd4840) |
| `sensitive_data_policy.kind` | [sensitive_data_policy.kind](data-sources--virtual_host--reference--group-003.md#canonical-5dbf72f59420fe0008d6be9d25abfeec0c60af0a64e7cce4ba2f579b49f0ad40) |
| `sensitive_data_policy.name` | [sensitive_data_policy.name](data-sources--virtual_host--reference--group-003.md#canonical-fb7b80c2084e04296b9476031b45d45d1257016e48038fe83a2956ee2ff70edd) |
| `sensitive_data_policy.namespace` | [sensitive_data_policy.namespace](data-sources--virtual_host--reference--group-003.md#canonical-6c45eff92d531b61f5edb5fa2df448def4235449878c4d8f30a163fa8025bc70) |
| `sensitive_data_policy.tenant` | [sensitive_data_policy.tenant](data-sources--virtual_host--reference--group-003.md#canonical-4b9a7fffafdd83bab5b7a77a97494f05785d7480cdcc6222aea2f49d5d33f6e5) |
| `sensitive_data_policy.uid` | [sensitive_data_policy.uid](data-sources--virtual_host--reference--group-003.md#canonical-a688442abdee1715b7c57af8bbc89022588f56d7fcb64f984b7448c83cdbb706) |
| `server_name` | [server_name](data-sources--virtual_host--reference--group-001.md#canonical-807e60475c19e2d90c79afea33ef50293340f687899143ddbdfe578428d0af2b) |
| `slow_ddos_mitigation` | [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-68f153655816b3c39b8aa77961cbd8d47fbeac2360184d417bbcdbbd10f9b2d0) |
| `slow_ddos_mitigation.disable_request_timeout` | [slow_ddos_mitigation.disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-e25bd100b78830900ed44406d8e365fa994d0bb74482f537f6ec2f77bb1ef8da) |
| `slow_ddos_mitigation.request_headers_timeout` | [slow_ddos_mitigation.request_headers_timeout](data-sources--virtual_host--reference--group-003.md#canonical-a1280aa5af5593f1bd6d811e3b1e3b56922ba5d4b750cff40542466b24e0cbad) |
| `slow_ddos_mitigation.request_timeout` | [slow_ddos_mitigation.request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-ecca1f5ac11fccb4eb6c52c6d18aceccd93f4783c179fbffc56453c135e70776) |
| `tls_cert_params` | [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-ce341a9fe12b157fbb01cf6ce4594fc747e96462f75c5bc8f2287a8680c7daa4) |
| `tls_cert_params.certificates` | [tls_cert_params.certificates](data-sources--virtual_host--reference--group-003.md#canonical-e9bced0e6666f9101a112e5d0f573ce569b961b131447e930552ca5527efdf1e) |
| `tls_cert_params.certificates.kind` | [tls_cert_params.certificates.kind](data-sources--virtual_host--reference--group-003.md#canonical-bddbabc1d90178c5af129f75ba0a98699412ee43afc3b279f2657e4e7f9c5770) |
| `tls_cert_params.certificates.name` | [tls_cert_params.certificates.name](data-sources--virtual_host--reference--group-003.md#canonical-057e6ec799c269b85c842358b178bde456d97e5c5ef0efb747aabed5dbe2d5ab) |
| `tls_cert_params.certificates.namespace` | [tls_cert_params.certificates.namespace](data-sources--virtual_host--reference--group-003.md#canonical-84df9e8497e0537f22b3ee7036f1a7db43893f33c8d658274c736486cc5f38fb) |
| `tls_cert_params.certificates.tenant` | [tls_cert_params.certificates.tenant](data-sources--virtual_host--reference--group-003.md#canonical-b2ee7bda0217ee35c0ed7371decbc76ab5be4d187fecc593d70464bfc30e247e) |
| `tls_cert_params.certificates.uid` | [tls_cert_params.certificates.uid](data-sources--virtual_host--reference--group-003.md#canonical-f567841f5d2cbf50768f705ed18625cadbb1beeac5d71827632a5964c39de1ab) |
| `tls_cert_params.cipher_suites` | [tls_cert_params.cipher_suites](data-sources--virtual_host--reference--group-003.md#canonical-f01eb6c194893e186067c2c7fe05e275286ea79ea3ced0473d41eef07de3a08f) |
| `tls_cert_params.client_certificate_optional` | [tls_cert_params.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-09220efd5109ca70e867a121cbbf3c0a07ae02dd29409e854f07d2b99d550fca) |
| `tls_cert_params.client_certificate_required` | [tls_cert_params.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-74f54396c29bfa3e059f074e601cb1e8e626bdf1326ad1517921371987b972ef) |
| `tls_cert_params.maximum_protocol_version` | [tls_cert_params.maximum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-5d53c1bda99569b9d4df1a6dbc8e27b0fbc08dbf9a3c4ac3d6a7849cbb35ca25) |
| `tls_cert_params.minimum_protocol_version` | [tls_cert_params.minimum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-11303134574ecc97ec5cd0995012082e83206e722b9466e24199120ee1a920c1) |
| `tls_cert_params.no_client_certificate` | [tls_cert_params.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-24556e8bb2bfac0fa5850456f26f05a60ad35a525662b21f6eaac9db86589fdf) |
| `tls_cert_params.validation_params` | [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-b0caac7847a38a1a85587cbb0319f454a5aa20fac3b51ff179a6ebbc7c73545b) |
| `tls_cert_params.validation_params.skip_hostname_verification` | [tls_cert_params.validation_params.skip_hostname_verification](data-sources--virtual_host--reference--group-003.md#canonical-f86d8236bcb22b09ea359ad60fb742c61947f15deb97deb7c3fe9c18688a9476) |
| `tls_cert_params.validation_params.trusted_ca` | [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-98d9f9782710c8dba62e79fef1708cbeb04118ce820579e058be56441a5bb196) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-ba86bf454a89d1f506dccf57615105110045d5ccbf5a8d09c499a26800c29c71) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--virtual_host--reference--group-003.md#canonical-ba7e1ae225c0883fd61397314bce4f8ddc8dd2a5df6e0ee2710ed0e593517e30) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--virtual_host--reference--group-003.md#canonical-2355f29906826f79ff26c65b860c4c7e4f1f192722f21c0a2dc93d29686124a8) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--virtual_host--reference--group-003.md#canonical-0843e82997e48b07187255008aacdfa0d37ad74fb1f1b95f44c7e13048df583e) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--virtual_host--reference--group-003.md#canonical-e4369d49371818866851a9d1b94abb972b38025f42c460c8f5b934330a12cf90) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--virtual_host--reference--group-003.md#canonical-f1215088f8189e7d1d1c5d9042a28ea72675b40f3807cae56c0ac85360bc4c10) |
| `tls_cert_params.validation_params.trusted_ca_url` | [tls_cert_params.validation_params.trusted_ca_url](data-sources--virtual_host--reference--group-003.md#canonical-d9563e38fceedcc5cd63427e960825ad0c365799f62e18e62f077d5d05d744c6) |
| `tls_cert_params.validation_params.verify_subject_alt_names` | [tls_cert_params.validation_params.verify_subject_alt_names](data-sources--virtual_host--reference--group-003.md#canonical-22140d719d51b247de5b0fb3999008a269de993beb6c00379a1b97e1cce5cde6) |
| `tls_cert_params.xfcc_header_elements` | [tls_cert_params.xfcc_header_elements](data-sources--virtual_host--reference--group-003.md#canonical-f4c81200f6e84bf480e43b51fa87012ab0e583a6c6ce8949372a3ffb0126426d) |
| `tls_parameters` | [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-0732653dbaddc2627d7093bf3f40a2dab1f66fb00e4e404d4c12f855ad744991) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-e999250ea6fc4e9c54c4d9562fc16d0e685f91c0ebe654c249b098e5d0d45994) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-7a64a241651dd0dfaf6aef17a3667c03defafc35a50df93736203606880b338f) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-5abc749c896aca2bc2343b0cc8e7002a7a1225039ab72ee0a3a32dddbbd4180e) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--virtual_host--reference--group-003.md#canonical-244de8fa402dab3c63864993b1419b1729ca0a57a6990fb3095bd059f4c79ce5) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-f50844da477b463c357255916fffb2e78bc30a539b92fc387d950ecb1be64526) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--virtual_host--reference--group-003.md#canonical-845e8423da4825cf72c81ccc638fd6af5ad5f01f4fd40401321595c03afc31ac) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-da96377d543ffd34477bb3798b91b095b9eb8017f9eb7b01ba2499b2390097c1) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--virtual_host--reference--group-003.md#canonical-1c760bddd207ebd1d5d41d9709fffbfbe17f13a8ddc6c80667c44949f06e5b64) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-47435d5adffad679050ca1070a72037a479d4f50728ce7e39e0c8283a3a98ae5) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-3fef8b5dd6b64e4d518e814e4c5e43df0e2972e5f251f8734dfb008169c030b0) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--virtual_host--reference--group-003.md#canonical-569e9f474f7c9152ec16013f5ca1c06161cff65d143891569e59f048296e4318) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-4e008c4e680a333ac574d1ad89fc42ce7f86908298613b2aca2e1ea3a2a690b4) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-1e9fe1f0fc7938955b696183a8c9e1918ed5a6a5a62e0e8eb42e7dfff8c3af34) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-5aa778016f56a7bb92a2a3c9f1a175eeab16884c0849063a3d075af3f3892b9f) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--virtual_host--reference--group-003.md#canonical-bc193035f1e242174c71222d70e893084312e98342fed174b5b47e857a13d6b4) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--virtual_host--reference--group-003.md#canonical-1bce9aec142dcb9f5adca8b979b33b5f9d1cf90e50c71a5957c05ea75d2f629b) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--virtual_host--reference--group-003.md#canonical-ad935fb7562b59702096e1d4773be02965f6c50af6d62614ca8b443af787226c) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-12b21048f3cba1abea5ec4182bb46b60338d94455dd3b24ecc3ff150ff00e002) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--virtual_host--reference--group-003.md#canonical-fe6becafb08a33b7176b9a069a60ea36939a6aa1816d0bcd6ea88e1e3dc13991) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--virtual_host--reference--group-003.md#canonical-3bc9433f4766825b7165736fdd227b89d2d6f95cb596181e50378085e9fa2847) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-94cf26757c1211fd5fe547448f667a26c6089f3e1fd46b4aaf89dba4e3b3e9ba) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-a61a6478d604eedb8f8444249f4c3cf7cf283966277563a3aff654948e3d80f4) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--virtual_host--reference--group-003.md#canonical-8c5f8be7287f0945daa7d32d468870a460ac9cd2d39d96503b6e7f71226e0e54) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-14abafaaf994cf232a1e931cadbb6ed2633b74e233a3653c3c6ac8322c3d6a9a) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-4f008fa20b87948309e8790c19d1485520a1e1d7e500fe18541d8ab07942bb47) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--virtual_host--reference--group-003.md#canonical-e74a68c645ec9ef69fee63a65468b890e275661be8b31f39cfd9eb33df0a6078) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--virtual_host--reference--group-003.md#canonical-f9334ddc1b34097cc5748915163ac21c179b00ef6320935a518b3569ca6bad79) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--virtual_host--reference--group-003.md#canonical-fd6f36998d075f73ae18de5b4645ed02e182216f61553c6a0958cce07fe341e6) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--virtual_host--reference--group-003.md#canonical-986ba2b6bc1abba4391175b663f694aa293880f11e31fe816e00be1aa2d81702) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--virtual_host--reference--group-003.md#canonical-7a717aafe94881c2b5c5dbaaad166d169f2939c80aadc63948ac4d38273ff81b) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--virtual_host--reference--group-003.md#canonical-a989b1cf31aa6aac024fa1a56b910f1afe94b77430aee89c74d63b70b185eae4) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--virtual_host--reference--group-003.md#canonical-8cf871809f50f833b3396fa80d2eb7b18df4781ed9d5701cd77c2ec2ec1e88a0) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-8367fcf195d958ca78e1cbf5cf0281bf212b1d0f766f063ce1b6d564e7181160) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](data-sources--virtual_host--reference--group-003.md#canonical-e0f94dab45d6986874c32e92e76bdfbc7cb82cff042f5a54e90c98d14cf3311a) |
| `user_identification` | [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-fed06eb3b1e4290fe0d4c0a7d0f9d0b0c96a25617fadf6321587fcb3903c8054) |
| `user_identification.kind` | [user_identification.kind](data-sources--virtual_host--reference--group-003.md#canonical-653bfa9c573b55bd7a58b82fa9912c8e157a2b2f37d515d9ad591507946273b0) |
| `user_identification.name` | [user_identification.name](data-sources--virtual_host--reference--group-003.md#canonical-722c06a06f8279863408879de60e03f5538c29a906b28e77114a0bab9547b473) |
| `user_identification.namespace` | [user_identification.namespace](data-sources--virtual_host--reference--group-003.md#canonical-6adfb7d5868d541895c97a73dca6d9c60f0067483160410479f0cb1f28ff825a) |
| `user_identification.tenant` | [user_identification.tenant](data-sources--virtual_host--reference--group-003.md#canonical-5d35978ae0af9718d1ec5e98682cffefcc75397501b163072365cd49b5c32948) |
| `user_identification.uid` | [user_identification.uid](data-sources--virtual_host--reference--group-003.md#canonical-2fd2dccda99d135b0fd73c524bf0ceec86dc7725ba1ccd25be6823c8a809f05d) |
| `waf_type` | [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-267765b8c41d2fdf4390bb490a14e7eb3457f4303e4a407fd7b594c3070758f3) |
| `waf_type.app_firewall` | [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-c0aa95543043bb54e77148cad240e23472fa45b6cd36dabf241432591dde6aa3) |
| `waf_type.app_firewall.app_firewall` | [waf_type.app_firewall.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-cde32b93374d2c59f4d3a3287daaff7afbd96c181ef33ad0ec3d34d97271f010) |
| `waf_type.app_firewall.app_firewall.kind` | [waf_type.app_firewall.app_firewall.kind](data-sources--virtual_host--reference--group-003.md#canonical-9ac7701ae769e8bd33e3767447bfab5a8f18a07446d6bb455ae1c0a50b1978f0) |
| `waf_type.app_firewall.app_firewall.name` | [waf_type.app_firewall.app_firewall.name](data-sources--virtual_host--reference--group-003.md#canonical-a424a221510dd1334f1750cff75b57c8a6059949509c13b5ef17131856de9d57) |
| `waf_type.app_firewall.app_firewall.namespace` | [waf_type.app_firewall.app_firewall.namespace](data-sources--virtual_host--reference--group-003.md#canonical-cf0b77a9ae94dcc6d45f2cffc9589b14f9ed79132ba6507b3241f9e649da68df) |
| `waf_type.app_firewall.app_firewall.tenant` | [waf_type.app_firewall.app_firewall.tenant](data-sources--virtual_host--reference--group-003.md#canonical-fd74eca47b642fe16cd45ad3197deaeb1cd55f21313b2f5df3e135e0357a95df) |
| `waf_type.app_firewall.app_firewall.uid` | [waf_type.app_firewall.app_firewall.uid](data-sources--virtual_host--reference--group-003.md#canonical-479481c905e506fa2daa7dfd8961cfd5e5cd29b99fb0519e2218a4d046d30118) |
| `waf_type.disable_waf` | [waf_type.disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-0761c85d42142623abe0fe0be2662756e0a9533ddd0030a1bf5db82bb9c58375) |
| `waf_type.inherit_waf` | [waf_type.inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-c4c69772973a07336ba269301194566088878696c7e14de3a7a67fcf98280702) |

<a id="canonical-696858006792b4e79a0a875452f6440024957f4824e62c4075cf61dfde57988d"></a>

## Next pages — Property reference / 4ebaf40b9c3f / 27

- [advertise_policies](data-sources--virtual_host--reference--group-001.md#canonical-1021cfe9bb5566ddd87b255a7eccf4493464e67a5a1f176e3667d41f4f67d516)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [buffer_policy](data-sources--virtual_host--reference--group-001.md#canonical-67c6bb4057c626103a8169b0ef713a2bdb3c67563949f25de13828f8d848654f)
- [captcha_challenge](data-sources--virtual_host--reference--group-001.md#canonical-cce1978f17446fe7bf3d9c0dd2a62ce4ea2a4f0e13f8fd7f2ffff4c33710c445)
- [coalescing_options](data-sources--virtual_host--reference--group-002.md#canonical-75fd922d8b2a0030acf5ff4db0e633324030eaf21e03d907149614a4050d5705)
- [compression_params](data-sources--virtual_host--reference--group-002.md#canonical-175e663c8201caf980e21bdc50984a06a8b563d7d07c88dfceb91ced493567cd)
- [cors_policy](data-sources--virtual_host--reference--group-002.md#canonical-a4ecd779378b174be7fa310c13ecda6cf2402208e2a96236ae26a14b61393c5b)
- [csrf_policy](data-sources--virtual_host--reference--group-002.md#canonical-4ee7bac88ab5ea17c5299f4f1d7b2559166a21de99f33fd89fcfb5ddd96c3cc6)
- [default_header](data-sources--virtual_host--reference--group-002.md#canonical-2ba5db0f8ddd27b58c29095c0986c994d4d2c03100a06100429e44c738a20fe4)
- [default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-be86a1bd888037fdc3242c0d59516ab12221888e99a192b774b252f4409bcb2d)
- [disable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-fbdf7fd14ab032fa59e6f0edf908260deddc86d23ea5ac6cb183d78a9e97ca96)
- [dynamic_reverse_proxy](data-sources--virtual_host--reference--group-002.md#canonical-b8e8c7fcb7321f9bdf837de330392a7f4162a7567cf7f68cb61c12c9040bbf3e)
- [enable_path_normalize](data-sources--virtual_host--reference--group-002.md#canonical-1171379e2864a9edbff9631c19170289cd393aea92e23cbd8fbf34b9e78ce7e2)
- [http_protocol_options](data-sources--virtual_host--reference--group-002.md#canonical-4aabb4290babb3f663ded1911fa1af29fbe8c0f2babc078ecc811d26b751f04a)
- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-35d4e92967faff2a93e6f02288638f65dd0c817f65630593c95b78f48dc76894)
- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-6fd9adf9858ac044d27f69605c0da0afab5ae1521dd18568d820e2c4370a249b)
- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-d44525cc14415a80f6e329d02112bdf067ba82809fab4500bda9866ea672ad6f)
- [no_request_limit_per_connection](data-sources--virtual_host--reference--group-002.md#canonical-545aca29b311fea151c0f133136e8abac4d81e741b6298394ac83aad08e117f3)
- [non_default_loadbalancer](data-sources--virtual_host--reference--group-002.md#canonical-8b5b859bf066db8c50680f61439eb6236af261b0f24ab581d2fa842959be914d)
- [pass_through](data-sources--virtual_host--reference--group-002.md#canonical-62dbd01f9aad94f490c77cc7b4d61c54a4fe212c25d4796a116058d8f3256475)
- [rate_limiter_allowed_prefixes](data-sources--virtual_host--reference--group-002.md#canonical-5fd02a3e53a4014bceaf5d9ab95dd66a94fd084b4552ea6670553171736272ca)
- [request_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-c607794e1950663e659931bd4c846516c6d6cb4b36d3764bdaaf2c464573208a)
- [request_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-f39c7adb3f264ebfc0320b7642ce2fe8977c6b9d0ce455f5c6130b30dd0f9a17)
- [response_cookies_to_add](data-sources--virtual_host--reference--group-002.md#canonical-b2eab755b936b68b42c7358fbe998033541c0bb72b8037e7642c6abed92cbbca)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-5ac4cc3e0fb3d2c552bd80fbd0e1e0fa0c015a37f28e884f2a9f6e809b842597)
- [retry_policy](data-sources--virtual_host--reference--group-002.md#canonical-4934cf1290ecb6b5c3257959b6affcf382a8a9611d4070478d2b7a96a328bd49)
- [routes](data-sources--virtual_host--reference--group-002.md#canonical-e51606ec3c960457562b22f8d862e9f0eb68742c034a1b4373e9a7b926a9d694)
- [sensitive_data_policy](data-sources--virtual_host--reference--group-003.md#canonical-5570059452e7a42407d2b110e7733a91dd1e1471402407dd0b9bd2b2184c3e65)
- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-88cb41571444c3a26fbfd8e039fa894468f209f6f9fb8b8397263ebd3c25f6e5)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-b9382842870e321acdf7016c1135e4f94dee90a091f8f06bcfcf3b5ed7401575)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-59249dd40b384286e5ea9bb6fe59217d37819dc7d52c503f41d2d418369ea741)
- [user_identification](data-sources--virtual_host--reference--group-003.md#canonical-42caf6807b285559de11020a7412783936bce26b48dc1c3b083cd332528d2cb1)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-f7cff9cfeb2dd8157227422588b288998707fc74cf1aa6fbdb624a059e568bc1)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-1021cfe9bb5566ddd87b255a7eccf4493464e67a5a1f176e3667d41f4f67d516"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c1a3edb54d89d59d5c2893ac5142439aa98200d50ecffc30c278f9444a2655a"></a>

## advertise_policies — advertise_policies / 1cba8960753f / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- advertise_policies

<a id="canonical-05494ad126ed1d3d8ba1432041ec4189d78e08a936c599cded20b69ba316aa25"></a>

Type: `"list"`. Computed.

Advertise Policy allows you to define networks or sites where you want a VIP for this virtual host
to be advertised. Each Policy rule can have different parameters, like TLS configuration, ports,
optionally IP address to be used for VIP. If advertise policy is not specified then no VIP is..

Upstream description:

Advertise Policy allows you to define networks or sites where you want a VIP for this virtual host
to be advertised. Each Policy rule can have different parameters, like TLS configuration, ports,
optionally IP address to be used for VIP. If advertise policy is not specified then no VIP is
assigned for this virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-44bcc430853d108fd939726b75697d1589bd4ef39684c37550cb3b90ac673857"></a>

## Direct properties — advertise_policies / 1cba8960753f / 3

<a id="canonical-c7b24cfc89036f33952d2fd6110ca9692c01fba3b742a6a66bad6b7fe9759928"></a>

<a id="canonical-8f01f57abff0b13bfd18f56fe28c7d3e2f62b2483e97cc347de7ecb468b737d3"></a>

## kind property — advertise_policies / 1cba8960753f / 4

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

<a id="canonical-c2ccbc9a8a9e9ba0c51d2d9c344fd40c818dba19fbedc83d5c7e5dda5402e09e"></a>

<a id="canonical-287766bbf51e5fd45ea1356397c62323c2d101ceb12af244b9f78c4103bc9c21"></a>

## name property — advertise_policies / 1cba8960753f / 5

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

<a id="canonical-d9eb86e615062680ce1019603c57547768149d62cb983cdfdc0cb423040af98f"></a>

<a id="canonical-db94bb8e8c92a8187483094a58fb1bd6dcae0e5fdcef340fc486963fd3158c17"></a>

## namespace property — advertise_policies / 1cba8960753f / 6

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

<a id="canonical-f7223ce543c2df2c1540a8e008b1e6acde62ab3d3c49c63effcbe42fffbe8005"></a>

<a id="canonical-0fef7be371a1650252b7f00338bc047e32420330d697dd9425aff748a5e12f5f"></a>

## tenant property — advertise_policies / 1cba8960753f / 7

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

<a id="canonical-c96b2044c69182ac5db579740ac35f3ed0dc7b7a06d6f9d6c68aaf4c58fafa2b"></a>

<a id="canonical-86f1c00e066f62de43d8773fac7465b33be7881aeefd2a7509a122cfde0978ff"></a>

## uid property — advertise_policies / 1cba8960753f / 8

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

<a id="canonical-74f8baa25ab1c70644707abce95836883116522b0f339bfa48d5e32cc360e5a9"></a>

## Next pages — advertise_policies / 1cba8960753f / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3465097e7a2d44e465b69946926917152e22282f101cdbbc733cd1b6d750c16"></a>

## authentication — authentication / 620cc112c0d9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- authentication

<a id="canonical-545d089bac57d022dd93f2f3a9229a961d1ed6d51b8979e737d8f6507e2aa40e"></a>

Type: `"single"`. Computed.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Upstream description:

Authentication related information. This allows to configure the URL to redirect after the
authentication Authentication Object Reference, configuration of cookie params etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cookie_params_choice": "[\"cookie_params\",\"use_auth_object_config\"]",
  "x-ves-oneof-field-redirect_url_choice": "[\"redirect_dynamic\",\"redirect_url\"]"
}
```

OneOf alternatives in this subsection:

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-545d089bac57d022dd93f2f3a9229a961d1ed6d51b8979e737d8f6507e2aa40e)
- [no_authentication](data-sources--virtual_host--reference--group-002.md#canonical-524320b0777e1f50d78a30949d30b94597acbf8a2389b0c8271854f1eef75fd4)

Select alternatives according to the provider validators above.

<a id="canonical-32bf07c0069c94d030d8cd198a3d8b5074f1acce37918a3af31aa68673bf5754"></a>

## Direct properties — authentication / 620cc112c0d9 / 3

- [auth_config](data-sources--virtual_host--reference--group-001.md#canonical-5b58e434648d227b0a63d957dd1de7176ad07e2d7541d734f7b8e1ab42c2276a): complete subsection reference.

- [cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc): complete subsection reference.

- [redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-57b2f3389add8acf9fb9d3e81b7bbd989667b811052db1c73b8e54fdd71547b6): complete subsection reference.

<a id="canonical-09f89f9369443451dc474e3f055fa399b3dbf6e24b754f3aeefb127b5d3cc6b7"></a>

<a id="canonical-991f1a93f6b2a9ffb8709afb2e554a9250cccc548eaf6b264d80a501f01f7c2e"></a>

## redirect_url property — authentication / 620cc112c0d9 / 4

Type: `"string"`. Computed.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

Upstream description:

Exclusive with \[redirect\_dynamic\]

user can provide a URL for e.g https&#58;//abc.xyz.com where user gets redirected. This URL
configured here must match with the redirect URL configured with the OIDC provider.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-29c07ee4e405fecf854bf57f6d3797440042eb6dac65df7d70403ab75fa12c51): complete subsection reference.

<a id="canonical-1bb4940e49a9e9193c7195da62f654da286e44b6f03003125ec426501d3a2a37"></a>

## Next pages — authentication / 620cc112c0d9 / 5

- [authentication.auth_config](data-sources--virtual_host--reference--group-001.md#canonical-5b58e434648d227b0a63d957dd1de7176ad07e2d7541d734f7b8e1ab42c2276a)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.redirect_dynamic](data-sources--virtual_host--reference--group-001.md#canonical-57b2f3389add8acf9fb9d3e81b7bbd989667b811052db1c73b8e54fdd71547b6)
- [authentication.use_auth_object_config](data-sources--virtual_host--reference--group-001.md#canonical-29c07ee4e405fecf854bf57f6d3797440042eb6dac65df7d70403ab75fa12c51)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-5b58e434648d227b0a63d957dd1de7176ad07e2d7541d734f7b8e1ab42c2276a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27b2b6d0396c6052d61a10e8121c05a33d62560efb3fd621207022a3f41fb6e8"></a>

## authentication.auth_config — authentication.auth_config / 7c6ff81e9156 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- authentication.auth_config

<a id="canonical-432e29e671ba2a926c08990f17d6508c1241450e6e6b14a0682008e22d75d3c9"></a>

Type: `"list"`. Computed.

Reference to Authentication Config Object.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-f252c7368859e5d99e58140c7e0e6f8f4cbce4faaa63518a182a485fe38ab354"></a>

## Direct properties — authentication.auth_config / 7c6ff81e9156 / 3

<a id="canonical-2b283c1a1d34b26c560827e967be9b9d7427d7b6cbb1a2548e9fd5daf1e87e78"></a>

<a id="canonical-1d970d5142862babeed2f720761bb702c05d2ed01c9703ec74d2392310120f8c"></a>

## kind property — authentication.auth_config / 7c6ff81e9156 / 4

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

<a id="canonical-93216c981497701474c8bb072c2e7d2d840b948c81d14b9502535ced6ce18b1d"></a>

<a id="canonical-6f4210b0bb4ac725b9ed7d7a54e9aabfe646da757fb9416a58238c8e3450bf59"></a>

## name property — authentication.auth_config / 7c6ff81e9156 / 5

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

<a id="canonical-ebf7f8d4f9606c5a964019f9b31958b4b94af96ae491d855775592378db883ed"></a>

<a id="canonical-c03fb500bb36100ef417206cea9167076bc354198db2b78b9af69a8c1e0688b4"></a>

## namespace property — authentication.auth_config / 7c6ff81e9156 / 6

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

<a id="canonical-f7831c240b154a2a979c7e4b426e09e32b9beb61d27086d011855baad5d048df"></a>

<a id="canonical-477be2af5b436f29be2ce36b3099a2141726375deef53792471e16f8320a56e9"></a>

## tenant property — authentication.auth_config / 7c6ff81e9156 / 7

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

<a id="canonical-7340a4ba0928ae551f1808689d2893076110559fbf13b564a5afdf087b1412d7"></a>

<a id="canonical-75a4cec5b7c78b4827ea6174726cd04505ecd8f77c28c2a47b782b81cb0c14b8"></a>

## uid property — authentication.auth_config / 7c6ff81e9156 / 8

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

<a id="canonical-c6de63fb5b074ffed06391ad4e45360c852b0ce2dd2372e53dd3fe120cfa8cd6"></a>

## Next pages — authentication.auth_config / 7c6ff81e9156 / 9

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f373ca2ad46a7d8322d37f3d8ff5c17b0f3e93ba1d6a47ce576755a48dd48281"></a>

## authentication.cookie_params — authentication.cookie_params / 716df73e8960 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- authentication.cookie_params

<a id="canonical-3f8c10ebccfe81c63b92370a682e1acc15a154076d82be89c809465b66f4fb6f"></a>

Type: `"single"`. Computed.

Specifies different cookie related config parameters for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

<a id="canonical-5f14a638a05cd57519293fdba4d8e1a6667f5a492c1289a2f72d795a4845c99d"></a>

## Direct properties — authentication.cookie_params / 716df73e8960 / 3

- [auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a): complete subsection reference.

<a id="canonical-034b5f1e3779a60c7bc855c5470479941faf8df4dca3e3c48ccbc43283878f7b"></a>

<a id="canonical-1c7b242f12a561178cae1a4c62ae3cdf854b521e086a43b38a581410b230d033"></a>

## cookie_expiry property — authentication.cookie_params / 716df73e8960 / 4

Type: `"number"`. Computed.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0660651eedcf79a126ac2ab14aa6c68c010766b222fa49ee182c580ef4b1127b"></a>

<a id="canonical-c2c9a6cf9d665dfb2e6e9c70af432bc222abfffabcb998076f4b8295230f39ea"></a>

## cookie_refresh_interval property — authentication.cookie_params / 716df73e8960 / 5

Type: `"number"`. Computed.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-f42563210d7658ba9f7670114469f1729e981cc5275f4be7ba185cc2d8e499cf): complete subsection reference.

<a id="canonical-fb85439fef7c781a9158124c8819940106ffb7351b41919ce09f759b3664577b"></a>

<a id="canonical-8b5e64eb872f25d0c8323f71e0f9351e0387cd0dd11a5a5a248703a49e73de8e"></a>

## session_expiry property — authentication.cookie_params / 716df73e8960 / 6

Type: `"number"`. Computed.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-8a20b6b39162f23ee8a680ab2fc96d3c43b54a85841208ac4e7c6e31bdd9bb27"></a>

## Next pages — authentication.cookie_params / 716df73e8960 / 7

- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [authentication.cookie_params.kms_key_hmac](data-sources--virtual_host--reference--group-001.md#canonical-f42563210d7658ba9f7670114469f1729e981cc5275f4be7ba185cc2d8e499cf)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d559dc79837dcf82d59772696eedc8cff0937d5e1ae34f2a0c3770948a92ec3"></a>

## authentication.cookie_params.auth_hmac — authentication.cookie_params.auth_hmac / 4e0e65f580fd / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- authentication.cookie_params.auth_hmac

<a id="canonical-aef7595d637d8c179f4df76f8bdea5911b7189356015cd5feef2a633b30ec7ef"></a>

Type: `"single"`. Computed.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0e3b6732c3aa25b1c3dcab30748895b4f7b04cbe2f0c2964128afb20d17dc8fb"></a>

## Direct properties — authentication.cookie_params.auth_hmac / 4e0e65f580fd / 3

- [prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455): complete subsection reference.

<a id="canonical-f11adc1269cc8a3196f37f64f73aad6c6d3e58b7e903187ca59dee2ef73a2c54"></a>

<a id="canonical-6b8ba43153fe899845126feff80ad46b653173bc440c60747a0d0e71716624a4"></a>

## prim_key_expiry property — authentication.cookie_params.auth_hmac / 4e0e65f580fd / 4

Type: `"string"`. Computed.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5): complete subsection reference.

<a id="canonical-d573333500f65ae28a46a28eb4613068fc9474d56542f5087e48db6c71c67d6b"></a>

<a id="canonical-029c3192950fdacdb4582dcde4215f79ee96e67f131372235bef97cbcd002133"></a>

## sec_key_expiry property — authentication.cookie_params.auth_hmac / 4e0e65f580fd / 5

Type: `"string"`. Computed.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

<a id="canonical-6a796abdd443898d86ad4773208f402c0007f0cb6aa0acf6a9e8aaeff1b020d5"></a>

## Next pages — authentication.cookie_params.auth_hmac / 4e0e65f580fd / 6

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d7e05c771896134df45bf1830cb409cef8cacc3d58b996dd3a88cf6df826d4"></a>

## authentication.cookie_params.auth_hmac.prim_key — authentication.cookie_params.auth_hmac.prim_key / bc63b0fc002c / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- authentication.cookie_params.auth_hmac.prim_key

<a id="canonical-7bddd40bc4d2064bef3369cd7299ca145681cf66e50eea985f2870fb0661f4e4"></a>

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

<a id="canonical-2af4ee5da35ce3319e4d6060d531f1e6549316b70497bf617edcbc159b6b2c00"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key / bc63b0fc002c / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-65188cca755d37cc85bfcf877741ebfc8b4a66c3c9e5fd96812fe8097bf584f8): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-637df4183bb59e8688ee9282c5ec06b7b47fc362e95fea6112126e1fb781153b): complete subsection reference.

<a id="canonical-f023218be33086c522153c145be5b735f7fc2f653b283b360b03bf135f69e13a"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key / bc63b0fc002c / 4

- [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-65188cca755d37cc85bfcf877741ebfc8b4a66c3c9e5fd96812fe8097bf584f8)
- [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-637df4183bb59e8688ee9282c5ec06b7b47fc362e95fea6112126e1fb781153b)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-65188cca755d37cc85bfcf877741ebfc8b4a66c3c9e5fd96812fe8097bf584f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7364009b683a3885a3d839b7d230ce745eae2bbeee7b8ec57b20941af97a6ca"></a>

## authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455)
- authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-c0341d65c47e460f596a8091270d8400c47ae79bf344553fcf337852cfc5007c"></a>

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

<a id="canonical-91496a515ef45c2a25d3faa99093dff9b500012b0df1f4d52556a2f5f5845378"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 3

<a id="canonical-4c5f77d1a5dbd3b7c97d8ffa77201c71efdc5e953b5f03f78fffea3065503bd0"></a>

<a id="canonical-99c3f26248cc2612d26cb646fb62b51778ccaa8a9caf08c207d5d1c0018dda89"></a>

## decryption_provider property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 4

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

<a id="canonical-0188008038a59052d496e926585d0c6ce6461346ececa785c4643d4e74cb1431"></a>

<a id="canonical-74fbf90b9f46b667c09e2efb07caf4ee6e29b1dc959bf6822794baf2b0509a8c"></a>

## location property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 5

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

<a id="canonical-ca9343ebb91dbb37a78509a4b79b83270cea1648be48656d5de5b16f6c1c5ad8"></a>

<a id="canonical-6482224a1823e3b7a821c6bf59f17b8a53048aef60ac67cf30be7f403db73fae"></a>

## store_provider property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 6

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

<a id="canonical-358b6f40c1ec59afb98263ae70be7dcb5f8c04363aa79b968edba3ce1d4916ed"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / 93f806ff4da0 / 7

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-637df4183bb59e8688ee9282c5ec06b7b47fc362e95fea6112126e1fb781153b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3907db176996887087d2f639915afa00ff73e4b97bebbaa0003ddd6485ba822"></a>

## authentication.cookie_params.auth_hmac.prim_key.clear_secret_info — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / add3f8923ad2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455)
- authentication.cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-09355611c22c3b0f9ae0e9772b3994622e8db963d66e7a9b530ba5823f3497e5"></a>

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

<a id="canonical-2ff028252df5104e1324b6adb349b8e22342c9b69aace3ddb0383c3619d969b3"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / add3f8923ad2 / 3

<a id="canonical-3544944703b5e2476d775b2431e4c242c0f892bf95dd62cb1df8fc5eebae309b"></a>

<a id="canonical-ef462a072ece264d5c662e48d3b2d3a235acfb6600fa9baf2ab24231393804b1"></a>

## provider_ref property — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / add3f8923ad2 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c4236e2869ef013afb8a35f5ef58a0d1b371aae075ed16a82bf2a56643725288"></a>

<a id="canonical-ad234961972ed29222cc7f78203f447cf2792b320dd92fab72401c743893ae35"></a>

## url property — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / add3f8923ad2 / 5

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

<a id="canonical-603cca01d1ea5a4962eca5b9f5c1757e989f416a796b919d218e82861a8893b8"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / add3f8923ad2 / 6

- [authentication.cookie_params.auth_hmac.prim_key](data-sources--virtual_host--reference--group-001.md#canonical-9670ae74a92e2f20295477469450dce0aced7459a39b3c12ca5b850983247455)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a27668233683954523960e84175970ccbc7ea656d16eef1ce68de33b93fd3619"></a>

## authentication.cookie_params.auth_hmac.sec_key — authentication.cookie_params.auth_hmac.sec_key / 1d278aa5e1a4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- authentication.cookie_params.auth_hmac.sec_key

<a id="canonical-a992fac48d63bcf50835428af8cc39e519ab2ea96e2a6e10c5fed98d015092f7"></a>

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

<a id="canonical-53ca72427c1960d3fcb43cbb9f834d94e9f1b586d34122c429ec83dac5775b73"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key / 1d278aa5e1a4 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-53e75377d5e1b1b1974e520236515fdc51945e7e8051b170a935574857b54c6e): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-45fab0a63ff9a712ea355cc4f9aa8b88dd70d6ea15133eb9a434aa1e5b541c69): complete subsection reference.

<a id="canonical-db6bf1fe020da8e8061790e5a296af48deee4ef5e8045182c4c2861408bfc831"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key / 1d278aa5e1a4 / 4

- [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-53e75377d5e1b1b1974e520236515fdc51945e7e8051b170a935574857b54c6e)
- [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--virtual_host--reference--group-001.md#canonical-45fab0a63ff9a712ea355cc4f9aa8b88dd70d6ea15133eb9a434aa1e5b541c69)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-53e75377d5e1b1b1974e520236515fdc51945e7e8051b170a935574857b54c6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ecf6f75ba0448356e6ef7dd6d78d0f67e4ffc98eee3f8ce7aea5fe9845a8f0e"></a>

## authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5)
- authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-0ca8b20ed853e4aaa551092427fd6dd3f3aaceaa5201dcff7be414c15bd785d3"></a>

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

<a id="canonical-3aafa804efb84e1def70c73836093e031101ec5ac647ff7c59f8f001b337f1e6"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 3

<a id="canonical-21ad37a05736aeb27698a9e5937fccf037e0e9beddbadfecda8b5cce2e5557a7"></a>

<a id="canonical-de2fbeca7e74a8e48e3614b554f3f1b6b1cafe992d55da4817646222827f1fd9"></a>

## decryption_provider property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 4

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

<a id="canonical-c791032b62709107e5e5022a153d5e282e44f6d0b23e41e945bc803d58c3dab1"></a>

<a id="canonical-097c40568d42fd2f22a6d84a85c52d205b3b857c23a3046dddd0cdbf5137d257"></a>

## location property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 5

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

<a id="canonical-b2b7869e53882c67bf684e140ba6a2e4c3fa2764cc7349f3b9a95b6c54d70a9a"></a>

<a id="canonical-960bca5577814c00840844e2178b6d77936d54b9fc513000c0cd5529c22175cb"></a>

## store_provider property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 6

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

<a id="canonical-ffb150690a3f5aee8a7de0cb69dc480b07d3884bad75d2eaf5ed3e00f857e77e"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 52c436c8700d / 7

- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-45fab0a63ff9a712ea355cc4f9aa8b88dd70d6ea15133eb9a434aa1e5b541c69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1c18396622574ffe8feedadeb8fc47a310332a178db3a4933c93e83d058a1c1"></a>

## authentication.cookie_params.auth_hmac.sec_key.clear_secret_info — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / f73733acea9b / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [authentication.cookie_params.auth_hmac](data-sources--virtual_host--reference--group-001.md#canonical-7b5992f323c105bb33e7da94791585e13bf471f314445823393dfe343000771a)
- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5)
- authentication.cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-7d85ecda68dbb92d34f1a32fc48486b062ed2410061dd28f04a57bdbbb4b2742"></a>

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

<a id="canonical-b221090296f3c7a3d081903ab84e209d3f8b16f51d29a73c894a0c9bf63b75a2"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / f73733acea9b / 3

<a id="canonical-f4c899bdcd373efcc298575bb1247eb7606b139157da8a4a321178d74c2cc779"></a>

<a id="canonical-d68f99b1d3ceff2696cc0e2998de7539edc4d133e10b000c868deef0b00c9059"></a>

## provider_ref property — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / f73733acea9b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-38bfcbee7fc3b6a2e43a666f44f4d250aec9343b5c40b2e9ada4137bf50515f1"></a>

<a id="canonical-628c00fbb612a226ca284fa2c779ba17aec3d7d2adfc6cf136bc30ceff8b23f1"></a>

## url property — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / f73733acea9b / 5

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

<a id="canonical-b80f374ec07c13bdd7f1d7efb2cbedffc0d58ee82b08d45fdfb672a9b34a0f49"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / f73733acea9b / 6

- [authentication.cookie_params.auth_hmac.sec_key](data-sources--virtual_host--reference--group-001.md#canonical-a95d490bec55138c8d669bfe9f7924deade651c18883d3a22821ad6a6c9adcb5)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f42563210d7658ba9f7670114469f1729e981cc5275f4be7ba185cc2d8e499cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9374ba287cbfcc2710588cd6d204b11171faafb8f0b417f9964cbe93029bf80"></a>

## authentication.cookie_params.kms_key_hmac — authentication.cookie_params.kms_key_hmac / 9dbbccce7630 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- authentication.cookie_params.kms_key_hmac

<a id="canonical-0bf2c24828e319c09f6253cc7cc80a51008086f3085993f4365caed4edf773e1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for kms key hmac.

Upstream description:

Reference to KMS Key Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-75d46ab5a51bcc9314b8487eea7880be91982d71bd63525ce90f2c5a9938eac2"></a>

## Direct properties — authentication.cookie_params.kms_key_hmac / 9dbbccce7630 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdb244636a2a4cb1a69af9b0304f69f085a9d376c3a41fdf05036b6e6ed23902"></a>

## Next pages — authentication.cookie_params.kms_key_hmac / 9dbbccce7630 / 4

- [authentication.cookie_params](data-sources--virtual_host--reference--group-001.md#canonical-b30c0c69d5135acb5537cf079df5ea35a4bc20131c6054a5f335e8184bbb20bc)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-57b2f3389add8acf9fb9d3e81b7bbd989667b811052db1c73b8e54fdd71547b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12d7f97e3ac3e5d8cac2d5d14e3744629b4d23c7fa49eec41ba35ea003b813c4"></a>

## authentication.redirect_dynamic — authentication.redirect_dynamic / d517029ff771 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- authentication.redirect_dynamic

<a id="canonical-2355951834bb37b5c69f3115561b18f77c431cb74e70a148410180684b3358e1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for redirect dynamic.

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

<a id="canonical-8fb8852c5afe062ad2ac245cdbdcd2abe335abcc5b860691fc88ffa5b9dbc149"></a>

## Direct properties — authentication.redirect_dynamic / d517029ff771 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a5a5281bede34edc62a0262468947c658c4619b863699e88384e4f5f1e8a48c"></a>

## Next pages — authentication.redirect_dynamic / d517029ff771 / 4

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-29c07ee4e405fecf854bf57f6d3797440042eb6dac65df7d70403ab75fa12c51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f6ede201884a1efee29879af9f49ac1b0c7d1325d6ba7730a15f4a363706df2"></a>

## authentication.use_auth_object_config — authentication.use_auth_object_config / fa4afbb1c8ad / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- authentication.use_auth_object_config

<a id="canonical-e435948b4c0aa0bf6868c064699abbad12773800cf5acccab7a1a5403bc7b4e5"></a>

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

<a id="canonical-7bb1329933320f77b20c3f229d44fa8342e6b9a33b4cd1b9cd30eb18b011f6f3"></a>

## Direct properties — authentication.use_auth_object_config / fa4afbb1c8ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66c640065890cd2f1d3e987d8b138070ed72bf6327f27ba103977dc0b12dbeed"></a>

## Next pages — authentication.use_auth_object_config / fa4afbb1c8ad / 4

- [authentication](data-sources--virtual_host--reference--group-001.md#canonical-15987fed2fcc063c79e277fbf6c52578580c27be2b1ae0d023263fc385c26660)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-67c6bb4057c626103a8169b0ef713a2bdb3c67563949f25de13828f8d848654f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa406910fefb4febd8fc02d89257a75bd1afbe2b8c20b9c192f252946fbb9112"></a>

## buffer_policy — buffer_policy / 0cf3ebc277d9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- buffer_policy

<a id="canonical-571f6f41d81c6e65e1d841cceb9d5ecdc4ec0e5603c345cbadcdf3ca663f1e90"></a>

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

<a id="canonical-c92731f3b3b5df3d27a879c5a3fbb5019bcd76f1bce6063218adeddbe7b97e7e"></a>

## Direct properties — buffer_policy / 0cf3ebc277d9 / 3

<a id="canonical-accdec73b60de3448f2e425634ac3f50af8bbd97b3c198235a0f274d7f81a096"></a>

<a id="canonical-6fc596d6a0935e8bff6025302a9967dd4f7c58b576030e52d801cb0a298bf02f"></a>

## disabled property — buffer_policy / 0cf3ebc277d9 / 4

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

<a id="canonical-ae0dc59b0331ea3e2799c396aa7f986ca394b3e4272cfc0e8ec7f26134e4e437"></a>

<a id="canonical-49218be744b80aedaa13f4ed149b031a7bcdf77a714d46400cb3700a870c7e60"></a>

## max_request_bytes property — buffer_policy / 0cf3ebc277d9 / 5

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

<a id="canonical-baeea205218eb540071415818789c496086b95003ab2244afb35e5d97eb7c60e"></a>

## Next pages — buffer_policy / 0cf3ebc277d9 / 6

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-cce1978f17446fe7bf3d9c0dd2a62ce4ea2a4f0e13f8fd7f2ffff4c33710c445"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e06db092117f39f124c889fad017da2841130ef6b914495e1866ff95fbd3678a"></a>

## captcha_challenge — captcha_challenge / 4bbc8cbfcf5d / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- captcha_challenge

<a id="canonical-3bff95472b997489389cf2f7c2c5183a9a42f9ac1cdcfee69125a1f00eae7087"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, js\_challenge, no\_challenge; Default: no\_challenge\] Enables
loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With
this feature enabled, only clients that pass the captcha challenge will be allowed to complete the
HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

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

OneOf alternatives in this subsection:

- [captcha_challenge](data-sources--virtual_host--reference--group-001.md#canonical-3bff95472b997489389cf2f7c2c5183a9a42f9ac1cdcfee69125a1f00eae7087)
- [js_challenge](data-sources--virtual_host--reference--group-002.md#canonical-6fead0d412383114980e7657466e05e6825f8ea21f77f0cfff01ac7f11f7cf40)
- [no_challenge](data-sources--virtual_host--reference--group-002.md#canonical-3f0bdb5749658baa498a33d910c44024358aa16e77e1a62fe2a2b49bfe09053c)

Select alternatives according to the provider validators above.

<a id="canonical-ff11df76d38dddafd2c3f51a971e88e62b8ebd1451b75f6f2cb5ee7eb27e0d54"></a>

## Direct properties — captcha_challenge / 4bbc8cbfcf5d / 3

<a id="canonical-b738d882bb776eb6864073de272b09b745e5bfc608c0c9dbd1d771508d5f6d19"></a>

<a id="canonical-479318266bade4165b308a3cbde61f22bf3b079c66ad3004fab31f9b476de7d4"></a>

## cookie_expiry property — captcha_challenge / 4bbc8cbfcf5d / 4

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

<a id="canonical-4bcc08a2b4476c9745ad1f75c41a0270c7b6dbb917dd7b7706f7c943a3762939"></a>

<a id="canonical-f3b2dbb99eba8221e1be15fb15e22020082cbf9e28d8da218ade2b259f647558"></a>

## custom_page property — captcha_challenge / 4bbc8cbfcf5d / 5

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
