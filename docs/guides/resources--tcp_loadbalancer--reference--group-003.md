---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-65a770cc745f955f1151fcae356addb1f91e3f6450b0a75e6db591fb38669c25"></a>

## tls_tcp_auto_cert.no_mtls — tls_tcp_auto_cert.no_mtls / 79a7ef6438e4 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-b76060dd354205d41b3cf6e0c381ec28490050e9e0011f219419dc62de7e8b07"></a>

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

<a id="canonical-90ea71ee3c0b5ca38bfeb8eff1b1a53378ba9e9c98cbeeb3327221ba6a4e2b78"></a>

## Direct properties — tls_tcp_auto_cert.no_mtls / 79a7ef6438e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50775cb5a1bf630e50e228b60e6f552d54e1dcad86df1082156ef2733b5fdc0c"></a>

## Next pages — tls_tcp_auto_cert.no_mtls / 79a7ef6438e4 / 4

- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba3455eb49f12d6a87588724cf7fbe72c0b394af6e4d6bebb9f16fd52f3a4e9d"></a>

## tls_tcp_auto_cert.tls_config — tls_tcp_auto_cert.tls_config / 973e513ec835 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- tls_tcp_auto_cert.tls_config

<a id="canonical-f9d219f6524fac96a79a280ef2b5143c27a7f5654bd547db08dd8dd670722b12"></a>

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

<a id="canonical-1ecef196a05c70c9e7833111f1fe4b9d50076210aa3002622d833de5651ed7f5"></a>

## Direct properties — tls_tcp_auto_cert.tls_config / 973e513ec835 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-b71c80c5ebb64e3e94d52958cf7bb0510593a74f1ea084cf03608bdba82d7076): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-609f494277c73abdb679671a70be302cea03af5a4b1b07af857727b6ba0322e8): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-a159989e00157f8cee2228389d61e6f03dc95fffb7a0164ad59a97b356d7177a): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-c0ed8b91291c4c14a7d46ee12778342c99ff0fb5512f4672d9189991fcd534f8): complete subsection reference.

<a id="canonical-4be5f87727d47c280e4fa731866be2d7f74d83fac1a6c0d1b68912fab182f5e9"></a>

## Next pages — tls_tcp_auto_cert.tls_config / 973e513ec835 / 4

- [tls_tcp_auto_cert.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-b71c80c5ebb64e3e94d52958cf7bb0510593a74f1ea084cf03608bdba82d7076)
- [tls_tcp_auto_cert.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-609f494277c73abdb679671a70be302cea03af5a4b1b07af857727b6ba0322e8)
- [tls_tcp_auto_cert.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-a159989e00157f8cee2228389d61e6f03dc95fffb7a0164ad59a97b356d7177a)
- [tls_tcp_auto_cert.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-c0ed8b91291c4c14a7d46ee12778342c99ff0fb5512f4672d9189991fcd534f8)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-b71c80c5ebb64e3e94d52958cf7bb0510593a74f1ea084cf03608bdba82d7076"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cb9e8c2018a989870004adbb55b166870f3b2be6e9ccb6c32331ce73d8bfab8"></a>

## tls_tcp_auto_cert.tls_config.custom_security — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-dae249d8ded308f90b8df7a98679aa4e5e244aacd4ffa94360524be0e25d493f"></a>

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

<a id="canonical-6057ea911f3dcb6fc0d846f3d229374a61c99d92ea8c33b184bd2242477dd287"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 3

<a id="canonical-d4b24d2945ef03ba1701cd848ab5b5cb4ac8b6a9d5c1f3e9d1e5fe90efa64960"></a>

<a id="canonical-1707c18771655902c771dd19ac92aab47c3b581d446a4b81f12c97efacc3aa54"></a>

## cipher_suites property — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 4

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

<a id="canonical-f2b238884819c50c885f0f18c697c98ea64ee7ed6ca247bff12ee0bb4b09eec6"></a>

<a id="canonical-910dee5761bb4f92c8a2bd49b49a0b211ae81772da428211926b57a46904ca54"></a>

## max_version property — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 5

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

<a id="canonical-bd1bd4bcaad7d6a1d24f63113817fa0e1379c208644c3b4c307365d7da5ba6b0"></a>

<a id="canonical-17e866a4b8dea74a89bd3c3276344923ad75cf13815a370cf4b82eb578d064df"></a>

## min_version property — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 6

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

<a id="canonical-e3c34f21b62c52f2eae8a0b5740555324975bbd9d81e57b954b8c568085f7289"></a>

## Next pages — tls_tcp_auto_cert.tls_config.custom_security / ed8a7ecb6ca3 / 7

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-609f494277c73abdb679671a70be302cea03af5a4b1b07af857727b6ba0322e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59cfe964ab165b67d195d555cf672e612642aeba41e335383381882d6172033a"></a>

## tls_tcp_auto_cert.tls_config.default_security — tls_tcp_auto_cert.tls_config.default_security / 655fb196e8fc / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-841321beacee51a236b218d2f005ccba7cd5d24217ca05579a9eee54a560dbc6"></a>

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

<a id="canonical-8e758a1abd1c7e9b0218ec77f75a277239f934524349a7d9353d56cedcc10a9a"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.default_security / 655fb196e8fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-242d92c3ba59001aa8cf88c54af5c3e19dc7c3d2800008b6efa4d88e8d20ac03"></a>

## Next pages — tls_tcp_auto_cert.tls_config.default_security / 655fb196e8fc / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a159989e00157f8cee2228389d61e6f03dc95fffb7a0164ad59a97b356d7177a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93a2afd94afe1ae4572c1c52e7662371851c3c10b949b9c74a0b44a50cdaaa13"></a>

## tls_tcp_auto_cert.tls_config.low_security — tls_tcp_auto_cert.tls_config.low_security / 060c3bf979ed / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-06f71cd71a1cd624d57fc08a6bb9c128274990f6ffcf5e2dbce0c27b5e18396e"></a>

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

<a id="canonical-0a7f5fa99730a88b60b7797e2885a96d9d38c996fce769c73f3a9e6a8f60f249"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.low_security / 060c3bf979ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40324a05b62d4f791a51d4e7b9aca8927643a224f8fe04741c1da5a2ce9cc101"></a>

## Next pages — tls_tcp_auto_cert.tls_config.low_security / 060c3bf979ed / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-c0ed8b91291c4c14a7d46ee12778342c99ff0fb5512f4672d9189991fcd534f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7b9c17530f2ed3d0787df2962f31cb8c632d1718b13df94d5fa20355cf59929"></a>

## tls_tcp_auto_cert.tls_config.medium_security — tls_tcp_auto_cert.tls_config.medium_security / 6e8d5c65c2ff / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-a3f5b47de0a00b09163da0873fba92dd4319be8641ea9a2180017413953fb436"></a>

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

<a id="canonical-581e9fbadd4c5f553614538fafa27118d95ebb3231f4c9472f5c1fd8739a6a05"></a>

## Direct properties — tls_tcp_auto_cert.tls_config.medium_security / 6e8d5c65c2ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-211313824a32d6671882eaab48d16f8ca9d65b8a422c25c88c227c9ac338abf2"></a>

## Next pages — tls_tcp_auto_cert.tls_config.medium_security / 6e8d5c65c2ff / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-053c13231bc39596a3fcdfa982f6c3b4827504577c4d740441df4aa4acfce9c2"></a>

## tls_tcp_auto_cert.use_mtls — tls_tcp_auto_cert.use_mtls / 4b42115940e4 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- tls_tcp_auto_cert.use_mtls

<a id="canonical-9a6f4d817ec8e7d816a0a0573e25a30a3db14016fc4bd4564b3b034a1f3ff74e"></a>

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

<a id="canonical-48a68cf0f9762be97727ae022ce871dc3a772cd1bbe7bd8fbb8ae4b057d37146"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls / 4b42115940e4 / 3

<a id="canonical-49fa2a4d05e8f50300ea151eb68e1a69744b9455bba66f79c09811cbf38d3f3e"></a>

<a id="canonical-79ec082fd9601eb5021adfb841977d8cf51c6d97f26bf6cd89b575b918058f9e"></a>

## client_certificate_optional property — tls_tcp_auto_cert.use_mtls / 4b42115940e4 / 4

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

- [crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-65755232402fdbedb08820a5b6e5420aaaeb05038c46dd7b98ddb34ddb4f7723): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-29f7befaaf5675599362549c967de47d909e319767ce9145967ac3018ab685f2): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-77e8938d9f5f65c737ff6c1100dcc1099140f0cd93b5064a752d94f13662da5c): complete subsection reference.

<a id="canonical-71bf92f1e9c5be87d8560415e77597a87a742704497c877daa92c73577cac3d8"></a>

<a id="canonical-72fbde94c10349995d57b8d438803b1c68d2c998218d9f678b21f50db013fe5c"></a>

## trusted_ca_url property — tls_tcp_auto_cert.use_mtls / 4b42115940e4 / 5

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-feb8f1c08ea9428f01c0fd9f25349bd414f425767544af5ceb7fbee6cd9bb934): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-5c7da21a63845fa217a1ca79cd68705365986cd3a450c5571f632e94549a0430): complete subsection reference.

<a id="canonical-5d1dc3df947d0a63668acac3720e0ae6f7a6785d5138f2ff1c2630ccee011eae"></a>

## Next pages — tls_tcp_auto_cert.use_mtls / 4b42115940e4 / 6

- [tls_tcp_auto_cert.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-65755232402fdbedb08820a5b6e5420aaaeb05038c46dd7b98ddb34ddb4f7723)
- [tls_tcp_auto_cert.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-29f7befaaf5675599362549c967de47d909e319767ce9145967ac3018ab685f2)
- [tls_tcp_auto_cert.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-77e8938d9f5f65c737ff6c1100dcc1099140f0cd93b5064a752d94f13662da5c)
- [tls_tcp_auto_cert.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-feb8f1c08ea9428f01c0fd9f25349bd414f425767544af5ceb7fbee6cd9bb934)
- [tls_tcp_auto_cert.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-5c7da21a63845fa217a1ca79cd68705365986cd3a450c5571f632e94549a0430)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-65755232402fdbedb08820a5b6e5420aaaeb05038c46dd7b98ddb34ddb4f7723"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a326cd0bfba054d9dd6e2f63127c369cc11f5490beda22992d2a3d852767c70"></a>

## tls_tcp_auto_cert.use_mtls.crl — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-484eab373ada77fa19c7bb4d3231256a300a3618044f45a2dc46807ac0a27b37"></a>

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

<a id="canonical-47e21ced6eeb706d094505c5c180640c68666a567fad935944d8ac96a51d3754"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 3

<a id="canonical-8a551dc437eaf6477d062c186b53c7164c7674c69e513261c7647e210413ea33"></a>

<a id="canonical-40019175027f871dc935e6bebde97a01739d677856da8abf9f3ebaf8eb35a359"></a>

## name property — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 4

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

<a id="canonical-9b9dd3d711ceb02556160599d4faae50e7281bad706998ac06d6c90fffe7aad5"></a>

<a id="canonical-59d4c10da44d17a25f991f30975ae6dd18ad19399b3034aa89a74f5cd4ca1b98"></a>

## namespace property — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 5

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

<a id="canonical-cbc06ec5d6997216e3888f99e2432b8dc0ec6f776508db4f7b2f5ac0c1a7a630"></a>

<a id="canonical-43db798010078a413a0e2b7fc8461cdcafb9c45a564739d8d94f79befb0d60b7"></a>

## tenant property — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 6

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

<a id="canonical-19005c832bb8daffba0dfd177cd038a8fad3f87db6ce99da3264c291b35e0a33"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.crl / 1c80edbed2f8 / 7

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-29f7befaaf5675599362549c967de47d909e319767ce9145967ac3018ab685f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af843fa9b641e031b3e1b7835543aa55f60e1d1d1e7dda7ccf32959b23b0bb60"></a>

## tls_tcp_auto_cert.use_mtls.no_crl — tls_tcp_auto_cert.use_mtls.no_crl / 41cf57b7c234 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-7d8e5434b81927930c100af8f0fd518584d1dc0044edb3d6449108d1b4e2baf2"></a>

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

<a id="canonical-aa4a33a5d1b992e5c07cefc81cd713d8201a4311e8d0a43953c545437d0f0134"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.no_crl / 41cf57b7c234 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4695bff685ec198171ce1d9f16a29de7b26b3c783553f64038398d04691be29f"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.no_crl / 41cf57b7c234 / 4

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-77e8938d9f5f65c737ff6c1100dcc1099140f0cd93b5064a752d94f13662da5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3db341044a13e804efc441d6ec6cddc760387245c143ca661d4fa5127875b51f"></a>

## tls_tcp_auto_cert.use_mtls.trusted_ca — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-34c229689c4c84321291ae7c8cd294a4e5575d5ef34e50ef7b8b247b53186745"></a>

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

<a id="canonical-1a10086d9b2b8915a1f674db390ad9527ca656a1be88be72964e7d34822f4655"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 3

<a id="canonical-1159fd3b503a0d5f85558686183d2cc853cbcb8d3b47d7db24eb2c7d175c7caa"></a>

<a id="canonical-13d31d7b76b9b613d26318f3c3b96f004055774e3213ef105f180de8bc1dda60"></a>

## name property — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 4

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

<a id="canonical-d0fd23e36c6885b1c42e7cc058d0332cdd196469ff2f174711a14f9c1cd838f8"></a>

<a id="canonical-ff98f8d90d6eaf41d19673d7f2aaf212c932ba4ec5acb7a774c003b30f46a454"></a>

## namespace property — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 5

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

<a id="canonical-de0afcf1d217c9a6617b28a1273ed47e1511808d3c0cd109e750547345351231"></a>

<a id="canonical-04f158abe3a97fabc77decab02087ae68693cd2fa545c894777f5b66f1810ea6"></a>

## tenant property — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 6

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

<a id="canonical-9fd51435d90e799d31ad565ff9007483348748cf7a00809b4f3b0bae9b729b4f"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.trusted_ca / 58495af53214 / 7

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-feb8f1c08ea9428f01c0fd9f25349bd414f425767544af5ceb7fbee6cd9bb934"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7b3d64b0cd3846203cbce6ea51e25596e2cb75d505a40a7a2d3ae24aea28a1a"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_disabled — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 034b1f850c16 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-074ec61f8e0c8fd8574c8335f022a9ff62b3f538641f246b7bea55c4f9e13271"></a>

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

<a id="canonical-d5d4f354e6dffc2ead6afbcd7b33c858f2d9d4e3f263f091893ae71bd3dfc027"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 034b1f850c16 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a28141e7cc51fa4224be5cfc3b919b8c5b5acff4d93c46bb9a5acd59a2138b1"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 034b1f850c16 / 4

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-5c7da21a63845fa217a1ca79cd68705365986cd3a450c5571f632e94549a0430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2146e57adf3bb6eddd49bc82f1b2e02163020b3b2946744d62b4745777aeb51"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_options — tls_tcp_auto_cert.use_mtls.xfcc_options / 690528251f51 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- tls_tcp_auto_cert.use_mtls.xfcc_options

<a id="canonical-a739c6c05f47696d1c43ab48dffb9464a25d74679fa7bbf0db8cd85f7ae33512"></a>

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

<a id="canonical-851f7c0e587b5a19554961eef7ff769d21f9c01c8ec6f1704e55f2e9ee1d0e81"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.xfcc_options / 690528251f51 / 3

<a id="canonical-5acb5161ef09d5e61528baf5c2050ff49e5741b721512e6cbe4ef24d6d96e5de"></a>

<a id="canonical-7518ec60253c7500676082ec4457d5581f92d9a1e79144b2da56f98d08e269c1"></a>

## xfcc_header_elements property — tls_tcp_auto_cert.use_mtls.xfcc_options / 690528251f51 / 4

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

<a id="canonical-25694f7d9dddb00ba705ddff1a8d1ac36d14a7876ffb6128a9f80a143dfa48f8"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.xfcc_options / 690528251f51 / 5

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
