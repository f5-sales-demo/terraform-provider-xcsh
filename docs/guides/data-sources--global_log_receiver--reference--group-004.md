---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-a5e36b08811d2d0da672aa5b7c987271115b8cf42e652aecabe8eba4fc62f8d6"></a>

## qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72)
- qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0fd96d72a5ac05a5c8841ea837aba13bb49e3b9431f8f34725a204b3347f2064"></a>

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

<a id="canonical-80134e3638524246469a4a726f0de85554293db25a10e6ea87e3585a0d547b22"></a>

## Direct properties — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 3

<a id="canonical-3af39a7d4779b04e9193ed4c07e0dfc08edec8760cd2dbf97d1ccb13f2be653f"></a>

<a id="canonical-e61c64e3bf8c1fb41b46a8e9f29cda542b95bbb98f40c2ca87f1149293d581cb"></a>

## decryption_provider property — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 4

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

<a id="canonical-e9b8ca6ec72623cbba7b9ddba0081b9d314d0f89caf15294107782a102b5a58a"></a>

<a id="canonical-52a36ab4e3f07704bd0db8258ab762a44330c45a7581df8e5d41ad19c1079730"></a>

## location property — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 5

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

<a id="canonical-626be832aa74e93743b31b7c1bcde5fa9a5fd61ab51880d8a905c7abce039846"></a>

<a id="canonical-9ef9ac7e8ebc87acc733ae6ae99b9668cbad50db5a4b6fa5ab1142df61f61126"></a>

## store_provider property — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 6

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

<a id="canonical-f5228c186fb4c4756006e9b22013a85d28b85cd0d4cbd11196c0eec629c3c770"></a>

## Next pages — qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 03bcdb37f883 / 7

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7e3cb6f28d0dd357bbb36d620bd9e5d500b3984fe74f976f4e40ac289dc48e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-643fc98ba933531f8faa86d474c2fa7dd7f65c6526584da09c881d1b4ddc48c2"></a>

## qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info — qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info / e47ccdf14dd9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72)
- qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-a1936f621178a74aaa10b9928b1c78d2f62fa15acf4857d48d3a76f3de174e32"></a>

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

<a id="canonical-969544b60a706a1132b0372b29c6c3738cf1ea723ee4a59c8b20e6da06504d3d"></a>

## Direct properties — qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info / e47ccdf14dd9 / 3

<a id="canonical-1fb2d311a83b4845f1c0c7bed9a9a547828c166b1029d3e45600de696c8bcff5"></a>

<a id="canonical-8d8ecd19f716d93aca1fbc19f407b41d4f284d2164216ae1aab1caf86098ffb5"></a>

## provider_ref property — qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info / e47ccdf14dd9 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-df9593fc2f1f916508281b6afc2f3bdc3a33105cabcffd8c543e8725800af392"></a>

<a id="canonical-29e4aeb1962450070b0ebc445b3e48a8619b99d567d46a32f44c7b81ff212261"></a>

## url property — qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info / e47ccdf14dd9 / 5

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

<a id="canonical-ed1927a7611da6d6dbf3271e3a32614c0e4801b12f4f91e668153bf56d19f7fc"></a>

## Next pages — qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info / e47ccdf14dd9 / 6

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-795c32cedffd28a0fc7e0221a4b413d357c43cc33256809536d1fd168f30a0ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe8b108314a578a97688d4e93322ead5ae4ef7af31070727bb8453b63cc2deae"></a>

## qradar_receiver.use_tls.no_ca — qradar_receiver.use_tls.no_ca / d3116a64fcf8 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.no_ca

<a id="canonical-8adb64c76c96940ebe498a228b0ab2ace2556cc658f955660a2edd4501ffae1b"></a>

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

<a id="canonical-42f8e0b1a17b34bf6feddd7205a60d55de7dd83628fe31a42cbb773c5dca675f"></a>

## Direct properties — qradar_receiver.use_tls.no_ca / d3116a64fcf8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-373938726a9a1a319cf90633c6eecaf003bf86c6910dcb8f3fb4ca52262e6072"></a>

## Next pages — qradar_receiver.use_tls.no_ca / d3116a64fcf8 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54670a4f12b5d069db049d24e13807aff48c9c6253328b85c7535f4ecffa8d0f"></a>

## request_logs — request_logs / 7b1e2c7def5c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- request_logs

<a id="canonical-4e2818b42395658eeb7c84785af4bf6b05a6a0bb11bef566421f96787bc2b52d"></a>

Type: `"single"`. Computed.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

<a id="canonical-4cdef02b80e0a563419b8cf480faecc5b070b0c55cbe8d189d535aac9cca6000"></a>

## Direct properties — request_logs / 7b1e2c7def5c / 3

- [sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-e042e6b2259d84c9e2e40775e6194ad463edab00c852fd5166b1dbcc2f5ea442): complete subsection reference.

- [unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-6f9648877f61e6fc3d5d2ba82227f0df634cedd2863510e4a76e02da2f775fe3): complete subsection reference.

<a id="canonical-2508fb678823c6db7b9c7be9d7d223585a868306d794e049f66410821ffeae98"></a>

## Next pages — request_logs / 7b1e2c7def5c / 4

- [request_logs.sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-e042e6b2259d84c9e2e40775e6194ad463edab00c852fd5166b1dbcc2f5ea442)
- [request_logs.unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-6f9648877f61e6fc3d5d2ba82227f0df634cedd2863510e4a76e02da2f775fe3)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e042e6b2259d84c9e2e40775e6194ad463edab00c852fd5166b1dbcc2f5ea442"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-734b44888300a30a0fe6256a88cd55705d41813c6eaa69066ce59b2b55852454"></a>

## request_logs.sampled — request_logs.sampled / 7fed2ec967ed / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b)
- request_logs.sampled

<a id="canonical-2700c0cb72e82e3c51ba6bf17f14cdaa2efb7da8c2bd97f35dee1cef21c11658"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7a2c9a6bc9dcee9f1915bda052a3fed8e283c169665be781fd3e90deae402e84"></a>

## Direct properties — request_logs.sampled / 7fed2ec967ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da0e69ed9ff7767a088ef6fe25257468e67161cce5249610d0172b8f4d37c944"></a>

## Next pages — request_logs.sampled / 7fed2ec967ed / 4

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6f9648877f61e6fc3d5d2ba82227f0df634cedd2863510e4a76e02da2f775fe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffa1b78e26b6bf2a2dcc54890b79f4af5befa788aa9e963cdd785ef9e4fde69b"></a>

## request_logs.unsampled — request_logs.unsampled / e049bde6dbd9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b)
- request_logs.unsampled

<a id="canonical-b2799c59167c5d2a69da737ee2143a2369004e5fc66edd7b671afc294694dfbc"></a>

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

<a id="canonical-8f6bc1eaeff5a8c41571ddfe088b3ed53771c2ba05c22edb176804b1b59ac1c0"></a>

## Direct properties — request_logs.unsampled / e049bde6dbd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07205dcc27770fa9d49e96263633785d30c81998376376be1e552f09f541efff"></a>

## Next pages — request_logs.unsampled / e049bde6dbd9 / 4

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-78587f316fc037494d3bdda63e54544614ce095432c2e1335fabfd6ee00f891b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfb53e92670a475d21c80a8a38f973f4dbfd0c4148e9f1e77939da4d833f9b18"></a>

## s3_receiver — s3_receiver / 80fe805d74ba / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- s3_receiver

<a id="canonical-e6569dee683a809947ec9064923b603e50489bad872e1b88bf8fec828a194c03"></a>

Type: `"single"`. Computed.

S3 Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-900022704a7c8b6ecb6595239401bb9f263351534acf86c65c4f96c996d50412"></a>

## Direct properties — s3_receiver / 80fe805d74ba / 3

- [aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-03e6fd43f15e519a131f912af05e505d1c5d48334eded448dd580c4eff1accc3): complete subsection reference.

<a id="canonical-dc58aa1b3a12744c8c068e45ea5de745c9274d0996fa328687dad41a0c31f924"></a>

<a id="canonical-a06fd7ceaf520f31f48089e349ae4bde6a3b6e7c24015441b7a3e153b3655a40"></a>

## aws_region property — s3_receiver / 80fe805d74ba / 4

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a): complete subsection reference.

<a id="canonical-049ee00c58c3abf4894c8cf5e8a7def624b9d156daed9d444ec13a868e1210f1"></a>

<a id="canonical-08c8b7266ab06999a785a32a47f1374d0e2a9797c291974f179a811476fede07"></a>

## bucket property — s3_receiver / 80fe805d74ba / 5

Type: `"string"`. Computed.

S3 Bucket Name. S3 Bucket Name.

Upstream description:

S3 Bucket Name.

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
    "pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
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
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc): complete subsection reference.

<a id="canonical-6de8c812186b4917772cdf7d3c9c07d94dd0c111598b43f086b7850e39cb9511"></a>

## Next pages — s3_receiver / 80fe805d74ba / 6

- [s3_receiver.aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-03e6fd43f15e519a131f912af05e505d1c5d48334eded448dd580c4eff1accc3)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-03e6fd43f15e519a131f912af05e505d1c5d48334eded448dd580c4eff1accc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d82b7cc2c30f60e966b659201388349c66fba706e9bdcbbf00d2f5b28a11b904"></a>

## s3_receiver.aws_cred — s3_receiver.aws_cred / fd847510f45d / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- s3_receiver.aws_cred

<a id="canonical-e218bb346ee7f6d594c482696a312c164acbf6a5022d8cdaf4ad10f26ed6b855"></a>

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

<a id="canonical-4cd3e56f3df7dc79bf7ca766f66126492b23522a45f826800cc9864791a9c9e6"></a>

## Direct properties — s3_receiver.aws_cred / fd847510f45d / 3

<a id="canonical-3a8ab356ce3acc9cbceaa924a6f01b6227bbe3611c0f725459a95bd32608d2ea"></a>

<a id="canonical-9c53108705fde238f9c4d18525c37dd8df0950af923046958a60f99aa8ebc6e6"></a>

## name property — s3_receiver.aws_cred / fd847510f45d / 4

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

<a id="canonical-705c04f98bbe168cd7e12a6297276290851bc9ffde7109e684299bd5b97e988a"></a>

<a id="canonical-e0e604d6e3fe6d149aa7002a669614a21d57ee6e32d203363870a7e5b8a5537f"></a>

## namespace property — s3_receiver.aws_cred / fd847510f45d / 5

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

<a id="canonical-77a207df33189e84bafe6781179130c012031593edebccffad7c671c13b65c88"></a>

<a id="canonical-cb52b58be139427341f9763ca6ba7dc43bfcbe5e389f35dd7cf8235dfe55c193"></a>

## tenant property — s3_receiver.aws_cred / fd847510f45d / 6

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

<a id="canonical-1ebf7ad77e00efb1c73b65c3b01419b4ed83d47dfa19ebd419c78dc3b1762823"></a>

## Next pages — s3_receiver.aws_cred / fd847510f45d / 7

- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82c12d8c8ed4858ed040f6678a1343656d7c7bd9d6ce0a5627645ea8b048eff3"></a>

## s3_receiver.batch — s3_receiver.batch / 8ffc7bd44530 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- s3_receiver.batch

<a id="canonical-c8b264cf9f82e39dcc4ada282650ae84579eb4df611cf8ddfb539c8d8b353f9d"></a>

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

<a id="canonical-773278ff58f0f37f2fc9b6af2e0e225a59b3ff6b8f81b4583a7330f8764ccf21"></a>

## Direct properties — s3_receiver.batch / 8ffc7bd44530 / 3

<a id="canonical-6e0f86194c793191f0d93310c6fe5075487c4fc28cac7f00f2c326f06da8f329"></a>

<a id="canonical-f500079f9cdbc4cf276cb6ef5ccbb31a2f0e0fbbff81da8c8b4a60abd7c72d9b"></a>

## max_bytes property — s3_receiver.batch / 8ffc7bd44530 / 4

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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-300b9cab0bad6458077f685529067636834b542cb03b361228448fc4d2de019c): complete subsection reference.

<a id="canonical-6903799bd4d481785dfa88f505f7f3c15257bb5916c190b7598a4ec06a393a43"></a>

<a id="canonical-aac7fc2fd548199130715315992b52f26ea5a191d08477dd180239677cc45b9a"></a>

## max_events property — s3_receiver.batch / 8ffc7bd44530 / 5

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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-b527578b7309ec0450357b828c45b33c6e012a606172e7b081d168d73291981d): complete subsection reference.

<a id="canonical-60675b24dae7f21318aa3102eed3991118bae1fd59bdb28d47caeeefed8c4147"></a>

<a id="canonical-3eed6569b8dd22c45a251ba74e81fed9ad55e97fe58f4aa470964393ef61c2df"></a>

## timeout_seconds property — s3_receiver.batch / 8ffc7bd44530 / 6

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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-d29a7bed931425b67efa5be3f3a92bbf3519d50231bad2390e3dcf7844acbb55): complete subsection reference.

<a id="canonical-124676760bb825f1cfacf5fb9a4d869da87590ff05b3a9d14b965823a56a4676"></a>

## Next pages — s3_receiver.batch / 8ffc7bd44530 / 7

- [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-300b9cab0bad6458077f685529067636834b542cb03b361228448fc4d2de019c)
- [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-b527578b7309ec0450357b828c45b33c6e012a606172e7b081d168d73291981d)
- [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-d29a7bed931425b67efa5be3f3a92bbf3519d50231bad2390e3dcf7844acbb55)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-300b9cab0bad6458077f685529067636834b542cb03b361228448fc4d2de019c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0d8948dd1b0b940d286c698b45a341fa76872a544a709b7144e1007f51ffeda"></a>

## s3_receiver.batch.max_bytes_disabled — s3_receiver.batch.max_bytes_disabled / 241ceaaf8c04 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- s3_receiver.batch.max_bytes_disabled

<a id="canonical-4ddbdb5f8c994a75895518eb944f7aa20735e06945deca9c6b342b1cb6d4d488"></a>

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

<a id="canonical-667e45e5083c952771ad0bffb6a722b9ebecfd933a9a5d682fc0f851b2c44cea"></a>

## Direct properties — s3_receiver.batch.max_bytes_disabled / 241ceaaf8c04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a6311d6b797dad629064c79572f1eaabecf0e9e48371627c9fed1c2994de603"></a>

## Next pages — s3_receiver.batch.max_bytes_disabled / 241ceaaf8c04 / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b527578b7309ec0450357b828c45b33c6e012a606172e7b081d168d73291981d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01ec57df1f101c82272f8c78f4dc9259503494c63b2107412cab0b890d912f66"></a>

## s3_receiver.batch.max_events_disabled — s3_receiver.batch.max_events_disabled / cf7618da3f91 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- s3_receiver.batch.max_events_disabled

<a id="canonical-dbf83ad9a7d045024b2d2950ae6fa9169f5cbd418924e4730172de376726d0b8"></a>

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

<a id="canonical-d7ec9ac43cf36718accbff7d829613ad83de7255616e78fa5cad38bc96534547"></a>

## Direct properties — s3_receiver.batch.max_events_disabled / cf7618da3f91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-253dfb1d9981521a6d102890645cd426f2b130e9b7f3321d994108ab47f5e661"></a>

## Next pages — s3_receiver.batch.max_events_disabled / cf7618da3f91 / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d29a7bed931425b67efa5be3f3a92bbf3519d50231bad2390e3dcf7844acbb55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e99eb26718555573fda5b61b1d830f9d26e349180fd24d5beec334bbf4568f"></a>

## s3_receiver.batch.timeout_seconds_default — s3_receiver.batch.timeout_seconds_default / 9e5a7e0c67eb / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- s3_receiver.batch.timeout_seconds_default

<a id="canonical-6bd760ceb829fb4969103ab06bc019e3f5b499462ee31956344b3484e831df37"></a>

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

<a id="canonical-b9ad38e99fb4d0fb027191f27e06c83bdb4b03c8bff13cb4a66a30705beb01a8"></a>

## Direct properties — s3_receiver.batch.timeout_seconds_default / 9e5a7e0c67eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9d0cb2093b9cf733a65af299a7680b10f3a4810ea388a8be6c568b13e741afa"></a>

## Next pages — s3_receiver.batch.timeout_seconds_default / 9e5a7e0c67eb / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0cd7758280b5d971be10f2b7f88c69a737c7b1a03e5dfe4a292700c885e77f8a)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a64f34254e83ab95531653da0af796be00a4637bef9c7e25dd03dc087136fa5e"></a>

## s3_receiver.compression — s3_receiver.compression / 0ecf133014bf / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- s3_receiver.compression

<a id="canonical-c1fe7d2192cafb05ceb0b440ee5022b15206de9f8c4fee42c2e2a1e1361d9c27"></a>

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

<a id="canonical-d89c6db97ff47bfcaf990af6f4eded82332f4d14b8977cee4885c07ff002456d"></a>

## Direct properties — s3_receiver.compression / 0ecf133014bf / 3

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-82c97c694d77703791a27b7cd741b785f13153f7a99ec9ebf279bb0939c61f56): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-e6798ea0b2f5f44995966ac4caf485f3b8854882786b489733f8de3de0a9bd32): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-c0cb1873d2dad00cac68f56250918e5df508af23ebf7ca1da5c391f369d3c078): complete subsection reference.

<a id="canonical-fc3c9cef4e0c0ef95366f6d5f290fdb5489608a952c3bbefcc930803e17d4d80"></a>

## Next pages — s3_receiver.compression / 0ecf133014bf / 4

- [s3_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-82c97c694d77703791a27b7cd741b785f13153f7a99ec9ebf279bb0939c61f56)
- [s3_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-e6798ea0b2f5f44995966ac4caf485f3b8854882786b489733f8de3de0a9bd32)
- [s3_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-c0cb1873d2dad00cac68f56250918e5df508af23ebf7ca1da5c391f369d3c078)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-82c97c694d77703791a27b7cd741b785f13153f7a99ec9ebf279bb0939c61f56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e8f3097ba120ced8b2920c261eace0de00c1b894166c48e8f57c2b648c7dfb7"></a>

## s3_receiver.compression.compression_default — s3_receiver.compression.compression_default / ed8a2e366541 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- s3_receiver.compression.compression_default

<a id="canonical-9115ec844c3bcafebd60303c09e849af9384d8584cf769da79a90437a21ea604"></a>

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

<a id="canonical-db04b53d4cd2cfc7cb7f45afee7e4fe30baaad7be33acca3a0635df381b34d18"></a>

## Direct properties — s3_receiver.compression.compression_default / ed8a2e366541 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-978474b798a96608c068ce908c99c07c45cb3a4b0dddb1da9d0cd6112a345738"></a>

## Next pages — s3_receiver.compression.compression_default / ed8a2e366541 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e6798ea0b2f5f44995966ac4caf485f3b8854882786b489733f8de3de0a9bd32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c22ce9b6bb0eebd16a57b33961560220111cf3c867563f2595da84e28d8d8c2"></a>

## s3_receiver.compression.compression_gzip — s3_receiver.compression.compression_gzip / 77319fdcb8d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- s3_receiver.compression.compression_gzip

<a id="canonical-680365dfa46a41d9c2dbb40dbe2bd251486924c93ab5ccc586ede4530852766d"></a>

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

<a id="canonical-42b205b5f3723edc00f2b8cc73862bc9f984306d0712b135a067bc4e621da3ff"></a>

## Direct properties — s3_receiver.compression.compression_gzip / 77319fdcb8d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b67cb53cac749dd1bbf5173ad26685dfe2f08f24ac67222ae7893beda11b8109"></a>

## Next pages — s3_receiver.compression.compression_gzip / 77319fdcb8d2 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c0cb1873d2dad00cac68f56250918e5df508af23ebf7ca1da5c391f369d3c078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61ce51d7f65a1b1a2d29209dfc0d762218a54ce45877b3171e14e1c3e2f8c68d"></a>

## s3_receiver.compression.compression_none — s3_receiver.compression.compression_none / 2b2f4b1875b3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- s3_receiver.compression.compression_none

<a id="canonical-2a27204842edaaec9929d7b4802e982575f1613ac5006309c58905c720a9c536"></a>

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

<a id="canonical-a28691980d007038e636524da16d84ff5530750a0546c1c58def4badafb78def"></a>

## Direct properties — s3_receiver.compression.compression_none / 2b2f4b1875b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5ef594322bdc492dec61983a0cc47ce6e9bb529153549093bf2b6aae864e0c6"></a>

## Next pages — s3_receiver.compression.compression_none / 2b2f4b1875b3 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-651124667d8e2e0ab5c977a3a69114bf13450026279a737345b93e69918b7c23)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43015820e44e0cc3f0ac6d573da6a9e9351dfc676bc0b3fd26a3f814b54dccf0"></a>

## s3_receiver.filename_options — s3_receiver.filename_options / 78e70b7436c6 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- s3_receiver.filename_options

<a id="canonical-4eb0d9f5b081ae8232b16d159840e808b152790389fb38557affa33ad738d2c6"></a>

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

<a id="canonical-c255efbbfdea3e2a0aff8439d8e4ae5d99e14c4fc6c4f4f763cbdac6f1809948"></a>

## Direct properties — s3_receiver.filename_options / 78e70b7436c6 / 3

<a id="canonical-ece9e20f1a3047ae68813023df4cbd133f3f9ab1e617b07aec01226213a1aaba"></a>

<a id="canonical-a2a8e2920203494b4e0bde3ebee95f107d5647388bf79cb41796829ef3f4ecdb"></a>

## custom_folder property — s3_receiver.filename_options / 78e70b7436c6 / 4

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

- [log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-1b8d30ac490f79b048797b991be18896ec3c0028bdb5afa46ffbf2858b72b23d): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-74b59d1566656e147a06b04d94ff40cc27985fe7b528161a446ec3305e5ef829): complete subsection reference.

<a id="canonical-2da1156715f0b3d338890962aee8a4f7c1ddb6e7c08ecef8715afdb4641b77bf"></a>

## Next pages — s3_receiver.filename_options / 78e70b7436c6 / 5

- [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-1b8d30ac490f79b048797b991be18896ec3c0028bdb5afa46ffbf2858b72b23d)
- [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-74b59d1566656e147a06b04d94ff40cc27985fe7b528161a446ec3305e5ef829)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-1b8d30ac490f79b048797b991be18896ec3c0028bdb5afa46ffbf2858b72b23d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5389846caa6d1ec8b2105333766957fe1a6aa78c8c7f9998941fc946921dd1d3"></a>

## s3_receiver.filename_options.log_type_folder — s3_receiver.filename_options.log_type_folder / 12c5a000422c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc)
- s3_receiver.filename_options.log_type_folder

<a id="canonical-babf622d80618bffd70b1e54007b2aeede94fd3aee7bcfb9a2dad24afa5819e9"></a>

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

<a id="canonical-d0ed320e04a1d2f70cdde9f8d827a105d1442bc3073e47f367531afb848035df"></a>

## Direct properties — s3_receiver.filename_options.log_type_folder / 12c5a000422c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2a719148a9b7f3ecd17d5d3486b38b7695e9afce2e4e616aedd353488c01e86"></a>

## Next pages — s3_receiver.filename_options.log_type_folder / 12c5a000422c / 4

- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-74b59d1566656e147a06b04d94ff40cc27985fe7b528161a446ec3305e5ef829"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a187f0bb1115676f7c25dc085d85f018b6ab1bbdcfed123b814a6c6efcfbc73"></a>

## s3_receiver.filename_options.no_folder — s3_receiver.filename_options.no_folder / 53cfe66177d5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-dae6dd6655685b107ef32055106d09b7b4dab7071a9959807699fd7c97a324e2)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc)
- s3_receiver.filename_options.no_folder

<a id="canonical-e5698fb7d9e33b0c2a77a9e282e83891ea358ed14c54cd8094393ab25ea3f344"></a>

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

<a id="canonical-e812288ac500a7e3988099038df8ff347b70c24bbebce1cf949ce194974c86b3"></a>

## Direct properties — s3_receiver.filename_options.no_folder / 53cfe66177d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa9395f668f3e69ca39fa0d95201212901b6523c29d18b4f234dcc2c34e6c5b5"></a>

## Next pages — s3_receiver.filename_options.no_folder / 53cfe66177d5 / 4

- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-107a168ef882d20a68233ab8d361bcea6dc73877663dc2b45a4db17c63358dfc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4234f3746f0500e78349b4be62bcfa5b49542d6219128d146579072b81afbd0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a94260c2cc2e606fa770a410c3c4c53d57801df53cdba92bbd23464483ecc01"></a>

## security_events — security_events / 5c40e5060488 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- security_events

<a id="canonical-998417cd4d62d05d5ed56692f9c2eee42b1e26b6d72d2278edfff8d76688f244"></a>

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

<a id="canonical-286101011da7f9f6caa0c62c3fcfee77cd0012635d9386c87a39425634c9305d"></a>

## Direct properties — security_events / 5c40e5060488 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7c5830b8b5b486370cbb5e24012d13b2472f33667977e6ff1b3e8738c5ef690"></a>

## Next pages — security_events / 5c40e5060488 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-653955d4eb6d184c66ed9d95df293a27c0054714aacdb05bc58b4233edb63976"></a>

## splunk_receiver — splunk_receiver / 30d857904475 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- splunk_receiver

<a id="canonical-9cea3aa814ee71326f9bd3e207fbff67b5828a4c37f9c9f8eefe23a77dc8bd31"></a>

Type: `"single"`. Computed.

Configuration for Splunk HEC Logs endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-7d453cef241833c43a4d21d11c482c7c35fb637c8fbb4c2484436fbec73dd7ef"></a>

## Direct properties — splunk_receiver / 30d857904475 / 3

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6): complete subsection reference.

<a id="canonical-1ad26994d6accd5bae3acbf0cc389f900e16db179300327e110ff93607d7cb29"></a>

<a id="canonical-f7b38d1d313db361599af831c561eac70c085321c2447320cb5782532804b4dd"></a>

## endpoint property — splunk_receiver / 30d857904475 / 4

Type: `"string"`. Computed.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Upstream description:

Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

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
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
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

- [no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-6f33c1576d6756db3a004947ea217ea5cab4df92dff33975881ed2e9c5f6b29e): complete subsection reference.

- [splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b): complete subsection reference.

<a id="canonical-d3305ba967590b4837b3dce8fb54100785720ce73ca8d90987c7050c00a864ee"></a>

## Next pages — splunk_receiver / 30d857904475 / 5

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- [splunk_receiver.no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-6f33c1576d6756db3a004947ea217ea5cab4df92dff33975881ed2e9c5f6b29e)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ad54e70ed3027781b8e2e50b34d221ab486cce0faf1fdf7edfbbb27fcbecba2"></a>

## splunk_receiver.batch — splunk_receiver.batch / a4884c29322f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- splunk_receiver.batch

<a id="canonical-34a2df733d0986cd3156d9a64bdd36da00f1e324c8764d9b601ce6d101e5b6f0"></a>

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

<a id="canonical-454f8c32680cf9833c3311e87cf848a3ca6204ef7c6eb579972f459744fe19fe"></a>

## Direct properties — splunk_receiver.batch / a4884c29322f / 3

<a id="canonical-776fc6e70314e46fd02d554c5e7138cb0bd5978b253b4387abaa5aab132b183e"></a>

<a id="canonical-00cd4261123193e17bad940bb1687b5c5366473ebb682468a811f58b023a8391"></a>

## max_bytes property — splunk_receiver.batch / a4884c29322f / 4

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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-d5ab90391877bd974fe9cc5100cbc2005719cb497592a62d5870c0b46055de79): complete subsection reference.

<a id="canonical-3e068b35448217058c7b3208e5e168674babc83794d9013fbae8bd3b9ec61d65"></a>

<a id="canonical-f2994199600098374057dd3c11b468f289352e1168b278178943b460366da5a0"></a>

## max_events property — splunk_receiver.batch / a4884c29322f / 5

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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-4f00a2a1536b54d17900b9aa35d0f580b9bddb0fcc06e6aa4959dc53f4b68d43): complete subsection reference.

<a id="canonical-24c9ff7b82d3935fb4c5ebd7ab232c12809da4f450aa0f11e4c1da50db2488a8"></a>

<a id="canonical-b183748972fb27a6ec512da0c49887dc9efc97509defc22af460334227bac022"></a>

## timeout_seconds property — splunk_receiver.batch / a4884c29322f / 6

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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-7b78b683bbfc5e0f41dc03c8ae7c74f184870ea3065156f1eadb1762e7e6fd0e): complete subsection reference.

<a id="canonical-9fc04562a6b81b46bc6b1983f6eb88ebfeab86cb70851dfee6153d2353ba87d3"></a>

## Next pages — splunk_receiver.batch / a4884c29322f / 7

- [splunk_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-d5ab90391877bd974fe9cc5100cbc2005719cb497592a62d5870c0b46055de79)
- [splunk_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-4f00a2a1536b54d17900b9aa35d0f580b9bddb0fcc06e6aa4959dc53f4b68d43)
- [splunk_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-7b78b683bbfc5e0f41dc03c8ae7c74f184870ea3065156f1eadb1762e7e6fd0e)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d5ab90391877bd974fe9cc5100cbc2005719cb497592a62d5870c0b46055de79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96554db28b7f15b829dbb943c806c10f01521c340bacb167e78eb38a8175daf5"></a>

## splunk_receiver.batch.max_bytes_disabled — splunk_receiver.batch.max_bytes_disabled / ad2e711d5c75 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- splunk_receiver.batch.max_bytes_disabled

<a id="canonical-6e2f37d1c855e45b12debbf7b019c6f49a4b2d9eb651431c0b680d7a93bb7625"></a>

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

<a id="canonical-243f27e76f46269e74ea85b51b52c7669c4668b1c540ca4c6beb446483a1c223"></a>

## Direct properties — splunk_receiver.batch.max_bytes_disabled / ad2e711d5c75 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d96bdddc6494cdecd29d5353f97cd08d4beb8c4784d1171a17f11190c3875913"></a>

## Next pages — splunk_receiver.batch.max_bytes_disabled / ad2e711d5c75 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4f00a2a1536b54d17900b9aa35d0f580b9bddb0fcc06e6aa4959dc53f4b68d43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72b335af5ae8e9a281c1fe50d3e7f67eb22029300e90eb325283d324f07f7b2b"></a>

## splunk_receiver.batch.max_events_disabled — splunk_receiver.batch.max_events_disabled / 1c90ce935c27 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- splunk_receiver.batch.max_events_disabled

<a id="canonical-758275a0ed18e57c894658ed4fff05162cdbb38a6b7a7a21df9a325320f3767b"></a>

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

<a id="canonical-1a208f228081d99e4b22cf6eddc354f8cf25e1969b9f6ab5305e889e16494763"></a>

## Direct properties — splunk_receiver.batch.max_events_disabled / 1c90ce935c27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1df80d7d8b809ff7a59d93ea9749d3723e367166656d2537afeaccfb475ea6d"></a>

## Next pages — splunk_receiver.batch.max_events_disabled / 1c90ce935c27 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-7b78b683bbfc5e0f41dc03c8ae7c74f184870ea3065156f1eadb1762e7e6fd0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fc449c192b54379d76e5014da4b603e39a936a894f058390a5886bc912d51ee"></a>

## splunk_receiver.batch.timeout_seconds_default — splunk_receiver.batch.timeout_seconds_default / 45b15ee63ad5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- splunk_receiver.batch.timeout_seconds_default

<a id="canonical-12d29d50f9c214f458879e12a035b24bb7d7dd9c8781826f08d19151e7eb7cd3"></a>

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

<a id="canonical-dfaa523f1fd5062f1f47f790c320ad082ed42bc0d2586eb0217504e5599d5c97"></a>

## Direct properties — splunk_receiver.batch.timeout_seconds_default / 45b15ee63ad5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c56d44e74a0edae0d0829f1e24f1556728a43b546b19946b8900188ea4de0367"></a>

## Next pages — splunk_receiver.batch.timeout_seconds_default / 45b15ee63ad5 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-5f0e20e0bc3f5f86705a951a194db9822e8119a2024c0ade70995464b81f8bfe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd059880dd000ee58d4040ad773fe195f9c45e7766bcee1223d5e042f8953c3c"></a>

## splunk_receiver.compression — splunk_receiver.compression / cda73306c88c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- splunk_receiver.compression

<a id="canonical-5eee87c1216fe43a4fd3c796e79c7e9a30199b5bc8e47e6efb9ff8f24c3a97ce"></a>

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

<a id="canonical-901fa19a43ebf90e74c2aaa11defb822658886b0975af412333059e77b0badf3"></a>

## Direct properties — splunk_receiver.compression / cda73306c88c / 3

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-283bcfc058112c5ac8a6875dc6f3dd9a273ea87dd28e1fb2d14b00e61cf9144c): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-a6726a2df91056f9d2052f5d009599ed9d8dcebd8a0df372467a6fdaf58a7794): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-5e9c5ef12386f8b613853fc14eee6febbd7d3e8d9ee5ebc62baaae711165c281): complete subsection reference.

<a id="canonical-e8921b1558fa40c6e6a7650e566f16f6ba1d301c28bf9403bae613eba2053d3d"></a>

## Next pages — splunk_receiver.compression / cda73306c88c / 4

- [splunk_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-283bcfc058112c5ac8a6875dc6f3dd9a273ea87dd28e1fb2d14b00e61cf9144c)
- [splunk_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-a6726a2df91056f9d2052f5d009599ed9d8dcebd8a0df372467a6fdaf58a7794)
- [splunk_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-5e9c5ef12386f8b613853fc14eee6febbd7d3e8d9ee5ebc62baaae711165c281)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-283bcfc058112c5ac8a6875dc6f3dd9a273ea87dd28e1fb2d14b00e61cf9144c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94ae02710667cb6c76f61df16535facc90ab3b4986d10b29479a3a481669c8d"></a>

## splunk_receiver.compression.compression_default — splunk_receiver.compression.compression_default / 81fb2d9ac536 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- splunk_receiver.compression.compression_default

<a id="canonical-f5f4d6b226d8a2af36388ba49e8f494d13d4127a42ab800be54f66e729d83ef2"></a>

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

<a id="canonical-989f49e004a6f22fd9986a6a9364a5937c73c3b66607aaf194ac581edd99b1ee"></a>

## Direct properties — splunk_receiver.compression.compression_default / 81fb2d9ac536 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-320eb654b0b643b05a0e05d4e7d60e138e52888e752e3f12804e0a000ac7d19d"></a>

## Next pages — splunk_receiver.compression.compression_default / 81fb2d9ac536 / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a6726a2df91056f9d2052f5d009599ed9d8dcebd8a0df372467a6fdaf58a7794"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e5e3b2f6028e82f064efce35f5af4081fe7092fcd0fb6b04c3863174e092b4"></a>

## splunk_receiver.compression.compression_gzip — splunk_receiver.compression.compression_gzip / 8cfc3dd7d97f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- splunk_receiver.compression.compression_gzip

<a id="canonical-9e73327291839cdc29e5594aaf0c836bfbd7853549a14a067338d6a3e54296d0"></a>

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

<a id="canonical-971cdc6dfd0d3236036b9a8a3c978f4033ec73ca713225a5fb8fcfd965a9b5b0"></a>

## Direct properties — splunk_receiver.compression.compression_gzip / 8cfc3dd7d97f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1a5fd54b69409443d7152f39e36c7c48f1840ba0fbc5c8d45cdd61c63ad3630"></a>

## Next pages — splunk_receiver.compression.compression_gzip / 8cfc3dd7d97f / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5e9c5ef12386f8b613853fc14eee6febbd7d3e8d9ee5ebc62baaae711165c281"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b2bc1715850a88f18ddd096e4df0e7a15947fef3733dfddd7a5ba3d24510c47"></a>

## splunk_receiver.compression.compression_none — splunk_receiver.compression.compression_none / a35778681a44 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- splunk_receiver.compression.compression_none

<a id="canonical-a94013c4d2eb50e7948ab142d6e91f813185cdd57ac347829813d58086cde01e"></a>

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

<a id="canonical-47264479b7fc503e22f291a5741fc01ba0230b3fe1982415ce5490a3501ee431"></a>

## Direct properties — splunk_receiver.compression.compression_none / a35778681a44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3666b2792455478ab4a01b6d35c117a2732eefb346b8e5ffcffeabb44b38b633"></a>

## Next pages — splunk_receiver.compression.compression_none / a35778681a44 / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-f362bb3a5632784678acac4ddc91123e14290dbe6c0532e649a2c047410ef7c6)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6f33c1576d6756db3a004947ea217ea5cab4df92dff33975881ed2e9c5f6b29e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ad26f23a03cb3fb6354bd42c605ccbe1302fe8e70615137de7fde8cf638cb3b"></a>

## splunk_receiver.no_tls — splunk_receiver.no_tls / 844e99bb4fb7 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- splunk_receiver.no_tls

<a id="canonical-7be2ff8bbf7dd1e630b7707ec28bc7b02d81d3aaebea54a485034a3a7602ea32"></a>

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

<a id="canonical-5878374e3bcbff69200345809d01a5e06dc177fe16c91e06c27d56ef55e99880"></a>

## Direct properties — splunk_receiver.no_tls / 844e99bb4fb7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9955f347b5c10ee15e414a4d67d9f7047835db8b2e589fa3acf5a9d44ccfcbe5"></a>

## Next pages — splunk_receiver.no_tls / 844e99bb4fb7 / 4

- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7022a0e3222683a3682d2ee9d9b503c55391df98081bf69b5e2856efb5fa0d09"></a>

## splunk_receiver.splunk_hec_token — splunk_receiver.splunk_hec_token / dbbae971add2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- splunk_receiver.splunk_hec_token

<a id="canonical-044e2dcbc57b8f62129b6426c180c925b5e20308fa3ec8dfcb4e062683f44e49"></a>

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

<a id="canonical-73366df57886770b02059284d2ae4483db7b8f59ea8cf5aa1d4377339ac93026"></a>

## Direct properties — splunk_receiver.splunk_hec_token / dbbae971add2 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-59356c85fef752b6d08f94cdf4053ae6dd20220a766cb38069f14017ea02113e): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-4994eae37efccd862472df6eca0c9fc5007ff76bbf9dc0ae356a67093215b9aa): complete subsection reference.

<a id="canonical-ff467a603de68576a786d5e48205b9a4cc251fefe8e4d1f414ea1548050293eb"></a>

## Next pages — splunk_receiver.splunk_hec_token / dbbae971add2 / 4

- [splunk_receiver.splunk_hec_token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-59356c85fef752b6d08f94cdf4053ae6dd20220a766cb38069f14017ea02113e)
- [splunk_receiver.splunk_hec_token.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-4994eae37efccd862472df6eca0c9fc5007ff76bbf9dc0ae356a67093215b9aa)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-59356c85fef752b6d08f94cdf4053ae6dd20220a766cb38069f14017ea02113e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8605c4b2a65b6a027ac7a4e451bc4d92f9e83322aec3c80c06c331ef393af6d8"></a>

## splunk_receiver.splunk_hec_token.blindfold_secret_info — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca)
- splunk_receiver.splunk_hec_token.blindfold_secret_info

<a id="canonical-de9221762147d6176ab8c8246c67fc3b6c452b8f08cb85a7b303f2d2cece57d8"></a>

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

<a id="canonical-699927e3f5fb4bd9c62c00f9ce7fa9d863ac5d7e765c582e8bc81946801d7c5c"></a>

## Direct properties — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 3

<a id="canonical-b345e3a32559b8bbe550c8231dae519c54fa46c5bed31f9324dc6e275578eed3"></a>

<a id="canonical-bfafe302be24486c7b2624363051ca13c053fe74f3b8c945188078925ad49a01"></a>

## decryption_provider property — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 4

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

<a id="canonical-13359d5dfbdca6ebe3555c27b0b72c222e633f8e513e3b71538ae494c72005e3"></a>

<a id="canonical-a0fe94e488781a5651551ca1b6d2efe095eaa838ce0fbd5f3552b9c7125f6a8d"></a>

## location property — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 5

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

<a id="canonical-75748dce5fdc52d1cb36abc510c089d1757ae5f3994269c65a41188fdf6e81fc"></a>

<a id="canonical-4219dfcbe6e4926c4a12d8fb0224df07bdce347736bb6eed647c82ab662bd60c"></a>

## store_provider property — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 6

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

<a id="canonical-da88422d92b8157178ab418d559c3b48eb4152868932b0892113e1d328c9fa4d"></a>

## Next pages — splunk_receiver.splunk_hec_token.blindfold_secret_info / cea377b071ae / 7

- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4994eae37efccd862472df6eca0c9fc5007ff76bbf9dc0ae356a67093215b9aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a76bfcad88cefb5c273d906a2a87eb7e0902613c81e8d8bf907d6793092282c2"></a>

## splunk_receiver.splunk_hec_token.clear_secret_info — splunk_receiver.splunk_hec_token.clear_secret_info / ba2c7a1ca205 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca)
- splunk_receiver.splunk_hec_token.clear_secret_info

<a id="canonical-202ef4d0c52ec01c10c6adfdd944ec782dc82c7ad10a9b144e1fad139ae2e257"></a>

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

<a id="canonical-c36588a28acac3720fa0ece3a2344efa0ea854e56188e1fb902ccc243d8ff1d3"></a>

## Direct properties — splunk_receiver.splunk_hec_token.clear_secret_info / ba2c7a1ca205 / 3

<a id="canonical-7fddef634373fa41eabcacc6eb6cde15b940369b5acd989f0fba2763916799e3"></a>

<a id="canonical-b0b0455d2eaa0554a6c0148e4c15cbafd5ab181b8d1e130895096b3a842a6d89"></a>

## provider_ref property — splunk_receiver.splunk_hec_token.clear_secret_info / ba2c7a1ca205 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2d736ecc5cdac838b855c2629d6e4c5378568366215045719a8c6ffff9db327f"></a>

<a id="canonical-4a481e2a3b3e2c0c6d551f31d7033fd16e3a9d0cb6a79412126641d30d9fd94a"></a>

## url property — splunk_receiver.splunk_hec_token.clear_secret_info / ba2c7a1ca205 / 5

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

<a id="canonical-32237d7bca99974077876b6287ebb3d96bc7f049a709d2bf4e6748b8d6569195"></a>

## Next pages — splunk_receiver.splunk_hec_token.clear_secret_info / ba2c7a1ca205 / 6

- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-436a02c00fe39470eaa51ad422676369925b7983809c1caaf63a11e028df61ca)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87be1eeaef0a0a32d109425309f1d05206839f00ad090924bb14f2e00d5cf775"></a>

## splunk_receiver.use_tls — splunk_receiver.use_tls / 0af8aa48018e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- splunk_receiver.use_tls

<a id="canonical-4d99715e85fec97e9bd6aff8fd89b9c95792a8daed8e69dc71e3587b77dbe4d2"></a>

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

<a id="canonical-a097cbb73f0102280bb6f80226ccc7ab00fe116009a5e710e7a7e911727a989a"></a>

## Direct properties — splunk_receiver.use_tls / 0af8aa48018e / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-177841fe5a8c5444143fb014115895a281766272420517e77eac6718f12e1f92): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-981d639ca1a818f91c8e062476c2519af7f71aebef43d9f0f8e36d1a4e444067): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-6ef2371675f3b700700c968f6ae035a33f354f66208569b4e2dcedbd5621bc59): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-ffabdeff8c1ccfbb3de64c18a66e143716deda5dcb88d4d07dbdf3e1f04995b9): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-5df0dc66ea0812046b75a25ce877e2d398461b4b972d18a0573afcc16f4120e0): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-eae57017adbeb7c68e1e092821240ed27d2ff35f1a975f08201392e6d50edeef): complete subsection reference.

<a id="canonical-4f815177ec88baea9133b6f243ef318d9b42ac2b1543fbc2add7e0b428b585a2"></a>

<a id="canonical-2e3f8bc8589272cefb9d852c3239bd4c3e49762f643aeb7d14f0ca3bd82d0cbd"></a>

## trusted_ca_url property — splunk_receiver.use_tls / 0af8aa48018e / 4

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

<a id="canonical-6b07c6395aa664129afda4d2b8db90adf85f7315c61c4acab29222076b4b1dd4"></a>

## Next pages — splunk_receiver.use_tls / 0af8aa48018e / 5

- [splunk_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-177841fe5a8c5444143fb014115895a281766272420517e77eac6718f12e1f92)
- [splunk_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-981d639ca1a818f91c8e062476c2519af7f71aebef43d9f0f8e36d1a4e444067)
- [splunk_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-6ef2371675f3b700700c968f6ae035a33f354f66208569b4e2dcedbd5621bc59)
- [splunk_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-ffabdeff8c1ccfbb3de64c18a66e143716deda5dcb88d4d07dbdf3e1f04995b9)
- [splunk_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-5df0dc66ea0812046b75a25ce877e2d398461b4b972d18a0573afcc16f4120e0)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638)
- [splunk_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-eae57017adbeb7c68e1e092821240ed27d2ff35f1a975f08201392e6d50edeef)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-177841fe5a8c5444143fb014115895a281766272420517e77eac6718f12e1f92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3873a863e9b7d1388bcf7e042c04886a2f1c2c287b93b2ad416cd6f78c17f9da"></a>

## splunk_receiver.use_tls.disable_verify_certificate — splunk_receiver.use_tls.disable_verify_certificate / a3bc5e3af640 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.disable_verify_certificate

<a id="canonical-abb7e0b70e2abd366dadfb109b660e27cfbaa253ba3b8ead5b41f74dc3e4a54d"></a>

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

<a id="canonical-fe00818ff54895eeb4931f687717e03179b6914c5d07fc7240fdd828323df06e"></a>

## Direct properties — splunk_receiver.use_tls.disable_verify_certificate / a3bc5e3af640 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b01051d0ca2d00677fae6d86f86bf1d4134c7be0984cc2bd211476104bf902d"></a>

## Next pages — splunk_receiver.use_tls.disable_verify_certificate / a3bc5e3af640 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-981d639ca1a818f91c8e062476c2519af7f71aebef43d9f0f8e36d1a4e444067"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8eaad2f84ad0a6b33c657daf0a132fb7855b50b159a915314c87fd65222ede9e"></a>

## splunk_receiver.use_tls.disable_verify_hostname — splunk_receiver.use_tls.disable_verify_hostname / adf13f506747 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.disable_verify_hostname

<a id="canonical-a82b9462cb3d11d7d7fa805695d852dab1f201c3d7bb5678dc7b82ff82438ece"></a>

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

<a id="canonical-ef5bf3f8aee477777d45c6388c9718e42913e8ae302d2f249c6b37f7750733f6"></a>

## Direct properties — splunk_receiver.use_tls.disable_verify_hostname / adf13f506747 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06a716414a3286b8aa796fe1a4d210e81fb9ede1e10052d2f4f9251b141cd39d"></a>

## Next pages — splunk_receiver.use_tls.disable_verify_hostname / adf13f506747 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6ef2371675f3b700700c968f6ae035a33f354f66208569b4e2dcedbd5621bc59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62b158fae6cf4719637faf0b45ca71acaa6aa5cf326eb22cda0698ae3492a98d"></a>

## splunk_receiver.use_tls.enable_verify_certificate — splunk_receiver.use_tls.enable_verify_certificate / 43d1a2bbe718 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.enable_verify_certificate

<a id="canonical-2ceaca628858702cdb41a68d5f289dda8d63dde178acbc7f490eac8486de74f9"></a>

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

<a id="canonical-0b029bb5b42dea141bf0733a3217853aa99c39b4b9917b1ba67074d462e512dc"></a>

## Direct properties — splunk_receiver.use_tls.enable_verify_certificate / 43d1a2bbe718 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2de3dbdc419a622829ad12946a6a14dfaedb083d09b9a383c94fb9c9dfb1863"></a>

## Next pages — splunk_receiver.use_tls.enable_verify_certificate / 43d1a2bbe718 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ffabdeff8c1ccfbb3de64c18a66e143716deda5dcb88d4d07dbdf3e1f04995b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3b0ec5c94d890d6ae08ef474e1244d3d13d08da64043061147e55763af30080"></a>

## splunk_receiver.use_tls.enable_verify_hostname — splunk_receiver.use_tls.enable_verify_hostname / b3b3761b49d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.enable_verify_hostname

<a id="canonical-b46c1dac09c40b284fb42e09aab194f30f2b606e9df403c8f09293dabb01728a"></a>

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

<a id="canonical-590127d40d3a9f3aa19f9d32f76c796477155b28621ceaa8c2772abca57a97e6"></a>

## Direct properties — splunk_receiver.use_tls.enable_verify_hostname / b3b3761b49d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c939cdaf462b0dd578a5ca6c797e81e60dca687aaa95417ad68de7359adcadb2"></a>

## Next pages — splunk_receiver.use_tls.enable_verify_hostname / b3b3761b49d2 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5df0dc66ea0812046b75a25ce877e2d398461b4b972d18a0573afcc16f4120e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1595ff3ad62b75fe216563343a9b35014287fdf04b692d670880b3af31e5b6b3"></a>

## splunk_receiver.use_tls.mtls_disabled — splunk_receiver.use_tls.mtls_disabled / 81d6424ad721 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.mtls_disabled

<a id="canonical-c60ee2a1d21d974a2e545ec9809e6618df78ffff3b433f5754e053fc53038507"></a>

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

<a id="canonical-815a4e1bc2ef9d79a02b573ad76f710c52aa173b2032911c703d4fe5a9ce4085"></a>

## Direct properties — splunk_receiver.use_tls.mtls_disabled / 81d6424ad721 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c253ac34413b547ac3dae7a1fcaa0d7f8968e064a24aebd99ecb99e99392b5cc"></a>

## Next pages — splunk_receiver.use_tls.mtls_disabled / 81d6424ad721 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fb938e8fa7e9e3372b157e108b6832b9cb9c660d392bc48aaeb1067c3e78c85"></a>

## splunk_receiver.use_tls.mtls_enable — splunk_receiver.use_tls.mtls_enable / 4ae6084a5d5b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.mtls_enable

<a id="canonical-84c913066466246d55d5f8a88979343ec606a4ea4bcef1c99154ebf41170aaef"></a>

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

<a id="canonical-cffca7d3297123eb04470d7a1b0bd5c1327d5102b05334d555553d687e6579a3"></a>

## Direct properties — splunk_receiver.use_tls.mtls_enable / 4ae6084a5d5b / 3

<a id="canonical-1811a5840278a7188ec0d5f1d0a5c9e3989f94e63bcafd57fc2ab3d38ebeeba8"></a>

<a id="canonical-82e1e846753e456509484237abd9fc8153dbd89a4c3254cfbd04289f3b8110ea"></a>

## certificate property — splunk_receiver.use_tls.mtls_enable / 4ae6084a5d5b / 4

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

- [key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0): complete subsection reference.

<a id="canonical-e91f6e53d06ab9e5ebc83c0cd9bf628d1d3748ccca326c5f761649bd70a0f10a"></a>

## Next pages — splunk_receiver.use_tls.mtls_enable / 4ae6084a5d5b / 5

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98fd720b362f2485b04e88aa036aafb51a475a9cae3b539eadefa94f942b9f5b"></a>

## splunk_receiver.use_tls.mtls_enable.key_url — splunk_receiver.use_tls.mtls_enable.key_url / 4f370abbd733 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638)
- splunk_receiver.use_tls.mtls_enable.key_url

<a id="canonical-d786d848070a35d93f6e2b2dd75b9b8e5e5bb54365896daed911e13fc64fbc5d"></a>

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

<a id="canonical-6287d2dddecdf9a07bb235ebdded55047814c6e27d8d0551dfff92560627f800"></a>

## Direct properties — splunk_receiver.use_tls.mtls_enable.key_url / 4f370abbd733 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-e27b9a5e29bf793398e65287c236f35a973cdb609f849da912c61531a070f7ab): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-4d3ff242879c20f944d5994bde66382a5f183021537e2ac5e5093c59694d94c9): complete subsection reference.

<a id="canonical-e5401b4b240b20819aa4fd27445c7c6e603ced004544baa8c5a7d035d97acd98"></a>

## Next pages — splunk_receiver.use_tls.mtls_enable.key_url / 4f370abbd733 / 4

- [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-e27b9a5e29bf793398e65287c236f35a973cdb609f849da912c61531a070f7ab)
- [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-4d3ff242879c20f944d5994bde66382a5f183021537e2ac5e5093c59694d94c9)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e27b9a5e29bf793398e65287c236f35a973cdb609f849da912c61531a070f7ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8be1fe7dfe9729db104530646af8b912d1d61fa6f2ccc069c259721a96963931"></a>

## splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0)
- splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-98e5aefc296f96fa068108af27931ca91875859c1f87ec1583e5f1bbef71cc1a"></a>

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

<a id="canonical-524e45f9b0c847d774c3a3c694012bb0ab6f77089a1b2ef75e9900d359ca3cef"></a>

## Direct properties — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 3

<a id="canonical-d2e97925f3c484802181d46258617a374eac6db0dc3fb34eb9a21ba6148353b1"></a>

<a id="canonical-01e6e93843ed9ce8bb4ee2437166f66b467783f76c027b4c76f4bc9d54724ff5"></a>

## decryption_provider property — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 4

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

<a id="canonical-3a51c28d4c72ce2e91c0c01f7bd39083af12defcc3e4cada6ccd615e0333c71c"></a>

<a id="canonical-d7794609ee1329602d2352e5d353ca1d392368e09178f85d13fe631df2ad4be0"></a>

## location property — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 5

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

<a id="canonical-7975a309044e820703085f20b8312c37c2e6ab157f071b7b7b722af305ee048a"></a>

<a id="canonical-a61d548667be68d2e6bd16e9a2d6695ada2bbb60ab0d4bc88314bb6b647fbf31"></a>

## store_provider property — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 6

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

<a id="canonical-62aada13950206d4b592f6a8cbd7347c950fd38332a45aaddbade2c92b33dae4"></a>

## Next pages — splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 68a0981e7150 / 7

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4d3ff242879c20f944d5994bde66382a5f183021537e2ac5e5093c59694d94c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ceb9006ae3774468d021e79979f8fbcb84d63980eb8f8e82fa28ce803b4f9b5e"></a>

## splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info — splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info / ab9d584fbfca / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-195b7ea2b496aa07bb13df03a6509b116fce65b1ce2ebac178f569fe01f92638)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0)
- splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-ee4f95880b5768442a7371a5a7c37686103dbe672fa10742f996359703e2eaaf"></a>

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

<a id="canonical-d180fba417f7264f124e5eeec5dcf7d20dba2297dc96893e28624ae9c6a36768"></a>

## Direct properties — splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info / ab9d584fbfca / 3

<a id="canonical-317d1fb57c20c83cb0f520804963901841283c437cc3b47299df8539d9a1d9ea"></a>

<a id="canonical-b86e16c0802fd641700f3146e77569e8abd5065a82d2c53dd0151b05ecbbdef3"></a>

## provider_ref property — splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info / ab9d584fbfca / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c62ef53090ef57d9b81fae6539bf2c0bfb8eeced53a4120ba067e24aaff7f93f"></a>

<a id="canonical-4f0cfbd2d23ec996c8b0ede028401af986738e203c51973f2f72831ec7b4a75a"></a>

## url property — splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info / ab9d584fbfca / 5

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

<a id="canonical-63752aaa8143efb0fb429b3a00706c2c97d94ab4be518fea73d89cf03fc087a5"></a>

## Next pages — splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info / ab9d584fbfca / 6

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-f909f109506303464fc99764682fb6862954dd55c387ca3523c685bbb8dee0e0)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-eae57017adbeb7c68e1e092821240ed27d2ff35f1a975f08201392e6d50edeef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b249d9dd251438f39b5c4dd21221a6b0d8e821248b025897bc8a1d7644528e18"></a>

## splunk_receiver.use_tls.no_ca — splunk_receiver.use_tls.no_ca / eb36849320de / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-6ed29c2aa1b2f50992c155d7c0c80fc7a5bf1c207674965cc79d023a9e7e0070)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- splunk_receiver.use_tls.no_ca

<a id="canonical-1960c66bf5a637fa3d0703679b811caebce661c566aab9c43f75747cb56a97db"></a>

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

<a id="canonical-cedeaa02d6064cc9684aaa438b289d0c4eeb0919c08f48f91073f70c4e00dc1d"></a>

## Direct properties — splunk_receiver.use_tls.no_ca / eb36849320de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f3a3f6c13c8205f7d830889b309437943dd79697dafbfa30310e985b9c52c83"></a>

## Next pages — splunk_receiver.use_tls.no_ca / eb36849320de / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-cc96ee4402fb6c658e8ebe116c70b3937604232dc99be0d9fc753069bc5ef44b)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b084449e7353ae7304dc406c9c1a045f862550094cba29245939bc8560d0240"></a>

## sumo_logic_receiver — sumo_logic_receiver / 231533809929 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- sumo_logic_receiver

<a id="canonical-95863243449b46ca8ba7bf4ea6f398b35ba9875c97fd3d58d4c6811920c9a0a1"></a>

Type: `"single"`. Computed.

Configuration parameter for sumo logic receiver.

Upstream description:

Configuration for SumoLogic endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-314de13ff02ff542ca155ce0e2d707efb37231e489f77b0683c6299f823275c2"></a>

## Direct properties — sumo_logic_receiver / 231533809929 / 3

- [url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129): complete subsection reference.

<a id="canonical-bd996f97f5537e97ae28206546bd5b2d21f92fe89849d296aaf0da8bb530234f"></a>

## Next pages — sumo_logic_receiver / 231533809929 / 4

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cde1f9959c99c35137b094c64b1a2498ef8209ff46ca70d3da0ba251d58e0681"></a>

## sumo_logic_receiver.url — sumo_logic_receiver.url / 74868ebfff1e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1)
- sumo_logic_receiver.url

<a id="canonical-cd5ed7f98721b5cfe5d7a06d3b8a10232ccec0455d9e782020bc57a294cfb487"></a>

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

<a id="canonical-8f41ee92f6d3f3daaba91e5a5dd53b1547572f3e6bb230c863de7dcfb254b471"></a>

## Direct properties — sumo_logic_receiver.url / 74868ebfff1e / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-2951b0df9223dae1cb50194b22afbf553827bed6bd2bb0a0c579896018380842): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-ca58f5dcd4679d6e5ad46edf1417b4654942e52e336c2ae7c61613ac1361a36a): complete subsection reference.

<a id="canonical-d7bb13ef880a9a171a7569b3f01a2288bc0ba07bbc921819b67b41f8143cd8b7"></a>

## Next pages — sumo_logic_receiver.url / 74868ebfff1e / 4

- [sumo_logic_receiver.url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-2951b0df9223dae1cb50194b22afbf553827bed6bd2bb0a0c579896018380842)
- [sumo_logic_receiver.url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-ca58f5dcd4679d6e5ad46edf1417b4654942e52e336c2ae7c61613ac1361a36a)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-2951b0df9223dae1cb50194b22afbf553827bed6bd2bb0a0c579896018380842"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25424fc5fc1f209122d87151fa1b6cdee0405478d2cd2130cc09943ea2c65e1c"></a>

## sumo_logic_receiver.url.blindfold_secret_info — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129)
- sumo_logic_receiver.url.blindfold_secret_info

<a id="canonical-bb24244f5fd3a931d8c6bbfde05acc9f8d03181b32c217d765bf00f2b6c10dc1"></a>

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

<a id="canonical-bc6b9a09feb18253707baa4e4596383b59966120d08618703568a15e869b35ca"></a>

## Direct properties — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 3

<a id="canonical-59bbe6f718ab3cd7b5aaacbc9a70813f725372cc9a251ea7d311a49714911d0e"></a>

<a id="canonical-1b39f834ed832ca6449ce63a14e2cf81547b3764b110a0d329ef6391e8f757b2"></a>

## decryption_provider property — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 4

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

<a id="canonical-b0f2dd09721aaf1b7475de5e952526890831b1c35f39b61012c7d44b5666b085"></a>

<a id="canonical-6e1a8961e6540254ad390adb60751e0e82b08f3f8dbe2b70e23df38e11a3b7aa"></a>

## location property — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 5

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

<a id="canonical-31af75e245dacacc8a8fe8344b49fdd6589e9acfa2f551eefaa4971ec7eccb47"></a>

<a id="canonical-062c436fe12d4165a456dc3dc1fe08344c246ce56a5c6b8799f31719dc323462"></a>

## store_provider property — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 6

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

<a id="canonical-722c15ba0aecb097986e52637f354ddd71f20ecf200f221b65e1664f79dd3678"></a>

## Next pages — sumo_logic_receiver.url.blindfold_secret_info / 65035515cc12 / 7

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ca58f5dcd4679d6e5ad46edf1417b4654942e52e336c2ae7c61613ac1361a36a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01feeea04ac054d46c9b1788a2bae288c671cd8837600e52169cff023ba1f099"></a>

## sumo_logic_receiver.url.clear_secret_info — sumo_logic_receiver.url.clear_secret_info / f9534b0a1862 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-60223cb26da3c4491af2cd9271ba7071122485608016a658e918ce37e12dc0c1)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129)
- sumo_logic_receiver.url.clear_secret_info

<a id="canonical-098ebdf4a57abe43e3c77acf8d241f3b1b88b22709eca7dc4b5c905b776489c2"></a>

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

<a id="canonical-bafad4f816d623fa5d541f53d7936ea07bb223b0774bb0432ba0e0f9ea17fdc2"></a>

## Direct properties — sumo_logic_receiver.url.clear_secret_info / f9534b0a1862 / 3

<a id="canonical-73c14d2eb6766b1fe5f820db4654175ce1ca0e1aed5cfe3238fe488df55aaa86"></a>

<a id="canonical-8f8000d2f3aeff29ee45ef96915252bc862d9e3743a3d41074a7eb55fdf4cc24"></a>

## provider_ref property — sumo_logic_receiver.url.clear_secret_info / f9534b0a1862 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6775b126b13e527e2ae655d822f2232d38804c4e5265f89e94b4e68f789d0a23"></a>

<a id="canonical-c34d89feacbf516cee8bfaa73982570eafe9a419028a65cb16a361df44d81958"></a>

## url property — sumo_logic_receiver.url.clear_secret_info / f9534b0a1862 / 5

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

<a id="canonical-4715857eba0e03fa88f9c041dabd3f04bc91068230e373bc7784bf768d9582a7"></a>

## Next pages — sumo_logic_receiver.url.clear_secret_info / f9534b0a1862 / 6

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-db8a76b49c8cc0f1794c21eb58f65de424cb0924e01da38dd2e7dc6701de1129)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
