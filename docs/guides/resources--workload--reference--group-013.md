---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-864b3c47575ad394b71661c2b79694b2f2a328fd0534d01d15418a6c0f697ff2"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 373d030d4354 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-7e65558f013baec4ae94037c071ec7f0f40794f51a3427aca5fbdda945308a91"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-6bcdfd94b745e79b956d0f1053d0b25effba37733684a4e73d59d8f001416c0f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 373d030d4354 / 3

- [custom_security](resources--workload--reference--group-013.md#canonical-04ec1c47d7ce4d61a4f64420cd23cfd0c54b6921620198f366f12e6eb80be766): complete subsection reference.

- [default_security](resources--workload--reference--group-013.md#canonical-abf5f6ea57c6292bf9bf001ebe8f77cb305100470cb80e2fe6c5157c8077eceb): complete subsection reference.

- [low_security](resources--workload--reference--group-013.md#canonical-01ed177659142ed2d95b186ca2d8d92847f656757895bd083fd7a87ad256bc53): complete subsection reference.

- [medium_security](resources--workload--reference--group-013.md#canonical-998cf42a9468497dba51963376a767067592023fe1f2343668efaa4ab7300ef4): complete subsection reference.

<a id="canonical-34998163128734b15914d9011f1ee7a48422cbcbbe0161659ac8fafbd6fdd963"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 373d030d4354 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-013.md#canonical-04ec1c47d7ce4d61a4f64420cd23cfd0c54b6921620198f366f12e6eb80be766)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-013.md#canonical-abf5f6ea57c6292bf9bf001ebe8f77cb305100470cb80e2fe6c5157c8077eceb)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-013.md#canonical-01ed177659142ed2d95b186ca2d8d92847f656757895bd083fd7a87ad256bc53)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-013.md#canonical-998cf42a9468497dba51963376a767067592023fe1f2343668efaa4ab7300ef4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-04ec1c47d7ce4d61a4f64420cd23cfd0c54b6921620198f366f12e6eb80be766"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27ea2f481257e988df4c8a45576d0b7194be129cd1871a16733ea0222b0bf97d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-dec626ebe2d73384d6744029f061b788b2b20b705992fc59249a70341a3be4b9"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-af134db6e9be4c47a6923854559b8ea3f6a6418441750eaf119bad66336b6b1a"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 3

<a id="canonical-4601969e633f17c36fb6ffb751679b4d8613242b85dffc48608740cf7dd05833"></a>

<a id="canonical-1c190a0e22c1f3d1075620335edca86c5f874a2b4871821c5d2eeadefef8ba85"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-15371684f3cbb391da28feae66614544b4bd93aacda93a2b4fc2fb95480fa1e4"></a>

<a id="canonical-b7673e7ae125ab0348f7441f7ec1481dde9368c9bd3e3ecbb70bcb725f97ac0d"></a>

## max_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-be947522ff20b43dbe96e7daafcd63b4d4361e3f711590eeabdf85f6ea698faa"></a>

<a id="canonical-84fd44b1835018505b51555a95f7e5555c8c77b47415fce1a8fdbc9c7b31693b"></a>

## min_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-42430ef9c6fb704717493ee7ec47cc5a34d8fb220a61621c22edae60b9116a08"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2efd934afdbf / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-abf5f6ea57c6292bf9bf001ebe8f77cb305100470cb80e2fe6c5157c8077eceb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0586b31f307d977ced8c62d83f9ee5496a3d4c2045790a6c05559a3cb09e07a9"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 99f3b811827e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-3a71cec8749e7828baee6038de72805ef050f7cb41806a4aa4b1c1dfb1a6bbfb"></a>

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
default_security = {}
```

<a id="canonical-ff9070f3627803313187b5c49d7c826a1b1d8d4b9fc2c342370db48bc56b5247"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 99f3b811827e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f196eb6cbcf6229b8bce296eae25f9852602b949254238e700772cbffa36457"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 99f3b811827e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-01ed177659142ed2d95b186ca2d8d92847f656757895bd083fd7a87ad256bc53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6a7a2e48239cd98a0f542159db341667b5474ad17824ecab50a3b243ec7dcca"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2d9097a25d52 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-9c63e3b1ea47d1ee287a8f58a1ee97f6f0f1bc6ea5c8c7f77d0da566c5f86b75"></a>

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
low_security = {}
```

<a id="canonical-8c351c25c46818875a048e7c6559a32edfffb38b73e21d754e73027e0127a3da"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2d9097a25d52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6af84fe4cd6a2b322804a67c1604a418ed7d2513764b747ca253f97e12c3bc0b"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 2d9097a25d52 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-998cf42a9468497dba51963376a767067592023fe1f2343668efaa4ab7300ef4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7b99158f3d29e4907c4d9de80724fad8020be8a8734066eb79cb38293940f06"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f821fb98ac1f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-5effdb7d0bceed7d1dee32d36f33ff4f4307041e8fc884cbb0276abbadd73a62"></a>

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
medium_security = {}
```

<a id="canonical-9c3f15b55ac10108b48c102a8f11e95deba5731cab87d26405e34de53793baf0"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f821fb98ac1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19f2d0c5064674459f07ebf7048276b2e482a544700e84fcbdb1a9dd2f699233"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f821fb98ac1f / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16ba9ee0846f30eb1b507229ea596ab87c3e4765af5e4f14e2156e626ffd00c6"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 614a6c43c303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3b5ff99eb13d468f77eede11f3bbf1e232d761639c97d3cc47e874a107193fa4"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-41255a9e74e889d86e1c71cc097fce9970d65422baf351d14dd3caca4cd9ce9f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 614a6c43c303 / 3

<a id="canonical-ed6699795e0682d830a9b5eee4200764c025a4cfefbe9d98ee406538c5458c85"></a>

<a id="canonical-9596b42f4e6538bbd984a7c46abad2a27c816951acda4065d10d62c65f3f5e75"></a>

## client_certificate_optional property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 614a6c43c303 / 4

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-013.md#canonical-bf1ead47c767d3f622052a503c403c0d4338d896a5fa33a69795480751ab0bbf): complete subsection reference.

- [no_crl](resources--workload--reference--group-013.md#canonical-066f1605da9105ce33b38e5858ebe963805cf07777e0d5228cef89be40c87537): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-013.md#canonical-a2cecb04ca53c847da75ef548504108634152805bca9b0863e7862816db6a889): complete subsection reference.

<a id="canonical-0a1eca29147493d329bba59c51593f6cc8db6f0d68a66d47fc776173ebd159d4"></a>

<a id="canonical-37e437944f85592fbf5da8e26d5fdbbbf4c5462ce490301a92bb1bee7312ca49"></a>

## trusted_ca_url property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 614a6c43c303 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--workload--reference--group-013.md#canonical-398c5378707aa0314dc642f7d3422ffaa2d8a33eddbe7d4604b66b4e62a692ff): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-013.md#canonical-41c3fcb98bc8914ebfcdd09b3d17d2c95ed495852c9d62f8adfd0e18d5f41064): complete subsection reference.

<a id="canonical-c50cd143eb21018b6097115d2129ce270597825caec067a4159e68cba82dedb0"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 614a6c43c303 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-013.md#canonical-bf1ead47c767d3f622052a503c403c0d4338d896a5fa33a69795480751ab0bbf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-013.md#canonical-066f1605da9105ce33b38e5858ebe963805cf07777e0d5228cef89be40c87537)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-013.md#canonical-a2cecb04ca53c847da75ef548504108634152805bca9b0863e7862816db6a889)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-013.md#canonical-398c5378707aa0314dc642f7d3422ffaa2d8a33eddbe7d4604b66b4e62a692ff)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-013.md#canonical-41c3fcb98bc8914ebfcdd09b3d17d2c95ed495852c9d62f8adfd0e18d5f41064)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bf1ead47c767d3f622052a503c403c0d4338d896a5fa33a69795480751ab0bbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1a2df5bd0c259cfcf600110473d2863d5685f79618eff563a768422d5afc0a9"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-a0197e8e4ee3fcc9672c93f210dfefcb30fce89dc9865c1c76e65a6ca38a8918"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-7d9b99af92c41e4aabc7fdbb953310b7442a570f776f5f54772f69d3c639e5de"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 3

<a id="canonical-9c6b670088cf7443e201e790c3eb84fe57cf52616e5a4c233c133aaef4cbcafd"></a>

<a id="canonical-51b43fdbece2922e3c734c4fa2d737833605d607e070e125a35c9f7a3a8d3a23"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 4

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

<a id="canonical-bc482e9902a8227541ca87c0390334fdd6330aea91d30c4270ddb3a794b6cb45"></a>

<a id="canonical-e201f321d55181310b6db2f54c3c64009849b425e8a50b30fac87c27057de273"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 5

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

<a id="canonical-1a658ede4f189b3553033437b5e1054d98067b9b3660d566953702ad3492b252"></a>

<a id="canonical-2e3a1067af8c22ef720aff178746e3e4ccfcd27a6e3b4d1f94690e86077f84b3"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 6

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

<a id="canonical-5cd19855a60634dff379b2073f9a5e7d3da73aed9755e4b8392b1b9c7d821cdb"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 36c3ce32284a / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-066f1605da9105ce33b38e5858ebe963805cf07777e0d5228cef89be40c87537"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86d5b47a2aa0fc095e761cefe4eb32c1beb75809a0eaa862fcd33a01cfe8bbbc"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 667cea615b2d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-ae2167f66d2429960f5f90277c36d274ba3f48ad6a98cfc3288ca073a3a4a00d"></a>

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
no_crl = {}
```

<a id="canonical-0724814c7cc25fee48ff5334749f7e08185e420814b1f5638c709bc28e5fc363"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 667cea615b2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4658b0390e3d8f29d759c034174e78041190f95bee51356a101a6251ee80493"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 667cea615b2d / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a2cecb04ca53c847da75ef548504108634152805bca9b0863e7862816db6a889"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d40ce0468a3a07794d10b521efd6b35a81f6f88053d6212f3d3c7331f90a5ea8"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-ecd3b356335f73c7e10f7aa338b4c6367c232583baa7895f0f59f82efcb7d877"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-93082d5d2c8c85d94bf95db9e600aa2e6017834732f0f5921f7ba3395e237f1f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 3

<a id="canonical-20953fcd8303e1d42208dfdfe9c85b6072682d55b4049f67dbe428c848c365b2"></a>

<a id="canonical-ef2b7f2c0a782b1b9184638d8afeba4ed40fe007f8ea68e39851ec1f55138309"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 4

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

<a id="canonical-522d276372a4fce78685b6039280a88b2d48f7bb87260fb1d43021ffb2edf418"></a>

<a id="canonical-67ec314d8290e4593c8bf01c3f20a1945a14f0042aecd5ec1984f6d135b71dc7"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 5

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

<a id="canonical-c4bd773814f3b43a9024724e5208aa62d1bf8abda70aa159cedbf0947ce0ff1b"></a>

<a id="canonical-d373c7c6780cdb27a006f20b36d48acc3f44604e09ed66c693e02ba35638ec39"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 6

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

<a id="canonical-bc7b6214f93f9df8777f9c9b269100470cdcc443d5d2226f0c6c43876a1b9fc3"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 60707d532430 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-398c5378707aa0314dc642f7d3422ffaa2d8a33eddbe7d4604b66b4e62a692ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-049718e7bf4375fe23afe46b2ba13da1e2f5c79d1964974773116414df70c0c3"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 118dbe8614b9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2211c7bd7f1ac8042405791d6af4412b372b1bfbc190028b304ae5f082000f17"></a>

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

<a id="canonical-4a74f18b758e912255d178fa2b3966a4cd7b69e0355919689b2328c3a552be27"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 118dbe8614b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da6bfdd54b009af16c971e35c9233cbfa45be985b43ed37f2dff999b85bc07ce"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 118dbe8614b9 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-41c3fcb98bc8914ebfcdd09b3d17d2c95ed495852c9d62f8adfd0e18d5f41064"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d766a9d67d9920b568b05c77a3aeeba74a5a5e011ad4d41a33d71c5a31c5ca1a"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67c01871dc42 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-8e3745edc31a8713b931982d03a16f649ac4ef5c87bf5f6a026afaf43e6197cd"></a>

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

<a id="canonical-e3fee0216c0d32e165693d2d176e7a02459e8e0a11390d38bec240132ea8e119"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67c01871dc42 / 3

<a id="canonical-9acaed97bc5af8c6f855d01eaf51b856a01ac95e23a86ceaa4b1051e4936227e"></a>

<a id="canonical-9fd22d0bb889039ede17c90207b6510e59e556ec84ffdb1f79832e5ad2415e1a"></a>

## xfcc_header_elements property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67c01871dc42 / 4

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

<a id="canonical-ee226b55f6278b99db594c9e1fe77ae3a217c371efd415a13d658148a8ab9b03"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 67c01871dc42 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5634816905061c9ff69bcdc11d4925d75209d1e2b54adc423cc4ff235fa7fc8b"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / c0debe23e017 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-d2c3f5f8a557ea835ab37332368c3917ed8f83bd7f754423f72e8a7e6f6130fd"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-9fa70da9620c68fa4855a15d77c8ed85ce07f12ea948999f7e90c55b6325618f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / c0debe23e017 / 3

- [no_mtls](resources--workload--reference--group-013.md#canonical-2624d5ac6a858f1cc916ba0760547882881eeb16bfd820dfed5d5c92b051fbab): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4): complete subsection reference.

- [tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780): complete subsection reference.

- [use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5): complete subsection reference.

<a id="canonical-2c1117fa028b3ab6fab51dc79a2c8164a3041b3663a423020859f8106b77a578"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / c0debe23e017 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-013.md#canonical-2624d5ac6a858f1cc916ba0760547882881eeb16bfd820dfed5d5c92b051fbab)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2624d5ac6a858f1cc916ba0760547882881eeb16bfd820dfed5d5c92b051fbab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-595d8ed6164d00c534e2953ee13a317c799a0bb0abd098ee9c1259c93e2c193d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a52bd14d6f5e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-916ba2db52eaf995211abad928d1c95462dcf314d1f773069ccf23a8f75bd1e1"></a>

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
no_mtls = {}
```

<a id="canonical-97fd53b830b5395a146d3261dd49f63ac702fd2d89f9a678a2a67efc3727e805"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a52bd14d6f5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-228026937a362b081820fbf185327c3effba9e312c3c40129772c4133239c3e9"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a52bd14d6f5e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-791826809a27228eba8803f6d3f1e29168e5624ebf116ee455c168a57b2e8261"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ae631133ab88 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-6398fe5652a9f35bae4a0aeb294ab1f0e37470af8a4fec002255d094a10af5a7"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-81bcae0e302bfd5750e4b2766718a484bc0fb260f3a3a223e36c638fc04adbec"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ae631133ab88 / 3

<a id="canonical-36a01519f903c198b04f75da631048894bf30e28e35ab240cb43835b2ac0035b"></a>

<a id="canonical-0cd06e017257764ff89f07ea23449abea5fe6ac1a3d374439a5719ec293c5acd"></a>

## certificate_url property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ae631133ab88 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--workload--reference--group-013.md#canonical-110342ac3454ede2c3f68c72a12082305dd0bbf7b2f6890dfc7125f4cc8cc0c7): complete subsection reference.

<a id="canonical-afd4006e19e1791f8a6948575675c5ee076df606cd9df899d3fc8e193bfaa791"></a>

<a id="canonical-aaef3cd0905620b0a4d1aff96e701f39a824cdc814f26a834a02b2bd13629d4c"></a>

## description_spec property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ae631133ab88 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-013.md#canonical-551ef11da9517718b6678d8ab3e54b81b5bfff23b3dc11d6ee4d469547ccbddc): complete subsection reference.

- [private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-013.md#canonical-e868583fc1b37728c7ad033a49c441b90eac7bf992a08b3c57bb74361b88e482): complete subsection reference.

<a id="canonical-7370a15879a93c4dfa915cd62418cb0145f0ab05791e5f262e6145118bb4416b"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ae631133ab88 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-013.md#canonical-110342ac3454ede2c3f68c72a12082305dd0bbf7b2f6890dfc7125f4cc8cc0c7)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-013.md#canonical-551ef11da9517718b6678d8ab3e54b81b5bfff23b3dc11d6ee4d469547ccbddc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-013.md#canonical-e868583fc1b37728c7ad033a49c441b90eac7bf992a08b3c57bb74361b88e482)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-110342ac3454ede2c3f68c72a12082305dd0bbf7b2f6890dfc7125f4cc8cc0c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e306bc3a1ed04b36bf51ae30128c2dd04c21e72c6020f026442f5d1743632c8"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7ff3976eca4c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-39b7dee89222d53f405ca0b462624d094406ac0cb9af4a2ce70275fcbb7afde7"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-ca153250698ef9b27736b6cdcb81b4c1f3d55f096333272508f3f15bc8f97d79"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7ff3976eca4c / 3

<a id="canonical-1a8e160e6e12e0f891e2b17c63742142cbd4e4a887950187366472f67dc84c91"></a>

<a id="canonical-5961d02d44ece97fcf94f5386ba1f30f90c20a7b7e2f92356d2d18f29f044796"></a>

## hash_algorithms property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7ff3976eca4c / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-a8767db442a47add2bf6c8e0603d4afa7870af09cf50c3c71117f350d1100c3b"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7ff3976eca4c / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-551ef11da9517718b6678d8ab3e54b81b5bfff23b3dc11d6ee4d469547ccbddc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a4261d46411666259829e2f2ffbb8714cebd337e95ca92180fa3de9e61c8730"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4607fd58d303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-7a51807539f48efa06fb03cc00dfe9c1cc0f5e84151a7056aee48946a9f2ddc4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-3f720ab3e126763e9641166f3869d7623e6fda518b59bf57cf39c6b475c3f50d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4607fd58d303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25781639280a73c2b2f07e804567586952105e818ba64fb4e635112443a4b3f0"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4607fd58d303 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edfc5491462eb7442526801fa455c4894f25066939f1c1157b7fcb7f818a8534"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / bd5536f49045 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-10e1463dad6a994c67b05884d857cd992b55cbf887894316d3608fd5ced1e1a0"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-9962df62e9a7413e6497ca8c04decd1cf9ea11807f7a5c83e51daa7a786bfa3f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / bd5536f49045 / 3

- [blindfold_secret_info](resources--workload--reference--group-013.md#canonical-cf9cff52bd9ac660925849b3e56bfa23ff1361349acc5e9c6da06149fa443729): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-013.md#canonical-7ff744b7873388b644fb53cd150108ac2310d7c628657579b630d837c951a42a): complete subsection reference.

<a id="canonical-d3b5d169d746716b584565fd3ab54d5db2eaa24a103109984d78b76daafb62cb"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / bd5536f49045 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-013.md#canonical-cf9cff52bd9ac660925849b3e56bfa23ff1361349acc5e9c6da06149fa443729)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-013.md#canonical-7ff744b7873388b644fb53cd150108ac2310d7c628657579b630d837c951a42a)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cf9cff52bd9ac660925849b3e56bfa23ff1361349acc5e9c6da06149fa443729"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba3adf465a843dca0f2a21da5c52c8930445d3954c6d9e833b6d13f6c5e4534a"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-44624b3864bc4a93738d3e259ea900b5cb1e1e1c12e20a63046cb3553089a3e9"></a>

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

<a id="canonical-fae65f437c508e39a6304bf7c648af84d40f357e1fb6a8327ced434a0639c8e0"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 3

<a id="canonical-186e3c9cb16ac30c8887a93c8a2ae44b60f6d9fe4d9324366101aeeb35e8a95e"></a>

<a id="canonical-6a2ffa0b7305b61f2d072ab979071a0bd29b963e1f59e23df8fc279a2763c6c2"></a>

## decryption_provider property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 4

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

<a id="canonical-6484dea9e1f4ba59342d2534f20a8e716360b85e849c45fffd653b9d612425a9"></a>

<a id="canonical-c5a8ce0ac075daa060e14c40c6c6acf86219cdba3bfd0c70171d70701e85b6fe"></a>

## location property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 5

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

<a id="canonical-1db838168a2e079e3979ba72c94f3484cd27b11ced5bb0288b0903de18f516be"></a>

<a id="canonical-0b304ba68b96267555b39c9fccd39f8b680f3f87e40737c38a7f9d8c32a7a0c6"></a>

## store_provider property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 6

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

<a id="canonical-ef030082045f0ccf28b464f34397aadc5ff8c1237675ec1d45a511899cdb9338"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 244173fd66dd / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7ff744b7873388b644fb53cd150108ac2310d7c628657579b630d837c951a42a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02f03620e1c894382cb9c1371f1a4e0c4b00a38a71df658a10f4fe280ada4797"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ff0df91f1338 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3ecbfffb2c4ecc1ba72df2d8747cfad23e78c22e6966035bda67a0dc9f5ccb33"></a>

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

<a id="canonical-7a4b6e105b8172fe3668d16e329d64c92fb867668a8f7351d1fb30147e92721b"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ff0df91f1338 / 3

<a id="canonical-85b53980b479670c4e179283bcc9f2fc5758692f9a7c93f2e3f72b44f6013cd1"></a>

<a id="canonical-b161b3073d1e8e9cef2ac272703b6eee429bc245612510f29d98e552eb68749c"></a>

## provider_ref property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ff0df91f1338 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-514afb2bec4c5e98045886e0ec7f7ea24cbd53942ed0620a0539fefaf3d93f10"></a>

<a id="canonical-9c8cedd69b917925efcbc1505d8bcad980d1f2e5a256b1a6091763b21922305b"></a>

## url property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ff0df91f1338 / 5

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

<a id="canonical-7bee739362887e73595a0fe42a55bea54e09db3f29c4f94e011967a3a3dcc413"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / ff0df91f1338 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-8f58a1deca8e6d475cde0eada61c14040558d767c2045321197c919546abeb30)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e868583fc1b37728c7ad033a49c441b90eac7bf992a08b3c57bb74361b88e482"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-576c99d7a63e8e1a2be16e925f83b52832443e04daee895bfdfefb955e251c94"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4bb2e3fd449e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-ec6ef8ceb748014e27ab61976e6fe23a12cb2625c13282d78b9bd143228864f4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-d5df66fbaf0d96ece66561e1326d65b60ebf1a22e7332e539ddefa36faf3673f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4bb2e3fd449e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48b1c75be665b01ef6d266ca92044231d649eee09e4bb73dd9bbe00bf080a2b0"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4bb2e3fd449e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-6aee5f255d1471c5045f12aa302d66531e2011a7573284566f056d74b93a6cf4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0158a7e800a5a23ca1b8257c0dc67cbcfd975caf51db3ae27f3d30ed358fbec5"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 0a427442d21a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-68fb10f66f689ec23396de14fb349ff0e876a8c6eb4c3e0bdd5859c7ac38f802"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd8caf206bb525aa3c38efe369aac2c071c01c0bb995534cb892d7669490955f"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 0a427442d21a / 3

- [custom_security](resources--workload--reference--group-013.md#canonical-25c1a61d6c4e8a1ad1bf0744fbd3697d1cd823c3197e3ca2ab6d8cd0de310aab): complete subsection reference.

- [default_security](resources--workload--reference--group-013.md#canonical-9be7815cba88a7236e2be6c6fa71cb5846e8f7badb90f8f30e50ae3750f9c7ff): complete subsection reference.

- [low_security](resources--workload--reference--group-013.md#canonical-7c61bb68548e33f4f2cfac06ded5db495eed51f8c58d5fcb1aab22d2df7ca71c): complete subsection reference.

- [medium_security](resources--workload--reference--group-013.md#canonical-2bf63ed21cc6e42b6fe2469a51cd1e941c012639debdec43cdc43957be035f6e): complete subsection reference.

<a id="canonical-dc8b24c4d3586f9574ad01368d8c16ac3c49373749c9081f956b2b835f98e76f"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 0a427442d21a / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-013.md#canonical-25c1a61d6c4e8a1ad1bf0744fbd3697d1cd823c3197e3ca2ab6d8cd0de310aab)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-013.md#canonical-9be7815cba88a7236e2be6c6fa71cb5846e8f7badb90f8f30e50ae3750f9c7ff)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-013.md#canonical-7c61bb68548e33f4f2cfac06ded5db495eed51f8c58d5fcb1aab22d2df7ca71c)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-013.md#canonical-2bf63ed21cc6e42b6fe2469a51cd1e941c012639debdec43cdc43957be035f6e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-25c1a61d6c4e8a1ad1bf0744fbd3697d1cd823c3197e3ca2ab6d8cd0de310aab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f94de314af7cc1c2bf283ef2d6c8bbe96b81fcf9fec75aff806f553cdad055a8"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-067b6e58035561a9e3a946028230f5ee0b867593eff6fd5addea4b675525ee5d"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-58a0af809f79500afdceb208734d74bcf9d7c2d9993603a60ba55e801ab35bc1"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 3

<a id="canonical-ea1dc47fc5aef270d55dee5ced995c66f7f1edb6df8b8b39fad709f41460a76f"></a>

<a id="canonical-084264023cb115e8b3a9c9154474414c2fd684a9022de14255732bc8edc6e5f0"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-4260857e956287094197cfb430b8cec2c2db876c9f0a7c7752b27643956853d4"></a>

<a id="canonical-db1eddf73999fe72d7c0f4173f27b24c381f55f0975b705b88c064d7f96a11d2"></a>

## max_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-b8d0af3daf34a4528465f42cf9860eeb0ca53b248454b4c4d3b31ac1efafe6dc"></a>

<a id="canonical-36a153231dcd8ea1d7c0811a80ddcba46a718f1181bebb2dae3a473586a2f5d1"></a>

## min_version property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-c3965f7b764e8775b8c86ace56efc477b8cd36ab966ae843a86c6dcc695244ff"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b594bb41354d / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9be7815cba88a7236e2be6c6fa71cb5846e8f7badb90f8f30e50ae3750f9c7ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f6fe43c0a3ee20e9e4b743e5eec960b4b073d34d6864def43a5ce2076219b2f"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a2591fa07952 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-85fe4376e4d775211f1578574fcd42d42c0a1c86e9714f3f4df0590d7f182484"></a>

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
default_security = {}
```

<a id="canonical-7c04632a4ac5d9fffb4d9fa90351121ea803abd2c55f9468b2fd7f28a6ca81c6"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a2591fa07952 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07bbab9ca5176c2f67d3c3137c78619474747f0fb86902ce5143e8e18568df77"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a2591fa07952 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7c61bb68548e33f4f2cfac06ded5db495eed51f8c58d5fcb1aab22d2df7ca71c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-056f9b0138a065c17089d401dd58966828221ed07ef22c3ccbdf0149d41b1386"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b61dfcb928fe / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-50fc689103cfbc9f592a282344f5953aee7a1f0cd912f3a56436e14063a45474"></a>

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
low_security = {}
```

<a id="canonical-971cef00d9b1fcd59e81db322d9f98338bddc249e3ab75c7cb9029f70f9052de"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b61dfcb928fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6354431d8b0ef0b5b5a64a9e389f71e3bd010b074f8da4f741690c5c0d7fe92"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / b61dfcb928fe / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2bf63ed21cc6e42b6fe2469a51cd1e941c012639debdec43cdc43957be035f6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e248f61697098ed4b38b71d78234006fcdb691058a69512dd5856bb9825f385d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / e0b052726a10 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-a353b2b987a7176bbd537874fa06f7b739cc8fa1f157846bfbb22532a1597498"></a>

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
medium_security = {}
```

<a id="canonical-0c54a45edd1e896d19f84f780310ba1ad78b8830042cdceef4c1c103e7b2db3c"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / e0b052726a10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83f541f1be1440208df2bcefc7206aed687e67bb8d77e0a0e9c95345c64a4728"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / e0b052726a10 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-ee78b0adea1b66cd0edc2bbae8be3330de1273d583d3cb797744471dc2803780)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba9efab6f2805bdb09c3074d7d220e3aead309e967e000da29f0c4dea9526603"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7adb88eae18b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-29714898dabf1584f0942b092e797c3b06a25c67ad5c4834692d33fb373a163f"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-add9b77d1d98ce32be5447b44d57b54ee8f07449ee42f20a90e3ff0ace010d1d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7adb88eae18b / 3

<a id="canonical-234195cd527dc0bf34b0895ee5fa7ad7f8b792f4154911f4b843e6e2158c37a0"></a>

<a id="canonical-3e96e8bd722b39303aaea1b9e7f2f3c9ac9bf507c176132ac8415f1c7f8dbdc3"></a>

## client_certificate_optional property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7adb88eae18b / 4

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-013.md#canonical-a4fa768a37dd6aa469d5ba88937611ba0eb9e4d808c2fa889e3b7bcd4c98696a): complete subsection reference.

- [no_crl](resources--workload--reference--group-013.md#canonical-a2fc14783ce32d4e3f7a2c33a99a68755c7c0b20ca66db2791eebe442fa9da43): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-013.md#canonical-2b39c1971ba82321b55d43e77a7970cb8ad0ee55d615abbfb68be6894d2c1b9d): complete subsection reference.

<a id="canonical-20384f29bd3e5d19c32ceb5503603bd22bb89a4b83616eb387a834b675144aba"></a>

<a id="canonical-d7864bd0365e1335249b3ee521d40952e025aff9243f50b572c8d34ba2e3c12f"></a>

## trusted_ca_url property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7adb88eae18b / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--workload--reference--group-013.md#canonical-2ab5c0265550fe27b34ea2476b7cafdfed6f740d8fa3868f2b85460fb7723fda): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-013.md#canonical-cf669e7ee2b4f59e251723664d753b0d28942004d3959586f6b17a068417eaa2): complete subsection reference.

<a id="canonical-5058732bfa6234e66719e1d3666d922ddb862368990183d9b351b05e52965435"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 7adb88eae18b / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-013.md#canonical-a4fa768a37dd6aa469d5ba88937611ba0eb9e4d808c2fa889e3b7bcd4c98696a)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-013.md#canonical-a2fc14783ce32d4e3f7a2c33a99a68755c7c0b20ca66db2791eebe442fa9da43)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-013.md#canonical-2b39c1971ba82321b55d43e77a7970cb8ad0ee55d615abbfb68be6894d2c1b9d)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-013.md#canonical-2ab5c0265550fe27b34ea2476b7cafdfed6f740d8fa3868f2b85460fb7723fda)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-013.md#canonical-cf669e7ee2b4f59e251723664d753b0d28942004d3959586f6b17a068417eaa2)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a4fa768a37dd6aa469d5ba88937611ba0eb9e4d808c2fa889e3b7bcd4c98696a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcc72cf22e187549d786ee014ad0912f2f8ff93c3b3b0dea31c0fb65770ce976"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-986ac8459d46970e8b860a7a0990926ac751be9bf7efd2f73232ea102d8e9414"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-3fdfadf6994f83dc38edf07a61220b409e17a65620d96b3f2a191a6d301e7d3d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 3

<a id="canonical-ad4d63b591e4ca1bd62ba8249bc9c2119b04956ce0f223f8073177d6a564eae5"></a>

<a id="canonical-2b331f95b926c4b4ee09cd3c23903ecf38462be9a43e6003c1e4d2e17c69921f"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 4

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

<a id="canonical-22685f901c2439db2f9d136dafea4bd6fdaf210c2743399d8e9a6e6d30a7bf38"></a>

<a id="canonical-01544da1d0a6cb100d6753cbc6c19806e491e2ffd6fb94105c93963e9854226e"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 5

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

<a id="canonical-151387b9eed1925a3265ef48035dbd35a5e41a056e94f8b558d1fa9dba492311"></a>

<a id="canonical-a5b89d49c4ddfa862e5ba66604749e013f75be79de7af4f4b986115879d4d399"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 6

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

<a id="canonical-e0cf27b9985ea393efa83087441b070fca4d01a36dc19ee26af2c1d6097d6428"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d4dae5ef851d / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a2fc14783ce32d4e3f7a2c33a99a68755c7c0b20ca66db2791eebe442fa9da43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13cdee0980f61d95f7797fc3e9aafbd65c2fa241ad569460a998c7dacbc61009"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4fee4e998c42 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-4548ee7bfd37db4843df558b255ef0c2446d69c60c8849c7f91f16b2ce1aa6c0"></a>

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
no_crl = {}
```

<a id="canonical-f1cbd6bb7ab293539330eda9f0dcb4467961615fa1fc8aea7cd8550f19f8082d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4fee4e998c42 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75cb5c95718895f983ad81ef50312016c935929079ea89eb5e6e863a9fa46fe1"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 4fee4e998c42 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2b39c1971ba82321b55d43e77a7970cb8ad0ee55d615abbfb68be6894d2c1b9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c411b490618532c7a7c08495f68ae64c4262d1fb2d54cad9d1ef16fa01464fd4"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-253d17646024389119c78cb8544671794cac0e41c7976095fadf1a1585348f7f"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf3c98c5d37855c59329b9a3fcd245e699ebc400905ed661b20218b41162a9ce"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 3

<a id="canonical-6ce7a4136f8383b923fb540e1a93fff992ebeda42fcc9678e87568b05236602c"></a>

<a id="canonical-62c257178a1d765dc0d7c647aa16b5136abf1b29bb07f982719ae13828d06a2b"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 4

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

<a id="canonical-781a3029cc78acf9c917af4f5d8c5245d3529069a75adf8868a33f74ed2f9b90"></a>

<a id="canonical-35ead8c5e7e3cba9e28ffed4434a08afa5c07aff1ec9ceebaa30af75332796e6"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 5

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

<a id="canonical-229f225b7e79e622a71efe4d323abe10f17863f09df69f3b72867e2fc731efd8"></a>

<a id="canonical-b0d0acf4b22a0eb30cd998d2b2d109b14c1da50f5768018a61bae8dcdeaeb1ac"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 6

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

<a id="canonical-7e6cec82d552adf24ce74035485acb7605f943278739f9983f15c832e8019087"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / a5dbfbf7532d / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2ab5c0265550fe27b34ea2476b7cafdfed6f740d8fa3868f2b85460fb7723fda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc375f0603b97ab47958971591b0d167a73082ae647cd3ade43474048fbe8c0e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 3fcb683331b4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-babb1695b737971ab9ad40804301b5fead682b33efd0ab242d3791183aa156ba"></a>

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

<a id="canonical-eab1a959b80725e72543336157fbb66ac332b04dddadee95f8f808b3a927db42"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 3fcb683331b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3ad8d1b135c3997938fbc1ce7a541de7165680629eb81075d2ea3aa6d126467"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / 3fcb683331b4 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cf669e7ee2b4f59e251723664d753b0d28942004d3959586f6b17a068417eaa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a82ac3875e69a1f61356d19cbeb07eca1eea7ebf317b94db72f80c8345c3055"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d9bee8c9d56c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-e05cf1eb6a1566a43595f8175582dfbaee462da61a48b10c9ada82b3d08521bf"></a>

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

<a id="canonical-be262861a1e30aaba3465d946ef16e06df37dfba61c8faa36f97312f2013facd"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d9bee8c9d56c / 3

<a id="canonical-39f6caddf0ef896ec3def03410ba15b08661345de613d0e6bf137f8369b0a7f2"></a>

<a id="canonical-e94e095eb1bbefcc43f7203732fad11733e0cda91dd655ed8cdc7811566b2a71"></a>

## xfcc_header_elements property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d9bee8c9d56c / 4

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

<a id="canonical-af0ad3e2233a9e6fbd24eabe7a624804dc18fd4f50aa9fac4ba48b912ba113d9"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_p / d9bee8c9d56c / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-a63de6fd4a2d1422068258fbd0ffeb8e80f640a3cb7760272ffe1e9b828824b5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-af097b6f6116d0cbb4580bc318179bf7da7a4923b0776b78bbdd457f4624a1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3389a960aaf0dd3d62cdcb5ed1f16d4ce4fc211090f33c8e158b60581a502f68"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-ade869ee11e8216dd9702720fdbeb2562dae0ce6b20601eef8fc1dce6f6fc925"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-7355105f11e6f029099777ebbb5044a0cfbd24e35b7869a36d7860d4cc5aa279"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 3

<a id="canonical-924ed70aba053f8c83e9f4271eb7342b4b884f5742b2453ad170a8454860e18f"></a>

<a id="canonical-7bfb7fbed358e42622f333b03586b8a67749337aafa88ced37a913dac9cdc044"></a>

## add_hsts property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 4

Type: `"bool"`. Optional.

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

<a id="canonical-93d245fdc35bc3bd0b1dbbb5c58211ee1fdc8a32d54c6ba21cf59dd24c7f9a7c"></a>

<a id="canonical-861dad0a18d1659763115e875a8074278fb07578477ddf943e25cd63ea30cfc0"></a>

## append_server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](resources--workload--reference--group-013.md#canonical-43df60f8741fd0c8b5162eb1d76b6ad29d4c5bded5317b68da3e660447f2fc24): complete subsection reference.

<a id="canonical-7dbfaa1ae0721542c4fad0389e176a97da629d2e956b0e98894e8bc93f058c04"></a>

<a id="canonical-b9aafbbc4adb461d9ec381d619932fdba9cfbfffe4c865d1fd4da688487f99a1"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 6

Type: `"number"`. Optional.

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

- [default_header](resources--workload--reference--group-014.md#canonical-a4b91fbfb8ca5986197e5376480964362ddc29a51cd5145b5c7a7c904f3ce283): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-014.md#canonical-4737b782314a34b56d5423ad7cf00dc78a6437a1357289bdf45470f7bbf18565): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-014.md#canonical-12f71b80987028aff185d496a589718293bb3c829b2300e813708580fa56319b): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-014.md#canonical-04b1897df8ce8845d6b4ad62c5e41d3df9425337b35b886614248ea08aabc41f): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-014.md#canonical-418679fcdbd8351e5158f24db842b57ee0e0436ff3767b3e04035e41155afba1): complete subsection reference.

<a id="canonical-597109857fb6bd35ad36dd4b1d7968274c460a246ab9844367ae7fee6b7bbc8e"></a>

<a id="canonical-3683631de5dafec6b93e8032219f30bd18e67acd2de799232a3bb2f0715a3826"></a>

## http_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 7

Type: `"bool"`. Optional.

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

- [no_mtls](resources--workload--reference--group-014.md#canonical-914101fc94cf12a4cf3bc89efb7ce3fe6d94ac2b2a9e81761d630096ead4d15a): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-014.md#canonical-35565bc9685b3784c8f4a955e918cfe895c9e426eec28d6f8dcbb462e8a7d704): complete subsection reference.

- [pass_through](resources--workload--reference--group-014.md#canonical-3c577e2e21e6be9320b85a79ee9df807c47a55299d55140f13ba0fe85d3d6ac3): complete subsection reference.

<a id="canonical-d87facc7066ca2215687df0263027e7a921ff85e015f76817e0f2c81f36aebae"></a>

<a id="canonical-1de4c4fcaa140d9becfb738f16a9849c14893d159111d92ea3d5f47c8ebff432"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-f39f3f5de37a4b4b0080e3e15a22637fc2e1dfe646b2ea44a667d40f2d07eab8"></a>

<a id="canonical-dab35a96e73e5eb46b1caf0af4b7e97f3a434bab4bf52032f707cac4eec54201"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-c343965fbdf373038e72a7127fd4b1888dba027a947dbbb42f2f4b34cabd7f6e"></a>

<a id="canonical-f7a552524b3e208422d5626a9ded9835ab17275e8d9d9e32c2b43959cc548091"></a>

## server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](resources--workload--reference--group-014.md#canonical-fe0aade2dbe509c9746f6387114b76fb8330d287a6ae1a7b238d32358e2c3f0b): complete subsection reference.

- [use_mtls](resources--workload--reference--group-014.md#canonical-51ca81c05ced80a9a750acf242412025e331589749bd3a48c893be28c5035d4a): complete subsection reference.

<a id="canonical-6d670dae2de1c0d67da462c81eb823ad5b5b01c9e755174662f889b2feda57f8"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / d8994835cfce / 11

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-013.md#canonical-43df60f8741fd0c8b5162eb1d76b6ad29d4c5bded5317b68da3e660447f2fc24)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-014.md#canonical-a4b91fbfb8ca5986197e5376480964362ddc29a51cd5145b5c7a7c904f3ce283)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-014.md#canonical-4737b782314a34b56d5423ad7cf00dc78a6437a1357289bdf45470f7bbf18565)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-014.md#canonical-12f71b80987028aff185d496a589718293bb3c829b2300e813708580fa56319b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-014.md#canonical-04b1897df8ce8845d6b4ad62c5e41d3df9425337b35b886614248ea08aabc41f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-418679fcdbd8351e5158f24db842b57ee0e0436ff3767b3e04035e41155afba1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-014.md#canonical-914101fc94cf12a4cf3bc89efb7ce3fe6d94ac2b2a9e81761d630096ead4d15a)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-014.md#canonical-35565bc9685b3784c8f4a955e918cfe895c9e426eec28d6f8dcbb462e8a7d704)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-014.md#canonical-3c577e2e21e6be9320b85a79ee9df807c47a55299d55140f13ba0fe85d3d6ac3)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-fe0aade2dbe509c9746f6387114b76fb8330d287a6ae1a7b238d32358e2c3f0b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-51ca81c05ced80a9a750acf242412025e331589749bd3a48c893be28c5035d4a)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-43df60f8741fd0c8b5162eb1d76b6ad29d4c5bded5317b68da3e660447f2fc24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-761e58c9166b2358f81218d0d042568159d9ce41a9537404940a3c48bf99b036"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_ / af4eae2c4cf0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-013.md#canonical-af097b6f6116d0cbb4580bc318179bf7da7a4923b0776b78bbdd457f4624a1cf)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-9e66eadebd9a3ffada84041661ebd9ce7a05db4f1a5f8acd36c6606eb7c0e7dc"></a>

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
