---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-acfeb2e07a127c916f5d98bca5f179e008b329b76d505cc67ee3c91be1ca9762"></a>

## Next pages — use_tls / 6f46ee02d5f8 / 6

- [use_tls.default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-db57aadc97b42dba89563240a91f6af47cc17e420b4d7893a0e0b8c69095750e)
- [use_tls.disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-2d518a92dd087f658f61307220cf744714465756ac60518a3f10c735c3896556)
- [use_tls.disable_sni](resources--origin_pool--reference--group-003.md#canonical-b5b7d33a1c24c0132e9e622b828b15c3ae409345b1aa93ba56b2dd9883dc17d3)
- [use_tls.no_mtls](resources--origin_pool--reference--group-003.md#canonical-48de757ca201189218824316bbd23f0667a07e3ace00bac6187f161d9aa35529)
- [use_tls.skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-12f882e6573470126ee15f0a7306222599f694c38b2a7952bb364bb049223f61)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- [use_tls.use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-f89b3835729201ab6bf3817b0e742a65db04a6aabcaff540a54ee8954ee9d28a)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-50025ee6e2504d3ccf7005acea544e797f91c7565591c23faad18c3469f04194)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-c0f1fd7fb6a8c6bd28b606ef7816026802d315812d220663bc1a5539c0c74402)
- [use_tls.volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-3320a35d9af7db62bdfa410461c6c31b0acb87e9b009de36d0b36da595233754)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-db57aadc97b42dba89563240a91f6af47cc17e420b4d7893a0e0b8c69095750e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9803893f290df54eafcd34e41f4a35b51849324dbce635cc66b794da2195456b"></a>

## use_tls.default_session_key_caching — use_tls.default_session_key_caching / d17ceec54d14 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.default_session_key_caching

<a id="canonical-ccba32fca58c04f30c6ea094640d7e46ebccb976307af9bf364cf292100b52bb"></a>

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

<a id="canonical-c84049be073796cba02d01aeea42e830973ec184c382bba795f980ce82a4e052"></a>

## Direct properties — use_tls.default_session_key_caching / d17ceec54d14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b98f81b3c0b898e70823cbeee608761d7b9a921dca9c9d0a6a9eda6a2d9e9a99"></a>

## Next pages — use_tls.default_session_key_caching / d17ceec54d14 / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-2d518a92dd087f658f61307220cf744714465756ac60518a3f10c735c3896556"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe6df983f844b4382eb44c1611988335cb5af05fd5783b02bd7d9cae9d92d693"></a>

## use_tls.disable_session_key_caching — use_tls.disable_session_key_caching / 4475ac9997b5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.disable_session_key_caching

<a id="canonical-1ff14ff691bc8669668470d527c065bcfe4584069a545f66929c609ae63a2723"></a>

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

<a id="canonical-7024aaad175b37c2bab243f835cb47facdddf80dc228435e8ca64d302a6a1e0e"></a>

## Direct properties — use_tls.disable_session_key_caching / 4475ac9997b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0311c40c2810e9a70702cb8d418a144bf860209793246b0faac15ed06e0918d"></a>

## Next pages — use_tls.disable_session_key_caching / 4475ac9997b5 / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-b5b7d33a1c24c0132e9e622b828b15c3ae409345b1aa93ba56b2dd9883dc17d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83c5762c76c09cf7cb65ad0e518f0911a36d2380505878b280b0fd7cbf87b0b0"></a>

## use_tls.disable_sni — use_tls.disable_sni / 8b4938064a5b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.disable_sni

<a id="canonical-d9a6a90c052f6d4027b4763ff9f1cba15a55f24c3aa09bcf7118abce3d56e6a1"></a>

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

<a id="canonical-608e2ab7632e15efe7e78c3d791991459238e70d60b5f6a2df1479b8da6f3592"></a>

## Direct properties — use_tls.disable_sni / 8b4938064a5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1252b84342dd8c108d9a0444bdf24b744c42b37a0520679b9992b74eda9aad29"></a>

## Next pages — use_tls.disable_sni / 8b4938064a5b / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-48de757ca201189218824316bbd23f0667a07e3ace00bac6187f161d9aa35529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-233f688fb22eabc4cd840fe3014e5f479cbf17e9104a8877317e694d82b27851"></a>

## use_tls.no_mtls — use_tls.no_mtls / 130568915cd9 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.no_mtls

<a id="canonical-c58eebd52426320bce8e844d54523f7bdfb4078e111256756b2327a147bb666a"></a>

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

<a id="canonical-9d098ad408e0e7738cb9bf9b218c6d865177ed94e3358a146bb0fbf0b91ef209"></a>

## Direct properties — use_tls.no_mtls / 130568915cd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2995a37f0e89a1aafb5f999a793c9966b9ae6218bfa64a5a3691e9b6fc090063"></a>

## Next pages — use_tls.no_mtls / 130568915cd9 / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-12f882e6573470126ee15f0a7306222599f694c38b2a7952bb364bb049223f61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0abcfa4c4d810ab7a7fe95c72d393f0dc1444cdd354bc0f2eed738629c57d943"></a>

## use_tls.skip_server_verification — use_tls.skip_server_verification / f230c3a6ba3b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.skip_server_verification

<a id="canonical-aabc3415320c442785d10159c2d66dbeb5df5ddca5b58f9f6bde8ba917846db1"></a>

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

<a id="canonical-0f47f599c3f30ec91d3e364d6d125518b0884e54ee92480dfbeb8d1f2bb63be7"></a>

## Direct properties — use_tls.skip_server_verification / f230c3a6ba3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67d3e74b95a73b1d479ddda2b851e1a8d177f19a8471910cc60a314ff0701e23"></a>

## Next pages — use_tls.skip_server_verification / f230c3a6ba3b / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24da27b3e19c217a2979b123700ff490360b525e6bf18d3396b2190077f8549d"></a>

## use_tls.tls_config — use_tls.tls_config / 6d3c609c6444 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.tls_config

<a id="canonical-48d80a573e5abff2da2251bc8116d41690f34e990bd6b2b0627dbbb7a2a1b7f0"></a>

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

<a id="canonical-aec8bb10ccee085e781115ae9251e444589b1c11b46f129163fd0fb5ccce1983"></a>

## Direct properties — use_tls.tls_config / 6d3c609c6444 / 3

- [custom_security](resources--origin_pool--reference--group-003.md#canonical-f9a30efc8510a1f6ebf6bcde79ec87a4e74c79d16af2b1a5f844aabcdad89e6d): complete subsection reference.

- [default_security](resources--origin_pool--reference--group-003.md#canonical-04d9c7a1003f68cd7c9eee24eb2321d95ac53f4277445515768da61e40786df6): complete subsection reference.

- [low_security](resources--origin_pool--reference--group-003.md#canonical-96baa02ca2e9e8087548d5a377cb0b02a62c887c481bde044ae23d7965775814): complete subsection reference.

- [medium_security](resources--origin_pool--reference--group-003.md#canonical-f7ca8aa6502404b3bcfb41be7e98b843c5cb20053ea23efb366bdb2e242e4629): complete subsection reference.

<a id="canonical-6a3332576bbe3b95ab1b6e39921c7fbe7f4fcb8f29587b948b1ad2a1e87543ab"></a>

## Next pages — use_tls.tls_config / 6d3c609c6444 / 4

- [use_tls.tls_config.custom_security](resources--origin_pool--reference--group-003.md#canonical-f9a30efc8510a1f6ebf6bcde79ec87a4e74c79d16af2b1a5f844aabcdad89e6d)
- [use_tls.tls_config.default_security](resources--origin_pool--reference--group-003.md#canonical-04d9c7a1003f68cd7c9eee24eb2321d95ac53f4277445515768da61e40786df6)
- [use_tls.tls_config.low_security](resources--origin_pool--reference--group-003.md#canonical-96baa02ca2e9e8087548d5a377cb0b02a62c887c481bde044ae23d7965775814)
- [use_tls.tls_config.medium_security](resources--origin_pool--reference--group-003.md#canonical-f7ca8aa6502404b3bcfb41be7e98b843c5cb20053ea23efb366bdb2e242e4629)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-f9a30efc8510a1f6ebf6bcde79ec87a4e74c79d16af2b1a5f844aabcdad89e6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6beacd493ba32ae6849e9b82f8de5b2b6532cef4c520dfe5341fa8d484104c3e"></a>

## use_tls.tls_config.custom_security — use_tls.tls_config.custom_security / 76488b324cb2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- use_tls.tls_config.custom_security

<a id="canonical-23cc1d74ba0a67be4461bb6e5702355aa1ba5d0a4e6128af3e2f16127bcb7fde"></a>

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

<a id="canonical-d6648759535a7e3d86310c1481d945697b131b456096d2ea04142d0ffff319e9"></a>

## Direct properties — use_tls.tls_config.custom_security / 76488b324cb2 / 3

<a id="canonical-e31e33f893cb760d94c69a584cf4704d395b8d26484c3ef982390073364e673c"></a>

<a id="canonical-245202b1adcb1328d9cbcd93f94a277b3c6fabc73632fc7424a392d9cccdc246"></a>

## cipher_suites property — use_tls.tls_config.custom_security / 76488b324cb2 / 4

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

<a id="canonical-1a9d483511b01fb81c22169ed7e12503ebd39d9eee7e4fd7bbe1d56cd108e786"></a>

<a id="canonical-e99327395147ff59aaf47076349410038b0bbd56b4cb9fcc7460e957f8c0c826"></a>

## max_version property — use_tls.tls_config.custom_security / 76488b324cb2 / 5

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

<a id="canonical-6fd72f8fd305bd79f7a14c37c6daf4fb65dab7221b073426afe7118bff8def62"></a>

<a id="canonical-d19dc4679cc9a33ed59886c9a9974b2004d817f62faa97e4e1be0a104aa567c1"></a>

## min_version property — use_tls.tls_config.custom_security / 76488b324cb2 / 6

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

<a id="canonical-4aafab8ccd99430345c63a46f2c5eaddf1fad4e3370985574d2155c2d08b075a"></a>

## Next pages — use_tls.tls_config.custom_security / 76488b324cb2 / 7

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-04d9c7a1003f68cd7c9eee24eb2321d95ac53f4277445515768da61e40786df6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d032f7df978bc7b542024a17b4e44cd06da5156c84b6396804f9a7724fe462dc"></a>

## use_tls.tls_config.default_security — use_tls.tls_config.default_security / 98183d296310 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- use_tls.tls_config.default_security

<a id="canonical-a0b2316e46d2d46972a8367a009cfcf47943e976e28bece8d1ef78c17dc5543e"></a>

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

<a id="canonical-9d1eabc0afeee647c4bbebb05068b58c8c3e40a44a222016617e0447007bd8ee"></a>

## Direct properties — use_tls.tls_config.default_security / 98183d296310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa9c0c66602605f66637c6ccec596ab7b682eb3015e376d401a78a11766a6a3b"></a>

## Next pages — use_tls.tls_config.default_security / 98183d296310 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-96baa02ca2e9e8087548d5a377cb0b02a62c887c481bde044ae23d7965775814"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fa1a33b82ac2d2b0d0613fdf592d9ff8c027d9e29636a61b3b7774be7bc8d36"></a>

## use_tls.tls_config.low_security — use_tls.tls_config.low_security / bb6057b3eaa3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- use_tls.tls_config.low_security

<a id="canonical-b094541db67703b9df23ab03ccee1430080c2a084d9929f4a23026c70eeb56a7"></a>

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

<a id="canonical-9667eaa7bbe4f3588f83dadf93fe1df8cf5b73a869a323ebeefb102319074043"></a>

## Direct properties — use_tls.tls_config.low_security / bb6057b3eaa3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0c020015d78bcac8eecb5d458b80cb5fb627f2be2235319a7a4ebe63077a45b"></a>

## Next pages — use_tls.tls_config.low_security / bb6057b3eaa3 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-f7ca8aa6502404b3bcfb41be7e98b843c5cb20053ea23efb366bdb2e242e4629"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-882d48b93a8e4ba8c6cae32e2b71b88e46a7ce0d6e5ea8239c159be09ff1ecc0"></a>

## use_tls.tls_config.medium_security — use_tls.tls_config.medium_security / 49d6e9d484d8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- use_tls.tls_config.medium_security

<a id="canonical-e88bbf350348e3ba2b3d193987f8c144a726cefa327ca141bc74a4b348cf7254"></a>

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

<a id="canonical-247639175f16f7c5e1cf95eaab8ec7ad25f07bf3300eb14a6988f4203082413c"></a>

## Direct properties — use_tls.tls_config.medium_security / 49d6e9d484d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-893d08f8cd2d7b5d2384b7a5103748befa5a566dbe2af1449a14ebe5d0b1c5d4"></a>

## Next pages — use_tls.tls_config.medium_security / 49d6e9d484d8 / 4

- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-56a9eaa2d7444586b13cd43e728c690562529568863e6fa1625bdcc3e182ea5d)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-f89b3835729201ab6bf3817b0e742a65db04a6aabcaff540a54ee8954ee9d28a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-020eb932eaa6e6033f0beefb352e3244f3ac5fa71cca75e6ae33ad986558eb0e"></a>

## use_tls.use_host_header_as_sni — use_tls.use_host_header_as_sni / db260fdb215d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.use_host_header_as_sni

<a id="canonical-c4bd4b873837edc5c2fe5b13eee5d85a6cd0d5231fa78694dd638250b3b95474"></a>

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

<a id="canonical-8e74008e22deadfc126d03136e929b24938a9657ae654a1a8ca860279a5a09a6"></a>

## Direct properties — use_tls.use_host_header_as_sni / db260fdb215d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d72a5bc3161f6f586de90d055e9130ad4e2cc00ea3c3717e972eea68b4a02f88"></a>

## Next pages — use_tls.use_host_header_as_sni / db260fdb215d / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ccc0ed32cb0b228014d4743b280f1a5a6b42c73feab25e517d8ed3f6db76eac"></a>

## use_tls.use_mtls — use_tls.use_mtls / 73acd756ca02 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.use_mtls

<a id="canonical-45af1ade87992d9a2998f6d5fd1be5d866864b3c6370f93333249e5b3f293625"></a>

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

<a id="canonical-d78ca82a66c294d6228763c739eb58ef4d431a1c22609e2c2bbd791307c5dac8"></a>

## Direct properties — use_tls.use_mtls / 73acd756ca02 / 3

- [tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8): complete subsection reference.

<a id="canonical-71629a90b537daf19bc7b03d2b9b90a7c9c766b78f21051d94c532bfe28a9c67"></a>

## Next pages — use_tls.use_mtls / 73acd756ca02 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a1e41595ea71e51c672baa917ff0355f949677dc3c08e2b871eb9ad37190830"></a>

## use_tls.use_mtls.tls_certificates — use_tls.use_mtls.tls_certificates / 7a0da201e6da / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- use_tls.use_mtls.tls_certificates

<a id="canonical-fab8c0218b7f95836c557b571310e6300576d4ef8640357c499f1a1398118ceb"></a>

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

<a id="canonical-4cdffed073a5c16ff6fc860dabbf67d1050ad19730e771be5d4cbd1e9705f372"></a>

## Direct properties — use_tls.use_mtls.tls_certificates / 7a0da201e6da / 3

<a id="canonical-48fff4f0e30652c53d4363ce7dee2613433b3defbafcbd7403893a56511ba1c4"></a>

<a id="canonical-a4de78ce23933181130dd900f36e0b997d02d9412a7f62e16f8089f5a3dbd689"></a>

## certificate_url property — use_tls.use_mtls.tls_certificates / 7a0da201e6da / 4

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

- [custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-a287dc25ab5c5449acb8d477f892a9bab863cb8bf5c48070016e00a4df641924): complete subsection reference.

<a id="canonical-b712e06a7926c4ea400834c0d529ccbba58f523d7888479a4708fe0edad0dff5"></a>

<a id="canonical-283b299534d5eab3e4e8db138f0d867db94b20c86ab08dedb428df38a48feabe"></a>

## description_spec property — use_tls.use_mtls.tls_certificates / 7a0da201e6da / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-81c2fbe4d08ab2ee25774d767023d7df84d16f20aa8972aca72eafb3f8ae1f6a): complete subsection reference.

- [private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6): complete subsection reference.

- [use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-5eadd0478059e99796d6d3c4220e6e8dd159d1280ddd94b58b0f85ca65f1a317): complete subsection reference.

<a id="canonical-c78712dbc33d53345f9ab11941ac5ee3c4c7d09d98dc2e0025535e4a622d982f"></a>

## Next pages — use_tls.use_mtls.tls_certificates / 7a0da201e6da / 6

- [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-a287dc25ab5c5449acb8d477f892a9bab863cb8bf5c48070016e00a4df641924)
- [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-81c2fbe4d08ab2ee25774d767023d7df84d16f20aa8972aca72eafb3f8ae1f6a)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6)
- [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-5eadd0478059e99796d6d3c4220e6e8dd159d1280ddd94b58b0f85ca65f1a317)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-a287dc25ab5c5449acb8d477f892a9bab863cb8bf5c48070016e00a4df641924"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bca5ac8794a9e06784341a6850422d95459be358075797f040fad0b7a7ce2784"></a>

## use_tls.use_mtls.tls_certificates.custom_hash_algorithms — use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 09d36d4b6a13 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-a7ae552bc9c62cf9c0475f71145d74000d32c0e8921159d83e83c046b4175cd6"></a>

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

<a id="canonical-e9753c288f1ad779517762f46692b0be0d317373b0a49f9312de65d8f88f644e"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 09d36d4b6a13 / 3

<a id="canonical-3951c323dbb70c7a10f00141e35bd0d2c9f717fa30f6a31bad718c18d9a7cfcf"></a>

<a id="canonical-bbd8b09367a241074f9f1589aa6fc5bdb986380b9d1a043867e9c7b1ae3251b8"></a>

## hash_algorithms property — use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 09d36d4b6a13 / 4

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

<a id="canonical-126b6e962b21106fed6e7a7fea00dc94f46f8e7390c702a7049720e4616f5c22"></a>

## Next pages — use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 09d36d4b6a13 / 5

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-81c2fbe4d08ab2ee25774d767023d7df84d16f20aa8972aca72eafb3f8ae1f6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-424c6a023655309f1ad7f2f94a2196cdb57a88fcd5acd59370a02eee6f5a6d21"></a>

## use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 4e23f9cb4d36 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-ffc2196f4ec59c808aafe69aab57d4a37b4fe4c8f2b0c6ae66718e49558f0f3e"></a>

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

<a id="canonical-096f157b0f9c6e8ae3e8c2be76d6ae94a489d28dd53e4a55c236d2e6d3aa8465"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 4e23f9cb4d36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c36d1b9f9acf927edaeb0e62adfd941e5c57d9922c5abea048870fca1503d551"></a>

## Next pages — use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 4e23f9cb4d36 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7045bf210c64d4e5fefe641708f0f36a12d12cbc79f2ab98024b128d78bc7af2"></a>

## use_tls.use_mtls.tls_certificates.private_key — use_tls.use_mtls.tls_certificates.private_key / 3ba256ee9205 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-b66b0564b88effdb1412803f9fa66e98fe6f99ef96c340068eb48d8e01e67294"></a>

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

<a id="canonical-23f8f6c406dbbe6531645eb0b7b67e9eac82897a4062c4ff3f67858f80292a9e"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.private_key / 3ba256ee9205 / 3

- [blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-aee57a25c98c72b7130fcaa378466e2f920cc3a096d19b2b397e151cd6a6cac0): complete subsection reference.

- [clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-66162bd1feecd97562f97dd794b40f31c366f5fc448f243f03b7e9bf9c034480): complete subsection reference.

<a id="canonical-a93d50c2adee40f71fe6a005353b3fb03c1ecb6719bfbadd943c09ad8d3038ec"></a>

## Next pages — use_tls.use_mtls.tls_certificates.private_key / 3ba256ee9205 / 4

- [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-aee57a25c98c72b7130fcaa378466e2f920cc3a096d19b2b397e151cd6a6cac0)
- [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-66162bd1feecd97562f97dd794b40f31c366f5fc448f243f03b7e9bf9c034480)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-aee57a25c98c72b7130fcaa378466e2f920cc3a096d19b2b397e151cd6a6cac0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6396b21a5569a27f5e04d2f8aa70d9ac94eb5792c772f498eba06910e94b8d87"></a>

## use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6)
- use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-793231fc84bd04d21af688065f28b33c17fd0b6bdd4917cfb3568e07c732f6e4"></a>

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

<a id="canonical-093d1df619e30c0e9ba4e2fb588963c6319204dad20e0d9f9b2af97b53ba3fd6"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 3

<a id="canonical-87f71613a458f22636023170d80dd4e7e31ed324fc36283cef0a61c931ad7681"></a>

<a id="canonical-b3b2a021e70704caa960d469c6aa3b9ea2ab70a1d1ca11ef5b88ab3fa4281a2d"></a>

## decryption_provider property — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 4

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

<a id="canonical-d2a40c00442c74541e0116d85d5ec572a7cbf3ee8c85a909a7bca28dcc06604c"></a>

<a id="canonical-3a05c10b0dea588e1e65d356cd6d039aa659114c5930d20c6b5f8c3d45e02010"></a>

## location property — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 5

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

<a id="canonical-0fd46ad65c6978409bdd43e53cacbaffa63fd305e3eb6ee7ae3c3859c7c9dd22"></a>

<a id="canonical-6b6a20eab0361d0d0c43fff83e2a9d35e7a7120fd34d13504696d2865d3b53f5"></a>

## store_provider property — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 6

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

<a id="canonical-604aee7e72f9cb380c457e64fed7e961a8974473964089e4af56f2745f05fb33"></a>

## Next pages — use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / d63812b9fc7b / 7

- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-66162bd1feecd97562f97dd794b40f31c366f5fc448f243f03b7e9bf9c034480"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f59e11754bc5f1db22421595a3166d30e36958002190af4e2a41e35f59d3c45"></a>

## use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1c4866469288 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6)
- use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-957b7d05b684c6015b27ea4e91117987e62b80b3b7f710fcbd91fde0cb771d8e"></a>

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

<a id="canonical-66c52b63744f1d6894f0d6106fc21aa7df4dc0f166d3bd9fbc3ec04893d8c232"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1c4866469288 / 3

<a id="canonical-0f3a90a79aec881baee92fee3b87e838da274eee2c1ad313de93926586837d0f"></a>

<a id="canonical-a31021aa5573ac8c79ae12d7fbb3a4b566ad02665fe5261c58ed2caa02544b71"></a>

## provider_ref property — use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1c4866469288 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-65f96d206e089ac3d81604af7e54ea32ba8a49c8208a5e0d84a04ca329c0d94c"></a>

<a id="canonical-40c38ebba65b4cf12aa18799575615e6f3780f8a93a646504806707125248fd1"></a>

## url property — use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1c4866469288 / 5

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

<a id="canonical-10aead5d6aa37ab7bb95a2cf10ed26107b0248bbd2504ed6143666a759453e81"></a>

## Next pages — use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 1c4866469288 / 6

- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-d7dc11f5f120022c17f5f44dbf0b3e8539afe7c237886641d210c694b3a5ecf6)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-5eadd0478059e99796d6d3c4220e6e8dd159d1280ddd94b58b0f85ca65f1a317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51aa32c93ce5edffa056178dba5dd055471afad60b88c93a3375799dfb574889"></a>

## use_tls.use_mtls.tls_certificates.use_system_defaults — use_tls.use_mtls.tls_certificates.use_system_defaults / 2d559697bda0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-58aaa3c6b66d808f551ad41b0ac3978d46b44a6122f283547e690bfcea5d4845)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-31d81d41c7a01d5944274ca77fc0db24145fb1c7e18607e9feb33a2af328836f"></a>

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

<a id="canonical-98fb31cedfe095bf45da81762f546b07efb46712f255684cc92c4d4293b5d032"></a>

## Direct properties — use_tls.use_mtls.tls_certificates.use_system_defaults / 2d559697bda0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac0903cad1b5fcdb7f137be2ffdb5401b325d297b784ae4205abc56be7894240"></a>

## Next pages — use_tls.use_mtls.tls_certificates.use_system_defaults / 2d559697bda0 / 4

- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-d4ee3bbe5f172423ca41d088bbe979ea90304ca2011ddb2d3b5858df37cc0db8)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-50025ee6e2504d3ccf7005acea544e797f91c7565591c23faad18c3469f04194"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31ea6d67ec3ac825ec1ad7e83d4bd14d1a86d38014db41eb7b9d700a2637b096"></a>

## use_tls.use_mtls_obj — use_tls.use_mtls_obj / d628f13ab1d3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.use_mtls_obj

<a id="canonical-3520e5e6634d0ff0d817cf349f31d5f4607e7cc113f149b83596a1bb40c78808"></a>

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

<a id="canonical-b1d76ceee7987f203840fb17cb878e96665a8434251ceb328ee5c332f0875a6a"></a>

## Direct properties — use_tls.use_mtls_obj / d628f13ab1d3 / 3

<a id="canonical-2d95cc8ea2403adddf24a1184da8ed67f4545f96d3a481e195360cbc3b337fc1"></a>

<a id="canonical-87de9156ac21eddd0d9628a8326cac89c4a232be53468f804e892da51641ffdc"></a>

## name property — use_tls.use_mtls_obj / d628f13ab1d3 / 4

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

<a id="canonical-e482b2717b9434b2eceefc9fe9480552eb61cea581b8b99a2fd961746fb02e58"></a>

<a id="canonical-604f64fe20104fa6b00ba8b6e3b714ee7cd33102e9ccb14d9e9a5718afa6e02d"></a>

## namespace property — use_tls.use_mtls_obj / d628f13ab1d3 / 5

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

<a id="canonical-e0528a56de892151cbc81f6a383a37469e4c6c3711eb71466f16d0157f1723d2"></a>

<a id="canonical-dfaa643e2dc0ed196d68cbdaeb16c5cefa683feef0fbf47ca8b0febcc8955d26"></a>

## tenant property — use_tls.use_mtls_obj / d628f13ab1d3 / 6

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

<a id="canonical-c07d02022a28ae01ec4e21e2da585028d77915bf78a1c5229b4f2d04dd1acd21"></a>

## Next pages — use_tls.use_mtls_obj / d628f13ab1d3 / 7

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-c0f1fd7fb6a8c6bd28b606ef7816026802d315812d220663bc1a5539c0c74402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec730a3ea859ffc4bc7f0495b35f0037b297f6db2b0b1255650d6acfd8d6a075"></a>

## use_tls.use_server_verification — use_tls.use_server_verification / 2fa61345189d / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.use_server_verification

<a id="canonical-6229777aebfda37586b438d4a060f4f7336d718149760aa4f2a091b2992fdacd"></a>

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

<a id="canonical-ab4100db6eb4a3b2f010f930274488156908ee3133ae9f5e3d29b5ddd69690dc"></a>

## Direct properties — use_tls.use_server_verification / 2fa61345189d / 3

- [trusted_ca](resources--origin_pool--reference--group-003.md#canonical-187939fc69c3853cb9c4f55f8702d318fd9b4a2d1db4fa5d82b6ab54b41ae565): complete subsection reference.

<a id="canonical-687cb5851f4fbe47acb950c0753c5f5b82facb5e24dff2b65fb896aa35fc7d78"></a>

<a id="canonical-8a17fc899deccbea22594e781abe7c50a6096d9e66156efd9fdf965164855d6c"></a>

## trusted_ca_url property — use_tls.use_server_verification / 2fa61345189d / 4

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

<a id="canonical-a636079a2b08d17fc3312d97aae2d61ad344efb99f9828e42cbdbdcecb4fb09b"></a>

## Next pages — use_tls.use_server_verification / 2fa61345189d / 5

- [use_tls.use_server_verification.trusted_ca](resources--origin_pool--reference--group-003.md#canonical-187939fc69c3853cb9c4f55f8702d318fd9b4a2d1db4fa5d82b6ab54b41ae565)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-187939fc69c3853cb9c4f55f8702d318fd9b4a2d1db4fa5d82b6ab54b41ae565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4ebac2947145c651da62d359c3fab50ef5687d961fc9ad924090909ff35921e"></a>

## use_tls.use_server_verification.trusted_ca — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-c0f1fd7fb6a8c6bd28b606ef7816026802d315812d220663bc1a5539c0c74402)
- use_tls.use_server_verification.trusted_ca

<a id="canonical-e5614161190c60fde0652313d0982a4dd055fc8f2696bf9d032e8eeb33f18bb2"></a>

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

<a id="canonical-6cce78a5f4b49a8adeaf05d3253171b4b94825aba6a3bf4be20d7755393a093e"></a>

## Direct properties — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 3

<a id="canonical-b001e6cf93141e04038668b59d889c1910c7e35cd03630d5e70397dde80553b9"></a>

<a id="canonical-78a8d0c77f88d319144c9480164af45d778a9d38669f1716b7882bf3157511e7"></a>

## name property — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 4

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

<a id="canonical-19956175ca098843607e4098d3d16c397dfbc1e97a52f759a9ed779e3c4e7c3d"></a>

<a id="canonical-898dfd5a872943c912be4be1123e2a6e612f612292ea43063cbc546a35d77838"></a>

## namespace property — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 5

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

<a id="canonical-fb45b5ebd6c0dde13bbe7f2a5a18934167d6dd64791bb758685f1df1ae6b4f84"></a>

<a id="canonical-1fae9fd9390c9eaef6bc91f079453831494631ae1438b99c713faed1ae144068"></a>

## tenant property — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 6

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

<a id="canonical-76c8a4b42527942eafc8df658c6f95780cf2ff2a2140535cc8e71db3c6fdc8af"></a>

## Next pages — use_tls.use_server_verification.trusted_ca / f4dad21a0584 / 7

- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-c0f1fd7fb6a8c6bd28b606ef7816026802d315812d220663bc1a5539c0c74402)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-3320a35d9af7db62bdfa410461c6c31b0acb87e9b009de36d0b36da595233754"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68f264e08dafb3a6dce4c574c31c9c8534431db7fe0d533b1b988e59c3dfae5e"></a>

## use_tls.volterra_trusted_ca — use_tls.volterra_trusted_ca / 93585f47c649 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-e61b24df1f4941cd9da6a249809a952cbb3c21cf96e76da5273da1bfdf5a1c47)
- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- use_tls.volterra_trusted_ca

<a id="canonical-42cb4c427226d97d0b39813647ae5c7b7b6687048b45cad7cba573a986cd3b8f"></a>

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

<a id="canonical-a567aba0896c8836cce4cac27fdf5c217fe3eabf0290dd9197424c9d6f6f45d0"></a>

## Direct properties — use_tls.volterra_trusted_ca / 93585f47c649 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b000eee956bd49988b95760c6d5d63b6eb352d3e00698b2aa125b562ff8f75a"></a>

## Next pages — use_tls.volterra_trusted_ca / 93585f47c649 / 4

- [use_tls](resources--origin_pool--reference--group-002.md#canonical-18c27b6d6857cf831ee951a3437c0ab831adbcbcb92c8fd70efb90f76fba4fdf)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
