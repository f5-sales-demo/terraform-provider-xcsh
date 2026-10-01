---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42a00796c5a165ad5a6bf7d65772005cd4d3fb2d3816108cbc0a6017c0d05dde"></a>

## Property reference — Property reference / 240c0194682b / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- Property reference

<a id="canonical-b2114281b94f58f6646764dd483828600fac3a60e804cf6855e40a3939069195"></a>

## Direct properties — Property reference / 240c0194682b / 3

<a id="canonical-07efce9bf0840964de4baccafeca145a2a52fc456bedfff492990677370d458e"></a>

<a id="canonical-488067d0e0c52c1b8fe281055aa2013a40bff1694bf23540684e79d9c469442c"></a>

## add_location property — Property reference / 240c0194682b / 4

Type: `"bool"`. Optional, Computed.

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

- [advertise_policies](resources--virtual_host--reference--group-001.md#canonical-d1bcc00d2c60025e105523da9df31743f7ec554f5fb4411aade5da40ab299997): complete subsection reference.

<a id="canonical-94834ee1a498fb5f552bf5a2cf4a1b0ea2d0c3785d6793f7eb8e8bc2c37a1de5"></a>

<a id="canonical-76dc68a665abb4f5f3d6f628271c78deb1665defaccf65c6f762799e95868885"></a>

## annotations property — Property reference / 240c0194682b / 5

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

<a id="canonical-8e78ffa6b3a09b733b1c91ed9783e13f24921a2f700b922f03184e151acacabf"></a>

<a id="canonical-d864d2fa7bc32a5048c684df6eefeeef67ea8c0ea7ec3284919598b3e52f7bbc"></a>

## append_server_name property — Property reference / 240c0194682b / 6

Type: `"string"`. Optional, Computed.

\[OneOf: append\_server\_name, default\_header, pass\_through, server\_name; Default:
default\_header\] Exclusive with \[default\_header pass\_through server\_name\] Specifies the value
to be used for Server header if it is not already present. If Server Header is already present it is
not overwritten. It is just passed.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Specifies the value to be used for
Server header if it is not already present. If Server Header is already present it is not
overwritten. It is just passed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [append_server_name](resources--virtual_host--reference--group-001.md#canonical-8e78ffa6b3a09b733b1c91ed9783e13f24921a2f700b922f03184e151acacabf)
- [default_header](resources--virtual_host--reference--group-002.md#canonical-d3517bd1b5330302df34bc045eb406ab70a915b5dc5b9d063f922f1931a4be85)
- [pass_through](resources--virtual_host--reference--group-002.md#canonical-1e592fae1bf8eabc3123b454db28a7736ef8eb9433c7b31d1763e1e457d09fe2)
- [server_name](resources--virtual_host--reference--group-001.md#canonical-e2292fae9f1acc474189c8cd06484fc02dcb49585114f5a5f353100f701544b8)

Select alternatives according to the provider validators above.

- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb): complete subsection reference.

- [buffer_policy](resources--virtual_host--reference--group-001.md#canonical-ee1ac983e9b8523722832c7e02ecb4b2baa4b4e1c2aeaf702fb8ed19891b69e7): complete subsection reference.

- [captcha_challenge](resources--virtual_host--reference--group-002.md#canonical-48dad9d72b604417cdc3ed26a4f083f5a6772b0a7439caf68f30c071b60ac123): complete subsection reference.

- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475): complete subsection reference.

- [compression_params](resources--virtual_host--reference--group-002.md#canonical-4f3d9a57e11ef2bc8d3c145dad841bdfe6ea9138ada5d30a43f2452f7f3e31e9): complete subsection reference.

<a id="canonical-ced14c2a809898672b8f5f29084541eef360c341847de7007b3cd49ad686848e"></a>

<a id="canonical-3a9b9cee503480a6a0a47615fec58d24ca0a71d7e4d486f6d9aedcd8faa800cb"></a>

## connection_idle_timeout property — Property reference / 240c0194682b / 7

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [cors_policy](resources--virtual_host--reference--group-002.md#canonical-ee76b09d5033a064e534dbcf23f0888d145a396376a78774d87c166b278ecd71): complete subsection reference.

- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888): complete subsection reference.

<a id="canonical-3545bd1d17380e9ebc190d19404cbde7e874f769bb00dd377772d91a1d957037"></a>

<a id="canonical-96057f39a5008ff7c13b39087f299422565dee9c25b56cbffc747c0c926f50dd"></a>

## custom_errors property — Property reference / 240c0194682b / 8

Type: `["map", "string"]`. Optional.

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

- [default_header](resources--virtual_host--reference--group-002.md#canonical-7e7c09152bd364c144bc494f521c88bb9c5ea0302e8aaa518b288b68868c63b7): complete subsection reference.

- [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-d92a03f468c5ca768933d2e076c2e61f6d479736aac01f2973d5793ad0715c36): complete subsection reference.

<a id="canonical-0c0affbded9891c9c61da5f1a9ca7ac92c6b4a0f0d8bf2335c2d83b161104f89"></a>

<a id="canonical-9f113e69d16ce36ce4b3fded840693564f4e4154258b96ebcd8d7992f80b0944"></a>

## description property — Property reference / 240c0194682b / 9

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

<a id="canonical-83a994a0cb16fed912a1c874c0f0341aa3c5cf51f3b29d99f62196e3549808ae"></a>

<a id="canonical-42467b33bcf768e662f09f4fd0828f55546f23034b583e8cc40b2ab7bfefdcda"></a>

## disable property — Property reference / 240c0194682b / 10

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

<a id="canonical-41ead0402edc451222a27aff2806fd8fdbb2015bb1257a487cf51129885c1574"></a>

<a id="canonical-46785ee67a58ca5f9727c8aa53e283d32f802a04c1a5dbd01bb8db480f49d42b"></a>

## disable_default_error_pages property — Property reference / 240c0194682b / 11

Type: `"bool"`. Optional, Computed.

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

<a id="canonical-077328adb2d3460a5165d479d11999e935bf3c21ec8b48cb06bc5960de258f49"></a>

<a id="canonical-12038cc706951dd488644ddc67c69e4ddb76b559d9a62f6cc04422b7416a7656"></a>

## disable_dns_resolve property — Property reference / 240c0194682b / 12

Type: `"bool"`. Optional, Computed.

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

- [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-022ed9122ed6c53cad261165c0a803d0d9b37c5773659acf7af71b9e28331800): complete subsection reference.

<a id="canonical-e6035f0ffeb305e97931fd37b47d8b2c9715c1542635b4f8b042c71b80e72f8e"></a>

<a id="canonical-db64caa044ee86610fc95fc4adf2f2a23ba896eb18e3d481f6830a16b1fb3686"></a>

## domains property — Property reference / 240c0194682b / 13

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 33),
}
```

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

- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-9f367c54b8943301f0e415a6089876b335b94d6f9d41d982b7c5e774a110fde6): complete subsection reference.

- [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-6a5f867f80659e50963f8f154bef91391a4c75b4124b8d0d73e2fb3cb58da05c): complete subsection reference.

- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f): complete subsection reference.

<a id="canonical-906935dd709057c7be6d858d25c065cf42cf6b7ea5b7484e3863502a8d793f97"></a>

<a id="canonical-88d4d8a88c1a9b4555079c06ae829f60e5b28bf482d7837d29cd58947a9a3b79"></a>

## id property — Property reference / 240c0194682b / 14

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-9cd5024a2dddb22085c624112628bd4a94e558eddfce99c768078e4cd6b850c3"></a>

<a id="canonical-e62207719afd22385f48d8de7c7dd0f0a9185a354a0a7cf9d3b1398dd15c0835"></a>

## idle_timeout property — Property reference / 240c0194682b / 15

Type: `"number"`. Optional, Computed.

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

- [js_challenge](resources--virtual_host--reference--group-002.md#canonical-11b72bc185ad52ef017b9eee296b6915e0b52fc8f5a62246894d92061978e8db): complete subsection reference.

<a id="canonical-0c572205ef7b1735ab23e6180abd6bbb532156d693417e4c836960d723661658"></a>

<a id="canonical-21e79f3043d12030b92f48f074e48d06403aacd2810725ce61a4520f61d94040"></a>

## labels property — Property reference / 240c0194682b / 16

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

<a id="canonical-b1a6f2824f020b76f2c54707b81f7e72932a325efaf8e42c3008032c75d0455c"></a>

<a id="canonical-80192dfac77b3b4d801518937294aba4f3a8af03aee62c36ed748fc8b61c5fd3"></a>

## max_request_header_size property — Property reference / 240c0194682b / 17

Type: `"number"`. Optional, Computed.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

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

<a id="canonical-f28c438a366dbe65e5109b2bdd3f10583658247cfb418a975696f922f1d2145b"></a>

<a id="canonical-9c2122e761d282af85a633518388a19520b7019e5c774527448aeddc70f3e7c1"></a>

## max_requests_per_connection property — Property reference / 240c0194682b / 18

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests a downstream client can send over a single connection to Envoy. Enter
a value &gt;=1 to define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [max_requests_per_connection](resources--virtual_host--reference--group-001.md#canonical-f28c438a366dbe65e5109b2bdd3f10583658247cfb418a975696f922f1d2145b)
- [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-68bc0e78a1914b2fca2c9c77dda8ff8cc9ed1a56b881e907164718b5308a1697)

Select alternatives according to the provider validators above.

<a id="canonical-ee2879083db9899b69970e21caf275d64a8cfe677a9f7fd968bacc4691e95b05"></a>

<a id="canonical-765a3b99980668ce35e8d2d5541e660b1dd844dba59cafc568d97b32185753ff"></a>

## name property — Property reference / 240c0194682b / 19

Type: `"string"`. Required.

Name of the Virtual Host. Must be unique within the namespace.

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

<a id="canonical-bb420faf8549f7016a5cf75a9885a9d337b85bf711cf04dcb7080fa076650dcd"></a>

<a id="canonical-c50461f6bb64d3677d37e4e4bd0e542f0bac20bd7cbee68737ced400f0139f9f"></a>

## namespace property — Property reference / 240c0194682b / 20

Type: `"string"`. Required.

Namespace where the Virtual Host is created.

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

- [no_authentication](resources--virtual_host--reference--group-002.md#canonical-c7e0df30d5d7698fa774bb1754d169c8cc175e23854bc454545eb10dd939da9b): complete subsection reference.

- [no_challenge](resources--virtual_host--reference--group-002.md#canonical-0eb2f80ed7b092db4204ce9381a1c03fc429243b175fa8ba9e9858ce9eb533ac): complete subsection reference.

- [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-6126e97857379c252caebeaf8772e51c7056956152579ddb07a9300d64c4052c): complete subsection reference.

- [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-9359da92c433a78ae5abf5c6f1b8a7b7386baf714874b872ba69c74b64b62989): complete subsection reference.

- [pass_through](resources--virtual_host--reference--group-002.md#canonical-3b0a98bd07e1b03da8584b4df7e151fde171e721bb49956600c642a226337ed3): complete subsection reference.

<a id="canonical-2941bf0427748d501ed594ba4d88002ab3a45ae6eee246e6ed2948a9ebdafb29"></a>

<a id="canonical-ca58c6a1f1d1a067d42eeee16ebf5ecbac2e2e5e36f4de9c09af2c038d8b5432"></a>

## proxy property — Property reference / 240c0194682b / 21

Type: `"string"`. Optional, Computed.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"),
}
```

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

- [rate_limiter_allowed_prefixes](resources--virtual_host--reference--group-002.md#canonical-461ae2d80e6a10ff94c81ee92edf8b528dfc035a1510afb918deb5d3b92610a3): complete subsection reference.

- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536): complete subsection reference.

<a id="canonical-5ff26c26fd4d5249aa959cf75e28af6b0612af1c7bdff47cda0b6d96aa528422"></a>

<a id="canonical-220de0dfad6dc303edca33a6b6cc2ee06c09a396fdc6694b755f2474d9eb9bf0"></a>

## request_cookies_to_remove property — Property reference / 240c0194682b / 22

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969): complete subsection reference.

<a id="canonical-50cbe48cbbb0a0c37ca39dfeba097a890248a580612ddba30d314b3890a0fcf2"></a>

<a id="canonical-885f35dcce27d8250eee969176f0444e79780f682dc0311ce97dc6d3bff727e8"></a>

## request_headers_to_remove property — Property reference / 240c0194682b / 23

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37): complete subsection reference.

<a id="canonical-0928f4da62f20bbfe105725c26cbbf872d63f36c20c03fa79a282b875e8fff29"></a>

<a id="canonical-7a962f2765f241b5426a4ccb43a41257d4777575e444a69eabf9eb8eb04ca872"></a>

## response_cookies_to_remove property — Property reference / 240c0194682b / 24

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a): complete subsection reference.

<a id="canonical-bf2211e3e2dacb77325bf329482d576352e6a67dd0a86fd01b3d3b9d7cca0dc7"></a>

<a id="canonical-25b5b506a8988d7a805339a21055284a198549bd4973ea4cfd5a72fc3eb5981d"></a>

## response_headers_to_remove property — Property reference / 240c0194682b / 25

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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

- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-d8937146e8606cccb8a5d0c883571ec51d383d5cf570f87b7edaa3bf109ee3e0): complete subsection reference.

- [routes](resources--virtual_host--reference--group-003.md#canonical-6c636d254e1181e17ba1647911e7ca1e748a8724cdbd5ecd9bfbcfb08ca38ceb): complete subsection reference.

- [sensitive_data_policy](resources--virtual_host--reference--group-003.md#canonical-578cb808eb1ecb7ba626831d3422f85119ed4ee3daff27fd256fecfc5c29a224): complete subsection reference.

<a id="canonical-e2292fae9f1acc474189c8cd06484fc02dcb49585114f5a5f353100f701544b8"></a>

<a id="canonical-0eb158cd422971e687b75c10ced1ff5519030fc6f373f4102de905256f6c91f4"></a>

## server_name property — Property reference / 240c0194682b / 26

Type: `"string"`. Optional, Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-bfd8c69950f38f4e58ba7969dd63cf2da90dd42a6c9450294fe6276246fe52f1): complete subsection reference.

- [timeouts](resources--virtual_host--reference--group-003.md#canonical-9fd42ea62fb46c99d5c861f931e67aeb0b60a2b0c3ecb3f21d9fea3e4fdd59b4): complete subsection reference.

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d): complete subsection reference.

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb): complete subsection reference.

- [user_identification](resources--virtual_host--reference--group-003.md#canonical-46844342ae47b9c04a530b6e2c0292516326b045b1dd72188227e28335b02bb4): complete subsection reference.

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd): complete subsection reference.

<a id="canonical-a7f138d148efec9014b3e0f99daef1488e4d4711aeb06afbc2553fd73cb78a19"></a>

## All schema paths — Property reference / 240c0194682b / 27

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `add_location` | [add_location](resources--virtual_host--reference--group-001.md#canonical-07efce9bf0840964de4baccafeca145a2a52fc456bedfff492990677370d458e) |
| `advertise_policies` | [advertise_policies](resources--virtual_host--reference--group-001.md#canonical-380ff1975fc5252b8d9dbf70d533256f94562e35fca1c7521c9059ea3d691751) |
| `advertise_policies.kind` | [advertise_policies.kind](resources--virtual_host--reference--group-001.md#canonical-8e64ae095c2e0295c63be8e4fe03b74da62aa0dd6efe09136caf7c42e090373c) |
| `advertise_policies.name` | [advertise_policies.name](resources--virtual_host--reference--group-001.md#canonical-0f3bfca14b947d54d992fca86b48d50934b07af9f564f463da4364753e92cb1f) |
| `advertise_policies.namespace` | [advertise_policies.namespace](resources--virtual_host--reference--group-001.md#canonical-ffc5094c8ef6598c9c7620d9bbec4d20fba14685fb115b39918fd7e46a129cef) |
| `advertise_policies.tenant` | [advertise_policies.tenant](resources--virtual_host--reference--group-001.md#canonical-d92394205e8e69c42a4f7ebb3ccc12b24bfaca67ba3a82679f56e4d337587856) |
| `advertise_policies.uid` | [advertise_policies.uid](resources--virtual_host--reference--group-001.md#canonical-f3b97f913adf30a11b568f54a97849765fbd4c822551d45fe233063c5b2be924) |
| `annotations` | [annotations](resources--virtual_host--reference--group-001.md#canonical-94834ee1a498fb5f552bf5a2cf4a1b0ea2d0c3785d6793f7eb8e8bc2c37a1de5) |
| `append_server_name` | [append_server_name](resources--virtual_host--reference--group-001.md#canonical-8e78ffa6b3a09b733b1c91ed9783e13f24921a2f700b922f03184e151acacabf) |
| `authentication` | [authentication](resources--virtual_host--reference--group-001.md#canonical-658ce56de17f01f8402aff9a9ca5b36a1985ebf8cfbfbe7365ec37927e54bddf) |
| `authentication.auth_config` | [authentication.auth_config](resources--virtual_host--reference--group-001.md#canonical-eafe29027e09406ab9740ea809f191fa46e570a2612b6ff631fbe22530f7d5d5) |
| `authentication.auth_config.kind` | [authentication.auth_config.kind](resources--virtual_host--reference--group-001.md#canonical-ef9cce8c192227ef11424d3a248c6ad135686d1abccdcee43f5306002972e7c0) |
| `authentication.auth_config.name` | [authentication.auth_config.name](resources--virtual_host--reference--group-001.md#canonical-a4efbb9f15f0719de37338ec664195cb2f3a44c0f155c6ca85ec8aa5b12f64eb) |
| `authentication.auth_config.namespace` | [authentication.auth_config.namespace](resources--virtual_host--reference--group-001.md#canonical-bedbdd99a593cd5cc6c4342aa052a42b69c3d5450a426703fc0e23baf5be759a) |
| `authentication.auth_config.tenant` | [authentication.auth_config.tenant](resources--virtual_host--reference--group-001.md#canonical-d4be94707271f2753c011c3816e338a13f3be4d2a8e659efdd82ae7b5d6761a5) |
| `authentication.auth_config.uid` | [authentication.auth_config.uid](resources--virtual_host--reference--group-001.md#canonical-9e14a4924dd1bf31465a63bb3736b02b19ea96584606f19e722412cea43c0ec1) |
| `authentication.cookie_params` | [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-a0cb57d51278b2b47199879ce97c2da4f0e227bebd21faf043ed6e185ba272f9) |
| `authentication.cookie_params.auth_hmac` | [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-4ad2e6e16a1bc63736effb1a2a9a1bcce84d64fb36b9bf42e7d87d8268873458) |
| `authentication.cookie_params.auth_hmac.prim_key` | [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-27584f058c60720017e88c8d3dd807c8ad0d0f70adc51d18b39b53038811e1be) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-c4fc0e8f6b2f9b6e5d2cf61e5e955ac19424c58a4a01d73808a04f9fe18c6586) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-001.md#canonical-82fa58929299d538bc4469ce9e035ddcd484694cfaaa08ab0fb71868b814e13b) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](resources--virtual_host--reference--group-001.md#canonical-0ca18d60bd0953b18381778e4ad6e73854f008da55c90abdc3801ff308bb291d) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-001.md#canonical-d759d1194202d914e0181ead701ebf36f92e46f57645472aeb8a00949ae7a549) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-6069b0fe0e3adab51c633db7adc0e86fb1480248fac3bb6fe2e75a5aecab9b84) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-001.md#canonical-31710687698e18d487a14786e9d902a52be2a87aed1deac897f881c0942da7ef) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url](resources--virtual_host--reference--group-001.md#canonical-540aa98290233c6ea5b4221c52b63c39a224b1bc5a577bba3b20bb51163b89d6) |
| `authentication.cookie_params.auth_hmac.prim_key_expiry` | [authentication.cookie_params.auth_hmac.prim_key_expiry](resources--virtual_host--reference--group-001.md#canonical-7865276d83abeba1be884079a586493ca934d6cef9079933182a77a75b9dd21b) |
| `authentication.cookie_params.auth_hmac.sec_key` | [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-29643fb49f55b242d51f91ddd9ad1171a7f031d3f95a62496b1a8bedb39b035c) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-c944bda9bfd6772de8b6970a462f95342adc8839f1f8afceb3bae75a5997ed63) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-001.md#canonical-201f1342fa0443b8afd8a4b39ccdd51e74672bf2cfbb21c2ba191100f5dff82f) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](resources--virtual_host--reference--group-001.md#canonical-f56f20e4a227ed948ce40e1631d407523221d3fc5164320f1c9114010dbb5c0a) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-001.md#canonical-698d100a61e1ee4d251f0893efcbea7fdecbfa00bb1b26d30c362d63cd319dfc) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-1cae42bcdf6fdf941e5e4279763913a530a240a3df24b5231ca7aa7ae21e97d3) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-001.md#canonical-28d378abca79f4dcad02dc1d6088e065c95fd91f48707356f6713cc4f649baa9) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url](resources--virtual_host--reference--group-001.md#canonical-6614af718e1bf93451c518dead2786f01e1b97ddee1b5a10723fba21d33b7af5) |
| `authentication.cookie_params.auth_hmac.sec_key_expiry` | [authentication.cookie_params.auth_hmac.sec_key_expiry](resources--virtual_host--reference--group-001.md#canonical-c7de10567edbefeab264a5b03f710c5f1d64de54cf1c37ea7dcf7d4aa8d2cff2) |
| `authentication.cookie_params.cookie_expiry` | [authentication.cookie_params.cookie_expiry](resources--virtual_host--reference--group-001.md#canonical-fefb30bf00e37834097eefb1626f1517ff5184a8c2b19a713b1107fc0f0b44bf) |
| `authentication.cookie_params.cookie_refresh_interval` | [authentication.cookie_params.cookie_refresh_interval](resources--virtual_host--reference--group-001.md#canonical-40b914404c6ac63dcf63b3765a9bb90b5e74c5ed30a2a783179e43b00ec7c0b3) |
| `authentication.cookie_params.kms_key_hmac` | [authentication.cookie_params.kms_key_hmac](resources--virtual_host--reference--group-001.md#canonical-79fc59c686ff0893f7c1ddec8d2bf79040ee6f4759c3fdf7159eb241ecfbc415) |
| `authentication.cookie_params.session_expiry` | [authentication.cookie_params.session_expiry](resources--virtual_host--reference--group-001.md#canonical-66b9f214734410caf921221cee8c1f24194c8d540f58ac598890c0ffff2a3ea0) |
| `authentication.redirect_dynamic` | [authentication.redirect_dynamic](resources--virtual_host--reference--group-001.md#canonical-e7493e214e93445166bea6c1d86f970947a3981ba0cb2af4d52db273780c1b18) |
| `authentication.redirect_url` | [authentication.redirect_url](resources--virtual_host--reference--group-001.md#canonical-aa49d328f42b7e70ce6c06c467a95a83767773a3e58d23e7c62068144d416b36) |
| `authentication.use_auth_object_config` | [authentication.use_auth_object_config](resources--virtual_host--reference--group-001.md#canonical-24a8521af663fb5c1b3224b81d5801ee4154288c0d9050a9c562952d2ce4c852) |
| `buffer_policy` | [buffer_policy](resources--virtual_host--reference--group-001.md#canonical-c5ac6ba00f7569a6353a1ec46c7d50866400b6cd58eb91c5c4478dad09008610) |
| `buffer_policy.disabled` | [buffer_policy.disabled](resources--virtual_host--reference--group-001.md#canonical-598b22405a958d7a3b3327176613dc7d3ed3be8bad0ff344a435660092c4c8cb) |
| `buffer_policy.max_request_bytes` | [buffer_policy.max_request_bytes](resources--virtual_host--reference--group-002.md#canonical-42ee9331eb65cc1b3b998ace1d6addb88025195a0dafadfb4c7cccc3e3b2a2c6) |
| `captcha_challenge` | [captcha_challenge](resources--virtual_host--reference--group-002.md#canonical-71d9e7281db43657da616ad5db03176a7c3ca5be648c834df4a0a2ed18309125) |
| `captcha_challenge.cookie_expiry` | [captcha_challenge.cookie_expiry](resources--virtual_host--reference--group-002.md#canonical-38071eb47639120938cbce673ccebec97c492f3ce8c9c5997bfbc484fb9d0413) |
| `captcha_challenge.custom_page` | [captcha_challenge.custom_page](resources--virtual_host--reference--group-002.md#canonical-e52128758479e4ab8f53a7ba37d6f68ac097bf24872448529e61123315f1c7d9) |
| `coalescing_options` | [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-d864915ac5cf4d377aaf3beb4101ffdefe5e0f208a4c36701effe4725a78d3d0) |
| `coalescing_options.default_coalescing` | [coalescing_options.default_coalescing](resources--virtual_host--reference--group-002.md#canonical-8ad3ef578434d136f52b938bd7ae95e28af32ac19679b5cdce9f2c552d436147) |
| `coalescing_options.strict_coalescing` | [coalescing_options.strict_coalescing](resources--virtual_host--reference--group-002.md#canonical-9e9290b8e194d9a3cef4fecbffe7d40980ec006f6c1317f6b3d3c8e4576ce59d) |
| `compression_params` | [compression_params](resources--virtual_host--reference--group-002.md#canonical-0d94d5639ee92675b97a1654fbc25d72b4359d53e4ec3722b4c7103199612e47) |
| `compression_params.content_length` | [compression_params.content_length](resources--virtual_host--reference--group-002.md#canonical-31259163bbe50684367325193a5440b660e1f2228035ae988f85eef8b3c1b2b6) |
| `compression_params.content_type` | [compression_params.content_type](resources--virtual_host--reference--group-002.md#canonical-147b6e69543b2102b6648d0ab8477c4cca3870fc3de085f4605da45d2d6a6d20) |
| `compression_params.disable_on_etag_header` | [compression_params.disable_on_etag_header](resources--virtual_host--reference--group-002.md#canonical-d5cd8841e3c78fb240db05f47b561a050eb8468d31a310cbd41fa65156f086a7) |
| `compression_params.remove_accept_encoding_header` | [compression_params.remove_accept_encoding_header](resources--virtual_host--reference--group-002.md#canonical-191343755cae05f769dc766b2888a4ba53cf6d6dfba407f2078b9849264b7ad1) |
| `connection_idle_timeout` | [connection_idle_timeout](resources--virtual_host--reference--group-001.md#canonical-ced14c2a809898672b8f5f29084541eef360c341847de7007b3cd49ad686848e) |
| `cors_policy` | [cors_policy](resources--virtual_host--reference--group-002.md#canonical-5c8133a21bd3f08e4013d3be4abef1876bb8c13e9965c482a00eb143ee8dd5d0) |
| `cors_policy.allow_credentials` | [cors_policy.allow_credentials](resources--virtual_host--reference--group-002.md#canonical-ddbe7f41570b57c00464a13380bdb6168fd9c7dd0a8a32c581a7b8b1d5616704) |
| `cors_policy.allow_headers` | [cors_policy.allow_headers](resources--virtual_host--reference--group-002.md#canonical-1a1edd3426519cdac3026cc6c2226745bed3c895598e14c43170dd0dad9dc72c) |
| `cors_policy.allow_methods` | [cors_policy.allow_methods](resources--virtual_host--reference--group-002.md#canonical-91956e99b04d09e89565eac7e429f1dc34365e62e7b6639659cdf322f297a870) |
| `cors_policy.allow_origin` | [cors_policy.allow_origin](resources--virtual_host--reference--group-002.md#canonical-5c8add71fcc0d2145a46c20f821f4bd083172287d225c629e027cdb4c0ee76fd) |
| `cors_policy.allow_origin_regex` | [cors_policy.allow_origin_regex](resources--virtual_host--reference--group-002.md#canonical-4a184787fd0582c91fcf99337f7e16dd2f8f6b3acf28254e09dc39866f7cf7d0) |
| `cors_policy.disabled` | [cors_policy.disabled](resources--virtual_host--reference--group-002.md#canonical-69685da36e59ad69d20f7c96824e46b64fe8583a667a59ef71136cbf67a15d2c) |
| `cors_policy.expose_headers` | [cors_policy.expose_headers](resources--virtual_host--reference--group-002.md#canonical-05eb25856b3c56403cb72c2b8078a081de98fb580cc373aa07aff6f54eea232f) |
| `cors_policy.maximum_age` | [cors_policy.maximum_age](resources--virtual_host--reference--group-002.md#canonical-65abf21750092dbdd936357f9506d9fb337e69d21559fd43c28e0459f0c1a5b2) |
| `csrf_policy` | [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-c0f4d2fadd53dc83b8c4a38344b84c235385083f39df30a8cbe10e2c5848cd1f) |
| `csrf_policy.all_load_balancer_domains` | [csrf_policy.all_load_balancer_domains](resources--virtual_host--reference--group-002.md#canonical-df4a57069beb62bc0b5582386d94895cf8a88ca6e96bee2136043e2719fceb77) |
| `csrf_policy.custom_domain_list` | [csrf_policy.custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-caf55fa50d756aa2dd32bf3214040e60ee88799078c8a66ec1be7920a0ecd585) |
| `csrf_policy.custom_domain_list.domains` | [csrf_policy.custom_domain_list.domains](resources--virtual_host--reference--group-002.md#canonical-91c12231918375fd28e4def6ac0d269cf1055d8ccfef3c0f0dbefe4e399103df) |
| `csrf_policy.disabled` | [csrf_policy.disabled](resources--virtual_host--reference--group-002.md#canonical-d94d625805316035b26db52bebdd7b419db1b75d227bffe54a628a797bf74e1a) |
| `custom_errors` | [custom_errors](resources--virtual_host--reference--group-001.md#canonical-3545bd1d17380e9ebc190d19404cbde7e874f769bb00dd377772d91a1d957037) |
| `default_header` | [default_header](resources--virtual_host--reference--group-002.md#canonical-d3517bd1b5330302df34bc045eb406ab70a915b5dc5b9d063f922f1931a4be85) |
| `default_loadbalancer` | [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-c0992ecfc4ca73dd753aaf85d2d75ef3edbd5200cebeaf46de2730ebe369ad90) |
| `description` | [description](resources--virtual_host--reference--group-001.md#canonical-0c0affbded9891c9c61da5f1a9ca7ac92c6b4a0f0d8bf2335c2d83b161104f89) |
| `disable` | [disable](resources--virtual_host--reference--group-001.md#canonical-83a994a0cb16fed912a1c874c0f0341aa3c5cf51f3b29d99f62196e3549808ae) |
| `disable_default_error_pages` | [disable_default_error_pages](resources--virtual_host--reference--group-001.md#canonical-41ead0402edc451222a27aff2806fd8fdbb2015bb1257a487cf51129885c1574) |
| `disable_dns_resolve` | [disable_dns_resolve](resources--virtual_host--reference--group-001.md#canonical-077328adb2d3460a5165d479d11999e935bf3c21ec8b48cb06bc5960de258f49) |
| `disable_path_normalize` | [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-b15ce0b210431e03aef6140a111dcf3f345d3c8c612723ffd0e042a9a4f464c1) |
| `domains` | [domains](resources--virtual_host--reference--group-001.md#canonical-e6035f0ffeb305e97931fd37b47d8b2c9715c1542635b4f8b042c71b80e72f8e) |
| `dynamic_reverse_proxy` | [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-77c60c14a57d0a083ba8717e235f41760c6226ce7b7e10d77628bf437aa959d4) |
| `dynamic_reverse_proxy.connection_timeout` | [dynamic_reverse_proxy.connection_timeout](resources--virtual_host--reference--group-002.md#canonical-2b587aee8579a510bc999d7fd90ba6dc3a5d8647f21ce0d48e39b7ab918ea2db) |
| `dynamic_reverse_proxy.resolution_network` | [dynamic_reverse_proxy.resolution_network](resources--virtual_host--reference--group-002.md#canonical-d2d5cbe4c90baf6495ed73c750e4697a36433908c0c46a531b70caef2fdaca1f) |
| `dynamic_reverse_proxy.resolution_network.kind` | [dynamic_reverse_proxy.resolution_network.kind](resources--virtual_host--reference--group-002.md#canonical-2e6fb3eacde1afa59355252a4e695c6725798546484f52a468cca61a1d37c087) |
| `dynamic_reverse_proxy.resolution_network.name` | [dynamic_reverse_proxy.resolution_network.name](resources--virtual_host--reference--group-002.md#canonical-82bcf98ed82eb6b6f42dcafa2e8956f5b8b94c285ea276e068cd30d61f3c965d) |
| `dynamic_reverse_proxy.resolution_network.namespace` | [dynamic_reverse_proxy.resolution_network.namespace](resources--virtual_host--reference--group-002.md#canonical-0d1f5257a13ef9005ea2ded6e746f0349975b6acb94ecffeacc677697847ef3f) |
| `dynamic_reverse_proxy.resolution_network.tenant` | [dynamic_reverse_proxy.resolution_network.tenant](resources--virtual_host--reference--group-002.md#canonical-1cdeca587ad8ed3d4ed3f7f41f82dbdb3599ebac2880172108180dd7c6da60ad) |
| `dynamic_reverse_proxy.resolution_network.uid` | [dynamic_reverse_proxy.resolution_network.uid](resources--virtual_host--reference--group-002.md#canonical-05faa2a379dc0eeb81048a1293cfffff679e1760daadb6e26f59d1f942c84dbe) |
| `dynamic_reverse_proxy.resolution_network_type` | [dynamic_reverse_proxy.resolution_network_type](resources--virtual_host--reference--group-002.md#canonical-802d08c2b696e2c417a762cac226c5ebe3cff5822a2d204274f705510cfaad8f) |
| `dynamic_reverse_proxy.resolve_endpoint_dynamically` | [dynamic_reverse_proxy.resolve_endpoint_dynamically](resources--virtual_host--reference--group-002.md#canonical-19b40f1d73c08cf53287ddb6c2389ee9a85d6a20cd41ae9d7fec3d136a22e520) |
| `enable_path_normalize` | [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-e84c6ededb183fb94482189c0162c4daad2f0a2f21804b67269d03dc631168e0) |
| `http_protocol_options` | [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-19820efcc037e4a9286c48ed8abe9523063812393871c30e76b76c8d2f37b859) |
| `http_protocol_options.http_protocol_enable_v1_only` | [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-fcd8be56fcdfe99c187c5941d3d89850c0f9122dcf264ec41d2ecf477afa8c94) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-ab4361e73e8d0eac5b2b4868762a4ffdd9d136005a703ec656ab1f9cdb6c5d5b) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--virtual_host--reference--group-002.md#canonical-6558312a5e932236e09e58cbeb7abafbddd85fcd7b1a08bdc8bfa3b33998e7b2) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-166b46b61c666b70336218c2238cb13ab2e3be1b360fa4894c48eeb6c0f72671) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-ff414ac6113ba2ec323fa906e05581ea4c679c23abedfcd2631bdd6eda5c1a15) |
| `http_protocol_options.http_protocol_enable_v1_v2` | [http_protocol_options.http_protocol_enable_v1_v2](resources--virtual_host--reference--group-002.md#canonical-033da88ec6584e96dc4209572acd9e365ebfbaf6a2b36bdd0fafe82dabbde871) |
| `http_protocol_options.http_protocol_enable_v2_only` | [http_protocol_options.http_protocol_enable_v2_only](resources--virtual_host--reference--group-002.md#canonical-b02a6bcb3a929d143103ddbdc18f48db9fb8be6d816e4149519b72935bbd4501) |
| `id` | [id](resources--virtual_host--reference--group-001.md#canonical-906935dd709057c7be6d858d25c065cf42cf6b7ea5b7484e3863502a8d793f97) |
| `idle_timeout` | [idle_timeout](resources--virtual_host--reference--group-001.md#canonical-9cd5024a2dddb22085c624112628bd4a94e558eddfce99c768078e4cd6b850c3) |
| `js_challenge` | [js_challenge](resources--virtual_host--reference--group-002.md#canonical-b153e6f9a1401eb6f4bb963290887b7f4e9e32e3f5774d78b7d046370c7a950f) |
| `js_challenge.cookie_expiry` | [js_challenge.cookie_expiry](resources--virtual_host--reference--group-002.md#canonical-a0590ffcb3bc05e7e678d87826ce229298611ee183979c37feecb2ed50996365) |
| `js_challenge.custom_page` | [js_challenge.custom_page](resources--virtual_host--reference--group-002.md#canonical-8cb150343847436d3988fdaa3f62c17b92d4be7946505dead7b72159d853b3fe) |
| `js_challenge.js_script_delay` | [js_challenge.js_script_delay](resources--virtual_host--reference--group-002.md#canonical-1b1c102a9f717a6df099e1e0c8769936bb365fe2d7f7b8e6c22da473b0b4152a) |
| `labels` | [labels](resources--virtual_host--reference--group-001.md#canonical-0c572205ef7b1735ab23e6180abd6bbb532156d693417e4c836960d723661658) |
| `max_request_header_size` | [max_request_header_size](resources--virtual_host--reference--group-001.md#canonical-b1a6f2824f020b76f2c54707b81f7e72932a325efaf8e42c3008032c75d0455c) |
| `max_requests_per_connection` | [max_requests_per_connection](resources--virtual_host--reference--group-001.md#canonical-f28c438a366dbe65e5109b2bdd3f10583658247cfb418a975696f922f1d2145b) |
| `name` | [name](resources--virtual_host--reference--group-001.md#canonical-ee2879083db9899b69970e21caf275d64a8cfe677a9f7fd968bacc4691e95b05) |
| `namespace` | [namespace](resources--virtual_host--reference--group-001.md#canonical-bb420faf8549f7016a5cf75a9885a9d337b85bf711cf04dcb7080fa076650dcd) |
| `no_authentication` | [no_authentication](resources--virtual_host--reference--group-002.md#canonical-13b5e85003583a216e5ac801072c91108bedac112a998229fb56dbad837663b9) |
| `no_challenge` | [no_challenge](resources--virtual_host--reference--group-002.md#canonical-4b397b8c08ef76993ac481552be2d5bb414ac9624b3a878236aebb1b761f9b4c) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-68bc0e78a1914b2fca2c9c77dda8ff8cc9ed1a56b881e907164718b5308a1697) |
| `non_default_loadbalancer` | [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-23c848563d952a2c307ae4d364f943917eb3cd1e45b888fdafa93d79e384c51e) |
| `pass_through` | [pass_through](resources--virtual_host--reference--group-002.md#canonical-1e592fae1bf8eabc3123b454db28a7736ef8eb9433c7b31d1763e1e457d09fe2) |
| `proxy` | [proxy](resources--virtual_host--reference--group-001.md#canonical-2941bf0427748d501ed594ba4d88002ab3a45ae6eee246e6ed2948a9ebdafb29) |
| `rate_limiter_allowed_prefixes` | [rate_limiter_allowed_prefixes](resources--virtual_host--reference--group-002.md#canonical-f1765662fe555b07979b10d8899ee4e402238c3c10599976c4aeef2196b6f15c) |
| `rate_limiter_allowed_prefixes.kind` | [rate_limiter_allowed_prefixes.kind](resources--virtual_host--reference--group-002.md#canonical-5616dc12f04047834ddd68a88895f8da833678f6f319bdede20c0c2f3e5a3700) |
| `rate_limiter_allowed_prefixes.name` | [rate_limiter_allowed_prefixes.name](resources--virtual_host--reference--group-002.md#canonical-1e689d78cbbeab12414606f4c2fd9c7cad7917f20975a658272b3c8c4ab4efd3) |
| `rate_limiter_allowed_prefixes.namespace` | [rate_limiter_allowed_prefixes.namespace](resources--virtual_host--reference--group-002.md#canonical-9fd8c48820a3471aae19878c33bf95f60e2a3f64a68fb438c8dddf055e8300ab) |
| `rate_limiter_allowed_prefixes.tenant` | [rate_limiter_allowed_prefixes.tenant](resources--virtual_host--reference--group-002.md#canonical-d40b393f8f163db734f09becd8d60b03d24e394b07806193e13367a7fe9f90eb) |
| `rate_limiter_allowed_prefixes.uid` | [rate_limiter_allowed_prefixes.uid](resources--virtual_host--reference--group-002.md#canonical-6753081bd0052fc66204235eeddebd526d39939414c7998f7b5286b014642c6d) |
| `request_cookies_to_add` | [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-5e82268618b708b5588b1b8c1f2d06783748ec16733b7e4eef6fe6e89bea6047) |
| `request_cookies_to_add.name` | [request_cookies_to_add.name](resources--virtual_host--reference--group-002.md#canonical-d4ff4d3abef1b7c9b959b4f679415f969271d07436f092fb47a1507e5a911f5a) |
| `request_cookies_to_add.overwrite` | [request_cookies_to_add.overwrite](resources--virtual_host--reference--group-002.md#canonical-3e43fe65fd7ee86cd6a092b7fd9d2c29a8d0186e6daf4b67e54eb2a4c7080fc3) |
| `request_cookies_to_add.secret_value` | [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-c0ce90cc5d5691600e6e04c7ae7de1777fb7d72e4cc2adc3cbbc2e6027581891) |
| `request_cookies_to_add.secret_value.blindfold_secret_info` | [request_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-5e9957e314b0e154ec1991275db07326e717a1528175261bb2cbb9ad540979d5) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-7b099b1fe5061dd5602f729b29519fa14fb28f61cafc2390d140d19c6b93618c) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.location` | [request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-eabd1300b96a64a9aae3bdc82d4914927cf749247056a25ea02cf9bf18382ded) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-59ea29ff434ca95d663dfd114da039dbd72c0fa827b1e8b5807437d988d16016) |
| `request_cookies_to_add.secret_value.clear_secret_info` | [request_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-d57d1832d6298fdf28f616b9096bd11d1c07684795370a8be102fe7104271e7b) |
| `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-3f66e1f5884613f26139ee91502bafd84752fd094cb9dbd6d584051c70912665) |
| `request_cookies_to_add.secret_value.clear_secret_info.url` | [request_cookies_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-a3af7155310b2209f8f7256424208703894307a03fe65c6024f42a6fb3c589b7) |
| `request_cookies_to_add.value` | [request_cookies_to_add.value](resources--virtual_host--reference--group-002.md#canonical-76c2e476f9e31247b349340204511102e4afed3844a8434e7c82c73a8302c523) |
| `request_cookies_to_remove` | [request_cookies_to_remove](resources--virtual_host--reference--group-001.md#canonical-5ff26c26fd4d5249aa959cf75e28af6b0612af1c7bdff47cda0b6d96aa528422) |
| `request_headers_to_add` | [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-c6debe68d96130a8c25541339d089eb45c40f3203fab02a8b30cd77e6f140187) |
| `request_headers_to_add.append` | [request_headers_to_add.append](resources--virtual_host--reference--group-002.md#canonical-84ec30612d462565985f1d33f67d0441a1cc8a46be3c1267f28dbbea46a0f843) |
| `request_headers_to_add.name` | [request_headers_to_add.name](resources--virtual_host--reference--group-002.md#canonical-63bc2676e5e5c681e0c6bb5b14b289d8c72981e1aa9f0379d8c5d66b7a1d6710) |
| `request_headers_to_add.secret_value` | [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-eeb22f339c06c6d25e5c5a2e8bc62431bb67aa7b3c3a7c54309b3144a1b92045) |
| `request_headers_to_add.secret_value.blindfold_secret_info` | [request_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-b74dd2d7450ac25ea744a2d7099f532ed70b3ba62846f60e4ed54fcacc8a5a6f) |
| `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-6fc812e6baa6d97684aa4899e83c7eff79de118d962e17bf32459cd32f4a1fdd) |
| `request_headers_to_add.secret_value.blindfold_secret_info.location` | [request_headers_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-f4cb2969084904d7f369f4c7eca62ddb58584bb4f9ff9bf32f7f86a2bfdb4f56) |
| `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-d294f7b1107fadb78bd641a0040181fe65388d4426cfa421833035f09acb019c) |
| `request_headers_to_add.secret_value.clear_secret_info` | [request_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-69991c6616646bbf8fc11db60a87e62d16af64982fcba7f66bd9d602fedce56c) |
| `request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-9609fd713b1f5bc7df53bc3a2674d86b6f2a5047a06a4338da046d624be75d8c) |
| `request_headers_to_add.secret_value.clear_secret_info.url` | [request_headers_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-bb96a57db04b0732854be69825969a7a1e05130de8ddc31b217818f61c1a47e3) |
| `request_headers_to_add.value` | [request_headers_to_add.value](resources--virtual_host--reference--group-002.md#canonical-07bf99e8a686c89eee9505fe3a7dd14eafc3ae0979ff6110bf41382a49e9a78b) |
| `request_headers_to_remove` | [request_headers_to_remove](resources--virtual_host--reference--group-001.md#canonical-50cbe48cbbb0a0c37ca39dfeba097a890248a580612ddba30d314b3890a0fcf2) |
| `response_cookies_to_add` | [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-aa0c5c4ce2a8d71f70d5ee670298dc0e6e9d0a9055f85fc832dadb6074c33a3e) |
| `response_cookies_to_add.add_domain` | [response_cookies_to_add.add_domain](resources--virtual_host--reference--group-002.md#canonical-89159365e42f4cb3888c656529631413e644b8c99c49996dbaf7634acc4bcd95) |
| `response_cookies_to_add.add_expiry` | [response_cookies_to_add.add_expiry](resources--virtual_host--reference--group-002.md#canonical-7eaed9ee5683fbb5e4c48f80e89aa7ac84d6b8eb1873874eaf61456364e0f4b0) |
| `response_cookies_to_add.add_httponly` | [response_cookies_to_add.add_httponly](resources--virtual_host--reference--group-002.md#canonical-d892b334ccbf5819c9430b09defc854f032adda946d0f20a5fc42705188a4b66) |
| `response_cookies_to_add.add_partitioned` | [response_cookies_to_add.add_partitioned](resources--virtual_host--reference--group-002.md#canonical-c2ab97ea65fff1f9956b761bfe25ef4fc279a974b56c7b0397e40ea232844fbd) |
| `response_cookies_to_add.add_path` | [response_cookies_to_add.add_path](resources--virtual_host--reference--group-002.md#canonical-50388e25685403efaa30a217587632c849e191262c34713467d83989661e05e3) |
| `response_cookies_to_add.add_secure` | [response_cookies_to_add.add_secure](resources--virtual_host--reference--group-002.md#canonical-2992c6444c5299cd66a09ae3c9c79a87b3e1ceafb53fee9637215dca62751c1a) |
| `response_cookies_to_add.ignore_domain` | [response_cookies_to_add.ignore_domain](resources--virtual_host--reference--group-002.md#canonical-ede24d9c283557b33c44104f81af0dcd72f0a10f8a93134da64e936d683238eb) |
| `response_cookies_to_add.ignore_expiry` | [response_cookies_to_add.ignore_expiry](resources--virtual_host--reference--group-002.md#canonical-6ae290ffb6080bffea0a5a2ab718868015dad866f9a1485bed5286b96e6c37c4) |
| `response_cookies_to_add.ignore_httponly` | [response_cookies_to_add.ignore_httponly](resources--virtual_host--reference--group-002.md#canonical-8f21f8fc8671cb8a896161517e68d77c7e5ac7ffd5064c936ecc9338d4642250) |
| `response_cookies_to_add.ignore_max_age` | [response_cookies_to_add.ignore_max_age](resources--virtual_host--reference--group-002.md#canonical-be5d70c87a46d9981e2355179b01477acaddf93463bbc5932451f4cda67fbeb7) |
| `response_cookies_to_add.ignore_partitioned` | [response_cookies_to_add.ignore_partitioned](resources--virtual_host--reference--group-002.md#canonical-05938624aea34b5800d44af33c5678b2b9471a8138b51ffe5840071858d314a3) |
| `response_cookies_to_add.ignore_path` | [response_cookies_to_add.ignore_path](resources--virtual_host--reference--group-002.md#canonical-e42aa3c9a27f8f7c43c55add844c6697a0aaa0aedf91d4f62ce650a180b6c9dd) |
| `response_cookies_to_add.ignore_samesite` | [response_cookies_to_add.ignore_samesite](resources--virtual_host--reference--group-002.md#canonical-8523ed0a210f00b8656e526ef78c0bf28f91d20c9db29abe9349d4f7b91acb33) |
| `response_cookies_to_add.ignore_secure` | [response_cookies_to_add.ignore_secure](resources--virtual_host--reference--group-002.md#canonical-91e800ea46ca2bef411cd39e8099987c85a357cce20e840b5a7a70288b71f4b7) |
| `response_cookies_to_add.ignore_value` | [response_cookies_to_add.ignore_value](resources--virtual_host--reference--group-002.md#canonical-2708f0fe15a923e87570fa7ddeb3a90a1c452f119808ee22f915770dce57c5bb) |
| `response_cookies_to_add.max_age_value` | [response_cookies_to_add.max_age_value](resources--virtual_host--reference--group-002.md#canonical-c6c44dac6a7f74df16039b67c7374023d70ae7d4712285632e5d8044be151482) |
| `response_cookies_to_add.name` | [response_cookies_to_add.name](resources--virtual_host--reference--group-002.md#canonical-e93f86f2ff5200ec9c2314017b4f7bf4a08f8afdd7fce8bc9687f05d815f0cce) |
| `response_cookies_to_add.overwrite` | [response_cookies_to_add.overwrite](resources--virtual_host--reference--group-002.md#canonical-c72440169fd9c6fa42def3e007aa4ec1e9563f8be9429dcaf3b4c10935151bc7) |
| `response_cookies_to_add.samesite_lax` | [response_cookies_to_add.samesite_lax](resources--virtual_host--reference--group-002.md#canonical-1ebe99e52e822f314a41eef13da4f42f75b5d8da00c0a080b88456d1cbbe179c) |
| `response_cookies_to_add.samesite_none` | [response_cookies_to_add.samesite_none](resources--virtual_host--reference--group-002.md#canonical-dac3de5390f74f3f61793c4d75ababef3acc4445e0d094630fa832086ac58e14) |
| `response_cookies_to_add.samesite_strict` | [response_cookies_to_add.samesite_strict](resources--virtual_host--reference--group-002.md#canonical-30e992d9e32e0234b265716eeeb788a3ed0fc8545ebfdde2074318b010294de5) |
| `response_cookies_to_add.secret_value` | [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-24fcc1c2910a2f42857eae9f247b3f0a26e7fb674ce744db40c9b00852b5b391) |
| `response_cookies_to_add.secret_value.blindfold_secret_info` | [response_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-e1ca003435aadf076ddf311fb106f18473e48b57ca614b1b2537b990d2a21e1a) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-f47a80bdccda091e5e432c8e371935b6385cfc320830a6f57dfc927485934324) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.location` | [response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-421d9511254068b896c56b682aff46ace0137157f49e79c72815a74468c24e8d) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-0ef9571ff933c09a8e2c67cc29682aa4e1449c374d22eeff38a72fb5560f8670) |
| `response_cookies_to_add.secret_value.clear_secret_info` | [response_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-a1620388edd67c967155e0a822fe77e809662969ccc418974aaa08b6c307a369) |
| `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-382d490c1e6002c483be03ae4f3ff48eb2bb5154a1b6ed3d283d19754b369f33) |
| `response_cookies_to_add.secret_value.clear_secret_info.url` | [response_cookies_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-1a4992058e61cfbfa06def0f2831cb931cdf6115e8b54cabfddea2e85c77ff03) |
| `response_cookies_to_add.value` | [response_cookies_to_add.value](resources--virtual_host--reference--group-002.md#canonical-7c63a0cad8e5467295784f4c5cdd9cf3c4f1c041af21864cd72fa2f5b6ea0cd1) |
| `response_cookies_to_remove` | [response_cookies_to_remove](resources--virtual_host--reference--group-001.md#canonical-0928f4da62f20bbfe105725c26cbbf872d63f36c20c03fa79a282b875e8fff29) |
| `response_headers_to_add` | [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-7e8501d5a5681371c96a7242d35ba1796367e5eb6f1c4fafb7070fc1770c3898) |
| `response_headers_to_add.append` | [response_headers_to_add.append](resources--virtual_host--reference--group-002.md#canonical-42ac71d39262b08c5e0756803456d9850c7898d360cb4368f89dbfec8860e701) |
| `response_headers_to_add.name` | [response_headers_to_add.name](resources--virtual_host--reference--group-002.md#canonical-5423c1783154b48c0370012b810cb4df0038ff735e97f61ab4a6d941c70af895) |
| `response_headers_to_add.secret_value` | [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4aa6b5724a2aee4a5808c96c51c75035d59d292e9ebea79c6de66e36df9d3571) |
| `response_headers_to_add.secret_value.blindfold_secret_info` | [response_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-1c01161ffc5340df4c82d9523e86386cd21f2bfa66c171a0a78f971eae374448) |
| `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-002.md#canonical-5554f91a516a34322905c3b0d228afcabc67f7ffae92cfa3976e81f7f7a9bbb4) |
| `response_headers_to_add.secret_value.blindfold_secret_info.location` | [response_headers_to_add.secret_value.blindfold_secret_info.location](resources--virtual_host--reference--group-002.md#canonical-557003721f25ba88b8f9ec2cf7861322a987fd9ebb1b3ddfdfed9a6e6e3bed3c) |
| `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-002.md#canonical-0a040f717f38daa8f8d79dce0cdebf8dc3c31f305344d7585e56fa0216861db9) |
| `response_headers_to_add.secret_value.clear_secret_info` | [response_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-e1c0f35ed29f6b14e06697024b818f70ec1d5c5e8cf131a7310df6c01ed724f4) |
| `response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--virtual_host--reference--group-002.md#canonical-313221894d5879fe081073011e85b803271d1118c56892dd052cea3ec54c14a5) |
| `response_headers_to_add.secret_value.clear_secret_info.url` | [response_headers_to_add.secret_value.clear_secret_info.url](resources--virtual_host--reference--group-002.md#canonical-7849f670e34e8fe0c4f76a395ffdecd455f02fb9a82e7e0dc3d111f11cee1a2a) |
| `response_headers_to_add.value` | [response_headers_to_add.value](resources--virtual_host--reference--group-002.md#canonical-488bf7fe06b700c17393ac108c2859349116b43e3951a51610ca0ab79015c919) |
| `response_headers_to_remove` | [response_headers_to_remove](resources--virtual_host--reference--group-001.md#canonical-bf2211e3e2dacb77325bf329482d576352e6a67dd0a86fd01b3d3b9d7cca0dc7) |
| `retry_policy` | [retry_policy](resources--virtual_host--reference--group-003.md#canonical-8c83455fcc3830d8c2b0807cf1af96b9766bfde824f7032efb14f33c75704eea) |
| `retry_policy.back_off` | [retry_policy.back_off](resources--virtual_host--reference--group-003.md#canonical-8a89e21a37a2e132fba037408d32ffecdda2b3e105a86f9616e2c1528b743cec) |
| `retry_policy.back_off.base_interval` | [retry_policy.back_off.base_interval](resources--virtual_host--reference--group-003.md#canonical-00b2fde48849acecf981836a73288d1447dfbb757f0e7a48c1d4c7b1c5a8d93e) |
| `retry_policy.back_off.max_interval` | [retry_policy.back_off.max_interval](resources--virtual_host--reference--group-003.md#canonical-ed756f6c8a09ba24d7163c62b2f34ca17ff962fcb4ef01957e2e50d2ce9587fd) |
| `retry_policy.num_retries` | [retry_policy.num_retries](resources--virtual_host--reference--group-003.md#canonical-f946855524cefd20fa278e7bfcad851b463b36d3aaf3420c79375fe88d7269b4) |
| `retry_policy.per_try_timeout` | [retry_policy.per_try_timeout](resources--virtual_host--reference--group-003.md#canonical-3c61f242a030156d4413297c547d3f86a6843a6518f4ed844960977ec75ad4c0) |
| `retry_policy.retriable_status_codes` | [retry_policy.retriable_status_codes](resources--virtual_host--reference--group-003.md#canonical-93bce8999d10370e971e402a55534617241cdd40fbf791614cf983a0c38a1797) |
| `retry_policy.retry_condition` | [retry_policy.retry_condition](resources--virtual_host--reference--group-003.md#canonical-b7d27dc2efa392cc1f44c31134c2b1d05b9083741f7771affc83ad7c853544a6) |
| `routes` | [routes](resources--virtual_host--reference--group-003.md#canonical-41b5509dc4a9a734c09b499ebed9ff21f44dfeeddab239772d62ba9847107338) |
| `routes.kind` | [routes.kind](resources--virtual_host--reference--group-003.md#canonical-7eca30a37b7561b2e6739bddaa22db86168a2624034019faabeb5dc286803c42) |
| `routes.name` | [routes.name](resources--virtual_host--reference--group-003.md#canonical-ef811af33dfcd0a5eaac3e197ab79dc61932b9eb3a7a8eeac982375f58be3360) |
| `routes.namespace` | [routes.namespace](resources--virtual_host--reference--group-003.md#canonical-bfbd04cbf024f588cd60e5c81a52b6c70279b64bcf1ecc77e586b1dd287de7ea) |
| `routes.tenant` | [routes.tenant](resources--virtual_host--reference--group-003.md#canonical-13722e8427c7026bd607f6ef1396cf15794f9883312c4bb9733878818e324452) |
| `routes.uid` | [routes.uid](resources--virtual_host--reference--group-003.md#canonical-92fa30813beaf4607d0fedd4eb1b07ee67690e31889f49b3a0071d9e7547f274) |
| `sensitive_data_policy` | [sensitive_data_policy](resources--virtual_host--reference--group-003.md#canonical-2a6d7a8d87b4164e518fce36ae94769fe2d960040666c4879153b1fec7a6026d) |
| `sensitive_data_policy.kind` | [sensitive_data_policy.kind](resources--virtual_host--reference--group-003.md#canonical-1d019c85a59158e2e2591f0af7f1919a89369d04d5571d10141f4c2ceafe2e82) |
| `sensitive_data_policy.name` | [sensitive_data_policy.name](resources--virtual_host--reference--group-003.md#canonical-efeaeb8b2391cf3d338119409e587bdbf8b6019ba4a377d2ae37ee99d0b5eeee) |
| `sensitive_data_policy.namespace` | [sensitive_data_policy.namespace](resources--virtual_host--reference--group-003.md#canonical-d41abdb37770c5883d42e63bf9063c4320bea817bb7c49707805065bafdf3450) |
| `sensitive_data_policy.tenant` | [sensitive_data_policy.tenant](resources--virtual_host--reference--group-003.md#canonical-178693dfdcfa963f4e6fc8da9a44620ed388abc9641b220e22798d9fada2218d) |
| `sensitive_data_policy.uid` | [sensitive_data_policy.uid](resources--virtual_host--reference--group-003.md#canonical-94f79d1068cb6968b79e07fe73a531df61ee81e962248d9ea06d5c1c133146fc) |
| `server_name` | [server_name](resources--virtual_host--reference--group-001.md#canonical-e2292fae9f1acc474189c8cd06484fc02dcb49585114f5a5f353100f701544b8) |
| `slow_ddos_mitigation` | [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-27587dc92427d66a0d0fef3287eb4a83b206273d0bb9fa4abc44a8b59e2a1236) |
| `slow_ddos_mitigation.disable_request_timeout` | [slow_ddos_mitigation.disable_request_timeout](resources--virtual_host--reference--group-003.md#canonical-0e63684fe038d15a057d6cd14ff450f2eee6bfbecabaed086ebf41f9219697ff) |
| `slow_ddos_mitigation.request_headers_timeout` | [slow_ddos_mitigation.request_headers_timeout](resources--virtual_host--reference--group-003.md#canonical-34663318a254c795ec1eb2052e85fb96e4eb454248878667a57b26c3030d8196) |
| `slow_ddos_mitigation.request_timeout` | [slow_ddos_mitigation.request_timeout](resources--virtual_host--reference--group-003.md#canonical-44f6d2bd5ea4f845ec7d1aa987679b03cd0e325b74d02eee86c55b3ec52ed9d6) |
| `timeouts` | [timeouts](resources--virtual_host--reference--group-003.md#canonical-6bcc0454c557758d4246132541c3e88f3d4b56202f2fcea983fdc6a5f4469b24) |
| `timeouts.create` | [timeouts.create](resources--virtual_host--reference--group-003.md#canonical-57f890124a18538cb7a04f07ffd3600b5a5dcde84756919b0e0bc47cc672649e) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_host--reference--group-003.md#canonical-ccbef58964b5aaa10a93faf707bfbeda05e51e865892f785f30182c8b74efa67) |
| `timeouts.read` | [timeouts.read](resources--virtual_host--reference--group-003.md#canonical-724705d1ae1d8a4e9c0556c3452e7232356159bf6b9407ff240d4a2cec7a54cc) |
| `timeouts.update` | [timeouts.update](resources--virtual_host--reference--group-003.md#canonical-744cc76a7ea111414623d33a7b51d9636d0362127dece7a75f19590ac12e8b0d) |
| `tls_cert_params` | [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-9bf14422de2f877266665152d496cd59590f4c859819e3fb92fcda32939edafd) |
| `tls_cert_params.certificates` | [tls_cert_params.certificates](resources--virtual_host--reference--group-003.md#canonical-42f37b46de16d2b3bce96e08ffac4b28ac23b329f145127601a596e7ea8bc40a) |
| `tls_cert_params.certificates.kind` | [tls_cert_params.certificates.kind](resources--virtual_host--reference--group-003.md#canonical-8c5bd782d7d1e0cebdb90b0020b05ee70a4dcd0f53725a7783c6f8de30c6af1f) |
| `tls_cert_params.certificates.name` | [tls_cert_params.certificates.name](resources--virtual_host--reference--group-003.md#canonical-804375d171de4405a0daf69aa9e33efa123b13350cc9ef6cbe88ab6b162efff9) |
| `tls_cert_params.certificates.namespace` | [tls_cert_params.certificates.namespace](resources--virtual_host--reference--group-003.md#canonical-b88a2c93e341db6da23f430bfb748feea29e2ca6e4abcc40b8c04fcc4765a715) |
| `tls_cert_params.certificates.tenant` | [tls_cert_params.certificates.tenant](resources--virtual_host--reference--group-003.md#canonical-244f9036d1288cc22a39d0a631f33f381486b630bbbf3bc3694f6098605a846f) |
| `tls_cert_params.certificates.uid` | [tls_cert_params.certificates.uid](resources--virtual_host--reference--group-003.md#canonical-0f4ff91add3d278b93002e35a1ae6132830a3decb5fc91b4c421533ee6608c53) |
| `tls_cert_params.cipher_suites` | [tls_cert_params.cipher_suites](resources--virtual_host--reference--group-003.md#canonical-feba6908239bd7267067749969afa176ff290e459c0dc488fe23728b39cba97d) |
| `tls_cert_params.client_certificate_optional` | [tls_cert_params.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-6bcaf662a7d595683aa98e82bff36fc931472e007d22f2a131ba28cb7a3cb653) |
| `tls_cert_params.client_certificate_required` | [tls_cert_params.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-92106834e14776aacb3887b41b88aefe28938c122cac350cd0e42da030028a7d) |
| `tls_cert_params.maximum_protocol_version` | [tls_cert_params.maximum_protocol_version](resources--virtual_host--reference--group-003.md#canonical-e293788e9390ea27bfdf414bf800a45d0edcebab112d08dcb8fe25b103ab2ef0) |
| `tls_cert_params.minimum_protocol_version` | [tls_cert_params.minimum_protocol_version](resources--virtual_host--reference--group-003.md#canonical-d73b9f5f7b646366e7237047f250e0b1a7df7e992c2a12d1beecb92a876b0d44) |
| `tls_cert_params.no_client_certificate` | [tls_cert_params.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-f49c06c4de642692d3b1e616f142089df494175cbe753559041b4f3b5eee2593) |
| `tls_cert_params.validation_params` | [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-4e8b2c25c1ffc84f962c974999294c0daf6cfa1dce7d1a80ca65e2aebd9554bd) |
| `tls_cert_params.validation_params.skip_hostname_verification` | [tls_cert_params.validation_params.skip_hostname_verification](resources--virtual_host--reference--group-003.md#canonical-14f9e5639f989f3ce643fe230a9d98c82f003e9ce551064e3662c72e7ba9642f) |
| `tls_cert_params.validation_params.trusted_ca` | [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-423dd4682577d152ea99df8ccc5a04cc99f052117f8da5183aec2ddd5d222e52) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-554594813952b6b9ac77a4683386b61fdbb2c23cee86ec0a010b9cbf429ffd83) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--virtual_host--reference--group-003.md#canonical-7f58bea9f7411fdf42e3929328e8496ee0ced1fce1ff419a3d07d6d02e2573a2) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name](resources--virtual_host--reference--group-003.md#canonical-fb74d5ccd7912de8d3b8a89cf3302b29fb7e51c9dcac0307c5ae4f251b96d43b) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--virtual_host--reference--group-003.md#canonical-d931c1bca4f2537f5ed06b27535c6912a7382232246d9a4c94d7d03f42f7e045) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--virtual_host--reference--group-003.md#canonical-50846b2f27c6f631cfed959f4ba905ea0285370950dd7a1766868568633d20b5) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--virtual_host--reference--group-003.md#canonical-2f7cb6edc37b5ee8946112cffbcf471f583266756e81ba43597238e20490583a) |
| `tls_cert_params.validation_params.trusted_ca_url` | [tls_cert_params.validation_params.trusted_ca_url](resources--virtual_host--reference--group-003.md#canonical-12b8f30b17f86840764bbfc98d7e534df7237e3f70bbd332f21eeefb45624ed3) |
| `tls_cert_params.validation_params.verify_subject_alt_names` | [tls_cert_params.validation_params.verify_subject_alt_names](resources--virtual_host--reference--group-003.md#canonical-7cf787ddec29db407850fd42070a43a8e18246c5e447945dc8c1a21cc2c71421) |
| `tls_cert_params.xfcc_header_elements` | [tls_cert_params.xfcc_header_elements](resources--virtual_host--reference--group-003.md#canonical-e0bf5b230e74820603e05c3c9093d5188b2fe5e78762b4583777a518ca312e75) |
| `tls_parameters` | [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-2910e5ae02d5215342ecba466be0588450d1c2ee31bd899f0945f871a1b7bdff) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-3f655a1f3b234364ff15174de1671876e5948d7f1b4b8b10cec8bfc2c4029462) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-dcb8ece32e17faf90c5ddeb537118789410100f199939d9d87d6041d77cba500) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-3361f7ef1ea137fa4acadf36723f081e3135c28bc782367ccd89835c042d7719) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--virtual_host--reference--group-003.md#canonical-b4e707c32c0133936d0d88a7d702791a5d5fd054ea6a8ef566a878bf27d9e421) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--virtual_host--reference--group-003.md#canonical-dffbc60bba666e379dd58dcbaaa41b08e95ba833eda099c00a41a0debfffeed4) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--virtual_host--reference--group-003.md#canonical-278f4578e3c4a057236a88c0e0c9b52aede56db873ff36bbeed6c3b63bd47379) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-45eeded18167500449953058a3c3596e2db4f63c12f3bf8a41df81306ad9d260) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--virtual_host--reference--group-003.md#canonical-f295dbfd63d722f6326764ae8f868e830fc85b54a556f036155a505a8691bbc2) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-71f409a81bd3389c626aa47709ab2477fc7ef40ac74d3f59824397324de2119f) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-91918049f6c58ecb50442bd8faa73b7cf803578d7cadc4c1f2107960fd342ff5) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--virtual_host--reference--group-003.md#canonical-4513e9bb77db3b600f35a7bccff34bc9cc08406f9beefeca98a71f31baa26ae0) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--virtual_host--reference--group-003.md#canonical-11c8400dcb15d587af4cb7859697a86a3f30795256c569b17b7c32b66af9f0cf) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-b298967600db69cf12af6b236de25221726b249bdc3358b76461f3d2c653d667) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-ed45b6b86142e48b849add86415f897ee702baaa873520f19ddd2c2bcbe549cf) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--virtual_host--reference--group-003.md#canonical-ccb2ebb6f35915b8e6304d0e441d026ccd69a7cfa8fd5b4883691849279fed68) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--virtual_host--reference--group-003.md#canonical-d0f566b5e8baf55b39c7a7118207d5fea2c83837fa382134bace13cf9730f18b) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--virtual_host--reference--group-003.md#canonical-33144ceee90fbd34d45599d41bf2be5fbbde5d5a6ff7c6700efeb28c2e313d63) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-3122101c7343ae3dab6220750a904a61d8253ea3c262d72c827be90f0ca76685) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--virtual_host--reference--group-003.md#canonical-f713c46bc014cc1077e4b6356ca34a448555f5ee7cc90b901d94ffda1203973e) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--virtual_host--reference--group-003.md#canonical-4a87f94fd2eb5900fb4fb36267c88465bc4f34c122015777eba5efae4ded24a6) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-efcda848f285218a82b1e8894adf645bb85e26a26bcd1d5fbb921d261e7973a3) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-8d4ee58b7177a6f9fa13c59abf709eaf5ceedec1ff0c993fd3777804b4e516b4) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--virtual_host--reference--group-003.md#canonical-b0252ef02943c090ff9e0e658d6de362d7c445832b4bf845f3d52d1ca7f8a981) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-9209c3fbbc62abaf32cab4acb9680b3e82a68a142a21a92a349d703069d27363) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-c6f256f24a3eb76ed761f1c000a638dbe489cfbce41c3046a2b8777a6ea991da) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--virtual_host--reference--group-003.md#canonical-e60aea31f7102649835da76b95c34bc80aa33385ac295251180d77b19cac6a6f) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--virtual_host--reference--group-003.md#canonical-b75214e91d0164a77bb472704a987ce62106d5246d0fecb15d22011c86f79f10) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--virtual_host--reference--group-003.md#canonical-4b2ef68cfcbe3b7c1e413d57cd1170fd366108ca6135363c15a59c12b89c00af) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--virtual_host--reference--group-003.md#canonical-cd9036b164eb1ca226dc36a897d9ef00886c0ab35917552493b8cdc3c16932fa) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--virtual_host--reference--group-003.md#canonical-fb36d170a925e7190f03f3c5d0f125947e64c1f4d7b9bbc28bfb94d958e79660) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--virtual_host--reference--group-003.md#canonical-a1f010dc5317da2f2a78e2b42271461bf71f176e90185a34b573c3dc34c053c7) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--virtual_host--reference--group-003.md#canonical-ba32d9898f625ee930f3fde08cdff8b0114df8b49223b685d89b280a27657507) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-bb5b2d7c60df78ea3b36937fd521914f2fb7279df3ff501900130a7fa722a10d) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](resources--virtual_host--reference--group-003.md#canonical-2783469c11c3d084d2beafc39a7bc24ce1506901dac126020d1ad1a0b042ca40) |
| `user_identification` | [user_identification](resources--virtual_host--reference--group-003.md#canonical-dbb523abe25819e8bb81bae3d99b15d764b7f9255b96454aeb275d56294fe312) |
| `user_identification.kind` | [user_identification.kind](resources--virtual_host--reference--group-003.md#canonical-18a95d2a38fcc67452c08ded4c7313fbe9455318669ff9f0636c40c825343c0f) |
| `user_identification.name` | [user_identification.name](resources--virtual_host--reference--group-003.md#canonical-9869dac62bc714e6eb7487342ebb4c0ea909a3bbee9511562dad9985e5427753) |
| `user_identification.namespace` | [user_identification.namespace](resources--virtual_host--reference--group-003.md#canonical-23a40709550f16e153f7555295f94632d25ebf1d676ec2b376ffe0a8e3e0a7d2) |
| `user_identification.tenant` | [user_identification.tenant](resources--virtual_host--reference--group-003.md#canonical-51a9e06f95220dc2419041eb04d9c9611773537b3534dbaf9f1621efa4c0bd28) |
| `user_identification.uid` | [user_identification.uid](resources--virtual_host--reference--group-003.md#canonical-3a64ef21e12d15b69149f7b6fe9ae130a19a08af42aa42ce89419bd33e72e0d2) |
| `waf_type` | [waf_type](resources--virtual_host--reference--group-003.md#canonical-ec1c32c04f956ddacb473484f769acc06292a25d4df21e8433fb6559f2e68938) |
| `waf_type.app_firewall` | [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-e728666a02f1d08cf43ef8ff4ff2553c9814007ca124886f3e6c0e38a623cc7f) |
| `waf_type.app_firewall.app_firewall` | [waf_type.app_firewall.app_firewall](resources--virtual_host--reference--group-003.md#canonical-9baf36fd3935fcfd8aa4c52631f0f0f546fe3fdbec53120c6915260124879e7f) |
| `waf_type.app_firewall.app_firewall.kind` | [waf_type.app_firewall.app_firewall.kind](resources--virtual_host--reference--group-003.md#canonical-043861fd334adfe6e17dd10bdf049520a281f0381cb12d5d7baeb70091eb1ea0) |
| `waf_type.app_firewall.app_firewall.name` | [waf_type.app_firewall.app_firewall.name](resources--virtual_host--reference--group-003.md#canonical-6037f3460bca75d2599b4da3f4e3fb4a39517a848cc8863e73bacfc144ebb6d0) |
| `waf_type.app_firewall.app_firewall.namespace` | [waf_type.app_firewall.app_firewall.namespace](resources--virtual_host--reference--group-003.md#canonical-8d6d5e7dfee3b786eef87f6566f553e8039b8e8159d98f6d6a9aea39b7e6328e) |
| `waf_type.app_firewall.app_firewall.tenant` | [waf_type.app_firewall.app_firewall.tenant](resources--virtual_host--reference--group-003.md#canonical-d1eaa2db01df304d0f8ff02803bc1008a6883d1076e222f361d968e21da83f99) |
| `waf_type.app_firewall.app_firewall.uid` | [waf_type.app_firewall.app_firewall.uid](resources--virtual_host--reference--group-003.md#canonical-5bce65512d6c3c884de73014da477c3a0c57e72a696fd1fd86628187fe750a42) |
| `waf_type.disable_waf` | [waf_type.disable_waf](resources--virtual_host--reference--group-003.md#canonical-32d2630735fdc0531a934161612f54962c3d67d8c19905addb2c9a7828926d24) |
| `waf_type.inherit_waf` | [waf_type.inherit_waf](resources--virtual_host--reference--group-003.md#canonical-05758b14baf08eb645c717336bea474329c00789729f0ea82de3b12db9a8873d) |

<a id="canonical-1a003b55da5064b7383874785fecd4a06944b43980f9b712584eefe3dcc4392a"></a>

## Next pages — Property reference / 240c0194682b / 28

- [advertise_policies](resources--virtual_host--reference--group-001.md#canonical-d1bcc00d2c60025e105523da9df31743f7ec554f5fb4411aade5da40ab299997)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [buffer_policy](resources--virtual_host--reference--group-001.md#canonical-ee1ac983e9b8523722832c7e02ecb4b2baa4b4e1c2aeaf702fb8ed19891b69e7)
- [captcha_challenge](resources--virtual_host--reference--group-002.md#canonical-48dad9d72b604417cdc3ed26a4f083f5a6772b0a7439caf68f30c071b60ac123)
- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475)
- [compression_params](resources--virtual_host--reference--group-002.md#canonical-4f3d9a57e11ef2bc8d3c145dad841bdfe6ea9138ada5d30a43f2452f7f3e31e9)
- [cors_policy](resources--virtual_host--reference--group-002.md#canonical-ee76b09d5033a064e534dbcf23f0888d145a396376a78774d87c166b278ecd71)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- [default_header](resources--virtual_host--reference--group-002.md#canonical-7e7c09152bd364c144bc494f521c88bb9c5ea0302e8aaa518b288b68868c63b7)
- [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-d92a03f468c5ca768933d2e076c2e61f6d479736aac01f2973d5793ad0715c36)
- [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-022ed9122ed6c53cad261165c0a803d0d9b37c5773659acf7af71b9e28331800)
- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-9f367c54b8943301f0e415a6089876b335b94d6f9d41d982b7c5e774a110fde6)
- [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-6a5f867f80659e50963f8f154bef91391a4c75b4124b8d0d73e2fb3cb58da05c)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [js_challenge](resources--virtual_host--reference--group-002.md#canonical-11b72bc185ad52ef017b9eee296b6915e0b52fc8f5a62246894d92061978e8db)
- [no_authentication](resources--virtual_host--reference--group-002.md#canonical-c7e0df30d5d7698fa774bb1754d169c8cc175e23854bc454545eb10dd939da9b)
- [no_challenge](resources--virtual_host--reference--group-002.md#canonical-0eb2f80ed7b092db4204ce9381a1c03fc429243b175fa8ba9e9858ce9eb533ac)
- [no_request_limit_per_connection](resources--virtual_host--reference--group-002.md#canonical-6126e97857379c252caebeaf8772e51c7056956152579ddb07a9300d64c4052c)
- [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-9359da92c433a78ae5abf5c6f1b8a7b7386baf714874b872ba69c74b64b62989)
- [pass_through](resources--virtual_host--reference--group-002.md#canonical-3b0a98bd07e1b03da8584b4df7e151fde171e721bb49956600c642a226337ed3)
- [rate_limiter_allowed_prefixes](resources--virtual_host--reference--group-002.md#canonical-461ae2d80e6a10ff94c81ee92edf8b528dfc035a1510afb918deb5d3b92610a3)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a)
- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-d8937146e8606cccb8a5d0c883571ec51d383d5cf570f87b7edaa3bf109ee3e0)
- [routes](resources--virtual_host--reference--group-003.md#canonical-6c636d254e1181e17ba1647911e7ca1e748a8724cdbd5ecd9bfbcfb08ca38ceb)
- [sensitive_data_policy](resources--virtual_host--reference--group-003.md#canonical-578cb808eb1ecb7ba626831d3422f85119ed4ee3daff27fd256fecfc5c29a224)
- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-bfd8c69950f38f4e58ba7969dd63cf2da90dd42a6c9450294fe6276246fe52f1)
- [timeouts](resources--virtual_host--reference--group-003.md#canonical-9fd42ea62fb46c99d5c861f931e67aeb0b60a2b0c3ecb3f21d9fea3e4fdd59b4)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-d491d27bfa51a6d195baf34dbf0b51a4b9ee43c2a9fb71fe8e124cf62c65a73d)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-3d9419c08a8022f83f841c549454835916486b5b937df67e68250c44d68452fb)
- [user_identification](resources--virtual_host--reference--group-003.md#canonical-46844342ae47b9c04a530b6e2c0292516326b045b1dd72188227e28335b02bb4)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-ba411891bd75569b20dc7ca5e038471e744c0a0b932ba6abb5c3108c4a98d2fd)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d1bcc00d2c60025e105523da9df31743f7ec554f5fb4411aade5da40ab299997"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f7b3cec4b37eea1ea5a2a2ffb005c627d124a0363cdf140a0838d5d7b577c8c"></a>

## advertise_policies — advertise_policies / f966f7e7a518 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- advertise_policies

<a id="canonical-380ff1975fc5252b8d9dbf70d533256f94562e35fca1c7521c9059ea3d691751"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
advertise_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6595ebbcdb1281334abab9b7d73073d72f2943a789f1c1a8ac2879c6fc190c61"></a>

## Direct properties — advertise_policies / f966f7e7a518 / 3

<a id="canonical-8e64ae095c2e0295c63be8e4fe03b74da62aa0dd6efe09136caf7c42e090373c"></a>

<a id="canonical-603a455b1703e736925c0c4dfa8ec4cf60b25f73ab6c5a03f06f265970f14b06"></a>

## kind property — advertise_policies / f966f7e7a518 / 4

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

<a id="canonical-0f3bfca14b947d54d992fca86b48d50934b07af9f564f463da4364753e92cb1f"></a>

<a id="canonical-286e153c62c89606907cba9fb50b6fb6a74a92330fad2c7464c263cd1a480e0a"></a>

## name property — advertise_policies / f966f7e7a518 / 5

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

<a id="canonical-ffc5094c8ef6598c9c7620d9bbec4d20fba14685fb115b39918fd7e46a129cef"></a>

<a id="canonical-3d21f49ab80752eeaf8efb812af8436cdaec86712bfab728f6afaf8d3bd4f628"></a>

## namespace property — advertise_policies / f966f7e7a518 / 6

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

<a id="canonical-d92394205e8e69c42a4f7ebb3ccc12b24bfaca67ba3a82679f56e4d337587856"></a>

<a id="canonical-d3119d39d6def5cee96005965eb831017b19305e222b9e3d1c269287e62e4eff"></a>

## tenant property — advertise_policies / f966f7e7a518 / 7

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

<a id="canonical-f3b97f913adf30a11b568f54a97849765fbd4c822551d45fe233063c5b2be924"></a>

<a id="canonical-f6970fbcb06b7ad7844fd2af315794c8c402ae739c3963293f179affd747c5f3"></a>

## uid property — advertise_policies / f966f7e7a518 / 8

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

<a id="canonical-4fdbdfc96202c3a519c10ecb65ebc692c3aadb89f6f90f29438be2ccbf6b35fb"></a>

## Next pages — advertise_policies / f966f7e7a518 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ce01cc20ca8db8a58824b4aafa1500806e303cc2aaf540f399109cad9b26ebb"></a>

## authentication — authentication / e6e4de2583c2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- authentication

<a id="canonical-658ce56de17f01f8402aff9a9ca5b36a1985ebf8cfbfbe7365ec37927e54bddf"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: authentication, no\_authentication; Default: no\_authentication\] Authentication related
information. This allows to configure the URL to redirect after the authentication Authentication
Object Reference, configuration of cookie params etc.

Upstream description:

Authentication related information. This allows to configure the URL to redirect after the
authentication Authentication Object Reference, configuration of cookie params etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_config"),
  validators.ConflictingObjectAttributes("cookie_params",
    "use_auth_object_config"),
  validators.ConflictingObjectAttributes("redirect_dynamic",
    "redirect_url")}
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
  "x-ves-oneof-field-cookie_params_choice": "[\"cookie_params\",\"use_auth_object_config\"]",
  "x-ves-oneof-field-redirect_url_choice": "[\"redirect_dynamic\",\"redirect_url\"]"
}
```

OneOf alternatives in this subsection:

- [authentication](resources--virtual_host--reference--group-001.md#canonical-658ce56de17f01f8402aff9a9ca5b36a1985ebf8cfbfbe7365ec37927e54bddf)
- [no_authentication](resources--virtual_host--reference--group-002.md#canonical-13b5e85003583a216e5ac801072c91108bedac112a998229fb56dbad837663b9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfb0c1e1fa1afc5e280a814357524ceb9f53f4c2167f714d6feb0244a8456080"></a>

## Direct properties — authentication / e6e4de2583c2 / 3

- [auth_config](resources--virtual_host--reference--group-001.md#canonical-5d6802e53bfdbeb481d4b4a839f3494d78a6f2b4c7bd02be1b07b75cd69ff36b): complete subsection reference.

- [cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99): complete subsection reference.

- [redirect_dynamic](resources--virtual_host--reference--group-001.md#canonical-371532deb04933472eae5a8c9848338b94f5efb94727652e6087f5469e349ad0): complete subsection reference.

<a id="canonical-aa49d328f42b7e70ce6c06c467a95a83767773a3e58d23e7c62068144d416b36"></a>

<a id="canonical-7fc89b5da54db29c6f90dc207990c01d77fd1e37c244f6e37e372a04649801f9"></a>

## redirect_url property — authentication / e6e4de2583c2 / 4

Type: `"string"`. Optional.

Exclusive with \[redirect\_dynamic\] user can provide a URL for e.g https&#58;//abc.xyz.com where
user gets redirected. This URL configured here must match with the redirect URL configured with the
OIDC provider.

Upstream description:

Exclusive with \[redirect\_dynamic\]

user can provide a URL for e.g https&#58;//abc.xyz.com where user gets redirected. This URL
configured here must match with the redirect URL configured with the OIDC provider.

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

- [use_auth_object_config](resources--virtual_host--reference--group-001.md#canonical-72ea9f91c6ae61b428fe85a9d2864eeadcd46a215d93d86bc9456760c159b3f2): complete subsection reference.

<a id="canonical-62627a24422a7936aebfa47fe86bc366a3bb24e2001613fa806b8dbf7748dc27"></a>

## Next pages — authentication / e6e4de2583c2 / 5

- [authentication.auth_config](resources--virtual_host--reference--group-001.md#canonical-5d6802e53bfdbeb481d4b4a839f3494d78a6f2b4c7bd02be1b07b75cd69ff36b)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.redirect_dynamic](resources--virtual_host--reference--group-001.md#canonical-371532deb04933472eae5a8c9848338b94f5efb94727652e6087f5469e349ad0)
- [authentication.use_auth_object_config](resources--virtual_host--reference--group-001.md#canonical-72ea9f91c6ae61b428fe85a9d2864eeadcd46a215d93d86bc9456760c159b3f2)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-5d6802e53bfdbeb481d4b4a839f3494d78a6f2b4c7bd02be1b07b75cd69ff36b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13a7771342fe3a7f2bdabe4fda8f3f08b55d240e253d0c528ea20a53041db8f6"></a>

## authentication.auth_config — authentication.auth_config / 4b57f779df5e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- authentication.auth_config

<a id="canonical-eafe29027e09406ab9740ea809f191fa46e570a2612b6ff631fbe22530f7d5d5"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
auth_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d53e1784770f7ddb4c3e2b40bc3f25d6e80c236a97abe3d4f0afe119afc2e07"></a>

## Direct properties — authentication.auth_config / 4b57f779df5e / 3

<a id="canonical-ef9cce8c192227ef11424d3a248c6ad135686d1abccdcee43f5306002972e7c0"></a>

<a id="canonical-dfc361f8359ebec8c655e7e19d208a339b06dff1bc4897cc75ae1f900bac764a"></a>

## kind property — authentication.auth_config / 4b57f779df5e / 4

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

<a id="canonical-a4efbb9f15f0719de37338ec664195cb2f3a44c0f155c6ca85ec8aa5b12f64eb"></a>

<a id="canonical-89d7ce155dfd9995b0e977fc76b9ebd67938fe0d776e29ba9e7dc62003a7045b"></a>

## name property — authentication.auth_config / 4b57f779df5e / 5

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

<a id="canonical-bedbdd99a593cd5cc6c4342aa052a42b69c3d5450a426703fc0e23baf5be759a"></a>

<a id="canonical-6ce0107f25a3c07955c88f0ee8669bc15e64e9d18b27510ded27329bac797eff"></a>

## namespace property — authentication.auth_config / 4b57f779df5e / 6

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

<a id="canonical-d4be94707271f2753c011c3816e338a13f3be4d2a8e659efdd82ae7b5d6761a5"></a>

<a id="canonical-7981c7fd973ed0687aab5d300c721d6949164019df2749c76404998510489162"></a>

## tenant property — authentication.auth_config / 4b57f779df5e / 7

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

<a id="canonical-9e14a4924dd1bf31465a63bb3736b02b19ea96584606f19e722412cea43c0ec1"></a>

<a id="canonical-fff12228da7f072a88f1f8961999afcf6c7df2a20d687c6c2d74a5832f4c5daa"></a>

## uid property — authentication.auth_config / 4b57f779df5e / 8

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

<a id="canonical-a366acff645f0c406586aeeb101b90473fae432c69ed5291cd3d49cbc55a153d"></a>

## Next pages — authentication.auth_config / 4b57f779df5e / 9

- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5374dd8bfcfad335c5e333b13336c30bafaa5bf260779a8c1cd4d5ed7cea853b"></a>

## authentication.cookie_params — authentication.cookie_params / 4c8b61119059 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- authentication.cookie_params

<a id="canonical-a0cb57d51278b2b47199879ce97c2da4f0e227bebd21faf043ed6e185ba272f9"></a>

Type: `"object"`. single nested block, Optional.

Specifies different cookie related config parameters for authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_hmac",
    "kms_key_hmac")}
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
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

Terraform syntax:

```terraform
cookie_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2411edebbb6a626927d541cc56b71c38bdb20d97bfffe2d975c6d534268504c2"></a>

## Direct properties — authentication.cookie_params / 4c8b61119059 / 3

- [auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e): complete subsection reference.

<a id="canonical-fefb30bf00e37834097eefb1626f1517ff5184a8c2b19a713b1107fc0f0b44bf"></a>

<a id="canonical-566bf1b18f3db6193942f6136538f52f1b2e58c7e2fca6f866265af2182acfcc"></a>

## cookie_expiry property — authentication.cookie_params / 4c8b61119059 / 4

Type: `"number"`. Optional.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

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

<a id="canonical-40b914404c6ac63dcf63b3765a9bb90b5e74c5ed30a2a783179e43b00ec7c0b3"></a>

<a id="canonical-676c3735f0398aeebef3d4a31c1cbaf48d3bc604d75dd1103bbea3fc6a5108df"></a>

## cookie_refresh_interval property — authentication.cookie_params / 4c8b61119059 / 5

Type: `"number"`. Optional.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

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

- [kms_key_hmac](resources--virtual_host--reference--group-001.md#canonical-ac3c3e93bff2afa576f4477dd3018f35f22db62eee75ecc505fc1454e3e2185a): complete subsection reference.

<a id="canonical-66b9f214734410caf921221cee8c1f24194c8d540f58ac598890c0ffff2a3ea0"></a>

<a id="canonical-8c417a7ae646ff081bae22e09d850bd1965bfea9777ded777624e8074b5bc77e"></a>

## session_expiry property — authentication.cookie_params / 4c8b61119059 / 6

Type: `"number"`. Optional.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1296000),
}
```

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

<a id="canonical-d6c4f969740b0fcf93959fdac30b765b056533da59e0a3345880573dc140b5cd"></a>

## Next pages — authentication.cookie_params / 4c8b61119059 / 7

- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [authentication.cookie_params.kms_key_hmac](resources--virtual_host--reference--group-001.md#canonical-ac3c3e93bff2afa576f4477dd3018f35f22db62eee75ecc505fc1454e3e2185a)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4ff46bf3ecdbb3fe44a5457201228074452181709e44dbd98c5de9fc2ce6598"></a>

## authentication.cookie_params.auth_hmac — authentication.cookie_params.auth_hmac / 0983320f28a1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- authentication.cookie_params.auth_hmac

<a id="canonical-4ad2e6e16a1bc63736effb1a2a9a1bcce84d64fb36b9bf42e7d87d8268873458"></a>

Type: `"object"`. single nested block, Optional.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prim_key_expiry",
    "sec_key_expiry")}
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
auth_hmac {
  # Configure direct properties listed below.
}
```

<a id="canonical-8e665a1c0657649b8f0031b4d56b9196611fe090db4f2200c274c960c281c92c"></a>

## Direct properties — authentication.cookie_params.auth_hmac / 0983320f28a1 / 3

- [prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9): complete subsection reference.

<a id="canonical-7865276d83abeba1be884079a586493ca934d6cef9079933182a77a75b9dd21b"></a>

<a id="canonical-b724ee87da1f99d309806f8c5d94923904e7c5d60486bb38c651c5c27b1266de"></a>

## prim_key_expiry property — authentication.cookie_params.auth_hmac / 0983320f28a1 / 4

Type: `"string"`. Optional.

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

- [sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271): complete subsection reference.

<a id="canonical-c7de10567edbefeab264a5b03f710c5f1d64de54cf1c37ea7dcf7d4aa8d2cff2"></a>

<a id="canonical-07ac35bb42635f234c0a2026d2dda3a5e0f1726de958c13b7cddf4cde6d5087b"></a>

## sec_key_expiry property — authentication.cookie_params.auth_hmac / 0983320f28a1 / 5

Type: `"string"`. Optional.

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

<a id="canonical-f52c4b722f62776b074c932700a70af3c6ee0edf5b777045cdeeb23d32aba624"></a>

## Next pages — authentication.cookie_params.auth_hmac / 0983320f28a1 / 6

- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9)
- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fad991eff7f0459bce3558fdb340599882bd23374eb0d6b75114eca1fd00f8e9"></a>

## authentication.cookie_params.auth_hmac.prim_key — authentication.cookie_params.auth_hmac.prim_key / b7f063a92bfc / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- authentication.cookie_params.auth_hmac.prim_key

<a id="canonical-27584f058c60720017e88c8d3dd807c8ad0d0f70adc51d18b39b53038811e1be"></a>

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
prim_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1d0dcb5c6cbe0023b5871eb62170ee2266f2c1b0a93036b9929fa62bf8316862"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key / b7f063a92bfc / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-638f4c8b62a62afc93b74ea763ec3627a4da321496303e3f4f4ea6334b590530): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-768e1a3414eec4bef087451a0c6e2fc03cb379fc42b3dffb8f9847ec89361449): complete subsection reference.

<a id="canonical-9de9bb7cf2dc8be19991dfb9d7e4e222a03afd0ab9e85f808a2802845bce5fce"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key / b7f063a92bfc / 4

- [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-638f4c8b62a62afc93b74ea763ec3627a4da321496303e3f4f4ea6334b590530)
- [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-768e1a3414eec4bef087451a0c6e2fc03cb379fc42b3dffb8f9847ec89361449)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-638f4c8b62a62afc93b74ea763ec3627a4da321496303e3f4f4ea6334b590530"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4a4681e515f5e0eb0b67c13f541e15ce671afd8afbb36307258c25fb4f3c38e"></a>

## authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9)
- authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-c4fc0e8f6b2f9b6e5d2cf61e5e955ac19424c58a4a01d73808a04f9fe18c6586"></a>

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

<a id="canonical-2d2768909bc375a02a72fc9a0b7d710b6169c7106f54b9ccab65078c11e672fc"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 3

<a id="canonical-82fa58929299d538bc4469ce9e035ddcd484694cfaaa08ab0fb71868b814e13b"></a>

<a id="canonical-7966c89ae43642e40330a1e79a433feae6f7a6bb137956a1e48043afee4ca11d"></a>

## decryption_provider property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 4

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

<a id="canonical-0ca18d60bd0953b18381778e4ad6e73854f008da55c90abdc3801ff308bb291d"></a>

<a id="canonical-34191c3121fd73077b3073ccfed9b8b919a7b7513d056c33a87f1cddf8e8e933"></a>

## location property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 5

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

<a id="canonical-d759d1194202d914e0181ead701ebf36f92e46f57645472aeb8a00949ae7a549"></a>

<a id="canonical-cb7484519ba9e06937d398a6f8e9afd65d7146e2afd552849c8d1b1866a2264f"></a>

## store_provider property — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 6

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

<a id="canonical-f9d714811ec5ea20d62744f779c597ae662192ebf14ac1cb8cb103936aab9c11"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info / f5ca228b22f0 / 7

- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-768e1a3414eec4bef087451a0c6e2fc03cb379fc42b3dffb8f9847ec89361449"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d5fecc6c66fe7ce44708276b19f361a66da4749dd4bf95642ff19e150db8c2d"></a>

## authentication.cookie_params.auth_hmac.prim_key.clear_secret_info — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / ba1adbb68f4e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9)
- authentication.cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-6069b0fe0e3adab51c633db7adc0e86fb1480248fac3bb6fe2e75a5aecab9b84"></a>

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

<a id="canonical-93555143f11d854a20e43e4671afd4bb8b843e795f4d6f26e44498084bb31e44"></a>

## Direct properties — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / ba1adbb68f4e / 3

<a id="canonical-31710687698e18d487a14786e9d902a52be2a87aed1deac897f881c0942da7ef"></a>

<a id="canonical-b414da1a676be126f574eccc6600c309abcfb8380f77ca8aa247b78da8e69eca"></a>

## provider_ref property — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / ba1adbb68f4e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-540aa98290233c6ea5b4221c52b63c39a224b1bc5a577bba3b20bb51163b89d6"></a>

<a id="canonical-b58bab51026c62aed3ec3c35eb0f998b7884daa2e7a00df441021716645ed99d"></a>

## url property — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / ba1adbb68f4e / 5

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

<a id="canonical-2c41940681ee771dfed6bd0d86ccf3dcc0b094dbb39c882447ea5354f980e882"></a>

## Next pages — authentication.cookie_params.auth_hmac.prim_key.clear_secret_info / ba1adbb68f4e / 6

- [authentication.cookie_params.auth_hmac.prim_key](resources--virtual_host--reference--group-001.md#canonical-e0134be364f4cae34574d32100775eb2dec24e2b32c975a2c65765a7e9f047a9)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cd14d63a4bf8630d0b71f88c2850f53d1e42e0801c82adcd3cce2e3a195bccc"></a>

## authentication.cookie_params.auth_hmac.sec_key — authentication.cookie_params.auth_hmac.sec_key / c7e6c4156d10 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- authentication.cookie_params.auth_hmac.sec_key

<a id="canonical-29643fb49f55b242d51f91ddd9ad1171a7f031d3f95a62496b1a8bedb39b035c"></a>

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
sec_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-bd9e1157c494c8eedf35feeac5904ed212f41c55c59888bfdc24a058e234d6fc"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key / c7e6c4156d10 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-d7761c128b31d151dee5f3f6a5ce50d3afde313a562738c3d3f30e15fd01d6ea): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-bbae9d33eb8eeccce25e035f4feb9f9db103e4a81309e264f1ce7bc70aae715e): complete subsection reference.

<a id="canonical-3343425455aa74e11a5226e6753261d5cc123b83bca3c6ab9c58bf9c391b943c"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key / c7e6c4156d10 / 4

- [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--virtual_host--reference--group-001.md#canonical-d7761c128b31d151dee5f3f6a5ce50d3afde313a562738c3d3f30e15fd01d6ea)
- [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](resources--virtual_host--reference--group-001.md#canonical-bbae9d33eb8eeccce25e035f4feb9f9db103e4a81309e264f1ce7bc70aae715e)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d7761c128b31d151dee5f3f6a5ce50d3afde313a562738c3d3f30e15fd01d6ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-825bd329a98dbb8477b2bbe8e6574a0ab29588b364ca26b894c371e1c1afa027"></a>

## authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271)
- authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-c944bda9bfd6772de8b6970a462f95342adc8839f1f8afceb3bae75a5997ed63"></a>

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

<a id="canonical-b0a596715c1cc4349d1dcf53801627fda6a3cd5a6741768931bab6d89e5f3daf"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 3

<a id="canonical-201f1342fa0443b8afd8a4b39ccdd51e74672bf2cfbb21c2ba191100f5dff82f"></a>

<a id="canonical-269e1a2eadf7238fe1ab02c78dbb26f2fa48b18850dcec06ce64aac51aa0f912"></a>

## decryption_provider property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 4

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

<a id="canonical-f56f20e4a227ed948ce40e1631d407523221d3fc5164320f1c9114010dbb5c0a"></a>

<a id="canonical-54990b9d2fd639a16269b4e7ad53b8955d0548190e8fa158cee00b12cd892998"></a>

## location property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 5

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

<a id="canonical-698d100a61e1ee4d251f0893efcbea7fdecbfa00bb1b26d30c362d63cd319dfc"></a>

<a id="canonical-fe3b6c0d40decc1cb3cb481a7562a2ad10a2ea569775594fdf4065d8d7ac6cc0"></a>

## store_provider property — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 6

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

<a id="canonical-5b1faec9c33e0af58dac5121f4882de566330161126af9af58e671b443c8785d"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info / 093b79ba3b0c / 7

- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-bbae9d33eb8eeccce25e035f4feb9f9db103e4a81309e264f1ce7bc70aae715e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c24752e03976d748568613347665369b2444e543d5916b4d73cccc1e7eb70a9"></a>

## authentication.cookie_params.auth_hmac.sec_key.clear_secret_info — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / bb81ba4a7aec / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [authentication.cookie_params.auth_hmac](resources--virtual_host--reference--group-001.md#canonical-16471d398abe309aced8e7f9c702f197076f8dac396f8fc1d2c5fb71c2b6cf2e)
- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271)
- authentication.cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-1cae42bcdf6fdf941e5e4279763913a530a240a3df24b5231ca7aa7ae21e97d3"></a>

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

<a id="canonical-8047fd44a859e1c668f1747ead7e21fa0a32b542d709c61733b0019371b18747"></a>

## Direct properties — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / bb81ba4a7aec / 3

<a id="canonical-28d378abca79f4dcad02dc1d6088e065c95fd91f48707356f6713cc4f649baa9"></a>

<a id="canonical-49ce80be35a78a3a3213284e27bce5f961560442e9c7181917909e94b10420ee"></a>

## provider_ref property — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / bb81ba4a7aec / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6614af718e1bf93451c518dead2786f01e1b97ddee1b5a10723fba21d33b7af5"></a>

<a id="canonical-24e36a017a0db36f60503e3b79f84cb42c891031a137fe6f947e02626edb653a"></a>

## url property — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / bb81ba4a7aec / 5

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

<a id="canonical-b4ebc799714d4877fef752c599966ee1ee96aa38b674cf537c4dddc91f58f5c8"></a>

## Next pages — authentication.cookie_params.auth_hmac.sec_key.clear_secret_info / bb81ba4a7aec / 6

- [authentication.cookie_params.auth_hmac.sec_key](resources--virtual_host--reference--group-001.md#canonical-485a89c987129854613014bfe3e6abb3761baf218c11113f2f3422ff0ab97271)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ac3c3e93bff2afa576f4477dd3018f35f22db62eee75ecc505fc1454e3e2185a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d68d250af1613ff0a2f4e34e61ad8001c500f7a68d3147c9ed976d9e3f4ad401"></a>

## authentication.cookie_params.kms_key_hmac — authentication.cookie_params.kms_key_hmac / fdd784d07d92 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- authentication.cookie_params.kms_key_hmac

<a id="canonical-79fc59c686ff0893f7c1ddec8d2bf79040ee6f4759c3fdf7159eb241ecfbc415"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
kms_key_hmac = {}
```

<a id="canonical-b91c7d967bd0acdc99cfa5e4cd59432a4c79376dea1d7c7d7ac6d4ea01dac8cb"></a>

## Direct properties — authentication.cookie_params.kms_key_hmac / fdd784d07d92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3933d12b677afedf4e44cc47756622a8a6ad1ad005655a3d65127bf687b546f0"></a>

## Next pages — authentication.cookie_params.kms_key_hmac / fdd784d07d92 / 4

- [authentication.cookie_params](resources--virtual_host--reference--group-001.md#canonical-5d779da8866c00b5134b9b88fd22babcc22d600d74f48b2498559bb28681ad99)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-371532deb04933472eae5a8c9848338b94f5efb94727652e6087f5469e349ad0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf575365a32b1c721c6322ef4d5f9948ab2667649582ef280f0e019f630b878e"></a>

## authentication.redirect_dynamic — authentication.redirect_dynamic / 4a7eafbeec26 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- authentication.redirect_dynamic

<a id="canonical-e7493e214e93445166bea6c1d86f970947a3981ba0cb2af4d52db273780c1b18"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
redirect_dynamic = {}
```

<a id="canonical-a0ed7e10215fdbb5eed34e36dab7eb4767200f702c3783cbf0f64a33ba305a4e"></a>

## Direct properties — authentication.redirect_dynamic / 4a7eafbeec26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e862cfc4d1be8e4d7935a1ce3357851da3cb672b8d11cc82dc480dd5ceb0c301"></a>

## Next pages — authentication.redirect_dynamic / 4a7eafbeec26 / 4

- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-72ea9f91c6ae61b428fe85a9d2864eeadcd46a215d93d86bc9456760c159b3f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2e5d65008503dd0c31daded65b9b210e0aade7aff5dfc9d4ba148842098cf03"></a>

## authentication.use_auth_object_config — authentication.use_auth_object_config / 99772ec1c982 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- authentication.use_auth_object_config

<a id="canonical-24a8521af663fb5c1b3224b81d5801ee4154288c0d9050a9c562952d2ce4c852"></a>

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
use_auth_object_config = {}
```

<a id="canonical-9704685e5a028d81c8d62bdc00111d67aaf630c0e7d071bad4606b74d1c9af4b"></a>

## Direct properties — authentication.use_auth_object_config / 99772ec1c982 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0a6eaed782aef7318eeef9f271fb78fc8d38a0ef8c784ce6c8fdaadb1dfe8e6"></a>

## Next pages — authentication.use_auth_object_config / 99772ec1c982 / 4

- [authentication](resources--virtual_host--reference--group-001.md#canonical-834a9e0f8267cdcb3d605c3f27442833c3f0f6463c508832efbf10489591cedb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ee1ac983e9b8523722832c7e02ecb4b2baa4b4e1c2aeaf702fb8ed19891b69e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e7dc7181463259743b2ba10fac27e832b2b8729c8c5763fe5eca3b4a46f9db6"></a>

## buffer_policy — buffer_policy / 5351efebb26c / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- buffer_policy

<a id="canonical-c5ac6ba00f7569a6353a1ec46c7d50866400b6cd58eb91c5c4478dad09008610"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc138d0e0bba800b84f6f2f6f24fad03b153e660509a7a806d3a5bfea14716d5"></a>

## Direct properties — buffer_policy / 5351efebb26c / 3

<a id="canonical-598b22405a958d7a3b3327176613dc7d3ed3be8bad0ff344a435660092c4c8cb"></a>
