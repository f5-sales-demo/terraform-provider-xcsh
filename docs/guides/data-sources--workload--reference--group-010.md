---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-6b388ce63450ab91f00cb69c81abd95bb4a7946f41582f8d7bf3bb40d89f5423"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 3f4de622121f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0fdd3e9623f347e1598c26ec769cfc0e95538a2145586a2e882ba7fe0771576c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-009.md#canonical-7240fc29921955794bc4be7014180c6c3dfcf165aea1332be809b30560c16f0d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-b98b2f2bec28d3ecf40af67a0be517b24035ede894d527492509379b8477a0d5"></a>

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

<a id="canonical-62bd5c91649873224467e9ad27e5fe7a8c6b432148923e96d4e798566f5e3019"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 3f4de622121f / 3

<a id="canonical-9f3d04d2967e9c4e25a370e5bc2a58462b94f353f9cc597322036f973b87975e"></a>

<a id="canonical-42c8812ce1feab846d62bdf2301c2ae54909f518fc23bcffa7f643246c5023d3"></a>

## provider_ref property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 3f4de622121f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a4d32ab5dd00ef8dad96ece15a66221f3f6c9f29d2409f47ae4608e58e269b8c"></a>

<a id="canonical-3679b66bead0cc6adf990746b050f0637559a6d3953add4146c8093b5d074c84"></a>

## url property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 3f4de622121f / 5

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

<a id="canonical-ef02027d1d04685517cd347cc62eba36c5c72fa9cdb750ed9659c89988ed0d91"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 3f4de622121f / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-009.md#canonical-7240fc29921955794bc4be7014180c6c3dfcf165aea1332be809b30560c16f0d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9a0a4be0574eaea185d202f2f0965fb33fc1383ae6b953a93d9e72510539b646"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e9e7ad9432c075de294c0004e3b19b673f84c44b2648532448c2ce36b12496f"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70732dac23b6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0fdd3e9623f347e1598c26ec769cfc0e95538a2145586a2e882ba7fe0771576c)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-016d0af669f875473583641d7b702f72ceb8d095f613970dcd48db8780f92453"></a>

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

<a id="canonical-c7ac2b7ccece3cc4194bd685d549b2d18ecd14ccb32806c0442a643a70c69c98"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70732dac23b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5084586f33fda6f180fcd5f36fd121d55812768a7be8f2eda148f56312f1eb1"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70732dac23b6 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0fdd3e9623f347e1598c26ec769cfc0e95538a2145586a2e882ba7fe0771576c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1bf2f98d8500ab394fe2e0801b1c8ca306b1ac918cac6b9eb3bd65d60ec5f2e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a0f8e46d62df / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-51a409637280a1250b72a19bc0b217744815d0498523987000064142613f65d8"></a>

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

<a id="canonical-34ebcd870e2cc694ddbf414c8ad5843cdf3bde58d82d054e3763df554f7e2a05"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a0f8e46d62df / 3

- [custom_security](data-sources--workload--reference--group-010.md#canonical-0953ad506ff44f23f9e4c2cfa8cf6a468153e21bb5bdcf39ce8748745c5b7804): complete subsection reference.

- [default_security](data-sources--workload--reference--group-010.md#canonical-2ad9ccba9695a9bad7029012ee1bfe29d46e8bc1b3e55509b0ae0a286d9181d6): complete subsection reference.

- [low_security](data-sources--workload--reference--group-010.md#canonical-b65cdfed655e17e8aac7e10847bea70ccf700d0fc531502f0879384d0cbd1020): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-010.md#canonical-d6a469c365c6f73727363aa02384fcc8121eb3e7930b450188b0b705ee1a4394): complete subsection reference.

<a id="canonical-f5010256ab161131f6d99d67d382660b2f5b082e217d0235ad394c1e905fd872"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a0f8e46d62df / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](data-sources--workload--reference--group-010.md#canonical-0953ad506ff44f23f9e4c2cfa8cf6a468153e21bb5bdcf39ce8748745c5b7804)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](data-sources--workload--reference--group-010.md#canonical-2ad9ccba9695a9bad7029012ee1bfe29d46e8bc1b3e55509b0ae0a286d9181d6)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](data-sources--workload--reference--group-010.md#canonical-b65cdfed655e17e8aac7e10847bea70ccf700d0fc531502f0879384d0cbd1020)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](data-sources--workload--reference--group-010.md#canonical-d6a469c365c6f73727363aa02384fcc8121eb3e7930b450188b0b705ee1a4394)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0953ad506ff44f23f9e4c2cfa8cf6a468153e21bb5bdcf39ce8748745c5b7804"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-104d0a0290c59f6d206704800d668e28c9b5a7276a733f5fed922c504ec43396"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-62d9e54ce17e8dfdf304a1bee46fb8daa2e61cf020fad4983218bfcc4eccb794"></a>

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

<a id="canonical-a0c28bdd5e70c6e741e08481cc8f32f62283e871f173c21e06a182eaa09ac0a7"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 3

<a id="canonical-c78e7c23bb8eb86ffbc82717145405bfd14ccba10b2639bc03193b46213c5d8c"></a>

<a id="canonical-f63cc2d317ccb86c9f85ee24fc3125c90a36aadc52f35af852e889d431b6441b"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 4

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

<a id="canonical-39e9e795752b9ed1120b0c9c2ff4afeeed9fc9652c4ba7d6916e9c5e540d77e2"></a>

<a id="canonical-c681b99fdb3e2f46220cbbfa257083c4ba96c4188baa072c128cc7dddf4300c7"></a>

## max_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 5

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

<a id="canonical-bbdde12ba9de793b516fdb3d19a52a596e0b2abaa68b18f134ab70b0b1be0410"></a>

<a id="canonical-1a605f6a6aa912c4814e66780250b5d9da7774ccbe7e36b64f52534e0dd60c07"></a>

## min_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 6

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

<a id="canonical-b1943f1bf020e20596064d1a09bfbcebc6415e936ac9b9f50f0bf545fcab2f14"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 504104922e55 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2ad9ccba9695a9bad7029012ee1bfe29d46e8bc1b3e55509b0ae0a286d9181d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34419d4535905e1f865a2668061a5b97568741e20afe9e54e4567c98fca4d0ff"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 63915b86a9e1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-eea7e1d64c2af0ed3c764f3d80eca05e4639cfd500f0402d915f4b2f0849da10"></a>

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

<a id="canonical-bbdb2ef46927eab0a147dfbb39b04d3bb67c79b498f43e667f838da7040fdf57"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 63915b86a9e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a0e9090d8d816e04590609a2ebc72a8f8db9d5cf47322693acca252e13784c9"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 63915b86a9e1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b65cdfed655e17e8aac7e10847bea70ccf700d0fc531502f0879384d0cbd1020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3f050b9823a3fbafd599383ac7cc3b6107daac2466a6137039e2aab0f0a2db4"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c593a684700c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-641ece18d6bfea2df8c4ee7bfb30b8e5bc1d87b581496b6118c93c177aefd3fc"></a>

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

<a id="canonical-44f3e877b95f96905b6db5eaca2d03ed627951d9c4d57ce603891314b9547d60"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c593a684700c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b350ab76860d77bdf4b92d5b1a8bcd35372ae61c37fb8453adc622cd45987b4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c593a684700c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d6a469c365c6f73727363aa02384fcc8121eb3e7930b450188b0b705ee1a4394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-246bdc7a537c9a44a57be723527edbebada59ce17fc9b8ab2a569cf74022593b"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edd1c94e14fe / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-fcf666504e49309e14470cc4d8942a1a1ddceb38b91597cdece2f4e59c7a69f1"></a>

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

<a id="canonical-43a94309e91e5a9c8b0d45f414f1cc4d1eaf4b7c45311fee88b9c732dbc3c0b0"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edd1c94e14fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a6f2c5610655f9b0125b40675270f4697155172784816327dc9a337e38e813b"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edd1c94e14fe / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-89601e4dbe1c4b703845350ac22dd3aabd069b27ed797e7d4251372177cabd02)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4522146f0e4576767da22a07be49b01e4630c9249d54ec1b3940bfde60855367"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f662739b29f1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-909ec27300bf6616c2e5a8bf9d3b2f21dc613a6b28e398cb8449a29fa6337b17"></a>

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

<a id="canonical-598297152e7fb8d8f60d9fe5e38bdb544871bdf4e5cec3b72d224dcb91aca341"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f662739b29f1 / 3

<a id="canonical-349d6c6e240ba209f4880c3518ba86f378de366bb8a0af53ff4a70c7b60f0219"></a>

<a id="canonical-791bdcd187ccdd399ea3788102bb0d80b5bea90d6141549d900395a434cc7c49"></a>

## client_certificate_optional property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f662739b29f1 / 4

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

- [crl](data-sources--workload--reference--group-010.md#canonical-84b568e4b9e6c36c4d8906e3f0f47568d569d025a182eaf6a671db4f8276c20a): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-010.md#canonical-76374effb8b2a2a7a25bcd402d026744250c3576650be1452245c50ca511587f): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-010.md#canonical-18bda0e176281a9db3ee71eb68fa96494ece1be5e007c4fc9a660a95527f5f74): complete subsection reference.

<a id="canonical-05853899c682d3cd959a6cb751afa6eebed1577425afde3377bb1339830cbbed"></a>

<a id="canonical-492742f10c6eeb3edbfd188f798cf0fd1d5ee0ee9cbc5800fc10ffc261706336"></a>

## trusted_ca_url property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f662739b29f1 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-010.md#canonical-a93062b4da6a42ee040774566c70a6f904195b61698217e1854c30a0691f67ed): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-010.md#canonical-d6799d198efea25ca80761689e40ec104aa4f3cab05330e94733073c8935a45f): complete subsection reference.

<a id="canonical-3eb9481dc7df3ea1722b8b5577e819e69fda4a82ab87c4c9b337eda44cf08b8b"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f662739b29f1 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-010.md#canonical-84b568e4b9e6c36c4d8906e3f0f47568d569d025a182eaf6a671db4f8276c20a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-010.md#canonical-76374effb8b2a2a7a25bcd402d026744250c3576650be1452245c50ca511587f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-010.md#canonical-18bda0e176281a9db3ee71eb68fa96494ece1be5e007c4fc9a660a95527f5f74)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-010.md#canonical-a93062b4da6a42ee040774566c70a6f904195b61698217e1854c30a0691f67ed)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-010.md#canonical-d6799d198efea25ca80761689e40ec104aa4f3cab05330e94733073c8935a45f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-84b568e4b9e6c36c4d8906e3f0f47568d569d025a182eaf6a671db4f8276c20a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f8e3d9255fab6b4818bf4e68a8f12694825c3e1ba08b48f4d4933c13ac93c19"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-cbffa27353c527e3f8aeb36e0ec8056e12a5219b86639fd662383a43579ec620"></a>

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

<a id="canonical-6cf96425062eae0a7eace6e28c37dd66a5f8b39b4ddc7926cd24068804a69ad2"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 3

<a id="canonical-37e1d08dff5cca81080c3f9e219b2a2e1290479474a8a9e6d1c8ef4260163d5a"></a>

<a id="canonical-e796064893b4a4f2546cfccc43858d195ff96fe8c706610beb3d06a8ad39e607"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 4

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

<a id="canonical-174cb80b6e5155c888eca9b59d691d10fffc3505797d107d50d9f889358bca15"></a>

<a id="canonical-d253398774d16bdf18b2f6b7cc095ebcea7288bf10965f9ed200d4336de574b5"></a>

## namespace property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 5

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

<a id="canonical-0a2d7e03ecd8e6cbf5fa5468595a7a07bb81449600921a12f300d9d422c622af"></a>

<a id="canonical-25f6f3ddcf3372326ab074d6ce5a87d1d13fd9b1c8a6f0ec9fb03dfa65e25bf1"></a>

## tenant property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 6

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

<a id="canonical-7112da58ff981ac6f70dea57ca37d93c41cc36831df9abf1d4d786b9c446c5e7"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6845b12e13f / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-76374effb8b2a2a7a25bcd402d026744250c3576650be1452245c50ca511587f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bfcce50af669c8ac44061975b920561d38129cb4d547f16982ad8a6b43d1550"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f89653e39b04 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-7b56b36eb2a0d5b8df7419880d8cb866a406a26768cb63bfc251d14ec4d6c328"></a>

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

<a id="canonical-26e800e78dbb8553bc971be64862dc756f31dd027a411c683810d8cbfa8dc5ae"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f89653e39b04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-947644ce1f9f6671e0f7dae590742b280797f5a0e33354a9568837b7db6ee37f"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / f89653e39b04 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-18bda0e176281a9db3ee71eb68fa96494ece1be5e007c4fc9a660a95527f5f74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e77d88e51838010a77c71620ebda4dfb085f64dbd52f8528a2f22715880591e7"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-6657636173fc294a628847503f8534c0b2627d8bc136f38c4f5eb36bb482eb2c"></a>

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

<a id="canonical-337b3ca1b8e15c1ff946310393532c6f13c5e5d7e4c30fd4fbeaec63c4eee574"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 3

<a id="canonical-965c3da49ac6d9a78a1a045f23aa57c0bfaa95f5a67c513ebb178b73f2d3940e"></a>

<a id="canonical-8b8a14fa009ffb0120562389dd2e9336862439666c153194f2cc951ea6775460"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 4

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

<a id="canonical-26cc0770c17b8ec3f1531e8f717e87329106121c51ad5701f8509d32e30cbe5e"></a>

<a id="canonical-086ef227cab83d825208bfd53fc80fdd146452c0d08b26e6933681b146a9c863"></a>

## namespace property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 5

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

<a id="canonical-b53e6a7b52f8e069c23dfa81b521580e7d7f7d5c42e9d2ff3569596d269cfb68"></a>

<a id="canonical-1e7794ddc967151fc961806eb52561c6bc5ea55313734b0033b64afb3ccbf8ac"></a>

## tenant property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 6

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

<a id="canonical-d81a7dff9d6cfe82b49205f5373e0386b5453b66cc9e7fc5cc10e98cb5fd7645"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e084a7e67972 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a93062b4da6a42ee040774566c70a6f904195b61698217e1854c30a0691f67ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0d847427ab836b4d1f32f2242f762f5f86e96f7ca0dd27b46d8543c0199b72c"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 20f337441c55 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-ce6bcc6bf30c5bb12163b78326b6700f67781e7bc33c1d260abfbecd6e7ce277"></a>

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

<a id="canonical-6aa237b5d565d4f8eca97d7197fb226a68241d3221d65670e361fab8347cde27"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 20f337441c55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9a22ee405402cc5b72ca9cad9b704a3cc4353ed78f7a863d956eee40748dec4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 20f337441c55 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d6799d198efea25ca80761689e40ec104aa4f3cab05330e94733073c8935a45f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e8ef8b9592c2d4b06a991c2d5058a75849195db0c5d1c8ad61c35092240ac30"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d5d2fcc71d17 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-6ae7f2b71d6a5a98e32ded4d2b787278e735d1f3c72f9f2d0a5affc2d4f91758)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-ccb6358bafbe4fd33ea49ce251d3f0473136052ee3906207ff885267e7a59217)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-ca5b10e76e4122a386756c472958d6923da80d36dc75e727ac9cb1323b5b6157"></a>

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

<a id="canonical-fb1186349a4ce1d0f3691af01a9693684b54e17abf21f5b6e89328f683805d81"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d5d2fcc71d17 / 3

<a id="canonical-6214cf8b54d834e2b057a915ee0abceb691aaf8b0b5a1e0584bdb83d3dc5a5c7"></a>

<a id="canonical-affaa3f6ce27a296cd537474e580754de3a0317a960bbbe936a53408e757678a"></a>

## xfcc_header_elements property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d5d2fcc71d17 / 4

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

<a id="canonical-bf9c9b482a25c001a7183b67f878d2cf48b9ecf1c4333227f5bc4e34c11d8fa0"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d5d2fcc71d17 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-699c84c751399ddadaae0676e637244a942ad05b8220951d302818798c9049eb)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bad20c0173fdd77328cceb05d463392f1afd1fc7ca30ee11137b6a98993a210"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-e9f789c62d9884340ee4c070b4f0b329ec2c103c13a1e547188d1ebd4bff3942"></a>

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

<a id="canonical-10a0aec08c8dc89765c8c0c372f0b879969d8e4b9d6fa9e2b54b7b61bf13dbf2"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 3

<a id="canonical-a1763641cbbea142041dab8bfeaded25ccaeb0efb0cea499e28954578d941b59"></a>

<a id="canonical-62ee0910845997d74ecc0039f3b5df489130de271605a407ff0071611a113f13"></a>

## add_hsts property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 4

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

<a id="canonical-7a4e19d2830e697e869726095339f421c63a05f0fa55d9156efd223abd005dcd"></a>

<a id="canonical-87f34f68fccff22691a98919063521872ed103208d1a4ff24b149cf6d558130c"></a>

## append_server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 5

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

- [coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff): complete subsection reference.

<a id="canonical-0d7c97df312681ded5be9faec36e5132b0e747b50f30ebfae854a54ff86d30ed"></a>

<a id="canonical-bf3df4c05674a12ca99ae40371b3915736d5caf638c310601c7fc297c94056c3"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 6

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

- [default_header](data-sources--workload--reference--group-010.md#canonical-5e90105ee1436064f6dcf3c3fb49e1433cc9080285230174833f09f824d6197b): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-010.md#canonical-ecf999a678baabceac58ab45a1f8f7de1fde8316ef25008f486438925ebcf359): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-010.md#canonical-68da8c4115bb0d6763522a8f1dab486e3242d91cb0d679669f323b41eb7ac42c): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-010.md#canonical-d651a030166b21b1822c40b87142af8aefd3b6820ab8b77bd6f02fa7ca883d73): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a): complete subsection reference.

<a id="canonical-cec529c2e9318fa317458d1c55088c0843aff52ba8da374243f072bc57c717ca"></a>

<a id="canonical-afbab3bdaa1bfe01ab5d7c2f435e0f394f6dee79f9492df68f385b0d30c92852"></a>

## http_redirect property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 7

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

- [no_mtls](data-sources--workload--reference--group-010.md#canonical-aa6929fcdc8bf3a1950d85cb9026271273f527a5fb8ab70d582b76e7e6c317d8): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-010.md#canonical-151e61f1010249d95659ea661d7880bdfa03d5b494b20272622a376352ac93cc): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-010.md#canonical-2eddcb36cd7403664f233120131c573f2ede644a32c40e89e89e9ce44c6131a4): complete subsection reference.

<a id="canonical-d1c9b606da050b936de83f26171d19cff09bc1eeafc63579575076ce33c46cb2"></a>

<a id="canonical-21bd44ee02d0be4ec65ed29e0aa600c7621c463c641043e3d72dd187f295f43b"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 8

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

<a id="canonical-5a1339fe180a2861c8482abd0e29d1ee583a413b2f8c241307417838bff353bd"></a>

<a id="canonical-18c73783a8d146037abad229d086bd7214f6e816399e4e98e007a2a68059b5ab"></a>

## port_ranges property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 9

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

<a id="canonical-dee3a3cf0d03640e70819633991ce81187b4c25aed149feace000cd8d5ef325b"></a>

<a id="canonical-5695aa9c3fa79ffae7846a569e3c7dfe40978e3c3729246bdc2309ba6c9d66af"></a>

## server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 10

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

- [tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-011.md#canonical-726c7188e648194c0e45105ef129e812ac182de57a3bc6f1626722aee77b4fcc): complete subsection reference.

<a id="canonical-6d68e0809b71c3df1345f5235452e9d6c32eaa6b08dd24e9fbfd787e2d036c72"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bf5581e94c93 / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-010.md#canonical-5e90105ee1436064f6dcf3c3fb49e1433cc9080285230174833f09f824d6197b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-010.md#canonical-ecf999a678baabceac58ab45a1f8f7de1fde8316ef25008f486438925ebcf359)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-010.md#canonical-68da8c4115bb0d6763522a8f1dab486e3242d91cb0d679669f323b41eb7ac42c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-010.md#canonical-d651a030166b21b1822c40b87142af8aefd3b6820ab8b77bd6f02fa7ca883d73)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-010.md#canonical-aa6929fcdc8bf3a1950d85cb9026271273f527a5fb8ab70d582b76e7e6c317d8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-010.md#canonical-151e61f1010249d95659ea661d7880bdfa03d5b494b20272622a376352ac93cc)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-010.md#canonical-2eddcb36cd7403664f233120131c573f2ede644a32c40e89e89e9ce44c6131a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-011.md#canonical-726c7188e648194c0e45105ef129e812ac182de57a3bc6f1626722aee77b4fcc)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9460172b8d53bb4c1bc6ce3b97cf3e4066847a05f23caae562b677ce7bf5f9e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 79d5de23d6fa / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-6ae83b06221f2688cb78cfd5cb70244c44a838b323feb1ebb712de011c3857e5"></a>

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

<a id="canonical-7759b8b497e3d972a9b6357b7e5baf13b5eb100d3a6f5c714e80511ca07801d8"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 79d5de23d6fa / 3

- [default_coalescing](data-sources--workload--reference--group-010.md#canonical-51cfd02df2966e285d5c8a154539f56a69f61a99ea3debc925807f907fedb8c9): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-010.md#canonical-0396ed368243145d80a833530d81f41d47f909dd204da16cb4a7973dae1c12bd): complete subsection reference.

<a id="canonical-c751ad1de2fc47dab91202963786671a856bb76d9ca1b6020c336721c6b8eab0"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 79d5de23d6fa / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-010.md#canonical-51cfd02df2966e285d5c8a154539f56a69f61a99ea3debc925807f907fedb8c9)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-010.md#canonical-0396ed368243145d80a833530d81f41d47f909dd204da16cb4a7973dae1c12bd)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-51cfd02df2966e285d5c8a154539f56a69f61a99ea3debc925807f907fedb8c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f2898ae85630bea183aca3eb4050da618472fa45c9b2cc71d4a38f4df236061"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 83cbd9418ed9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-4ea0fdefad461a44eef9bd74d91f7337ff27be1f4263e0bb879ecb115d616951"></a>

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

<a id="canonical-707aa69e586f13439242992fa989857d78f5ea59a0f621700a489501907e03db"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 83cbd9418ed9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d789856588f8c6f4dfa97c4efcca9544e8e6c1c61cc1232ca5ce5db17560d854"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 83cbd9418ed9 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0396ed368243145d80a833530d81f41d47f909dd204da16cb4a7973dae1c12bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46dc752c130abd325be71ae11aab879e1f7527753b73c8a33d5107e83a5e50a4"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b564759327c4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-a37f0860a11bc001394a02959b4108e1ef8af72916f688af4d89ce99312a4c10"></a>

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

<a id="canonical-eccb94ca46a5ad7a4b870e74a5144c680b86aa66177095e0b0b3c3e09bbc8f0f"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b564759327c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3b4395542182d77c014178729df8a0a9882f18f45d07d8a0ccc47c163e3790e"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b564759327c4 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-010.md#canonical-54fb36d762e6112ac080b06b7acb4e10e51fe38b0885b7756b10e93212bac8ff)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5e90105ee1436064f6dcf3c3fb49e1433cc9080285230174833f09f824d6197b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c407722b8f3d5b8d72d6cc24d8362ad402b1961d75c6343f0e8d622602e39ad1"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 58cd3ace32f1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-a581de3ec1ac9bb752ee09b6216c8fed37de8f53b3b9421e84dd06fb229938a8"></a>

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

<a id="canonical-eac2f6ae002786a1e8154d7f918c5051b0c9c911d95de1bc50667a4818f6f732"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 58cd3ace32f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2603c3ff3d573656e9ec43faa0d7d3d9cc9074d3121bfac54c525a08289c6f2"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 58cd3ace32f1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ecf999a678baabceac58ab45a1f8f7de1fde8316ef25008f486438925ebcf359"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e97bc89bbd3a1a9378d8c2cf31c8b3f74981e2ea66e48fe475a64bede00d0cee"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0afc81f03fe2 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-3769389cbe4b1240a7a9465f875d9748d00526bdc3c619ac0cdabe06d415b066"></a>

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

<a id="canonical-1fde9f43df0013238eff8908797ef2b5af0e00be33f5aae1011230fa36b2e848"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0afc81f03fe2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-312d07f70764d191101dd0163134270c8ae60a9e47340042b424cced8bc092f1"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0afc81f03fe2 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-68da8c4115bb0d6763522a8f1dab486e3242d91cb0d679669f323b41eb7ac42c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa1249cbe9b9a94b6aa700fe23bcf184971532c995ca1d2f124b71e3b80d6c06"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2ae8219ed829 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-38f5c3ac983f09a04880fc7bfab79d9a82aa9be0cbaaf03bdcbc8ab1b81b6432"></a>

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

<a id="canonical-a3fff68adb42a5d36cae7997ab8241f0b2c69f29cde0575912f9665f69a3f5f0"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2ae8219ed829 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c7028c32bf765916f6961154f4dc3b3ca564b04898ec784bcb5de36f24f021b"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2ae8219ed829 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d651a030166b21b1822c40b87142af8aefd3b6820ab8b77bd6f02fa7ca883d73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d68a2d1556afbcb1878f4f7b82f3c9c1bc4e3bbf3322f4bdd2c379b08a22165"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 038b8ff670c6 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-67fa0ff857fd8604a9b9b8e8157a995e5abfba5f7c9b051ac900f098edfcffd9"></a>

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

<a id="canonical-53a77bf16a9ffc56e67a658597d642dbbda6715571cb91c5481c1b8c897c0fc3"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 038b8ff670c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9db72043dba6317bd2e565c4598e394d395e29b1eaf5906c5739e0d9056b7400"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 038b8ff670c6 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7364a8ba0a8a2cf32139d04ffb51702f5e73d18bfb6f13475f93104d0338549"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b6180cebabe1 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-eeafe005559c4ac5326635c26d02dfa5e23d7a1e7dd528388d01d561254103f2"></a>

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

<a id="canonical-0c346e19f6bda8bc1f44287bcda544899810220033cda56ce1758e3cdf0dd9b0"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b6180cebabe1 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-010.md#canonical-113ac7685e165f46fd87671fb389809596e65aeeade4d4f73cb309db404eac7e): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-010.md#canonical-0f217599b6cda057d16abe66823be70b3ce1af0450016b8dd16c2bb79b0514cd): complete subsection reference.

<a id="canonical-d6fce30990ce96bb2a5e9d7fb2725a0aefa0f8dc870ea087340ffdad6111e599"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / b6180cebabe1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-010.md#canonical-113ac7685e165f46fd87671fb389809596e65aeeade4d4f73cb309db404eac7e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-010.md#canonical-0f217599b6cda057d16abe66823be70b3ce1af0450016b8dd16c2bb79b0514cd)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cde367cd2a9b136afd5552b57b72ca1ea52c10b82d82b57128669a5b929cbdd0"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70409e250546 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-39d1ee4d3bc5018ee427df6ea916b5718f1e1aeb56cd75fac7a0f5b2fe75a8f5"></a>

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

<a id="canonical-ea8085260837aba69e8b38e7ff8cfc60fe01f8530ebfad9ce30c38980e730e46"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70409e250546 / 3

- [header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b): complete subsection reference.

<a id="canonical-99c5c1829f0fdc54233e7169106b5a993188745002a838e397788a8119ba6821"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 70409e250546 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c03c45746131299a8be68e5136196408cbd18bf268d384b35dddf4fdfbdd2a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 416a6a6513b4 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-788f65338749f786c826abe65b9412cbfa5bc1f8187ca864e241715dfa9142a4"></a>

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

<a id="canonical-84573a3439520a33f16cc4a5f573ba52c25d6f8b77c3adc39059aeea9510e343"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 416a6a6513b4 / 3

- [default_header_transformation](data-sources--workload--reference--group-010.md#canonical-9d7769de4e6dbc6e270237dea6cbe7a46e135196616217a93d83b5ffa860111b): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-010.md#canonical-00f26c44d5ce1963a30e9d2bede712d52ae92d7b32d367e8bc12e29b103e38da): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-010.md#canonical-6ce8b61b1b8632b67e47b8e342e3e6971ad3dd964ae8f6cd2027b3963fe71451): complete subsection reference.

<a id="canonical-114e8964fe1f0b6f83c71f6852727ff5acf6a7bd1b4b7eabb345a0a1e328310d"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 416a6a6513b4 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-010.md#canonical-9d7769de4e6dbc6e270237dea6cbe7a46e135196616217a93d83b5ffa860111b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-010.md#canonical-00f26c44d5ce1963a30e9d2bede712d52ae92d7b32d367e8bc12e29b103e38da)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-010.md#canonical-6ce8b61b1b8632b67e47b8e342e3e6971ad3dd964ae8f6cd2027b3963fe71451)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9d7769de4e6dbc6e270237dea6cbe7a46e135196616217a93d83b5ffa860111b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02526e16c6fe283c30e5060bbc57927345f23f9c4ec200cef949b03fe97bf1ce"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0d26433af227 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-b19cf1af3b44453eb4130d80f29ace16542b34bf1fbd0a59ee34334a9264e43b"></a>

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

<a id="canonical-6cbdd083028072bf2f4a4cbf083338e0266f20488c0e888061acd450b0287c24"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0d26433af227 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a15b8d6f2631126358dda331ad40e66af4c3a288af56574cc3732ac7d7121a3b"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0d26433af227 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-00f26c44d5ce1963a30e9d2bede712d52ae92d7b32d367e8bc12e29b103e38da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaac1aa824aa0c6ea691b06602bdcce9de9ee0032c6068e2a0f0f00501b674b4"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e73c803b40bd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-181a0f838357248ddb57e0a9b7064a5974966d84c76439226e95a05d73bf4567"></a>

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

<a id="canonical-9e62c12b34a34a2d0b8a0f2dfedb06333cf58fbdf4f83acfb9173f99990445c1"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e73c803b40bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9107bf3b5bc7f903a6735b048e6c06da5d2577ff13fe84f1a0eb8fa6f9287fd"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e73c803b40bd / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-6ce8b61b1b8632b67e47b8e342e3e6971ad3dd964ae8f6cd2027b3963fe71451"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-211fcd63c22244807dc151abc2538472c43d3981a5fe864ffe0cbc585c3afe84"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5e13142c0a5c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-010.md#canonical-1294687ea3a0d1206542700ed5f7973da4218d85197e3ad357cf4cd0a6e2dd2d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-4bd25411c0c07d27640a67c2adb2a7f983e65dda4e75c852e768a17f8fe87a52"></a>

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

<a id="canonical-14c05b5927ea094cd164fdb0a47e349417ef085e6312e0b36d897eae5d879539"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5e13142c0a5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ccf71ce0d9a938cb8dc2da60aa0658ce12d7580cdfd787f71fb14488a1689bbd"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 5e13142c0a5c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-010.md#canonical-dfd12220ac0b48b03c64fe0c95443cf159248a5ed0770874b3ed4627c831801b)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-113ac7685e165f46fd87671fb389809596e65aeeade4d4f73cb309db404eac7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b49415e975236c5db1378fc9be5bd61c0d273ed87fc2bf0e91c6bcce7f812c08"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4eb67908474b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-910b9cb8df1646071f4ee18e4425424b1e50ef3a525d2e2c0df7a686b921b1ea"></a>

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

<a id="canonical-dbba9727fb594cf45353aab59e8332c8fed923dfe790b92a71d16f87c7867188"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4eb67908474b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4de54ea8c4b64ce078237b95ea92e740ea3e982330abdc876ab83ffc3cc75a28"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4eb67908474b / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0f217599b6cda057d16abe66823be70b3ce1af0450016b8dd16c2bb79b0514cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58ab5e687336223b97ffec4056712570c07bdcc1e6ff0a626a44215375a69915"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 71ec469794a7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-740b3f9ec667a2b8a8a5eade3807adfc3bb432745bede70f92ec8c62da41adc1"></a>

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

<a id="canonical-5f529c90cb7b0d140572c5ce51c9d1ab5a551bdde2984de9cf614f864e1956a6"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 71ec469794a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3cf59e7a1b937f15c54f8ec37fa32ea6be4a6694a0c707d86927fedd68ca39c4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 71ec469794a7 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-010.md#canonical-2ac9c3eb41c506bcde236fe8ac978828188a4595285a59df090c993b13592e8a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-aa6929fcdc8bf3a1950d85cb9026271273f527a5fb8ab70d582b76e7e6c317d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60c5b5c1c443d00af83afdfa1264f1d8605a5f26bf77afe6994d75abfeb9380b"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 02609b00ba6a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-b80a17aab8c54280bf83af24bd3406b87937c9e4d4904e5415c74f8b7d2ae35f"></a>

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

<a id="canonical-2e5485d38a880a3563a3c47f5d5a9b65886341d0ba61ba4e435d2855112cfb7b"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 02609b00ba6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7beff4d6acb3b0c04915985bf19a84e1f861e78f831453028534785d5d3377f3"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 02609b00ba6a / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-151e61f1010249d95659ea661d7880bdfa03d5b494b20272622a376352ac93cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7f9831ce398824306a5326e1f35b7fe573106416fe4c37de161c3b5c446e6c8"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / db3183e0d651 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-53761695accc3e69955d9c90c4e21d04266aebc3f867841d30d929c980463a47"></a>

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

<a id="canonical-c178b06d09bea8670b6669436fff81eb70ee084b7718b73d44f535538bd79026"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / db3183e0d651 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b27fe0c01450f91d78b899610681b7ca42c8922d2bdae20f795c8c22dab8889"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / db3183e0d651 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-2eddcb36cd7403664f233120131c573f2ede644a32c40e89e89e9ce44c6131a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5baeb3d1e2cedce06452cb56037c1e963bc967037cada982034820ec2a3f604"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2fd1903455c3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-5a65e5e50cdd288739b7e26e691e281116979fcaf79fbb0c02b50df2bb270a84"></a>

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

<a id="canonical-2375e1275dfda84d6396fd7465520e7abd9ede91bc799c8bdc8aeee72f9210f1"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2fd1903455c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8deb69f55eefb8bf10900a68f3d0cefe98a0889be0b07bfddd5d72c00f5523b8"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2fd1903455c3 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53ebbc36f3aa8ace2623c4ad1b95c0ba43be596fc41b4a29e22bc11e051ed918"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 32a23e7c0706 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-8978b7d6b42e45bf7cd1fd3ed4af67047fc4edaae573b8a2bd7fa6f46fa01b7b"></a>

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

<a id="canonical-e6a3aa755012b9632b83e4b3efc66db72243f058fd16f83ab2c265cf754f9075"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 32a23e7c0706 / 3

- [custom_security](data-sources--workload--reference--group-010.md#canonical-9b1a0556ce8fdf8cdbcff4a9ec15bf08289dce4a8661e8b54399b297843701c6): complete subsection reference.

- [default_security](data-sources--workload--reference--group-010.md#canonical-984ba926a461753f8acaf4086ba46e90b0b9a05b0529b61b00a06c5e2785bcfd): complete subsection reference.

- [low_security](data-sources--workload--reference--group-010.md#canonical-260c8ed2956bee25e1e7827682b935d2f4bfab49ca12366f231bd03fb8fc1bc7): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-011.md#canonical-f39fdebe8ff8dc6ddefa8976e0a9bab8902491a10fbe474ee719389d4025ba16): complete subsection reference.

<a id="canonical-3e52329ec8c50c8de76196727df5af623963a64bd3b7d44b9c4c3b55dfaaaf63"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 32a23e7c0706 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](data-sources--workload--reference--group-010.md#canonical-9b1a0556ce8fdf8cdbcff4a9ec15bf08289dce4a8661e8b54399b297843701c6)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](data-sources--workload--reference--group-010.md#canonical-984ba926a461753f8acaf4086ba46e90b0b9a05b0529b61b00a06c5e2785bcfd)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](data-sources--workload--reference--group-010.md#canonical-260c8ed2956bee25e1e7827682b935d2f4bfab49ca12366f231bd03fb8fc1bc7)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](data-sources--workload--reference--group-011.md#canonical-f39fdebe8ff8dc6ddefa8976e0a9bab8902491a10fbe474ee719389d4025ba16)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-9b1a0556ce8fdf8cdbcff4a9ec15bf08289dce4a8661e8b54399b297843701c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2336ecf9566d85c4434e03fbaf851ef5da87250a4d683241a318411fe528d20"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-5d12ad4698c39acfc83dd1ca8d9236ffebcf1bddf50e35b60e7ae46c63175cb0"></a>

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

<a id="canonical-28c191625e992355777fd19bef17e2f1cf82f6ab3f8d2c4038766dfe9c4e7e7c"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 3

<a id="canonical-6e5eb0d66ac9a10941ab2d14259a34fa8768c37763fff3c8445c82f13a8acd05"></a>

<a id="canonical-5c8b29243d71cdfaceabff8b60bbdecac65af25abaebb3f020276353a492bcd4"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 4

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

<a id="canonical-97a6ae34afa368da1449437c51e3b2697b6319fda644fcb5863b68c59679976a"></a>

<a id="canonical-e72c45f084c3727cfdcf79a5cada863a68998efdaaad6a0f5fd7e94c4a9aced4"></a>

## max_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 5

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

<a id="canonical-c3d110292f10fe6b1922d70e4a4d4a5e17f9942a3333889dad9b2779df24265a"></a>

<a id="canonical-b17b43fef11fb0051ed4cbe6b2df03c43c1637d9be0da950ed2deb63009116d4"></a>

## min_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 6

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

<a id="canonical-edfece217588cad3c03a1ea85d0b6f39bb44309f613ebda230e24a3f0af0b985"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 8f363b1da2b5 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-984ba926a461753f8acaf4086ba46e90b0b9a05b0529b61b00a06c5e2785bcfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8bbe88f8bbacf0956c948737909f4dcd29dcb2b710c54c1af5a5de66fb2571e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1aa0eef7bdf5 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [service](data-sources--workload--reference--group-005.md#canonical-62f73a2321a24c94ecbdc8e6251a0c5b9ebab48950d28f103cc9e4f5fd742025)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-6c2c050a44ac31c47f219b2da9b91a659c0f7e07240a1d9ed26c328ae20319ff)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-7450468fa4e0351b28d1df2c51158b161e8d59b8946ec729f0b1b7f16217441c)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-dd7ee64823cb466d5d5ec98c954f713c9db16ae3b5d676ba8ba400b91a109286)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-097d167f64e2b428df3091a67123e0cba76128c80bc50f5aa467387925f7ffd8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-cb558922372d639b8f935627c0dc3027a3586d18d6ce937b77b1623f0a87a4b1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-3c6fe7acbb3e7d43101bb33001feeb507501240dfc6045bf42be34bbd5c362a4)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-4d91abc0be525145a9bd5d7a422db816f5bd5e3076376111b3c6b7b40259217b"></a>

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

<a id="canonical-a51e70dfd895083ea9b0c03929a7e2d673bef32c3f46ff692281699639cefb20"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1aa0eef7bdf5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82d33a71fb61c6361fe026030fb37fb640b7a104231a20ca74226d1771f3bc7d"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1aa0eef7bdf5 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-010.md#canonical-b15e182bca7d185919e739cb6411a2d5c76ef123cda968190f0dbde40953bf26)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-260c8ed2956bee25e1e7827682b935d2f4bfab49ca12366f231bd03fb8fc1bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
