---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-707860f591dcc04b115200d998cfdc3af19f9b96f926c693b78a04eaadb75deb"></a>

## Direct properties — tls_intercept.custom_certificate.private_key / 1a3e99cf0cba / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-3a4fee5c67c61a509017bdbb1c6652401c65e5fe854ebb3a00982dc7e62b09ae): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-7a9b9b1f9fcdee81fcea694f6eead012ed7a841f1ea4199ef91729dea077fba2): complete subsection reference.

<a id="canonical-e26fb30678a254735ff68fed30da636aa0c273b6d053ca0dc0a8b9efb2910ba0"></a>

## Next pages — tls_intercept.custom_certificate.private_key / 1a3e99cf0cba / 4

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-3a4fee5c67c61a509017bdbb1c6652401c65e5fe854ebb3a00982dc7e62b09ae)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-7a9b9b1f9fcdee81fcea694f6eead012ed7a841f1ea4199ef91729dea077fba2)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-3a4fee5c67c61a509017bdbb1c6652401c65e5fe854ebb3a00982dc7e62b09ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d63970c9199e5fcabd9fa8038b41cd3d68c0a8510260e37587acb4a5507f7cb8"></a>

## tls_intercept.custom_certificate.private_key.blindfold_secret_info — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-9c6019608b52faeb2a90e3946f800ba6b5c42ab618dd32cd402002aa5b31e6dd"></a>

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

<a id="canonical-0019891e30ec63919c23122b3bd5bee3dc62fb70d343b6281d03db487834ab0a"></a>

## Direct properties — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 3

<a id="canonical-df23aea7bffa0d5e1cd368399a94be42625cd91623358525f1ae5749748de698"></a>

<a id="canonical-de43994cb8da54a6aac2c8724f94d88b826f00e927bc010f91cf3457e77a4e17"></a>

## decryption_provider property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 4

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

<a id="canonical-f0323c83c9282e20a6f18e8dcb121c749b182f5d9fd0d95aad83869942df13ba"></a>

<a id="canonical-1199aa41256a10a21e674d0894ba9f119fbcd72199838f8c3bd63b04c19a4548"></a>

## location property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 5

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

<a id="canonical-fbb5e1c961f78125244d0d8fc5073f391e7e116f730c0923b1076b6b9c7c87f7"></a>

<a id="canonical-a13dad88b9daeb44d8d7feb557ebcc40a0e18179e91cf799c3af574ca6c1781c"></a>

## store_provider property — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 6

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

<a id="canonical-abe49002cf21b3faab907eee72745a46b6819f58a7b98a355d757b1ccb0fb6b5"></a>

## Next pages — tls_intercept.custom_certificate.private_key.blindfold_secret_info / b8e69404d57a / 7

- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-7a9b9b1f9fcdee81fcea694f6eead012ed7a841f1ea4199ef91729dea077fba2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4216835e7c7fc9f4c3d2c58d08b3d351aeb5636babd5536dc65d5f0fc6accf9d"></a>

## tls_intercept.custom_certificate.private_key.clear_secret_info — tls_intercept.custom_certificate.private_key.clear_secret_info / 145f40e839c5 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-68bccfbb040ca55c6936c712c969bcb9a9aab7fb1f04abb189a5191cbf971e02"></a>

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

<a id="canonical-3ed3a81de53f87ef39b848d5be6de01b9585592750f3e2543629f6739d3e01e0"></a>

## Direct properties — tls_intercept.custom_certificate.private_key.clear_secret_info / 145f40e839c5 / 3

<a id="canonical-8b18ec08dc71a8ec20569200c95edf01af504c7e23f48e4df0c17b9f631ebaea"></a>

<a id="canonical-3ba9fd867ef3e5f505de6bc47b1b477673427e09e37721cb8a07d98097a336a7"></a>

## provider_ref property — tls_intercept.custom_certificate.private_key.clear_secret_info / 145f40e839c5 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c3fabdadf1ee11f0618330660c453f7d721785f3639de7b8ab2e654379de8e08"></a>

<a id="canonical-5181eaa8fe00201006ac78b9b79f0d61c22d729e9c3c83197124d68a5a034be0"></a>

## url property — tls_intercept.custom_certificate.private_key.clear_secret_info / 145f40e839c5 / 5

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

<a id="canonical-a438e2780d25859d83864cd48fcdd6e3e23b5d229ab695294994bd34a5ac3c9d"></a>

## Next pages — tls_intercept.custom_certificate.private_key.clear_secret_info / 145f40e839c5 / 6

- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-d4600eb811072dda7d896b3aa2c7d3020bda6a91d40a2b89cbcde7fc2eff12c7)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-72ba0999c99d0ccb921817e6c9637334daa45195b2895807ee7fcd77e78d66ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25012eb32180a6d6d357eff0dd01cca3e6542586dd14a73c50c3aef3ef0c773c"></a>

## tls_intercept.custom_certificate.use_system_defaults — tls_intercept.custom_certificate.use_system_defaults / 9ba3e068d75e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-e44449250351256d37b8fb1d78e992d925bdb943135a3d39f8548a6c17a97621"></a>

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

<a id="canonical-1ddf511804313e5212137248701b2254def76827302a391c9e001e9c7b325ee3"></a>

## Direct properties — tls_intercept.custom_certificate.use_system_defaults / 9ba3e068d75e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26ed9ccd60c00cdf92c9014b534ba65e766b2c72dde997fcbc923ca9355f29c2"></a>

## Next pages — tls_intercept.custom_certificate.use_system_defaults / 9ba3e068d75e / 4

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-6d6d054d6803562e3ce4095d6419b6c6f969015d8d089433882368a9009d5c28)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-ecb2a177ffba54be67bf4286a0deaed73e261813227e5337de84698479edc76b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26c822e81da1b6437f6644cbdc8104f68f07442eb08310a217918b4db385d6ae"></a>

## tls_intercept.enable_for_all_domains — tls_intercept.enable_for_all_domains / 7ad6e115b0b5 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- tls_intercept.enable_for_all_domains

<a id="canonical-7c29bb7cba68e6f1366727e73adec2130f869b775fc6dc3b061371e27294af75"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable for all domains.

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

<a id="canonical-561da31dc7d05d70d887e948ef0f7492ccefd7baf60001e04c8f8fd59636854c"></a>

## Direct properties — tls_intercept.enable_for_all_domains / 7ad6e115b0b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8392d093a0708d5c65c41c4fb7ade44b45c9831fcd1759c2c8fb206981d4b06f"></a>

## Next pages — tls_intercept.enable_for_all_domains / 7ad6e115b0b5 / 4

- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81de04b50939ed53329fb565b20456b06e033eeb43831f6fe2837ba043961c2a"></a>

## tls_intercept.policy — tls_intercept.policy / b13432354ee5 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- tls_intercept.policy

<a id="canonical-91367051670db5eb78bea3f7cf78a1486c6a3990ec69f7eee49d4d4ca1d05264"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9b32f0d29dff2a99fe9a7100e08414b7d093a82bcf06c4dadb87fc8846d49fe8"></a>

## Direct properties — tls_intercept.policy / b13432354ee5 / 3

- [interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696): complete subsection reference.

<a id="canonical-c9fe452369cf211c1e21ae0b67a98eb1f5ede720c9fcb47dac4e847cc5d119e5"></a>

## Next pages — tls_intercept.policy / b13432354ee5 / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-285d5b94ea7622e344237a668cb115df703af93ec4847b2e22ca5a547c9fe512"></a>

## tls_intercept.policy.interception_rules — tls_intercept.policy.interception_rules / 2d6e99394a00 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- tls_intercept.policy.interception_rules

<a id="canonical-573829d6fa9cb7a6d326732b36c369af33adbe5cedf9f370e6fb889971ef8949"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

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

<a id="canonical-cbe708b9ffb3ade02efe775262aeff1b17a147e4293a758babcac3433b0e0d49"></a>

## Direct properties — tls_intercept.policy.interception_rules / 2d6e99394a00 / 3

- [disable_interception](data-sources--proxy--reference--group-005.md#canonical-980bd2f4b00b4175cbb587fbdbf5d620545222bbc8da4bd272c4f407971784d7): complete subsection reference.

- [domain_match](data-sources--proxy--reference--group-005.md#canonical-31978c3e1db579b430e3a99fafff19261cb7991869fce92c062b0ab433579876): complete subsection reference.

- [enable_interception](data-sources--proxy--reference--group-005.md#canonical-4baed728f63a631cc553143528c873b7011627c48dbfd01ce9679ac6cc9033db): complete subsection reference.

<a id="canonical-52856745344cb316c14ac42b4b01b7df74e3ba189b3273254e2e2dbba8a241e3"></a>

## Next pages — tls_intercept.policy.interception_rules / 2d6e99394a00 / 4

- [tls_intercept.policy.interception_rules.disable_interception](data-sources--proxy--reference--group-005.md#canonical-980bd2f4b00b4175cbb587fbdbf5d620545222bbc8da4bd272c4f407971784d7)
- [tls_intercept.policy.interception_rules.domain_match](data-sources--proxy--reference--group-005.md#canonical-31978c3e1db579b430e3a99fafff19261cb7991869fce92c062b0ab433579876)
- [tls_intercept.policy.interception_rules.enable_interception](data-sources--proxy--reference--group-005.md#canonical-4baed728f63a631cc553143528c873b7011627c48dbfd01ce9679ac6cc9033db)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-980bd2f4b00b4175cbb587fbdbf5d620545222bbc8da4bd272c4f407971784d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cfe1a885b6ba0d805f2e02fca1c73c405f92d0704be4084fbb50b25d1cfe4a5"></a>

## tls_intercept.policy.interception_rules.disable_interception — tls_intercept.policy.interception_rules.disable_interception / 95472e9e4d38 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-ecd35760a2153b535ff3a9e6c55b8118b7bb39c052eab1226a1afb742738f03e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable interception.

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

<a id="canonical-7c07720e2b6792374908342cdc3caa9f530c5fd9bc049d48a0bf08e9e5fc0212"></a>

## Direct properties — tls_intercept.policy.interception_rules.disable_interception / 95472e9e4d38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db9a91f5f8b1dfec90478bea27d6486551278a1b02153c8a0c84c485aec3d2b5"></a>

## Next pages — tls_intercept.policy.interception_rules.disable_interception / 95472e9e4d38 / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-31978c3e1db579b430e3a99fafff19261cb7991869fce92c062b0ab433579876"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d157732a8c8935648eaff014699d65f3472e041e293fd243999e0c6669e23045"></a>

## tls_intercept.policy.interception_rules.domain_match — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-809662283b159958ca30107dbf4d980050faaef8b39c3bf9d084a5351eba0d59"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-ae16f3f7db23b277e46a13fc2414663e95f42429ccf14397285d32109d043cfa"></a>

## Direct properties — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 3

<a id="canonical-d97971ff85147bb7c85d4db406cb3b23321cabc5e43c8f73928114ec63c623da"></a>

<a id="canonical-be0977d3927e945bf0ba1ebebd875f10f2a43c19884bab1854f391df1c160406"></a>

## exact_value property — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-6452460b530b3855535d152fdecfc14423c1721feabea069acd26d5790d8e1b8"></a>

<a id="canonical-bfe7151128a679222917cfab4b057fd83696470e2741f80ff948a859d21c5aed"></a>

## regex_value property — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-44899ebbc98d6ae2972ed59388b62873928e13304739e72d70f46570b3259c25"></a>

<a id="canonical-487432bf88858a43f4998849473983b2ad186171e2753ecfd4bdfd20869810a3"></a>

## suffix_value property — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-4b39af95faa1f0d88274386c091285518e5eb068c3a895accbc809883b5e2029"></a>

## Next pages — tls_intercept.policy.interception_rules.domain_match / d4c029e49b66 / 7

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-4baed728f63a631cc553143528c873b7011627c48dbfd01ce9679ac6cc9033db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8ec9c77fcfd6a7546dad88a48c56e79460d5216e885175dfb69d518bf831ea7"></a>

## tls_intercept.policy.interception_rules.enable_interception — tls_intercept.policy.interception_rules.enable_interception / 6c8af515354b / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-d3e6ddf4d6986b6d894e3b44e376c3a2b59bd597a3816550a7f041694b523110)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-5022b4b0baab045c96d34bc868b54831c813c7ec712bb86f786979f1b4a9af17"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable interception.

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

<a id="canonical-ce8e62534afb8e3497ad1fd92be5ca65eba4049b1f97eaa92293d938ef97d393"></a>

## Direct properties — tls_intercept.policy.interception_rules.enable_interception / 6c8af515354b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efb732e85ace79b807f2468ecf61a7fc3aaf7f9d1f717279eaf2b7a17a19a341"></a>

## Next pages — tls_intercept.policy.interception_rules.enable_interception / 6c8af515354b / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-4b1c6177ccb23395c82f1936e152988a6437d3e4af3cc6172f543812384c4696)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-c81186c0316cc9113970ea7b9b30ccea8316c7a3acae7a83d09ea2fc188a7dd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7f016877ac9daf5724d78b7845033ecc768cb6f89188234c8f4ee1fcf8a8a21"></a>

## tls_intercept.volterra_certificate — tls_intercept.volterra_certificate / 3e0fca05f817 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- tls_intercept.volterra_certificate

<a id="canonical-ee31197a0753d7526ecde5972f9533e04f0a9eb7a80f672bf1a7bbbefd3aade9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra certificate.

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

<a id="canonical-12337b1f79c45102360255a02ce327afdc1f2823583e60fb62f752275f75d613"></a>

## Direct properties — tls_intercept.volterra_certificate / 3e0fca05f817 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9f9bcf010867a7aba8e0877d98034167d1e5ce6386130e50ac10b81d057fb9f"></a>

## Next pages — tls_intercept.volterra_certificate / 3e0fca05f817 / 4

- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-20a058a6b9bcb6d8622f44922a7417163865c4159331dfa258fabb72315c782d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b56433dd00427f56e97c086011d2d8c16dc04aba596550472745ed3768675f2c"></a>

## tls_intercept.volterra_trusted_ca — tls_intercept.volterra_trusted_ca / 90006238774e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- tls_intercept.volterra_trusted_ca

<a id="canonical-f4c3e1060dc323eb83f9baeea12eba19b7f8710a12497c25cb732b09f41cc7f1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

<a id="canonical-5230e79519aa2f320288b394dfed1d5a0d398620f08c517ccf7d65449c62074a"></a>

## Direct properties — tls_intercept.volterra_trusted_ca / 90006238774e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab3a876bceb4f20a81f68b4815303fefa5810c4a1e0a86c5c9e2cd1529eaaaad"></a>

## Next pages — tls_intercept.volterra_trusted_ca / 90006238774e / 4

- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
