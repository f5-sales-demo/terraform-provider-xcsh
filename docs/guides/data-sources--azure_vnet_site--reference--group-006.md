---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-4064bd98763e6e3e93519fb11227d6c1a8aee49cda27d8ca4d8525b96de19338"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-5868a5951725e76be735867caf00a52b90339253731333c16b2b04e680c38a8f"></a>

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

<a id="canonical-0cc50259db9e9f2d3223dae309686988a4128f93c4196a61eb1e7a04cb68a8c9"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 3

<a id="canonical-bfb23ddcfe275c02c7328de53f3c8a925c62f4c562534ccefead7964444726c3"></a>

<a id="canonical-3eecbc74b37f8c530e23ee0486b6d7b44da00cee53042397237b50ed4c1ab626"></a>

## decryption_provider property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 4

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

<a id="canonical-687af310fc1b63ed5e13ddf9066744d2da0260f03542afedff0c9ead574e9f9b"></a>

<a id="canonical-4f346726197800321e36dcf36b59047c2ba8cd8c3fb031af9254536f30f7464e"></a>

## location property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 5

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

<a id="canonical-d70317339a129825de8774d443aabf4eb1b0a160f1f2f2aef31a1b37c0820eee"></a>

<a id="canonical-c5665bc2197f4dd8430b2f3867c2176fa5a89ede0d6db7c52f2f18cc1deec244"></a>

## store_provider property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 6

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

<a id="canonical-411acba06ae8e4caf90f839cbe9690d178a667e7f1397990dd168949a203c756"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / b55c91cae981 / 7

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-69777777604c41afdbd849d3c55ee081f2a6c4069540639264a034cb44904a06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc2b15aba18a95097d60a2744e882462eb011ba142a01509f3b8f5e8794cb695"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 5213ef6721a9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-991535880a4c0bed47506ac20e0321cee2e18874307d6acd4cff079565abf28c"></a>

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

<a id="canonical-6b87ed6c97f40b08d53ea2ba1f1a62602c037aa543e61415cd072c86b1718585"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 5213ef6721a9 / 3

<a id="canonical-8168ee82defa97b3dd6e60ff5d38643bf5fd2c6c8832952309b8dfb6927ab830"></a>

<a id="canonical-86ba90d3b6c834f40cd4890954d9392204d8c49bd9eb566f00a8ed093378cf44"></a>

## provider_ref property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 5213ef6721a9 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c06c8f604e947b882989214eccf12ed829fe867800092a2675cc3fd705e66054"></a>

<a id="canonical-3d40647efce67ae4e8424929b37fd68988aa532193497b551f407d17de7c4b1d"></a>

## url property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 5213ef6721a9 / 5

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

<a id="canonical-c7a496a547eea36c53e28ccbd6f296acb2d06476d097db11efaae22d60df7301"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 5213ef6721a9 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8b3fd92aacfcb186818fbd2890ea4b7fa6fd03bd4af97c65aaa065c4dbbbacd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d48642cf2f5bb27d722942cef5ec7898171af7aaf86cf01e9f9f1f48bd6a5161"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / d992ee05c9c0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-db4636e5aee7a770591d9fa773a18a381bffd89bd82a2e631fc5f875d5d52677"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise to route server.

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

<a id="canonical-19aa58b7f593fd3fbd05861a4ac0864f12403318aeadfc38738408bec8153d0a"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / d992ee05c9c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8c8541ad569b446ac028c4947aeaef77bd555425c24463ee83d2b2b1f8b7738"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / d992ee05c9c0 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e1f2cf316a7cb29e597119b6382ea72e98ccc4750192a3af03b43bc4fb100a7"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / f940f30abb97 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="canonical-71e95bc4f1b5d0c9cbd820c6080d948e4b4a1a911601a7d993d23b683fa4df16"></a>

Type: `"single"`. Computed.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-8b5b3d7ccf1aca8644891b4c9188b0828fb4ba7c1ea5ff2cfdacebca75b02721"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / f940f30abb97 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-96641229840dde3d88f6c74c8461f472fc785d3cf513e2e806d2af5844778198): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-98f16a6fe0757fd25f213260ebab152839253bcf465a74d0020e1439d9f72ccb): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-5efb590e6cbcaab93c6586701cf195bf240c0b6b6d37bf6c0bae2ae3dc994bea): complete subsection reference.

<a id="canonical-ea2c69757aed77870e5c0e5e60b6a9ee2aa400387f195c5c72a14fd370731a36"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / f940f30abb97 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-96641229840dde3d88f6c74c8461f472fc785d3cf513e2e806d2af5844778198)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-98f16a6fe0757fd25f213260ebab152839253bcf465a74d0020e1439d9f72ccb)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-5efb590e6cbcaab93c6586701cf195bf240c0b6b6d37bf6c0bae2ae3dc994bea)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-96641229840dde3d88f6c74c8461f472fc785d3cf513e2e806d2af5844778198"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0709e0cc13a9f06d5c9eb42202c676c35afad6e6b3f46882123ba0fd54c5d591"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / d629f4366a36 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-065bfaa6075076b177589026413129606088c373ce242fbc6bf1e64bb3f82f5c"></a>

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

<a id="canonical-5e5398ebd2541ca5c1a13b2318f1d894c2755ea3c975ced09399073031c9d9cf"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / d629f4366a36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fbb5320f31f993aeadf45a2a31fe113de4cca0b49864593f26a18430506b1af9"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / d629f4366a36 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-98f16a6fe0757fd25f213260ebab152839253bcf465a74d0020e1439d9f72ccb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80c963d44b45d4193040e78b3602aff8188665fa37e57b90e9baba51dab573c2"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / f82c6b056737 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-3e3714cc6589f40017e6e7dd9dd63a23ada4310449659882a50d92e6df6aa992"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-731f8f3bd498972d7763024e3abbe28abf7e36257f7cf87eb03b114831261910"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / f82c6b056737 / 3

<a id="canonical-938135a4e2031a4c4bf2c1496700603d7fe3de22243a63a5da24dfd6dc60ecaa"></a>

<a id="canonical-2a49b1268cdab39928c021b1ce1530960b66a8c0360797fe2f2c8fc2c3a82eda"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / f82c6b056737 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-adb4a96b4b837795ec89610d6541239b1b59552be25559c0d1f818c7e7397b05): complete subsection reference.

<a id="canonical-869594fe95a6a2315fdde93e35963cfcdc07dbf9d6a584d4b0fdeb2f250b30f5"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / f82c6b056737 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-adb4a96b4b837795ec89610d6541239b1b59552be25559c0d1f818c7e7397b05)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-adb4a96b4b837795ec89610d6541239b1b59552be25559c0d1f818c7e7397b05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3786ba410ae9fdff2c6cb47bc0de58da2d4c8d3a6967bc7be3d0a6e30548c879"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / 92648aa38152 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-98f16a6fe0757fd25f213260ebab152839253bcf465a74d0020e1439d9f72ccb)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-52481be121ecbe98be3754a2a9f1715c4f0dbec16183a04224e5457ba03d0c30"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-74bd7796339482bc37a2874b2378b01e8c7caf8dfda3db8ed934e008aaf23b3c"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / 92648aa38152 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fc32221f21d1927707c57352930ef1ae80b611388e2d36e702b54db27cac86c"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / 92648aa38152 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-98f16a6fe0757fd25f213260ebab152839253bcf465a74d0020e1439d9f72ccb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-5efb590e6cbcaab93c6586701cf195bf240c0b6b6d37bf6c0bae2ae3dc994bea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-038127496446fca2a6ec43cdf772b580bdab3f9162251b9de8b1a3d3f73b27ae"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / 5ea805926d0d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-2f6184df51c153076fae61db0468d13fcc41106fe5587a6da765fe2e9fbeac16"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fa8b369f7b1f141da3d32f95096ae709e8691e2f7267fc6e45686961ccbc3b7c"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / 5ea805926d0d / 3

<a id="canonical-d2f7ab292866f8b23da9f9aa062ed6a7b0c2cc9bbdcb7f4ae8cbcde2fd73e46d"></a>

<a id="canonical-cbb2c5f869b5497765f2e92aeb806e7a6812dd28a5d3faa88219a9ac003a5712"></a>

## ipv4 property — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / 5ea805926d0d / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-05570b017ecebf3bf7da1a297be6292b4a6a0fc8f8ac5ec667b07d57df4247fd"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / 5ea805926d0d / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f07ef72f623c86e88abe948dd366014976ec25786244f630cbf39abd58d69e4f"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 62f43d122f98 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

<a id="canonical-a1bc6e4ec0e0891345f4b6591e716a604a14bdf34dfdf0240fec6df70542856b"></a>

Type: `"single"`. Computed.

Configuration parameter for route server subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-ec23df0d40521e21433e4cc6ec30e62a491b9ecd5cd2580a73fb4b4c4f22b1bc"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 62f43d122f98 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-52b5928aee3004d37f768900daddff3b151b3a83b0984d6f613dfeb57d6455bd): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-e581022db77f8d9b3a2003006ade38e71fbdfa4a7b36ede85c9b18dc97ca8535): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-8b2754e031ca410dc9bd92c7230d5a7007fe7577af277fd615d608cd46994430): complete subsection reference.

<a id="canonical-13cb3e1c6ef04fbd95dcf60261929ec19ed519a5b71f0d36f7bd8322b2415e13"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 62f43d122f98 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-52b5928aee3004d37f768900daddff3b151b3a83b0984d6f613dfeb57d6455bd)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-e581022db77f8d9b3a2003006ade38e71fbdfa4a7b36ede85c9b18dc97ca8535)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-8b2754e031ca410dc9bd92c7230d5a7007fe7577af277fd615d608cd46994430)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-52b5928aee3004d37f768900daddff3b151b3a83b0984d6f613dfeb57d6455bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ac6c217ab511e6cafcdb923c099fa0a5af2933ead9d158d6b6bfd4b819c7b36"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 8c575c4c31da / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-1f3f9ee632afb7e7c4a1010e94a16f2df1b490573625de7a15c6d56dcaaebb4f"></a>

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

<a id="canonical-0a50883a9f31d09925111685f1355ce4cdfd5103581a35930b81ac77ece7d30c"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 8c575c4c31da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65f62244f4e74330fd29abd0cb8f23c7803837697c02f411b3839d4e03c57b93"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 8c575c4c31da / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-e581022db77f8d9b3a2003006ade38e71fbdfa4a7b36ede85c9b18dc97ca8535"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33fb606d4aab5f98bfe84f336316cb6c06df1a8feee14d40164887725a4eb717"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 69118afa6582 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-41be28b7ddb91f84f10a01c31830143bb4f2172f8434c6c60a071f44086c0e6a"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-25ab77f37e3bd97721422536dd9922aaf4f5fda96fe41d5cb41132a3dfd7e2d9"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 69118afa6582 / 3

<a id="canonical-23d25ea4aa9074f1a579343be8828c473e7cf97e115543e6cc195a6cd56e743c"></a>

<a id="canonical-cfbd9736932c50ff16b5833c08d90f2ecdd3073b50832bedfab7cc21aa563d32"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 69118afa6582 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-72b3b64eb62b8c68b88a574f371bc95cbd8bbcdebf3e41d32649700e609f247e): complete subsection reference.

<a id="canonical-5c96a6a6130c65d2d26c7ea268003e056b48ae62a765539c255d78fdf8c7385c"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 69118afa6582 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-72b3b64eb62b8c68b88a574f371bc95cbd8bbcdebf3e41d32649700e609f247e)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-72b3b64eb62b8c68b88a574f371bc95cbd8bbcdebf3e41d32649700e609f247e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0171ce76b026c1aa1d04d9606c9dce74cb95a363337ff641483358f12711fcb"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / 0bc0d5c40568 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-e581022db77f8d9b3a2003006ade38e71fbdfa4a7b36ede85c9b18dc97ca8535)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-d84d2cd4362d4caaf385a69885f9b526a2e8572b2b41a96a35c5953f4245965a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-28ecbd9ab143eafd224825cbd992c14067b405e66f308f608dc731f54e534475"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / 0bc0d5c40568 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f46e08943c09058ef5e53cfff87b237855aee35f63c510830e64402296553bb"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / 0bc0d5c40568 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-e581022db77f8d9b3a2003006ade38e71fbdfa4a7b36ede85c9b18dc97ca8535)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8b2754e031ca410dc9bd92c7230d5a7007fe7577af277fd615d608cd46994430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ada9d448d3f7c78711a060913ae7181e3d822c7c78d8a4d6ff2899f811167cea"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / 3b8d59c01990 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-4ed7696880322d396b1e81762fa66d5bfb5fb8090ee8bf7a7d75cdc3b3632f51"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3cd50e5411417786a463e256e8b21413bc86ddcd5043d39418e28a103e3114df"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / 3b8d59c01990 / 3

<a id="canonical-898fe0eb4f11ad86513a676825cd74ebac286836d84152ec3962fb378469feb5"></a>

<a id="canonical-ede32acdf74b15c8cda9fb94b029984a11256031e14704b360416b2efed1c798"></a>

## ipv4 property — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / 3b8d59c01990 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-02d70cabde5bac72b2f9f1fe3c1a670afe8c755a431277956a1be4223ffeaf29"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / 3b8d59c01990 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-299278ea065d5b3db3d839922a4b47db34f8cd9b753d64361f5a360120c8d2e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-692f32d0bdfc9808e2738999d3d5d5ebd2fc1d7b70cd64f6e04d8e39c64e2b6b"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / 67c3ff63dcb8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-13fd89a153eac70038a2957e179f034eb8a622f746422332fceed1b2b98574a4"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b961235cbd7c9e415ffbe54599e48a8f4ed6db6897a5e4cdc92cbb6d4fa8e143"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / 67c3ff63dcb8 / 3

<a id="canonical-7017b3f9ae77b2f2767f6f7c0cd4fc46064b27703e5e7773c72d110ac26ec8d6"></a>

<a id="canonical-f7b1f464c3f366f3d0a1c47d058edbe1bdd49d4f12c683a23684368ff3a02c1a"></a>

## cloudlink_network_name property — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / 67c3ff63dcb8 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-5ef14163e077e35426d7915d5ef2d5ad72ab0d72bfaa8a4c83bd32531fd1ea35"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / 67c3ff63dcb8 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-fa5a94ca7a39b2d36fc2759b911feb90e360b9105d49563df9d195b3c8b998fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c476e2a12ead58a97c544d501eea9d483fc1cff2c4d8c9fa2e8522df2de65f0f"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 4be1513a9b23 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-d4bc921f96af03a8bda0df97efde807cd1c04533bf10318f37e3d2203649d21b"></a>

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

<a id="canonical-f42897aa60b300f55210c50f42cb5fa56925884e1675c0ed0c19fa8047c8eb1b"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 4be1513a9b23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-052f066fdf136b4cd892e7ff0a72b1ab052b938ef2145a04426646cb96bf53a8"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 4be1513a9b23 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-c5430ea326d3a0e63f9d8011d6bc6a8aa46668371d4c21e9817c8d4947740b2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b848d25d62e40d042ed1954a38e8452e750f67f4936b360d432b6de0321cc257"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 3d0c9cd7a631 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az

<a id="canonical-f5bc3c8dca35336d2f168029da0195b493bc4e41d16fd4dd9cc150e1e0701a27"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw1az.

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

<a id="canonical-db2634d6a1eec69c4c97098fe0575d100614f9149b625891ea7e1123644d0561"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 3d0c9cd7a631 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d97ac88c4920726f68f249f09645e823e3519dcc888d9fde84638bcbc62de6c2"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 3d0c9cd7a631 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-aad40f15c0a8bfb1b5fd48281b49ceec4dd12e554d021a431c09c5c570bb97fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2be5141f318e349cf9dccc5dd1f18f1c2fea7892cca501e2216e9f23c848d9de"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / f88fe4204715 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az

<a id="canonical-a500d1a767a49aa6c330566328841ba0b9f174467c236704ad6df81e89ca3743"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw2az.

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

<a id="canonical-34f14f77d0afd11afc3069fd0c49e41c87a076813c05f7f4faec27163e916390"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / f88fe4204715 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17deb82963eb3f89cccef8e8a1f00e56c41c8bc44186adbd19dc394134a390d8"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / f88fe4204715 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-259b6cdc06ea3436858053d9c540e2da17f985c597c8a108a157046f71159958"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8607a08649ee844305ef3844bbb5b7cca2fef803a0fdb9bd690964983fc0e161"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 5eb45e121c4e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf

<a id="canonical-5994c16f1b1f2df1fce7b48aea80d5d43c5143bd31a0375dd165fa90f9b76ed6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku high perf.

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

<a id="canonical-78dfcdf964248c394e956e97da2a6c720949698a611650671ede861ca3f8cc92"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 5eb45e121c4e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32538a9dbb3b9b33a9e17df9141f8b89fc4b15e456e7ad83319494ec01b4e741"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 5eb45e121c4e / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-647fd346ef093b5ddf530586a56202953953f81023603e85f6041ed7b47ab228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38106af50fdcca5c311310910c75097d86e08b89764359aab7ae6e644d7fb6ca"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_standard — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / f0c6598f678a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_standard

<a id="canonical-9b5720f60528f8eb6e7545c04344403f3942f47eb0eeca72e0f1bf52427119a3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku standard.

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

<a id="canonical-3011b456edd8bbda647317db986b2dae7a440fef5c4126071e2f86758cc531fa"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / f0c6598f678a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69b85d6caadce79d4d79c806e3523d59211bf0f4577739b7b28854ed2fa5706a"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / f0c6598f678a / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9931fcbfe69b6ce52d59c36554c1bdd93a56e1b893fa4a7ba5b08b790d37505f"></a>

## ingress_egress_gw_ar.hub.spoke_vnets — ingress_egress_gw_ar.hub.spoke_vnets / df0e380c08d7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- ingress_egress_gw_ar.hub.spoke_vnets

<a id="canonical-47c07db7908af6ab9b51a5c2c74dc6e9979cdfe56c1309f73c93e4b2f53af1e8"></a>

Type: `"list"`. Computed.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-200456e0a143a1c50d236418f08079ceb12522bf09a7cb1c5e47d5d82a9c8431"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets / df0e380c08d7 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-a8ba1366a491e8ad9ae4fedafb979e77b87022ef96137a55f011939b2eb33417): complete subsection reference.

- [labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-85582baa4e1a04b33d2ef3b8feadcfc30381da40cc977e8a3006db25e9852ac9): complete subsection reference.

- [manual](data-sources--azure_vnet_site--reference--group-006.md#canonical-3e3483f8772fc289ac74f5205c83076443261e8becabdc8d365af6ba48b2bb5c): complete subsection reference.

- [vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f): complete subsection reference.

<a id="canonical-645063d1585202e616b41857d5e517bf276edf60ee4b462334ef9f0ace3e2ded"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets / df0e380c08d7 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-a8ba1366a491e8ad9ae4fedafb979e77b87022ef96137a55f011939b2eb33417)
- [ingress_egress_gw_ar.hub.spoke_vnets.labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-85582baa4e1a04b33d2ef3b8feadcfc30381da40cc977e8a3006db25e9852ac9)
- [ingress_egress_gw_ar.hub.spoke_vnets.manual](data-sources--azure_vnet_site--reference--group-006.md#canonical-3e3483f8772fc289ac74f5205c83076443261e8becabdc8d365af6ba48b2bb5c)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a8ba1366a491e8ad9ae4fedafb979e77b87022ef96137a55f011939b2eb33417"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-341a886ceef293e1e8471d123952b5725fa03665f6076292343d396d0922f4c4"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.auto — ingress_egress_gw_ar.hub.spoke_vnets.auto / 869a9508b371 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- ingress_egress_gw_ar.hub.spoke_vnets.auto

<a id="canonical-aaff1def33186de96e90e039c46d9ccc74ce215b0c1385d0829c524550d0d707"></a>

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

<a id="canonical-dc0ccf36cabfc6ad81f6d402d1df1ad5015e9714fea01f6a864828f49f11aa14"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.auto / 869a9508b371 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f16b0766fbaf05dfd2cfb3712d5b104e6f7a7df3126b3291d26b49759f9dfece"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.auto / 869a9508b371 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-85582baa4e1a04b33d2ef3b8feadcfc30381da40cc977e8a3006db25e9852ac9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-454ed3b1ecb4da6447b5bd1c231bb19da444e82da613e3efdc9a1c7d458d557f"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.labels — ingress_egress_gw_ar.hub.spoke_vnets.labels / 462e12678a43 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="canonical-a214d2be2b7bf10223498b6005f6b0dc0e75219058212764dbdf669ef9d5d2a2"></a>

Type: `"single"`. Computed.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-02f8bed488827df97744e8ddd9a87b73a70d70d5de526a0944ee1e164a685f89"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.labels / 462e12678a43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-979279e1ac8be37e9caeb97867cd6f695a950f12f96c2830bf466447b8f97fef"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.labels / 462e12678a43 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-3e3483f8772fc289ac74f5205c83076443261e8becabdc8d365af6ba48b2bb5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae0bb69df7af3bf26faaa280c710ca46e4f74d07ac5f096c7c05beb006ce1d17"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.manual — ingress_egress_gw_ar.hub.spoke_vnets.manual / d8b8d85b5c92 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- ingress_egress_gw_ar.hub.spoke_vnets.manual

<a id="canonical-d3b23d0dc0799d59004c1a08c766ff3a4513afcdc1c54f2f22109aafc06af8a8"></a>

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

<a id="canonical-d36f60a626237ab4752f9be25a1b8d1709d6ced4ff60f53f65f2464c5a83dd1a"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.manual / d8b8d85b5c92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0df691f090fb5247172327be4d6f0e0ac4f4ce74720d8b645cad9752ca6e4eb8"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.manual / d8b8d85b5c92 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a76a8735f0d141547dcd7e3584d9cc8bc59d40d5ec38955ef36ccd9969f3b641"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet — ingress_egress_gw_ar.hub.spoke_vnets.vnet / 62763c2c05f3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet

<a id="canonical-ec92ce10e72dd8921bd94a38af7b17dd16453b14fc841d52c827abd33e16aa93"></a>

Type: `"single"`. Computed.

Resource group and name of existing Azure VNet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

<a id="canonical-b4bfd816b0f93ceafab6d4e1811194382a820ff78ca45df4ebe6a67009f3aab5"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet / 62763c2c05f3 / 3

- [f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-0284c4b889b12b922cfa6ed0b320f66f5cde4dd44fd8d93dbc7b9a6dce5185cb): complete subsection reference.

- [manual_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-ed247a5c4a0d3a22e0e4d6dbb343601a826033fa8629aff1b3d73b623ca11cff): complete subsection reference.

<a id="canonical-fafc4272759764d086fff81c5e7d0a463c77555bc9eddb5716c9dcf81634ba81"></a>

<a id="canonical-fd58092cb3dfdd30aa8d70cc88c1f30792d317e6803fb4d429931e6fd4a0787d"></a>

## resource_group property — ingress_egress_gw_ar.hub.spoke_vnets.vnet / 62763c2c05f3 / 4

Type: `"string"`. Computed.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9cc7a05a76aa3fb7d467781bade71a2d59455e1c056d900863f8d3f3672e2147"></a>

<a id="canonical-77c05d5b70dc92fb1b5064d951584374545146259416c1a2e3ed003a4206281f"></a>

## vnet_name property — ingress_egress_gw_ar.hub.spoke_vnets.vnet / 62763c2c05f3 / 5

Type: `"string"`. Computed.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-24c24ef2418bbbd0f49725b22ca07562d685f4ed0029900c8768fd11fed4847f"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet / 62763c2c05f3 / 6

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-0284c4b889b12b922cfa6ed0b320f66f5cde4dd44fd8d93dbc7b9a6dce5185cb)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-ed247a5c4a0d3a22e0e4d6dbb343601a826033fa8629aff1b3d73b623ca11cff)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-0284c4b889b12b922cfa6ed0b320f66f5cde4dd44fd8d93dbc7b9a6dce5185cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b2e8acc4d50b2e4855551599761303c58dd6e58263984ca81168326c12a22d9"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / deb6025b2222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-0d73c86874a38505b71099f9341b3dd3e44c82a38c4915ee76a58ab76fc98ffb"></a>

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

<a id="canonical-1aac30718019ee655355fd378d5a86607af5f41663a2dbee193ec84986471c7c"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / deb6025b2222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332ac4383e122ed102cb14a35937c4cdcd9e18aeb0cdda29b2b10ad67751bb5"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / deb6025b2222 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ed247a5c4a0d3a22e0e4d6dbb343601a826033fa8629aff1b3d73b623ca11cff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9967cc294a369214df3a965480098ef96b3b7f280ccd9679aa790e0ece4b5367"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 2d4346235f16 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-f6a9686c3660a85c77169546b84b460783dc475a62cb48649a042126ec4c4f3f"></a>

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

<a id="canonical-2b822b8c9fec1ba3a3fce3af4a684adfceb2fd9a9a61bc8d6ddce5de67fba1b5"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 2d4346235f16 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c702aa792c46e10df5714e4a4ae4977ef02370107f6a30c7110f76e12fe5a8c"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 2d4346235f16 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-614f2623089b68c4bb1d3a119611a59a42a92d94174288530516fbe8663c5a8f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-126b95779e8e162cb81bfc1c43b9c4ce5e34fbb9c9b1472e474636626b55d9f7"></a>

## ingress_egress_gw_ar.inside_static_routes — ingress_egress_gw_ar.inside_static_routes / c69a311a887c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.inside_static_routes

<a id="canonical-de6d9b7a29b0bd7b78a3165d5f8c7a09ab67260d167a7ab6077404d212d70636"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f9704f43b4488e61b5e9db24fbe87c1df5a4146bc01b4d968a590afc699ecf82"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes / c69a311a887c / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501): complete subsection reference.

<a id="canonical-0ab39fa04e4f17914ea46ff3d56d32782960df5aae0e929424edc760ae51059a"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes / c69a311a887c / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaa9f6ae16af2ecfcc8ebdd5c75869ed80502b4bd8a0879bf739d5ce9b00f829"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list — ingress_egress_gw_ar.inside_static_routes.static_route_list / ae5802607dfc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- ingress_egress_gw_ar.inside_static_routes.static_route_list

<a id="canonical-a28509feeeeac188be2aa3328f9bf7a21018cfb10fbd71ea0b399fa3b460b674"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-307b047237025ce7b94ffab2ccffb433eddc7e07d1fb952ae8dfada5f588ba06"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list / ae5802607dfc / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b): complete subsection reference.

<a id="canonical-debe3651f631b26c30174b0a99a0f5886c41e8c253399880c582957ff7b895da"></a>

<a id="canonical-85f414afa94ae75113e0ed7ee7f59a409051e21a481b75f41b5e58d652e4415e"></a>

## simple_static_route property — ingress_egress_gw_ar.inside_static_routes.static_route_list / ae5802607dfc / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-089a4bd2c371998ba1d3be8b0ee9d3a43df426ad82dac97c7d1338b0de863ded"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list / ae5802607dfc / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4049771cc4325aa3beafa26ef3b2da5b1d25a2f7bc66c08cf1b2e7b60b372a1b"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 2545ba1bdfff / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-169d9a01712ba702de077ed869d99bf2becdcb835136962de062b1b136c6f1ed"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f8334358779c0a798ce054f8aaf2818839f7bf4b7aa065e810920cd1c6ed39ec"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 2545ba1bdfff / 3

<a id="canonical-ec4a945280f291220a109c7ac5034ec891eb36aac1fd4a9dbc3b7c7a88a8aab7"></a>

<a id="canonical-713652cf13aa3d53bb76cf790a32468fb4d51226a830ea3d75d90b470598ae21"></a>

## attrs property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 2545ba1bdfff / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5cc6fd8430f5f15f527a5b95acc80d01682cfbfe926eacf6e6c31bf60774115): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e): complete subsection reference.

<a id="canonical-9299bff2910d4bc0f490de89db2d9ba1f6a9614494a1cd0e34922f0504ffc45a"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 2545ba1bdfff / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5cc6fd8430f5f15f527a5b95acc80d01682cfbfe926eacf6e6c31bf60774115)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d5cc6fd8430f5f15f527a5b95acc80d01682cfbfe926eacf6e6c31bf60774115"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22b77f13f1eaf2073fea3ff460db977ab70e2a1efffcc6b8ee555e25967e6778"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 70469fd061d0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-738216325238379a507c253059b278fcbe21ff0921f806bc783c103316b4e5a9"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-31bddbe90f5256fd2460aaa924ac65b5ce146c2320a3930a520c0c3ea13ae301"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 70469fd061d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4e1a10ef7b30438a2f7d6461b520b3f105cb4878fc6471617864e3911a4a72b"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 70469fd061d0 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbf9d6e388307e5c111f9809a31d525e47793e4dbb89aa3e4597f8c64a86d481"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 0fa80684267c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-62cbf967861ffb5ce3396c8a19a60e1bc2e68b78d2f60fbe082348c60ec3bd3b"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-75737426f3c34d778c7af08d653250997e1c5e61518a57c853a2a5f5d8f8b91e"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 0fa80684267c / 3

- [interface](data-sources--azure_vnet_site--reference--group-006.md#canonical-c9bf0dd837e0d5657ad3c30a97a2e18f4e435d0d2a3fc8860017b1918ebf30b0): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff): complete subsection reference.

<a id="canonical-469b308520778422587c986a5bfb94ca1b58f30bf38af39696c5d5000ecead08"></a>

<a id="canonical-3fa2886eae81aea5b7ed84d07583335fe7a24e1243de78d1da9946bd5a716ad9"></a>

## type property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 0fa80684267c / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-78d0fa5d6f6b38baed9d2e941f422c3c69ce2c2fe7e8f65573110b02243c7ced"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 0fa80684267c / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-006.md#canonical-c9bf0dd837e0d5657ad3c30a97a2e18f4e435d0d2a3fc8860017b1918ebf30b0)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-c9bf0dd837e0d5657ad3c30a97a2e18f4e435d0d2a3fc8860017b1918ebf30b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca456e5b7692de6dce23906f32ddb65c85a58b35c4ab2d99fc4f87ca891cc2d4"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-462dc121ec314481d0ee17aab52aa2700925ec3a18d71e89fb69c251eb30b0a7"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-bdbe5bb169078c09547cb5a2b859cc1d0f4e55e8fb4de6ec12c519713dbcf3a3"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 3

<a id="canonical-88429411747a1440522b89a4e8513666c0ee32fe16ecdec86804d80c62a0a3bd"></a>

<a id="canonical-aac938235b0d4552208744a49a428960d558079f5ecb78ba888fae88a6dd5655"></a>

## kind property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-5874d32c7bef7938e0bafca9f8d84e0e3af21a6d95b9f4a49a34cf30c056a191"></a>

<a id="canonical-74eda759f51719f8b5ae0cbae6f0dc4075e4f9b31f94e25136c8aafe68924a1d"></a>

## name property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-c86df88a9dbd2f02eee5fe50fbfe76077883a8b63f0b282995616e8f9ed542f5"></a>

<a id="canonical-e8e67a07285c6709705a0fff8f8662294f0d831ae528bf1117353c7c74b5808b"></a>

## namespace property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-75912e643db8fb09c8f17f2d2812822e31d706fd6ce046bc61688f39ddae1d95"></a>

<a id="canonical-755c6dfdb191c743b333f95a0b3cfdb79de94c2c8c09471b21c069f11764465b"></a>

## tenant property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-cec4b057c8a62bbe87d855e8813ec12b208f4a73b4139f891530ae2db42dda0b"></a>

<a id="canonical-a3d56c96078fbaa2a5db9d08de0a428d7003e6ab8ece54880cd45b82af77a8c0"></a>

## uid property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-9c619f64a8991b9c90c3e2207367a6504cde90704dfed5c2e05f96ace5653603"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f29b0145793f / 9

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9b975fd49609eecdbe4a1df914779606ad4a5c8ef32566578ca8496ca7637c4"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fd760c9d1bfd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-af99394e7faafc27b85bb08cbee7bf66189bda321eb83009e96b976a9fd6ac6f"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-0efa39deeb88cfcd9422b5ffd32a12f500ceaa17d12b5cdb494e2f40a5a5aefc"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fd760c9d1bfd / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-43e04c3036f0169ceaf3c03fe0752e9e8f166d73b14dcaeada686a8f3402b6ef): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-03a8a76b9a22d099b6867a9f4950462506e49a0a89c01f09bc0f9996902ad364): complete subsection reference.

<a id="canonical-83c90fb46df39da985a7532a7682cdd27c0443c0e85c457be297ca855d656368"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fd760c9d1bfd / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-43e04c3036f0169ceaf3c03fe0752e9e8f166d73b14dcaeada686a8f3402b6ef)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-03a8a76b9a22d099b6867a9f4950462506e49a0a89c01f09bc0f9996902ad364)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29c77404c0e9251cea372a58576d7494768b91295cd052384576dec2687e41c1"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d2b3895e425d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-ddd4be72bb40f003f241719cc9309f920234340e55ff7de32c03fdcc304fc3d2"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5ee160402693176f137e22db8b9c8fe1b1873f9ee410e5c8b02c74f46ddaab44"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d2b3895e425d / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-50d96774483931c74fe58f6f8bbca556acec749c85fb012a9331a07edaec245b): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-17afbfef1f198d621cda563fd6e135ef85de947e112753f81bd9550294dd8b83): complete subsection reference.

<a id="canonical-d378fd7990183f7cf66e7920c658185249a46b7e801de7edb052726a15e21ce8"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d2b3895e425d / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-50d96774483931c74fe58f6f8bbca556acec749c85fb012a9331a07edaec245b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-17afbfef1f198d621cda563fd6e135ef85de947e112753f81bd9550294dd8b83)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-50d96774483931c74fe58f6f8bbca556acec749c85fb012a9331a07edaec245b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06105894004754a300cb4ab49b1aa35bb6dfd2090d72ce0688779b9ef83974b4"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e9ebd166e855 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-d2c333c558a2eb479349d767fbd2463f29ecc951b835641add46a7cd76b7a931"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b7fe920fac9b1b85eb5cf0567bec186b64c06b847b4c57ea2c74758e9a2f53f2"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e9ebd166e855 / 3

<a id="canonical-47b91ad405e511f1a3477b066478057e12b40612930b5bf2420110a89343dd32"></a>

<a id="canonical-0bc272b03cd5869a6d9f5d2d1464df847c933b62583f7ddc480399deece9dd30"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e9ebd166e855 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3477538bf0e4f3434543a0d0d531b487857337ebdbbf6a0c688d1a25c94d1559"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e9ebd166e855 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-17afbfef1f198d621cda563fd6e135ef85de947e112753f81bd9550294dd8b83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-698095712ea0e31c668a0793307357fa10b8a18605dd64d769935b7687779e17"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b361db8601b6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-34cb782153e104ae2a9d63cee9473a51f07f79aa48c233b3b09065295bddbe67"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-146805bb44f7842a0f6eb89a15bd7dad8c350a4989193035f3a507bc81ee6332"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b361db8601b6 / 3

<a id="canonical-049f4ce4fc76b7e94a836a523addedc8228a80336a20d35947b4a4b52decbb55"></a>

<a id="canonical-9f370a6d3fdf9d4d3214cbf2d65df934c257d47cc6bbcfb27b372f576c9979d9"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b361db8601b6 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-94401054a62c627f6d7a62bd1d92c40ffafaf3923b6a9571e23bcce843199345"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b361db8601b6 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-845ebdfaca5a8882dcfb7b91fc6b5ba3704c5b4f15dcd7b381c99caa19c70f3f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-43e04c3036f0169ceaf3c03fe0752e9e8f166d73b14dcaeada686a8f3402b6ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28a718465dfb9bad6dd7043b14ccaa53ad82103b4a449b35b27ac841a5997720"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 509b5213765b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-7ea2c88a3b1143202de707b632b0a0dcd99e346e88a87185b938770c275abe27"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-68db531eff17ad1bf3d54eef531664a591f8396a11b6fd62534118675c181fbf"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 509b5213765b / 3

<a id="canonical-817abbce9b6e25541b49786fe624c406791754294bf3d1f1d26e01ec7d237118"></a>

<a id="canonical-414abb7b9b6697ca1deef3456d6113231c58d7f05a7e1b8bba29f3897608ae2a"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 509b5213765b / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2bd56ed8ca56fe4e3953b804cb4edeeadf06e85a1cd67735236ac4cc6d22da9a"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 509b5213765b / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-03a8a76b9a22d099b6867a9f4950462506e49a0a89c01f09bc0f9996902ad364"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a1137902bad1365944192e0364e7f4d86fa87009006d01334a5e502885d5f90"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e6d6155fd968 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-3c9e9f2095d33a7de9def368989327d7808339129d972e6bbe86b92376940ee7)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-7716db93203f86a53369105bb362887f47a36363e1d7c404edf825fc13d0e893"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-060a00147a426bce72f3d5cf248265ab81b492b0ff4992eb3337790fee97e0d8"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e6d6155fd968 / 3

<a id="canonical-d7e785fe968446b3ca91833d06e939e73e66eb861112e2819fe030693d258365"></a>

<a id="canonical-fe889001b47cbda51fe6e14cf194945108fac0b954f7c89a4300faae4234203a"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e6d6155fd968 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3b934a5c3b1f4061e072b3c2f56c594ef7a372839171f5ddf910c45db0d3251f"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e6d6155fd968 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-975131928a9cf3abaa26925ec83180f8619ad4d33e709348f18eb202921b4cff)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5896d51675257cdf21d65759185b4c7f0762ac875a1540cf9de0a5e1bac6ef9f"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e075ae8afb1b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-f897ac39a7e041227fd4efbeba7f49c7fdc04956c200e9589f165b5bf6cadfa8"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-e0205d7a5ba5946b7692b69e48f23988b6eebae7ec16c4b7a08d0c8d6290911f"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e075ae8afb1b / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-23dfcc2bec7e0dde98d0bd03c18859a1c1cbf924d9a675707ee417b6e1173d9c): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-782e6a60baa5a4eec46d0e28c4574b27f7bcc90355629d79e9e19786364b60a8): complete subsection reference.

<a id="canonical-5073a8eac296fc1caca1e5fcbe863dae5acbf16adfeb22c9b40daa0fd1dcf879"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e075ae8afb1b / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-23dfcc2bec7e0dde98d0bd03c18859a1c1cbf924d9a675707ee417b6e1173d9c)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-782e6a60baa5a4eec46d0e28c4574b27f7bcc90355629d79e9e19786364b60a8)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-23dfcc2bec7e0dde98d0bd03c18859a1c1cbf924d9a675707ee417b6e1173d9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7facc6b19cbf7258545af99a40b04422df6268c72cf666ba70ca5e0dc6f7d128"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 754b05ce2f8a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-402547f13663617cb33902dc2f513870cb354be8f59a07e919964ef468b26b78"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a5269010322fb727da28412208aafa5d2d30f7c9fe29999e88643dbc3c458fe1"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 754b05ce2f8a / 3

<a id="canonical-5513391607e8f8ea251c05e1de30ded2a9aaf05ad5e1303fd047e92017e8e314"></a>

<a id="canonical-ebe748f22260bda174451a0885c21583b6dd6e0ee948d5ef70bc843b27a724b1"></a>

## plen property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 754b05ce2f8a / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
    }
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

<a id="canonical-72aacb5391e3581e8d9c24a1fb1bb3422119c845b3eed191bcdf80e31d8c132f"></a>

<a id="canonical-cdb68b3d618e406863ae39c37bfea20d830ef4d2d180ef7c7c62ff79de9212ba"></a>

## prefix property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 754b05ce2f8a / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-65f8ecda04e022d0a1ed7ddefc60e0189f3a389e565a5bd56d576fe4ed8cd1a6"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 754b05ce2f8a / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-782e6a60baa5a4eec46d0e28c4574b27f7bcc90355629d79e9e19786364b60a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61e819517701db4629e3c003021f8ed664eb3205046e42eaf6001bd41bd69d1b"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d679c7a73c1f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-220355e757a9c7e9c6021b05ab03d6a2ed4312e1b30744a3a4de7cf9f22b3501)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-cdbca5162fb6706b88feb932d97c892088e4f8ab7465b88bfb2e27dd2a6f494b)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-29eaa6597dc389986c87df38d3ffa0dbd270fff28d7ade0de011e1be700aee96"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-845e658ad19394bdf50e27ded4fdbfcb82328596064020d04eddca70ddb807ad"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d679c7a73c1f / 3

<a id="canonical-e7a00d77e6c857259b26276d5ebe95520ac2e31cbfa3cfd48e6518160559306f"></a>

<a id="canonical-fd53cb4dfbf8f5b6a7931e55034703fdc5448c43399dd25b3f6fdb555d8375c3"></a>

## plen property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d679c7a73c1f / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-58a0246c4af046f1c14209e834e21d04646fe0af7ff59bc6ae4f067b07e884fd"></a>

<a id="canonical-0971aa8bd89d5ae152287118ca4d1a4a952d1186f8008a48b4142a650d2b5ad1"></a>

## prefix property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d679c7a73c1f / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-ec15f74a9dedbfbc364e44a00cc00d221042554af060ea6abf83aee660d19da9"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / d679c7a73c1f / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-aa4ee012108051db66573444e1787d8b76ce01c0e4894f203f31ab87d6d16f6e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-1adc55cf4e964eca81eda4ae9033b5497809c0dd2e417faee5856f72d7912936"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75e774510f31e743ff8de421f0f1e2eb8ea1586b951dcc9718b198cd965afc08"></a>

## ingress_egress_gw_ar.no_dc_cluster_group — ingress_egress_gw_ar.no_dc_cluster_group / 3b64e17e5c92 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_dc_cluster_group

<a id="canonical-ad989afa3d942c21ffa6334b2ac2d7939eb78ba66bed0571f5bdbe973dc949c8"></a>

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

<a id="canonical-25e43b887e8368cad2d485f3a47144c4cfa18c4187049ca2682cb6e6c69b49e1"></a>

## Direct properties — ingress_egress_gw_ar.no_dc_cluster_group / 3b64e17e5c92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54dbf1f8c00d374e4ebbcd4478c72de3148fb4067fe3d33ba2e211a222782fc3"></a>

## Next pages — ingress_egress_gw_ar.no_dc_cluster_group / 3b64e17e5c92 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9104b6a30f4a654bf9bcdc4f16f1561bf3b0e8ee75a843ef3c1de495ac60ea7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c54427aaa34b09e599cfc27fe5178c05de7d54def960d7c5c1f497d77be70ca"></a>

## ingress_egress_gw_ar.no_forward_proxy — ingress_egress_gw_ar.no_forward_proxy / aa7da71f703e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_forward_proxy

<a id="canonical-f1c923e9f085c1da4f7e0afb4bf11693c27c46ef6dc59c823128764f8120e188"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-e731e8716488fff17ef7be4e1e17ccd18b9901b5195dec7f47153ea99d497fbe"></a>

## Direct properties — ingress_egress_gw_ar.no_forward_proxy / aa7da71f703e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa0b5d0a3cf83af10fafffb5fa68c184608beeb0a9ecdea8c891b8553f6c3a2f"></a>

## Next pages — ingress_egress_gw_ar.no_forward_proxy / aa7da71f703e / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ae5945d1bb672f90650e08c6886fd4978545223f3e8e51d824d7579ae7d2f074"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7deaf2f8f9f4ff24221e3bdb01f5b0428e48e11d672bd96cda143bf85f59cf4f"></a>

## ingress_egress_gw_ar.no_global_network — ingress_egress_gw_ar.no_global_network / 63af2994f289 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_global_network

<a id="canonical-50f141db4edb52b799cf9c4823eacf0e308f60368faa94aa311ecdd7ad936fe9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-30b088951d707e110b0376ebe94938f32cb5ed7156d083c06db416ba39f0c1f5"></a>

## Direct properties — ingress_egress_gw_ar.no_global_network / 63af2994f289 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-970fb166ba71e66ae76959a73f2ee36faf88b334e188fe3cff0c09f2be3406b6"></a>

## Next pages — ingress_egress_gw_ar.no_global_network / 63af2994f289 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b1ffdac3082b8b06a008d52c4072ccac2a90903b13767ca88f126b4ca6138998"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5cb6735c55157ae99a28ba6386f754c9424d4220f00a004301cee0978dc4e706"></a>

## ingress_egress_gw_ar.no_inside_static_routes — ingress_egress_gw_ar.no_inside_static_routes / 2f935a333072 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_inside_static_routes

<a id="canonical-7ed02a93e2881c457b9e06360a3a4cfbe5c447f523235b91d6326ca170d1aec5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no inside static routes.

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

<a id="canonical-d534e6534491f1bffde07e87fcf20e1fc5d165aef4321ee41964da5e7471fc2b"></a>

## Direct properties — ingress_egress_gw_ar.no_inside_static_routes / 2f935a333072 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b0f33abb20b4a2a5c0aaae05467fe93a6c6b866a9ffda0a425ee0813e5210fa"></a>

## Next pages — ingress_egress_gw_ar.no_inside_static_routes / 2f935a333072 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-fdda130c92b94013b60116a4bbba0410e68076b4a47c6dc701bcc29bbb1bb8ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-330b5e7a79d6306fea015972445f56cab3c4b48c06b26c2be1f214117f1f7e3d"></a>

## ingress_egress_gw_ar.no_network_policy — ingress_egress_gw_ar.no_network_policy / 58a2b121f1c9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_network_policy

<a id="canonical-6ea0184306068b04594322903cf70012378688d97f2a2b32a9b5641f677cf72d"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-336f5befdad1478ad5b05b3f4e981dff1a42794d8389bd2c8b6b6a1ebf6c0f9c"></a>

## Direct properties — ingress_egress_gw_ar.no_network_policy / 58a2b121f1c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b891f9bc78e4a8f897fec171d771c24be327719e310a36e2bab24c0be2a52577"></a>

## Next pages — ingress_egress_gw_ar.no_network_policy / 58a2b121f1c9 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6a379509da9c1ee5a61e9229fb3c0c0b629543b3a032d49ff588e2d82f923fb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5217f3374552e88ebbf300e10fe72e46bf2ed3b67faa5b613d7f68eb427f53cb"></a>

## ingress_egress_gw_ar.no_outside_static_routes — ingress_egress_gw_ar.no_outside_static_routes / 8a31c7d7137f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.no_outside_static_routes

<a id="canonical-c5c86bcf77f8dc73f2e6cfe64bb6572bede4840749b81c277df16d2d87e4c42e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no outside static routes.

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

<a id="canonical-cb90b38b141dd798a54f2ec3385f8c5b19866b6beac2731185dcc7b7b96d1b10"></a>

## Direct properties — ingress_egress_gw_ar.no_outside_static_routes / 8a31c7d7137f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5d3396e9fac513d29642abcefbb3c088d20c8c476dc41a8692944f3956c2ed8"></a>

## Next pages — ingress_egress_gw_ar.no_outside_static_routes / 8a31c7d7137f / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c881d193f074ff8da2bfdcb04fcc1d9bc3a0ac9c4894bf1d05b908b9e5e6235f"></a>

## ingress_egress_gw_ar.node — ingress_egress_gw_ar.node / 428bcbd35dfa / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.node

<a id="canonical-25c13f12f7f930acb95474df27adf6b0837844c74071a11f5e27119d8875a071"></a>

Type: `"single"`. Computed.

Parameters for creating two interface Node in one AZ.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fe392a94681de0b6ae86e1a50f0333376e6003825793054fb964d69736e70efa"></a>

## Direct properties — ingress_egress_gw_ar.node / 428bcbd35dfa / 3

<a id="canonical-a06c495a3e72e58fc2ee45effc70b35f6edaa1e8ca3c7e56af5e79fe8aeabd36"></a>

<a id="canonical-afcbb2dded939cce1f49e17a9eaf95399e60f12bed82aa31fdec5ab2450d7b57"></a>

## fault_domain property — ingress_egress_gw_ar.node / 428bcbd35dfa / 4

Type: `"number"`. Computed.

Namuber of fault domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f): complete subsection reference.

<a id="canonical-b1d6575166d93f4e55c86fbe1b07bae5946f06e1ca7d53a716fbdf07ac54cfa1"></a>

<a id="canonical-3e6590a967b848949f0815305aff2925c5298366420b9695df33c529804e54b9"></a>

## node_number property — ingress_egress_gw_ar.node / 428bcbd35dfa / 5

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-6eb4bbb8a8881bd84845936e87bdcdfab71c34c1f2381c80c4a05e990a714fd8): complete subsection reference.

<a id="canonical-68c20489c95eed2a89d98d1f6643f3c21c585a22112a79af607f6a635b9721ce"></a>

<a id="canonical-039479b803e266c728affb93efa957e75d24dc5926ea90e5ed976de368c5461c"></a>

## update_domain property — ingress_egress_gw_ar.node / 428bcbd35dfa / 6

Type: `"number"`. Computed.

Namuber of update domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-e2d48700679d4fede3101f8d8091e90e5a58480eccb5f22a9d3d00570f16be75"></a>

## Next pages — ingress_egress_gw_ar.node / 428bcbd35dfa / 7

- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- [ingress_egress_gw_ar.node.outside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-6eb4bbb8a8881bd84845936e87bdcdfab71c34c1f2381c80c4a05e990a714fd8)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12573c7a1123803dbe48e9414ebd4466c869f96e80fbe554b9dfa1cadf0a99b0"></a>

## ingress_egress_gw_ar.node.inside_subnet — ingress_egress_gw_ar.node.inside_subnet / ba1bf079ebd0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- ingress_egress_gw_ar.node.inside_subnet

<a id="canonical-3d6b160e845751007db01e375282cd74347f5922678e8e5c2be428c8e2a25b9f"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-3baf376d9ced76e6bb1af0582263954067982e5a4d8d3d4632bcaec0b901391a"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet / ba1bf079ebd0 / 3

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-cce6b587f21a6707cd9ef0ad85110fc3e4bc6f5892f3b4ef054f79d9572afc60): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-4b93f047b4e39752a556cb5a764f229c61abc6878b9d33e1dde2ba5ee1337c44): complete subsection reference.

<a id="canonical-531e46b1ce99d53ce5760ef545ca6b762364d2db4e9963a65270c63fb9a48665"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet / ba1bf079ebd0 / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-cce6b587f21a6707cd9ef0ad85110fc3e4bc6f5892f3b4ef054f79d9572afc60)
- [ingress_egress_gw_ar.node.inside_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-4b93f047b4e39752a556cb5a764f229c61abc6878b9d33e1dde2ba5ee1337c44)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-cce6b587f21a6707cd9ef0ad85110fc3e4bc6f5892f3b4ef054f79d9572afc60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b6df857892ac6c61236d1e87d6294c6934311e94ecdea1aabf0f44fef8086e7"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet — ingress_egress_gw_ar.node.inside_subnet.subnet / 733c7e7882d6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- ingress_egress_gw_ar.node.inside_subnet.subnet

<a id="canonical-ca54452e3fbbf75f33c338e49c53c40839ce3ad197b641703c648f755a307a33"></a>

Type: `"single"`. Computed.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-50958e8dff8d1bf128d0d118255fbcaa262936f007b309a9d57724981b629fb0"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet / 733c7e7882d6 / 3

<a id="canonical-80b2c35955c7f81e92e8bcd6297d32365ef3c758ff3acb1dfb5770fa19cfd2e7"></a>

<a id="canonical-fe70fa34acaf84d1f2f0301454ed1fe1113af8b1efb788bd1be482462b8b7e11"></a>

## subnet_name property — ingress_egress_gw_ar.node.inside_subnet.subnet / 733c7e7882d6 / 4

Type: `"string"`. Computed.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2fe5d9a08fb7f38f697c4e12629082c2df1a036d5aff7d97a62bf7ce49d35b20"></a>

<a id="canonical-a8abf562ebd9c3a0829238e9de9c78258d95a1c7d240f6490de161514579865d"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.node.inside_subnet.subnet / 733c7e7882d6 / 5

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-8a856bde273a85a6d4fc42f8154e838a3b05e3d73d761b97db6e34cd0c789eb1): complete subsection reference.

<a id="canonical-a466055d2ab0ea8357f7914f332147637f7c1e165cdaf9ee21256753744b4c17"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet / 733c7e7882d6 / 6

- [ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-8a856bde273a85a6d4fc42f8154e838a3b05e3d73d761b97db6e34cd0c789eb1)
- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8a856bde273a85a6d4fc42f8154e838a3b05e3d73d761b97db6e34cd0c789eb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-586ae527add7c9e5cfe2f052bb8fc9d19b8017bedc3b7a21216d8ab65b7fa987"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 12674864979d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- [ingress_egress_gw_ar.node.inside_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-cce6b587f21a6707cd9ef0ad85110fc3e4bc6f5892f3b4ef054f79d9572afc60)
- ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group

<a id="canonical-d1654d53d4d2351d32ee42759ab51141ed284883ce0e975b601af47b893fb451"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-48eb2723377b3a92c0a6c4780f6d0319e4df27137e5eb5d14af03b23de18220e"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 12674864979d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9aaeea527b8f63f216a0ae5aaac27a0682b0a00a2298438cf6f46593a0db03e"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 12674864979d / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-cce6b587f21a6707cd9ef0ad85110fc3e4bc6f5892f3b4ef054f79d9572afc60)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-4b93f047b4e39752a556cb5a764f229c61abc6878b9d33e1dde2ba5ee1337c44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf8b79387fcd5b9d3479654e0a24bc8dcb48b193e70b3ad31b3b9b53105d0795"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet_param — ingress_egress_gw_ar.node.inside_subnet.subnet_param / e8fa52450ae1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- ingress_egress_gw_ar.node.inside_subnet.subnet_param

<a id="canonical-9bdc510ef92e4220301971b255683d473a842402c1c83b474d41945466944813"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42b4f9d183b566481c13a7595d4f135cfac0f41b52ab1f029d013a08fcde4e85"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet_param / e8fa52450ae1 / 3

<a id="canonical-18ef7d940e22d4d3267f29c0766400d5f7abcb823c005207ebe0c8c8296dfd2d"></a>

<a id="canonical-967c00e36eaceb37d866b8985a93d4157c46f9a2d40970cf6a0d971d129d7806"></a>

## ipv4 property — ingress_egress_gw_ar.node.inside_subnet.subnet_param / e8fa52450ae1 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-f49b957886cf47a5a56a8ca0b480463559b2628f8f4d4c052cf50b91e6f38302"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet_param / e8fa52450ae1 / 5

- [ingress_egress_gw_ar.node.inside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fc307eccace2eebc155045859303fbef4a03c28b4e4892eb44186e9f4129095f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6eb4bbb8a8881bd84845936e87bdcdfab71c34c1f2381c80c4a05e990a714fd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0e001946c478b30553a1e4d187013d69b6c2d553ceaa94fdf92a4d4363669e2"></a>

## ingress_egress_gw_ar.node.outside_subnet — ingress_egress_gw_ar.node.outside_subnet / bb710820c267 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- ingress_egress_gw_ar.node.outside_subnet

<a id="canonical-1abb798706dc191aee3c16c77043b6ebaca73257e96f1758a7b37a908775ec48"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-a692ae5189c54407dbc6390a726bf0526175361f35b15ba164a0b21e1d21e1a4"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet / bb710820c267 / 3

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-c673d7e0eddf5d3bd34759dbd162a4df3472bc4b707851df6dc2c5241013ef22): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-007.md#canonical-19869e022df63d74bf5588a1237604083ab102636b007651178e44543affc895): complete subsection reference.

<a id="canonical-241bfb66d05a152c4d544beec0452045847556acb2aa6a801eea907217331157"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet / bb710820c267 / 4

- [ingress_egress_gw_ar.node.outside_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-c673d7e0eddf5d3bd34759dbd162a4df3472bc4b707851df6dc2c5241013ef22)
- [ingress_egress_gw_ar.node.outside_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-007.md#canonical-19869e022df63d74bf5588a1237604083ab102636b007651178e44543affc895)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-c673d7e0eddf5d3bd34759dbd162a4df3472bc4b707851df6dc2c5241013ef22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-928a5fed2cdb123720cab1b5b40135c1a038494f8abfa4e4acae756e494e34cf"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet — ingress_egress_gw_ar.node.outside_subnet.subnet / fb8b7e2b8874 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [ingress_egress_gw_ar.node.outside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-6eb4bbb8a8881bd84845936e87bdcdfab71c34c1f2381c80c4a05e990a714fd8)
- ingress_egress_gw_ar.node.outside_subnet.subnet

<a id="canonical-8b4e6194b44d45070c12bfd673e732f1e3028a2188a43c0207af22e012e8598f"></a>

Type: `"single"`. Computed.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-58159c9593cbd973696c6c54654ec37da107f9fd32be311692051ce59e8787fb"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet.subnet / fb8b7e2b8874 / 3

<a id="canonical-40717a2886d0017f2ec0c47c9535a223dde53b82f916e24cc02987df5d1e1a81"></a>

<a id="canonical-d52770fa4b1a63e91bcf0e747882913fdf5a69857174e070390d157c258b683f"></a>

## subnet_name property — ingress_egress_gw_ar.node.outside_subnet.subnet / fb8b7e2b8874 / 4

Type: `"string"`. Computed.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-07ea6b995a4c28132d4a82863e0ccd738cdf3ba2107341964362420f77db4a6c"></a>

<a id="canonical-8139f4b91a994b904b379189af57f07d589fae214fe50c8f2795ca40d6661abe"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.node.outside_subnet.subnet / fb8b7e2b8874 / 5

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-0fa8f0ae3fafc41f2f09c97d40476503e477bfa6436d43e35713e54332bbe2d4): complete subsection reference.

<a id="canonical-dded8ccb417329edd96a7f9f0aac805ff3ac6e0d85e480475ac8dea440592b5e"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet.subnet / fb8b7e2b8874 / 6

- [ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-0fa8f0ae3fafc41f2f09c97d40476503e477bfa6436d43e35713e54332bbe2d4)
- [ingress_egress_gw_ar.node.outside_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-6eb4bbb8a8881bd84845936e87bdcdfab71c34c1f2381c80c4a05e990a714fd8)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-0fa8f0ae3fafc41f2f09c97d40476503e477bfa6436d43e35713e54332bbe2d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
