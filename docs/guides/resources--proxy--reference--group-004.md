---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-76c7fb48766116708100414b6b69ff9da5cbd348a3b233f0c27f59bee7e58db3"></a>

## tenant property — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 6

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

<a id="canonical-1bd3cdadc66004a949d0a62874471f4331d0d66b1f7c5ed823c68769ec2f36b7"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 7

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f1c470c6447008992c3730d09548f77620e03c2e703777ad78af0a410626b77a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5de9885b8c17128f77033031d1f75235d5f4ba4bf2a6e383f51f53fc3c0dcf37"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled / 5e27a3386c34 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled

<a id="canonical-d3b70e43bee2f79116ce851707a8bdef4a1332d0aad2d11e5765e6657fd55b06"></a>

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
xfcc_disabled = {}
```

<a id="canonical-2919a7bf1f0f3e3ff38add04ab020c87c739dca37be1ea540254dc820ee52e99"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled / 5e27a3386c34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2d58330a8af06c47ee78be67f106835f23e615acbe0fd92c50f2b2545ba6852"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled / 5e27a3386c34 / 4

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-063cbc226b3095e04be5ff936f17986f6d6a54844d01ff228964e1712f8e24ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f7398a74fcd5f0209c2788d9ac59b46bd8fd66c28f3e13ea22e2aa72018346"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options / a14cec6c30a1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options

<a id="canonical-e50879c809f6638e947e2d9adcf192224823a7ff4a700ba131c5d83ca30bdc7d"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-041e844d6cd9ac19b167e68b10972eded0e2c866e877646f46fefdc214405089"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options / a14cec6c30a1 / 3

<a id="canonical-337b0ccefb6a87c397016ccf9b8a8df42d7f6ff31bac9f1cb1d24b0f43f731d0"></a>

<a id="canonical-09617a7f0317c77c39427601de548ad02741250296fbffb067c6d233fb209a5c"></a>

## xfcc_header_elements property — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options / a14cec6c30a1 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-dc2d408a74b6e7195971a23892ee65d68ae441a3eaaec5ec1763c38c59b9d55a"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options / a14cec6c30a1 / 5

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-d6201277c11d95edb6c35bf0110b8dedb94e654460ba57d9257c01c6cadaca1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5e053031861c4e824fa71bd65924da6f0c9d818fbb26d620199e0c5e7959c1b"></a>

## dynamic_proxy.sni_proxy — dynamic_proxy.sni_proxy / 4ccc12847c25 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- dynamic_proxy.sni_proxy

<a id="canonical-60310432f74632f6f49cd5f4c5ae0a2e93d92b3f77a6213e158264d2162e76cf"></a>

Type: `"object"`. single nested block, Optional.

Dynamic SNI Proxy Type. Parameters for dynamic SNI proxy.

Upstream description:

Parameters for dynamic SNI proxy.

Receipt-pinned upstream constraints:

```json
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
sni_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0646fac0a6aa2235e8d8d658ea9915fc67aa3eefcba2cf8c4e915217c80d8767"></a>

## Direct properties — dynamic_proxy.sni_proxy / 4ccc12847c25 / 3

<a id="canonical-d29cb9b33cdc3b7ff857f92c53c87f9356ff478d992ed75c2ebab679f1486b92"></a>

<a id="canonical-9e2b4a9f8c633c67c09a004f8ab60dd725425ca3f77c9196d4059c1210a04de6"></a>

## idle_timeout property — dynamic_proxy.sni_proxy / 4ccc12847c25 / 4

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400000,
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
    "ves.io.schema.rules.uint32.lte": "86400000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  }
}
```

<a id="canonical-9912c10ae139f21ad0857e069e9cd44e9874e814c9a7a08dab8d2c465af74302"></a>

## Next pages — dynamic_proxy.sni_proxy / 4ccc12847c25 / 5

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-017395461bf2437f60b2399b20ffb07f0ce14b2a9854eaeee9ba0eaa98df894e"></a>

## http_proxy — http_proxy / 23e30187ced8 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- http_proxy

<a id="canonical-0b32a135f0eac087814b5eeb5a2baa7a8860790f0ef07edeea58f4e6eb6bfb9e"></a>

Type: `"object"`. single nested block, Optional.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Upstream description:

Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

Terraform syntax:

```terraform
http_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-f82930d27dd4080f51e128abfc625510ca10e7caa75750b80708cccdd7a49cd5"></a>

## Direct properties — http_proxy / 23e30187ced8 / 3

- [enable_http](resources--proxy--reference--group-004.md#canonical-826640e1098d48e62d6ea2fd7ae99da796c90c6075ef36053e750b3279c86984): complete subsection reference.

- [more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43): complete subsection reference.

<a id="canonical-735696aaaf4e5f583e9b39df681bc30fa81534e4643f16f1105334505e47ff81"></a>

## Next pages — http_proxy / 23e30187ced8 / 4

- [http_proxy.enable_http](resources--proxy--reference--group-004.md#canonical-826640e1098d48e62d6ea2fd7ae99da796c90c6075ef36053e750b3279c86984)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-826640e1098d48e62d6ea2fd7ae99da796c90c6075ef36053e750b3279c86984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88751d2536f576a90fe4c53ca8d70e7b7c2961a4a23015d34b91632db6fef18d"></a>

## http_proxy.enable_http — http_proxy.enable_http / 123d52a7678e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- http_proxy.enable_http

<a id="canonical-a0c518db07a4bc8a6e1d8649c58607ad6c1eb53adb2847382ab8c4d4a9882b2a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable http.

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
enable_http {}
```

<a id="canonical-25cab81a173647b5add03e27f0bcf8fe34bb524470b853de780684986ba30350"></a>

## Direct properties — http_proxy.enable_http / 123d52a7678e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ebc72c1bd3c7c96ac38f39dc17cbe7595231ab07bb8d552fe82b8d91199139ff"></a>

## Next pages — http_proxy.enable_http / 123d52a7678e / 4

- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08b8bc442fec275d9811d238d80de481fea94dbccac30506676cd240512640ac"></a>

## http_proxy.more_option — http_proxy.more_option / 0300ced90bbd / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- http_proxy.more_option

<a id="canonical-a906b91650ee34c26970211cfc973a24f9ce7defdf483e764392b6e79f5daddd"></a>

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

<a id="canonical-356f703bce437f011ebfb41aff680128e5040d8ce5a03f10c64c2f91211dca31"></a>

## Direct properties — http_proxy.more_option / 0300ced90bbd / 3

- [buffer_policy](resources--proxy--reference--group-004.md#canonical-3055703a316b36cfbd35dda2ae1f5cf39c51274ffeb49731a71712c8ff3f5ef3): complete subsection reference.

- [compression_params](resources--proxy--reference--group-004.md#canonical-f82d5a616089bb7a2e42b8cfb0947749953c75154da9ae4efd055ab96bafe117): complete subsection reference.

<a id="canonical-3728d6163c87af12a231ff6334c66ac90542e13399074237d086616f0335c037"></a>

<a id="canonical-b66ba8a68404c1945a86fdb70d932ceb87e1768bbe98ed3d5a688e16d47178f5"></a>

## custom_errors property — http_proxy.more_option / 0300ced90bbd / 4

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

<a id="canonical-ede00cd446bc94f2926bd2f32d27b560d0500fb9ebe92759d423cccce4a0b27c"></a>

<a id="canonical-422b37f2a14a0fe90fbf7101dfd9d4c451c93333d20776c9e6b196e3320f1003"></a>

## disable_default_error_pages property — http_proxy.more_option / 0300ced90bbd / 5

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

- [disable_path_normalize](resources--proxy--reference--group-004.md#canonical-2f1f1fd90308cfaf2907ec9f177bb2e93712076492bc24f6ddc0fad75216ad56): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-004.md#canonical-b1e4bad1866dfe74d87dadbcfa5219689850a8792cce4cc6d5e162b8b099533b): complete subsection reference.

<a id="canonical-165c2731825ddf391ef2c730dfd54987e049fba6d39ed8e099734fd1fe7c55c6"></a>

<a id="canonical-e2cb903e3c32e3f08f347f2442ee9520bef7ca53dcc568dcd3ed13d8b1e77999"></a>

## idle_timeout property — http_proxy.more_option / 0300ced90bbd / 6

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

<a id="canonical-746169c1a1f2cae9081b4f96861ec304f68eb411cc68cefcdcfd40c979ea1fe1"></a>

<a id="canonical-092a04c02f5621bea7385b56baac3bfaf9b427f0318278b85a15f28a6f6d30c7"></a>

## max_request_header_size property — http_proxy.more_option / 0300ced90bbd / 7

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

<a id="canonical-42725721913a1a8958e594f270a1311cd06a7f059f56064a59d3d55aa2535032"></a>

<a id="canonical-f8c09e0f2bb04c4ff84a8be578dd45b74bbe8a5ef2e197e728c8a5b10ed8acfe"></a>

## max_requests_per_connection property — http_proxy.more_option / 0300ced90bbd / 8

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

- [no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-436bac797f51b156e9aa78e64ba27e4f8b08334dc461faaf4133f163f1cc29ad): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6): complete subsection reference.

<a id="canonical-18bac30cb907dff889c8d5e6def87368feec0df166b34f5bc7e39d15f601627f"></a>

<a id="canonical-8890e47dcca9d2254e5a1736db8caa6796c80e871f5b877e3f88cff3da163f13"></a>

## request_cookies_to_remove property — http_proxy.more_option / 0300ced90bbd / 9

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

- [request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50): complete subsection reference.

<a id="canonical-1b4517345ca4208fa81826bceb895c183ffe03fbc339aa122946fcf724fcb1a7"></a>

<a id="canonical-5d8f17370285c4c4dbf67cc1e7143aeb3051b74e1da6470e5f6ed0316a4d42c3"></a>

## request_headers_to_remove property — http_proxy.more_option / 0300ced90bbd / 10

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

- [response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375): complete subsection reference.

<a id="canonical-9070dc7d65283745bd6925624f39c68f02e94ae86937a7e7359d4288cce94be5"></a>

<a id="canonical-973f95fd7d826b3daba2af4fedf2ec66d9710f50579f42cd455df9e3cfe109d2"></a>

## response_cookies_to_remove property — http_proxy.more_option / 0300ced90bbd / 11

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

- [response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e): complete subsection reference.

<a id="canonical-8259051d13c027d004f95b106fc6e679ae4e5b2fc2dd31ad457742da3b35833b"></a>

<a id="canonical-3e27f952457ef0ad4f5646ea29cf75fc9317cda7eeff9d0b268108d52b8b9400"></a>

## response_headers_to_remove property — http_proxy.more_option / 0300ced90bbd / 12

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

<a id="canonical-34940389c328e62faa9956d7a529471b71c2126f145cf33f94a1719b6e1b9f6d"></a>

## Next pages — http_proxy.more_option / 0300ced90bbd / 13

- [http_proxy.more_option.buffer_policy](resources--proxy--reference--group-004.md#canonical-3055703a316b36cfbd35dda2ae1f5cf39c51274ffeb49731a71712c8ff3f5ef3)
- [http_proxy.more_option.compression_params](resources--proxy--reference--group-004.md#canonical-f82d5a616089bb7a2e42b8cfb0947749953c75154da9ae4efd055ab96bafe117)
- [http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-004.md#canonical-2f1f1fd90308cfaf2907ec9f177bb2e93712076492bc24f6ddc0fad75216ad56)
- [http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-004.md#canonical-b1e4bad1866dfe74d87dadbcfa5219689850a8792cce4cc6d5e162b8b099533b)
- [http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-436bac797f51b156e9aa78e64ba27e4f8b08334dc461faaf4133f163f1cc29ad)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3055703a316b36cfbd35dda2ae1f5cf39c51274ffeb49731a71712c8ff3f5ef3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6ce5d3de850758450a3b8c87c8c33d0f9dea1cead09c4e5510b929874247e06"></a>

## http_proxy.more_option.buffer_policy — http_proxy.more_option.buffer_policy / dbd0c5d85ef2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.buffer_policy

<a id="canonical-8cf6b5e271a99c0e7f2e0cacab9346cc939a5348c869dd44e61585532a70b8e8"></a>

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

<a id="canonical-9045e7728ccdf4131581d1dd91583ae7a48f7b7d0d94d01356ff1993132a6ebe"></a>

## Direct properties — http_proxy.more_option.buffer_policy / dbd0c5d85ef2 / 3

<a id="canonical-cc9e4c73ff321a392c44588b8329db1491c83cd4c1613a86fff4cfe7e0fe3581"></a>

<a id="canonical-5bcf6a20db66b71d4ba06eecc1b009fb28a37be4216cff1363318584a8b773f7"></a>

## disabled property — http_proxy.more_option.buffer_policy / dbd0c5d85ef2 / 4

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

<a id="canonical-cae7da4dd1f6cff109156d7129d872c79e37fc47bdd30ea80d9e8c905fd28284"></a>

<a id="canonical-042d45c1c71b6ad68d93c0bbf21b99e5ad0fa8f4bd9381d25dd7b64aa9149276"></a>

## max_request_bytes property — http_proxy.more_option.buffer_policy / dbd0c5d85ef2 / 5

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

<a id="canonical-b077319e2033c1277481ce6df7a545c2f0425dab2a5707828bb0f6130c49dffa"></a>

## Next pages — http_proxy.more_option.buffer_policy / dbd0c5d85ef2 / 6

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f82d5a616089bb7a2e42b8cfb0947749953c75154da9ae4efd055ab96bafe117"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7dd72420309b27e856dc39e2ddcd2b29c913528c4641c4db9d4117fe07bd555"></a>

## http_proxy.more_option.compression_params — http_proxy.more_option.compression_params / 620b86725d40 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.compression_params

<a id="canonical-383e31266524d5de80038a80d4c0a2af4836046f08e784b34e2e0d796d92f0d1"></a>

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

<a id="canonical-0378de311a47aa0aa5dc709fd063de3114d23c6452ca392008492561b813a4cd"></a>

## Direct properties — http_proxy.more_option.compression_params / 620b86725d40 / 3

<a id="canonical-adf50e60cd45e3caf5ff2d843263680e33c1f77e62b7a2e5383e4b166c454833"></a>

<a id="canonical-820172738f2d1000d6c05a512da65051a7d59af968583d22fc5be342bde16c0b"></a>

## content_length property — http_proxy.more_option.compression_params / 620b86725d40 / 4

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

<a id="canonical-1e338a7438e4d257263a6704514c050c016c9a79c3fbf3d3c1fc01215e5e8d3a"></a>

<a id="canonical-e83d65926d4b8cc2864fa7c8df7e709851c41484360ef06321ebb5c1d0069fee"></a>

## content_type property — http_proxy.more_option.compression_params / 620b86725d40 / 5

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

<a id="canonical-4595fe55730030125bab9d3260a5d9f4d82cd6e018808036cdba6635e6aa734e"></a>

<a id="canonical-302baf96260906fd28bdc929cc027d934e77a56a0cae71913774357bc44802d7"></a>

## disable_on_etag_header property — http_proxy.more_option.compression_params / 620b86725d40 / 6

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

<a id="canonical-9e2c93b3da8c0d5b37d695de58e6471660458df76eeea02a3936912dacfcab5d"></a>

<a id="canonical-1b52d528660eaa38d70947a78b852e4d14b8174f6abafd2b00e6e31cb0137b25"></a>

## remove_accept_encoding_header property — http_proxy.more_option.compression_params / 620b86725d40 / 7

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

<a id="canonical-93001a9d23eae7b07ebd9929b54a0c7acaed706a6125af7002183b835d731e11"></a>

## Next pages — http_proxy.more_option.compression_params / 620b86725d40 / 8

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-2f1f1fd90308cfaf2907ec9f177bb2e93712076492bc24f6ddc0fad75216ad56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b79811ced453c4a7b738a803951f64f4e052a69110b80819118670753a668bd1"></a>

## http_proxy.more_option.disable_path_normalize — http_proxy.more_option.disable_path_normalize / e2d1fee012c7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.disable_path_normalize

<a id="canonical-16a6f43d1719a4a5c5833cc54d3a8cdd50ee26e5af32b8677d6af23280ffa70f"></a>

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

<a id="canonical-f5c2787899ff8db3213909f7e041c8c6e7c9933803796b56738e1facad42c566"></a>

## Direct properties — http_proxy.more_option.disable_path_normalize / e2d1fee012c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40e1b20facad625c877878ac87c646b82153fb3ae151167ad6ccd0e93f42f6b9"></a>

## Next pages — http_proxy.more_option.disable_path_normalize / e2d1fee012c7 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b1e4bad1866dfe74d87dadbcfa5219689850a8792cce4cc6d5e162b8b099533b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cfaca6797c4f1e12ae16c7b60f95d4665356f4ef962390fd825d61ec8fa8c6b"></a>

## http_proxy.more_option.enable_path_normalize — http_proxy.more_option.enable_path_normalize / ba3eefd47795 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.enable_path_normalize

<a id="canonical-23463db2b6894d74a8c81043382e9770c15ba12f66815b70ea8ed86bb1ca7a15"></a>

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

<a id="canonical-0fa3c851350b39429b0413cd9ea10c9831386a71d204113b7d93895014c0c357"></a>

## Direct properties — http_proxy.more_option.enable_path_normalize / ba3eefd47795 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae56c3d057a4f677d97ce2605b065c9c33f5d0d9ca494dd04ea92b0e9334650d"></a>

## Next pages — http_proxy.more_option.enable_path_normalize / ba3eefd47795 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-436bac797f51b156e9aa78e64ba27e4f8b08334dc461faaf4133f163f1cc29ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6102826846941ee3e5a8d7b692b3694d83a66492c0f124525550e539bf8da158"></a>

## http_proxy.more_option.no_request_limit_per_connection — http_proxy.more_option.no_request_limit_per_connection / 6e810208efc4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-870c0cd6a16e0ac2df872480cf12d204243ea9fb3121cb8811ebd78c717abf03"></a>

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

<a id="canonical-a514800d28752570bbbfd3c4785bb7423d13f05401e032c29890827ba64746dc"></a>

## Direct properties — http_proxy.more_option.no_request_limit_per_connection / 6e810208efc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0373c33c6d0804d34a54fd150189d69cc8b6f1f8b6a5e6f4ee962d8957a9545"></a>

## Next pages — http_proxy.more_option.no_request_limit_per_connection / 6e810208efc4 / 4

- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05597abc7ef5b6613c672b1afc84b6a8b2e8caf0f0fe7a16fc03772f122949ef"></a>

## http_proxy.more_option.request_cookies_to_add — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.request_cookies_to_add

<a id="canonical-6403d5afd60fc9062c8773971be2b52db10f6ab3ce7f450ab29bffedcf0d1653"></a>

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

<a id="canonical-b7ca86a7dce504a2f17cfb6d8e80d916e3d5576d8b0510ef7c33c36223499afc"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 3

<a id="canonical-d289451206d36ab6ff33d535f276c6b225d8cc21f7b7447e37b479116546f6f4"></a>

<a id="canonical-bfebb6fd0fe0ea64db8cb3a160e8988ce8e50a28e0c2a14ba3bb7be0b92d7218"></a>

## name property — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 4

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

<a id="canonical-23e8e497e2526f48a1c78cd7e3e10fd19975b7be2988151c2a315835d7280474"></a>

<a id="canonical-fec80a17c79c62369a299a56c3696998270719bcdce01464ce1ca92303ec6168"></a>

## overwrite property — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 5

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

- [secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd): complete subsection reference.

<a id="canonical-958fb0cf0cdf75baa1783e15e13f9a7ccb4547fe42c3874524c601f1825ee886"></a>

<a id="canonical-10ada42e5f8081093831ebfe3289a5ae62a3d3f95c136dcc5259abf509180371"></a>

## value property — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 6

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

<a id="canonical-53c5209b77a465bf3609f7ff44d0b70770a2948bd7e065230d16e7cc49cd0e32"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add / 2fa496ba4e96 / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9f6ddbf2fec0c97a8daae3ee13dd6371ee7c24110cdc5b39b06f15264d0628a"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value — http_proxy.more_option.request_cookies_to_add.secret_value / 1f56f18e54c2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6)
- http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0ecb558ae890d89890aa26e9d4fdefd65434cbebef436b8d0e7084704008d7c8"></a>

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

<a id="canonical-f8b7fcf2954c6fae9fb070c71c79bacf3eb757871e07583041f224482c5078ec"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value / 1f56f18e54c2 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-b088c719f2d9dccc7859ec710bfd031b1b5dddcd688a1e83ca9fdb59da0a27aa): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-cee96d303c7bd7fdd63c61e503f4bdde382479ee4628d2ec8fa342bcbac6af1d): complete subsection reference.

<a id="canonical-81eb6aeab56d33e863d2398209a453180ad8a5dc5fd90bce23c31aa9dfb17e55"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value / 1f56f18e54c2 / 4

- [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-b088c719f2d9dccc7859ec710bfd031b1b5dddcd688a1e83ca9fdb59da0a27aa)
- [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-cee96d303c7bd7fdd63c61e503f4bdde382479ee4628d2ec8fa342bcbac6af1d)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b088c719f2d9dccc7859ec710bfd031b1b5dddcd688a1e83ca9fdb59da0a27aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-551c310897b954fbecf922cafba17021951bdddc75df15c759e6d56cec5eea56"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd)
- http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-4412b42b2341068db086c5263e12538a9505eec39a6977245624b9047d824154"></a>

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

<a id="canonical-d4bc6c9485e86a35ad226dadff33ff0c6b71214b03d2fb3d9a72b696a963ea20"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 3

<a id="canonical-fdf6faad09357baa950a2006960baef9b8033b58bb5b678fd6921fa241176c93"></a>

<a id="canonical-c97476d68008b84fb1dd443d859e77379ce85dbc90f680e51d259879b4e75332"></a>

## decryption_provider property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 4

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

<a id="canonical-443bac7ee5f9f1c073fa62c6f8f5806f4323204f696018134019f1e61614c982"></a>

<a id="canonical-ae17c4b760fcd0130090c2782df653897d41aa03f65284a85f7f2802f190bc90"></a>

## location property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 5

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

<a id="canonical-f4804d0d3c5a25c2e33640e82f66380b4779eedd3b06df4460f846dc44024c05"></a>

<a id="canonical-b38ba362d6a58016fcc75584181e8ac1853e2750cb6f986f394b794ce2b1265a"></a>

## store_provider property — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 6

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

<a id="canonical-84c536e1b3b1db3a7b0966d985e9d6fb91bddcaa2d0a78f1f3ebb44c7a0868a9"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info / 2fcbfe0a6f03 / 7

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-cee96d303c7bd7fdd63c61e503f4bdde382479ee4628d2ec8fa342bcbac6af1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc76865c70ebf1dcb3f5c84fbf62df4e087b2c9df62ad598198807b13f1879d3"></a>

## http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / 0371248b741e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-753d9845eb049cf37c4371f270fdce9ad9098e09049b31434e15b350ec9e6df6)
- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd)
- http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-c9751314a80a655a95a209f0614bc6cc169cb290ee8bf88de5aabd1a874325d7"></a>

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

<a id="canonical-b2cc2c8020d6e31316ea071c671cfcae8bc32fd4af1a7843fe3cf15a90016bca"></a>

## Direct properties — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / 0371248b741e / 3

<a id="canonical-e85708957d0043bcce26015b337d44a5606ed9c5061620937fc92e3377e43787"></a>

<a id="canonical-7447a1d17632aac54dba49cc243ed4d2054f0e0e3f56774dc83c538fce28686b"></a>

## provider_ref property — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / 0371248b741e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a5be61a8bf79f9f239a20feaf2c04e209f4f07936a2cefc53cd5c08dfa2a20a2"></a>

<a id="canonical-56ed0deb0aa1e5842771bc697b5d6c0ab94493ddba7da182f5d6b1a499e5d4d4"></a>

## url property — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / 0371248b741e / 5

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

<a id="canonical-1e050c09fa5f0939605ac5d3f737368fd1ac426c8accb63d6ec487f3290656b3"></a>

## Next pages — http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info / 0371248b741e / 6

- [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-c4f845bc6f75b7ddc0ae3fedbc1a92c8a63402631b209b5421f9e521179a69bd)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2240c2dba2d541f0aea258d782cb4a2ca8f7ded88e08022fd849a907c02aded2"></a>

## http_proxy.more_option.request_headers_to_add — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.request_headers_to_add

<a id="canonical-6c196bb5ba382e4fa32433b6fe72c471d45388eb4b2632512f12726a724ec014"></a>

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

<a id="canonical-00b07f8562432725c0bc3f7740a6d59eab2440f0c2f4472bb377b179174b0d9a"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 3

<a id="canonical-c7a39f624ceb748ea3c07b8116996f9a2b442ca2a9dc3d8c72ae136b26ec3c6d"></a>

<a id="canonical-ab1461385f98347993a45e4c899b6b4af0c634233b7c93fed9d8807abd21fd73"></a>

## append property — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 4

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

<a id="canonical-631940ef57f7b03fdcb7f34a12e9a97088f458343fc46cdd98ffcdccd2e41b55"></a>

<a id="canonical-98c52eb8b613d6e1d7357a7ed31ab9da3deca33dfb5145dd33ff6731727aa967"></a>

## name property — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 5

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

- [secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8): complete subsection reference.

<a id="canonical-83b85d9cf7952cfb34ef6d3b2c59f4278fec4b489d7997cb3d12e07fd5421089"></a>

<a id="canonical-e08f65ee865827a3a88c9a2fa73ccfa48dee58c54512478305e1ce132b30d289"></a>

## value property — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 6

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

<a id="canonical-c678059b83c3c0c5a252eea9d3f7b006ad583ea53dbd7309f59c1f8881580298"></a>

## Next pages — http_proxy.more_option.request_headers_to_add / adb302fba6d0 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25f4fcf8c048208891406bb5cdc4c0236226bb1584459ea66b0537eee85015e6"></a>

## http_proxy.more_option.request_headers_to_add.secret_value — http_proxy.more_option.request_headers_to_add.secret_value / 8afc2d911454 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50)
- http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-b0b5802aa48c22bac8eb685faeac1038628495e389e245ce9c02ffe1eec374f7"></a>

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

<a id="canonical-72a35e0cf72c0eb9681a120a00b205aca67e96ea826dbdd6c1684e9e70a44632"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value / 8afc2d911454 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-948829310fa833971d462861f788dfffb67cb75264533b6ebfe78bc2dc5871b2): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-603639ceeddd51e623c85649b4d95b36195f66d8440144559ed709078232777c): complete subsection reference.

<a id="canonical-3a19a4fa68a23ac44faaa0c5052178956801176bc8fea47e46a9f43360fd92b0"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value / 8afc2d911454 / 4

- [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-948829310fa833971d462861f788dfffb67cb75264533b6ebfe78bc2dc5871b2)
- [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-603639ceeddd51e623c85649b4d95b36195f66d8440144559ed709078232777c)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-948829310fa833971d462861f788dfffb67cb75264533b6ebfe78bc2dc5871b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb5a7936ba7794711d3e74f5fbedf9f08bf8e611f428fffa715b02c9d2747f8a"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8)
- http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-4de9cf8db78a01aef63995366708a9be56f800f74d410afe21aa9543b8beeb3f"></a>

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

<a id="canonical-3466a5f680419bbb444b24d3e1b8f833b1cfa6bb292633a12385e08364df8774"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 3

<a id="canonical-d71f1817893768c60aea0c9008d95ac24d60d8256e232c5986cbe355a8323957"></a>

<a id="canonical-10ee2a7c34e5b085a106701a2ec151f9c26600c98a554c021da904f988fd73f4"></a>

## decryption_provider property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 4

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

<a id="canonical-ed547e4a40cc20e190d97588259f508cfcadcbf2ed5aef67eaf6207098d5a538"></a>

<a id="canonical-31326082f1cfbf928899c8ea9b4180c546cc702ae745878d4d21b95e9988b626"></a>

## location property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 5

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

<a id="canonical-87cfb1899da2b9bbd8a79577b722a55c290f88654e01a8ac6cd8ad24ca02fe08"></a>

<a id="canonical-338bf260b90b39a3780050dd6879ace124c1ee85341e6697519ef4436004b132"></a>

## store_provider property — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 6

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

<a id="canonical-c9aa1f61cc15f5ff7b1196560d8bc06dc90487312c2dac43d206737e10618e7d"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info / 1ddf86569a17 / 7

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-603639ceeddd51e623c85649b4d95b36195f66d8440144559ed709078232777c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b530deacaed7a82d8dae1e00bb5f039657323138d35cffdac44e183ca1b625f1"></a>

## http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / 8ba4dcbacacc / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-b40b6d89ddde8b7bd30344866e1adcf44a1bd04a682e682fb6f0832fde3d1a50)
- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8)
- http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-b5d408b3cbe5f94c165c4d6c7ed9fb2d9c562580db3605cada8378c481020c3a"></a>

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

<a id="canonical-aac769ed3c571276d01994de9c1d7b37825754a733aef2ed260f773709f35a46"></a>

## Direct properties — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / 8ba4dcbacacc / 3

<a id="canonical-065d036ffe024b38b38b931c6b87a1657c64cd6e6a4fd70eb794796488247451"></a>

<a id="canonical-9454530518e9119f87af0c3e6c7e4d326da4d819011274133b4b126eb59fe375"></a>

## provider_ref property — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / 8ba4dcbacacc / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-348534e651d7c123fef76b4607e86a2527ba62b13a8274f303f18d9b44788c94"></a>

<a id="canonical-a904c6e190ef570c86118f5f3e9a731594fcb525fbbaf2f563f19a98ce95e3a1"></a>

## url property — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / 8ba4dcbacacc / 5

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

<a id="canonical-8885bd3d51b2dc1ea1b07be748f9d705dc4af381ab55fd306919d68961c55f2f"></a>

## Next pages — http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info / 8ba4dcbacacc / 6

- [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-69224cdbbbfa17144aae083dc9796bceee915ba7375fc993f9731295ffca97f8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a11779065de2d43cebef18043a67c31ce8775446042c99897ee6b0a6590f8cbc"></a>

## http_proxy.more_option.response_cookies_to_add — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.response_cookies_to_add

<a id="canonical-545ef34733b875f9551ae14b2d9baab2b50a6796d237dffd48333a42cb4e9e09"></a>

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

<a id="canonical-672d629a94985443993f3ec9b884d869d63654bacbb40a637baf578d7472680f"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 3

<a id="canonical-c7076a9671d4e51dd81f4fcc993dbe058b861a27f8879a4a5768878c8a0cf517"></a>

<a id="canonical-95bf39fe241cd11281254f3f5b1d90b95bfbbf271167cb59818e24600cf4f9a6"></a>

## add_domain property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 4

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

<a id="canonical-ee46a5c02108c1869d9762111b8777db1cd77392ac05c468dad1d18fc7900fee"></a>

<a id="canonical-3bc99eddfe8e1a39242a16f4665b5afe0da4958a5aa3b2805ce96dd6eff4c319"></a>

## add_expiry property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 5

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

- [add_httponly](resources--proxy--reference--group-004.md#canonical-3a1e2f848cb536adedcee8f38e8f25d68d300f8ddef2776207bfe7095da62558): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-004.md#canonical-d8ff0808e81656a824b9344a15df27ff8f538183e31e2945c3aa98654d3b4ab9): complete subsection reference.

<a id="canonical-0bd5977ae42cd2014585edd03699e1049b769681d8eeb70f7ed30eca80bb1b42"></a>

<a id="canonical-a7e481b4a2e1d84e146edc1973e57d09443ee150c043ec6de1ec6cddf0346099"></a>

## add_path property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 6

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

- [add_secure](resources--proxy--reference--group-004.md#canonical-585cd48ff5af07a20355c59f4dcb95e481dfd4e7caeb4008f96b15b9d6571186): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-004.md#canonical-2c574b680ab2ee0f81594e9e09e09c70c6601258ed797b5f8048fc7628a696fb): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-004.md#canonical-8fa82314ef8979e297f632ad3865418dbebab19723259976e40d39054e9642bf): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-004.md#canonical-38054b9cf6e1c3659a5c36f3a607ad1ae1daee11b2a6fe154bf760e8073dcc19): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-004.md#canonical-c8fc7d6189d4c0c04f3ec6bc8563fde7037b55fe4f38411ce445ed3050bc127c): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-004.md#canonical-1b54582132ad744de60807ff15166e895399f67a261fafcf2c9db8bc19c010da): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-004.md#canonical-39ee7d8e05a72451e143a4e39e5016bf53903893b9431e41f0da984ffa27ccad): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-004.md#canonical-95801219d4880a6c0bcd0c1e5a30e1744963749192cb0b9dff15e3a8e663df23): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-004.md#canonical-0e61c272211cf02822fab53c8d1e0f5b50409f94df32248da793da0accab1a6c): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-004.md#canonical-4da8b42edfe31b7c36d9558644ccdbd5c380008f0859125c698fd1b6584d7278): complete subsection reference.

<a id="canonical-b43bee81e88230fe4a110156e8b7ae5d667201aa40719ca3fedd06f842566d3b"></a>

<a id="canonical-dead48bd236d80c8c0b776b8b661906008dedc7436cd625af5bca9ca54850280"></a>

## max_age_value property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 7

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

<a id="canonical-8e6b681d95f879a349f4dc37b7d654fef2b3ab9aca1ac0993df170266f538586"></a>

<a id="canonical-ac24ea50992bcf9abf4d140354508b54f58cfcbb8ea50e3859723b89c968eccf"></a>

## name property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 8

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

<a id="canonical-dbcbed562d99be4ef0311c142458f5395cbffbc7e4b3bc5df8df0dd886f3c90c"></a>

<a id="canonical-284ba120b66e82ab4064566d895534adf60e9ff69f991ffa772d255bbac398aa"></a>

## overwrite property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 9

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

- [samesite_lax](resources--proxy--reference--group-004.md#canonical-664c563f2b049b5f79cea95590cba93629ff5180eb8dd09abe3a067d2535c579): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-004.md#canonical-defa6f3a22d2fdea832882b2d7b27c192088599748d136007b444bc2f6aab988): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-004.md#canonical-aa3a6a043e813e529d14b052f43f97a713ea290428dad971f4f2a81bf89009f9): complete subsection reference.

- [secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea): complete subsection reference.

<a id="canonical-b7a32287ce5ecf96d5293f3450c02fb1523e8565486ade4a103ab9ed1f2f05f8"></a>

<a id="canonical-664f908addc154b97297290fe8de03084df5a9c985954d9bb3ff73a2f07cf91d"></a>

## value property — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 10

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

<a id="canonical-2787ee811040b0efc56e6da124d7b690f3e3b019fba0bb385cc040bf5d51e3fc"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add / 19730d32323c / 11

- [http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-004.md#canonical-3a1e2f848cb536adedcee8f38e8f25d68d300f8ddef2776207bfe7095da62558)
- [http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-004.md#canonical-d8ff0808e81656a824b9344a15df27ff8f538183e31e2945c3aa98654d3b4ab9)
- [http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-004.md#canonical-585cd48ff5af07a20355c59f4dcb95e481dfd4e7caeb4008f96b15b9d6571186)
- [http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-004.md#canonical-2c574b680ab2ee0f81594e9e09e09c70c6601258ed797b5f8048fc7628a696fb)
- [http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-004.md#canonical-8fa82314ef8979e297f632ad3865418dbebab19723259976e40d39054e9642bf)
- [http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-004.md#canonical-38054b9cf6e1c3659a5c36f3a607ad1ae1daee11b2a6fe154bf760e8073dcc19)
- [http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-004.md#canonical-c8fc7d6189d4c0c04f3ec6bc8563fde7037b55fe4f38411ce445ed3050bc127c)
- [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-004.md#canonical-1b54582132ad744de60807ff15166e895399f67a261fafcf2c9db8bc19c010da)
- [http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-004.md#canonical-39ee7d8e05a72451e143a4e39e5016bf53903893b9431e41f0da984ffa27ccad)
- [http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-004.md#canonical-95801219d4880a6c0bcd0c1e5a30e1744963749192cb0b9dff15e3a8e663df23)
- [http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-004.md#canonical-0e61c272211cf02822fab53c8d1e0f5b50409f94df32248da793da0accab1a6c)
- [http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-004.md#canonical-4da8b42edfe31b7c36d9558644ccdbd5c380008f0859125c698fd1b6584d7278)
- [http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-004.md#canonical-664c563f2b049b5f79cea95590cba93629ff5180eb8dd09abe3a067d2535c579)
- [http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-004.md#canonical-defa6f3a22d2fdea832882b2d7b27c192088599748d136007b444bc2f6aab988)
- [http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-004.md#canonical-aa3a6a043e813e529d14b052f43f97a713ea290428dad971f4f2a81bf89009f9)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3a1e2f848cb536adedcee8f38e8f25d68d300f8ddef2776207bfe7095da62558"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23dee913b09b80b7727c8581797c7804d8e4308aedcd7a98ecebf022856cfeb6"></a>

## http_proxy.more_option.response_cookies_to_add.add_httponly — http_proxy.more_option.response_cookies_to_add.add_httponly / f340be83c805 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-b0a5ea0e66cc529b1215fcefa6a4a6ea0304a00f807109194a1c4b5ca7a69c7c"></a>

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

<a id="canonical-103ca3c0f0f0bc0fe9ed5beb333df957da716eb9b152eb48ba757cfdee975e8b"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_httponly / f340be83c805 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3aeb64f8d34c907a8447a0df7d5146a5a16b3b89807a07b0a97120c7d6b93c47"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_httponly / f340be83c805 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-d8ff0808e81656a824b9344a15df27ff8f538183e31e2945c3aa98654d3b4ab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dff5bbcd503e8bbb64f3c51fb682c1e14cfd113f1e2ee2177164a76ee9dc1972"></a>

## http_proxy.more_option.response_cookies_to_add.add_partitioned — http_proxy.more_option.response_cookies_to_add.add_partitioned / c305ded3af96 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-00a00748160b4aac76f84f81e8b0ad10be7a37de46f37d3135e8bd5c33c7e769"></a>

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

<a id="canonical-f0d6433e238ebfc0e57961616a145147452b81b96cf9d416b6f7b4a1d7b9e7c9"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_partitioned / c305ded3af96 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-555c5068a1e0221ab3213455c0add8208992d805e91664648cdb11d03a6f6f11"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_partitioned / c305ded3af96 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-585cd48ff5af07a20355c59f4dcb95e481dfd4e7caeb4008f96b15b9d6571186"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c57f3be1cd4ff7670024622b6599494956d42ac74bd8af63c140868b137b8f6"></a>

## http_proxy.more_option.response_cookies_to_add.add_secure — http_proxy.more_option.response_cookies_to_add.add_secure / 200efd196272 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-7f7ce3ae882c837c8200be2c7302c9be18c64833590395e0752e5a126d776b2a"></a>

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

<a id="canonical-61ee7455cd3eaaabb4bf601edbfe77b02818129fd2bc04d3d56c9eca5423e13e"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.add_secure / 200efd196272 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-144f19684c941e5eb35f3256213fd6ef5cef26051d5c54bb2c62a62a3d85d218"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.add_secure / 200efd196272 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-2c574b680ab2ee0f81594e9e09e09c70c6601258ed797b5f8048fc7628a696fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd47bf12f0c26971471ab0d340a09ce6e6e8dcc9b82b87a54569ee1aee4f73f4"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_domain — http_proxy.more_option.response_cookies_to_add.ignore_domain / 7a0e914e03c4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-2986c2f629ee517bfe63ebd4be28a6cd1a25ef5e9524877152b07f903531329c"></a>

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

<a id="canonical-e4ec9cd4e7d2e43e0a87208fc0735818ea0c62ac3c6c31b13af263799b84b1fb"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_domain / 7a0e914e03c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8946c7fe8a55ce0ddde84f0143c7e0bf7167935d8ebe12ea9a43617646809314"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_domain / 7a0e914e03c4 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-8fa82314ef8979e297f632ad3865418dbebab19723259976e40d39054e9642bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b4d30556cf0d97efbc1fc4c579a9291475c097ecd72a6d83b9c51eed9ab41c7"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_expiry — http_proxy.more_option.response_cookies_to_add.ignore_expiry / 22d743e0d4bd / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-f1a5a9ed7d3dd7237b088543cdb298ed354b760a8198f7c5806452965f275eec"></a>

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

<a id="canonical-e87a3181fd105d4d21ff3a67a96fa30a7ae82e6d95c230dba40a6b9b4b6a8023"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_expiry / 22d743e0d4bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4c4fbb1e7341916868db25f7c4b2b2fbe9453d246f4c2296e09cfc66e3ab581"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_expiry / 22d743e0d4bd / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-38054b9cf6e1c3659a5c36f3a607ad1ae1daee11b2a6fe154bf760e8073dcc19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6350c97b21925ce674153301d81cebc648585dd6cfa11045543ec7c7e661463"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_httponly — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 65b69044a0eb / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-c98625cec1cde80b40a34e910d2db392edbd9ba2bf2cd60bd89e3b88b37a2fbe"></a>

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

<a id="canonical-03d9960e916072e70cba8f40fcb1687f8cf779c8242b3968c77a91869c02f06f"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 65b69044a0eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14487ddbfe19865d68a40d78ec7f10660eaf1a6bdd6e7bba810c18849e6113d3"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_httponly / 65b69044a0eb / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c8fc7d6189d4c0c04f3ec6bc8563fde7037b55fe4f38411ce445ed3050bc127c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-268eaf424c34fa590bc64727d8bb44d88e592b1ea69f1f3fa5bbde45018ea0f2"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_max_age — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 3bc77680e2ec / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-6844a4b0437cade383668be426bd198e6943d9c15d12c971cf791093caef32b1"></a>

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

<a id="canonical-3727bbd83e6d08b60df1549ba26c0137833962ddd8aa5722e784efdf964539d2"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 3bc77680e2ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76bcb2b8e691f459ea0b737b7801aa0661322ae5ac0de6c473d1dfe7ad20717a"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_max_age / 3bc77680e2ec / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1b54582132ad744de60807ff15166e895399f67a261fafcf2c9db8bc19c010da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d24f1df61945e124919dffe9cf850cc9470fcb6c059f1864c850c19efc53f8f9"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_partitioned — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / 2b2b8d529405 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-9046fece5c9e5c4c01c186be7433ca4aa81a4d5f7c91f29c2d321dc3c0947cf3"></a>

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

<a id="canonical-eb569aaca50619bb0fc897d9584aa20a19a965a3ce9c1ceb099fff365c1459b6"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / 2b2b8d529405 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89530543d873033d6af2d56eaf22c247e6624ab84e0610acaabf3f73239f6fb6"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_partitioned / 2b2b8d529405 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-39ee7d8e05a72451e143a4e39e5016bf53903893b9431e41f0da984ffa27ccad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89defb04917a98cae43b1eba8adf63bb82609628b423f0bc53378ed70c3caf8a"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_path — http_proxy.more_option.response_cookies_to_add.ignore_path / ce63b5fd4ad1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-dcbae0a311485a0c2b8f64799facb1b761f380afeda5f02eadb2f24028a9e837"></a>

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

<a id="canonical-ed41d6d7cd619a3c7d1e1af1bd2ddd5dd3db468a9d837f6ae2df3ce1a983ab47"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_path / ce63b5fd4ad1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f335688b4e421c83f2407fb692d3108a915557510ba55f92982ce4301b055a0b"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_path / ce63b5fd4ad1 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-95801219d4880a6c0bcd0c1e5a30e1744963749192cb0b9dff15e3a8e663df23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9097e566eeb65447863c47cb7034e95917f34788b1f97aa5d0d48c569090daf3"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_samesite — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 9f8117ce7474 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-df6c6072c3b545418789d6dc942451d41d6f02d65db1baa387d86d0aeb5433f4"></a>

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

<a id="canonical-f35fd64c35ce60098f9b53974aa802dd45a9ccca6c7e2e49b15e238fced0b806"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 9f8117ce7474 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afce0268fc129a26ab86ecb5717a9495b118e722fa7d98be97e0bcc8cdc921dc"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_samesite / 9f8117ce7474 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0e61c272211cf02822fab53c8d1e0f5b50409f94df32248da793da0accab1a6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9209e523228669b033c50ea5fb72533e040f76f70f8a9f76416aa08c1fdadf0f"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_secure — http_proxy.more_option.response_cookies_to_add.ignore_secure / 16957c181fb6 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-6fff4dc454de26eeca67acf979f13907a0aeed98404df50adb79c05de433e6be"></a>

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

<a id="canonical-4eed5db3d48c871d05bad8abeb4e20852027a44015f0534435596e34ca88fae1"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_secure / 16957c181fb6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3913f05f9497f7e530e8593a3bf2a32af638658f4d854c240895fe3318c86690"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_secure / 16957c181fb6 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-4da8b42edfe31b7c36d9558644ccdbd5c380008f0859125c698fd1b6584d7278"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d09c6b73dc7b8387eebd4d6610edc5878ccc898fea070895276d23d7a3a1f98"></a>

## http_proxy.more_option.response_cookies_to_add.ignore_value — http_proxy.more_option.response_cookies_to_add.ignore_value / 33d3cb2f76e7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-8014117b70769794f36939ff45149cc9869abba103ba5bfd5e7c7eb408b20a19"></a>

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

<a id="canonical-a6e0c431e39696a0d4ff5bba6bed29370a94445399200aaa48b988a5a3723841"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.ignore_value / 33d3cb2f76e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-489d3f94620a075b6c74d9522160af3e8fdc5226f8bf7ad2079b6812ae6d4550"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.ignore_value / 33d3cb2f76e7 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-664c563f2b049b5f79cea95590cba93629ff5180eb8dd09abe3a067d2535c579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bed31d67cf7fe7b96d86bd1f963079c937ea19dde2dcfc748aedc59e25d3c8b4"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_lax — http_proxy.more_option.response_cookies_to_add.samesite_lax / e45c142f464a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-3b02cdad8394d48d53050bc062d51fe6c8bb3cd583b353c2865d54219afa2925"></a>

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

<a id="canonical-8d62db8a5e11dffe5abcea1985491e988eab3e5c7b8a136f7569338a4387b4e3"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_lax / e45c142f464a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75e7cfdbfa60aa7848777c4402fb377769ac66f7cf62d3839d1b27071034e431"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_lax / e45c142f464a / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-defa6f3a22d2fdea832882b2d7b27c192088599748d136007b444bc2f6aab988"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-564c9c28b36cb29349a55b4ad8ecb86886a5daedbe1bf498d3c216388c57a28a"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_none — http_proxy.more_option.response_cookies_to_add.samesite_none / 40eb311a41b2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-3f3285a4892d5b931a61e9ef8f89f44854da746ffd93f795df9e55511cfa5153"></a>

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

<a id="canonical-ee48cc18b1aa469fb949d5c382b8daa218bbb413989ae89bbd45e7d30ad454e8"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_none / 40eb311a41b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-742c3a8dc6a2b24109301ef6741b3d1d00852922f9b2488eaa70e8ffbb2be4cb"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_none / 40eb311a41b2 / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-aa3a6a043e813e529d14b052f43f97a713ea290428dad971f4f2a81bf89009f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8cf0e2aba86991c5d81ff1771d000d5413fa316ff87258464c3650c99876128"></a>

## http_proxy.more_option.response_cookies_to_add.samesite_strict — http_proxy.more_option.response_cookies_to_add.samesite_strict / 906a4d6017fe / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-55a9c9a31daec7b5c4d6ebe8d5b44bbaf1a4461ce50c676b33a306f6324a406a"></a>

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

<a id="canonical-e9980f321dd2952679de3ff196579252930abe21ed0065c03dbb2d9be8994d63"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.samesite_strict / 906a4d6017fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b920ab7cde61696b28dfc13d8a0f2473fdfce4a309d61eefb990fb2c88852108"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.samesite_strict / 906a4d6017fe / 4

- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81422762bac25bdbbeaf79f6e2ecc8ccde9c42e3120afc85e149b9112bbe00a1"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value — http_proxy.more_option.response_cookies_to_add.secret_value / 0bc07683aa25 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-70978771232ddc64734761c156cc39a218f50403811d217b907d5be3d1172728"></a>

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

<a id="canonical-404d0fa9af869ef2ad581f6c2f225ec2c27cd5bb08c2468967fc6dda606e02d5"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value / 0bc07683aa25 / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-645efa6c63edac168653af349b9166f19460eb9520afe674daa51ff25eb5104d): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-95e21122411bbea12fd9050e9b3260a9fe6fd0da1a432f3faf440ddc858010f2): complete subsection reference.

<a id="canonical-f65eff7d525d4a3b7e6c5b6cd5a9910a40851b19d35ad1160cdc947a4809cec7"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value / 0bc07683aa25 / 4

- [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-645efa6c63edac168653af349b9166f19460eb9520afe674daa51ff25eb5104d)
- [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-95e21122411bbea12fd9050e9b3260a9fe6fd0da1a432f3faf440ddc858010f2)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-645efa6c63edac168653af349b9166f19460eb9520afe674daa51ff25eb5104d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10f36c602af3bb3054ad5b932e83708a4fb0a8b63afdc7199d85328ba095acfa"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea)
- http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-ff6dc40090338ddf4be9cc24946acb4b4353475cad299ef22189e400371deace"></a>

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

<a id="canonical-6803497b4cc8e6b1a682504627d32902867600e3fcf6e1873148f1e8c7693424"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 3

<a id="canonical-7a467bfb090ea249374ef8a3367f5e171df5c48a194d569201bc9f03cc3fd545"></a>

<a id="canonical-0ee4fb4d1400019bf56e88d572d8b1ecd573cbd292cac69f230200da6edc0bae"></a>

## decryption_provider property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 4

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

<a id="canonical-f45a00ef3cd26d8656bd2278d5a86d4731122e304b276a350e3968ae36a6e6d1"></a>

<a id="canonical-c1dd66fe054dbd9d0cfbb0739d29c31fa8dfb8ac540f33f64f917e47ff8f56b1"></a>

## location property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 5

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

<a id="canonical-19360d722940b59c5ca068921a526103fa7689af6482775757326067b34301a5"></a>

<a id="canonical-1149402cffb43fdf9d7abd866cc45f1f1c248556ebf08a028218d5f2354f046c"></a>

## store_provider property — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 6

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

<a id="canonical-4ebaf7253b4dc50d39b7167114dd0c79de32db54020d9d76d5810382d2ed9bba"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_inf / 4a4e2a44610d / 7

- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-95e21122411bbea12fd9050e9b3260a9fe6fd0da1a432f3faf440ddc858010f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3573c55776e8c1bf165b601f2668ddc8dfc2c8219acae55f4736a4f71651ce3"></a>

## http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / 8ffea88930e5 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0006d2e7bfbba80251af65207e38a2ebe4732051e49606c5e342c2b0d4035375)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea)
- http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-904e054a6ed6ecddf28200b90c3776b51c50868debd6786e53a83289d6ddfaf3"></a>

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

<a id="canonical-e27373894d01505ddc5b0559d27591b34477d67430b103333210b9e6b9ea09a7"></a>

## Direct properties — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / 8ffea88930e5 / 3

<a id="canonical-3e609e559185c932697f9d7758d9514932b449cd2556274ea4d1fd767afc7c42"></a>

<a id="canonical-bfd5c881d2a172cba66c43af01fdfd19b05cf998f2f6c40c6a426a27412eaf9b"></a>

## provider_ref property — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / 8ffea88930e5 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6ef38a2804404b4943c54ff293a25b10a32f222908e0e537b15ece3554f550cb"></a>

<a id="canonical-b30bb0f86c2540c068a68d8dec4441f38292d8751b3a3220ed2a2f3d83f5909a"></a>

## url property — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / 8ffea88930e5 / 5

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

<a id="canonical-111d997b5ad7db4e9b2f10f320c901eec5ba1c809632ada4b406b3462735e512"></a>

## Next pages — http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info / 8ffea88930e5 / 6

- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0eae8f799a4ffeba16cb9e2f6a6356d8dae1043b8047d2ad381ebd4e143f18ea)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74edabe5004ba1217e391ee3a4ded24a581c91817a3928bf5cfe4dac76214c2c"></a>

## http_proxy.more_option.response_headers_to_add — http_proxy.more_option.response_headers_to_add / def95620fec1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- http_proxy.more_option.response_headers_to_add

<a id="canonical-7ff679e8770a6c67354618df1271d147d18c7c7349c4f600e7907b9369eab9bd"></a>

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

<a id="canonical-d29d922cfbd631597dc8f1965b485296c7f496e2d2a66d81973d609507a27365"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add / def95620fec1 / 3

<a id="canonical-f3e18e8971ea7dc96f9b6d473be8ee1766ec117c3224c928a17859a676355333"></a>

<a id="canonical-8c40ec40a3d69d39efe7fdab2379679ea8608d109a3827ae6725e6698cebf580"></a>

## append property — http_proxy.more_option.response_headers_to_add / def95620fec1 / 4

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

<a id="canonical-f7a4fe2bf2a5f52eca30d61cf8551572dafb6ee45a9b312118650c5f913d6f79"></a>

<a id="canonical-fe0d96901e013288eb6ca625adccac573cc963e3d5c0b203c4dfd9aaa95edcc6"></a>

## name property — http_proxy.more_option.response_headers_to_add / def95620fec1 / 5

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

- [secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87): complete subsection reference.

<a id="canonical-acaa10f37ce096d6085ab6742d515ddeca971deadd3f1dc80d6712f787e463ca"></a>

<a id="canonical-42cf8703c18566d2081d128aba53ee6c0af5859f8fb0e4f32588a2b68337c495"></a>

## value property — http_proxy.more_option.response_headers_to_add / def95620fec1 / 6

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

<a id="canonical-6c9592a2e51ff600c702820dcf0af5c58cd13aef8d7845ded5fbc9470a2851e3"></a>

## Next pages — http_proxy.more_option.response_headers_to_add / def95620fec1 / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4feb5bfe370be5289b51b8990098dcbc43ff6744e62f79010f46a4a48b98089a"></a>

## http_proxy.more_option.response_headers_to_add.secret_value — http_proxy.more_option.response_headers_to_add.secret_value / a0ba6884a66c / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-d16dd2fe6edda61da9387cff1a885013e71a7e5578f538b762793190950bab9a"></a>

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

<a id="canonical-feb3e2833907f74b64d425d0e427947d63b125397a230f68ffeae186c7f71d74"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value / a0ba6884a66c / 3

- [blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-416561f381abd5541d2f1d8820b5a8846bda0a8eecf3ee2cb5023e649e681ecb): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-004.md#canonical-441c87824f9a2a45bea093f3370802f8d7ec71695354c362ef8b6b81fca4e20c): complete subsection reference.

<a id="canonical-8b59fd7bdfac874f400966753487316fa257c75a307c632a8a8f87469c631ee1"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value / a0ba6884a66c / 4

- [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-416561f381abd5541d2f1d8820b5a8846bda0a8eecf3ee2cb5023e649e681ecb)
- [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-441c87824f9a2a45bea093f3370802f8d7ec71695354c362ef8b6b81fca4e20c)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-416561f381abd5541d2f1d8820b5a8846bda0a8eecf3ee2cb5023e649e681ecb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6bd4c31f5ac01b9b911938a5e74999c7d79b7f31c7e937346f2d30fcf1422e4"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87)
- http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-197c254a4202938d805c9298d2c8aa9a58bbbccb7fdb18736c3596a264e16b77"></a>

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

<a id="canonical-d88301eb87757014dc280864db175e1783e5bb8f1b645c080ffe91e143a01697"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 3

<a id="canonical-9499b04e51c9f08592afab1e3b17b580f61869cdac9ec543e954bfb9347aeba4"></a>

<a id="canonical-17e79f2ef517e05deffaeb9fb200c6daec9576363c11eb7c54aeb7649ca904ba"></a>

## decryption_provider property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 4

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

<a id="canonical-ed693583f7c8f43d7942802f47a000a37d66c00e14b741f502963120f4d835ea"></a>

<a id="canonical-b8660e7ec4e4d66666d2d4f66a5de3af01817ecce1caeb81ad387403440b67ad"></a>

## location property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 5

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

<a id="canonical-c74c03120b6c938b79d3f2deda3af57da22ce1dee6bc0db7ac0a0a1b2b190075"></a>

<a id="canonical-aaf0769887c4028ee2cfda5635c6966c7383803abad7c26cbbc65cb414b12b47"></a>

## store_provider property — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 6

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

<a id="canonical-85d42f7fe585a02e47b768a5f3c0488f75f1d0cdbe7cb828875335de38a94412"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_inf / d0caa37fc843 / 7

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-441c87824f9a2a45bea093f3370802f8d7ec71695354c362ef8b6b81fca4e20c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-293f8b4bc48743eee1736bdf0d12fc4628db666c1280d0d98105c2765fa0e848"></a>

## http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 29028e75474f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-c11a7131e95fa79f43b25ed5a9688c6666ef887b9398e6433e72e59a3614cc43)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-39ca663e43d03e11bd97a83c6d29b48c01e9f4ca62898f4736593741eff6004e)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87)
- http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-a09d6bb0f7f4ef2962e010690849f16d7077f28128846beb7e7a91a3c3c7423c"></a>

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

<a id="canonical-115eb101c9491dee9b98933416dd8a90ad1432e5c834af762b14176e8f3cdf41"></a>

## Direct properties — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 29028e75474f / 3

<a id="canonical-48b2c057aa8af7cdeb4d8f98c2eb90c892d26b12e738f3f8eec0c8801a1d6c08"></a>

<a id="canonical-93a7e1aba926e928f2b2af65da4a7d83610cada6a1cdb283cc03b4dc61c332c4"></a>

## provider_ref property — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 29028e75474f / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3c1363464d9600ce5745864013e910035bb508737c6a6e4287aff270f5dfc59a"></a>

<a id="canonical-8064c5ae076f4e2ee7990407711a573088c6f8c542b871229487675e07f58cd7"></a>

## url property — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 29028e75474f / 5

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

<a id="canonical-707d1b9e22db97850c5b2f011782dfb48a76d1b7a5bc05396787cb5b0aac4c0c"></a>

## Next pages — http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info / 29028e75474f / 6

- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b014f92ce02e66b2df3a235bcdd810b3a5238cd5e41702b8bd46b6825ce9db87)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-ff779654e9d7a3b0554f6fef391a194052fe154808cfcc51b7979ca7298d8d2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d87a6543f4586ee224ece439bc3349b888c87dc071caecd87090d5edf94864a"></a>

## no_forward_proxy_policy — no_forward_proxy_policy / 9b44ac546c91 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- no_forward_proxy_policy

<a id="canonical-72d597129d8e04f873ef78a4172e75b919979521b55b6cd8614ee84456067dde"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_forward_proxy_policy = {}
```

<a id="canonical-5ba5646c3a9ee25b3a972a9c8c0147129c11d12292053defc1015aefeb3706d4"></a>

## Direct properties — no_forward_proxy_policy / 9b44ac546c91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-666741a2372b3d41ae5cd1d39a2822240b3af252056724c4ff1b7d0280dd83eb"></a>

## Next pages — no_forward_proxy_policy / 9b44ac546c91 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c1d4991763e98c65835bf4a35d131a34666e553de1f223df41092c700befd9fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e23b79edac8b7c3a84628f0f34d64e2fc39a3349449955847a9fdc11d431b15"></a>

## no_interception — no_interception / 5bf1c1d32870 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- no_interception

<a id="canonical-fe2f7d86c6d293b9d1825c60aa5a53f34fd4db2312fcfee1d1e64b58eff6901c"></a>

Type: `["object", {}]`. Optional.

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

- [no_interception](resources--proxy--reference--group-004.md#canonical-fe2f7d86c6d293b9d1825c60aa5a53f34fd4db2312fcfee1d1e64b58eff6901c)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-b09947c5b3a5d27f69f31c0a2da06055276dbe6a003232d163a89b3739b4716c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_interception = {}
```

<a id="canonical-92a77cfb76c73bbcaa9f16c292c78bcb01adbcbd2541e85d018fd8e21892394a"></a>

## Direct properties — no_interception / 5bf1c1d32870 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80790110f881fa62537af7e140569925d1391154aed224d79a9c53194f12928f"></a>

## Next pages — no_interception / 5bf1c1d32870 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b0a1caa8391902d4f6610c4dbed1b5392a4fd7331ae15afa906a7878ea3a2f40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a54020a1b8e3dfb4088de3e215ebe4904936c9b282fa244267387a1695d7bb3b"></a>

## site_local_inside_network — site_local_inside_network / 4423b353ab2e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- site_local_inside_network

<a id="canonical-d96f7e5101cc20f99673e7fa523e66f77389734b46fd6469a694eb553bc6fe30"></a>

Type: `["object", {}]`. Optional.

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

- [site_local_inside_network](resources--proxy--reference--group-004.md#canonical-d96f7e5101cc20f99673e7fa523e66f77389734b46fd6469a694eb553bc6fe30)
- [site_local_network](resources--proxy--reference--group-004.md#canonical-e8ebfe418e106ba12def83255b7186488931618fa70331fc01027565e775fa8b)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
site_local_inside_network = {}
```

<a id="canonical-5c509bffb4416cda2e00b147d21e46c3a7fb952d289d52e09f234653a43b3b7c"></a>

## Direct properties — site_local_inside_network / 4423b353ab2e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ca8c1b61842690ace685722f3a83bb8c25e56f0a99e74f8b215e9b8b61aae72"></a>

## Next pages — site_local_inside_network / 4423b353ab2e / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-73781e83eabf108576e0ba5c0ee0ead9ada60a44a59a56ea627eaf9a0e9570c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e36d73335a3cd5c66f8e5e0c57f5fb2ab77b66659b26befb0c2b9ce58a72c15"></a>

## site_local_network — site_local_network / 88a64b2d5c95 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- site_local_network

<a id="canonical-e8ebfe418e106ba12def83255b7186488931618fa70331fc01027565e775fa8b"></a>

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
site_local_network = {}
```

<a id="canonical-8dbf0402c3bb715c759537c4beb3880646fc86efc8d98fd3742a35bd612701b0"></a>

## Direct properties — site_local_network / 88a64b2d5c95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ec0c7b9a347bc6a483cd6897f7a7518c1bfc1eb653ec02fa64c8faab1f5e675"></a>

## Next pages — site_local_network / 88a64b2d5c95 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b2a2f8acd2d8b49033cd0495a231c8ff8d285b615ee07ec6d8a571fe7c1e852"></a>

## site_virtual_sites — site_virtual_sites / aaa0f4a63e02 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- site_virtual_sites

<a id="canonical-3f1690011579c5952d206b3af1687ef51b17cc9a189ce572f1a79976d6cc93d4"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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
site_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f5364676742f27eb72e51b77a6a93444d37148b62967599c3ac7ba090881e9d"></a>

## Direct properties — site_virtual_sites / aaa0f4a63e02 / 3

- [advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402): complete subsection reference.

<a id="canonical-e48dec8ef981b82fc6ca4eadf132767505657fb00c20b44a2010b1c0263f6498"></a>

## Next pages — site_virtual_sites / aaa0f4a63e02 / 4

- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c6835d3fc93f26829ac35a305e658fdb12ea2707cd997c042d89552dc1347c3"></a>

## site_virtual_sites.advertise_where — site_virtual_sites.advertise_where / 02b330599615 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- site_virtual_sites.advertise_where

<a id="canonical-90521e29fc9d52b0d5804df53d52724d9bbca2cbd91a51baea9dde19237f4bc8"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site")}
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
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-683f9d7c0ad4476d5105e89c8e76a0dd276a3e7f803386e6a686ec81d91aa8e9"></a>

## Direct properties — site_virtual_sites.advertise_where / 02b330599615 / 3

<a id="canonical-b7f570ad2bdce578996dc5d596091f48c794c7143894120436ccc71bd0f42ad7"></a>

<a id="canonical-0f31585800d0d3edfba01969ec353545709a07f277a3da11f454571c67161f3b"></a>

## port property — site_virtual_sites.advertise_where / 02b330599615 / 4

Type: `"number"`. Optional.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Upstream description:

Exclusive with \[use\_default\_port\] TCP port to Listen.

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

- [site](resources--proxy--reference--group-004.md#canonical-948dc3edcc83805f304b7510f2dd275054200aab5bd1fa20497144d25b107707): complete subsection reference.

- [use_default_port](resources--proxy--reference--group-004.md#canonical-86401157f683146dca76cf73b8a7219a29a80e1b91a563f555f505e876a90e74): complete subsection reference.

- [virtual_site](resources--proxy--reference--group-004.md#canonical-56cb8cc6f75ca94245ca58189fe0b09826f50a728714c3ad33c01761194c46d4): complete subsection reference.

<a id="canonical-99f890450eb69cd9687e0d7176925ae1cb0d9ebeedd9b1faf7414e29c2f1f98b"></a>

## Next pages — site_virtual_sites.advertise_where / 02b330599615 / 5

- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-004.md#canonical-948dc3edcc83805f304b7510f2dd275054200aab5bd1fa20497144d25b107707)
- [site_virtual_sites.advertise_where.use_default_port](resources--proxy--reference--group-004.md#canonical-86401157f683146dca76cf73b8a7219a29a80e1b91a563f555f505e876a90e74)
- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-004.md#canonical-56cb8cc6f75ca94245ca58189fe0b09826f50a728714c3ad33c01761194c46d4)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-948dc3edcc83805f304b7510f2dd275054200aab5bd1fa20497144d25b107707"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a240f243e4597eb22baae69b9a30441801d9dd9ae1ad8ec2b4bb66240e74f5f"></a>

## site_virtual_sites.advertise_where.site — site_virtual_sites.advertise_where.site / 97b322a08c1a / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- site_virtual_sites.advertise_where.site

<a id="canonical-fae667cfd5b27beb8ab1677ee24b938d147ff538f23a4dfe33f6680ef4114ba8"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-53ec2f8ecb31c5e4d8f388d2b190c118925111be32518dbe390189d291e59c78"></a>

## Direct properties — site_virtual_sites.advertise_where.site / 97b322a08c1a / 3

<a id="canonical-c86ee1fbe48534a1931be608c257b28e80136b75ee4d0bdb582a35eb146a97f5"></a>

<a id="canonical-8b2a7f5a68bf63d330742f2688cf2651154cdf460a4a2d9b3edfe71b3d4eba8d"></a>

## ip property — site_virtual_sites.advertise_where.site / 97b322a08c1a / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-fec5e059bc8261415fd9ed2c9b217b77e44463577ae9c3f2e4a8b76f4b635510"></a>

<a id="canonical-c186e41b696ad52b125879898ee9c5f55185c390fca83a94cabd3c6c84055171"></a>

## network property — site_virtual_sites.advertise_where.site / 97b322a08c1a / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [site](resources--proxy--reference--group-004.md#canonical-0ad282064d4692b0e4369d19ce33db032987bcc7c7ff395bdc94ca7bef08f6fc): complete subsection reference.

<a id="canonical-100582468f659cc30dfadd6f1570736e870a1bcb1f69df40047a2fa81daa26de"></a>

## Next pages — site_virtual_sites.advertise_where.site / 97b322a08c1a / 6

- [site_virtual_sites.advertise_where.site.site](resources--proxy--reference--group-004.md#canonical-0ad282064d4692b0e4369d19ce33db032987bcc7c7ff395bdc94ca7bef08f6fc)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0ad282064d4692b0e4369d19ce33db032987bcc7c7ff395bdc94ca7bef08f6fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc6a342eb02b7437ea5dc971d04a351775fc378e872ce4df788c786730c00c97"></a>

## site_virtual_sites.advertise_where.site.site — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-004.md#canonical-948dc3edcc83805f304b7510f2dd275054200aab5bd1fa20497144d25b107707)
- site_virtual_sites.advertise_where.site.site

<a id="canonical-93e6f0d8a18186a4ddd871408da2e1942152ee62df065c8cdece33315b7191e7"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-862b73f0551d66d8b2042211d8e5a734c48e1275fcf9753a2508ab2a7d62166f"></a>

## Direct properties — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 3

<a id="canonical-d9d38728bdc52588685fa3990dcbe15a083d9fc126c94853ad41b57f2efca797"></a>

<a id="canonical-758441bbc2bef9a5d194b90fd815b19ae7cc2ce695450913dd9c96b16f453edc"></a>

## name property — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 4

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

<a id="canonical-44a50821a7077ed29926153de2f10702f52e99cbd50ce40dbc31fd345572136f"></a>

<a id="canonical-23c3ae2f9db70674fc5a86b82041775052a138e857911c1524986ce44138d211"></a>

## namespace property — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 5

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

<a id="canonical-df75a0c6723909da34cba46c6a9e8f3854b5bef6ac96d18d18de3f50960a4d0e"></a>

<a id="canonical-81092e43901e70813f432c97d92379d4372bc147ec1f3542bab307f9a8a95bbb"></a>

## tenant property — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 6

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

<a id="canonical-5928fd3186525ba521a900c652630dd620940958261ce0b1866a0644a06a22af"></a>

## Next pages — site_virtual_sites.advertise_where.site.site / a29d5c4cc59d / 7

- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-004.md#canonical-948dc3edcc83805f304b7510f2dd275054200aab5bd1fa20497144d25b107707)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-86401157f683146dca76cf73b8a7219a29a80e1b91a563f555f505e876a90e74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e25d982c19a31f9f06a424611001c65b9cda83b2947fa8a23a63513fd61b671c"></a>

## site_virtual_sites.advertise_where.use_default_port — site_virtual_sites.advertise_where.use_default_port / 859e2caa57b5 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- site_virtual_sites.advertise_where.use_default_port

<a id="canonical-1a6325d2b67e0303382e9e7565890c30eed1e770c388a69406d994f15396c174"></a>

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
use_default_port = {}
```

<a id="canonical-0b30bcd23216ac002f595a4fbbf368230e4132596a0f46c89993841eff78f109"></a>

## Direct properties — site_virtual_sites.advertise_where.use_default_port / 859e2caa57b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e305ccd2ae84fd6872c28d4c35ede7504e119276b36c3e57277a5d4a687e6d7"></a>

## Next pages — site_virtual_sites.advertise_where.use_default_port / 859e2caa57b5 / 4

- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-56cb8cc6f75ca94245ca58189fe0b09826f50a728714c3ad33c01761194c46d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-783e20202776e2578a7a7551c9feae5ff61804478b497d233246c394a5032b66"></a>

## site_virtual_sites.advertise_where.virtual_site — site_virtual_sites.advertise_where.virtual_site / 1b1f32d97505 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-09def8c1ed4434ecd861127657544cd2810ba7bcbcc7fdf746874bbdee930402)
- site_virtual_sites.advertise_where.virtual_site

<a id="canonical-2667ec23946d0986525e114ebffae9ceff274b3b064d723b7a26d63485ecbf4a"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0b578f15be923544e537abe8279b0f23bfc4957ce718b2fa0e54eb5b68b002ac"></a>

## Direct properties — site_virtual_sites.advertise_where.virtual_site / 1b1f32d97505 / 3

<a id="canonical-4bd8157050ec7704dec50e2e9127025580d4cbf3dab2919ccb0db878bcdf9104"></a>

<a id="canonical-797e67ef7a7c32db1d143716565fdcc73a3402a2ae3ff6167f77fcc0641d1060"></a>

## network property — site_virtual_sites.advertise_where.virtual_site / 1b1f32d97505 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [virtual_site](resources--proxy--reference--group-005.md#canonical-2446fec07bb39696ba8c35542910b9a8ada175babf7e63d3661ee4fbc8b9c7f9): complete subsection reference.
