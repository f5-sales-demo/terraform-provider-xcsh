---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-d01e6663edef2b30dd58b6b8d1f6edad3dd35c9634cb198bb112d44c03124c05"></a>

## circuit_id property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / 38f538e3653f / 4

Type: `"string"`. Computed.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-10ca3c3b83260544d9761983b332f70f3eac2ad055f99ae5e23d451acd858d50"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / 38f538e3653f / 5

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-e40b9fb4053045f7fe5c589b339514843637634386071c210841581ed8ea4460)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b825a0df8c6bbdb4edd1424d897fba5d8a8e93607b810a8000d12add342691"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 50f8498eb46a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-e40b9fb4053045f7fe5c589b339514843637634386071c210841581ed8ea4460)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-003.md#canonical-fe1aab1fdc6149e277d35b5754a36c9f439d48b5bc2a2e2d9c9138d57f78ca9d)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-1afedfc35bc0a4bd862d048604e349bf1d57d28fcb46710a7b028fb15aedb67a"></a>

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

<a id="canonical-91166fc695b469787344b46bf054957c5c9452e3ed3c526b7494146a1bf9677e"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 50f8498eb46a / 3

- [blindfold_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-22f43ed890024fb66f7efc62b828b2ac6d69c18e4889e5b5deed8ad09316a737): complete subsection reference.

- [clear_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-d3d20ed9eeca09d3f0dada5a4fa8475aca3c1965e81cabaced51d9f14c3cddd2): complete subsection reference.

<a id="canonical-fa6bb0dfa190024e6529e29ebd50b0cac36ea061d72c016072109e026b74b15b"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 50f8498eb46a / 4

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-22f43ed890024fb66f7efc62b828b2ac6d69c18e4889e5b5deed8ad09316a737)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--reference--group-004.md#canonical-d3d20ed9eeca09d3f0dada5a4fa8475aca3c1965e81cabaced51d9f14c3cddd2)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-003.md#canonical-fe1aab1fdc6149e277d35b5754a36c9f439d48b5bc2a2e2d9c9138d57f78ca9d)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-22f43ed890024fb66f7efc62b828b2ac6d69c18e4889e5b5deed8ad09316a737"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64f7a6ae48db846e4e18767ea8a82fd78ec4da4fd0255a6e72f565400d5299bc"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-e40b9fb4053045f7fe5c589b339514843637634386071c210841581ed8ea4460)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-003.md#canonical-fe1aab1fdc6149e277d35b5754a36c9f439d48b5bc2a2e2d9c9138d57f78ca9d)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-2e951bd1af5a5f7ced167d270e125d24abe08e5d34bfb5d94d59bd7ccec4c313"></a>

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

<a id="canonical-7df05e3952533a6132f2538ff9902fe045bbf90a0ee3736f0bcf5cb0710b3b3d"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 3

<a id="canonical-80a9e42c5a136d87aee60d4c5c498e0ae8ad62350d9dad0e265c09a1ef78dc93"></a>

<a id="canonical-01d86c33291f0cd40c573868f8e96416338ef2d2b486777595b1202da844147a"></a>

## decryption_provider property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 4

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

<a id="canonical-16fd8bc651fd2db796c686c9c62e6ae2030db5c9a25a18bfbfc9200e3a0871fd"></a>

<a id="canonical-3daeef173a1905ee3674306ece90915abf3d8b1c6eef58c65b98195fa0678f6b"></a>

## location property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 5

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

<a id="canonical-f11ae3213c25c56c62bc9d3240bfec0d96fa3315f6ddc094538037f870c6458f"></a>

<a id="canonical-9e6b51dc219dfcbb36af001dd409f9c264a43fca3589d7489c077f7b0b023281"></a>

## store_provider property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 6

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

<a id="canonical-5709284db1e0443eb5eebe21f360c5be4dddbf116e2cb736debf523689dcb503"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / cac82361549d / 7

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d3d20ed9eeca09d3f0dada5a4fa8475aca3c1965e81cabaced51d9f14c3cddd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccb40e3fbafd4978e35e7c20dc41c6ff4d219f8a09fff148b84fe8a5f81adf12"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 8a4dbadcbab0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-003.md#canonical-e40b9fb4053045f7fe5c589b339514843637634386071c210841581ed8ea4460)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-003.md#canonical-fe1aab1fdc6149e277d35b5754a36c9f439d48b5bc2a2e2d9c9138d57f78ca9d)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-6619b4d4d37726dc84dece0465e7ea0d1218a7cb55cf76b00c9a057894e79341"></a>

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

<a id="canonical-435fded362ca3212e987ffdf5f3de3a39244f95dc3f52562e52af40f6a76665b"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 8a4dbadcbab0 / 3

<a id="canonical-cbc35dc63b95872bde3ef9b4cb9b99b944a9e4484ebf0cb636c50815eb03fcc8"></a>

<a id="canonical-0235677c7dcb08e5a19c9972c8bb3b8332cd45c35cf0795f398b4d07b565156e"></a>

## provider_ref property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 8a4dbadcbab0 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1e22ea020298dbdef378d7b60ab5aa6cdcf911e4d86a709fc04601e8f9680415"></a>

<a id="canonical-14f63cbc51b2fa9d9ba3cbccec84bcd87d9204e59403ea1e131341e23345b53e"></a>

## url property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 8a4dbadcbab0 / 5

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

<a id="canonical-ba5e9f2014ed938d2dbc04013b84de94352139419481a217f73332c096bce6db"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 8a4dbadcbab0 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-004.md#canonical-948dde4b94321d8ee1a68e9d1457d641fe9c20224890092afa843747089ee14a)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ffc74a657b4e1abcdf11b303152a50b115874bc9988dd611f7b3dc2070a89aa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8ae29014aaf0297850495ff9203f5a0d041e367b5c4183bcbaaa31c3cb62dd1"></a>

## ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / ce582d956a53 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-f001150f24be1154ef0f5fe60f6d4f718c8cd20ff6b47c3a59e1326d51278d6c"></a>

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

<a id="canonical-7024f63f4b9460441887ec66663229e23d2eb47ba19fa51234ebf76508e2fb72"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / ce582d956a53 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34861814efba4c86dcd8e92c324d49696b3f0ddc5abd8f89c61bf7ec0413449a"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / ce582d956a53 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97ca01c6e38319278d71c2e6588bb7622b478fe9384ff1633ee4c75462b3a68b"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / faa00ab7a1c9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet

<a id="canonical-9aebd5b051d55bfc4d40e509aacaec9f474464617007f022ca1c781542ad8b78"></a>

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

<a id="canonical-033dac489dd445e9a0ac3be0bff9e68c78089a441d698401bf5da98727ca58c4"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / faa00ab7a1c9 / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-55417bbd79d995047ea617d85b496e167c53b8b45afc1505f5b8658e4371bc79): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-79385f0566f07b3c5f13e322179538255e44a3b131b767b4424ed1220d849fd2): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-3db3a61c4f8a8ac27fa548528b8aa59d96d68647ea4a2f32006b10cb60b1af06): complete subsection reference.

<a id="canonical-9bcacdadff8f7da7af6d8acd5608af3af303b96da6d20b5b8ed1d1d58f8993c3"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / faa00ab7a1c9 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-55417bbd79d995047ea617d85b496e167c53b8b45afc1505f5b8658e4371bc79)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-79385f0566f07b3c5f13e322179538255e44a3b131b767b4424ed1220d849fd2)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-3db3a61c4f8a8ac27fa548528b8aa59d96d68647ea4a2f32006b10cb60b1af06)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-55417bbd79d995047ea617d85b496e167c53b8b45afc1505f5b8658e4371bc79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f08f8d4fbb5cc437096947d7bb50a1f51984d95156b98208640b466630548e7"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 5794f1e1112a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-fd73ac3541d3f2a14103f9be76cf02ea01b89856ec42be8170d00928db0d559d"></a>

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

<a id="canonical-86badd0e41390611acc67e8f839d6faedcee42b2e137ef7cec02f29d346715fe"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 5794f1e1112a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-977e291b08b8401d6b6ed99841f2b23c324e1115dca9b13d0e2a3e8f6c4b40f5"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 5794f1e1112a / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-79385f0566f07b3c5f13e322179538255e44a3b131b767b4424ed1220d849fd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-368dfdfea0b0bdef32fdb161737bf7d2eb3b247b219b7e7b7369d66b5591b284"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 372cdd6a2ecc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-954a4459a1a1ff7f6fa5b4458959112954781681e45f0ef754d177fe350ba5f1"></a>

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

<a id="canonical-1e65b178b985638a0a29680a072bb7b8298570d99f0d9b00762f67e75a1e64e9"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 372cdd6a2ecc / 3

<a id="canonical-3214507b86366bfab18d7aea60560ede91ca856af7c3838453bdf3a33c15f4a2"></a>

<a id="canonical-e3ae70d16f699e5edb405e052631ed5b06d891e01707e87b3b8fe878c85b4fad"></a>

## subnet_resource_grp property — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 372cdd6a2ecc / 4

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-cc6a87884c7d21e62190cb756ddae65ff13bcb3b05ce7d61046c0e631dfb6215): complete subsection reference.

<a id="canonical-304ca1ae6c2fa50c26b149a6f96c98233e511d08266bdb1d0ccf5843d10b60f5"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 372cdd6a2ecc / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-cc6a87884c7d21e62190cb756ddae65ff13bcb3b05ce7d61046c0e631dfb6215)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-cc6a87884c7d21e62190cb756ddae65ff13bcb3b05ce7d61046c0e631dfb6215"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b361c7e16d7598234d66b66f384a62cb3038cd66382ca97261e41916255cc32"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / 18ddb2aa4b6c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-79385f0566f07b3c5f13e322179538255e44a3b131b767b4424ed1220d849fd2)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-ee50eed551f77930902779bca28a8e0cbc39148982fc900ec03fb6eb9ce1beb7"></a>

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

<a id="canonical-ea1b402393fbf92b59cb97ea591e0d75d96c6c729d3f8bb63648e69dfe5781e0"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / 18ddb2aa4b6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-608d4a47c2a73c364af89370c6a5190c7460e57ff968856c09f53a85a6682a22"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / 18ddb2aa4b6c / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-79385f0566f07b3c5f13e322179538255e44a3b131b767b4424ed1220d849fd2)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-3db3a61c4f8a8ac27fa548528b8aa59d96d68647ea4a2f32006b10cb60b1af06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df52f3a271269c484219d137e388f874cece1562047800c534201ca3c7cdf70e"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 433a866fcf72 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-dc2ac841b85990a82d872e86adca998465f1219dae4659d81e4cdac95ec8366a"></a>

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

<a id="canonical-43dab02107132fe72e35235026c679bb4a77cd5a562eb8c514c7d36ae0234b48"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 433a866fcf72 / 3

<a id="canonical-0359e6a5d84bcbe16694592d3e42987a3d6d35c5484af692fc684de968070037"></a>

<a id="canonical-5807aeac8cb7e8067e8203c373fc0cbc9f091b850566772691c8e633f5db75ca"></a>

## ipv4 property — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 433a866fcf72 / 4

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

<a id="canonical-9a4d2cad2cacf0fd40d241d172a75fd4c8c5b1889765eca62b1fe8d32e030789"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 433a866fcf72 / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-0c6eead6a7b8f23c709b8c1a6f166cab4060e58a1ee61822da8aa2e197339f9f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6369e325bf42a23f51024db8e5177e921a0997d994d0a0d33b0ba53b67503a0c"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / 47d354e2060d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet

<a id="canonical-16705c49166a8cb0d0f31f2e365a24af96a07f76fd55490ce5e02a34d988bfee"></a>

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

<a id="canonical-d93922cf10c31c72f02eda0562e4959091e4d2c5b1cf75c7c17d0e0ca9a318e3"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / 47d354e2060d / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-49c007f18619bb606284c4ab05e326c6afe26bcb0708a3b7fb071b941b150191): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1a119331e3f81272b601b256f911553c1655bbff80ce366e28acf4291f32e38b): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-6824d260428214200461ecd72ffcfb5d48a72a09d3b1619dbb6c7b5d33d15be8): complete subsection reference.

<a id="canonical-70bc651d54591647fe7c138fc9cc76fb9b9053e4a7c9899dadf9c82be22bfa63"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / 47d354e2060d / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-49c007f18619bb606284c4ab05e326c6afe26bcb0708a3b7fb071b941b150191)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1a119331e3f81272b601b256f911553c1655bbff80ce366e28acf4291f32e38b)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-004.md#canonical-6824d260428214200461ecd72ffcfb5d48a72a09d3b1619dbb6c7b5d33d15be8)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-49c007f18619bb606284c4ab05e326c6afe26bcb0708a3b7fb071b941b150191"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ea950b6da3a13252c175082f31042ec914861daebf297b3f727ea1945585607"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / 63a89bc3a39b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-12aa37f33f00f578078a2b2de40891b75ebd50c6a46cfdc5808d2dfe7d88a25c"></a>

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

<a id="canonical-fa79cc887f978867f75896e80b44a8ddff22980d99e6931328d4573508815f94"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / 63a89bc3a39b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c566deb0ad4c15c3bf25d2dfb2645ec0202db8fe271dc1f06523bb0ede914a88"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / 63a89bc3a39b / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-1a119331e3f81272b601b256f911553c1655bbff80ce366e28acf4291f32e38b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c4a8f3906342dc273d7efbda9bcf4d1bbf15c35fe07a1b0e03604350adaf293"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / c1b494d63190 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-26a23f73c4090bda3b1074ee725f44da9637f72fa7a604cab55fa1c5aa038249"></a>

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

<a id="canonical-3264d900446421b59d97bd22b1eb4310970cf3abccc131f34b43892211664024"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / c1b494d63190 / 3

<a id="canonical-81e402749420e2cb327328357fbcc5145d1f9df9a9828ff2d78a014165117eb1"></a>

<a id="canonical-c6fcaa6f03a4dc27d80a8b9eacf72dde47b6fe3f115b25cfba3856fad3da84d4"></a>

## subnet_resource_grp property — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / c1b494d63190 / 4

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

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-7b7d000f2714aec391e327494fdc51a184dce75d3f8a573c67a6715231de5508): complete subsection reference.

<a id="canonical-8e5102c267a0f16d5a8d64425c047c07acf755e4171ebb505dc7c777efdff63b"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / c1b494d63190 / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-004.md#canonical-7b7d000f2714aec391e327494fdc51a184dce75d3f8a573c67a6715231de5508)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-7b7d000f2714aec391e327494fdc51a184dce75d3f8a573c67a6715231de5508"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a36bbe8f9c6a77f5731ca2c849dd2e05dd721d0f4911fa7dc9d59b13ba66fb"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / e84343dd5757 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1a119331e3f81272b601b256f911553c1655bbff80ce366e28acf4291f32e38b)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-9527d9e1f9cc276246af45359d77b56156888aa8e500bc908d6fe0218c3eb3ae"></a>

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

<a id="canonical-6e0574a30fd4115cb83dfad80ada5535bad6fb5889fad075cb16ba6df57fccde"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / e84343dd5757 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad4c0caf4050d1c8fcf2e589f6dfe757f9a5da10ac6ded21709edc2e0abfd07e"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / e84343dd5757 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-1a119331e3f81272b601b256f911553c1655bbff80ce366e28acf4291f32e38b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6824d260428214200461ecd72ffcfb5d48a72a09d3b1619dbb6c7b5d33d15be8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3287a5b3e44f119a001f9f8b7ab10cc289c536e74f5e6bd50e6f51e3505abe94"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 584939fe363e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-14f7035318a6ce72d30e889599545e4a9257d97d37467a8d05df021e4a8b702f"></a>

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

<a id="canonical-673e410b5c3122a20dc710cdff7b3612206e8d1ad99a95336947970b0c65562f"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 584939fe363e / 3

<a id="canonical-cfbd73ff4894aaf6915f3c6831e3ae787727d6938c281a57f62016911e39b689"></a>

<a id="canonical-a1b84107b9b3ccf3bd8dcb6b4e7f9a646af4aeaf0a4f14375494a1706e32634f"></a>

## ipv4 property — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 584939fe363e / 4

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

<a id="canonical-0ba6c7169b356f32daf68efbfbd22de4b555e3b42f4bf69ee02d1656ae7a1f0e"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 584939fe363e / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-788554d43e98deb8b38cf87872fbe1405c11713fd136c77f4752121bca6a7414)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9376e85c096dd463137445027ad6a386a6a3d9c940fe5f79333c6fbe54b15597"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e6d7b34bfedac39d77ef559ab7fe1c3dd8c33461c1b66f78a159eec750d2950"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / 9369c7febda9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-a23a592ad0bd024f8e0aaec815a7bd777a207085355139e55b230bb786112412"></a>

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

<a id="canonical-6fbeb8ecf28ab068c162949a62b9ee30a5fb901b7415b9944e3a2a40604c03ef"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / 9369c7febda9 / 3

<a id="canonical-542d72ba9b159b9594af2f05557605432981a3c22b4a0e3ecde09f9ee48b26ce"></a>

<a id="canonical-96a477412737517791651eca728567ec0c303b8fbb4212797013d80f14df10cd"></a>

## cloudlink_network_name property — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / 9369c7febda9 / 4

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

<a id="canonical-9fda4feb83a71e89eac434a4127ed3a231e7166c32d306505848866f07a611a4"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / 9369c7febda9 / 5

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9e4b51e70c0672d172dd59ae1f3cbc668f6af8339053fe5f182c40262ec1be1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-052db72f3e08bb97c3147633cdd95146aae1ab287e05738251372186c7c997ca"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / e721790417ca / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-a88b1d4a63619753f0d26c78aa43b1ee08e243cb2719293fd4941c415a081de1"></a>

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

<a id="canonical-9ca7d4ec76163c46b77a40b4602a550078ae53c1fcea53ba31b7ca1955e56490"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / e721790417ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cf28f547279b155b8dd5a304d06a3ed79b8fbcec3063a820de9bea851adbe6b"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / e721790417ca / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-dcf40d2f20d797c2830f58a00cfb62d77726b8893e1a4096cf17af9c6d516d03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e4c5be3b83f9c4be8937797409a1de4df877142e305978e1ce9ce2017a91329"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw1az — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / 78aaec1fcd46 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw1az

<a id="canonical-f5c296b694bd42a7ed2e1bf4bd36c9791e47183611900322edea12849b9374d4"></a>

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

<a id="canonical-6ce146c56adf40cede8ecaa99566d9bc2aa76fe2cb14d4bf1d316e493f53dba6"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / 78aaec1fcd46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d13e280f65ac089fdc66a12795afb56c65e5886e99fdfe3b889afd06e6fd4e1b"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / 78aaec1fcd46 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-03d5cebcc9552675a93221af2c1bef69e50196763233e22ddd68c202db29eac2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81bd707204da3e567117eb5c227cc57adf999f1976bc3694b7c492494dade18c"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw2az — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / c8324505098b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw2az

<a id="canonical-04c5ecda725073aa37955ccbb3770e09c6cecd41e1dbe6a133caeb7867417fb0"></a>

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

<a id="canonical-a04bebfc0c95edaa93efa4b31fb64c633ac4fae30ea1f41b08f85af501294194"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / c8324505098b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fb7b655084dd6cf4eaea99c3dbe8b2b4895ee1cad163b8bfc16d66210fc8db5"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / c8324505098b / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-06967f271e89bddf54d30f296e0ccbbcdb7b260fd37bbc6fc3f646f94aef3015"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-defb9c1e3749075cbcddd4b07507c27b6c28633f0ba5c394b268681c5ff4b020"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_high_perf — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / 9062a05cd4a8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.sku_high_perf

<a id="canonical-1abf2df1a156201bee54d8adb988e0ecefe113f983306a4749cb7192d0be98d6"></a>

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

<a id="canonical-aacb01eac0f0ead44d223ada04b2610203f2ff20340a37828fbce2b2307b8165"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / 9062a05cd4a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5e15551c71c310a6a90b9a2c5286479ff41107d0e9279ee7f64b8e0b4807a6a"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / 9062a05cd4a8 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-42f97c1b2369c3b92a25330536342e093a4719028fd4c7306c4089b118568348"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0334e0b3dfc4ee15fcd17510d610ff5c9c8452e19670a19e2a3520c4e7c17c40"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_standard — ingress_egress_gw.hub.express_route_enabled.sku_standard / 75964aeb6067 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- ingress_egress_gw.hub.express_route_enabled.sku_standard

<a id="canonical-85505d81165d2edb6bebc6259e37689cb1d23da24393b19f7c032b5a85f87816"></a>

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

<a id="canonical-75202680598e6b9a529281e0e3b8a2b4e5ead1ebb2240d1293b92df8656e1dd5"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_standard / 75964aeb6067 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-220283540e93eb469df798e2d57375519b019d73f088626dce35819d99baa78c"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_standard / 75964aeb6067 / 4

- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-003.md#canonical-d7f7be373cbfb2f815e63d4360a11d4242d7ee813a6f1e8a20263ae779e39e0c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddbf2abdf8deb30dec7961f2809132f1de02ba4fc553cc8222740a77b8230c6d"></a>

## ingress_egress_gw.hub.spoke_vnets — ingress_egress_gw.hub.spoke_vnets / b43c0485fe5e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- ingress_egress_gw.hub.spoke_vnets

<a id="canonical-4041cbfd662655ac20a6c33e1a3ad0ed41e07c08ad8cc093b3e6e51b9a7cd56f"></a>

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

<a id="canonical-7dfc2040ab218e051ea4c4e30015b1fc17a611e8570da878f78575d1cf7ea340"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets / b43c0485fe5e / 3

- [auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-045292edb6b381086c7302f7d028be672fbad79670138a8388970528ec82fef0): complete subsection reference.

- [labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-f274efe743ab8d927622a20c6f73dfad2eaa11769b83d5cb44b57050640e8b36): complete subsection reference.

- [manual](data-sources--azure_vnet_site--reference--group-004.md#canonical-eb638b52b7032a3f557c996e68e340795a2aaebdf437f5739417c9ab23efd609): complete subsection reference.

- [vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989): complete subsection reference.

<a id="canonical-4a9810eeb2c77bcc6542951605761a538650aaf908aa66fadba32538a0cb4871"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets / b43c0485fe5e / 4

- [ingress_egress_gw.hub.spoke_vnets.auto](data-sources--azure_vnet_site--reference--group-004.md#canonical-045292edb6b381086c7302f7d028be672fbad79670138a8388970528ec82fef0)
- [ingress_egress_gw.hub.spoke_vnets.labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-f274efe743ab8d927622a20c6f73dfad2eaa11769b83d5cb44b57050640e8b36)
- [ingress_egress_gw.hub.spoke_vnets.manual](data-sources--azure_vnet_site--reference--group-004.md#canonical-eb638b52b7032a3f557c996e68e340795a2aaebdf437f5739417c9ab23efd609)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-045292edb6b381086c7302f7d028be672fbad79670138a8388970528ec82fef0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67d9d3c3b12b379c89f6e6a5feb9dee2b7cb242a3b0d055b5125369579973d63"></a>

## ingress_egress_gw.hub.spoke_vnets.auto — ingress_egress_gw.hub.spoke_vnets.auto / a14d81020f50 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- ingress_egress_gw.hub.spoke_vnets.auto

<a id="canonical-831833d391df06384eb796720ceefb1c15b7bfa1a3b356291e20d07799404b05"></a>

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

<a id="canonical-7909ef23b482ab7a8d47b88d2d1f1e9feac91aebe6e3407c8b1eba441943fd5b"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.auto / a14d81020f50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f10d98fff09f67c5367a8a059a35c9bcbadd63189ec47294024dff1c1306cc8b"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.auto / a14d81020f50 / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f274efe743ab8d927622a20c6f73dfad2eaa11769b83d5cb44b57050640e8b36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5527e3bdd3d64cfc3e1b0d81f3262f5dc3c2c7a987eb405b0d08dba95b745958"></a>

## ingress_egress_gw.hub.spoke_vnets.labels — ingress_egress_gw.hub.spoke_vnets.labels / be3608e58c1b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- ingress_egress_gw.hub.spoke_vnets.labels

<a id="canonical-0ebffcf435902aae46383a221d3c3af26738ffcae5c60861ac040bbe77ae3f2e"></a>

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

<a id="canonical-9551e21af94cf7d088e4d32c656a57cb71fc67bfe9bb565b32356cbade037d39"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.labels / be3608e58c1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-895300a2532c0d491051184748392f17888342ba07a86c4397679ce186b1fb95"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.labels / be3608e58c1b / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-eb638b52b7032a3f557c996e68e340795a2aaebdf437f5739417c9ab23efd609"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d148095928977944b76bc27ef735d6a790998149dd4f6d0f5802fb2f19915d"></a>

## ingress_egress_gw.hub.spoke_vnets.manual — ingress_egress_gw.hub.spoke_vnets.manual / 39207412f27d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- ingress_egress_gw.hub.spoke_vnets.manual

<a id="canonical-8ddc0c5087e4667c0d27f22fc4540721851780b7b8970219c171d6cdb42c57a1"></a>

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

<a id="canonical-f21cad832c0cbabfc670e1adccb1cce17f10a7563985e1fff69344fc603f4bc6"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.manual / 39207412f27d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7423bea2c2cedba433437389566f6920ca465d85a2a1dbff33ab321ad2ad185"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.manual / 39207412f27d / 4

- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9217f529eebc79eaca3272733b77f8a2296bbda30341440cefa9f0b3f6ba13a"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet — ingress_egress_gw.hub.spoke_vnets.vnet / a56f43191492 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- ingress_egress_gw.hub.spoke_vnets.vnet

<a id="canonical-00f705ce5e145b7b34a7c846b75e52f5b73081dd8f8b24de748e8b819dbd1e16"></a>

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

<a id="canonical-8b50b2e6dd5602f2e598e77e26533872afcac8457a26f2dbac63309bedc1a079"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet / a56f43191492 / 3

- [f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-104b73bbb31e10143dc7547d635fbb8a30afd8800a751c4338ff31aa085ff9da): complete subsection reference.

- [manual_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-dde5227ef689a3e1d34f470841d98f01ea5947a075f6791467bdc5baa9bd2848): complete subsection reference.

<a id="canonical-ad04f6eabe6c283f969c00c737662fc6a1ed0404db2447b2b0c14a274c386e06"></a>

<a id="canonical-5da78561cafc6373b04b329bc384441a27bd24f3e6aacb8bfe92d7b9f5253420"></a>

## resource_group property — ingress_egress_gw.hub.spoke_vnets.vnet / a56f43191492 / 4

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

<a id="canonical-55b5ba97e586e5970fc627275b3875535e8358d063e91ed13009372c1b04340e"></a>

<a id="canonical-b59adf17bde63590a75bc85b482ba05a50ab2304606a26f6996a11dbe5e69f1f"></a>

## vnet_name property — ingress_egress_gw.hub.spoke_vnets.vnet / a56f43191492 / 5

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

<a id="canonical-dd5fd6ad3dc38c4a17404b69bcf8309d0a962ce07520bb12dbb7cbd6aade1606"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet / a56f43191492 / 6

- [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-104b73bbb31e10143dc7547d635fbb8a30afd8800a751c4338ff31aa085ff9da)
- [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--reference--group-004.md#canonical-dde5227ef689a3e1d34f470841d98f01ea5947a075f6791467bdc5baa9bd2848)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-104b73bbb31e10143dc7547d635fbb8a30afd8800a751c4338ff31aa085ff9da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ecd903bf8f21b4d84551329009c13324bdd976877e1a9b8c061df646c8b4bf8"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 76a46eb37e9d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989)
- ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-f4bce8c769d4c29280572eceef34edebca7728c8c0cf93cdf5247f3601687d4c"></a>

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

<a id="canonical-afd69b465e35039566d7dd828350aaae646346de8848392bf7fa433dc8f897f3"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 76a46eb37e9d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2b0258679e1bae8ba307b0242e062dae19c768bb4c291f754bee559a6a59bc5"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 76a46eb37e9d / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-dde5227ef689a3e1d34f470841d98f01ea5947a075f6791467bdc5baa9bd2848"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24f7fef82ef97708f678a512eb5eaf6eed4c659a29000077ae63f2326a7448df"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 11b97eeef9f7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--reference--group-003.md#canonical-3001287fdc22d2cdc4540318cebaecd1f647edfa597b42f34a0a67e978ba3965)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-85dd5e3f83185de6682264651764e950e8f67740b4ced6b2b7b093433d95fd50)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989)
- ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-32a4f9c2fcded3f6169e488f0fc49ed4d0fa75a8980dfaa5fbaaeed6de556d9d"></a>

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

<a id="canonical-39af6ed8872c35cf23df637942cd1006169eb08a390f331fe29f8f3d3f898da7"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 11b97eeef9f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc5df18888abdb699430e42649fa02549da2c73d71b175b2fafa5bd148be3dca"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 11b97eeef9f7 / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-004.md#canonical-902f2ce1090330ed58a980b71268e19129652cacd0f9dce8ae38358fc098e989)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6048c55c2c5175b6e9c89779be5bf9a4c4133279668f4b9579d96104b61ac86"></a>

## ingress_egress_gw.inside_static_routes — ingress_egress_gw.inside_static_routes / 3a3e64756521 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.inside_static_routes

<a id="canonical-d021145e2b9ca09065c23b956825c980bc13681d6b20549bc0ed91add5aa0c7e"></a>

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

<a id="canonical-14e81433b3adedd361c100fa610911cd5a8fb811087ddfc60352af3338578d86"></a>

## Direct properties — ingress_egress_gw.inside_static_routes / 3a3e64756521 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf): complete subsection reference.

<a id="canonical-940ea2ff101881bdcff8fb7f50c3d01e14898037260fed345c0e6b8169aea0ce"></a>

## Next pages — ingress_egress_gw.inside_static_routes / 3a3e64756521 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd0ba70ea3b7d8ac0d2528add616eff3c59014394d0922ae5ca02397ea3d0898"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — ingress_egress_gw.inside_static_routes.static_route_list / 0f7a15cc2828 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-ee07506c4d0f45c266f685784581f440c979bca32ecbab6b8744114f7c703a41"></a>

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

<a id="canonical-d6e15afb82338ee0e1357d0fb19a8011cdffe9b47b71d1e0fd76d0f481bee43a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list / 0f7a15cc2828 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0): complete subsection reference.

<a id="canonical-b8df698808dd02e0512ad901aa3e4be4da37ff5e1b40f7d24c32a64a2585ce6a"></a>

<a id="canonical-1c59ab625268fc758bb32e499c5f27a917fff100366ea8c02487f69df7b14863"></a>

## simple_static_route property — ingress_egress_gw.inside_static_routes.static_route_list / 0f7a15cc2828 / 4

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

<a id="canonical-d7c0c6cfa4dde8b1922008a6615de69dc69f85b1a0d5f52b908aa23fc7d3dba0"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list / 0f7a15cc2828 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf18f94ff2706fa8705dd6cc9bb680ba4bf89fc441918390aa0c597a962dd019"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / d6a20ed66625 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-23910b4c34c94e4526f39ade05700fdd19d4e1e2781e24fab477fadd9202cb77"></a>

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

<a id="canonical-c59d97c621b8c348d2ff5505cf44d04fe85ec610464ae11669216d0ba1069115"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / d6a20ed66625 / 3

<a id="canonical-d391e5f87096415f86a0307f2fdf63573e3f2d24911fa503fda7219343d8acde"></a>

<a id="canonical-3b846159aaa9fda0d3bf2bce627ce1bfeed15e5a3660713464f97af11a67cab0"></a>

## attrs property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / d6a20ed66625 / 4

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

- [labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-b78ac2d6085de0fcf7fac63b7ceba092bc2f8dd76c3471befa754f088dac78c9): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38): complete subsection reference.

<a id="canonical-0727f95f9143c5499b58d664aaece5ef843a96ab1b89b7de7ce8aeac5a1c8211"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / d6a20ed66625 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-b78ac2d6085de0fcf7fac63b7ceba092bc2f8dd76c3471befa754f088dac78c9)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b78ac2d6085de0fcf7fac63b7ceba092bc2f8dd76c3471befa754f088dac78c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b43c4fba66012c8fc0fce9b0caa672af864807f125b2a93f52bbed470e0fd410"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 45ad1162823e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3188f6f889a9ae1ad75485b63139d270eac9ee4b1692192c18d218d014365f4f"></a>

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

<a id="canonical-17e357118527cc657b8099e436e100c6c6de5001a67d582550ed361c36c8d8e1"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 45ad1162823e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4aba44e7ff3a7da7fc006447ae55d30b37cef1088293a843270ac0ccb333c8ef"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 45ad1162823e / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f868b3e76d82f2367289a609fcae6cb924bee428c1b1c1eeeba004d545ec9d78"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 8221cdc6dc68 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-49cbc72f8c14b7b7880c21f910aa0e4dd7fee7af3cbb044ec6dfcf3de069b28c"></a>

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

<a id="canonical-2d3ccd67cfcb1df22501d676f96737209faa90abca7d94d5a2ca60591f6f1d45"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 8221cdc6dc68 / 3

- [interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-f2ac7cd3e248ba5284945e275c941957340b1e316d71ebb3cc60886151ec8194): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea): complete subsection reference.

<a id="canonical-3b04c861724f4d3edfae5a8ab0ce1427c403f32cd533f7e128cd256dd3625f1a"></a>

<a id="canonical-19d8dc50e0ca9f8dc84dfb6cfbc4b8a9053b9821edddd5f3b816f2ea67a82e54"></a>

## type property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 8221cdc6dc68 / 4

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

<a id="canonical-1ee93d78b780bf918450bc20948018755cfe72ba45f50abb88a721b95eb35e50"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 8221cdc6dc68 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-f2ac7cd3e248ba5284945e275c941957340b1e316d71ebb3cc60886151ec8194)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f2ac7cd3e248ba5284945e275c941957340b1e316d71ebb3cc60886151ec8194"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eff7d05359482faaefc86827f75a018bde90c8dbc87c3548fc79b84a2c977ec"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-c2f5af347ae84b9be819c3941ab87bafeeb89d78ab3553a9aa75b1369c7e622d"></a>

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

<a id="canonical-907236b7a012f5546559467cf0615122af41946e606f01832c73d56f7ad03edc"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 3

<a id="canonical-4cdace68e9e24af792e2e3908c2de71fd9da1cfe4c8d3206170d9231af381391"></a>

<a id="canonical-61f46dbfd94c378f59727b21bd977327b6f750c18eff758954145843ea2b9bdd"></a>

## kind property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 4

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

<a id="canonical-7e6743e97849f720dbef5055dce8c633deffd90a9992fe5e57d02c7960dd2806"></a>

<a id="canonical-718a2415e6e2a724659378fecfa6b80052818db114e449fd9d480dc6d8a01677"></a>

## name property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 5

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

<a id="canonical-eb116a100cb27fc8b7f4f74ec7dde6d92b6d3a5ae52b3936a129a63e926dba14"></a>

<a id="canonical-24728906792f221f218358868dfab84fe94e70e4bc45b113ffef5728f86d3662"></a>

## namespace property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 6

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

<a id="canonical-434e16b42bf074f205dc86234eb2b7779506eecbaea9c893dd509c5075ce782e"></a>

<a id="canonical-4e2a5ecb4517b8205966758778f6df1c3d9597847d514b52766040bfc68069fd"></a>

## tenant property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 7

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

<a id="canonical-221a9b0e8820cb72d224e035a42d59f8fd8e1098d31eba3492981dd0f7b008c1"></a>

<a id="canonical-f18e9d32febdf4d42269277a2a7e10a491e3c94154f5930072c68148c52805ab"></a>

## uid property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 8

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

<a id="canonical-8a4d0efd3f3b80cba4162ce1739f2dd3fa4dc76024173990f5667b2d40e7ecea"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 37e0c79dd6b1 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32d205b94e272e23ae8b562548deab8d2138fdd377cdfad8f03d5faa0e6384a2"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / dc52b9cfa13c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-c627b0b291f911d32d7b8566a5764713b7e9181eea82b5dbcaee3735409d7865"></a>

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

<a id="canonical-8fd8ae9c06e3564e546eed2fc2150ad63365a616d4762c9ffa51023ec673e49c"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / dc52b9cfa13c / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-484da7eed6ddfccfaa5143bd5304021079285094621408cbd1a5db4d53c879b0): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-d6870e5c1478a282607e98cad79c62fae257da60a794d8d05196000905e1abd8): complete subsection reference.

<a id="canonical-e6b3dafee908d4beb3cc8b86a2357f8ebd5bd1d2c3d2fd7257b09516dae8ceb5"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / dc52b9cfa13c / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-484da7eed6ddfccfaa5143bd5304021079285094621408cbd1a5db4d53c879b0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-d6870e5c1478a282607e98cad79c62fae257da60a794d8d05196000905e1abd8)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79499a2f00507bae4e6d8e7ff27b024172b606c95849ee97ee81349a58611707"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / b561d1b595df / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-38787ea87eb4ccf01fac42a7e64f6239f0932395966f22341480491c45f90256"></a>

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

<a id="canonical-c3631413891f1e1924e849deba02ea3fb4df4e595ee8076b15b2ccb880164a57"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / b561d1b595df / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-7bc14a16ed2bf7ef5eea6b5545821337ae33698dd9992bb7bb28e11da30b4ef9): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-4ff5ba9bf2195abe87c9b933a3eb9058f6a7cf776e97fb6b8583816e48918835): complete subsection reference.

<a id="canonical-706768b0da976f831a80a5451a19d9553eaf82fbd24faf8325e72b979e65a317"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / b561d1b595df / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-7bc14a16ed2bf7ef5eea6b5545821337ae33698dd9992bb7bb28e11da30b4ef9)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-4ff5ba9bf2195abe87c9b933a3eb9058f6a7cf776e97fb6b8583816e48918835)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-7bc14a16ed2bf7ef5eea6b5545821337ae33698dd9992bb7bb28e11da30b4ef9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-713fb95beef9d4b96f60baa75ca55b2594f00550c595a6606b9eba6ae61846f7"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 9e8d2cab5e20 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-1af24e56aba5dffefe3e4f04bee7ac07f503c5a13be322cf0bed770ebdc8dda0"></a>

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

<a id="canonical-6795709292e5d319bc97302c631a9596b6586cb5f6fbea99ae75ef7aeb2aa17b"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 9e8d2cab5e20 / 3

<a id="canonical-bac90e5a4720b95dade50c6c5e1f0fc2d87965c7427cda6052a2252e0b3f76c1"></a>

<a id="canonical-0d815b921a6f89fb4145353022937c7ab146158dda2f7183b41906a5984e6af6"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 9e8d2cab5e20 / 4

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

<a id="canonical-fee2f08c83b3ec087f2d41e173aeb1ed3114e91bc37916828add6ae3e9f864da"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 9e8d2cab5e20 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-4ff5ba9bf2195abe87c9b933a3eb9058f6a7cf776e97fb6b8583816e48918835"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a46cbee286954234cbbc64625af708b5274d936d12cd90ee4775f64403666fd"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ae081ce6399c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-252215748ddce9c5c3350b9b3995ced781f7d7fb63748e0439484e56dc22e2ce"></a>

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

<a id="canonical-912a2bceffce97dd8a624cfd400baeb268762026267588d2d45e8060e22149eb"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ae081ce6399c / 3

<a id="canonical-7d607bbc71a2be998664e9862e9af26cbc21690385958e77011d7f738e9cadf0"></a>

<a id="canonical-07617130575b52b1e78cb7cb9152fdecac1eec25b490bfbbbe40b5f057b323d7"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ae081ce6399c / 4

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

<a id="canonical-149ef22abcac673f0d26174cee0eb5b9ec0fcdfebae27f4dc92e050a600e6838"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ae081ce6399c / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-004.md#canonical-d55e613f2245cbb2bbe797605d6d98aa6bed655ac8e42154ab747793a4797efb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-484da7eed6ddfccfaa5143bd5304021079285094621408cbd1a5db4d53c879b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6649beee06187e3d12384ce69890c5470940abc8c34931b9504b0504aeb9e44b"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ea2c08764971 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-e589ed770f835442a1b679b8c0788ac289fd66257bb8548b615805e5894dcab0"></a>

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

<a id="canonical-40bee877070549548feba1c39dbf6e53f65a1755954446dfbdf6cce128c224f0"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ea2c08764971 / 3

<a id="canonical-ed51ed9d11f1630d928f68ef362433689b1b4cde9b917f7c2d4d120a848afcd5"></a>

<a id="canonical-a79a7e024844869681f0d9d90d63007a1a2dc5423713338241617b1047836193"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ea2c08764971 / 4

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

<a id="canonical-6185ab2d2d7e77842fed7ddeb308af4fba7694d674903eed17412f45ecdd3567"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ea2c08764971 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d6870e5c1478a282607e98cad79c62fae257da60a794d8d05196000905e1abd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4fbbeb884456ecbdf1e477b15a85d52d6e9a1b0de496c5936aed3eabb351b62"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 224b285caf5b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-8ef01c2f839c196287646dedc927fe5c70272a43aec63564b05fc75550d98f8f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-83f78b5e7fdbdd11b62f7789971a3918d7bf48d7b8ebd7459e90744c032be5ba"></a>

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

<a id="canonical-4b1a4406d294fd4867b0781ab536ceec2ee676442006722737e8f7571cbb5f64"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 224b285caf5b / 3

<a id="canonical-bf09418241d94f63d72cfd64c2ba89a0ecc41bccb43b3134c83fe6182e3d0720"></a>

<a id="canonical-cd8adbf78706d32e2f6954cfee8259ba0a57e6a8ec5db296dfabe9bf078dde61"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 224b285caf5b / 4

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

<a id="canonical-78b3148120cea918aca2e647723da924df6eaeb1654068fc31d62b37bfe896a7"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 224b285caf5b / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-004.md#canonical-dc0f4e75f2d30a063f2c5c746817181ef5020ebf44907c5daa5c15cc96833eea)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dddea3d5e8ada10dcdb25c2dae16779be48c19f6c782a0c0f1cf954bd966fc9"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 7351478dabe2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-837255f250b386ac3e589a3d26db5d7f8efc2e3e2c137f26b9be7ddfb512ef89"></a>

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

<a id="canonical-49e1e27295b5ddfd7091e0c9548291f40e06fef47d995ec6887fcd3e769e26d1"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 7351478dabe2 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-6f1358f1359110e89d1b8381d62f0fd34f45904c33c4f7730b8b7bb958b4bc1e): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-4a35a39d695646d80e76256b332e0bace37f6576627177b08ce6e7b8ce2e6c77): complete subsection reference.

<a id="canonical-c876bbb920ea8c752ac64c89d242d7c3c6e9de47871ec9b71614bded3b046b62"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 7351478dabe2 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-004.md#canonical-6f1358f1359110e89d1b8381d62f0fd34f45904c33c4f7730b8b7bb958b4bc1e)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-004.md#canonical-4a35a39d695646d80e76256b332e0bace37f6576627177b08ce6e7b8ce2e6c77)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6f1358f1359110e89d1b8381d62f0fd34f45904c33c4f7730b8b7bb958b4bc1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5f041e9acfcb0e63a8da9158a7621e1d9214f830d52da1f9dec1990dd4eaace"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / e32310d09a76 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-c3b0f31ab7c84978b68f9dfc79d806299ac2d3c1c13017d1feab2a4297359a57"></a>

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

<a id="canonical-9c624923f1b7139bbbc1154b43ba3bcd8a052230fa4e630fcc0d29a9cc5873aa"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / e32310d09a76 / 3

<a id="canonical-b60925c72b6aea331313c68814a9cdac6f23c4ab78d7155b85fe5b462572b0a8"></a>

<a id="canonical-84ee4104271a5974e519cee95446b80fda01fafd217e571cc07082b2286bf12b"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / e32310d09a76 / 4

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

<a id="canonical-2fa60d6e88f188bc2de9edda519e66f5bf29c9a8ed779d014d635f47631dd390"></a>

<a id="canonical-e040a0fcd296032cc8199f7fba71f3580c7a54849190720527d7da7e4f035344"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / e32310d09a76 / 5

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

<a id="canonical-a2e2b96fdd7477cf971591b4b1af8fcccedd771a28a66febeaa4b09fa070730c"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / e32310d09a76 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-4a35a39d695646d80e76256b332e0bace37f6576627177b08ce6e7b8ce2e6c77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c23577fe5ffa165c84039f7012fc4e3d0acc133a18deec54923d3636dd4be8b0"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 16dcb5af4b73 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.inside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-ac9e14e624a1abe552dc271041b30eb36e7b9c421ea7d6e9633239fb810b3e5d)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-260c8e6ec995af3dbba4a5553faab2d3bca2c461d17269e5f800c9c1b155febf)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-9e979ed089d2e549f39777a4d8b382c606a8a0c6bd47196513d344a2e1a094a0)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-8bec618e3c50be8b0d561a5729aae8047a6c1e5154954f3a876f373f6f80dccd"></a>

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

<a id="canonical-6814628bdddbe5cbadb792665b9bc44c69f2ee7fc0ef7d0095c0139df118cbe8"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 16dcb5af4b73 / 3

<a id="canonical-adc72f9e0c762dd8be968e46a03c29e0b2a4045c50c010e7e49d950211b61fd2"></a>

<a id="canonical-200346a5a73981c84e87c2eae8ee20136aa66ed18e03e12aba327a84c3de6aa3"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 16dcb5af4b73 / 4

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

<a id="canonical-64ef6e5b0b9816f0e2bdb1262a2c6b52b4097f78189ca8c23b62acbfe9b82a9d"></a>

<a id="canonical-5a7f61a3a3cf51ad9477d671dad3b49e809385af40cc30694184a8f82d31da23"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 16dcb5af4b73 / 5

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

<a id="canonical-80c4b709e916773fae2ea9a3d1c438d5d640a3ee619f12b0066b88b7a6d628ec"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 16dcb5af4b73 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-004.md#canonical-348deb69f9c7adf832104e0af6327e450bdfa675f3e2057a21f91b8bcf9e2c38)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-1157bf1fc3849f50346b8bbde307b90ef8e6269249bfcec73b51d239d3729608"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f227c66bf3d10d7cfe9e71b6df57fabced2320e8998c1cd5763611e1c9fb3d3"></a>

## ingress_egress_gw.no_dc_cluster_group — ingress_egress_gw.no_dc_cluster_group / 8921e48ebfb1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-ed56b4d74ab7f3e5466c562782c4de38bd6feef2a7b4e0e6e2647f658307a058"></a>

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

<a id="canonical-1931f6b337c981d52de543f05299f7de7d8b3b4db98edfec7416f65de69ec7f1"></a>

## Direct properties — ingress_egress_gw.no_dc_cluster_group / 8921e48ebfb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f2dded908a2420d509ab720590dd78b3435a37298a9af3fa8378ce2f1c7c652"></a>

## Next pages — ingress_egress_gw.no_dc_cluster_group / 8921e48ebfb1 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d3352522899102734197c804ef380720e3c775eccd31783b266584126d43b8f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5995904ae1e304d01370e9a7352f7ef3a62e6268d4fa650779cf2077941708a4"></a>

## ingress_egress_gw.no_forward_proxy — ingress_egress_gw.no_forward_proxy / 6eeeb7889824 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-da835fe567daa197aa5c80090199fe1a287064aa04596ed2606c9a29c952308f"></a>

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

<a id="canonical-8c4808490644b931da22ddf189fdc6e82339bcd92324af47db8563a7b1b08e9d"></a>

## Direct properties — ingress_egress_gw.no_forward_proxy / 6eeeb7889824 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4db352185ee2f04d6e05437b0cd4c2e2646bdc56f48fe4e1f7cb6125cadf415"></a>

## Next pages — ingress_egress_gw.no_forward_proxy / 6eeeb7889824 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b36ff120f5b0b07aec3fe6a9aefe8ff0510a9705dd34bcd90db8d81ff183d2f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d72e26156dd59870205632ef7cb5b3d1dcf0a61d25fb9a5dd345aa9cdaad8be3"></a>

## ingress_egress_gw.no_global_network — ingress_egress_gw.no_global_network / 0ecdc5d7d457 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_global_network

<a id="canonical-c74f09af08bcb52c22428aecf7860e14309eea223116f23f9173242271961cfa"></a>

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

<a id="canonical-58aaa46e8fd4c69938298e9ea5e53f8322a3aedd77330ff822f40c6c0dc34e7b"></a>

## Direct properties — ingress_egress_gw.no_global_network / 0ecdc5d7d457 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2be88fa1403b8c3c1b8bfcf363d6649975491f023620e4e237385b85aafc2ea4"></a>

## Next pages — ingress_egress_gw.no_global_network / 0ecdc5d7d457 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b423610de43db2fadb63eb73e0edbeb01d34934cd89dc2c45e11d90f3733768a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a927361ce387b941e0575aec0eb93f5d1171be1405481f35baaf6ebd601e9664"></a>

## ingress_egress_gw.no_inside_static_routes — ingress_egress_gw.no_inside_static_routes / cac099ea688e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-0298a549e2e4f84a683da24c1382fef9c266d9440ed49a77d9433b498e2256e4"></a>

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

<a id="canonical-68b3697f25d2ed91502ae362a5fa22768c38d8e903cd3302eba85b8f621ccb8c"></a>

## Direct properties — ingress_egress_gw.no_inside_static_routes / cac099ea688e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68434dcaeb6e9ebdf46f2703fd125bb7453fcb9b90ae21ba2729b51e458ff9f2"></a>

## Next pages — ingress_egress_gw.no_inside_static_routes / cac099ea688e / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-3a783c4069972120e21688650e441a04e1e7e7c1487b70747f4f0e96231b1c6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69483ab13c2d6980a107bbfc72043b814b9cd7a51c3fc008ff5eb8f58417aa78"></a>

## ingress_egress_gw.no_network_policy — ingress_egress_gw.no_network_policy / 024ec74d6dd4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_network_policy

<a id="canonical-0dde107a31a087bd48ab7b26fe92378ec4178545cefef0a19f70dfd20cb64967"></a>

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

<a id="canonical-f59d0faf053afe6cd1d60454a9d02ea5dd15d1ed055bcb1073bb4f2ef1935616"></a>

## Direct properties — ingress_egress_gw.no_network_policy / 024ec74d6dd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b86f7927f2047bb38f55e0c4a2157dd1055dfee9239392401da01c88b38c572b"></a>

## Next pages — ingress_egress_gw.no_network_policy / 024ec74d6dd4 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-002a76c05b3656d52e09091c927d25372511c2bdb8a6390cabfc6a075b09b992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eab1e5c929a57923db8517a28e1875acff1a897d5ad61c0dee866dc39def509f"></a>

## ingress_egress_gw.no_outside_static_routes — ingress_egress_gw.no_outside_static_routes / e9aa443ab872 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-9610426e168410f3a353e0746feab5b0ea7f1043ed66b70e8c42b4164a267cfc"></a>

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

<a id="canonical-61fc9d517208b38ad6929f5252e8d2b5f25c42d70f4e8a8f4b8eff0e449a4f0e"></a>

## Direct properties — ingress_egress_gw.no_outside_static_routes / e9aa443ab872 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d81a0941d590bc68c5645c62d0ffd81f8886fab50a2901410c8daf3ef8ccb06"></a>

## Next pages — ingress_egress_gw.no_outside_static_routes / e9aa443ab872 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-e3bc8be9facef978ebe6bc4f1306c0deef221171f6f593178c459f069831860e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0df0e284240080657122fb50d4313fdff5abfb5d04c6353092e169de9b27da16"></a>

## ingress_egress_gw.not_hub — ingress_egress_gw.not_hub / 508d19066556 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.not_hub

<a id="canonical-3b78350fa9545642a1ebd2a68a8af257fed6472056b2bde5c5afd989032bb4f9"></a>

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

<a id="canonical-fc71f575377af5c2388f3e5b0a373e74121ac454f17760db5a09742d9e1672c6"></a>

## Direct properties — ingress_egress_gw.not_hub / 508d19066556 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0cbaccc3ef9109c5bb62d3440d2015f66d8c7ca550355237c2b25343f566201f"></a>

## Next pages — ingress_egress_gw.not_hub / 508d19066556 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75b7c51aaa109f8f0123bc9de0799bd8620ffb05e0446f16d23bc5d42e852ed6"></a>

## ingress_egress_gw.outside_static_routes — ingress_egress_gw.outside_static_routes / 24fe76fe5274 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.outside_static_routes

<a id="canonical-9cf803d6aa07a05b620d7065a1820b17385d9d6ceade3f11e87525b5e80584f7"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

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

<a id="canonical-f6cc6d3975eb2a0dddce80ffadca84948baa5a3bf9a9582286bd1db5f32e4a10"></a>

## Direct properties — ingress_egress_gw.outside_static_routes / 24fe76fe5274 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5): complete subsection reference.

<a id="canonical-5ac45955d906e6561545542da9e6d01d9bfdaa10a3984bb9beabbff740582d5d"></a>

## Next pages — ingress_egress_gw.outside_static_routes / 24fe76fe5274 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b494f6d9ab6470be5a5b57eb7cf3a8577556a63dc1613f7faf06f52b50be901d"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — ingress_egress_gw.outside_static_routes.static_route_list / 2b8cd7c8fc15 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-9b58c03b3387b3021e8026014fbd7846abe90129d38251321c0ff50e6c6a31be"></a>

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

<a id="canonical-734b4538a747b36916f4469a2aa7de7583ac7b19cf8581831a12ac78793cebd7"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list / 2b8cd7c8fc15 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130): complete subsection reference.

<a id="canonical-3d42671a3fc2333adfe134dbd0aba069d96111fa69128173ff17611e0b434124"></a>

<a id="canonical-457c5a81ddbb555736b755033c9c50361f049528ea5d6a4cecb76471f1deb8e5"></a>

## simple_static_route property — ingress_egress_gw.outside_static_routes.static_route_list / 2b8cd7c8fc15 / 4

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

<a id="canonical-b6b4b3160d062b40315d5ba65ae4017a0d1deec291191def1731dd3b2f3fbc9c"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list / 2b8cd7c8fc15 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9358d6a21b0226e75370d94cea9c3efd0c297804746fd5e92c4b2cf22027ec71"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / e32a66a951e6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-a33e1e202d6d1d24754851f583bd0cd583ef1446d99a4f20dd9dc1fd6e068128"></a>

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

<a id="canonical-cfc682321b93b82b499d2b9f8c30d77651c20eda8acd740336c311a9866d4aca"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / e32a66a951e6 / 3

<a id="canonical-21fce68e0fbc5815c47b2e5ea40b2e79395df68f129e17aa74ac9c6792c19957"></a>

<a id="canonical-ede0d8129b7a97c5a9f5c8d90a685415db06e06dbce26a1459ca396bed788d1b"></a>

## attrs property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / e32a66a951e6 / 4

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

- [labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-48d69695eb2c482e7e74258ff61c646d36cad597a60c655e11952d2150de72f4): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00): complete subsection reference.

<a id="canonical-6207b4a0c7a265bf47ee780f6b839d730a700fda202e71bd41f250dfece20b61"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / e32a66a951e6 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-004.md#canonical-48d69695eb2c482e7e74258ff61c646d36cad597a60c655e11952d2150de72f4)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-48d69695eb2c482e7e74258ff61c646d36cad597a60c655e11952d2150de72f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5b5bd928cf1c322b304f3a0e2e001ee4f27d5e9d0a35f9f98b87b2464d9b50f"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 5c3ba01c5e5d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-61a55d14005f07077bfccbe3f9a1e4675a5a3c9dd8feb5802b53d833ce1d8cd5"></a>

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

<a id="canonical-08309fa849bc5b43ef91896cc10062f3ef4603254e066a54027f6f6afdeb6a90"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 5c3ba01c5e5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-914e3dcf68dfd1e774cec52e5b507db15ed16f83fbd8b2241e72aece041f09c6"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 5c3ba01c5e5d / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f17ea52114a153c771a2b2ae97206d0f13d3991673b806e327e9121d3bfa7bd"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / a13efb98a635 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-517bc5dd5263786ca6191337a35b820b1ade1cd2c1470a880571717fbd99c134"></a>

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

<a id="canonical-a5df569128c7e9041b2ae6be2302b2247bcb2a52b64632abe53eea94192697df"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / a13efb98a635 / 3

- [interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-54628cd6601d0b6af5e56b7a1803bd13601f89280ab8f5e04af12b5abaf9f8b2): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b): complete subsection reference.

<a id="canonical-c69134e7550043b6e6a6f3cedfccfed6495211d2b0749e81f27ff54d521de8e2"></a>

<a id="canonical-7ebeeeda3ceff62640b7038c25fe932caaa98b227375719986df5b53b02c5bfa"></a>

## type property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / a13efb98a635 / 4

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

<a id="canonical-1804e7a1ded509d5547cd342f9625ad58bca2c6d8b15fd44a04cd126212e8a21"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / a13efb98a635 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-004.md#canonical-54628cd6601d0b6af5e56b7a1803bd13601f89280ab8f5e04af12b5abaf9f8b2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-54628cd6601d0b6af5e56b7a1803bd13601f89280ab8f5e04af12b5abaf9f8b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
