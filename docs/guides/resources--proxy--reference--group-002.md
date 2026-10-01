---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-3321978182fc49202b82ca5d0f3f3a81adeb3c1fa03957fc2b7325deaa2769b9"></a>

## dynamic_proxy.http_proxy.more_option.compression_params — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.compression_params

<a id="canonical-213e7b68c13f0b23f51adbe941d36c0be190aced7d6965ceb4fe9211b7e2d29b"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-22ac68567abb6815eca6ee453dda19d597a3ca0d0423090265e150f9c39bb9da"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 3

<a id="canonical-1ade19f0532e84732c2f7033ae7186b0acc8c69e642b75fb0591c53bed0e398a"></a>

<a id="canonical-5a924434ec3996ecfec34ae11f54b2fd23e87510986ee395520b33f24c3b34e7"></a>

## content_length property — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
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

<a id="canonical-d5a48af64dcc846fb7bf88f6f5cd0f43d002614351c39836b6ac02d326b8d2ae"></a>

<a id="canonical-9a7da162c87c2c6ab490f730ecff3c2f658d0c17899539084ae1aed36fd6acc9"></a>

## content_type property — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/javascript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/javascript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

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

<a id="canonical-ca3bc4ffead0383e7f7326a9fccc72d366b9160a6e526037870d8bd3b49f5c0e"></a>

<a id="canonical-cccddd0d5cae6c7f00090dd842af4c39d2a060b125aba5169a8fa936874b97e2"></a>

## disable_on_etag_header property — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 6

Type: `"bool"`. Optional.

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

<a id="canonical-70da857101199fd3ff3a66c26937c982eed85bcfbbaae4aee99c5fcbb3b5d301"></a>

<a id="canonical-9971053419bebe7729dd1df4e7ed76b2499757b0fa3fdd5baed3ec3f285e0238"></a>

## remove_accept_encoding_header property — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 7

Type: `"bool"`. Optional.

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

<a id="canonical-9405c635b192e86e7a7bfadf70f484991ffb0a1d58d75feaadd52f024b995c3f"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.compression_params / 28dd2d48d5fe / 8

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c4ed3745a27a9c81ec440e1267ace3b87d85ff6c39d101721de6286825f46fb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fae3d9cf589b061052d9dc6ebf0120a6133070c8031fb7ab940a6deec2bd03c"></a>

## dynamic_proxy.http_proxy.more_option.disable_path_normalize — dynamic_proxy.http_proxy.more_option.disable_path_normalize / ed61d0c0d3a3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="canonical-514171c0d87873a53b8b1347eb29a4e06ca64ce549e757979e79b160690e016c"></a>

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
disable_path_normalize = {}
```

<a id="canonical-8b4f25e87c1c99fbf6e74d7eed24a90bc6033d525dab412d6acf91f5d981cb92"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.disable_path_normalize / ed61d0c0d3a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f5be6789bc91c367d7bdec0631ef572473cd1ec8bb5b2f4c1c09befe92a53af"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.disable_path_normalize / ed61d0c0d3a3 / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a7f49efacc90aa2781df43b8b4993ca69b62106dab54afc06f452dd61e52af5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ccb81ac58d068a88cf79ac90dbe31fc29648452782634aa4b3d95fb2dd2e92f"></a>

## dynamic_proxy.http_proxy.more_option.enable_path_normalize — dynamic_proxy.http_proxy.more_option.enable_path_normalize / 27696fbeb91f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.enable_path_normalize

<a id="canonical-22767d0b9a6cbfe9bd81c83311122aae19c737be37f93e7594f0bd5a16611f91"></a>

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
enable_path_normalize = {}
```

<a id="canonical-6a228be46039fbc27844b4254711e3ee69abbd43d91fbce9808114d696018956"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.enable_path_normalize / 27696fbeb91f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-893707096931f245fb9db64b3244c1f63eba0a4ce7ce27e44ff87fe54f1ddec1"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.enable_path_normalize / 27696fbeb91f / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-7079c85d332c593da75961184f9050dd1376f461677914bad6e33fc70f9ad482"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60665246be53d6e56abebb7bf27662ac2ddb82689a391ba0f7c863c7ade9201c"></a>

## dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection — dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection / 2ea35f7500ed / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-6d74ee486240c9b92f6a4044fc64d8cf4754ad39715c4328c92cff3947178213"></a>

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

<a id="canonical-d26978503975cdcc7d38f61e0d6dfb4411479cb592277597060526cf5e0105ff"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection / 2ea35f7500ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f9b6363b100d750c08a0cd338920f04a3aa970d3ad01d152e384ec3ab7d20f9"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection / 2ea35f7500ed / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-049ed31a3bfe139c4c9f32f89e275a55792c82589295529ae4b5240379666d2b"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add

<a id="canonical-8aab627787a43e6a87aa344b972343838c51e5d1fbf2c144192c3f11e94ecf85"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-12acae6f7d95498655314f5205775055a4b2b40891b7dcd5096590021e01ddfa"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 3

<a id="canonical-24a3f5cc95b0244665752f1f7315b6367d64fc899068e9baf29ddad7b9c439c4"></a>

<a id="canonical-43b96094d30bd59ee9f4ad7a3da3e5e8ff393428a5779c47e3bfa2e610851d7f"></a>

## name property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-b18ef25c6a0cd60cd8fbb47f82056a85f4c15e54566b2ce3fd233a64f1dc8f55"></a>

<a id="canonical-8f739313d2b3bf5f76345244da67b4e48502fea16a3cffd71b3cd87d9328212e"></a>

## overwrite property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 5

Type: `"bool"`. Optional.

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331): complete subsection reference.

<a id="canonical-ee5350ec06390e91a58e2e5c95f00bb6ab63b5e63dc5a258969da32e3747c0d8"></a>

<a id="canonical-50946bbc07865d7c0f6d72c62b6d00a5c1bd19bd06409404e6c4e522e0673164"></a>

## value property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

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

<a id="canonical-a613e422f3e1bfddb02f0b325b72227716b4d6fe6968841841477b54d10c635c"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_cookies_to_add / bc2fc0e03f82 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5722f433bbf50396fc09a4924b1febaacb38aad9f717fe5baea224fb9d02871"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value / 1e53d8427490 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-01c74856ea4d34518868cb26f7635104fd9057051ec8d73296c570fcf31692d6"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2abd94520a8ea919e9dd08fd0906b6d9e73b97513246e41914edcb0ab674eb93"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value / 1e53d8427490 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-836dccc20f6c7be37f11aeac08060d85183b91522eb13f15f7b1ee97f3d8826f): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-afc9fd3643960b8e71eba8bf63fe59365c2650b87bbabc376bdac28cf6fa773d): complete subsection reference.

<a id="canonical-ce9ab8bdb6ff9acfb7fb39a5872312df9a2e314aedfe6f46264cc1a22b93e034"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value / 1e53d8427490 / 4

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-836dccc20f6c7be37f11aeac08060d85183b91522eb13f15f7b1ee97f3d8826f)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-afc9fd3643960b8e71eba8bf63fe59365c2650b87bbabc376bdac28cf6fa773d)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-836dccc20f6c7be37f11aeac08060d85183b91522eb13f15f7b1ee97f3d8826f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9e3fae356467ceccfaf80b580077d153cba5fb1c4e84b2ce4e1f0121217227e"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-eb50d4cc3a2d7527e65dda6369bedd22a369db8c7bd5150d157d1f227860dfaa"></a>

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

<a id="canonical-c7a150e8960a7c4d58ac657b116783406b52f03df58ad8594b93eecd7826f1f7"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 3

<a id="canonical-744de0473da1b0382ec08b39e73894c635bcfef56caece6ac673fc2e35bd9b7c"></a>

<a id="canonical-247788003267d64bc5e91350dc3b60321b9ee219e5f0f7a99f6c8b4e3525505c"></a>

## decryption_provider property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 4

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

<a id="canonical-d526c7751195e4364d3602722f6c1704e6019684e5dad79faabd0d800875ad45"></a>

<a id="canonical-5eb6265267b1ac2ffb1a3ca6181b904cc1ca87de4b68e3653940d0d2444395f5"></a>

## location property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 5

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

<a id="canonical-5eebcd0c7e213846afc98db0a339c0529f1b623e03d1c86a2ce767a9814475e4"></a>

<a id="canonical-58e75ffdcdbbbdad85450e5772c5f550025f43b812b76aa651ccc082afb26fee"></a>

## store_provider property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 6

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

<a id="canonical-637949eb061b15a92093a55ba644a147568e6d30cdd61e0c7938939478b63046"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfo / 770c8f5e04b8 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-afc9fd3643960b8e71eba8bf63fe59365c2650b87bbabc376bdac28cf6fa773d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca027149daed0dc1fb3ee2b542dd25d51064d105bdc661e3a92ea5b30fb474a8"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_s / c6b8ce0dd50a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-a9c20b2353b22ba658439b51d362dd35c2eb5626ec402fcb037675585ddd6e4e"></a>

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

<a id="canonical-e0e418e880c2caa2a7208f52bfa7bbb7e27d114da6fa517c755e73663585bcd1"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_s / c6b8ce0dd50a / 3

<a id="canonical-67d20dbcf357e78728dab51dcb495e19e69a2dcffbf11658291b73babf2df95e"></a>

<a id="canonical-f30083474e1e9b8c1553d88923cd9ab0c4ed56b13b7233657bef77524f4c17c3"></a>

## provider_ref property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_s / c6b8ce0dd50a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bc7923339a2797bf0374ec93ad3fb4a589c47c0d9f7d8abc5cc3edca3ff84f99"></a>

<a id="canonical-252f8b919df1017964fc7c66581700721e6d59499a4a2272b9734cfa8773769a"></a>

## url property — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_s / c6b8ce0dd50a / 5

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

<a id="canonical-9f219fd78c8068f6c19b1f220e539b78e77276925ad8d77a649abe8dd931fcbe"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_s / c6b8ce0dd50a / 6

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-a965fe737134d58b55b54254ab971c015709616a6557f99266f526bd83924331)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3435fe7c469fa7213138a8cffa4e26d2605045db1016727a810e6052ab8994d"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add

<a id="canonical-a27640bfabd62f0fc60548a6b293d63a34cad50d2aed5976695dd017416318d6"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-93107347a28da60660f839ead0fcd12a9e81e1532c5a08cdfb41ee44b2626a34"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 3

<a id="canonical-a41f98db6bc763e27ace2163d1333c0a9eb8a622fda3fdf7351c325d1e2dc9bc"></a>

<a id="canonical-8a98cc5b339a98b0558a8cbebbb2c33f680143e1483fe9f3a6c959c6f05d8a2c"></a>

## append property — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-33eb189dfaf28c867c6c0ac5b7b240ec5acca8a5ea103655b339e3bfa239ae88"></a>

<a id="canonical-27d0579b0fe46c7b455272633cf76d100778d2d60fdf3ca4e4417e6bd92e771a"></a>

## name property — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd): complete subsection reference.

<a id="canonical-fab2f2e257d2add31da82e8ced5aac3ed54d5d2f923716c3255873a661afb05f"></a>

<a id="canonical-79b87ce29cbf239afcf88c4c245eabf80b7abd804f99c7e35b0301efc613bc90"></a>

## value property — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-3861d0a783dbd6d3f9c9100fcb76900d5f0dbde353cfb0d8c7b1f0265827a021"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_headers_to_add / c21fe0c54150 / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea8eb168cb81767f94d74006fbea17aedf768620deb17ebdc1812e6dda27579f"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value / 66e7840faab2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-1297978fcaccffe3e8b3e46c7acfd978245b9897225ccb84c832cff8cb6b974a"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a86aa154d7f6d0b26df4ccadf611a024aae0914f4b70b36ede8561432547ad5"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value / 66e7840faab2 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-a57003ee837c86f5e34f143a8daf13d95faf95bc5018aa5568339c3bee87ba12): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-0b40ea52c0df70481ada13e8ad2a660a45738d6ff8694e9fb26a89560311c8c1): complete subsection reference.

<a id="canonical-2bb53d2045eaabc3117fb3a3f64875bc0abb29dfa1799665c784059867993fbd"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value / 66e7840faab2 / 4

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-a57003ee837c86f5e34f143a8daf13d95faf95bc5018aa5568339c3bee87ba12)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-0b40ea52c0df70481ada13e8ad2a660a45738d6ff8694e9fb26a89560311c8c1)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a57003ee837c86f5e34f143a8daf13d95faf95bc5018aa5568339c3bee87ba12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a75f80c304e49c1fe470a54a3ca975f759c55e9a5d6c78307c8862b309beaeb0"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-5dcacd39c1117c5334e9b2761e2e884d03c9c6c1367a47f09aebfc02ecc119d9"></a>

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

<a id="canonical-3e3c77f75f7d92a8f679cd32ab442b2a83129994904d7a9c737d16394c6b9024"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 3

<a id="canonical-f7f42434e3aaf79acae9472d6cabfac44dbfefd111a6501e83a1ad14bd692e8e"></a>

<a id="canonical-c6090d8a8280b05d46009d576b29718e828571b9f05b19acf6390144cceccbf6"></a>

## decryption_provider property — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 4

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

<a id="canonical-6905e4d9c6c7d1f471567398ff8e62cadc489a2c9c511fc703bfa1a9cb4a9414"></a>

<a id="canonical-fd4aafbf8248803543cf98e3717a6d9ef297d59a684af6538c16728b0666ccf1"></a>

## location property — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 5

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

<a id="canonical-4cf63e417dbe54475131b2766eeadd7eceec8192d4c68054746233ee29f8ce94"></a>

<a id="canonical-99d7514c44dd772a47ba8ec3b1be583942ba36133bb56edd0d2421565f4f2d65"></a>

## store_provider property — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 6

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

<a id="canonical-c7fd61eb755c3c556671a81368079ee79b4e8c7eaaa881bf8573e8f4449b5db9"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfo / 331cad5dddfb / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0b40ea52c0df70481ada13e8ad2a660a45738d6ff8694e9fb26a89560311c8c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7845a9deb0cf46bbcb44fb46e35429e29c90e4eb652adc0d6497f0984063c2c1"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_s / 0fd3ae42b07e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-eeb8058ae3004af59224908e0331d5a56dad4fae4f2d73377b907f71db37540d"></a>

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

<a id="canonical-5e56fb3d090d0d5cf1b6b19754cc365aaff68e947cece74ada9f89a611a367f5"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_s / 0fd3ae42b07e / 3

<a id="canonical-48ea85867d4f64879632794f1917c6c829ed1991326146f546c29d3d6a09c7f4"></a>

<a id="canonical-b43fa7a3e1bf1edd8ba988a7e8320d4928f5dd142c57980e2372497d2eb55cf8"></a>

## provider_ref property — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_s / 0fd3ae42b07e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e329d30a4bb4af93995301173833541b2c53d5f76fd6c28f88349069f5f5c4f8"></a>

<a id="canonical-9f0d6ee418d56d452a5925f8eef690b7e6f3794c1dc46453c6c4c9149731e4eb"></a>

## url property — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_s / 0fd3ae42b07e / 5

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

<a id="canonical-956874a0c988a37a6cf614a723310ef80abfe173c518db9157152a7524319337"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_s / 0fd3ae42b07e / 6

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-11c29a0f1cbff9dd8e5af63be56c4e97e6d13a261ab7daa5769fb3aa14dae1cd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-990142ed0ea15d6f1345ee4024625755d5f309fa041a85cf7194f75aa1f3c7b6"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add

<a id="canonical-c9d5fef1f1dccaac93bbdb2c88b9b873a1b2e445aefe859f91e09853e1b59567"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-7012faaccf388d9c43d74b6a226ee6cea296cfd2f0a98ab3d97b17f38d220e5a"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 3

<a id="canonical-90dda50f44b36dc30653cc414194ee2bc14ebb3350d1550b9d4ec9d64be7fceb"></a>

<a id="canonical-594ad706d4bebdf36ec2a04c804eee641a7f1d3fe54433c823d12e1a7671a907"></a>

## add_domain property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0267e9fdb908a2902164573b845e0bede8daaac7d964f4bddf631e8421c0e008"></a>

<a id="canonical-757948b5c462c1369c525fa12123698e325acdf03140a3943d8ae14ae9257fbc"></a>

## add_expiry property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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

- [add_httponly](resources--proxy--reference--group-002.md#canonical-cf9c542eefbe698a70cf03d7bc03f5bc4202c9a85dc80511d3a9b849831ba3ae): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-002.md#canonical-bd68d7f3473fc5e3e6742b6439970f69d1ef202772cc1acb1d43f2556ec2cc2f): complete subsection reference.

<a id="canonical-2a2b7025d1d09bd75f702aca6c7d0cfbc88b2b9479b5bb457a909a9b15a3f235"></a>

<a id="canonical-bf99aebbee52feba6470a347ca9664142bcd4c417320a7f4204684b7826d2eb1"></a>

## add_path property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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

- [add_secure](resources--proxy--reference--group-002.md#canonical-ca7baaa997ed67e052b727796f0f333884d1560f8d5cb42d681cf7d2a8798380): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-002.md#canonical-aee8fea49bd5223f665733dfe2c67a2f6b1b7e39d5a7f92347fdc9eb5b6a65cc): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-002.md#canonical-210334523b56e1361b76d426119ae803084e64aef9d7ad37cb532adbaecc91f4): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-002.md#canonical-3b58d981cde4ddb6f5eda414174fb546de04af935f8b8bec6f2080107e19e31c): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-002.md#canonical-2e734523e773003c0d44a42955404602bfb814d7b734f896027241421d471838): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-002.md#canonical-a810a733abc3e517ebbb39f95b00e0dc5b11d590ceee9166eaa976806598eade): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-002.md#canonical-11cfed8de651ae404db033a1a83bb50612da3b4edbb399408c21472f85f8335f): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-002.md#canonical-088aef2bd84e53755fe94840236931c94b3b81dc3724610c78ed46168564646f): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-002.md#canonical-baf211f49a3ac74e5aa2dff10f3d7b6eef6e61401167643b2886b8970bbe9412): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-002.md#canonical-3a22f2bbf12bdeab50f0a7fc442d1e980ef1fca7f6f77be180f9c772412d29f1): complete subsection reference.

<a id="canonical-e8f2aa39c4bfc7b46ddac685037c0456c23e3d1e91bebfd95410d4907a8ace4f"></a>

<a id="canonical-29fea3fdca425abcccf953c3bc0bad216b2e6146bc12ced2810b88b80d10fd0d"></a>

## max_age_value property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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

<a id="canonical-5b09c170e04f634fcb5d14ecacef5da066bfd5021daf59be4b6565c9dba4259e"></a>

<a id="canonical-9d5601b081aa8667cae5a0a4b2b3268bb3131ac15baabb6116c3b8afecff4a41"></a>

## name property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-03e23172864cc0716eef0bdff6a6f2f77b574b457add46302baedfe843580704"></a>

<a id="canonical-875ee2eab57b2235e98cfb176f2c421831287a9b65fb000d0c50851cb0f332a5"></a>

## overwrite property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 9

Type: `"bool"`. Optional.

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

- [samesite_lax](resources--proxy--reference--group-002.md#canonical-204a1167aab293b7b5956e4ad577a2de73d7543607e13cd04fe785f764be96ce): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-002.md#canonical-e23f9bae6d6ec39805e36e0430a6d959e0cf3cfd092214a3d5b590bc7b1c8d1e): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-002.md#canonical-8183b8021024da9657126994656d2e368e154b677f3ef84b7cc212db1607845d): complete subsection reference.

- [secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f): complete subsection reference.

<a id="canonical-cff5654e0cd1d2d9bfbb5b692003b3e1b45d9d21381e1cbe915c6aa2774ed06d"></a>

<a id="canonical-da9d74775df3962834060d16043e7e0c4dd625b8665841dab5cc4c29843a7a56"></a>

## value property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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

<a id="canonical-fbfa70d07bd46e82823bad8214b1d9d0e2f4e02e26cb1d22f3f8e8b8f348e17b"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add / 6aaff69e00d2 / 11

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-002.md#canonical-cf9c542eefbe698a70cf03d7bc03f5bc4202c9a85dc80511d3a9b849831ba3ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-002.md#canonical-bd68d7f3473fc5e3e6742b6439970f69d1ef202772cc1acb1d43f2556ec2cc2f)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-002.md#canonical-ca7baaa997ed67e052b727796f0f333884d1560f8d5cb42d681cf7d2a8798380)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-002.md#canonical-aee8fea49bd5223f665733dfe2c67a2f6b1b7e39d5a7f92347fdc9eb5b6a65cc)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-002.md#canonical-210334523b56e1361b76d426119ae803084e64aef9d7ad37cb532adbaecc91f4)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-002.md#canonical-3b58d981cde4ddb6f5eda414174fb546de04af935f8b8bec6f2080107e19e31c)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-002.md#canonical-2e734523e773003c0d44a42955404602bfb814d7b734f896027241421d471838)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-002.md#canonical-a810a733abc3e517ebbb39f95b00e0dc5b11d590ceee9166eaa976806598eade)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-002.md#canonical-11cfed8de651ae404db033a1a83bb50612da3b4edbb399408c21472f85f8335f)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-002.md#canonical-088aef2bd84e53755fe94840236931c94b3b81dc3724610c78ed46168564646f)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-002.md#canonical-baf211f49a3ac74e5aa2dff10f3d7b6eef6e61401167643b2886b8970bbe9412)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-002.md#canonical-3a22f2bbf12bdeab50f0a7fc442d1e980ef1fca7f6f77be180f9c772412d29f1)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-002.md#canonical-204a1167aab293b7b5956e4ad577a2de73d7543607e13cd04fe785f764be96ce)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-002.md#canonical-e23f9bae6d6ec39805e36e0430a6d959e0cf3cfd092214a3d5b590bc7b1c8d1e)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-002.md#canonical-8183b8021024da9657126994656d2e368e154b677f3ef84b7cc212db1607845d)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-cf9c542eefbe698a70cf03d7bc03f5bc4202c9a85dc80511d3a9b849831ba3ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07448a242a2f71093f393944c1be651609bbb991b2b6d5a9bb709bbe585eb8db"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly / b018b44548d3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-1e9ea30cac410586e3152382aeb0992181d9d796b1c64dadcb305ba92c40249d"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_httponly = {}
```

<a id="canonical-21d8bc10bdfc054d0a68a35b29fb288a1a61fc62d9cbf79e4f9a7740cf244931"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly / b018b44548d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd239349a72826c015064735215683ae37244610f2d91b00f13a30d1d874c00d"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly / b018b44548d3 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-bd68d7f3473fc5e3e6742b6439970f69d1ef202772cc1acb1d43f2556ec2cc2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e3ad6a885729cbc67ae682134f75360769af7777e55194d1fc213a034d5f2b3"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned / 07c266c5e33a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-457cc8e3666c259ea9a59caa709695b60ec62e1028c3420d71d0438cdc4b4d75"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_partitioned = {}
```

<a id="canonical-b222b0a54a333515688f8e9554741eead2fce174bf2f5213251ed126b14cfaa2"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned / 07c266c5e33a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ad74b1673918a4ec67ef3a10678bd22d06f0508bb607d7f96ae8e30a0ec8726"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned / 07c266c5e33a / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-ca7baaa997ed67e052b727796f0f333884d1560f8d5cb42d681cf7d2a8798380"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-077239a75d324678fecad769186cde676f848c399d18585ab71c4fe7ac8c6498"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure / 6f258f73b174 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-a05ff8b0e1868e53f63af923df183709231c3e80bd08c8769c92483e90fe80d2"></a>

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
add_secure = {}
```

<a id="canonical-1c747c09c4303805ae89511f11efdc00646d32c5f357a1dc2cee5d7cda130350"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure / 6f258f73b174 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0063d94c608d8d65022930d99673c79a4ad6edaad8184da9f248757b96a81cf"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure / 6f258f73b174 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-aee8fea49bd5223f665733dfe2c67a2f6b1b7e39d5a7f92347fdc9eb5b6a65cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d44b73bec65e36d3a6ef58dc1f283cda7d13a6a684f2a9e72467ef5af3aa0a7a"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain / 1efdda55d878 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-48e69eb70eef49eb1202a97f1942267ac56a25b75884becc1095114cd74b63c6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

<a id="canonical-eb55b28ca0876d8586734c6daf3fb781865851e250403ecf36c0af3b0b6612f7"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain / 1efdda55d878 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-256b86833d69b7a888701d76ee3c9550a5e6eb7777443ebc95f1c2426e18a113"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain / 1efdda55d878 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-210334523b56e1361b76d426119ae803084e64aef9d7ad37cb532adbaecc91f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8106f2aedb525807822b7b95093561d833e7e0079fe9bf9e2b07da9f4c73b87a"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry / bcfd741b3327 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-b9b59ed04e5568f76bbacef32b31893c84310a7e561bf2a34e14d0957412a088"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_expiry = {}
```

<a id="canonical-2317cc8f939fbaa45d1c7f6fd3233bc7ab3be962ae62286fe83c70664fdd11ab"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry / bcfd741b3327 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2b24a70df7e8c84f99c4f044d4d5c388268b7bb06ac262b3891359350abb9e4"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry / bcfd741b3327 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3b58d981cde4ddb6f5eda414174fb546de04af935f8b8bec6f2080107e19e31c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a71fca0e21ffcaee834bbb0f9a5da4ea9ca157257009445216779d9b545fa4dd"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly / e9775751bc35 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-2a777b9a5099769dc6d77e949022bfbc8dd4dede555c148a02c016c79324b982"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

<a id="canonical-3bb411101aedf23bccf9bdd0450cfa4c6e651219c9e8185717090fc1adce8e50"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly / e9775751bc35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a3b8cb1fba9afbf55ff089a43540153bdb5fd1d78275be1798c64d6bb40eeb0a"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly / e9775751bc35 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-2e734523e773003c0d44a42955404602bfb814d7b734f896027241421d471838"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f8eb4c5f7030a7151e4a38c128ed7b4c1578b7a815a120b76b3084591d27a91"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age / 589ae942a460 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-2563e709019f874ff61be5a968891f6a71f2b344946bae25df8f7abe1caa787f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_max_age = {}
```

<a id="canonical-514d642461ac1a01b14f0727f7dfef6c160e5fd8ad269c21f0bc7bd51fe17c14"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age / 589ae942a460 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-27a22a05c7354a8b0f27ac7125d06978faf03bdf8925c47644848946a0390a4d"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age / 589ae942a460 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a810a733abc3e517ebbb39f95b00e0dc5b11d590ceee9166eaa976806598eade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4603e165288f2e7cd4852ba33e0c1731bae260d0c45861e317ed9da0260d3a11"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned / f055f02352e3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-b6e513a5a9f1d0269bd4b962bbf54295423c3ec1d874c688a499e468fcdd84c4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_partitioned = {}
```

<a id="canonical-24802570c5c979689c66f9d78618d0e69eaff1ae4836e5aa9e786da3babf2645"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned / f055f02352e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6250422df8481624237ca841a0ce8c49f5bedcdd4e529aba0493e968ab603016"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned / f055f02352e3 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-11cfed8de651ae404db033a1a83bb50612da3b4edbb399408c21472f85f8335f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88b780c71fb929c5ebfe36b04b66d11b66079e06b10b187d8da33fefe4dc2682"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path / d805e300bc2a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-0be8ce327b355ef34b455ac46b56bd37458d0185c77768e1a06ddff65da825d1"></a>

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
ignore_path = {}
```

<a id="canonical-e3fe7309095644c892548c395a596c0e3a11c803e5cf7d5922c61877ad14d18e"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path / d805e300bc2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ac12232ae00829641b13f87924f9aed7d7bc1b19a9844af1be045e6f6cce0c8"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path / d805e300bc2a / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-088aef2bd84e53755fe94840236931c94b3b81dc3724610c78ed46168564646f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb320c5e518740c4e70c7c9ac972f43e90a2b55220b6589a277d1691f375d5e5"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4fc175719ac3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-d37a3ef59c111fd8f308278216a034ff918491e43cb5a28ec818a939a04c8529"></a>

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
ignore_samesite = {}
```

<a id="canonical-d554325f6bbc5532344e3c03fc706b1fb3161a088e7c02353e401fc2d6f918e6"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4fc175719ac3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c954278d3ff385f50419b857d4f1a4c3df2ac41697c66e7850b8f65762d67653"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite / 4fc175719ac3 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-baf211f49a3ac74e5aa2dff10f3d7b6eef6e61401167643b2886b8970bbe9412"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b9f1f50b6f8d07d1c5ca03988b2daed96bc5d36bc38d45e1286632b3eff1f8b"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure / b64e71aff424 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-97f8050df33607417bce799b772f2d60568ff9c0db4967b6dc98a29fe2f9dadd"></a>

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
ignore_secure = {}
```

<a id="canonical-986b5f3c2965916ad5dc2a460182ba324911f2533b8a83d7afc9ae9b9a5c604f"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure / b64e71aff424 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71695bc2e51784f3b93215c47a2f3aac09eb968724eb73031715b96ce3e3930d"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure / b64e71aff424 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3a22f2bbf12bdeab50f0a7fc442d1e980ef1fca7f6f77be180f9c772412d29f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-458aa49f8ab8d42ba339bdf9a667d4ce6cfce21208dc734539e1752dbbe7f595"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value / 35e2268afb35 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-7a8535c7e42feec9e64840995d770a65e4a14eeff4171de4c346ad5b123ced92"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_value = {}
```

<a id="canonical-f1cd3470e72c22e680806b28424055344cd2b0870d90cf0ad48d499ee6cc16ac"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value / 35e2268afb35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4cc84cc87b717750b27742e3457736339b3cc56e24689ceaf60fa5c08b1350da"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value / 35e2268afb35 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-204a1167aab293b7b5956e4ad577a2de73d7543607e13cd04fe785f764be96ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a57fc872226d94f5f0ebc4323c79cc9049bf7d5757523160a205e2a2c02bcc2"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax / bfaf16ce67a1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-95cd35b2bab01038c4409635e3ec86135073c1a06ec4597bf78c82a86a8f271d"></a>

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
samesite_lax = {}
```

<a id="canonical-b17f6cb2f04303d4257e1c155df00405d29e88e630f409a2a0bd89431314baf0"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax / bfaf16ce67a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-329ebd4f06dcb9b01aa4be9dd48044e6eb6bc530ed4063f8fa37b0d1f378c5a9"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax / bfaf16ce67a1 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e23f9bae6d6ec39805e36e0430a6d959e0cf3cfd092214a3d5b590bc7b1c8d1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61be2e69a6d99d1ca6a50a782eca1a201c3fbaf8160ec0e94df77e512676fb25"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none / eb4c96bdfc95 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-326b0ab4b5938c858fe532efe200d1ba6e8f76cd6c0fcbdd5db63ca53f557198"></a>

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
samesite_none = {}
```

<a id="canonical-476e6b50a8e2c49abcb2598a811faa17b0c959362bc09381f5f26b4d2f6500d8"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none / eb4c96bdfc95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2533b51dab9cdbb21fef49ccaadccfd4d4cbb88045e566c67f1dd151b294e225"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none / eb4c96bdfc95 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-8183b8021024da9657126994656d2e368e154b677f3ef84b7cc212db1607845d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8aa3da7f86c898901438e58a1be1583b8808a5ab0d21612c5f4d9abb657b251"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict / e7172295914f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-3342ac460f6f6e0d66003eea0e7963c91aa135d773ec597459a1d7fb158dcb1c"></a>

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
samesite_strict = {}
```

<a id="canonical-bbf9e01e195d1e82da7013d77dc18377debeb7693e57db7d6d4cbeb8b16e979a"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict / e7172295914f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc7121e2bbc3530ebb18dbeb0a2eb7a78b45ea8af5a600321e53e12dcd2455c3"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict / e7172295914f / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34cf31105cab5719d926f7ceb76c60e99af2e72945604a6ee15bbf5799f23796"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value / 9a86222ba4ee / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-b01e1826f76ab13443fc6182520de2732384094c575f43d8dc886d68d3a7c19f"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d31fb7b1eb0471fab8c52311e1d1d83f6c7b24820d5ccf5fde0b35bf6a4986f"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value / 9a86222ba4ee / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0ad09c04829dba49411a3bf79c110b1636c858d119feb4becc3f944e8fde8b35): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-7545ae4b49a99fd53bedadfb2ebe2ed525f6bbc20037ee08bcef5e784fee7555): complete subsection reference.

<a id="canonical-67bf854b18d3551e2c68794f0cc13f0932a39e6d4f0c38fe3d78d095a1579b1c"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value / 9a86222ba4ee / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0ad09c04829dba49411a3bf79c110b1636c858d119feb4becc3f944e8fde8b35)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-7545ae4b49a99fd53bedadfb2ebe2ed525f6bbc20037ee08bcef5e784fee7555)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0ad09c04829dba49411a3bf79c110b1636c858d119feb4becc3f944e8fde8b35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0031e895cfdd62dda42d6e41c01cf9ecbe0fc97e0c1fb30de5764160342b7ba"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-8e84acb9e8061441d943276ee9cda520ea4f22d58f29532b7fc839c1829d2917"></a>

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

<a id="canonical-cad8b75eac0a12ded7036cf849ea2676627c304c7cab00d883ad77ab3cbaa521"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 3

<a id="canonical-e0f29b3c330c7151e1bd878015843304400cc8b106449e390ea394d8acd8f2a3"></a>

<a id="canonical-6db8b3ff552543aebb42722d8966ab11b6171fb8dff33806c04ee1e5e996bde5"></a>

## decryption_provider property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 4

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

<a id="canonical-93a958d9c9accc210d3b69f11c4dd254f2d8a231f5757f1ee7e19fb218b5020f"></a>

<a id="canonical-cb64a1c8bbe8e56fe8d5a330db18d489592cf9978d2ae8c9718c695e003f0a64"></a>

## location property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 5

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

<a id="canonical-606f3509da6702954ca899cd8af1d29cd1ad7a160a3d55ea06740ec5b6fda24e"></a>

<a id="canonical-d7403c63c0ea197532703e4ab6995313559dd05d81534d41ccc29293eca2a0e7"></a>

## store_provider property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 6

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

<a id="canonical-f30eda8d708114f55184c0644c34ce32d27cec180d965f9b220e9c5cf60993a4"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindf / 3af840c3a32c / 7

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-7545ae4b49a99fd53bedadfb2ebe2ed525f6bbc20037ee08bcef5e784fee7555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75ffd704e506aebf7ca627b1749430731b4a179c9a503dbef99f317d5d591f29"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_ / e1b8b41eb08f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-fc6f42dff80d3194a3bfdb42ccfe102ee4a2e11e5d2fa2fcbe1acb6cebb1580a"></a>

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

<a id="canonical-5c0e9d876ad647a65292c58bbf0709101633738ff245b9876329704d77225a0a"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_ / e1b8b41eb08f / 3

<a id="canonical-1b0fb27b140b5ed39a278b56b48ba74f050c41c32e007035a3b94fed7598bb07"></a>

<a id="canonical-503f58c3e114c561a51c76f13eccb1e929e6aa3ae4ca095edaccf81be7c17094"></a>

## provider_ref property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_ / e1b8b41eb08f / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-216fa4893e4d460c33762adac444a33ecae71be89da4186a954614713c3e35ea"></a>

<a id="canonical-9484da4a094f552219efecf8574c98a57e941965fc52ed271cbf75fec60199f6"></a>

## url property — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_ / e1b8b41eb08f / 5

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

<a id="canonical-17c68de53b8e16a3a9a055d0e9c57a4f6a2ad27b980802ff37220bd1402da69c"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_ / e1b8b41eb08f / 6

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-e7e2743b4e4e9b621fa4c589edd5549d6be132d4b131cfa27d8d42652ae0817f)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-006b40b380af54c93563173edbec112c119122cdc2f372275a4d51987f41893d"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add

<a id="canonical-22419f03471bc9f58e816611e4c854b7c4bdce0654048c00748cf95a03a536e8"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-c500e51fd3f576fbd43dfe53e154fa32dab3d0aa770381b9fa9f0219d6bb5970"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 3

<a id="canonical-375c5fa5654c8ef753e008a0b587ae5d476900dd1d75ca5fe8b72358022045d9"></a>

<a id="canonical-4f7213a96e7812778ac8463fedad5e354ab231603c226156c95680f92608a6a9"></a>

## append property — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-41fa3575e1f303eaa8f291aec3788ed92bc829811a3efbaed805ebd09fda7eac"></a>

<a id="canonical-1ef1489a5f7aadadd3708e9de6939aa941a38c9817f85ddae1b5a6bdcc22d79a"></a>

## name property — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6): complete subsection reference.

<a id="canonical-60098772833f15c0b17fdd4163d4b21fbf4a0d1e1b5e159a8aa1eb6fc08f7bc2"></a>

<a id="canonical-1a4a8a37a740ac5716cc192019fe41c1d1011c0b3ac1683a304a0c9671adc924"></a>

## value property — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-268fab72779045a77de095ed3cc439747f0d5b518b1c7433bb38ca098ec2a548"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_headers_to_add / 6802b8120e50 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-557956a3cf44ac5572fb054fcdd088e154c6efb78ef598d9b9e34eb980cb3a36"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value / 108146f0d30a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-7b67e3f04332d65dba95f0c4b3e18b94134d208f1949d7cc2c1179301cb21550"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-87dbdcf0aaf012fa670870a4871ff48d3b9aae2800ad3373170bd61c05c3dff9"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value / 108146f0d30a / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-d279f601ba3a075dd28366cb95fcd0a34639011084dad06e784e464ea5334160): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-627ce20fee842ac3becf0ed02e13b9ad38f141acf0317693d71ffb5253eeb6a0): complete subsection reference.

<a id="canonical-ef111e7af71e1ba6692188c06cfbca33fc388908f47089a36296240087eb1c78"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value / 108146f0d30a / 4

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-d279f601ba3a075dd28366cb95fcd0a34639011084dad06e784e464ea5334160)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-627ce20fee842ac3becf0ed02e13b9ad38f141acf0317693d71ffb5253eeb6a0)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-d279f601ba3a075dd28366cb95fcd0a34639011084dad06e784e464ea5334160"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-640cd1d9cd5de2f0728548b3d78657e01f0681ffd16539d356da9d4e486da23b"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-934a99e5a31560aeead0ee0eef71e88b3f89a15c73aec62be511860bea3c1477"></a>

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

<a id="canonical-e3f4d79ab2d7a2bddc3602fc8bdbde47734d994fe3cbad6cb412caa01cff2433"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 3

<a id="canonical-15d8b4062131f5709de4c99484a4126b544d94f116e7fa1e29242b15aa7d6678"></a>

<a id="canonical-b16546e2da2795e3450940fc43642c91c38d8837a675859c0fe05a5ff23ee612"></a>

## decryption_provider property — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 4

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

<a id="canonical-c72084d45953e553ee02a5e12a1a718fb0cf64c5e32a38aee41a91c939f2062c"></a>

<a id="canonical-8a2c814776f081561566aa9ee705cefaf3527d72df34213ea8e3b7a4cb6e3036"></a>

## location property — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 5

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

<a id="canonical-54b460c3455474051a6454365edb304fab676af063bc2e435c4fb0a708b92216"></a>

<a id="canonical-cbe0dfbd0e1b7c18922f9c3fa73631419e609464fa3e9b264f2a4e78f3801d26"></a>

## store_provider property — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 6

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

<a id="canonical-741cfdcc895ff28c54f04c7a22ace9928d18380da1294dc91e0b2a236e78f173"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindf / 811c715219a7 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-627ce20fee842ac3becf0ed02e13b9ad38f141acf0317693d71ffb5253eeb6a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-188e83f81e23f0c57a7be7d936544e4c853c421101dc2898668b6afc79e685f1"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_ / 5ea50c84b7e2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3eddc912eefb51342772ab0a8c39c42d5c22a809c1f7d9c3477f626412c8d037"></a>

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

<a id="canonical-87a2fe893c88914a4e543c09364fbcd043146d400913d5e8e2999431ba481654"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_ / 5ea50c84b7e2 / 3

<a id="canonical-01eef4c09d4a692c095d380476bc4a5a0c622ea37457254ba8d6298ea1bab298"></a>

<a id="canonical-cd815e27ccc651620c46e411e3bcf6c1db8e35a801946780568b5134017eb92c"></a>

## provider_ref property — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_ / 5ea50c84b7e2 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7347bb8fb5c7af03d5f33031d5f82d446c61a947a3e1e5822e0b1187146afe44"></a>

<a id="canonical-0834ee43d5489de282fbc9f30aa434eccaad266f8b4abf08c5d3403e016c507f"></a>

## url property — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_ / 5ea50c84b7e2 / 5

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

<a id="canonical-405f3bcfef6896ebd0b6e432a321e9fc6a8fe01e7eba0746b426c0e934d0d6cb"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_ / 5ea50c84b7e2 / 6

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-50e120659a480e8b9bb6ca30bc6fb760936ec1e2b369d418590102238a600bd6)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5c10770f55736e94ce6ff67fbd67658a9bf5f6ee999430d35ef3ad6ec80f4b0"></a>

## dynamic_proxy.https_proxy — dynamic_proxy.https_proxy / cab42de5291f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- dynamic_proxy.https_proxy

<a id="canonical-901bf8f71f3a9b0d1a7c0d6eec9a181156d1b4f066eb5f637974f2549b792a1d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https proxy.

Upstream description:

Parameters for dynamic HTTPS proxy.

Receipt-pinned upstream constraints:

```json
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
https_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfbf13b9f94410b8309e0400358c1f8f99675069657704a6a9ce3560cf994e85"></a>

## Direct properties — dynamic_proxy.https_proxy / cab42de5291f / 3

- [more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22): complete subsection reference.

- [tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8): complete subsection reference.

<a id="canonical-6ae7b298103a5ec1495090dbf31bce1b90a83fe17c17b89a8610a544e8c12ce8"></a>

## Next pages — dynamic_proxy.https_proxy / cab42de5291f / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a352841926af632b662cc075199f782f9819d8b1d119578d1fcdbccf4e51e0a"></a>

## dynamic_proxy.https_proxy.more_option — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- dynamic_proxy.https_proxy.more_option

<a id="canonical-af3207efedb07304c822887e9dd4d8f5dc42ea5596ef958de8dcc8a0a2cb14fd"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b813d428e39cbcc6e48d1faf2c3806a1f42803e5bc05999ad7ea3756c865247"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 3

- [buffer_policy](resources--proxy--reference--group-002.md#canonical-1b7e56bb8bbce609f6db5c32804bb5d89312e27281086cda3649e8dc4db8e4db): complete subsection reference.

- [compression_params](resources--proxy--reference--group-002.md#canonical-6a7977d1c65f1f3741b61eb4d2bacf32b5fc289ad109626731f5cd976b5dbf24): complete subsection reference.

<a id="canonical-18c587539eab37a4ac13157411c6eb696cad636d7862b24656a65237238d9635"></a>

<a id="canonical-1c4dd99ee1ea0c3159452c4e560685d85169ff6b6ef3eb1567d480167f8ba0e9"></a>

## custom_errors property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-7731b930a79e803c4efb076ed102764252b164e3523858cd00fa429764bbcc80"></a>

<a id="canonical-52bbfdd1d6493c667a9880c7664cf1597163fa233d144fb8380076bb09b8af8c"></a>

## disable_default_error_pages property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 5

Type: `"bool"`. Optional.

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

- [disable_path_normalize](resources--proxy--reference--group-002.md#canonical-86e4db9ad3d146d8a6eabbc9e94032e7d3e0509974b2fa25211fb558b4e9fa46): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-002.md#canonical-85c7bd4210e9ad7b0f8025b69c7d13e8506950f03e9c2ec5ea687dff834057f5): complete subsection reference.

<a id="canonical-880f7da5f45a05a1a84608ab2d5e5151c58a5b496bf5141c34c01f7635b37301"></a>

<a id="canonical-f0a38da9a4c8e26bc591221eb5abeb7c3634c6de97f90833909767b2fc601183"></a>

## idle_timeout property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

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

<a id="canonical-5f34c5c1e56a9d1ad080df6d2630a7968855db34146d8d40b40b960b3972e0e2"></a>

<a id="canonical-adb74ea69727df711d2ef33ce058695a7b21fbb2e5bd6fed0440f70683726a01"></a>

## max_request_header_size property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

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

<a id="canonical-95de99c88625409f62b647d01ff7139a8bebdcde264567f13afdb464e2be990c"></a>

<a id="canonical-a7e7a40a09e733e305bc2fff77067f436bfe802236939250bc16599d01638888"></a>

## max_requests_per_connection property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-a49105370a4fcb3c99bf1cf42a1cbbae63e47ec9359a25209fc7766354f3ce1d): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3): complete subsection reference.

<a id="canonical-affa7ab2b4c1a5bf9a0d60ccb68c04db8faf0630dc8de94af550fb40b8c03b59"></a>

<a id="canonical-af80df19a28c680f4ada61bf296b0066cbd34fe42b24867d5d34d6e72e3f3530"></a>

## request_cookies_to_remove property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 9

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

- [request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9): complete subsection reference.

<a id="canonical-3faa52c82256c455077bb47d3f09382fd7f505423c449a9075b4b1384c7ca17f"></a>

<a id="canonical-6b77221f20ec500903cbe2523fafc79230229fedf43cbf63856c5f06f0df8a9c"></a>

## request_headers_to_remove property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 10

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

- [response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74): complete subsection reference.

<a id="canonical-7f668ddd4211b48f0b6f60bdb8d8ca54724ecfde3737441672e6526352972732"></a>

<a id="canonical-7d2083ea60bd93fca28ff5e3acb301f08ef9cafc7705b910efe910ace0e8920e"></a>

## response_cookies_to_remove property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 11

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

- [response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4): complete subsection reference.

<a id="canonical-e0d3993566ef7b4c0ff6cb5d56c1f19da2874a23b6e233b12ec846cc43bd335c"></a>

<a id="canonical-df7be9051d64a787937401a36e58226c0cc729fd38037be02710eaa79d3e655d"></a>

## response_headers_to_remove property — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 12

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

<a id="canonical-c7bd16391c34e840cad406e9f0d8d661d186e9f494cb5edbe20177f3c52ffe41"></a>

## Next pages — dynamic_proxy.https_proxy.more_option / 5fcfbfd846b9 / 13

- [dynamic_proxy.https_proxy.more_option.buffer_policy](resources--proxy--reference--group-002.md#canonical-1b7e56bb8bbce609f6db5c32804bb5d89312e27281086cda3649e8dc4db8e4db)
- [dynamic_proxy.https_proxy.more_option.compression_params](resources--proxy--reference--group-002.md#canonical-6a7977d1c65f1f3741b61eb4d2bacf32b5fc289ad109626731f5cd976b5dbf24)
- [dynamic_proxy.https_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-86e4db9ad3d146d8a6eabbc9e94032e7d3e0509974b2fa25211fb558b4e9fa46)
- [dynamic_proxy.https_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-85c7bd4210e9ad7b0f8025b69c7d13e8506950f03e9c2ec5ea687dff834057f5)
- [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-a49105370a4fcb3c99bf1cf42a1cbbae63e47ec9359a25209fc7766354f3ce1d)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1b7e56bb8bbce609f6db5c32804bb5d89312e27281086cda3649e8dc4db8e4db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb416e3d67812ea4075eb91224598d7d1fd4b16c8c263b23b836d1f83b871c97"></a>

## dynamic_proxy.https_proxy.more_option.buffer_policy — dynamic_proxy.https_proxy.more_option.buffer_policy / 6e9bf4458f56 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.buffer_policy

<a id="canonical-a164a7cc116b986334ada46dea5e5274d820168e10ea0e96dbaf5a9737e7ad33"></a>

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

<a id="canonical-1f891b485c69cd2953a972a1a1c341aa4cb001d82e1af1fe92d2036fab98479e"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.buffer_policy / 6e9bf4458f56 / 3

<a id="canonical-8225328ae2f824fe4c47841f962edce062c37b9f1c841de8d2e9678491cb7a5f"></a>

<a id="canonical-609ab6eed8fadbfddda604e211247b477186b3788bf5c44354af45a360308abb"></a>

## disabled property — dynamic_proxy.https_proxy.more_option.buffer_policy / 6e9bf4458f56 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-d43c9fdb03a8752b15de1b69abd385791c26206c741cc08dddf640ac9797803f"></a>

<a id="canonical-f846d637897e434373e642a34ed6eda8edde5d730d9f14b92f6e30fb6394da39"></a>

## max_request_bytes property — dynamic_proxy.https_proxy.more_option.buffer_policy / 6e9bf4458f56 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

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

<a id="canonical-02d344dd727ce3b42789a994cbb8fd5d1cedfbb38a154daa5408f01cd2669596"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.buffer_policy / 6e9bf4458f56 / 6

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-6a7977d1c65f1f3741b61eb4d2bacf32b5fc289ad109626731f5cd976b5dbf24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b34d89fd8c55249275ede768a4710455c068bcebb2445107d97a310b0ed967f"></a>

## dynamic_proxy.https_proxy.more_option.compression_params — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.compression_params

<a id="canonical-e8a3d9b2dac40a251c92b50b8f9c49fce6db87e565f5f2ab7a7757dd36fdfe87"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-deaaaac4df69865393ba6f9eab9935b94a98ceb38242f3c243758f70bca07818"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 3

<a id="canonical-965f2186d92305ebb91e60c7c0eae6bb9429600e74fddb7902c62560bd4bdfc2"></a>

<a id="canonical-9f4992a4ccfea8ff063676d0022c07b6a50c0e4bfa675a51e9dc8582557d7450"></a>

## content_length property — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
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

<a id="canonical-263b3c19085c17b141339029b5b606c9353d732fcaf84828d86d52f1bef080b6"></a>

<a id="canonical-0f1615d61938752cd772d7bbf94f0088f7c2eb339bcbb114e216b57bf609ba52"></a>

## content_type property — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/javascript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/javascript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

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

<a id="canonical-1296370495ecf496080d4e9162d54b13d9b280aab882b45df968e983425161c9"></a>

<a id="canonical-fe851164bf584f520d10eb6b79af6ffb78865f03671a681035be0e37394d746c"></a>

## disable_on_etag_header property — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 6

Type: `"bool"`. Optional.

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

<a id="canonical-1d99ec991b74b3c75b2c0520f3c4297dea7376dde0e465b373b89a8f7abdd97c"></a>

<a id="canonical-2232bb28d1fe3fdbcf0e5deffaeb6733d668afd442718fc9b3665197e420fe9f"></a>

## remove_accept_encoding_header property — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 7

Type: `"bool"`. Optional.

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

<a id="canonical-3031da8d410c53c310137946b2f8d0c3c321837eef097ce935df9a50ca4b3c9a"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.compression_params / 7188c95918a7 / 8

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-86e4db9ad3d146d8a6eabbc9e94032e7d3e0509974b2fa25211fb558b4e9fa46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78d22a304a6145273cc1981c897b6facff945adefe9b52fe83aec460a78ef384"></a>

## dynamic_proxy.https_proxy.more_option.disable_path_normalize — dynamic_proxy.https_proxy.more_option.disable_path_normalize / 896e17980696 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.disable_path_normalize

<a id="canonical-bfabaf4feafcaedba2eaf63374e4db9d8015f535e4072f5c8f6310d87c0655f4"></a>

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
disable_path_normalize = {}
```

<a id="canonical-ca557ce9b924bfbcb2556f2c6d6a68a0821d845fbe6c062d211bbd35c477745e"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.disable_path_normalize / 896e17980696 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e08c98d8a359778339eb4a463454b4d3130277240458ee79f33fbfc5934fc18c"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.disable_path_normalize / 896e17980696 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-85c7bd4210e9ad7b0f8025b69c7d13e8506950f03e9c2ec5ea687dff834057f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc53a94830c5b02ab537427d1a1d706ce2e4754a9434a0f309f8a1650c1c68b4"></a>

## dynamic_proxy.https_proxy.more_option.enable_path_normalize — dynamic_proxy.https_proxy.more_option.enable_path_normalize / a7c1cb279662 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.enable_path_normalize

<a id="canonical-ffcae2af9f51ab79a55332320f6476afff8f0363ba6c04709f451e9b4cc93421"></a>

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
enable_path_normalize = {}
```

<a id="canonical-87b2823a17a23c252dabd7c6d1f322d46a5bdec9812f29a66fc14b65f894ef54"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.enable_path_normalize / a7c1cb279662 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-908ed3e0f9f704b8c271be508acf7f39d692146e480a328b551bbf3d4387dfb2"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.enable_path_normalize / a7c1cb279662 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a49105370a4fcb3c99bf1cf42a1cbbae63e47ec9359a25209fc7766354f3ce1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2df7315e4ba96703314c9e100860eda07458a20a20a9281674b170ec19c38643"></a>

## dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection — dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection / f6514eb81264 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection

<a id="canonical-a2a1ec7e2f95add1eff045afa9d96d463fea1089f70ad48a8444cccc47dcc284"></a>

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

<a id="canonical-c4c93a5424e55a422e3546a08e1b4a4cf10f94d30ea77f6528c403ff4266ac96"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection / f6514eb81264 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a0bef4473fad0ac7e8751c305dcd24c9b1edabec6ebf3dccaff76a3134c57d9"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection / f6514eb81264 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64eb82081602f6ac6a87d7f20f18908444550975b98dac3aa6b59ede3b65901b"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add

<a id="canonical-8d92f6b4b3d166d1198d8ecdc7cecdd88442c9b083731fb40160260fe624d6e8"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2205b50423550b65bdfd46ff9774205e910e9177736198419fa61eecb758feee"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 3

<a id="canonical-2f6759ecb869e8a8f4542ca70ce1ad7678556500a35b7c1c5142fe8d259b593c"></a>

<a id="canonical-f871ae2b6f77b5b1579845b2c57da855432ddeaa6b52ed3ada14b170888c20b8"></a>

## name property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-6285ca628622fa6dace80fe6adab9f202d5945149089d4dc1aeb8db5999e4b59"></a>

<a id="canonical-369501e21748e73619d5ffb091fb9fb82e00f2c5e320b22c433a5428df82efe3"></a>

## overwrite property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 5

Type: `"bool"`. Optional.

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0): complete subsection reference.

<a id="canonical-f1d4ea809e0dcd922aa81ca08b87ea6323017104a4673e2b9ba5b79fd1d5ca00"></a>

<a id="canonical-5970c99992efc9bea6df0dc6c86d3a057aa60109bb1b517d41b778939517bbef"></a>

## value property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

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

<a id="canonical-67657ac4b2a0fcd3661a0b0213d3e5f7510a9b79f397e5e5101867cb0a843010"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_cookies_to_add / 3c2bb7ebb2b1 / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31402707466df747d412eeade2c7608784cc92a84d470c867278f182110fa011"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value / fadfd8e9af75 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-863914f07b423f2a2059cac8875a0fb1ed79f6c1df2d14cf247dc21f7dcd09e3"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-537457ec12cd03f51261e86b17aaf7f1d5671949ed38981bc1ba1da32eac2378"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value / fadfd8e9af75 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-1d87313d9d590e611581ac40b2c8b07286c8c1557a52941a841c1edb3e28ef9d): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-a6d44936c1eeaa514e00676cb4cfba3c369dadfcf076895c35f49e7713164600): complete subsection reference.

<a id="canonical-2a8aac8158e6c91394ef3cab909681a48b698c178777f0e7a1c0307be4fd1b30"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value / fadfd8e9af75 / 4

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-1d87313d9d590e611581ac40b2c8b07286c8c1557a52941a841c1edb3e28ef9d)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-a6d44936c1eeaa514e00676cb4cfba3c369dadfcf076895c35f49e7713164600)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1d87313d9d590e611581ac40b2c8b07286c8c1557a52941a841c1edb3e28ef9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24feeaf7d8af85befaf2657b7d7be9c2b86560fe42d2b692892c937d8de02b33"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0e10d68e221b7d4ed0883ea99aa77c7fdffa7a8cbb98defaaf93445a83e6e450"></a>

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

<a id="canonical-7176c3d6b670113a6018e297d2ac2a60d89cc42ac2316d81607e77bff9bd59ac"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 3

<a id="canonical-da49d0e219b193b891550c18c98ba6827f955dba06e7160a6f6781542dc85384"></a>

<a id="canonical-fe3c62eaddd8f5e725a4fa5b671452f478fb9fea9fe92b3d56d5ff562c19a6cb"></a>

## decryption_provider property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 4

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

<a id="canonical-6ed108ade63a87d8c558128a8ac9e02f7c47b4680f1033191b0e2b53cfeb28a5"></a>

<a id="canonical-be3fe9b6e37d7b1c24e687372f74ea8761ac869385fe1460d3e2f77b76edc822"></a>

## location property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 5

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

<a id="canonical-c94398ae3df12b047785ac0715753c563e8ed136b3425e9ee4cfb36b2533cd8c"></a>

<a id="canonical-6bd71c61fc6645110f4d82c4dbe0870a0109fa8070aa5a2bdbeda71988e59a6f"></a>

## store_provider property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 6

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
