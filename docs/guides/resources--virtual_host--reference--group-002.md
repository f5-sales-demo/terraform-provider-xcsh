---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-74b8dccd82e85dbcb2fd1251d5e3e805f4f13cbccecfd4f0b6d0d4e3d1bd59fe"></a>

## disabled property — buffer_policy / 5351efebb26c / 4

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

<a id="canonical-42ee9331eb65cc1b3b998ace1d6addb88025195a0dafadfb4c7cccc3e3b2a2c6"></a>

<a id="canonical-3db908c7e3f5c455a123a8a443d0c489f47061c143399abdc87f3e0e389817ac"></a>

## max_request_bytes property — buffer_policy / 5351efebb26c / 5

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

<a id="canonical-dad3b61f51fe697f99f7f8ab9a55602a2d83e5a357f644d2fe7da84715fae729"></a>

## Next pages — buffer_policy / 5351efebb26c / 6

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-48dad9d72b604417cdc3ed26a4f083f5a6772b0a7439caf68f30c071b60ac123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf05dcb6234c4f9717e1fe8f48cd9d160ebaa52f0098904d7eafb3e203c7d3bc"></a>

## captcha_challenge — captcha_challenge / c4db09fb23c0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- captcha_challenge

<a id="canonical-71d9e7281db43657da616ad5db03176a7c3ca5be648c834df4a0a2ed18309125"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

OneOf alternatives in this subsection:

- [captcha_challenge](resources--virtual_host--reference--group-002.md#canonical-71d9e7281db43657da616ad5db03176a7c3ca5be648c834df4a0a2ed18309125)
- [js_challenge](resources--virtual_host--reference--group-002.md#canonical-b153e6f9a1401eb6f4bb963290887b7f4e9e32e3f5774d78b7d046370c7a950f)
- [no_challenge](resources--virtual_host--reference--group-002.md#canonical-4b397b8c08ef76993ac481552be2d5bb414ac9624b3a878236aebb1b761f9b4c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-e223939dac34b209e26ff44ecb3cb37c5860382b20c57f6e2593544de175121b"></a>

## Direct properties — captcha_challenge / c4db09fb23c0 / 3

<a id="canonical-38071eb47639120938cbce673ccebec97c492f3ce8c9c5997bfbc484fb9d0413"></a>

<a id="canonical-522d43e4dfa54b7567ec40b3dd0c3d1e4ccfd254da0768ec386bf4a079061773"></a>

## cookie_expiry property — captcha_challenge / c4db09fb23c0 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
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

<a id="canonical-e52128758479e4ab8f53a7ba37d6f68ac097bf24872448529e61123315f1c7d9"></a>

<a id="canonical-6c77821ed6d1951bcd30fd172b065d178c20a30681cb1a21d58ed4c297fde412"></a>

## custom_page property — captcha_challenge / c4db09fb23c0 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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

<a id="canonical-57c2e11e8ef0a442a21d27fdcef3cbf69132fc99348cc8846e02d3759a92619a"></a>

## Next pages — captcha_challenge / c4db09fb23c0 / 6

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4b51e5ca53e22360f5aa7320e8bef962069fac12326158dba8f5298e59a4dd9"></a>

## coalescing_options — coalescing_options / 1e7038081efd / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- coalescing_options

<a id="canonical-d864915ac5cf4d377aaf3beb4101ffdefe5e0f208a4c36701effe4725a78d3d0"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fb2a168fc77833287b13db2b067017a3c59aefc789dc66f563575c0e1087ac3"></a>

## Direct properties — coalescing_options / 1e7038081efd / 3

- [default_coalescing](resources--virtual_host--reference--group-002.md#canonical-8666c6890b3a720f6fcdee0fda2f50678426718d6a03d0f748818b1186b51fe3): complete subsection reference.

- [strict_coalescing](resources--virtual_host--reference--group-002.md#canonical-a6de454e9b91683a18a163799de6473151951ecffa3cac99f793122003b2edcc): complete subsection reference.

<a id="canonical-2f015c7d52749572122a4b607cbf202b8d7ee5c960c46f74b8dd064bcf9dd96a"></a>

## Next pages — coalescing_options / 1e7038081efd / 4

- [coalescing_options.default_coalescing](resources--virtual_host--reference--group-002.md#canonical-8666c6890b3a720f6fcdee0fda2f50678426718d6a03d0f748818b1186b51fe3)
- [coalescing_options.strict_coalescing](resources--virtual_host--reference--group-002.md#canonical-a6de454e9b91683a18a163799de6473151951ecffa3cac99f793122003b2edcc)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-8666c6890b3a720f6fcdee0fda2f50678426718d6a03d0f748818b1186b51fe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3c63f08bb357deb010db1fcd85151263887b3319b2c90b462b18e62d22c85ed"></a>

## coalescing_options.default_coalescing — coalescing_options.default_coalescing / da56d4f451d5 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475)
- coalescing_options.default_coalescing

<a id="canonical-8ad3ef578434d136f52b938bd7ae95e28af32ac19679b5cdce9f2c552d436147"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-3d9a42e9580c8b99e3ae16891e23096919eea7bf7a4b1541b7e7fd627b23c0f2"></a>

## Direct properties — coalescing_options.default_coalescing / da56d4f451d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3a182c8139a476b233803d9b5d272be7a7e7bf4a9c7b90fbe18e22998b67d5f"></a>

## Next pages — coalescing_options.default_coalescing / da56d4f451d5 / 4

- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-a6de454e9b91683a18a163799de6473151951ecffa3cac99f793122003b2edcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-052ae95d7e2168df06cfea761638884e30f73dbd1daf145726d0a0e506fd7e9f"></a>

## coalescing_options.strict_coalescing — coalescing_options.strict_coalescing / 5af97aa52df8 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475)
- coalescing_options.strict_coalescing

<a id="canonical-9e9290b8e194d9a3cef4fecbffe7d40980ec006f6c1317f6b3d3c8e4576ce59d"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-345cc029109927ec6aee0eb54ece884340337cbe657fc3dff5a67b2f101e0740"></a>

## Direct properties — coalescing_options.strict_coalescing / 5af97aa52df8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7299b341c3159b1db0252eb60eeb5782068f2e4b781dde39770c30eb07025f96"></a>

## Next pages — coalescing_options.strict_coalescing / 5af97aa52df8 / 4

- [coalescing_options](resources--virtual_host--reference--group-002.md#canonical-5f4c528aa3ed1b4dbf291db435ba909e5691ec9cd19682efc5c8b262cf42d475)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-4f3d9a57e11ef2bc8d3c145dad841bdfe6ea9138ada5d30a43f2452f7f3e31e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40148c3456303de34f797b43265a0349829294d0037878d31b9f551433f521ee"></a>

## compression_params — compression_params / c685269d2170 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- compression_params

<a id="canonical-0d94d5639ee92675b97a1654fbc25d72b4359d53e4ec3722b4c7103199612e47"></a>

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

<a id="canonical-eb2ab4510959dc8ed48368154784cc97b6fa36e6c08738e5ea0a35ae56b78220"></a>

## Direct properties — compression_params / c685269d2170 / 3

<a id="canonical-31259163bbe50684367325193a5440b660e1f2228035ae988f85eef8b3c1b2b6"></a>

<a id="canonical-6a94d3f16d3c5a81feba5263a54c7a2a564c69e573777a061ca3632156816fb8"></a>

## content_length property — compression_params / c685269d2170 / 4

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

<a id="canonical-147b6e69543b2102b6648d0ab8477c4cca3870fc3de085f4605da45d2d6a6d20"></a>

<a id="canonical-b2901201464ddd3ff26997aa6d0d5e575ffc1b16d84c0fc39902c7a0252243fa"></a>

## content_type property — compression_params / c685269d2170 / 5

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

<a id="canonical-d5cd8841e3c78fb240db05f47b561a050eb8468d31a310cbd41fa65156f086a7"></a>

<a id="canonical-91bde7bb332190a9fa59b2fd4ca2fe179035c2b59f711c5b8c0b2000ed2ed713"></a>

## disable_on_etag_header property — compression_params / c685269d2170 / 6

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

<a id="canonical-191343755cae05f769dc766b2888a4ba53cf6d6dfba407f2078b9849264b7ad1"></a>

<a id="canonical-39f69b287ad2d5811d752059544fa22b486e77ebfd0088a069da3a9488b616ef"></a>

## remove_accept_encoding_header property — compression_params / c685269d2170 / 7

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

<a id="canonical-bc78ba87629ea2c913e0b219af66fe0f123557a91a87d426dfdea9f58d2ad034"></a>

## Next pages — compression_params / c685269d2170 / 8

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ee76b09d5033a064e534dbcf23f0888d145a396376a78774d87c166b278ecd71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96c360a48e8cedd605d6b40f722847456db3f609fb0e09625723dac89a205317"></a>

## cors_policy — cors_policy / 4a73129d4ca6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- cors_policy

<a id="canonical-5c8133a21bd3f08e4013d3be4abef1876bb8c13e9965c482a00eb143ee8dd5d0"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-09562ebdd2c03512e8124eccd26a11ac184cc64cf68352ca132fb771e506b4c4"></a>

## Direct properties — cors_policy / 4a73129d4ca6 / 3

<a id="canonical-ddbe7f41570b57c00464a13380bdb6168fd9c7dd0a8a32c581a7b8b1d5616704"></a>

<a id="canonical-e296bb0e078d78f21dbbb9e8b6172af8c9d930fcc00fabd01fac7de15b59322b"></a>

## allow_credentials property — cors_policy / 4a73129d4ca6 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-1a1edd3426519cdac3026cc6c2226745bed3c895598e14c43170dd0dad9dc72c"></a>

<a id="canonical-fee6f99a143662d213c288be1864e3db8b6d2db91264610b897ba8a4a1fdd6b4"></a>

## allow_headers property — cors_policy / 4a73129d4ca6 / 5

Type: `"string"`. Optional.

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

<a id="canonical-91956e99b04d09e89565eac7e429f1dc34365e62e7b6639659cdf322f297a870"></a>

<a id="canonical-3c4daae6f4ad4714710d608337d7741265b7a7febbf8ebe3f001081a5e1bb5ba"></a>

## allow_methods property — cors_policy / 4a73129d4ca6 / 6

Type: `"string"`. Optional.

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

<a id="canonical-5c8add71fcc0d2145a46c20f821f4bd083172287d225c629e027cdb4c0ee76fd"></a>

<a id="canonical-5e52e5aae4718a382bf7a27a8d68f839ca05bc008b450bbf3f965802083f9017"></a>

## allow_origin property — cors_policy / 4a73129d4ca6 / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-4a184787fd0582c91fcf99337f7e16dd2f8f6b3acf28254e09dc39866f7cf7d0"></a>

<a id="canonical-1f6f3999b8c718c543f535796c28603d628fad7223560cc954344e5919260895"></a>

## allow_origin_regex property — cors_policy / 4a73129d4ca6 / 8

Type: `["list", "string"]`. Optional.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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

<a id="canonical-69685da36e59ad69d20f7c96824e46b64fe8583a667a59ef71136cbf67a15d2c"></a>

<a id="canonical-62ebd84801a2aa6a58bd8f3e580f84d09fdadfb674742e8cac429fb73e84b63f"></a>

## disabled property — cors_policy / 4a73129d4ca6 / 9

Type: `"bool"`. Optional.

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

<a id="canonical-05eb25856b3c56403cb72c2b8078a081de98fb580cc373aa07aff6f54eea232f"></a>

<a id="canonical-77279c8ad37e50a776a66e7d8a6d1463fb351e0ec884e6d82ab19a79c362c40c"></a>

## expose_headers property — cors_policy / 4a73129d4ca6 / 10

Type: `"string"`. Optional.

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

<a id="canonical-65abf21750092dbdd936357f9506d9fb337e69d21559fd43c28e0459f0c1a5b2"></a>

<a id="canonical-ce903bf759b71155426aad8fe57a64bb0ce6ddfeb46951948adc13202d7e9866"></a>

## maximum_age property — cors_policy / 4a73129d4ca6 / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
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

<a id="canonical-6c35ac90c66c5d5129321ca8c83103b58ed0ad1a52698d867db3f65b76f188a8"></a>

## Next pages — cors_policy / 4a73129d4ca6 / 12

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b5b21904cbcfd0cd14d17104520ed9af967225f7a125165397d0bd5dc35df47"></a>

## csrf_policy — csrf_policy / 0823e2bd4dcc / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- csrf_policy

<a id="canonical-c0f4d2fadd53dc83b8c4a38344b84c235385083f39df30a8cbe10e2c5848cd1f"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f6b274a438a7971581916ad491b9d8baee35afeefc975f99af3e26a15ab4aa4"></a>

## Direct properties — csrf_policy / 0823e2bd4dcc / 3

- [all_load_balancer_domains](resources--virtual_host--reference--group-002.md#canonical-146bbc05e2454d936692e7a0db1a3a33060e8de05018cadfc5a030e99d0bf627): complete subsection reference.

- [custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-24cb91d070a035487c3f1953c9f425d8b5f850c990204291f8645c3fcc6baec1): complete subsection reference.

- [disabled](resources--virtual_host--reference--group-002.md#canonical-424c7daa750210eb0733c796066bfa519874ad460cac46ab31e9bb13a4c6e06b): complete subsection reference.

<a id="canonical-fc242f919fc7979a61941c530b53aa28b468334daec8f2aab98cf562ca1acd6f"></a>

## Next pages — csrf_policy / 0823e2bd4dcc / 4

- [csrf_policy.all_load_balancer_domains](resources--virtual_host--reference--group-002.md#canonical-146bbc05e2454d936692e7a0db1a3a33060e8de05018cadfc5a030e99d0bf627)
- [csrf_policy.custom_domain_list](resources--virtual_host--reference--group-002.md#canonical-24cb91d070a035487c3f1953c9f425d8b5f850c990204291f8645c3fcc6baec1)
- [csrf_policy.disabled](resources--virtual_host--reference--group-002.md#canonical-424c7daa750210eb0733c796066bfa519874ad460cac46ab31e9bb13a4c6e06b)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-146bbc05e2454d936692e7a0db1a3a33060e8de05018cadfc5a030e99d0bf627"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e01641fb27ad6de9f3ae0935d37b7797c274cdf4737c3369c484f3523f3e195f"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / 70d7f45d5d48 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- csrf_policy.all_load_balancer_domains

<a id="canonical-df4a57069beb62bc0b5582386d94895cf8a88ca6e96bee2136043e2719fceb77"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_load_balancer_domains = {}
```

<a id="canonical-553538e939fe2602522f0a8475ed7527fe6b64b5eb8c44e42987c11e2c6932d8"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / 70d7f45d5d48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0da9fb4d4f5e0fb7af87f47e196bb907fa74b28eca3014a6ae9036e1cc60d57"></a>

## Next pages — csrf_policy.all_load_balancer_domains / 70d7f45d5d48 / 4

- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-24cb91d070a035487c3f1953c9f425d8b5f850c990204291f8645c3fcc6baec1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29e1b6e422a7813e860a9793ecc6d0fcb64ad9157186239ac59814e1ad03e474"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / eff91f8abca2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- csrf_policy.custom_domain_list

<a id="canonical-caf55fa50d756aa2dd32bf3214040e60ee88799078c8a66ec1be7920a0ecd585"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-25c87fb22f079aeb7dcd3c079e069fba0aeab1a41acc3e12bf4200921ce05938"></a>

## Direct properties — csrf_policy.custom_domain_list / eff91f8abca2 / 3

<a id="canonical-91c12231918375fd28e4def6ac0d269cf1055d8ccfef3c0f0dbefe4e399103df"></a>

<a id="canonical-d156b8c747351c9baf64a5c05036ab08f8b8126ca68e00adf2e7ff85e9d07d69"></a>

## domains property — csrf_policy.custom_domain_list / eff91f8abca2 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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

<a id="canonical-e9954527b49f075337d576daf6cea23584eb17026f74e74b65718e704dc938a4"></a>

## Next pages — csrf_policy.custom_domain_list / eff91f8abca2 / 5

- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-424c7daa750210eb0733c796066bfa519874ad460cac46ab31e9bb13a4c6e06b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcedc8d405643b6b02f72674d3ab6486945b83cfe4500a3cf78eea15aa94223b"></a>

## csrf_policy.disabled — csrf_policy.disabled / c91d1dcb2ab6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- csrf_policy.disabled

<a id="canonical-d94d625805316035b26db52bebdd7b419db1b75d227bffe54a628a797bf74e1a"></a>

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
disabled = {}
```

<a id="canonical-a7aa63d9be41e44c51d53f5a230c0ecc090e37a40ef29d576b07add107dec74f"></a>

## Direct properties — csrf_policy.disabled / c91d1dcb2ab6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-05b28fb356ff7a31b9e4ec0d237ed1aac821b9d67520a626f723fce38693db45"></a>

## Next pages — csrf_policy.disabled / c91d1dcb2ab6 / 4

- [csrf_policy](resources--virtual_host--reference--group-002.md#canonical-0efc46423ef4fc26823c8e68be887cd7b77ed43d03abf0d0df82eb23199d2888)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-7e7c09152bd364c144bc494f521c88bb9c5ea0302e8aaa518b288b68868c63b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c4e079ba21bb829d659deab72ae49d831baa5938e14f3a76204c265b8f0f8c7"></a>

## default_header — default_header / 09dcb8535d8a / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- default_header

<a id="canonical-d3517bd1b5330302df34bc045eb406ab70a915b5dc5b9d063f922f1931a4be85"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-4bfd151524db62aa34ff14f2b4ab2844296000772abf95cadadc0778dbb7fe63"></a>

## Direct properties — default_header / 09dcb8535d8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-616ca4ff2a153327b2bd42966845392e36b4499934d5fa97aff281d89b73477b"></a>

## Next pages — default_header / 09dcb8535d8a / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-d92a03f468c5ca768933d2e076c2e61f6d479736aac01f2973d5793ad0715c36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35541fd584aaff6d586dc1622fc7dd2b4a6b0d04421c3413957a7ebb238e2e54"></a>

## default_loadbalancer — default_loadbalancer / 6fb154d246da / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- default_loadbalancer

<a id="canonical-c0992ecfc4ca73dd753aaf85d2d75ef3edbd5200cebeaf46de2730ebe369ad90"></a>

Type: `["object", {}]`. Optional.

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

- [default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-c0992ecfc4ca73dd753aaf85d2d75ef3edbd5200cebeaf46de2730ebe369ad90)
- [non_default_loadbalancer](resources--virtual_host--reference--group-002.md#canonical-23c848563d952a2c307ae4d364f943917eb3cd1e45b888fdafa93d79e384c51e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-3d95d651b037b72154c142169d43a9a170667b0dbcf2b85c6af41b3f6f855e10"></a>

## Direct properties — default_loadbalancer / 6fb154d246da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-482798d2eee7d4c023dc938106078b2433ff27bd07b60e4e8c95f59ffef6cc56"></a>

## Next pages — default_loadbalancer / 6fb154d246da / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-022ed9122ed6c53cad261165c0a803d0d9b37c5773659acf7af71b9e28331800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1abde4044970d454c7a23edf6e3a1037d550f126416003bc26e2774c002f739c"></a>

## disable_path_normalize — disable_path_normalize / 3c9f57a7b53e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- disable_path_normalize

<a id="canonical-b15ce0b210431e03aef6140a111dcf3f345d3c8c612723ffd0e042a9a4f464c1"></a>

Type: `["object", {}]`. Optional.

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

- [disable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-b15ce0b210431e03aef6140a111dcf3f345d3c8c612723ffd0e042a9a4f464c1)
- [enable_path_normalize](resources--virtual_host--reference--group-002.md#canonical-e84c6ededb183fb94482189c0162c4daad2f0a2f21804b67269d03dc631168e0)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_path_normalize = {}
```

<a id="canonical-adfde2e5f9a07ece6c92d0784555911317a20140b3914dbe1db46ff38d565e11"></a>

## Direct properties — disable_path_normalize / 3c9f57a7b53e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b498cf5109e22ec71fd95048b097cac9b3e01144f83327e732b0292c4a9ea68"></a>

## Next pages — disable_path_normalize / 3c9f57a7b53e / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9f367c54b8943301f0e415a6089876b335b94d6f9d41d982b7c5e774a110fde6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c508820c26a71dc9c7693166cae9356f8b08f878efa2dcc1e97e3ebeb5cc5e17"></a>

## dynamic_reverse_proxy — dynamic_reverse_proxy / f036335c8e86 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- dynamic_reverse_proxy

<a id="canonical-77c60c14a57d0a083ba8717e235f41760c6226ce7b7e10d77628bf437aa959d4"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dynamic_reverse_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad705fb29aa6c2867bc4d36e778b38a67255bee49789a290d5761db35e60373d"></a>

## Direct properties — dynamic_reverse_proxy / f036335c8e86 / 3

<a id="canonical-2b587aee8579a510bc999d7fd90ba6dc3a5d8647f21ce0d48e39b7ab918ea2db"></a>

<a id="canonical-fa0b02a644033eee49b7e302755fa5f0caadba254d503bc55be6d19a8b987420"></a>

## connection_timeout property — dynamic_reverse_proxy / f036335c8e86 / 4

Type: `"number"`. Optional.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

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

- [resolution_network](resources--virtual_host--reference--group-002.md#canonical-5c4ce57f993eb1bbd90764d9b828a271a5cef46a1832075cb48f8b530c50e0eb): complete subsection reference.

<a id="canonical-802d08c2b696e2c417a762cac226c5ebe3cff5822a2d204274f705510cfaad8f"></a>

<a id="canonical-8c7633be7df8ecfcb150b8664d31cd2c35ddbceaceedb655f7d8059fd206b2ef"></a>

## resolution_network_type property — dynamic_reverse_proxy / f036335c8e86 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

<a id="canonical-19b40f1d73c08cf53287ddb6c2389ee9a85d6a20cd41ae9d7fec3d136a22e520"></a>

<a id="canonical-a87309b17ab8abc495327ff95967b8cc9fe2b30efde5db5883592d0741759025"></a>

## resolve_endpoint_dynamically property — dynamic_reverse_proxy / f036335c8e86 / 6

Type: `"bool"`. Optional.

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

<a id="canonical-9d63d1aac35956a45b74beb10ea15c2b3c2eca928c5f276115d874ba0a51f239"></a>

## Next pages — dynamic_reverse_proxy / f036335c8e86 / 7

- [dynamic_reverse_proxy.resolution_network](resources--virtual_host--reference--group-002.md#canonical-5c4ce57f993eb1bbd90764d9b828a271a5cef46a1832075cb48f8b530c50e0eb)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-5c4ce57f993eb1bbd90764d9b828a271a5cef46a1832075cb48f8b530c50e0eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebf76de46306b389cb3fcf4d6efb3b96a2c999d118c957b28d293ca0b34d1836"></a>

## dynamic_reverse_proxy.resolution_network — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-9f367c54b8943301f0e415a6089876b335b94d6f9d41d982b7c5e774a110fde6)
- dynamic_reverse_proxy.resolution_network

<a id="canonical-d2d5cbe4c90baf6495ed73c750e4697a36433908c0c46a531b70caef2fdaca1f"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
resolution_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f0a439935019dd1d3a9eb43ba6ffc75a197fcbe4f0372817dee57b0e7118e68"></a>

## Direct properties — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 3

<a id="canonical-2e6fb3eacde1afa59355252a4e695c6725798546484f52a468cca61a1d37c087"></a>

<a id="canonical-0de3ee6d192e8117e703517e80f233d7179cc07e72b7e3048a42621d8a8354c5"></a>

## kind property — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 4

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

<a id="canonical-82bcf98ed82eb6b6f42dcafa2e8956f5b8b94c285ea276e068cd30d61f3c965d"></a>

<a id="canonical-db74d34b10440735684b3cdcad3b62f078fc425611d897ccd3f511e45edf59e1"></a>

## name property — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 5

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

<a id="canonical-0d1f5257a13ef9005ea2ded6e746f0349975b6acb94ecffeacc677697847ef3f"></a>

<a id="canonical-062bc7048e8af3ba78fe1571c00762313c08c184696863a9ae84f4521a3291a7"></a>

## namespace property — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 6

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

<a id="canonical-1cdeca587ad8ed3d4ed3f7f41f82dbdb3599ebac2880172108180dd7c6da60ad"></a>

<a id="canonical-e4657b423fef4eb08461a9abc42c4cb137eaf2a605b5cbfe50a4f3a86df36f20"></a>

## tenant property — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 7

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

<a id="canonical-05faa2a379dc0eeb81048a1293cfffff679e1760daadb6e26f59d1f942c84dbe"></a>

<a id="canonical-8c88f6e9d3c266647a83f2c2f4c3cf9b21a759cbd59feba67f7ae8e1ea3fe3ad"></a>

## uid property — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 8

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

<a id="canonical-4b4011ab2c0d4c4ce951524b4f75b5623cf056bae37427bfe5be1470891cf928"></a>

## Next pages — dynamic_reverse_proxy.resolution_network / 3b2ecc8b5d9a / 9

- [dynamic_reverse_proxy](resources--virtual_host--reference--group-002.md#canonical-9f367c54b8943301f0e415a6089876b335b94d6f9d41d982b7c5e774a110fde6)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-6a5f867f80659e50963f8f154bef91391a4c75b4124b8d0d73e2fb3cb58da05c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a75905a0cc08de10cacbeea87488635cbd6d54536ced126f6fd38a171694cce"></a>

## enable_path_normalize — enable_path_normalize / 73f836988dc2 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- enable_path_normalize

<a id="canonical-e84c6ededb183fb94482189c0162c4daad2f0a2f21804b67269d03dc631168e0"></a>

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

<a id="canonical-6d8c0793a61db005220a3ea1fe94f9df1af39b48e10dfee29cc6ca75f2b11def"></a>

## Direct properties — enable_path_normalize / 73f836988dc2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59b99d4b46b8c49f14a06e9eb8b46d016ece3664aa3bc6a26b29e650ffa1fdbb"></a>

## Next pages — enable_path_normalize / 73f836988dc2 / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-083477a2c78c2e909ab6522a1dae3c033dd6566a8875cc00e7e27c63e9dc4123"></a>

## http_protocol_options — http_protocol_options / a9c4361cc5ab / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- http_protocol_options

<a id="canonical-19820efcc037e4a9286c48ed8abe9523063812393871c30e76b76c8d2f37b859"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb46fe992ea6b6aef70be0d051133ac995f7c84f103ab70e6f68df1bcfeca075"></a>

## Direct properties — http_protocol_options / a9c4361cc5ab / 3

- [http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--virtual_host--reference--group-002.md#canonical-0d87b7edbbbd3d15835e9d5ff101e286a44695169eb9f219a8dfbd196c73b3a3): complete subsection reference.

- [http_protocol_enable_v2_only](resources--virtual_host--reference--group-002.md#canonical-dec10c871c77edd1c8d864a829ed58d14deedc4be762c114100c40a62b6951c9): complete subsection reference.

<a id="canonical-e8ba1038b83f7e22dbd85fcfec2a544c3aabd756c5ac2755d782029d562cf350"></a>

## Next pages — http_protocol_options / a9c4361cc5ab / 4

- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- [http_protocol_options.http_protocol_enable_v1_v2](resources--virtual_host--reference--group-002.md#canonical-0d87b7edbbbd3d15835e9d5ff101e286a44695169eb9f219a8dfbd196c73b3a3)
- [http_protocol_options.http_protocol_enable_v2_only](resources--virtual_host--reference--group-002.md#canonical-dec10c871c77edd1c8d864a829ed58d14deedc4be762c114100c40a62b6951c9)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a452e127d43c04583e524e73b00ab581d04215d038ab5d227b15a902abcb4574"></a>

## http_protocol_options.http_protocol_enable_v1_only — http_protocol_options.http_protocol_enable_v1_only / 9f1266536636 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-fcd8be56fcdfe99c187c5941d3d89850c0f9122dcf264ec41d2ecf477afa8c94"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-714367448178c093fa8b70008a301489bfdf5d6c847fe5d0e21fb7b784d95c11"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only / 9f1266536636 / 3

- [header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554): complete subsection reference.

<a id="canonical-d2cc4e3bda3b01e7500bbfe204323359ec3b7f823157ca067a92c5a1546aa535"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only / 9f1266536636 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3de28a0b49b358d927c4980d735615644f98f6b2120014e3df79642afb929e5"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 46e43c328070 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-ab4361e73e8d0eac5b2b4868762a4ffdd9d136005a703ec656ab1f9cdb6c5d5b"></a>

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

<a id="canonical-142bd00c5dbc5d820cf3d2c557d5c0ecec92961650b65ff7a4c12c191aa7ba1e"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 46e43c328070 / 3

- [default_header_transformation](resources--virtual_host--reference--group-002.md#canonical-c9f9d91be383ef39892a2a88e6d28ee4e515e6559cde0848131b6d3afedbd1a8): complete subsection reference.

- [preserve_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-53e4b1f63aa07fb2b5e453bd6be92373db9d7b0d35f7f2aa753a6c2e03802ffb): complete subsection reference.

- [proper_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-70f3495b42a5326d7ea23c7b3a95ece40cb837e6e88487b02e627f684b38743c): complete subsection reference.

<a id="canonical-851506ae2887f7b72fad449d159395684e9c4cf1852da4e5494119bd8b6c20c1"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation / 46e43c328070 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--virtual_host--reference--group-002.md#canonical-c9f9d91be383ef39892a2a88e6d28ee4e515e6559cde0848131b6d3afedbd1a8)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-53e4b1f63aa07fb2b5e453bd6be92373db9d7b0d35f7f2aa753a6c2e03802ffb)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--virtual_host--reference--group-002.md#canonical-70f3495b42a5326d7ea23c7b3a95ece40cb837e6e88487b02e627f684b38743c)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c9f9d91be383ef39892a2a88e6d28ee4e515e6559cde0848131b6d3afedbd1a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57b74b3527ad89b5e5c1d892863aa6c26976dd939f09e71b70202462eeab1379"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / 239c2b2855bd / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-6558312a5e932236e09e58cbeb7abafbddd85fcd7b1a08bdc8bfa3b33998e7b2"></a>

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

<a id="canonical-66a82fd622ece9500a4bec117de020a58c26ea6091920cc705d68c827ea7fab4"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / 239c2b2855bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0e7776c037fd892630d073cc5ed476f8c2dd8a09f9558cba04e75f2ca2f0c36"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.default / 239c2b2855bd / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-53e4b1f63aa07fb2b5e453bd6be92373db9d7b0d35f7f2aa753a6c2e03802ffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20598617361798cf7c5ab8d3fde699b2748f3581330f09e3edd558604a53651"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / 938955c9769d / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-166b46b61c666b70336218c2238cb13ab2e3be1b360fa4894c48eeb6c0f72671"></a>

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

<a id="canonical-c5176178bd1e610cf39d3fb75e1be0b56d628e49c65dcb58ba757cf30d373072"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / 938955c9769d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2fa6a18c6928e749b917f68202d973463141c5d0219ff38ac2b53afa844bca2d"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserv / 938955c9769d / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-70f3495b42a5326d7ea23c7b3a95ece40cb837e6e88487b02e627f684b38743c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d62bb40d1571f3eb0ac330a50ed57d49becf95f69be64beb6cf4a2eedc692d"></a>

## http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / 23d492422df1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [http_protocol_options.http_protocol_enable_v1_only](resources--virtual_host--reference--group-002.md#canonical-e4292f4c52c5a8ef4e6844ebf1ca0f28dfc97ed43395d0973b93c20cbe025a4c)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-ff414ac6113ba2ec323fa906e05581ea4c679c23abedfcd2631bdd6eda5c1a15"></a>

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

<a id="canonical-2a9e73cc248e6c0ed980cac1ce098602209dd8c9640cbef343b6c9f1f8875307"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / 23d492422df1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5652133cf77cf98dd45e94b33e9a6acfd8a0b65bff58780d5f697934e57eacf1"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_ / 23d492422df1 / 4

- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--virtual_host--reference--group-002.md#canonical-f528d33f44fe525c587248188278189d9312bc96e5624c4ba4a1d6415d201554)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-0d87b7edbbbd3d15835e9d5ff101e286a44695169eb9f219a8dfbd196c73b3a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9069970576b42ec45adcf4fa19a413a6a147fbe268fc9a5d99075b27f632e98f"></a>

## http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_options.http_protocol_enable_v1_v2 / 29d855870147 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-033da88ec6584e96dc4209572acd9e365ebfbaf6a2b36bdd0fafe82dabbde871"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-97a26bd816ebc0c0eb540a2a6abbca0ab511872d6cfd77fbb58cf3f6712c067f"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v1_v2 / 29d855870147 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66e76e5e4d6ef9f2b7af532ec5178d099f3795e932de21b1b3712fd7c6e09e3c"></a>

## Next pages — http_protocol_options.http_protocol_enable_v1_v2 / 29d855870147 / 4

- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-dec10c871c77edd1c8d864a829ed58d14deedc4be762c114100c40a62b6951c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db389347a965e6f34823964daed319d6eeec20960628bb3862d8202bb499c790"></a>

## http_protocol_options.http_protocol_enable_v2_only — http_protocol_options.http_protocol_enable_v2_only / 725e217557df / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-b02a6bcb3a929d143103ddbdc18f48db9fb8be6d816e4149519b72935bbd4501"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

<a id="canonical-5f6127eacc77fd622cf58f42e8ead9eee846f4db20e24a9b5d4c09b0eca9f39d"></a>

## Direct properties — http_protocol_options.http_protocol_enable_v2_only / 725e217557df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-41fb833d369debefc3a6396628d2e3f4f265a2a454c2aedde9a77b602f1a2b6d"></a>

## Next pages — http_protocol_options.http_protocol_enable_v2_only / 725e217557df / 4

- [http_protocol_options](resources--virtual_host--reference--group-002.md#canonical-268ac157a37a8bd17977cf048c9893badf2c2aa09fbbc91ff6b3e4ffdb57811f)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-11b72bc185ad52ef017b9eee296b6915e0b52fc8f5a62246894d92061978e8db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45ba9af82f887a78f9ce6699b709579684e259a0761336f728c520dee29f2282"></a>

## js_challenge — js_challenge / 774511976156 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- js_challenge

<a id="canonical-b153e6f9a1401eb6f4bb963290887b7f4e9e32e3f5774d78b7d046370c7a950f"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-81db8b7dc44470991e2a57e48a9a7f4b3635cc869183e43d7a5854c34aa48395"></a>

## Direct properties — js_challenge / 774511976156 / 3

<a id="canonical-a0590ffcb3bc05e7e678d87826ce229298611ee183979c37feecb2ed50996365"></a>

<a id="canonical-8924ee064c81ce1517822b77ad84e7a203b697d59752eba600c347db2456a39a"></a>

## cookie_expiry property — js_challenge / 774511976156 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
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

<a id="canonical-8cb150343847436d3988fdaa3f62c17b92d4be7946505dead7b72159d853b3fe"></a>

<a id="canonical-9b5dd65e20f1d97a776d00f0d8369d55844dbb5961926256a277145744d89369"></a>

## custom_page property — js_challenge / 774511976156 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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

<a id="canonical-1b1c102a9f717a6df099e1e0c8769936bb365fe2d7f7b8e6c22da473b0b4152a"></a>

<a id="canonical-ecce65e24dcf82641cd7fbef7a20be7338f1b3e3b7f62b832979d3506219c8e6"></a>

## js_script_delay property — js_challenge / 774511976156 / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

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

<a id="canonical-75257cdfe6c6261048c6d9ca69cb45dfd0766bb736c708d5ea963bbca34e0b03"></a>

## Next pages — js_challenge / 774511976156 / 7

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c7e0df30d5d7698fa774bb1754d169c8cc175e23854bc454545eb10dd939da9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfe5fb5e39c55e90f001685eea0faa2bec210d5c7c52258a429ecd0b82abcd59"></a>

## no_authentication — no_authentication / 5eceeb19df46 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- no_authentication

<a id="canonical-13b5e85003583a216e5ac801072c91108bedac112a998229fb56dbad837663b9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_authentication = {}
```

<a id="canonical-dab3f3602455b288ce22c6eff834955be1f842a396b9748da67e9b92e8e3350e"></a>

## Direct properties — no_authentication / 5eceeb19df46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d441207bb200abc91796358ad298c6c0580e14c5273f331f980f8ece438987d"></a>

## Next pages — no_authentication / 5eceeb19df46 / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-0eb2f80ed7b092db4204ce9381a1c03fc429243b175fa8ba9e9858ce9eb533ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aaf8b743d97c5c53e0a3072ba1cdab5d6033f4aebd7cd56635654c5108a2d4f3"></a>

## no_challenge — no_challenge / 5d1d0df3992e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- no_challenge

<a id="canonical-4b397b8c08ef76993ac481552be2d5bb414ac9624b3a878236aebb1b761f9b4c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_challenge = {}
```

<a id="canonical-aa831d174677286e13dfb22a59722baad7385b00ca042849b4b4c5ddd6e80a51"></a>

## Direct properties — no_challenge / 5d1d0df3992e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85e83163de81bebc6f212cc6960d0d6a33c2ce1ec0dd3a0fdb247366dc1bab88"></a>

## Next pages — no_challenge / 5d1d0df3992e / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-6126e97857379c252caebeaf8772e51c7056956152579ddb07a9300d64c4052c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a929360bf0d248a4cee637a5a5359776c8d99145faf1050e7bcf0a0819ba4d1"></a>

## no_request_limit_per_connection — no_request_limit_per_connection / d16ec748c55f / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- no_request_limit_per_connection

<a id="canonical-68bc0e78a1914b2fca2c9c77dda8ff8cc9ed1a56b881e907164718b5308a1697"></a>

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

<a id="canonical-291727cce7474bd528dd330bd11026e250411c7f7a42f3bbc76a11cd27723e6a"></a>

## Direct properties — no_request_limit_per_connection / d16ec748c55f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91cc35e19bcef841b6efdec35d6112fe6cf8f6a4c8caf01805330c78d1466b4f"></a>

## Next pages — no_request_limit_per_connection / d16ec748c55f / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9359da92c433a78ae5abf5c6f1b8a7b7386baf714874b872ba69c74b64b62989"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4399efcbe7122c551c248cfd4b9b37eae2ff049ff68e8689b5648abcc55be0d"></a>

## non_default_loadbalancer — non_default_loadbalancer / 35e8db4ed83e / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- non_default_loadbalancer

<a id="canonical-23c848563d952a2c307ae4d364f943917eb3cd1e45b888fdafa93d79e384c51e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

<a id="canonical-cc109e118c5d5dd8255c00a15f3e051b5e72ed2840dec91723f1c9abc95b5ba0"></a>

## Direct properties — non_default_loadbalancer / 35e8db4ed83e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7d6f9d9a07912a636c466816296a5def6bc6a6a4c6d9d05e3e60f27f7d7c9af"></a>

## Next pages — non_default_loadbalancer / 35e8db4ed83e / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-3b0a98bd07e1b03da8584b4df7e151fde171e721bb49956600c642a226337ed3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d11455599499373d255fbf352f0de69f3c37c6746d1015db227dd5b9d61803f0"></a>

## pass_through — pass_through / 2a7d9f911929 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- pass_through

<a id="canonical-1e592fae1bf8eabc3123b454db28a7736ef8eb9433c7b31d1763e1e457d09fe2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

<a id="canonical-c59320fc59d70d401aa62ace886d4d53926ef1660b654d3bfaf817ccbc923a57"></a>

## Direct properties — pass_through / 2a7d9f911929 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11f5fdeb3bab22e28bc829a47d7f9d832ed9588f4ce724f8878898dbca25a15f"></a>

## Next pages — pass_through / 2a7d9f911929 / 4

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-461ae2d80e6a10ff94c81ee92edf8b528dfc035a1510afb918deb5d3b92610a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1da091d8fcd113d2f641fb225fa1600eca1153543f14df0123cdf32e0ebe7ae5"></a>

## rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- rate_limiter_allowed_prefixes

<a id="canonical-f1765662fe555b07979b10d8899ee4e402238c3c10599976c4aeef2196b6f15c"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-832aabb51bdac5be501dd585025c1144ef86a4a8c94c79add9e39971a104f595"></a>

## Direct properties — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 3

<a id="canonical-5616dc12f04047834ddd68a88895f8da833678f6f319bdede20c0c2f3e5a3700"></a>

<a id="canonical-52c09fcf1d96908cabc35e70dafca463de849a44d5a0e20119969d3877a4a9a3"></a>

## kind property — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 4

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

<a id="canonical-1e689d78cbbeab12414606f4c2fd9c7cad7917f20975a658272b3c8c4ab4efd3"></a>

<a id="canonical-7a7b0c76cfe48fbc2629c880ba75824f85ce410e48faa0816853ffc2abf3a402"></a>

## name property — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 5

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

<a id="canonical-9fd8c48820a3471aae19878c33bf95f60e2a3f64a68fb438c8dddf055e8300ab"></a>

<a id="canonical-5f8d513919cae6886bf7793fb15f6689167a3d6c255bc3c623a52e13f5570e8a"></a>

## namespace property — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 6

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

<a id="canonical-d40b393f8f163db734f09becd8d60b03d24e394b07806193e13367a7fe9f90eb"></a>

<a id="canonical-7787cdcd82b6345cb8aa5a20b0a39237527e4c89d2f72fd3cabf980549cb71d3"></a>

## tenant property — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 7

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

<a id="canonical-6753081bd0052fc66204235eeddebd526d39939414c7998f7b5286b014642c6d"></a>

<a id="canonical-02a49a5fed00e3d2476e521c4547b883bb9a80df90fd112c0319d39b84a9f709"></a>

## uid property — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 8

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

<a id="canonical-e83e6db64c5df70e70e56ee993ac7d3695c4296963978d0e16251bf73e20df38"></a>

## Next pages — rate_limiter_allowed_prefixes / d0ecdb09fbf6 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9c9441bd1c64d695088dfbb04ed657857550e50f642d7fc008fb2321c476fc5"></a>

## request_cookies_to_add — request_cookies_to_add / 5f8be85cbe26 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- request_cookies_to_add

<a id="canonical-5e82268618b708b5588b1b8c1f2d06783748ec16733b7e4eef6fe6e89bea6047"></a>

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

<a id="canonical-784c433466fc7a7cf38a25915580326bbddbe5c21a15d1cd57ab819bb279690a"></a>

## Direct properties — request_cookies_to_add / 5f8be85cbe26 / 3

<a id="canonical-d4ff4d3abef1b7c9b959b4f679415f969271d07436f092fb47a1507e5a911f5a"></a>

<a id="canonical-e87fa18dff506694599d5e16579eb90e9c06144b9de5700c754820529b689965"></a>

## name property — request_cookies_to_add / 5f8be85cbe26 / 4

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

<a id="canonical-3e43fe65fd7ee86cd6a092b7fd9d2c29a8d0186e6daf4b67e54eb2a4c7080fc3"></a>

<a id="canonical-ea8046decdf311b4f501c077c94342c8bbfb3e85fe4c7bfa94f4acc2673f3a73"></a>

## overwrite property — request_cookies_to_add / 5f8be85cbe26 / 5

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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13): complete subsection reference.

<a id="canonical-76c2e476f9e31247b349340204511102e4afed3844a8434e7c82c73a8302c523"></a>

<a id="canonical-5f76db893001d2f6da1e0c892534441992b51057f3ba90b8658d03a7be42367b"></a>

## value property — request_cookies_to_add / 5f8be85cbe26 / 6

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

<a id="canonical-069162044fb7efb734477de8830c996c4746c9ab656d19899adf0818a39cb396"></a>

## Next pages — request_cookies_to_add / 5f8be85cbe26 / 7

- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc0e1b6f2f48d9cddb13ac1b8a333e51a102b2c4d678e882f6bc8823a5e5c7fb"></a>

## request_cookies_to_add.secret_value — request_cookies_to_add.secret_value / 86c01ed514d6 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536)
- request_cookies_to_add.secret_value

<a id="canonical-c0ce90cc5d5691600e6e04c7ae7de1777fb7d72e4cc2adc3cbbc2e6027581891"></a>

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

<a id="canonical-abd7f24dd3ac9ac9c5e5bfd4d9c7e2321ae951647c333c161aefe8f232ecd928"></a>

## Direct properties — request_cookies_to_add.secret_value / 86c01ed514d6 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-de1e5df912907368e3f42353c8cc5fd312687654c60c8d4ba9edc9ab8d4a3e50): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-806957e0d4d1cd7fd580b2bb40a399cb47c95fc751179450bc820ecae258ec62): complete subsection reference.

<a id="canonical-d7b53cb1bdb7a0aeaee8a155261bd795e87a3703f05fc9da144aa7b8d3adbb16"></a>

## Next pages — request_cookies_to_add.secret_value / 86c01ed514d6 / 4

- [request_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-de1e5df912907368e3f42353c8cc5fd312687654c60c8d4ba9edc9ab8d4a3e50)
- [request_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-806957e0d4d1cd7fd580b2bb40a399cb47c95fc751179450bc820ecae258ec62)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-de1e5df912907368e3f42353c8cc5fd312687654c60c8d4ba9edc9ab8d4a3e50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab1d7f55ebbac18f5f333dfe1fd28712eadc12332e71c7bf3ec3551f6cc1039e"></a>

## request_cookies_to_add.secret_value.blindfold_secret_info — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536)
- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13)
- request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-5e9957e314b0e154ec1991275db07326e717a1528175261bb2cbb9ad540979d5"></a>

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

<a id="canonical-94b2a7cb638dc4abd617a0f0b3db031852f9b42b6936ac88f9dd97065bddfd84"></a>

## Direct properties — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 3

<a id="canonical-7b099b1fe5061dd5602f729b29519fa14fb28f61cafc2390d140d19c6b93618c"></a>

<a id="canonical-72898c5715d1309b2585d397b0a615f08c5334388ac877398bd839b63d5ac224"></a>

## decryption_provider property — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 4

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

<a id="canonical-eabd1300b96a64a9aae3bdc82d4914927cf749247056a25ea02cf9bf18382ded"></a>

<a id="canonical-9ca7cfd3bcee3e95c2cbcce7a09445924ccd3dace9624fb186900f2f97fd2a1b"></a>

## location property — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 5

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

<a id="canonical-59ea29ff434ca95d663dfd114da039dbd72c0fa827b1e8b5807437d988d16016"></a>

<a id="canonical-92a22d6ddf1e7194a28fc5e4c1f1b32f4dd095cf5e85f318b212917f74ef290e"></a>

## store_provider property — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 6

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

<a id="canonical-57ba4b8f273db37b7025f7fbfa95252614b8f775fc8993b76a674bdd204c1b7b"></a>

## Next pages — request_cookies_to_add.secret_value.blindfold_secret_info / 350e5516f196 / 7

- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-806957e0d4d1cd7fd580b2bb40a399cb47c95fc751179450bc820ecae258ec62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17be16eb130eddb5ff28dfc449fc7e845577002c8fefefc3afc2e200420f91bc"></a>

## request_cookies_to_add.secret_value.clear_secret_info — request_cookies_to_add.secret_value.clear_secret_info / 04a0cc843f25 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1574a375122f042b98d7d2a028b2b752f8280d81499b0f51618c56357b193536)
- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13)
- request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-d57d1832d6298fdf28f616b9096bd11d1c07684795370a8be102fe7104271e7b"></a>

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

<a id="canonical-49c53fc1a6848a1d02499aed64d82e268f97d7cf1c3017d2de4d4bfe8a0c15a5"></a>

## Direct properties — request_cookies_to_add.secret_value.clear_secret_info / 04a0cc843f25 / 3

<a id="canonical-3f66e1f5884613f26139ee91502bafd84752fd094cb9dbd6d584051c70912665"></a>

<a id="canonical-ed8b4c73b9746ef99979891e29363175036a0d66b3eaa0f83ced51ee1bc71700"></a>

## provider_ref property — request_cookies_to_add.secret_value.clear_secret_info / 04a0cc843f25 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a3af7155310b2209f8f7256424208703894307a03fe65c6024f42a6fb3c589b7"></a>

<a id="canonical-8a74a52c2598705bc30be2246a163f8c902fae9fd367d9f7deecdafdd09e41da"></a>

## url property — request_cookies_to_add.secret_value.clear_secret_info / 04a0cc843f25 / 5

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

<a id="canonical-19d962064695367ad24fa112d842384f6a1304817d9a56511cc4bb5e60e9ff7f"></a>

## Next pages — request_cookies_to_add.secret_value.clear_secret_info / 04a0cc843f25 / 6

- [request_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-4274cf0e1466e26a139710d7b822d61a4048f955dceb500923ca09b11067ce13)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d226ccc48993b4df0e920ab022c8c8d9d4dd95d2c9b630f9ce8b3e99f3f7dad"></a>

## request_headers_to_add — request_headers_to_add / 4b40dabdb397 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- request_headers_to_add

<a id="canonical-c6debe68d96130a8c25541339d089eb45c40f3203fab02a8b30cd77e6f140187"></a>

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

<a id="canonical-3d703439f879b3f210156aac6faf56ee036671703d6d1af359204634efb4e493"></a>

## Direct properties — request_headers_to_add / 4b40dabdb397 / 3

<a id="canonical-84ec30612d462565985f1d33f67d0441a1cc8a46be3c1267f28dbbea46a0f843"></a>

<a id="canonical-a1e1167d7d3056a372614352c56ab98b354d216d39e3e25f69d0ef860a652c2d"></a>

## append property — request_headers_to_add / 4b40dabdb397 / 4

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

<a id="canonical-63bc2676e5e5c681e0c6bb5b14b289d8c72981e1aa9f0379d8c5d66b7a1d6710"></a>

<a id="canonical-57a3e729ae02a508ee9c670ace0f22aa3e73d75965ed2557b3b7ab4c6acd09fa"></a>

## name property — request_headers_to_add / 4b40dabdb397 / 5

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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39): complete subsection reference.

<a id="canonical-07bf99e8a686c89eee9505fe3a7dd14eafc3ae0979ff6110bf41382a49e9a78b"></a>

<a id="canonical-84f38ecca7ecf56372ac5464767ea5aa471435299609ccda7d99eadbefe2a56d"></a>

## value property — request_headers_to_add / 4b40dabdb397 / 6

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

<a id="canonical-a94f7fda6a9f8195e06e4d27faa543a863cdcde075e9aa130b2c21be7e68666c"></a>

## Next pages — request_headers_to_add / 4b40dabdb397 / 7

- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbec0e1587a6532fd68ed5ae7207e12b1de51b35c0d220bdd9bb8aaeb7ef8677"></a>

## request_headers_to_add.secret_value — request_headers_to_add.secret_value / ecfe67af92ad / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969)
- request_headers_to_add.secret_value

<a id="canonical-eeb22f339c06c6d25e5c5a2e8bc62431bb67aa7b3c3a7c54309b3144a1b92045"></a>

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

<a id="canonical-6d513e892ccab0823f4bab1b291db681eb2303ee93bd88f30f05a07a7e19bda7"></a>

## Direct properties — request_headers_to_add.secret_value / ecfe67af92ad / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-bded8e8a8889af833b718d0c9a6a15f6b977cbd73e72b6c77245387e3fe0b423): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-4177eb77332f305939749636b8713ef512fe23ef716a8249f6b48d4718b85c3d): complete subsection reference.

<a id="canonical-6a8b9583b5de5c7b51dc0fa86b785696f9468b347170db07a41893288152d656"></a>

## Next pages — request_headers_to_add.secret_value / ecfe67af92ad / 4

- [request_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-bded8e8a8889af833b718d0c9a6a15f6b977cbd73e72b6c77245387e3fe0b423)
- [request_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-4177eb77332f305939749636b8713ef512fe23ef716a8249f6b48d4718b85c3d)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-bded8e8a8889af833b718d0c9a6a15f6b977cbd73e72b6c77245387e3fe0b423"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71bb73fb4e30a9cdf6278bb2c6f2a5a1ae91f6d3ebfa75cc40baacadad4595f6"></a>

## request_headers_to_add.secret_value.blindfold_secret_info — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969)
- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39)
- request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-b74dd2d7450ac25ea744a2d7099f532ed70b3ba62846f60e4ed54fcacc8a5a6f"></a>

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

<a id="canonical-22484f5e7a35a52408202388ce6101157c4ecfc598acfa7a8fcdcc9670080092"></a>

## Direct properties — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 3

<a id="canonical-6fc812e6baa6d97684aa4899e83c7eff79de118d962e17bf32459cd32f4a1fdd"></a>

<a id="canonical-bb821afae5a45260dd4d8aa90bc9905e136788c0dcf9428bf64e4190988ad7aa"></a>

## decryption_provider property — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 4

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

<a id="canonical-f4cb2969084904d7f369f4c7eca62ddb58584bb4f9ff9bf32f7f86a2bfdb4f56"></a>

<a id="canonical-28563c70691f06aef1b916e7708d2b641c941e747b57d99afabc2010d40064f8"></a>

## location property — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 5

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

<a id="canonical-d294f7b1107fadb78bd641a0040181fe65388d4426cfa421833035f09acb019c"></a>

<a id="canonical-cb47d662de793d17a9ebecffa4c3ef55eeec222c38fee3df0d0c9c9cf859e5c2"></a>

## store_provider property — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 6

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

<a id="canonical-d3d83874c4ba9140c9cfa0b1629cf688097c3b3ce25aff487d8b113d29aad6f0"></a>

## Next pages — request_headers_to_add.secret_value.blindfold_secret_info / 3fefacbd11eb / 7

- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-4177eb77332f305939749636b8713ef512fe23ef716a8249f6b48d4718b85c3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92f8eca733d035f2135c81b013120189ebffa25a2fc17ce78b8a92926ee4fbc0"></a>

## request_headers_to_add.secret_value.clear_secret_info — request_headers_to_add.secret_value.clear_secret_info / 6020ad60748c / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [request_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-cb41615eb2310e078c40712816ffcb0f146606d2867237b529be4b01a6ce8969)
- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39)
- request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-69991c6616646bbf8fc11db60a87e62d16af64982fcba7f66bd9d602fedce56c"></a>

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

<a id="canonical-76f9ce157e21469388f5e8da3242abc6f21525f2947227ce7e8f898c18eed478"></a>

## Direct properties — request_headers_to_add.secret_value.clear_secret_info / 6020ad60748c / 3

<a id="canonical-9609fd713b1f5bc7df53bc3a2674d86b6f2a5047a06a4338da046d624be75d8c"></a>

<a id="canonical-fb14e422197ecef06a96eca6b511c1e3e4a63a24de3fa2664898001962109c51"></a>

## provider_ref property — request_headers_to_add.secret_value.clear_secret_info / 6020ad60748c / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bb96a57db04b0732854be69825969a7a1e05130de8ddc31b217818f61c1a47e3"></a>

<a id="canonical-89e1a299fc8d7a555dd3517fc9c9a33cd095c80584086d1aee2a016407d42c48"></a>

## url property — request_headers_to_add.secret_value.clear_secret_info / 6020ad60748c / 5

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

<a id="canonical-a6c7dc172c08149d316279be5b573ac7bd123b2d4e6fc7749feb0d1f88e44b1e"></a>

## Next pages — request_headers_to_add.secret_value.clear_secret_info / 6020ad60748c / 6

- [request_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-b12b431a85cb9701667bf503268013bc563f31f2d02cb284925c5bdb49c83d39)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-674bd581a732e9e1189ee48ec62dd24cca388f54322a54edb3dcc59e1c2724ce"></a>

## response_cookies_to_add — response_cookies_to_add / 51f05a6eaf20 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- response_cookies_to_add

<a id="canonical-aa0c5c4ce2a8d71f70d5ee670298dc0e6e9d0a9055f85fc832dadb6074c33a3e"></a>

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

<a id="canonical-7a0e5565ddb29eac3bfa9c6654e43a545688eb0539ec67bee613bc994733732d"></a>

## Direct properties — response_cookies_to_add / 51f05a6eaf20 / 3

<a id="canonical-89159365e42f4cb3888c656529631413e644b8c99c49996dbaf7634acc4bcd95"></a>

<a id="canonical-5f4477c6b2182520eff0f14ece49bc28c34c6cba56c024db779a17849fb7aa19"></a>

## add_domain property — response_cookies_to_add / 51f05a6eaf20 / 4

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

<a id="canonical-7eaed9ee5683fbb5e4c48f80e89aa7ac84d6b8eb1873874eaf61456364e0f4b0"></a>

<a id="canonical-4bcc0efe8cb0f223e5997aa9dc540643dea9e9feff41208faa3bc19e278df20a"></a>

## add_expiry property — response_cookies_to_add / 51f05a6eaf20 / 5

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

- [add_httponly](resources--virtual_host--reference--group-002.md#canonical-19f32417d81d7298854959ae5db61cf1b10863b8f0638427ee063e5cdbb22f7d): complete subsection reference.

- [add_partitioned](resources--virtual_host--reference--group-002.md#canonical-8b0b03f2fd1f80019347e2f0e75ab4702f0deb1345751ced3d6b5f620fa530f9): complete subsection reference.

<a id="canonical-50388e25685403efaa30a217587632c849e191262c34713467d83989661e05e3"></a>

<a id="canonical-c7b14f965791f116629a1c991c98f9cc1b0bff169cbaf8a7be93a8356cec4cb3"></a>

## add_path property — response_cookies_to_add / 51f05a6eaf20 / 6

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

- [add_secure](resources--virtual_host--reference--group-002.md#canonical-a8cad4e4c8f9caf2fb1aedebec93c4c5e433593a84059ebb2724b5e6a6b3b499): complete subsection reference.

- [ignore_domain](resources--virtual_host--reference--group-002.md#canonical-e14373fdc9c8cfb84d289951160b5f3652014057bb099e9cf1e920f1528bac26): complete subsection reference.

- [ignore_expiry](resources--virtual_host--reference--group-002.md#canonical-249d7e7b0a123c956421e57f3be8b3269e4931e03a350b15923e27f85e448ca0): complete subsection reference.

- [ignore_httponly](resources--virtual_host--reference--group-002.md#canonical-c9240233de110d27eebfbec16e51991e2a0da789ddfae7c379e723a069c0ce62): complete subsection reference.

- [ignore_max_age](resources--virtual_host--reference--group-002.md#canonical-cc4659c786d3c76d56524ac7f9d814c3cc700c497d756bfbc53be2101c08e60d): complete subsection reference.

- [ignore_partitioned](resources--virtual_host--reference--group-002.md#canonical-3b3a49d56f482e1fab35880702db7fdf679184201d665b6093cc4e714ed732e0): complete subsection reference.

- [ignore_path](resources--virtual_host--reference--group-002.md#canonical-1d16344d6e339d1ce0f61cedb96ff9f836bfca776ebbac7b308a8e1337b77df8): complete subsection reference.

- [ignore_samesite](resources--virtual_host--reference--group-002.md#canonical-f41b08ca9afb2ac17e384112b38f8edd3840ec55f84ef7ea75bd6fce9cf3b454): complete subsection reference.

- [ignore_secure](resources--virtual_host--reference--group-002.md#canonical-f5ebc7c702a9b26830957419174bd852361ff12a45e5d7bd064e4b6a54a989f2): complete subsection reference.

- [ignore_value](resources--virtual_host--reference--group-002.md#canonical-c87b1c50be423a75254c67b99929b2425ad6b87dacfbf85c61b02ef10fd6c84d): complete subsection reference.

<a id="canonical-c6c44dac6a7f74df16039b67c7374023d70ae7d4712285632e5d8044be151482"></a>

<a id="canonical-5434085565a197fce9a7751a6b7fd55f24282d18705e553108df27d583c9a122"></a>

## max_age_value property — response_cookies_to_add / 51f05a6eaf20 / 7

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

<a id="canonical-e93f86f2ff5200ec9c2314017b4f7bf4a08f8afdd7fce8bc9687f05d815f0cce"></a>

<a id="canonical-00df44807d48a1a8e54227a0ee7b21645b92731fb681fe3e42332d2fd42012b8"></a>

## name property — response_cookies_to_add / 51f05a6eaf20 / 8

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

<a id="canonical-c72440169fd9c6fa42def3e007aa4ec1e9563f8be9429dcaf3b4c10935151bc7"></a>

<a id="canonical-9a47449a2e00379da0c3b6408369d71054d7b919f031f481da93bdd52845144b"></a>

## overwrite property — response_cookies_to_add / 51f05a6eaf20 / 9

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

- [samesite_lax](resources--virtual_host--reference--group-002.md#canonical-c395cc4c9ba3ad97b83a249c1dee0e9c498c12565c6ad66fe45815942fa73d4a): complete subsection reference.

- [samesite_none](resources--virtual_host--reference--group-002.md#canonical-1ec088bd3d81f0d9170294927fb481e8b9207a796862a2b2e5b220bdf519f7b8): complete subsection reference.

- [samesite_strict](resources--virtual_host--reference--group-002.md#canonical-ff9ae6de674e066c34391ee2919e274cc8360089a5c500cfb64d25f40ecdadc8): complete subsection reference.

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c): complete subsection reference.

<a id="canonical-7c63a0cad8e5467295784f4c5cdd9cf3c4f1c041af21864cd72fa2f5b6ea0cd1"></a>

<a id="canonical-83e5f202ec0c6c98ba8661b59871222cc7ec423cb7ff85443052dc0ac4ecf474"></a>

## value property — response_cookies_to_add / 51f05a6eaf20 / 10

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

<a id="canonical-9c55e586574534ea34e26c03310a1e19eff2eafc20b04772e9bfd0273a100139"></a>

## Next pages — response_cookies_to_add / 51f05a6eaf20 / 11

- [response_cookies_to_add.add_httponly](resources--virtual_host--reference--group-002.md#canonical-19f32417d81d7298854959ae5db61cf1b10863b8f0638427ee063e5cdbb22f7d)
- [response_cookies_to_add.add_partitioned](resources--virtual_host--reference--group-002.md#canonical-8b0b03f2fd1f80019347e2f0e75ab4702f0deb1345751ced3d6b5f620fa530f9)
- [response_cookies_to_add.add_secure](resources--virtual_host--reference--group-002.md#canonical-a8cad4e4c8f9caf2fb1aedebec93c4c5e433593a84059ebb2724b5e6a6b3b499)
- [response_cookies_to_add.ignore_domain](resources--virtual_host--reference--group-002.md#canonical-e14373fdc9c8cfb84d289951160b5f3652014057bb099e9cf1e920f1528bac26)
- [response_cookies_to_add.ignore_expiry](resources--virtual_host--reference--group-002.md#canonical-249d7e7b0a123c956421e57f3be8b3269e4931e03a350b15923e27f85e448ca0)
- [response_cookies_to_add.ignore_httponly](resources--virtual_host--reference--group-002.md#canonical-c9240233de110d27eebfbec16e51991e2a0da789ddfae7c379e723a069c0ce62)
- [response_cookies_to_add.ignore_max_age](resources--virtual_host--reference--group-002.md#canonical-cc4659c786d3c76d56524ac7f9d814c3cc700c497d756bfbc53be2101c08e60d)
- [response_cookies_to_add.ignore_partitioned](resources--virtual_host--reference--group-002.md#canonical-3b3a49d56f482e1fab35880702db7fdf679184201d665b6093cc4e714ed732e0)
- [response_cookies_to_add.ignore_path](resources--virtual_host--reference--group-002.md#canonical-1d16344d6e339d1ce0f61cedb96ff9f836bfca776ebbac7b308a8e1337b77df8)
- [response_cookies_to_add.ignore_samesite](resources--virtual_host--reference--group-002.md#canonical-f41b08ca9afb2ac17e384112b38f8edd3840ec55f84ef7ea75bd6fce9cf3b454)
- [response_cookies_to_add.ignore_secure](resources--virtual_host--reference--group-002.md#canonical-f5ebc7c702a9b26830957419174bd852361ff12a45e5d7bd064e4b6a54a989f2)
- [response_cookies_to_add.ignore_value](resources--virtual_host--reference--group-002.md#canonical-c87b1c50be423a75254c67b99929b2425ad6b87dacfbf85c61b02ef10fd6c84d)
- [response_cookies_to_add.samesite_lax](resources--virtual_host--reference--group-002.md#canonical-c395cc4c9ba3ad97b83a249c1dee0e9c498c12565c6ad66fe45815942fa73d4a)
- [response_cookies_to_add.samesite_none](resources--virtual_host--reference--group-002.md#canonical-1ec088bd3d81f0d9170294927fb481e8b9207a796862a2b2e5b220bdf519f7b8)
- [response_cookies_to_add.samesite_strict](resources--virtual_host--reference--group-002.md#canonical-ff9ae6de674e066c34391ee2919e274cc8360089a5c500cfb64d25f40ecdadc8)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-19f32417d81d7298854959ae5db61cf1b10863b8f0638427ee063e5cdbb22f7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ba41dc55ee8c5b6c030c7eeafbc7ae824c676c03876d7b20630e0ab5244a7c6"></a>

## response_cookies_to_add.add_httponly — response_cookies_to_add.add_httponly / d6c7e49e8044 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.add_httponly

<a id="canonical-d892b334ccbf5819c9430b09defc854f032adda946d0f20a5fc42705188a4b66"></a>

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

<a id="canonical-57701b785ecfda91bedceb09d855bf40e5568bacb343e246fe8c4684cd36c6d9"></a>

## Direct properties — response_cookies_to_add.add_httponly / d6c7e49e8044 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1276d7879d1570e44ee0d5bf86e86caee124d848fa9537e6542639e8c23eab83"></a>

## Next pages — response_cookies_to_add.add_httponly / d6c7e49e8044 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-8b0b03f2fd1f80019347e2f0e75ab4702f0deb1345751ced3d6b5f620fa530f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46617bf0e0afdde36891fd854e399bf0c49594b16cefb076c94a2e4d510da520"></a>

## response_cookies_to_add.add_partitioned — response_cookies_to_add.add_partitioned / 56efc5e013e5 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.add_partitioned

<a id="canonical-c2ab97ea65fff1f9956b761bfe25ef4fc279a974b56c7b0397e40ea232844fbd"></a>

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

<a id="canonical-3a680f403d63ddb9ca09a17a327231141b11f05429af81634dff9bb2e8520d78"></a>

## Direct properties — response_cookies_to_add.add_partitioned / 56efc5e013e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2f95dfeffb34572f98eacbec33d69df992412d2d87cae5f5d46d570c8643b55"></a>

## Next pages — response_cookies_to_add.add_partitioned / 56efc5e013e5 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-a8cad4e4c8f9caf2fb1aedebec93c4c5e433593a84059ebb2724b5e6a6b3b499"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f20f974e835d085f4d9ad99b14fd6502ae0a99ea1f3491f40c2ea72c7257fe61"></a>

## response_cookies_to_add.add_secure — response_cookies_to_add.add_secure / ea27eee923c0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.add_secure

<a id="canonical-2992c6444c5299cd66a09ae3c9c79a87b3e1ceafb53fee9637215dca62751c1a"></a>

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

<a id="canonical-8fb2a5c38af3d1ac43aba1b9c86814b2602d421ce8279c107786432c0ccbaedb"></a>

## Direct properties — response_cookies_to_add.add_secure / ea27eee923c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6071ebf60c9e8d01455499ab1e50e0d094954fe826aaf10827c1a7b18a6ba59b"></a>

## Next pages — response_cookies_to_add.add_secure / ea27eee923c0 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-e14373fdc9c8cfb84d289951160b5f3652014057bb099e9cf1e920f1528bac26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7c2c25f2b59d8784202d24bee5e831a08f25e41dbe538d0e5b1cfe0a08f2667"></a>

## response_cookies_to_add.ignore_domain — response_cookies_to_add.ignore_domain / fee7e9729cb4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_domain

<a id="canonical-ede24d9c283557b33c44104f81af0dcd72f0a10f8a93134da64e936d683238eb"></a>

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

<a id="canonical-7b97a10bd1caa51be4b501ccc0f253b2bdb83811c2bdac701c621524d04981d4"></a>

## Direct properties — response_cookies_to_add.ignore_domain / fee7e9729cb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6484267be31081cda05a27edd4900bcb5896e2464af10752972d389d05646287"></a>

## Next pages — response_cookies_to_add.ignore_domain / fee7e9729cb4 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-249d7e7b0a123c956421e57f3be8b3269e4931e03a350b15923e27f85e448ca0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b359c8b3b9120fc8fe6522a82768122915cc29a42c51f27b1c3caa25ac80ae9"></a>

## response_cookies_to_add.ignore_expiry — response_cookies_to_add.ignore_expiry / 4bb1a2bb6d5c / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_expiry

<a id="canonical-6ae290ffb6080bffea0a5a2ab718868015dad866f9a1485bed5286b96e6c37c4"></a>

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

<a id="canonical-471ef584e8b7e539295ca03b8fbc37cd1305fc89b48a51d1708ccb3700ff4f36"></a>

## Direct properties — response_cookies_to_add.ignore_expiry / 4bb1a2bb6d5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c28353e2154df1e3f2615190fb6850d79a9c4e76c5bd3da1b32349e57c2f930"></a>

## Next pages — response_cookies_to_add.ignore_expiry / 4bb1a2bb6d5c / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c9240233de110d27eebfbec16e51991e2a0da789ddfae7c379e723a069c0ce62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce54b3eb32cd931782b58b1505e0ecda2cf34b352ff6bcd9dee80d182cd08c5b"></a>

## response_cookies_to_add.ignore_httponly — response_cookies_to_add.ignore_httponly / 1528d52adfd9 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_httponly

<a id="canonical-8f21f8fc8671cb8a896161517e68d77c7e5ac7ffd5064c936ecc9338d4642250"></a>

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

<a id="canonical-e500637569317db05c26ace971006b22c0c340f24acd382065969821992950f4"></a>

## Direct properties — response_cookies_to_add.ignore_httponly / 1528d52adfd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7654031f646d3b91bafbb0c19820a4c26189c4f415ac8fcaacfb946764021bf"></a>

## Next pages — response_cookies_to_add.ignore_httponly / 1528d52adfd9 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-cc4659c786d3c76d56524ac7f9d814c3cc700c497d756bfbc53be2101c08e60d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4955b2c22a422911eb55d45a8e2176a8fbce4097cce130fece647393c06fed4c"></a>

## response_cookies_to_add.ignore_max_age — response_cookies_to_add.ignore_max_age / 63691d0498f5 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_max_age

<a id="canonical-be5d70c87a46d9981e2355179b01477acaddf93463bbc5932451f4cda67fbeb7"></a>

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

<a id="canonical-30924c6e63e0922d22bcff8bffa213c9606694a8b28a6e3c7dab69bc3e3bf7dc"></a>

## Direct properties — response_cookies_to_add.ignore_max_age / 63691d0498f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86162e4fc16aa53f957176fdb515d9ff01e284165d6ba4accec5389866eb9723"></a>

## Next pages — response_cookies_to_add.ignore_max_age / 63691d0498f5 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-3b3a49d56f482e1fab35880702db7fdf679184201d665b6093cc4e714ed732e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31aaaa7eba303784663d43803186e15cd47ed0103323cbcba78ad9cdcdf8cb96"></a>

## response_cookies_to_add.ignore_partitioned — response_cookies_to_add.ignore_partitioned / 9a8bb6532e3b / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_partitioned

<a id="canonical-05938624aea34b5800d44af33c5678b2b9471a8138b51ffe5840071858d314a3"></a>

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

<a id="canonical-9c1e2718a29ad62ec310fd15c2df88cbddb43d40fd30ba449e051e7873aa61d0"></a>

## Direct properties — response_cookies_to_add.ignore_partitioned / 9a8bb6532e3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2072953ac3118bbeba11df27eade4d3db669b64fab0c4fe2b2bbce6b6ae77ae"></a>

## Next pages — response_cookies_to_add.ignore_partitioned / 9a8bb6532e3b / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-1d16344d6e339d1ce0f61cedb96ff9f836bfca776ebbac7b308a8e1337b77df8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcefdba095398f32389c0775bdecd40989788cbb2b1e685ee6208a07c8732807"></a>

## response_cookies_to_add.ignore_path — response_cookies_to_add.ignore_path / b8bc6dca48d1 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_path

<a id="canonical-e42aa3c9a27f8f7c43c55add844c6697a0aaa0aedf91d4f62ce650a180b6c9dd"></a>

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

<a id="canonical-0168da0ca7f976f65897ae7e378c175a5890e2b519df2a7d82d471f2cded6f0c"></a>

## Direct properties — response_cookies_to_add.ignore_path / b8bc6dca48d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17b5f7451676cb179f3b0102dc6ecb2d6c280a486c6273fb15c4bf4664dd6262"></a>

## Next pages — response_cookies_to_add.ignore_path / b8bc6dca48d1 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-f41b08ca9afb2ac17e384112b38f8edd3840ec55f84ef7ea75bd6fce9cf3b454"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f74ab192b7fd1d6ca5ac1ff4cb56f1e64a0b5c794c26c835d295078c6d9d6cc"></a>

## response_cookies_to_add.ignore_samesite — response_cookies_to_add.ignore_samesite / a01aa5308550 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_samesite

<a id="canonical-8523ed0a210f00b8656e526ef78c0bf28f91d20c9db29abe9349d4f7b91acb33"></a>

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

<a id="canonical-7ae7dc51c1f06becc1336e7094259db6e05ad950a83e271237f6848f8aceae10"></a>

## Direct properties — response_cookies_to_add.ignore_samesite / a01aa5308550 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-396004bc13954dbcb1b217d37112402accb6a5786506d11f8e7acca43d7808c8"></a>

## Next pages — response_cookies_to_add.ignore_samesite / a01aa5308550 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-f5ebc7c702a9b26830957419174bd852361ff12a45e5d7bd064e4b6a54a989f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac99fc28a6c97dbda5c3e238c98c0b1bd66da67c4f6aed3565372758fe720341"></a>

## response_cookies_to_add.ignore_secure — response_cookies_to_add.ignore_secure / cc8bb508338b / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_secure

<a id="canonical-91e800ea46ca2bef411cd39e8099987c85a357cce20e840b5a7a70288b71f4b7"></a>

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

<a id="canonical-283287df476e2e3b1db311eb63783e29a78da552ee1799a4671dd5e1f9087318"></a>

## Direct properties — response_cookies_to_add.ignore_secure / cc8bb508338b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed537892e668e596266b3d92ce6cc4cf7704814e293491b9e28df491705485bb"></a>

## Next pages — response_cookies_to_add.ignore_secure / cc8bb508338b / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c87b1c50be423a75254c67b99929b2425ad6b87dacfbf85c61b02ef10fd6c84d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d881dde579f1a54b5e3e97182771e8cb30f941ee6b7a020b1db91b46b4681c4"></a>

## response_cookies_to_add.ignore_value — response_cookies_to_add.ignore_value / a178ea8844a3 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.ignore_value

<a id="canonical-2708f0fe15a923e87570fa7ddeb3a90a1c452f119808ee22f915770dce57c5bb"></a>

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

<a id="canonical-865ec740c1570a81e998716eb8bd171afef6c3c6d575c5868313ca4599cd45aa"></a>

## Direct properties — response_cookies_to_add.ignore_value / a178ea8844a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1993f9e73c78994d39a16b05c9cbf4520e4a4725f4d3e3ebca71f5fe26cfac1"></a>

## Next pages — response_cookies_to_add.ignore_value / a178ea8844a3 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c395cc4c9ba3ad97b83a249c1dee0e9c498c12565c6ad66fe45815942fa73d4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c11b3cfcd1e902bf289e0c4aee6e9d7f9b77871ea7507b9c311b899611c24fc0"></a>

## response_cookies_to_add.samesite_lax — response_cookies_to_add.samesite_lax / 8a6dbd903fd0 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.samesite_lax

<a id="canonical-1ebe99e52e822f314a41eef13da4f42f75b5d8da00c0a080b88456d1cbbe179c"></a>

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

<a id="canonical-9d4d2da430d6c7d2af1375ba22ed87736e420b46d019a063226305abf4e01011"></a>

## Direct properties — response_cookies_to_add.samesite_lax / 8a6dbd903fd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dfcebddd428a0a8d2fb8d6aeb3f519bd155103797af8f8f451228cd11f670cd"></a>

## Next pages — response_cookies_to_add.samesite_lax / 8a6dbd903fd0 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-1ec088bd3d81f0d9170294927fb481e8b9207a796862a2b2e5b220bdf519f7b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e080cbd6e909afcec08adf272792e9c27c9a4737b0437b468394ffe97f054c6b"></a>

## response_cookies_to_add.samesite_none — response_cookies_to_add.samesite_none / 8413dec5cfb4 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.samesite_none

<a id="canonical-dac3de5390f74f3f61793c4d75ababef3acc4445e0d094630fa832086ac58e14"></a>

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

<a id="canonical-13d920615dfe3c936e3eb2b69220f6093cb8694bc708cf01585654cd01e89c2d"></a>

## Direct properties — response_cookies_to_add.samesite_none / 8413dec5cfb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fed43b9f07ad6e6f97c542a34d9e9cb7620d7d259edb8394c34ef56002ecdc4c"></a>

## Next pages — response_cookies_to_add.samesite_none / 8413dec5cfb4 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-ff9ae6de674e066c34391ee2919e274cc8360089a5c500cfb64d25f40ecdadc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74db2d1158718dde584737020fb1ea9350807e98395da3905ba002fa6c7173b1"></a>

## response_cookies_to_add.samesite_strict — response_cookies_to_add.samesite_strict / d275dbdb1376 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.samesite_strict

<a id="canonical-30e992d9e32e0234b265716eeeb788a3ed0fc8545ebfdde2074318b010294de5"></a>

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

<a id="canonical-0023b90088f55c659b8107c6b2ec58a924c4cd54decd464fb733645647c4011a"></a>

## Direct properties — response_cookies_to_add.samesite_strict / d275dbdb1376 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4870f55746651d1a1e5303b5d4ed3f1ee5ee6ae04cb064bd4daa89364daca03"></a>

## Next pages — response_cookies_to_add.samesite_strict / d275dbdb1376 / 4

- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f401f720380874c808ad061d42ce70d3ecfa29ef43690e0bd8d93b76f48f1288"></a>

## response_cookies_to_add.secret_value — response_cookies_to_add.secret_value / f232566af35d / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- response_cookies_to_add.secret_value

<a id="canonical-24fcc1c2910a2f42857eae9f247b3f0a26e7fb674ce744db40c9b00852b5b391"></a>

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

<a id="canonical-e5224ffeb9d24dea7e7ace0df999db947043be50f68a9406127aff93cd80a4d8"></a>

## Direct properties — response_cookies_to_add.secret_value / f232566af35d / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-c104faef23e5040adda8866a06e38fd95dcabd4477d16f62c83143f8b867ce1b): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-6a33fad4e09912a6befcb8600e638ae3e178702f056545ca749d9284b2d8123e): complete subsection reference.

<a id="canonical-bf04a7314778544d5180f6deeb53478999cbfe851bf1634ff675ab84d00ec35e"></a>

## Next pages — response_cookies_to_add.secret_value / f232566af35d / 4

- [response_cookies_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-c104faef23e5040adda8866a06e38fd95dcabd4477d16f62c83143f8b867ce1b)
- [response_cookies_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-6a33fad4e09912a6befcb8600e638ae3e178702f056545ca749d9284b2d8123e)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-c104faef23e5040adda8866a06e38fd95dcabd4477d16f62c83143f8b867ce1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c697f6355ea8ee0195485400fa49741368487aaf49662747fd92376b81ef767"></a>

## response_cookies_to_add.secret_value.blindfold_secret_info — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c)
- response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-e1ca003435aadf076ddf311fb106f18473e48b57ca614b1b2537b990d2a21e1a"></a>

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

<a id="canonical-b2e2c4b50c228581dd9347fd5b0985b6f08e3775885755da684a5ccf4837ae07"></a>

## Direct properties — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 3

<a id="canonical-f47a80bdccda091e5e432c8e371935b6385cfc320830a6f57dfc927485934324"></a>

<a id="canonical-66c23ede62a8aa936a5d9bf410770eed892ae9b14c36b6dbb21f719e4a1427c4"></a>

## decryption_provider property — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 4

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

<a id="canonical-421d9511254068b896c56b682aff46ace0137157f49e79c72815a74468c24e8d"></a>

<a id="canonical-41d58dd8b00a940791c7a63c248b1d1d0e78e35499667038d98dd854a21829d9"></a>

## location property — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 5

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

<a id="canonical-0ef9571ff933c09a8e2c67cc29682aa4e1449c374d22eeff38a72fb5560f8670"></a>

<a id="canonical-4f61d17ccb87efa09094a22ce71a169fa3fe31706f38b6c47d3ee0ff81d5c66c"></a>

## store_provider property — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 6

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

<a id="canonical-db6a78b91e658aa2b8c94f0deb9b8b26ad736b9f34e76aa18cef6fdc8af13e54"></a>

## Next pages — response_cookies_to_add.secret_value.blindfold_secret_info / 20766d08208a / 7

- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-6a33fad4e09912a6befcb8600e638ae3e178702f056545ca749d9284b2d8123e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fb1daee204c5320dd5fee250cba5c49ba31eae9a1a61a76e66c0b27c3c3b3f3"></a>

## response_cookies_to_add.secret_value.clear_secret_info — response_cookies_to_add.secret_value.clear_secret_info / 77f07d7c37cd / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-4e7cf59e1041b6abaeebc939d7938c20eca265136fa1a3d4848470b84b158f37)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-a1620388edd67c967155e0a822fe77e809662969ccc418974aaa08b6c307a369"></a>

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

<a id="canonical-463ce70f8020c0228c926ac6c96739dd092d27eff40cfe455ae1e58a8e722dfb"></a>

## Direct properties — response_cookies_to_add.secret_value.clear_secret_info / 77f07d7c37cd / 3

<a id="canonical-382d490c1e6002c483be03ae4f3ff48eb2bb5154a1b6ed3d283d19754b369f33"></a>

<a id="canonical-9203decd397f80d5395d32889e040a28d1f615d099000a8c67c4d38b5dd0985c"></a>

## provider_ref property — response_cookies_to_add.secret_value.clear_secret_info / 77f07d7c37cd / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1a4992058e61cfbfa06def0f2831cb931cdf6115e8b54cabfddea2e85c77ff03"></a>

<a id="canonical-bb0ab845b61e92d76207c3b79b54cb72edd3e030bc0ce31d12a952c096290f45"></a>

## url property — response_cookies_to_add.secret_value.clear_secret_info / 77f07d7c37cd / 5

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

<a id="canonical-560517ba021a14a19191365cf7cd9cf5222c7251cbfb51b5296d0bab4a9e2b5a"></a>

## Next pages — response_cookies_to_add.secret_value.clear_secret_info / 77f07d7c37cd / 6

- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-44cdd6b7faac8da88d089ef6b9a11482aac9727a9803fdca38fa9cc414df229c)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1a21ebd13ca7ef765a5aa2d6beca71ca0884f5fa9ccec1461b1138fc998c62b"></a>

## response_headers_to_add — response_headers_to_add / 00642a87c820 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- response_headers_to_add

<a id="canonical-7e8501d5a5681371c96a7242d35ba1796367e5eb6f1c4fafb7070fc1770c3898"></a>

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

<a id="canonical-37ad4ce120ede349471de5cbce8599240e86a0cea4da736dcfebce16342c4800"></a>

## Direct properties — response_headers_to_add / 00642a87c820 / 3

<a id="canonical-42ac71d39262b08c5e0756803456d9850c7898d360cb4368f89dbfec8860e701"></a>

<a id="canonical-29d7e7e9f237213619ba59f7b1f36fa7e11b377c65a374e0be4cd75de8d113c7"></a>

## append property — response_headers_to_add / 00642a87c820 / 4

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

<a id="canonical-5423c1783154b48c0370012b810cb4df0038ff735e97f61ab4a6d941c70af895"></a>

<a id="canonical-c115c6cbd06b76925ebdfeee49188a17bb9154628525a385dff7cf2fb7ec0ead"></a>

## name property — response_headers_to_add / 00642a87c820 / 5

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

- [secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6): complete subsection reference.

<a id="canonical-488bf7fe06b700c17393ac108c2859349116b43e3951a51610ca0ab79015c919"></a>

<a id="canonical-5252e5167035b717039ba37342363f2f2c5aa690a1d26fcee88322e155d15d21"></a>

## value property — response_headers_to_add / 00642a87c820 / 6

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

<a id="canonical-39c8ba3e78fe64b208a53f09f5c23bddc0fd682f0b036fcbcd6fcdd25559acf0"></a>

## Next pages — response_headers_to_add / 00642a87c820 / 7

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2ae59dc793b2f0360a34070430b753f6fb33f4850956b0ada0f55e94aa54dad"></a>

## response_headers_to_add.secret_value — response_headers_to_add.secret_value / cf89de3a4e55 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a)
- response_headers_to_add.secret_value

<a id="canonical-4aa6b5724a2aee4a5808c96c51c75035d59d292e9ebea79c6de66e36df9d3571"></a>

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

<a id="canonical-eef098813818c520d429e92bd23d47759cf082150b83a76a6fcc9309004d1f4d"></a>

## Direct properties — response_headers_to_add.secret_value / cf89de3a4e55 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-6f1e1abdcc7f00a5e1b14e294096a5aa2efe844f7a225f3e243673db4afba62f): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-64cc288b7a349e49fdbb27b3e173ef4b3357331027a491ab6ab2fe6c9bbe0d9d): complete subsection reference.

<a id="canonical-5aad2aebb9e6f30e89b91b35851ab8c881c36975ee3189222899a71014bf1528"></a>

## Next pages — response_headers_to_add.secret_value / cf89de3a4e55 / 4

- [response_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-002.md#canonical-6f1e1abdcc7f00a5e1b14e294096a5aa2efe844f7a225f3e243673db4afba62f)
- [response_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-002.md#canonical-64cc288b7a349e49fdbb27b3e173ef4b3357331027a491ab6ab2fe6c9bbe0d9d)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-6f1e1abdcc7f00a5e1b14e294096a5aa2efe844f7a225f3e243673db4afba62f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0576b0145aff5360906fafbb16c76d361c68b68f5b092a524a1b3e7751f9b2a6"></a>

## response_headers_to_add.secret_value.blindfold_secret_info — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1c01161ffc5340df4c82d9523e86386cd21f2bfa66c171a0a78f971eae374448"></a>

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

<a id="canonical-0996a688ef089e2d6a817f719148ee92b063b57bfb635855d17395b9f2b11681"></a>

## Direct properties — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 3

<a id="canonical-5554f91a516a34322905c3b0d228afcabc67f7ffae92cfa3976e81f7f7a9bbb4"></a>

<a id="canonical-67340a61956a5a83816882f220edc7786bcad41e3d5a39731c1e2d00a3797a0e"></a>

## decryption_provider property — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 4

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

<a id="canonical-557003721f25ba88b8f9ec2cf7861322a987fd9ebb1b3ddfdfed9a6e6e3bed3c"></a>

<a id="canonical-028ae644c618175f508c842b271248ec645ab9b08f4b879358b1c255e5330652"></a>

## location property — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 5

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

<a id="canonical-0a040f717f38daa8f8d79dce0cdebf8dc3c31f305344d7585e56fa0216861db9"></a>

<a id="canonical-69655d3e1c32842e9542bf15d05875119f1f07d7588909495d9f3ccf5fbbc4cd"></a>

## store_provider property — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 6

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

<a id="canonical-dd256e17a146d79d7cd770bf316763a2efb9ab4054f849a9f2fb3a33c50c11b6"></a>

## Next pages — response_headers_to_add.secret_value.blindfold_secret_info / bb9b60e2fd70 / 7

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-64cc288b7a349e49fdbb27b3e173ef4b3357331027a491ab6ab2fe6c9bbe0d9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90480e5961b765a91597fce4507747aff27022918459f327a60d8697c5684f91"></a>

## response_headers_to_add.secret_value.clear_secret_info — response_headers_to_add.secret_value.clear_secret_info / 221066718a6a / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-27ae1f7eddbb41243c64362e9882a8363bf5158e484ca5ea4bae281e3b7113a8)
- [response_headers_to_add](resources--virtual_host--reference--group-002.md#canonical-8d2d03e2bc424cfad7ef3ddf8c556266729057ed299f5e7f243e78f51f32530a)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-9598e203cfe8eef3b9e6de40749607514cfe0108d15230ec47178797f692a7f6)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-e1c0f35ed29f6b14e06697024b818f70ec1d5c5e8cf131a7310df6c01ed724f4"></a>

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

<a id="canonical-fc6af512efbffade71b5881042acef9a766724d3595a94a094b8cf9c3ec17f22"></a>

## Direct properties — response_headers_to_add.secret_value.clear_secret_info / 221066718a6a / 3

<a id="canonical-313221894d5879fe081073011e85b803271d1118c56892dd052cea3ec54c14a5"></a>

<a id="canonical-6f38b8a6a23294187bc5627db027f5c4ef6688af3b48e121fca53523a27b6e56"></a>

## provider_ref property — response_headers_to_add.secret_value.clear_secret_info / 221066718a6a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7849f670e34e8fe0c4f76a395ffdecd455f02fb9a82e7e0dc3d111f11cee1a2a"></a>
