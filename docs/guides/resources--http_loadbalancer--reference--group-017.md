---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-52449bbb7d86d80ccb048ab920455a08d393a19d34f4ec944940019133ea5076"></a>

## Next pages — default_pool.use_tls / 9968264bf97b / 6

- [default_pool.use_tls.default_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-a2acfe534bdcaa0e5053d2f480bea97b11a74d89d2f878881191a1949391d9f1)
- [default_pool.use_tls.disable_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-47f48b7f228f8c096ef34438b03482235bf3677bc15e0db31a30286aef3ed799)
- [default_pool.use_tls.disable_sni](resources--http_loadbalancer--reference--group-017.md#canonical-b09ff81e1b6e78eb20e610fe5f76673a18292c7aded52fc2609604bbeff8d198)
- [default_pool.use_tls.no_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-33f4e0a4361011987c5811f404e04aba9e70703e6d77518bc949bafdbcaac10e)
- [default_pool.use_tls.skip_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-33f9e5356bcfd9228c0a7d2d9124d4c2ea9ac37da7b1cf02f374d37484d15c4c)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- [default_pool.use_tls.use_host_header_as_sni](resources--http_loadbalancer--reference--group-017.md#canonical-7a219fecc30a63d0bec0125952244ad9dd6f6eb848785eec222ba376f7eb9bc7)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-85d339096a8a94d45c3ee4ac5d310e910ea0d15e221b1ad1c7454382122a6ec6)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-f26b369f3bbe3d3bc658c0b08395510e8df06b4445cd510817ab56c342ad0a0c)
- [default_pool.use_tls.volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-f913ba70e9bad1bf406bf57ac712f432b527b62659c57f4aea342e0ad11610cb)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a2acfe534bdcaa0e5053d2f480bea97b11a74d89d2f878881191a1949391d9f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-515c902c6b4e65971f90ce35bfc2564b3c2da93611564f45eb36d489d9c7b58e"></a>

## default_pool.use_tls.default_session_key_caching — default_pool.use_tls.default_session_key_caching / 0b7871973559 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-19809b076452817f840870dc3ef4f67de4a394df03f5b81d0df9da8eff06f29d"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
default_session_key_caching = {}
```

<a id="canonical-491cd024ca2008a4c5601ba438f93f0eb57db7a2a51f93069bc5977ee0e40cff"></a>

## Direct properties — default_pool.use_tls.default_session_key_caching / 0b7871973559 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64d60f5c84f3e6fb950be88f00b957bd21f767f3d48170b705fbb25f132be3b3"></a>

## Next pages — default_pool.use_tls.default_session_key_caching / 0b7871973559 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47f48b7f228f8c096ef34438b03482235bf3677bc15e0db31a30286aef3ed799"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f5ef1295bd8512b1cc4c16b0b1a25bf529bf2d2f6d25d44fa62dfdfbae243e7"></a>

## default_pool.use_tls.disable_session_key_caching — default_pool.use_tls.disable_session_key_caching / 438345452837 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-df27e1d3781ff49e69715eb1939ac9250349a2cb25164948696d9740ef425a53"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

<a id="canonical-2e1965f00d084f2b30b72da6a2fdce3b465403e67f0544642063d1a27dd833b9"></a>

## Direct properties — default_pool.use_tls.disable_session_key_caching / 438345452837 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26e00b434f228619dadaba5c039f16ffc30a3ddcac78e97f9883f9bdf997bdb0"></a>

## Next pages — default_pool.use_tls.disable_session_key_caching / 438345452837 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b09ff81e1b6e78eb20e610fe5f76673a18292c7aded52fc2609604bbeff8d198"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83359e47601941658dea503713a06cdae31ae43ba898d840b28620660ee7d0d5"></a>

## default_pool.use_tls.disable_sni — default_pool.use_tls.disable_sni / 181171a69ed5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.disable_sni

<a id="canonical-a899e8778c5bddd5843eb007267ed38ae4235b31e519a1ffd78a0fdb9a7c69cc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-cf3592fe7cd46454fd6890f81e0cce95764ac5da310077a90d6bea527b1f493b"></a>

## Direct properties — default_pool.use_tls.disable_sni / 181171a69ed5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73bd8fea6ac4ebdbba513b36520e02de12ea780a2545851a49176d85a425ba59"></a>

## Next pages — default_pool.use_tls.disable_sni / 181171a69ed5 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-33f4e0a4361011987c5811f404e04aba9e70703e6d77518bc949bafdbcaac10e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-538ec2ab3aff9e92e3a4557d37bb941b68b6721a20135ef15305e98ef29444bf"></a>

## default_pool.use_tls.no_mtls — default_pool.use_tls.no_mtls / 095898c72504 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.no_mtls

<a id="canonical-7e45c482870ccf63b9df795977dd64cdde49bb635738d6a84558e11fecbf9c5d"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-8f16662c2b1ecd9f203062c842fa5c780bc03eb2c6f0a3cac2b890d93f3077bd"></a>

## Direct properties — default_pool.use_tls.no_mtls / 095898c72504 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-624dd9f370481a10b0e4574a353e5f331d690bc7828f8dbedcf210fffef2e93a"></a>

## Next pages — default_pool.use_tls.no_mtls / 095898c72504 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-33f9e5356bcfd9228c0a7d2d9124d4c2ea9ac37da7b1cf02f374d37484d15c4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f027e4f17db165f55fed39be6023c69809cf9441940b293cdd94cf036cb654e4"></a>

## default_pool.use_tls.skip_server_verification — default_pool.use_tls.skip_server_verification / 19991172db4d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.skip_server_verification

<a id="canonical-660a6146f17f6cd74736853bf3ac0568c6dee88f9e7ab62a3ad45da2023a1cb6"></a>

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
skip_server_verification = {}
```

<a id="canonical-f2eace3b3b3fc52f6aee4ff506b63609aa62466c1ba8f6146c1bd5182ec83310"></a>

## Direct properties — default_pool.use_tls.skip_server_verification / 19991172db4d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-308876fec3095b4658541bc0c77385a8f764a030032c3444a60dde66912dae26"></a>

## Next pages — default_pool.use_tls.skip_server_verification / 19991172db4d / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d1cef8bbb1b984e17bced0219430048a1c7f3650ed315062f5593039d8e55ea"></a>

## default_pool.use_tls.tls_config — default_pool.use_tls.tls_config / d9ca5a449ca6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.tls_config

<a id="canonical-e0e9ee040676e4c17ae4f64d0d3d1cae77f094d6d2c65a84413ad199b86dd1f7"></a>

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
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c6ad6727dd1f3412af55a22dc9f054c91079dda5566a1e819b4505355a5b5f7"></a>

## Direct properties — default_pool.use_tls.tls_config / d9ca5a449ca6 / 3

- [custom_security](resources--http_loadbalancer--reference--group-017.md#canonical-a0fcdedac61610badcf2b1fd6fd1374d8c181cde339ad6016148583944fabb6f): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-017.md#canonical-108ebffdea5f4f4dcbe5ddf6f43e92047772b3f34282dd42c8fe376de82424d4): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-017.md#canonical-3343eceaef9eff1addaf5280798ec5b60da6e41a8ddb1c07a66e440c132b2d5d): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-017.md#canonical-2a650cfe0bed8fbe3c0456406420d31ca836d6d818fc23c1a9e8fdfa1d0f0770): complete subsection reference.

<a id="canonical-e55969d5ce38b8db78b44aa0edaa0f4a6674b5430298f96d1a2fe0b83b9983f8"></a>

## Next pages — default_pool.use_tls.tls_config / d9ca5a449ca6 / 4

- [default_pool.use_tls.tls_config.custom_security](resources--http_loadbalancer--reference--group-017.md#canonical-a0fcdedac61610badcf2b1fd6fd1374d8c181cde339ad6016148583944fabb6f)
- [default_pool.use_tls.tls_config.default_security](resources--http_loadbalancer--reference--group-017.md#canonical-108ebffdea5f4f4dcbe5ddf6f43e92047772b3f34282dd42c8fe376de82424d4)
- [default_pool.use_tls.tls_config.low_security](resources--http_loadbalancer--reference--group-017.md#canonical-3343eceaef9eff1addaf5280798ec5b60da6e41a8ddb1c07a66e440c132b2d5d)
- [default_pool.use_tls.tls_config.medium_security](resources--http_loadbalancer--reference--group-017.md#canonical-2a650cfe0bed8fbe3c0456406420d31ca836d6d818fc23c1a9e8fdfa1d0f0770)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a0fcdedac61610badcf2b1fd6fd1374d8c181cde339ad6016148583944fabb6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92a0689bf9ae72a7f19a611b3f818e2c4f167284f9467ff97e27025b74b84ac8"></a>

## default_pool.use_tls.tls_config.custom_security — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-687d6b53df8862ac3d6cd206d613cd562244351267a7190230fc94ce59e8c5c3"></a>

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

<a id="canonical-d60e6898c95c5dc52fbe915f4522d1f46d81d207a8a5b70f33fb946424774a4b"></a>

## Direct properties — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 3

<a id="canonical-b2775e74b7567097104c0a37b8cfec2db328338fde7f3833aa9034113973b41a"></a>

<a id="canonical-8239d9184d41cac7c807c02e2ad5ac503b9a9d11c2ed7a7ed2b38380a4be9ea3"></a>

## cipher_suites property — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 4

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

<a id="canonical-c9585af80978fdf14bd29531f536f954717b4f40846be4d49fba59fc6320c435"></a>

<a id="canonical-469525e9565bfc0dc90353e6fa59ae35b1ff8dabdd80195d372b2fdd3d549806"></a>

## max_version property — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 5

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

<a id="canonical-715768d9f1dc7537d1867024ec7f4e0c0dae74c65c6c5afc4439cccafe1b0dfb"></a>

<a id="canonical-de0585b56200f207b35bc2ab3e14eaaccd785ba5ebbaa6f111c31a2afb67f94b"></a>

## min_version property — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 6

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

<a id="canonical-713b9523ddf01889a466698cdc4e565bba8d7bd9c407ccedb991cfd89260b6f6"></a>

## Next pages — default_pool.use_tls.tls_config.custom_security / 51a55b4a3dbf / 7

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-108ebffdea5f4f4dcbe5ddf6f43e92047772b3f34282dd42c8fe376de82424d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6b97f2a50f50dc852880b355fc6d0adeb3d246710f5ca7a58f74ef4f119f63e"></a>

## default_pool.use_tls.tls_config.default_security — default_pool.use_tls.tls_config.default_security / 03daa35c2a35 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-8166d5fba2c8ad877d1aef69bb4d530d5ca8d7dd17de6ee730ba6ebff8013059"></a>

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

<a id="canonical-60564ee08b36dab1b6ea6d28b983b5df44d3a38bfa5708fe2a4440af543cecd1"></a>

## Direct properties — default_pool.use_tls.tls_config.default_security / 03daa35c2a35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-714f91e46682314a6d21ba091b727765bf6fac0b8f1e1e7f89db1b58349a4fb8"></a>

## Next pages — default_pool.use_tls.tls_config.default_security / 03daa35c2a35 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3343eceaef9eff1addaf5280798ec5b60da6e41a8ddb1c07a66e440c132b2d5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cffaf420c9fc04c0e85bffe0de8ca72a0f24f7f4f009931afb03e36f8607f2fb"></a>

## default_pool.use_tls.tls_config.low_security — default_pool.use_tls.tls_config.low_security / 4e27073f9bb9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-1cfaaf4205a521fe31e8de72ca608bc5b2a75542f264d35aecfaf0c20387c5b5"></a>

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

<a id="canonical-9faa051a8e762f507839d4011e0aaa42d42b42ad4f313a574e933d9709af458f"></a>

## Direct properties — default_pool.use_tls.tls_config.low_security / 4e27073f9bb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ea514e37d80986bb52983410073de349efcc9f5182dfe1928dd7dde548deb1e"></a>

## Next pages — default_pool.use_tls.tls_config.low_security / 4e27073f9bb9 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2a650cfe0bed8fbe3c0456406420d31ca836d6d818fc23c1a9e8fdfa1d0f0770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c50510759dc4c06d9b7614c2bc216173aa545df4dc74f56b6bb422a9b066963"></a>

## default_pool.use_tls.tls_config.medium_security — default_pool.use_tls.tls_config.medium_security / 7473d3b40d49 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-40cb3b4afa107a751bc204b7a2a03626d5d6addeaa68cac67621ca51d6bb79e3"></a>

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

<a id="canonical-d42b776a78edd2b41eed9a734c183b7a2a51eb64bec857b0eef01fbf066aaded"></a>

## Direct properties — default_pool.use_tls.tls_config.medium_security / 7473d3b40d49 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3dd0fc497642c18d43b7157ac39b6a938d36c5ee0eadc96cb78ed1fec9d5ddfe"></a>

## Next pages — default_pool.use_tls.tls_config.medium_security / 7473d3b40d49 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a219fecc30a63d0bec0125952244ad9dd6f6eb848785eec222ba376f7eb9bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75bd1cc2775d73912cdc86ecfbb429f02431e3dfc424db7a7b5f6705e9be43a5"></a>

## default_pool.use_tls.use_host_header_as_sni — default_pool.use_tls.use_host_header_as_sni / 602d2f94f335 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-37346b65427b75e3197def84fb3f37c1b8ad45732d7057d8753604caf14e10b2"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_host_header_as_sni = {}
```

<a id="canonical-d09d8b5a31d7c321b4ef656f7340aa2c4942ef8c5860abea1fce4ba2a8f6773d"></a>

## Direct properties — default_pool.use_tls.use_host_header_as_sni / 602d2f94f335 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2edb997b6b0f9d8fdb9371ecc97d7882b66accd72d93526f37cc3608e9d4634d"></a>

## Next pages — default_pool.use_tls.use_host_header_as_sni / 602d2f94f335 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f26a396899fbb96912dc6a510f7f98564305061c40428927b74f28859f4cce14"></a>

## default_pool.use_tls.use_mtls — default_pool.use_tls.use_mtls / 767b2ed92832 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.use_mtls

<a id="canonical-063fe4d2805a59e9a26c594a56966c27fb169c0563fd7cd7d20e910071fee2b4"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-ffccc19bfcc97a27642ce22bd0dc185b5244e331addad8ec89cbbec7bf6ab1a4"></a>

## Direct properties — default_pool.use_tls.use_mtls / 767b2ed92832 / 3

- [tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b): complete subsection reference.

<a id="canonical-56b94cd0bedf8f7f7a60c88225ec2849bee85e0696b72e4e12f996507581448c"></a>

## Next pages — default_pool.use_tls.use_mtls / 767b2ed92832 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faf8632791e44ec53bb04c0f35d8e20c684c1149cb14c9ca8791fe0cd676b40b"></a>

## default_pool.use_tls.use_mtls.tls_certificates — default_pool.use_tls.use_mtls.tls_certificates / e717bb66e71a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-a67dcddd26d94d934197b5f71a358b8489d818bdd5e3823aa5f7935a56c4056a"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
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

<a id="canonical-132a1662727b313d0bee9f693aafc9fdc570b7e21761a71f2f1f2dbeff241a4c"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates / e717bb66e71a / 3

<a id="canonical-0881bd6ed24e2ffd026c9663cd62fc7290eeee108838ac71d09e6120e56118d7"></a>

<a id="canonical-8f12c592dc9d058a288c33a7342cbe4b13d587dcb03f935af3e14b65304a7eb7"></a>

## certificate_url property — default_pool.use_tls.use_mtls.tls_certificates / e717bb66e71a / 4

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

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-017.md#canonical-889d7c28bf52fc63f0a9c32d9cceee0c8e5569a1b0abb1e214a8e0227855654d): complete subsection reference.

<a id="canonical-a94eb108d0eca1acce73817a97b0d97d6aa1ed6b2541633c329917c421748c1b"></a>

<a id="canonical-ab37770b0d53c323a64e222a5ecad323a917802f4174de67a91e73300c723fb6"></a>

## description_spec property — default_pool.use_tls.use_mtls.tls_certificates / e717bb66e71a / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-017.md#canonical-7afc7202f8478ebc6c12b91c3e06065f08e1b3240638bc290d46a259a496edee): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-5633d1016bce337c3f1149f9e273bca626a75dd24c8ff7c176ba2ea12bcc173c): complete subsection reference.

<a id="canonical-844ca1f8dc69680aeb869cc2b3ed59810e8639992f2dd0b0ac5ae623938e92db"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates / e717bb66e71a / 6

- [default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--http_loadbalancer--reference--group-017.md#canonical-889d7c28bf52fc63f0a9c32d9cceee0c8e5569a1b0abb1e214a8e0227855654d)
- [default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--http_loadbalancer--reference--group-017.md#canonical-7afc7202f8478ebc6c12b91c3e06065f08e1b3240638bc290d46a259a496edee)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb)
- [default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-5633d1016bce337c3f1149f9e273bca626a75dd24c8ff7c176ba2ea12bcc173c)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-889d7c28bf52fc63f0a9c32d9cceee0c8e5569a1b0abb1e214a8e0227855654d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab34d94f38d592127417b8e9a5b38c4e7e29f94adf11f948790a27d482c6c653"></a>

## default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3ebbd4eacf53 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-4c2471af5e2d32033a3749b252ede80d7cb4fbb351b62fe320660afe3bfdc2ab"></a>

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

<a id="canonical-0f9e3c73032c30c72be26890eee4eba68a65f2bef3afa2b5f24c350245e0540b"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3ebbd4eacf53 / 3

<a id="canonical-fd08d3918173fad132277640cfef7c54afdea2973449dfabc4d513979ff93c1d"></a>

<a id="canonical-54a6c86025b07ebadee87f4e24adc20d4e84f05489ac8102d08e3b9ccfdd2bcf"></a>

## hash_algorithms property — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3ebbd4eacf53 / 4

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

<a id="canonical-265f826808b2996503608c05b4da0006eea29b831bae859fa4d5e7a3ea1407cd"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3ebbd4eacf53 / 5

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7afc7202f8478ebc6c12b91c3e06065f08e1b3240638bc290d46a259a496edee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2dbc50e9be5af72c448eb4d0b7016870ef205e7528f35b0f266af92d46272f1"></a>

## default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 962c26fbca40 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-703898940534edc6309863a6ffc96b813d88df43fc11d8f361017930e50534fd"></a>

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

<a id="canonical-02b87e1b8f913e7c8eacdc0eb85cc6c09773789e68eccd7878b27c9e76647f9f"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 962c26fbca40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca3b23718cdd331b215be9ed896e1f611feb4c4fa59311e04b09c04196b0eeba"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 962c26fbca40 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00224f5a5e1390739dcc7b5cb631272697c480fdbda30cee8b540f455a9ab861"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key — default_pool.use_tls.use_mtls.tls_certificates.private_key / c3133157afbd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-a392df95cd7b679b5ccaf7f3145dd3deb4fa0d783c5c23cec08b07bed5738e37"></a>

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

<a id="canonical-c633055623a2be8ac3862430008965419b32f593e7b78a4abc464c423e9bd9c3"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key / c3133157afbd / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-e2766a0af4f6a0ff1e4f3477da9c0030be1f13c6e70d79c2fa2e50b292c72dcf): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-f3056c94ab28b0a62fbd29518536fae1aa9e5868997b04eac261916afb207012): complete subsection reference.

<a id="canonical-8914f2c36745ab09fca535feecef720c226b9e224ba45688bcf9b2de46c15397"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key / c3133157afbd / 4

- [default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-e2766a0af4f6a0ff1e4f3477da9c0030be1f13c6e70d79c2fa2e50b292c72dcf)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-f3056c94ab28b0a62fbd29518536fae1aa9e5868997b04eac261916afb207012)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e2766a0af4f6a0ff1e4f3477da9c0030be1f13c6e70d79c2fa2e50b292c72dcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca6c778e658066d1e37d0d577c927b8a5dd8e7c9ec0d78778b50b2985daaded2"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0a2e42e294d3f0f846b12aa4af5edd411f28d395e112a4054e1154d0c14436e9"></a>

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

<a id="canonical-ed59065e8c8212708331b7fa8383cefb83daba2813597e6bb66560343f636a33"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 3

<a id="canonical-bf5a7e7581cb5c98edcc387acad42ef8cc92be0e41f86db3162cd06905829eb7"></a>

<a id="canonical-0e35da731e9031a2ff4668f730cbaf481caa3ebf1c488a9fb340efde512a4ae4"></a>

## decryption_provider property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 4

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

<a id="canonical-957988fe7570c31be71e61a18dfec235efc4bdbcd7f062339e05b3f097b7f207"></a>

<a id="canonical-a5c9f6a01fe505088fdf69c29f945de42dd002e7d21306b3ecc5cfa0649bc8aa"></a>

## location property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 5

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

<a id="canonical-20cc49c75e0125ac830bd275e73dd8a9f437285ca75506980a2e884d9587a5ee"></a>

<a id="canonical-a7748cc5cb30cd897f1c5ee82ea9c8eeedf9aa422702a3765062ddb2e8b4dd2e"></a>

## store_provider property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 6

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

<a id="canonical-eedaf8afe7c9ddd863f79f4d33ea03e55e41de5db41316e63509824b1a1e659b"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 0956d603eeb0 / 7

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f3056c94ab28b0a62fbd29518536fae1aa9e5868997b04eac261916afb207012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58b3db044ea6deadace221efdbeb6121430bbd117417ed1d73223d66833c2b01"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1d98ef790a8e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-f0019f4f8c103a7e0fb483bc8bac96fb4125932c4f052f6fb9483c0b0684e63c"></a>

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

<a id="canonical-8e60d30dd37c405b9d83164f634ed3a2837f30cf2e3223befef79c1ecaa698aa"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1d98ef790a8e / 3

<a id="canonical-20edf8bbc52755c573526f645646bf608de363ff29aa0ed9477096d110ee197f"></a>

<a id="canonical-3cca4724ff11a3385f1eab3273ae9aee7dafa7e91c14f37d7e4a15c49de5714d"></a>

## provider_ref property — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1d98ef790a8e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e2630afeca8404ef44b24cd58a9b239fdf60760f4e3692e87217d2336efdcfe5"></a>

<a id="canonical-e1349eeffa4ea75f60bd3e0d80fcda65951a58fd2798810fc20635a91f1dacaa"></a>

## url property — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1d98ef790a8e / 5

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

<a id="canonical-636e66da0cbecff89184d791f62094f22260642771d4337f0b203d51283c5c1e"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1d98ef790a8e / 6

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-9d8a5c503d6619879fdb85cbcfdd1dc2fac921901ba94d2b3935e442c90a10eb)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5633d1016bce337c3f1149f9e273bca626a75dd24c8ff7c176ba2ea12bcc173c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58345123e6851d5c5ad2896e8855aefc40f7f0f501d02128feec6d7b8ad74837"></a>

## default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / f30125797550 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-1f016e6c54a985ae3f2718dcfb234dbc0b640743524542ddfc6bfaac2b2746c1"></a>

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

<a id="canonical-f9fff4a974ae0d1eee268598011444398a6897dc8d3f84871b54c1103fff96b6"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / f30125797550 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46ae7d3a2e43fb5010c6d1623950b7eab5d14636eeaf58b78300d9fba830d667"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / f30125797550 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-a96e16d8b19784455d81069fd6836a3ee74e3fb94a00d34ae1f085355929c30b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-85d339096a8a94d45c3ee4ac5d310e910ea0d15e221b1ad1c7454382122a6ec6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d697cbe82ceff093e57ecf6cb4f37bcbe4efc61e8653c8a58f2d18e3862d11a"></a>

## default_pool.use_tls.use_mtls_obj — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-0a596715c52d251d3f34cd338c29e582dff1bacd5e8a100f1c3445ff0fcb9911"></a>

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-32881d1451944cc94e96be806e6768c3fef04d4539789f68037c7ded873dbafc"></a>

## Direct properties — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 3

<a id="canonical-1995348e3cdc1e9856beffb81c258ac123cb12de73c76a5e6965c3b594ad6ddc"></a>

<a id="canonical-d2e93e6e1a8d54eb92baef508ebf31718a3f3d2cc22c6d4dc55828a57187fe9f"></a>

## name property — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 4

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

<a id="canonical-1bf88f912e7d8b8c8002fe1e7852ae4f28218b79399e7041e5c42ef37c210e14"></a>

<a id="canonical-50d6c144c8c70eea0318ae6a9d08d70e0caf2197ccf0d25025ed69cfab85e109"></a>

## namespace property — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 5

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

<a id="canonical-887ef6c2ba1a83566056548e97732ca08449ac80344b6a028c8a0aa4e90fc0fb"></a>

<a id="canonical-47714755259a6f5f956f29c4f238b8171bd3878fe2bb6fc46bc3bd7c262920f7"></a>

## tenant property — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 6

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

<a id="canonical-6fda14588de14fd2d19e1785f171591a58cacfd21f9eff3cfd4fe4cb7608cd02"></a>

## Next pages — default_pool.use_tls.use_mtls_obj / e7c2d4f75aee / 7

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f26b369f3bbe3d3bc658c0b08395510e8df06b4445cd510817ab56c342ad0a0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0b4343c1f51a03e6acbebb19fd547a00965b99b5bc3a194eff36848ddf374d"></a>

## default_pool.use_tls.use_server_verification — default_pool.use_tls.use_server_verification / 00faf457c49b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.use_server_verification

<a id="canonical-9c4c8bb24d6409fd60555fb85b3e67996e09147fe5533c7de47f36b1d2f4b3fe"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9f123c99ac4d5f63518daac99b38d4c6aa9fb2c705f169f356a4ea6a3ebeff5"></a>

## Direct properties — default_pool.use_tls.use_server_verification / 00faf457c49b / 3

- [trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-ed460d2f4b6ce7ccaaae7633f2d8e43111b7c0ff6c32c0d0ad0e47b1863d1700): complete subsection reference.

<a id="canonical-f26610ab5ede11985a246762b94a96adb5b4fa744f333d59b75e8213acad6ffd"></a>

<a id="canonical-ad3d15aabaab16a4030639edfa74d43a2e1ae74795466150628e47f5e5eb3079"></a>

## trusted_ca_url property — default_pool.use_tls.use_server_verification / 00faf457c49b / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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

<a id="canonical-446e9bca86c58ca15a70965b61de35b43e58f14e2c28ce2cdcc1c52e601b5ed3"></a>

## Next pages — default_pool.use_tls.use_server_verification / 00faf457c49b / 5

- [default_pool.use_tls.use_server_verification.trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-ed460d2f4b6ce7ccaaae7633f2d8e43111b7c0ff6c32c0d0ad0e47b1863d1700)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ed460d2f4b6ce7ccaaae7633f2d8e43111b7c0ff6c32c0d0ad0e47b1863d1700"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d05eae1c002ca61ecc6a4dc15e29cb90ca5c7c1c3f6f1d285bb4b46f2d12b2b"></a>

## default_pool.use_tls.use_server_verification.trusted_ca — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-f26b369f3bbe3d3bc658c0b08395510e8df06b4445cd510817ab56c342ad0a0c)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-6d807ad7f77018e9300f7cd1b92128ffc2c652f115b38c13551386d9ffab2882"></a>

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

<a id="canonical-6ae7e857528ecf942f4e0a0201e6cca94ebe6fbbc9fa0ebc4092a532e0d8f3ce"></a>

## Direct properties — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 3

<a id="canonical-e793822d037bbc60ee9eebbbf26f68748e9ac7827fe2f049773e36f7e3dbff12"></a>

<a id="canonical-1f42317a4e1fb65e86abda9d7e5cc9a3505ad9c0c80fa8cdab81dd4adaee5112"></a>

## name property — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 4

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

<a id="canonical-e2485cea0a249620f716ffd80a7c8b3ecb056507db7609f325e4b72f7deb9b44"></a>

<a id="canonical-0e0eed8f6dd7e77e8b92bf0f2676117d427d8a8f06e3e1848dd80406cea62f27"></a>

## namespace property — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 5

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

<a id="canonical-659ce85b51525e35c21c25488b845698da2b6fb72d1d75a63ec5f879a3e8d246"></a>

<a id="canonical-a21bb91a2fb54bbedb0b9cfc6431e4b700725fa89effe9684ddb49f349bbe2f2"></a>

## tenant property — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 6

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

<a id="canonical-1c6adc0d06e3aca77e0ba6c47409bcb4f5714f9af8cf2f5c10614f3d05989751"></a>

## Next pages — default_pool.use_tls.use_server_verification.trusted_ca / 0ce597d670a3 / 7

- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-f26b369f3bbe3d3bc658c0b08395510e8df06b4445cd510817ab56c342ad0a0c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f913ba70e9bad1bf406bf57ac712f432b527b62659c57f4aea342e0ad11610cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc6e67edf89a422684447fb041571d9fa84defa4e3c36f41dab33a2684830bbf"></a>

## default_pool.use_tls.volterra_trusted_ca — default_pool.use_tls.volterra_trusted_ca / 97ece1a7f5f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-5b752b4df30b55faabb5a1836933ef65eb56833c4ed31114edbd15c094192ff7"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
volterra_trusted_ca = {}
```

<a id="canonical-bae23dc54569cc45c7e136ee3a4c9789dbadd3124b6e679f5bfd12ab7973e753"></a>

## Direct properties — default_pool.use_tls.volterra_trusted_ca / 97ece1a7f5f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-040c7e83ed0210668401b8658c3df8192921485dd2136afd345dd93b590bfe2d"></a>

## Next pages — default_pool.use_tls.volterra_trusted_ca / 97ece1a7f5f4 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2437a25c846b3d84ee97df7a1eb4205a5507408f60577babcbd742049657433e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fecaabc24a87e6315c260f0fb96b2599b6dd6cf281c4348e4e02695ccdcc0662"></a>

## default_pool.view_internal — default_pool.view_internal / 11747303ba54 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- default_pool.view_internal

<a id="canonical-d5d7005169ad626129101d59ba7f51d72210125e33823157920e8a84a7ad700c"></a>

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
view_internal {
  # Configure direct properties listed below.
}
```

<a id="canonical-391ac636ccf514ba27a861bbd955df8f252c56085107203acc7a18a61e2ffd9f"></a>

## Direct properties — default_pool.view_internal / 11747303ba54 / 3

<a id="canonical-bc8497b8a24f23e96b53132664706b27e479260241ed308ca83d37eb9d574ebc"></a>

<a id="canonical-9faea58c785b52167a29db1869f2dd2b7c13c9325f602931d8b346dd94343662"></a>

## name property — default_pool.view_internal / 11747303ba54 / 4

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

<a id="canonical-491cb361c8f70da44b9fd59b0ed063f326ca37e41b81471d81696b114ba202ea"></a>

<a id="canonical-d7a8ba2f1317e4cdefd846305c6104e87eb477e433a9f0434b053d919fe59418"></a>

## namespace property — default_pool.view_internal / 11747303ba54 / 5

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

<a id="canonical-f09ab20d629852808d93afdc8f65e240ec7ae6d0a5a31efa395d1a465c9f6628"></a>

<a id="canonical-5a0f736743f435155873ff5cf8502f74ade4b00b5f7228580e8e8359b1dfb633"></a>

## tenant property — default_pool.view_internal / 11747303ba54 / 6

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

<a id="canonical-10436e1395c15a4f58b02f04139f80d44778eac48d6dbb2ce31f7adc89f463be"></a>

## Next pages — default_pool.view_internal / 11747303ba54 / 7

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffac39bf8e507ffe1e8bc9ca094a0a805a6bb73b96cc7fbe019e239d2499166a"></a>

## default_pool_list — default_pool_list / 21f27efb52d6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- default_pool_list

<a id="canonical-1c3b80d6c7b7a785845d51a8f61cf52e21c31fc4a0b1e1b79a3d23ffd159ecbc"></a>

Type: `"object"`. single nested block, Optional.

Origin Pool List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
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
default_pool_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3f49d7b86f00cd69349dc6b4e07716944bf82512b6defc8ed52a0d17486a57d"></a>

## Direct properties — default_pool_list / 21f27efb52d6 / 3

- [pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456): complete subsection reference.

<a id="canonical-c4a79d54c6d621789a9a5bbd0764c741ee3cf4078891968dc83ac5cec442c1fe"></a>

## Next pages — default_pool_list / 21f27efb52d6 / 4

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8af8f79e915bc3bf20003f8d77153703b3eaa3b79d06c7ded3ca629e838c11e"></a>

## default_pool_list.pools — default_pool_list.pools / 913b0da2d4ff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675)
- default_pool_list.pools

<a id="canonical-5c7daccfd74fe9b2ab542e9af60cb187793d4fdca7bdb93abc21490f1398e91b"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-178f2bd64168dccab7e4be384edc96d6862b620f7458818eeee529d90169ba84"></a>

## Direct properties — default_pool_list.pools / 913b0da2d4ff / 3

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-2394f1a23c354bcbc36e88ee3e1ef2e5e8daefe19d5d8cb771abed4d0587ac1d): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-e4fb4e5c17aa7041971e9ebce645b3d2c10ac0b39e8bfeb06185c0a6586d17fb): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-c767841b2c58bdaab7af0a0116d093c04167b7f2962a45e9774f41b2e6d0ce5c): complete subsection reference.

<a id="canonical-e81d8f94bcf90863e11a37f3b5718d9e6f73f05c584c3cd458896d611effcabc"></a>

<a id="canonical-f07dfb0c5103ca979e38d962d47cfd8098b8c0ae21db07f9b38403816db06e96"></a>

## priority property — default_pool_list.pools / 913b0da2d4ff / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-271d4ddf032a3327cf004e0b97b48dd708f1feece36ea43acd12b42808ee205e"></a>

<a id="canonical-02a686e4d3ca5ab6b7ee8aba005305d4dc4dbca3dab36ad45d09a605cae9c10b"></a>

## weight property — default_pool_list.pools / 913b0da2d4ff / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-171c5147dbe4aaf3ef7624533b88982522b3c753f19cc437a1e5201e6d32cf4b"></a>

## Next pages — default_pool_list.pools / 913b0da2d4ff / 6

- [default_pool_list.pools.cluster](resources--http_loadbalancer--reference--group-017.md#canonical-2394f1a23c354bcbc36e88ee3e1ef2e5e8daefe19d5d8cb771abed4d0587ac1d)
- [default_pool_list.pools.endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-e4fb4e5c17aa7041971e9ebce645b3d2c10ac0b39e8bfeb06185c0a6586d17fb)
- [default_pool_list.pools.pool](resources--http_loadbalancer--reference--group-017.md#canonical-c767841b2c58bdaab7af0a0116d093c04167b7f2962a45e9774f41b2e6d0ce5c)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2394f1a23c354bcbc36e88ee3e1ef2e5e8daefe19d5d8cb771abed4d0587ac1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cec2b607195a49c94d7afabda09794b112deb3b08a12070dbaa1e93212435576"></a>

## default_pool_list.pools.cluster — default_pool_list.pools.cluster / 2dd161097656 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- default_pool_list.pools.cluster

<a id="canonical-8a381c24105f246aea39bdb69387c9189139245488c1eb7f5d91ac5278e878ed"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-04f40142abceb3fe6ffb5397bbed7e3173542575b1cb01baa98e6ced80f54226"></a>

## Direct properties — default_pool_list.pools.cluster / 2dd161097656 / 3

<a id="canonical-3569ab8301f26b4e3b327fb906a65c9a833c1324cbd23aa259d0047a02552ba7"></a>

<a id="canonical-74c8fe2e09708a7eaaf60ff13b3eff9a9000c6cf4619485953bdb3daf89491bd"></a>

## name property — default_pool_list.pools.cluster / 2dd161097656 / 4

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

<a id="canonical-0de0df9acc11fdcd0ff6106bbc22bcefe5b8623350c9ff52a98121eb7f8ab129"></a>

<a id="canonical-d122357940dcbb99f8edea54292d0f6b8c1f47ad0bf3b3fef6bc9a538accc46d"></a>

## namespace property — default_pool_list.pools.cluster / 2dd161097656 / 5

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

<a id="canonical-a37ac4e0cc05f24737fb8f3ae975265ea3752bdde51c6fab8455b7c63532ff89"></a>

<a id="canonical-bc65faab5af2e6b2dc9f8d6fc03e174623873dcec05e203f7299cef17de6de66"></a>

## tenant property — default_pool_list.pools.cluster / 2dd161097656 / 6

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

<a id="canonical-f762ee8122512ba1046599dc2f35238dd00293fc8fb65f095716cb2b2c5ce36d"></a>

## Next pages — default_pool_list.pools.cluster / 2dd161097656 / 7

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e4fb4e5c17aa7041971e9ebce645b3d2c10ac0b39e8bfeb06185c0a6586d17fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59b03147b2071ae29180b2fa7c60981e1962425705c768cff943adc3767cd62d"></a>

## default_pool_list.pools.endpoint_subsets — default_pool_list.pools.endpoint_subsets / 92ab86d12e08 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-7b10e6db074efcc26bb6af1d55ad8a15be5f5585e29409a8f99f62c932560ada"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-fba317ef112321396605ec2f1de4720418332f8e2562e12399ed1e04da16fee7"></a>

## Direct properties — default_pool_list.pools.endpoint_subsets / 92ab86d12e08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60a411436033e68d0c505ba570596fb38fef5dfbae1613038f0853c3158afa53"></a>

## Next pages — default_pool_list.pools.endpoint_subsets / 92ab86d12e08 / 4

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c767841b2c58bdaab7af0a0116d093c04167b7f2962a45e9774f41b2e6d0ce5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b26e2d95f131d20be2bddd80aaa510187ab1f18b423df23278345e2e31907e6d"></a>

## default_pool_list.pools.pool — default_pool_list.pools.pool / 621eb2e35daf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-2e2c6f117545e8507f9a557e78ac7c1395c119a6fefa3b39e4e67a70cc292675)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- default_pool_list.pools.pool

<a id="canonical-79bf30639570053a25aa86a59b2d978e4ee06e681dce617e51c4b587031f0538"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f03770d264908e3d82cb5c58b64e7c6ca26a3a8ff20c0b518aa9f2b5eb8d105"></a>

## Direct properties — default_pool_list.pools.pool / 621eb2e35daf / 3

<a id="canonical-53a8462d1082f502681d7f23ff91a00b69da2a66ce34386bf24c44c0eb07bce5"></a>

<a id="canonical-dcd33bd557380f7dfd4004d74267dc091b90f244cf41ad217379902309951e90"></a>

## name property — default_pool_list.pools.pool / 621eb2e35daf / 4

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

<a id="canonical-90bc2637e131e99b1ff6f47cca59e9a92dfa5359bb087525177718e83c15849f"></a>

<a id="canonical-0ac3e9760da51a68cf5bc321a4c9e3be42ec5da9ebc48915f41ba16efcfa8750"></a>

## namespace property — default_pool_list.pools.pool / 621eb2e35daf / 5

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

<a id="canonical-f46f73d6be33995c8ab047cbb7018b59be0a1af81731d70eb4b4e688eab388fa"></a>

<a id="canonical-9ea2051760be8e64149aedf2abeb030d1abcca432c21eaa4e7ca4102ff3540ce"></a>

## tenant property — default_pool_list.pools.pool / 621eb2e35daf / 6

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

<a id="canonical-0cc9bf598bb280cc43c4ad771e9fc82430aec7908204d88905c17391598c3572"></a>

## Next pages — default_pool_list.pools.pool / 621eb2e35daf / 7

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-c109fe2c98b6ac2a39483a41d46722bb285516a2fc57fb98e5255612e6d72456)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9597f97a85370936a062a5c9e9d71862786e96397075679046d0bcfb496f3e91"></a>

## default_route_pools — default_route_pools / 4f59b3a0d06f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- default_route_pools

<a id="canonical-8c0805dc479f068a3e7984c2d798275c056763653004d02f7f3d311dcaf59270"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools used when no route is specified (default route).

Upstream description:

Origin Pools used when no route is specified (default route)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_route_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-bb7a8decba78ece84f5cf043d9144ff32dfabca7455898dc9edf968b57a638b2"></a>

## Direct properties — default_route_pools / 4f59b3a0d06f / 3

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-4ad272eae48bf50fe957de790b260e62ed195406edce8910238f3c1aa252e4fe): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-9a2ed094868853075ed52fee02ee5b063906e6835f9e7554690e9e0d81c7f5b9): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-8be0ce5996480a221823344dd1e42757f28ff33d0d1fa1a9471980da1aaf6440): complete subsection reference.

<a id="canonical-3cde1cd85864dd11a4602f3c54b85899bc82db045c0fc104d8e9c771cf019a09"></a>

<a id="canonical-e07a73725537af7a87b7724b37da7c82a4a8fd6274019cb65c3e2b2cdddfc137"></a>

## priority property — default_route_pools / 4f59b3a0d06f / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-dff31785d480fad2c34887fb2a62d81dcf39d56ace83bf49904ee4590aba01a3"></a>

<a id="canonical-0930977645b8279f0b85ea1e3bfb3b76737f980ec90671b19491beeb256fa787"></a>

## weight property — default_route_pools / 4f59b3a0d06f / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-08ce494d44ef24418f98e32674c687fd4a0639a98404402af59d083873e49caa"></a>

## Next pages — default_route_pools / 4f59b3a0d06f / 6

- [default_route_pools.cluster](resources--http_loadbalancer--reference--group-017.md#canonical-4ad272eae48bf50fe957de790b260e62ed195406edce8910238f3c1aa252e4fe)
- [default_route_pools.endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-9a2ed094868853075ed52fee02ee5b063906e6835f9e7554690e9e0d81c7f5b9)
- [default_route_pools.pool](resources--http_loadbalancer--reference--group-017.md#canonical-8be0ce5996480a221823344dd1e42757f28ff33d0d1fa1a9471980da1aaf6440)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4ad272eae48bf50fe957de790b260e62ed195406edce8910238f3c1aa252e4fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca7ca7963ec29dc8c40bdebbe09bcdba39ac3a228989dc39b0284a6dbaad119c"></a>

## default_route_pools.cluster — default_route_pools.cluster / 5ca07c5db91f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- default_route_pools.cluster

<a id="canonical-82a429666b7618dfa84df53159a0aeaa8bc7849c9e6bf5a317566651b583ea26"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-81aaeeed9f90e22f0b11e604ecfeb4b94b5fa95af18591b7b54cb33d3ad61362"></a>

## Direct properties — default_route_pools.cluster / 5ca07c5db91f / 3

<a id="canonical-b9734be6c56117511ad798ba491aed3977f13b65a66760fa7d43dc84d5bb2777"></a>

<a id="canonical-dbaebc5e1804b6bd385a7178ad2abc1030ef5d6eca3a2276bcd89d4bdc17d212"></a>

## name property — default_route_pools.cluster / 5ca07c5db91f / 4

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

<a id="canonical-9e26002d234b6bd0c9dff5d6a2fa796bd1c2eeaa3efaa15a800c3efb333c1c6d"></a>

<a id="canonical-08651c03993f3f3264295c96f8190d5a256233830714430e134dd4497ef5e2fc"></a>

## namespace property — default_route_pools.cluster / 5ca07c5db91f / 5

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

<a id="canonical-9294ed215c3af483498d292b56fcbd1b5f191b3a7bd522de4b4b69fadf513689"></a>

<a id="canonical-fe691da211aee55fd62f0c9b2bc71cf6fccadbccf8fe686761e97628f0224da0"></a>

## tenant property — default_route_pools.cluster / 5ca07c5db91f / 6

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

<a id="canonical-34bde1c2f82074ca33dfe49e72278567d5bce2acd2c1e2ad324e97f835389a90"></a>

## Next pages — default_route_pools.cluster / 5ca07c5db91f / 7

- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9a2ed094868853075ed52fee02ee5b063906e6835f9e7554690e9e0d81c7f5b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df21cdfdb2b34e62eba82c8a9681166150af79bc45629fd8d4e6669a9bae1642"></a>

## default_route_pools.endpoint_subsets — default_route_pools.endpoint_subsets / 02f35bd22506 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- default_route_pools.endpoint_subsets

<a id="canonical-1663718e01cea0e92585385c9a6250053b530762274e442c58b6b63652c4d15f"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-09d587b32ec2cc5c58f673c2ed2b2c3bb961d69f57150889e59a2408dd527ddd"></a>

## Direct properties — default_route_pools.endpoint_subsets / 02f35bd22506 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-327885d01ee5405c9ae030c6151a578487d8787d7bf78a626676a413af9aee78"></a>

## Next pages — default_route_pools.endpoint_subsets / 02f35bd22506 / 4

- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8be0ce5996480a221823344dd1e42757f28ff33d0d1fa1a9471980da1aaf6440"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d38e953ceeab833f4e7a8d457aaadb35d40d7e5efcfdf7fe6161e8df894e0089"></a>

## default_route_pools.pool — default_route_pools.pool / 3c7cd35a7f6b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- default_route_pools.pool

<a id="canonical-63389a9597293f49707775566e89c0cb14f3691d1f7b6e943fb19402918473c5"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-e4c6f8dee8869c79825552af07dfe6c0874f68b69ab762c81ff94cae01e100f6"></a>

## Direct properties — default_route_pools.pool / 3c7cd35a7f6b / 3

<a id="canonical-da6dc026e28bef9006d00a842d00bcd8542c13fbe92f44386ad437f0e182d07f"></a>

<a id="canonical-7f821414eb2e25bc4557a9171173148363e9695abbeeb9cceac5efd249fc2f42"></a>

## name property — default_route_pools.pool / 3c7cd35a7f6b / 4

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

<a id="canonical-75560727159a259bc364bc99684ec5f599094c4778a8e3b55b574826092ebb66"></a>

<a id="canonical-74f49c9889549d4678a27d909f5f58238770348b92a7f2a5201c38242dd29e09"></a>

## namespace property — default_route_pools.pool / 3c7cd35a7f6b / 5

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

<a id="canonical-7afb8a8febb29bc8162653d0908e9130d36a56d4f66426fa6a75b85ea70c196f"></a>

<a id="canonical-54c696cc712a085416ba54387a555fd2fedd03c9cc28400ca4562ca83a302c22"></a>

## tenant property — default_route_pools.pool / 3c7cd35a7f6b / 6

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

<a id="canonical-66b5448fab4fb899cf09303177610b352bf3ee815a611f6d6e317d7761c57427"></a>

## Next pages — default_route_pools.pool / 3c7cd35a7f6b / 7

- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-2b8c60fa9dceb82aef881e64259e895cadc5174aabddb1c7bc55794ad8d2c41f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-865cc01ec6425b69dc0f336f2b1df4508cba7e06e1863d551ccd67564a9f8773"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a38d8334b5ff27e6d3a06fbaf5cddcfb9bd4dac67f51257c83de6de25f2a623"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / b0265f58cccc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- default_sensitive_data_policy

<a id="canonical-8f29b9dbbb5cd046a12490ade9de501c6246f62b8825494d93259c9081a20257"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-017.md#canonical-8f29b9dbbb5cd046a12490ade9de501c6246f62b8825494d93259c9081a20257)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-026.md#canonical-bde87177d5899e4b570b68001e43fcceb5df0837c3f0b2e71fe8eeebe618be2a)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

<a id="canonical-7a810b6e16a292e44f251904a0a5cc2bb9a4749f97d22f07b818065080dbdd13"></a>

## Direct properties — default_sensitive_data_policy / b0265f58cccc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8979cbe1415b4c16ecad24033f012b86fbdfaae9e629cd36204b0862ab1e217"></a>

## Next pages — default_sensitive_data_policy / b0265f58cccc / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1b55f3059b0d40854d39c4144a9207fcf3f11031231709c548cd57f59eedb026"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed627141cc501796c3438a169cd68554cd9db0a015aa1eb5ae599fd499b553c9"></a>

## disable_api_definition — disable_api_definition / ac1b3c207072 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_api_definition

<a id="canonical-199baa9cd32513f95b003364d3600c1e425b2483c83890657c8fe21f70f09bf7"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_api_definition = {}
```

<a id="canonical-192419b093c5e6c600af1844215e9c17c337d73abf0a323ce7ac3a189c672df5"></a>

## Direct properties — disable_api_definition / ac1b3c207072 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f038f52fa5b1fb3bc7a3c6546c5876d9fc523111c66708196265c278f6f9cb1e"></a>

## Next pages — disable_api_definition / ac1b3c207072 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2e06259edaa0293b9b5a461b956cd0a71e61ff8a95f34885a81c11c6239f00ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26a856a26d595af3e2ba62f2da5d1dce7292e33324f2baff1eaf88cd39f7b7b8"></a>

## disable_api_discovery — disable_api_discovery / 59e49d4199cc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_api_discovery

<a id="canonical-c03239bd0d4e98462d8170c6c141efb7fd9be246fa31af22c119ca6901a8b9ea"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c03239bd0d4e98462d8170c6c141efb7fd9be246fa31af22c119ca6901a8b9ea)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-fe40879040cf7b0f20c6abe35a749e588fd8ac197bf41b3a7856b4a8e58bed24)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

<a id="canonical-d3b4a033c03bc6e6748708c5a6e74a95f01382ed167fa543e4bba53cc78e5df5"></a>

## Direct properties — disable_api_discovery / 59e49d4199cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94e6abcadf0dd5a7d5d8df48b744b87e620077e209c3b24d31ad983adec5256d"></a>

## Next pages — disable_api_discovery / 59e49d4199cc / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f789378c079cb5e3adaa3102def219654d096f04b8c8a2cb49e577b8a9d534a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a667b6b43a3c03c549c386ab55bdd46c46be8342e37dd64f7db8c5d285db58d9"></a>

## disable_api_testing — disable_api_testing / 145e004b8b95 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_api_testing

<a id="canonical-92fdffa0592551285acab640901a3e81c0a77f5c572709339b273f36a12809fc"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
disable_api_testing = {}
```

<a id="canonical-cf17aa566a84dc5d56dd7b92fde3812927b85f9527ce0d0c78cff14209d4678e"></a>

## Direct properties — disable_api_testing / 145e004b8b95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c11873a6f76f71eb3911ccf7891b68bdb2fceb3574bee3cee85bf60347163e3"></a>

## Next pages — disable_api_testing / 145e004b8b95 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-db0f00840fedd3d4d10fca6b6e6a38c90e6d3c86f97d953ed4754b50821630fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dbc08e2bf84fb93168b1753b9717ebbe651d2d89cc0add2e6ca6144295d6e33"></a>

## disable_bot_defense — disable_bot_defense / 684d993c95fe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_bot_defense

<a id="canonical-0dea985b2422bf8166d9f9cfcc2f6abf7e8192e24f50ef1844e9974ffd7ea15f"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable bot defense. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
disable_bot_defense = {}
```

<a id="canonical-2cd9549a55b861d4fc5a4499707b3c45e7e2dee39171f837304f0ebd156a0795"></a>

## Direct properties — disable_bot_defense / 684d993c95fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b14b9fba753bbdbf2be7df7b520e135ea99162e2d51ef9a436872999e405dbc2"></a>

## Next pages — disable_bot_defense / 684d993c95fe / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b944ef368907204f33191c57e912db818db4f0d91232de5993ce8c168e2f8289"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7df64c14bde0c5ceb8c57df8b862b74dfaf3649de14e72ba7dc9ae03faf03beb"></a>

## disable_caching — disable_caching / b48320d98736 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_caching

<a id="canonical-b413cccae648cb2bf5e5fe30f01c130252244e7839d08a39058ea4f65a9b067a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable caching.

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
disable_caching = {}
```

<a id="canonical-4ae95b68a3ee72ff1e420f86328197f38a51ce52e81f4a8f2c050c04a9648f1a"></a>

## Direct properties — disable_caching / b48320d98736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d175fd77e0b25a917d81fb2d5797a4b5a1e3a12ec1dbd1b8c02cfbd1a322efd"></a>

## Next pages — disable_caching / b48320d98736 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9ae7b6d81df251cf2b47e0945f1a03143010e544df39ca300e07b415f04488f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25443af2a560822ca19bb2784ca6d2c8f8d44854b5503372b1da3f1d08c7f338"></a>

## disable_client_side_defense — disable_client_side_defense / 5ba621f6ef46 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_client_side_defense

<a id="canonical-7701f8ceb4f61af7f586280ab35cd8654495e369b75bf9f3819499706530bef2"></a>

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
disable_client_side_defense = {}
```

<a id="canonical-6f9ac846fb1d8e6962909601445a0e3ef5797c072f0476cef95ae117c8c2f3d2"></a>

## Direct properties — disable_client_side_defense / 5ba621f6ef46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4d734d34b6acc9c9c496221d18bb2d617e69d3dc234b3120b4985a4fab2fb21"></a>

## Next pages — disable_client_side_defense / 5ba621f6ef46 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-626bc9a9e1f5a8dbcff2cfc47f51f44452f307d56b7e4446911d766b3cdaaa42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-707dd2db00ba21259b303b4707ba37bde10285c3f915976973e86076c402e5cb"></a>

## disable_ip_reputation — disable_ip_reputation / c52422582c6f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_ip_reputation

<a id="canonical-7046ee33d8cb3357cb1df0ecbfd3b2c779e4bd3b346c9189aec8b4a045623e70"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-017.md#canonical-7046ee33d8cb3357cb1df0ecbfd3b2c779e4bd3b346c9189aec8b4a045623e70)
- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-7d35a14fb275919af7cef8f9b7c0ef1bbd579d5a6905ae11cf7f28882c2e175e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

<a id="canonical-497e721b1e14a889d762742a083f4322d4742f4517a60a79393e0a879cefc0fa"></a>

## Direct properties — disable_ip_reputation / c52422582c6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a3c276854cf5d369e634074bc71daebaa4c848496cefbef661026e74fe56a79a"></a>

## Next pages — disable_ip_reputation / c52422582c6f / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c6b6bcee92fa30fbd299d474868a4374a98620ff5e7a8c91390225b379dfd71a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d519545020fb5e214e4a982ce394a014b13abbfbe8cbcaef48714bf6eea26da"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / 85e34257f6bd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_malicious_user_detection

<a id="canonical-1d46997755c0b0cc92fea227685d37bc94e74ce26cb8989fa01ede0ba48a426d"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-017.md#canonical-1d46997755c0b0cc92fea227685d37bc94e74ce26cb8989fa01ede0ba48a426d)
- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-607b3508bdae831af0f1ad3c8f873a93b57362c88115ae2f8fcd196f375107a3)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

<a id="canonical-7898fe56d057a597d01ba484c6c3a2d55522113344c64ce8270fa16bbadbf79d"></a>

## Direct properties — disable_malicious_user_detection / 85e34257f6bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec20c8bd17b43b7158c8849f2299606a45ef5666f062c3e338a97e18dd44217a"></a>

## Next pages — disable_malicious_user_detection / 85e34257f6bd / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-73fd4dd9ed35740311b05e7c058fa47c122713897fcd26a1db5a9492c44f1311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3aaeab280395faee324c6027f6349a2c348408ac59aae648e5dac1d65e46db2"></a>

## disable_malware_protection — disable_malware_protection / 1fa910bcebfa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_malware_protection

<a id="canonical-92d1fabbf35b865552aea9361c156b6cb87d89f9785addcaf2d2774f8384ffba"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malware\_protection, malware\_protection\_settings; Default:
disable\_malware\_protection\] Configuration parameter for disable malware protection. Defaults to
\`map\[\]\`. Server applies default when omitted.

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

- [disable_malware_protection](resources--http_loadbalancer--reference--group-017.md#canonical-92d1fabbf35b865552aea9361c156b6cb87d89f9785addcaf2d2774f8384ffba)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-9d794a0e2d06dac24c5c9a6b68a3b7bc38c90d0c4ab591c9432720f0abd30c29)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malware_protection = {}
```

<a id="canonical-1ab0288cdb4ae51f3dc59b5888ab8342af74a9f51adbec723d070cc2d1ae45b7"></a>

## Direct properties — disable_malware_protection / 1fa910bcebfa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-973b26e751bab5241db15584b32b75e90217bebed25d43862876774560b8aaf6"></a>

## Next pages — disable_malware_protection / 1fa910bcebfa / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c48e1d3b25523022258a53db20abda4b84439f175704d24090277c6e2ae78317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e910420079dc9a65b4d97d060511620e3bc1806560b61adb4f7c76122ed4fd24"></a>

## disable_rate_limit — disable_rate_limit / 1dc3bc94431e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_rate_limit

<a id="canonical-47adf84a77fa93ae7d8206d6c0c59c560f0728d1b6ff16f895e03e19b5d03eca"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable rate limit. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_rate_limit = {}
```

<a id="canonical-583c6f4e71e6128a5f2176f00814b4c6f178c26f5614293ed1b478e64db214d6"></a>

## Direct properties — disable_rate_limit / 1dc3bc94431e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9465f9c5c6baab1294e7c59cfbedf564243aeae30befd06732ee53798123bb28"></a>

## Next pages — disable_rate_limit / 1dc3bc94431e / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a144aa0b53927f9806853d0586a2b3c5b599fac28431d3cf714604a41e03584b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e31270fb1c1675cf4b038869d54f7cd5f3bb48f3ac5a52db004ed7b712c5c888"></a>

## disable_threat_mesh — disable_threat_mesh / 7f317860d1b1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_threat_mesh

<a id="canonical-88d24e407817001dd90ec7f5d8f98c1a3b32fc725e7cb490cd57d5ef2d4b5bbe"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-017.md#canonical-88d24e407817001dd90ec7f5d8f98c1a3b32fc725e7cb490cd57d5ef2d4b5bbe)
- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-9028820520dd7f20c408a5367458bb5254abd65af3b8c3287c02692c0cce4864)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

<a id="canonical-60f991d77e15b4830c9bdc061e7afae89ece2cb29539d07847536a6af798b009"></a>

## Direct properties — disable_threat_mesh / 7f317860d1b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31fecb6973d4933b4c59ca97b54039ef9b9a93778a9186b27220b2a2cbb97c89"></a>

## Next pages — disable_threat_mesh / 7f317860d1b1 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8472f96112920b65bd27a1c5e159bbf163ddeb4a5ca3802a1d9872ad3c0607ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2955979343de9e953b9f896711e4d8492bd7792f350eee81c7b340ca271bbf1"></a>

## disable_trust_client_ip_headers — disable_trust_client_ip_headers / 20747f1748d8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_trust_client_ip_headers

<a id="canonical-4b0997b4f8f14478b389eb1774e1154b0669ccf69c40f71f7b1aeb0763928225"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-017.md#canonical-4b0997b4f8f14478b389eb1774e1154b0669ccf69c40f71f7b1aeb0763928225)
- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-fd61051873fceb8231a37e8be01c074bfaeea1bfd44309d8f68dff8951e1ac19)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_trust_client_ip_headers = {}
```

<a id="canonical-cb8bdb4268f11d4b12b50f40cc19b7c414ecfd220f6c17d82557c4cad8b114d9"></a>

## Direct properties — disable_trust_client_ip_headers / 20747f1748d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6aeac6050f49fd42e81f95161643ee80c1025d1aee5a4ccc48ac0a0a8747b00"></a>

## Next pages — disable_trust_client_ip_headers / 20747f1748d8 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2d3dceada3af713c60e3a5c97f2df4fffd6ce15af8979fc63fd04464363b4849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a19589346da23f5c9bd91fdbb979338573b3c06da533f4345bef3eed772ea64d"></a>

## disable_waf — disable_waf / 4e4d0f9a2af6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- disable_waf

<a id="canonical-09d5ef0dc74c89796b7c2991846a1462d1dfd53787270c8da7a0fde3c05fe341"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable waf. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_waf = {}
```

<a id="canonical-69ebb53197e8a8584e03e6c8a7e46ea797250cbe51597f52adb2b906f78e1605"></a>

## Direct properties — disable_waf / 4e4d0f9a2af6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a0906733f1c0c39dd68b721618e189fdd85fd542885b88de5c1def018a461b6"></a>

## Next pages — disable_waf / 4e4d0f9a2af6 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b29f363d3a20a25f933614c6079ff33efdf1edf8965dfa2ff873b90ad86f557d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb644810e90ddec1b834566f0c052f1c6b3f86727070377685c79e22d5d3fb9f"></a>

## do_not_advertise — do_not_advertise / e052e1f93fc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- do_not_advertise

<a id="canonical-c9d7a37ec1718c7f11749c43b483812a1f36bd674178532d404c7763d40cd375"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-452fa4880492a963bba42056c45ae2db9b0fb866efb4c2e6914a16c9e56b9e27"></a>

## Direct properties — do_not_advertise / e052e1f93fc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7df768309331fa1df2ff80eceafba9f913a4e4400a400d3e6fd449b79f83bf35"></a>

## Next pages — do_not_advertise / e052e1f93fc8 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc09dcc10bcd2b53a67505f1d1e7d096f381d084ff83e0fac73fa7d0a0b5c616"></a>

## enable_api_discovery — enable_api_discovery / 2bbe7d71fd3f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- enable_api_discovery

<a id="canonical-fe40879040cf7b0f20c6abe35a749e588fd8ac197bf41b3a7856b4a8e58bed24"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-46348a10fef1d6df444a028e6e73c0c31751af85f297576c03af0fdeb6f47104"></a>

## Direct properties — enable_api_discovery / 2bbe7d71fd3f / 3

- [api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-001b95fe7fcb9807071346b2bc36a98014302211e6ed07d4435d791bff84e6e8): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-953ae11281fa4529ec902f19117d94823bf805fb4f08835530b9fa97f9129c29): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-f3d58fdbeec479613663bd12a57c81d9c2af56ed18a38ff598298ff0e1fdc477): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-e5e6ccc33d87a4887505b909d80117ad702076966af90990dd849b764a7fb68b): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3c3fef9df001bf6ec2a41b05448ae5c42a59901077d5b170d4f4d85f3157184c): complete subsection reference.

<a id="canonical-f22a140d1336851c632560b74bbd49deff433b5764ca90591a501d6c2313d7a9"></a>

## Next pages — enable_api_discovery / 2bbe7d71fd3f / 4

- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169)
- [enable_api_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-001b95fe7fcb9807071346b2bc36a98014302211e6ed07d4435d791bff84e6e8)
- [enable_api_discovery.default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-953ae11281fa4529ec902f19117d94823bf805fb4f08835530b9fa97f9129c29)
- [enable_api_discovery.disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-f3d58fdbeec479613663bd12a57c81d9c2af56ed18a38ff598298ff0e1fdc477)
- [enable_api_discovery.discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-e5e6ccc33d87a4887505b909d80117ad702076966af90990dd849b764a7fb68b)
- [enable_api_discovery.enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3c3fef9df001bf6ec2a41b05448ae5c42a59901077d5b170d4f4d85f3157184c)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649a12122834bdbe1a67ae909f692e5d8db29316debd4b5d3777f10b2bf4eeba"></a>

## enable_api_discovery.api_crawler — enable_api_discovery.api_crawler / 5cb29e394f1a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.api_crawler

<a id="canonical-f79d179a627a97307233df03354bf0350be97374335086e8bdaaba8e60bc09b8"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f318b66825660bfec6f054704ae9d347463c1b60517db207011e9df3a262eac"></a>

## Direct properties — enable_api_discovery.api_crawler / 5cb29e394f1a / 3

- [api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-9c6bc2080dd8970f5408fce5825dc385b6d9e93a3e2839ea37c26fad4f54eca4): complete subsection reference.

<a id="canonical-3476fbc48ad1647eafd7846d6b54a6b0c262cc27c61dc485f7849e7cec004ab1"></a>

## Next pages — enable_api_discovery.api_crawler / 5cb29e394f1a / 4

- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [enable_api_discovery.api_crawler.disable_api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-9c6bc2080dd8970f5408fce5825dc385b6d9e93a3e2839ea37c26fad4f54eca4)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93e596382b8c90780e3502d4614b5ed5a1883e848d132cf56a0d8136700a82c1"></a>

## enable_api_discovery.api_crawler.api_crawler_config — enable_api_discovery.api_crawler.api_crawler_config / 6f6f840aa1c4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-051c113b1287ea03f6eab924400d9498944f071f2a59e07ff737930ddb204549"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bc7236b88eaea4c100300c765b56ae77f1edd1358a908510b97c0d726e6dceb"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config / 6f6f840aa1c4 / 3

- [domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82): complete subsection reference.

<a id="canonical-9df81c26967ac99e87f041fc112afe03ce95cfa742bac7ffc079e103d60a7092"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config / 6f6f840aa1c4 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b98b72f4c355beb51e1df112ab7bf0fd25f9369f37ff114ac95f06f3eb6f4b1"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — enable_api_discovery.api_crawler.api_crawler_config.domains / 082c07664be2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-402520f9f68828b56e0b04e033b23df26a7b1535e31703b8ae51f3cad0a12219"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-534fe9042298a08936ec96038e55c41e203f62695cdd788adf5af78f65fb4435"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains / 082c07664be2 / 3

<a id="canonical-ed92bb7ebacdf931974b4b74443d62fdb4481c8314484ff445e9bfe3c89445f2"></a>

<a id="canonical-1a0697f24aae2e2c52eff6c88d781982a32fd8a8fedeb4189df886cfaff0777b"></a>

## domain property — enable_api_discovery.api_crawler.api_crawler_config.domains / 082c07664be2 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72): complete subsection reference.

<a id="canonical-1ad69522e732f84c0dd640bdfc52ed85ecad886debf5c8c7da2df1a5bd9edded"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains / 082c07664be2 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-704fa82bf078880e32ee4e0f31e6dd56fa051d72836f0591c8e01d31235d6eb6"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 479c633801d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-ad3195d63ff64fbe8fb95162e9f76d0e297dd94321b79fb5b37017993d4eab6b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3475ac762e149f23559b0c5d2edc004c97352153fa3c4e117e3ec792dfc296f"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 479c633801d9 / 3

- [password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609): complete subsection reference.

<a id="canonical-a5a55a5dcaa2712c995556e1f4a1ad6c6e64042b783844076f27d8c8c296276b"></a>

<a id="canonical-ce239a6f4652df72759b63f8d5c03bf0d283412edf27b267085b6d82cec7a07a"></a>

## user property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 479c633801d9 / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-cc531b420756618d279cc06e6bb2f6803742bbfe54263b497a43007e0daed77b"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 479c633801d9 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4dbe8d25b92c9bd0ec6d5d0d405ccbad1de6278203431f9fd6110f5027492637"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 5ea7d95629aa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-07a7661919483062341b9a58f84708ab0c5e68173ce461cc3dcab7e3614783ad"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fc17ac4cf6dad4a8ee77ac3991a5b04817dfe5e11d87bafbb38f513373b10c4"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 5ea7d95629aa / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-fbbc6f522407165759fac9659500088fb7ded09914b057bf36528571c67fc69f): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-6022a933f064a6e14af21bace655327b95fb93891ca8617a6da1b0f138492889): complete subsection reference.

<a id="canonical-a3d785cf23cc55b7d5399052cab935400856a2b90b4a5335ac0010b95407965e"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 5ea7d95629aa / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-fbbc6f522407165759fac9659500088fb7ded09914b057bf36528571c67fc69f)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-6022a933f064a6e14af21bace655327b95fb93891ca8617a6da1b0f138492889)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fbbc6f522407165759fac9659500088fb7ded09914b057bf36528571c67fc69f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81e4cde34503dba80cef24ec1ab5bb2841061dffd96ac2b5829695a6ba8e3379"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-5d3a78a6ef4518f7a1c1047afb5f66b5c96c3ae898a9a6610639d1b93581ea9e"></a>

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

<a id="canonical-c6cb5eb79840566019f2d9f4890cf0bf48b9c7f7049a7218927e622ec210bf4b"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 3

<a id="canonical-ea8b35f68397adc967ee49748994cd960d677800ccde2ddf07d9cc77e21f7044"></a>

<a id="canonical-512a63a57a020807be5a74bd871145cc681db59df26b8d8cc99638abad7ff645"></a>

## decryption_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 4

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

<a id="canonical-3371bff55a36087e82952d6879c4e9909c58b38520e567899d13dbb001cb5988"></a>

<a id="canonical-78cf836c54e8d21a5af97a3172cd7a0b0923fab9bdb135adae93af9edaca8de6"></a>

## location property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 5

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

<a id="canonical-f0dc0969622385a30ad9e5542618dfb087992b383ec37d70c5a6dab57193bb17"></a>

<a id="canonical-a4a61895b77238f6168955f0d68a0d14dd19cfe4fb74c902ce71537ddf19b520"></a>

## store_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 6

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

<a id="canonical-d62469c7f6c32c818256d683e32ca96a07fb265825e07b1e504a7423b4eb8f45"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 3998a028e5e4 / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6022a933f064a6e14af21bace655327b95fb93891ca8617a6da1b0f138492889"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2f9a03ab434b4fa639c9a85279b0ba69280893e0017655887c95cac1fc45e3b"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 47daa4cc8f47 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-5c401ce4b6ae58f9dd471af2381ac042a839e80758e2dbc50c39895be4a2a9f4)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-3dfee91e26b2b41df1182ea6c2621cdccfc73c1289cc3adde5499db7b64eee82)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-5af5f8ce581f3b81407ad5e84aba942390f2a997c9a6d67d203a0d2c3cdbae72)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-3624f5dde3eeb9266a87b13914b253db1a1dea4e09c5f4da434aef998dbb94b0"></a>

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

<a id="canonical-0164c7e71ad6fa14a453a2ebcd6804c9c91d7b45a1c89be008764e6a2444ab34"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 47daa4cc8f47 / 3

<a id="canonical-372e370f6f8d295a3f12e574e740d4a909d0aa28db572f54e9bac9176f8bb8a6"></a>

<a id="canonical-60408d0402de5f0dca2cd3d7efbf63bc7d4c3f98db8cc6114fa857d4808694a8"></a>

## provider_ref property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 47daa4cc8f47 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-15bdf5f93220b950f1b646465847da97c9685dece6992cb10ae6a4214bfa0653"></a>

<a id="canonical-d883e41713e56c5a3ea7a6332830e37e77d7a9c8548ebe44151d2b4ef1eb0390"></a>

## url property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 47daa4cc8f47 / 5

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

<a id="canonical-f5e3e59ba297ce624d8ed30becf30f4a15f700cbfb7d878e080bae29ad1838f7"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 47daa4cc8f47 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0610bec694f1041db05a8815bc645b626c180fb4ebf6d3892f4bf31bc1134609)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9c6bc2080dd8970f5408fce5825dc385b6d9e93a3e2839ea37c26fad4f54eca4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89b5ffd230efba6b2dac8ecf008e1a76f5221f7a5c020f66aa3497f663b47b1c"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — enable_api_discovery.api_crawler.disable_api_crawler / 11e83d68da0c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-a7e75872a169e41342a0369a10ab5c718d57a4e21533fc6bd07aba6ef210bf57"></a>

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
disable_api_crawler = {}
```

<a id="canonical-904a8924d1b89da2112931ee22f8192d48679c84743a822aa6e8c06a1ed1150b"></a>

## Direct properties — enable_api_discovery.api_crawler.disable_api_crawler / 11e83d68da0c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04ec57bbef66c5b5b34c62c0b058149bffec818d627eb0c24f56491f47403876"></a>

## Next pages — enable_api_discovery.api_crawler.disable_api_crawler / 11e83d68da0c / 4

- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-08c25fb3858586594e54203b6943179aa610b8ff76e66f7be23af64df1570b8b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8a958211739e545afd1a84be8278b9eb87cd211a15e499b08f6fd6775b543169"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfa73ea43ea5dfdbc92c40765efe8a82b3452dae4dc869215825d1ebfb771283"></a>

## enable_api_discovery.api_discovery_from_code_scan — enable_api_discovery.api_discovery_from_code_scan / f9d98d29f7da / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-628aa19f9f36a820828680903d497ee1aca25b069ab51f1780efe056dc82b9d9"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-92c9c682d43a5e357fad3015cddc6a86044b37357b0f2aa47e030254b7012f9c"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan / f9d98d29f7da / 3

- [code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3): complete subsection reference.

<a id="canonical-e612c2603411bc552cfd568ab6aad0d8a51a4f38e7a0620e49411f8b7ddc145d"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan / f9d98d29f7da / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-c92592437695ce4e484e16cf1b000e2558346e492b9be3ef743d389861095739)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-06644f1edd1899894148c2f317ecb380c5ef66abd2f2f0786f21206bf906d7b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
