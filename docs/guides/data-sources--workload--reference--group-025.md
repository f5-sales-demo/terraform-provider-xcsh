---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-871e05d79030ae6d0bc0a651afcf7ff45926552d3d5b5bdcd8e4f96258817462"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc1cc5e52067 / 5

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

<a id="canonical-56448d7e498604d3c6f40934ee89e6d4bae9f9fed6014d331c24e25c6d871f66"></a>

<a id="canonical-3973397db66bed24f97637ce095a463c3b9802c2a4a954926833f61f2d75e857"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc1cc5e52067 / 6

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

<a id="canonical-304c32a4ec61bbdf0ba4c0e441a29f40129f25ead4802542f6de33dbe39c81a2"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dc1cc5e52067 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-f5ef0c84d18b08fcdc5fc94e396bae60018d4485f5d70afd07cfc9a36a646036)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-c8091537571bec74634b9ff0b4547c2e693e737faa6946c6aeb3c10459e3d624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b470c49a98106aeb1fa1fdf86f44e599f268bcce60a2200e8d806f9ba0bfb1ae"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 30a7e9a16f24 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-fa6b2fc456b8a511fe9191f3231d37b4d96ee04c3cfbbd58e9e6b57ee4881cc7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-f5ef0c84d18b08fcdc5fc94e396bae60018d4485f5d70afd07cfc9a36a646036)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-f2f3c0bf523472086a2323a4b6600f67a928f63e22f98a4ee0867067b15b8922"></a>

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

<a id="canonical-1c233df704d447a1373c361757a1612a6fdaca65f917d742b26b40e3276771cf"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 30a7e9a16f24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-88ad9ae8aa906ca0cccf6d87adeeb5024d2448e91b82fd6024e7f20d7f085bb8"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 30a7e9a16f24 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-f5ef0c84d18b08fcdc5fc94e396bae60018d4485f5d70afd07cfc9a36a646036)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a6bc4423f41bf167bfb6cf7a76ae9e525dda25d9ff8999a83cf4d95d678abbec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e48ff3678afb6786de80c6638e84b0de7c3710732c5c5434b76e4b67b4e3fcf"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 859330590b87 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-fa6b2fc456b8a511fe9191f3231d37b4d96ee04c3cfbbd58e9e6b57ee4881cc7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-f5ef0c84d18b08fcdc5fc94e396bae60018d4485f5d70afd07cfc9a36a646036)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-234aa6b36493620df5174d153a8daf24f8a2eefc8b95919712f0b7593175d35d"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-aec140f426a71f3cc5287c0b52c990ebd238f7952d19e723c629a2c91fb2e9ba"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 859330590b87 / 3

<a id="canonical-3d6fed6c0dd507254835493e99a18957699b1d77c53f1991e371a18b61db0ebb"></a>

<a id="canonical-80f63834d0c82d2278220aabe073d9433923b208128016797ac69e171171eb9a"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 859330590b87 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-b0a14020996c5ad4138d1528f86657d52c3a6f4fa48c518da48e66c24dfd60cc"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 859330590b87 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-f5ef0c84d18b08fcdc5fc94e396bae60018d4485f5d70afd07cfc9a36a646036)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07c3c43a9a71c39a1ab20a1c99eec788871da8af73a1e3f23548d9881f304233"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 24c47769d301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-5d75b474e5a29591f362d870ffa8eb5872e4017fd969a4be648427606e343bc9"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-d8291fa754e012fd262e961772a37a24198855507930d631b5f42b83f7d61504"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 24c47769d301 / 3

- [no_mtls](data-sources--workload--reference--group-025.md#canonical-51847bc2d95a158b8642b15e5f6f5af8b80bad271a6f34a33eaed8e6e4c1fbf2): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed): complete subsection reference.

<a id="canonical-9607768792a1761484351a7277e975a86485d42db406d68ee6b56b80df659fb4"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 24c47769d301 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--reference--group-025.md#canonical-51847bc2d95a158b8642b15e5f6f5af8b80bad271a6f34a33eaed8e6e4c1fbf2)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-51847bc2d95a158b8642b15e5f6f5af8b80bad271a6f34a33eaed8e6e4c1fbf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f312e51f93806734326cd935624e9cd2e1cf62d79b2eeba4824a87a0198ecc9e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ffb17ce0ed06 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-c502db14c8418092c66f78a234d94f60ea2986ad77ae6c64c1e03bfd37a4918b"></a>

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

<a id="canonical-3305c02367f6f580b29de64b8b0d5d9531413af2d7b9dc8c09402fd3ea94563d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ffb17ce0ed06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53e0bbe334abe98daf0dc5a41313393b4556b128a355a47040e323faf801a05e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ffb17ce0ed06 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a471df6096210425244c1bbff278f087470c9494d789256f961280282c7bdff"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dd32712256b7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-347006b237acbc4717566ff9876be27c8c6da4d7078ff0f20c9dfaddfc06415f"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2880d62ea0513bcde7544ce10b5d53a159eb55db613c57e09d872cc81237af4b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dd32712256b7 / 3

<a id="canonical-6dbea8554b129f43d61136757aaa74bac5232cad833624954ec155301c6190d9"></a>

<a id="canonical-c98fe73d630a0778baf688b8b6366e03409836f5ec422f9bd15051d9c77a2cd2"></a>

## certificate_url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dd32712256b7 / 4

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

- [custom_hash_algorithms](data-sources--workload--reference--group-025.md#canonical-02497f72fb83c246ff98a07a742c4944fa3c9d5aa2fc81ffc22058fdaede5fb9): complete subsection reference.

<a id="canonical-b5b9026cdb2d32aa558f0b05bb2b7b1d3cbb1a540bcbd1f3eb910d406d9fc2a5"></a>

<a id="canonical-19741da53cede3fa0f7e34990a22a09dd8cfc9465c43427d8a093ecc73c63eb4"></a>

## description_spec property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dd32712256b7 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-025.md#canonical-11cc62da278ef06bf07f548ff1c8885ef4067adb30e3da771a7e39bab87acb11): complete subsection reference.

- [private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-025.md#canonical-2d3c7e92102b7ffec33851a9665ce0677d5f48ad2bc0a3486ac404602a043a15): complete subsection reference.

<a id="canonical-08de27519aaa7ef27661c0e27b083a4c0839aa4ba6b3f0ba212ea6d61c1a476a"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / dd32712256b7 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--workload--reference--group-025.md#canonical-02497f72fb83c246ff98a07a742c4944fa3c9d5aa2fc81ffc22058fdaede5fb9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--workload--reference--group-025.md#canonical-11cc62da278ef06bf07f548ff1c8885ef4067adb30e3da771a7e39bab87acb11)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--workload--reference--group-025.md#canonical-2d3c7e92102b7ffec33851a9665ce0677d5f48ad2bc0a3486ac404602a043a15)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-02497f72fb83c246ff98a07a742c4944fa3c9d5aa2fc81ffc22058fdaede5fb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eb7358512f5dcadc35a36a70d6cb7ebc4f1af302bfde1dc4f4408c9791c527a"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bfc8609ac636 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-161d801c02f2b87a874083b3d281073fd198307d5ecc11738b097ff5bba6282b"></a>

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

<a id="canonical-bae4b336755c99fb90b530f8d5eabc1ee1eaf28cf86c273879a59eb0d5e88b10"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bfc8609ac636 / 3

<a id="canonical-2759d40a4659ca3294fe6d5b7df4931bf21e08d66d6886138feb4f995ddb8550"></a>

<a id="canonical-0d1e18965ef4e9e4b873106571b1da737610d34d3ea08d5a0bab6c623158c5e9"></a>

## hash_algorithms property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bfc8609ac636 / 4

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

<a id="canonical-3febb07ec4c2a7a84569b69b37abff41f3fb5cf6aaa61a1041b1257692578cb1"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bfc8609ac636 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-11cc62da278ef06bf07f548ff1c8885ef4067adb30e3da771a7e39bab87acb11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e8abc2abac53d055636cb7a2bf3be990d4b97bb8e94da11a96c2f1c7fc42cd8"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0c5994bb9010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-e49a25c8789984113c1cd52e23322788cb577d82fd38cc2cd4de5827b7c082da"></a>

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

<a id="canonical-439bcf2c340076335e28cc524ee5192b5ff04540195b4d1c4232eb7cd5cf8f00"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0c5994bb9010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6eb0dfdbd933d33aa3c718554a77555a49c0fdfb1ae856eb1a9bba260df41741"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0c5994bb9010 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99c434aec2802d3f53bbd36027871649c7106cb784fc09911a3900a1c8368b87"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31b18c7c5225 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-c83594dab643abd0b0e6fe7af798cb49fabbf4b9af999f5e2fe62bfbd35c5bcc"></a>

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

<a id="canonical-7ff325f9b0ddbd72848de9e1849564c67197f4eeb004c38cb3303a74c0500a09"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31b18c7c5225 / 3

- [blindfold_secret_info](data-sources--workload--reference--group-025.md#canonical-b2e3d530d6b1cdee81bce41c07eb7535475c674ae8eb9b3ad1dc5944b06cf719): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-025.md#canonical-7954134f0cc0baad58222f168b693a37820e6f96cd9b7e022820d89b3f6f67e3): complete subsection reference.

<a id="canonical-ef2d35fe108fa3dd9600bd2231780f31ec0135bc1854d3eccf40c4fefe636fc8"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 31b18c7c5225 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--reference--group-025.md#canonical-b2e3d530d6b1cdee81bce41c07eb7535475c674ae8eb9b3ad1dc5944b06cf719)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--reference--group-025.md#canonical-7954134f0cc0baad58222f168b693a37820e6f96cd9b7e022820d89b3f6f67e3)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b2e3d530d6b1cdee81bce41c07eb7535475c674ae8eb9b3ad1dc5944b06cf719"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af2ef1c4cd7f85dfd7beeb9c037702b6c17f32a209a6974f0606b29bd9975933"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-599a13728c2baaee4d173d5c6b89aec650423e4712b2053f5d345a09b99e6405"></a>

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

<a id="canonical-aeac486fe00ca0e8a7b97f57cbeb8fdaa0943d52e624a37c2bdd23872131d759"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 3

<a id="canonical-cbb427cca759d9a8d8302ffe575c99f0ff2763685fc38b03f31f79e2514de7c0"></a>

<a id="canonical-8a67c9470762b90a9b8ba9a054ba68a0b80c39bffc243bdddd6e4e60823cc83b"></a>

## decryption_provider property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 4

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

<a id="canonical-18bb90ac08fbd8d3ae234bad4501d42de62d80402379c2e0fdf72b8366fa2670"></a>

<a id="canonical-3fdf475fdc59f3865d5b465381e1b582ff017d65dd1309fdeee820c7c3f0994d"></a>

## location property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 5

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

<a id="canonical-df2095ed76c66e49cee2803e05683061953943103285a85abc7a9c045e598e25"></a>

<a id="canonical-82886abe77ae6b3f0817e4cbf9f15a05a7f9faa0d3e70bde10849a8fe96a9ea8"></a>

## store_provider property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 6

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

<a id="canonical-0ec6d8df4dad3816f5137da71b181abdf428be325ad14da6bb67f8eaec297fab"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / f6fb3f6363a5 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7954134f0cc0baad58222f168b693a37820e6f96cd9b7e022820d89b3f6f67e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61985720f9a4f4a85d45d0cb4f782d1b9bb8bf8a41f5ac2b822b18edfbe28196"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / d9c63259fe05 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-be72d4e394947d609995087e243680f812d4ba4663678e84f6483af4762fa705"></a>

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

<a id="canonical-596b8664054d981c0f5f8de7566128b7bf82b2c1885d63a309bd11ff04c5500c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / d9c63259fe05 / 3

<a id="canonical-e1c7a2d2c17f02c1979ca60d30fbea9569e40396d78a2477541d4629112c1f71"></a>

<a id="canonical-e613194b6bb726a7b43fbc3d78aa8bfbb868ae5935acdca9d6d563afa7baae38"></a>

## provider_ref property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / d9c63259fe05 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2ab533bc2f616e77d4c04faa48a7bf8df516453e95f24c2c5bb70e1426c51f37"></a>

<a id="canonical-6c11303d7008c7e0ffe2c9c6830e3f1da98905074c18b68c058e0a5a652463ba"></a>

## url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / d9c63259fe05 / 5

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

<a id="canonical-641dfbfa25e00b42780e6a8e5b7f95e05fec6c52fd76552b756887f280f21cf9"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / d9c63259fe05 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-025.md#canonical-78380052c327012caa2501bba897c404eaa0af8c2f0bacbe9110faa426ae3611)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2d3c7e92102b7ffec33851a9665ce0677d5f48ad2bc0a3486ac404602a043a15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-147e9bf952f8e57ef05152db7e7e8bbfa4da2c549965fb843e4a9bbf83556a97"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c1378d2af989 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-dbb56384f7064564e277f8dfb7c92bb62923330fec3675c5c6e4ea78a035c2d6"></a>

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

<a id="canonical-158b05c458c2957a6826d804e5d5e2dee46b70888261b391aebc1b7dd6511ca9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c1378d2af989 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90da75134c8b2e8f318f27061fae3333cf0d840c69c725863ec075b6a5310509"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c1378d2af989 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-025.md#canonical-b1c7829dcce7ec189f5a5faa5f8b9db883ae54c650491caaaca0e2101359e05c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55285d6aefd63044a380c0aecfdaa13b62e8b7a7d11ff7ef002549b399ff155f"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c36b8c5e3cbc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-436d2f57cd7c780b8bc2b495c2919850b94c97cf7373bcdd9461f40400e31474"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-affc4a91c43c959279638f514897737280f0fbee4972c8a304d1501c5225d18d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c36b8c5e3cbc / 3

- [custom_security](data-sources--workload--reference--group-025.md#canonical-f6573bfc70d7f0b5fd2d7a6e4306cc64e1fe23bc3ff31199648f6400d94ccd20): complete subsection reference.

- [default_security](data-sources--workload--reference--group-025.md#canonical-998d872d2d813d609ef0fd38879ccb8c8b46abf4b851ab4cd0e82ddba1c5afdd): complete subsection reference.

- [low_security](data-sources--workload--reference--group-025.md#canonical-30c7b023b1857b182c11d128ff170c95a4b21176af0f86c752c91b7b9c2393fe): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-025.md#canonical-6c3038fadcdaca1b499a3d454eb2261ffeb5a9b99faa1d307842a7edd6e74dc2): complete subsection reference.

<a id="canonical-20caf9140c4a85050ef18cbf388a8be334f802b3ca4f751066c7a0dd21928fab"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c36b8c5e3cbc / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security](data-sources--workload--reference--group-025.md#canonical-f6573bfc70d7f0b5fd2d7a6e4306cc64e1fe23bc3ff31199648f6400d94ccd20)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security](data-sources--workload--reference--group-025.md#canonical-998d872d2d813d609ef0fd38879ccb8c8b46abf4b851ab4cd0e82ddba1c5afdd)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security](data-sources--workload--reference--group-025.md#canonical-30c7b023b1857b182c11d128ff170c95a4b21176af0f86c752c91b7b9c2393fe)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security](data-sources--workload--reference--group-025.md#canonical-6c3038fadcdaca1b499a3d454eb2261ffeb5a9b99faa1d307842a7edd6e74dc2)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f6573bfc70d7f0b5fd2d7a6e4306cc64e1fe23bc3ff31199648f6400d94ccd20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6374d9fd48b0a573d77cec1c305f3dfd6f82cc9a8fd747f502ff2931c3b02a4"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-12ad71d755e5fbdbbe4fdae32e735ee571baa14fd00135b4e1e28fb016dd8c52"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-235de56efc709435b8bb0d54e25d14881d0a55f7d49eccfb58192eb0fa8ffab5"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 3

<a id="canonical-e670260d60db17a54d76bc0b7d5d8790a064ffa6f65ed3e2923145cc0ffcf868"></a>

<a id="canonical-53e7f3fd5352d835c917f7816c97458385f7f195d9ab12c8d79f668beaf42c00"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1d8dc178d075893c8aab6e95c0364d0d658a1a0de4c9d14efa000992d2347a92"></a>

<a id="canonical-8ecfc533e5b57548598b9a8f5b4ee31663983c29ffae7e6898e9219d40a8d4e8"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 5

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

<a id="canonical-21135dc27a645c362a0a5e1935ba81340fd0adc13fe3ee075ea6b37f5e6a6032"></a>

<a id="canonical-77133a5733e713c4fd2511f7be8557208c756de84ee30e6f24380121d3a5ce12"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 6

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

<a id="canonical-4a961fa7d95835c0766b3e27f3153af3bb76135f9685fc9240024b6084883c1c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / da78a9379fe2 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-998d872d2d813d609ef0fd38879ccb8c8b46abf4b851ab4cd0e82ddba1c5afdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3606f30976b6f1c372d3acb98299873712090cfc67c5f261ed026dc6bc27c788"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c4a7aedded / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-14a5ae0ff3067e3c3ae418dd34ce5c8f320b765223ef72c78d3b54b839c1d9c5"></a>

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

<a id="canonical-878414c2270e4478c58b43dc0034b1743c79cb0a88b4a646837c9ab88c47e528"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c4a7aedded / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc2b2f7e6a254aefe342a4f9b47486a8c7bf9d9ddf0a63c9bcc5eb78283105e3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 45c4a7aedded / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-30c7b023b1857b182c11d128ff170c95a4b21176af0f86c752c91b7b9c2393fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e3e528b582ccf262b70dff5fb0c091946e5489f6a672512fea8332b86cdb7c2"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ae8043dee4b8 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-192db1cc1e149a822dc5abff34b8537dc25e724a9f19537e5ebeca23cee6c270"></a>

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

<a id="canonical-497ef7b83beb2fecaed2c25b52ce5839e188bff789431eb0d619e830a4145eb3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ae8043dee4b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2336d176303a6c3223f68d7f95bc58a0f95435506841f3964b3df7ed0f66666"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ae8043dee4b8 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6c3038fadcdaca1b499a3d454eb2261ffeb5a9b99faa1d307842a7edd6e74dc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c96053d61c3222cb2969362d7dbddcb02b9b1396ca8b0556bcd93b4ce4710928"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29f6ec311c84 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-fedf6b00f42b7e623dfee035cd74f61afab1b449a6f52ee7e0723985fe2c3c2b"></a>

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

<a id="canonical-f6317aed631f89871cdcdd670cbb9fcaf9accb1f99b703f92441e03703aa1623"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29f6ec311c84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42cc830ade4d53b119488bd99288d66a11c50bc816f148e2b9ca702616154b0e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 29f6ec311c84 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-025.md#canonical-025a98dcea0528288ea88679d37803d18e4a879528b54d18d7adda20aa085cec)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6531b582cf0d47928a44f170abd8373162071e89f6007dec3d128ab2d3212f9"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2de1754f1651 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-033cd8fee88c7d18ed8b643da9303f27e9381849cc3df09324df60e12fc87818"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-313a3f7037699757fd8fb4974340d2e2ed9469fa1031fdb39bb56f96d229f256"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2de1754f1651 / 3

<a id="canonical-c0968ff600997f9c7c957b554f8154a11cea68f66c5d7f16888808b06ee823bf"></a>

<a id="canonical-741270d2b09e36fc51e6dec4349c4d363d1c0e760bda4030a6be239a41575bde"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2de1754f1651 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--workload--reference--group-025.md#canonical-e7154f14a3b87d2a60462b70135f5837763bcb08396f6ff107b4f480c6fe92dc): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-025.md#canonical-0ee0a1887f5c3e3cb0483b2dc0d5f2084ed2ad9087b8fc70be06490130f4ac50): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-025.md#canonical-44a4051691b4dc6ae66735b7305614111b5d79729967d9860b3a36e61d6c0d02): complete subsection reference.

<a id="canonical-2276828e7097e3046342b2402a27ef2526fed682e77c8da4444b32cbddf1a3d7"></a>

<a id="canonical-010ffe9f921eb588dfc2a09ac53a4bc0d48e1094a72fec72b075cecd29545eb9"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2de1754f1651 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--workload--reference--group-025.md#canonical-3d834538721180604ea1ac2773980baae3443a64e657e7a04b61af7163f3719e): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-025.md#canonical-7aa068b0be0696ccb92e036000d54d7b2e38012d4a3ae0cc4158fa447396f990): complete subsection reference.

<a id="canonical-14ddba6be14b9e61adac60596fafaed5024ebfa82f0a38ef42a243f9fc72ce3b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2de1754f1651 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-025.md#canonical-e7154f14a3b87d2a60462b70135f5837763bcb08396f6ff107b4f480c6fe92dc)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-025.md#canonical-0ee0a1887f5c3e3cb0483b2dc0d5f2084ed2ad9087b8fc70be06490130f4ac50)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-025.md#canonical-44a4051691b4dc6ae66735b7305614111b5d79729967d9860b3a36e61d6c0d02)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-025.md#canonical-3d834538721180604ea1ac2773980baae3443a64e657e7a04b61af7163f3719e)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-025.md#canonical-7aa068b0be0696ccb92e036000d54d7b2e38012d4a3ae0cc4158fa447396f990)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e7154f14a3b87d2a60462b70135f5837763bcb08396f6ff107b4f480c6fe92dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d85eeda6478b34564cfa3d71784e276dc14727ba4b7dbdf06101b4372ca11081"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-d35e62c1d7bf39dded36b3fc7c68fb7f9fbe616ff18fd5be1688a07d7b1ff59a"></a>

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

<a id="canonical-68896ecdd5d45f486100433875d1f839954b1e07076c67eb547d21a55cb46d2d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 3

<a id="canonical-21b990f524db5fd5b40dc7239c81af9c8fff4637de4d17dcb5a60e4a43bd224d"></a>

<a id="canonical-982a88b6cff9bbaf6b3f8bfdfb69892dab5d3fc329d2cd59afae3d6e1cdac328"></a>

## name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 4

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

<a id="canonical-8ebf039b821df614651b7fdd9b023dfcba18e847e84698d212edc5c2a43230d4"></a>

<a id="canonical-72d128b9e20bb781e9e57d9dd9d1108afc7193883303a8746450e4d0fde60313"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 5

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

<a id="canonical-27f7b720aafd5d2a055c8a474c4f6c5fc40a882bd0057bdccf6b4651922ff84d"></a>

<a id="canonical-d2f4286d6b59045311c1d856a405ad75620b0aefede0df558b2f7be40baa24b2"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 6

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

<a id="canonical-fb2a5b9d7f6aa352cf5cd82a601f59e079ee2243b8f3ec08cf270e020f1888e2"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 9b4d39e34b67 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0ee0a1887f5c3e3cb0483b2dc0d5f2084ed2ad9087b8fc70be06490130f4ac50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a520572f55081a258b48aa03beb61fc275b0252551dbd5835be2b396cc4dd95b"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3daf2bf08291 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-e1e0a6d66a2952c969f8ceb75c210b7765c4fd486c44cdf4cb3753bb86992a92"></a>

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

<a id="canonical-f5b2561147a337757ca87e56fc63409381126a8c3d1a6fd15180a0f9d0587b81"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3daf2bf08291 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91d979e6cda782adc449da48f50a46f380918969b8fb1d5a4104d398d7054a2c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3daf2bf08291 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-44a4051691b4dc6ae66735b7305614111b5d79729967d9860b3a36e61d6c0d02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84ad4030f5bcb37e363844ac9ebcaddbf289d16c315fc9fac834c2e2d95e14ca"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-330ca873e81e8c192acee2df118dfd705a8665839e6a9180fd012377f423008c"></a>

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

<a id="canonical-b0be8e2248b368348f3c219f5cd636c90f68b1d464e3ac1c53d5b930c453a216"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 3

<a id="canonical-397efbd12e329c4f61a50fb599ad910bc84bb19bcceb20ce9d641aeb453d8f01"></a>

<a id="canonical-fe9c19fb7925e4f90bced7ddca193f9976bc3a94c452a15901bd51b5a59fee7e"></a>

## name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 4

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

<a id="canonical-a0e15994a77bcb78552cefe80c0958b81c5e805963bdccfea038383a8349e29f"></a>

<a id="canonical-dcec01ab6b1329a5403a49e9d7f156d919ce2391fddfb5af4ccb202315bbb77a"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 5

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

<a id="canonical-a7ca7b8d0356b08e6d6374df9bcef42171e4ddaed5c15e3e7508c5bab32cc29a"></a>

<a id="canonical-a6e3d678dfe1992bbeeb833c28a537ebfea2901cfac811aa2a4da3da508b9d3e"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 6

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

<a id="canonical-1870d9306f08612ab4f1a8462bc02afd13cda181e32ca64e80c83a6b1a78ac89"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 011ad4d6ae14 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3d834538721180604ea1ac2773980baae3443a64e657e7a04b61af7163f3719e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3267794593a43da4bb6793f2400992d66d3e7d267b4f74687c254dbf1aa995ef"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bf2dece85955 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-54d735ee8b674574a60fe51759579f85e29ee64b802144fc2b85f49e21c53905"></a>

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

<a id="canonical-87526a32f59e6356d51ea9a66ef09e3fa996ea9fe80bcdffd1733f3c2802ff31"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bf2dece85955 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9f84ba1092c728fff2aa57073412f383a8b0dd7bd710277487dda0fb6a06887"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / bf2dece85955 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7aa068b0be0696ccb92e036000d54d7b2e38012d4a3ae0cc4158fa447396f990"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d846118c04b45f31c5edffa848987f4efdf21ff902a0983c59857e0491c047e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c685fb05daca / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-dc8957b03198c3e5f6ce88fc4279971dfaf07884c7e06c624d85397fe1cf7d75)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-67e73671527ce3192955bb6d9c00b12333900aff5023b57c18a5e3b548ded95f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-bd47e0ae8672889271e9351c1359080c1cf68c3b557d5d1ddffdbc77f8790864"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6bbe630dd33be6c415263b91eb6106b789e9cceb2485d2043dec85aed9260be3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c685fb05daca / 3

<a id="canonical-dabd1fd841367d8f53fee52bd3d47963cc2b0347efa5d386d6fec98c0efc0dde"></a>

<a id="canonical-a929e62a897150da1fa729d72302768e3229fa3f00ea4e2f18220fde59e8280f"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c685fb05daca / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-91f565cdd4cef9b8f5403a88a9d10de8c14f25eb7202bbb3ed3337228d092380"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c685fb05daca / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-025.md#canonical-2db3d276ff46c9e73418807bb833046bc1d1d5d2203515f9fce167acb7c046ed)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e447258cc408ff08ece4248a75f978acd7100e0228c23ac56fa04da742960e32"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-93cb8afceb3568c78e38b150cdf8f4ca311bc110b88951e5f9b08a4f4aae79c6"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

<a id="canonical-b0d0a299d54353f9ec88a440c0a68b249164a0f7f449c7372acc3e554b240718"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 3

<a id="canonical-ff956f64ce7607c78f22f26e8463d076e0e84739476d30381352e68b5eecc5d6"></a>

<a id="canonical-f5f3815e353bd5cce3795c34a41bd63b2d471c00ddbac173d286263e20f64b94"></a>

## add_hsts property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d524651f1a26c023e08cdabe89d782cbec9cdc396686a7edfaf36fb049664eb4"></a>

<a id="canonical-6147eb62bb2a940669f6dcbf9be472f5b4bf79400d41ca89442d6fd76d729d97"></a>

## append_server_name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1): complete subsection reference.

<a id="canonical-54acbf7bc786c305c1d0eb2a3423ecba89c6a5ae97ed754ebc6083115f7ae708"></a>

<a id="canonical-c39e37ffbc3db636add9edebc67bfec1e212965d44c1eb99394e5971d116d78a"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 6

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

- [default_header](data-sources--workload--reference--group-025.md#canonical-95a04bd6b49bf1aaeb9f8f4021c1a446c41d8645b28ffcd99b8a83569d5a39bd): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-e0b20041727b3d717441031ba62b14f1d37bc29bbda616b0b836602e5e174c7f): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-025.md#canonical-159ecbabb87c21aa5c8126d9efcdc7dba2fc84eb764c4f0b38f0f8e66a16b99b): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-025.md#canonical-7fb6391a7ac85374bbf826bd9faef09d0d855240ac749a9300248c9f9ffd2624): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7): complete subsection reference.

<a id="canonical-47fae51b64a8572c5b08a06875e78e9dc7bdc32d040a8e6930f249efcd74397b"></a>

<a id="canonical-76e39b198dc7417f2f0ce37fd61b88c88eb41c4e06e89e1ac6b677bfd4585196"></a>

## http_redirect property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](data-sources--workload--reference--group-026.md#canonical-cdfae73757c0fcf2e28e977a401a5d962d6e1e0b377d369f21b508e9037ad5e5): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-026.md#canonical-3e25122adcb2c2bba3c8490ad9e74107d7f2d11e462a945e180225f12cd2aa6c): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-026.md#canonical-bff6faaf24054b5ebfc4b8ea8d74735cb7ee94bf4291a1f9e7c377ba1bbc33be): complete subsection reference.

<a id="canonical-c5191a75d89877cb583d86b75d26a9e9fde039fb40955d91b839813fa192a746"></a>

<a id="canonical-53cbd7e7b26f077e238ad2a44b9fa4ce7644829ee09650d20f0a4d9033d55a8f"></a>

## port property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3e4c6209276cf7b77b93d3ee90ee989cf31017cecd8208ae95e8f1c1333e4f2b"></a>

<a id="canonical-1371eeefd32a238b15c129e0e049f177eece7d83461f3c2d2c80b5a627c60122"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-26e8abf22a4899e011245845b34a94ff0e0fd22fb031fc4c889c2a9b48d4d520"></a>

<a id="canonical-15c3ab1ebeeb3741c678f8a575113813d4ea6ac5c27600ab03b2faf9694b3cec"></a>

## server_name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](data-sources--workload--reference--group-026.md#canonical-4865dfd6b4f6cea302d2b8f8f5276d01c2f42efd0cad3524d7a483f5ed6d0528): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-026.md#canonical-a8494ebff457818ef6376d45f5b4881d6d741f68987d97540e9c408020f95cee): complete subsection reference.

<a id="canonical-706f3c7b202d0c55219e536421b8a67af41e8ccc1e35240d5321df9bd78b2a00"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 6a39a2427907 / 11

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-025.md#canonical-95a04bd6b49bf1aaeb9f8f4021c1a446c41d8645b28ffcd99b8a83569d5a39bd)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-e0b20041727b3d717441031ba62b14f1d37bc29bbda616b0b836602e5e174c7f)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-025.md#canonical-159ecbabb87c21aa5c8126d9efcdc7dba2fc84eb764c4f0b38f0f8e66a16b99b)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-025.md#canonical-7fb6391a7ac85374bbf826bd9faef09d0d855240ac749a9300248c9f9ffd2624)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-026.md#canonical-cdfae73757c0fcf2e28e977a401a5d962d6e1e0b377d369f21b508e9037ad5e5)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-026.md#canonical-3e25122adcb2c2bba3c8490ad9e74107d7f2d11e462a945e180225f12cd2aa6c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-026.md#canonical-bff6faaf24054b5ebfc4b8ea8d74735cb7ee94bf4291a1f9e7c377ba1bbc33be)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-026.md#canonical-4865dfd6b4f6cea302d2b8f8f5276d01c2f42efd0cad3524d7a483f5ed6d0528)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-026.md#canonical-a8494ebff457818ef6376d45f5b4881d6d741f68987d97540e9c408020f95cee)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13e4b14b588ecd196ce077a2930cf9893e5fa7a4e8fce30a2e45b1406ee91670"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2d6ed1786ab / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-2447e1a0080d8a10f198858dfe589957edce9cd0abb286b992c2870a73b92ab2"></a>

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

<a id="canonical-f84dd3125271e197adbe0bb6a17144a56e839c81810795f66a2d5959515a1f0e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2d6ed1786ab / 3

- [default_coalescing](data-sources--workload--reference--group-025.md#canonical-0becd624a3434e666947f751999ff96b50500a0d7d6d8a1e7e68e4c949543408): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-025.md#canonical-dec923c10b729b90f0442f928b67c1bc617f78bf4b6c1f8d72b09b8e85012b25): complete subsection reference.

<a id="canonical-bfec29d000ff166c4fa9126f26375ebfa6a08cef1cb4bb36a67f87591d2d2de3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2d6ed1786ab / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-025.md#canonical-0becd624a3434e666947f751999ff96b50500a0d7d6d8a1e7e68e4c949543408)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-025.md#canonical-dec923c10b729b90f0442f928b67c1bc617f78bf4b6c1f8d72b09b8e85012b25)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0becd624a3434e666947f751999ff96b50500a0d7d6d8a1e7e68e4c949543408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-364dd0e3a306f47466c925b09545692ee20c3514fccaebf505f14034c0e00ca8"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / fda7d907d25e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-543e7897b528a61ae8da2330f49a36ea43c5afe3b2b2fc592dd97fd8b2955087"></a>

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

<a id="canonical-b6fbdec37f8a4b87c1cc24c00d92f37784ec67a28dda0e0e509c1052a0873d12"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / fda7d907d25e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-326342817f8b171a90a30cce09b8276fcbee4301f825375a84a9e9dfbdaecd75"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / fda7d907d25e / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-dec923c10b729b90f0442f928b67c1bc617f78bf4b6c1f8d72b09b8e85012b25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f0fbcde307400ecfbcdbb9c36435d0b517fbe12e288fb3bbecc2dbf2e30e89e"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 840189e19779 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-00eebba6dc1d44d280275f7371ce088e114ac32f8b703a012a8723330f8ba5b1"></a>

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

<a id="canonical-8568eee3efecb9db7c1c820c8f0f6d647bc6993ce0c3aec085ee62c6764fff75"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 840189e19779 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-95951402ddb389d94a9225cd396a6a9473c2c3dc2207efb97d3384ff1fcfb038"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 840189e19779 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-9fd59e572a4daaff9a19c2ae1c0f7dbfe8ab370c444ac3c473ed1b97d223d7a1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-95a04bd6b49bf1aaeb9f8f4021c1a446c41d8645b28ffcd99b8a83569d5a39bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d81c8595530b07d91e70fb26dde965477a9801a99d110057395214b034d4422"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3c458f15f487 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-50ffaf6e58f9e4bfb1f7edce6e27ab721484015a7547f6556658e47f5bdba586"></a>

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

<a id="canonical-86891c1ae5e8bee15016176e9819a149cd83e446a49154f63390ac17cc756f7d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3c458f15f487 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bde407d1c7034b05259e5f36558ac576df5de573aebf8a002ee1cfd49a2102d3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 3c458f15f487 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e0b20041727b3d717441031ba62b14f1d37bc29bbda616b0b836602e5e174c7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ae1ab8694c65f6507f87fb80ae2854ebdb3309510eac5c89487759dc4a2d6bc"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 76b670c7781d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-c90d05c3ea8ed0a7363fc784df591c994db239aff54e84c5d8965ab3f55f8d7e"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8907d1b3a21bf6b677189ca2cba3bf5086ddfd7f2c02a789f878a1f6b917abed"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 76b670c7781d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d492e20e642aaa72a6e3ad5b29dab02ffb58e29895a871626d1f8f7939d87e8c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 76b670c7781d / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-159ecbabb87c21aa5c8126d9efcdc7dba2fc84eb764c4f0b38f0f8e66a16b99b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa06b32769b99f956d0f54ed44b6df257d1c1a17875e516c728da5fb831d1de6"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / b6de5c46d038 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-78367c246b986f82ea209f77816bb5c35eac8c847b0eaa4bd75f681ff32012a6"></a>

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

<a id="canonical-f2abc67dac426be5a8cfb172fef25f412901043044989863261036ade108aa0c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / b6de5c46d038 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35a0375cd6432e169ecb77dde2fbc82906f017134a0d30ee0820847df8ace5de"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / b6de5c46d038 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7fb6391a7ac85374bbf826bd9faef09d0d855240ac749a9300248c9f9ffd2624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c67aae5dcc707f862951b9b17bed42c86699a336604104cac3734749129eced2"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0d8169d041a9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-3a4d8a685e54f4d7141d9ac26dbc74d32823ebc514063b319487407c4f4e68ed"></a>

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

<a id="canonical-6c599083cb4c0d398ebe3d878e285e96f230f2b02753c61e1e2f8ddda9abfc75"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0d8169d041a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21c174e3c6028eb992453fa496d8d92246c813090fe098bffd675642adfc3a1a"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 0d8169d041a9 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d342306dd0e64485739523ba9490df17b7bdec2c85722a0bf31662b4e037b174"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8161bcc9651c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-555e1574f4b177073a5514245f1461e698c655a047a75ae968c2e514fb66d805"></a>

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

<a id="canonical-b21671b6fb8ae60611fc428c789f253fbd3c299d1beedf833b581535c27976e1"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8161bcc9651c / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-026.md#canonical-ecb2b4e1fc149f237f76193cabc3e645b1de6fcf4fb5bde6c12fd0614fbd780e): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-026.md#canonical-b6520c6963fb3d7a2804b9ba061567110544af0d1ce531a515af2539c9e3a454): complete subsection reference.

<a id="canonical-fc910759b0999898ebddbecf333436634ad13175d75ef68f1561fdd0b0844034"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8161bcc9651c / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-026.md#canonical-ecb2b4e1fc149f237f76193cabc3e645b1de6fcf4fb5bde6c12fd0614fbd780e)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-026.md#canonical-b6520c6963fb3d7a2804b9ba061567110544af0d1ce531a515af2539c9e3a454)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1930b9f188401add079d7fc6e1993c2c05b007bb853f682622607e144dae43c0"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2660ad4d8da / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-ef0a6807457f29e17c198dee70866889b9676abd90be8257ee26b80f3a2e017b"></a>

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

<a id="canonical-72b64986a4a75a3facf1246ff081442de4c4c877f2538aadfa3a5e06798d8200"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2660ad4d8da / 3

- [header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c): complete subsection reference.

<a id="canonical-d2f57bd158b2a150229627e945a4945aa08361504d6e1d6b4bbdf7f8787bb792"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c2660ad4d8da / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326538713c9c21d1ad2a58a901fa932629e0e744634e81e75ae5503dfee44425"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / adbf56b9d63e / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2fb1d2431e9986cff9df3d850320dede45eea34c1330455c2d9d1b4a37bcd2a8"></a>

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

<a id="canonical-3f8045717c814a7b94d8af9e8c1d14c8afb12fd4708dd12b8fab37b781f704c7"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / adbf56b9d63e / 3

- [default_header_transformation](data-sources--workload--reference--group-025.md#canonical-9505ef340014876e57abf1fb15c1ca61b186e69ad3bcbe04275237aaf020a46e): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-fe70fcf97d7f9efdf04fe9c43cc97011d6c698bb9914296a01371b0432a37182): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-942f57b3b66ca3e9ecaad7acaf930f7050e1d4c0329e6ed153890737959afe10): complete subsection reference.

<a id="canonical-81f6401e083ca2f26a7f386e60b641b2efa231d4d67f7da4b327f8e422081622"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / adbf56b9d63e / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-025.md#canonical-9505ef340014876e57abf1fb15c1ca61b186e69ad3bcbe04275237aaf020a46e)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-fe70fcf97d7f9efdf04fe9c43cc97011d6c698bb9914296a01371b0432a37182)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-942f57b3b66ca3e9ecaad7acaf930f7050e1d4c0329e6ed153890737959afe10)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9505ef340014876e57abf1fb15c1ca61b186e69ad3bcbe04275237aaf020a46e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd3ee086546d0343a2c8e5685f383c5ea03eb4f5273a7907358ccb15242a0343"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2b88908ce219 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0d80c2f4ad61a7ea5974ed4be8516682a9edd22c0d354ea9108019dda3bdd150"></a>

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

<a id="canonical-6516bb62187951dcecd5eeb6a5388e977a55a9c382ff53f1b25dcfcc2fdd81aa"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2b88908ce219 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c77ef721e77cfa7c54c7549f271a33fc543bd147bd76839e6331b6572917ee97"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 2b88908ce219 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-fe70fcf97d7f9efdf04fe9c43cc97011d6c698bb9914296a01371b0432a37182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9b8ee85b08701a8e19e3131af50d27f81a191803d80823ce2867c4fcf08f708"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 59f34f9269dc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-857b34aacbc6367caf1f8bf01657730653de332a42a54e42d42b8445bff1df33)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-1aa566d80b1190fccf9045fc1043f09025ab2cd4e225c03acb6fcb4e8ffbe3a9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-a25d8f2b136582f1feaaafd5965e36173d34041a94ff027affb2c471253a6620)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-1cf533a36c897fb5bc5ce6b994300ea59f47ae0399bf64c0813ad7cf175525e7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-a4fb2196c48af6373c983e1a64aa1eccd11ee89dacfcedf9fd43ec756ec72d58)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-c2130a67a0256ea741dc17b7d6ceaf7142036519e1de100a3e898786f6716527"></a>

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

<a id="canonical-8530bfcade1ce62f05f7c5fadc1e2ed5786202a6c4681f2e3585fb3b1f9f2cda"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 59f34f9269dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6ca944ef196833e4927d51de18f46d2b1bb9bc9e2fd0a2ae2ff43167f943fb6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 59f34f9269dc / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-e5ef4cb048281bad0d2b2b7a33627a34c08fe1bffe16d3f55d20274e393e227c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-942f57b3b66ca3e9ecaad7acaf930f7050e1d4c0329e6ed153890737959afe10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
