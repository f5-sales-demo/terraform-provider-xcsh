---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-6b37aa57c6679830c76bec3ad2eedc923bb0bbd78855e7cb8f9d85cc6da5184f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-6acb655df79c463a0d2a4bc8074afb10fbab82c9912455482cf77566f6d2d31c"></a>

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

<a id="canonical-b258dbf59917a075bec71d78bedd68fd268298f7fe3b9eb861cf7f6db5d3c836"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 3

<a id="canonical-8dba1018372a1325f27d7853fd2d5f2aa0d85beac972f1267c5a10578336a245"></a>

<a id="canonical-3262130fd098118d7c791776ec5e38acaa26da427c99419279fd48aa669a42c4"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 4

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

<a id="canonical-0092e51e5b5da57a762ab0870ab6308cbb5eb6b45423e1aeb44af7236ba34117"></a>

<a id="canonical-dc7381936786b6042ea18b10160682d19ac47ec6ac99d5cd93781b00a2533c34"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 5

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

<a id="canonical-43f08c09737b0b71365ae90191089fb720893f62c75aa700d255afad473166b6"></a>

<a id="canonical-8263c5dd2cae5a44e48332747447e58a3b6936bae51d51ea01dc12250719de87"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 6

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

<a id="canonical-71f6dc5be4f855e55cb5c5c7f230a5e18f2c896befb3c97c2b48dc060fcc1360"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3f5c13ffe93d / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-68b162074fb03e93110ea5aee82a92ac322077e3b4cd13a890843fe83d6c3669"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c59fc9178ec191d78ced9e43b7203ba724bced79ba9152739b4737cd474177f2"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 324164eb80fc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-528e076deb03d642ef2e818c917f5811ac5f242017cbd1f09c11f5174e829199"></a>

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

<a id="canonical-9630f937cfad9f60950e7a6c2b53d5138dc9a41a3d4ae2f3ed88d2c1274f6b75"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 324164eb80fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33f3c3a5826992c3605e69c12b082035960ce9cce27164034ef6b98f92276d14"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 324164eb80fc / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-751f208d550f356d3e308219de8fbfcf7834b40c262336d3fcc6cfa18af8e82e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b93f08e960723b033d43070f5d6751d3847a5cd18b3f42a67a6b1172c98ef45"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cdd82528bcdd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-5787639bf4689548a4b19c207dd107a288731d3ddb88989855a331c8770067a0"></a>

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

<a id="canonical-610f22fbe3938670fd5d101a9b2cf9fc8589cb5fb8b2855a78bf6aec1e8b5c82"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cdd82528bcdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f3cbbd4f4089a2692e90a709c1855ab1251fe7ae46892c80f0ffafd3608ebba"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / cdd82528bcdd / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-cc3c026bd5204ba2e1dcc6d830ff663ff1a8c3469ffd5e89c1ddd5db18f877a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53584805e3edb5e73ef5d2b37fd4ac517cafbacf05f6b531814c381755746f9e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7786f906cea3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-51a44fbb018465fb6cba520cb3fa788e4fbec74ce58f8c695d4fa9640ac02b5c"></a>

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

<a id="canonical-a7529ea2ae6abd2a11e8d45fc86866e18e3dd7d20d758d50fad8c9b50ace56b4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7786f906cea3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6e8a73d8d17d26d4bbc9cf9a1207b725c6523c971736a779b2f1d0163bb09d3c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7786f906cea3 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-021.md#canonical-d6e93062fa3efd6d52325cf5c6017f02b76e0d639d5c6f04a7d6267e69fc123e)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af197f5d09fbdf59a1332a5408e151f14f40f2ad8c59c5007af854967c339c20"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f4ba69fc863 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-cda18e2e51af9591d74c5e9dac50753cdf59c7f6bafe15086d3871be8551d37c"></a>

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

<a id="canonical-fe66c9acee500ec4f54a7bfca7c2be4554e2dcf4424307ec70600dc5c3f2c1ec"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f4ba69fc863 / 3

<a id="canonical-cb24a6eadb233c9606a04783761fe58d3adf7eaa842dc27cbb917a86fa54ab0f"></a>

<a id="canonical-599bcbd7e673cb672a8ade0f03895f178e8325cb34f23215e9e1bb54ab45f7b0"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f4ba69fc863 / 4

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

- [crl](data-sources--workload--reference--group-022.md#canonical-72f8f253ec2e64e30a90946f459dda3e213814b4a36a48952ba33840a2a28d9e): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-022.md#canonical-333376c044f37321a4531b04db3b198f97899f04e714648e8b2fba27934430cd): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-022.md#canonical-22e24ecff0b8e0395a12187b1f1864d8ec87a24df06196f67e0129b0c956248b): complete subsection reference.

<a id="canonical-9c47696a61040dfeb334acd951b3c637d4a2f72982db024ec990ef30f6716484"></a>

<a id="canonical-9002d6dd459bdac9a87106bca2ecff404560ebd3633284d4a414fcdb80a54330"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f4ba69fc863 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-3c97dbf4020434089401c870af1d07cc2fe5aa949cbaf0e2062efceef2d2ef43): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-022.md#canonical-4dcf4fd42a31b544cce87014bcd79512d1a09305484d7f5f9ca570e993e54629): complete subsection reference.

<a id="canonical-9ad94add6d53311edf32b11fc70ba64d36b7ce867c92003894a9ca50a00655b0"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f4ba69fc863 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-022.md#canonical-72f8f253ec2e64e30a90946f459dda3e213814b4a36a48952ba33840a2a28d9e)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-022.md#canonical-333376c044f37321a4531b04db3b198f97899f04e714648e8b2fba27934430cd)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-022.md#canonical-22e24ecff0b8e0395a12187b1f1864d8ec87a24df06196f67e0129b0c956248b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-3c97dbf4020434089401c870af1d07cc2fe5aa949cbaf0e2062efceef2d2ef43)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-022.md#canonical-4dcf4fd42a31b544cce87014bcd79512d1a09305484d7f5f9ca570e993e54629)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-72f8f253ec2e64e30a90946f459dda3e213814b4a36a48952ba33840a2a28d9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc4f9e0c9f4d4cf9d71cba70a1aaa943a2d93745e7637dd2e174aff0da1c70a5"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-0570c35e362cfe2b11ad702c272d817ac94cdc32ce7ae45efe66044e52329975"></a>

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

<a id="canonical-a0e28731403e838aac6c7b5aa5271e26412eb4dc6a4bdb1d69f75a0f8f397910"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 3

<a id="canonical-002b0eea3ee540ee7eb72b3e88838b34362954fc4c18afad58d052c1ac225a4e"></a>

<a id="canonical-7d8adedadd39dfac0e36f25254bf9bcf7fa417a18d8019d8ad7644a268c7f70d"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 4

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

<a id="canonical-554957003175362ca415aa5ca62b928e86cc595dfc02abf8db9bae860442be8a"></a>

<a id="canonical-4849ebfa75e070e2c79d2eab6d95c7f14e5db23cdb6f7d7ed4f38a7c095337e0"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 5

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

<a id="canonical-5aafc0d4dcd6edcaa5afa17c7d22b2d005d3024a7b509f6709101e0e552e907e"></a>

<a id="canonical-d8f593722efe717a37ebf93b5134e8306c4bd18e66ca28ce819c5bfe94c0b9b5"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 6

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

<a id="canonical-565bc80f2833858a75415abccc13c329e65a3a1ac036b4346d63574f5aae4623"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2afa3cec37dd / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-333376c044f37321a4531b04db3b198f97899f04e714648e8b2fba27934430cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c00e3103deea2f3eaf13baabc9e462c1aa0cb989e5aec3bca4f3c5421a9e823a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e091db9e0d7c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-5819afbdee6d159c7f5690d780d155686be6420cf4239962ce85a5e7819df95a"></a>

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

<a id="canonical-8469cb6728f680d9d18742f71346e6d9be7ad9073f8e183fc6d6a3304938b19c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e091db9e0d7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c445604adce41f91b55b251046a3e649c4d5d78adea6ada340ca0a0c7617ac87"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e091db9e0d7c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-22e24ecff0b8e0395a12187b1f1864d8ec87a24df06196f67e0129b0c956248b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c27ca98c641eae4c19b08e1a4ed9be471e4c1e9831ac7c5647a88de308201a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-75aeb548d5d238ae0c54cd159a4c95a208ea2472daed6aec97e5f7fe4264a6a0"></a>

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

<a id="canonical-60b521d3a944c0ca7bd70d45e6d8f69026a15a0da607ee3db1fa268afa9e270a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 3

<a id="canonical-40d9ae9a82b9539241eacc8760ad5776ef889736a69240fd90553692bdc416d3"></a>

<a id="canonical-ffa231da4da3b885f4c1c111064d90e00f20c98519d84e29f99f826c6846386b"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 4

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

<a id="canonical-9b01a34f6e7b73877050f4de1b33278770ee2973f0bb55a7f37233b340d6fe5e"></a>

<a id="canonical-c64ba6d3e910162f4ba92ed2f162328e39563774a34d531e822e3b40dd3d4ef8"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 5

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

<a id="canonical-11e5db2f01153e91b291a43d945549d096860b59d5329a17ced432791d4bc945"></a>

<a id="canonical-4923ff944949a7ea1cdfe406d8b5e52ac4088aee43b66bd5cf728899573ee65d"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 6

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

<a id="canonical-d55f9e937a696a6ed605915ba7ec406298c9bf0f57c9359eb384248cdc50a9c9"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ca51bff34c09 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-3c97dbf4020434089401c870af1d07cc2fe5aa949cbaf0e2062efceef2d2ef43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-beb8a946023c3cbc800dd2c4004f747e92dc8d20e8a3c9486b7036841df6fd06"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b75b30cf54e9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3327079b4471342e8c6498402cff344b4e23481bf7ec6298a35e3d215284baca"></a>

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

<a id="canonical-1cc9d245d8c6da42cc0b7ed3bc6f489ea86872d43fc98b2b649b49c8389e4d74"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b75b30cf54e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a89912c8ff432976943662f182b8b02bcab454dc71bbc882d7e6d562d2cded8"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / b75b30cf54e9 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4dcf4fd42a31b544cce87014bcd79512d1a09305484d7f5f9ca570e993e54629"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d208d8fcefc6bb15aa3daf0e76823e75bb305397fdbca23ca2e2adcce600c120"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d39289590816 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-92118fa3f0734d579f2e8a192ad8683688d7dab049fa27b13c9aa605f29d51ad)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-021.md#canonical-046ebc3bda3ee61d4f97bc91221bb74207478726685c4f094e266e35303744b3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-792069e2afd28b5a13f7d9894e22aa070816f467a5c1b4acf715ed8510cbb79c"></a>

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

<a id="canonical-2f8e19ce109a2c3d939d98bc655242d0147e90cc1370b5690dd8f1765ea1a7f0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d39289590816 / 3

<a id="canonical-6f3e1a2d4a5ab6d111a1ac1cc8887f8e3952e8a2e296c278b08b82442e9512de"></a>

<a id="canonical-1f9e69fbfe9fd3bc95b2e73da7f6f3686a1941c73fc8cb284730aa3745dcce54"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d39289590816 / 4

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

<a id="canonical-48b78ed66f2f68b1bd8c925b71784126ebb924c8a1e8d60e7d4c45dc66be6c99"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d39289590816 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-f01ebb9827d06cd0a121e1e00bb67b5540bc1dd54a05af7de08fc923b2f75d7a)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-788a7c53e489af02dd0da5a51269aa1b7452d291002754926736ddc507adbe76"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-a5ea66b40b4c45efd17328993f38ca287f8fa97fba4d8f60ec230bf0602b5d77"></a>

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

<a id="canonical-e200faf20074f35de51fb79940e6a236f00ef79e0627200fd300234211f3538c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 3

<a id="canonical-5157140c88ed3eb7f3d62be9c8dc28aa921b8381adf573797d75d4efa9a7e66b"></a>

<a id="canonical-3862a320e25256e992b0898182b6058b363075729d7a607ad4fd9adad3ac5c16"></a>

## add_hsts property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 4

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

<a id="canonical-4e81f36c0249ef004bac64563b68c3116671ebc7d6136a75def454c76f0e33ba"></a>

<a id="canonical-3e25cf37572c6e0e1a2f9a3e9dad6b16ee5071917d2a5f218a3894f4de4d60f6"></a>

## append_server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 5

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

- [coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3): complete subsection reference.

<a id="canonical-c060bb8fc8e4f08d8064f7175e4d36174ccf66c6d5b4a69041343c6a1ea29957"></a>

<a id="canonical-b61e2f96be675515dfb32de584a30c46f1c70e0c45f5e7e1278c242a3aed0b29"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 6

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

- [default_header](data-sources--workload--reference--group-022.md#canonical-d3b41aaef4b68df394be86e89509b6f47e8db37ea825a386e54cbca50af67ffa): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-53ca9f4cbe3ce11cb17ec776e35d6399fc8ea35de66c1280f6087a97bf37f76f): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-022.md#canonical-ee3627d95268155aa7ec8e4cb7c8da0ca532b9abce82d843b9dcb872be7ebdc6): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-022.md#canonical-d80af29b8130accac43ef298443483cb53a95c5274bd61e8efaa5e20c0d4b02e): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38): complete subsection reference.

<a id="canonical-b59c52af096957f98e01b05bec6548a41adf196baf8fd2742da9d54aa2cd9b62"></a>

<a id="canonical-98d1b7efb353c463317b854249c5c232c0860781b4109c320b534969d65b74eb"></a>

## http_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 7

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

- [no_mtls](data-sources--workload--reference--group-022.md#canonical-120264537e890821082bbc2433af30fc6ddac7601ff4de2b9f3241e1c97d6429): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-cd9bea2d395c62b620a7e929de128930cdb0bf97832ba0847b1c65422c912a9d): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-022.md#canonical-d9c14ec6b5955fad7d3fd7dd87d25048200447bdb08461e00633753ca1866017): complete subsection reference.

<a id="canonical-249d08658781b3d9f219bf8b480a29de5cc9fb1d627574faf137f67224c2a215"></a>

<a id="canonical-7102587b2ea65f929698fc095dbabc4922c6ffcd951b85149e64b66bb743bff0"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 8

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

<a id="canonical-b946764e3d06fe9a6f03ac51f946a77213e8c04385b577c7cf78a5f129cc377f"></a>

<a id="canonical-781699786751cafd7a624ff11d0474f32a59cdf919cd4f8c59bf1b8aabd9a8a0"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 9

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

<a id="canonical-c04a94e3ba093b7f2ffa577dd27c06b67626ecd1af0c57126c5717f7f2f6e06a"></a>

<a id="canonical-0c3f7343eb9c97136c18e708fa81b89f3a34efd68bba7c72504fc83013b5f2c9"></a>

## server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 10

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

- [tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd): complete subsection reference.

<a id="canonical-f1132fc8e1047013ec5511450ec4f8c68b6214a4084565efb03860de7ca9640d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8a3b5a920d8d / 11

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-022.md#canonical-d3b41aaef4b68df394be86e89509b6f47e8db37ea825a386e54cbca50af67ffa)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-53ca9f4cbe3ce11cb17ec776e35d6399fc8ea35de66c1280f6087a97bf37f76f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-022.md#canonical-ee3627d95268155aa7ec8e4cb7c8da0ca532b9abce82d843b9dcb872be7ebdc6)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-022.md#canonical-d80af29b8130accac43ef298443483cb53a95c5274bd61e8efaa5e20c0d4b02e)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-022.md#canonical-120264537e890821082bbc2433af30fc6ddac7601ff4de2b9f3241e1c97d6429)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-cd9bea2d395c62b620a7e929de128930cdb0bf97832ba0847b1c65422c912a9d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-022.md#canonical-d9c14ec6b5955fad7d3fd7dd87d25048200447bdb08461e00633753ca1866017)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-022.md#canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2316d4c9454dab063ea2db4197b70f87ee7d26341a02a3eb0b298a71ba619cbd"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7402e69a3501 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-fa0b70a82788c338e5b7daa7ec4e7f6739c7125803b55d78747353f5d854d691"></a>

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

<a id="canonical-0300919c9f9308405f3d636abae420984787ced7ad686db2b454d42c78651f4e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7402e69a3501 / 3

- [default_coalescing](data-sources--workload--reference--group-022.md#canonical-5201342e894197f233b2b1bbe754534bad098516c0409857c15f811fee364185): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-022.md#canonical-327b4a58fade3884b979cffdfbdb7ae2669cc13274bf1404855d26b4a97ce450): complete subsection reference.

<a id="canonical-49f66cea3d211ce2956c73974af7eb882462c07828c24fefb2bebf8b2c3e1aff"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7402e69a3501 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-022.md#canonical-5201342e894197f233b2b1bbe754534bad098516c0409857c15f811fee364185)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-022.md#canonical-327b4a58fade3884b979cffdfbdb7ae2669cc13274bf1404855d26b4a97ce450)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5201342e894197f233b2b1bbe754534bad098516c0409857c15f811fee364185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f617bba6d43b267ad7c63706414c6bf9cc1511c5ba881e521621a7c49cd8db94"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d2015c3c4813 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-f429c5beaceb9c8e0959f3081ef50e34754742264d2642e73c136f090d7c803e"></a>

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

<a id="canonical-900e59ddf555a27684f29a6a69d1240776ff595b180fb5fa60f5fb230c475476"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d2015c3c4813 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04d8fadb470b6ef390226a586fef9554e18328a870229c69d8431f6acaf2331c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d2015c3c4813 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-327b4a58fade3884b979cffdfbdb7ae2669cc13274bf1404855d26b4a97ce450"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a65cc6e057226e42c113d32f428b71f4c56395606e6a6b098c35a364edcb4b2f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c40d2a2b74dc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-9884f5415d0b12153834b56fa2710ac3bb8cb6ef58130aee45a4b466af96f908"></a>

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

<a id="canonical-ef1375ba38f5c6750a32fc539693472cf970ade835f63707f8a3aa136fe25bee"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c40d2a2b74dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f833a77a311d7c711b9ad42bb163d502708022c589601d16c51548989adce70"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c40d2a2b74dc / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-e8a706cdd6d34906194aa827f83f7f9cc3eccaacf8f51cc93d72f33aca0c6dd3)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d3b41aaef4b68df394be86e89509b6f47e8db37ea825a386e54cbca50af67ffa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bfafcd6916061ee8085139691be362fcb7ea91feedf081582b22546ed411038"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4e591e6d55bc / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-c6027ef2c5837f9f5dc06498d7ae1f9199e3fa16f526655ab299c243d6afc40c"></a>

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

<a id="canonical-05d39ff1e5173b0d390e6f483b01efa72712a8b895fa872b2887bc7f0e4e1573"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4e591e6d55bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9541e379eb48480084e5fdc309ee66fe8c419a4e97398d3771c42e7804042c1f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4e591e6d55bc / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-53ca9f4cbe3ce11cb17ec776e35d6399fc8ea35de66c1280f6087a97bf37f76f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1d229b2a1d692f448d2affce1f7d13053362d49becd293e79d23946e7977280"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d29dc49ee2cf / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-c992aa982a9184912e2fb33148e12b342ef4b196510616187ad4195786032c85"></a>

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

<a id="canonical-58edcea64a1ee96c11b243dee75ce9e946c15dfae0c5074bab7100e781e22a82"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d29dc49ee2cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e1afb24d0df11b7cfeaf6539e4a24d8290e1cbfa0817c75b28a8cd46b247d79"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / d29dc49ee2cf / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-ee3627d95268155aa7ec8e4cb7c8da0ca532b9abce82d843b9dcb872be7ebdc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f80210a8bbee1fc39543b744efa42209aee4211b751af3555339ed07c10afde"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7620f1d3f3d0 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-7a61b8951d6ed82f5d05b4cb12ef11cc903b09accd4eeac0ea319418cc8d3a94"></a>

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

<a id="canonical-c14a9966f5af3672b6dc2fec3d75aac7e6c2f1f6e8db895b1b8c199060bdf230"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7620f1d3f3d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-256a88aff2350a4d8805784909781db238803c331f6079b2a85b2238c1569aea"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 7620f1d3f3d0 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d80af29b8130accac43ef298443483cb53a95c5274bd61e8efaa5e20c0d4b02e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-188c24f0230f99eaf92919092f42385592bd752896d1d03e667b5ec9e5313920"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 15e07c49738f / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-8b61ddd1024c2755a52c9d0eb7b552d512078d5b688fc3a702058c626d3c5a65"></a>

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

<a id="canonical-12c392e566225e5ae93f8d422a3f6e3b95180f8a6a1ab40fa3104e62aa9f5123"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 15e07c49738f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f87d3a4224a5edf31f9af64fadf6c79840d0dc3153cc6dc91e26b7eb1c4a0b0d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 15e07c49738f / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28258527fd23990566d6df8204d28eccacfb78775a5fae5ea520454467a13300"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 209821083259 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-ecd89205673ab449c4481815d231186311a9f357573e8bd9da407b4dd158cceb"></a>

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

<a id="canonical-1308d9afbfc76b206061d8060f77773660f37b3a3275c2d0bbcda229490d7668"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 209821083259 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-022.md#canonical-4273619d9c1c3239c81b4dbc0305b731d71922b80d94528625862ebd5212a7f0): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-022.md#canonical-e35890a083713808d2a568844186375c7ed1f78cde3620e8b4f3b5902812077d): complete subsection reference.

<a id="canonical-28b52b6bdf732958bd4eb814a5f22e60507462deb31b7806e7562b7f758724ea"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 209821083259 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-022.md#canonical-4273619d9c1c3239c81b4dbc0305b731d71922b80d94528625862ebd5212a7f0)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-022.md#canonical-e35890a083713808d2a568844186375c7ed1f78cde3620e8b4f3b5902812077d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c1ef61234085a556cbf910c9ea2f2f507a158ad060496b6a467426bc48f3148"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb3294784bba / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-50ec55b7199dadf15814e707c0e9b4a84fd93bb7cfce0e5f7beb9740affab335"></a>

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

<a id="canonical-c9ea727f183bd5d3080d12eacb97578bf535e1acac1f4780005184c04fa6cbe2"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb3294784bba / 3

- [header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07): complete subsection reference.

<a id="canonical-2971c9a895f08e54ba5a69c542906ec3a63a945770a314a00fb98f8c917a152b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb3294784bba / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08fb733596996e63e59b0d4084e2b1fca131cbf87910315883d4ce101e07f273"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 418333d3203c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-66ace451fdd4f0ece4267cd17cbde9d69ee73aa5391fc8aa8dc593dcd2ccf2d5"></a>

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

<a id="canonical-9018ce1aa1d3b598ddda014b56ddb2a7b9617a4224bf69f4395b5e6a6cab133e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 418333d3203c / 3

- [default_header_transformation](data-sources--workload--reference--group-022.md#canonical-a2de1ed99c2545917c4a63b4fa5b40e416e5df47ec2a1340e434fab9627721f0): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-022.md#canonical-7280b26d8fcbfa471956567517363fa2570e1188197e10da6326fc39df015285): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-022.md#canonical-bb55e10202b6ccc43a9e5ee8394239016bb7eb363956e4c8fa314e9d89f61a74): complete subsection reference.

<a id="canonical-efd0541201ce94371ebecdce8e6c60dcd39215a5dd111199b62ce153ed44c611"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 418333d3203c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-022.md#canonical-a2de1ed99c2545917c4a63b4fa5b40e416e5df47ec2a1340e434fab9627721f0)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-022.md#canonical-7280b26d8fcbfa471956567517363fa2570e1188197e10da6326fc39df015285)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-022.md#canonical-bb55e10202b6ccc43a9e5ee8394239016bb7eb363956e4c8fa314e9d89f61a74)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a2de1ed99c2545917c4a63b4fa5b40e416e5df47ec2a1340e434fab9627721f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d81dce3449348d3776a41354632cda4bca5b6f2e7c5816f02e9a1d276c7dd232"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f84f7f332643 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-c46ed05c7dcc880ac07b0e048fb45f509c26b5f1b12058966b2e47a6d349021c"></a>

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

<a id="canonical-262e36e9d93e3bf85aae78a9d5fe0d1b7f0089ce4231ef45c45f4923f64795bb"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f84f7f332643 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a936a26c99c49720775332fc49d9f3f86480a56328f9a05b63923ad59e8ea31"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f84f7f332643 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-7280b26d8fcbfa471956567517363fa2570e1188197e10da6326fc39df015285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11e9794170c542ae66bbf6ff046567fe39b1d34fa339e070363d20eaf1ace23a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e02ef2f363d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-02f366164c8c37c3e6a1d30dd8ed55b9c22c84acb0eed1c27c80ea693ef35790"></a>

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

<a id="canonical-bc5381cdf5fa550a13b94f973a81b76a90829d35c0a58711f06984de988c94e4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e02ef2f363d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80ff45ea7d44857a81890cc587c31487b3e44c51c96af21b5c33f4eb438841b8"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e02ef2f363d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bb55e10202b6ccc43a9e5ee8394239016bb7eb363956e4c8fa314e9d89f61a74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab07f59306cbf162b1c1109f75ce0ffefb91e246790cc8b6265ff67348d867c8"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3671f9ca527b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-e5db547a05bdb05faae6e6595c77531a7d2dce56229ce13b80ab8a87b43f7cea)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0741817a194cfae0554b9ac6c1a272bbf755d27dab1efa3f50ebadf33d655b94"></a>

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

<a id="canonical-e4ab4377bdfd8703a8f0127ee5ab69229eabc277a7c122732ef81bc925370b8a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3671f9ca527b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f39c1db1a15ee9bf78beb9676076d65445540338dae5000c97fbdb81a85de7aa"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3671f9ca527b / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-370272ec537bfdcb9aef9bf8d721e0e216ff6034861a920d59f6f73b372f0e07)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-4273619d9c1c3239c81b4dbc0305b731d71922b80d94528625862ebd5212a7f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a85481b93fa9a45bc270065d0eb90ae644c9ddb6bfb05cf13732f2201a2ea67"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0aef0b57bcea / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-e840d801150db55075f2d24ea960786bee568ec84ca1e8f9445b3c73f7fef137"></a>

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

<a id="canonical-c8a5a3d78951770399d37bc9c2d7fa7e33d6fbe5910a512efcc73afe7f952a91"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0aef0b57bcea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3be0f660b3ecbad28a509fa38173dd41a967ad27a9bb689dac135cf8cb3f6bf3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0aef0b57bcea / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-e35890a083713808d2a568844186375c7ed1f78cde3620e8b4f3b5902812077d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6903bae07c1cfa5aab04dfb0efd81ed192dc015cc89d5421d9b82c0f0b300760"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c5a0def67660 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-92c72d05c6549dac6404d1c97f9121bc87c120b55c9057db91c161fab1deff63"></a>

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

<a id="canonical-cf3e7d83401b3e831fc8ac0310b833e73981b231287ba2f332c8ddf856bdce3c"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c5a0def67660 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c20c18dd5348c3ad94beb2e023ee8e7496ca07c4b09ddcb3c69d7c378a515edf"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c5a0def67660 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-d3195ad2bb80b9bd047f5ed2728750a82f0dd6fd41498c45fe1a7c5a41026f38)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-120264537e890821082bbc2433af30fc6ddac7601ff4de2b9f3241e1c97d6429"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afa9342f23e6cd6dc383475f15df71f73cb78b4615c2e94806ac8e39e086d565"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f547b8594707 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-150d9ab6c8ddbb8856017989a7eddd1ab58e5410c072597fcbd7bd6f0742ea7b"></a>

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

<a id="canonical-23b6fb4626dabe3ff4897dcfbfc34eace02726b343c19c51da2002a041700944"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f547b8594707 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-debc5478eae61d44ff050103291b75f102ce33d5ff9230104fa4a97206754b10"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f547b8594707 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-cd9bea2d395c62b620a7e929de128930cdb0bf97832ba0847b1c65422c912a9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfe85aeabab65f7ac4442433bd61d776bc5b3ecb186f3794d4817584c9289c2c"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 569927182d3d / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-de9272ff5c430c45d0be289c98e22ab4b3dfd8469ee934b45b4cf37c1bb593f2"></a>

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

<a id="canonical-0cf4a4939956177d5c92cd6ed001d1a85648587c1829149efb65a2ed6f1f0a04"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 569927182d3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6e18060b6b05ff5fefa1316dc14ad0b8056b87f7f585bd05bebdf1b4ac5edd9"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 569927182d3d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-d9c14ec6b5955fad7d3fd7dd87d25048200447bdb08461e00633753ca1866017"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a00a079b46477596f42358a1e0d7c31588737db625e94238342e17680e5a7e4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ca11982d23c / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-177bf40f7e5c578178c33bedf3a5c7bbdd52642cdb48bf87d36f21aa07324c47"></a>

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

<a id="canonical-c40c352f732449e8e4cf96041cd8f24858d2ffb30708d0c7f115271a46a52cd4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ca11982d23c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-357b968176a6570b730c2dc58be9526968d6be6532c4ed75d45aceee6748ee0e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ca11982d23c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13847f793d4f7d87ea1825a855f95b195db3d1dbb79eaaae6417dfbf2c81ca3f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb8d703df2d9 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-fe340f5cd9558fc45246b0485f21a06470b3ccb2c81926b12985d8041c5a9aa6"></a>

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

<a id="canonical-8fe8d6a5baf2ec59574be087bcaf45d7452c20f204e2e1db4e5a8b993c3e3378"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb8d703df2d9 / 3

- [custom_security](data-sources--workload--reference--group-022.md#canonical-f9523198c0158504f833f7cd2678ee85723513849ddb2579feed4b1a52c265ca): complete subsection reference.

- [default_security](data-sources--workload--reference--group-022.md#canonical-a444e6a1cec524d232debd37bb63babe88dfc8e3e1add493ea05f36439fdc6d7): complete subsection reference.

- [low_security](data-sources--workload--reference--group-022.md#canonical-0746c0bfc3abae7df9472011ca63c9df59eae0a8940e31483341773737d0527a): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-022.md#canonical-5ea74c4f888da4075e3fbdf65096adcb5e794b6f8a6f002e47fd0aa09aa270ae): complete subsection reference.

<a id="canonical-e3d90e98a2babe72f1dd8236d56a7d531fd9547046beb9c0b1171287294a703c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / eb8d703df2d9 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](data-sources--workload--reference--group-022.md#canonical-f9523198c0158504f833f7cd2678ee85723513849ddb2579feed4b1a52c265ca)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](data-sources--workload--reference--group-022.md#canonical-a444e6a1cec524d232debd37bb63babe88dfc8e3e1add493ea05f36439fdc6d7)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](data-sources--workload--reference--group-022.md#canonical-0746c0bfc3abae7df9472011ca63c9df59eae0a8940e31483341773737d0527a)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](data-sources--workload--reference--group-022.md#canonical-5ea74c4f888da4075e3fbdf65096adcb5e794b6f8a6f002e47fd0aa09aa270ae)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-f9523198c0158504f833f7cd2678ee85723513849ddb2579feed4b1a52c265ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3fa67bac29327116be75e211998bb415387474c87029c43f7a78d49a29cea06"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-ca095ffbce0c50778d0a07b67fb0f5bd9ccdc5e380f1931d4edf6c58830be211"></a>

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

<a id="canonical-8954ef228bf4b035b4b9779837955b7e1c3f1fa233f766cae32c57b75355cb43"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 3

<a id="canonical-18732e27ca37e99938b635f0a88502b940b26dad80e216a4278239f6425790ed"></a>

<a id="canonical-d051e8fe7b855d8f4abcea69201c5b25d89f63f875caa1c85665e152e5f7b941"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 4

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

<a id="canonical-631c74f5ea751f4421da5406818b599c56e0d9a8026cdb8c755027ef1e08c60e"></a>

<a id="canonical-22c43a8580bd5e6a01c6d7bf78361ae2845a380ea257290a5878a693ca69e0db"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 5

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

<a id="canonical-c6e7ea7cf1c313bf4b251627a165b76ba12ff2116e606c8c4ecbfca497cd0a23"></a>

<a id="canonical-7959792d110a132bf66e00a12bd3d29d10936c5699b367b16106b764eebb60bf"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 6

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

<a id="canonical-c75911d37f5a7c93b0ea447114ba1299602073540c86ded8b57a98f8069ef4b3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / ed3de5614021 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-a444e6a1cec524d232debd37bb63babe88dfc8e3e1add493ea05f36439fdc6d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbfe0a0eae81e956d0968415a71d2592d782025126310d4a5f09755d0626a12e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4481e76c33e3 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-02d377ec1ce5518c39f2c76017268e7aee128925ad2a426d1028365b17e784b8"></a>

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

<a id="canonical-3a0b765c96ffbaf37bd427d069bc3b084f8ee9a3d8a00aa351b13e27f8492220"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4481e76c33e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03e92df40b769419e5f3ecdf0c57081c375e798f78ef3a3a4f2eb18519fe9322"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4481e76c33e3 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0746c0bfc3abae7df9472011ca63c9df59eae0a8940e31483341773737d0527a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af3e222531bc95de4d73420027aac792d672b8e8d88a3ef13872719f3b51a783"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3b360ef5e151 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-eae70d905803bf08abe34acb5fe8d0b3d161e31d310b51c969d780b79cbd6023"></a>

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

<a id="canonical-768761ca02abbbae958d7d6025a5f134b6c60a2e985fb095a1f89fc743ddcc91"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3b360ef5e151 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2ce7acb684e3f87483ceb96523bd329b003aa9ad90293e946b5c68fcf67508f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3b360ef5e151 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-5ea74c4f888da4075e3fbdf65096adcb5e794b6f8a6f002e47fd0aa09aa270ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b4d0b8524a42df396c5a9689d6dc262359bf959ae1cd24ef1f2ab50c4e43fd9"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2e7f6f4c1b30 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-bacee3d43d7dbf5702d0381af0041e7b43fc10cd6b25346b20ec39d36f2a152c"></a>

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

<a id="canonical-6f41d2c6607a3d207f0feed464f0e4edd6613c9ce1b154c67fcb4052d6629a45"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2e7f6f4c1b30 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85593bc6abafc0d324dba0a920d1b74a6d02b8c872d7ca593064a249cd87ac34"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2e7f6f4c1b30 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-022.md#canonical-bb6c8abb21ffb56c6da63b0cfc40d92f2a40457af7484dd7d627d4b79fc62490)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-feafac3c30f911a98096a4b3d0da61b9fe41e35c73a4efce957bc606373292cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88e91d1addc79445ce4540670b0bc9c7ee321b84d36d2fc611dd58b0b103a74b"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3226130dc64a / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-9ecc5193e89f3fd03b58e4ca87b12e44ca929d7b2fb971f37a8bbb71e6ebc497)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-8d2488744f1696f511350def83db625e6eaaa2b2cec3e1e4b29b815f55be6068)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-4a9dc0954793a1f95ff2c36f5f4000c83e1f22720e6b24258b5247c265d15f28)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-68d37c6a4a27ced5c021322801ecb411660ca2db93a3d72569b3403243bfd4ff)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-1251236759af652e9d20c05e70b1d0ac993f0d06e405d11834bbaf998fa6e5d4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-b6c3da2d123319b70cd035f38530e7c4f981086a06d3b6a85a8e62257f31c52c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-d839a30793886c61fd06909d6937c45fe1a29205269cf1c2d29af93120a0ce44)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-878c5bf9d642d38989c8904c5d406f0a0d1a003fd744ad88f71f5cb48f712ab6"></a>

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

<a id="canonical-67420dbbcc9ca9e20d6f14e68a3cebbd91280fd9a37c880903a8b262e575f0f8"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3226130dc64a / 3

<a id="canonical-97664577882bc42921172c1b9f5fc25af0abd58585afe4ba894855d9cefa2182"></a>

<a id="canonical-a6a308a755ff711665891784eb0122250558ccc1b4c11782ef6c0191b3928aa2"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3226130dc64a / 4

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

- [crl](data-sources--workload--reference--group-023.md#canonical-72c9ae85c6d7c1c9441815f4287691df1e2e2984d0e9850bb106edcba3233093): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-023.md#canonical-65f7b89b229ae56736233e983d6ee84a3ea49139621115cd4fc0526e39a3d497): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-023.md#canonical-c25f679cf13053c21e8478bd715145451461b80c6055eba199793e7b405917b3): complete subsection reference.

<a id="canonical-a0ec6978a8fd14e9fe3e0d0aaad8e1c50f6dfc822a7d5a2537238bb4cc637607"></a>

<a id="canonical-91baadf52b28a9749ff59572a4ed3c0492c024503ed82556f8806a28435ccb03"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3226130dc64a / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-023.md#canonical-212ea2af65c31256d594f079baeb4d219b208bb2e8bbab3c9c9819f3a2431ca3): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-023.md#canonical-633e70a684fd7390b2777fe958b675b0d21a9abc3e7ee2ad6a239c93f3c92472): complete subsection reference.
