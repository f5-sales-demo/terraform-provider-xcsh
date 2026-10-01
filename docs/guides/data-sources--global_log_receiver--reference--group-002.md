---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-d621ca3695e72eb3fe1adbafaaf111a155bf725452b48905ef5b2733fda984f8"></a>

## Next pages — azure_receiver.batch.max_events_disabled / 73ca8935150f / 4

- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-def1ef14667fdcb585f83fae4b18ed18e48417029785d272b1f408aa4e16a157"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7342bf0948d14b34aaf0c2833b114c666e4484472d87ee34ddd376a1ab1dda3b"></a>

## azure_receiver.batch.timeout_seconds_default — azure_receiver.batch.timeout_seconds_default / 01195f07c9e0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- azure_receiver.batch.timeout_seconds_default

<a id="canonical-5b31381aef3de6031f6b10cff74861e476d214c0f187a72adf69e3e94d69dc8b"></a>

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

<a id="canonical-d943ce6f8393eed879e1e5e0aeee36e4d8c0c7508672fda904a9e620d9eb23ad"></a>

## Direct properties — azure_receiver.batch.timeout_seconds_default / 01195f07c9e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6cab0825e26b6d39e45ce77c70de2c64d59b4f1a3b55dbadca3434aefac5bf4"></a>

## Next pages — azure_receiver.batch.timeout_seconds_default / 01195f07c9e0 / 4

- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-94b5d6cae427d6372c732fb99af31e900e01b115c060f422b7aab9e92d0a220e)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25fffd2eab6d1b3651e30d9b3fa4fc46db7f76f1250e9fd6e1fa161723c405b8"></a>

## azure_receiver.compression — azure_receiver.compression / ac8c25f4fd60 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- azure_receiver.compression

<a id="canonical-d2738f44957da101adde93ab23d91ca58bc172fb4c8798282b46621f00d67597"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-d48033e2781f1dfded835d1103ad6ba305f2754abdd59934b02c0b0c39e261cd"></a>

## Direct properties — azure_receiver.compression / ac8c25f4fd60 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-6aff064e59637634f3c38f88ab53c5e898db6f07dbcf49e215c859bf297688c9): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-6f6b153242d32379144214ad80d313efa3d7247ecba7545ad55301cea30ba742): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-ca2e5a8f121089006934c6f00345ff800a0e7b65542422f8bb377784104a9493): complete subsection reference.

<a id="canonical-542ba324cc276afd86c773e7e54e52a3a4e758700f02d6abe7facbd84944edeb"></a>

## Next pages — azure_receiver.compression / ac8c25f4fd60 / 4

- [azure_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-6aff064e59637634f3c38f88ab53c5e898db6f07dbcf49e215c859bf297688c9)
- [azure_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-6f6b153242d32379144214ad80d313efa3d7247ecba7545ad55301cea30ba742)
- [azure_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-ca2e5a8f121089006934c6f00345ff800a0e7b65542422f8bb377784104a9493)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6aff064e59637634f3c38f88ab53c5e898db6f07dbcf49e215c859bf297688c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f457ecfb1cf967a1588e75107ca63df7d8c0152ef8b7a7fb0f4a0266f17a94"></a>

## azure_receiver.compression.compression_default — azure_receiver.compression.compression_default / 41705667a606 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- azure_receiver.compression.compression_default

<a id="canonical-086288c086c32802a48bd111599752ee62c14655880db336654c3d9cb03f2435"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-c4f3f9b79afbd0430010725b99874ecbcc154609e4171a647f75c0167be8b0bf"></a>

## Direct properties — azure_receiver.compression.compression_default / 41705667a606 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69a157ed9e897c9f94f5071c702fd6b2993884ae0b0435e2f225ce0c5735a793"></a>

## Next pages — azure_receiver.compression.compression_default / 41705667a606 / 4

- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6f6b153242d32379144214ad80d313efa3d7247ecba7545ad55301cea30ba742"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1559862e21fbac8761b5ada3f3a239dc68c26388f9834b0de5aecdb926c88d5e"></a>

## azure_receiver.compression.compression_gzip — azure_receiver.compression.compression_gzip / be86caca4b25 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- azure_receiver.compression.compression_gzip

<a id="canonical-8f0646fe1ad8c33548e76b5dc1943a7a82609c23691e5d80761e32777e035288"></a>

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

<a id="canonical-0ef796e82e5c6dc91d34c7007ecdb5b4c3f220155d01f138fb74d2d58c41df81"></a>

## Direct properties — azure_receiver.compression.compression_gzip / be86caca4b25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76016abe0fe3914d327faeb1898d90eaac77dad4460eada24be3b33d064aae00"></a>

## Next pages — azure_receiver.compression.compression_gzip / be86caca4b25 / 4

- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ca2e5a8f121089006934c6f00345ff800a0e7b65542422f8bb377784104a9493"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f12064444ffe1e49c272a7b4905b768970dcb3ebef59101c9399975cb8950d0c"></a>

## azure_receiver.compression.compression_none — azure_receiver.compression.compression_none / d287e9693194 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- azure_receiver.compression.compression_none

<a id="canonical-c1cf177ac80c70e866c17182bf34a06edea969f5c1b9915e34b7c3af97a0e816"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-7c6da91aad6bc2ad954460721c02a398ed37469b8c591922744cd43e1d9b0d39"></a>

## Direct properties — azure_receiver.compression.compression_none / d287e9693194 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-750ad09d5d07b40b204f6c3351fcfe82dddb38cbe7a4bb6e0689247f68d0178c"></a>

## Next pages — azure_receiver.compression.compression_none / d287e9693194 / 4

- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-71a878ee1d36ee303f1eda176a382c17ae04e2922f1d2d38102d6f0138b7d5fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8ddcfb576156aa8724444e8fab2c0b77e10b75ccb617c377290eba1d1c18456"></a>

## azure_receiver.connection_string — azure_receiver.connection_string / 305e5fde15ab / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- azure_receiver.connection_string

<a id="canonical-9a8ab2b67481917b6269a6805c892ff08b34454c6553603e9b1d289629620eb7"></a>

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

<a id="canonical-c9fe6d1289724dcbf15e4b42cf33424c1638220a98758e863c2f4a61314f4135"></a>

## Direct properties — azure_receiver.connection_string / 305e5fde15ab / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-aa5ff27e32c29c60b9fc86aa006c8c8b3fcd8e54e4d3893694abc45edbe5d9ad): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-dac23c246116cc6ea5c375e84d5bfe520b8f298afe3fbaf73e59432d1bc314e1): complete subsection reference.

<a id="canonical-2dd0089b80ca1dba383901277e512557ed63b4bd6d8569afffc22a2bd412c5b5"></a>

## Next pages — azure_receiver.connection_string / 305e5fde15ab / 4

- [azure_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-aa5ff27e32c29c60b9fc86aa006c8c8b3fcd8e54e4d3893694abc45edbe5d9ad)
- [azure_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-dac23c246116cc6ea5c375e84d5bfe520b8f298afe3fbaf73e59432d1bc314e1)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-aa5ff27e32c29c60b9fc86aa006c8c8b3fcd8e54e4d3893694abc45edbe5d9ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe22ac6e859df7cb492fb3989167a4b91e8f072936662adda110468e283c4fde"></a>

## azure_receiver.connection_string.blindfold_secret_info — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164)
- azure_receiver.connection_string.blindfold_secret_info

<a id="canonical-9b62ae478506050f71f7fc7bc9d3e4408d7ae700e0162609e71ef22e1e22a85a"></a>

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

<a id="canonical-377629e31f61603261a7e488c44ff31d96acc1427a9bf7d4fdc8dd605a41b6c6"></a>

## Direct properties — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 3

<a id="canonical-a744ebc9550404dd61dcea5bf66772b90ce57d076e11589fa435dcf40978b101"></a>

<a id="canonical-97864b6e51af16c34027aac2c16b240f90b495974bc84604db353261148f1057"></a>

## decryption_provider property — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 4

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

<a id="canonical-a9a97fd5f79a84523a0049a312619d06f06c3de7bd0f28b6b6e8fe2e0ced3568"></a>

<a id="canonical-f068c8d7e520ff0427cb85d7dfc658bd2d12deecb80ba80c8cdec1eef9b56eb8"></a>

## location property — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 5

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

<a id="canonical-54df62adb361edbeec6fecd7b0a9b62a6544759b6d6dc0e8516862672955919d"></a>

<a id="canonical-0c7f896e6a85e8941bbf2ff443442dd9ec6bba123df1c922c5b507be309cd01d"></a>

## store_provider property — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 6

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

<a id="canonical-1d9de67f8252333f52b7d401c5b272d55d45584533da26f686bcda979080285c"></a>

## Next pages — azure_receiver.connection_string.blindfold_secret_info / d820604a2e1e / 7

- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-dac23c246116cc6ea5c375e84d5bfe520b8f298afe3fbaf73e59432d1bc314e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d775ac8481d55eee51ec120011bfec60cd0773c10ddf7ce1c2951e685260ba0"></a>

## azure_receiver.connection_string.clear_secret_info — azure_receiver.connection_string.clear_secret_info / 4f4b804114cd / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164)
- azure_receiver.connection_string.clear_secret_info

<a id="canonical-926c7a50ad8d0a896ce6777b5734de431fb118c5ef72ed0ec9a97401d8959e1e"></a>

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

<a id="canonical-13f8e65e70c5593e74e15ddb76c1950ea5f866c8a64bc973d5bda74190cc49b0"></a>

## Direct properties — azure_receiver.connection_string.clear_secret_info / 4f4b804114cd / 3

<a id="canonical-facb44a29e39230970a32f42c195dcb032350f5dcbc3b042b2b94bf0cfa39dd5"></a>

<a id="canonical-424fdf1723050a7e1b6852c900380b6e16d461cb80829a79e82dd29a2833df83"></a>

## provider_ref property — azure_receiver.connection_string.clear_secret_info / 4f4b804114cd / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8dad2fa3a7256b4d1522a2a7b83b4d9bcb65269c7be85f4acb018dd0f050062c"></a>

<a id="canonical-33ee98ad685ecd0dd573faf096d84e7a6adeeb4e688aabd47f98b938041c42bc"></a>

## url property — azure_receiver.connection_string.clear_secret_info / 4f4b804114cd / 5

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

<a id="canonical-f0847d10a194973e66397eaffbabba8dab275f5b424d9ea21d04f181d8eda664"></a>

## Next pages — azure_receiver.connection_string.clear_secret_info / 4f4b804114cd / 6

- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-b1d4378fc6ac36df67d7e84d8d66a74cc58a0fe156db7dbb626ca26b57bbe164)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d73569264aecf8e9909a0ccd32a32da8da892a7af51caa83fbbb48b9a1ee43"></a>

## azure_receiver.filename_options — azure_receiver.filename_options / 240c6695c909 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- azure_receiver.filename_options

<a id="canonical-d3e22600ed554da8cfaaea2fc08b5a052ac2a723867835c0276889a548aa8000"></a>

Type: `"single"`. Computed.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

<a id="canonical-dd5067d733b16568600f279be771e216c02a911da5b4ee8c3a5f61219d8465ae"></a>

## Direct properties — azure_receiver.filename_options / 240c6695c909 / 3

<a id="canonical-25913b195b678e1c346147d6998f11b4fc7b93dbd4dde9c079d92e0bfb5610fd"></a>

<a id="canonical-dce68850e5eb41aa580b54e2db0e3cece449cfc86502bb10070b7fa179df7eaa"></a>

## custom_folder property — azure_receiver.filename_options / 240c6695c909 / 4

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-b214b8028ceaaa7af500f8ad08156727822268533cc8c079116eac3ef83edfdc): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-e4b9629c270283b8675c0b695c3f5bbd5ce7e457eb75a0f4abda87d2a94536b5): complete subsection reference.

<a id="canonical-5cf162ab4f967686f7152fd543d289d6741d1c14b3a78614abe62edc5ec5ebd7"></a>

## Next pages — azure_receiver.filename_options / 240c6695c909 / 5

- [azure_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-b214b8028ceaaa7af500f8ad08156727822268533cc8c079116eac3ef83edfdc)
- [azure_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-e4b9629c270283b8675c0b695c3f5bbd5ce7e457eb75a0f4abda87d2a94536b5)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b214b8028ceaaa7af500f8ad08156727822268533cc8c079116eac3ef83edfdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e0369d42d7c95fc3f3733f59f24dfd8e368e6ea83d46748f3754450476e4b6"></a>

## azure_receiver.filename_options.log_type_folder — azure_receiver.filename_options.log_type_folder / c941e8cd1e31 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f)
- azure_receiver.filename_options.log_type_folder

<a id="canonical-8b206a92d0a96a4a933c67f6ccdb4eabc53984ce64e88b3164cc1a3fa4e4cc23"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for log type folder.

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

<a id="canonical-e21ed47697b6ec953266c7fa0b472b7b2ba584a5e43395429bef987cdcfbb61b"></a>

## Direct properties — azure_receiver.filename_options.log_type_folder / c941e8cd1e31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-394cef76d867558ce95dc6aa85e98e3b6dda3a0646b7300d3500fdb229650415"></a>

## Next pages — azure_receiver.filename_options.log_type_folder / c941e8cd1e31 / 4

- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e4b9629c270283b8675c0b695c3f5bbd5ce7e457eb75a0f4abda87d2a94536b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d196a113beb0e5da6560f4df0d83ee843d34b461462fff5d69f79444ec9164b6"></a>

## azure_receiver.filename_options.no_folder — azure_receiver.filename_options.no_folder / 5a9f6fc202f2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-99c68f0c69d8f9c338e720b39f9b655bfc69129b1f91704ea1303907e530b2f6)
- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f)
- azure_receiver.filename_options.no_folder

<a id="canonical-a5c03b4b29286e10a758840756f91c035ba813671cbe3b7bcae84faadcf0cb80"></a>

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

<a id="canonical-e93d19e9711c175910ec003346913051f4cfd8d57dddd40f3d4582f37c873078"></a>

## Direct properties — azure_receiver.filename_options.no_folder / 5a9f6fc202f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8878a2f39dc1286609439d86b3668026900b54e18eadf8d9e1f3fae778840a31"></a>

## Next pages — azure_receiver.filename_options.no_folder / 5a9f6fc202f2 / 4

- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-67e28dd59f9f134efb8387a9fd4b71f4c99385a9db5a25fa28ea0da0aa22330f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-229e171193ff781c733d4041aa61683fbe6fe10ca04a244715fcbe88dcdc431a"></a>

## datadog_receiver — datadog_receiver / 1233f9656628 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- datadog_receiver

<a id="canonical-4c6892ae68db2e1b186f351c385d1019fefce792671cb3df607e0554ec2ddc23"></a>

Type: `"single"`. Computed.

Datadog Configuration. Configuration for Datadog endpoint.

Upstream description:

Configuration for Datadog endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"endpoint\",\"site\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-697ff2ee41970174bb2ba3f9b08a64055f38313b85041c9434b6bf191fd0c32b"></a>

## Direct properties — datadog_receiver / 1233f9656628 / 3

- [batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3): complete subsection reference.

- [datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26): complete subsection reference.

<a id="canonical-32e4c9055a23446d8b7669634b242c92d7193d89e342b97e1af0e9d033c63647"></a>

<a id="canonical-355e0baf5358604719c7f584001d553872405f306bee33a791444f604ea8465e"></a>

## endpoint property — datadog_receiver / 1233f9656628 / 4

Type: `"string"`. Computed.

Exclusive with \[site\] Datadog Endpoint,.

Upstream description:

Exclusive with \[site\] Datadog Endpoint,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-251c2db095e7640abbf5ea252e3015234c33e5ee1d2c36f74dbd33cdd292a64c): complete subsection reference.

<a id="canonical-5f00d9798ce5f10dfa4fcda10cae8bf619585e963754f81cf4e3deda5aa2f64e"></a>

<a id="canonical-1b9b9981dbd9bc4ab5f3285c21dab8318a3ad90942d2a9080f194febf06a2646"></a>

## site property — datadog_receiver / 1233f9656628 / 5

Type: `"string"`. Computed.

Exclusive with \[endpoint\] Datadog Site,.

Upstream description:

Exclusive with \[endpoint\] Datadog Site,.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f): complete subsection reference.

<a id="canonical-b30c445b5476a76e8f1bfb4ad79b629ac4128104ab024dfd019fe45e5e61dbf7"></a>

## Next pages — datadog_receiver / 1233f9656628 / 6

- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26)
- [datadog_receiver.no_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-251c2db095e7640abbf5ea252e3015234c33e5ee1d2c36f74dbd33cdd292a64c)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b95a84e0e88facad1c16a37ccc04da67bac925c4891d9965860cc198c0a33e7a"></a>

## datadog_receiver.batch — datadog_receiver.batch / eb8bc726bd82 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- datadog_receiver.batch

<a id="canonical-8ecf8fe844568640d846280cb3a0ce7dbc084d5104a8e3e9f01c9c2eb5c70b15"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-35c43f2f10fb492b0ac9d043dcff54bebd932482c210513d307846b0bb070100"></a>

## Direct properties — datadog_receiver.batch / eb8bc726bd82 / 3

<a id="canonical-a18701e55c6e13a2cf1dd7e6ad1ef31c0644cc1a3d85fb083143d0621d111ede"></a>

<a id="canonical-c82db16983dbe016fd3526138caca3c2d1bde9baf5c9e3470a7dff91b4b2dadb"></a>

## max_bytes property — datadog_receiver.batch / eb8bc726bd82 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

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
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-f882a4a12dd5fe21c8330a0e880ec7237fcafbad58baacc99beaf4fa6ad8ec42): complete subsection reference.

<a id="canonical-cdcd43eec8d6b858a9ebd27e5440b0367156e09e3907460598279755ee3195b3"></a>

<a id="canonical-bf5cfe770bb1a6b011d613bcfc964770651e65a53940b069090d6af8c733e8e0"></a>

## max_events property — datadog_receiver.batch / eb8bc726bd82 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0c2114e3f491c32edc51249e137eaa8598b0c246c1173185798285b888f8af1c): complete subsection reference.

<a id="canonical-59ddef35fd30dcd9c2b1244eedfd147c7486e94fde5ddef5df41e4d6d516befc"></a>

<a id="canonical-6f602aa1b67054c0d6b2914ceecbdbbc76f1a4e97497ade6ea9c3fddb7f60dc4"></a>

## timeout_seconds property — datadog_receiver.batch / eb8bc726bd82 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-c10243e1fb5bd812a02289eea4562a42ec5b3448645b400bfddb56db473c587c): complete subsection reference.

<a id="canonical-0bd7d116e4959ee6b7648603e7d108594ee43f36f1590c6984c34c851d18ce15"></a>

## Next pages — datadog_receiver.batch / eb8bc726bd82 / 7

- [datadog_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-f882a4a12dd5fe21c8330a0e880ec7237fcafbad58baacc99beaf4fa6ad8ec42)
- [datadog_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0c2114e3f491c32edc51249e137eaa8598b0c246c1173185798285b888f8af1c)
- [datadog_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-c10243e1fb5bd812a02289eea4562a42ec5b3448645b400bfddb56db473c587c)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f882a4a12dd5fe21c8330a0e880ec7237fcafbad58baacc99beaf4fa6ad8ec42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f44789f33374bb089a15f313ef780424016200eb75ee732fef05b6c471e746a"></a>

## datadog_receiver.batch.max_bytes_disabled — datadog_receiver.batch.max_bytes_disabled / 389bba040d10 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- datadog_receiver.batch.max_bytes_disabled

<a id="canonical-30637618a2a579faf2ba0f1de44caa34b7520f55b744e1b15ea7fcf1d94c89bc"></a>

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

<a id="canonical-bc54437915a2010f8af16d688aa38e7d37c1280b3eb2b7e8a59f00141ac699b0"></a>

## Direct properties — datadog_receiver.batch.max_bytes_disabled / 389bba040d10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66a0ed42f8398ee331cf8104a6aed56317573882186b4fd68dd5be2a7048cc87"></a>

## Next pages — datadog_receiver.batch.max_bytes_disabled / 389bba040d10 / 4

- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-0c2114e3f491c32edc51249e137eaa8598b0c246c1173185798285b888f8af1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e125ff0918b005a58505f6e64f126e4beb678d5c553498822db68de1fee3132"></a>

## datadog_receiver.batch.max_events_disabled — datadog_receiver.batch.max_events_disabled / 5f77b8a77320 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- datadog_receiver.batch.max_events_disabled

<a id="canonical-63b7bab2b36f2ac5fb91b3c4bd94e4ccd37c11789152f2c059e8395ccfc3ff1d"></a>

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

<a id="canonical-84649e8252f42f432c9883bc2da818b5a155c843261a002482200ce45c7ed259"></a>

## Direct properties — datadog_receiver.batch.max_events_disabled / 5f77b8a77320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9801eb08ef39558dbcb1be9cb515511b272dd8ed8a150b02296de24e9d8d6f0e"></a>

## Next pages — datadog_receiver.batch.max_events_disabled / 5f77b8a77320 / 4

- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c10243e1fb5bd812a02289eea4562a42ec5b3448645b400bfddb56db473c587c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-615e7bafe0366892cc2855838140a06a964327769b9d4fb518b805f329bf93a9"></a>

## datadog_receiver.batch.timeout_seconds_default — datadog_receiver.batch.timeout_seconds_default / 611f1cb575d0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- datadog_receiver.batch.timeout_seconds_default

<a id="canonical-b55f689a0c07b6da9a2a30289d24e56333b3c8cf2e58b81e9deea3db36a4062f"></a>

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

<a id="canonical-647f1a8adb6ee860d2be5be71b4e60ba004ded790c488901d99f6e0f65710599"></a>

## Direct properties — datadog_receiver.batch.timeout_seconds_default / 611f1cb575d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6cf0338333e9eb2ab5f11598e92d0a28d5ede3edbf81a05c0e50b9aee50b1f8a"></a>

## Next pages — datadog_receiver.batch.timeout_seconds_default / 611f1cb575d0 / 4

- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-370c75ff3ae46457ae92eb9d5b5b67313b6e397125d76f7dd01e73a9f8da25e4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b76dbf037bde655431116838b790c79bc8090996ae8405079e28eb3090625f66"></a>

## datadog_receiver.compression — datadog_receiver.compression / 515b3ac36803 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- datadog_receiver.compression

<a id="canonical-951ad4f64255801c1bbe1236d09be1aa809874f2dd5455e4bf0391f6610c9cb9"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-c9e71dc17e7d7a37c595f9ecabaf265c6bb34e5887adddaa847183e9619ce81d"></a>

## Direct properties — datadog_receiver.compression / 515b3ac36803 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-5e25707ab7d20c6e93f5dea25a2e23228f81c1fe1c1894d7dbd47cd072fe8334): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-a198c55a648abd4f520bb663f8615de6290b6646fa436ec3773ec179e818f232): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-b719a09ecfa97de6fce9c570ed89fe9f9d314c380c9b863441ad21c8877d4866): complete subsection reference.

<a id="canonical-e865dc10b0e6fcdb6baa070da5f3d86f320300ab1f71650ffcb508364866acae"></a>

## Next pages — datadog_receiver.compression / 515b3ac36803 / 4

- [datadog_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-5e25707ab7d20c6e93f5dea25a2e23228f81c1fe1c1894d7dbd47cd072fe8334)
- [datadog_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-a198c55a648abd4f520bb663f8615de6290b6646fa436ec3773ec179e818f232)
- [datadog_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-b719a09ecfa97de6fce9c570ed89fe9f9d314c380c9b863441ad21c8877d4866)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5e25707ab7d20c6e93f5dea25a2e23228f81c1fe1c1894d7dbd47cd072fe8334"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2590ea8042ee235c0f508af9f61c2df8b0571c4b41d63aa73c772d17097b27df"></a>

## datadog_receiver.compression.compression_default — datadog_receiver.compression.compression_default / d72ec55f0712 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- datadog_receiver.compression.compression_default

<a id="canonical-e43da377baf4adaed52f9bf8799918b4fb6d37bf56dcb4451ec694f157175bfb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-5ac46b6890cbe2a6cb9c836067ae5a2edd555199f10f45850e95535e28f1365d"></a>

## Direct properties — datadog_receiver.compression.compression_default / d72ec55f0712 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26490a9b11d5b76b96e15430372cdf77d8b8cd36eb68d5da3bd578e3871c97de"></a>

## Next pages — datadog_receiver.compression.compression_default / d72ec55f0712 / 4

- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a198c55a648abd4f520bb663f8615de6290b6646fa436ec3773ec179e818f232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d225751a89821caec701ac0bfc8a0cbadd00446c8bcefafd0dc0846d5c126d1"></a>

## datadog_receiver.compression.compression_gzip — datadog_receiver.compression.compression_gzip / c72b1541b1cb / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- datadog_receiver.compression.compression_gzip

<a id="canonical-43dd18c502649a825da456d90cfc1fb475675c4df863728a82bbc1978029da4a"></a>

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

<a id="canonical-145c1f57b3dad5d10fe4c78e43ef9f0d403bafba617e027caa5f8f789d2f4610"></a>

## Direct properties — datadog_receiver.compression.compression_gzip / c72b1541b1cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-821fc3f6b98c7b0a6fba18c6ac2e779cd2e00dd0f24cbb2bc306c3035e78a0c2"></a>

## Next pages — datadog_receiver.compression.compression_gzip / c72b1541b1cb / 4

- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b719a09ecfa97de6fce9c570ed89fe9f9d314c380c9b863441ad21c8877d4866"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ecba033e38a4be09bc9105d60a29d44875ed71deb33f58f56bfedb34eba4844"></a>

## datadog_receiver.compression.compression_none — datadog_receiver.compression.compression_none / 96cb5f973d7c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- datadog_receiver.compression.compression_none

<a id="canonical-8e9cf508f57f3d052f6a4a5c3ec9269e35d1e246edd34202dfb8bf3d19b75155"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-eb7ac4bc02f5c0ade98fcd30e25fec7082721a520086e80401436a394b4a0da7"></a>

## Direct properties — datadog_receiver.compression.compression_none / 96cb5f973d7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab7848b406219faf9741828e8321dcdcb1ac85fc874ccbb93734c2036e0d9afe"></a>

## Next pages — datadog_receiver.compression.compression_none / 96cb5f973d7c / 4

- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-15faf461dfa00005914fdfe10517cf3cb84dbcd56c62f6e29fda5dd342ada4a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-624955da577a99f1f6fceb6323bfddac00c7366f2ceadc793804db7c057d2c93"></a>

## datadog_receiver.datadog_api_key — datadog_receiver.datadog_api_key / f472baad1782 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- datadog_receiver.datadog_api_key

<a id="canonical-3b4764f72c0139bb246e25e69b1c49172d79a115dc699afa8aebc35d23e51dd6"></a>

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

<a id="canonical-20a3fe467787e3eaf6667e114ad4a06fde6fe0ed7040bcbda724a579b398ab50"></a>

## Direct properties — datadog_receiver.datadog_api_key / f472baad1782 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-67a7b4f2a203ece70cdc2da64b1c1bb5dc186b364e6989bf0cd3d2829998365d): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-8cac806313e7252e3d3ce0406cab1be14ef9eca350528bb4b4436ae84a9e6c63): complete subsection reference.

<a id="canonical-0596549216acec2678b6d0a11871ce43e92fb63ec65828cd30aca9d0adfa27de"></a>

## Next pages — datadog_receiver.datadog_api_key / f472baad1782 / 4

- [datadog_receiver.datadog_api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-67a7b4f2a203ece70cdc2da64b1c1bb5dc186b364e6989bf0cd3d2829998365d)
- [datadog_receiver.datadog_api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-8cac806313e7252e3d3ce0406cab1be14ef9eca350528bb4b4436ae84a9e6c63)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-67a7b4f2a203ece70cdc2da64b1c1bb5dc186b364e6989bf0cd3d2829998365d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbcd868e27e97ee3c6a2b1a3e41cbe31f3ee440c9d4bcffed6535d09f1a0f3e9"></a>

## datadog_receiver.datadog_api_key.blindfold_secret_info — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26)
- datadog_receiver.datadog_api_key.blindfold_secret_info

<a id="canonical-fe83c79b3b5a2a87778a0bba976b8aa21b2371cc049fa2c9e43b45cc456fc56c"></a>

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

<a id="canonical-415817e6bcb68eb33ec15ea1e437508185800ae2f8ed92e34307aa87e12817e8"></a>

## Direct properties — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 3

<a id="canonical-06793d20358936e63313d7c3b3d11fd1d791f28c903e376e3d0f094d3e796802"></a>

<a id="canonical-e227b386df0792a6c35f583d0e45b611fd7824944c4c5220e1af3e75ff33842e"></a>

## decryption_provider property — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 4

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

<a id="canonical-193837d52a6573f6f2c01a654e754e2d1635f3232757d415ab49d03667890bb0"></a>

<a id="canonical-f19e659cbdfe9e324b7fcae094b30f633a38320de02c43f934a27c8d9cf40aca"></a>

## location property — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 5

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

<a id="canonical-3d12e1b7d780d698f1d7ab07ce2c4f2839d35edc301a38d1507f0b69af58d952"></a>

<a id="canonical-07a56aa0e6e9b52ef60b83bf4eda058e32c24ce2e0211bec7774c80343f436bc"></a>

## store_provider property — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 6

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

<a id="canonical-f809255b16783a73e214bf937c590bd2e6fc0b9863bf1d66c9c693397c0c28d4"></a>

## Next pages — datadog_receiver.datadog_api_key.blindfold_secret_info / c7e761fc6dd4 / 7

- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-8cac806313e7252e3d3ce0406cab1be14ef9eca350528bb4b4436ae84a9e6c63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8991912b73e7d49a64d683e4433eb76d2d3ac95105e837ee3fcd7f049ca6b702"></a>

## datadog_receiver.datadog_api_key.clear_secret_info — datadog_receiver.datadog_api_key.clear_secret_info / f1b28f6e5bc2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26)
- datadog_receiver.datadog_api_key.clear_secret_info

<a id="canonical-1f9417400ee0546c88a35a96ab86efb120b36233c50c0e8175278996ff480cf3"></a>

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

<a id="canonical-b39926c76013fdd616f766a671df9534d95671684b30503562ec968bec14f06a"></a>

## Direct properties — datadog_receiver.datadog_api_key.clear_secret_info / f1b28f6e5bc2 / 3

<a id="canonical-0988e20d2c700e9c4bb198c13b22c4508c21cd40ea55d677a731936c9da75348"></a>

<a id="canonical-2763128772dedb2e0c96ffa81f12eaa0498860c57f141f3e57293ec875a4225d"></a>

## provider_ref property — datadog_receiver.datadog_api_key.clear_secret_info / f1b28f6e5bc2 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-18a53425696b0cb46dff9524c18a00c376e9e13af0af5f9847a54c5224a6a0a7"></a>

<a id="canonical-49525891078f86e8b14143c74c6e5a1da6b03e71e3e2fc965ca7322a8c5d4988"></a>

## url property — datadog_receiver.datadog_api_key.clear_secret_info / f1b28f6e5bc2 / 5

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

<a id="canonical-db096c1b1e0ada8c9f9404712de9a3375a3d0cefb9402df0fe2336e66a687eb7"></a>

## Next pages — datadog_receiver.datadog_api_key.clear_secret_info / f1b28f6e5bc2 / 6

- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-b98aa356c35292d2e931c6bd2f2f3f8bbd0d79c55ad12bc747e2e121cc128a26)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-251c2db095e7640abbf5ea252e3015234c33e5ee1d2c36f74dbd33cdd292a64c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61b9dc1feeccbe321ff2e2e75a70d6489111ad2ec321e422c6b10a6fbc0cec62"></a>

## datadog_receiver.no_tls — datadog_receiver.no_tls / cfe25723988f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- datadog_receiver.no_tls

<a id="canonical-8bddf8deae3d8010a00b52e06e73acd601fb8b978ce18652636dd280d8c4ce39"></a>

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

<a id="canonical-7fb5170cea0701ec914773af2216a0bafa21df9a36efd33d14db3e59f94d4fd3"></a>

## Direct properties — datadog_receiver.no_tls / cfe25723988f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b06c3b81963f3c2e0ba5cd399ddbd9d3bdced864be23a29a853e48bf25e669ab"></a>

## Next pages — datadog_receiver.no_tls / cfe25723988f / 4

- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-687a36809df6e19ce8d48df55bffba79af2c85efbfe7824e471f67eefed9d684"></a>

## datadog_receiver.use_tls — datadog_receiver.use_tls / d20c2e53d258 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- datadog_receiver.use_tls

<a id="canonical-d3ec2a9b3a86cef3bbc3253d1e4c17ebc021362047bf28e4ec948704a41f95c7"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-467e935e80ac1e338a2d495d4e1e5513761833ca22f482787fb71730268b070b"></a>

## Direct properties — datadog_receiver.use_tls / d20c2e53d258 / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-3fa8d9a1a0d17f093d8ada1aaa8d13377fdd18dd23eaa003721cf8e68e645a4b): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-8d326ab3d90eaabb8af723a6736585e91a3e22184e3abc72c97a7c26427085cf): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-d1cae5a810f6f879ad18517110153c9d7ff36e9a91f2d805c085c03c90bfca30): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-a9410a7cdc60a85b4152dfeffe6f4c50a332d53885319d78caa50c2eccdbddc7): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-c78162594b67e9803fb9b4e97270aca3370400ee133b9c6b676d280d0d0fd3ed): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-002.md#canonical-bac46d4ca622e0699daedc00ae9ae1977743b0a7044c299edec971222267190e): complete subsection reference.

<a id="canonical-f079ce660e3481ad4433bc031d11254834dee7b14a1a0eb759b9ae9841279d65"></a>

<a id="canonical-330b31ff8655702ef545c2b1565c5ff5941a67d73755ec0571495d0e77fb2dd4"></a>

## trusted_ca_url property — datadog_receiver.use_tls / d20c2e53d258 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-daf007aee4cfeeccb18d038b791ac186a30729f5413b20fa5c152e51adbbaeee"></a>

## Next pages — datadog_receiver.use_tls / d20c2e53d258 / 5

- [datadog_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-3fa8d9a1a0d17f093d8ada1aaa8d13377fdd18dd23eaa003721cf8e68e645a4b)
- [datadog_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-8d326ab3d90eaabb8af723a6736585e91a3e22184e3abc72c97a7c26427085cf)
- [datadog_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-d1cae5a810f6f879ad18517110153c9d7ff36e9a91f2d805c085c03c90bfca30)
- [datadog_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-a9410a7cdc60a85b4152dfeffe6f4c50a332d53885319d78caa50c2eccdbddc7)
- [datadog_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-c78162594b67e9803fb9b4e97270aca3370400ee133b9c6b676d280d0d0fd3ed)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2)
- [datadog_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-002.md#canonical-bac46d4ca622e0699daedc00ae9ae1977743b0a7044c299edec971222267190e)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3fa8d9a1a0d17f093d8ada1aaa8d13377fdd18dd23eaa003721cf8e68e645a4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59d13806ee6de8a865417545faa6111ab1ea3cd14cf90d7dcb344515b0d7b663"></a>

## datadog_receiver.use_tls.disable_verify_certificate — datadog_receiver.use_tls.disable_verify_certificate / 9108a3dffea9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.disable_verify_certificate

<a id="canonical-e51991fd316ad751946194cb45351a404124461893fc381578495a9af12b5e2b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-c7d038d81a5652959f87aef4125b7a79784839535129ff6ea44df4e743e62a82"></a>

## Direct properties — datadog_receiver.use_tls.disable_verify_certificate / 9108a3dffea9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-273b2acb9e8d7dfbac215ce8fafc64f68ee9551e7dceb040a91ab60dcc146d59"></a>

## Next pages — datadog_receiver.use_tls.disable_verify_certificate / 9108a3dffea9 / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-8d326ab3d90eaabb8af723a6736585e91a3e22184e3abc72c97a7c26427085cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4622e614de91828f97a2b1b1b1f655559cab28019f54d25d5f74084e8850f73a"></a>

## datadog_receiver.use_tls.disable_verify_hostname — datadog_receiver.use_tls.disable_verify_hostname / 908aee59936d / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.disable_verify_hostname

<a id="canonical-bdfae2de990360c37d6d31230df3b3830d0d16a1156b53d9d55db22267216a45"></a>

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

<a id="canonical-86e514753100df7e108a87a0341dd54aa233b888b54b954351d127c969964201"></a>

## Direct properties — datadog_receiver.use_tls.disable_verify_hostname / 908aee59936d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94227d5de91928570ae70140aba8c4f15e07768e13ceef776a3c6b8101335476"></a>

## Next pages — datadog_receiver.use_tls.disable_verify_hostname / 908aee59936d / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d1cae5a810f6f879ad18517110153c9d7ff36e9a91f2d805c085c03c90bfca30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f45beba6ad71ac38f35b476dc5144e834043bb0c90dba467538463cc74d441e"></a>

## datadog_receiver.use_tls.enable_verify_certificate — datadog_receiver.use_tls.enable_verify_certificate / 36da19f8f989 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.enable_verify_certificate

<a id="canonical-9773376843b913e9e41565bc24bfc07ba2da6ac70f1f85f3ea36a4bad418ef09"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-a0c65a27a57e966894006ae158fb27c4fc864c84568cba79b2825653d1403522"></a>

## Direct properties — datadog_receiver.use_tls.enable_verify_certificate / 36da19f8f989 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47dfcceae721f1b34ddbf14abfd2abf420472c288e430e199bd8aadccca18851"></a>

## Next pages — datadog_receiver.use_tls.enable_verify_certificate / 36da19f8f989 / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a9410a7cdc60a85b4152dfeffe6f4c50a332d53885319d78caa50c2eccdbddc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70ae723d7588517ff2b89d139d3028a69eee925131881f6a68f93ec1a933f687"></a>

## datadog_receiver.use_tls.enable_verify_hostname — datadog_receiver.use_tls.enable_verify_hostname / 39dab66fb3d6 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.enable_verify_hostname

<a id="canonical-a5921b2bc139d637ffeb4827cf7088024989142b79824731724c9c0d9026667a"></a>

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

<a id="canonical-87ef3b87aafb2db063e02fe3a09378d55c41f83a0451c94613fd8774369cb690"></a>

## Direct properties — datadog_receiver.use_tls.enable_verify_hostname / 39dab66fb3d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d26e961cb390f8393aecdada879727509da5de0930ac6107f51f429e10600e1d"></a>

## Next pages — datadog_receiver.use_tls.enable_verify_hostname / 39dab66fb3d6 / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c78162594b67e9803fb9b4e97270aca3370400ee133b9c6b676d280d0d0fd3ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a535dd15059647d3424d3c79b0cf080e2c60db4b30dc5aee4b2293cb52fab1aa"></a>

## datadog_receiver.use_tls.mtls_disabled — datadog_receiver.use_tls.mtls_disabled / 695e67edb753 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.mtls_disabled

<a id="canonical-5ee0eefaf8ab38af5e1f102d5d789c2bb55f42f945de35e6c7b6f2bffde61589"></a>

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

<a id="canonical-1ff06edc706067e8ff6b3511b6a3e194dcfd412e3b5cc6d0de82ae6725525e3c"></a>

## Direct properties — datadog_receiver.use_tls.mtls_disabled / 695e67edb753 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80374daa95993bf26a1bb3720cf7b3e6fa33e2cee48461439e1eb6c230e61013"></a>

## Next pages — datadog_receiver.use_tls.mtls_disabled / 695e67edb753 / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55ed9ca2d73f191c39dc2d637259ea950fe136047c78453ac09860646645ae15"></a>

## datadog_receiver.use_tls.mtls_enable — datadog_receiver.use_tls.mtls_enable / b616938b55ab / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.mtls_enable

<a id="canonical-aa4734897cfe92c5c9c6617dd19a7263c59f3ffca3863e76cd2c598702b07ca4"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7b276358d2288b2796f693e4fad54fae1d92af04674799b17aac5eee0abc36d7"></a>

## Direct properties — datadog_receiver.use_tls.mtls_enable / b616938b55ab / 3

<a id="canonical-14dbf16c960950556a755cb0f85d2f10a2b1e418117b46617ab0e9b3bb44b206"></a>

<a id="canonical-9e500db3bcf98c5ddcf10edf5edea86058b8dc84fd7556142d9a0b720a500523"></a>

## certificate property — datadog_receiver.use_tls.mtls_enable / b616938b55ab / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844): complete subsection reference.

<a id="canonical-3c52e9b7bd495471bfee1b4dbd4734a36b1e6a57a1efa46e657c4ed54b902f88"></a>

## Next pages — datadog_receiver.use_tls.mtls_enable / b616938b55ab / 5

- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea3a33dbeb819ba2ca0ab7207788607e5dcb6ae52611835d327126c72609c801"></a>

## datadog_receiver.use_tls.mtls_enable.key_url — datadog_receiver.use_tls.mtls_enable.key_url / f0bdc31c1045 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2)
- datadog_receiver.use_tls.mtls_enable.key_url

<a id="canonical-b17a03c634b6cb250c4c6730fa7f8be7d179fa2e3fae5975161876bd6783a656"></a>

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

<a id="canonical-b4a140afe2efd167071002f47bc4ade5f466ee7502269b193206aa7a3d77c6ea"></a>

## Direct properties — datadog_receiver.use_tls.mtls_enable.key_url / f0bdc31c1045 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-588c7741e54cddf827b3127a2c0db77876ac53ff13602baae2b19a52a1ed48dc): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-e99a76942ee7c01739f410e8bce42441ffd340d8a31b73567eb94aab7c29ae04): complete subsection reference.

<a id="canonical-ec7fd48eae41e5b64710bd988209b6f47e7c81abe8cf459c2da94cc87eb89a4e"></a>

## Next pages — datadog_receiver.use_tls.mtls_enable.key_url / f0bdc31c1045 / 4

- [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-588c7741e54cddf827b3127a2c0db77876ac53ff13602baae2b19a52a1ed48dc)
- [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-e99a76942ee7c01739f410e8bce42441ffd340d8a31b73567eb94aab7c29ae04)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-588c7741e54cddf827b3127a2c0db77876ac53ff13602baae2b19a52a1ed48dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3856f40dd7c9bace7cefe621dcd554db80ad0411c1de4829c6e6e6ae6c1cb180"></a>

## datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2)
- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844)
- datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-fbea8ff3604339cd1032d95248e8701d7c86f391b586318e8260c714c5503238"></a>

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

<a id="canonical-206cfc9900d4142f7ab601256095c2e897adbb70cefec527da444eabc528d6c0"></a>

## Direct properties — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 3

<a id="canonical-71666d17008ac5c410cced558c067a2a6a8c48c069c110af97d7d632b85d3a47"></a>

<a id="canonical-38f2a1ce3329af82eee8134640a02be6355a545393c2fa1cab6d11afdeb4faa0"></a>

## decryption_provider property — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 4

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

<a id="canonical-bd8ff3c7bb88e5b7cc55874a3e8dd18130e5204ddd02baad0dcd034f574a717a"></a>

<a id="canonical-75b504fa067653491627e3e899d01ef418141bebe685cf3177d5980eaa2c7c6b"></a>

## location property — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 5

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

<a id="canonical-dcf6b54e5941ededbfcbe14c0df740a5364e9534896d42abd811a135ee509928"></a>

<a id="canonical-485464d4ba128c5554cb97f126c238ce515a3ed966ab8a9919e7dc3bfa443105"></a>

## store_provider property — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 6

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

<a id="canonical-3b91cf0bd161d94747eae58c95fe74fd60c592377ee85390bdf15856fd4fccf1"></a>

## Next pages — datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / a737308e5171 / 7

- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e99a76942ee7c01739f410e8bce42441ffd340d8a31b73567eb94aab7c29ae04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7cd3800cb7e24745391489e9cb9ac65155e15e7a1242c07b94cf559979fdd54"></a>

## datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info — datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 42822de4e91f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-ff5706714e0fa2ce1bcdc4b0cd6bfa1141623e3ebfea85b322f4cb802a6457d2)
- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844)
- datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-817eb7fe5c07b09131147f957abab9349fa67fed0733f3c727705cee91673d2c"></a>

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

<a id="canonical-8060e3db927273d8d77e4bb2fced8584c746d2a371780579a5803b427644861b"></a>

## Direct properties — datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 42822de4e91f / 3

<a id="canonical-57e2ab3b560dfad4133e8a72ee6fb2229bb48d6918671c98d9557b9b25798b4c"></a>

<a id="canonical-5602a29c8eb56860194eb94a4513ed7b296938aa71f3f2b7738e0659b300fc73"></a>

## provider_ref property — datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 42822de4e91f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-97a88cb14e10cd4fbfc91d0861e9adb9bc816979613e5219ae37ccfb3da97776"></a>

<a id="canonical-3d31bddd779a7f9c310ff25c7c7fc5c703603132df39c8fd36996d9d9853473f"></a>

## url property — datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 42822de4e91f / 5

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

<a id="canonical-93cddd6122d05389830fea6d890da8d052d16faa198bab76ec9cf5553a487cc8"></a>

## Next pages — datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 42822de4e91f / 6

- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3008851f2f5102ace609b107e0b66d768f8c02b5f20673bc4e5bfc5e48b77844)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-bac46d4ca622e0699daedc00ae9ae1977743b0a7044c299edec971222267190e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40415397fd5b3e4fe21a68af207832a2e0b638da75a023f555e47b9fd8259ebd"></a>

## datadog_receiver.use_tls.no_ca — datadog_receiver.use_tls.no_ca / d713202bcab7 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-7241d05aab97e71040e3653bc7552f4bc7ac3ac847d65183ed5db6d89ed82052)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- datadog_receiver.use_tls.no_ca

<a id="canonical-8a89e65daa4a2c84b0b05f9f32c5d7f0c0ac0edb524e3c6f4a47b6b4947d77d1"></a>

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

<a id="canonical-758597f75d9646088978775b2d4fcc863be19e66cf95b316731658aa06403be7"></a>

## Direct properties — datadog_receiver.use_tls.no_ca / d713202bcab7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-783f1d6dcd28bccd7713358e0d142ad12a36580596a76eb06d5c4c13616634dc"></a>

## Next pages — datadog_receiver.use_tls.no_ca / d713202bcab7 / 4

- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-7febd75ebadf2dd3f6e8fe8e27acd81d2605a234fc2996e32fc43b7f56be143f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b7191829d34da96f0435f5aec41d4e561e626957b75cdf343f72630ff0e93240"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9212bf0584ea6eeacff21fc5e7318fbbdd12bdb6f8093434dc02da6e1fe6fa68"></a>

## dns_logs — dns_logs / 68c78a35e425 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- dns_logs

<a id="canonical-3c0cf4cfae7a37e36a098964bdb7ac68726ba850bef0307a663b00b39428e737"></a>

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

<a id="canonical-08f88081c3f3dea5121945ad34ad2686b1165265e871f5b8cf7dfbb1e90ce4f3"></a>

## Direct properties — dns_logs / 68c78a35e425 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4105b16f4ec7ac92ac45c701ed7b34ba418cf4ce35e5c27a1dbf2fd851727338"></a>

## Next pages — dns_logs / 68c78a35e425 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c91ba8066a71a57fe39f5a3d829a0a3d500b58660be7b47404d4443534be6e40"></a>

## gcp_bucket_receiver — gcp_bucket_receiver / e74312a9ea85 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- gcp_bucket_receiver

<a id="canonical-1128734e088b037c58c9603f737bff23c86df51c827cfc4f1932c9f3cf5ed364"></a>

Type: `"single"`. Computed.

GCP Bucket Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b14520a803d82b05d0f43fe657528d97073bc4cc09189b6d36e6a9268b3c37cb"></a>

## Direct properties — gcp_bucket_receiver / e74312a9ea85 / 3

- [batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad): complete subsection reference.

<a id="canonical-c6194b69f5e744f8b3d20234e8adc6281e06ba05e9ee2ac5fc7e84896cf86492"></a>

<a id="canonical-a6ccc39b29c0c329022464e549113675778a35fb1c225b2318c00cc4f51c41af"></a>

## bucket property — gcp_bucket_receiver / e74312a9ea85 / 4

Type: `"string"`. Computed.

GCP Bucket Name. GCP Bucket Name.

Upstream description:

GCP Bucket Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0): complete subsection reference.

- [gcp_cred](data-sources--global_log_receiver--reference--group-002.md#canonical-c9ea57caaf346544494a27ca6dcaf584f1349a4a2ac00684a9e93848201b28fe): complete subsection reference.

<a id="canonical-b30cc930a5d576f35d05a444954b3019499fe4d01a2e77b41464183404e29af4"></a>

## Next pages — gcp_bucket_receiver / e74312a9ea85 / 5

- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0)
- [gcp_bucket_receiver.gcp_cred](data-sources--global_log_receiver--reference--group-002.md#canonical-c9ea57caaf346544494a27ca6dcaf584f1349a4a2ac00684a9e93848201b28fe)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c72ce5aa8d3ee8841314ca4b455f51ebc0d19089da2c36984593591932c3ee3a"></a>

## gcp_bucket_receiver.batch — gcp_bucket_receiver.batch / e0f552ff6ffb / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- gcp_bucket_receiver.batch

<a id="canonical-b60b7fc7098719dcbb05bd18ccf5b42822e76a931101db459cea2c6c927d3a99"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-f898684efb75cc4a3988ddc5134faf9e5ae0a08417cc17b78165093bcf17385f"></a>

## Direct properties — gcp_bucket_receiver.batch / e0f552ff6ffb / 3

<a id="canonical-81ad38888e85f02ca1447590b568cf903a758c5464a0e1ba37a1e39916dc59fc"></a>

<a id="canonical-fce5a3dc5a121ea70adff916ac48b7e21c6cd22ff78d839d708925e34faa92f1"></a>

## max_bytes property — gcp_bucket_receiver.batch / e0f552ff6ffb / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

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
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-762f8b5f676c0fb349c2276a66277fe4066570bb005cbb52972938536bb59458): complete subsection reference.

<a id="canonical-6b9291527985f632a629818a6d9e9bc7f3ec23ff57533fb170773a5d3d8636ba"></a>

<a id="canonical-00a7dca32e1bae30c6feee79bbd0f07983b5f9137c576e462aebdbe7e78de333"></a>

## max_events property — gcp_bucket_receiver.batch / e0f552ff6ffb / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-6a45c1ea6d9cb7d119bb27c4c6bdd4986f16505a108f3375b98979c96de57977): complete subsection reference.

<a id="canonical-c4ef0f3ed46534b17bda510943e96532d5971cd421dc898e97bb2d4cccf1f603"></a>

<a id="canonical-14557c4c851ee2e019ec7c25bbc24c61ba70141476c70665458bb3622f57a40e"></a>

## timeout_seconds property — gcp_bucket_receiver.batch / e0f552ff6ffb / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-0231819e92365e375fb7ad5759bf330d30f92dd05972f2d1284f54b8862a16ad): complete subsection reference.

<a id="canonical-ef48bea4aa5e029ca2beab0dfabac9a3dc274c195c8b3155931070f4a856476c"></a>

## Next pages — gcp_bucket_receiver.batch / e0f552ff6ffb / 7

- [gcp_bucket_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-762f8b5f676c0fb349c2276a66277fe4066570bb005cbb52972938536bb59458)
- [gcp_bucket_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-6a45c1ea6d9cb7d119bb27c4c6bdd4986f16505a108f3375b98979c96de57977)
- [gcp_bucket_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-0231819e92365e375fb7ad5759bf330d30f92dd05972f2d1284f54b8862a16ad)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-762f8b5f676c0fb349c2276a66277fe4066570bb005cbb52972938536bb59458"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6da63984a9d8c7698e838f43974019d85b4bbc7f451947c0aea72374547d6aa"></a>

## gcp_bucket_receiver.batch.max_bytes_disabled — gcp_bucket_receiver.batch.max_bytes_disabled / c590b227a09c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- gcp_bucket_receiver.batch.max_bytes_disabled

<a id="canonical-edd512431d149cc6d8d74a94071804e1c087f519285a4fd4da99a727efda1084"></a>

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

<a id="canonical-1599aee9c06eb6a48a0be419503874af8704701b57469e66d590429e354797c7"></a>

## Direct properties — gcp_bucket_receiver.batch.max_bytes_disabled / c590b227a09c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a68acc058248c5c83b4b8e1ed09e14a206bdafcf8d0e7c98abd1e0cc424e7b58"></a>

## Next pages — gcp_bucket_receiver.batch.max_bytes_disabled / c590b227a09c / 4

- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6a45c1ea6d9cb7d119bb27c4c6bdd4986f16505a108f3375b98979c96de57977"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b85aef34d014fc51527c69034881fa673a4c9a6d602a533b90f2a0e637eccc53"></a>

## gcp_bucket_receiver.batch.max_events_disabled — gcp_bucket_receiver.batch.max_events_disabled / 1a208767b295 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- gcp_bucket_receiver.batch.max_events_disabled

<a id="canonical-41278e2ff25ec6ccef86e8017b4964126ce35de5c548450e72cf852a3cb608c3"></a>

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

<a id="canonical-e743465915d3a543c08e3b4793210b6578c40dd7401fc86e9a52e196e80ef22e"></a>

## Direct properties — gcp_bucket_receiver.batch.max_events_disabled / 1a208767b295 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a53d92115b3f3cf5c9e4ca2fe45b869599aa4cb6ccedbf8ba093cfbb4174a55e"></a>

## Next pages — gcp_bucket_receiver.batch.max_events_disabled / 1a208767b295 / 4

- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-0231819e92365e375fb7ad5759bf330d30f92dd05972f2d1284f54b8862a16ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ccdce507ef85dd9649ca7a5b0986ddd1f566cb631ef00dfb02aaa3f0058b659"></a>

## gcp_bucket_receiver.batch.timeout_seconds_default — gcp_bucket_receiver.batch.timeout_seconds_default / be491654eb3b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- gcp_bucket_receiver.batch.timeout_seconds_default

<a id="canonical-22434fc273c739d43c31bc44965c48c9cfb522ffd013ed5ca75d4bdeb9aeab73"></a>

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

<a id="canonical-67484da2fd7bb7eba2a46dc996fb22497b9951bbd0a0bc8c4b365f17a5a6f914"></a>

## Direct properties — gcp_bucket_receiver.batch.timeout_seconds_default / be491654eb3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72dead844849dea32926e0b18791447d71774cca9ad2780a94077eeb4f012728"></a>

## Next pages — gcp_bucket_receiver.batch.timeout_seconds_default / be491654eb3b / 4

- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-94b526786afe3c76e6a5ac3722a958d0149ab0efc0dbc840ff5f8ef0edcb8dad)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad5f32fd315a462dff537a74198cc647f95020878926089459021ba10adbf068"></a>

## gcp_bucket_receiver.compression — gcp_bucket_receiver.compression / 323dce26c7ad / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- gcp_bucket_receiver.compression

<a id="canonical-fbfc0f6fb8c878954065cdeb459f311f202261ae70b9fe876789f9bf680d0705"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-19aa0167157ceeba89ce40993b3f1f4da0226f4b61b88b49e76265f42f5f51c9"></a>

## Direct properties — gcp_bucket_receiver.compression / 323dce26c7ad / 3

- [compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-53bd0a30876a7751226894776e815caa4186984374b29a4e6ed1cdd6a3efa40f): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-69b30389c0f9e065d6ac67f6aa548729cfd49d39d37a598f861dbc057831460f): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-2c8bec669e7a15bb2b1f3000a745cfb30c7282b69494dbe0e47efa302f17d934): complete subsection reference.

<a id="canonical-3a173b554c406075517c4ba9f74dacda5f9d3a6f01181bf28754860573dfc04c"></a>

## Next pages — gcp_bucket_receiver.compression / 323dce26c7ad / 4

- [gcp_bucket_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-53bd0a30876a7751226894776e815caa4186984374b29a4e6ed1cdd6a3efa40f)
- [gcp_bucket_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-69b30389c0f9e065d6ac67f6aa548729cfd49d39d37a598f861dbc057831460f)
- [gcp_bucket_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-2c8bec669e7a15bb2b1f3000a745cfb30c7282b69494dbe0e47efa302f17d934)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-53bd0a30876a7751226894776e815caa4186984374b29a4e6ed1cdd6a3efa40f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-873b62cc23d0bda6d5fad849aef8ea624ba56dc51b912eacd11bdbb9c113bfcc"></a>

## gcp_bucket_receiver.compression.compression_default — gcp_bucket_receiver.compression.compression_default / 420f3848c17a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- gcp_bucket_receiver.compression.compression_default

<a id="canonical-847f5fd5904cda37f611bf0a290cff3ed2c5fbac75eb3d96739cdacd63ad8e4f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-105395f76cf5c23dc5874358f0768ea2125d67aa4d82da690d3fc7e7fe6e3ecd"></a>

## Direct properties — gcp_bucket_receiver.compression.compression_default / 420f3848c17a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-857209120546c01674daae2a7a942de91289c70a7e3a8d0eb1d835c958be66ec"></a>

## Next pages — gcp_bucket_receiver.compression.compression_default / 420f3848c17a / 4

- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-69b30389c0f9e065d6ac67f6aa548729cfd49d39d37a598f861dbc057831460f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a614735e00f0f3e1c8fad0c0ba7466628eefb198899a0c3c92f20837c552e6a9"></a>

## gcp_bucket_receiver.compression.compression_gzip — gcp_bucket_receiver.compression.compression_gzip / e033778ae264 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- gcp_bucket_receiver.compression.compression_gzip

<a id="canonical-0f5423887de056959545ed6540659231afd490d2f4bdb9436b9a502dc6e7da4f"></a>

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

<a id="canonical-de79c3866b4d420cb4b15be7e671c75ad21948895fbe5ce25fce25b154a498af"></a>

## Direct properties — gcp_bucket_receiver.compression.compression_gzip / e033778ae264 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c41cc3d791d021de558bfd514e76e2718afa4d2190d668f3e0dc1bec274c7999"></a>

## Next pages — gcp_bucket_receiver.compression.compression_gzip / e033778ae264 / 4

- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-2c8bec669e7a15bb2b1f3000a745cfb30c7282b69494dbe0e47efa302f17d934"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-128aa555d61d3b6e1bec419463ad5b301da1e472498085314f9d400ed5afbc2f"></a>

## gcp_bucket_receiver.compression.compression_none — gcp_bucket_receiver.compression.compression_none / 432d924353b4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- gcp_bucket_receiver.compression.compression_none

<a id="canonical-415d117efdb1a3cad59447afb8489555eeaaf24a6771f91e970bda40180b2d3c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-a732f63e317932443e800d9ac113d92f4b08ba5b9024a3c3dd3a2086fdce6169"></a>

## Direct properties — gcp_bucket_receiver.compression.compression_none / 432d924353b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4194f150cd190d87a6ae29cbb2d5baff6406ac65d8c4e026d62f340d338574e4"></a>

## Next pages — gcp_bucket_receiver.compression.compression_none / 432d924353b4 / 4

- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-c7ec9136f89705016491b927a98cced2b9e5b1f2cb2de1f36160f83edc69d1c1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a88e3458a49115648d905a764a7f8fe47c6b8d3fccb64891427dfcf3275be04c"></a>

## gcp_bucket_receiver.filename_options — gcp_bucket_receiver.filename_options / d50839b72b71 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- gcp_bucket_receiver.filename_options

<a id="canonical-f4e52b63f489e36499eedc242691bf4039e4a60219e5760737085866128630d1"></a>

Type: `"single"`. Computed.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

<a id="canonical-9515a0e4ec199b7b0c928ba19dfacb0c81e5d1b97f19499c078c80a0e9e5e8d2"></a>

## Direct properties — gcp_bucket_receiver.filename_options / d50839b72b71 / 3

<a id="canonical-7df8ea6a557623f4bde559320187ac7a4f50ded1d98899e53f46dc18ecc5bdd6"></a>

<a id="canonical-430181a16fe7943181f2c343ca1097a15fca90e760644dc2cd5c9f6a8a18c696"></a>

## custom_folder property — gcp_bucket_receiver.filename_options / d50839b72b71 / 4

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-4eab30c6bf6e69bba7847d857c9b6f2d4be54ae411d6424ce3d4564f3d7c8775): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-fdf90d118aebfa8c59fb028daf15ceee703630450b4c740bad17358c613ff0aa): complete subsection reference.

<a id="canonical-0a6e036c8fc5f1a681b34ee77c9ab8f07786b728c5f78e48ff9ab4ab2f7b2297"></a>

## Next pages — gcp_bucket_receiver.filename_options / d50839b72b71 / 5

- [gcp_bucket_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-4eab30c6bf6e69bba7847d857c9b6f2d4be54ae411d6424ce3d4564f3d7c8775)
- [gcp_bucket_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-fdf90d118aebfa8c59fb028daf15ceee703630450b4c740bad17358c613ff0aa)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4eab30c6bf6e69bba7847d857c9b6f2d4be54ae411d6424ce3d4564f3d7c8775"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-080c6442b4d365426571435f8ef80e45e1f7907e4e24c3fa663125308993c81e"></a>

## gcp_bucket_receiver.filename_options.log_type_folder — gcp_bucket_receiver.filename_options.log_type_folder / 5bb4ce1453a0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0)
- gcp_bucket_receiver.filename_options.log_type_folder

<a id="canonical-821ba9dc5ad1754203bb7e8ee30b8b48c8a2e83d29ec1e31bd644a42ec88e02d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for log type folder.

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

<a id="canonical-e56845630b16bb3388d34ca41181b0fe4ca3f622cffe447a7746b95eff10d565"></a>

## Direct properties — gcp_bucket_receiver.filename_options.log_type_folder / 5bb4ce1453a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc3e9d39e87be1a034245f0cf58e6dd7558be0d2fb76af451d6eeef4bda199b4"></a>

## Next pages — gcp_bucket_receiver.filename_options.log_type_folder / 5bb4ce1453a0 / 4

- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-fdf90d118aebfa8c59fb028daf15ceee703630450b4c740bad17358c613ff0aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-335e26943e3f8b740b67bcd2b32bab42127803ab76e958db6a39bb0fdd129a44"></a>

## gcp_bucket_receiver.filename_options.no_folder — gcp_bucket_receiver.filename_options.no_folder / 06b123734124 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0)
- gcp_bucket_receiver.filename_options.no_folder

<a id="canonical-4bd851431589b06ac623a71a18fe0eb7c56c3e4393b41f0f3fa0200c677b0299"></a>

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

<a id="canonical-e4946ac62ce6be018677b289e344144d16aee81c3af04e9cc47f72f35993b246"></a>

## Direct properties — gcp_bucket_receiver.filename_options.no_folder / 06b123734124 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17a24c72cb3168a0b23524d369cc26ae53cda39acee715bc303baca7288fa8c7"></a>

## Next pages — gcp_bucket_receiver.filename_options.no_folder / 06b123734124 / 4

- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-15239fbfdaebe06e08ea2c925638380e0000c0e6c0dde051c338a110559956e0)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c9ea57caaf346544494a27ca6dcaf584f1349a4a2ac00684a9e93848201b28fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fb264924d177daefced8884f41b8cc596a56b2270b0005163aa5d795e6b6189"></a>

## gcp_bucket_receiver.gcp_cred — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- gcp_bucket_receiver.gcp_cred

<a id="canonical-aa836cb86a1ce4924fdd598a108474605169187fdf189a89e3a72fcaae3c33a9"></a>

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

<a id="canonical-0ea796979f68fc93511a761057151a663c45ca7554dda6737985f5f66529c97b"></a>

## Direct properties — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 3

<a id="canonical-1aa6ded0200a2f42c76148303d6419283bf2a97c68afb62fb20a7f76a2f483d3"></a>

<a id="canonical-b187a073e1a664c8fb3e9517c1b56343c565180602fff5c07fff746cc28b9b0f"></a>

## name property — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 4

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

<a id="canonical-4e867ba467d430eb854764492612e68ba4a413d36b6baf1eb0114ca7debebdae"></a>

<a id="canonical-1e9cd51eda4eea5acd37b0d3fc206147265935d99a0b4666234eb3142a7b21a1"></a>

## namespace property — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 5

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

<a id="canonical-11470f0aaf7da9afce507a3a7512eff548a9f896365a6de14b2a0fac0636acd9"></a>

<a id="canonical-d318c27d03201403fc2503a415f36ff282c6bf101465e3d38c3fa8928be310cd"></a>

## tenant property — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 6

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

<a id="canonical-cc9a082b098098d366b851b567530931962c818b58d59e877477c57f4158f141"></a>

## Next pages — gcp_bucket_receiver.gcp_cred / 1ad880e60862 / 7

- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-58684b364e0b8e5011351cd9b3602148b448064d1831c1a27d06df7e605452b2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e813775c669b80516b1261fa87e907be7cccc54dcda73445ea3a20bd80eece7"></a>

## http_receiver — http_receiver / 2b5b33f61aad / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- http_receiver

<a id="canonical-1128fb5d3d174ae457010bed22d2085ce139b22d5a8cc9a82b46392b50accfab"></a>

Type: `"single"`. Computed.

Configuration parameter for http receiver.

Upstream description:

Configuration for HTTP endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_choice": "[\"auth_basic\",\"auth_none\",\"auth_token\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-ce498db6568e3262dced5a97e3e4fdf2838873ba905039f4a12f8917d6e5829f"></a>

## Direct properties — http_receiver / 2b5b33f61aad / 3

- [auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045): complete subsection reference.

- [auth_none](data-sources--global_log_receiver--reference--group-002.md#canonical-f2d65e99dd5506f6a302b666009b3fc6acfe9d8a71b0716c0d8681582c3e9c8f): complete subsection reference.

- [auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06): complete subsection reference.

- [batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962): complete subsection reference.

- [no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-e3aaf3b727494a85b078661c9bb0be0422abd4b2ac37efada8cbf6b9841ccf2a): complete subsection reference.

<a id="canonical-a853fd4459b3bd75e7e8aabb18dbef36d8ca17e39d1e568b827290abda57a64c"></a>

<a id="canonical-2fa5cb392c3d5fb12240f6b6ff3e28f41f196457eb0b3bcb91efd5d83b23ac27"></a>

## uri property — http_receiver / 2b5b33f61aad / 4

Type: `"string"`. Computed.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc): complete subsection reference.

<a id="canonical-e281763646b25465e98061b6685234424c1b4b8cce30ef5741245a228a1fbcd5"></a>

## Next pages — http_receiver / 2b5b33f61aad / 5

- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045)
- [http_receiver.auth_none](data-sources--global_log_receiver--reference--group-002.md#canonical-f2d65e99dd5506f6a302b666009b3fc6acfe9d8a71b0716c0d8681582c3e9c8f)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- [http_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-e3aaf3b727494a85b078661c9bb0be0422abd4b2ac37efada8cbf6b9841ccf2a)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eec48754afceb9a32a4e7a69fa4c7cd93719325a145176ca515852be118a1ac5"></a>

## http_receiver.auth_basic — http_receiver.auth_basic / 4c7656f899aa / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.auth_basic

<a id="canonical-58acaf9a0fbd810a4eb6afb7f9dd39d3156cd8752b7e8d7eea1666f869a922ee"></a>

Type: `"single"`. Computed.

Authentication parameters to access HTPP Log Receiver Endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c641902d24948851cb85d0e9d170cf4f58c52283f208bbede44c1166a92eed00"></a>

## Direct properties — http_receiver.auth_basic / 4c7656f899aa / 3

- [password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db): complete subsection reference.

<a id="canonical-3cc30d5948f55a5dc362c3702ad64e11970e515ebb389ae15e90bbef5c9dca6f"></a>

<a id="canonical-2b4fd339b9d777aebc9570443145aba3ff31083d2cc47625f5eccbe20e06805b"></a>

## user_name property — http_receiver.auth_basic / 4c7656f899aa / 4

Type: `"string"`. Computed.

User Name. HTTP Basic Auth User Name.

Upstream description:

HTTP Basic Auth User Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-78cd330c9cfcc3120226642e79689e1c154bee5f6f443cfe81d6de03b2629f78"></a>

## Next pages — http_receiver.auth_basic / 4c7656f899aa / 5

- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d374219cf50f7151732f482ca658e81ab9dfe6bac301d13987bb1a650d813dd"></a>

## http_receiver.auth_basic.password — http_receiver.auth_basic.password / 5d98e8ab6ad3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045)
- http_receiver.auth_basic.password

<a id="canonical-c0125065ecf75fa9f1dddc29d858199d093e8dbf998aa8adb276ea2766de3d81"></a>

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

<a id="canonical-e86312c5c7dcd453320111e752dc552ec97a50f33b366c6a5ffee94c1c2d6c46"></a>

## Direct properties — http_receiver.auth_basic.password / 5d98e8ab6ad3 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-742916f18a21d9e402d3d92a123589551839c07b12071903f6db9d77c52d1dc0): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-21d79609e5ebd6c0b538dfdb408960e769c2d21bb436ed91eff34a817fe0fc5d): complete subsection reference.

<a id="canonical-de73ba2170b6afd1b98bf5f31b8cc6b515f9e0664c1491f59f3de24a635a47f8"></a>

## Next pages — http_receiver.auth_basic.password / 5d98e8ab6ad3 / 4

- [http_receiver.auth_basic.password.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-742916f18a21d9e402d3d92a123589551839c07b12071903f6db9d77c52d1dc0)
- [http_receiver.auth_basic.password.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-21d79609e5ebd6c0b538dfdb408960e769c2d21bb436ed91eff34a817fe0fc5d)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-742916f18a21d9e402d3d92a123589551839c07b12071903f6db9d77c52d1dc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-465ff20a9bcc271b0bed04b9110633ae8a28148a729329d56d3e46a8bc87c717"></a>

## http_receiver.auth_basic.password.blindfold_secret_info — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045)
- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db)
- http_receiver.auth_basic.password.blindfold_secret_info

<a id="canonical-21cb144e32d156083fb56db24f7dc3213ebc85c2dbb50ea2105abf966f38d7c5"></a>

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

<a id="canonical-22ecbc2c84bc78ce22bb6e66ea7a9ee3c5fc1a1cd2eaac649b3c02919f8a72ca"></a>

## Direct properties — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 3

<a id="canonical-3c52ed6dcb408ff9a11ce0a703f3ecd53b22ca4457331178d14522aa06cb2cd0"></a>

<a id="canonical-eafb6f39fb9ec4ddd43d9e69a0f0af871c00ad0788827f67b5ca010246418730"></a>

## decryption_provider property — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 4

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

<a id="canonical-5089fb9a7f79ab2b14afa16b30872727260979e05cf65eee9fb6a0dbac6fe62f"></a>

<a id="canonical-75734497fe5b2cf698972dbd48aa1277e5436d5b9d1fd5394b3e9815c56ceda1"></a>

## location property — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 5

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

<a id="canonical-caea25fba2ebbd7984ee329cb830ecf71145a968e0cb59280128314a90a95132"></a>

<a id="canonical-757007df2ca82ddec3ea942b6e5e3543a2a2679605b469da7be7cadbac93ec7d"></a>

## store_provider property — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 6

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

<a id="canonical-39a946d963d040e19b1d72146496c1aa9b10813143fbc0e6cd5ad3e8078121b8"></a>

## Next pages — http_receiver.auth_basic.password.blindfold_secret_info / 41daae5ea99b / 7

- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-21d79609e5ebd6c0b538dfdb408960e769c2d21bb436ed91eff34a817fe0fc5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e16992fa3aadd90aa8fe4fc0e4c3238866caa48b4c8a70f7f56fc7b46287d9c3"></a>

## http_receiver.auth_basic.password.clear_secret_info — http_receiver.auth_basic.password.clear_secret_info / f44dec125656 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-3d302c4f7d55ca389d4ff5f95c7881c1f886c6fd3e6fe98f1b9bdbfb5b223045)
- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db)
- http_receiver.auth_basic.password.clear_secret_info

<a id="canonical-0760a48cc9e4330a5ef077135d96fd56583d3308cd101729e9e534c0b4953135"></a>

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

<a id="canonical-a695b1ed820417e12aad6c7eb657145ae22a3b77339707271db93d4e37df2341"></a>

## Direct properties — http_receiver.auth_basic.password.clear_secret_info / f44dec125656 / 3

<a id="canonical-b7dc171f74029b26a6eea8218796d2f55d42f2a4e81674506858516e78d724fd"></a>

<a id="canonical-3da727ea958e7658fa95b9cfb4d78d97b2d36d081b361952bd94882f6539f4b5"></a>

## provider_ref property — http_receiver.auth_basic.password.clear_secret_info / f44dec125656 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6ad8789e73c15e6f8d1a77233995b779b0f924bbd301bfba05895af0f3080e11"></a>

<a id="canonical-661c8ef8bc66616dac6f4d2d9109fde1a734fc6f32eff6239105a44a027469cd"></a>

## url property — http_receiver.auth_basic.password.clear_secret_info / f44dec125656 / 5

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

<a id="canonical-9a358f286bd912f0a9d0ca4b9b940abdf69bb08830e6245affb26e294facd417"></a>

## Next pages — http_receiver.auth_basic.password.clear_secret_info / f44dec125656 / 6

- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-bff7d32117accd38d7a8877c13feafdd6a40b2569933ba97bd30f99aca4186db)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f2d65e99dd5506f6a302b666009b3fc6acfe9d8a71b0716c0d8681582c3e9c8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6005fd78556b5872ba4df201035a521b53fd1030888e0fce1be78ff181f0709"></a>

## http_receiver.auth_none — http_receiver.auth_none / cb889865bd61 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.auth_none

<a id="canonical-01513ad57d8235f3bfebd49bfca8fba002802cff8dc3fb29186ab4cb5dd58a25"></a>

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

<a id="canonical-e893667ee6c4f90d5b4e0a9146a08770d7a1bf914eb37605bb200ba12bf1279c"></a>

## Direct properties — http_receiver.auth_none / cb889865bd61 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8d24cfdd6960780ecffda010cf4307c01d1f5112ec51867b0c98f906b61758a"></a>

## Next pages — http_receiver.auth_none / cb889865bd61 / 4

- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddfb3ec8d6df340f4c3aa499a7a2160766dd6b67b2d0ae04ab694d244ff301a7"></a>

## http_receiver.auth_token — http_receiver.auth_token / 9fa6d4d59d94 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.auth_token

<a id="canonical-a30225a85f2219910dbf42f18d8ab75bbe92ee9bf83086659cc65a94d26a2e26"></a>

Type: `"single"`. Computed.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-75692d5387bae4ff036a5ac72aff434795bb4bceb69b0b6427ec5234e460376d"></a>

## Direct properties — http_receiver.auth_token / 9fa6d4d59d94 / 3

- [token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f): complete subsection reference.

<a id="canonical-4e0c2bb181df25c5b344e7b1a737d79299cdf7fd304eb3fc85feba0bd2c81096"></a>

## Next pages — http_receiver.auth_token / 9fa6d4d59d94 / 4

- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df66dec24082b7fbe348f331b33d894c4f1cad613c30222491af75e6a79f295c"></a>

## http_receiver.auth_token.token — http_receiver.auth_token.token / 0714b3a4f964 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06)
- http_receiver.auth_token.token

<a id="canonical-212714edc1f0c237204594a28c58e49ab0f6e860365434f4e797b451dc54135d"></a>

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

<a id="canonical-3dbd8d600b22a5455a2bcac419a3aab129f567160c4e680f8a98740b3e745eee"></a>

## Direct properties — http_receiver.auth_token.token / 0714b3a4f964 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-da7538bdc34da2de2b7d62bdae17f16b008eda000dd0b94b619ee1e09e524bc9): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-212db569f3bf230a7a205a82377e5966255182cb257769e0475d257755c162e6): complete subsection reference.

<a id="canonical-3a4b2f707fcf9795230da20cc87239561bc93144e54e812df97da98c6c374228"></a>

## Next pages — http_receiver.auth_token.token / 0714b3a4f964 / 4

- [http_receiver.auth_token.token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-da7538bdc34da2de2b7d62bdae17f16b008eda000dd0b94b619ee1e09e524bc9)
- [http_receiver.auth_token.token.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-212db569f3bf230a7a205a82377e5966255182cb257769e0475d257755c162e6)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-da7538bdc34da2de2b7d62bdae17f16b008eda000dd0b94b619ee1e09e524bc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-143beb4c38d3b7707ca27520080d6ab769891c551738f0d56223516248620d1e"></a>

## http_receiver.auth_token.token.blindfold_secret_info — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06)
- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f)
- http_receiver.auth_token.token.blindfold_secret_info

<a id="canonical-461b1cc272eb18994019e3576da6fc74eadfbad679bda4e2635a78d0c1e51672"></a>

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

<a id="canonical-0bd41bf661e47e15175e255ed312b425c0bd58206518be760756f8ad74524ff5"></a>

## Direct properties — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 3

<a id="canonical-7139a81111a08aa2ae64a8e6f57e411b98de84ce67ca03ae64db1dae64191976"></a>

<a id="canonical-36e6ea6ffcfc64a9c4692b584edf219c416d50bf69e156c92c0922ceb752d63e"></a>

## decryption_provider property — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 4

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

<a id="canonical-2f6c7fb752f6ef91c882f02751482b8eab988d37381e19993f27291386843d3a"></a>

<a id="canonical-dc7250961d4f918027bd1c9c5f8ba9d0d40a30c53017c5ca91937d451db0d0ab"></a>

## location property — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 5

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

<a id="canonical-880478675d62ab19f6102e7d5b200e787677ee63499eab852c414c472875e819"></a>

<a id="canonical-fae23ec00041ecb1a4738646e9a2ec8c2016cfc46903c48f715abaa532be3b7c"></a>

## store_provider property — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 6

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

<a id="canonical-e4c0d4d7a7079e9a2698794a09e2ab8a114763829ca1a63097bec27e90ece2cf"></a>

## Next pages — http_receiver.auth_token.token.blindfold_secret_info / 09d109fc4666 / 7

- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-212db569f3bf230a7a205a82377e5966255182cb257769e0475d257755c162e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52a84dab16f2f9b0eb272cff6551485300623a6052440cee7a2971b396cab9fd"></a>

## http_receiver.auth_token.token.clear_secret_info — http_receiver.auth_token.token.clear_secret_info / e86263761157 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-fb3b3586caa5ed75842ef652dc7dc256a52fab755d9cf26ae8207129c4310d06)
- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f)
- http_receiver.auth_token.token.clear_secret_info

<a id="canonical-51132397636dc8f3e8697e429737107d1b58bd09bdb750448d0b6316e347db1a"></a>

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

<a id="canonical-eacc7e3626cb7eba5f9e70331c4ff962ed343cfd05ef100732cea72172deaa61"></a>

## Direct properties — http_receiver.auth_token.token.clear_secret_info / e86263761157 / 3

<a id="canonical-df0027f393ad7732156561351b77d9d36ba81c887178d1d96481b4d92fa89869"></a>

<a id="canonical-42f3c9b27932a930af14374686524623ab904d8c7f04f7f14626fec53fef6531"></a>

## provider_ref property — http_receiver.auth_token.token.clear_secret_info / e86263761157 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ff152eb7f27ee50bcdd347e166e3ed1c907751210713d49bae9bd59bef073f91"></a>

<a id="canonical-4f3542f829cde555e89feda11874f07b789079c298a1f59d666ad4baead9f7c9"></a>

## url property — http_receiver.auth_token.token.clear_secret_info / e86263761157 / 5

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

<a id="canonical-f71d342e912b8a8cc0ebbf2353475d6854bde871929f89394b45caa2c4a82826"></a>

## Next pages — http_receiver.auth_token.token.clear_secret_info / e86263761157 / 6

- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-a17c99a0b3c6b02dd2c01f57cb53b380e7dd225174adba4a84dc490e3b32a08f)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a28f1ee1539d233662b8209a2cee0a81f1fb7f1b31c4caee68d4fc16d017b2d6"></a>

## http_receiver.batch — http_receiver.batch / 5a4001d9fbae / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.batch

<a id="canonical-cf0bc96796259fa40eafcbd02e580a711195bfa0959a1d485c68356561b21890"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-9d9e034323f8535b181caa067f00bec204d4e15f498c7da331172a67eb5d5c5b"></a>

## Direct properties — http_receiver.batch / 5a4001d9fbae / 3

<a id="canonical-d8c364a6a300974a5800494b4b190177aca6112710e1a57944f2564814832330"></a>

<a id="canonical-a925b1f9d77d920f9f6d8d6d19173c5439e1444d5fcb0d8b48d6dad248c27b44"></a>

## max_bytes property — http_receiver.batch / 5a4001d9fbae / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

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
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-9fc6a8ca510ed80063c8999508888dba28b03c5e83e5c0d702bcc2255964faa3): complete subsection reference.

<a id="canonical-a9470306f908b1367a9c4143e0d7043208823c5491d02b87c04b93535cb1da12"></a>

<a id="canonical-82f5963685374f9f4402c7b96b47fd1e1e546ddb6fd901c56c4c2191c379d5e8"></a>

## max_events property — http_receiver.batch / 5a4001d9fbae / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-7595719a54f74f12e78a4f3b7f6f0d785ac128a17e39e4b886c9b1508a14802c): complete subsection reference.

<a id="canonical-120df0f073d7a4a7e6fae3730d1d41c2a7238895459c504c90857edd353bc6bf"></a>

<a id="canonical-7c98aaa8e4f827612d3fc728f518431c0d3ec407848f3f7deeb4a22efe517ed6"></a>

## timeout_seconds property — http_receiver.batch / 5a4001d9fbae / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-f110068ecff5526412e84df55b8269c1fdd92791435ab9eccf142fa854c650c0): complete subsection reference.

<a id="canonical-4c8afb350ecfb8703ece07ff2b906a3663dfe9c61dae5b44722e92f456a0ca4f"></a>

## Next pages — http_receiver.batch / 5a4001d9fbae / 7

- [http_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-9fc6a8ca510ed80063c8999508888dba28b03c5e83e5c0d702bcc2255964faa3)
- [http_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-7595719a54f74f12e78a4f3b7f6f0d785ac128a17e39e4b886c9b1508a14802c)
- [http_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-f110068ecff5526412e84df55b8269c1fdd92791435ab9eccf142fa854c650c0)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-9fc6a8ca510ed80063c8999508888dba28b03c5e83e5c0d702bcc2255964faa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b255b300279667660b67ef6adabfcc7cb159bd1fa6f9ce34f6299968a5a22b93"></a>

## http_receiver.batch.max_bytes_disabled — http_receiver.batch.max_bytes_disabled / c0750accb404 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- http_receiver.batch.max_bytes_disabled

<a id="canonical-0301b27b63919cd15c5bf72226d99de8090673c65d63b875b9885e637b6f1fd5"></a>

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

<a id="canonical-40d38c5c574eaa80476d734ca77c46d406bad910af95e976224c5d7055a15679"></a>

## Direct properties — http_receiver.batch.max_bytes_disabled / c0750accb404 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0309a94250d272a98666678d16d0c3850ef4c9505aad81ac8cc18767fd4ac595"></a>

## Next pages — http_receiver.batch.max_bytes_disabled / c0750accb404 / 4

- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7595719a54f74f12e78a4f3b7f6f0d785ac128a17e39e4b886c9b1508a14802c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33b21ab29223772c51d5b14cbcfc68709c485ab346a771c6d339538e9ed8b504"></a>

## http_receiver.batch.max_events_disabled — http_receiver.batch.max_events_disabled / 7787f6f59c3b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- http_receiver.batch.max_events_disabled

<a id="canonical-477b5254a598775f0a0afbcfb3dc450666dcf59fc419258c1bd683e374b11b00"></a>

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

<a id="canonical-211dc4825ebaafa115a519700a0478c544284af0fa3aaad7e24d416a5ee7ec4d"></a>

## Direct properties — http_receiver.batch.max_events_disabled / 7787f6f59c3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36ba70e43a94d0626fc3bcfc25d2ddcb02cedea25679fa1994dab13636f82175"></a>

## Next pages — http_receiver.batch.max_events_disabled / 7787f6f59c3b / 4

- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
